package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32) uint {
    v4 := []int{-390, -9223372036854775808, -9223372036854775808, -9223372036854775808}
    var v5 string = ("\n\t" + ("a" + "hello"))
    var v6 uint8 = 127
    v6 = v6
    if ((1.5 >= float32((0.1 - float64((1e10 + 3.14159))))) || (true || false)) {
        var v7 int32 = p0
        var v8 float32 = 0.1
        _ = v7
        _ = v8
    } else {
        switch 0 {
        case 0:
            var v9 int = 2147483647
            _ = v9
        case -1:
            fmt.Println(v6, v6)
        }
        var v10 uint64 = 18446744073709551615
        v10 += (v10 * v10)
        _ = v10
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := []int{2147483647}
    v2 := []uint32{}
    var v3 int16 = 0
    var v11 uint8 = 1
    _ = len(v2)
    var v12 bool = false
    var v13 float32 = (((0.0 / (0.1 * -2.25)) - (3.14159 * float32(3.14159))) / 1.5)
    var v14 int16 = v3
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v13, v15)
        fmt.Println(v15, v3, v13)
        v17 := GenS1{F0: (1 <= 845), F1: (1 < -128), F2: v14}
        _ = fi16
        _ = v17
    }
    _ = len(v2)
    v15 = 1
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            v11 = v11
            var v20 int16 = 1
            _ = fi19
            _ = v20
        }
        v21 := GenS2{F0: -10}
        _ = fi18
        _ = v21
    }
    switch -512 {
    case 42:
        switch 32768 {
        case -9223372036854775808:
            v15 -= -1
        case 209:
            var v22 int = 4294967295
            _ = v22
        case 2147483647:
            fmt.Println(v13, v12)
        }
    case 1:
        v1 = append(v1, 854)
    default:
        v12 = (!((4294967295 == 251) || v12))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
