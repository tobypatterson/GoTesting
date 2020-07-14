package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/hello/domain/location"
	"example.com/hello/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*location.Country, *errors.ApiError) {

	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))

	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient error when getting country %s", countryId),
		}
	}

	if response.StatusCode > 209 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {

			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
			}
		}

		return nil, &apiErr
	}

	var result location.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &result, nil
}
