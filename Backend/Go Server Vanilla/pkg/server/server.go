package server

import "net/http"

var indexPage = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>My Heading</h1>
    <p>My Paragraph</p>
</body>
</html>
`

var userInfo = `
{	
	"name": "testuser",
	"age": 21
}
`

type Server struct {
}

func New() *Server {
	return &Server{}
}
func (s *Server) RespHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(indexPage))
}

func (s *Server) RespUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(userInfo))
}
