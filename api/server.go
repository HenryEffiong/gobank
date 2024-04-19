package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/henryeffiong/gobank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	// config     util.Config
	store *db.Store
	// tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
