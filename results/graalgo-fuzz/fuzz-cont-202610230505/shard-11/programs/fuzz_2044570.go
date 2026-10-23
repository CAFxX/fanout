package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 uint8, p1 uint16) uint16 {
    fmt.Println(p0, p0)
    p0 = 127
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 uint16) int8 {
    var v7 float32 = float32(float64(((1e10 * (float32(0.1) - (-2.25 / -2.25))) / float32(1.5))))
    p0 = p0
    v7 = ((0.0 - (0.1 * ((((v7 / ((v7 * 1e10) - v7)) / (((3.14159 * (v7 + (v7 / v7))) * v7) / 1.0)) - 3.14159) - v7))) - -2.25)
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 int = -319
    v2 := make(map[uint32]bool)
    var v3 string = "\n\t"
    var v4 string = ""
    v5 := []int{0, (v1 & 297), (v1 / ((-2147483648) | 1))}
    var v6 uint16 = 65535
    var v8 uint64 = 4294967295
    if false {
        var v9 uint8 = 2
        _ = v9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 := GenS1{F0: genFn1((v6 * (v6 - genFn0(2, (v6 & 32767))))), F1: 71}
    v2[26] = (true || false)
    var v11 uint8 = 42
    if len(v5) > 0 {
        v5[0] = 32768
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int = (fi12 / (((fi12 / ((9223372036854775807) | 1))) | 1))
        _ = fi12
        _ = v13
    }
    switch v1 {
    case 9223372036854775807:
        var v14 float64 = float64(float32(3.14159))
        _ = v14
    case 128:
        fmt.Println(v8)
    default:
        var v15 bool = false
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
