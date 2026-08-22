let taskList = document.getElementById("taskList")

class Task {
    constructor(name,description) {
        this.name = name;
        this.description = description;
        this.checked = true
    }
    getHTMLtask() {
        return `<li class="taskCard">
            <h2>${this.name}</h2><p>${this.description}</p>
            <input type="checkbox" id="compleated" name="compleated" value="${this.checked?"yes":"no"}">
        </li>`
    }
}

let tasks = [new Task("Tarea1", "Esta es la primera tarea"),
    new Task("Tarea2", "Esta es la segunda tarea"),
    new Task("Tarea3", "Esta es la tercera tarea")
]

taskList.innerHTML = tasks.map((x) => x.getHTMLtask()).join('')