package main

import (
	"log"
	"testing"
)

func Test_asd(t *testing.T) {
	rsrp := 93

	switch {
	case rsrp <= 80:
		log.Println("1")
	case rsrp <= 81 || rsrp <= 90:
		i := float64(75 / 90 * rsrp)
		log.Println("2", i)
	case rsrp <= 91 || rsrp <= 99:
		i := 50 * rsrp / 99

		log.Println("3", i, rsrp)
	case rsrp >= 100:
		log.Println("4")
	}

}
