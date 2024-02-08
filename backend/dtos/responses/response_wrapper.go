package responses

import "github.com/labstack/echo/v4"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Use for general response
func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

// Use when we want to return message in a success response
func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

// Use when we want to return error message in a failed response
func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
