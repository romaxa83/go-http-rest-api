package apiserver

//
type APIServer struct {

}

// инициализируем apiserver
func New() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error  {
	return nil
}