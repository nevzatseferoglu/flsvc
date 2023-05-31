# flsvc

CLI tool which acting as a client for https://fl-service-api-doc.netlify.app/.

```bash
This API caters to data scientists, simplifying remote host communication with service endpoints.It allows users to efficiently manage
flower federated learning clusters.

API doc: https://fl-service-api-doc.netlify.app/

Usage:
  flsvc [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  docker       Install and get the states of the docker dependencies
  help         Help about any command
  ping         Ping the remote host which has the given IPAddress
  remote-hosts Remote host operations

Flags:
  -c, --config string   config file (default is $HOME/.flsvc.yaml)
  -h, --help            help for flsvc

Use "flsvc [command] --help" for more information about a command.
```

## ToDo

- [x] Add remaning routes
- [ ] BubbleTea integration
- [ ] Refactoring, create unique package for each child of the root command (encapsulation)
