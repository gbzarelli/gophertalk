# GopherTalk

![](./images/gophertalk.png)

GopherTalk is a multi-user chat powered by GO to explore its standard 
library and features like sockets, goroutines, channels 
and sync package.

---

## The Project

The project was two binaries, the [server](./cmd/server.go) located in {project-dir}/cmd/server.go and the [client](./cmd/client.go) in {project-dir}/cmd/client.go

![Alt Text](./images/show.gif)

### Server

The server is responsible to maintain the channels with your clients open, and, redirect the messages to them.

The traffic between client-server-client is done by payloads in JSON formats established in the package internal / dto

Was use pure Socket channel to accept new clients. No specific protocol was used for communication.

### Client

The client just connect in the server (informed when start) and login with an uniq username, there is no authentication. 
If the username was exists in the server, its reject and asked to reconnect.

The client can talk with all people in the server or a specific user defined by command

#### Commands:

- `/help`         : to show the help message
- `/users`        : for list connected users
- `/to {user}`    : to define the user to send the message
- `/all`          : to define all people to send the message

# Run

Build the project

```shell
>make build
```

Run the server at port 8080:

```shell
>make run-server
```

Run client:

```shell
>make run-client
```

Enjoy!

# Contribute

Pull Requests are welcome. For important changes, open an 'issue' first to discuss what you would like to change. Be sure to update tests as appropriate.

# Developer

Guilherme Biff Zarelli

- Blog/Site - https://helpdev.com.br
- LinkedIn - https://linkedin.com/in/gbzarelli/
- GitHub - https://github.com/gbzarelli
- Medium - https://medium.com/@guilherme.zarelli
- Email - gbzarelli@helpdev.com.br
