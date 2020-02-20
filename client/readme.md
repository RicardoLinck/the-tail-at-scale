# Client
Web application that fires requests to the replicas of `external-service`.

Endpoints:
 - `/simple` - Simple call to the first replica url. Should have the same performance as each of the replicas.
 - `/fanout` - Fanout 3 requests (one to each replica) for each request received. Should have the best performance P99 around 20ms. Firing 3x more requests.
 - `/hedged` - Fires a hedged request after the first one outstanding the expected P95 (21ms). Should increase tail performance to P99 around 40ms. Firing up to 5% more requests only.

## Setup

```bash
go build

# running and saving output to file for later inspection
./client &> simple.log
```

## Testing and getting metrics
```bash
echo "GET http://localhost:8080/simple" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
echo "GET http://localhost:8080/hedged" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
echo "GET http://localhost:8080/fanout" | vegeta attack -duration=5s -rate 100/1s -workers 10 -connections 10 --insecure | tee results.bin | vegeta report
```