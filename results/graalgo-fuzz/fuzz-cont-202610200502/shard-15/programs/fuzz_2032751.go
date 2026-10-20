package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    var v3 int8 = 42
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    if true {
        if true {
            v6 := v4[342]
            _ = len(v5)
            _ = v6
        }
    } else {
        fmt.Println(v3, v1)
    }
    var v7 int32 = 2147483647
    if (65535 != 42) {
        fmt.Println(v7, v3, v1)
    } else {
        var v8 uint8 = 2
        _ = v8
    }
    switch v1 {
    case "a":
        v9 := GenS1{F0: 127, F1: (v7 - 1), F2: 42}
        _ = v9
    default:
        v10 := []int{-1, -1}
        _ = v10
    }
    var v11 uint32 = 2
    fmt.Println(v1)
    v11 = 280
    v3 = v3
    for fi12 := 0; fi12 < 4; fi12++ {
        switch 0 {
        case 582:
            if (!false) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v3 += (v3 ^ 60)
                    if true {
                        if false {
                            _ = len(v5)
                            v1 = "\n\t"
                            fmt.Println(fi12)
                        }
                        v14 := GenS1{F0: 127, F1: v7, F2: 0}
                        _ = v14
                    }
                    var v15 int32 = v7
                    _ = fi13
                    _ = v15
                }
                v11 -= v11
                v16 := v4[fi12]
                _ = v16
            }
        case 128:
            v17 := []int{128}
            _ = v17
        case -1:
            var v18 float64 = 1e10
            _ = v18
        default:
            v19 := GenS1{F0: 8, F1: v7, F2: 0}
            _ = v19
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
