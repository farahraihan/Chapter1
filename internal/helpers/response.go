package helpers

type Meta struct {
	TotalItems   int `json:"totalItems"`
	ItemsPerPage int `json:"itemsPerPage"`
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
}

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

func ResponseWithMetaFormat(code int, status string, message string, data any, meta Meta) map[string]any {
	var result = make(map[string]any)

	result["code"] = code
	result["status"] = status
	result["message"] = message
	if data != nil {
		result["data"] = data
	}
	result["meta"] = meta

	return result
}
