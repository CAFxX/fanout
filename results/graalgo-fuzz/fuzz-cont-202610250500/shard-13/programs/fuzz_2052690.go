package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 165
    v3 := []int{(^(v2))}
    var v4 int32 = 358
    v5 := []int{(v2 ^ -169), v2, 128, v2}
    var v6 uint16 = 502
    if false {
        if len(v5) > 0 {
            v5[0] = v2
        }
        if true {
            switch "" {
            case "foo bar":
                switch v2 {
                case 938:
                    if false {
                        v3 = append(v3, v2)
                    }
                }
            case "a":
                v6 = v6
            }
        }
    }
    fmt.Println(v2, v6, v6)
    var v7 int16 = 42
    if (v4 > 1) {
        if ((!true) || (1 == 929)) {
            fmt.Println(v4, v6)
            v6 -= 342
        } else {
            v2 += v2
        }
        v6 += (v6 | 220)
        v2 += v2
    } else {
        var v8 uint32 = 866
        v7 -= 1000
        var v9 int8 = 0
        _ = v8
        _ = v9
    }
    var v10 int8 = 67
    v11 := GenS1{F0: int64(v7), F1: v6}
    v7 -= (v7 * (^(v7)))
    v12 := GenS1{F0: -9223372036854775808, F1: (^(v6))}
    var v13 int32 = (v4 | v4)
    v13 = (v13 + v4)
    v13 += (v4 * v13)
    fmt.Println(v10, v10, v13)
    var v14 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
