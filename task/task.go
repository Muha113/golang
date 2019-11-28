package task

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	BirthDay  time.Time
}

type People []Person

const timeLayout = "2006-Jan-02"

func (s Person) String() string {
	var resultStr string
	resultStr += fmt.Sprint("[", s.FirstName, ", ")
	resultStr += fmt.Sprint(s.LastName, ", ")
	resultStr += fmt.Sprint(s.BirthDay.Format(timeLayout), "]")
	return resultStr
}

func (p People) String() string {
	var str string
	for _, i := range p {
		str += i.String() + "\n"
	}
	return str
}

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].BirthDay.Equal(p[j].BirthDay) {
		strI := p[i].FirstName + p[i].LastName
		strJ := p[j].FirstName + p[j].LastName
		return strI < strJ
	}
        return p[i].BirthDay.After(p[j].BirthDay)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Task() {
	a, err := time.Parse(timeLayout, "2005-Aug-15")
	if err != nil {
		fmt.Println(err.Error())
	}
	b, err := time.Parse(timeLayout, "2006-Aug-15")
	if err != nil {
		fmt.Println(err.Error())
	}
	c, err := time.Parse(timeLayout, "2007-Aug-15")
	if err != nil {
		fmt.Println(err.Error())
	}
	p := People{
		{FirstName: "Vlad", LastName: "Petrov", BirthDay: c},
		{FirstName: "Sasha", LastName: "Polyakov", BirthDay: a},
		{FirstName: "Sasha", LastName: "Polynkov", BirthDay: a},
		{FirstName: "Andrej", LastName: "Nehay", BirthDay: b},
	}
	sort.Sort(p)
	fmt.Println(p)
}
