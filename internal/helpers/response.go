package helpers

func ResponseFormat(code int, status string, message string, data any) map[string]any {
	var result = make(map[string]any)

	result["code"] = code
	result["status"] = status
	result["message"] = message
	if data != nil {
		result["data"] = data
	}

	return result
}
