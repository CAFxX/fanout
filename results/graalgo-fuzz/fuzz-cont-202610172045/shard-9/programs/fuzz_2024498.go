package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 string; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 uint16; F1 int; F2 uint8 }
type GenS6 struct { F0 int; F1 uint32; F2 int }
type GenS7 struct { F0 uint64; F1 int8; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 int64 }
func genFn0() uint32 {
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: "", F1: 128, F2: 501}
            var v8 string = "\n\t"
            v9 := GenS1{F0: "hello", F1: 12, F2: 194}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        v4 += (v4 >> ((1) & 63))
        _ = fi5
    }
    switch 1 {
    case 128:
        if false {
            var v10 uint32 = 0
            _ = v10
        }
    case 9223372036854775807:
        v11 := []int{32767, -906, 223, 707}
        _ = v11
    case -1:
        if false {
            v4 = v4
        }
    default:
        var v12 float32 = 3.14159
        _ = v12
    }
    _ = v4
    return 876
}
func genFn1(p0 string, p1 uint, p2 int64) int16 {
    var v13 int = 42
    if false {
        v14 := GenS1{F0: (("\n\t" + "") + "a"), F1: 128, F2: 2147483647}
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int16 = 0
            if false {
                fmt.Println(fi15)
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 string = "a"
                    p1 = 2
                    switch (fi15 & fi17) {
                    case 9223372036854775807:
                        fmt.Println(fi15)
                    }
                    _ = fi17
                    _ = v18
                }
                v19 := []int{(-(fi15))}
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        v20 := GenS1{F0: "", F1: 255, F2: 0}
        _ = v14
        _ = v20
    }
    var v21 uint32 = genFn0()
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v21
    return 567
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{}
    var v3 uint64 = 454
    switch 128 {
    case 42:
        if len(v2) > 0 {
            v2[0] = 1
        }
    case 883:
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 += v3
            var v23 bool = true
            var v24 uint32 = 0
            _ = fi22
            _ = v23
            _ = v24
        }
    case 128:
        v3 += (v3 << ((65535) & 63))
    default:
        switch "hello" {
        case "foo bar":
            var v25 float64 = 1e10
            _ = v25
        case "hello":
            var v26 int16 = genFn1((("a" + "hello") + (("foo bar" + ("hello" + ("foo bar" + ""))) + "hello")), 18446744073709551615, 0)
            _ = v26
        }
    }
    var v27 float64 = 0.0
    switch 1 {
    case -9223372036854775808:
        v28 := GenS2{F0: "\n\t", F1: -9223372036854775808, F2: 434}
        _ = v28
    case -1:
        if (-2147483648 > 2147483647) {
            var v29 int = -2147483648
            v1[4294967295] = 949
            _ = v29
        } else {
            v27 -= (v27 + float64(3.14159))
            var v30 int = 4294967295
            _ = v30
        }
    case 1:
        var v31 int64 = -9223372036854775808
        _ = v31
    }
    var v32 int64 = 807
    v33 := GenS2{F0: "foo bar", F1: v32, F2: (v3 / ((v3) | 1))}
    v34 := GenS3{F0: 2147483647}
    if true {
        if (true || false) {
            if (v27 == v27) {
                var v35 uint8 = 128
                _ = v35
            }
            fmt.Println(v32, v3, v32)
        }
        v36 := GenS4{F0: "foo bar", F1: 9223372036854775807, F2: 42}
        _ = v36
    } else {
        if (624 < 18446744073709551615) {
            v3 = (^(v3))
            v37 := GenS5{F0: 42, F1: -2147483648, F2: 2}
            _ = v37
        }
        fmt.Println(v32)
        for fi38 := 0; fi38 < 12; fi38++ {
            for fi39 := 0; fi39 < 15; fi39++ {
                fmt.Println(v32, fi38, v3)
                v40 := GenS6{F0: fi39, F1: 2, F2: (^(fi39))}
                var v41 float32 = 3.14159
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        v27 = -2.25
        for fi43 := 0; fi43 < 14; fi43++ {
            if true {
                if (44 != 42) {
                    var v44 uint8 = 2
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    switch "foo bar" {
                    case "hello":
                        if ((4294967295 == 255) || (true || (!((!true) || true)))) {
                            switch fi43 {
                            case -2147483648:
                                switch (fi43 * 9223372036854775807) {
                                case -252:
                                    if (1 == (fi43 & (fi42 * (-(fi42))))) {
                                        var v45 float64 = ((float64((1e10 / 0.1)) / 1.0) / v27)
                                        v1[(fi43 << ((65535) & 63))] = genFn0()
                                        _ = v45
                                    } else {
                                        v44 += (v44 << ((919) & 63))
                                    }
                                    v32 = -9223372036854775808
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v44, v44, v3)
                                    var v46 float64 = float64(1.5)
                                    v44 += (^(v44))
                                    _ = v46
                                default:
                                    var v47 string = "\n\t"
                                    _ = v47
                                }
                            default:
                                fmt.Println(v32, v44)
                            }
                        }
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v48 float32 = 1e10
                        _ = v48
                    case "a":
                        var v49 uint8 = 148
                        _ = v49
                    default:
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                    _ = v44
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v50 int = 128
                        _ = v50
                    case "":
                        fmt.Println(v32)
                    default:
                        if true {
                            switch fi43 {
                            case 32767:
                                var v51 int64 = 634
                                _ = v51
                            case 128:
                                v2 = append(v2, 95)
                            case 208:
                                if false {
                                    var v52 int8 = 43
                                    _ = v52
                                } else {
                                    var v53 float32 = (0.1 * 0.1)
                                    _ = v53
                                }
                            default:
                                var v54 int = fi43
                                _ = v54
                            }
                        } else {
                            v55 := GenS6{F0: 670, F1: 42, F2: -511}
                            v32 -= (^(v32))
                            v56 := GenS7{F0: (v3 * 18446744073709551615), F1: -128, F2: (!true)}
                            _ = v55
                            _ = v56
                        }
                    }
                }
                if false {
                    v1[fi42] = 42
                    fmt.Println(v32)
                    var v57 int16 = -32768
                    _ = v57
                } else {
                    fmt.Println(v32)
                    fmt.Println(v3, v27, fi42)
                    v27 = 0.1
                    v58 := GenS2{F0: ("hello" + "\n\t"), F1: -690, F2: (v3 ^ (^(v3)))}
                    _ = v58
                }
            }
            if (!false) {
                v59 := v1[(^(fi42))]
                v3 -= v3
                var v60 float32 = -2.25
                if true {
                    v61 := GenS8{F0: 55}
                    v62 := GenS8{F0: 59}
                    var v63 int = fi43
                    if false {
                        var v64 uint16 = 2
                        fmt.Println(v59)
                        _ = v64
                    } else {
                        var v65 float64 = (((float64((3.14159 / 1.0)) * -2.25) / v27) - float64(v60))
                        _ = len(v1)
                        _ = v65
                    }
                    switch (v63 & fi43) {
                    case 9223372036854775807:
                        if len(v2) > 0 {
                            v2[0] = (^(v59))
                        }
                    case 0:
                        if true {
                            v27 = v27
                            _ = len(v2)
                        } else {
                            v66 := GenS2{F0: ("foo bar" + ((("foo bar" + "\n\t") + ("hello" + "\n\t")) + "a")), F1: (^(v32)), F2: (v3 ^ v3)}
                            switch "" {
                            case "\n\t":
                                var v67 uint64 = v3
                                _ = v67
                            default:
                                if len(v2) > 0 {
                                    v2[0] = v59
                                }
                            }
                            _ = v66
                        }
                    default:
                        v60 = v60
                    }
                    if ((v59 / ((992) | 1)) == 4294967295) {
                        var v68 float32 = 1.5
                        v69 := GenS2{F0: "hello", F1: v32, F2: v3}
                        _ = v68
                        _ = v69
                    } else {
                        var v70 int16 = 0
                        _ = v70
                    }
                    _ = v61
                    _ = v62
                    _ = v63
                }
                _ = v59
                _ = v60
            } else {
                if (genFn1("", 65535, v32) <= 1000) {
                    switch "a" {
                    case "":
                        v27 = (float64(3.14159) * 1.5)
                    }
                } else {
                    v2 = append(v2, 2147483648)
                }
                v27 -= v27
                v71 := GenS9{F0: (v32 % (((^(v32))) | 1)), F1: v32}
                v72 := GenS3{F0: 949}
                v73 := GenS4{F0: "", F1: (fi42 | (fi43 ^ fi42)), F2: -77}
                var v74 uint16 = 2
                v3 = (v3 & v3)
                _ = v71
                _ = v72
                _ = v73
                _ = v74
            }
            var v75 float32 = -2.25
            _ = fi43
            _ = v75
        }
        fmt.Println(v32)
        _ = fi42
    }
    var v76 uint64 = v3
    delete(v1, int(v32))
    v32 = (v32 - v32)
    v77 := GenS2{F0: ("foo bar" + "foo bar"), F1: v32, F2: (v3 & 0)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v76
    _ = v77
}
