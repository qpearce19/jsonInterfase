package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Error bool	`json:"error"`
	Message string `json:"message"`

	Data json.RawMessage	`json:"data"`
}

type deviceList struct {
	DeviceId int `json:"deviceId"`
	State bool	`json:"state"`
	LastRssi int	`json:"lastRssi"`
}

type gateList struct {
	GatewayId int	`json:"gatewayId"`
	State bool	`json:"state"`
	LastOnline int	`json:"lastOnline"`
}

func main() {
	var response response
	var topic string = "device/list"
	var jsonBlob = []byte(`{
  "error": false,                   
  "message": "abs",
  
  "data":
  [
    {
      "deviceId": 15,
      "state": true,
	  "lastRssi": 55
    }
  ]
}`)

	err := json.Unmarshal(jsonBlob, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("fine")
	if response.Error == true {
		log.Fatal(response.Error)
	}

	switch topic {
	case "device/list":
		var dl []deviceList
		if err := json.Unmarshal([]byte(response.Data), &dl); err != nil {
			log.Fatal(err)
		}
		fmt.Println(dl[0].DeviceId)
	case "gateway/list":
		var gl []gateList
		if err := json.Unmarshal([]byte(response.Data), &gl); err != nil {
			log.Fatal(err)
		}
	}

}
