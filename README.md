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
Set up a collection by creating a yaml or json file (e.g. collections/example.yaml/json) with the following format:

#### Yaml example
```yaml
requests:
    GetGoogle: 
        method: GET
        url: https://google.com
    GetDiscord: 
        method: GET
        url: https://discord.com
```

#### Json example
```json
{
    "requests": {
        "GetGoogle": {
            "method": "GET",
            "url": "https://google.com"
        },
        "GetDiscord": {
            "method": "GET",
            "url": "https://discord.com"
        }
    }
}
```

The main command will run a request from the collection when provided with the following arguments:
```
-c, --collection-directory string   The directory where collections are stored (default "collections")
-f, --file-name string              The file name of the collection
-h, --help                          help for swarm
-n, --request-name string                   The name of the request to run
```

Example of the command
```
swarm  -c collections -f example.json -n GetGoogle
swarm  -c collections -f example.yaml -n GetGoogle
```

### Running a get request
```
swarm get --url https://google.com
```

