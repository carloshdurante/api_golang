package server

import (
	"log"

	"github.com/carloshdurante/api_golang/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8090",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server running on port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
