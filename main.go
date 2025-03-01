package main
import (
	"TM/router"
)
func main(){
	router := router.InitializeRouter()
	router.Run("localhost:8081")
}