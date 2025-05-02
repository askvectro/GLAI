module gla-core

go 1.17

require (
    github.com/gorilla/websocket v1.5.0
    github.com/labstack/echo/v4 v4.11.1
    github.com/sirupsen/logrus v1.9.3
    github.com/joho/godotenv v1.5.1
    firebase.google.com/go/v4 v4.15.2
)

replace (
    google.golang.org/grpc => github.com/grpc/grpc-go v1.42.0
    google.golang.org/protobuf => github.com/golang/protobuf v1.5.2
)
