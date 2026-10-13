package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint64 = 758
    var v4 uint8 = 127
    if true {
        v5 := GenS1{F0: false}
        _ = v5
    } else {
        v6 := []int{-2147483648, 42}
        v4 += (v4 * (v4 - v4))
        _ = v6
    }
    switch 0 {
    case -319:
        fmt.Println(v3, v4)
    case 791:
        v3 += (v3 * (v3 + (v3 - (v3 | (v3 * v3)))))
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v4 = v4
        fmt.Println(v4)
        v8 := GenS2{F0: v4, F1: 9223372036854775807}
        _ = fi7
        _ = v8
    }
    var v9 int8 = -128
    var v10 uint64 = v3
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint16 = 157
        v10 = v10
        var v13 float32 = ((0.0 + 0.1) * -2.25)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v3, v10)
    v4 = v4
    if false {
        fmt.Println(v4)
        switch 9223372036854775807 {
        case -419:
            var v14 uint8 = 128
            _ = v14
        case 9223372036854775807:
            var v15 bool = (1e10 == float32((float64(1.5) * 1.5)))
            _ = v15
        case 32768:
            switch -21 {
            case -1:
                v9 = v9
            case -9223372036854775808:
                v10 += (v10 << ((207) & 63))
            default:
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 int8 = 0
                    v17 += (v17 | v9)
                    _ = fi16
                    _ = v17
                }
            }
        default:
            v18 := []int{}
            _ = v18
        }
        v4 = 2
    } else {
        var v19 int32 = -1
        _ = v19
    }
    var v20 uint32 = 550
    v4 = (v4 + 42)
    var v21 bool = (((881 <= 42) && (!false)) || true)
    var v22 int32 = 42
    var v23 uint16 = 62
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
