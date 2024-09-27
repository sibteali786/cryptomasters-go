package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/datatypes"
)

const apiURL = "http://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	// 3 because B is 1 Char
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters minimum: %d received", len(currency))
	}
	// converting currency into uppercase
	upCase := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upCase))
	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		// res.Body is a reader (tcp connection return data in chunks )
		bodyBytes, err := io.ReadAll(res.Body) // kind of like await in js ( synchronous ). Will wait here until all data comes in
		if err != nil {
			return nil, err
		}

		// json := string(bodyBytes) // converts bytes to string
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		// returning new Error becuase the *err* is nil that is why we reached this else condition
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid} // variable
	return &rate, nil
}
