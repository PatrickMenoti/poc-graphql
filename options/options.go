package options

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/machinebox/graphql"
)

type HTTPEvent struct {
	Status int `json:"status"`
	Count  int `json:"count"`
}

type Response struct {
	HTTPEvents []HTTPEvent `json:"httpEvents"`
}

type HTTPEvent2 struct {
	HTTPUserAgent string `json:"httpUserAgent"`
	Count         int    `json:"count"`
}

type APIResponse struct {
	HTTPEvents []HTTPEvent2 `json:"httpEvents"`
}

type HTTPMetric struct {
	Ts                 time.Time `json:"ts"`
	HTTPRequestsTotal  int       `json:"httpRequestsTotal"`
	HTTPSRequestsTotal int       `json:"httpsRequestsTotal"`
	EdgeRequestsTotal  int       `json:"edgeRequestsTotal"`
}

type MetricsResponse struct {
	HTTPMetrics []HTTPMetric `json:"httpMetrics"`
}

func Option1(tokenvalue string) {
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

func Option2(tokenvalue string) {
	graphqlClient := graphql.NewClient("https://api.azionapi.net/events/graphql")

	graphqlRequest := graphql.NewRequest(`
	query Top10UserAgent {
		httpEvents(
		  limit: 5
		  filter: {
			tsRange: {begin:"2024-01-20T10:10:10", end:"2024-01-26T10:10:10"}
		  }
		  aggregate: {count: rows} 
		  groupBy: [httpUserAgent]
		  orderBy: [count_DESC]
		  )
		{
		  httpUserAgent
		  count
		}
	  }
`)
	token := "Token " + tokenvalue

	graphqlRequest.Header.Set("Authorization", token)

	// var graphqlResponse interface{}
	var response APIResponse
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &response); err != nil {
		panic(err)
	}

	PrettyPrint(response)
	// spew.Dump(response)
	// fmt.Println(response)
}

func Option3(tokenvalue string) {
	graphqlClient := graphql.NewClient("https://api.azionapi.net/metrics/graphql")

	graphqlRequest := graphql.NewRequest(`
	query HttpCalculatedTotalRequests {
		httpMetrics(
		  limit: 1000
		  filter: {
			tsRange: {begin:"2024-01-20T10:10:10", end:"2024-01-26T10:10:10"}
		  }
		  groupBy: [ts]
		  orderBy: [ts_ASC]
		) 
		{        
		  ts
		  httpRequestsTotal
		  httpsRequestsTotal
		  edgeRequestsTotal
		}
	  }
`)
	token := "Token " + tokenvalue

	graphqlRequest.Header.Set("Authorization", token)

	// var graphqlResponse interface{}
	var response MetricsResponse
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &response); err != nil {
		panic(err)
	}

	PrettyPrint(response)
	// spew.Dump(response)
	// fmt.Println(response)
}

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
