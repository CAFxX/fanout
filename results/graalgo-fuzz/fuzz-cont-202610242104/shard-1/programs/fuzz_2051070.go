package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int8; F2 int32 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 string; F1 uint32; F2 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 uint }
type GenS13 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 int = -938
    var v5 int = (v4 | 349)
    switch v4 {
    case 127:
        if ((true || true) && (true && false)) {
            fmt.Println(v5)
            v3 = v3
        }
    case -916:
        var v6 int8 = 0
        _ = v6
    case -522:
        if ((!false) && (42 < 2147483647)) {
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v3 -= (v3 + (1.5 / (v3 + float32(1e10))))
                    _ = fi8
                }
                _ = fi7
            }
        } else {
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(v3)
                _ = fi9
            }
            if false {
                fmt.Println(v5)
                if (-41 > -128) {
                    fmt.Println(v5)
                }
            } else {
                fmt.Println(v3, v4, v3)
            }
            fmt.Println(v5, v4)
        }
    }
    _ = len(v1)
    v10 := []int{32767}
    _ = len(v10)
    if len(v10) > 0 {
        v10[0] = (v4 / (((v5 % ((v4) | 1))) | 1))
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v5 = (v5 * fi11)
        if false {
            var v12 uint32 = 2147483647
            v5 = (v4 % (((v5 & v4)) | 1))
            var v13 uint32 = v12
            _ = v12
            _ = v13
        }
        if len(v10) > 0 {
            v10[0] = -1
        }
        _ = fi11
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v5 = (v5 + v4)
    var v14 uint32 = 4294967295
    v15 := GenS1{F0: 127, F1: v14, F2: 713}
    if (!true) {
        fmt.Println(v14, v4)
    } else {
        v16 := GenS1{F0: (v5 % (((v4 / ((298) | 1))) | 1)), F1: v14, F2: 18446744073709551615}
        v17 := GenS2{F0: -9223372036854775808, F1: ("a" >= "a"), F2: 4294967295}
        _ = v16
        _ = v17
    }
    v3 = ((1e10 * -2.25) - ((0.0 - v3) + (3.14159 * float32(-2.25))))
    var v18 int = (^(v4))
    if (!false) {
        v5 -= -915
    } else {
        v4 += 128
        fmt.Println(v5, v5)
        v18 = v4
    }
    v19 := GenS3{F0: 42}
    if (0 < 255) {
        v20 := GenS2{F0: -9223372036854775808, F1: ((18446744073709551615 <= 923) || ((v3 - v3) <= 1.5)), F2: 255}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS3{F0: 4294967295}
            v3 = -2.25
            for fi23 := 0; fi23 < 14; fi23++ {
                v14 -= (v14 << ((563) & 63))
                fmt.Println(v5)
                if false {
                    v24 := GenS3{F0: 347}
                    v4 += v18
                    if (false || (!(true || (v14 <= v14)))) {
                        v25 := GenS3{F0: 2}
                        if (v14 > v14) {
                            if false {
                                var v26 uint32 = 2147483647
                                v14 = 42
                                switch v4 {
                                case 128:
                                    var v27 uint16 = 32767
                                    _ = v27
                                default:
                                    _ = len(v10)
                                }
                                _ = v26
                            } else {
                                v28 := GenS3{F0: 1}
                                if len(v10) > 0 {
                                    v10[0] = 128
                                }
                                _ = v28
                            }
                        } else {
                            fmt.Println(fi21, v4, v3)
                            v5 -= (fi23 / (((v4 + v4)) | 1))
                            var v29 float64 = (0.0 * (1e10 / 3.14159))
                            _ = v29
                        }
                        _ = v25
                    } else {
                        if false {
                            var v30 bool = ((("foo bar" + "") >= "hello") && (1 >= -1))
                            _ = v30
                        }
                        fmt.Println(fi23, fi23, v3)
                    }
                    var v31 float32 = ((((float32(-2.25) + 0.0) + 0.1) / 1.5) * (1e10 + v3))
                    if ((0 != -421) || true) {
                        var v32 string = "foo bar"
                        _ = v32
                    }
                    _ = v24
                    _ = v31
                } else {
                    if false {
                        v4 -= v18
                        var v33 uint16 = 0
                        switch "foo bar" {
                        case "":
                            fmt.Println(v33, v4)
                        case "foo bar":
                            var v34 int8 = 42
                            _ = v34
                        }
                        _ = v33
                    } else {
                        switch -187 {
                        case 1:
                            switch (v18 * 32767) {
                            case -2147483648:
                                v35 := v1[(-(v5))]
                                _ = v35
                            case 0:
                                if (0 != 1) {
                                    var v36 uint16 = 2
                                    _ = len(v2)
                                    var v37 int64 = 1
                                    _ = v36
                                    _ = v37
                                }
                                if (255 != 255) {
                                    v38 := GenS2{F0: -1, F1: ((-37 != 1) || (true || false)), F2: 42}
                                    _ = len(v10)
                                    v39 := GenS2{F0: -9223372036854775808, F1: true, F2: 255}
                                    _ = v38
                                    _ = v39
                                } else {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    v40 := GenS3{F0: 42}
                                    v18 = v5
                                    _ = v40
                                }
                                if false {
                                    v5 += (v4 / ((0) | 1))
                                }
                                v41 := GenS3{F0: 18446744073709551615}
                                fmt.Println(fi23, v18)
                                _ = v41
                            }
                        case 32768:
                            if false {
                                if true {
                                    delete(v1, 0)
                                } else {
                                    var v42 bool = false
                                    if v42 {
                                        if false {
                                            var v43 uint = 255
                                            _ = v43
                                        } else {
                                            v44 := GenS2{F0: 42, F1: v42, F2: 2}
                                            switch v18 {
                                            case 127:
                                                var v45 float32 = v3
                                                _ = v45
                                            case 992:
                                                var v46 uint32 = v14
                                                _ = v46
                                            }
                                            v14 += v14
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            _ = v44
                                        }
                                    }
                                    fmt.Println(v3)
                                    _ = v42
                                }
                                if false {
                                    if true {
                                        fmt.Println(v3, v3)
                                    }
                                }
                                v47 := GenS2{F0: -1, F1: (!true), F2: 85}
                                _ = v47
                            }
                        }
                        var v48 string = "hello"
                        _ = v48
                    }
                    var v49 uint16 = 310
                    if false {
                        if false {
                            v50 := v1[v4]
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v50
                        } else {
                            if false {
                                fmt.Println(fi21, fi21)
                                if (false && true) {
                                    if false {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        if (-46 >= 42) {
                                            fmt.Println(v49, v14, v4)
                                            if (!true) {
                                                var v51 uint8 = uint8(v14)
                                                v52 := v2[v4]
                                                _ = v51
                                                _ = v52
                                            }
                                            v3 = 1.5
                                        } else {
                                            switch v18 {
                                            case 194:
                                                var v53 uint8 = 3
                                                _ = v53
                                            case 9223372036854775807:
                                                fmt.Println(fi23, v18)
                                            default:
                                                v3 = (v3 / v3)
                                            }
                                        }
                                        var v54 int = fi23
                                        var v55 uint16 = v49
                                        if false {
                                            if false {
                                                v49 += v55
                                            }
                                            v14 = v14
                                            switch "" {
                                            case "foo bar":
                                                if len(v10) > 0 {
                                                    v10[0] = v18
                                                }
                                            }
                                            if true {
                                                switch fi23 {
                                                case -9223372036854775808:
                                                    v10 = append(v10, -963)
                                                    v56 := GenS4{F0: -32768, F1: -1}
                                                    _ = v56
                                                case 32768:
                                                    v10 = append(v10, (v5 << ((0) & 63)))
                                                case -928:
                                                    if true {
                                                        var v57 bool = true
                                                        var v58 int16 = 96
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                default:
                                                    v14 += (^(v14))
                                                }
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    switch v4 {
                                                    case -924:
                                                        if false {
                                                            if false {
                                                                switch "hello" {
                                                                case "a":
                                                                    var v59 float32 = v3
                                                                    _ = v59
                                                                }
                                                                v55 += v49
                                                            } else {
                                                                v54 = int(v49)
                                                                v14 = (v14 << ((2) & 63))
                                                            }
                                                            fmt.Println(v14, fi23)
                                                            v60 := GenS2{F0: -9223372036854775808, F1: (!true), F2: 2}
                                                            _ = v60
                                                        } else {
                                                            fmt.Println(v3, v5, v4)
                                                            v3 -= (v3 / ((v3 / v3) - (v3 / 1.0)))
                                                        }
                                                    case 1:
                                                        var v61 string = "hello"
                                                        var v62 int8 = -128
                                                        if (42 != 306) {
                                                            var v63 uint32 = (^(v14))
                                                            var v64 uint64 = 536
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            var v65 bool = false
                                                            _ = v65
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                    default:
                                                        v66 := GenS5{F0: ("foo bar" + "\n\t")}
                                                        _ = v66
                                                    }
                                                case "hello":
                                                    v49 += (v49 << ((0) & 63))
                                                default:
                                                    if true {
                                                        switch "hello" {
                                                        case "":
                                                            v14 = v14
                                                        case "a":
                                                            var v67 float32 = (1.5 + v3)
                                                            _ = v67
                                                        case "hello":
                                                            fmt.Println(v14, v49, v5)
                                                        default:
                                                            if (((0.0 == 1e10) || (((^(v55)) > (v55 & 1)) || false)) && (!((42 != 42) || (float32(1.5) > 1e10)))) {
                                                                var v68 float32 = 3.14159
                                                                if false {
                                                                    v69 := GenS6{F0: 386, F1: 9223372036854775807}
                                                                    if true {
                                                                        fmt.Println(v14)
                                                                        v70 := GenS2{F0: int64(v14), F1: (-1 != -9223372036854775808), F2: 909}
                                                                        v10 = append(v10, (v18 | (-(fi23))))
                                                                        var v71 int64 = -9223372036854775808
                                                                        v18 -= (fi23 << ((18446744073709551615) & 63))
                                                                        _ = len(v2)
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        var v72 uint = 255
                                                                        if false {
                                                                            switch v18 {
                                                                            case -9223372036854775808:
                                                                                if (2 != 4294967295) {
                                                                                    fmt.Println(fi23, v14)
                                                                                    if ((1.5 - 0.0) == 3.14159) {
                                                                                        v73 := GenS6{F0: (^(v72)), F1: fi21}
                                                                                        _ = v73
                                                                                    } else {
                                                                                        if (!(uint8(v14) != 0)) {
                                                                                            switch v5 {
                                                                                            case 0:
                                                                                                var v74 uint64 = 916
                                                                                                _ = v74
                                                                                            default:
                                                                                                fmt.Println(v72, fi23, fi21)
                                                                                            }
                                                                                            v75 := GenS7{F0: v5}
                                                                                            _ = v75
                                                                                        }
                                                                                        v76 := GenS1{F0: (fi23 | v18), F1: v14, F2: 4294967295}
                                                                                        v72 = v72
                                                                                        _ = v76
                                                                                    }
                                                                                    switch (("a" + "\n\t") + "hello") {
                                                                                    case "hello":
                                                                                        _ = len(v1)
                                                                                    case "a":
                                                                                        if false {
                                                                                            switch fi21 {
                                                                                            case -2147483648:
                                                                                                if (!false) {
                                                                                                    switch fi21 {
                                                                                                    case 42:
                                                                                                        delete(v1, (v18 / ((v18) | 1)))
                                                                                                    case 2147483647:
                                                                                                        if (false || (277 >= 42)) {
                                                                                                            switch "hello" {
                                                                                                            case "\n\t":
                                                                                                                _ = len(v2)
                                                                                                            default:
                                                                                                                fmt.Println(v18)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v77 := v2[v18]
                                                                                                            _ = v77
                                                                                                        }
                                                                                                    case 127:
                                                                                                        fmt.Println(v14)
                                                                                                    }
                                                                                                    if false {
                                                                                                        v2[(fi23 - (v54 >> ((uint(v49)) & 63)))] = -487
                                                                                                        fmt.Println(fi23, v18)
                                                                                                    } else {
                                                                                                        fmt.Println(v49, v3, v68)
                                                                                                    }
                                                                                                    var v78 int8 = 42
                                                                                                    _ = v78
                                                                                                }
                                                                                            case 4294967295:
                                                                                                var v79 bool = true
                                                                                                _ = v79
                                                                                            }
                                                                                            var v80 uint16 = 888
                                                                                            _ = v80
                                                                                        } else {
                                                                                            delete(v2, (v18 + (v4 / (((fi23 - -2147483648)) | 1))))
                                                                                            v49 = (^(v55))
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v3 = 0.0
                                                                                    default:
                                                                                        v2[(v5 & 921)] = (v54 + v5)
                                                                                    }
                                                                                } else {
                                                                                    v18 += v4
                                                                                    v4 += (-(fi21))
                                                                                    switch (((((("" + ("a" + "a")) + "hello") + "\n\t") + (("\n\t" + ("" + ("a" + "\n\t"))) + "foo bar")) + ("foo bar" + "hello")) + ("foo bar" + "")) {
                                                                                    case "a":
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v81 := GenS8{F0: -32768, F1: 0, F2: -933}
                                                                                            switch (fi21 * v54) {
                                                                                            case -564:
                                                                                                delete(v1, (v54 * v4))
                                                                                            case -178:
                                                                                                v54 += 32768
                                                                                            default:
                                                                                                v10 = append(v10, 0)
                                                                                            }
                                                                                            _ = v81
                                                                                        case "hello":
                                                                                            var v82 int16 = 844
                                                                                            _ = v82
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v5 += v4
                                                                                    case "hello":
                                                                                        switch fi23 {
                                                                                        case 2147483647:
                                                                                            v4 += (fi23 / ((366) | 1))
                                                                                        case 128:
                                                                                            var v83 int64 = 42
                                                                                            v84 := GenS9{F0: 0, F1: (v14 ^ (v14 | 2147483648))}
                                                                                            v10 = append(v10, fi21)
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case -963:
                                                                                v85 := GenS7{F0: (fi21 >> (((v72 / (((v72 >> ((2) & 63))) | 1))) & 63))}
                                                                                _ = v85
                                                                            }
                                                                            var v86 uint64 = 2
                                                                            switch (v18 - fi23) {
                                                                            case -406:
                                                                                v87 := GenS7{F0: -2147483648}
                                                                                _ = v87
                                                                            }
                                                                            _ = v86
                                                                        }
                                                                        _ = len(v2)
                                                                        _ = v72
                                                                    }
                                                                    _ = v69
                                                                }
                                                                _ = v68
                                                            }
                                                        }
                                                        var v88 uint32 = v14
                                                        _ = v88
                                                    }
                                                }
                                                v4 = 2147483647
                                                var v89 uint = 1
                                                _ = v89
                                            }
                                        } else {
                                            v54 = (v4 >> ((661) & 63))
                                            delete(v1, v54)
                                        }
                                        var v90 uint16 = v49
                                        var v91 string = "foo bar"
                                        _ = v54
                                        _ = v55
                                        _ = v90
                                        _ = v91
                                    }
                                    fmt.Println(v49, v5)
                                }
                                fmt.Println(v3, v5, v4)
                            } else {
                                v3 = (((v3 * 1.5) + ((0.0 + 1.5) - 3.14159)) - v3)
                                v18 = 4294967295
                                v18 = (v5 & 127)
                            }
                            v92 := GenS10{F0: (("hello" + (("foo bar" + "") + ((((("\n\t" + "") + "a") + "\n\t") + ("" + "a")) + ""))) + "\n\t"), F1: v14, F2: ("\n\t" + "foo bar")}
                            _ = v92
                        }
                        delete(v1, 313)
                        if false {
                            v4 = (v4 ^ (v5 * v5))
                            if (1 >= 106) {
                                v4 -= fi21
                                if false {
                                    var v93 int = v4
                                    v4 = (v93 ^ (-(v4)))
                                    if true {
                                        v94 := GenS10{F0: ("foo bar" + "hello"), F1: v14, F2: "hello"}
                                        if ((^(v49)) > (v49 / (((v49 | v49)) | 1))) {
                                            _ = len(v1)
                                            var v95 uint = 1
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            _ = v95
                                        }
                                        fmt.Println(fi23)
                                        _ = v94
                                    }
                                    if false {
                                        var v96 int32 = 94
                                        if false {
                                            if (!true) {
                                                v5 = (v18 >> ((18446744073709551615) & 63))
                                                v5 = (^(fi23))
                                                var v97 float32 = -2.25
                                                _ = v97
                                            }
                                            v98 := GenS4{F0: 0, F1: 9223372036854775807}
                                            _ = v98
                                        } else {
                                            v18 = (v93 & v93)
                                        }
                                        fmt.Println(fi21)
                                        _ = v96
                                    }
                                    v99 := GenS11{F0: 1}
                                    var v100 uint32 = (v14 * (v14 + v14))
                                    switch v93 {
                                    case -305:
                                        switch "" {
                                        case "a":
                                            if (true && true) {
                                                v18 = (v93 + v4)
                                            } else {
                                                var v101 int8 = 1
                                                _ = v101
                                            }
                                        case "":
                                            v5 += (v18 >> ((42) & 63))
                                        }
                                    case -101:
                                        v5 += (v93 * -554)
                                    case -53:
                                        var v102 string = "foo bar"
                                        _ = v102
                                    default:
                                        var v103 int8 = -90
                                        if (612 >= 2) {
                                            v104 := GenS12{F0: v14, F1: 1}
                                            if (!(1 > -2147483648)) {
                                                fmt.Println(v100)
                                                switch "foo bar" {
                                                case "a":
                                                    var v105 float64 = 3.14159
                                                    _ = v105
                                                case "":
                                                    var v106 uint8 = 127
                                                    _ = v106
                                                default:
                                                    if (0 >= (v14 / ((v100) | 1))) {
                                                        var v107 uint64 = 2
                                                        _ = v107
                                                    } else {
                                                        var v108 uint8 = 99
                                                        _ = v108
                                                    }
                                                }
                                            } else {
                                                v109 := GenS5{F0: (("\n\t" + "\n\t") + "\n\t")}
                                                _ = v109
                                            }
                                            _ = v104
                                        }
                                        _ = v103
                                    }
                                    _ = v100
                                    _ = v93
                                    _ = v99
                                } else {
                                    var v110 float64 = (((float64((1.5 + (v3 * 0.0))) - 1e10) / -2.25) + (0.1 + ((1e10 * 1e10) - 1.5)))
                                    _ = len(v1)
                                    _ = v110
                                }
                                fmt.Println(v18)
                            } else {
                                switch "a" {
                                case "":
                                    var v111 int64 = -1
                                    _ = v111
                                default:
                                    v5 += fi23
                                }
                                v5 = (^(v4))
                                v5 = -346
                            }
                        } else {
                            var v112 uint = 2
                            v4 = fi23
                            _ = v112
                        }
                    }
                    v113 := GenS10{F0: "foo bar", F1: (v14 << ((2) & 63)), F2: ("hello" + "a")}
                    if false {
                        v114 := GenS11{F0: v49}
                        _ = v114
                    }
                    var v115 int32 = 2147483647
                    _ = v113
                    _ = v115
                    _ = v49
                }
                switch (("" + "foo bar") + "") {
                case "":
                    var v116 string = "foo bar"
                    var v117 bool = (false && (false && true))
                    _ = v116
                    _ = v117
                case "hello":
                    v10 = append(v10, (^(v4)))
                case "a":
                    v118 := GenS13{F0: 111}
                    _ = v118
                }
                v119 := GenS3{F0: 2}
                switch "hello" {
                case "\n\t":
                    var v120 float64 = float64(float32((float64(float32(-2.25)) - (0.0 * 1e10))))
                    _ = v120
                default:
                    v121 := GenS3{F0: 42}
                    _ = v121
                }
                var v122 string = ("a" + "foo bar")
                v10 = append(v10, 127)
                _ = fi23
                _ = v119
                _ = v122
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
