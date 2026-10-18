package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }

func main() {
    var v1 int64 = -98
    v2 := make(map[string]uint32)
    var v3 bool = (true || (466 > 0))
    v4 := make(map[int]uint32)
    v1 += v1
    var v5 uint16 = 411
    if v3 {
        if (222 < 79) {
            v6 := GenS1{F0: 4294967295, F1: 42}
            _ = v6
        } else {
            v5 = (^(v5))
        }
        v7 := GenS2{F0: (v5 - (v5 * v5)), F1: 2147483647, F2: (v1 & 42)}
        _ = v7
    } else {
        fmt.Println(v1, v3)
        v8 := GenS1{F0: 734, F1: -2147483648}
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int = 4294967295
            v11 := v2["foo bar"]
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    var v12 uint64 = 1
    var v13 uint64 = v12
    var v14 float64 = 3.14159
    switch "foo bar" {
    case "\n\t":
        v15 := GenS2{F0: (v5 % (((v5 - (v5 % ((v5) | 1)))) | 1)), F1: -2147483648, F2: 1}
        _ = v15
    case "":
        v16 := []int{-2147483648, -242, 9223372036854775807}
        _ = v16
    default:
        switch 0 {
        case 128:
            v5 = v5
        case 9223372036854775807:
            v17 := GenS2{F0: v5, F1: 2147483647, F2: -802}
            _ = v17
        case -1000:
            var v18 bool = false
            _ = v18
        default:
            var v19 float32 = 1e10
            _ = v19
        }
    }
    var v20 int8 = -128
    var v21 int = 2147483647
    var v22 int64 = (v1 << ((1) & 63))
    delete(v4, (^(v21)))
    var v23 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
