## local weighted load balancer testing
### Setup
We have three services running on different ports:

8080: mock healthy service in AZ1

8081: mock healthy service in AZ2

8082: mock healthy service in AZ3

Started an envoy to add the above three services as lb_endpoints

8080 Server 1, priority 0

8081 Server 2, priority 1

8082 Server 3, priority 2

### Test
#### Scenario 1: 0 healthy endpoints in AZ1
Start envoy:
```
envoy -c envoy-config-az1-0%healthy.yaml
```
then make 100 Http calls:
```
./make_http_calls.sh
```
The result shows all request went to AZ2, because AZ2 has priority of 2
```
vanggie@88665a3efcfc AZRoutingPoc % ./batch_call.sh
Made 100 Http calls
AZ1 OK traffic:        0
AZ2 OK traffic:      100
AZ3 OK traffic:        0
```
#### Scenario 2: 50% healthy endpoints in AZ1
Start envoy:
```
envoy -c envoy-config-az1-50%healthy.yaml
```
then make 100 Http calls:
```
./make_http_calls.sh
```
The result shows 69% requests went to AZ1
and 31% requests went to AZ2.
All requests are succeeded
This is close to description here:
https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/locality_weight
```
vanggie@88665a3efcfc AZRoutingPoc % ./make_http_calls.sh 
Made 100 Http calls
AZ1 OK traffic:       69
AZ2 OK traffic:       31
AZ3 OK traffic:        0
```

#### Scenario 3: 100% healthy endpoints in AZ1
Start envoy:
```
envoy -c envoy-config-az1-50%healthy.yaml
```
then make 100 Http calls:
```
./make_http_calls.sh
```
The result shows 100% requests went to AZ1, because AZ1 has the highest priority
```
vanggie@88665a3efcfc AZRoutingPoc % ./make_http_calls.sh
Made 100 Http calls
AZ1 OK traffic:      100
AZ2 OK traffic:        0
AZ3 OK traffic:        0
```


