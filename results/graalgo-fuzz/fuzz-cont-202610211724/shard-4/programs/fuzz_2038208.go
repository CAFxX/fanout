package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 int16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 int16, p2 int16) uint16 {
    v5 := GenS1{F0: 858}
    switch 1 {
    case -238:
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                p0 = ((p0 + "") + (p0 + p0))
                v8 := []int{fi6, fi7, (fi6 & (fi6 + (fi6 | -2147483648)))}
                var v9 uint64 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    default:
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int32 = -1
            _ = fi10
            _ = v11
        }
    }
    switch (((((((p0 + "foo bar") + (("\n\t" + p0) + ((p0 + (("hello" + (p0 + (p0 + "a"))) + (("\n\t" + (p0 + p0)) + (p0 + p0)))) + "\n\t"))) + p0) + "foo bar") + p0) + p0) + p0) {
    case "foo bar":
        v12 := GenS2{F0: 255, F1: 0, F2: p2}
        _ = v12
    case "\n\t":
        var v13 string = "hello"
        _ = v13
    default:
        var v14 float64 = (float64(float32(3.14159)) - 1e10)
        _ = v14
    }
    v15 := []int{127, 445}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    return 0
}
func genFn1(p0 uint8, p1 int64) int64 {
    var v16 int16 = 1
    v17 := []int{2147483647, 2147483647}
    p0 -= (p0 % (((p0 | p0)) | 1))
    var v18 bool = false
    switch ("" + ("foo bar" + "\n\t")) {
    case "":
        p0 -= p0
    case "foo bar":
        if (0.1 <= 1.5) {
            if len(v17) > 0 {
                v17[0] = 128
            }
            switch "hello" {
            case "foo bar":
                if len(v17) > 0 {
                    v17[0] = 2147483647
                }
            case "":
                v16 -= (v16 * 1)
            case "a":
                v19 := GenS3{F0: -267, F1: 387}
                _ = v19
            }
        }
    default:
        var v20 float32 = ((0.0 / 1e10) / 1.5)
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return (p1 + (p1 | (p1 << ((597) & 63))))
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.1
    var v3 int = 4294967295
    v4 := make(map[int]int)
    fmt.Println(v3, v3)
    v2 += float64(float32(3.14159))
    var v21 uint = 4294967295
    v22 := GenS4{F0: genFn0("hello", 42, 0)}
    v23 := GenS2{F0: genFn0("", 898, -32768), F1: v21, F2: 1}
    var v24 uint8 = 242
    fmt.Println(v2, v3)
    var v25 int64 = 0
    v21 = v21
    switch "foo bar" {
    case "foo bar":
        v1[v3] = v3
    }
    var v26 int32 = -2147483648
    v27 := v1[v3]
    v28 := GenS2{F0: genFn0("a", 207, 768), F1: v21, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
