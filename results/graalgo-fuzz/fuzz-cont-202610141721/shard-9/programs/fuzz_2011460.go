package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 int32; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 uint32, p2 uint8) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 += -23
        _ = fi6
    }
    var v7 int = 32767
    var v8 uint16 = 42
    if (-2147483648 != int32(p0)) {
        var v9 uint8 = p2
        v10 := GenS1{F0: "a", F1: 42}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 101
}
func genFn1(p0 bool, p1 int8) uint64 {
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := []int{32768}
        _ = fi11
        _ = v12
    }
    var v13 int16 = 0
    v14 := []int{9223372036854775807, 4294967295, -1, 4294967295}
    for fi15 := 0; fi15 < 15; fi15++ {
        v14 = append(v14, int(p1))
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 700
}

func main() {
    v1 := []uint32{1, 4294967295, 65535, 2147483647}
    v2 := make(map[int]string)
    var v3 float32 = (float32(3.14159) - 0.1)
    var v4 uint32 = 642
    v5 := make(map[string]string)
    var v16 uint8 = 149
    var v17 string = ""
    v1 = append(v1, 192)
    switch "" {
    case "":
        var v18 float32 = -2.25
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 1; fi19++ {
            v4 = 1
            _ = fi19
        }
    default:
        var v20 int = -356
        _ = v20
    }
    _ = len(v1)
    if false {
        v4 += v4
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            if (0 >= 433) {
                v4 = (v4 * 0)
                var v22 uint32 = 4294967295
                _ = v22
            } else {
                v17 = ""
                var v23 bool = false
                _ = v23
            }
            if len(v1) > 0 {
                v1[0] = 2
            }
            _ = fi21
        }
        var v24 uint32 = v4
        v4 += v4
        _ = v24
    }
    v25 := GenS2{F0: 32768, F1: 888, F2: 872}
    var v26 int8 = 19
    switch v17 {
    case "hello":
        for fi27 := 0; fi27 < 14; fi27++ {
            if true {
                fmt.Println(v16)
            } else {
                var v28 uint = 18446744073709551615
                v29 := GenS3{F0: -1, F1: 255, F2: -1}
                _ = v28
                _ = v29
            }
            _ = fi27
        }
    case "":
        var v30 uint8 = 78
        _ = v30
    default:
        for fi31 := 0; fi31 < 7; fi31++ {
            if true {
                _ = len(v5)
            } else {
                v26 -= (^(v26))
                v32 := GenS1{F0: "\n\t", F1: (v4 / ((v4) | 1))}
                if false {
                    fmt.Println(fi31, v4)
                }
                _ = v32
            }
            v16 = v16
            _ = fi31
        }
    }
    if ((1e10 * (((1e10 + (float64(3.14159) * 0.1)) * float64(v3)) - (3.14159 / ((1.5 + float64((0.1 + (((v3 + (v3 - (0.0 * -2.25))) / float32(0.1)) + 0.0)))) + 1.5)))) != (float64((1e10 - ((((3.14159 - v3) + ((v3 - v3) - (((-2.25 * (v3 * v3)) / ((-2.25 + 1e10) / (1e10 + ((v3 + ((1.5 - (0.1 * 0.1)) + v3)) * v3)))) * float32(3.14159)))) + (v3 / -2.25)) / v3))) + 0.1)) {
        v33 := GenS3{F0: -1, F1: v16, F2: 2147483647}
        v5["a"] = ("a" + v17)
        _ = v33
    } else {
        var v34 int32 = 2147483647
        switch (v17 + v17) {
        case "foo bar":
            var v35 uint16 = 255
            _ = v35
        case "a":
            v17 = v17
        case "":
            if (65535 <= 255) {
                v36 := GenS4{F0: v16, F1: (v17 + "a"), F2: genFn1(((true || true) || (v17 < (v17 + v17))), -1)}
                _ = v36
            } else {
                var v37 float64 = float64(((1.5 + 3.14159) / (v3 + (float32((0.0 + 1e10)) + (v3 / 3.14159)))))
                _ = v37
            }
        default:
            delete(v5, v17)
        }
        if (0 != 0) {
            for fi38 := 0; fi38 < 17; fi38++ {
                for fi39 := 0; fi39 < 12; fi39++ {
                    v40 := v2[-670]
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
            v41 := GenS5{F0: 65535, F1: 1}
            _ = len(v2)
            _ = v41
        }
        _ = v34
    }
    v3 = -2.25
    if false {
        v42 := GenS1{F0: "foo bar", F1: (v4 / (((v4 % ((v4) | 1))) | 1))}
        if ((v3 - v3) < v3) {
            var v43 uint16 = 42
            _ = v43
        } else {
            if false {
                fmt.Println(v16, v3)
                var v44 float64 = (0.0 + 1.5)
                v16 = (v16 / (((v16 & v16)) | 1))
                _ = v44
            } else {
                var v45 string = v17
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 uint8 = v16
                    switch (fi46 & (fi46 >> ((uint(v47)) & 63))) {
                    case 0:
                        for fi48 := 0; fi48 < 17; fi48++ {
                            v49 := GenS6{F0: -1}
                            _ = fi48
                            _ = v49
                        }
                    default:
                        v50 := GenS1{F0: v17, F1: 486}
                        _ = v50
                    }
                    _ = fi46
                    _ = v47
                }
                _ = v45
            }
            v26 += (-(v26))
            v17 = "\n\t"
        }
        var v51 int16 = 747
        _ = v42
        _ = v51
    }
    var v52 int32 = 550
    var v53 int8 = (v26 % ((genFn0((v26 / (((v26 << ((255) & 63))) | 1)), v4, v16)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
