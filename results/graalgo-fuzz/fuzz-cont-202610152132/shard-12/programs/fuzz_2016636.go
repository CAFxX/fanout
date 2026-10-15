package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint32; F2 int }
type GenS6 struct { F0 uint32; F1 string; F2 string }
func genFn0(p0 int16, p1 uint32, p2 int64) int8 {
    v6 := []int{-410, 4294967295, -959, 32767}
    v6 = append(v6, 0)
    if (!(127 == 127)) {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int = (fi7 ^ int(p2))
            v8 = -2147483648
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (0 != 255), F1: (p0 ^ p0), F2: (p1 ^ (p1 + p1))}
        _ = v9
    } else {
        if true {
            if true {
                var v10 int8 = 0
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 string = "hello"
                    if len(v6) > 0 {
                        v6[0] = fi11
                    }
                    _ = fi11
                    _ = v12
                }
                v10 = v10
                _ = v10
            }
        }
        v13 := GenS1{F0: (!(true && true)), F1: 860, F2: 42}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 int32) uint {
    p0 += (p0 & -1)
    v14 := GenS1{F0: false, F1: -32768, F2: 991}
    switch -519 {
    case 127:
        v15 := []int{878, -342, 32767, 127}
        _ = v15
    default:
        v16 := []int{-1, 266, 9223372036854775807, 4294967295}
        _ = v16
    }
    v17 := []int{32767, 0}
    p0 += (p0 + (^(p0)))
    _ = p0
    _ = v14
    _ = v17
    return 42
}

func main() {
    var v1 float64 = (float64(float32(-2.25)) / -2.25)
    var v2 uint = 263
    var v3 bool = (!true)
    var v4 string = (("" + "foo bar") + "a")
    var v5 float32 = 1.5
    fmt.Println(v5, v3)
    var v18 uint16 = 627
    if (v3 && (-32768 <= -1)) {
        v19 := []int{-1}
        _ = v19
    } else {
        v2 += uint(v18)
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        if (v3 || (v3 || v3)) {
            switch ("" + (v4 + "foo bar")) {
            case "":
                var v21 uint32 = 2147483647
                _ = v21
            }
        }
        v22 := []int{fi20, fi20, fi20, fi20}
        fmt.Println(v2, v2, v5)
        _ = fi20
        _ = v22
    }
    v23 := GenS2{F0: 21}
    switch v4 {
    case "foo bar":
        v24 := GenS3{F0: ((135 != 784) || (v3 || ((v3 || true) && (356 > -2147483648)))), F1: 65535}
        _ = v24
    case "\n\t":
        var v25 float32 = ((-2.25 + (0.0 / ((float32(v1) - (0.0 * 0.1)) + (0.1 / v5)))) / (1e10 + v5))
        _ = v25
    default:
        v4 = (((("a" + "foo bar") + (v4 + v4)) + v4) + (v4 + "a"))
    }
    if v3 {
        if false {
            var v26 float32 = 0.0
            if (-980 <= 42) {
                switch -180 {
                case 2147483647:
                    v27 := GenS2{F0: genFn0(0, 1, 42)}
                    _ = v27
                }
            } else {
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 float32 = (v5 / v5)
                    v5 += v26
                    var v30 uint = 255
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v31 := GenS4{F0: 42, F1: (v18 - v18)}
                _ = v31
            }
            _ = v26
        } else {
            v18 -= 31
            v32 := []int{127, 2147483647}
            _ = v32
        }
    } else {
        v33 := GenS5{F0: 9223372036854775807, F1: 711, F2: -9223372036854775808}
        _ = v33
    }
    if true {
        v5 += float32((v1 - v1))
    }
    v34 := GenS6{F0: 2, F1: (v4 + (v4 + (v4 + (v4 + v4)))), F2: v4}
    v35 := []int{-96, -9223372036854775808, -731}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
