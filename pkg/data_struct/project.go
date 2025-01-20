package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
}

type ListTask struct {
	tasks []Task
}

func (l *ListTask) AddTask(task Task) {
	l.tasks = append(l.tasks, task)
}

func (l *ListTask) MarkCompleted(index int) {
	l.tasks[index].completed = true
}

func (l *ListTask) EditTask(index int, task Task) {
	l.tasks[index] = task
}

func (l *ListTask) DeleteTask(index int) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}

func main() {
	lista := ListTask{}

	leer := bufio.NewReader(os.Stdin)

	for {
		var option int
		fmt.Println("Seleccione una opcion:\n,"+
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir\n")

		fmt.Println("Ingresar la Opcion: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var t Task
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.description, _ = leer.ReadString('\n')
			lista.AddTask(t)
			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scan(&index)
			lista.MarkCompleted(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Task
			fmt.Print("Ingrese el indice de la tarea que desea editar: ")
			fmt.Scan(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.name, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.description, _ = leer.ReadString('\n')
			lista.EditTask(index, t)
			fmt.Println("Tarea editada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scan(&index)
			lista.DeleteTask(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo...")
			return

		}

		fmt.Println("Lista de tareas: ")
		fmt.Println("===================================================================")
		for i, task := range lista.tasks {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, task.name, task.description, task.completed)
		}
		fmt.Println("===================================================================")
	}

}
