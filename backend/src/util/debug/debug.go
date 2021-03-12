package debug

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

func Error(context string, message string) {
	errorMessage := fmt.Sprintf("%s - %s - Error - %s", time.Now().String(), context, message)
	log.Errorf(errorMessage)
}

func Warning(context string, message string) {
	warningMessage := fmt.Sprintf("%s - %s - Info - %s", time.Now().String(), context, message)
	log.Warn(warningMessage)
}

func Info(context string, message string) {
	infoMessage := fmt.Sprintf("%s - %s - Info - %s", time.Now().String(), context, message)
	log.Info(infoMessage)
}
