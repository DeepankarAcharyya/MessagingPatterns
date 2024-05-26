# Producer Consumer Messaging Patterns
Repo : Exploring and implementing different messaging patterns in Go.

* Messaging patterns help connect services in a loosely coupled manner.
* Very beneficial in a micro-service setup.
* One microservice can generate a message ( data to be shared ) => pass it to the broker and another microservice can pick it up.
* The decoupling makes the services more robust, reliable, easily scalable(independently) and fault tolerant.

### Components :
##### Producer 
- The process or service generating the data.
##### Consumer
- The process or service responsible to receive the data and process it.
##### Shared Data Store
- The common bridge, decoupling the Producer and Consumer. It can be a message queue or data store where producers write to and consumers read from it without data loss. Emphasis on `WITHOUT DATA LOSS`.

### Types :
##### 1. Competing Consumers Pattern
Multiple consumer processes are ready to process the incoming data and they compete with each other to process each message. With multiple concurrent consumers, a system can process multiple messages concurrently to optimize throughput, to improve scalability and availability and distribute the workload among the available consumer processes. The pattern is scalable by default.

##### 2. Priority Queue Pattern
The pattern is designed to address scenarios where different SLAs are provided to the consumers based on priority/subscription plans (etc).
Either use a priority queue that reorganizes to put the high priority data at the very front or maintain separate queues for different priorities.

![](https://miro.medium.com/v2/resize:fit:640/format:webp/1*sDCgUnfCAPLpDS6M-pZFSQ.png)

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*cSWsS86Q0-IVV7ORQhdiTw.png)

##### 3. Queue-Based Load Leveling Pattern
This patterns works on reducing the impact of spikes in the influx of data to be processed. The intermediary shared data store acts as a buffer storing the data till the consumers are ready to process it. In this pattern the consumers processes the data in a (more or less) uniform rate, so it can be said that this pattern works on reducing the load on the consumer processes.
This pattern is not exactly ideal for services that needs to process data with minimal latency.


References : 
* [Cloud Design Patterns: 3 messaging patterns](https://medium.com/@dmosyan/cloud-design-patterns-3-messaging-patterns-15e1e01da92)
