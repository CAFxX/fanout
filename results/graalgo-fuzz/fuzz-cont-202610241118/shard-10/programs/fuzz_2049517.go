package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0() int64 {
    var v4 bool = true
    v5 := GenS1{F0: 100, F1: 1}
    var v6 int = 9223372036854775807
    v6 = (v6 & (v6 * 608))
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 71, F1: 7}
            if v4 {
                var v10 int16 = 835
                v11 := GenS2{F0: 42, F1: 32767, F2: v10}
                switch 1 {
                case 127:
                    v6 = fi7
                default:
                    var v12 int8 = -128
                    _ = v12
                }
                switch "hello" {
                case "a":
                    v13 := []int{9223372036854775807, fi8, (^(fi8)), fi8}
                    _ = v13
                }
                _ = v10
                _ = v11
            } else {
                var v14 string = "foo bar"
                v14 = "a"
                v15 := GenS2{F0: 126, F1: 32767, F2: 577}
                _ = v14
                _ = v15
            }
            _ = fi8
            _ = v9
        }
        v6 = 903
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return -588
}
func genFn1(p0 int8, p1 uint, p2 int16) uint {
    fmt.Println(p0)
    var v16 bool = false
    if v16 {
        v17 := []int{}
        if (v16 || (42 >= 2)) {
            fmt.Println(p0, p1, p0)
        }
        p2 += p2
        _ = v17
    }
    var v18 uint16 = 194
    if (49 == -498) {
        p2 -= -1
        p1 += (^(p1))
        v19 := GenS2{F0: 255, F1: 255, F2: p2}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v18
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = -526
    var v3 uint64 = 42
    if false {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS2{F0: 128, F1: 577, F2: (v1 + -774)}
            _ = fi20
            _ = v21
        }
        if true {
            var v22 float32 = 0.1
            _ = v22
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                switch 517 {
                case -2147483648:
                    switch "hello" {
                    case "a":
                        if (42 <= -1) {
                            v2 += 1
                            v3 = v3
                            v24 := GenS3{F0: (v1 | v1), F1: -81}
                            _ = v24
                        } else {
                            if (702 < 42) {
                                v25 := []int{fi23}
                                v26 := GenS1{F0: -77, F1: 66}
                                _ = v25
                                _ = v26
                            }
                        }
                    case "":
                        v1 = -824
                    case "foo bar":
                        fmt.Println(v3)
                    }
                case 4294967295:
                    v2 = (v2 % ((v2) | 1))
                default:
                    v3 = v3
                }
                v27 := []int{(-(fi23)), int(v3), 42}
                _ = fi23
                _ = v27
            }
            v28 := GenS3{F0: (v1 << ((genFn1(0, 65535, -391)) & 63)), F1: -1}
            _ = v28
        }
        var v29 int8 = 42
        _ = v29
    }
    if false {
        if (195 <= 42) {
            fmt.Println(v1, v2)
            v30 := GenS3{F0: v1, F1: -123}
            _ = v30
        }
    }
    v3 += v3
    var v31 float32 = 1e10
    var v32 uint = genFn1(121, 1, v1)
    var v33 float64 = 0.0
    var v34 uint8 = 1
    v3 += (v3 - uint64(v34))
    var v35 int16 = (^(v1))
    fmt.Println(v35, v31, v2)
    var v36 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
