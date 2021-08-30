# Flaky API: Dealing with errors when fetching data

We have a simple pagianted API hosted at [http://app-homevision-staging.herokuapp.com/api_project/houses?page=1](http://app-homevision-staging.herokuapp.com/api_project/houses?page=1) that returns a list of houses along with some metadata. The task will be to write a script that accomplishes the following tasks:

1. Requests the first 10 pages of results from the API
2. Parses the JSON returned by the API
3. Downloads the photo for each house and saves it in a file with the format: `id-[NNN]-[address].[ext]`

There are a few gotchas to watch out for:

1. This is a *flaky* API! That means that it will likely fail with a non-200 response code. Your code *must* handle these errors correctly so that all photos are downloaded
2. Downloading photos is slow so please think a bit about how you would optimize your downloads, making use of concurrency

## Implementation using golang

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

I choose Golang cuz it allows us to write concurrent systems without effort

### Requirement

- Golang 16.6

### Run

    $ go build flaky.go
    $ ./build/flaky

#### or 

    $ ./build/flaky

#### or

    $ go run flaky.go

## Run test cases

    $ go test -v  ./application

Note: if you wanna see the time duration of the program just use 
the `time` before the command, example : `$ time ./build/flaky`

## Architecture

I'm using [Clean architecture](https://www.amazon.com/-/es/Robert-Martin/dp/0134494164) so we can scale our solution without fear :)

![alt text](./docs/clean-architecture.png)

- By: Samuel Bonilla
- Email: Pythonners@gmail.com


