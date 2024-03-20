package main

import "fmt"

func main() {

	println("estou aqui")

	// variavel
	var variavel int
	var variavelString string

	// toda variavel deve ser utilizada, caso contrario não irá compilar
	// valor default como 0 para int
	println(variavel)
	println(variavelString)

	// declaracao e atribuição usando o operador :=
	maisUmaVariavel := 123

	println(maisUmaVariavel)

	// if simples

	segundaVariavel := 12

	if segundaVariavel < 10 {
		println("aaaa")
	} else {
		println("bbbb")
	}
	// if com inicializacao
	if i := 5; i <= 5 {
		println("valor maior ou igual que 5")
	} else {
		println("valor menor que 5")
	}

	// switch

	terceiraVariavel := "draven"

	switch terceiraVariavel {
	case "aaaa":
		println("aaaa")
	case "bbbb":
		println("bbbb")
	default:
		println("draven")
	}

	// switch com fallthrough

	println("------------ Switch com fallthrough ----------------")
	switch terceiraVariavel {

	case "draven":
		println("2222222")
		fallthrough
	case "":
		println("iiiiiiii")
		fallthrough
	default:
		println("draven")
	}

	// switch true

	println("------------ Switch true ----------------")
	switch {
	case terceiraVariavel == "draven":
		println("mundo do draven")
	case terceiraVariavel == "bardo":
		println("own")
	default:
		println("deu ruim")
	}

	// arrays possuem tamanho fixo e iniciam com zero
	println("--------------- arrays ----------------")

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

}
