package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{42, 702, 578}
    v2 := []int{2147483647}
    var v3 int64 = -938
    v4 := GenS1{F0: 1}
    var v5 int64 = (v3 / (((v3 << ((255) & 63))) | 1))
    v5 = v5
    v1 = append(v1, 2147483647)
    v3 += v5
    if (32767 != 0) {
        v3 += (v5 % ((v5) | 1))
        v3 = 1
    }
    var v6 uint16 = 0
    v2 = append(v2, 127)
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 += (v5 + (v5 % ((v5) | 1)))
        _ = fi7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
