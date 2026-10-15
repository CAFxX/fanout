package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 bool }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int16; F1 uint8; F2 int }
type GenS6 struct { F0 uint; F1 int64 }
func genFn0(p0 int8, p1 uint32, p2 uint8) uint64 {
    v6 := []int{}
    var v7 string = ("\n\t" + "foo bar")
    p1 = 2147483647
    switch "" {
    case "hello":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: true, F1: v7}
            _ = fi8
            _ = v9
        }
    case "":
        var v10 bool = false
        _ = v10
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: 32767, F1: 478}
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    var v13 uint8 = 128
    v13 = v13
    _ = v13
    return "\n\t"
}

func main() {
    var v1 int64 = 0
    var v2 uint = 2
    var v3 uint16 = 992
    v4 := []string{"a", "", ((((("a" + "") + "") + "") + "a") + "foo bar"), ""}
    v5 := []int{42, 32767, 856}
    var v14 float64 = (((0.1 / 1.5) * float64(float32(((0.1 + ((0.1 / 1.0) - 1.5)) * (float64((-2.25 - 0.1)) - 1.5))))) - (float64(-2.25) / (float64((((((3.14159 * (0.1 + 3.14159)) * (3.14159 / 3.14159)) / 1.0) / (1.5 + 0.0)) - (((1.5 * 1.5) / (-2.25 - float32(1.5))) / 1e10))) + 0.1)))
    v15 := GenS1{F0: (v1 > v1), F1: "hello"}
    var v16 uint8 = 128
    for fi17 := 0; fi17 < 17; fi17++ {
        if len(v5) > 0 {
            v5[0] = 32767
        }
        v3 -= 47
        for fi18 := 0; fi18 < 17; fi18++ {
            v3 += (v3 << (((^(v2))) & 63))
            if ("foo bar" < "\n\t") {
                var v19 uint64 = 2
                if (-657 >= 0) {
                    var v20 int32 = 42
                    if (v20 != (v20 | (v20 >> (((v2 % ((65535) | 1))) & 63)))) {
                        if true {
                            if len(v5) > 0 {
                                v5[0] = fi18
                            }
                            fmt.Println(v2)
                        } else {
                            var v21 uint64 = (v19 / (((v19 << ((v2) & 63))) | 1))
                            _ = v21
                        }
                        if (32767 == 1) {
                            v22 := GenS2{F0: v3, F1: (v2 | 344)}
                            v23 := GenS1{F0: (false || (42 != 943)), F1: "foo bar"}
                            _ = len(v5)
                            v24 := GenS2{F0: 65535, F1: 1}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        var v25 float64 = v14
                        _ = v25
                    }
                    switch (fi18 * (-(fi17))) {
                    case 2147483647:
                        _ = len(v5)
                    case 9223372036854775807:
                        var v26 uint64 = (v19 ^ v19)
                        _ = v26
                    case 4294967295:
                        v2 = v2
                    }
                    _ = v20
                } else {
                    var v27 uint = 18446744073709551615
                    var v28 string = "\n\t"
                    _ = v27
                    _ = v28
                }
                if ((0 < 50) && true) {
                    if false {
                        v1 = (v1 / (((v1 | (v1 / ((v1) | 1)))) | 1))
                        v14 -= float64(-2.25)
                    }
                    if true {
                        var v29 int16 = int16(v1)
                        _ = v29
                    } else {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    }
                }
                v2 -= v2
                if false {
                    v16 = (v16 >> ((18446744073709551615) & 63))
                } else {
                    v5 = append(v5, fi17)
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    v3 += (^(v3))
                }
                _ = len(v5)
                _ = v19
            } else {
                switch ("" + "a") {
                case "foo bar":
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        var v30 uint8 = (v16 >> ((v2) & 63))
                        _ = v30
                    default:
                        v1 = v1
                    }
                case "":
                    if false {
                        var v31 float64 = v14
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        _ = v31
                    }
                }
            }
            switch "a" {
            case "a":
                v32 := GenS1{F0: (false || true), F1: "hello"}
                _ = v32
            case "":
                if true {
                    switch "a" {
                    case "a":
                        if false {
                            var v33 int8 = 100
                            _ = v33
                        } else {
                            var v34 int32 = 651
                            _ = v34
                        }
                    case "\n\t":
                        switch fi17 {
                        case -9223372036854775808:
                            var v35 int8 = 100
                            _ = v35
                        case 981:
                            switch "\n\t" {
                            case "hello":
                                v14 = ((3.14159 / (v14 * v14)) - ((((0.1 * v14) + 3.14159) / 1.0) / v14))
                            }
                        case 2147483647:
                            v36 := GenS2{F0: v3, F1: (^(v2))}
                            _ = v36
                        default:
                            _ = len(v5)
                        }
                    default:
                        if (965 < 42) {
                            v1 += (v1 - (v1 / (((v1 / ((v1) | 1))) | 1)))
                            v2 = v2
                        }
                    }
                }
            default:
                v1 = 729
            }
            _ = fi18
        }
        _ = fi17
    }
    var v37 uint64 = 4294967295
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        if (!(true || false)) {
            switch 0 {
            case 847:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 9223372036854775807:
                if true {
                    fmt.Println(v16)
                    switch "hello" {
                    case "":
                        var v38 float64 = 3.14159
                        _ = v38
                    case "a":
                        var v39 int32 = 42
                        _ = v39
                    case "foo bar":
                        if true {
                            v37 -= (v37 * (^(v37)))
                            v1 = v1
                        } else {
                            v40 := GenS3{F0: 42, F1: 0, F2: (uint(v3) > v2)}
                            v41 := GenS2{F0: (^(v3)), F1: (v2 ^ 229)}
                            _ = v40
                            _ = v41
                        }
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                }
            case 384:
                if ((v14 / v14) <= 1e10) {
                    v42 := GenS4{F0: 42, F1: "\n\t"}
                    _ = v42
                } else {
                    v5 = append(v5, -1)
                }
            }
            v16 = (v16 / (((v16 | 2)) | 1))
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        }
        switch "\n\t" {
        case "":
            var v43 int64 = v1
            _ = v43
        case "hello":
            switch "foo bar" {
            case "foo bar":
                var v44 bool = (v1 > -528)
                _ = v44
            default:
                v45 := GenS5{F0: 1, F1: (^(v16)), F2: -2147483648}
                _ = v45
            }
        case "foo bar":
            v46 := GenS6{F0: (^(v2)), F1: v1}
            _ = v46
        }
    }
    v1 += v1
    fmt.Println(v2, v16, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
