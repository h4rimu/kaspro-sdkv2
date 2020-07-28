package utils

import (
	"encoding/json"
	"encoding/xml"
)

func JsonToString(uniqueID string, param interface{}) string {
	result, err := json.MarshalIndent(&param, "", "\t")
	if err != nil {
		log.Fatalf(uniqueID, "Error occurred ", err.Error())
		return ""
	}

	strResult := string(result)
	return strResult
}

func XmlToString(uniqueID string, param interface{}) string {
	result, err := xml.MarshalIndent(&param, "", "\t")
	if err != nil {
		log.Errorf(uniqueID, "Error occurred ", err.Error())
		return ""
	}

	strResult := string(result)
	return strResult
}
