package routing_api

import "github.com/tedsuo/rata"

const (
	UpsertRoute      = "Upsert"
	DeleteRoute      = "Delete"
	ListRoute        = "List"
	EventStreamRoute = "EventStream"
	ListRouterGroups = "ListRouterGroups"
)

var Routes = rata.Routes{
	{Path: "/v1/routes", Method: "POST", Name: UpsertRoute},
	{Path: "/v1/routes", Method: "DELETE", Name: DeleteRoute},
	{Path: "/v1/routes", Method: "GET", Name: ListRoute},
	{Path: "/v1/events", Method: "GET", Name: EventStreamRoute},
	{Path: "/v1/router_groups", Method: "GET", Name: ListRouterGroups},
}
