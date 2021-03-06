package main

import (
  "os"
  service "github.com/mmsikora/gogo-service/service"
)

func main() {
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }

  server := service.NewServer()
  server.Run(":" + port)
}
