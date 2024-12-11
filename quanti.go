import (
	"fmt"
	"math"
	"math/cmplx"
)

type Qubit struct {
	alpha complex128
	beta  complex128
}

func NewQubit(alpha, beta complex128) *Qubit {
	return &Qubit{alpha: alpha, beta: beta}
}

func (q *Qubit) Normalize() {
	norm := cmplx.Abs(q.alpha)*cmplx.Abs(q.alpha) + cmplx.Abs(q.beta)*cmplx.Abs(q.beta)
	if norm != 0 {
		q.alpha /= complex(math.Sqrt(norm), 0)
		q.beta /= complex(math.Sqrt(norm), 0)
	}
}

func Hadamard(q *Qubit) {
	alpha := q.alpha
	beta := q.beta

	q.alpha = (alpha + beta) / complex(math.Sqrt(2), 0)
	q.beta = (alpha - beta) / complex(math.Sqrt(2), 0)
}

func main() {
	q := NewQubit(1, 0)
	q.Normalize()
	fmt.Printf("Before Hadamard: alpha: %v, beta: %v\n", q.alpha, q.beta)

	Hadamard(q)
	fmt.Printf("After Hadamard: alpha: %v, beta: %v\n", q.alpha, q.beta)
}