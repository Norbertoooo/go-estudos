package main

import (
	"fmt"
)

func main() {

	println("estou aqui")

	// variavel
	var variavel int
	var variavelString string

	const n = 500000000
	const s string = "constante"

	fmt.Println(n)
	fmt.Println(s)

	// toda variavel deve ser utilizada, caso contrario não irá compilar
	// valor default como 0 para int
	fmt.Println(variavel)
	fmt.Println(variavelString)

	// declaracao e atribuição usando o operador :=
	maisUmaVariavel := 123

	fmt.Println(maisUmaVariavel)

	// existe 2 formas de controles, if e switch

	// if simples

	segundaVariavel := 12

	if segundaVariavel < 10 {
		fmt.Println("aaaa")
	} else {
		fmt.Println("bbbb")
	}
	// if com inicializacao
	if i := 5; i <= 5 {
		fmt.Println("valor maior ou igual que 5")
	} else {
		fmt.Println("valor menor que 5")
	}

	// switch

	terceiraVariavel := "draven"

	switch terceiraVariavel {
	case "aaaa":
		fmt.Println("aaaa")
	case "bbbb":
		fmt.Println("bbbb")
	default:
		fmt.Println("draven")
	}

	// switch com fallthrough

	fmt.Println("------------ Switch com fallthrough ----------------")
	switch terceiraVariavel {

	case "draven":
		fmt.Println("2222222")
		fallthrough
	case "":
		fmt.Println("iiiiiiii")
		fallthrough
	default:
		fmt.Println("draven")
	}

	// switch true

	fmt.Println("------------ Switch true ----------------")
	switch {
	case terceiraVariavel == "draven":
		fmt.Println("mundo do draven")
	case terceiraVariavel == "bardo":
		fmt.Println("own")
	default:
		fmt.Println("deu ruim")
	}

	// Existe 4 estruturas de dados, arrays, slices, maps e structs

	// arrays possuem tamanho fixo e iniciam com zero
	fmt.Println("--------------- arrays ----------------")

	// sem inicializar, valores default 0
	var array [3]int

	fmt.Println("valor default no array: ", array)
	fmt.Println("valor no indice 2: ", array[2])
	fmt.Println("tamanho do array: ", len(array))
	fmt.Println("capacidade do array: ", cap(array))

	array[1] = 12

	fmt.Println("valor no array: ", array)

	// inicializando
	outroArray := [3]float64{1.2, 2.2, 3.2}
	fmt.Println("valor do array: ", outroArray)

	// bidimencional so por curiosidade mesmo, nunca vai ser usado
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("--------------- slices ----------------")

	// slice nao possui tamanho especifico
	var pedaco []int
	fmt.Println(pedaco)
	fmt.Printf("tamanho: %d capacidade: %d ", len(pedaco), cap(pedaco))
	fmt.Println("slice é igual a nil: ", pedaco == nil)

	// quando criado com a funcao make, o default 0 eh atribuido
	sliceComMake := make([]int, 5)

	fmt.Println(sliceComMake)

	fmt.Println("slice é igual a nil: ", sliceComMake == nil)
	fmt.Printf("tamanho: %d capacidade: %d ", len(sliceComMake), cap(sliceComMake))

	sliceComMake[2] = 12

	fmt.Println(sliceComMake[2])

	sliceComMake = append(sliceComMake, 1)

	fmt.Println(sliceComMake)
	fmt.Printf("tamanho: %d capacidade: %d ", len(sliceComMake), cap(sliceComMake))

	sliceComMake = append(sliceComMake, 1, 2, 3, 4, 5, 6, 7, 8, 10)

	fmt.Println(sliceComMake)
	fmt.Printf("tamanho: %d capacidade: %d ", len(sliceComMake), cap(sliceComMake))

	outroSlice := make([]int, len(sliceComMake))

	// para copiar o slice de destino precisa ter o mesmo tamanho e tipo do original
	copy(outroSlice, sliceComMake)
	fmt.Println("Outro slice copiado: ", outroSlice)
	fmt.Printf("tamanho: %d capacidade: %d ", len(outroSlice), cap(outroSlice))

	// consegue recortar um slice de acordo com o indice inicial e final
	cortado := sliceComMake[0:5]

	fmt.Println("cortado: ", cortado)

	cortado2 := sliceComMake[5:10]

	fmt.Println("cortado2: ", cortado2)

	clear(sliceComMake)
	fmt.Println("slice zerado com clear", sliceComMake)
	fmt.Println("tamanho: ", len(sliceComMake), " capacidade:", cap(sliceComMake))

	// slice iniciando
	sliceIniciado := []string{"a", "b", "c"}

	fmt.Println(sliceIniciado)
	fmt.Printf("tamanho: %d capacidade: %d ", len(sliceIniciado), cap(sliceIniciado))

	// sempre que usa append para adicionar um novo valor, o valor é adicionado no proximo indice e dobra a capacidade do slice
	sliceIniciado = append(sliceIniciado, "d")

	fmt.Println(sliceIniciado)
	fmt.Printf("tamanho: %d capacidade: %d ", len(sliceIniciado), cap(sliceIniciado))

	fmt.Println("--------------- maps ----------------")

	var mapaSimples map[int]string

	fmt.Println(mapaSimples)
	fmt.Println(len(mapaSimples))
	fmt.Println("eh nil? ", mapaSimples == nil)

	mapaComMake := make(map[int]string)

	fmt.Println(mapaComMake)
	fmt.Println(len(mapaComMake))
	fmt.Println("eh nil? ", mapaSimples == nil)

	mapaComMake[0] = "aaaa"
	fmt.Println("mapa com make: ", mapaComMake)
	fmt.Println("tamanho: ", len(mapaComMake))
	fmt.Println("primeiro valor: ", mapaComMake[0])

	mapIniciado := map[int]string{0: "a", 1: "b", 2: "c"}
	fmt.Println("mapa iniciado: ", mapIniciado)
	fmt.Println("tamanho: ", len(mapIniciado))

	delete(mapIniciado, 2)
	fmt.Println("mapa iniciado: ", mapIniciado)
	fmt.Println("tamanho: ", len(mapIniciado))

	// se a chave nao existe no mapa ou se o mapa eh nil, nada acontece
	delete(mapIniciado, 3)

	fmt.Println("mapa iniciado: ", mapIniciado)
	fmt.Println("tamanho: ", len(mapIniciado))

	// zera o mapa, tanto valores como tamanho
	clear(mapIniciado)

	fmt.Println("mapa iniciado: ", mapIniciado)
	fmt.Println("tamanho: ", len(mapIniciado))
	fmt.Println("eh nil? ", mapaSimples == nil)

	fmt.Println("---------------- Estrutura de repeticao ---------------")
	// so existe uma estrutura de repeticao (for) que possui 3 variacoes

	fmt.Println("---------------- tradicional ---------------")
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	fmt.Println("---------------- com range ---------------")
	for i, v := range array {
		fmt.Println("indice: ", i, " valor: ", v)
	}

	i := 5
	for i != 0 {
		fmt.Println(i)
		i--
	}

	// loop
	for {
		fmt.Println("loop")
		break
	}
}
