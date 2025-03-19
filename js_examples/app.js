// Callback example
let callbackFunc = (arg, callback) => {
    let arr = []
    for(i=0; i<arg; i++){
        arr.push({name:`User ${i+1}`, id:i+1})
    }
    callback(arr)
}

function printArrToConsole(arr){
    console.log(arr)
}

setTimeout(()=>{
    callbackFunc(2, printArrToConsole)
}, 1_000)

console.log("This log statement is on the line below callback")

// Promises example