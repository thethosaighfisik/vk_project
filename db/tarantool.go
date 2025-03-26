package db

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	_ "github.com/tarantool/go-tarantool/datetime"
	_ "github.com/tarantool/go-tarantool/decimal"
	_ "github.com/tarantool/go-tarantool/uuid"
	//"time"
	"log"
<<<<<<< HEAD
	"os"
=======
>>>>>>> 886bde79 (First commit)
)

func ConnectDB() (*tarantool.Connection){

	opts := tarantool.Opts{
		User: "sampleuser",
		Pass: "123456",
	}
<<<<<<< HEAD
	host := os.Getenv("TARANTOOL_HOST")  // Будет "tarantool" из переменной окружения
	port := os.Getenv("TARANTOOL_PORT")  // Будет "3301"

// Используйте эти переменные для подключения

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := tarantool.Connect(addr, opts)
=======

	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
>>>>>>> 886bde79 (First commit)
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
