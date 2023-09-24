package elastic

import (
	"github.com/op/go-logging"
	"os"
)

func GetElasticLogger() *logging.Logger {
	var elasticLogger = logging.MustGetLogger("elastic")
	var backend = logging.NewLogBackend(os.Stdout, "[Elastic] ", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.DEBUG, "")

	return elasticLogger
}

var logger = GetElasticLogger()
