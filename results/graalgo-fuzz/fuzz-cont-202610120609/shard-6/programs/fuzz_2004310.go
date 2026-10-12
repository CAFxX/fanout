package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    v4 := []string{("\n\t" + "\n\t")}
    v5 := []uint32{2147483648}
    v6 := make(map[string]int)
    _ = len(v4)
    var v7 bool = (374 > 639)
    _ = len(v5)
    v7 = (593 > 330)
    v3 += v3
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v8 int64 = 42
    v9 := GenS1{F0: (v3 << ((2) & 63)), F1: 2, F2: 18446744073709551615}
    v1 -= ((0.1 - ((v1 * 3.14159) * 1.5)) * ((float64(float32(1e10)) / float64((-2.25 * (-2.25 + 0.1)))) - ((v1 + v1) - float64(3.14159))))
    v4 = append(v4, ("a" + "a"))
    fmt.Println(v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
