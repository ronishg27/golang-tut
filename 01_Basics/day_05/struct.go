package day_05

import "fmt"

type College struct {
	Name              string
	Address           string
	Students          []string
	TotalNoOfStudents int
}

// making struct methods
func (clg College) DisplayCollegeInfo() {
	fmt.Println("Name: ", clg.Name)
	fmt.Println("Address: ", clg.Address)
	fmt.Println("Total No of Students: ", clg.TotalNoOfStudents)
	fmt.Println("Students: ", clg.Students)
}

func CheckStruct() {

	samriddhi := College{
		Name:              "Samriddhi College",
		Address:           "Lokanthali",
		Students:          []string{"Ram", "Shyam", "Hari"},
		TotalNoOfStudents: 3,
	}

	// DisplayCollegeInfo(samriddhi)
	samriddhi.DisplayCollegeInfo()
}
