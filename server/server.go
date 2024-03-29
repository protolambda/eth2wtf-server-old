package server

import (
	"eth2wtf-server/clh"
	"eth2wtf-server/client"
	. "eth2wtf-server/common"
	"eth2wtf-server/hub"
	"eth2wtf-server/world"
	"log"
	"net/http"
	"os"
)

type Server struct {
	clientHub *hub.Hub
	world     *world.World
}

func NewServer() *Server {
	return &Server{
		clientHub: hub.NewHub(),
		world:     world.NewWorld(log.New(os.Stdout, "world ", log.LstdFlags)),
	}
}

func (s *Server) Run() {
	s.clientHub.Run()
}

func (s *Server) ServeWs(w http.ResponseWriter, r *http.Request) {
	s.clientHub.ServeWs(w, r, s.NewClientHandler)
}

func (s *Server) NewClientHandler(send chan<- []byte) client.ClientHandler {
	return clh.NewClientHandler(s.world, send)
}

func (s *Server) World() *world.World {
	return s.world
}

func (s *Server) GetViewers() []Viewer {
	return s.clientHub.GetViewers()
}
