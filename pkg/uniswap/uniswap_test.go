package uniswap

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Liquidity(t *testing.T) {
	testCases := []struct {
		// In a pool ETH/USDC, x is the initial amount of ETH.
		x float64
		// In a pool ETH/USDC, y is the initial amount of USDC.
		y float64
		// In a pool ETH/USDC, p is the initial price of a position's range.
		p float64
		// In a pool ETH/USDC, a is the lower boundary of a position's range.
		a float64
		// In a pool ETH/USDC, b is the upper boundary of a position's range.
		b float64
		// In a pool ETH/USDC, c is the new price after it has changed.
		c float64
		// In a pool ETH/USDC, L is the initial amount of a position's
		// liquidity.
		L float64
		// In a pool ETH/USDC, X is the amount of ETH after c has been applied.
		X float64
		// In a pool ETH/USDC, Y is the amount of USDC after c has been applied.
		Y float64
		// R is the realized loss incurred after c has been applied.
		R float64
	}{
		// Case 0
		{
			x: 7.302967433402214,
			y: 10000,
			p: 1250,
			a: 1000,
			b: 1500,
			c: 1001,
			L: 2679.1246264404454,
			X: 15.504330210158221,
			Y: 42.350094896836815,
			R: 0.1864487876359352,
		},
		// Case 1
		{
			x: 7.302967433402214,
			y: 10000,
			p: 1250,
			a: 1000,
			b: 1500,
			c: 1000,
			L: 2679.1246264404454,
			X: 15.546659145875672,
			Y: 0,
			R: 0.18726042051470126,
		},
		// Case 2
		{
			x: 7.302967433402214,
			y: 10000,
			p: 1250,
			a: 1000,
			b: 1500,
			c: 999,
			L: 2679.1246264404454,
			X: 15.546659145875672,
			Y: 0,
			R: 0.18807316009418662,
		},
		// Case 3
		{
			x: 7.302967433402214,
			y: 10000,
			p: 1250,
			a: 1000,
			b: 1500,
			c: 500,
			L: 2679.1246264404454,
			X: 15.546659145875672,
			Y: 0,
			R: 0.5936302102573506,
		},
		// Case 4
		{
			x: 7.302967433402214,
			y: 10000,
			p: 1250,
			a: 1000,
			b: 1500,
			c: 1500,
			L: 2679.1246264404454,
			X: 0,
			Y: 19040.69105618437,
			R: 0.09132952613313927,
		},
		// Case 5
		{
			x: 7.698003589195008,
			y: 10000,
			p: 1250,
			a: 1150,
			b: 1350,
			c: 1100,
			L: 6926.698897495369,
			X: 15.736341464510037,
			Y: 0,
			R: 0.11785085217457547,
		},
		// Case 6
		{
			x: 7.698003589195008,
			y: 10000,
			p: 1250,
			a: 1150,
			b: 1350,
			c: 1400,
			L: 6926.698897495369,
			X: 0,
			Y: 19607.380428618937,
			R: 0.05630327057242024,
		},
		// Case 7
		{
			x: 7.698003589195008,
			y: 10000,
			p: 1250,
			a: 1150,
			b: 1350,
			c: 1200,
			L: 6926.698897495369,
			X: 11.435708089194486,
			Y: 5052.100301048237,
			R: 0.04319298176211461,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var l float64
			{
				l = liquidity1(tc.y, tc.p, tc.b)
			}

			var x float64
			{
				x = AmountOfX(l, tc.p, tc.a, tc.b)
			}

			if x != tc.x {
				t.Fatalf("x\n\n%s\n", cmp.Diff(tc.x, x))
			}

			var L float64
			{
				L = Liquidity(x, tc.y, tc.p, tc.a, tc.b)
			}

			if L != tc.L {
				t.Fatalf("L\n\n%s\n", cmp.Diff(tc.L, L))
			}

			var X float64
			{
				X = AmountOfX(L, tc.c, tc.a, tc.b)
			}

			if X != tc.X {
				t.Fatalf("X\n\n%s\n", cmp.Diff(tc.X, X))
			}

			var Y float64
			{
				Y = AmountOfY(L, tc.c, tc.a, tc.b)
			}

			if Y != tc.Y {
				t.Fatalf("Y\n\n%s\n", cmp.Diff(tc.Y, Y))
			}

			var R float64
			if tc.c <= tc.p {
				var s float64
				var e float64
				{
					s = (x * tc.p) + tc.y
					e = (X * tc.c) + Y
				}

				{
					R = 1 - (e / s)
				}
			} else {
				var s float64
				var e float64
				{
					s = (x * tc.c) + tc.y
					e = (X * tc.c) + Y
				}

				{
					R = 1 - (e / s)
				}
			}

			if R != tc.R {
				t.Fatalf("R\n\n%s\n", cmp.Diff(tc.R, R))
			}
		})
	}
}
