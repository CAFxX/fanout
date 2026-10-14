package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint32) uint8 {
    var v5 uint32 = p0
    if true {
        v5 = (p0 << ((2) & 63))
    }
    v6 := []int{2147483647, -9223372036854775808}
    v5 -= p0
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    _ = p0
    _ = v5
    _ = v6
    return 131
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    var v3 int8 = 0
    var v4 int8 = (v3 << ((0) & 63))
    fmt.Println(v2, v3, v2)
    var v7 string = "a"
    v3 = (v3 - v4)
    var v8 int64 = -871
    _ = len(v1)
    v9 := []int{837, 128, 1, 650}
    var v10 bool = (!((v4 <= (v3 & (v3 >> ((1) & 63)))) || ((255 != 0) && true)))
    if v10 {
        var v11 int16 = 32767
        _ = v11
    } else {
        var v12 float32 = ((1e10 * ((0.1 * 0.1) * ((-2.25 * (float32((-2.25 + 0.0)) * 1.5)) * (0.1 * ((3.14159 - 3.14159) / 1.0))))) + 0.0)
        _ = v12
    }
    if v10 {
        v13 := GenS1{F0: (v4 % (((^(v4))) | 1)), F1: 0}
        v14 := GenS1{F0: v3, F1: genFn0(65535)}
        var v15 int16 = 1
        _ = v13
        _ = v14
        _ = v15
    }
    if v10 {
        var v16 float32 = 1.5
        v4 -= v4
        _ = v16
    }
    var v17 float32 = 0.1
    v18 := GenS1{F0: v4, F1: 127}
    fmt.Println(v4, v7)
    switch -2147483648 {
    case -2147483648:
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v2, v4, v3)
            v20 := v1[fi19]
            _ = fi19
            _ = v20
        }
    case 32768:
        v4 -= 1
    case 127:
        fmt.Println(v4)
    default:
        v21 := GenS1{F0: v3, F1: genFn0(2)}
        _ = v21
    }
    v10 = false
    fmt.Println(v3)
    v8 = 12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
