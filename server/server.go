package server

import (
	"log"

	"github.com/flavio-foa-dev/adopet/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigureRoutes(s.server)
	log.Print("Starting server at port:", s.port)
	log.Fatal(router.Run(":" + s.port))

}
