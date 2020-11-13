package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"

)

 func main(){
	 fmt.Println("Servidor gRPC")
	 lis, err := net.Listen("tcp",":9000")
	 if err != nil {
		 log.Fatalf("Failed to listen 9000 %v", err)
	 }
	 grpcServer := grpc.NewServer()
	 if err := grpcServer.Serve(lis); err != nil{
		 log.Fatalf("fallo el servidor grcp en el puerto 9000: %v",err)
	 }
 }