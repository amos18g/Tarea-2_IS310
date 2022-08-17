package ColaPrioridad

import (
	"fmt"
	"testing"
)

func TestInsertarYMind(t *testing.T) {
	cp1, err := NewColaPrioridad(5)
	if err != nil {
		fmt.Println(err)
	}

	e0, err := cp1.ElementoMin()
	//tendra un error al no contener tareas
	if err == nil {
		t.Error("Esperado nill, obtenido: ", err)
		fmt.Println(e0)
	}
	fmt.Println(err)

	t1 := NewTarea("Limpiar equipos", 3)
	err = cp1.InsertarEnPrioridad(t1)

	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	t2 := NewTarea("Instalar antivirus", 0)
	err = cp1.InsertarEnPrioridad(t2)
	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	t3 := NewTarea("Instalar antivirus", 0)
	err = cp1.InsertarEnPrioridad(t3)
	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	e1, err := cp1.ElementoMin()

	if err == nil {
		fmt.Println(err)
	}

	//sera igual  a t2
	if e1 != t2 {
		t.Error("Esperado tarea", t2, ", obtenido: ", e1)
	}

}

func TestNewColaPrioridad(t *testing.T) {

	cp1, err := NewColaPrioridad(-1)

	//error sera distinto de nil
	if err == nil {
		t.Error("Esperado nill, obtenido: ", err)
		fmt.Println(nil, "  ", cp1)
	}
}

func TestQuitarMin(t *testing.T) {
	cp1, err := NewColaPrioridad(5)
	if err != nil {
		fmt.Println(err)
	}

	e0, err := cp1.QuitarMin()
	//tendra un error al no contener tareas
	if err == nil {
		t.Error("Esperado nill, obtenido: ", err)
		fmt.Println(e0)
	}
	if err != nil {
		fmt.Println(err)
	}

	t1 := NewTarea("Limpiar equipos", 5)
	err = cp1.InsertarEnPrioridad(t1)

	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	t2 := NewTarea("Instalar antivirus", 2)
	err = cp1.InsertarEnPrioridad(t2)
	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	t3 := NewTarea("Instalar antivirus", 0)
	err = cp1.InsertarEnPrioridad(t3)
	//tendra como valor nil
	if err != nil {
		t.Error("Esperado nill, obtenido: ", err)
	}

	e1, err := cp1.QuitarMin()

	if err != nil {
		fmt.Println(err)
	}

	e2, err := cp1.QuitarMin()

	if err != nil {
		fmt.Println(err)
	}

	//sera igual  a t2
	if e2 != t2 {
		t.Error("Esperado tarea", t2, ", obtenido: ", e1)
	}

	e3, err := cp1.QuitarMin()

	if err != nil {
		fmt.Println(err)
	}

	//sera igual  a t1
	if t1 != e3 {
		t.Error("Esperado tarea", t1, ", obtenido: ", e3)
	}

}

func TestColaPVacia(t *testing.T) {

	cp1, err := NewColaPrioridad(5)
	if err != nil {
		fmt.Println(err)
	}

	//tendra como valor true
	if !cp1.ColaPrioridadVacia() {
		t.Error("Esperado valor", !cp1.ColaPrioridadVacia(), ", obtenido: ", cp1.ColaPrioridadVacia())
	}

	t1 := NewTarea("Limpiar equipos", 5)
	err = cp1.InsertarEnPrioridad(t1)
	if err != nil {
		fmt.Println(nil)
	}

	//tendra como valor false
	if cp1.ColaPrioridadVacia() {
		t.Error("Esperado valor", cp1.ColaPrioridadVacia(), ", obtenido: ", !cp1.ColaPrioridadVacia())
	}

	cp1.QuitarMin()

	//tendra como valor true
	if !cp1.ColaPrioridadVacia() {
		t.Error("Esperado valor", !cp1.ColaPrioridadVacia(), ", obtenido: ", cp1.ColaPrioridadVacia())
	}

}
