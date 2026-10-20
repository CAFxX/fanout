package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
func genFn0(p0 int16, p1 int, p2 bool) int8 {
    var v6 int16 = p0
    var v7 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64) uint8 {
    for fi8 := 0; fi8 < 2; fi8++ {
        p0 = p0
        var v9 float32 = 3.14159
        _ = fi8
        _ = v9
    }
    var v10 int16 = -1
    p0 += p0
    fmt.Println(p0, p0)
    _ = p0
    _ = v10
    return 2
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    v3 := []int{-980, 42, 32768}
    var v4 int8 = -77
    v5 := make(map[uint32]uint32)
    var v11 uint = 65535
    fmt.Println(v4, v1)
    v11 = v11
    v1 = v1
    fmt.Println(v4, v1, v4)
    switch 32767 {
    case -963:
        v11 = v11
    }
    var v12 int32 = 2147483647
    switch 128 {
    case -358:
        if (2147483648 < 686) {
            for fi13 := 0; fi13 < 17; fi13++ {
                if (float32(1.5) <= 0.1) {
                    fmt.Println(v1)
                } else {
                    if (false || false) {
                        var v14 string = ""
                        _ = v14
                    }
                    switch "" {
                    case "hello":
                        v12 -= -1
                    default:
                        var v15 uint64 = 617
                        _ = v15
                    }
                    v1 = v1
                }
                _ = len(v3)
                _ = fi13
            }
            var v16 int8 = 1
            _ = v16
        } else {
            var v17 int = -1
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(v11, v1)
            fmt.Println(v1, v4)
            _ = fi18
        }
    }
    v19 := GenS1{F0: 9223372036854775807, F1: 845, F2: 42}
    if false {
        v12 -= -2147483648
        var v20 uint32 = 4294967295
        v20 += v20
        _ = v20
    } else {
        var v21 float32 = -2.25
        switch "a" {
        case "hello":
            fmt.Println(v11)
        case "\n\t":
            var v22 bool = (!false)
            _ = v22
        default:
            v23 := GenS1{F0: 187, F1: v11, F2: (v11 * v11)}
            _ = v23
        }
        _ = v21
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v11 = (^(v11))
        _ = fi24
    }
    v1 = 861
    delete(v2, 9223372036854775807)
    v3 = append(v3, -516)
    v4 = (v4 | v4)
    v4 += (v4 % ((v4) | 1))
    var v25 int16 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
