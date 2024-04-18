package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

type Cloud struct {
	Cover string `json:"cover"`
	Base  int64  `json:"base"`
	Type  int64  `json:"type"`
}

// Forecast represents the fcsts object in the JSON.
type Forecast struct {
	TimeGroup   int      `json:"timeGroup"`
	TimeFrom    int64    `json:"timeFrom"`
	TimeTo      int64    `json:"timeTo"`
	TimeBec     *string  `json:"timeBec"`
	FcstChange  *string  `json:"fcstChange"`
	Probability *int     `json:"probability"`
	Wdir        int      `json:"wdir"`
	Wspd        int      `json:"wspd"`
	Wgst        *int     `json:"wgst"`
	Visib       string   `json:"visib"`
	Altim       *string  `json:"altim"`
	Clouds      []Cloud  `json:"clouds"`
	IcgTurb     []string `json:"icgTurb"`
	Temp        []string `json:"temp"`
}

// Response represents the entire JSON structure.
type Response struct {
	TafID         int        `json:"tafId"`
	IcaoID        string     `json:"icaoId"`
	DbPopTime     string     `json:"dbPopTime"`
	BulletinTime  string     `json:"bulletinTime"`
	IssueTime     string     `json:"issueTime"`
	ValidTimeFrom int64      `json:"validTimeFrom"`
	ValidTimeTo   int64      `json:"validTimeTo"`
	RawTAF        string     `json:"rawTAF"`
	MostRecent    int        `json:"mostRecent"`
	Remarks       string     `json:"remarks"`
	Lat           float64    `json:"lat"`
	Lon           float64    `json:"lon"`
	Elev          int        `json:"elev"`
	Prior         int        `json:"prior"`
	Name          string     `json:"name"`
	Fcsts         []Forecast `json:"fcsts"`
	RawOb         string     `json:"rawOb"`
}

func GetTAF(tafRequest Request) (jsonData []Response, err error) {
	// Validate the struct with reflection
	inputType := reflect.TypeOf(tafRequest)

	if inputType.Kind() != reflect.Struct {
		return jsonData, errors.New("GetTAF func needs struct as input")
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

	// TODO: Could probably abstract the URL generation into a separate function.
	// Might need several different structs for each different API request
	baseURL := "https://aviationweather.gov/api/data/taf"

	params := url.Values{}

	icaoString := strings.Join(tafRequest.IDs, ",")

	params.Set("ids", icaoString)
	params.Set("format", tafRequest.Format)
	params.Set("metar", tafRequest.Metar)
	params.Set("date", tafRequest.Date)

	finalURL := baseURL + "?" + params.Encode()

	httpRes, err := http.Get(finalURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(httpRes.Body)
	httpRes.Body.Close()

	if httpRes.StatusCode > 209 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", httpRes.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", jsonData[0].IcaoID)
	return jsonData, err
}
