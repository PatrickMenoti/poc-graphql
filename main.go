package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/machinebox/graphql"
)

type HTTPEvent struct {
	Status int `json:"status"`
	Count  int `json:"count"`
}

type Response struct {
	HTTPEvents []HTTPEvent `json:"httpEvents"`
}

func main() {
	var tokenvalue string
	flag.StringVar(&tokenvalue, "token", "", "API token for authorization")
	// Parse the command-line flags
	flag.Parse()

	// Check if the token is provided
	if tokenvalue == "" {
		fmt.Println("Please provide a valid API token using the -token flag.")
		return
	}

	graphqlClient := graphql.NewClient("https://api.azionapi.net/events/graphql")

	graphqlRequest := graphql.NewRequest(`
	query Top10StatusCodes {
		httpEvents(
		  limit: 5
		  filter: {
			tsRange: { begin:"2024-01-20T10:10:10", end:"2024-01-26T10:10:10" }
		  }
		  aggregate: {count: status}
		  groupBy: [status]
		  orderBy: [count_DESC]
		  )
		{
		  status
		  count
		}
	  }
`)
	token := "Token " + tokenvalue

	graphqlRequest.Header.Set("Authorization", token)

	// var graphqlResponse interface{}
	var response Response
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &response); err != nil {
		panic(err)
	}

	PrettyPrint(response)
	// spew.Dump(response)
	// fmt.Println(response)

}

// func init() {
// 	flag.StringVar(&flagvar, "token", "", "personaltoken")
// }

// print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
