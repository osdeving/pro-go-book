package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	/*
	*
	* Definição
	*
	 */

	type Product struct {
		Name, Category string
		Price          float64
	}

	// Podemos embutir um tipo em outro, o nome é opcional
	type StockLevel struct {
		Product
		Count int
	}

	// Podemos criar sem usar o nome do tipo, isso é um campo embudito
	type Foo struct {
		int
		string
	}

	myF := func(user struct {
		string
		int
	}) {
		fmt.Println(user.string)
		fmt.Println(user.int)
	}

	myF(struct {
		string
		int
	}{"joao", 123})

	/*
	*
	*	Inicialização
	*
	 */

	p := Product{
		Name:     "Mouse",
		Category: "Informática",
		Price:    53.99,
	}

	// é permitido iniciar parcialmente e os valores não definidos serão
	// zero, string vazia ou false
	p2 := Product{Name: "Teclado"}

	// podemos iniciar sem mencionar o nome do campo, mas daí a ordem importa
	p3 := Product{"Teclado", "Informática", 150.99}
	// p4 := Product{"Mouse", 13.00} // ❌ Errado, o segundo campo é uma string
	p4 := Product{"Mouse", "Informática", 13.00} // ✅ Ok

	// Criar vazio (ambas formas abaixo são equivalentes)
	var p5 = new(Product)
	var p6 = &Product{}

	/*
	*
	* Acesso aos campos
	*
	 */
	p.Price = 1400.99

	p.Name = "Computador"

	fmt.Println(p.Name, p.Category, p.Price)
	fmt.Println(p2.Name, p2.Category, p2.Price)
	fmt.Println(p3.Name, p3.Category, p3.Price)
	fmt.Println(p4.Name, p4.Category, p4.Price)
	fmt.Println(p5.Name, p5.Category, p5.Price)
	fmt.Println(p6.Name, p6.Category, p6.Price)

	foo := Foo{1, "foo"}
	foo.string = "teste"
	fmt.Println(foo.string)

	stock1 := StockLevel{
		Product: Product{
			Name:     "Mouse",
			Category: "Informática",
			Price:    53.99,
		},
		Count: 10,
	}

	// Podemos iniciar o tipo embutido diretamente
	sock2 := StockLevel{Product{"Teclado", "Informática", 150.99}, 5}

	// Podemos acessar os campos do tipo embutido diretamente
	// sem precisar usar o nome do tipo como em stock1.Product.Name
	fmt.Println(stock1.Name, stock1.Category, stock1.Price, stock1.Count)

	fmt.Println(sock2.Name, sock2.Category, sock2.Price, sock2.Count)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct{Nome_direfenre_do_produto, CategoriaDoProduto string; Preco_Do_Produto float64} {"Mouse","Informática",53.99})
	fmt.Println("Builder Json:", builder.String())
}
