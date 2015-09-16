package db

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cloudfoundry/gunk/workpool"
	"github.com/cloudfoundry/storeadapter"
	"github.com/cloudfoundry/storeadapter/etcdstoreadapter"
)

//go:generate counterfeiter -o fakes/fake_db.go . DB
type DB interface {
	ReadRoutes() ([]Route, error)
	SaveRoute(route Route) error
	DeleteRoute(route Route) error

	ReadTcpMappings() ([]TcpMapping, error)
	SaveTcpMapping(tcpMapping TcpMapping) error

	Connect() error
	Disconnect() error
	WatchRouteChanges() (<-chan storeadapter.WatchEvent, chan<- bool, <-chan error)
}

type Feature string

type RouterGroup struct {
	Guid     string    `json:"guid"`
	Name     string    `json:"name"`
	Features []Feature `json:"features"`
}

type Route struct {
	Route           string `json:"route"`
	Port            uint16 `json:"port"`
	IP              string `json:"ip"`
	TTL             int    `json:"ttl"`
	LogGuid         string `json:"log_guid"`
	RouteServiceUrl string `json:"route_service_url,omitempty"`
}

type TcpMapping struct {
	TcpRoute tcpRoute `json:"route"`
	HostPort uint16   `json:"host_port"`
	HostIP   string   `json:"host_ip"`
}

type tcpRoute struct {
	RouterGroupGuid string `json:"router_group_guid"`
	ExternalPort    uint16 `json:"external_port"`
}

const (
	TCP_MAPPING_BASE_KEY string = "/v1/tcp_mapppings/router_groups"
	HTTP_ROUTE_BASE_KEY  string = "/routes"
)

type etcd struct {
	storeAdapter *etcdstoreadapter.ETCDStoreAdapter
}

func NewETCD(nodeURLs []string, maxWorkers uint) (*etcd, error) {
	workpool, err := workpool.NewWorkPool(int(maxWorkers))
	if err != nil {
		return nil, err
	}

	storeAdapter := etcdstoreadapter.NewETCDStoreAdapter(nodeURLs, workpool)
	return &etcd{
		storeAdapter: storeAdapter,
	}, nil
}

func (e *etcd) Connect() error {
	return e.storeAdapter.Connect()
}

func (e *etcd) Disconnect() error {
	return e.storeAdapter.Disconnect()
}

func (e *etcd) ReadRoutes() ([]Route, error) {
	routes, err := e.storeAdapter.ListRecursively(HTTP_ROUTE_BASE_KEY)
	if err != nil {
		return []Route{}, nil
	}

	listRoutes := []Route{}
	for _, node := range routes.ChildNodes {
		route := Route{}
		json.Unmarshal([]byte(node.Value), &route)
		listRoutes = append(listRoutes, route)
	}
	return listRoutes, nil
}

func (e *etcd) SaveRoute(route Route) error {
	key := generateKey(route)
	routeJSON, _ := json.Marshal(route)
	node := storeadapter.StoreNode{
		Key:   key,
		Value: routeJSON,
		TTL:   uint64(route.TTL),
	}

	return e.storeAdapter.SetMulti([]storeadapter.StoreNode{node})
}

func (e *etcd) DeleteRoute(route Route) error {
	key := generateKey(route)
	err := e.storeAdapter.Delete(key)
	if err != nil && err.Error() == "the requested key could not be found" {
		err = DBError{Type: KeyNotFound, Message: "The specified route could not be found."}
	}
	return err
}

func (e *etcd) WatchRouteChanges() (<-chan storeadapter.WatchEvent, chan<- bool, <-chan error) {
	return e.storeAdapter.Watch("/routes")
}

func generateKey(route Route) string {
	return fmt.Sprintf("%s/%s,%s:%d", HTTP_ROUTE_BASE_KEY, url.QueryEscape(route.Route), route.IP, route.Port)
}

func (e *etcd) ReadTcpMappings() ([]TcpMapping, error) {
	tcpMappings, err := e.storeAdapter.ListRecursively(TCP_MAPPING_BASE_KEY)
	if err != nil {
		return []TcpMapping{}, nil
	}

	listMappings := []TcpMapping{}
	for _, routerGroupNode := range tcpMappings.ChildNodes {
		for _, externalPortNode := range routerGroupNode.ChildNodes {
			for _, mappingNode := range externalPortNode.ChildNodes {
				tcpMapping := TcpMapping{}
				json.Unmarshal([]byte(mappingNode.Value), &tcpMapping)
				listMappings = append(listMappings, tcpMapping)
			}
		}
	}
	return listMappings, nil
}

func (e *etcd) SaveTcpMapping(tcpMapping TcpMapping) error {
	key := generateTcpMappingKey(tcpMapping)
	tcpMappingJson, _ := json.Marshal(tcpMapping)
	node := storeadapter.StoreNode{
		Key:   key,
		Value: tcpMappingJson,
	}
	return e.storeAdapter.SetMulti([]storeadapter.StoreNode{node})
}

func generateTcpMappingKey(tcpMapping TcpMapping) string {
	// Generating keys following this pattern
	// /v1/tcp_routes/router_groups/{router_guid}/{external_port}/{host-ip}:{host-port}
	return fmt.Sprintf("%s/%s/%d/%s:%d", TCP_MAPPING_BASE_KEY,
		tcpMapping.TcpRoute.RouterGroupGuid, tcpMapping.TcpRoute.ExternalPort, tcpMapping.HostIP, tcpMapping.HostPort)
}

func NewTcpMapping(routerGroupGuid string, externalPort uint16, hostIP string, hostPort uint16) TcpMapping {
	return TcpMapping{
		TcpRoute: tcpRoute{RouterGroupGuid: routerGroupGuid, ExternalPort: externalPort},
		HostPort: hostPort,
		HostIP:   hostIP,
	}
}
