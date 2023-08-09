# Reviewing System Configuration

Initially, it's beneficial to inspect the system's memory and CPU configuration.

To view memory details:
`free`{{exec}}

To gather CPU information:
`lscpu`{{exec}}

You should observe that our system is equipped with approximatively 2048 MB of memory and a single CPU.

# Now, Examine the JVM Configuration

`make metrics-j`{{exec}}

You'll find that while the JVM recognizes 1 CPU, it's allocated a maximum heap size of just 480 MB.

# Let's now play with the java app

Start by familiarizing yourself with the documentation:  [Readme](https://github.com/techlabfdj/killercoda/blob/main/java-mem-block-reserver/README.md)

Using the Java application, we'll incrementally allocate larger memory blocks until we hit the JVM heap's threshold. This will help us understand its behavior under such conditions. After each request, we'll monitor the memory usage.

 1. Allocate a memory block:`curl -X POST -H "Content-Type: application/json" -d '{"size": 32, "unit": "MEGABYTES"}' http://localhost:8080/segments`{{exec}}  
 2. Review the memory metrics:`make metrics-j`{{exec}}

Continue the loop until you hit the Heap threshold. Once you approach this limit, you should encounter a 400 HTTP response code.  

It's important to observe that the JVM doesn't terminate but remains operational. This resilience is a standard characteristic of Java applications.

Ok, now let's stop the java application ang go to next step: `make stop-j`{{exec}}
