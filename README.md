# mongo-ping

```bash
export CGO_ENABLED=0
GOOS=linux GOARCH=amd64 go build -v -o /bin/mongo-ping .

docker build -t shaad7/mongo-ping:0.0.1 .
docker push shaad7/mongo-ping:0.0.1
```

