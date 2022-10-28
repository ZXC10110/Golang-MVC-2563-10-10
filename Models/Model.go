package Models

func (b *Student) TableName() string {
	return "student"
}

type Student struct {
	StudentId int    `json:"student_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Year      int    `json:"year"`
}

type Count struct {
	Count int `json:"count"`
}

type Sport struct {
	SportNo int `json:"sport_no"`
	Student Student
}
