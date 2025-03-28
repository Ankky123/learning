package main

import (
	"fmt"
	"reflect"
	"strings"
)

type organisation struct {
	Name            string
	Director        director
	establishedDate string
	TotalEmployees  int
	Employees       []employee
	Departments     struct {
		HR      *deparment
		Finance *deparment
	}
}

type director struct {
	Name string
	Age  int
}

type employee struct {
	Emp_id      int
	Name        string
	JoiningDate string
	Age         int
	Department  string
	Address     *struct {
		Country string
		State   string
		City    string
		Pincode int
	}
}

type deparment struct {
	Head_Emp_id int
	Task        string
}

func main() {
	//https://www.geeksforgeeks.org/reflection-in-golang/
	//https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
	//https://www.nutanix.dev/2022/04/22/golang-the-art-of-reflection/

	ADP := organisation{
		Name: "ADP",
		Director: director{
			Age:  27,
			Name: "Rishabh",
		},
		establishedDate: "2001-05-05",
		TotalEmployees:  3,
		Employees: []employee{
			{
				Emp_id:      2,
				Name:        "Ankit",
				JoiningDate: "2022-04-17",
				Age:         28,
				Department:  "BRF",
				Address: &struct {
					Country string
					State   string
					City    string
					Pincode int
				}{
					Country: "India",
					State:   "MP",
					City:    "Sheopur",
					Pincode: 476337,
				},
			},
			{
				Emp_id:      1,
				Name:        "Parv",
				JoiningDate: "2021-03-15",
				Age:         28,
				Department:  "HR",
				Address: &struct {
					Country string
					State   string
					City    string
					Pincode int
				}{
					Country: "India",
					State:   "Gujarat",
					City:    "Ahmedabad",
					Pincode: 593011,
				},
			},
			{
				Emp_id:      3,
				Name:        "Satyaprakash",
				JoiningDate: "2022-05-15",
				Age:         26,
				Department:  "Auxilio",
				Address: &struct {
					Country string
					State   string
					City    string
					Pincode int
				}{
					Country: "India",
					State:   "UP",
					City:    "Varanasi",
					Pincode: 777555,
				},
			},
		},
		Departments: struct {
			HR      *deparment
			Finance *deparment
		}{
			HR: &deparment{
				Head_Emp_id: 1,
				Task:        "Manage employees and fulfill their duty related requirements",
			},
			Finance: &deparment{
				Head_Emp_id: 2,
				Task:        "Manage payrolls of employees and company finance",
			},
		},
	}
	tDir := director{
		Name: "Rishabh Kanabar",
		Age:  29,
	}
	// wrappperDecoder(ADP, 0)
	//trying to change the name in the ADP to Rishabh Kanabar
	fmt.Println(ADP)
	refVal := reflect.ValueOf(&ADP)
	refVal.Elem().FieldByName("Name").SetString("ADP Pvt Ltd")
	refVal.Elem().FieldByName("Director").FieldByName("Name").SetString("Rishabh K")
	refVal.Elem().FieldByName("Employees").Index(0).FieldByName("Age").SetInt(30)
	refVal.Elem().FieldByName("Employees").Index(0).FieldByName("Emp_id").SetInt(4)
	refVal.Elem().FieldByName("Director").Set(reflect.ValueOf(tDir))
	//Add an element in the array and replace the director object

	empArr := refVal.Elem().FieldByName("Employees")
	some := refVal.Elem().FieldByName("Employees").Type().Elem()
	fmt.Println("Contained type: ", some.Name())

	tEmp := employee{
		Emp_id:      5,
		Name:        "Manisha",
		JoiningDate: "2023-05-15",
		Age:         25,
		Department:  "BRF ADP",
		Address: &struct {
			Country string
			State   string
			City    string
			Pincode int
		}{
			Country: "India",
			State:   "Rajasthan",
			City:    "Bikaner",
			Pincode: 111666,
		},
	}

	empArr = reflect.Append(empArr, reflect.ValueOf(tEmp))
	refVal.Elem().FieldByName("Employees").Set(empArr)

	v1 := refVal.Elem().FieldByName("Name")
	i1 := v1.Interface()
	fmt.Println("v1: ", v1)
	fmt.Println("i1: ", i1)
	wrappperDecoder(i1, 0)
	v2 := refVal.Elem().FieldByName("Employees").Index(0)
	i2 := v2.Interface()
	fmt.Println("v2: ", v2)
	fmt.Println("i2: ", i2)
	wrappperDecoder(i2, 0)

	fmt.Println(ADP)
	fmt.Println("Reached End")
}

func wrappperDecoder(data interface{}, depth int) {

	dataV := reflect.ValueOf(data)
	dataT := reflect.TypeOf(data)
	dataK := dataT.Kind()
	fmt.Println(strings.Repeat("\t", depth), "Type is :", dataT, " Kind is :", dataK, " Value is :", data)
	decodeInterface(dataV, depth+1)
}
func decodeInterface(data reflect.Value, depth int) {

	if data.Kind() == reflect.Struct {
		for i := 0; i < data.NumField(); i++ {
			field := data.Field(i)
			fieldK := field.Kind()
			fmt.Println(strings.Repeat("\t", depth), "Field name is :", data.Type().Field(i).Name, " Type is :", field.Type().Name(), "  Kind is :", field.Kind(), "  Value is :", field)
			if fieldK == reflect.Struct {
				decodeInterface(field, depth+1)
			}
			if fieldK == reflect.Slice {
				decodeInterface(field, depth+1)
			}
			if fieldK == reflect.Ptr {
				decodeInterface(field, depth+1)
			}
		}
	} else if data.Kind() == reflect.Slice {
		for i := 0; i < data.Len(); i++ {
			ele := data.Index(i)
			fmt.Println(strings.Repeat("\t", depth), "Element Type is :", ele.Type().Name(), "  Kind is :", ele.Kind(), "  Value is: ", ele)
			decodeInterface(ele, depth+1)
		}
	} else if data.Kind() == reflect.Ptr {
		val := data.Elem()
		fmt.Println(strings.Repeat("\t", depth), "Contained Type is :", val.Type().Name(), "  Kind is :", val.Kind(), "  Value is: ", val)
		decodeInterface(val, depth+1)
	}
}
