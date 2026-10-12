package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
func genFn0() string {
    v4 := []int{-9223372036854775808}
    _ = len(v4)
    v4 = append(v4, 979)
    _ = v4
    return "\n\t"
}
func genFn1() uint16 {
    switch 32767 {
    case -504:
        var v5 bool = false
        _ = v5
    default:
        if true {
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := []int{(fi7 | (fi6 >> ((702) & 63))), fi7, 354}
                    v9 := GenS1{F0: -128}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = fi6
            }
            v10 := GenS2{F0: -322}
            _ = v10
        } else {
            v11 := GenS1{F0: -128}
            var v12 float64 = 1e10
            v12 += (((v12 - float64(1.5)) / ((float64((1.5 / (0.1 / -2.25))) * v12) / (0.1 + v12))) * (-2.25 - v12))
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 uint = 443
        v14 = (v14 | v14)
        v14 = v14
        _ = fi13
        _ = v14
    }
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 float32 = ((float32((3.14159 + float64(0.1))) - -2.25) - 3.14159)
    var v15 uint32 = 660
    v16 := GenS3{F0: 0, F1: -863}
    v15 += (v15 * v15)
    _ = len(v2)
    v17 := GenS3{F0: 100, F1: -1}
    var v18 uint8 = 0
    var v19 float32 = v3
    switch 32767 {
    case -9223372036854775808:
        v18 = v18
    }
    v20 := []int{127, 0, 233, 127}
    v21 := GenS4{F0: true, F1: 42, F2: 32768}
    v19 = 3.14159
    if len(v20) > 0 {
        v20[0] = 128
    }
    v22 := GenS3{F0: -1, F1: 32767}
    delete(v2, 0)
    if (65535 >= 32767) {
        v23 := GenS2{F0: -2147483648}
        for fi24 := 0; fi24 < 15; fi24++ {
            v3 += (1e10 * (v3 / 1.0))
            v15 = 4294967295
            v18 = (v18 ^ v18)
            _ = fi24
        }
        var v25 uint8 = v18
        _ = v23
        _ = v25
    }
    var v26 uint32 = (v15 / (((^(v15))) | 1))
    v26 = 839
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
}
