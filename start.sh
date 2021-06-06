#!/bin/bash

cd app && go env -w GO111MODULE=off && go get github.com/satori/go.uuid && go get github.com/go-sql-driver/mysql && go run cmd/*.go