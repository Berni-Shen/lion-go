package entities

type Service struct {
	BaseModel
	Name        string
	Version     int
	Description string
}

type Instance struct {
	BaseModel
	ServiceID string
	Address   string
	Port      string
}

type Resouce struct {
	BaseModel
	ServiceID    string
	GetMethod    string
	PostMethod   string
	PutMethod    string
	DeleteMethod string
}
