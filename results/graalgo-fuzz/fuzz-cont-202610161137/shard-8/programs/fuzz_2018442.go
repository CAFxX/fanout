package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
func genFn0() bool {
    v4 := []int{}
    v5 := GenS1{F0: 169, F1: -166, F2: 2}
    var v6 int8 = -60
    _ = v4
    _ = v5
    _ = v6
    return true
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483648, 45, 103}
    v3 := make(map[uint32]uint32)
    v7 := GenS1{F0: 128, F1: 0, F2: 166}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 127, F1: 42, F2: 92}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 float32 = 1.5
            _ = fi10
            _ = v11
        }
        var v12 int8 = 127
        _ = fi8
        _ = v12
        _ = v9
    }
    var v13 int = 127
    v13 = (^(v13))
    fmt.Println(v13)
    var v14 int16 = 32767
    switch 1 {
    case -876:
        fmt.Println(v13, v13)
    case 126:
        var v15 int32 = -1
        _ = v15
    default:
        v14 += v14
    }
    v16 := GenS2{F0: ((42 <= 42) || (640 < 18446744073709551615)), F1: 128, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
