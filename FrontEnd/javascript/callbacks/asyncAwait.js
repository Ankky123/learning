
async function myFunction(a, b) {

    //let value = await promise
    //await will make the function pause the execution and wait for a resolved promise before it continues
    let myPromise = new Promise(function (resolve, reject) {
        //I promise to add the two numbers
        let sum = ""
        sum = a+b
        if (sum == "") {
            //we got a error
            reject("An error occured")
        } else {
            //the promise was resolved
            resolve(sum)
        }
    })

    let somePromise = await myPromise //waiting for my promise to be resolved.
    // let val = somePromise.then((c) => c)
    return somePromise
}

var a = myFunction(1,2)
console.log(a)
a.then(console.log).catch(console.log)
console.log("At the end of line")

//async will make a function return promise


