package main

iimport (
	"fmt"
	"log"
	"os/exec"
)

// main object of oberver pattern is to notify the changes in the object to any other object that is connected

//This is the observer interface and has two Update() and GetID() functions.

type Observer interface {
	Update(string)
	GetID() string
}

type Instructor struct {
	Id           string
	fName        string
	isAccepting  bool
	observerList []Observer
}

func NewInstructor(name string) *Instructor {
	newId, err := exec.Command("uuidgen").Output()

	if err != nil {
		log.Fatal(err, "newID is not created")
	}

	return &Instructor{Id: string(newId), fName: name, isAccepting: false}
}

func (I *Instructor) NotifyAll() {
	for _, observer := range I.observerList {
		observer.Update("new syllabus has been posted")
		fmt.Printf("observer %s has been notified \n", observer.GetID())
	}
}

func (I *Instructor) Register(observer Observer) {
	I.observerList = append(I.observerList, observer)
}

func (I *Instructor) RegisterList(oList []Observer) {
	I.observerList = append(I.observerList, oList...)
}

func RemoveFromList(observerList []Observer, removeObserver Observer) []Observer {
	for index, observer := range observerList {
		if observer.GetID() == removeObserver.GetID() {
			// remove the observer from the slice
			return append(observerList[:index], observerList[index+1:]...)
		}
	}

	return observerList
}

func (I *Instructor) Unregister(o Observer) []Observer {
	I.observerList = RemoveFromList(I.observerList, o)
	return I.observerList
}

type Student struct {
	studentId string
	fName     string
	lName     string
}

func (s *Student) Update(update string) {
	fmt.Printf("%s : %s  your Instructor has posted information %s \n", s.studentId, s.fName, update)
}

func (s *Student) GetID() string {
	return s.studentId
}

type StudentManage interface {
	Register(Observer Observer)
	RegisterList(Observer []Observer)
	Unregister(Observer Observer) []Observer
	NotifyAll()
}

func main() {
	// create a instructor
	adam := NewInstructor("adam bate")

	// create a two student structure
	student1 := &Student{studentId: "2006xxxx", fName: "John", lName: "Doe"}
	student2 := &Student{studentId: "2006xxxx", fName: "Van", lName: "Rockey"}

	adam.RegisterList([]Observer{student1, student2})
	adam.NotifyAll()
}
