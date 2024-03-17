package api

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Request struct {
	IDs    []string
	Format string
	Metar  string
	Date   string
}

func GetTAF(tafRequest Request) (quest Request, err error) {
	// Validate the struct with reflection
	inputType := reflect.TypeOf(tafRequest)

	if inputType.Kind() != reflect.Struct {
		return tafRequest, errors.New("GetTAF func needs struct as input")
	}

	structVal := reflect.ValueOf(tafRequest)
	fieldNum := structVal.NumField()

	// Check if fields are set
	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := inputType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			err = fmt.Errorf(fmt.Sprintf("%v%s is not set", err, fieldName))
		}

	}

	baseURL := "https://aviationweather.gov/data/api/taf"

	params := url.Values{}

	icaoString := strings.Join(tafRequest.IDs, ",")

	params.Set("ids", icaoString)
	params.Set("format", tafRequest.Format)
	params.Set("metar", tafRequest.Metar)
	params.Set("date", tafRequest.Date)

	finalURL := baseURL + "?" + params.Encode()

	// response, err := http.Get(finalURL)

	fmt.Println(finalURL)

	return tafRequest, err
}
