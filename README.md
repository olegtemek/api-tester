## Api-tester
This utility personally helps me to test my Api, for example by sending multiple requests at the same time


### Get started

1. Clone repository
```bash
git clone https://github.com/olegtemek/api-tester
```

2. Open repository folder and going to cmd in terminal
```bash
cd test-api/cmd
```

3. Start testing your api with custom flags
* Flag -u = Target url
* Flag -m = Method (allowed: "GET", "POST", "PUT", "PATCH", "DELETE")
* Flag -t = Timeout duration (sec)
* Flag -c = Requests count
* Flag -w = Workers count

```bash
go run main.go -u=https://example.com -w=2 -c=1000 -t=10
```