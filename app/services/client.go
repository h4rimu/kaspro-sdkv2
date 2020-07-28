package services

import (
	"github.com/h4rimu/kaspro-sdkv2/logging"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var log = logging.MustGetLogger("kaspro-sdkv2")

type ThirdParty interface {
	HitClient(url url.URL) (*ClientResponse, *error)
}

type ClientParty struct {
	UniqueID    string
	ClientName  string
	HttpMethod  string
	Debug       bool
	Retry       int
	UrlApi      url.URL
	HttpClient  http.Client
	Headers     []map[string]string
	RequestBody io.Reader
}

type ClientResponse struct {
	HttpCode     int
	ByteResponse []byte
}

func (c *ClientParty) HitClient() (*ClientResponse, *error) {

	request, err := http.NewRequest(c.HttpMethod, c.UrlApi.String(), c.RequestBody)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	for _, element := range c.Headers {
		for key, value := range element {
			if strings.ToLower(key) == "baseauth" {
				baseAuth := strings.Split(value, ":")
				request.SetBasicAuth(baseAuth[0], baseAuth[1])
			} else {
				request.Header.Set(key, value)
			}
		}
	}

	if c.Debug {
		log.Debugf(c.UniqueID, "===================================================== HIT "+strings.ToUpper(c.ClientName)+" API =======================================================")
		log.Debugf(c.UniqueID, "HTTP METHOD %s ", request.Method)
		log.Debugf(c.UniqueID, "URL %s ", request.URL)
		log.Debugf(c.UniqueID, "HEADER %s ", request.Header)
		if request.Body != nil {
			log.Debugf(c.UniqueID, "BODY %s ", request.Body)
		}
		log.Debugf(c.UniqueID, "==========================================================================================================================")
	}

	if c.Retry > 1 {
		for i := 0; i < c.Retry; i++ {

			log.Debugf(c.UniqueID, "Retry Hit To Client ", i)
			response, err := c.HttpClient.Do(request)

			if err != nil {
				log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
				return nil, &err
			}

			if response.StatusCode != http.StatusRequestTimeout || response.StatusCode != http.StatusGatewayTimeout {
				byteResult, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
					return nil, &err
				}
				printClientResponse(c.Debug, c.UniqueID, c.ClientName, string(byteResult))
				clientResponse := ClientResponse{
					HttpCode:     response.StatusCode,
					ByteResponse: byteResult,
				}
				return &clientResponse, nil
			}

			byteResult, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
				return nil, &err
			}
			printClientResponse(c.Debug, c.UniqueID, c.ClientName, string(byteResult))
		}
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	printClientResponse(c.Debug, c.UniqueID, c.ClientName, string(byteResult))

	clientResponse := ClientResponse{
		HttpCode:     response.StatusCode,
		ByteResponse: byteResult,
	}
	return &clientResponse, nil
}

func printClientResponse(debug bool, uniqueID string, clientName string, strResult string) {
	if debug {
		log.Debugf(uniqueID, "================================================ RESPONSE "+strings.ToUpper(clientName)+" API ================================================")
		log.Debugf(uniqueID, "GET RESPONSE FROM "+strings.ToUpper(clientName)+" API %s", strResult)
		log.Debugf(uniqueID, "======================================================================================================================")
	}
}
