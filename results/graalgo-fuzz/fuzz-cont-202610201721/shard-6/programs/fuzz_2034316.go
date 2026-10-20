package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8) uint {
    var v6 int8 = (p0 ^ p0)
    v7 := GenS1{F0: 127, F1: 140}
    switch (("" + "") + "hello") {
    case "\n\t":
        v8 := GenS1{F0: 2, F1: 1}
        _ = v8
    case "":
        var v9 int64 = -569
        _ = v9
    case "foo bar":
        v10 := []int{}
        _ = v10
    }
    fmt.Println(p0)
    if (true && (false && true)) {
        p0 += 1
        fmt.Println(p0)
        v11 := GenS1{F0: 255, F1: 2}
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint64) bool {
    var v12 uint16 = 0
    v13 := GenS1{F0: 207, F1: 138}
    p1 = p1
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 float32 = ((1.5 * (3.14159 - -2.25)) * float32(((((float64(0.0) - 1e10) + 0.1) + 3.14159) / float64(3.14159))))
    var v2 uint = 4294967295
    v3 := make(map[int]uint32)
    v4 := []int{-9223372036854775808, -2147483648, 20}
    var v5 int8 = 0
    switch 2147483647 {
    case 1:
        v14 := v3[0]
        _ = v14
    default:
        delete(v3, 212)
    }
    v15 := GenS1{F0: 2, F1: 255}
    var v16 float64 = 3.14159
    if len(v4) > 0 {
        v4[0] = -72
    }
    v5 += (v5 - (v5 - (v5 % ((v5) | 1))))
    var v17 uint64 = 412
    fmt.Println(v5)
    fmt.Println(v17, v2)
    if (true && (((v5 ^ (v5 ^ (v5 - (v5 | (v5 - 4))))) >= (v5 * v5)) || (!(v2 > genFn0(v5))))) {
        v3[4294967295] = 2
        v18 := GenS1{F0: 0, F1: 2}
        _ = v18
    } else {
        delete(v3, -25)
        fmt.Println(v1, v1)
        var v19 int8 = 78
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(fi20)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
