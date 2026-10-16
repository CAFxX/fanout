package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 uint }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0(p0 bool, p1 uint16) int64 {
    if (4294967295 != 127) {
        switch "hello" {
        case "foo bar":
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 string = "\n\t"
                switch "a" {
                case "hello":
                    var v7 uint64 = 656
                    _ = v7
                case "a":
                    for fi8 := 0; fi8 < 19; fi8++ {
                        if p0 {
                            p1 -= 327
                        } else {
                            if (0 <= 1) {
                                v9 := []int{(fi8 * fi5), (fi5 | (fi5 << ((18446744073709551615) & 63))), 241}
                                p1 = (p1 - p1)
                                fmt.Println(fi8, fi8)
                                p1 = p1
                                _ = v9
                            }
                            p1 = p1
                        }
                        _ = fi8
                    }
                case "foo bar":
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v11 := []int{(fi10 / (((fi10 >> ((314) & 63))) | 1)), -961}
                        _ = fi10
                        _ = v11
                    }
                }
                _ = fi5
                _ = v6
            }
        case "\n\t":
            fmt.Println(p1, p0)
        case "hello":
            for fi12 := 0; fi12 < 18; fi12++ {
                if ((uint(p1) != 0) && true) {
                    var v13 float32 = 0.0
                    v14 := []int{fi12, 128}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int16 = 0
                    p1 -= (^(p1))
                    _ = v15
                }
                _ = fi12
            }
        default:
            p0 = p0
        }
        fmt.Println(p0)
        fmt.Println(p1, p1)
    }
    p0 = (234 < 42)
    _ = p0
    _ = p1
    return 893
}
func genFn1(p0 string, p1 uint, p2 uint) int64 {
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint = (p2 & 18446744073709551615)
        var v18 uint = p1
        _ = fi16
        _ = v17
        _ = v18
    }
    p0 = ""
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: "foo bar", F1: 128, F2: -83}
        p1 += (^(p2))
        _ = fi19
        _ = v20
    }
    p1 -= (p1 - 255)
    v21 := GenS2{F0: -583, F1: 279, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 267
}

func main() {
    var v1 float32 = 3.14159
    var v2 string = "a"
    v3 := []string{v2, v2}
    var v4 int32 = -220
    fmt.Println(v1, v4, v1)
    if (((float64(-2.25) * float64(((float32(3.14159) + (1e10 * v1)) - float32(0.0)))) / (((((float64((v1 - 1.5)) * (((0.1 + 0.1) * 3.14159) * float64(((1.5 * ((v1 / (v1 + 0.1)) * 3.14159)) / v1)))) * 1e10) / float64(((float32(1.5) * (1e10 * v1)) - 3.14159))) / float64(1e10)) * (1.5 * 3.14159))) == 1e10) {
        fmt.Println(v1, v2, v4)
    }
    var v22 float64 = ((0.1 * 1e10) + 3.14159)
    var v23 uint64 = 62
    if ((v4 | v4) == v4) {
        v4 = 992
        v22 = (v22 / ((0.0 + v22) * v22))
        for fi24 := 0; fi24 < 15; fi24++ {
            if (false && false) {
                var v25 int64 = 42
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v3)
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 uint = 18446744073709551615
                _ = fi26
                _ = v27
            }
            _ = fi24
        }
    } else {
        var v28 int16 = 92
        _ = v28
    }
    fmt.Println(v22, v4)
    v29 := GenS3{F0: 127, F1: 362, F2: 42}
    var v30 int32 = (v4 % ((v4) | 1))
    if ((v22 / v22) >= 1.5) {
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS2{F0: -98, F1: -9223372036854775808, F2: 1}
            _ = fi31
            _ = v32
        }
        if false {
            if (2147483647 != v4) {
                v33 := GenS4{F0: (v23 ^ 2), F1: (v30 / (((v4 >> ((1) & 63))) | 1))}
                _ = v33
            } else {
                v34 := GenS5{F0: 42, F1: 127, F2: (v30 - v30)}
                _ = v34
            }
        }
        v22 += float64(3.14159)
    }
    var v35 uint32 = 1
    var v36 int16 = 514
    var v37 int32 = (v30 / (((v30 ^ (v4 % ((v4) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
