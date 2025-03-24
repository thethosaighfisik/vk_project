package db

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	_ "github.com/tarantool/go-tarantool/datetime"
	_ "github.com/tarantool/go-tarantool/decimal"
	_ "github.com/tarantool/go-tarantool/uuid"
	//"time"
	"log"
)

func ConnectDB() (*tarantool.Connection){

	opts := tarantool.Opts{
		User: "sampleuser",
		Pass: "123456",
	}

	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
		return nil
	}
	resp, err := conn.Ping()
	if err != nil {
    		log.Fatalf("Ошибка подключения: %s", err)
	}
	log.Println("Tarantool подключен:", resp.Code)
	return conn
}
