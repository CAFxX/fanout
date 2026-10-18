package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 uint32) string {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: true, F1: 0, F2: 255}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return (("\n\t" + "foo bar") + "hello")
}
func genFn1(p0 int64, p1 string) int64 {
    p0 = (p0 * (p0 ^ 42))
    var v7 int32 = 753
    if false {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int = fi8
            var v10 uint8 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: false, F1: 411, F2: 163}
        p0 = p0
        _ = v11
    }
    var v12 uint64 = 189
    v13 := []int{-479}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 32767
    var v3 uint16 = v2
    var v4 bool = false
    v14 := []int{470, int(v1)}
    var v15 float32 = (1e10 + (0.0 * 1.5))
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 int32 = -1
        if v4 {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        } else {
            v14 = append(v14, 2147483647)
        }
        var v18 uint64 = 2
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v1, v15)
    switch 32767 {
    case -967:
        v19 := GenS1{F0: (!(v4 || ((v15 + 1.5) != 0.1))), F1: 0, F2: 1}
        _ = v19
    case -9223372036854775808:
        var v20 float32 = v15
        _ = v20
    }
    var v21 int64 = 1
    if ("foo bar" <= "hello") {
        v22 := GenS2{F0: 42}
        _ = v22
    }
    if v4 {
        if len(v14) > 0 {
            v14[0] = -1
        }
        var v23 float64 = ((0.0 - ((-2.25 * ((1.5 + 0.1) + 1e10)) / 1.5)) + (0.1 + 3.14159))
        switch -9223372036854775808 {
        case -740:
            v3 = (v3 | (v3 & v3))
        }
        _ = v23
    } else {
        v14 = append(v14, 353)
        if v4 {
            var v24 float64 = 1e10
            v4 = v4
            v25 := GenS1{F0: v4, F1: 586, F2: 42}
            _ = v24
            _ = v25
        } else {
            var v26 uint32 = 211
            var v27 int16 = v1
            if (v4 && (v4 && true)) {
                var v28 int32 = 42
                v1 += (-(v27))
                for fi29 := 0; fi29 < 9; fi29++ {
                    v15 = (v15 + ((v15 + float32(0.0)) * (float32(0.0) / (v15 / (v15 / v15)))))
                    _ = fi29
                }
                _ = v28
            }
            _ = v26
            _ = v27
        }
        switch 2147483647 {
        case 127:
            if v4 {
                v30 := GenS3{F0: (v21 - (v21 + (v21 % ((v21) | 1)))), F1: 127}
                for fi31 := 0; fi31 < 19; fi31++ {
                    if (!v4) {
                        v3 = v2
                    }
                    _ = fi31
                }
                fmt.Println(v2, v2)
                _ = v30
            } else {
                v1 += v1
                v15 -= v15
                for fi32 := 0; fi32 < 3; fi32++ {
                    for fi33 := 0; fi33 < 7; fi33++ {
                        var v34 int32 = 2147483647
                        v15 = 3.14159
                        v1 = v1
                        v34 = (v34 + (v34 | v34))
                        v35 := GenS1{F0: v4, F1: 47, F2: 2}
                        v36 := GenS2{F0: 696}
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v21 = (v21 + -597)
                    var v37 uint16 = v2
                    _ = fi32
                    _ = v37
                }
            }
        case 1:
            v38 := GenS3{F0: (v21 | v21), F1: -75}
            _ = v38
        default:
            var v39 string = (("" + ("a" + ("hello" + "a"))) + "a")
            _ = v39
        }
    }
    if (3 == -1) {
        var v40 int16 = v1
        if (v4 || v4) {
            var v41 uint32 = 1
            _ = v41
        }
        _ = v40
    } else {
        for fi42 := 0; fi42 < 18; fi42++ {
            if v4 {
                fmt.Println(v21)
            }
            var v43 uint = 217
            _ = fi42
            _ = v43
        }
        v44 := GenS2{F0: 700}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
