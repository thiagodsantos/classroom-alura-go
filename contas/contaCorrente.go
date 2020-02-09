package contas

import "clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor do depósito é menor que zero", c.saldo
}

func (c *ContaCorrente) Tranferir(valorTransferencia float64, contaDestino *ContaCorrente) {
	if valorTransferencia <= c.saldo {
		c.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
