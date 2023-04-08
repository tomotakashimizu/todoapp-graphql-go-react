package infrastructure

import "time"

type Todo struct {
	ID        string `xorm:"'id' pk"`
	Text      string
	Done      bool
	UserID    string    `xorm:"'user_id'"`
	CreatedAt time.Time `xorm:"created TIMESTAMPZ"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMPZ"`
}

func (t Todo) TableName() string {
	return "todos"
}
