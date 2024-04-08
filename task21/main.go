package main

// Реализовать паттерн «адаптер» на любом примере

// Пример 1
// В этом примере у нас есть два различных интерфейса: Printer и Scanner.
// Создается адаптер PrinterAdapter, который оборачивает объект Scanner и
// предоставляет методы, совместимые с интерфейсом Printer

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type ScannerDevice struct{}

func (s *ScannerDevice) Scan() {
	fmt.Println("Сканирую документ")
}

type PrinterAdapter struct {
	Scanner Scanner
}

func (a *PrinterAdapter) Print() {
	a.Scanner.Scan()
}

func main() {
	scanner := &ScannerDevice{}
	adapter := &PrinterAdapter{
		Scanner: scanner,
	}
	adapter.Print()
}

// Пример 2
// В этом примере у нас есть интерфейс PaymentProcessor и структура BankPaymentProcessor,
// которая не совместима с интерфейсом. Мы создаем адаптер PaymentGatewayAdapter,
// который оборачивает объект BankPaymentProcessor и предоставляет методы,
// совместимые с интерфейсом PaymentProcessor.

/*
import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type BankPaymentProcessor struct{}

func (b *BankPaymentProcessor) ProcessPayment(amount int) {
	fmt.Println("Processing payment via bank:", amount)
}

type PaymentGatewayAdapter struct {
	BankProcessor BankPaymentProcessor
}

func (a *PaymentGatewayAdapter) ProcessPayment(amount float64) {
	a.BankProcessor.ProcessPayment(int(amount))
}

func main() {
	bankProcessor := &BankPaymentProcessor{}
	adapter := &PaymentGatewayAdapter{
		BankProcessor: bankProcessor,
	}
	adapter.ProcessPayment(100.50)
}
*/
