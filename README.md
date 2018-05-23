# Sample go applications for use with [efgh](https://github.com/evankanderson/efgh)

To use:

```shell
go get
go build
export PORT=8080  # set PORT=8080 on windows
efgh-samples
```

To send some test events:
```shell
curl -v -H "CE-EventID: 313373" \
  -H "CE-EventTime: 2018-05-21T22:14:11Z" \
  -d "{\"a_field\": \"Blaa\", \"a_number\": 4321}" \
  http://localhost:8080
```