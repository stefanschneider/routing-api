package handlers

import (
	"encoding/json"
	routing_api "github.com/cloudfoundry-incubator/routing-api"
	"github.com/cloudfoundry-incubator/routing-api/authentication"
	"github.com/cloudfoundry-incubator/routing-api/db"
	"github.com/pivotal-golang/lager"
	"net/http"
)

type TcpMappingsHandler struct {
	token  authentication.Token
	db     db.DB
	logger lager.Logger
}

func NewTcpMappingsHandler(token authentication.Token, database db.DB, logger lager.Logger) *TcpMappingsHandler {
	return &TcpMappingsHandler{
		token:  token,
		db:     database,
		logger: logger,
	}
}

func (h *TcpMappingsHandler) Upsert(w http.ResponseWriter, req *http.Request) {
	log := h.logger.Session("create-tcp-mapping")
	decoder := json.NewDecoder(req.Body)

	var tcpMappings []db.TcpMapping
	err := decoder.Decode(&tcpMappings)
	if err != nil {
		handleProcessRequestError(w, err, log)
		return
	}

	log.Info("request", lager.Data{"tcp_mapping_creation": tcpMappings})

	// err = h.token.DecodeToken(req.Header.Get("Authorization"), AdvertiseRouteScope, AdminRouteScope)
	// if err != nil {
	// 	handleUnauthorizedError(w, err, log)
	// 	return
	// }

	err = h.validateTcpMappings(tcpMappings)
	if err != nil {
		handleProcessRequestError(w, err, log)
		return
	}

	for _, tcpMapping := range tcpMappings {
		err = h.db.SaveTcpMapping(tcpMapping)
		if err != nil {
			handleDBCommunicationError(w, err, log)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TcpMappingsHandler) validateTcpMappings(tcpMappings []db.TcpMapping) error {
	for _, tcpMapping := range tcpMappings {
		err := h.validateTcpMapping(tcpMapping)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *TcpMappingsHandler) validateTcpMapping(tcpMappings db.TcpMapping) error {

	if tcpMappings.TcpRoute.RouterGroupGuid == "" {
		err := routing_api.NewError(routing_api.TcpMappingInvalidError, "Each tcp mapping requires a valid router group guid")
		return &err
	}

	if tcpMappings.TcpRoute.ExternalPort <= 0 {
		err := routing_api.NewError(routing_api.TcpMappingInvalidError, "Each tcp mapping requires a positive external port")
		return &err
	}

	if tcpMappings.HostIP == "" {
		err := routing_api.NewError(routing_api.TcpMappingInvalidError, "Each tcp mapping requires a non empty host ip")
		return &err
	}

	if tcpMappings.HostPort <= 0 {
		err := routing_api.NewError(routing_api.TcpMappingInvalidError, "Each tcp mapping requires a positive host port")
		return &err
	}

	return nil
}
