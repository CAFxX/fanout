package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint, p1 bool, p2 int32) int8 {
    p2 = -1
    p0 = (p0 - p0)
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() int64 {
    var v7 int64 = 24
    v7 = (^(v7))
    v8 := GenS1{F0: 32768, F1: 2147483647}
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    v3 := []int{781, 831, -1}
    var v4 string = "hello"
    var v5 float64 = 1.5
    v6 := []int{-380, 32767, 1, -824}
    var v9 float32 = float32(float64(((-2.25 / 1.0) * 1e10)))
    v6 = append(v6, 127)
    fmt.Println(v4, v4)
    v10 := GenS2{F0: 1}
    v11 := GenS3{F0: 4294967295, F1: true, F2: 1}
    _ = len(v6)
    var v12 uint = 255
    fmt.Println(v4, v5, v9)
    v3 = append(v3, 106)
    v9 = float32(0.0)
    var v13 int32 = 1
    v4 = v4
    v1[v4] = (true || false)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v9 -= v9
            _ = fi15
        }
        _ = fi14
    }
    v5 = v5
    if true {
        _ = len(v1)
        switch 127 {
        case -9223372036854775808:
            if (true || (((!false) || (3.14159 > v5)) && true)) {
                if (9223372036854775807 < 123) {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 uint32 = 87
                        _ = fi16
                        _ = v17
                    }
                    v12 += v12
                    if false {
                        if (!false) {
                            var v18 float64 = v5
                            var v19 uint = v12
                            _ = v18
                            _ = v19
                        }
                    }
                } else {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 string = v4
                        _ = fi20
                        _ = v21
                    }
                }
                v9 = ((((0.0 / 1.0) + (v9 + (float32((v5 + v5)) / ((1e10 / v9) * 0.0)))) - 3.14159) / 3.14159)
            }
        case 0:
            var v22 int8 = 1
            _ = v22
        default:
            fmt.Println(v9)
        }
    } else {
        if (true || (!true)) {
            for fi23 := 0; fi23 < 3; fi23++ {
                if true {
                    fmt.Println(v9)
                }
                switch 0 {
                case 648:
                    v3 = append(v3, fi23)
                case 1:
                    v9 = v9
                default:
                    if (42 > 32767) {
                        switch "" {
                        case "a":
                            var v24 uint32 = 1
                            _ = v24
                        }
                        if false {
                            v12 -= 255
                            fmt.Println(v9)
                            var v25 int64 = -1
                            _ = v25
                        }
                        v26 := v1[v4]
                        _ = v26
                    } else {
                        var v27 uint8 = 128
                        v28 := GenS4{F0: 0}
                        _ = v27
                        _ = v28
                    }
                }
                var v29 int8 = -1
                _ = fi23
                _ = v29
            }
        } else {
            delete(v1, v4)
        }
        var v30 bool = (32767 < -1)
        v31 := GenS3{F0: 2, F1: v30, F2: 122}
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
