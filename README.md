# Golang JWT Example
This is a sample of JWT using golang.

## Glide
To install missing packages, simply install glide and run:
```glide up```

### Compile
To compile the project, run
```go build```

### API End Points
There are 2 API end points, the first one accepts a JWT token. You can generate another sample here [jwt.io](https://jwt.io/)
```eyJhbGciOiJUaGlzIGlzIG15IHRva2VuIiwidHlwIjoiSldUIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZXhwIjoiMjAxNy0xMS0wNCJ9.GyjQItZyp2wbzYjMzH0C_SfTr187bv1eYnEgCkcrhYA```.
The other one generates a JWT token based on a successful login.

- Static JWT Token decoding

```URL: /jwt/validate/```
```
Request:
{
	"jwt": "eyJhbGciOiJUaGlzIGlzIG15IHRva2VuIiwidHlwIjoiSldUIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZXhwIjoiMjAxNy0xMS0wNCJ9.GyjQItZyp2wbzYjMzH0C_SfTr187bv1eYnEgCkcrhYA"
}
```
```
Response:
{
	"jwtValid": true
}
```

- Login, used to generate a JWT token

```URL: /jwt/login/```
```
Request:
{
	"username": "user",
	"password": "pass"
}
```
```
Response:
{
	"jwt": 'Some JWT Token'
}
```
