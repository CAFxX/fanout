package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 42
    v3 := []uint32{v2, v2}
    v4 := []uint32{}
    var v5 uint16 = 846
    v5 += (v5 ^ v5)
    v3 = append(v3, v2)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: v2, F1: (v2 << ((1) & 63)), F2: (v1 + v1)}
        var v8 uint32 = 462
        var v9 uint = 469
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    var v11 float32 = float32(0.1)
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v5 -= 255
            _ = len(v3)
            _ = fi13
        }
        var v14 string = "hello"
        _ = fi12
        _ = v14
    }
    v10 -= (^(v10))
    if false {
        switch 236 {
        case 937:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -1:
            v15 := GenS2{F0: 62}
            _ = v15
        case 553:
            v5 += v5
        default:
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 string = "\n\t"
                if false {
                    v1 += v1
                    fmt.Println(v5, v2, v2)
                }
                v10 = v10
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS2{F0: 8}
        _ = v18
    }
    var v19 bool = true
    _ = len(v3)
    var v20 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
