package shared

// SharedService provides business logic shared across all platforms
// This can include common utilities, data access, etc.
type SharedService struct{}

func (s *SharedService) GetAppInfo() map[string]any {
	return map[string]any{
		"name":    "multi-platform-app",
		"version": "1.0.0",
		"author":  "awesome-goose",
	}
}

func (s *SharedService) FormatResponse(data any) map[string]any {
	return map[string]any{
		"success": true,
		"data":    data,
	}
}
