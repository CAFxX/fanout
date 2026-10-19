package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0() int {
    v5 := []int{1, 1}
    var v6 uint16 = 42
    _ = v5
    _ = v6
    return -645
}
func genFn1(p0 string, p1 string) int8 {
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 184, F1: 0}
        v10 := GenS2{F0: 0}
        if (float64(float32(0.1)) > (1.5 * 1.5)) {
            p0 = p1
            v11 := []int{genFn0(), (fi8 << ((4294967295) & 63)), fi8, fi8}
            for fi12 := 0; fi12 < 3; fi12++ {
                p0 = "\n\t"
                if (-1 > 9223372036854775807) {
                    if false {
                        v7 += (v7 << ((42) & 63))
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v13 int8 = 42
                            _ = v13
                        default:
                            switch p0 {
                            case "a":
                                var v14 uint = 4294967295
                                _ = v14
                            case "\n\t":
                                var v15 float32 = float32(1.5)
                                _ = v15
                            default:
                                fmt.Println(fi8, p0, v7)
                            }
                        }
                    }
                    if true {
                        var v16 uint64 = 0
                        switch fi8 {
                        case 661:
                            switch p1 {
                            case "\n\t":
                                if (!(3.14159 > 3.14159)) {
                                    var v17 int64 = -9223372036854775808
                                    var v18 float32 = 3.14159
                                    v19 := GenS1{F0: (v7 ^ 127), F1: (-(v17))}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            case "a":
                                _ = len(v11)
                                v20 := GenS3{F0: 32767}
                                var v21 uint = 255
                                _ = v20
                                _ = v21
                            case "hello":
                                v22 := GenS3{F0: 1}
                                _ = v22
                            default:
                                var v23 int16 = 1
                                _ = v23
                            }
                        case 1:
                            p0 = (((p0 + "hello") + p0) + (("" + (p1 + "a")) + "hello"))
                        case 128:
                            switch genFn0() {
                            case 32767:
                                var v24 int8 = -4
                                _ = v24
                            default:
                                switch genFn0() {
                                case 0:
                                    p1 = ((("a" + (p0 + p0)) + ("foo bar" + "\n\t")) + (p1 + p1))
                                case 128:
                                    _ = len(v11)
                                case 4294967295:
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                default:
                                    var v25 float32 = (1e10 * 3.14159)
                                    p0 = (p1 + p0)
                                    _ = v25
                                }
                            }
                        default:
                            var v26 uint16 = 255
                            _ = v26
                        }
                        v7 = v7
                        _ = v16
                    }
                    if (!false) {
                        var v27 uint32 = 2
                        p0 = "hello"
                        _ = v27
                    }
                } else {
                    var v28 float64 = (1e10 - float64((0.1 + -2.25)))
                    _ = v28
                }
                _ = fi12
            }
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v29 float64 = 3.14159
    v30 := []int{4294967295, 9223372036854775807, -282}
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    _ = v7
    return 127
}

func main() {
    var v1 int16 = 300
    var v2 int16 = (v1 | -40)
    v3 := []int{4294967295, 0, -912, 2147483647}
    v4 := make(map[string]uint32)
    v2 += v1
    v2 += (v2 - 155)
    v1 += v1
    var v31 int8 = int8(v1)
    switch -9223372036854775808 {
    case 128:
        for fi32 := 0; fi32 < 12; fi32++ {
            v2 = int16(v31)
            var v33 uint64 = 4294967295
            _ = fi32
            _ = v33
        }
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    }
    fmt.Println(v2, v31, v2)
    v34 := GenS4{F0: 0, F1: -212}
    v2 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
