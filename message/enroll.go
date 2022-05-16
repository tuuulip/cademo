package message

type Enroll struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
