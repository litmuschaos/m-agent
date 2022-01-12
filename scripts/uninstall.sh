#!/usr/bin/env bash
: ${BINARY_NAME:="m-agent"}
: ${M_AGENT_INSTALL_DIR:="/usr/local/bin"}
: ${SERVICE_NAME:="m-agent.service"}
: ${SERVICE_DIR:="/etc/systemd/system"}
: ${USE_SUDO:="true"}

HAS_M_AGENT_BINARY=$(test -f $M_AGENT_INSTALL_DIR/$BINARY_NAME && echo true || echo false)
HAS_M_AGENT_SERVICE=$(test -f $SERVICE_DIR/$SERVICE_NAME && echo true || echo false)
HAS_SYSTEMD="$(type "systemctl" &> /dev/null && echo true || echo false)"

# runs the given command as root (detects if we are root already)
runAsRoot() {
  if [ $EUID -ne 0 ] && [ "$USE_SUDO" == "true" ]; then
    sudo "${@}"
  else
    "${@}"
  fi
}

# checkRequirements checks if files to be deleted exist and systemctl is present
checkRequirements() {
  
  if [ "${HAS_M_AGENT_BINARY}" != "true" ]; then
    echo "$BINARY_NAME binary not found"
    exit 1
  fi

  if [ "${HAS_M_AGENT_SERVICE}" != "true" ]; then
    echo "$SERVICE_NAME unit file not found"
    exit 1
  fi

  if [ "${HAS_SYSTEMD}" != "true" ]; then
    echo "Systemd is required for installation"
    exit 1
  fi
}

uninstall() {
  if [ "$(systemctl is-active $SERVICE_NAME)" == "active" ]; then
    runAsRoot systemctl stop $SERVICE_NAME
    echo "Stopped $BINARY_NAME service"
  fi

  if [ "$(systemctl is-enabled $SERVICE_NAME)" == "enabled" ]; then
    runAsRoot systemctl disable $SERVICE_NAME
    echo "Disabled $BINARY_NAME service"
  fi

  runAsRoot rm "$SERVICE_DIR/$SERVICE_NAME"
  echo "Removed $SERVICE_NAME unit file"

  runAsRoot rm "$M_AGENT_INSTALL_DIR/$BINARY_NAME"
  echo "Removed $BINARY_NAME binary file"
}

checkRequirements
uninstall
