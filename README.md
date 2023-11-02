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
    PostExample: 
        method: POST
        url: https://localhost:8080
        content-type: application/json
        body: "{\"test\":true}"
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
        },
        "PostExample": {
            "method": "POST",
            "url": "https://localhost:8080",
            "contentType": "application/json",
            "body": "{\"test\":true}"
        }
    }
}
```

These arguments are used to run a request from a collection:
```
-c, --collection-directory string   The directory where collections are stored (default "collections")
-f, --file-name string              The file name of the collection
-n, --request-name string           The name of the request to run
```

These arguments are used to run a request straight from the CLI:
```
-b, --body string                   Body in string format
--content-type string               Content type of the request body
-m, --method string                 Method for request
-u, --url string                    URL for request
```

Example of the command
```
swarm -c collections -f example.json -n GetGoogle
swarm -c collections -f example.yaml -n GetGoogle
```

### Running a get request
```
swarm --method GET --url https://google.com
```
### Running a post request
```
swarm --method POST --url https://google.com --content-type application/json --body="{\"test\":true}"
```
