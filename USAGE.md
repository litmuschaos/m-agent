# m-agent Usage Docs for process-kill experiment

A simple chaos experiment that you can execute using m-agent is process-kill experiment. As the name suggests, a process-kill chaos causes any number of processes in your target machine to be killed ungracefully. In this document, we'll learn how we can execute this chaos experiment using ChaosCenter.

**Note:** The killed processes are not recovered by the experiment i.e. there's no revert of the chaos injection as part of the experiment execution.  

## How to Execute process-kill Experiment?
### STEP-1: Install m-agent in Target Machine
To install m-agent, execute the following commands in your target machine:
```
$ curl -fsSL -o get_m-agent.sh https://raw.githubusercontent.com/litmuschaos/m-agent/master/scripts/install.sh
$ chmod 700 get_m-agent.sh
$ ./get_m-agent.sh
```

You can find the advanced installation options [here](./README.md#installation).

### STEP-2: Generate and Add Agent Token
An authentication token is necessary for communicating with the m-agent from the experiment pod. It ensures that only the experiment pod will be authorized to access the m-agent.

To generate the authentication token, execute the following command:
```
m-agent -get-token
```

Choose the duration through which you want the token to remain valid in the prompt. If you want to specify a different expiry duration for the token, you can use [non-interactive](./README.md#usage) token generation.

Once you have the token and agent endpoint, create a k8s secret YAML manifest `agent-secret.yaml` by updating the following template:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: agent-secret
type: Opaque
stringData:
  AGENT_ENDPOINT: <YOUR_AGENT_ENDPOINT>
  AUTH_TOKEN: <TOKEN>
```

Finally, apply this manifest in the `litmus` namespace:
```
kubectl apply -f agent-secret.yaml -n litmus
```

### STEP-3: Add m-agent ChaosHub
Presently, the process-kill experiment charts are not located in the public Litmus ChaosHub. Therefore, to access the experiment the ChaosHub, we need to add the following public GitHub repository as ChaosHub:

- **Hub Name:** m-agent
- **Git URL:** https://github.com/litmuschaos/chaos-charts 
- **Branch:** m-agent

Upon adding the ChaosHub, you will be able to access the os/process-kill experiment charts.

### STEP-4: Create a Chaos Workflow for m-agent
Create a Chaos Workflow, and choose the m-agent ChaosHub which you had previously added. Then add the os/process-kill experiment and tune the experiment by adding the target PIDs in `PROCESS_IDS` as comma-separated values. You can leave the other settings as default.

You can add any of the Litmus Probes for this experiment as well. The cmdProbe in this case will execute in the target machine so that you can validate your chaos hypothesis condition within the remote target machine.

You can now execute the chaos workflow that you have created.

### STEP-5: Observe the Experiment Outcome
As a result of this experiment, the target PIDs will be killed and the target processes will be removed from the list of currently running processes, that can be obtained using the `ps aux` command.

## Conclusion
In conclusion, we observed how to use m-agent for performing the process-kill experiment. We saw how to install and use m-agent, how the experiment charts can be added to ChaosCenter and how to tune the experiment for targetting the processes executing in a remote target machine. Feel free to raise all your queries and concerns in the [Litmus Kubernetes Slack channel](https://kubernetes.slack.com/?redir=%2Farchives%2FCNXNB0ZTN).