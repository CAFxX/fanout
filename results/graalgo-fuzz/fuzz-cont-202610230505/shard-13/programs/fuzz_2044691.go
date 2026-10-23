package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int16 }
func genFn0(p0 uint8) uint {
    p0 -= p0
    v4 := GenS1{F0: 1, F1: false}
    switch "" {
    case "foo bar":
        var v5 float32 = 0.0
        _ = v5
    case "hello":
        if (p0 <= 1) {
            var v6 uint32 = 215
            if false {
                for fi7 := 0; fi7 < 12; fi7++ {
                    v6 = v6
                    _ = fi7
                }
                v6 = uint32(p0)
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        p0 = p0
                    case "hello":
                        v6 -= (v6 | v6)
                    case "":
                        for fi8 := 0; fi8 < 17; fi8++ {
                            fmt.Println(v6, v6, v6)
                            fmt.Println(p0)
                            _ = fi8
                        }
                    }
                    v9 := GenS1{F0: (v6 << ((4294967295) & 63)), F1: true}
                    _ = v9
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v10 int64 = -9223372036854775808
                        _ = v10
                    }
                }
            } else {
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS1{F0: (^(v6)), F1: ((p0 * p0) > p0)}
                    _ = fi11
                    _ = v12
                }
                v13 := GenS2{F0: -1}
                fmt.Println(p0, p0)
                _ = v13
            }
            _ = v6
        }
    }
    v14 := []int{-1, 1}
    _ = p0
    _ = v14
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    v15 := GenS3{F0: true, F1: 1, F2: 1}
    v2 = (((((v2 + 0.0) / ((1e10 + v1) / v1)) - v1) - 3.14159) / (v1 / v1))
    var v16 bool = (255 == 1)
    v17 := []int{127, 42}
    if true {
        v1 += -2.25
        var v18 int = 0
        _ = v18
    } else {
        if (1000 < 658) {
            v16 = (-1 < -72)
        }
    }
    var v19 int8 = 57
    var v20 int32 = 0
    switch ("hello" + ("hello" + "hello")) {
    case "foo bar":
        v21 := GenS4{F0: (!v16), F1: 2, F2: -904}
        _ = v21
    default:
        var v22 int32 = v20
        _ = v22
    }
    v19 += 100
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
