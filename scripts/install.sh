#!/usr/bin/env bash

: ${BINARY_NAME:="m-agent"}
: ${USE_SUDO:="true"}
: ${DEBUG:="false"}
: ${M_AGENT_INSTALL_DIR:="/usr/local/bin"}
: ${SERVICE_NAME:="m-agent.service"}
: ${SERVICE_DIR:="/etc/systemd/system"}
: ${PORT:="41365"}

HAS_CURL="$(type "curl" &> /dev/null && echo true || echo false)"
HAS_WGET="$(type "wget" &> /dev/null && echo true || echo false)"
HAS_SYSTEMD="$(type "systemctl" &> /dev/null && echo true || echo false)"

HAS_STRESS_NG="$(type "stress-ng" &> /dev/null && echo true || echo false)"

# initArch discovers the architecture for this system
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="arm";;
    armv6*) ARCH="arm";;
    armv7*) ARCH="arm";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

# initOS discovers the operating system for this system
initOS() {
  OS=$(uname|tr '[:upper:]' '[:lower:]')
}

# runs the given command as root (detects if we are root already)
runAsRoot() {
  if [ $EUID -ne 0 ] && [ "$USE_SUDO" == "true" ]; then
    sudo "${@}"
  else
    "${@}"
  fi
}

# verifySupported checks that the os/arch combination is supported for
# binary builds, as well whether or not necessary tools are present
verifySupported() {
  local supported="linux-386\nlinux-amd64\nlinux-arm\nlinux-arm64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    exit 1
  fi

  if [ "${HAS_SYSTEMD}" != "true" ]; then
    echo "Systemd is required for installation"
    exit 1
  fi

  if [ "${HAS_CURL}" != "true" ] && [ "${HAS_WGET}" != "true" ]; then
    echo "Either curl or wget is required"
    exit 1
  fi

  if [ "$(netstat -tulpn | grep $PORT)" != "" ]; then
    echo "$PORT port is not available"
    exit 1
  fi

  if [ "${HAS_STRESS_NG}" != "true" ]; then
    printf "stress-ng not found, please install stress-ng:\nhttps://snapcraft.io/install/stress-ng/ubuntu"
    exit 1
  fi
}

# checkDesiredVersion checks if the desired version is available
checkDesiredVersion() {

  # Get tag using github api
  local api_release_tags="https://api.github.com/repos/litmuschaos/m-agent/tags"
    
  if [ "${HAS_CURL}" == "true" ]; then
    TAGS=$(curl -Ls $api_release_tags | grep '"name":' | awk '{print $2}' | tr -d '"' | tr -d "v" | tr -d ",")
  elif [ "${HAS_WGET}" == "true" ]; then
    TAGS=$(wget $api_release_tags -O - 2>&1 | grep '"name":' | awk '{print $2}' | tr -d '"' | tr -d "v" | tr -d ",")
  fi

  if [ "x$DESIRED_VERSION" == "x" ]; then

    # fallback to master version if no release tags are found
    if [ -z "$TAGS" ]; then
      TAG="master"
    else
      TAG=$(echo $TAGS | head -1)
    fi
  else
    if [[ $TAGS == *"$DESIRED_VERSION"* ]]; then
      TAG=$DESIRED_VERSION
    else
      echo "$DESIRED_VERSION not found"
      exit 1
    fi
  fi
}

# isMAgentInstalled checks if m-agent is already installed
isMAgentInstalled() {
  if [[ -f "${M_AGENT_INSTALL_DIR}/${BINARY_NAME}" ]]; then
    echo "$BINARY_NAME already exits."
    return 0
  else
    return 1
  fi
}

# downloadFile downloads the latest binary package 
downloadFile() {
  M_AGENT_DIST="m-agent-$OS-$ARCH-$TAG.tar.gz"
  BIN_DOWNLOAD_URL="https://m-agent.s3.amazonaws.com/$M_AGENT_DIST"
  M_AGENT_TMP_ROOT="$(mktemp -dt m-agent-installer-XXXXXX)"
  M_AGENT_TMP_FILE="$M_AGENT_TMP_ROOT/$M_AGENT_DIST"
  echo "Downloading $BIN_DOWNLOAD_URL"
  if [ "${HAS_CURL}" == "true" ]; then
    curl -SsL "$BIN_DOWNLOAD_URL" -o "$M_AGENT_TMP_FILE"
  elif [ "${HAS_WGET}" == "true" ]; then
    wget -q -O "$M_AGENT_TMP_FILE" "$BIN_DOWNLOAD_URL"
  fi
}

# installFile installs the m-agent binary
installFile() {
  M_AGENT_TMP="$M_AGENT_TMP_ROOT/$BINARY_NAME"
  mkdir -p "$M_AGENT_TMP"
  tar xf "$M_AGENT_TMP_FILE" -C "$M_AGENT_TMP"
  M_AGENT_TMP_BIN="$M_AGENT_TMP/m-agent"
  echo "Preparing to install $BINARY_NAME into ${M_AGENT_INSTALL_DIR}"
  runAsRoot cp "$M_AGENT_TMP_BIN" "$M_AGENT_INSTALL_DIR/$BINARY_NAME"
  echo "$BINARY_NAME installed into $M_AGENT_INSTALL_DIR/$BINARY_NAME"
}

createConfigFile() {
  runAsRoot mkdir "/etc/$BINARY_NAME"
  runAsRoot echo $PORT > "$M_AGENT_TMP_ROOT/PORT"
  runAsRoot cp "$M_AGENT_TMP_ROOT/PORT" "/etc/$BINARY_NAME/"
  echo "$BINARY_NAME server PORT config file created"
}

# setupService downloads the service unit file, set it up and start the service
setupService() {
  M_AGENT_TMP_SERVICE="$M_AGENT_TMP_ROOT/$SERVICE_NAME"
  SERVICE_DOWNLOAD_URL="https://m-agent.s3.amazonaws.com/$SERVICE_NAME"
  echo "Downloading $SERVICE_DOWNLOAD_URL"
  if [ "${HAS_CURL}" == "true" ]; then
    curl -SsL "$SERVICE_DOWNLOAD_URL" -o "$M_AGENT_TMP_SERVICE"
  elif [ "${HAS_WGET}" == "true" ]; then
    wget -q -O "$M_AGENT_TMP_SERVICE" "$SERVICE_DOWNLOAD_URL"
  fi
  echo "Setting up $BINARY_NAME into ${SERVICE_DIR}"
  runAsRoot cp "$M_AGENT_TMP_SERVICE" "$SERVICE_DIR/$SERVICE_NAME"
  runAsRoot chmod 755 "$SERVICE_DIR/$SERVICE_NAME"
  runAsRoot systemctl enable "$SERVICE_NAME"
  runAsRoot systemctl start "$SERVICE_NAME"
}

# fail_trap is executed if an error occurs
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    if [[ -n "$INPUT_ARGUMENTS" ]]; then
      echo "Failed to install $BINARY_NAME with the arguments provided: $INPUT_ARGUMENTS"
      help
    else
      echo "Failed to install $BINARY_NAME"
    fi
    echo -e "\tFor support, go to https://github.com/litmuschaos/m-agent."
  fi
  cleanup
  exit $result
}

# help provides possible cli installation arguments
help () {
  echo "Accepted cli arguments are:"
  echo -e "\t[--help|-h ] ->> prints this help"
  echo -e "\t[--version|-v <desired_version>] ->> When not defined it fetches the latest release from GitHub"
  echo -e "\te.g. --version 1.0.0 or -v master"
  echo -e "\t[--no-sudo]  ->> install without sudo"
  echo -e "\t[--port|-p <port>] ->> custom port for m-agent server"
}

# cleanup temporary files 
cleanup() {
  if [[ -d "${M_AGENT_TMP_ROOT:-}" ]]; then
    rm -rf "$M_AGENT_TMP_ROOT"
  fi
}

# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e

# Set debug if desired
if [ "${DEBUG}" == "true" ]; then
  set -x
fi

# Parsing input arguments (if any)
export INPUT_ARGUMENTS="${@}"
set -u
while [[ $# -gt 0 ]]; do
  case $1 in
    '--version'|-v)
       shift
       if [[ $# -ne 0 ]]; then
           export DESIRED_VERSION="${1}"
       else
           echo -e "Please provide the desired version. e.g. --version v3.0.0 or -v canary"
           exit 0
       fi
       ;;
    '--port'|-p)
       shift
       if [[ $# -ne 0 ]]; then
          PORT="${1}"
       else
          echo -e "Please provide the port for m-agent server. e.g. 41365"
          exit 0
       fi
       ;;
    '--no-sudo')
       USE_SUDO="false"
       ;;
    '--help'|-h)
       help
       exit 0
       ;;
    *) exit 1
       ;;
  esac
  shift
done
set +u

initArch
initOS
verifySupported
checkDesiredVersion
if ! isMAgentInstalled; then
  downloadFile
  installFile
  createConfigFile
  setupService
fi
cleanup