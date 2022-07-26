package main

import "fmt"

func main() {
	//var frequency map[int]int // hashmap
	frequency := map[int]int{} // hashmap
	var array = []int{1, 2, 5, 3, 4, 2, 7, 10, 3}
	for _, v := range array {
		fmt.Println(v)
	}

	println()

	//for _, v := range array {
	//if ()
	//}

	frequency[4] = 4
	frequency[10] = -1
	fmt.Println(frequency)

	println()

	// iterare peste hashmap
	for key, val := range frequency {
		fmt.Println(key, val)
	}

	x := frequency[4]
	fmt.Println(x)

	// nil

	data := map[string]string{}
	data["ana"] = "mere"

	/*x, exists := data["pere"]

	if !exists {
		fmt.Println("nu exista")
	} else {
		fmt.Println(x)
	}*/

	if x, exist := data["ana"]; exist {
		fmt.Println("")
	} else {
		fmt.Println("valoarea este: ", x)
	}

	// delete from hashmap
	delete(frequency, 4)
	fmt.Println(frequency)

	// structs
	structs()

	// empty int
	emptyInt()
}

/// structs

type Student struct {
	name   string `json: Name`
	age    int    `json: Age`
	grades map[string]int
}

type Student2 struct {
	name   string `json: Name`
	age    int    `json: Age`
	school string `json: School`
}

type School struct {
	name     string
	students []*Student
}

func changeAge(s *Student, newAge int) {
	s.age = newAge
}

func (s *Student) setAge(newAge int) {
	s.age = newAge
}

func newStudent(n string, a int) *Student {
	//stud := Student{name: n, age: a}
	//return &stud
	return &Student{name: n, age: a, grades: map[string]int{}}
}

func (s *School) accept(visitor Visitor) {
	visitor.visitSchool(s)
}

type myInt int32

func (m myInt) addToAnother(m2 myInt) {
	m += m2
	fmt.Printf("Sum is %d\n", m)
}

func structs() {
	s1 := Student{name: "Ana", age: 12, grades: map[string]int{}}
	fmt.Printf("Studentul %s are varsta %d\n", s1.name, s1.age)

	// pointeri pt referinte
	changeAge(&s1, 16)
	fmt.Println(s1.age)

	s1.setAge(14)
	fmt.Println(s1.age)

	println()

	s2 := newStudent("Maria", 18)
	fmt.Printf("Student: %v\n", *s2)

	// afisaza si numele campurilor
	fmt.Printf("Student (cu numele campurilor): %+v\n", *s2)

	println()

	// sau pe ultima linie trb sa fie mereu virgula pe ultima linie
	s3 := Student{
		name:   "Ana 3",
		age:    111,
		grades: map[string]int{},
	}

	fmt.Println(s3)

	println()

	var x int32

	myInt(x).addToAnother(2)

	// interfaces

	students := []*Student{&s1, s2, &s3}

	school := newSchool("UPB", students)

	pisaExam := newPisaExam("pisa", []int{12, 15})
	bacExam := newBacExam("bac", 18)

	school.accept(pisaExam)
	school.accept(bacExam)

	for _, student := range school.students {
		fmt.Println("Name: ", student.name)
		for k, v := range student.grades {
			fmt.Println(k, v)
		}
	}

}

// interfaces

// asa il obligi sa implementeze
//var _ Visitor = &Student{}

func newSchool(n string, studs []*Student) *School {
	return &School{name: n, students: studs}
}

// pentru supraincarcare / suprascriere
type College struct {
	School
	name string
}

// empty interfaces

func emptyInt() {
	frecquency := map[int]interface{}{
		4:  "mere",
		10: -1,
	}

	fmt.Println(frecquency)

	println()

	frecc := map[interface{}]interface{}{
		"mere": "pere",
		1:      -1,
		12:     Student{name: "Dan", age: 15, grades: map[string]int{}},
		15:     Student2{name: "Cosmin", age: 25, school: "UPB"},
	}

	fmt.Println(frecc)

	println()

	// jsons sa zicem

	for k, v := range frecc {
		if value, exists := v.(Student); exists {
			fmt.Println(value)
		}

		fmt.Printf("%v, %T\n", k, v)
	}

	//json.Unmarshal("JSON DE PE NET", Student{})

	for _, v := range frecc {
		switch s := v.(type) {
		case Student:
			fmt.Printf("Este student 1: %s, %d\n", s.name, s.age)
		case Student2:
			fmt.Printf("Este student 2: %s %d %q\n", s.name, s.age, s.school)
		case string:
			fmt.Printf("Found string: %s\n", s)
		default:
			fmt.Printf("No match for %v\n", v)
		}
	}
}
