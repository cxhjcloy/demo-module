package show

import "fmt"

type student struct {
	name    string
	age     int
	address string
}

func (s *student) showStudentInfo() (info string) {
	info = fmt.Sprintf("the student name is %s,age is %d address is %s", s.name, s.age, s.address)
	return
}
