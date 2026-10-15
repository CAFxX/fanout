package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0(p0 int64) int16 {
    var v4 int = 42
    v5 := []int{int(p0), (^(v4)), (v4 | (v4 | 4294967295)), v4}
    _ = p0
    _ = v4
    _ = v5
    return 902
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]uint32)
    v3 := []uint32{157, 2147483647}
    v6 := GenS1{F0: 4294967295, F1: 32767, F2: 368}
    v1 -= 1e10
    var v7 int8 = 24
    if len(v3) > 0 {
        v3[0] = 1
    }
    v8 := GenS1{F0: 2147483648, F1: 648, F2: 980}
    v3 = append(v3, 0)
    v7 = v7
    fmt.Println(v1, v7)
    var v9 int32 = 42
    switch 1 {
    case 127:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case 32767:
        var v10 uint32 = 87
        _ = v10
    case -2147483648:
        var v11 int32 = (v9 ^ (v9 | (v9 / (((v9 % ((873) | 1))) | 1))))
        _ = v11
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        delete(v2, 42)
        _ = fi12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: 2147483648, F1: 1000, F2: v9}
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    var v15 int8 = (^(v7))
    delete(v2, 994)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
