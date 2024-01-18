package configuration

func ApiProjectConfig() string {
	var baseUrl = "https://operation-center-service.azurewebsites.net"
	var apiPath = "/api/v1/get-all-project-user"
	var token = "?code=ILAxQhIak5OZDhqR9Mm_yqmrkGIsyGtu7Xs1BZfDRFl0AzFuU2bl-A=="
	res := baseUrl + apiPath + token
	return res
}

// https://operation-center-service.azurewebsites.net/api/v1/get-all-project-user?code=ILAxQhIak5OZDhqR9Mm_yqmrkGIsyGtu7Xs1BZfDRFl0AzFuU2bl-A==
