package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-2147483648, 676, -447, 591}
    v3 := make(map[uint32]string)
    var v4 uint = 742
    _ = len(v1)
    fmt.Println(v4)
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            switch fi6 {
            case -738:
                var v7 uint = 18446744073709551615
                _ = v7
            case 32767:
                var v8 float64 = -2.25
                v2 = append(v2, (fi6 % ((-2147483648) | 1)))
                v9 := GenS1{F0: 185}
                _ = v8
                _ = v9
            default:
                v4 = v4
            }
            v10 := GenS1{F0: 255}
            _ = len(v2)
            var v11 int64 = 0
            _ = fi6
            _ = v10
            _ = v11
        }
        v2 = append(v2, (-(fi5)))
        _ = fi5
    }
    fmt.Println(v4)
    if true {
        var v12 int32 = 42
        _ = v12
    }
    var v13 uint32 = 65535
    var v14 int = -507
    if (!true) {
        if true {
            v13 -= (v13 / ((v13) | 1))
            v15 := v1[v14]
            if ((v4 >> ((v4) & 63)) != v4) {
                v16 := GenS2{F0: (v14 + -9223372036854775808), F1: 2147483647, F2: -1}
                var v17 int32 = 42
                var v18 uint = v4
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            _ = len(v1)
            v19 := GenS2{F0: v14, F1: -98, F2: -74}
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
