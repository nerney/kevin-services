package util

import (
	"kevin_services/models"
	"net/http"
	"strings"
)

// NormalizeString trims leading and trailing whitespace then converts the string to all uppercase.
func NormalizeString(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}

// MethodNotPermitted checks if the request method is permitted. If not permitted, it return Response with status code 405
func MethodNotPermitted(method string) *models.Response {
	if method == "GET" || method == "POST" || method == "PUT" || method == "DELETE" {
		return nil
	}
	return &models.Response{StatusCode: http.StatusMethodNotAllowed}
}
