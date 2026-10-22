package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 int16, p1 string) uint16 {
    if (!(598 <= 348)) {
        v6 := GenS1{F0: 248, F1: 0, F2: 1}
        var v7 uint8 = 12
        p0 += 714
        _ = v6
        _ = v7
    } else {
        var v8 float64 = (0.0 * 0.1)
        _ = v8
    }
    v9 := GenS1{F0: 94, F1: -192, F2: 255}
    v10 := []int{1}
    var v11 uint16 = 40
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (v11 * (v11 / ((255) | 1)))
}
func genFn1(p0 int, p1 uint16) int64 {
    p1 = (p1 - (p1 ^ 242))
    switch (p0 + 975) {
    case 2147483647:
        v12 := []int{p0}
        _ = v12
    case 554:
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS2{F0: -2147483648, F1: 1, F2: -2147483648}
            _ = fi13
            _ = v14
        }
    }
    var v15 uint8 = 0
    var v16 int8 = 42
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    v1 := []string{(("hello" + "hello") + "foo bar")}
    v2 := []string{("hello" + "a"), ((("\n\t" + "") + "foo bar") + (("hello" + "\n\t") + "foo bar")), "hello"}
    v3 := []int{-1, 127, 9223372036854775807, 1}
    var v4 int64 = 42
    v5 := make(map[uint32]int)
    var v17 uint64 = 18446744073709551615
    v18 := GenS3{F0: -1, F1: -1, F2: 65535}
    fmt.Println(v4, v17)
    v17 = (v17 / (((v17 % (((^(v17))) | 1))) | 1))
    var v19 float64 = (-2.25 - (float64(0.0) / 1.5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS4{F0: 0, F1: 1000}
    var v21 int8 = -112
    v19 -= v19
    if (908 >= 868) {
        if (0.0 <= (1e10 / (float32(v19) * 0.1))) {
            switch "foo bar" {
            case "\n\t":
                if (!false) {
                    var v22 int32 = 340
                    _ = v22
                } else {
                    v23 := GenS5{F0: ("" + "foo bar"), F1: 1}
                    _ = v23
                }
            }
        }
    }
    v3 = append(v3, 655)
    v24 := GenS1{F0: 224, F1: 42, F2: 0}
    switch "\n\t" {
    case "":
        v4 = genFn1(-763, 42)
    case "foo bar":
        v19 -= v19
    }
    var v25 int64 = 42
    var v26 bool = (v17 > (v17 - 80))
    v26 = true
    v21 += 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
