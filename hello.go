package hello

import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

func Go() string {
	return quote.GoV3()
}
