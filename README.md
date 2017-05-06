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

File requests.yaml should contain your request. For example:

```yaml
- url: /hello?test=test
  method: POST
  headers:
    Accept:
    - '*/*'
    Accept-Encoding:
    - gzip, deflate, br
    Accept-Language:
    - en-US,en;q=0.8
    Cache-Control:
    - no-cache
    Connection:
    - keep-alive
    Content-Length:
    - "39"
    Content-Type:
    - application/json
    Origin:
    - chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop
    Test:
    - test
    User-Agent:
    - Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)
      Chrome/57.0.2987.133 Safari/537.36
  body: "{\n\t\"request\":\"test\",\n\t\"hello\":\"world\"\n}"
```
