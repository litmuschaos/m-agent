[![Go Report Card](https://goreportcard.com/badge/github.com/litmuschaos/m-agent)](https://goreportcard.com/report/github.com/litmuschaos/m-agent)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Flitmuschaos%2Fm-agent.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Flitmuschaos%2Fm-agent?ref=badge_shield)
![GitHub Workflow](https://github.com/litmuschaos/m-agent/actions/workflows/push.yml/badge.svg?branch=master)

# m-agent
Machine Agent a.k.a. m-agent is a lightweight, platform-generic daemon agent that can remotely inject faults into machine scoped resources, as part of the LitmusChaos Experiments.

# Requirements
- Linux OS [Tested on Ubuntu, CentOS, RHEL, and SUSE Linux]
- systemd

# Installation
To install m-agent in your target machine, you can execute the following commands in the target machine:
```
$ curl -fsSL -o get_m-agent.sh https://raw.githubusercontent.com/litmuschaos/m-agent/master/scripts/install.sh
$ chmod 700 get_m-agent.sh
$ ./get_m-agent.sh
```

You can specify any particular m-agent version for installation:
```
$ ./get_m-agent.sh --version <VERSION>
```

By default, `41365` port is utilized by m-agent. You can specify a custom port at which m-agent should listen for client messages:
```
$ ./get_m-agent.sh --port <PORT>
```

Finally, you can specify if the installation can take place without using `sudo`, if it is not present:
```
$ ./get_m-agent.sh --no-sudo
```

# Usage
```
Usage: m-agent [options]

options:
  -get-token
    	generates a token to be used for the authentication of the requests made to the agent
  -token-expiry-duration string
    	token expiry duration (non-interactive mode)
  -update-port string
      update the m-agent server port
```

Upon installing m-agent, you can use it to generate a token for your Chaos Experiment. It step will require you to specify an expiry duration for your token. Tokens are valid through a minimum duration of 1 minute to a maximum of 30 days. The token can be generated in two modes:
1. Interactive Mode
2. Non-Interactive Mode

To generate a token in an interactive mode, use the `-get-token` boolean flag, which will prompt you to select the expiry duration for the token:
```
m-agent -get-token
``` 

The non-interactive mode can be used to generate token with more flexibility in terms of its expiry duration. Use the `-token-expiry-duration` flag along with the `-get-token` flag to use this mode. `-token-expiry-duration` is a string flag which expects the expiry duration of the token to be specified in the form of a numeric value suffixed with a single alphabet out of 'm' or 'M', 'h' or 'H', and 'd' or 'D' denoting minutes, hours, and days respectively.

For minutes, the corresponding value must lie in between 1 and 60, inclusively. For hours, the corresponding value must lie in the range of 1 to 24, inclusively. Lastly, for days, the corresponding value must lie between 1 to 30, inclusively.

As an instance, to create a token with a validity of 30 minutes, one can use the following command:
```
m-agent -get-token -token-expiry-duration 30m
```

Similarly, a token valid for 15 days from the time of creation can be specified as:
```
m-agent -get-token -token-expiry-duration 15D
```

You can also update the port at which m-agent should listen for the client messages:
```
sudo m-agent -updated-port <NEW_PORT>
sudo systemctl restart m-agent
```  

# Uninstallation
If you wish to uninstall m-agent, you can execute the following commands in the target machine:
```bash
$ curl -fsSL -o remove_m-agent.sh https://raw.githubusercontent.com/litmuschaos/m-agent/master/scripts/uninstall.sh
$ chmod 700 remove_m-agent.sh
$ ./remove_m-agent.sh
```
