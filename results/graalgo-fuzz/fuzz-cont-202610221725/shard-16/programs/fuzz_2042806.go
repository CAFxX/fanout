package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint) string {
    var v7 int16 = -354
    v7 -= v7
    var v8 uint64 = 214
    if ((v8 << (((p0 - p0)) & 63)) > 0) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v7 -= (v7 & v7)
            v10 := GenS1{F0: fi9, F1: 1, F2: (fi9 | fi9)}
            var v11 uint = 453
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return ("" + "\n\t")
}
func genFn1() uint {
    var v12 int16 = 602
    var v13 float32 = 1.5
    v14 := []int{32768, -9223372036854775808, 42, 2147483647}
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(fi15, v12, v12)
        _ = len(v14)
        _ = fi15
    }
    _ = v12
    _ = v13
    _ = v14
    return 6
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 473
    var v3 float32 = float32(float64(3.14159))
    v4 := make(map[string]string)
    var v5 uint8 = v1
    var v6 float64 = 0.0
    v16 := GenS1{F0: int(v2), F1: -1, F2: 9223372036854775807}
    v17 := GenS2{F0: v1, F1: 42, F2: 2}
    var v18 float64 = ((v6 * float64(v3)) * 1e10)
    var v19 int16 = -32768
    v19 += (v19 % (((v19 / (((v19 - (v19 * v19))) | 1))) | 1))
    var v20 float64 = (v18 - -2.25)
    v21 := GenS3{F0: "foo bar", F1: 2, F2: -100}
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
