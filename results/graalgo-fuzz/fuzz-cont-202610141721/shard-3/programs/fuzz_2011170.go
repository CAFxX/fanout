package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 int }
type GenS5 struct { F0 int32; F1 int16; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 int64; F2 uint }
type GenS8 struct { F0 int }
func genFn0() int16 {
    switch "a" {
    case "":
        if true {
            v7 := []int{127}
            var v8 uint32 = 4294967295
            v8 = v8
            _ = v7
            _ = v8
        }
    case "hello":
        var v9 float64 = 0.0
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: (("foo bar" + "hello") + ("\n\t" + ("\n\t" + ("foo bar" + (("\n\t" + "") + "foo bar"))))), F1: -128, F2: 2}
        _ = v10
    default:
        var v11 bool = false
        _ = v11
    }
    v12 := []int{}
    var v13 uint64 = 2
    v14 := GenS1{F0: "foo bar", F1: 42, F2: 127}
    _ = v12
    _ = v13
    _ = v14
    return 0
}
func genFn1(p0 uint8, p1 int64, p2 int32) int64 {
    v15 := []int{}
    v15 = append(v15, 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = (!(18446744073709551615 != 2))
    var v3 string = ("\n\t" + "foo bar")
    var v4 uint16 = 0
    v5 := make(map[int]uint32)
    var v6 int = -980
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int = int(v1)
        v18 := GenS1{F0: "\n\t", F1: -128, F2: 0}
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = v3
    v4 += v4
    fmt.Println(v4, v2)
    v1 += (v1 / (((v1 ^ v1)) | 1))
    if false {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int16 = -32768
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            v4 = (v4 >> ((0) & 63))
            _ = fi21
        }
        v22 := []int{(v6 + v6), v6, v6, (v6 | int(v1))}
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int = (^(v6))
        if (42 != 244) {
            if v2 {
                var v25 string = ((v3 + (v3 + v3)) + v3)
                fmt.Println(v25, fi23)
                _ = v25
            }
            delete(v5, (v24 - (v6 | v6)))
            var v26 uint32 = uint32(v4)
            _ = v26
        } else {
            delete(v5, (v24 | fi23))
            var v27 bool = false
            if (-497 >= -1) {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v29 := GenS2{F0: "a", F1: 1}
                    v1 = (v1 << ((2) & 63))
                    _ = fi28
                    _ = v29
                }
                v30 := []int{32768, v6, 0, 4294967295}
                var v31 uint64 = 42
                _ = v30
                _ = v31
            }
            _ = v27
        }
        for fi32 := 0; fi32 < 17; fi32++ {
            switch (fi32 * 66) {
            case 128:
                if v2 {
                    switch (v24 | v6) {
                    case 42:
                        v33 := GenS2{F0: ((v3 + v3) + v3), F1: 2147483647}
                        _ = v33
                    }
                } else {
                    var v34 int = fi23
                    fmt.Println(v6)
                    var v35 int64 = v1
                    _ = v34
                    _ = v35
                }
            case 1:
                fmt.Println(fi23, v2)
            case -413:
                v5[(v24 * v6)] = 2
            default:
                fmt.Println(v3)
            }
            fmt.Println(v6, v24)
            if (v2 || (false && v2)) {
                var v36 string = "foo bar"
                _ = v36
            }
            _ = fi32
        }
        _ = fi23
        _ = v24
    }
    v6 = (-(v6))
    _ = len(v5)
    var v37 int8 = -1
    if (481 >= 503) {
        if false {
            if (((0.1 * (3.14159 * 0.0)) - (-2.25 + 1.5)) < (0.0 - 1e10)) {
                var v38 int = (-(v6))
                for fi39 := 0; fi39 < 18; fi39++ {
                    v6 = (v38 << ((933) & 63))
                    _ = fi39
                }
                switch v3 {
                case "foo bar":
                    fmt.Println(v3, v38, v38)
                case "":
                    switch v6 {
                    case -2147483648:
                        for fi40 := 0; fi40 < 5; fi40++ {
                            v41 := []int{-155, (fi40 - (v38 & 0)), v6, (v6 - 719)}
                            _ = fi40
                            _ = v41
                        }
                    }
                case "hello":
                    for fi42 := 0; fi42 < 9; fi42++ {
                        v2 = true
                        if (v2 || v2) {
                            if ((244 < 128) || (1e10 <= ((0.1 - 3.14159) + 1e10))) {
                                var v43 float64 = 0.1
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    var v45 int16 = 0
                                    _ = fi44
                                    _ = v45
                                }
                                _ = v43
                            }
                            fmt.Println(v37)
                            v46 := GenS2{F0: "a", F1: 2147483647}
                            _ = v46
                        }
                        _ = len(v5)
                        _ = fi42
                    }
                default:
                    if (v37 < v37) {
                        v47 := v5[(v38 + (v6 >> ((944) & 63)))]
                        var v48 bool = true
                        v6 += (v6 & (v6 >> ((2) & 63)))
                        _ = v47
                        _ = v48
                    }
                }
                _ = v38
            } else {
                var v49 int8 = (v37 | (v37 ^ 100))
                v50 := GenS3{F0: 65535, F1: (v1 * (v1 ^ genFn1(255, v1, 42)))}
                switch v6 {
                case 127:
                    switch v3 {
                    case "hello":
                        if v2 {
                            var v51 float32 = (1.5 / -2.25)
                            var v52 uint64 = 1
                            _ = v51
                            _ = v52
                        }
                    case "\n\t":
                        switch v6 {
                        case 127:
                            var v53 int16 = 32767
                            _ = v53
                        case -1:
                            if ("foo bar" < "\n\t") {
                                v54 := GenS3{F0: 237, F1: v1}
                                var v55 string = (("hello" + v3) + (v3 + v3))
                                _ = v54
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 1; fi56++ {
                                    v1 = (^(v1))
                                    _ = fi56
                                }
                                if v2 {
                                    v49 = v49
                                    switch ("hello" + ((v3 + v3) + "\n\t")) {
                                    case "foo bar":
                                        fmt.Println(v37, v1)
                                    case "hello":
                                        v57 := GenS4{F0: 0, F1: 2, F2: v6}
                                        _ = v57
                                    case "":
                                        fmt.Println(v37)
                                    default:
                                        var v58 float32 = -2.25
                                        _ = v58
                                    }
                                } else {
                                    if (!v2) {
                                        delete(v5, 853)
                                        switch v3 {
                                        case "hello":
                                            v2 = v2
                                        case "foo bar":
                                            if (!v2) {
                                                switch (v6 & (v6 ^ 4294967295)) {
                                                case 516:
                                                    v59 := []int{}
                                                    _ = v59
                                                case -1:
                                                    switch v6 {
                                                    case 937:
                                                        var v60 int32 = -283
                                                        _ = v60
                                                    }
                                                case 127:
                                                    for fi61 := 0; fi61 < 5; fi61++ {
                                                        v62 := []int{v6, (^(v6)), (v6 * v6)}
                                                        v63 := GenS5{F0: 1, F1: 0, F2: 985}
                                                        for fi64 := 0; fi64 < 13; fi64++ {
                                                            if false {
                                                                var v65 bool = true
                                                                var v66 uint8 = 92
                                                                if false {
                                                                    var v67 uint64 = 0
                                                                    var v68 int8 = 1
                                                                    delete(v5, (fi64 << ((276) & 63)))
                                                                    if ((42 <= 42) && (!(!false))) {
                                                                        switch fi64 {
                                                                        case 289:
                                                                            var v69 int32 = -2147483648
                                                                            _ = v69
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            switch 42 {
                                                                            case -1:
                                                                                if (v65 || (int8(v1) <= 1)) {
                                                                                    var v70 int8 = -128
                                                                                    _ = v70
                                                                                }
                                                                                fmt.Println(v1, v1)
                                                                                v71 := GenS1{F0: "", F1: 127, F2: v66}
                                                                                _ = v71
                                                                            case -9223372036854775808:
                                                                                if ((v65 && ((0.1 + 3.14159) == 1.5)) || true) {
                                                                                    v72 := GenS6{F0: 252}
                                                                                    switch v3 {
                                                                                    case "a":
                                                                                        v73 := v5[-967]
                                                                                        _ = v73
                                                                                    case "\n\t":
                                                                                        v49 = 100
                                                                                    default:
                                                                                        var v74 uint32 = 211
                                                                                        _ = v74
                                                                                    }
                                                                                    _ = v72
                                                                                } else {
                                                                                    v62 = append(v62, 985)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    fmt.Println(v65, v3, v4)
                                                                    if v65 {
                                                                        if v2 {
                                                                            v75 := GenS2{F0: "foo bar", F1: uint32(v4)}
                                                                            switch (((v3 + v3) + (v3 + (v3 + v3))) + v3) {
                                                                            case "a":
                                                                                v76 := GenS6{F0: 63}
                                                                                _ = v76
                                                                            case "\n\t":
                                                                                v49 = v37
                                                                            case "hello":
                                                                                var v77 float32 = 0.1
                                                                                _ = v77
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                        if false {
                                                                            if (v3 != "foo bar") {
                                                                                v78 := GenS4{F0: 1, F1: v66, F2: v6}
                                                                                var v79 uint = 255
                                                                                v80 := GenS7{F0: v1, F1: (v1 + int64(fi64)), F2: v79}
                                                                                _ = v78
                                                                                _ = v79
                                                                                _ = v80
                                                                            } else {
                                                                                var v81 uint8 = (v66 >> ((0) & 63))
                                                                                _ = v81
                                                                            }
                                                                            v6 -= fi61
                                                                        } else {
                                                                            fmt.Println(v66)
                                                                            v65 = v2
                                                                            var v82 bool = false
                                                                            var v83 bool = true
                                                                            var v84 float64 = 1.5
                                                                            _ = v82
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                        if ((v2 && false) || (!false)) {
                                                                            if false {
                                                                                var v85 string = ((v3 + v3) + ("" + ((v3 + (v3 + v3)) + v3)))
                                                                                if v2 {
                                                                                    v2 = v2
                                                                                }
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                switch v3 {
                                                                case "a":
                                                                    v6 = int(v4)
                                                                }
                                                                var v86 string = ("hello" + v3)
                                                                if v2 {
                                                                    var v87 uint = 255
                                                                    v88 := GenS8{F0: (fi64 / (((fi61 - (fi61 + int(v49)))) | 1))}
                                                                    v89 := GenS3{F0: 116, F1: (v1 | v1)}
                                                                    _ = v87
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    var v90 float64 = ((0.0 + (1e10 / (0.1 + 0.0))) * (-2.25 * 1e10))
                                                                    _ = v90
                                                                }
                                                                _ = v86
                                                            }
                                                            if (472 > 0) {
                                                                var v91 int16 = genFn0()
                                                                v91 = v91
                                                                _ = v91
                                                            } else {
                                                                if v2 {
                                                                    v3 = (v3 + v3)
                                                                }
                                                            }
                                                            delete(v5, int(v4))
                                                            _ = fi64
                                                        }
                                                        _ = fi61
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                v2 = v2
                                                var v92 int64 = -195
                                                _ = v92
                                            } else {
                                                fmt.Println(v3, v4, v49)
                                            }
                                        case "\n\t":
                                            switch (v6 / ((32768) | 1)) {
                                            case -1:
                                                v93 := GenS3{F0: 350, F1: (v1 + v1)}
                                                _ = v93
                                            case -9223372036854775808:
                                                v1 += (v1 & 0)
                                            }
                                        }
                                        v49 = (v49 + v37)
                                    } else {
                                        var v94 int = (v6 / ((v6) | 1))
                                        _ = v94
                                    }
                                    var v95 uint = 255
                                    _ = v95
                                }
                                fmt.Println(v2, v3)
                            }
                        }
                    case "":
                        for fi96 := 0; fi96 < 14; fi96++ {
                            v97 := []int{fi96, v6, -465, (fi96 << ((840) & 63))}
                            fmt.Println(v3, v1, v49)
                            _ = fi96
                            _ = v97
                        }
                    default:
                        v98 := GenS2{F0: v3, F1: 4294967295}
                        _ = v98
                    }
                case 9223372036854775807:
                    switch (v6 << ((255) & 63)) {
                    case 32767:
                        v37 += 127
                    default:
                        if (!v2) {
                            fmt.Println(v6, v2)
                            v1 = v1
                        }
                    }
                case -869:
                    var v99 uint64 = 4294967295
                    _ = v99
                default:
                    for fi100 := 0; fi100 < 5; fi100++ {
                        var v101 int16 = int16(v6)
                        _ = fi100
                        _ = v101
                    }
                }
                _ = v49
                _ = v50
            }
            for fi102 := 0; fi102 < 9; fi102++ {
                var v103 uint32 = 399
                switch ((v3 + (v3 + "a")) + "hello") {
                case "hello":
                    var v104 bool = v2
                    _ = v104
                case "a":
                    switch v3 {
                    case "a":
                        var v105 uint16 = (v4 * v4)
                        _ = v105
                    default:
                        v103 += (v103 & v103)
                    }
                default:
                    var v106 int64 = (v1 / (((v1 / (((v1 << ((18446744073709551615) & 63))) | 1))) | 1))
                    _ = v106
                }
                _ = fi102
                _ = v103
            }
            var v107 int8 = 100
            _ = v107
        }
        v108 := []int{-883, v6, (v6 - (v6 * (^(v6))))}
        delete(v5, v6)
        _ = v108
    } else {
        var v109 int = 537
        if (!v2) {
            if (!(v4 != v4)) {
                var v110 int64 = (v1 & (v1 * (v1 - -182)))
                v111 := GenS5{F0: -1, F1: genFn0(), F2: 4294967295}
                delete(v5, v109)
                _ = v110
                _ = v111
            } else {
                var v112 uint16 = 65535
                v112 -= v112
                switch 2147483647 {
                case -293:
                    v112 = 734
                }
                _ = v112
            }
        }
        _ = v109
    }
    v1 = genFn1(203, (^(v1)), 189)
    delete(v5, v6)
    for fi113 := 0; fi113 < 10; fi113++ {
        v3 = v3
        fmt.Println(v6, v2)
        if true {
            var v114 float32 = (0.1 + 0.0)
            _ = v114
        } else {
            v5[-1] = 0
        }
        _ = fi113
    }
    delete(v5, v6)
    var v115 int16 = 42
    v37 = (v37 / ((v37) | 1))
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v115
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
