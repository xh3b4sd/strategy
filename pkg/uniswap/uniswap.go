package uniswap

import "math"

// Liquidity returns the amount of liquidity given a position's parameters. In
// the pool ETH/USDC, ETH is x (token0) and USDC is y (token1). The middle of
// the price range at position creation is p and the lower and upper bounds of
// said price range are a and b respectively.
func Liquidity(x float64, y float64, p float64, a float64, b float64) float64 {
	if p <= a {
		return liquidity0(x, a, b)
	}

	if p >= b {
		return liquidity1(y, a, b)
	}

	return math.Min(liquidity0(x, p, b), liquidity1(y, a, p))
}

// AmountOfX returns the amount of token0 given a position's parameters. In a
// pool ETH/USDC, AmountOfX calculates the amount of ETH, given a position's
// liquidity initialized according to Liquidity. Here p might very well be
// outside the price range of a and b.
func AmountOfX(l float64, p float64, a float64, b float64) float64 {
	p = math.Max(math.Min(p, b), a)
	return l * ((math.Sqrt(b) - math.Sqrt(p)) / (math.Sqrt(p) * math.Sqrt(b)))
}

// AmountOfY returns the amount of token1 given a position's parameters. In a
// pool ETH/USDC, AmountOfY calculates the amount of USDC, given a position's
// liquidity initialized according to Liquidity. Here p might very well be
// outside the price range of a and b.
func AmountOfY(l float64, p float64, a float64, b float64) float64 {
	p = math.Max(math.Min(p, b), a)
	return l * (math.Sqrt(p) - math.Sqrt(a))
}

func liquidity0(x float64, a float64, b float64) float64 {
	return x * ((math.Sqrt(a) * math.Sqrt(b)) / (math.Sqrt(b) - math.Sqrt(a)))
}

func liquidity1(y float64, a float64, b float64) float64 {
	return y / (math.Sqrt(b) - math.Sqrt(a))
}
