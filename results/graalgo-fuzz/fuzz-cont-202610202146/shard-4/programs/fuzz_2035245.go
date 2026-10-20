package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint32) string {
    v4 := GenS1{F0: (p1 ^ p1), F1: true}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    return ""
}
func genFn1(p0 int32) int8 {
    var v5 int16 = 804
    v6 := GenS1{F0: 2, F1: (394 != 2)}
    _ = p0
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]bool)
    var v3 uint8 = 64
    fmt.Println(v1, v3)
    switch -9223372036854775808 {
    case -2147483648:
        v7 := GenS1{F0: 2, F1: (32767 == 32767)}
        _ = v7
    case 2147483647:
        fmt.Println(v1, v1)
    default:
        v1 += v1
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v3 = (v3 + v3)
        _ = fi8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: (!true)}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint8 = v3
        for fi13 := 0; fi13 < 11; fi13++ {
            v1 += ((float32((3.14159 + 0.0)) / v1) * (v1 / 1e10))
            fmt.Println(v1, fi13)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: 180}
    var v15 int32 = 2147483647
    v16 := GenS1{F0: 42, F1: (-1 >= 0)}
    if (1e10 <= (3.14159 - float64((((v1 - v1) - (float32(0.1) / ((float32((((1e10 * 1e10) * float64(v1)) * ((0.0 / float64((v1 + -2.25))) * 0.1))) - float32(0.1)) - (((3.14159 * 0.0) / float32(3.14159)) + v1)))) / (v1 - ((v1 - v1) + v1)))))) {
        fmt.Println(v1, v3)
        v17 := GenS1{F0: 736, F1: (!(!false))}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
