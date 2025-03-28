console.log("hi")
let data = {
        "name": "Geeks for Geeks",
        "age": "23"
}
let options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(data)
  }

let fetchResp = fetch("http://dummy.restapiexample.com/api/v1/create", options)
fetchResp.then(res=>res.json()).then(console.log)