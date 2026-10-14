package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    var v3 uint16 = 1
    if false {
        for fi4 := 0; fi4 < 3; fi4++ {
            v3 = v3
            var v5 float64 = 3.14159
            _ = fi4
            _ = v5
        }
        var v6 bool = false
        for fi7 := 0; fi7 < 1; fi7++ {
            if (131 > 67) {
                v6 = false
            } else {
                var v8 uint64 = 18446744073709551615
                v9 := GenS1{F0: 42}
                _ = v8
                _ = v9
            }
            v6 = (3.14159 != 3.14159)
            fmt.Println(fi7, v6, v6)
            _ = fi7
        }
        _ = v6
    }
    var v10 int16 = 429
    v11 := []int{127}
    var v12 uint16 = (v3 | 42)
    switch int(v10) {
    case 503:
        switch "\n\t" {
        case "hello":
            var v13 uint64 = 245
            _ = v13
        default:
            if (((v12 ^ (^(v3))) <= (v12 * 778)) || (1 != v12)) {
                var v14 bool = false
                _ = v14
            } else {
                switch 1 {
                case -9223372036854775808:
                    v2 = (v2 % (((v2 % (((v2 / ((42) | 1))) | 1))) | 1))
                case 127:
                    v2 += (v2 << ((42) & 63))
                case 445:
                    v10 += -974
                }
                if ((1.5 - 0.0) > (1.5 - 0.1)) {
                    v10 = v10
                }
                delete(v1, 2147483647)
            }
        }
    case 128:
        fmt.Println(v10, v2, v3)
    default:
        v2 -= -2147483648
    }
    v10 += v10
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint8 = 44
        var v17 uint = 42
        _ = fi15
        _ = v16
        _ = v17
    }
    v3 -= (v3 / (((v3 % ((1) | 1))) | 1))
    var v18 int8 = -84
    var v19 uint16 = v3
    var v20 uint64 = 4294967295
    if len(v11) > 0 {
        v11[0] = 32768
    }
    v21 := GenS2{F0: "hello", F1: 471, F2: 933}
    v18 = v18
    var v22 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
