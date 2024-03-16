package main

var students = []*Student{}

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade int32  `json:"gradee"`
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}
