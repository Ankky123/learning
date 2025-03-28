
console.log("In the promise file")


//Make an array for storing the students 
var str = ""
const students = [{ name: "Std1", subject: "Javascript" },
{ name: "Std2", subject: "Python" }
]



//console.log()
//showRandom("Rick")
//appendStudent({ name: "Std3", subject: "Go" }).then(getStudents).catch(()=>console.log("Some Error occured"))
//consuming code must wait for the fulfilled or rejected promise
appendStudentUtil({name:"Std3", subject:"Go"})
appendStudent({ name: "Std4", subject: "C++" }).then(getStudents).catch(()=>console.log("Some Error occured"))

function getStudents(dum) {

        students.forEach(concatStudent)
        document.getElementById("StudentList").innerHTML = str
        console.log("students have been fetched")
        console.log("the string is :" + str)
        str = ""
}

function appendStudentUtil(newStudent) {
    appendStudent(newStudent).then(getStudents).catch(()=>console.log("Some Error occured"))
}
function appendStudent(newStudent) {
    //I promise to append a student
    return new Promise(function(resolve, reject) {
        //sending a request - Producing code
        let preLength = students.length
        students.push(newStudent) //the student is added so no error
        let postLength = students.length
        let error = false
         //setting a timeline
    setTimeout(function () {
        
        if (preLength == postLength) {
            error = true
        } else {
            error = false
        }
        if (!error) {
            console.log("student has been enrolled")
            resolve()  // executing the call back here
        } else {
            console.log("Failed to append the student")
            reject()
        }
        
    }, 2000)
    })
   

}

//One liner arrow function
// var hello = name => "Hello " + name;
// console.log(hello("Ankit"))
function concatStudent(student) {
    str += `<li> name:${student.name} subject:${student.subject} </li>`
}

// function showRandom(name) {
//     setTimeout(()=>document.getElementById("testArrow").innerHTML=name, 3000)
// }
//1. defining a function -> can be done using arrow function
//2. passing a function -> can be passed and defined on the same time using arrow function
//3. calling a function -> arrow cannot be used
// function a(){
//     console.log("a")
// }
// let s = new Promise(function(r){
//     setTimeout(function(){r("Hello8");console.log("Hello2");},2000)
//     console.log("hello5")
//     if (false) {
//        throw new Error("hello7")
//     }
//    })
// let promiseObject = new Promise(async function(myResolve, myReject){
//     let x = 0

//     //check for something
//     //check for something

//     if (x==0) {
//       // look like the callback myResolve should execute here but where is its definition??
//         //from here myResolve() callback is only getting a value, it will be called later

//         myResolve("OK")
        
//         console.log("In the the explicit resolved : OK2")

//         await s.then(val => console.log(val)).catch(err => console.log(err.message))
//         console.log("Hello3")

        
         
        
//     } else {
//         myReject("Error is found") //look like the callback myResolve should execute here but where is its definition??
//     }
// });

// function ifResolved(value) {
//     console.log("In the the explicit resolved :" + value)
// }

// function ifRejected(value) {
//     console.log("In the explicit rejected" + value)
// }
// promiseObject.then(
//     //function(value){ifResolved(value)} //
//     (value)=>ifResolved(value) //defining the myResolve() here but it will also be executed here
// ).catch(
//     function(value){ifRejected(value)} //defining the myReject() here but it will also be executed here
// )

// Learnt callbacks, arrowfunctions, promises, setTimeout, 

