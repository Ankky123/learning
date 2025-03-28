
//Make an array for storing the students 
var str = ""
const students = [{ name: "Std1", subject: "Javascript" },
{ name: "Std2", subject: "Python" }
]


appendStudent({ name: "Std3", subject: "Go" }, getStudents)

//getStudents()

appendStudent({ name: "Std4", subject: "C++" }, getStudents)



function getStudents(dum) {

        students.forEach(concatStudent)
        document.getElementById("StudentList").innerHTML = str
        console.log("students have been fetched")
        console.log("the string is :" + str)
        str = ""

}

function appendStudent(newStudent, myCallBack) {
    //setting a settime
    setTimeout(function () {
        students.push(newStudent)
        console.log("student has been enrolled")
        let dum = 0
        myCallBack(dum)  // executing the call back here
    }, 2000)

}

function concatStudent(student) {
    str += `<li> name:${student.name} subject:${student.subject} </li>`
}

