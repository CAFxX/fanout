package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1.5
    v5 := make(map[int]int)
    v6 := []string{("a" + "hello"), "\n\t"}
    v1 = (^(v1))
    v5[32767] = 0
    switch 1 {
    case 0:
        for fi7 := 0; fi7 < 1; fi7++ {
            v4 += 1.5
            v1 += (v1 & 1000)
            var v8 bool = false
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(v4)
    }
    if true {
        v9 := GenS1{F0: v1, F1: 42}
        v4 -= (v4 * (0.1 / (1e10 * 1.5)))
        _ = v9
    }
    switch -9223372036854775808 {
    case 933:
        _ = len(v5)
    case 1:
        v10 := GenS2{F0: 117, F1: (65535 <= 577)}
        _ = v10
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v11 := GenS1{F0: v1, F1: 0}
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := v2[2]
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: v1, F1: 1}
    v16 := GenS2{F0: -108, F1: (!false)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
