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

### Hands-on Example

### Server
We have an HTTP server running on port 8080 as the upstream service in our example. To simulate the upstream service is down, we return a 500 error code to the client in the first 5 seconds on startup.

<img width="1038" alt="Captura de Tela 2022-10-20 às 01 26 45" src="https://user-images.githubusercontent.com/48101122/196856412-bd3e0ba0-c11b-4b58-8e06-fc95e778bcab.png">

### Client

We define a simple function to call the upstream service on the client side.

<img width="1405" alt="Captura de Tela 2022-10-20 às 01 28 45" src="https://user-images.githubusercontent.com/48101122/196856636-0f57446b-47ed-48bc-8da6-feddd3945de7.png">

### Main Function
In main function, we first initialize a circuit breaker with its configuration.

- Name is the name of the circuit breaker
- MaxRequests is the maximum number of requests allowed to pass through when the circuit breaker is half-open.
- Interval is the cyclic period of the closed state for the circuit breaker to clear the internal Counts.
- Timeout is the period of the open state, after which the state of the circuit breaker becomes half-open.
- ReadyToTrip is called whenever a request fails in the closed state. The circuit breaker will come into the open state if this function returns true.
- OnStateChange is called whenever the state of the CircuitBreaker changes.

Then we make 100 requests to the upstream service and observe the circuit breaker state change.

![Captura de Tela 2022-10-20 às 01 33 32](https://user-images.githubusercontent.com/48101122/196857229-fd53a5c2-6684-471a-ae06-b8effba4b2f8.png)

### Explanation
In the first second, the circuit breaker found that the upstream service consecutively failed more than three times; it switched to the open state from the closed state.

<img width="1369" alt="Captura de Tela 2022-10-20 às 01 41 34" src="https://user-images.githubusercontent.com/48101122/196858027-4a2ee814-1d87-4387-a1ff-88515444b7d3.png">

After the timeout period passes, the circuit breaker will switch to the half-open state.

The circuit breaker switches to an open state when the following request return error from the upstream service.

<img width="1385" alt="Captura de Tela 2022-10-20 às 01 42 27" src="https://user-images.githubusercontent.com/48101122/196858127-709aa131-5d5b-43ef-944a-3948fe3fccb4.png">

When upstream recovers, the circuit breaker observes this change by making three continuous successful requests. Then it switches to the closed state.

<img width="1384" alt="Captura de Tela 2022-10-20 às 01 43 18" src="https://user-images.githubusercontent.com/48101122/196858215-57594717-5a3d-4be4-8120-20b5c9dde6c9.png">

### Conclusion
Hopefully, this article has made you realize how the circuit breaker pattern is used in real-world applications.

Run this project in your machine and have a good time!!
Please, follow this project and my github.