package server

type WebServer struct {
	Port    int
	Version int
}

func NewWebServer() *WebServer {
	tmp := &WebServer{
		Port:    8080,
		Version: 1,
	}
	return tmp
}
