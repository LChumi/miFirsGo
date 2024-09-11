package main

// go get github.com/aws/aws-lambda-go/lambda 
// goos=linux go build -o main main.go para compilaciones 
// en caso de funvionar en cmd 
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Link string `json:"link"`
	Key string `json:"key"`
}

type Response struct {
	Link string `json:"link"`
	Key string `json:"key"`
}

func main(){
	lambda.Start(Handler)
}

func Handler (event InputEvent)(Response,error){
	fmt.Print(event) //para efectos de registro mando a mostrar cada vez que entren se imprime 

	var slink string 
	if event.Link == "" {
		slink="vacio"
	} else{
		slink = event.Link
	}

	return Response{
		Link: slink,
		Key: event.Key,
	}, nil
}