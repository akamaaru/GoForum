This is a server application on Go allowing users to register/login, create and comment posts on forum.

## Stack
- Client and sever communication: `net/http`
- Router: `gorilla/mux`
- Authentication: `jwt/v5`
- Password hashing: `bcrypt`
- Environment variables: `godotenv`
- Database: `go-sql-driver/mysql`
- Logger: `log/slog`
- Validation: `validator/v10`
- Test: `testing`, `testify`, `net/http/httptest`

## Requests
### User service
- Registration
```
POST /api/v1/register 
{
    "first_name": "user_name",
    "last_name": "user_last_name",
    "email": "asdf@mail.com",
    "password": "asdf"
}
```
- Login
```
POST /api/v1/login 
{
    "email": "asdf@mail.com",
    "password": "asdf"
}
```
### Posts service
- Get feed
```
GET /api/v1/feed
```
- Get post
```
GET /api/v1/post/{id}
```
- Create post
```
POST /api/v1/post/
Authorization = <JWT token>
{
	"title": "my first post!",
	"text": "hello. this is my first post"
}
```
### Comments service
- Get comments of a post
```
GET /api/v1/post/{id}/comments
```
- Leave comment
```
POST /api/v1/post/{id}/comments
Authorization = <JWT token>
{
	"text": "my first comment!"
}
```
