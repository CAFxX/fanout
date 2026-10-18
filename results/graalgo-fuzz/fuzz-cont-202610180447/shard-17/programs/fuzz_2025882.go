package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 bool) bool {
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!(2147483648 == 2147483647)) {
            p0 = (true && (!p0))
            fmt.Println(fi6, p0)
        } else {
            if (-2147483648 < 0) {
                p0 = p0
                var v7 float64 = -2.25
                switch "\n\t" {
                case "":
                    fmt.Println(v7)
                case "hello":
                    v8 := []int{-462, 128, (fi6 >> ((923) & 63))}
                    _ = v8
                default:
                    if (!true) {
                        v7 += v7
                        switch fi6 {
                        case 319:
                            if (p0 && true) {
                                for fi9 := 0; fi9 < 13; fi9++ {
                                    v7 = (v7 + (((((-2.25 * (v7 + (v7 - (((v7 - v7) - v7) - float64((((0.1 - (0.0 * float32(1.5))) * (-2.25 - ((1.5 / 0.1) - (0.0 - (3.14159 + (3.14159 - (-2.25 * (((-2.25 * 0.1) + 0.1) * ((0.1 / 1.0) / 3.14159))))))))) * -2.25)))))) * (float64(3.14159) - v7)) / (v7 - ((v7 * -2.25) - (-2.25 - v7)))) - v7) / v7))
                                    var v10 float64 = v7
                                    if true {
                                        v11 := []int{32768, fi6, fi9}
                                        var v12 uint8 = 2
                                        switch "hello" {
                                        case "a":
                                            v13 := GenS1{F0: -5, F1: 0}
                                            _ = v13
                                        }
                                        _ = v11
                                        _ = v12
                                    }
                                    _ = fi9
                                    _ = v10
                                }
                                p0 = p0
                            }
                        }
                        fmt.Println(p0, v7, fi6)
                    }
                }
                _ = v7
            } else {
                p0 = (-2147483648 <= -2147483648)
                var v14 float32 = 1e10
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{}
            _ = fi15
            _ = v16
        }
        var v17 uint32 = 65535
        _ = fi6
        _ = v17
    }
    var v18 uint8 = 1
    for fi19 := 0; fi19 < 8; fi19++ {
        if p0 {
            v18 += v18
        }
        v20 := []int{fi19, (fi19 & fi19), (^(fi19)), fi19}
        v18 = (v18 ^ (v18 % ((v18) | 1)))
        _ = fi19
        _ = v20
    }
    if (4294967295 <= 837) {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int16 = 42
            if true {
                fmt.Println(v22, v22, p0)
            }
            _ = fi21
            _ = v22
        }
    }
    if p0 {
        v23 := []int{}
        var v24 uint8 = v18
        _ = v23
        _ = v24
    } else {
        if (172 > 2) {
            var v25 int16 = 618
            _ = v25
        } else {
            fmt.Println(p0, v18)
            if p0 {
                p0 = false
            } else {
                v26 := GenS1{F0: 2147483647, F1: 2147483647}
                _ = v26
            }
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := []int{}
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 uint16 = 97
                v18 -= v18
                _ = fi30
                _ = v31
            }
            var v32 float32 = 3.14159
            v33 := GenS1{F0: -615, F1: 103}
            _ = fi29
            _ = v32
            _ = v33
        }
        v34 := GenS2{F0: 1, F1: -1}
        _ = v34
    }
    _ = p0
    _ = v18
    return (true && ((p0 && true) && true))
}
func genFn1(p0 bool, p1 int8, p2 uint16) int {
    var v35 int16 = -1
    var v36 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v36
    return -77
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{-9223372036854775808, 32768, 265, 128}
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 uint32 = 745
    v37 := GenS2{F0: -1, F1: 2147483647}
    fmt.Println(v5)
    _ = len(v2)
    v38 := GenS2{F0: 0, F1: 571}
    v5 += 2
    var v39 int64 = -9223372036854775808
    v40 := GenS1{F0: -854, F1: (^(v5))}
    v39 = v39
    if true {
        var v41 uint64 = 414
        _ = v41
    }
    v39 += (v39 | v39)
    var v42 string = "\n\t"
    v42 = v42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
