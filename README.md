# nsq-auth
nsq http auth service

```text
 ./nsq-auth -h
Usage:

2021/12/25 17:10:56 Usage:
  nsq-auth [OPTIONS]

Application Options:
  -a, --address=  api port default :1325 (default: :1325)
  -i, --identity= identity default zhimiaox-nsq-auth (default: zhimiaox-nsq-auth)
  -u, --auth-url= auth-url (default: http://localhost:1325)
  -t, --ttl=      auth expire duration unit s, default 60 (default: 60)
  -s, --secret=   root secret allow all push and sub topic and channel
  -f, --csv=      csv secret file path

Help Options:
  -h, --help      Show this help message

```

### Example

> root secret

```shell
./nsq-auth --secret "123456"
./nsqd --auth-http-address "localhost:1325"
```

> csv file secret

./example.csv

| secret           | topic | channel | allow             |
|:-----------------|:------|:--------|:------------------|
| PrNQuOLcyAwDPJO7 | t1    |         | publish           |
| PrNQuOLcyAwDPJO7 | t1    | c1      | subscribe         |
| PrNQuOLcyAwDPJO7 | t2    | .\*     | publish subscribe |


```shell
./nsq-auth --csv "./example.csv"
./nsqd --auth-http-address "localhost:1325"
```

### Test

```text
goos: windows
goarch: amd64
pkg: github.com/nsq-auth
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkApi_Auth
BenchmarkApi_Auth-16                2074            686007 ns/op
PASS
```