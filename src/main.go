package main

import (
	 "log"
// 	 "fmt"
)
// structure:
// main
// -- store instance
// -- newApiServer instance
// --== apiserver struct instance
//
// -- newApiServer.RUN
// --== router(mux package)
// --== custom makeHTTPHandleFunc
// --==-- handleRequest (GET discrimination)
// --==--== handleGetGap - returns table from sql


func main(){
	store, err := newPGStore()
	if err !=nil{
		log.Fatal(err)
	}
	// troubleshooting store connection
	// fmt.Printf("%+v/n", store)
	server := NewAPIServer(":3000", store)
	server.Run()

}