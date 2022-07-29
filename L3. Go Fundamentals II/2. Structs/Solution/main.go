package main

import "fmt"

type Classroom struct {
	id          uint16
	capacity    uint8
	subject     string
	studentList []Student
}

type Student struct {
	id   uint16
	name string
}

func main() {
	c1 := Classroom{
		id:       1,
		capacity: 200,
		subject:  "Art",
		studentList: []Student{
			{
				id:   20,
				name: "Eric",
			},
			{
				id:   30,
				name: "Sloan",
			},
		},
	}

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 100
	c2.subject = "Theater"
	c2.studentList = []Student{
		{
			id:   40,
			name: "Vince",
		},
		{
			id:   50,
			name: "Johnny",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}
