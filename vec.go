package rt2

type Vec struct {
	X, Y float64
}

// Add sets v to a+b and returns v.
func (v *Vec) Add(a, b *Vec) *Vec {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	return v
}

// Sub sets v to a-b and returns v.
func (v *Vec) Sub(a, b *Vec) *Vec {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	return v
}

// Mul sets v to a*f and returns v.
func (v *Vec) Mul(a *Vec, f float64) *Vec {
	v.X = a.X * f
	v.Y = a.Y * f
	return v
}

// Div sets v to a/f and returns v.
func (v *Vec) Div(a *Vec, f float64) *Vec {
	v.X = a.X / f
	v.Y = a.Y / f
	return v
}

// Neg sets v to -a and returns v.
func (v *Vec) Neg(a *Vec) *Vec {
	v.X = -a.X
	v.Y = -a.Y
	return v
}
