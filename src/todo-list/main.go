package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *ListaTareas) marcarCompletado(index int) {
	lt.tareas[index].completado = true
}

func (lt *ListaTareas) editarTarea(index int, t Tarea) {
	lt.tareas[index] = t
}

// **
//
// */
func (lt *ListaTareas) eliminarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
}

func main() {
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Mostrar lista de tareas\n",
			"6. Salir\n")

		fmt.Println("Ingrese su opción:")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada con éxito.")
		case 2:
			var index int
			fmt.Println("Ingrese el índice de la tarea a marcar como completada:")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada.")

		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el índice de la tarea a editar:")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea editada con éxito.")

		case 4:
			var index int
			fmt.Println("Ingrese el índice de la tarea a eliminar:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada con éxito.")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}

		fmt.Println("Lista de Tareas")
		fmt.Println("===============")
		for i, t := range lista.tareas {
			fmt.Println("%d. %s - %s (Completado: %t)\n", i, t.nombre, t.descripcion, t.completado)
		}
		fmt.Println("===============")
	}
}
