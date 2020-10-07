package main

import (
	"fmt"
	"github.com/h4rimu/kaspro-sdkv2/app/services"
	"github.com/h4rimu/kaspro-sdkv2/caching"
	opLogging "github.com/h4rimu/kaspro-sdkv2/logging"
	"github.com/h4rimu/kaspro-sdkv2/utils"
	"net/http"
)

var log = opLogging.MustGetLogger("kaspro-sdkv2")

func main() {

	testClientRetry()
}

func loggingTesting() {
	log.Infof("INFO", "INFO MESSAGE")
	log.Fatalf("FATAL", "FATAL MESSAGE")
	log.Errorf("ERROR", "ERROR MESSAGE")
	log.Debugf("DEBUG", "DEBUG MESSAGE")
}

func simpleCacheTesting() {
	simpleCache := caching.SimpleCache{
		ExpiredTime: 5,
		PurgeTime:   10,
	}

	globaclCache := simpleCache.CreateNewCache()

	simpleCache.SetValue("A", "B", *globaclCache)
	asdfaf := simpleCache.GetValue("A", *globaclCache)
	fmt.Println("ASDKASLASJDJ ", *asdfaf)
}

func testClientRetry() {

	urlApi := utils.SetApiUrl("INTERNAL", "http://dev.kaspro.id/DCZ4DmJMPVKsX75/081513245412/validate")

	clientRetry := services.ClientRetry{
		MaxRetry:      10,
		HttpToRetry:   SetHttpToRetry(),
		BeginInterval: 5,
		EndInterval:   5,
	}

	services := services.ClientParty{
		UniqueID:    "INTERNAL",
		ClientName:  "KASPRO",
		HttpMethod:  http.MethodGet,
		Debug:       true,
		ClientRetry: &clientRetry,
		UrlApi:      *urlApi,
		HttpClient:  http.Client{},
		Headers:     *SetKasproHeaderContent(),
		RequestBody: nil,
	}

	response, err := services.HitClient()
	if err != nil {
		fmt.Println("ERROR ", *err)
	}

	fmt.Println("HTTP CODE ", string(response.ByteResponse))
}

func SetKasproHeaderContent() *[]map[string]string {

	var headers []map[string]string
	mapContentType := make(map[string]string)
	mapContentType["Content-Type"] = "application/json"
	headers = append(headers, mapContentType)

	mapAcceptLanguage := make(map[string]string)
	mapAcceptLanguage["Accept-Language"] = "EN"
	headers = append(headers, mapAcceptLanguage)

	headerToken := make(map[string]string)
	headerToken["Token"] = "WLu28cXFYvrdtQ7KFNxDUI3hpufmj+EbNknAEL9i7pfdjx69s/lnu3YSScaxUv+7Iere9Or5f1AvNC3rO8l+U3gkcU87vUrlHu6llGJeZiolpM2mD1ZePTlPyjVrArkmlK5Ui8vnGmu55anh2jq2Y4KD9HIj2FI8ENzfFqPX3/vmVH2e8ImkxsDuK1Ot+oH6BVxUKThhqcVPFfv3Qe52AA=="
	headers = append(headers, headerToken)
	return &headers
}

func SetHttpToRetry() *[]map[int]string {

	var retries []map[int]string
	httpCode200 := make(map[int]string)
	httpCode200[http.StatusBadRequest] = "Bad Request"
	retries = append(retries, httpCode200)

	return &retries
}
