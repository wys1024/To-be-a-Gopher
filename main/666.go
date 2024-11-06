package main

func romanToInt(s string) int {
	lmmap := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	if len(s) == 2 {
		v, e := lmmap[s]
		if e {
			return v
		}
	}
	sum := 0
	for s1 := range s {
		sum = sum + lmmap[string(s1)]
	}
	return sum
}
