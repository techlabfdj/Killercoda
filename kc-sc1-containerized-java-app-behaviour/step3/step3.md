# Step 3 - Reviewing System Configuration

Initially, it's beneficial to inspect the system's memory and CPU configuration.

## Memory details

To view memory detail, you could use the command `free`{{exec}}  

The Mem metric includes several fields that give an overview of the system's memory usage, such as:  

1. Total: This is the total amount of physical RAM on your system.  

2. Used: This shows the amount of memory that has been used up or the amount of RAM that is currently being utilized by running programs and processes.  

3. Free: This is the amount of physical memory that is not currently being used by any running processes and is ready to be allocated to new processes.  

4. Shared: This displays the total amount of memory used by the temporary tmpfs file system. Tmpfs is a file system that stores files in the computer's main memory (RAM) making it faster to access compared to traditional storage methods like a hard drive.  

5. Buff/cache: This is the memory that the kernel (operating system) uses to store recently used data so that it can be accessed quickly. It is used to speed up the performance of the computer by reducing the amount of time it takes to access data from the hard drive. Think of it like a temporary storage area where the computer stores data that it might need soon, so that it doesn't have to search for it again later.  

6. Available: This shows an estimated value of how many memory resources are still open for use. This value can fluctuate as processes start and stop and memory is freed up and allocated. So, while it may not actively be used by a process at the moment, it is still available to be allocated to a process if needed.  

 => extract from [how to use the linux free command](https://www.turing.com/kb/how-to-use-the-linux-free-command)

## CPU details

To gather CPU information you could use the command `lscpu`{{exec}}

- CPU(s) represents the number of logical cores, which equals “Thread(s) per core” × “Core(s) per socket” × “Socket(s)”.   
- One socket is one physical CPU package (which occupies one socket on the motherboard)
- Each socket hosts a number of physical cores, and each core can run one or more threads

## Sum up of system informations

Using these 2 tools, you should observe that our system is equipped with: 
- approximatively 2048 MB of memory 
- a single CPU.

# Now, let's examine the JVM Configuration

We have seen details about memory and cpu configuration on our host.
Keep in mind that these details could be differents from the JVM perspective.
we could restrict memory and cpu usage 

`make metrics-j`{{exec}}

You'll find that :
- the JVM recognizes 1 CPU
- the JVM allocates a maximum heap size of 480 MB (vs 2048 MB on the host).

# Let's now play with the java app

Start by familiarizing yourself with the documentation:  [Readme](https://github.com/techlabfdj/killercoda/blob/main/java-mem-block-reserver/README.md)

Using the Java application, we'll incrementally allocate larger memory blocks until we hit the JVM heap's threshold. This will help us understand its behavior under such conditions. After each request, we'll monitor the memory usage.

 1. Allocate a memory block:`curl -X POST -H "Content-Type: application/json" -d '{"size": 32, "unit": "MEGABYTES"}' http://localhost:8080/segments`{{exec}}  
 2. Review the memory metrics:`make metrics-j`{{exec}}

Continue the loop until you hit the Heap threshold. Once you approach this limit, you should encounter a 400 HTTP response code.  

It's important to observe that the JVM doesn't terminate but remains operational. This resilience is a standard characteristic of Java applications.

Ok, now let's stop the java application ang go to next step: `make stop-j`{{exec}}
