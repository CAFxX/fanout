package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    v1 := []int{}
    v2 := []string{"hello", "\n\t", ""}
    var v3 int8 = -103
    v4 := []uint32{0}
    var v5 float64 = -2.25
    v6 := make(map[uint32]string)
    var v7 int16 = -1
    v7 -= (v7 & -795)
    switch -765 {
    case 4294967295:
        var v8 int8 = (^(v3))
        _ = v8
    case -2147483648:
        switch -1 {
        case 127:
            var v9 float64 = (-2.25 + v5)
            _ = v9
        case -9223372036854775808:
            var v10 uint16 = 853
            _ = v10
        default:
            v6[2] = "a"
        }
    case 1:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 4294967295, F1: (int(v3) < fi11)}
            _ = fi11
            _ = v12
        }
    }
    switch 420 {
    case 4294967295:
        for fi13 := 0; fi13 < 1; fi13++ {
            v3 += (-(v3))
            _ = fi13
        }
    case -597:
        var v14 uint32 = 546
        _ = v14
    }
    v15 := GenS1{F0: 255, F1: true}
    v7 = 366
    v7 -= v7
    var v16 float32 = ((-2.25 - (-2.25 + (float32(v5) - float32(float64(0.1))))) - 1e10)
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS2{F0: ("a" + "foo bar"), F1: 218}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int8 = (v3 / (((v3 ^ 127)) | 1))
        fmt.Println(v5, v3)
        v2 = append(v2, "hello")
        _ = fi19
        _ = v20
    }
    v7 -= v7
    v7 = -91
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
