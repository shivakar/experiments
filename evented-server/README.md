## Evented servers

There are two common solutions to handling concurrent requests to a server:
* *Threaded* servers use threads that each handle one client request. e.g.
  Apache HTTP server
* *Evented* servers that run a single event loop process that handles events for
  all connected clients. e.g. Nginx

Typically, the performance of *threaded* servers degrades rapidly beyond a
certain number of concurrent requests, while *evented* servers are known for
their ability to scale.

The rule of thumb is that if the workload is CPU bound, then *threaded* server
should perform better than *evented* servers. If on the other hand, the workload
is I/O bould, *evented* servers should fair better.

## Experiments

The following programs were created:

1. Client which connects to a TCP server using hostname and port, sends and
   receives messages.
2. Threaded server which echoes back the client message.
3. Evented server which echoes back the client message.

## Notes

### Client 

The client's algorithm looks as follows:

```
GET hostname and port from the commandline arguments
CONNECT to the server
WHILE not terminated by user and connection to server good:
  GET message to send as user input from console
  SEND message to the server
  RECEIVE message from the server
  DISPLAY message from the server to the user
```