package models

// TemplateData holds template data
type TemplateData struct {
	IntMap    map[string]int
	StringMap map[string]string
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
