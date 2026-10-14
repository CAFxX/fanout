package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 uint16, p2 string) uint8 {
    fmt.Println(p0)
    switch "a" {
    case "\n\t":
        v7 := GenS1{F0: 1, F1: 0}
        _ = v7
    case "":
        var v8 float64 = (1e10 + -2.25)
        _ = v8
    case "foo bar":
        p1 -= (p1 & (p1 * p1))
    }
    p1 += 255
    v9 := GenS2{F0: 42}
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 float64 = 0.1
    var v4 int32 = 0
    var v5 float64 = v3
    var v6 int = 0
    var v10 uint8 = 127
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 string = "foo bar"
        switch "foo bar" {
        case "hello":
            var v13 uint16 = uint16(v10)
            _ = v13
        case "":
            if (0 <= 18446744073709551615) {
                var v14 int16 = 0
                v2[fi11] = 2147483648
                _ = v14
            } else {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v10 = 127
                    var v16 int8 = -58
                    v12 = ""
                    _ = fi15
                    _ = v16
                }
            }
        case "foo bar":
            v12 = v12
        default:
            switch v12 {
            case "a":
                for fi17 := 0; fi17 < 9; fi17++ {
                    _ = len(v1)
                    var v18 int64 = 0
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 string = "a"
                _ = v19
            }
        }
        fmt.Println(v3, v5, v10)
        _ = fi11
        _ = v12
    }
    v20 := GenS1{F0: 295, F1: (v10 ^ (v10 % (((v10 / (((v10 + genFn0(244, 33, "foo bar"))) | 1))) | 1)))}
    v21 := GenS2{F0: 1}
    v22 := GenS3{F0: 905, F1: (v4 / ((v4) | 1)), F2: 65535}
    v23 := GenS4{F0: v6}
    v24 := GenS4{F0: (v6 % ((v6) | 1))}
    var v25 uint64 = 18446744073709551615
    var v26 float64 = ((v5 - ((0.0 - 1e10) - (v3 * (v3 * 0.1)))) + (v5 + v3))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
