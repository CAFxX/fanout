package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 uint64, p2 uint64) int64 {
    var v4 int8 = 0
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 bool = true
    v2 := []string{("" + "")}
    v3 := []int{895, 0}
    fmt.Println(v1)
    v1 = v1
    switch "" {
    case "foo bar":
        v1 = (65535 <= 1)
    case "hello":
        var v5 int = 1
        _ = v5
    case "\n\t":
        var v6 bool = v1
        _ = v6
    }
    var v7 uint8 = 1
    var v8 uint16 = uint16(v7)
    v9 := GenS1{F0: uint16(v7), F1: -1, F2: v8}
    v10 := GenS1{F0: v8, F1: -10, F2: 353}
    if v1 {
        switch 128 {
        case 128:
            v11 := GenS2{F0: 42}
            _ = v11
        case 0:
            v1 = ((0.0 * 1.5) > float64(1.5))
        default:
            var v12 int8 = 100
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 684, F1: 0, F2: (v8 - (v8 - v8))}
        switch -253 {
        case 452:
            v14 := GenS2{F0: 18446744073709551615}
            _ = v14
        case 42:
            v1 = (!v1)
        default:
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int32 = 183
                var v17 string = "a"
                v8 -= v8
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = v13
    }
    v7 -= 12
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
