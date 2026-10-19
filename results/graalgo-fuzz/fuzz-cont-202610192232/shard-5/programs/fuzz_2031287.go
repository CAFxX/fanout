package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16; F1 int32; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int16 }
func genFn0(p0 int64, p1 uint8) string {
    if (false || (!(!false))) {
        v6 := []int{}
        fmt.Println(p0)
        _ = v6
    }
    v7 := GenS1{F0: 255, F1: -1}
    _ = p0
    _ = p1
    _ = v7
    return "a"
}
func genFn1(p0 int) int64 {
    if true {
        if (true || (0 != -54)) {
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := []int{fi9, (p0 & (^(fi9))), fi9, (fi8 & fi8)}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        }
        if (-1 >= 42) {
            if (42 != 42) {
                fmt.Println(p0)
                var v11 int64 = 900
                _ = v11
            } else {
                switch (^(p0)) {
                case 42:
                    var v12 int8 = 0
                    _ = v12
                case 2147483647:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 int32 = 42
                        v14 = v14
                        var v15 float64 = 3.14159
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    var v16 string = "foo bar"
                    _ = v16
                }
                p0 = p0
            }
        } else {
            p0 += (p0 | (^(p0)))
            var v17 int32 = 1
            var v18 int = p0
            _ = v17
            _ = v18
        }
    }
    v19 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v19
    return -1
}

func main() {
    v1 := []string{(("" + "") + ("a" + ((("hello" + "hello") + "\n\t") + "hello"))), ("\n\t" + "\n\t"), ("" + "hello"), "hello"}
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 uint16 = 42
    if (!((!((v5 <= 2) && (false || true))) || (-690 == 513))) {
        v5 = 32767
        var v20 int8 = -1
        v21 := GenS1{F0: 38, F1: 0}
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: 127, F1: -2147483648}
        _ = v22
    }
    if (false && (32767 != -995)) {
        switch 2147483647 {
        case 729:
            v4["a"] = 4294967295
        case 32768:
            var v23 float32 = 3.14159
            _ = v23
        case -9223372036854775808:
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 uint16 = v5
                v26 := GenS1{F0: 0, F1: 1}
                v2 = -1
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        if ((-2.25 - 0.0) < (0.0 + float32(-2.25))) {
            var v27 int = 32767
            v5 += (v5 << ((726) & 63))
            _ = v27
        }
        var v28 bool = false
        _ = v28
    } else {
        v2 = 71
        var v29 float32 = 0.1
        _ = v29
    }
    v5 = 696
    fmt.Println(v5, v2)
    switch 452 {
    case -178:
        if false {
            switch "" {
            case "hello":
                if false {
                    if false {
                        v30 := GenS1{F0: 88, F1: 1}
                        _ = v30
                    } else {
                        if true {
                            v2 += genFn1(128)
                            v2 -= (v2 / (((-(v2))) | 1))
                            v31 := GenS2{F0: 127, F1: 4294967295, F2: (v2 | (v2 << ((2) & 63)))}
                            _ = v31
                        } else {
                            if true {
                                var v32 int32 = 2147483647
                                var v33 int = -27
                                v34 := GenS2{F0: 255, F1: 42, F2: 9223372036854775807}
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                switch ("a" + "a") {
                                case "\n\t":
                                    v5 += v5
                                case "a":
                                    if false {
                                        v5 += v5
                                    }
                                case "foo bar":
                                    var v35 uint64 = 15
                                    _ = v35
                                default:
                                    switch 1 {
                                    case -58:
                                        v2 -= genFn1(9223372036854775807)
                                    case -1:
                                        var v36 string = "\n\t"
                                        _ = v36
                                    }
                                }
                            }
                        }
                        switch 4294967295 {
                        case -593:
                            v37 := GenS2{F0: 2, F1: 2, F2: genFn1(-1)}
                            _ = v37
                        default:
                            if (-1 <= 127) {
                                var v38 uint16 = (^(v5))
                                _ = v38
                            }
                        }
                    }
                } else {
                    if (!(231 >= 1)) {
                        v39 := v4["foo bar"]
                        _ = v39
                    }
                    fmt.Println(v5, v2)
                    v40 := v4["foo bar"]
                    _ = v40
                }
            case "a":
                var v41 float64 = 0.0
                _ = v41
            case "":
                fmt.Println(v5)
            default:
                v2 += v2
            }
            fmt.Println(v2, v5)
        } else {
            for fi42 := 0; fi42 < 19; fi42++ {
                v43 := GenS3{F0: (fi42 << ((65535) & 63)), F1: 1}
                v1 = append(v1, ("\n\t" + ""))
                switch (fi42 * fi42) {
                case 1:
                    fmt.Println(v5, v5, v2)
                case 32768:
                    v5 = (v5 & v5)
                case 127:
                    v2 -= -988
                }
                _ = fi42
                _ = v43
            }
            v5 -= (v5 * v5)
            var v44 uint8 = 103
            _ = v44
        }
    case 32767:
        if (("hello" + "foo bar") >= "") {
            v5 = (v5 * (v5 % ((v5) | 1)))
            var v45 uint = 823
            if (false || ((true && false) && ((float64((0.0 / (1e10 - (0.1 + float32(float64(-2.25)))))) - 3.14159) > ((0.0 + ((float64((float32((((-2.25 * (-2.25 / (1e10 / -2.25))) * ((1.5 * 3.14159) - (0.1 / 0.1))) + 0.1)) * (0.0 * float32(1.5)))) / (1e10 + 0.1)) / 3.14159)) + 0.1)))) {
                if ("" < "foo bar") {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    v46 := GenS4{F0: -1, F1: -138, F2: (v2 - (v2 % (((v2 - (v2 - 1))) | 1)))}
                    v47 := GenS5{F0: 0}
                    _ = v46
                    _ = v47
                }
                for fi48 := 0; fi48 < 5; fi48++ {
                    v2 += (-(v2))
                    var v49 uint16 = 405
                    _ = fi48
                    _ = v49
                }
            }
            _ = v45
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                var v51 float64 = (1e10 / (-2.25 / -2.25))
                v51 = v51
                for fi52 := 0; fi52 < 4; fi52++ {
                    v1 = append(v1, "\n\t")
                    var v53 int64 = v2
                    _ = fi52
                    _ = v53
                }
                _ = fi50
                _ = v51
            }
            v4["a"] = 2147483647
            for fi54 := 0; fi54 < 16; fi54++ {
                _ = len(v4)
                _ = fi54
            }
        }
    default:
        fmt.Println(v5, v2)
    }
    if false {
        v2 -= v2
        var v55 int = 128
        _ = v55
    } else {
        var v56 string = "a"
        _ = v56
    }
    if (true || true) {
        for fi57 := 0; fi57 < 16; fi57++ {
            var v58 int32 = 42
            v58 = 823
            _ = fi57
            _ = v58
        }
    } else {
        v59 := v4[""]
        _ = v59
    }
    v60 := GenS6{F0: 9223372036854775807, F1: -74}
    var v61 uint8 = 41
    if (((v61 << ((18446744073709551615) & 63)) > v61) || false) {
        var v62 string = "a"
        v62 = v62
        _ = v62
    } else {
        var v63 int = 32767
        _ = v63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
}
