package main

import (
	"fmt"
	"github.com/h4rimu/kaspro-sdkv2/caching"
	opLogging "github.com/h4rimu/kaspro-sdkv2/logging"
)

var log = opLogging.MustGetLogger("kaspro-sdkv2")

func main() {
	loggingTesting()
	simpleCacheTesting()
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
	fmt.Println("ASDKASLASJDJ ",*asdfaf)
}
