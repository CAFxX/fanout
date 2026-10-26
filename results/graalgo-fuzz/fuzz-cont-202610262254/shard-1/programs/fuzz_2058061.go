package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
func genFn0() uint8 {
    var v7 uint16 = 42
    switch "hello" {
    case "":
        v7 += v7
    case "hello":
        fmt.Println(v7)
    case "\n\t":
        v8 := GenS1{F0: -758, F1: false, F2: 128}
        _ = v8
    default:
        var v9 uint = 255
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = (((1e10 / (-2.25 / float64(float32(0.1)))) + (1e10 / 1.0)) * (-2.25 - 0.0))
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", "a"}
    v5 := []uint32{v1}
    v6 := make(map[string]int)
    v10 := GenS1{F0: 168, F1: (!false), F2: 0}
    v1 += (^(v1))
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 = 2
        _ = fi11
    }
    var v12 bool = true
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: 1, F1: (false && (0 < 0)), F2: 1}
        fmt.Println(v2, v1, v1)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint32 = v1
        v5 = append(v5, v1)
        fmt.Println(v16, fi15, fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println(v12, v1, v2)
    var v17 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
