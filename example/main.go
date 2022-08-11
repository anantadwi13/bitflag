package main

import (
	"github.com/anantadwi13/bitflag"
	"log"
)

const (
	FeatureOne bitflag.Flag = iota
	FeatureTwo
	FeatureThree
)

func main() {
	moduleA := bitflag.Default

	moduleA.Set(FeatureOne, FeatureTwo)
	if !moduleA.Has(FeatureOne, FeatureTwo) {
		log.Panicln()
	}

	moduleA.Toggle(FeatureOne, FeatureThree)
	if !moduleA.Has(FeatureTwo, FeatureThree) || moduleA.Has(FeatureOne) {
		log.Panicln()
	}

	moduleA.Remove(FeatureTwo)
	if !moduleA.Has(FeatureThree) || moduleA.Has(FeatureOne, FeatureTwo) {
		log.Panicln()
	}

	moduleB := bitflag.New(FeatureTwo, FeatureThree)

	if !moduleB.Has(FeatureTwo, FeatureThree) {
		log.Panicln()
	}

	moduleB.Set(FeatureOne)
	if !moduleB.Has(FeatureOne, FeatureTwo, FeatureThree) {
		log.Panicln()
	}

	moduleB.Toggle(FeatureOne, FeatureThree)
	if !moduleB.Has(FeatureTwo) || moduleB.Has(FeatureOne, FeatureThree) {
		log.Panicln()
	}

	moduleB.Remove(FeatureTwo)
	if moduleB.Has(FeatureOne) || moduleB.Has(FeatureTwo) || moduleB.Has(FeatureThree) {
		log.Panicln()
	}
}
