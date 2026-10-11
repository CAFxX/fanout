package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 int) int8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 127, F1: 2}
        _ = fi4
        _ = v5
    }
    v6 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 string, p2 int) int32 {
    fmt.Println(p2)
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    v2 = v2
    var v7 int = -2147483648
    v2 += 255
    v2 -= v2
    v2 += v2
    v2 -= (v2 ^ 32767)
    delete(v3, 0)
    if true {
        v8 := GenS1{F0: 2, F1: 1}
        _ = v8
    }
    v9 := GenS2{F0: (v7 << ((18446744073709551615) & 63)), F1: -1}
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
