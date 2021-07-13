package repo

//TaskModel ...
type TaskModel struct {
	ID      uint64 `gorm:"column:id" json:"id,omitempty"`
	Name    string `gorm:"column:name" json:"name,omitempty"`
	StartAt string `gorm:"column:startAt" json:"startAt,omitempty"`
	EndAt   string `gorm:"column:endAt" json:"endAt,omitempty"`
	Status  uint8  `gorm:"column:status" json:"status,omitempty"`
}

//TableName ...
func (t *TaskModel) TableName() string {
	return "tasks"
}
