package configuration

func ApiProjectConfig() string {
	var baseUrl = "https://jeab-project.azurewebsites.net"
	var apiPath = "/api/get-project"
	var token = "?code=h8lIZJ3SmmTidunpnFLntDrq29ZAK0-QhXKMLX1exD14AzFujfVguA=="
	res := baseUrl + apiPath + token
	return res
}
