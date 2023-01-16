package main

func main() {
	server := NewServer(":3000")
	server.Route("/a", server.AddMiddleware(HandleRoot, LoggerMiddleware()), server.routers.home_router)
	server.Route("/b", server.AddMiddleware(HandleHome, LoggerMiddleware()), server.routers.api_router)
	server.Listen()
}
