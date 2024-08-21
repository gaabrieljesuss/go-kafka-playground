package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	_ "go_kafka_playground/api-gateway/src/apps/api/docs"
	"go_kafka_playground/api-gateway/src/apps/api/routes"
	"go_kafka_playground/api-gateway/src/utils"
	"log"
	"strconv"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	godotenv.Load(".env")
	api := NewAPI(getServerHostAndPort())
	api.Serve()
}

func getServerHostAndPort() (string, int) {
	host := utils.GetenvWithDefault("SERVER_HOST", "0.0.0.0")
	portStr := utils.GetenvWithDefault("SERVER_PORT", "8000")
	var port int
	if v, err := strconv.Atoi(portStr); err != nil {
		log.Fatal("The server port env variable must be a number (e.g 8000)")
	} else {
		port = v
	}
	return host, port
}

type API interface {
	Serve()
}

type api struct {
	host   string
	port   int
	server *echo.Echo
}

// @title Go Kafka Playground API
// @version 1.0
// @description Go Kafka Playground API
// @contact.name Gabriel de Jesus
// @contact.email gjs8@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func NewAPI(host string, port int) API {
	server := echo.New()
	return &api{host, port, server}
}

func (a *api) Serve() {
	a.loadRoutes()
	a.start()
}

func (a *api) rootEndpoint() *echo.Group {
	return a.server.Group("/api")
}

func (a *api) loadRoutes() {
	manager := routes.New()
	manager.Load(a.rootEndpoint())
}

func (a *api) start() {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	if err := a.server.Start(address); err != nil {
		a.server.Logger.Fatal(err)
	}
}
