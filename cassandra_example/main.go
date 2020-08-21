package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

type Emp struct {
	id       string
	name     string
	lastname string
	age      int
}

func createEmp(e Emp) {
	fmt.Println("Creating employee")
	if err := Session.Query("insert into emps(empid, first_name, last_name, age) values(?, ?, ?, ?)",
		e.id, e.name, e.lastname, e.age).Exec(); err != nil {
		fmt.Println("error inserting emp")
		fmt.Println(err)
	}
}

func getEmps() []Emp {
	fmt.Println("getting all employees")
	var emps []Emp
	m := map[string]interface{}{}

	iter := Session.Query("select * from emps").Iter()
	for iter.MapScan(m) {
		emps = append(emps, Emp{
			id:       m["empid"].(string),
			name:     m["first_name"].(string),
			lastname: m["last_name"].(string),
			age:      m["age"].(int),
		})
		m = map[string]interface{}{}
	}
	return emps
}

func updateEmp(e Emp) {
	fmt.Printf("Updating emp with id = %s\n", e.id)
	if err := Session.Query("update emps set first_name = ?, last_name = ?, age = ? where empid = ?",
		e.name, e.lastname, e.age, e.id).Exec(); err != nil {
		fmt.Println("Error updating emp")
		fmt.Println(err)
	}
}

func deleteEmp(id string) {
	fmt.Printf("Deleting emp with id %s\n", id)
	if err := Session.Query("delete from emps where empid = ?", id).Exec(); err != nil {
		fmt.Println("error deleting emp")
		fmt.Println(err)
	}
}

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "test"
	cluster.ProtoVersion = 4
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}

func main() {
	e1 := Emp{"e1", "pepe", "keke", 30}
	e2 := Emp{"e2", "pepega", "kekejojo", 10}
	createEmp(e1)
	fmt.Println(getEmps())
	createEmp(e2)
	fmt.Println(getEmps())
	e3 := Emp{"e1", "nono", "veve", 14}
	updateEmp(e3)
	fmt.Println(getEmps())
	deleteEmp(e2.id)
	fmt.Println(getEmps())
}
