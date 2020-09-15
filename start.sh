#!/bin/bash

cd transport/graphql/schema/
go generate
cd ../../../
# docker-compose up --build

go run cmd/todo/main.go