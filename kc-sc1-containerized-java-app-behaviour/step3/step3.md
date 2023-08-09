# Reviewing System Configuration

Initially, it's beneficial to inspect the system's memory and CPU configuration.

To view memory details:
`free`{{exec}}

To gather CPU information:
`lscpu`{{exec}}

You should observe that our system is equipped with 2048 MB of memory and a single CPU.

# Now, Examine the JVM Configuration

`make metrics-j`{{exec}}

You'll find that while the JVM recognizes 1 CPU, it's allocated a maximum heap size of just 480 MB.