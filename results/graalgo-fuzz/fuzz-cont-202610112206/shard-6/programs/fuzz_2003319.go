package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 string) int16 {
    var v4 string = (p0 + (p0 + "a"))
    if false {
        switch 127 {
        case 32767:
            fmt.Println(v4)
        case 873:
            v5 := []int{2147483647, 127}
            _ = v5
        }
        if false {
            var v6 uint64 = 1
            _ = v6
        } else {
            v4 = p0
            var v7 uint32 = 19
            var v8 uint64 = 18446744073709551615
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 463, F1: 180}
        _ = v9
    } else {
        v10 := []int{42}
        _ = v10
    }
    _ = p0
    _ = v4
    return 623
}
func genFn1(p0 uint64) int {
    var v11 float64 = ((1.5 - float64((1.5 - float32(1e10)))) - 1.5)
    fmt.Println(v11, p0)
    if false {
        if (true && false) {
            fmt.Println(p0, p0)
        }
        var v12 int16 = -386
        _ = v12
    } else {
        v13 := []int{127, 1, 127, 410}
        var v14 int16 = 671
        _ = v13
        _ = v14
    }
    v15 := []int{4294967295, 9223372036854775807, -9223372036854775808, -573}
    _ = p0
    _ = v11
    _ = v15
    return 0
}

func main() {
    var v1 string = "hello"
    var v2 uint = 255
    v3 := make(map[uint32]string)
    v16 := GenS1{F0: 0, F1: 2147483647}
    v17 := []int{-87}
    var v18 int32 = -2147483648
    if false {
        switch -1 {
        case 32768:
            _ = len(v3)
        case 32767:
            if true {
                v18 -= -2147483648
                fmt.Println(v2, v2)
                var v19 uint32 = 241
                _ = v19
            }
        case 0:
            if true {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v18 -= v18
                    switch genFn1(2) {
                    case -1:
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                    _ = fi20
                }
                switch 32768 {
                case 32768:
                    if (!true) {
                        if (1 >= (v2 >> ((v2) & 63))) {
                            v21 := GenS1{F0: uint64(v2), F1: 2147483647}
                            for fi22 := 0; fi22 < 8; fi22++ {
                                if true {
                                    for fi23 := 0; fi23 < 14; fi23++ {
                                        var v24 float32 = (-2.25 - 1.5)
                                        var v25 uint = (^(v2))
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    for fi26 := 0; fi26 < 16; fi26++ {
                                        v27 := GenS2{F0: 176, F1: 0}
                                        fmt.Println(fi26, fi22, v1)
                                        _ = fi26
                                        _ = v27
                                    }
                                    fmt.Println(v1, fi22, v18)
                                }
                                var v28 uint8 = uint8(v2)
                                _ = fi22
                                _ = v28
                            }
                            var v29 uint16 = 255
                            _ = v21
                            _ = v29
                        } else {
                            delete(v3, 0)
                        }
                        _ = len(v17)
                        v2 += (v2 ^ (v2 + 495))
                    }
                case -818:
                    fmt.Println(v1, v18, v1)
                case 32767:
                    if (!false) {
                        v30 := GenS3{F0: (!(false && (false && (uint8(v2) > uint8(v2))))), F1: -9223372036854775808}
                        _ = v30
                    }
                default:
                    if len(v17) > 0 {
                        v17[0] = genFn1(610)
                    }
                }
                _ = len(v3)
            }
        default:
            var v31 uint64 = 2
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 float32 = (1e10 + -2.25)
        var v34 bool = false
        var v35 string = v1
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println(v2, v2)
    switch genFn1(1) {
    case 450:
        v18 = v18
    case 127:
        var v36 float64 = 0.1
        _ = v36
    case 32767:
        v37 := GenS3{F0: false, F1: 127}
        _ = v37
    }
    v2 -= v2
    var v38 float64 = (0.1 * ((((((1.5 + 1e10) * ((1.5 - 3.14159) * ((0.1 * 0.0) - -2.25))) * 0.1) - 3.14159) / (float64(((3.14159 + float32((((3.14159 - float64(float32((0.0 * (float64(0.0) * 0.1))))) - float64(((float32((0.1 / -2.25)) + ((0.0 + (((1.5 + (0.1 + 0.0)) - 1.5) / (((1e10 / 3.14159) * ((3.14159 * 3.14159) / 3.14159)) - 0.0))) / (((1.5 / -2.25) * (3.14159 * ((float32(1e10) / 3.14159) / 0.1))) - 0.0))) + float32(float64((0.1 / 3.14159)))))) + 1e10))) / 1.0)) - (0.1 * ((3.14159 * 0.1) - 0.1)))) * 1e10))
    fmt.Println(v18, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
}
