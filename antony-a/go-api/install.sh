#!/bin/sh
go install github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

