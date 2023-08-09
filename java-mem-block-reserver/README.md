# Java Memory Block Reserver

## Overview

`java-mem-block-reserver` is a simple web service written with java v17 and based on springboot  v3.1.0 that allows you to reserve blocks of memory. This can be useful for testing how your system or other applications behave under low memory conditions.

## Usage

To start the service, simply run the application. By default, it will start a web server on port 8080.

The service exposes several HTTP endpoints:

- `POST /segments`: Reserve a block of memory. The size and unit of the block are specified in the request body as JSON. For example: `{"size": 32, "unit": "MEGABYTES"}`. The service will return a UUID that identifies the reserved block.

- `GET /segments/{id}`: Get information about a reserved block of memory. Replace `{id}` with the UUID of the block.

- `PUT /segments/{id}`: Update the size of a reserved block of memory. Replace `{id}` with the UUID of the block. The new size and unit are specified in the request body as JSON.

- `DELETE /segments/{id}`: Release a reserved block of memory. Replace `{id}` with the UUID of the block.

- `GET /segments`: List all reserved blocks of memory.

- `GET /gc`: Trigger garbage collection in the Go runtime. This can be used to release memory that is no longer in use.

- `GET /stats`: Get statistics about the Go runtime and memory usage.

## Examples

Here are some examples of how to use the service with `curl`:

Reserve a block of memory:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"size": 32, "unit": "MEGABYTES"}' http://localhost:8080/segments
```

Get information about a reserved block of memory:
```bash
curl http://localhost:8080/segments/{id}
```

Update the size of a reserved block of memory:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"size": 64, "unit": "MEGABYTES"}' http://localhost:8080/segments/{id}
```

Release a reserved block of memory:
```bash
curl -X DELETE http://localhost:8080/segments/{id}
```

List all reserved blocks of memory:
```bash
curl http://localhost:8080/segments
```

Trigger garbage collection:
```bash
curl http://localhost:8080/gc
```

Get runtime and memory usage statistics:
```bash
curl http://localhost:8080/stats
```

=> Possible units are `BYTES`, `KILOBYTES`, `MEGABYTES` and `GIGABYTES`  
=> Replace {id} with the UUID of a reserved block of memory.  



