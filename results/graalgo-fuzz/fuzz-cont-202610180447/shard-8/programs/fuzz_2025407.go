package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 bool }
func genFn0(p0 uint32, p1 int64, p2 uint16) int32 {
    fmt.Println(p2)
    if (42 > 127) {
        var v4 uint16 = 255
        v5 := GenS1{F0: 1}
        fmt.Println(p1)
        _ = v4
        _ = v5
    } else {
        p1 -= 42
    }
    v6 := []int{127, -2147483648}
    v7 := GenS2{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint32) int64 {
    fmt.Println(p1, p1)
    var v8 bool = false
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    var v1 int = 42
    var v2 string = ("a" + "foo bar")
    var v3 float32 = -2.25
    v2 = ("hello" + v2)
    v2 = ((v2 + v2) + (("\n\t" + "a") + "hello"))
    v9 := []int{2147483647, 32767, (v1 & v1)}
    if false {
        v1 -= (-(v1))
    }
    if (true || (!false)) {
        v3 = ((1e10 + v3) + (v3 * v3))
        v10 := GenS3{F0: 1, F1: 9223372036854775807}
        _ = v10
    } else {
        v3 -= (-2.25 * v3)
        var v11 int64 = 42
        if true {
            var v12 int = (v1 | (v1 * v1))
            v3 -= v3
            _ = v12
        } else {
            v13 := GenS1{F0: 42}
            var v14 int16 = 32767
            _ = v13
            _ = v14
        }
        _ = v11
    }
    if ((-38 > -1) || (genFn1(885, 2147483647) != 0)) {
        if true {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v2 = (v2 + "hello")
        } else {
            var v15 uint16 = 255
            if len(v9) > 0 {
                v9[0] = v1
            }
            _ = v15
        }
        v16 := GenS1{F0: 255}
        var v17 float64 = (-2.25 * 0.0)
        _ = v16
        _ = v17
    } else {
        var v18 int8 = -1
        _ = v18
    }
    _ = len(v9)
    if (1 <= 30) {
        var v19 float64 = (0.1 + 0.1)
        _ = v19
    }
    if (-2147483648 >= -2147483648) {
        if (!((!(1 == genFn1(205, 42))) || (((((3.14159 + v3) + -2.25) / v3) / 1.0) != (((v3 - (v3 - (0.0 - v3))) - (1.5 / 1e10)) - (((1e10 + 1.5) - (v3 / v3)) * (0.1 + v3)))))) {
            if (true || (2147483647 <= genFn0(649, -9223372036854775808, 1))) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    if false {
                        if false {
                            for fi21 := 0; fi21 < 20; fi21++ {
                                if len(v9) > 0 {
                                    v9[0] = -9223372036854775808
                                }
                                v9 = append(v9, v1)
                                switch fi20 {
                                case -9223372036854775808:
                                    var v22 string = ((v2 + "\n\t") + v2)
                                    _ = v22
                                }
                                _ = fi21
                            }
                        } else {
                            if false {
                                if len(v9) > 0 {
                                    v9[0] = (fi20 | -853)
                                }
                                var v23 int32 = genFn0(1, 42, 2)
                                _ = v23
                            } else {
                                var v24 float32 = float32((1.5 * -2.25))
                                v25 := GenS3{F0: genFn1(1, 394), F1: 9223372036854775807}
                                _ = v24
                                _ = v25
                            }
                        }
                        v2 = v2
                    }
                    var v26 float64 = -2.25
                    _ = fi20
                    _ = v26
                }
            }
            switch v1 {
            case 2147483647:
                v3 -= -2.25
            case 127:
                fmt.Println(v1, v1)
            default:
                v1 = v1
            }
            v27 := GenS1{F0: 427}
            _ = v27
        }
    }
    if (true || (genFn1(907, 2147483647) < -1)) {
        if (65535 >= 65535) {
            switch 9223372036854775807 {
            case -17:
                v3 = v3
            case 1:
                v3 -= v3
            }
        } else {
            fmt.Println(v1, v3)
            var v28 string = (v2 + (("foo bar" + ((("\n\t" + v2) + "hello") + ("a" + ("a" + (v2 + (v2 + v2)))))) + ((v2 + "") + "\n\t")))
            _ = v28
        }
    }
    v1 += (v1 * v1)
    v2 = v2
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println(v2)
    v1 = (v1 ^ 42)
    v29 := GenS4{F0: false, F1: 2, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v9
}
