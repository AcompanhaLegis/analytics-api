package dadosabertos

import (
	"github.com/op/go-logging"
	"os"
)

func GetDadosAbertosLogger() *logging.Logger {
	var dadosabertosLogger = logging.MustGetLogger("dadosabertos")
	var backend = logging.NewLogBackend(os.Stdout, "[DadosAbertos] ", 0)
	var backendLeveled = logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.DEBUG, "")

	dadosabertosLogger.SetBackend(backendLeveled)
	return dadosabertosLogger
}

var logger = GetDadosAbertosLogger()
