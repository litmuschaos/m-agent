# M-Agent API for CPU Stress Experiment v1.0.0 documentation

CPU Stress experiment causes the CPUs of the target machines to be stressed at a defined rate for a fixed interval of time.

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
      <td>ws://&lt;NODE_EXTERNAL_IP&gt;:&lt;PORT&gt;/cpu-stress</td>
      <td>ws</td>
      <td>The agent endpoint is exposed at a port, which by default is **41365** but can be changed, where the client can attempt to establish a connection, given the agent is publicly accessible. For the connection to be established, an authentication token is required which can be generated using m-agent itself.
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
                  <ul><li>&lt;PORT&gt;</li></ul>
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
                <td><p>The authentication token obtained from the agent is to be put in the request header with key &quot;Authorization&quot; and value &quot;Bearer &lt;AUTHENTICATION_TOKEN&gt;&quot;.</p>
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

*The message consists of a &quot;feedback&quot; of type string, with the value &quot;ACTION_SUCCESSFUL&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string which was obtained from the client message. The payload can therefore contain any kind of object that may suitably be sent as part of the feedback.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-ERROR"></a>

Sends an error message to the client to indicate that the requested "action" has failed, as well as return the error string.

#### Channel Parameters




###  `publish` ERROR

*The message consists of a &quot;feedback&quot; of type string, with the value &quot;ERROR&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string which was obtained from the client message. In this case, the payload will contain the error message, which will be of type string.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-CHECK_STEADY_STATE"></a>

Validates the steady state for the cpu-stress experiment by ensuring m-agent is working properly and streess-ng is available.

#### Channel Parameters




###  `subscribe` CHECK_STEADY_STATE

*The message consists of an &quot;action&quot; of type string, with the value &quot;CHECK_STEADY_STATE&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will consist of an integer array which will contain the PIDs of the target processes.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-EXECUTE_EXPERIMENT"></a>

Stress a defined number of CPUs with a defined load percentage for the specified Chaos Interval.

#### Channel Parameters




###  `subscribe` EXECUTE_EXPERIMENT

*The message consists of an &quot;action&quot; of type string, with the value &quot;EXECUTE_EXPERIMENT&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will consist of a structure CPUStressParams containing field variables for number of target CPUs, load percentage and Chaos Interval.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-EXECUTE_COMMAND"></a>

Executes a bash command as part of the Litmus inline cmdProbe execution.

#### Channel Parameters




###  `subscribe` EXECUTE_COMMAND

*The message consists of an &quot;action&quot; of type string, with the value &quot;EXECUTE_EXPERIMENT&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will contain the bash script command of type string.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-CHECK_LIVENESS"></a>

Validates the liveness of m-agent and the stress-ng process for the Chaos Interval.

#### Channel Parameters




###  `subscribe` CHECK_LIVENESS

*The message consists of an &quot;action&quot; of type string, with the value &quot;CHECK_LIVENESS&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will be nil.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-REVERT_CHAOS"></a>

Reverts the chaos injection by waiting on the defunct stress-ng process and checking for any erroneous exit of the process.

#### Channel Parameters




###  `subscribe` REVERT_CHAOS

*The message consists of an &quot;action&quot; of type string, with the value &quot;REVERT_CHAOS&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will be nil.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-ABORT_EXPERIMENT"></a>

Aborts CPU stress chaos during the Chaos Interval by forcefully killing the active stress-ng process or removing the defunct process if the execution is already over.

#### Channel Parameters




###  `subscribe` ABORT_EXPERIMENT

*The message consists of an &quot;action&quot; of type string, with the value &quot;ABORT_EXPERIMENT&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will be nil.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





<a name="channel-CLOSE_CONNECTION"></a>

Closes the websocket connection after echoing a CLOSE_CONNECTION message back to the client.

#### Channel Parameters




###  `subscribe` CLOSE_CONNECTION

*The message consists of an &quot;action&quot; of type string, with the value &quot;CLOSE_CONNECTION&quot; and a &quot;payload&quot; of type object (Go Interface type), along with a Request ID string. In this case, the payload will be nil.* 

#### Message


Message encapsulates an &quot;action&quot; or a &quot;feedback&quot; of type string and a payload of type object (Go Interface type), along with a Request ID string, which can then be sent to the server from the client or vice-versa. It also contains a Request ID to appropriately map the response message by m-agent for the corresponding message.



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
</tr>

<tr>
  <td>ReqID </td>
  <td>string</td>
  <td> </td>
  <td><em>Any</em></td>
</tr></tbody>
</table>


###### Example of payload _(generated)_

```json
{
  "Action": "string",
  "Payload": {},
  "ReqID": "string"
}
```





