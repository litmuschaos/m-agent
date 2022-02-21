package log

import (
	"log"
	"os"

	"github.com/litmuschaos/m-agent/internal/m-agent/errorcodes"
)

// GetTokenErrorLogger returns a logger for handling token errors
func GetTokenErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetTokenErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetClientMessageReadErrorLogger returns a logger for handling client message read errors
func GetClientMessageReadErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetClientMessageReadErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetSteadyStateCheckErrorLogger returns a logger for handling steady state check errors
func GetSteadyStateCheckErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetSteadyStateCheckErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetExecuteExperimentErrorLogger returns a logger for handling execute experiment errors
func GetExecuteExperimentErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetExecuteExperimentErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetCommandProbeExecutionErrorLogger returns a logger for handling command probe execution errors
func GetCommandProbeExecutionErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetCommandProbeExecutionErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetInvalidActionErrorLogger returns a logger for handling invalid action errors
func GetInvalidActionErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetInvalidActionErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

// GetChaosAbortErrorLogger returns a logger for handling chaos abort errors
func GetChaosAbortErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetChaosAbortErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func GetLivenessCheckErrorLogger() *log.Logger {

	return log.New(os.Stdout, errorcodes.GetLivenessCheckErrorPrefix(), log.Ldate|log.Ltime|log.Lmsgprefix)
}
