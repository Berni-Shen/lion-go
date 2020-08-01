package request

type Service struct {
	Name         string
	Version      int
	Description  string
	Address      string
	Port         string
	GetMethod    string
	PostMethod   string
	PutMethod    string
	DeleteMethod string
}
