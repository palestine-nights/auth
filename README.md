[license]: ./LICENSE
[docs]: ./docs
[docker]: ./Dockerfile

# Auth

> JWT Auth service.

Create to separate logic for auth from API.

## Development

Compile source code

```sh
$> go build -o main src/*.go
```

Run server

```
$> ./main
```

## Usage

Build and deploy using [docker][docker].

See [API documentation][docs] for more information.

Example with [httppie](https://httpie.org).

```sh
$> http GET http://localhost:8080/auth
```

```json
[
    {

    }
]
```

## License
Project released under the terms of the MIT [license][license].
