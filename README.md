# circuit-breaker-go-example
### Circuit breaker implementation example in GO


In Microservices architecture, a service usually calls other services to retrieve data, and there is a chance that the upstream service may be down. If the problem is caused by transient network issues or temporal unavailability, the client service can retry the request several times to solve the issue.

However, other serious problems may occur, such as a database outage or a service not responding quickly. In such cases, too many repeated requests which doomed to fail might lead to cascading failures throughout the system.

So, here comes the rescue: the circuit breaker. It is a mechanism that allows you to protect your service from performing too many requests in a short period.

In this post, I will show you how to use the gobreaker library to implement a circuit breaker pattern in a real example.

### Concepts Overview

### Three States

The circuit breaker has three distinct states, Closed, Open, and Half-Open.

![cb_example1](https://user-images.githubusercontent.com/48101122/196518376-e4a2ccc3-ef0b-4202-9a53-ef22a5873748.jpeg)

- Closed — all the requests are allowed to pass through to the upstream service.

- Open — All the requests are not allowed to pass through to the upstream service.

- Half-Open — To determine if the upstream service recovered, the circuit breaker will allow only a small number of requests to pass through in this state.

### State Changes

![cb_example2](https://user-images.githubusercontent.com/48101122/196519044-324be6c5-e24f-480e-b4fc-942c52eb91c6.png)

- Close to Open — The circuit breaker is in the closed state; when failed requests exceed the threshold, the circuit breaker will change to the open state.
- Open to Half-Open — The circuit breaker is in the open state; when a certain timeout period passes, the circuit breaker will change to the half-open state.
- Half-Open to Open — The circuit breaker is in the half-open state; when a request to upstream service still fails, the circuit breaker will change to the open state again.
- Half-Open to Closed — The circuit breaker is in a half-open state; when a certain number of predefined requests are successful, the circuit breaker will change to the closed state.


### Request Count

The Circuit Breaker holds the number of requests and their successes/failures.

```
type Counts struct {
    Requests             uint32
    TotalSuccesses       uint32
    TotalFailures        uint32
    ConsecutiveSuccesses uint32
    ConsecutiveFailures  uint32
}
```