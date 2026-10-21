package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int64; F2 uint }
type GenS6 struct { F0 uint16; F1 int64; F2 bool }
type GenS7 struct { F0 uint8; F1 bool; F2 string }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int; F1 string; F2 bool }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 bool; F1 bool }
type GenS15 struct { F0 int; F1 uint32; F2 uint8 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint; F1 int }
type GenS18 struct { F0 int16; F1 uint16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int32; F1 int8 }
type GenS21 struct { F0 int64; F1 uint }
type GenS22 struct { F0 string; F1 uint; F2 uint }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 int64; F1 uint }
type GenS25 struct { F0 int64; F1 int8 }
type GenS26 struct { F0 string; F1 bool; F2 string }
type GenS27 struct { F0 bool; F1 int32; F2 string }
type GenS28 struct { F0 int; F1 int; F2 uint32 }
type GenS29 struct { F0 string; F1 string; F2 int32 }
type GenS30 struct { F0 uint; F1 int8; F2 int16 }
type GenS31 struct { F0 int32; F1 uint }
type GenS32 struct { F0 uint; F1 uint }
type GenS33 struct { F0 uint16 }
type GenS34 struct { F0 int; F1 int32; F2 uint16 }
type GenS35 struct { F0 int }
type GenS36 struct { F0 int8; F1 string }
func genFn0() string {
    var v7 bool = (42 == 0)
    if (((!v7) || (-2.25 >= 1.5)) || false) {
        v8 := []int{740, 0, 571, 474}
        v7 = ((true || (0 <= 942)) || (1 >= 284))
        _ = v8
    } else {
        v7 = ((-1 != 1) && false)
    }
    v7 = (1 != 117)
    _ = v7
    return "hello"
}
func genFn1(p0 uint, p1 uint64, p2 int16) uint32 {
    v9 := GenS1{F0: 127, F1: ("foo bar" + "foo bar")}
    v10 := GenS2{F0: -16}
    switch ("hello" + "hello") {
    case "a":
        fmt.Println(p0)
    case "\n\t":
        p1 = (p1 << (((p0 + p0)) & 63))
    case "foo bar":
        var v11 int = -842
        _ = v11
    default:
        var v12 uint32 = 2147483647
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{65535, 2, 4294967295, 4294967295}
    v2 := make(map[int]int)
    var v3 int32 = -1
    var v4 uint32 = 0
    var v5 bool = false
    var v6 bool = (1 > 9223372036854775807)
    if true {
        fmt.Println(v6)
    } else {
        v3 = (-(v3))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if true {
        switch 127 {
        case -9223372036854775808:
            if v6 {
                v13 := v2[-2147483648]
                var v14 int32 = (v3 * (v3 / (((v3 % ((v3) | 1))) | 1)))
                _ = v13
                _ = v14
            } else {
                var v15 uint8 = 255
                v16 := v2[9223372036854775807]
                _ = v15
                _ = v16
            }
        case -1:
            v17 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
            _ = v17
        default:
            switch "\n\t" {
            case "\n\t":
                v3 += v3
            case "foo bar":
                if v5 {
                    if ((-9223372036854775808 != 503) || ((!v6) || ((1 < 13) || (18446744073709551615 != 18446744073709551615)))) {
                        if v5 {
                            if true {
                                var v18 float64 = float64(1e10)
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    v5 = v6
                                    var v20 uint = 263
                                    _ = fi19
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                for fi21 := 0; fi21 < 20; fi21++ {
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        v5 = v6
                                        fmt.Println(fi22, fi21)
                                        _ = fi22
                                    }
                                    for fi23 := 0; fi23 < 8; fi23++ {
                                        v4 += 705
                                        switch fi21 {
                                        case 128:
                                            if true {
                                                var v24 int8 = -128
                                                _ = v24
                                            } else {
                                                var v25 float32 = 3.14159
                                                _ = v25
                                            }
                                        case 32767:
                                            var v26 uint16 = 1
                                            _ = v26
                                        case 2147483647:
                                            var v27 float64 = ((1.5 * (0.1 - 3.14159)) - 1.5)
                                            _ = v27
                                        }
                                        if false {
                                            _ = len(v2)
                                            _ = len(v2)
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483647
                                                }
                                                if ((!(fi23 >= (fi23 >> ((0) & 63)))) && v5) {
                                                    fmt.Println(v5, v3)
                                                }
                                                if v6 {
                                                    if (42 > 9223372036854775807) {
                                                        fmt.Println(fi23)
                                                    }
                                                    v6 = v6
                                                } else {
                                                    if (42 >= 1) {
                                                        var v28 uint16 = 149
                                                        if v5 {
                                                            fmt.Println(v6)
                                                        } else {
                                                            v6 = false
                                                            fmt.Println(fi21, v3, fi21)
                                                            if (fi23 < fi23) {
                                                                switch fi23 {
                                                                case 9223372036854775807:
                                                                    if (-128 >= -1) {
                                                                        var v29 int32 = v3
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v29
                                                                    } else {
                                                                        if v6 {
                                                                            if (42 >= 0) {
                                                                                v30 := GenS2{F0: (fi23 - fi21)}
                                                                                delete(v2, (fi23 / ((fi23) | 1)))
                                                                                _ = v30
                                                                            } else {
                                                                                var v31 uint = uint(v4)
                                                                                var v32 int64 = -9223372036854775808
                                                                                if (v4 == (v4 * v4)) {
                                                                                    var v33 uint = 0
                                                                                    v32 = (v32 << (((v33 + (v33 + 4294967295))) & 63))
                                                                                    switch int(v4) {
                                                                                    case 128:
                                                                                        if v5 {
                                                                                            v28 = (v28 - (v28 ^ v28))
                                                                                        }
                                                                                    case -515:
                                                                                        fmt.Println(v28, v4)
                                                                                    case -2147483648:
                                                                                        v2[(-(fi23))] = (-(fi23))
                                                                                        v1 = append(v1, (^(v4)))
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            if (float32((0.0 / (1e10 * 0.1))) > (1.5 / 0.1)) {
                                                                                                switch (-(fi21)) {
                                                                                                case 2147483647:
                                                                                                    if v6 {
                                                                                                        if (2 <= 2) {
                                                                                                            _ = len(v2)
                                                                                                        }
                                                                                                        switch -168 {
                                                                                                        case 42:
                                                                                                            v31 += 65535
                                                                                                        case 4294967295:
                                                                                                            switch fi23 {
                                                                                                            case -499:
                                                                                                                var v34 int64 = v32
                                                                                                                _ = v34
                                                                                                            default:
                                                                                                                switch "\n\t" {
                                                                                                                case "":
                                                                                                                    v2[(fi23 % (((fi23 >> ((917) & 63))) | 1))] = (fi21 ^ fi23)
                                                                                                                    var v35 uint32 = (^(v4))
                                                                                                                    _ = v35
                                                                                                                case "hello":
                                                                                                                    switch (-(fi21)) {
                                                                                                                    case 1:
                                                                                                                        fmt.Println(v5, v28)
                                                                                                                    default:
                                                                                                                        var v36 uint32 = v4
                                                                                                                        _ = v36
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    if v6 {
                                                                                                                        v6 = false
                                                                                                                        v33 -= v33
                                                                                                                        v2[fi21] = (fi23 % (((fi21 >> ((v33) & 63))) | 1))
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v37 uint32 = 42
                                                                                                                    _ = v37
                                                                                                                }
                                                                                                            }
                                                                                                        case 1:
                                                                                                            delete(v2, fi23)
                                                                                                        }
                                                                                                        var v38 float32 = 3.14159
                                                                                                        _ = v38
                                                                                                    }
                                                                                                case 128:
                                                                                                    var v39 uint64 = 2
                                                                                                    fmt.Println(v28, v3, v3)
                                                                                                    v6 = (uint64(v4) >= (v39 * (v39 - v39)))
                                                                                                    _ = v39
                                                                                                case 4294967295:
                                                                                                    delete(v2, fi21)
                                                                                                default:
                                                                                                    if ((v4 ^ (v4 << ((42) & 63))) != genFn1((v31 << (((v31 / ((v31) | 1))) & 63)), 937, 756)) {
                                                                                                        v2[0] = int(v4)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            if ((125 < 7) || (2147483647 > (fi23 >> ((v33) & 63)))) {
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v40 := GenS3{F0: ("hello" + "hello"), F1: ""}
                                                                                                    _ = v40
                                                                                                case "foo bar":
                                                                                                    var v41 int32 = (v3 % (((v3 % (((-(v3))) | 1))) | 1))
                                                                                                    _ = v41
                                                                                                }
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                v42 := GenS2{F0: fi21}
                                                                                                _ = v42
                                                                                            } else {
                                                                                                if ((v4 / ((genFn1(v33, 310, -135)) | 1)) <= (v4 ^ 1)) {
                                                                                                    v43 := GenS4{F0: "foo bar"}
                                                                                                    fmt.Println(v28, v4, v3)
                                                                                                    if false {
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = uint32(v33)
                                                                                                        }
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            delete(v2, fi21)
                                                                                                        case "hello":
                                                                                                            var v44 int32 = 487
                                                                                                            _ = v44
                                                                                                        }
                                                                                                        if v6 {
                                                                                                            if ((0.0 <= 3.14159) || (((!v5) && v6) && (!v5))) {
                                                                                                                v6 = v6
                                                                                                                v3 = int32(fi23)
                                                                                                                fmt.Println(v6)
                                                                                                            } else {
                                                                                                                if len(v1) >= 2 {
                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                }
                                                                                                                var v45 int = fi21
                                                                                                                var v46 int8 = 93
                                                                                                                _ = v45
                                                                                                                _ = v46
                                                                                                            }
                                                                                                            v33 -= 65535
                                                                                                            var v47 bool = v5
                                                                                                            _ = v47
                                                                                                        } else {
                                                                                                            var v48 float64 = 1e10
                                                                                                            v49 := GenS5{F0: int8(fi21), F1: (v32 + -9223372036854775808), F2: (v31 & 18446744073709551615)}
                                                                                                            switch fi23 {
                                                                                                            case 128:
                                                                                                                v4 = 731
                                                                                                            case -124:
                                                                                                                fmt.Println(v6, v48, v4)
                                                                                                            default:
                                                                                                                _ = len(v1)
                                                                                                            }
                                                                                                            _ = v48
                                                                                                            _ = v49
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v5, fi23)
                                                                                                        fmt.Println(v6)
                                                                                                        var v50 int = fi21
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    _ = v43
                                                                                                }
                                                                                                v1 = append(v1, v4)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        delete(v2, (fi23 - fi23))
                                                                                    }
                                                                                    _ = v33
                                                                                } else {
                                                                                    if (450 <= 42) {
                                                                                        fmt.Println(v6)
                                                                                    }
                                                                                    if v6 {
                                                                                        if ((v3 / (((^(v3))) | 1)) < int32(v28)) {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = uint32(v31)
                                                                                            }
                                                                                            v51 := GenS3{F0: "a", F1: "\n\t"}
                                                                                            var v52 int8 = 42
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        }
                                                                                        v3 = 0
                                                                                    } else {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v5 = (-1 <= 32767)
                                                                                    }
                                                                                }
                                                                                _ = v31
                                                                                _ = v32
                                                                            }
                                                                        }
                                                                    }
                                                                case 703:
                                                                    var v53 uint16 = v28
                                                                    _ = v53
                                                                default:
                                                                    v54 := v2[(^(fi23))]
                                                                    _ = v54
                                                                }
                                                            } else {
                                                                if ((v5 || v6) || (1 > 427)) {
                                                                    var v55 int64 = -9223372036854775808
                                                                    var v56 uint = 251
                                                                    fmt.Println(v56)
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                                var v57 int16 = 1
                                                                _ = v57
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = (v4 + (v4 << ((255) & 63)))
                                                            }
                                                            if ("\n\t" == ("a" + ("a" + ("\n\t" + "\n\t")))) {
                                                                var v58 float64 = 0.1
                                                                _ = v58
                                                            }
                                                        }
                                                        fmt.Println(v5, v3)
                                                        _ = v28
                                                    } else {
                                                        v3 = 1
                                                        v4 = v4
                                                    }
                                                    v59 := GenS1{F0: (fi21 % ((fi23) | 1)), F1: "hello"}
                                                    if true {
                                                        if (true && (9223372036854775807 > -762)) {
                                                            v5 = v5
                                                            if (42 < 42) {
                                                                switch (fi21 | (fi23 % (((fi23 << ((0) & 63))) | 1))) {
                                                                case -2147483648:
                                                                    v3 = v3
                                                                case -9223372036854775808:
                                                                    switch (fi21 * fi23) {
                                                                    case 335:
                                                                        v60 := GenS2{F0: (^(fi23))}
                                                                        _ = v60
                                                                    default:
                                                                        var v61 uint32 = 42
                                                                        var v62 bool = v5
                                                                        var v63 uint = 65535
                                                                        v4 -= genFn1(v63, 18446744073709551615, 32767)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        switch 42 {
                                                                        case 9223372036854775807:
                                                                            v62 = (v6 || v62)
                                                                        case 32768:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v61 + (v61 | (v4 | (^(v4)))))
                                                                            }
                                                                        }
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                case 32767:
                                                                    v3 += v3
                                                                default:
                                                                    v64 := GenS3{F0: (("foo bar" + "\n\t") + "hello"), F1: "hello"}
                                                                    _ = v64
                                                                }
                                                                v4 = 2147483647
                                                            } else {
                                                                var v65 uint16 = 42
                                                                var v66 int32 = v3
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        } else {
                                                            fmt.Println(fi21)
                                                            var v67 int16 = 32767
                                                            v2[(fi23 - (-(fi21)))] = (fi23 / (((fi23 | (^(fi21)))) | 1))
                                                            _ = v67
                                                        }
                                                        v68 := GenS6{F0: 0, F1: -1, F2: (!(!(v5 && (!(!(-9223372036854775808 <= 1))))))}
                                                        _ = v68
                                                    } else {
                                                        if true {
                                                            fmt.Println(v3)
                                                            v69 := GenS7{F0: 255, F1: (float32(((1e10 + 0.0) + -2.25)) == (0.1 + (-2.25 - 1.5))), F2: "\n\t"}
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v69
                                                        }
                                                    }
                                                    var v70 float32 = 1e10
                                                    var v71 uint64 = 4294967295
                                                    var v72 uint8 = 127
                                                    v73 := GenS8{F0: 741, F1: -1}
                                                    _ = v59
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v74 := GenS4{F0: ("\n\t" + (("" + "hello") + ""))}
                                                v75 := GenS6{F0: 0, F1: -60, F2: v5}
                                                switch (fi21 - fi21) {
                                                case 128:
                                                    if true {
                                                        v6 = false
                                                        v6 = (!false)
                                                    } else {
                                                        _ = len(v1)
                                                        fmt.Println(v6, v6, v6)
                                                        if (235 <= 127) {
                                                            var v76 bool = (29 >= -128)
                                                            _ = v76
                                                        } else {
                                                            if v6 {
                                                                var v77 string = (("hello" + "foo bar") + "\n\t")
                                                                v78 := GenS2{F0: fi21}
                                                                fmt.Println(fi21, v5, v6)
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            if v6 {
                                                                v79 := GenS4{F0: ("a" + ((("\n\t" + "hello") + ("" + "hello")) + ((((("" + "a") + "a") + "hello") + ("hello" + "\n\t")) + "hello")))}
                                                                _ = v79
                                                            } else {
                                                                switch "a" {
                                                                case "a":
                                                                    fmt.Println(v5, fi21, fi23)
                                                                default:
                                                                    v3 += 1
                                                                }
                                                                var v80 uint8 = uint8(v4)
                                                                if false {
                                                                    var v81 uint16 = 42
                                                                    v5 = (42 >= 0)
                                                                    _ = v81
                                                                }
                                                                if (65535 < 778) {
                                                                    if false {
                                                                        switch (fi21 * (^(fi23))) {
                                                                        case 128:
                                                                            if v6 {
                                                                                switch "a" {
                                                                                case "a":
                                                                                    if ((float64(-2.25) - 1e10) < (1e10 / float64(1e10))) {
                                                                                        var v82 uint32 = v4
                                                                                        var v83 string = ""
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                    }
                                                                                case "foo bar":
                                                                                    switch (fi21 & (fi23 ^ (fi21 - fi21))) {
                                                                                    case -9223372036854775808:
                                                                                        switch 127 {
                                                                                        case 565:
                                                                                            switch ("hello" + "") {
                                                                                            case "a":
                                                                                                v84 := GenS9{F0: 667}
                                                                                                _ = v84
                                                                                            case "":
                                                                                                v3 = (-(v3))
                                                                                            case "foo bar":
                                                                                                var v85 int = (fi21 | fi21)
                                                                                                _ = v85
                                                                                            }
                                                                                        }
                                                                                    case 32767:
                                                                                        v5 = true
                                                                                    case -740:
                                                                                        fmt.Println(v3, v3)
                                                                                    }
                                                                                case "\n\t":
                                                                                    v3 = 1
                                                                                default:
                                                                                    var v86 int64 = -9223372036854775808
                                                                                    _ = v86
                                                                                }
                                                                                v80 = (v80 << ((65535) & 63))
                                                                            } else {
                                                                                v6 = (255 < 255)
                                                                                var v87 string = ("\n\t" + "\n\t")
                                                                                var v88 int16 = 0
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                        case -251:
                                                                            v89 := v2[fi23]
                                                                            _ = v89
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(v3, v4)
                                                                                var v90 float32 = 1.5
                                                                                if (1 == 615) {
                                                                                    var v91 int16 = 0
                                                                                    _ = v91
                                                                                } else {
                                                                                    v92 := GenS10{F0: v3, F1: (fi21 + (fi21 & (fi23 % (((fi23 >> ((18446744073709551615) & 63))) | 1))))}
                                                                                    v6 = true
                                                                                    _ = v92
                                                                                }
                                                                                _ = v90
                                                                            } else {
                                                                                fmt.Println(fi21, fi21)
                                                                                v6 = (!v6)
                                                                                v93 := GenS5{F0: int8(v80), F1: -1, F2: 255}
                                                                                var v94 int16 = 0
                                                                                v6 = v6
                                                                                _ = v93
                                                                                _ = v94
                                                                            }
                                                                        }
                                                                        delete(v2, -264)
                                                                    } else {
                                                                        fmt.Println(fi23)
                                                                        v5 = v5
                                                                        var v95 int8 = 100
                                                                        _ = v95
                                                                    }
                                                                    if ("hello" >= ("a" + "hello")) {
                                                                        _ = len(v1)
                                                                        var v96 int32 = -862
                                                                        if false {
                                                                            switch ("foo bar" + (("a" + "a") + "foo bar")) {
                                                                            case "":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            var v97 int16 = -32768
                                                                            _ = v97
                                                                        }
                                                                        if true {
                                                                            v96 += -1
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v4
                                                                            }
                                                                            if v5 {
                                                                                var v98 int32 = 2147483647
                                                                                var v99 int8 = 13
                                                                                v5 = (!(7 > (^(v80))))
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                        _ = v96
                                                                    }
                                                                }
                                                                _ = v80
                                                            }
                                                        }
                                                    }
                                                case -512:
                                                    if (42 == (v4 | v4)) {
                                                        v3 -= v3
                                                        v6 = v6
                                                        var v100 int16 = 32767
                                                        var v101 int64 = 9223372036854775807
                                                        switch "a" {
                                                        case "":
                                                            var v102 float64 = float64(0.1)
                                                            _ = v102
                                                        }
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                case 0:
                                                    fmt.Println(v5)
                                                default:
                                                    var v103 uint = 1
                                                    _ = v103
                                                }
                                                _ = v74
                                                _ = v75
                                            }
                                            if (0 >= (v4 % (((v4 + v4)) | 1))) {
                                                _ = len(v1)
                                                v3 -= v3
                                                switch "hello" {
                                                case "\n\t":
                                                    switch "a" {
                                                    case "\n\t":
                                                        fmt.Println(v5)
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            v104 := GenS4{F0: ("hello" + "foo bar")}
                                                            _ = v104
                                                        case "foo bar":
                                                            switch ("hello" + ("" + (("a" + "") + "foo bar"))) {
                                                            case "hello":
                                                                v6 = false
                                                            }
                                                        default:
                                                            v4 = (v4 - (v4 / ((522) | 1)))
                                                        }
                                                        v3 += (v3 - 2147483647)
                                                        v105 := v2[fi21]
                                                        _ = v105
                                                    }
                                                case "a":
                                                    v106 := GenS11{F0: -20}
                                                    v5 = (322 >= 42)
                                                    _ = v106
                                                case "":
                                                    v107 := v2[fi23]
                                                    switch "" {
                                                    case "":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "foo bar":
                                                        var v108 int32 = (^(v3))
                                                        _ = v108
                                                    case "a":
                                                        var v109 float32 = -2.25
                                                        _ = v109
                                                    }
                                                    var v110 float64 = 1e10
                                                    _ = v107
                                                    _ = v110
                                                }
                                                switch fi21 {
                                                case 9223372036854775807:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v111 := GenS12{F0: fi21, F1: ("hello" + ""), F2: v5}
                                                        v6 = (!v5)
                                                        _ = v111
                                                    case "\n\t":
                                                        v3 += -1
                                                    case "":
                                                        v1 = append(v1, 65535)
                                                    default:
                                                        v112 := GenS13{F0: v6}
                                                        _ = v112
                                                    }
                                                case 575:
                                                    v113 := GenS14{F0: (115 != -128), F1: true}
                                                    _ = v113
                                                }
                                                fmt.Println(v3, v6)
                                            }
                                            if ((-1 >= 0) && v5) {
                                                var v114 int16 = -818
                                                _ = v114
                                            }
                                        }
                                        if v5 {
                                            var v115 uint = 2
                                            v5 = (!v5)
                                            var v116 int16 = int16(v3)
                                            _ = v115
                                            _ = v116
                                        }
                                        v4 -= v4
                                        _ = fi23
                                    }
                                    if v6 {
                                        if (1 >= (^(v4))) {
                                            var v117 uint = 421
                                            var v118 int8 = 42
                                            for fi119 := 0; fi119 < 16; fi119++ {
                                                fmt.Println(v5, v117, v5)
                                                v5 = v5
                                                _ = fi119
                                            }
                                            _ = v117
                                            _ = v118
                                        } else {
                                            var v120 uint32 = 65535
                                            v121 := GenS15{F0: 1, F1: v4, F2: 1}
                                            _ = v120
                                            _ = v121
                                        }
                                        var v122 uint = 18446744073709551615
                                        _ = v122
                                    } else {
                                        v123 := GenS4{F0: "hello"}
                                        switch ("hello" + ("hello" + ((((("hello" + "a") + "a") + "foo bar") + ("\n\t" + "foo bar")) + "a"))) {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "hello":
                                            for fi124 := 0; fi124 < 4; fi124++ {
                                                switch 421 {
                                                case 1:
                                                    fmt.Println(v6, fi21)
                                                case -9223372036854775808:
                                                    var v125 int = -1
                                                    _ = v125
                                                default:
                                                    _ = len(v2)
                                                }
                                                var v126 int8 = 127
                                                v127 := GenS16{F0: 1}
                                                _ = fi124
                                                _ = v126
                                                _ = v127
                                            }
                                        case "foo bar":
                                            if (!false) {
                                                for fi128 := 0; fi128 < 17; fi128++ {
                                                    if false {
                                                        var v129 string = "foo bar"
                                                        var v130 uint16 = 32767
                                                        _ = v129
                                                        _ = v130
                                                    } else {
                                                        if ((!v5) || false) {
                                                            v6 = v5
                                                            v131 := GenS13{F0: v5}
                                                            var v132 int = 32768
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                        switch (fi128 | (fi21 & int(v4))) {
                                                        case -877:
                                                            v6 = true
                                                        case -1:
                                                            if len(v1) > 0 {
                                                                v1[0] = (v4 >> ((18446744073709551615) & 63))
                                                            }
                                                        }
                                                    }
                                                    v133 := GenS13{F0: ((!(true && ("\n\t" >= "\n\t"))) && (!(((float64((1.5 * -2.25)) - -2.25) + 0.1) <= 0.1)))}
                                                    v4 -= 2
                                                    _ = fi128
                                                    _ = v133
                                                }
                                                for fi134 := 0; fi134 < 19; fi134++ {
                                                    _ = len(v2)
                                                    _ = fi134
                                                }
                                                fmt.Println(v5, fi21, v5)
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    v135 := GenS8{F0: 645, F1: -127}
                                                    _ = v135
                                                case "\n\t":
                                                    switch ((("foo bar" + "foo bar") + "a") + ("hello" + "")) {
                                                    case "":
                                                        switch (fi21 ^ (-(fi21))) {
                                                        case 127:
                                                            v136 := v2[fi21]
                                                            _ = v136
                                                        case 128:
                                                            switch "hello" {
                                                            case "":
                                                                var v137 int8 = -113
                                                                _ = v137
                                                            case "foo bar":
                                                                for fi138 := 0; fi138 < 8; fi138++ {
                                                                    var v139 uint16 = 0
                                                                    v140 := GenS17{F0: 2, F1: fi21}
                                                                    _ = fi138
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            case "a":
                                                                v3 -= (v3 % ((0) | 1))
                                                            default:
                                                                v5 = true
                                                            }
                                                        default:
                                                            for fi141 := 0; fi141 < 7; fi141++ {
                                                                v142 := GenS16{F0: genFn1(0, 2, 32767)}
                                                                _ = fi141
                                                                _ = v142
                                                            }
                                                        }
                                                    case "\n\t":
                                                        delete(v2, (fi21 & -177))
                                                    case "foo bar":
                                                        for fi143 := 0; fi143 < 7; fi143++ {
                                                            if v5 {
                                                                delete(v2, fi21)
                                                                var v144 int16 = -630
                                                                if v6 {
                                                                    v145 := GenS13{F0: (!v6)}
                                                                    if true {
                                                                        if v6 {
                                                                            var v146 float32 = (((1.5 - ((1e10 + ((0.0 - 1e10) - 3.14159)) / (3.14159 + 3.14159))) + 0.0) * ((-2.25 / (1.5 * -2.25)) * 1.5))
                                                                            var v147 int8 = -128
                                                                            v4 = v4
                                                                            _ = v146
                                                                            _ = v147
                                                                        } else {
                                                                            v3 -= v3
                                                                        }
                                                                    }
                                                                    if v5 {
                                                                        var v148 uint32 = v4
                                                                        _ = v148
                                                                    }
                                                                    if (!(true || v6)) {
                                                                        var v149 float64 = (-2.25 / 1.0)
                                                                        v6 = true
                                                                        _ = v149
                                                                    } else {
                                                                        v5 = (9223372036854775807 != -191)
                                                                        switch ("\n\t" + "foo bar") {
                                                                        case "foo bar":
                                                                            var v150 uint64 = 18446744073709551615
                                                                            _ = v150
                                                                        case "a":
                                                                            switch (fi143 | fi143) {
                                                                            case 128:
                                                                                if true {
                                                                                    var v151 uint8 = 0
                                                                                    var v152 bool = false
                                                                                    _ = v151
                                                                                    _ = v152
                                                                                } else {
                                                                                    if false {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v153 := GenS18{F0: -51, F1: 0}
                                                                                        _ = v153
                                                                                    }
                                                                                    v154 := v2[128]
                                                                                    v3 = v3
                                                                                    _ = v154
                                                                                }
                                                                            case -1:
                                                                                fmt.Println(v5, v6)
                                                                            default:
                                                                                var v155 float64 = (3.14159 + 3.14159)
                                                                                _ = v155
                                                                            }
                                                                        case "\n\t":
                                                                            switch 1 {
                                                                            case 611:
                                                                                var v156 int8 = 0
                                                                                _ = v156
                                                                            }
                                                                        }
                                                                    }
                                                                    if ((v4 + 42) <= 0) {
                                                                        v157 := GenS19{F0: (v4 >> ((0) & 63))}
                                                                        var v158 uint = 42
                                                                        _ = v157
                                                                        _ = v158
                                                                    }
                                                                    v3 += (^(v3))
                                                                    _ = v145
                                                                } else {
                                                                    var v159 string = (("" + "a") + "\n\t")
                                                                    _ = v159
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = v4
                                                                }
                                                                v144 = (v144 | (v144 * -32768))
                                                                v3 += v3
                                                                _ = v144
                                                            }
                                                            v5 = ((v6 && (((3.14159 * float32(float64(1e10))) / 1.0) > ((3.14159 * (float32(-2.25) / 1.5)) + ((0.1 / 0.1) * 1e10)))) && true)
                                                            v160 := GenS7{F0: 3, F1: v5, F2: ("\n\t" + "a")}
                                                            var v161 float64 = 0.1
                                                            if (v161 == (v161 * v161)) {
                                                                v5 = false
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    if (!v6) {
                                                                        var v162 uint32 = genFn1(65535, 2, -32768)
                                                                        _ = v162
                                                                    } else {
                                                                        v3 += 1
                                                                    }
                                                                case "":
                                                                    fmt.Println(v6)
                                                                case "a":
                                                                    var v163 int16 = -190
                                                                    _ = v163
                                                                default:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            }
                                                            _ = fi143
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                    default:
                                                        v164 := GenS14{F0: v6, F1: (v5 || (91 > 133))}
                                                        _ = v164
                                                    }
                                                }
                                                for fi165 := 0; fi165 < 20; fi165++ {
                                                    v5 = v5
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if false {
                                                            switch (fi165 >> ((65535) & 63)) {
                                                            case 0:
                                                                v166 := GenS4{F0: ""}
                                                                _ = v166
                                                            default:
                                                                var v167 bool = v6
                                                                _ = v167
                                                            }
                                                            v168 := GenS20{F0: v3, F1: 100}
                                                            _ = v168
                                                        } else {
                                                            delete(v2, (fi165 - 4294967295))
                                                            v3 = (v3 & v3)
                                                        }
                                                    case "hello":
                                                        var v169 uint8 = 25
                                                        _ = v169
                                                    default:
                                                        if ((float64((0.1 - 0.1)) != 1e10) || (4294967295 > 2147483648)) {
                                                            var v170 bool = (!false)
                                                            _ = v170
                                                        } else {
                                                            var v171 float32 = 1e10
                                                            var v172 int = 0
                                                            _ = v171
                                                            _ = v172
                                                        }
                                                    }
                                                    _ = fi165
                                                }
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = (v4 << ((25) & 63))
                                                    }
                                                    if (v4 < 42) {
                                                        v173 := GenS13{F0: true}
                                                        _ = len(v2)
                                                        switch fi21 {
                                                        case 127:
                                                            for fi174 := 0; fi174 < 14; fi174++ {
                                                                v1 = append(v1, v4)
                                                                v175 := v2[42]
                                                                fmt.Println(fi21)
                                                                if v5 {
                                                                    var v176 bool = (!((!(v5 && (1 > 9223372036854775807))) || true))
                                                                    if (!v176) {
                                                                        v5 = ((true && true) && (-614 != 0))
                                                                    }
                                                                    _ = v176
                                                                }
                                                                v6 = false
                                                                _ = fi174
                                                                _ = v175
                                                            }
                                                        }
                                                        _ = v173
                                                    } else {
                                                        v2[fi21] = (^(fi21))
                                                        switch "" {
                                                        case "":
                                                            var v177 int8 = 5
                                                            _ = v177
                                                        case "foo bar":
                                                            for fi178 := 0; fi178 < 15; fi178++ {
                                                                var v179 uint8 = 42
                                                                v179 += v179
                                                                v1 = append(v1, v4)
                                                                _ = fi178
                                                                _ = v179
                                                            }
                                                        }
                                                    }
                                                    v2[(fi21 + (fi21 | fi21))] = (^(fi21))
                                                } else {
                                                    switch fi21 {
                                                    case 400:
                                                        for fi180 := 0; fi180 < 10; fi180++ {
                                                            var v181 int = fi21
                                                            _ = fi180
                                                            _ = v181
                                                        }
                                                    case 4294967295:
                                                        if v5 {
                                                            if (!(0.1 > -2.25)) {
                                                                var v182 int = fi21
                                                                var v183 int8 = -128
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v184 := GenS21{F0: 0, F1: 4294967295}
                                                                    _ = v184
                                                                case "hello":
                                                                    v1 = append(v1, (v4 & (^(v4))))
                                                                case "a":
                                                                    var v185 int = v182
                                                                    _ = v185
                                                                default:
                                                                    v2[v182] = (^(v182))
                                                                }
                                                                _ = v182
                                                                _ = v183
                                                            } else {
                                                                var v186 float64 = float64(float32(0.1))
                                                                var v187 bool = v6
                                                                if (!true) {
                                                                    fmt.Println(v186)
                                                                } else {
                                                                    for fi188 := 0; fi188 < 20; fi188++ {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (v4 + (^(v4)))
                                                                        }
                                                                        v5 = (1e10 == (0.0 + v186))
                                                                        _ = fi188
                                                                    }
                                                                }
                                                                _ = v186
                                                                _ = v187
                                                            }
                                                            fmt.Println(fi21, v3, v6)
                                                            fmt.Println(v5, v3, fi21)
                                                        } else {
                                                            fmt.Println(fi21, v3)
                                                        }
                                                    default:
                                                        if (v6 || (v6 || v5)) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v3 -= 1
                                                            case "":
                                                                if ((v5 || v5) || (!v6)) {
                                                                    fmt.Println(v5, v4, fi21)
                                                                } else {
                                                                    v6 = (42 == -972)
                                                                }
                                                            }
                                                            if v5 {
                                                                var v189 int16 = 631
                                                                v3 = v3
                                                                var v190 uint32 = (v4 & genFn1(4294967295, 42, (v189 * (v189 * (v189 & (^(v189)))))))
                                                                _ = v189
                                                                _ = v190
                                                            } else {
                                                                var v191 int16 = 752
                                                                _ = v191
                                                            }
                                                            if ((v3 >= v3) || (!(v3 >= (-(v3))))) {
                                                                for fi192 := 0; fi192 < 20; fi192++ {
                                                                    fmt.Println(v5)
                                                                    if ("" <= "foo bar") {
                                                                        v3 = (v3 << ((uint(v4)) & 63))
                                                                        switch ((("hello" + "a") + ("a" + "foo bar")) + "") {
                                                                        case "":
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        case "foo bar":
                                                                            v4 += (v4 << ((65535) & 63))
                                                                        }
                                                                    } else {
                                                                        var v193 uint8 = 83
                                                                        _ = len(v2)
                                                                        _ = v193
                                                                    }
                                                                    _ = fi192
                                                                }
                                                                v3 = -2147483648
                                                            }
                                                        } else {
                                                            if ((1e10 < (1.5 + (float32(0.0) * -2.25))) || (1 <= 121)) {
                                                                if ((v4 & genFn1(729, 550, 244)) >= (^(v4))) {
                                                                    var v194 int32 = (v3 - (-(v3)))
                                                                    v195 := GenS22{F0: "\n\t", F1: 2, F2: 4294967295}
                                                                    v3 = (v194 + v3)
                                                                    _ = v194
                                                                    _ = v195
                                                                }
                                                                for fi196 := 0; fi196 < 18; fi196++ {
                                                                    var v197 uint8 = 42
                                                                    fmt.Println(v4, v4)
                                                                    _ = fi196
                                                                    _ = v197
                                                                }
                                                            }
                                                            if false {
                                                                if (v3 > (-(v3))) {
                                                                    var v198 int8 = 1
                                                                    v4 += (v4 - (^(v4)))
                                                                    _ = v198
                                                                }
                                                                var v199 int = (fi21 + (fi21 >> ((18446744073709551615) & 63)))
                                                                v6 = (!false)
                                                                _ = v199
                                                            } else {
                                                                v200 := GenS23{F0: (2 > 198)}
                                                                _ = v200
                                                            }
                                                            v201 := v2[fi21]
                                                            _ = v201
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            v6 = v6
                                        }
                                        _ = v123
                                    }
                                    _ = fi21
                                }
                                if len(v1) > 0 {
                                    v1[0] = (^(v4))
                                }
                            }
                            var v202 uint32 = (v4 >> ((0) & 63))
                            _ = v202
                        }
                    } else {
                        switch 244 {
                        case -955:
                            if (1000 == 1000) {
                                v203 := GenS5{F0: 1, F1: int64(v3), F2: 18446744073709551615}
                                v204 := GenS16{F0: 4294967295}
                                _ = len(v2)
                                _ = v203
                                _ = v204
                            } else {
                                v205 := GenS14{F0: false, F1: ((0 == -274) && (0 <= 1))}
                                v206 := GenS1{F0: 128, F1: ("\n\t" + "a")}
                                _ = v205
                                _ = v206
                            }
                        case -1:
                            delete(v2, int(v3))
                        case 1:
                            var v207 uint32 = (v4 + (^(v4)))
                            _ = v207
                        default:
                            fmt.Println(v4, v3, v3)
                        }
                        var v208 int16 = 0
                        _ = v208
                    }
                    if (3.14159 == 1.5) {
                        _ = len(v2)
                    } else {
                        v209 := GenS11{F0: 42}
                        v4 = genFn1(uint(v4), 278, 32767)
                        for fi210 := 0; fi210 < 2; fi210++ {
                            if len(v1) > 0 {
                                v1[0] = 352
                            }
                            _ = fi210
                        }
                        _ = v209
                    }
                } else {
                    fmt.Println(v4)
                    v3 = (v3 * v3)
                }
            case "hello":
                v4 += v4
            default:
                var v211 uint16 = 0
                _ = v211
            }
        }
        switch 0 {
        case -877:
            v212 := GenS24{F0: -9223372036854775808, F1: 2}
            _ = v212
        }
        if false {
            for fi213 := 0; fi213 < 7; fi213++ {
                delete(v2, fi213)
                var v214 uint8 = 209
                v215 := v2[32768]
                _ = fi213
                _ = v214
                _ = v215
            }
            for fi216 := 0; fi216 < 3; fi216++ {
                var v217 bool = v6
                switch "" {
                case "\n\t":
                    for fi218 := 0; fi218 < 11; fi218++ {
                        v219 := GenS2{F0: fi218}
                        if (v217 && v6) {
                            fmt.Println(fi216, v3, fi218)
                            v5 = v6
                        } else {
                            fmt.Println(fi216, v6)
                            var v220 float64 = (0.1 * (0.1 - -2.25))
                            fmt.Println(v217)
                            _ = len(v1)
                            _ = v220
                        }
                        v3 = v3
                        _ = fi218
                        _ = v219
                    }
                case "foo bar":
                    v221 := GenS18{F0: 0, F1: 429}
                    _ = v221
                case "a":
                    delete(v2, fi216)
                }
                v4 += (v4 | (v4 >> ((691) & 63)))
                _ = fi216
                _ = v217
            }
        } else {
            v222 := GenS18{F0: -1, F1: 42}
            var v223 float64 = (float64(0.1) / 0.1)
            if v6 {
                v224 := GenS16{F0: (^(v4))}
                if (42 != -945) {
                    if (true || v6) {
                        var v225 int64 = -900
                        switch 2147483647 {
                        case -9223372036854775808:
                            if v5 {
                                switch -444 {
                                case 0:
                                    v3 = (v3 * (v3 * (v3 & (v3 >> ((42) & 63)))))
                                case 32768:
                                    var v226 int32 = (v3 | 1)
                                    _ = v226
                                case -807:
                                    v227 := GenS25{F0: (v225 + v225), F1: 45}
                                    _ = v227
                                default:
                                    v1 = append(v1, 2147483647)
                                }
                            }
                        case 1:
                            v3 = -1
                        }
                        if true {
                            v223 += float64(3.14159)
                        } else {
                            _ = len(v1)
                        }
                        _ = v225
                    }
                    v228 := GenS26{F0: "\n\t", F1: (78 == 42), F2: (("" + "a") + "a")}
                    var v229 int8 = 127
                    _ = v228
                    _ = v229
                }
                if (v6 && false) {
                    if ((220 == 1) || false) {
                        v1 = append(v1, 2147483648)
                        if (!(!(!true))) {
                            fmt.Println(v6, v4, v223)
                        } else {
                            v230 := GenS27{F0: (1 <= -117), F1: v3, F2: ""}
                            var v231 int16 = 42
                            var v232 uint16 = 770
                            _ = v230
                            _ = v231
                            _ = v232
                        }
                        if ((v3 << ((42) & 63)) < (v3 / ((-801) | 1))) {
                            var v233 uint8 = 2
                            _ = v233
                        } else {
                            v4 = genFn1(65535, 1, 32767)
                            if v5 {
                                fmt.Println(v3, v3)
                                var v234 string = ("\n\t" + (("\n\t" + "\n\t") + "\n\t"))
                                for fi235 := 0; fi235 < 19; fi235++ {
                                    switch (^(fi235)) {
                                    case -1:
                                        var v236 string = (v234 + v234)
                                        _ = v236
                                    case -261:
                                        var v237 int = -1
                                        _ = v237
                                    case 32767:
                                        switch v234 {
                                        case "":
                                            v238 := GenS25{F0: -1, F1: 0}
                                            _ = v238
                                        case "\n\t":
                                            for fi239 := 0; fi239 < 4; fi239++ {
                                                var v240 uint8 = 127
                                                if ((^(fi239)) <= (fi235 % (((^(fi235))) | 1))) {
                                                    var v241 uint8 = (^(v240))
                                                    v2[(fi239 / ((42) | 1))] = (fi235 - (fi235 / ((fi235) | 1)))
                                                    _ = v241
                                                }
                                                _ = fi239
                                                _ = v240
                                            }
                                        case "a":
                                            v242 := GenS13{F0: v6}
                                            _ = v242
                                        }
                                    default:
                                        var v243 bool = v6
                                        _ = v243
                                    }
                                    v234 = "foo bar"
                                    _ = fi235
                                }
                                _ = v234
                            } else {
                                var v244 uint32 = (v4 % ((v4) | 1))
                                _ = v244
                            }
                            var v245 float32 = (float32((((v223 / (v223 * v223)) - (v223 / ((3.14159 - v223) * float64(1.5)))) * v223)) + 1.5)
                            _ = v245
                        }
                    } else {
                        if false {
                            fmt.Println(v5, v5, v3)
                        } else {
                            v5 = (v5 && ((-774 > -51) && (true && true)))
                        }
                        if false {
                            if v5 {
                                var v246 float64 = (1e10 - v223)
                                v246 -= -2.25
                                _ = v246
                            } else {
                                switch 9223372036854775807 {
                                case 305:
                                    var v247 uint64 = 1
                                    _ = v247
                                case -24:
                                    if true {
                                        v223 = v223
                                        var v248 int8 = 100
                                        _ = v248
                                    }
                                case 4294967295:
                                    var v249 string = ("hello" + "a")
                                    _ = v249
                                }
                                for fi250 := 0; fi250 < 20; fi250++ {
                                    v5 = true
                                    _ = fi250
                                }
                                for fi251 := 0; fi251 < 13; fi251++ {
                                    v252 := GenS28{F0: (fi251 + (fi251 / ((fi251) | 1))), F1: fi251, F2: genFn1(65535, uint64(v4), 32767)}
                                    _ = fi251
                                    _ = v252
                                }
                            }
                            _ = len(v2)
                            v253 := GenS2{F0: -9223372036854775808}
                            _ = v253
                        } else {
                            var v254 uint16 = 255
                            for fi255 := 0; fi255 < 2; fi255++ {
                                v256 := GenS29{F0: "foo bar", F1: "hello", F2: v3}
                                if false {
                                    fmt.Println(v6, fi255)
                                    var v257 uint64 = 481
                                    _ = v257
                                } else {
                                    var v258 uint32 = v4
                                    _ = v258
                                }
                                _ = fi255
                                _ = v256
                            }
                            v1 = append(v1, (v4 << ((18446744073709551615) & 63)))
                            _ = v254
                        }
                    }
                    v6 = (!v5)
                    var v259 float64 = float64(((0.1 + (3.14159 - (float32(v223) * 1.5))) / 1.5))
                    _ = v259
                } else {
                    switch -9223372036854775808 {
                    case 32767:
                        v260 := GenS30{F0: 65535, F1: -1, F2: 32767}
                        _ = v260
                    }
                    if false {
                        if len(v1) > 0 {
                            v1[0] = v4
                        }
                    } else {
                        v2[32768] = -9223372036854775808
                        v261 := GenS31{F0: -801, F1: 255}
                        _ = v261
                    }
                }
                _ = v224
            } else {
                v5 = (v3 > (v3 >> ((42) & 63)))
                v262 := GenS32{F0: 259, F1: 18446744073709551615}
                _ = v262
            }
            _ = v222
            _ = v223
        }
    } else {
        var v263 bool = (((42 >= 571) && v6) && false)
        if false {
            fmt.Println(v263, v3)
            _ = len(v2)
        }
        switch 127 {
        case 128:
            var v264 int8 = -128
            _ = v264
        case 2147483647:
            for fi265 := 0; fi265 < 18; fi265++ {
                v3 += v3
                _ = fi265
            }
        default:
            if (266 != 0) {
                v266 := GenS33{F0: 42}
                _ = v266
            }
        }
        _ = v263
    }
    v267 := GenS34{F0: 42, F1: -66, F2: 42}
    fmt.Println(v5, v3, v6)
    v3 = v3
    _ = len(v1)
    if (false || (!v6)) {
        v268 := GenS35{F0: 9223372036854775807}
        _ = v268
    } else {
        v3 = -2147483648
        if v6 {
            v6 = (1e10 != -2.25)
            fmt.Println(v5, v4)
            var v269 int32 = (v3 * v3)
            _ = v269
        } else {
            var v270 float64 = (float64(0.1) * -2.25)
            _ = v270
        }
    }
    for fi271 := 0; fi271 < 5; fi271++ {
        _ = len(v2)
        _ = fi271
    }
    var v272 bool = (v6 && false)
    v273 := GenS36{F0: 100, F1: "\n\t"}
    var v274 uint = 42
    var v275 uint32 = (^(v4))
    v6 = (v4 >= genFn1((v274 % ((0) | 1)), 123, 919))
    if (v274 != v274) {
        _ = len(v2)
    } else {
        var v276 uint8 = 255
        for fi277 := 0; fi277 < 17; fi277++ {
            v278 := GenS4{F0: "foo bar"}
            _ = fi277
            _ = v278
        }
        _ = v276
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v272:", v272)
    fmt.Println("v274:", v274)
    fmt.Println("v275:", v275)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v267
    _ = v272
    _ = v273
    _ = v274
    _ = v275
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
