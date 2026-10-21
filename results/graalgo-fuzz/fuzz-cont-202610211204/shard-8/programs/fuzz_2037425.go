package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 bool; F2 uint }
type GenS4 struct { F0 bool; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint8; F1 int; F2 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 int16, p1 uint8, p2 int16) int8 {
    if false {
        v7 := GenS1{F0: 42, F1: 0}
        var v8 uint = 2
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p2)
    case "":
        var v9 float32 = 3.14159
        _ = v9
    case "a":
        var v10 int = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 bool = false
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 2147483647, F1: -654}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -1
}
func genFn1(p0 uint64, p1 int8) bool {
    v14 := GenS1{F0: -2147483648, F1: 1}
    if ((true && (true && (9223372036854775807 < 9223372036854775807))) && true) {
        v15 := GenS1{F0: -1, F1: 8}
        if (((-1 == -9223372036854775808) && (false || false)) || true) {
            v16 := GenS1{F0: -420, F1: 32767}
            _ = v16
        }
        _ = v15
    } else {
        var v17 int = 42
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v14
    return ("hello" != "hello")
}

func main() {
    v1 := []int{4294967295, 246}
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    v4 := []uint32{4294967295, 447}
    v5 := make(map[uint32]bool)
    v6 := []int{9223372036854775807, 4294967295}
    if (!(!false)) {
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(v2, fi18)
            v2 -= (1.5 * -2.25)
            fmt.Println(v2, fi18)
            _ = fi18
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int8 = -1
            _ = fi20
            _ = v21
        }
        v2 -= v2
        _ = fi19
    }
    if (919 != 32767) {
        v22 := GenS1{F0: 0, F1: 128}
        _ = v22
    }
    _ = len(v3)
    for fi23 := 0; fi23 < 8; fi23++ {
        v2 = 0.1
        if (!false) {
            switch fi23 {
            case 1:
                v2 = v2
            case 128:
                v24 := GenS2{F0: 0, F1: 32768}
                _ = v24
            default:
                v2 = v2
            }
            if len(v6) > 0 {
                v6[0] = (^(fi23))
            }
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        switch "foo bar" {
        case "":
            for fi26 := 0; fi26 < 15; fi26++ {
                v5[2147483647] = true
                switch 127 {
                case 508:
                    v2 += float64(1e10)
                case 557:
                    var v27 uint64 = 401
                    _ = v27
                default:
                    _ = len(v5)
                }
                var v28 uint = 349
                _ = fi26
                _ = v28
            }
        case "hello":
            var v29 float64 = v2
            _ = v29
        case "a":
            if true {
                var v30 uint8 = 204
                fmt.Println(v2, v2, v30)
                _ = v30
            }
        }
        v31 := GenS1{F0: 42, F1: fi25}
        _ = fi25
        _ = v31
    }
    v32 := GenS1{F0: -908, F1: -9223372036854775808}
    if (!((4294967295 < 577) || true)) {
        if false {
            if len(v6) > 0 {
                v6[0] = 32767
            }
            var v33 uint8 = 2
            var v34 float64 = (v2 - v2)
            _ = v33
            _ = v34
        } else {
            switch 32767 {
            case 9223372036854775807:
                v35 := GenS2{F0: 9223372036854775807, F1: 808}
                _ = v35
            case 0:
                v36 := GenS3{F0: -9223372036854775808, F1: (0 <= -537), F2: 169}
                _ = v36
            }
        }
    }
    switch -248 {
    case -2147483648:
        if false {
            v2 = v2
            v2 = v2
        } else {
            v2 = v2
            for fi37 := 0; fi37 < 9; fi37++ {
                for fi38 := 0; fi38 < 10; fi38++ {
                    if false {
                        v2 -= (-2.25 * v2)
                    } else {
                        if (1 <= 118) {
                            var v39 bool = false
                            if v39 {
                                if (true || v39) {
                                    fmt.Println(v39, fi37, v2)
                                    var v40 bool = (911 <= 722)
                                    v41 := GenS4{F0: false, F1: 0, F2: 9223372036854775807}
                                    var v42 uint = 282
                                    var v43 uint = 493
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    switch 851 {
                                    case 127:
                                        v39 = v39
                                    case -9223372036854775808:
                                        var v44 uint32 = 65535
                                        _ = v44
                                    case 128:
                                        if false {
                                            if v39 {
                                                var v45 int8 = -18
                                                v6 = append(v6, (^(fi38)))
                                                _ = v45
                                            }
                                        }
                                    default:
                                        if v39 {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (!((v39 || false) || v39)) {
                                                    var v46 bool = v39
                                                    _ = v46
                                                } else {
                                                    var v47 int8 = -49
                                                    fmt.Println(v39, v39, v39)
                                                    if len(v1) > 0 {
                                                        v1[0] = -2147483648
                                                    }
                                                    var v48 int64 = 458
                                                    if true {
                                                        if ((141 != 42) && (!false)) {
                                                            switch "hello" {
                                                            case "a":
                                                                var v49 int = fi38
                                                                _ = v49
                                                            case "":
                                                                v50 := GenS5{F0: 0, F1: fi38, F2: -32768}
                                                                _ = v50
                                                            }
                                                        } else {
                                                            v47 = genFn0(-494, 1, -642)
                                                            fmt.Println(fi37, v47)
                                                            if v39 {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = len(v4)
                                                            }
                                                        }
                                                        if v39 {
                                                            switch "foo bar" {
                                                            case "a":
                                                                if true {
                                                                    var v51 string = "hello"
                                                                    _ = len(v4)
                                                                    v52 := v3[(fi38 & (fi37 ^ -1))]
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                            var v53 bool = (float64(3.14159) < v2)
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS6{F0: (("a" + (("hello" + "\n\t") + ("foo bar" + "hello"))) + ""), F1: (v47 / ((v47) | 1))}
                                                            var v55 uint = 581
                                                            var v56 uint32 = 491
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = fi38
                                                        }
                                                    } else {
                                                        var v57 uint8 = 127
                                                        switch fi38 {
                                                        case 1:
                                                            v2 = (v2 - v2)
                                                            v47 = v47
                                                            v39 = true
                                                            v58 := v5[925]
                                                            _ = v58
                                                        }
                                                        v39 = ("\n\t" >= "")
                                                        _ = v57
                                                    }
                                                    if (-728 != 42) {
                                                        v59 := GenS6{F0: "hello", F1: -128}
                                                        var v60 int32 = -1
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        var v61 int8 = (v47 >> ((723) & 63))
                                                        _ = v61
                                                    }
                                                    switch fi38 {
                                                    case 208:
                                                        v39 = ("hello" > (((("foo bar" + "foo bar") + ("" + "a")) + "hello") + ((("\n\t" + "hello") + ("hello" + "a")) + "hello")))
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                }
                                            case "\n\t":
                                                if len(v1) > 0 {
                                                    v1[0] = (fi38 * fi38)
                                                }
                                            default:
                                                var v62 float64 = -2.25
                                                _ = v62
                                            }
                                        }
                                    }
                                }
                                v2 -= (((1.5 - (0.1 * v2)) - v2) + v2)
                            }
                            _ = v39
                        } else {
                            if false {
                                var v63 uint = 255
                                var v64 uint64 = 444
                                _ = v63
                                _ = v64
                            } else {
                                fmt.Println(v2, v2)
                                v65 := GenS2{F0: 9223372036854775807, F1: (fi38 - fi38)}
                                _ = v65
                            }
                            var v66 uint32 = 110
                            _ = v66
                        }
                        v1 = append(v1, (fi38 + 42))
                    }
                    v2 += float64((0.0 + (((float32((v2 + 1.5)) + (1.5 * 0.1)) / 3.14159) + float32((1.5 / float64((3.14159 * (float32(v2) - (float32(v2) + (0.0 - 1.5))))))))))
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            if (730 < 18446744073709551615) {
                                var v67 uint = 4294967295
                                var v68 uint16 = 321
                                fmt.Println(fi37, v67, v67)
                                _ = v67
                                _ = v68
                            } else {
                                v69 := GenS4{F0: true, F1: 767, F2: 684}
                                if false {
                                    v2 -= 1e10
                                    v2 -= v2
                                } else {
                                    v2 += 1e10
                                    v70 := GenS7{F0: (fi37 / ((fi38) | 1)), F1: 0}
                                    _ = v70
                                }
                                var v71 int8 = 42
                                _ = v69
                                _ = v71
                            }
                        case "a":
                            v72 := GenS3{F0: (-(fi37)), F1: true, F2: 2}
                            _ = v72
                        }
                    }
                    _ = fi38
                }
                _ = fi37
            }
        }
    default:
        fmt.Println(v2)
    }
    for fi73 := 0; fi73 < 18; fi73++ {
        var v74 uint32 = 356
        fmt.Println(v2, v2, v2)
        _ = fi73
        _ = v74
    }
    for fi75 := 0; fi75 < 19; fi75++ {
        fmt.Println(fi75)
        _ = fi75
    }
    var v76 int8 = -27
    if ((0 == 65535) && (295 != 498)) {
        if (186 == -1) {
            v77 := v5[0]
            _ = v77
        } else {
            v76 += 1
            var v78 int32 = 219
            _ = v78
        }
    } else {
        switch ("foo bar" + ((("" + "a") + ("hello" + "hello")) + "a")) {
        case "\n\t":
            v79 := GenS8{F0: v76, F1: 1}
            _ = v79
        case "a":
            if true {
                for fi80 := 0; fi80 < 18; fi80++ {
                    if ("\n\t" < ((("\n\t" + "foo bar") + "") + "\n\t")) {
                        for fi81 := 0; fi81 < 15; fi81++ {
                            v76 += (v76 << ((42) & 63))
                            var v82 int = fi81
                            var v83 uint8 = 53
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                        if true {
                            v2 += v2
                            delete(v5, 2147483648)
                            v76 = genFn0(42, 127, -469)
                        } else {
                            switch "" {
                            case "\n\t":
                                if (4294967295 != 2) {
                                    if ((true || true) && (65535 != 408)) {
                                        for fi84 := 0; fi84 < 15; fi84++ {
                                            var v85 string = "hello"
                                            _ = fi84
                                            _ = v85
                                        }
                                        v86 := GenS1{F0: -2147483648, F1: (fi80 << ((748) & 63))}
                                        _ = v86
                                    } else {
                                        for fi87 := 0; fi87 < 8; fi87++ {
                                            v88 := v5[375]
                                            v76 = 127
                                            _ = fi87
                                            _ = v88
                                        }
                                    }
                                    fmt.Println(v76, v76)
                                }
                            case "":
                                var v89 int = 0
                                _ = v89
                            default:
                                if (!false) {
                                    if false {
                                        v90 := v5[2147483647]
                                        _ = v90
                                    } else {
                                        v91 := GenS9{F0: 42, F1: 65535}
                                        _ = v91
                                    }
                                    fmt.Println(v76, v76)
                                    var v92 int32 = -1
                                    _ = v92
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v2 += v2
                                    case "\n\t":
                                        if false {
                                            v93 := GenS10{F0: 0}
                                            _ = v93
                                        } else {
                                            delete(v5, 2147483647)
                                            var v94 uint32 = 42
                                            _ = v94
                                        }
                                    default:
                                        v95 := GenS10{F0: (fi80 | 32768)}
                                        _ = v95
                                    }
                                    var v96 int32 = 786
                                    _ = v96
                                }
                            }
                            if true {
                                v2 = v2
                                for fi97 := 0; fi97 < 17; fi97++ {
                                    v2 += float64(1.5)
                                    var v98 int32 = 42
                                    _ = fi97
                                    _ = v98
                                }
                                if true {
                                    var v99 uint8 = 246
                                    var v100 int16 = 1000
                                    _ = v100
                                    _ = v99
                                } else {
                                    v6 = append(v6, (fi80 / (((fi80 | 4294967295)) | 1)))
                                    fmt.Println(v76, v2, fi80)
                                }
                            }
                            var v101 int32 = -1
                            _ = v101
                        }
                    } else {
                        var v102 float64 = v2
                        for fi103 := 0; fi103 < 8; fi103++ {
                            var v104 float64 = (v2 / float64(3.14159))
                            v105 := v3[143]
                            v76 = (v76 ^ (-(v76)))
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        for fi106 := 0; fi106 < 11; fi106++ {
                            v107 := GenS8{F0: v76, F1: -1}
                            v76 += v76
                            _ = fi106
                            _ = v107
                        }
                        _ = v102
                    }
                    var v108 uint32 = 42
                    _ = fi80
                    _ = v108
                }
                for fi109 := 0; fi109 < 8; fi109++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v110 uint16 = 266
                    _ = fi109
                    _ = v110
                }
            } else {
                v2 = (-2.25 / float64(1e10))
                switch "" {
                case "":
                    for fi111 := 0; fi111 < 2; fi111++ {
                        if (false || (true && false)) {
                            for fi112 := 0; fi112 < 12; fi112++ {
                                fmt.Println(fi111, v76, fi112)
                                v113 := GenS8{F0: 42, F1: 42}
                                if (4294967295 <= 4294967295) {
                                    if (127 <= 2) {
                                        fmt.Println(fi112, v76)
                                        v76 = (v76 >> ((255) & 63))
                                        switch "hello" {
                                        case "\n\t":
                                            v4 = append(v4, 2)
                                        case "":
                                            v114 := GenS11{F0: true, F1: "a", F2: "a"}
                                            _ = v114
                                        }
                                    }
                                }
                                _ = fi112
                                _ = v113
                            }
                            v5[2147483648] = ((float32(v2) * 3.14159) >= (3.14159 / (1.5 + 0.1)))
                            v115 := GenS9{F0: 2147483647, F1: 42}
                            _ = v115
                        }
                        _ = fi111
                    }
                case "hello":
                    v76 += -114
                }
            }
        case "":
            v116 := v3[9223372036854775807]
            _ = v116
        default:
            v76 -= genFn0(-108, 0, -1)
        }
        v2 -= ((v2 / (v2 - 1.5)) + 1e10)
        v2 += v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
