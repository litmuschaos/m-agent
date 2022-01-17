package log

import (
	"log"
	"os"

	"github.com/litmuschaos/m-agent/internal/m-agent/errorcodes"
)

func GetTokenErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetTokenErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetClientMessageReadErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetClientMessageReadErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetSteadyStateCheckErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetSteadyStateCheckErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetExecuteExperimentErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetExecuteExperimentErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetCommandProbeExecutionErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetCommandProbeExecutionErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetInvalidActionErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetInvalidActionErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}
