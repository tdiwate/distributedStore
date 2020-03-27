package main

func hashFunc(s string, mod int) int {
	loc :=0
	for _, i := range s {
		loc += int(i)
	}
	return loc%mod
}
