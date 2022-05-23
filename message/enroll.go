package message

type Enroll struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Type         string `json:"type"`
	Organization string `json:"organization"`
	Domain       string `json:"domain"`
}
