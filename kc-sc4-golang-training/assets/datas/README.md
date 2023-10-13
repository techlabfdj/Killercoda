# A sample Web Service based on GIN framework

## Installation

```shell
[ ! -z "$GOPATH" ] &&
( [ -d $GOPATH/src/gitlab-tech/training ] || mkdir -p $GOPATH/src/gitlab-tech/training) &&
cd $GOPATH/src/gitlab-tech/training &&
git clone https://gitlab.techlabfdj.io/techlab/training/golang.git &&
cd $GOPATH/src/gitlab-tech/training/golang/gin-samples/datas &&
go build &&
./datas
```

## Using it

Web server will start on :8080 and exposes a /datas path.

### `GET /datas?offset=x&limit=y`

* offset defaults to 0 and cannot be negative
* limit defaults to 10 and cannot be negative or exceed 20
* returns a JSON object with
  * offset and limit properties set to used value during request
  * totalCount gives the total number of records in the system
  * array of records matchind the request in the system
  * count is the number of records returned in datas array

*Example:*

* HTTP request sent `GET /datas?limit=2`
* JSON response with 200 (OK) HTTP status code

  ```json
  {
    "count": 2,
    "datas": [
      {
        "id": "b8984156-5767-48db-b6ae-833016cdbbc1",
        "created_at": "2020-02-14T18:06:41.3767325+01:00",
        "modified_at": "2020-02-14T18:06:41.3767325+01:00",
        "aaa": "bbb",
        "ccc": "ddd"
      },
      {
        "id": "4de731dc-4d80-4432-be45-f8564d07582d",
        "created_at": "2020-02-14T18:06:41.5897325+01:00",
        "modified_at": "2020-02-14T18:06:41.5897325+01:00",
        "aaa": "bbb",
        "ccc": "ddd"
      }
    ],
    "limit": 2,
    "offset": 0,
    "totalCount": 50
  }
  ```

*Curl command line example :*  
``

### `POST /datas`

* accepts application/json content
* will store (in-memory) the data and assign it an id (uuidv4)
* store is limited to 100 entries
* returns 200 with added data with its id and created_at/modified_at dates

*Example:*

* JSON data sent in request body

  ```json
  {
    "aaa": "bbb",
    "ccc": "ddd"
  }
  ```

* JSON returned response with 200 (OK) HTTP status code

  ```json
  {
    "id": "b8984156-5767-48db-b6ae-833016cdbbc1",
    "created_at": "2020-02-14T18:06:41.3767325+01:00",
    "modified_at": "2020-02-14T18:06:41.3767325+01:00",
    "aaa": "bbb",
    "ccc": "ddd"
  }
  ```

*Curl command line example :*  
`curl -d '{"aaa": "bbb","ccc": "ddd"}' -H "Content-Type: application/json" -X POST http://localhost:8080/datas`

### `GET /datas/:id`

* id must be a uuidv4
* returns 404 if id is not a valid data reference or the requested data
* returns 200 with data

*Curl command line example :*  
`curl http://localhost:8080/datas/dcfffde8-ad11-4b89-a779-8ea30772d854`

### `PUT /datas/:id`

* id must be a uuidv4
* updated data shall be sent in request body as a JSON object
* returns 404 if id is not a valid data reference or the requested data
* returns 200 with newly recorded data in the system

*Example for `POST /datas/b8984156-5767-48db-b6ae-833016cdbbc1`:*

* JSON data sent in request body

  ```json
  {
    "aaa": "bbb",
    "ccc": "eee"
  }

* JSON response returned with 200 (OK) HTTP status code

  ```json
  {
    "id": "b8984156-5767-48db-b6ae-833016cdbbc1",
    "created_at": "2020-02-14T18:06:41.3767325+01:00",
    "modified_at": "2020-02-14T18:07:42.6767325+01:00",
    "aaa": "bbb",
    "ccc": "eee"
  }
  ```

*Curl command line example :*  
`curl -d '{"aaa": "bbb","ccc": "eee"}' -H "Content-Type: application/json" -X PUT http://localhost:8080/datas/b8984156-5767-48db-b6ae-833016cdbbc1`

### `DELETE /datas/:id`

* id must be a uuidv4
* returns 404 if id is not a valid data reference or the requested data
* returns 204 if successful

*Example:*

* HTTP request sent `DELETE /datas/b8984156-5767-48db-b6ae-833016cdbbc1`
* no JSON content returned wih 204 (NO CONTENT) HTTP status code

*Curl command line example :*  
`curl -X DELETE http://localhost:8080/datas/b8984156-5767-48db-b6ae-833016cdbbc1`
