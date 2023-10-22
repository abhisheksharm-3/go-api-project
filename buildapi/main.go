package main

//model for course - file

type Course struct {
	CourseId    string  `json: "courseId"`
	CourseName  string  `json: "coursename"`
	CoursePrice int     `json: "price"`
	Author      *Author `json: "author"`
}

type Author struct {
	Fullname string `json: "fullname"`
	Website  string `json: "website"`
}

//fake database
var courses []Course

//helpers/middleware -file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
