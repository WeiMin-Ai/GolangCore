package model

type student struct {
	Name string
	age  int
}

func (s *student) GetAge() int {
	return s.age
}

func NewStudent(name string, age int) *student {
	return &student{Name: name, age: age}
}
