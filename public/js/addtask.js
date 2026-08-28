const sendTaskRequest = async (name, description) => {
    try {
        const response = await fetch("http://localhost:3000/addtask", {
            method: "POST",
            body: JSON.stringify({
                name,
                description
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        })
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
    } catch (err) {
        console.error('Fetch error:', err);
    }
}

const submitTask = async () => {
    taskName = document.getElementById("inputName").value
    taskDescription = document.getElementById("inputDescription").value
    console.log(taskName)
    console.log(taskDescription)
    try {
        await sendTaskRequest(taskName,taskDescription)
        window.location.href = '/'
    } catch (err) {
        alert("Failed to add task")
    }
    
}