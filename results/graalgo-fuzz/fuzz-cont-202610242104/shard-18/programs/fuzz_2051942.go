package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
func genFn0(p0 uint) bool {
    p0 -= (p0 / (((p0 | (p0 % ((65535) | 1)))) | 1))
    var v4 int64 = 591
    var v5 uint16 = 255
    _ = p0
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int32 = 326
    var v3 float32 = 1e10
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            v7 := []int{fi6, 128, (fi6 * fi6), (fi6 * fi6)}
            _ = v7
        }
        var v8 int8 = 127
        if ((v8 % ((v8) | 1)) == -13) {
            v9 := GenS1{F0: v1, F1: 100, F2: -1}
            v10 := []int{}
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v8
    }
    v3 = float32((1e10 - -2.25))
    v11 := []int{9223372036854775807, 497, -47, -21}
    var v12 uint16 = 32767
    v3 = (v3 / 0.1)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v12 -= v12
    v13 := GenS1{F0: (v1 | (v1 / (((v1 >> (((v1 * 0)) & 63))) | 1))), F1: -27, F2: 1}
    v14 := GenS2{F0: "a", F1: 231, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
