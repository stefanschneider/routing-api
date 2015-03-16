package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudfoundry-incubator/routing-api/authentication"
	"github.com/cloudfoundry-incubator/routing-api/config"
	"github.com/cloudfoundry-incubator/routing-api/db"
	"github.com/cloudfoundry-incubator/routing-api/handlers"
	"github.com/cloudfoundry/dropsonde"
	"github.com/pivotal-golang/lager"

	"github.com/tedsuo/rata"
)

var Routes = rata.Routes{
	{Path: "/v1/routes", Method: "POST", Name: "Upsert"},
	{Path: "/v1/routes", Method: "DELETE", Name: "Delete"},
	{Path: "/v1/routes", Method: "GET", Name: "List"},
}

var maxTTL = flag.Int("maxTTL", 120, "Maximum TTL on the route")
var port = flag.Int("port", 8080, "Port to run rounting-api server on")
var cfg_flag = flag.String("config", "", "Configuration for routing-api")

func route(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(f)
}

func main() {

	flag.Parse()
	if *cfg_flag == "" {
		fmt.Fprintf(os.Stderr, "starting: No configuration file provided")
		os.Exit(1)
	}

	cfg, err := config.NewConfigFromFile(*cfg_flag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "starting: %v", err)
		os.Exit(1)
	}

	logger, err := getLogger(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Log file creation error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("omg wtf %v", logger)
	logger.Info("ballllaaaaaaa!")

	err = dropsonde.Initialize(cfg.MetronConfig.Address+":"+cfg.MetronConfig.Port, cfg.LogGuid)
	if err != nil {
		logger.Error("Dropsonde failed to initialize:", err)
		os.Exit(1)
	}

	logger.Info("database", lager.Data{"etcd-addresses": flag.Args()})
	database := db.NewETCD(flag.Args())

	token := authentication.NewAccessToken(cfg.UAAPublicKey)
	err = token.CheckPublicToken()
	if err != nil {
		logger.Error("starting", err)
		os.Exit(1)
	}

	validator := handlers.NewValidator()

	routesHandler := handlers.NewRoutesHandler(token, *maxTTL, validator, database, logger)

	actions := rata.Handlers{
		"Upsert": route(routesHandler.Upsert),
		"Delete": route(routesHandler.Delete),
		"List":   route(routesHandler.List),
	}

	handler, err := rata.NewRouter(Routes, actions)
	if err != nil {
		panic("unable to create router: " + err.Error())
	}

	handler = handlers.LogWrap(handler, logger)

	logger.Info("starting", lager.Data{"port": *port})
	err = http.ListenAndServe(":"+strconv.Itoa(*port), handler)
	if err != nil {
		panic(err)
	}
}

func getLogger(config config.Config) (lager.Logger, error) {
	logger := lager.NewLogger("routing-api")

	if config.LogFile == "" {
		sink := lager.NewReconfigurableSink(lager.NewWriterSink(os.Stdout, lager.DEBUG), lager.DEBUG)
		logger.RegisterSink(sink)
	} else {
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}

		sink := lager.NewReconfigurableSink(lager.NewWriterSink(file, lager.DEBUG), lager.DEBUG)
		logger.RegisterSink(sink)
	}

	return logger, nil
}
