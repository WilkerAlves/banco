package main

import (
	"fmt"

	"github.com/wilker/banco/clientes"
	"github.com/wilker/banco/contas"
)

func main() {
	testesContaCorrente()
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(50)
	PagarBoleto(&contaDoDenis, 10)
	fmt.Println("conta do denis:", contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(100)
	contaDaLuisa.Sacar(50)
	PagarBoleto(&contaDaLuisa, 10)
	fmt.Println("conta da luiza:", contaDaLuisa.ObterSaldo())
}

func testesContaCorrente() {
	clienteSilvia := clientes.Titular{
		Nome:      "Silvia",
		CPF:       "533.845.390-87",
		Profissao: "DEV",
	}

	contaDaSilvia := contas.ContaCorrente{
		Titular:       clienteSilvia,
		NumeroAgencia: 123,
		NumeroConta:   1234356,
	}
	contaDaSilvia.Depositar(500)

	contaDoGustavo := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Gustavo",
			CPF:       "906.163.810-00",
			Profissao: "PO",
		},
		NumeroAgencia: 456,
		NumeroConta:   765312,
	}
	contaDoGustavo.Depositar(500)

	fmt.Println(contaDaSilvia.Sacar(300))

	status, valor := contaDaSilvia.Depositar(300)
	fmt.Println(status, valor)

	fmt.Println("saldo antes da transferencia:", contaDaSilvia.ObterSaldo())
	fmt.Println("saldo antes da transferencia:", contaDoGustavo.ObterSaldo())
	statusTransferencia := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(statusTransferencia)
	fmt.Println("saldo depois da transferencia:", contaDaSilvia.ObterSaldo())
	fmt.Println("saldo depois da transferencia:", contaDoGustavo.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}
