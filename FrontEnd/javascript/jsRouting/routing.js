const appDiv = "app"
// let obj1 = {key1: "Data1", key2:"Data2"}
// let obj2 = {"key3":"Data3", "key4":"Data4"}
// console.log(obj1)
// console.log(obj2)
// console.log(obj1["key1"])
// console.log(obj2["key3"])

let routes = {}
let templates = {}

//Register the templates
template("template1", templateFunc1)
template("template-view1", templateView1)
template("template-view2", templateView2)
console.log("templates map/object is :" + templates)
console.log(templates)

//define the mappings route->template
route("/", "template1")
route("/view1", "template-view1")
route("/view2", "template-view2")
console.log("routes map/object is :" + routes)
console.log(routes)

//For the first load or when the routes are changed in the browser url box
window.addEventListener("load", router)
window.addEventListener("hashchange",   router)

//Register a template
function template(name, templateFunction) {
    return templates[name] = templateFunction
}

//Register the routes
function route (path, template) {
    if (typeof template == "function") {
        return routes[path] = template 
    } else if (typeof template == "string") {
        return routes[path] = templates[template]
    } else {
        return;
    }
}

function templateFunc1() {
    let myDiv = document.getElementById(appDiv)
    myDiv.innerHTML = "Home"
    const link1 = createLink("view1", "Go to view1", "#/view1")
    const link2 = createLink("view2", "Go to view2", "#/view2")

    myDiv.appendChild(document.createElement("BR"))
    myDiv.appendChild(link1)
    myDiv.appendChild(document.createElement("BR"))
    myDiv.appendChild(link2)
    return myDiv
}

function templateView1() {
    let myDiv = document.getElementById(appDiv)
    myDiv.innerHTML =""
    const link1 = createDiv("view1", "<div><h1>This is view 1</h1><a href='#/'>Go back to Index</a></div>")
    return myDiv.appendChild(link1)

}

function templateView2() {
    let myDiv = document.getElementById(appDiv)
    myDiv.innerHTML =""
    const link2 = createDiv("view2", "<div><h1>This is view 2</h1><a href='#/'>Go back to Index</a></div>")
    return myDiv.appendChild(link2)
}

function createLink(title, text, href) {
    let a = document.createElement("a")
    let linkText = document.createTextNode(text)
    a.appendChild(linkText)
    a.title = title
    a.href = href
    return a 
}

function createDiv(id, xmlString) {
    let d = document.createElement("div")
    d.id = id
    d.innerHTML = xmlString 
    return d.firstChild 
}

//Give the corresponding route of fail
function resolveRoute(route) {
    try {
        return routes[route]
    } catch (error) {
        throw new Error("The route is not defined")
    }
}

//The actual router, get the current URL and generate the corresponding template
function router(evt) {
    const url = window.location.hash.slice(1) || "/"
    const routeResolved = resolveRoute(url)
    routeResolved()
}



  





