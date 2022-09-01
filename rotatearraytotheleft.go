package main

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func rotateL(a []int, i int) {
	//rotate the array to the left * i
    i = i % len(a)
    if i < 0 {
        i += len(a)
    }
    for c := 0; c < gcd(i, len(a)); c++ {
        t := a[c]
        j := c
        for {
            k := j + i
            if k >= len(a) {
                k -= len(a)
            }
            if k == c {
                break
            }
            a[j] = a[k]
            j = k
        }
        a[j] = t
    }
}
