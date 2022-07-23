package message

func RenderResponse(data interface{}, status string) interface{} {
	outputData := Response{
		Status: status,
		Data:   data,
	}
	return outputData
}
