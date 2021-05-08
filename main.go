package goarea

import "math"

const Pi = 3.416

// Circ é responsavel por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é uma função que retorno o calculo de área
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
