package handlers

import (
	//"log"
	"fmt"
	"vk_project/models"
	"vk_project/db"
	"encoding/json"
	"net/http"
	"github.com/tarantool/go-tarantool"
	_ "github.com/tarantool/go-tarantool/datetime"
	_ "github.com/tarantool/go-tarantool/decimal"
	_ "github.com/tarantool/go-tarantool/uuid"
	                "github.com/go-chi/chi/v5"
)


func HandlerPost(w http.ResponseWriter, r *http.Request){
	var data models.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Invalid JSON:", http.StatusBadRequest)
		return
	}
	conn := db.ConnectDB()

	defer conn.Close()

	Data, err := conn.Do(tarantool.NewSelectRequest("bands").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{data.Key}),).Get()
	if len(Data.Data) != 0 {
		fmt.Println("The key exists:", http.StatusConflict)
		return
	}

	
	Data, err = conn.Do(tarantool.NewInsertRequest("bands").Tuple([]interface{}{data.Key, data.Value}),).Get()
	if err != nil {
		fmt.Println("Got an error:", err)
		return
	} else {
		fmt.Println("Inserted tuple:", Data)
	}

}
func HandlerPut(w http.ResponseWriter, r *http.Request){
	var Value models.PutStruct
	err := json.NewDecoder(r.Body).Decode(&Value)
	if err != nil {
		fmt.Println("Invalid JSON:", http.StatusBadRequest)
		return 
	}
	conn := db.ConnectDB()

        defer conn.Close()

        var key string
        key = chi.URLParam(r, "id")
        fmt.Println("Received key:", key)
        data, err := conn.Do(tarantool.NewSelectRequest("bands").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{key}),).Get()
        if len(data.Data) == 0 {
                fmt.Println("The key doesn't exist:", http.StatusNotFound)
                return
        }

	data, err = conn.Do(
    		tarantool.NewReplaceRequest("bands").
        	Tuple([]interface{}{key, Value.Value}),
		).Get()

	if err != nil {
		fmt.Println("Got an error:", err)
		return
	}
	fmt.Println("Updated tuple:", data)
}


func HandlerGet(w http.ResponseWriter, r *http.Request){
	var key string
	key = chi.URLParam(r, "id")
	fmt.Println("Received key:", key)

	conn := db.ConnectDB()

        defer conn.Close()

	
	data, err := conn.Do(tarantool.NewSelectRequest("bands").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{key}),).Get()

	if err != nil || len(data.Data) == 0{
		fmt.Println("The key doesn't exist:", http.StatusNotFound)
		return
	}
	fmt.Println("Tuple:", data)

}

func HandlerDelete(w http.ResponseWriter, r *http.Request){
	var key string
	key = chi.URLParam(r, "id")
	fmt.Println("Received key:", key)

	conn := db.ConnectDB()

	defer conn.Close()
	
	data, err := conn.Do(tarantool.NewSelectRequest("bands").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{key}),).Get()
        if len(data.Data) == 0 {
                fmt.Println("The key doesn't exist:", http.StatusNotFound)
                return
        }

	data, err = conn.Do(tarantool.NewDeleteRequest("bands").Key([]interface{}{key}),).Get()
	if err != nil {
		fmt.Println("Got an error:", err)
		return
	}
	fmt.Println("Deleted tuple:", data)

}
