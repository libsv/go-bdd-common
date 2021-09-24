# Protocol Documentation

To generate the go files locally on your machine, run `./proto_go.sh`, or `make proto`. See more detail inside `proto_go.sh` file documentation.

To securely generate the go files using docker, run

```terminal
docker container run --rm --name proto_gen --mount type=bind,source=/your/local/path/to/platformservices/proto,target=/development/ps/proto -it chainkins/ps_gobuilder:1.16.6 /bin/bash -c "cd /development/ps/proto; ./proto_go.sh"
```

To generate the documentation for these protobuf definitions, run `make gen` command in this project.  This will output documentation to the docs folder
