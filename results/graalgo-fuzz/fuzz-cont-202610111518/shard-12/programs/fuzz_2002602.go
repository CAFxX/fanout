package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint32) uint {
    if (false && false) {
        v6 := GenS1{F0: 118}
        _ = v6
    } else {
        p0 = 1
    }
    var v7 uint32 = 0
    p1 += (v7 % ((1) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 503
}

func main() {
    var v1 int8 = 100
    var v2 float32 = 0.0
    var v3 uint16 = 65535
    v4 := make(map[uint32]uint32)
    var v5 uint = 722
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int32 = 1
        _ = fi8
        _ = v9
    }
    var v10 int16 = 0
    v10 -= (v10 + (v10 | v10))
    v5 = 707
    fmt.Println(v5, v2)
    var v11 int32 = -712
    switch "foo bar" {
    case "a":
        var v12 uint32 = 2
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint32 = 4294967295
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    if ((v14 ^ v14) < v14) {
                        _ = len(v4)
                        var v16 uint16 = 255
                        _ = v16
                    }
                    switch "a" {
                    case "\n\t":
                        v17 := GenS2{F0: 255, F1: 463}
                        _ = v17
                    case "":
                        if (float32(0.1) > ((1e10 * 3.14159) - 3.14159)) {
                            v18 := GenS3{F0: v1, F1: v3}
                            fmt.Println(v2, v10, fi13)
                            if true {
                                if (-999 < 747) {
                                    v5 += v5
                                    var v19 uint32 = (v14 | 2147483647)
                                    v20 := []int{fi13, 2147483647, -18}
                                    _ = v19
                                    _ = v20
                                }
                                v21 := GenS4{F0: ("\n\t" + ""), F1: (v11 & v11)}
                                _ = v21
                            }
                            if false {
                                switch "" {
                                case "hello":
                                    v22 := GenS5{F0: (v14 & v14)}
                                    _ = v22
                                }
                                var v23 bool = true
                                _ = v23
                            } else {
                                delete(v4, 2147483647)
                            }
                            _ = v18
                        }
                    default:
                        var v24 float64 = ((((0.1 - (3.14159 - (((3.14159 * 0.0) - float64(float32(0.1))) * -2.25))) * (((float64(v2) + 1.5) / (((1e10 + (float64(1e10) * 3.14159)) / 3.14159) / -2.25)) + (-2.25 - (float64(v2) + 1e10)))) - 1.5) * (float64(float32((1.5 * 3.14159))) / 3.14159))
                        _ = v24
                    }
                    if (0 == (v14 & 2147483647)) {
                        v25 := GenS1{F0: 453}
                        _ = v25
                    }
                } else {
                    var v26 uint8 = 128
                    fmt.Println(fi15, fi15)
                    v27 := v4[(v14 | (v14 * uint32(v26)))]
                    _ = v26
                    _ = v27
                }
                fmt.Println(v1)
                v14 = (^(v14))
                _ = fi15
            }
            var v28 int32 = v11
            _ = fi13
            _ = v14
            _ = v28
        }
    default:
        v2 -= ((0.1 + float32(0.0)) / -2.25)
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS6{F0: 0}
        switch "hello" {
        case "\n\t":
            v31 := []int{32767, (^(fi29)), fi29}
            _ = v31
        case "a":
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 bool = true
                if false {
                    var v34 uint = v5
                    _ = v34
                }
                v5 = (^(v5))
                _ = fi32
                _ = v33
            }
        case "":
            delete(v4, 319)
        default:
            fmt.Println(v11)
        }
        _ = fi29
        _ = v30
    }
    v2 -= float32(float64(v2))
    fmt.Println(v3, v1, v3)
    var v35 bool = true
    v36 := GenS6{F0: v1}
    var v37 int64 = int64(v5)
    fmt.Println(v35)
    switch "hello" {
    case "\n\t":
        v3 = 607
    default:
        var v38 int8 = v1
        _ = v38
    }
    var v39 uint16 = v3
    var v40 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
