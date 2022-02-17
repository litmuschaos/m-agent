package experiment

import (
	"bytes"
	"log"
	"net/http"
	"os/exec"

	"github.com/litmuschaos/m-agent/internal/m-agent/errorcodes"
	logger "github.com/litmuschaos/m-agent/internal/m-agent/log"
	"github.com/litmuschaos/m-agent/internal/m-agent/messages"
	"github.com/litmuschaos/m-agent/internal/m-agent/upgrader"
	"github.com/litmuschaos/m-agent/pkg/cpu"
	"github.com/litmuschaos/m-agent/pkg/probes"
)

var (
	cmd            *exec.Cmd
	stdout, stderr bytes.Buffer
)

// CPUStress listens for the client actions and executes them as appropriate
func CPUStress(w http.ResponseWriter, r *http.Request) {

	// upgrade the connection to a websocket connection
	upgrader := upgrader.GetConnectionUpgrader()

	clientMessageReadLogger := logger.GetClientMessageReadErrorLogger()
	steadyStateCheckErrorLogger := logger.GetSteadyStateCheckErrorLogger()
	executeExperimentErrorLogger := logger.GetExecuteExperimentErrorLogger()
	commandProbeExecutionErrorLogger := logger.GetCommandProbeExecutionErrorLogger()
	invalidActionErrorLogger := logger.GetCommandProbeExecutionErrorLogger()
	chaosRemediationErrorLogger := logger.GetChaosRemediationErrorLogger()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to establish connection with client, %v", err)
		return
	}

	for {

		action, payload, err := messages.ListenForClientMessage(conn)
		if err != nil {
			if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetClientMessageReadErrorPrefix()+err.Error()); err != nil {

				clientMessageReadLogger.Printf("Error occured while sending error message to client, %v", err)
			}
			conn.Close()
			return
		}

		switch action {

		case "CHECK_STEADY_STATE":
			if err := cpu.CheckForStressNG(); err != nil {

				if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetSteadyStateCheckErrorPrefix()+err.Error()); err != nil {
					steadyStateCheckErrorLogger.Printf("Error occured while sending error message to client, %v", err)
				}

				conn.Close()
				return
			}

			if err := messages.SendMessageToClient(conn, "ACTION_SUCCESSFUL", messages.Message{}); err != nil {

				steadyStateCheckErrorLogger.Printf("Error occured while sending feedback message to client, %v", err)
				conn.Close()
				return
			}

		case "EXECUTE_EXPERIMENT":
			cmd, err = cpu.StressCPU(payload, &stdout, &stderr, conn)
			if err != nil {
				if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetExecuteExperimentErrorPrefix()+err.Error()); err != nil {
					executeExperimentErrorLogger.Printf("Error occured while sending error message to client, %v", err)
				}
				conn.Close()
				return
			}

			if err := messages.SendMessageToClient(conn, "ACTION_SUCCESSFUL", messages.Message{}); err != nil {
				executeExperimentErrorLogger.Printf("Error occured while sending feedback message to client, %v", err)
				conn.Close()
				return
			}

		case "EXECUTE_COMMAND":
			cmdProbeStdout, err := probes.ExecuteCmdProbeCommand(payload)

			if err != nil {
				if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetCommandProbeExecutionErrorPrefix()+err.Error()); err != nil {
					commandProbeExecutionErrorLogger.Printf("Error occured while sending error message to client, %v", err)
					conn.Close()
					return
				}
			}

			if err := messages.SendMessageToClient(conn, "ACTION_SUCCESSFUL", cmdProbeStdout); err != nil {
				commandProbeExecutionErrorLogger.Printf("Error occured while sending feedback message to client, %v", err)
				conn.Close()
				return
			}

		case "ABORT_STRESS_PROCESS":
			if err := cpu.AbortStressNGProcess(cmd, stderr); err != nil {
				if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetChaosRemediationErrorPrefix()+err.Error()); err != nil {
					chaosRemediationErrorLogger.Printf("Error occured while sending error message to client, %v", err)
				}
				conn.Close()
				return
			}

			if err := messages.SendMessageToClient(conn, "ACTION_SUCCESSFUL", stdout.String()); err != nil {
				chaosRemediationErrorLogger.Printf("Error occured while sending feedback message to client, %v", err)
				conn.Close()
				return
			}

		default:
			if err := messages.SendMessageToClient(conn, "ERROR", errorcodes.GetInvalidActionErrorPrefix()+"Invalid action: "+action); err != nil {
				invalidActionErrorLogger.Printf("Error occured while sending error message to client, %v", err)
			}
			conn.Close()
			return
		}
	}
}
