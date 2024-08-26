package lista

type Lista[T any] interface {

	//EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un elemento en la primer posicion de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta un elemento en la ultima posicion de la lista
	InsertarUltimo(T)

	//BorrarPrimero borra el primer elemento de la lista y lo devuelve . Si la lista esta vacia, entra en panico
	//con el mensaje: "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacia, entra en panico
	//con el mensaje: "La lista esta vacia"
	VerPrimero() T

	//VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacia, entra en panico
	//con el mensaje: "La lista esta vacia"
	VerUltimo() T

	//Largo devuelve la cantidad de elementos en la lista
	Largo() int

	//Iterar es el iterador interno, permite al usuario pasarle una funcion como parametro para controlar
	// la iteracion
	Iterar(visitar func(T) bool)

	//Iterador devuelve un iterador externo que permite iterar sobre la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve elemento actual, en el momento de la iteracion, de la lista.
	//Si todos los elementos fueron iterados, entra en panico con "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve true si hay un elemento por ver y false en caso contrario.
	HaySiguiente() bool

	//Siguiente pasa al siguiente elemento de la lista
	//Si todos los elementos fueron iterados, entra en panico con "El iterador termino de iterar"
	Siguiente()

	//Insertar inserta en la lista, entre el elemento actual y el anterior, el valor pasado por parametro
	Insertar(T)

	//Borrar devuelve el elemento actual de la lista, y lo elimina de la lista
	//Si todos los elementos fueron iterados, entra en panico con "El iterador termino de iterar"
	Borrar() T
}
