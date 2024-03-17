package main

import (
	"github.com/sirroderickyammins/avweather/api"
)

func main() {
	request := api.Request{
		IDs:    []string{"KBWI", "KMCI", "KMTN", "KIAD"},
		Format: "json",
		Metar:  "true",
		Date:   "20240310_234422Z",
	}

	api.GetTAF(request)
}
