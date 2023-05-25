package cola_prioridad

type heap[T comparable] struct {
	arreglo     []T
	funcion_cmp func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.funcion_cmp = funcion_cmp
	return heap
}
func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.arreglo = arreglo
	for i := len(arreglo); i >= 0; i-- {
		heap.downHeap(i)
	}
	return heap
}
func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {

}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arreglo) == 0
}
func (h *heap[T]) swap(viejo, nuevo int) {
	aux := h.arreglo[viejo]
	h.arreglo[viejo] = h.arreglo[nuevo]
	h.arreglo[nuevo] = aux
}
func (h *heap[T]) comparacionPadreHijo(padre, hijo int) bool {
	if h.funcion_cmp(h.arreglo[padre], h.arreglo[hijo]) > 0 {
		return false
	} else {
		return true
	}
}

func (h *heap[T]) upHeap(hijo int) {
	if hijo == 0 {
		return
	}
	posPadre := uint((hijo - 1) / 2)
	if h.comparacionPadreHijo(int(posPadre), hijo) {

		h.swap(int(posPadre), hijo)

		h.upHeap(int(posPadre))
	}
}
func (h *heap[T]) maximoDeTres(hIzq, hDer, padre int) int {
	if h.funcion_cmp(h.arreglo[padre], h.arreglo[hDer]) > 0 && h.funcion_cmp(h.arreglo[padre], h.arreglo[hIzq]) > 0 {
		return len(h.arreglo)
	} else if h.funcion_cmp(h.arreglo[padre], h.arreglo[hIzq]) > 0 && h.funcion_cmp(h.arreglo[hDer], h.arreglo[hIzq]) < 0 {
		return hIzq
	} else {
		return hDer
	}
}
func (h *heap[T]) downHeap(padre int) {
	if padre >= len(h.arreglo)-1 {
		return
	}
	hIzq := 2*padre + 1
	hDer := 2*padre + 2

	max := padre

	if hDer <= (len(h.arreglo) - 1) {
		max = h.maximoDeTres(hIzq, hDer, padre)
	} else if hIzq <= (len(h.arreglo)-1) && h.comparacionPadreHijo(int(padre), hIzq) {
		max = hIzq
	}

	h.swap(max, padre)

	h.downHeap(max)

}

func (h *heap[T]) Encolar(nuevo T) {
	h.arreglo = append(h.arreglo, nuevo)

	h.upHeap(len(h.arreglo) - 1)
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.arreglo[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := h.arreglo[0]
	h.swap(len(h.arreglo)-1, 0)

	h.arreglo = h.arreglo[:len(h.arreglo)-1]
	if !h.EstaVacia() {
		h.downHeap(0)
	}
	return dato
}

func (h *heap[T]) Cantidad() int {
	return len(h.arreglo)
}
