package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 uint, p2 uint32) uint {
    if false {
        fmt.Println(p1, p0, p1)
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 255}
            for fi8 := 0; fi8 < 2; fi8++ {
                switch fi6 {
                case -362:
                    p1 -= (p1 << ((2) & 63))
                case 438:
                    var v9 int16 = -32768
                    if true {
                        p0 = ((((((p0 + "foo bar") + (p0 + "")) + p0) + p0) + p0) + p0)
                        var v10 int32 = -1
                        _ = v10
                    }
                    p1 -= p1
                    switch (fi6 >> ((p1) & 63)) {
                    case 42:
                        p2 = (p2 << (((^(p1))) & 63))
                    case 0:
                        var v11 float32 = (3.14159 * float32((float64(0.0) / ((-2.25 - 0.1) * float64(((1.5 + -2.25) + 1e10))))))
                        _ = v11
                    }
                    if (!(197 < 42)) {
                        if ((p1 & 255) > (p1 & p1)) {
                            p0 = p0
                            fmt.Println(fi6, p1, p2)
                            p0 = p0
                        }
                        switch "a" {
                        case "foo bar":
                            if (100 != 38) {
                                p2 = 2147483648
                                p1 = p1
                                v12 := []int{(fi8 ^ 42), fi6, 1}
                                _ = v12
                            }
                        default:
                            v13 := GenS2{F0: 834, F1: -879}
                            _ = v13
                        }
                        fmt.Println(v9, p1, p0)
                    } else {
                        p1 += 4294967295
                    }
                    var v14 int32 = -411
                    var v15 int = (fi8 * (fi6 >> (((^(p1))) & 63)))
                    switch ((((p0 + p0) + p0) + (("\n\t" + "foo bar") + (p0 + p0))) + "hello") {
                    case "a":
                        v16 := GenS1{F0: 0}
                        _ = v16
                    default:
                        fmt.Println(p2, p1)
                    }
                    _ = v14
                    _ = v15
                    _ = v9
                case 925:
                    var v17 uint16 = 1
                    _ = v17
                }
                _ = fi8
            }
            fmt.Println(p2, p1, p1)
            _ = fi6
            _ = v7
        }
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        p0 = p0
        v19 := GenS3{F0: fi18}
        var v20 uint32 = 0
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint = (^(p1))
    fmt.Println(v21)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return (v21 - v21)
}
func genFn1(p0 bool, p1 uint32) int16 {
    fmt.Println(p0, p1)
    var v22 string = "\n\t"
    v23 := []int{9223372036854775807, 42, 1, 2147483647}
    var v24 uint32 = (p1 % ((42) | 1))
    var v25 bool = (!false)
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return 708
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float64 = float64((float32(0.0) - (-2.25 - (0.1 / float32(0.1)))))
    v5 := make(map[uint32]int)
    if ((-596 <= 42) && (595 < 42)) {
        v1 += 127
        var v26 uint = 2
        if (42 < 313) {
            v4 = 1e10
            v4 = (1.5 / ((float64((-2.25 + -2.25)) / v4) + -2.25))
        } else {
            _ = len(v3)
            var v27 uint8 = 235
            var v28 float64 = v4
            _ = v27
            _ = v28
        }
        _ = v26
    } else {
        var v29 uint32 = 2
        _ = v29
    }
    if (!((255 < 128) || true)) {
        fmt.Println(v4, v1)
    } else {
        v30 := GenS3{F0: 32768}
        _ = v30
    }
    switch "\n\t" {
    case "foo bar":
        var v31 int32 = -666
        _ = v31
    case "a":
        var v32 int8 = 100
        _ = v32
    }
    v1 = v1
    v4 = (((-2.25 - float64((0.1 - (((-2.25 + 1e10) / 0.1) * 0.0)))) / (((v4 - (v4 * (0.0 / 1.0))) - v4) / 1e10)) / ((v4 - (0.0 * v4)) + v4))
    fmt.Println(v1)
    var v33 uint8 = 0
    var v34 int = 9223372036854775807
    v34 += 696
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
