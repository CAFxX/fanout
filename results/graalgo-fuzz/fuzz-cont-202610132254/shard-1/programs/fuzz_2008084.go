package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int; F1 bool; F2 string }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 bool; F1 uint16; F2 int16 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS12 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0(p0 int32, p1 int, p2 uint32) uint8 {
    p2 = 2147483647
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = 0
        v9 := GenS1{F0: 128, F1: "a"}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.1
    if (0 < 1) {
        var v11 uint16 = 32767
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 uint32, p1 uint16) int32 {
    v12 := []int{-229, 1, int(p1)}
    p1 = p1
    p0 = p0
    fmt.Println(p1, p0)
    p1 = 255
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    var v1 int32 = -341
    var v2 string = "a"
    var v3 int16 = -1
    v4 := []uint32{1, 2147483648, 1}
    v5 := make(map[int]bool)
    var v6 uint64 = 0
    var v13 uint32 = 65535
    v14 := GenS1{F0: 209, F1: (v2 + v2)}
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 = "hello"
        if (v6 >= uint64(v13)) {
            for fi16 := 0; fi16 < 15; fi16++ {
                switch (fi15 << ((42) & 63)) {
                case 770:
                    v3 = v3
                case 127:
                    v17 := v5[(^(fi16))]
                    _ = v17
                }
                _ = fi16
            }
            v18 := GenS1{F0: genFn0(genFn1((v13 % ((v13) | 1)), 255), fi15, 2147483648), F1: v2}
            if true {
                var v19 string = ""
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 float64 = (float64(((3.14159 / 1.0) / 1.0)) / 1e10)
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS2{F0: v2, F1: genFn1(v13, 888)}
                    _ = fi22
                    _ = v23
                }
                _ = v19
            }
            _ = v18
        } else {
            v24 := GenS3{F0: v1, F1: false, F2: false}
            for fi25 := 0; fi25 < 3; fi25++ {
                if false {
                    var v26 uint32 = v13
                    var v27 uint32 = v26
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(fi25)
                    fmt.Println(v2, v2)
                    v6 = v6
                }
                if true {
                    v2 = "foo bar"
                    if ((42 < -124) || true) {
                        v28 := GenS1{F0: 139, F1: "foo bar"}
                        v13 = 0
                        _ = v28
                    }
                    switch (v2 + v2) {
                    case "":
                        v5[-2147483648] = (!(!(2 >= 255)))
                    case "foo bar":
                        if (v3 > v3) {
                            v1 = -2147483648
                            var v29 uint16 = 65535
                            _ = v29
                        } else {
                            if (65535 <= 0) {
                                switch v2 {
                                case "a":
                                    switch fi25 {
                                    case 4294967295:
                                        v30 := GenS1{F0: 116, F1: v2}
                                        _ = v30
                                    case -9223372036854775808:
                                        v3 = (v3 & v3)
                                    default:
                                        v1 -= v1
                                    }
                                }
                            }
                            var v31 uint64 = v6
                            v32 := GenS4{F0: v31}
                            fmt.Println(v1, v13)
                            if true {
                                v2 = "a"
                                var v33 string = v2
                                fmt.Println(fi25)
                                _ = v33
                            }
                            var v34 uint16 = 1
                            _ = v31
                            _ = v32
                            _ = v34
                        }
                    default:
                        var v35 bool = true
                        _ = v35
                    }
                } else {
                    if true {
                        if true {
                            var v36 string = (("foo bar" + ("" + "")) + "foo bar")
                            if true {
                                if true {
                                    v13 = 2147483648
                                    var v37 int16 = v3
                                    v2 = v2
                                    var v38 int32 = v1
                                    if (585 < 32767) {
                                        v6 = (v6 ^ v6)
                                        fmt.Println(v36, v36, fi15)
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = (v13 / ((42) | 1))
                                        }
                                    }
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                var v39 string = v36
                                _ = v39
                            }
                            if false {
                                if false {
                                    v40 := GenS2{F0: (v2 + (v36 + ("hello" + v36))), F1: v1}
                                    _ = v40
                                } else {
                                    v2 = (v2 + v36)
                                }
                            } else {
                                v41 := GenS4{F0: v6}
                                _ = v41
                            }
                            switch "hello" {
                            case "a":
                                var v42 bool = true
                                _ = v42
                            case "\n\t":
                                v13 += v13
                            case "foo bar":
                                switch "" {
                                case "foo bar":
                                    v2 = v36
                                default:
                                    v43 := GenS5{F0: true, F1: -58}
                                    _ = v43
                                }
                            default:
                                var v44 string = v2
                                _ = v44
                            }
                            _ = v36
                        }
                        if (47 == 213) {
                            switch 9223372036854775807 {
                            case 709:
                                v1 += v1
                            case 32768:
                                if ((^(v6)) != 42) {
                                    switch fi25 {
                                    case 2147483647:
                                        v6 = v6
                                    case -661:
                                        var v45 uint64 = (v6 & (^(v6)))
                                        _ = v45
                                    default:
                                        switch -1 {
                                        case 219:
                                            v3 += v3
                                            switch (^(fi25)) {
                                            case 850:
                                                var v46 int8 = -1
                                                _ = v46
                                            case 9223372036854775807:
                                                var v47 uint32 = v13
                                                _ = v47
                                            case 1:
                                                switch v2 {
                                                case "\n\t":
                                                    if ((!false) || (v13 > v13)) {
                                                        switch 42 {
                                                        case 127:
                                                            v6 = (^(v6))
                                                        default:
                                                            if ((1e10 + 0.0) < (((-2.25 + float32(0.0)) - -2.25) * 0.1)) {
                                                                if true {
                                                                    if true {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        var v48 int = 954
                                                                        var v49 int16 = v3
                                                                        switch 1 {
                                                                        case 32767:
                                                                            delete(v5, (^(v48)))
                                                                        case -2147483648:
                                                                            v50 := v5[1]
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            _ = len(v5)
                                                                            _ = v50
                                                                        default:
                                                                            v51 := GenS4{F0: 108}
                                                                            _ = v51
                                                                        }
                                                                        v5[(fi25 / ((4294967295) | 1))] = true
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        v5[2147483647] = false
                                                                        if (true && false) {
                                                                            var v52 uint8 = 128
                                                                            var v53 int32 = (v1 >> ((65535) & 63))
                                                                            _ = len(v4)
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                    var v54 uint8 = 128
                                                                    var v55 uint = 2
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                            } else {
                                                                var v56 string = ("\n\t" + v2)
                                                                if (false || true) {
                                                                    var v57 uint16 = 65535
                                                                    var v58 uint32 = v13
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                fmt.Println(v2, v6, v1)
                                                                _ = v56
                                                            }
                                                        }
                                                        v13 = (v13 << ((719) & 63))
                                                        var v59 bool = (v3 != v3)
                                                        var v60 int64 = 9223372036854775807
                                                        var v61 bool = false
                                                        var v62 bool = (-128 <= -128)
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        if (((0.1 - ((1.5 / 1.0) / 0.1)) * 1e10) > 0.0) {
                                                            v63 := GenS6{F0: (^(fi15)), F1: (0 >= 180)}
                                                            _ = v63
                                                        } else {
                                                            var v64 uint64 = v6
                                                            _ = v64
                                                        }
                                                    }
                                                default:
                                                    if (v13 >= v13) {
                                                        var v65 int = -1
                                                        var v66 int8 = 42
                                                        v6 -= (v6 | (v6 << ((2) & 63)))
                                                        var v67 uint32 = v13
                                                        if (0.1 >= 1.5) {
                                                            var v68 int32 = v1
                                                            _ = v68
                                                        } else {
                                                            fmt.Println(v66)
                                                            var v69 uint = 42
                                                            _ = v69
                                                        }
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        v70 := GenS7{F0: fi25, F1: false, F2: v2}
                                                        if true {
                                                            delete(v5, 32768)
                                                            var v71 uint = 1
                                                            _ = v71
                                                        }
                                                        if false {
                                                            if (299 >= 405) {
                                                                v1 += (v1 | genFn1(42, 32767))
                                                                var v72 float32 = 1e10
                                                                if len(v4) > 0 {
                                                                    v4[0] = v13
                                                                }
                                                                _ = v72
                                                            } else {
                                                                v3 = (v3 >> ((1) & 63))
                                                                fmt.Println(v6, v6)
                                                            }
                                                            var v73 int8 = -127
                                                            var v74 int64 = 1
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        _ = v70
                                                    }
                                                }
                                            }
                                        case 127:
                                            switch (fi25 << ((768) & 63)) {
                                            case -116:
                                                if (v1 != 2147483647) {
                                                    delete(v5, 42)
                                                } else {
                                                    if true {
                                                        switch fi25 {
                                                        case 128:
                                                            v75 := GenS5{F0: (9223372036854775807 >= -9223372036854775808), F1: 0}
                                                            _ = v75
                                                        case 190:
                                                            var v76 int16 = int16(v13)
                                                            _ = v76
                                                        case 0:
                                                            v4 = append(v4, 359)
                                                            fmt.Println(v6, v3, fi15)
                                                        }
                                                        if true {
                                                            v1 -= (v1 & 1)
                                                            v3 += (v3 % ((1000) | 1))
                                                        } else {
                                                            var v77 bool = (-9223372036854775808 >= -52)
                                                            if v77 {
                                                                var v78 uint64 = v6
                                                                _ = v78
                                                            } else {
                                                                var v79 uint = 888
                                                                _ = v79
                                                            }
                                                            _ = v77
                                                        }
                                                        switch (fi25 & -1) {
                                                        case 4294967295:
                                                            switch fi15 {
                                                            case -902:
                                                                switch "a" {
                                                                case "a":
                                                                    if true {
                                                                        v6 += v6
                                                                        switch (("" + ("a" + ("\n\t" + v2))) + "") {
                                                                        case "a":
                                                                            if ((v3 - v3) > 32767) {
                                                                                _ = len(v4)
                                                                            } else {
                                                                                if true {
                                                                                    v80 := GenS8{F0: 2, F1: false}
                                                                                    v3 += v3
                                                                                    switch v2 {
                                                                                    case "hello":
                                                                                        v1 -= genFn1(23, 32767)
                                                                                    case "foo bar":
                                                                                        var v81 int = (fi15 + (fi15 >> ((4294967295) & 63)))
                                                                                        _ = v81
                                                                                    case "a":
                                                                                        v1 = -1
                                                                                    }
                                                                                    _ = v80
                                                                                } else {
                                                                                    v3 += int16(v6)
                                                                                    var v82 int = (fi25 % ((fi25) | 1))
                                                                                    var v83 uint = 412
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                }
                                                                                v6 = v6
                                                                            }
                                                                        case "\n\t":
                                                                            v1 = -1
                                                                        case "":
                                                                            v2 = v2
                                                                        default:
                                                                            v84 := GenS2{F0: (((v2 + "a") + v2) + "\n\t"), F1: v1}
                                                                            _ = v84
                                                                        }
                                                                        v2 = "foo bar"
                                                                    } else {
                                                                        fmt.Println(v6, v13)
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            fmt.Println(v1, v3)
                                                                        default:
                                                                            if (false && ((v3 << ((4294967295) & 63)) <= v3)) {
                                                                                _ = len(v4)
                                                                                v85 := GenS1{F0: 181, F1: "hello"}
                                                                                fmt.Println(v13)
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        v86 := GenS9{F0: 156, F1: (v3 | -784)}
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        _ = v86
                                                                    }
                                                                default:
                                                                    switch (v2 + v2) {
                                                                    case "foo bar":
                                                                        v87 := GenS10{F0: true, F1: 2, F2: v3}
                                                                        var v88 int16 = (v3 << ((650) & 63))
                                                                        _ = v87
                                                                        _ = v88
                                                                    case "a":
                                                                        if (v3 != (v3 << ((107) & 63))) {
                                                                            v89 := GenS11{F0: true, F1: 1, F2: 42}
                                                                            var v90 string = ("" + v2)
                                                                            _ = len(v5)
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (v13 | (v13 - v13))
                                                                            }
                                                                            v5[fi15] = true
                                                                            _ = v89
                                                                            _ = v90
                                                                        } else {
                                                                            v91 := GenS5{F0: (v3 == (v3 + v3)), F1: -1}
                                                                            if true {
                                                                                v3 -= (-(v3))
                                                                                var v92 uint32 = (v13 << ((1) & 63))
                                                                                _ = v92
                                                                            } else {
                                                                                if true {
                                                                                    if true {
                                                                                        v2 = (v2 + v2)
                                                                                        v6 -= (v6 + 555)
                                                                                        v93 := GenS10{F0: (false || true), F1: 640, F2: (-(v3))}
                                                                                        _ = v93
                                                                                    }
                                                                                    if (!((v2 >= "") && false)) {
                                                                                        fmt.Println(v1, fi15, v3)
                                                                                        v6 = v6
                                                                                        v94 := GenS7{F0: -273, F1: true, F2: v2}
                                                                                        _ = v94
                                                                                    } else {
                                                                                        v95 := GenS2{F0: v2, F1: (v1 & (^(v1)))}
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2, v13, v6)
                                                                                v96 := GenS1{F0: 255, F1: "foo bar"}
                                                                                _ = v96
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                    case "":
                                                                        if (!true) {
                                                                            var v97 int = -9223372036854775808
                                                                            _ = v97
                                                                        } else {
                                                                            fmt.Println(v3)
                                                                            v13 = v13
                                                                            fmt.Println(v2, v3, fi15)
                                                                        }
                                                                    default:
                                                                        v1 = 42
                                                                    }
                                                                }
                                                            case -847:
                                                                if false {
                                                                    v3 = (v3 / ((v3) | 1))
                                                                    v13 = v13
                                                                    if (((v2 + (v2 + v2)) + "") != v2) {
                                                                        v98 := GenS5{F0: true, F1: 42}
                                                                        _ = v98
                                                                    }
                                                                } else {
                                                                    var v99 uint8 = 42
                                                                    if true {
                                                                        v2 = (v2 + v2)
                                                                        fmt.Println(fi15)
                                                                    }
                                                                    v1 += genFn1((v13 & v13), 0)
                                                                    _ = v99
                                                                }
                                                            default:
                                                                if true {
                                                                    fmt.Println(v6)
                                                                    v13 = (v13 | v13)
                                                                    var v100 uint16 = 244
                                                                    v13 = (v13 & 729)
                                                                    var v101 uint32 = 2147483647
                                                                    _ = v100
                                                                    _ = v101
                                                                } else {
                                                                    v13 = uint32(v6)
                                                                    if true {
                                                                        v1 -= v1
                                                                        v13 -= v13
                                                                    } else {
                                                                        v2 = ""
                                                                        v102 := GenS12{F0: v2, F1: (v6 >> ((42) & 63)), F2: v6}
                                                                        fmt.Println(v6, v2, v6)
                                                                        _ = v102
                                                                    }
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            var v103 bool = true
                                                            _ = v103
                                                        case -9223372036854775808:
                                                            if ((1.5 + 0.0) != 0.0) {
                                                                var v104 uint64 = v6
                                                                _ = v104
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                fmt.Println(v6, fi15)
                                            }
                                        default:
                                            var v105 uint16 = 785
                                            _ = v105
                                        }
                                    }
                                    v13 = (v13 & 0)
                                    fmt.Println(v6, v2, v1)
                                    v106 := GenS7{F0: (fi25 ^ fi15), F1: true, F2: (v2 + "hello")}
                                    _ = v106
                                } else {
                                    var v107 uint8 = uint8(v6)
                                    _ = v107
                                }
                                fmt.Println(fi25, v3, fi15)
                                v13 = v13
                            }
                        } else {
                            v1 -= v1
                            fmt.Println(fi25, fi15)
                            switch (v2 + "foo bar") {
                            case "foo bar":
                                v6 += v6
                            default:
                                var v108 uint64 = v6
                                _ = v108
                            }
                        }
                        var v109 uint8 = 128
                        v5[32768] = false
                        v6 -= (v6 & (v6 - v6))
                        _ = v109
                    }
                    fmt.Println(v13)
                }
                _ = fi25
            }
            _ = v24
        }
        if (v1 <= v1) {
            v1 -= v1
        }
        _ = fi15
    }
    switch 127 {
    case -162:
        if false {
            var v110 uint8 = 42
            _ = v110
        }
    default:
        v6 = (v6 / ((18446744073709551615) | 1))
    }
    v13 -= v13
    delete(v5, 13)
    fmt.Println(v13, v2, v2)
    if ((v1 << ((42) & 63)) < (v1 >> ((1) & 63))) {
        v3 -= (v3 ^ v3)
        var v111 uint = 538
        _ = v111
    } else {
        var v112 uint = 0
        v13 -= 0
        _ = v112
    }
    var v113 float32 = 0.1
    var v114 int32 = v1
    v115 := v5[-355]
    var v116 uint32 = (v13 % (((v13 & v13)) | 1))
    var v117 float64 = (0.0 / 1e10)
    var v118 int8 = 73
    v119 := GenS6{F0: -128, F1: v115}
    var v120 int8 = (-(v118))
    v116 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v120:", v120)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v116
    _ = v117
    _ = v118
    _ = v119
    _ = v120
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
