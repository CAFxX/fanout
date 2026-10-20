package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32) int32 {
    v6 := GenS1{F0: -1}
    p0 = 4294967295
    if false {
        var v7 uint64 = 4294967295
        v8 := GenS1{F0: 42}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    v4 := []int{128}
    var v5 string = "\n\t"
    v9 := GenS1{F0: -774}
    v10 := GenS2{F0: (!false), F1: int16(v2)}
    if true {
        v11 := GenS2{F0: false, F1: 380}
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                if true {
                    v2 = (v2 + (v2 % ((v2) | 1)))
                    fmt.Println(fi13, fi12, v5)
                    var v14 float64 = float64(-2.25)
                    _ = v14
                }
                _ = fi13
            }
            _ = fi12
        }
        var v15 uint8 = 0
        _ = v11
        _ = v15
    }
    v16 := GenS1{F0: -1}
    fmt.Println(v2, v5)
    var v17 float32 = 0.0
    v2 -= (v2 - 297)
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float64 = float64((0.0 - 0.1))
        var v20 uint16 = 426
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
