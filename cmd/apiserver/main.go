package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/romaxa83/go-http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

// парсим (при запуске сервер), файл крнфига,
//который можно задавать в командной строке при запуске
func init()  {
	/*
	параметры -
		1 - в какую переменую парсим
		2 - какой флаг при запуске команды
		3 - значение по дефолту
		4 - описание
	*/
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()  {
	// парсим переданные флаги при запуски сервера
	flag.Parse()

	// инициализируем конфиг
	config := apiserver.NewConfig()

	// парсим файл с конфигом (1 пар. - путь к файлу, 2 - куда парсим)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// создаемнаш апи сервер
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
