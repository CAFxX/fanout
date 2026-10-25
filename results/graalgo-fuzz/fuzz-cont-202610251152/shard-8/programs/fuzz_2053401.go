package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
type GenS4 struct { F0 uint; F1 int16; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int32, p1 uint8, p2 string) bool {
    v6 := []int{279, -2147483648, -9223372036854775808, -354}
    for fi7 := 0; fi7 < 6; fi7++ {
        p2 = p2
        p1 = p1
        p2 = (p2 + (p2 + (p2 + (p2 + p2))))
        _ = fi7
    }
    if (!((42 != 4294967295) && true)) {
        p2 = p2
        if (257 <= 1) {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS1{F0: 42, F1: 255, F2: 65535}
                p2 = ("foo bar" + p2)
                _ = fi8
                _ = v9
            }
            fmt.Println(p0, p2, p0)
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: 1, F1: 2, F2: 42}
                p1 = 128
                _ = fi10
                _ = v11
            }
        } else {
            var v12 bool = (((9223372036854775807 == -44) && ((true || (3.14159 < (3.14159 * -2.25))) && (false && (p0 == p0)))) && false)
            fmt.Println(p1, p1)
            v13 := GenS1{F0: 0, F1: 42, F2: 866}
            _ = v12
            _ = v13
        }
    } else {
        if true {
            var v14 float32 = 1.5
            switch (p2 + ((p2 + (p2 + "")) + ("" + "hello"))) {
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "a":
                p2 = p2
            case "":
                v15 := GenS1{F0: 32767, F1: 32767, F2: 0}
                _ = v15
            default:
                for fi16 := 0; fi16 < 12; fi16++ {
                    p1 = (^(p1))
                    switch fi16 {
                    case 127:
                        if true {
                            v17 := GenS2{F0: 65535, F1: 818}
                            for fi18 := 0; fi18 < 8; fi18++ {
                                var v19 uint16 = 0
                                v20 := GenS2{F0: 18446744073709551615, F1: 847}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = v17
                        } else {
                            var v21 uint32 = 4294967295
                            var v22 float32 = (float32(-2.25) / v14)
                            _ = v21
                            _ = v22
                        }
                    case -232:
                        if ((-2.25 * 1e10) >= (float64(float32(1.5)) / -2.25)) {
                            for fi23 := 0; fi23 < 9; fi23++ {
                                var v24 int16 = -648
                                if ((p1 * p1) < (^(p1))) {
                                    switch p2 {
                                    case "hello":
                                        var v25 int64 = 9223372036854775807
                                        _ = v25
                                    }
                                } else {
                                    v14 += 0.1
                                    p0 = (p0 >> ((0) & 63))
                                }
                                var v26 uint8 = (^(p1))
                                _ = fi23
                                _ = v24
                                _ = v26
                            }
                            p0 = (p0 + (p0 - p0))
                            for fi27 := 0; fi27 < 4; fi27++ {
                                if true {
                                    if (!false) {
                                        fmt.Println(p2, p2, v14)
                                        p1 = p1
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                                if (!(true && false)) {
                                    if true {
                                        fmt.Println(p1)
                                    }
                                    var v28 int32 = (p0 ^ p0)
                                    _ = v28
                                } else {
                                    p0 -= (-(p0))
                                }
                                _ = fi27
                            }
                        } else {
                            switch p2 {
                            case "hello":
                                if true {
                                    if (!false) {
                                        for fi29 := 0; fi29 < 9; fi29++ {
                                            if false {
                                                v30 := GenS3{F0: 255, F1: -418, F2: -1}
                                                p2 = (("hello" + p2) + "foo bar")
                                                _ = v30
                                            }
                                            _ = fi29
                                        }
                                        fmt.Println(fi16)
                                    } else {
                                        var v31 int64 = -9223372036854775808
                                        if true {
                                            if ((false && false) || true) {
                                                switch fi16 {
                                                case -441:
                                                    for fi32 := 0; fi32 < 3; fi32++ {
                                                        var v33 string = p2
                                                        v34 := GenS1{F0: 0, F1: 1, F2: 42}
                                                        var v35 uint = 255
                                                        v33 = "foo bar"
                                                        switch v33 {
                                                        case "\n\t":
                                                            v6 = append(v6, (fi16 % ((0) | 1)))
                                                        default:
                                                            if ((^(fi16)) > -80) {
                                                                var v36 uint64 = 526
                                                                p2 = "a"
                                                                _ = v36
                                                            }
                                                        }
                                                        v37 := GenS2{F0: 0, F1: 609}
                                                        _ = fi32
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                        _ = v37
                                                    }
                                                case 32767:
                                                    switch (p2 + "foo bar") {
                                                    case "\n\t":
                                                        v38 := GenS2{F0: 18446744073709551615, F1: 1}
                                                        _ = v38
                                                    case "a":
                                                        for fi39 := 0; fi39 < 14; fi39++ {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v31 = 1
                                                            _ = fi39
                                                        }
                                                    case "foo bar":
                                                        var v40 int32 = (p0 - (^(p0)))
                                                        _ = v40
                                                    default:
                                                        v14 = v14
                                                    }
                                                case 2147483647:
                                                    v41 := GenS4{F0: 627, F1: 0, F2: 494}
                                                    _ = v41
                                                default:
                                                    var v42 int = (^(fi16))
                                                    _ = v42
                                                }
                                            } else {
                                                var v43 int8 = 56
                                                var v44 uint8 = (p1 - 253)
                                                for fi45 := 0; fi45 < 19; fi45++ {
                                                    if (!true) {
                                                        v31 += v31
                                                        v46 := GenS5{F0: 2147483647}
                                                        _ = v46
                                                    } else {
                                                        v44 = v44
                                                        p1 = 1
                                                        var v47 uint = uint(v44)
                                                        _ = v47
                                                    }
                                                    v48 := GenS6{F0: 42, F1: v44}
                                                    _ = fi45
                                                    _ = v48
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                            switch p2 {
                                            case "":
                                                for fi49 := 0; fi49 < 11; fi49++ {
                                                    var v50 int64 = v31
                                                    if (0 <= 60) {
                                                        if (2147483647 >= (fi16 ^ fi49)) {
                                                            if false {
                                                                var v51 int32 = p0
                                                                fmt.Println(v31, p1)
                                                                if (false || (4294967295 != 333)) {
                                                                    v52 := GenS3{F0: p1, F1: -9223372036854775808, F2: 32767}
                                                                    var v53 uint = 1
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                _ = v51
                                                            } else {
                                                                var v54 bool = (1 == 42)
                                                                _ = v54
                                                            }
                                                            v31 = (-(v31))
                                                        }
                                                        v50 -= (v31 * (v31 >> ((810) & 63)))
                                                    }
                                                    v55 := GenS7{F0: (fi16 - fi49), F1: ((100 > 60) || true)}
                                                    _ = fi49
                                                    _ = v50
                                                    _ = v55
                                                }
                                            case "\n\t":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                _ = len(v6)
                                            }
                                            v14 = ((3.14159 - v14) / v14)
                                        } else {
                                            _ = len(v6)
                                        }
                                        _ = v31
                                    }
                                    var v56 bool = (p0 > (p0 / (((p0 >> ((0) & 63))) | 1)))
                                    _ = v56
                                } else {
                                    fmt.Println(p0, v14, fi16)
                                    var v57 int16 = -32768
                                    p0 -= p0
                                    _ = v57
                                }
                            case "\n\t":
                                p0 = p0
                            case "a":
                                v6 = append(v6, int(p1))
                            default:
                                _ = len(v6)
                            }
                            fmt.Println(fi16, fi16)
                        }
                    case 0:
                        if (!(p2 < "")) {
                            v6 = append(v6, (fi16 | (-(fi16))))
                            for fi58 := 0; fi58 < 7; fi58++ {
                                fmt.Println(p2)
                                var v59 int64 = -388
                                var v60 uint64 = 984
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            var v61 bool = ((p2 + "a") >= ((p2 + ("" + "a")) + p2))
                            _ = v61
                        } else {
                            if ((p1 * (^(p1))) < p1) {
                                for fi62 := 0; fi62 < 20; fi62++ {
                                    v14 = (v14 + v14)
                                    fmt.Println(p1, p0, p2)
                                    _ = fi62
                                }
                                for fi63 := 0; fi63 < 15; fi63++ {
                                    p2 = "\n\t"
                                    v6 = append(v6, fi16)
                                    v64 := GenS5{F0: 4294967295}
                                    _ = fi63
                                    _ = v64
                                }
                            }
                            _ = len(v6)
                        }
                    }
                    var v65 string = p2
                    _ = fi16
                    _ = v65
                }
            }
            var v66 uint = 2
            _ = v14
            _ = v66
        } else {
            if true {
                var v67 int16 = -162
                var v68 bool = (-1 < 32767)
                for fi69 := 0; fi69 < 10; fi69++ {
                    var v70 uint = 4294967295
                    var v71 int32 = int32(v70)
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                _ = v67
                _ = v68
            } else {
                p1 += 253
                var v72 bool = false
                _ = v72
            }
            p1 -= p1
            for fi73 := 0; fi73 < 19; fi73++ {
                for fi74 := 0; fi74 < 11; fi74++ {
                    if (false && (p0 == (p0 / ((-1) | 1)))) {
                        var v75 int = fi73
                        v76 := GenS8{F0: (p1 % ((124) | 1)), F1: 2}
                        var v77 uint16 = 2
                        _ = v75
                        _ = v76
                        _ = v77
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    fmt.Println(fi73)
                    _ = fi74
                }
                _ = fi73
            }
        }
        if (261 <= 32767) {
            fmt.Println(p0, p0)
            var v78 int8 = 42
            var v79 string = (("" + (((p2 + "\n\t") + (p2 + p2)) + (p2 + p2))) + p2)
            _ = v78
            _ = v79
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 >= (p0 / ((p0) | 1)))
}
func genFn1() string {
    switch (("foo bar" + "a") + ("a" + "hello")) {
    case "foo bar":
        switch 4294967295 {
        case 127:
            var v80 uint = 4294967295
            _ = v80
        case 4294967295:
            v81 := GenS2{F0: 65535, F1: 0}
            _ = v81
        default:
            v82 := GenS5{F0: 2147483647}
            _ = v82
        }
    case "":
        var v83 bool = (-100 <= 0)
        _ = v83
    case "a":
        var v84 int8 = 1
        _ = v84
    default:
        var v85 uint = 18446744073709551615
        _ = v85
    }
    var v86 uint = 42
    fmt.Println(v86)
    v87 := []int{-1, 1, 32768}
    _ = v86
    _ = v87
    return ""
}

func main() {
    v1 := []int{-841, -55}
    var v2 uint32 = 0
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    v5 := []int{32767, 128, 4294967295, 127}
    v2 = (v2 >> ((65535) & 63))
    for fi88 := 0; fi88 < 5; fi88++ {
        v2 = 1
        v2 = v2
        v4[(fi88 % ((fi88) | 1))] = true
        _ = fi88
    }
    if true {
        switch 128 {
        case 2147483647:
            v3 = ((0.0 * v3) / v3)
        default:
            var v89 int64 = -9223372036854775808
            _ = v89
        }
        v90 := GenS8{F0: 111, F1: v2}
        _ = v90
    }
    if (!(!((false || false) && ((9223372036854775807 == 42) && (806 == 42))))) {
        var v91 string = "a"
        _ = v91
    }
    v3 += (v3 / (v3 * float32(0.1)))
    if (((!(240 == 226)) || (false || (((false || false) && (!true)) && false))) && true) {
        for fi92 := 0; fi92 < 2; fi92++ {
            for fi93 := 0; fi93 < 20; fi93++ {
                switch (fi92 | fi92) {
                case -1:
                    v3 += ((float32((1e10 + (float64((((-2.25 * (0.0 * float32(((((float64(1e10) * 0.0) + float64(v3)) - -2.25) + -2.25)))) * float32((0.0 / 0.1))) * (float32(-2.25) + v3))) / (0.1 * -2.25)))) + (1e10 + -2.25)) + 1e10)
                }
                v2 += (v2 % ((207) | 1))
                var v94 string = ("a" + ("foo bar" + "a"))
                _ = fi93
                _ = v94
            }
            _ = fi92
        }
        var v95 int16 = 1
        if true {
            if (95 == 1) {
                v96 := GenS4{F0: 1, F1: (v95 | (^(v95))), F2: 19}
                _ = len(v5)
                _ = v96
            }
        } else {
            var v97 int16 = 0
            fmt.Println(v97, v2, v95)
            var v98 uint8 = 241
            _ = v97
            _ = v98
        }
        _ = v95
    } else {
        fmt.Println(v3, v2)
        for fi99 := 0; fi99 < 16; fi99++ {
            fmt.Println(v2)
            var v100 int = (^(fi99))
            _ = fi99
            _ = v100
        }
    }
    v2 -= (v2 & 4294967295)
    switch 4294967295 {
    case -411:
        v2 = 65535
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
