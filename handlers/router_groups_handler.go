package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cloudfoundry-incubator/routing-api/authentication"
	"github.com/cloudfoundry-incubator/routing-api/db"
	"github.com/pivotal-golang/lager"
)

const (
	RouterGroupsReadScope     = "router_groups.read"	
)

type RouterGroupsHandler struct {
	token     authentication.Token
	logger    lager.Logger
}

func NewRouteGroupsHandler(token authentication.Token, logger lager.Logger) *RouterGroupsHandler {
	return &RouterGroupsHandler{
		token:     token,		
		logger:    logger,
	}
}

func (h *RouterGroupsHandler) ListRouterGroups(w http.ResponseWriter, req *http.Request) {
	log := h.logger.Session("list-router-groups")

	err := h.token.DecodeToken(req.Header.Get("Authorization"), AdminRouteScope)
	if err != nil {
		handleUnauthorizedError(w, err, log)
		return
	}
	
	
}
