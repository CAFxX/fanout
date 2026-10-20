package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 bool; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 string; F2 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int64; F2 string }
type GenS11 struct { F0 uint16; F1 bool }
func genFn0() uint {
    var v5 uint8 = 255
    if (0 <= 65535) {
        if true {
            if (-9223372036854775808 == -532) {
                var v6 uint8 = (^(v5))
                v6 = 145
                v7 := []int{-544}
                _ = v6
                _ = v7
            }
        } else {
            v5 = (v5 * 223)
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint64 = 4294967295
            v9 = 278
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 bool = false
            var v12 int32 = -544
            if (!v11) {
                v13 := GenS1{F0: (fi10 & 9223372036854775807), F1: 42, F2: 779}
                var v14 bool = ((2147483648 < 2147483648) && (9223372036854775807 < -843))
                _ = v13
                _ = v14
            } else {
                var v15 uint64 = 802
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v5)
    var v16 float32 = ((-2.25 * 3.14159) + float32(float64(0.1)))
    v17 := GenS1{F0: 317, F1: 2147483648, F2: 618}
    _ = v16
    _ = v17
    _ = v5
    return 4294967295
}
func genFn1(p0 int8) uint16 {
    var v18 uint32 = 2147483648
    for fi19 := 0; fi19 < 20; fi19++ {
        p0 -= -128
        fmt.Println(v18, p0, fi19)
        var v20 float32 = 1.5
        _ = fi19
        _ = v20
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            v21 := GenS1{F0: 265, F1: (v18 - (^(v18))), F2: -32768}
            v18 += (v18 - 130)
            _ = v21
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                v18 = (v18 % (((v18 - (v18 % ((42) | 1)))) | 1))
                if ("a" == "foo bar") {
                    v18 += 0
                }
                switch "hello" {
                case "hello":
                    p0 = 0
                default:
                    v23 := []int{(^(fi22)), (fi22 ^ fi22), fi22}
                    _ = v23
                }
                _ = fi22
            }
            v24 := []int{4294967295}
            _ = v24
        }
    case "hello":
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := GenS1{F0: (fi25 >> ((genFn0()) & 63)), F1: v18, F2: -32768}
            v27 := []int{128, -853}
            if false {
                if (952 == 892) {
                    switch -433 {
                    case 310:
                        for fi28 := 0; fi28 < 6; fi28++ {
                            var v29 uint64 = 2
                            switch (fi25 - fi28) {
                            case 964:
                                if false {
                                    fmt.Println(v18)
                                    var v30 uint16 = 833
                                    _ = v30
                                } else {
                                    var v31 int16 = 32767
                                    v27 = append(v27, -363)
                                    v29 += 995
                                    _ = v31
                                }
                            default:
                                var v32 uint = 1
                                _ = v32
                            }
                            _ = fi28
                            _ = v29
                        }
                    case -1:
                        if false {
                            v18 += (v18 - (v18 | (v18 - 65535)))
                            v33 := GenS1{F0: fi25, F1: 42, F2: 315}
                            if false {
                                v18 = 692
                                for fi34 := 0; fi34 < 5; fi34++ {
                                    v18 = 4294967295
                                    var v35 int = fi25
                                    if true {
                                        v18 = v18
                                        v36 := GenS2{F0: 42, F1: (v18 | (v18 ^ v18)), F2: 127}
                                        _ = v36
                                    } else {
                                        if len(v27) > 0 {
                                            v27[0] = -2147483648
                                        }
                                        v37 := GenS3{F0: 233, F1: false, F2: true}
                                        fmt.Println(v18)
                                        _ = v37
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                            } else {
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    v39 := GenS3{F0: 1, F1: (!true), F2: (!true)}
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            _ = v33
                        } else {
                            if (42 == 210) {
                                switch "\n\t" {
                                case "a":
                                    switch "foo bar" {
                                    case "a":
                                        var v40 float32 = ((0.0 - -2.25) + 3.14159)
                                        _ = v40
                                    case "\n\t":
                                        v18 = v18
                                    case "foo bar":
                                        v41 := GenS1{F0: fi25, F1: (v18 - v18), F2: 1}
                                        _ = v41
                                    }
                                case "hello":
                                    fmt.Println(fi25, fi25, v18)
                                case "\n\t":
                                    fmt.Println(v18, p0, fi25)
                                }
                            }
                        }
                    case 2147483647:
                        var v42 int32 = -1
                        _ = v42
                    default:
                        switch (("hello" + (("foo bar" + (("" + ("\n\t" + ("foo bar" + ("" + "hello")))) + "")) + "")) + ("a" + "foo bar")) {
                        case "":
                            var v43 int8 = 42
                            _ = v43
                        case "\n\t":
                            v18 = v18
                        case "a":
                            _ = len(v27)
                        default:
                            var v44 float64 = 1.5
                            _ = v44
                        }
                    }
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                }
            } else {
                var v45 float32 = 3.14159
                var v46 int = fi25
                switch fi25 {
                case 923:
                    v47 := GenS4{F0: 0}
                    _ = v47
                default:
                    var v48 int = 143
                    _ = v48
                }
                _ = v45
                _ = v46
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    case "\n\t":
        v49 := GenS5{F0: 128, F1: 32767}
        _ = v49
    default:
        v50 := []int{-1, 32768}
        _ = v50
    }
    switch 1000 {
    case 9223372036854775807:
        if false {
            v18 = 2
            fmt.Println(p0)
            fmt.Println(p0)
        } else {
            v18 = (v18 & (v18 - v18))
            p0 = p0
            if (uint16(v18) != 201) {
                v51 := []int{int(v18)}
                _ = v51
            } else {
                var v52 uint64 = 2
                v53 := GenS3{F0: 128, F1: true, F2: false}
                v54 := []int{127, 128, 42, 128}
                _ = v52
                _ = v53
                _ = v54
            }
        }
    case 127:
        var v55 int8 = (p0 | 100)
        _ = v55
    case 922:
        v18 += v18
    }
    _ = p0
    _ = v18
    return 46
}

func main() {
    var v1 int16 = -1
    var v2 int32 = 2147483647
    var v3 string = ""
    v4 := make(map[string]string)
    switch v3 {
    case "hello":
        if (42 > -341) {
            if (true || (4294967295 == 42)) {
                v1 = v1
                fmt.Println(v3, v3, v1)
            } else {
                if (0.0 >= (0.1 / 0.1)) {
                    v56 := GenS2{F0: 4294967295, F1: 831, F2: 7}
                    v57 := GenS1{F0: -1, F1: 1000, F2: -32768}
                    v4[""] = v3
                    _ = v56
                    _ = v57
                }
            }
            switch v3 {
            case "":
                var v58 uint64 = 64
                _ = v58
            default:
                var v59 float64 = 0.1
                _ = v59
            }
        }
    }
    for fi60 := 0; fi60 < 12; fi60++ {
        var v61 uint32 = 719
        v3 = v3
        fmt.Println(v61, v61, v1)
        _ = fi60
        _ = v61
    }
    v62 := []int{0}
    if len(v62) >= 2 {
        v62 = v62[1:len(v62)]
    }
    if false {
        fmt.Println(v3)
    } else {
        v63 := v4[v3]
        var v64 string = (("hello" + (v3 + v3)) + "hello")
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 15; fi65++ {
        for fi66 := 0; fi66 < 1; fi66++ {
            if (32767 <= 1) {
                var v67 int8 = 0
                var v68 uint16 = 1
                if true {
                    v3 = v3
                    _ = len(v4)
                }
                var v69 uint8 = 0
                v69 -= v69
                _ = v67
                _ = v68
                _ = v69
            } else {
                switch (fi65 / ((fi66) | 1)) {
                case -9223372036854775808:
                    var v70 string = "a"
                    _ = v70
                case 4294967295:
                    v2 -= 208
                }
                if ((221 <= 109) || true) {
                    v2 -= v2
                }
                switch 32768 {
                case 0:
                    var v71 uint64 = 2
                    _ = v71
                default:
                    _ = len(v4)
                }
            }
            v72 := GenS6{F0: 255, F1: 403}
            if false {
                var v73 int32 = (v2 << ((0) & 63))
                if len(v62) > 0 {
                    v62[0] = fi66
                }
                var v74 bool = ((!false) || (!false))
                _ = v73
                _ = v74
            }
            _ = fi66
            _ = v72
        }
        for fi75 := 0; fi75 < 6; fi75++ {
            if (((v3 + "a") + "a") <= (v3 + "a")) {
                if ((65535 >= 421) || (v1 == 964)) {
                    if ((-(v2)) > (v2 % (((v2 + (v2 << ((genFn0()) & 63)))) | 1))) {
                        var v76 bool = (((42 >= 2) && ((0.1 != -2.25) || ((fi75 * -2147483648) < (fi75 - 32768)))) || true)
                        if ((!true) && (fi75 <= (-(fi65)))) {
                            switch v3 {
                            case "a":
                                if (65535 >= 255) {
                                    v1 = -484
                                    v3 = v3
                                }
                            default:
                                v2 = (v2 << ((18446744073709551615) & 63))
                            }
                            v2 = v2
                            v3 = (v3 + "")
                        } else {
                            var v77 int32 = v2
                            v2 -= -63
                            v4["foo bar"] = "a"
                            var v78 int8 = -128
                            _ = v77
                            _ = v78
                        }
                        _ = v76
                    }
                    var v79 uint8 = 210
                    _ = v79
                } else {
                    if false {
                        fmt.Println(fi75, v3, fi75)
                        switch (fi75 << ((1) & 63)) {
                        case -590:
                            switch (fi75 + (-(fi65))) {
                            case 32768:
                                if (3 > (^(v1))) {
                                    if false {
                                        var v80 int = -820
                                        v80 += (v80 | (fi75 & (v80 | v80)))
                                        _ = v80
                                    }
                                }
                            case 127:
                                var v81 float64 = 3.14159
                                _ = v81
                            case -2147483648:
                                if (!(true && true)) {
                                    v2 -= int32(v1)
                                    v82 := GenS7{F0: 4294967295}
                                    _ = v82
                                } else {
                                    var v83 string = (v3 + ((v3 + ("a" + "hello")) + ((("hello" + v3) + "hello") + v3)))
                                    _ = v83
                                }
                            default:
                                fmt.Println(fi75, v3, v3)
                                switch v3 {
                                case "a":
                                    var v84 int = (fi75 ^ (fi75 % (((fi75 / (((fi65 >> ((0) & 63))) | 1))) | 1)))
                                    _ = v84
                                default:
                                    var v85 int32 = (v2 - 0)
                                    _ = v85
                                }
                                var v86 bool = true
                                _ = v86
                            }
                        case 991:
                            v1 = v1
                        case 128:
                            _ = len(v62)
                        }
                    } else {
                        v87 := GenS8{F0: 255, F1: v3, F2: 192}
                        switch v3 {
                        case "hello":
                            switch fi65 {
                            case 32767:
                                var v88 string = "foo bar"
                                _ = v88
                            case -2147483648:
                                var v89 uint = 65535
                                _ = v89
                            case 9223372036854775807:
                                if ((!true) && true) {
                                    if false {
                                        if (float64(0.1) <= ((0.1 * (((0.1 - 3.14159) / (((float64(0.0) * (0.0 - -2.25)) - -2.25) + 1.5)) - 0.0)) * (3.14159 - 0.0))) {
                                            var v90 uint64 = 1
                                            _ = v90
                                        } else {
                                            v91 := GenS1{F0: (fi65 << ((65535) & 63)), F1: 4294967295, F2: (v1 << ((65535) & 63))}
                                            if false {
                                                fmt.Println(fi65, v1)
                                            } else {
                                                if (!(true && false)) {
                                                    fmt.Println(fi75)
                                                } else {
                                                    _ = len(v4)
                                                }
                                            }
                                            v3 = ((v3 + v3) + "")
                                            _ = v91
                                        }
                                    } else {
                                        if ((fi65 << ((genFn0()) & 63)) != (fi75 << ((255) & 63))) {
                                            var v92 uint8 = 128
                                            v93 := GenS1{F0: 518, F1: 0, F2: (v1 % (((v1 << ((4294967295) & 63))) | 1))}
                                            if len(v62) > 0 {
                                                v62[0] = (fi65 / ((-807) | 1))
                                            }
                                            _ = v92
                                            _ = v93
                                        } else {
                                            v94 := GenS9{F0: v3}
                                            _ = v94
                                        }
                                        v95 := GenS10{F0: ((!(false && false)) || false), F1: 9223372036854775807, F2: v3}
                                        var v96 uint = 1
                                        _ = v95
                                        _ = v96
                                    }
                                } else {
                                    delete(v4, v3)
                                    v97 := GenS5{F0: (fi75 + -233), F1: 0}
                                    v98 := GenS3{F0: 2, F1: ((3.14159 - 0.1) >= 1.5), F2: false}
                                    _ = v97
                                    _ = v98
                                }
                            }
                        case "\n\t":
                            v3 = v3
                        }
                        v3 = v3
                        _ = v87
                    }
                    if false {
                        v1 = v1
                        if len(v62) > 0 {
                            v62[0] = (fi75 * fi65)
                        }
                    } else {
                        var v99 uint16 = 864
                        if ((false || (42 != 918)) || (fi65 != fi75)) {
                            v100 := GenS5{F0: fi65, F1: (v99 | v99)}
                            _ = v100
                        }
                        var v101 int64 = 9223372036854775807
                        _ = v101
                        _ = v99
                    }
                    var v102 float32 = 1.5
                    fmt.Println(v2, v1)
                    _ = v102
                }
                v62 = append(v62, (^(fi65)))
            }
            _ = fi75
        }
        _ = fi65
    }
    for fi103 := 0; fi103 < 4; fi103++ {
        var v104 int32 = (v2 << ((1) & 63))
        fmt.Println(v1)
        _ = fi103
        _ = v104
    }
    var v105 int16 = (v1 - (v1 ^ v1))
    fmt.Println(v3)
    for fi106 := 0; fi106 < 5; fi106++ {
        fmt.Println(v3, v2, v2)
        v107 := GenS2{F0: 2147483647, F1: 97, F2: 42}
        fmt.Println(v105, v105, fi106)
        _ = fi106
        _ = v107
    }
    delete(v4, (v3 + v3))
    v3 = "\n\t"
    v4[(v3 + "hello")] = v3
    if (-128 != -128) {
        v108 := GenS6{F0: 2, F1: 1}
        _ = v108
    } else {
        if false {
            v2 = v2
            v109 := GenS11{F0: 255, F1: true}
            var v110 int16 = 42
            _ = v109
            _ = v110
        }
    }
    for fi111 := 0; fi111 < 6; fi111++ {
        var v112 uint8 = 213
        var v113 int = fi111
        var v114 uint16 = 42
        _ = fi111
        _ = v112
        _ = v113
        _ = v114
    }
    v115 := GenS3{F0: 255, F1: true, F2: (4294967295 > 0)}
    var v116 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v105
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v62
}
