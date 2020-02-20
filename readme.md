# The Tail at Scale 
Presentation about ["The Tail at Scale"](https://www2.cs.duke.edu/courses/cps296.4/fall13/838-CloudPapers/dean_longtail.pdf) paper. 

## Presentation
The presentation was built with [Reveal.js](https://revealjs.com/#/). And it can be found on `/presentation`. To run the presentation you only need to open [/presentation/index.html](/presentation/index.html).

## Demo
The demo consists of two web applications. The tool used to test it and get metrics from the applications is [vegeta](https://github.com/tsenart/vegeta)

### external-service
Web application writen in go found under `/external-service`. This service emulates a service with a tail problem. P95 is under 20ms but P99 is over 120ms. This is achieve with an internal sleep. This service is meant to be replicated using docker. Specific information about it can be found in [/external-service/readme.md](/external-service/readme.md).

### client
This is a web application writen in go exposing different endpoints for each of the techniques presented (fanout, hedged requests). The `client` application depends upon the replicas of `external-service`. Specific information can be found in [/client/readme.md](/client/readme.md).