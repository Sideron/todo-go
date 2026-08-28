let taskList = document.getElementById("taskList")

class Task {
    constructor(name,description) {
        this.name = name;
        this.description = description;
        this.completed = true
    }
    getHTMLtask() {
        return `<li id="${this.name}Task" class="taskCard">
            <h2>${this.name}</h2><p>${this.description}</p>
            <input type="checkbox" id="compleated" name="compleated" value="${this.completed?"yes":"no"}">
        </li>`
    }
}

let tasks = []

const renderTasks = () => {
    taskList.innerHTML = tasks.map((x) => x.getHTMLtask()).join('')
}

const getTasks = async () => {
    try {
        const response = await fetch('http://localhost:3000/tasks')

        if (!response.ok) {
            console.error('Error: ',response.status)
            throw new Error(`Error: ${response.status}`)
        }

        const tasksRetrived = await response.json()
        tasks = []
        tasks = tasksRetrived.map(t => new Task(t.name,t.description,t.completed))

        renderTasks(); 

    } catch (err) {
        console.error('Fetch error:', error);
    }
}

renderTasks()

getTasks()