# M-Agent API for Process Kill Experiment v1.0.0 documentation

M-Agent is a platform-generic, OS scoped agent for aiding with the injection and orchestration of the faults, as part of the LitmusChaos Experiments. It can also be used to inject chaos into any physical node installed with a Linux OS. 
Process Kill experiment causes target processes, identified by their PIDs, to be killed in either serial or parallel mode.

## Table of Contents

* [Servers](#servers)
* [Channels](#channels)

<a name="servers"></a>

## Servers

<table>
  <thead>
    <tr>
      <th>URL</th>
      <th>Protocol</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
  <tr>
      <td>ws://&lt;node-external-ip&gt;:41365/process-kill</td>
      <td>ws</td>
      <td>The agent endpoint is exposed at port **41365**, where the client can attempt to establish a connection, given the agent is publicly accessible. For the connection to be established, an authentication token is required which can be generated using the agent itself.
</td>
    </tr>
    <tr>
      <td colspan="3">
        <details>
          <summary>URL Variables</summary>
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Default value</th>
                <th>Possible values</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>port</td>
                <td>
                    <em>None</em>
                  </td>
                <td>
                  <ul><li>41365</li></ul>
                  </td>
                <td></td>
              </tr>
              </tbody>
          </table>
        </details>
      </td>
    </tr>
    <tr>
      <td colspan="3">
        <details>
          <summary>Security Requirements</summary>
          <table>
            <thead>
              <tr>
                <th>Type</th>
                <th>In</th>
                <th>Name</th>
                <th>Scheme</th>
                <th>Format</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody><tr>
                <td>http</td>
                <td></td>
                <td></td>
                <td>bearer</td>
                <td>JWT</td>
                <td><p>The authentication token obtained from the agent is to be put in the request header with key &quot;Authorization&quot; and value &quot;Bearer &lt;authentication-token&gt;&quot;.</p>
</td>
              </tr></tbody>
          </table>
        </details>
      </td>
    </tr>
    </tbody>
</table>




## Channels



<a name="channel-ACTION_SUCCESSFUL"></a>

Sends a message to the client to indicate that the request for an "action" to be performed has been successfully accomplished.

#### Channel Parameters




###  `publish` ACTION_SUCCESSFUL

*The message consists of a &quot;feedback&quot; of type string, with the value &quot;ACTION_SUCCESSFUL&quot; and a &quot;payload&quot; of type object (Go Interface type). The payload can therefore contain any kind of object that may suitably be sent as part of the feedback.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), which can then be sent to the server from the client or vice-versa.



##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>

<tr>
  <td>Action </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr>

<tr>
  <td>Payload </td>
  <td>object</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {}
}
```





<a name="channel-ERROR"></a>

Sends an error message to the client to indicate that the requested "action" has failed, as well as return the error string.

#### Channel Parameters




###  `publish` ERROR

*The message consists of a &quot;feedback&quot; of type string, with the value &quot;ERROR&quot; and a &quot;payload&quot; of type object (Go Interface type). In this case, the payload will contain the error message, which will be of type string.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), which can then be sent to the server from the client or vice-versa.



##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>

<tr>
  <td>Action </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr>

<tr>
  <td>Payload </td>
  <td>object</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {}
}
```





<a name="channel-CHECK_STEADY_STATE"></a>

Validates the steady state of the target Processes i.e. whether all the target Processes exist and running.

#### Channel Parameters




###  `subscribe` CHECK_STEADY_STATE

*The message consists of an &quot;action&quot; of type string, with the value &quot;CHECK_STEADY_STATE&quot; and a &quot;payload&quot; of type object (Go Interface type). In this case, the payload will consist of an integer array which will contain  the PIDs of the target processes.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), which can then be sent to the server from the client or vice-versa.



##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>

<tr>
  <td>Action </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr>

<tr>
  <td>Payload </td>
  <td>object</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {}
}
```





<a name="channel-EXECUTE_EXPERIMENT"></a>

Kills the target processes.

#### Channel Parameters




###  `subscribe` EXECUTE_EXPERIMENT

*The message consists of an &quot;action&quot; of type string, with the value &quot;EXECUTE_EXPERIMENT&quot; and a &quot;payload&quot; of type object (Go Interface type). In this case, the payload will consist of an integer array which will contain  the PIDs of the target processes.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), which can then be sent to the server from the client or vice-versa.



##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>

<tr>
  <td>Action </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr>

<tr>
  <td>Payload </td>
  <td>object</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {}
}
```





<a name="channel-EXECUTE_COMMAND"></a>

Executes a bash script command as part of the Litmus cmdProbe execution.

#### Channel Parameters




###  `subscribe` EXECUTE_COMMAND

*The message consists of an &quot;action&quot; of type string, with the value &quot;EXECUTE_EXPERIMENT&quot; and a &quot;payload&quot; of type object (Go Interface type). In this case, the payload will contain the bash script command of type string.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), which can then be sent to the server from the client or vice-versa.



##### Payload


<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
      <th>Accepted values</th>
    </tr>
  </thead>
  <tbody>

<tr>
  <td>Action </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr>

<tr>
  <td>Payload </td>
  <td>object</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {}
}
```





