package message

type Enroll struct {
	User         string `json:"user"`
	Organization string `json:"organization"`
	Type         string `json:"type"`
}
