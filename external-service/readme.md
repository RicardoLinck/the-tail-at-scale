# External service

HTTP web server with a `/ishealthy` endpoint with internal sleep.

## Setup

```bash
docker build -t external-service:lt .
```

```bash
docker run -d --name external-service-1 -p 8081:8080 external-service:lt
docker run -d --name external-service-2 -p 8082:8080 external-service:lt
docker run -d --name external-service-3 -p 8083:8080 external-service:lt
```

## Testing and getting metrics
```bash
echo "GET http://localhost:8081/ishealthy" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
echo "GET http://localhost:8082/ishealthy" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
echo "GET http://localhost:8083/ishealthy" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
```

## Tear-down

```bash
docker rm -f external-service-1
docker rm -f external-service-2 
docker rm -f external-service-3
```

