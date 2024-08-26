package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaniciada(t *testing.T) {
	t.Log("Inicializamos una lista y probamos funcionamiento basico inicial")
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.NotNil(t, lista)

}
func TestInsertaPrimero(t *testing.T) {
	t.Log("Se insertan pocos valores al principio de la lista y se analiza si salen en el orden correcto")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

}

func TestInsertarUltimo(t *testing.T) {
	t.Log("Se insertan pocos valores al final de la lista y se analiza si salen en el orden correcto")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())

	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())

}

func TestInsertarYQuitarVarios(t *testing.T) {
	t.Log("Vamos insertando y quitando varios elemento intercaladamente y vemos si salen en el mismo orden y el largo de la lista es correcto")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 16; i++ {
		if i%3 == 0 {
			lista.BorrarPrimero()
		} else {
			lista.InsertarPrimero(i)
		}
	}
	require.EqualValues(t, 5, lista.Largo())
	for i := 13; i != 1; i -= 3 {
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, i-3, lista.VerPrimero())
		require.EqualValues(t, 1, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
	}
}

func TestVerPrimeroYULtimo(t *testing.T) {
	t.Log("Se prueba que el primero y el ultimo se actualizan correctamente hasta tener la lista vacia")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("hola")
	lista.InsertarPrimero("algoritmos")
	lista.InsertarUltimo("1234")
	require.EqualValues(t, "algoritmos", lista.VerPrimero())
	require.EqualValues(t, "1234", lista.VerUltimo())
	lista.BorrarPrimero()
	require.EqualValues(t, "hola", lista.VerPrimero())
	require.EqualValues(t, "1234", lista.VerUltimo())
	lista.BorrarPrimero()
	require.EqualValues(t, "1234", lista.VerPrimero())
	require.EqualValues(t, "1234", lista.VerUltimo())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

}

func TestVolumenInsertandoUltimo(t *testing.T) {
	t.Log("Se aniaden muchos elementos a la lista insertandolos al final, luego se los van eliminando viendo que se mantenga el orden y el largo de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 1000000; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, 0, lista.VerPrimero())
		require.False(t, lista.EstaVacia())
	}
	require.EqualValues(t, 1000000, lista.Largo())
	for i := 0; !lista.EstaVacia(); i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, 999999-i, lista.Largo())
		//Si la lista no esta vacia vemos el nuevo tope, condicion borde del ciclo
		if !lista.EstaVacia() {
			require.EqualValues(t, i+1, lista.VerPrimero())
		}
	}
}

func TestVolumenInsertandoPrimero(t *testing.T) {
	t.Log("Se aniaden muchos elementos a la lista insertandolos al principio, luego se los van eliminando viedno que se mantenga ek orden y el largo de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 1000000; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
	}
	require.EqualValues(t, 1000000, lista.Largo())
	for i := 999999; !lista.EstaVacia(); i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, i, lista.Largo())
		//Si la lista no esta vacia vemos el nuevo tope, condicion borde del ciclo
		if !lista.EstaVacia() {
			require.EqualValues(t, i-1, lista.VerPrimero())
		}
	}
}

func TestOperacionesInvalidasListaVacia(t *testing.T) {
	t.Log("Se testean las operaciones que deberian dar Panic en una lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestIteradorInternoTodaLaLista(t *testing.T) {
	t.Log("Iteramos toda la lista y almacenamos la suma de cada dato en una variable")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 11; i++ {
		lista.InsertarUltimo(i)
	}
	suma_elementos := 0
	lista.Iterar(func(dato int) bool {
		suma_elementos += dato
		return true
	})
	require.EqualValues(t, 55, suma_elementos)
}

func TestIteradorInternoPrimerElemento(t *testing.T) {
	t.Log("Iteramos solo el primer elemento de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 5; i > 0; i-- {
		lista.InsertarUltimo(i)
	}
	multiplicacion := 5
	lista.Iterar(func(dato int) bool {
		multiplicacion *= dato
		return false
	})
	require.EqualValues(t, 25, multiplicacion)
}

func TestIterarAlgunoElementos(t *testing.T) {
	t.Log("Iteramos algunos elementos y cortamos la iteracion por alguna condicion")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(10)
	numero_impar := -1
	lista.Iterar(func(dato int) bool {
		if dato%2 != 0 {
			numero_impar = dato
			return false
		}
		return true
	})
	require.EqualValues(t, 5, numero_impar)
}

func TestIterarListaVacia(t *testing.T) {
	t.Log("Probamos iterar una lista vacia y que no falle al intentarlo")
	lista := TDALista.CrearListaEnlazada[string]()
	cadena := "Se itero sin fallas"
	lista.Iterar(func(dato string) bool {
		cadena = dato
		return true
	})
	require.EqualValues(t, "Se itero sin fallas", cadena)
}

func TestIterarListaLlena(t *testing.T) {
	t.Log("Probamos iterar una lista con elementos y que no falle al intentarlo")
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < 11; i++ {
		lista.InsertarUltimo(i)
	}
	valor := 0
	lista.Iterar(func(v int) bool {
		if v == valor {
			valor++
			return true
		}
		return false
	})
	require.EqualValues(t, 11, valor)
}

func TestIteradorExtInsertarInicioyFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)

	iter := lista.Iterador()

	iter.Insertar(1)

	require.EqualValues(t, 1, lista.VerPrimero())

	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(4)

	require.EqualValues(t, 4, lista.VerUltimo())

}

func TestIteradorExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()

	iter.Siguiente()

	iter.Insertar(2)

	require.EqualValues(t, 2, iter.VerActual())

	iter.Siguiente()

	require.EqualValues(t, 3, iter.VerActual())

	iter.Siguiente()
	iter.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

}

func TestIteradorExtBorrarInicioFinal(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()

	borrado := iter.Borrar()

	require.EqualValues(t, 1, borrado)
	require.EqualValues(t, 2, lista.VerPrimero())

	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(5)
	iter.Insertar(4)

	require.EqualValues(t, 4, iter.VerActual())

	borrado = iter.Borrar()

	require.EqualValues(t, 4, borrado)
	require.EqualValues(t, 5, lista.VerUltimo())

	iter.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorExtBorrarMedio(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()

	iter.Siguiente()
	iter.Siguiente()

	borrado := iter.Borrar()

	require.EqualValues(t, 3, borrado)
	require.EqualValues(t, 4, iter.VerActual())

}

func TestIterBorrarUnicoELemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		iter.Borrar()
		if !iter.HaySiguiente() {
			break
		}
	}
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())

}

func TestIterInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	i := 4
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if i == 4 {
			iter.Insertar(4)
		}
		require.EqualValues(t, i, iter.VerActual())
		i--
	}

	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())

}

func TestIterBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.EqualValues(t, 5, iter.Borrar())

	i := 1
	for !lista.EstaVacia() {
		require.EqualValues(t, 4, lista.VerUltimo())
		require.EqualValues(t, i, lista.BorrarPrimero())
		i++
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())

}

func TestIterInsertarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())

}
