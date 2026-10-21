package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 int32, p1 uint32) uint {
    var v4 float64 = (0.1 - -2.25)
    p0 = (-(p0))
    var v5 int = 2147483647
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 bool = (false || ((128 >= -1) && true))
    v2 := make(map[string]string)
    var v3 int = 32768
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 bool = (!v1)
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(v3, v3)
            var v9 float32 = 3.14159
            _ = fi8
            _ = v9
        }
        v7 = v1
        _ = fi6
        _ = v7
    }
    v3 -= v3
    v10 := []int{(v3 ^ v3)}
    v2["hello"] = "foo bar"
    if v1 {
        fmt.Println(v3)
        var v11 float32 = ((0.0 / 3.14159) + ((0.1 * float32(((-2.25 + float64(3.14159)) * 1e10))) - ((1e10 * ((0.1 + 1e10) / 0.1)) * float32(1.5))))
        _ = v11
    } else {
        v3 -= -9223372036854775808
        if (65535 < 4294967295) {
            if (42 <= -350) {
                var v12 bool = (42 < v3)
                var v13 float32 = 0.0
                _ = v12
                _ = v13
            }
            fmt.Println(v1, v3)
        } else {
            v14 := v2["\n\t"]
            v14 = (v14 + v14)
            _ = v14
        }
    }
    if (291 > 181) {
        if len(v10) > 0 {
            v10[0] = v3
        }
        var v15 uint8 = 255
        _ = v15
    }
    if (!v1) {
        fmt.Println(v3, v1)
    } else {
        var v16 string = ("\n\t" + "\n\t")
        _ = v16
    }
    var v17 bool = false
    switch "hello" {
    case "hello":
        v1 = v1
    case "foo bar":
        fmt.Println(v17)
    case "\n\t":
        v17 = true
    default:
        v3 = v3
    }
    var v18 int16 = 42
    switch "foo bar" {
    case "hello":
        var v19 int16 = (v18 | v18)
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 13; fi20++ {
            v2[""] = "\n\t"
            v3 += (v3 % ((fi20) | 1))
            _ = fi20
        }
    case "\n\t":
        v18 -= v18
    default:
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int64 = 384
            _ = fi21
            _ = v22
        }
    }
    v3 += v3
    var v23 int64 = 9223372036854775807
    if (!false) {
        var v24 bool = (42 < -1)
        var v25 int8 = 42
        _ = v24
        _ = v25
    }
    if (42 >= -519) {
        if (false || v17) {
            for fi26 := 0; fi26 < 3; fi26++ {
                v17 = v17
                var v27 uint16 = 1
                var v28 bool = v17
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 int8 = -128
            if ("\n\t" < ("a" + "")) {
                v18 = v18
            } else {
                fmt.Println(v1)
                fmt.Println(v23)
            }
            _ = v29
        }
        v30 := GenS1{F0: false, F1: 1}
        v1 = (3.14159 >= (((((0.1 * 0.1) / -2.25) + 0.0) / 3.14159) * float64((0.1 / 0.1))))
        _ = v30
    } else {
        if (v1 || ((("foo bar" + "\n\t") <= "foo bar") || v17)) {
            switch v3 {
            case 0:
                fmt.Println(v1, v23, v18)
            default:
                switch "" {
                case "":
                    switch v3 {
                    case -444:
                        v3 = (v3 | v3)
                    case -766:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v23 += (v23 << ((4294967295) & 63))
                            _ = fi31
                        }
                    default:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            fmt.Println(v23)
                            _ = fi32
                        }
                    }
                case "a":
                    for fi33 := 0; fi33 < 11; fi33++ {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            v35 := GenS1{F0: v1, F1: 255}
                            _ = fi34
                            _ = v35
                        }
                        fmt.Println(v1, v1)
                        _ = fi33
                    }
                case "foo bar":
                    v18 = 0
                default:
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if (42 < 42) {
                            fmt.Println(v1, fi36)
                            switch 128 {
                            case 127:
                                v23 = v23
                            }
                            v37 := GenS2{F0: 2147483647}
                            _ = v37
                        }
                        _ = fi36
                    }
                }
            }
            v38 := v2[((("a" + "\n\t") + "") + ("a" + "a"))]
            v18 += 1000
            _ = v38
        }
    }
    v3 -= (v3 & v3)
    if len(v10) > 0 {
        v10[0] = v3
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := GenS3{F0: -128, F1: genFn0(42, 479)}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
