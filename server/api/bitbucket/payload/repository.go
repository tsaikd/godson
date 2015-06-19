package payload

type Repository struct {
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}
