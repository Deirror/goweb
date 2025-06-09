# gRPC-kitchen

![images](https://github.com/user-attachments/assets/b9b711f6-9b1b-416d-be3c-7ae513fbcd1a)

> This project was based in [
Complete Golang and gRPC Microservices (Project Course)](https://www.youtube.com/watch?v=ea_4Ug5WWYE)

Description
-

A project demonstrating how microservices communicate using gRPC, instead of HTTP(JSON (un)marshaling is a heavy task).

Install
-

-  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
-  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
-  `make` should be installed to run servers

Steps to run
-

-  `git clone` on the repo
-  run `make gen` to generate gRPC protobuf Go code
-  run `make run-orders` and `make run-kitchen` on other terminal
-  open your browser with the following URL -> `http://localhost:1000`
