package show

func Show() string {
	s := student{name: "陈嘉", age: 19, address: "安徽合肥"}
	return s.showStudentInfo()
}
