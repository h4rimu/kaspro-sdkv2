package utils

import (
	"github.com/h4rimu/kaspro-sdkv2/app/types"
	"github.com/h4rimu/kaspro-sdkv2/database"
	opLogging "github.com/h4rimu/kaspro-sdkv2/logging"
)

var MMEN *types.MessageMap
var MMID *types.MessageMap
var DBModel *database.DBModel

var log = opLogging.MustGetLogger("kaspro-sdkv2")
