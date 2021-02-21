package main

import (
	"log"
	"github.com/romaxa83/go-http-rest-api/internal/app/apiserver"
)

func main()  {
	// создаемнаш апи сервер
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
