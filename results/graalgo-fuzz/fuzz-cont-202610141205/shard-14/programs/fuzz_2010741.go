package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint; F2 string }
type GenS8 struct { F0 uint64; F1 bool; F2 int8 }
func genFn0(p0 string, p1 uint, p2 string) int {
    var v7 int8 = int8(p1)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}

func main() {
    var v1 int32 = 95
    var v2 int64 = -1
    var v3 uint64 = 2
    var v4 uint8 = 2
    v5 := []uint32{2147483648, 92}
    var v6 uint = 18446744073709551615
    v5 = append(v5, 2)
    v4 = v4
    for fi8 := 0; fi8 < 18; fi8++ {
        v4 = v4
        fmt.Println(v1, fi8, fi8)
        _ = fi8
    }
    if (0 >= genFn0("\n\t", 18446744073709551615, "\n\t")) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v5 = append(v5, 42)
            _ = fi9
        }
        _ = len(v5)
        switch 32767 {
        case -2147483648:
            fmt.Println(v6, v2, v4)
        case -677:
            v10 := GenS1{F0: uint16(v4), F1: 65535, F2: (v1 | v1)}
            _ = v10
        }
    }
    var v11 int64 = (-(v2))
    if true {
        for fi12 := 0; fi12 < 18; fi12++ {
            v1 = v1
            _ = fi12
        }
        if (true && true) {
            if true {
                v2 -= -1
                v5 = append(v5, 143)
                switch "\n\t" {
                case "a":
                    if false {
                        v4 += v4
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                default:
                    v6 = v6
                }
            } else {
                v13 := GenS2{F0: 2, F1: v4}
                var v14 uint16 = 255
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                _ = v13
                _ = v14
            }
            if false {
                v5 = append(v5, 2147483647)
                if (!(9223372036854775807 == 127)) {
                    var v15 float64 = -2.25
                    switch "a" {
                    case "":
                        if true {
                            fmt.Println(v15, v2, v3)
                        }
                    case "\n\t":
                        if true {
                            v5 = append(v5, 776)
                        }
                    }
                    if false {
                        v1 = (v1 / ((v1) | 1))
                        if (v1 > 0) {
                            if true {
                                v16 := GenS2{F0: v6, F1: v4}
                                v5 = append(v5, 50)
                                if false {
                                    if len(v5) > 0 {
                                        v5[0] = 2147483648
                                    }
                                    v1 -= v1
                                    v17 := GenS3{F0: (v3 + v3), F1: 0, F2: (v6 | (v6 - v6))}
                                    _ = v17
                                } else {
                                    v6 = v6
                                    if (false && false) {
                                        var v18 string = "foo bar"
                                        _ = v18
                                    }
                                }
                                _ = v16
                            } else {
                                var v19 uint8 = 9
                                if true {
                                    v20 := GenS4{F0: (v3 >> (((v6 - 708)) & 63)), F1: v6}
                                    _ = v20
                                }
                                _ = v19
                            }
                        } else {
                            fmt.Println(v3, v15)
                            v21 := GenS1{F0: 1, F1: 65535, F2: (v1 | (v1 & v1))}
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 4; fi22++ {
                            fmt.Println(v15, v3, v2)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v2 += (v2 - -1)
                            _ = fi22
                        }
                    }
                    _ = v15
                }
                if false {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi23
                    }
                    if (int(v1) != genFn0("a", (v6 ^ 2), "hello")) {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            v6 -= (v6 ^ (v6 & (v6 * v6)))
                            v3 += v3
                            _ = fi24
                        }
                    } else {
                        switch 32768 {
                        case 42:
                            v4 -= 17
                        case 454:
                            var v25 uint8 = (v4 << ((382) & 63))
                            _ = v25
                        case 212:
                            var v26 float32 = 3.14159
                            _ = v26
                        default:
                            v2 -= v2
                        }
                        var v27 int8 = 42
                        v6 = (v6 >> ((v6) & 63))
                        _ = v27
                    }
                }
            }
        }
    } else {
        if true {
            v28 := GenS5{F0: v4, F1: (v6 ^ (v6 & (v6 - (v6 ^ v6)))), F2: genFn0((("hello" + ("a" + "\n\t")) + "hello"), v6, "hello")}
            for fi29 := 0; fi29 < 7; fi29++ {
                switch -53 {
                case -2147483648:
                    v30 := GenS4{F0: uint64(v6), F1: v6}
                    _ = v30
                }
                _ = fi29
            }
            _ = v28
        } else {
            v1 += v1
        }
        v3 = v3
    }
    v31 := GenS6{F0: v1}
    for fi32 := 0; fi32 < 4; fi32++ {
        v5 = append(v5, 323)
        switch "hello" {
        case "foo bar":
            v6 -= 346
        }
        var v33 int16 = 372
        _ = fi32
        _ = v33
    }
    switch ("hello" + "a") {
    case "hello":
        fmt.Println(v11)
    case "":
        switch -597 {
        case 4294967295:
            v34 := GenS1{F0: 42, F1: 42, F2: v1}
            _ = v34
        case 32768:
            if true {
                v35 := GenS7{F0: (^(v4)), F1: v6, F2: ("hello" + "")}
                _ = v35
            } else {
                var v36 string = ("foo bar" + "foo bar")
                v37 := GenS8{F0: (v3 << ((v6) & 63)), F1: true, F2: 42}
                v1 += -2147483648
                _ = v36
                _ = v37
            }
        }
    default:
        var v38 float32 = (-2.25 + (0.1 * (0.0 * 0.1)))
        _ = v38
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -2147483648 {
    case 128:
        if (!(413 <= (v2 / (((v11 * v11)) | 1)))) {
            var v39 uint16 = 42
            var v40 int64 = 1
            v41 := GenS7{F0: (v4 >> ((uint(v3)) & 63)), F1: 512, F2: "foo bar"}
            _ = v39
            _ = v40
            _ = v41
        }
    case -9223372036854775808:
        var v42 float64 = (3.14159 / 0.1)
        _ = v42
    }
    var v43 bool = (("\n\t" != (("" + ("hello" + "")) + (("\n\t" + "foo bar") + ("\n\t" + (("foo bar" + "hello") + ("a" + "")))))) || (!(((((((1e10 / (1.5 + (((1.5 + ((0.0 + 0.1) - float32(float64((float32(0.0) * 1e10))))) - 1e10) / (((3.14159 / ((float32(1e10) - 1e10) - 1.5)) - (0.1 - 0.0)) + 0.1)))) + (0.0 * 1.5)) / 1.5) * (0.1 + (1e10 / (3.14159 / 0.1)))) / 1e10) / 0.1) == ((0.1 / 1.5) * 1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
