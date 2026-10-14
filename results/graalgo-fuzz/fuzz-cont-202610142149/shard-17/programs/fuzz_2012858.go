package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int32 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 bool; F1 bool; F2 int64 }
func genFn0(p0 int16, p1 uint16, p2 uint) string {
    v6 := GenS1{F0: (-128 > 27), F1: "\n\t"}
    p0 += p0
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(fi7 * (fi7 << (((p2 << ((p2) & 63))) & 63))), fi7, -2147483648}
        if (0 != 123) {
            if len(v8) > 0 {
                v8[0] = (fi7 >> ((p2) & 63))
            }
            switch "\n\t" {
            case "a":
                v9 := GenS2{F0: ("\n\t" + "hello"), F1: false}
                _ = v9
            case "\n\t":
                var v10 bool = (false || true)
                _ = v10
            default:
                _ = len(v8)
            }
        } else {
            var v11 int8 = -68
            _ = v11
        }
        v12 := GenS3{F0: 18446744073709551615, F1: 576, F2: 0}
        _ = fi7
        _ = v12
        _ = v8
    }
    if true {
        var v13 uint8 = 0
        _ = v13
    } else {
        fmt.Println(p2)
        v14 := GenS4{F0: (p1 | p1), F1: 1, F2: 1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint8, p1 uint16) uint {
    fmt.Println(p1, p0)
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            switch "hello" {
            case "a":
                fmt.Println(p1, fi15)
            case "hello":
                p1 = 255
            default:
                v17 := GenS2{F0: (("a" + "hello") + ""), F1: (p1 >= 65535)}
                _ = v17
            }
            _ = fi16
        }
        p1 += uint16(p0)
        _ = fi15
    }
    switch "\n\t" {
    case "hello":
        if (false || (((0.0 * (1.5 * ((float32(3.14159) / 1.0) + (-2.25 * (1.5 * (-2.25 / (float32(1.5) + 1.5))))))) / 3.14159) != 3.14159)) {
            if (!(0 > 1)) {
                var v18 uint = 171
                var v19 uint64 = 733
                p0 -= p0
                _ = v18
                _ = v19
            }
            var v20 float32 = (3.14159 / 1.0)
            if (42 == 18446744073709551615) {
                var v21 int64 = 42
                _ = v21
            }
            _ = v20
        } else {
            if (-1 <= 32767) {
                v22 := []int{128, -9223372036854775808, 0, 1}
                v22 = append(v22, int(p0))
                v22 = append(v22, 1)
                _ = v22
            } else {
                var v23 int16 = 32767
                var v24 int32 = 0
                _ = v23
                _ = v24
            }
            p1 = p1
            var v25 int = 9223372036854775807
            _ = v25
        }
    case "\n\t":
        p1 = p1
    default:
        var v26 uint8 = (p0 | p0)
        _ = v26
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []int{1}
    var v2 uint32 = 65535
    v3 := make(map[int]bool)
    var v4 uint16 = 255
    var v5 int = -181
    if false {
        v27 := GenS1{F0: (float32(0.1) == float32(1.5)), F1: "\n\t"}
        _ = v27
    } else {
        var v28 bool = ((1.5 - ((0.0 - 1.5) / 3.14159)) >= 0.1)
        var v29 uint8 = 128
        _ = v28
        _ = v29
    }
    _ = len(v1)
    v3[1] = false
    v2 = (v2 & (^(v2)))
    var v30 float32 = 1.5
    var v31 float32 = v30
    if (1 > 695) {
        v5 -= v5
        var v32 float64 = float64(3.14159)
        _ = v32
    } else {
        for fi33 := 0; fi33 < 20; fi33++ {
            fmt.Println(v5, v31)
            for fi34 := 0; fi34 < 5; fi34++ {
                v35 := GenS5{F0: (^(fi34)), F1: v5}
                v30 = 3.14159
                _ = fi34
                _ = v35
            }
            v4 = 458
            _ = fi33
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            fmt.Println(fi36)
            var v37 int64 = -184
            var v38 int = v5
            _ = fi36
            _ = v37
            _ = v38
        }
        v3[v5] = (((!(!((v5 == 32768) && (int32(v2) <= 889)))) || (false && false)) || (true && false))
    }
    if (1 <= 1) {
        if (9223372036854775807 == -567) {
            var v39 int16 = 0
            delete(v3, (v5 ^ v5))
            _ = v39
        } else {
            v40 := GenS1{F0: true, F1: (("hello" + "a") + "\n\t")}
            if (!true) {
                if true {
                    var v41 int8 = -128
                    _ = v41
                }
            }
            _ = v40
        }
        fmt.Println(v31, v2)
        if (genFn1(42, (v4 % ((65535) | 1))) >= 18446744073709551615) {
            v5 -= v5
            for fi42 := 0; fi42 < 10; fi42++ {
                v2 += (v2 << ((18446744073709551615) & 63))
                switch 32767 {
                case -9223372036854775808:
                    for fi43 := 0; fi43 < 20; fi43++ {
                        var v44 uint = 1
                        var v45 int16 = 42
                        v5 = -970
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                default:
                    if (v30 <= (3.14159 / ((0.1 - (v31 * (float32(-2.25) / v30))) / 1.0))) {
                        v46 := GenS1{F0: (-2147483648 > (v5 / ((v5) | 1))), F1: ("\n\t" + ("" + ""))}
                        v3[fi42] = ((0 < 2147483647) && true)
                        _ = v46
                    } else {
                        switch ("" + "hello") {
                        case "foo bar":
                            switch 9223372036854775807 {
                            case -468:
                                if (94 != 0) {
                                    v47 := GenS1{F0: true, F1: "foo bar"}
                                    var v48 int8 = 127
                                    if false {
                                        if ((-1 < 664) || false) {
                                            var v49 int64 = 42
                                            _ = v49
                                        }
                                        fmt.Println(v48, v48, v4)
                                    }
                                    _ = v47
                                    _ = v48
                                }
                            case -9223372036854775808:
                                v50 := GenS6{F0: true, F1: false, F2: 0}
                                _ = v50
                            case 127:
                                var v51 int32 = 0
                                _ = v51
                            default:
                                _ = len(v3)
                            }
                        case "hello":
                            fmt.Println(v5, fi42)
                        }
                    }
                }
                _ = fi42
            }
            var v52 bool = true
            _ = v52
        }
    } else {
        v31 += (1.5 * 1e10)
        v31 = (((v31 - v30) / 1.0) - (0.0 / float32((1e10 * 1e10))))
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v4 = (v4 * (v4 - v4))
    v5 = 930
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
