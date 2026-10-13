package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint64, p1 int) uint32 {
    v7 := GenS1{F0: true, F1: 79}
    var v8 uint8 = 127
    v9 := []int{-624}
    switch "a" {
    case "":
        p1 += 865
    case "\n\t":
        if (true || (2 <= (v8 / ((v8) | 1)))) {
            if (0.1 > -2.25) {
                p0 += (p0 % (((^(p0))) | 1))
            }
            var v10 uint32 = 4294967295
            p0 = uint64(v8)
            _ = v10
        }
    default:
        v11 := GenS2{F0: 65535}
        _ = v11
    }
    p1 -= 32768
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}
func genFn1() uint32 {
    var v12 uint64 = 1
    v12 -= v12
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint = 1
            fmt.Println(fi14, v15, fi14)
            var v16 uint16 = 2
            if false {
                v15 += 18446744073709551615
                if (!false) {
                    v17 := []int{}
                    var v18 uint64 = (v12 - v12)
                    switch (("\n\t" + "a") + "") {
                    case "foo bar":
                        v17 = append(v17, fi14)
                    case "\n\t":
                        if (float64(((0.1 - 1e10) * 3.14159)) == 1.5) {
                            var v19 int16 = -132
                            _ = v19
                        }
                    case "hello":
                        v18 = v12
                    default:
                        fmt.Println(v18, fi14)
                    }
                    _ = v17
                    _ = v18
                } else {
                    v15 += (v15 >> (((v15 >> ((v15) & 63))) & 63))
                    var v20 uint8 = 2
                    var v21 bool = (uint(v20) == 42)
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi14)
                if true {
                    if (!false) {
                        v22 := GenS1{F0: (false || (-1 != 42)), F1: 0}
                        v23 := []int{fi13, 4294967295, (fi13 / ((fi13) | 1))}
                        _ = v22
                        _ = v23
                    }
                    var v24 int16 = 689
                    _ = v24
                } else {
                    if false {
                        v25 := GenS3{F0: 1, F1: 255, F2: 2147483647}
                        v26 := []int{(fi14 + fi14)}
                        var v27 float32 = 1.5
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 uint8 = 200
                    _ = v28
                }
                var v29 uint64 = (^(v12))
                v12 = (v29 << ((v15) & 63))
                _ = v29
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    v12 += v12
    _ = v12
    return genFn0((v12 / (((v12 - (v12 | (v12 + (v12 / ((1) | 1)))))) | 1)), 4294967295)
}

func main() {
    v1 := []int{816, 32768, -183, -582}
    var v2 bool = false
    var v3 bool = false
    var v4 uint32 = 42
    var v5 uint64 = 0
    v6 := make(map[int]string)
    v2 = false
    v3 = (311 > 42)
    v1 = append(v1, 9223372036854775807)
    if (-1 <= 1) {
        v4 += 591
        var v30 uint64 = (v5 ^ 42)
        var v31 uint8 = 255
        _ = v30
        _ = v31
    }
    delete(v6, 32767)
    fmt.Println(v5, v5, v2)
    v5 = 171
    v32 := GenS4{F0: -128, F1: v4}
    v3 = v2
    v2 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
