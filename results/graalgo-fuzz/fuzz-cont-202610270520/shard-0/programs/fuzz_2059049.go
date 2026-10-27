package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
func genFn0(p0 int, p1 int32, p2 int32) bool {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: fi6, F1: 1000}
        v8 := GenS2{F0: 918, F1: 210}
        p0 = fi6
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = p0
    switch p0 {
    case 32767:
        p2 += p1
    case -387:
        if (-1 <= -9223372036854775808) {
            p0 -= 127
        } else {
            var v9 uint = 255
            _ = v9
        }
    case -156:
        var v10 int32 = (p1 * p1)
        _ = v10
    default:
        fmt.Println(p2, p1)
    }
    var v11 int32 = 2147483647
    v12 := []int{p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 int8 = 97
    v2 := []string{"hello", "", ("" + ("\n\t" + "\n\t"))}
    var v3 float32 = (-2.25 * ((3.14159 * 0.1) + (-2.25 / 0.1)))
    v4 := make(map[string]int)
    var v5 int32 = 42
    var v13 uint16 = 1
    v14 := GenS3{F0: true, F1: (^(v13)), F2: "foo bar"}
    fmt.Println(v5, v5)
    _ = len(v2)
    v13 = (v13 >> ((18446744073709551615) & 63))
    _ = len(v4)
    var v15 uint8 = 127
    for fi16 := 0; fi16 < 12; fi16++ {
        if (false && true) {
            var v17 uint64 = 18446744073709551615
            for fi18 := 0; fi18 < 13; fi18++ {
                v3 += ((v3 + (v3 - v3)) - (float32((float64(3.14159) - float64(v3))) * ((v3 + (v3 - float32((0.1 + 0.1)))) - (-2.25 / (v3 + ((-2.25 - v3) * (v3 / ((v3 + v3) + 0.0))))))))
                _ = fi18
            }
            _ = v17
        } else {
            v15 -= (v15 + 149)
        }
        v1 += (v1 | -88)
        _ = fi16
    }
    fmt.Println(v13, v15)
    switch -539 {
    case -9223372036854775808:
        v19 := GenS2{F0: (v5 % ((v5) | 1)), F1: 2}
        _ = v19
    case 128:
        v13 -= v13
    }
    var v20 int8 = (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
