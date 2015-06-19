package payload

type User struct {
	Username    string `json:"username,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}
