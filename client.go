package evengclient

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	// path to the eve_ng api endpoint
	endpointPath = "api/"
)

type client struct {
	*clientData
}

/*
clientData - Contains data of a client
*/
type clientData struct {
	baseURL  string
	username string
	password string

	resty   *resty.Client
	useAuth bool
}

/*
NotValidError - Is returned when the client was not initialized properly
*/
type NotValidError struct{}

func (m *NotValidError) Error() string {
	return "client was not created properly with the func New...Client(baseURL string)"
}

/*
isValid - returns true if a client is valid and false if a client is invalid
*/
func (c *client) isValid() bool {
	return c.clientData != nil
}

/*
SetUsernameAndPassword - Is used to set a username and password for https auth
*/
func (c *client) SetUsernameAndPassword(username, password string) error {
	if !c.isValid() {
		return &NotValidError{}
	}
	if username == "" {
		return errors.New("invalid username")
	}
	if password == "" {
		return errors.New("invalid password")
	}
	c.username = username
	c.password = password
	c.useAuth = true
	return nil
}

/*
init is used for initialising the config file
*/
func init() {
	// Search config in home directory with name "eve-ng-api" (without extension).
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("eve-ng-api")

	//Set env var prefix to only match certain vars
	viper.SetEnvPrefix("EVE_NG_API")

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	viper.ReadInConfig()
}

/*
request - Is used to send either GET, POST, PUT or DELETE requests
*/
func (c *client) request(method string, path string, body string, header, queryParams map[string]string) (*resty.Response, error) {
	request := c.resty.R()
	request.SetHeader("Content-Type", "application/json")

	if header != nil {
		request.SetHeaders(header)
	}

	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	if body != "" {
		request.SetBody(body)
	}

	if c.useAuth {
		request.SetBasicAuth(c.username, c.password)
	}

	var response *resty.Response
	response = nil

	var err error
	err = nil

	switch method {
	case "GET":
		response, err = request.Get(c.baseURL + urlEscapePath(path))
	case "POST":
		response, err = request.Post(c.baseURL + urlEscapePath(path))
	case "PUT":
		response, err = request.Put(c.baseURL + urlEscapePath(path))
	case "DELETE":
		response, err = request.Delete(c.baseURL + urlEscapePath(path))
	default:
		return nil, errors.New("invalid http method: " + method)
	}
	if err != nil {
		return nil, errors.Wrap(err, "error during http request")
	}
	if response.StatusCode() != 200 && response.StatusCode() != 201 {
		return nil, errors.Wrap(getHTTPError(response), "http request responded with an error")
	}

	return response, err
}

//Http error handling

/*
HTTPError - Represents an http error returned by the api.
*/
type HTTPError struct {
	StatusCode int
	Status     string
	Body       *ErrorResponse
}

/*
ErrorResponse - Contains error information.
*/
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (h HTTPError) Error() string {
	msg := "http error: status code: " + strconv.Itoa(h.StatusCode) + " // status: " + h.Status
	if h.Body != nil {
		msg += " // message: " + h.Body.Message
	}
	return msg
}

func getHTTPError(response *resty.Response) error {
	httpError := HTTPError{
		StatusCode: response.StatusCode(),
		Status:     response.Status(),
	}
	var errorResponse ErrorResponse
	err := json.Unmarshal(response.Body(), &errorResponse)
	if err != nil {
		return httpError
	}
	httpError.Body = &errorResponse
	return httpError
}

//---------- helper functions ----------//

/*
unmarshalDataIntoStruct - Is used to unmarshal responses from the eve-ng REST API into certain structs
*/
func (c *client) unmarshalDataIntoStruct(responseBody []byte, i interface{}) error {
	isResponseDataEmpty, err := checkForEmptyResponseData(responseBody)
	if err != nil {
		return err
	}
	if !isResponseDataEmpty {
		var basicResponse BasicResponse
		basicResponse.Data = i
		err := json.Unmarshal(responseBody, &basicResponse)
		if err != nil {
			return err
		}
	}

	return nil
}

//the following lines of code have to be done because otherwise a empty data provided in the api response could not be correctly unmarshaled
func checkForEmptyResponseData(responseBody []byte) (bool, error) {
	var basicResponse BasicResponse
	err := json.Unmarshal(responseBody, &basicResponse)
	if err != nil {
		return false, err
	}
	emptyData := make([]interface{}, 0)
	return reflect.DeepEqual(emptyData, basicResponse.Data), nil
}

/*
urlEscapePath - Escapes special characters of a given url path
*/
func urlEscapePath(unescaped string) string {
	arr := strings.Split(unescaped, "/")
	for i, partString := range strings.Split(unescaped, "/") {
		arr[i] = url.QueryEscape(partString)
	}
	return strings.Join(arr, "/")
}

/*
jsonEscape - Escapes special characters of a given json string
*/
func jsonEscape(unescaped string) (string, error) {
	escaped, err := json.Marshal(unescaped)
	if err != nil {
		return "", errors.Wrap(err, "json marshal failed")
	}
	return string(escaped)[1 : len(escaped)-1], nil
}
