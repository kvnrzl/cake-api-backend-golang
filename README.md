# Cake API Backend Golang Example

The model of Cake has 3 layer :

- Repository Layer
- Service Layer
- Controller Layer

### How To Run This Project

#### Run the Testing

Go to test file in every layer and run the test

```bash
$ go test
```

#### Run the Project

The steps to run it with `docker-compose`

```bash
#clone the project
$ git clone https://github.com/kvnrzl/cake-api-backend-golang.git

#move to project
$ cd cake-api-backend-golang

#build the docker
$ docker build -t cake-api-backend-golang .

#run the docker
$ docker-compose up -d

#check the docker
$ docker ps

#check the logs
$ docker logs cake-api-backend-golang

#stop the docker
$ docker-compose down
```

### API Documentation

#### Cake API

The Cake API allows you to create, retrieve, update, and delete cakes.

### HTTP Methods

The Cake API supports the following HTTP methods:

- `POST`: create a new cake
- `GET`: retrieve a list of all cakes or a specific cake by ID
- `PATCH`: update a specific cake by ID
- `DELETE`: delete a specific cake by ID

### Endpoints

The Cake API has the following endpoints:

#### `POST /cake`

Create a new cake.

**Request Body**:

The request body should contain the following JSON object:

```bash
{
"title":"string",
"description":"string",
"rating":float,
"image":"string"
}
```

**Response**:

If the request is successful, the response will be a `201 Created` status code.

#### `GET /cake`

Retrieve a list of all cakes.

**Response**:

If the request is successful, the response will be a `200 OK` status code.

#### `GET /cake/:id`

Retrieve a specific cake by ID.

**Path Parameters**:

- `id`: the ID of the cake to retrieve

**Response**:

If the request is successful, the response will be a `200 OK` status code.

If the cake with the specified ID does not exist, the response will be a `404 Not Found` status code.

#### `PATCH /cake/:id`

Update a specific cake by ID.

**Path Parameters**:

- `id`: the ID of the cake to update

**Request Body**:

The request body should contain the following JSON object:

```bash
{
"title":"string",
"description":"string",
"rating":float,
"image":"string"
}
```

**Response**:

If the request is successful, the response will be a `200 OK` status code.

If the cake with the specified ID does not exist, the response will be a `404 Not Found` status code.

#### `DELETE /cake/:id`

Delete a specific cake by ID.

**Path Parameters**:

- `id`: the ID of the cake to delete

**Response**:

If the request is successful, the response will be a `200 OK` status code.

If the cake with the specified ID does not exist, the response will be a `404 Not Found` status code.

### Tools Used:

In this project, I use some tools listed below. But you can use any simmilar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need.

- All libraries listed in [`go.mod`](https://github.com/kvnrzl/backend-engineer-test-privy/blob/main/go.mod)
- To Generate Mocks for testing needs ["github.com/vektra/mockery"](https://github.com/vektra/mockery)
