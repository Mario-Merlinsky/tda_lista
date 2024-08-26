package lista

const ERROR_LISTA_VACIA = "La lista esta vacia"
const ERROR_ITERADOR_FINALIZO = "El iterador termino de iterar"

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero  *nodoLista[T]
	ultimo   *nodoLista[T]
	cantidad int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func crearNodoLista[T any](dato T, prox *nodoLista[T]) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	nodo.prox = prox
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.cantidad == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodoLista[T](dato, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = nodo
	}
	lista.primero = nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodoLista[T](dato, nil)
	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.prox = nodo
	}
	lista.ultimo = nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.prox
	lista.cantidad--

	if lista.EstaVacia() {
		lista.ultimo = nil
	}

	return dato
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(ERROR_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.cantidad
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (iter iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(ERROR_ITERADOR_FINALIZO)
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(ERROR_ITERADOR_FINALIZO)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.prox

}

func (iter *iterListaEnlazada[T]) Insertar(elem T) {

	nodo := crearNodoLista[T](elem, iter.lista.primero)

	if !iter.HaySiguiente() {
		iter.lista.ultimo = nodo
	}
	if iter.anterior == nil {
		iter.lista.primero = nodo
	} else {
		nodo.prox = iter.actual
		iter.anterior.prox = nodo

	}
	iter.actual = nodo
	iter.lista.cantidad++

}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(ERROR_ITERADOR_FINALIZO)
	}

	dato := iter.actual.dato

	if iter.anterior == nil {
		iter.lista.primero = iter.actual.prox

	} else {
		if iter.actual.prox == nil {
			iter.lista.ultimo = iter.anterior
		}
		iter.anterior.prox = iter.actual.prox
	}
	iter.actual = iter.actual.prox
	iter.lista.cantidad--
	return dato

}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; actual = actual.prox {
		if !visitar(actual.dato) {
			break
		}
	}
}
