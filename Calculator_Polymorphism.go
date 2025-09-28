package main

type Operation interface {
	Execute(a, b float64) float64
	Name() string
}

// Structهای مختلف برای هر عملیات
type Add struct{}

func (op Add) Execute(a, b float64) float64 { return a + b }
func (op Add) Name() string                 { return "جمع" }

type Subtract struct{}

func (op Subtract) Execute(a, b float64) float64 { return a - b }
func (op Subtract) Name() string                 { return "تفریق" }

type Multiply struct{}

func (op Multiply) Execute(a, b float64) float64 { return a * b }
func (op Multiply) Name() string                 { return "ضرب" }

type Divide struct{}

func (op Divide) Execute(a, b float64) float64 { return a / b }
func (op Divide) Name() string                 { return "تقسیم" }
