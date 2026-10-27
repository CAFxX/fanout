package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint, p1 int8) uint32 {
    v7 := GenS1{F0: 18446744073709551615, F1: -491, F2: 1}
    v8 := GenS1{F0: 42, F1: -32768, F2: 42}
    switch 32767 {
    case 1:
        var v9 float32 = 1e10
        _ = v9
    case 32767:
        p0 -= p0
    case -2147483648:
        var v10 int16 = -1
        _ = v10
    }
    var v11 int16 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 1
}
func genFn1() int16 {
    var v12 int32 = -1
    fmt.Println(v12)
    v13 := []int{}
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (!(false || (false && false)))
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    var v5 uint = v4
    var v6 bool = true
    fmt.Println(v4, v5, v2)
    v14 := GenS1{F0: 222, F1: -1, F2: 2147483647}
    var v15 string = "a"
    v3[-742] = 128
    v4 += 0
    if ((0 <= 9223372036854775807) || true) {
        var v16 float64 = (-2.25 + ((float64(3.14159) + 1.5) - 1e10))
        _ = v16
    } else {
        fmt.Println(v2, v4)
        v3[-586] = 32767
        _ = len(v3)
    }
    v17 := GenS2{F0: -128, F1: 167}
    v18 := GenS1{F0: (v1 ^ 4294967295), F1: 42, F2: 596}
    v6 = ((-2.25 * (0.0 / -2.25)) < ((1e10 + 1e10) + 3.14159))
    var v19 int16 = 32767
    var v20 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
