package main

import (
	"clientes"
	"contas"
	"fmt"
)

func PagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteThiago := clientes.Titular{
		Nome:      "Thiago",
		CPF:       "123.123.123-12",
		Profissao: "Desenvolvedor",
	}

	contaDoThiago := contas.ContaCorrente{Titular: clienteThiago}
	contaDoThiago.Depositar(30000)

	PagarBoleto(&contaDoThiago, 2000)

	fmt.Println(contaDoThiago.ObterSaldo())
}
