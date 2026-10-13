package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64, p1 uint, p2 int16) int32 {
    p0 += (p0 / ((p0) | 1))
    var v6 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}
func genFn1() string {
    var v7 float32 = 1e10
    v7 += 1e10
    var v8 float64 = (3.14159 / 1.5)
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := []uint32{12, 42, 42, 2147483648}
    v2 := []int{128, 4294967295, 2147483647}
    v3 := make(map[int]bool)
    var v4 uint = 42
    v5 := []string{"a", ("foo bar" + ("a" + "")), ("foo bar" + "hello")}
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: 1, F1: (228 == -293)}
        _ = fi10
        _ = v11
    }
    if (true && (!true)) {
        var v12 uint = (^(v4))
        _ = v12
    } else {
        v13 := v3[42]
        v13 = v13
        _ = v13
    }
    v4 = 730
    if (3.14159 <= (0.1 - ((((1e10 / ((3.14159 / 1e10) - 0.0)) - 0.0) * 3.14159) + (-2.25 / 1.0)))) {
        fmt.Println(v4)
        v14 := GenS2{F0: -1, F1: 42}
        _ = v14
    }
    v15 := GenS2{F0: 306, F1: 255}
    v4 = v4
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: -2147483648, F1: false}
        fmt.Println(fi16, v4)
        _ = fi16
        _ = v17
    }
    var v18 float64 = 0.1
    var v19 float64 = v18
    var v20 float64 = float64(float32(float64(((0.1 / (0.1 + (0.1 / (0.1 + 3.14159)))) * (1.5 / 1e10)))))
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint8 = 127
            _ = fi22
            _ = v23
        }
        fmt.Println(v19)
        _ = fi21
    }
    v19 = v19
    var v24 uint16 = 818
    if ((3.14159 > 1.5) && (42 < 2)) {
        fmt.Println(v20)
        v25 := GenS3{F0: v24}
        _ = v25
    } else {
        var v26 string = ("foo bar" + "foo bar")
        if (false || false) {
            v27 := GenS4{F0: 0}
            if ((0.0 * 1.5) == ((0.0 / 1e10) / -2.25)) {
                var v28 float64 = v20
                _ = v28
            } else {
                var v29 uint = 274
                v30 := GenS2{F0: 0, F1: v24}
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                switch -2147483648 {
                case -1:
                    if len(v2) > 0 {
                        v2[0] = fi31
                    }
                case 0:
                    if ((fi31 * fi31) < (fi31 >> ((249) & 63))) {
                        if (1000 < 0) {
                            var v32 uint32 = 2147483647
                            v24 -= v24
                            _ = v32
                        } else {
                            var v33 uint8 = 1
                            v18 = (v19 - v18)
                            _ = v33
                        }
                    } else {
                        v34 := GenS2{F0: 310, F1: 0}
                        _ = v34
                    }
                default:
                    v1 = append(v1, 2)
                }
                fmt.Println(v26, v18, v4)
                for fi35 := 0; fi35 < 9; fi35++ {
                    _ = len(v1)
                    var v36 int8 = 127
                    v19 -= v19
                    if true {
                        if ((!(39 != 0)) && (4294967295 == 936)) {
                            fmt.Println(v26, v18, v26)
                            if len(v2) > 0 {
                                v2[0] = fi31
                            }
                        }
                    } else {
                        var v37 int8 = v36
                        _ = v37
                    }
                    _ = fi35
                    _ = v36
                }
                _ = fi31
            }
            _ = v27
        }
        _ = v26
    }
    v38 := GenS2{F0: 253, F1: (v24 % (((v24 ^ (v24 ^ v24))) | 1))}
    if true {
        switch 4294967295 {
        case 9223372036854775807:
            var v39 uint = (v4 * (v4 / ((v4) | 1)))
            _ = v39
        case 575:
            var v40 uint32 = 2
            _ = v40
        default:
            fmt.Println(v20)
        }
        var v41 float64 = (v18 / float64((0.0 - 3.14159)))
        _ = v41
    } else {
        var v42 uint8 = 127
        var v43 int = 42
        _ = v42
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
