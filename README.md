# go-MockHttp
Http server that mocks http responses

## Usage example
Have example.yaml file in current directory

```yaml
/hello:
  POST:
  - select:
      in: query
      key: test
      value: test
    response:
      code: 200
      headers:
        Content-Type:
        - application/json
      body: | 
        {
          "test":"test"
        }
```

Start mock server

```bash
$ go-mockhttp -mocks example.yaml -out requests.yaml
```

Send POST request to url http://localhost:8080/hello?test=test.
