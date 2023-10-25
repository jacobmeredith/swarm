# Swarm
A simple command line utility for grouping and calling HTTP requests

You can run requests directly from the command line or you can set up a collection and run a request from the collection file.

## Quick start
First install dependencies and then build the executable.
```
make install_dependencies
make build_cli
```

After doing these steps you can run the executable
```
./bin/swarm
```

## Usage
### Running a request from collection
Set up a collection by creating a yaml file (e.g. collections/example.yaml) with the following format:

```yaml
requests:
    GetGoogle: 
        method: GET
        url: https://google.com
    GetDiscord: 
        method: GET
        url: https://discord.com
```

The main command will run a request from the collection when provided with the following arguments:
```
-c, --collection-directory string   The directory where collections are stored (default "collections")
-f, --file-name string              The file name of the collection
-h, --help                          help for swarm
-n, --name string                   The name of the request to run
```

Example of the command
```
swarm  -c collections -f example -n GetGoogle
```

### Running a get request
```
swarm get --url htts://google.com
```

