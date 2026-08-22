const submitTask = () => {
    taskName = document.getElementById("inputName").value
    taskDescription = document.getElementById("inputDescription").value
    console.log(taskName)
    console.log(taskDescription)
    window.location.href = '/'
}