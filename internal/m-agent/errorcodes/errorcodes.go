// Copyright 2022 LitmusChaos Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package errorcodes

const (
	tokenErrorCode   = "2000"
	tokenErrorString = "TOKEN_ERROR"

	authErrorCode   = "2001"
	authErrorString = "AUTH_ERROR"

	clientMessageReadErrorCode   = "2002"
	clientMessageReadErrorString = "CLIENT_MESSAGE_READ_ERROR"

	steadyStateCheckErrorCode   = "2003"
	steadyStateCheckErrorString = "STEADY_STATE_CHECK_ERROR"

	executeExperimentErrorCode   = "2004"
	executeExperimentErrorString = "EXPERIMENT_EXECUTION_ERROR"

	commandProbeExecutionErrorCode   = "2005"
	commandProbeExecutionErrorString = "COMMAND_EXECUTION_ERROR"

	invalidActionErrorCode   = "2006"
	invalidActionErrorString = "INVALID_ACTION_ERROR"

	chaosAbortErrorCode   = "2007"
	chaosAbortErrorString = "CHAOS_ABORT_ERROR"

	livenessCheckErrorCode   = "2008"
	livenessCheckErrorString = "LIVENESS_CHECK_ERROR"

	closeConnectionErrorCode   = "2009"
	closeConnectionErrorString = "CLOSE_CONNECTION_ERROR"

	chaosRevertErrorCode   = "2010"
	chaosRevertErrorString = "CHAOS_REVERT_ERROR"
)

func GetTokenErrorPrefix() string {
	return tokenErrorCode + ": " + tokenErrorString + ": "
}

func GetAuthErrorPrefix() string {
	return authErrorCode + ": " + authErrorString + ": "
}

func GetClientMessageReadErrorPrefix() string {
	return clientMessageReadErrorCode + ": " + clientMessageReadErrorString + ": "
}

func GetSteadyStateCheckErrorPrefix() string {
	return steadyStateCheckErrorCode + ": " + steadyStateCheckErrorString + ": "
}

func GetExecuteExperimentErrorPrefix() string {
	return executeExperimentErrorCode + ": " + executeExperimentErrorString + ": "
}

func GetCommandProbeExecutionErrorPrefix() string {
	return commandProbeExecutionErrorCode + ": " + commandProbeExecutionErrorString + ": "
}

func GetInvalidActionErrorPrefix() string {
	return invalidActionErrorCode + ": " + invalidActionErrorString + ": "
}

func GetChaosAbortErrorPrefix() string {
	return chaosAbortErrorCode + ": " + chaosAbortErrorString + ": "
}

func GetLivenessCheckErrorPrefix() string {
	return livenessCheckErrorCode + ": " + livenessCheckErrorString + ": "
}

func GetCloseConnectionErrorPrefix() string {
	return closeConnectionErrorCode + ": " + closeConnectionErrorString + ": "
}

func GetChaosRevertErrorPrefix() string {
	return chaosRevertErrorCode + ": " + chaosRevertErrorString + ": "
}
