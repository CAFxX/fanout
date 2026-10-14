package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 int32; F2 bool }
type GenS7 struct { F0 bool; F1 int64; F2 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 uint16; F2 bool }
type GenS12 struct { F0 int64; F1 int64 }
type GenS13 struct { F0 uint32 }
func genFn0(p0 int) int64 {
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int32 = -2147483648
            _ = fi7
            _ = v8
        }
        if false {
            p0 = 4294967295
            v9 := []int{(-(fi6)), (fi6 + p0)}
            var v10 int8 = -48
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 float64 = 0.0
                if true {
                    v13 := []int{(fi6 + fi6), 2147483647}
                    if len(v13) > 0 {
                        v13[0] = fi11
                    }
                    _ = v13
                }
                v14 := []int{2147483647}
                _ = fi11
                _ = v12
                _ = v14
            }
        }
        v5 = 42
        _ = fi6
    }
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 int32) int64 {
    if true {
        p0 = 2147483647
    }
    switch 0 {
    case 2147483647:
        p0 += (p0 >> ((2) & 63))
    default:
        switch (("a" + "") + (("hello" + "a") + "\n\t")) {
        case "foo bar":
            var v15 int64 = 42
            _ = v15
        }
    }
    _ = p0
    return 1
}

func main() {
    var v1 float64 = (0.1 + (1.5 + -2.25))
    var v2 bool = ((42 <= 42) || true)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v16 uint32 = 2147483647
    var v17 int64 = -837
    fmt.Println(v2, v17)
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v2, v2)
        v19 := GenS1{F0: (v16 | v16)}
        _ = len(v3)
        _ = fi18
        _ = v19
    }
    v20 := []int{-26, -2147483648, 4294967295}
    var v21 uint16 = 2
    var v22 uint32 = 1000
    if v2 {
        switch 995 {
        case 128:
            delete(v3, (v16 | (v16 + (v16 & 1))))
        case -976:
            switch 128 {
            case 32767:
                fmt.Println(v17, v21, v21)
            case -9223372036854775808:
                if (1000 == -32768) {
                    if (v2 || (v22 >= v22)) {
                        switch 2147483647 {
                        case 115:
                            v1 = (v1 * (v1 + ((3.14159 * ((1.5 - (v1 + (v1 - (v1 / v1)))) - v1)) * (v1 / v1))))
                        case 128:
                            if false {
                                var v23 float32 = 0.0
                                v23 += v23
                                _ = v23
                            }
                        case -418:
                            _ = len(v20)
                        default:
                            switch 2147483647 {
                            case 413:
                                v16 = (v22 / ((v22) | 1))
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    if v2 {
                                        var v24 int16 = 1
                                        _ = v24
                                    }
                                case "a":
                                    _ = len(v3)
                                case "hello":
                                    v17 += genFn0(-296)
                                default:
                                    var v25 int16 = 1000
                                    _ = v25
                                }
                            }
                        }
                        fmt.Println(v22, v2)
                        if (true || true) {
                            v1 += v1
                        } else {
                            v1 = float64(3.14159)
                        }
                    } else {
                        var v26 bool = (-31 == 127)
                        _ = v26
                    }
                    var v27 int8 = 59
                    v4[42] = "\n\t"
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v29 := GenS1{F0: (v22 * (v22 - (^(v22))))}
                        if v2 {
                            for fi30 := 0; fi30 < 9; fi30++ {
                                var v31 string = ("hello" + ("hello" + ""))
                                _ = fi30
                                _ = v31
                            }
                            v1 = v1
                            delete(v3, v22)
                        } else {
                            var v32 uint32 = v16
                            var v33 uint8 = 1
                            fmt.Println(v16)
                            _ = v32
                            _ = v33
                        }
                        _ = fi28
                        _ = v29
                    }
                    for fi34 := 0; fi34 < 3; fi34++ {
                        switch "hello" {
                        case "a":
                            v16 = v22
                        case "\n\t":
                            var v35 int8 = int8(v22)
                            _ = v35
                        case "":
                            v17 -= int64(v21)
                        default:
                            for fi36 := 0; fi36 < 14; fi36++ {
                                v37 := GenS1{F0: (v16 << ((65535) & 63))}
                                _ = fi36
                                _ = v37
                            }
                        }
                        _ = fi34
                    }
                }
            case 282:
                if (!v2) {
                    v17 += (v17 * v17)
                    v38 := GenS2{F0: 0, F1: -547, F2: (v17 | -1)}
                    var v39 string = "a"
                    _ = v38
                    _ = v39
                } else {
                    if (v21 == 103) {
                        if v2 {
                            v16 = 2147483647
                            for fi40 := 0; fi40 < 3; fi40++ {
                                v20 = append(v20, 127)
                                var v41 int8 = 75
                                if len(v20) > 0 {
                                    v20[0] = 32768
                                }
                                _ = fi40
                                _ = v41
                            }
                            v42 := GenS2{F0: 2, F1: 127, F2: v17}
                            _ = v42
                        } else {
                            _ = len(v20)
                        }
                        v43 := GenS3{F0: 1, F1: v22}
                        _ = v43
                    }
                    var v44 int32 = 2147483647
                    _ = v44
                }
            }
        default:
            if v2 {
                v21 = (v21 * v21)
                if false {
                    if ("" > "") {
                        if v2 {
                            v1 -= (v1 * (float64((-2.25 * (0.0 - ((1e10 - 0.1) / (1.5 + 1.5))))) - v1))
                        }
                        fmt.Println(v1, v16, v21)
                    }
                    v1 -= (v1 + v1)
                } else {
                    v22 = (v22 % (((v16 | v16)) | 1))
                    v3[v22] = v2
                    var v45 uint8 = 42
                    _ = v45
                }
            } else {
                v16 += (v16 - 42)
            }
        }
        v46 := GenS4{F0: 58, F1: 2, F2: "a"}
        _ = v46
    } else {
        var v47 uint8 = 1
        if v2 {
            var v48 int32 = 42
            if (int8(v47) < 8) {
                var v49 uint64 = 4294967295
                fmt.Println(v17)
                _ = v49
            }
            _ = v48
        } else {
            if ((-(v17)) != -9223372036854775808) {
                for fi50 := 0; fi50 < 16; fi50++ {
                    fmt.Println(v16, v21)
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v16, v2)
                    default:
                        if (false || v2) {
                            for fi51 := 0; fi51 < 12; fi51++ {
                                v20 = append(v20, (fi50 << ((1) & 63)))
                                if (!v2) {
                                    if true {
                                        var v52 string = ""
                                        var v53 bool = v2
                                        if v2 {
                                            v53 = (false || v2)
                                            if v2 {
                                                var v54 int8 = 122
                                                _ = v54
                                            }
                                            _ = len(v3)
                                            var v55 int32 = -2147483648
                                            _ = v55
                                        } else {
                                            v22 = (v22 + (v16 % ((v22) | 1)))
                                            _ = len(v3)
                                            fmt.Println(fi50)
                                        }
                                        _ = v52
                                        _ = v53
                                    }
                                    v22 = (v22 >> ((uint(v21)) & 63))
                                    switch fi51 {
                                    case 32767:
                                        if true {
                                            v56 := GenS3{F0: 1, F1: v16}
                                            var v57 int64 = genFn0(fi51)
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 float32 = 0.1
                                            var v59 uint32 = (v22 | 4294967295)
                                            _ = v58
                                            _ = v59
                                        }
                                    case 9223372036854775807:
                                        switch (fi51 & (fi51 - fi51)) {
                                        case 1:
                                            v60 := GenS5{F0: fi51}
                                            _ = v60
                                        case -343:
                                            switch fi50 {
                                            case -1:
                                                v17 -= -72
                                            case 967:
                                                v61 := GenS6{F0: false, F1: 42, F2: v2}
                                                _ = v61
                                            }
                                        }
                                    case 42:
                                        if v2 {
                                            var v62 uint64 = 227
                                            if len(v20) >= 2 {
                                                v20 = v20[1:len(v20)]
                                            }
                                            v47 -= (v47 * 42)
                                            var v63 int8 = 100
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v64 := GenS5{F0: -8}
                                            fmt.Println(v21, v22, fi50)
                                            if (!(!v2)) {
                                                v65 := GenS4{F0: v47, F1: 0, F2: ""}
                                                v2 = (-138 < 32767)
                                                _ = v65
                                            }
                                            _ = v64
                                        }
                                    }
                                    var v66 int16 = -32768
                                    var v67 uint = 255
                                    v68 := v3[v22]
                                    if (!v68) {
                                        v47 = uint8(v16)
                                        switch ("\n\t" + "\n\t") {
                                        case "hello":
                                            if v2 {
                                                var v69 int64 = -9223372036854775808
                                                v16 += v16
                                                v22 = (v22 ^ 0)
                                                _ = v69
                                            } else {
                                                var v70 uint16 = (v21 + (v21 - 2))
                                                v22 = uint32(v70)
                                                var v71 int64 = v17
                                                var v72 float32 = (-2.25 * -2.25)
                                                if v68 {
                                                    var v73 int16 = (v66 % (((v66 - v66)) | 1))
                                                    var v74 int32 = -631
                                                    if true {
                                                        fmt.Println(v73, v22)
                                                        if v68 {
                                                            v73 = (v73 ^ (-(v66)))
                                                        } else {
                                                            fmt.Println(v2, v72, v2)
                                                        }
                                                        var v75 string = "foo bar"
                                                        _ = v75
                                                    } else {
                                                        v22 = (v16 >> (((v67 / (((v67 >> ((v67) & 63))) | 1))) & 63))
                                                        v71 -= v17
                                                        if (v17 <= (v71 & v71)) {
                                                            var v76 int32 = v74
                                                            var v77 float32 = (v72 - v72)
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            v74 = v74
                                                        }
                                                    }
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                            var v78 int32 = 806
                                            if v68 {
                                                if (((fi50 - (fi50 | fi50)) < fi51) && v2) {
                                                    if ((v68 || v68) || true) {
                                                        fmt.Println(v67, v17)
                                                        if (true && ((!((v17 + 534) > v17)) || v68)) {
                                                            if v2 {
                                                                if (v68 || true) {
                                                                    switch (-(fi51)) {
                                                                    case 128:
                                                                        switch -317 {
                                                                        case -353:
                                                                            var v79 uint = 4294967295
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    _ = len(v20)
                                                                    switch (("hello" + "a") + "\n\t") {
                                                                    case "a":
                                                                        switch 0 {
                                                                        case 0:
                                                                            v22 = (v16 - 2)
                                                                        default:
                                                                            v78 += v78
                                                                        }
                                                                    default:
                                                                        if (602 >= 18446744073709551615) {
                                                                            v80 := GenS7{F0: ((v47 * (v47 / ((2) | 1))) == 200), F1: genFn0(fi50), F2: 127}
                                                                            fmt.Println(v17, fi50)
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                } else {
                                                                    v2 = (v66 != -1)
                                                                    fmt.Println(v47)
                                                                    v21 += uint16(v47)
                                                                    v2 = (2 > (v47 >> ((v67) & 63)))
                                                                }
                                                                if v2 {
                                                                    _ = len(v20)
                                                                } else {
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                }
                                                                fmt.Println(fi51)
                                                            } else {
                                                                _ = len(v4)
                                                                fmt.Println(v2, v1, v22)
                                                                if len(v20) > 0 {
                                                                    v20[0] = (fi51 / ((int(v21)) | 1))
                                                                }
                                                            }
                                                            v47 = v47
                                                            if false {
                                                                var v81 uint16 = v21
                                                                v67 += v67
                                                                v82 := GenS8{F0: (("" + "") <= ("foo bar" + ("a" + "a")))}
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                        } else {
                                                            if (3.14159 < 0.1) {
                                                                fmt.Println(v22, v66)
                                                                v22 -= 4294967295
                                                                v67 += v67
                                                                switch (fi51 / ((32768) | 1)) {
                                                                case -697:
                                                                    var v83 uint64 = 0
                                                                    _ = v83
                                                                case 4294967295:
                                                                    fmt.Println(fi51)
                                                                }
                                                            } else {
                                                                var v84 int32 = (v78 / ((v78) | 1))
                                                                v20 = append(v20, fi51)
                                                                _ = v84
                                                            }
                                                            v67 -= (v67 * (v67 ^ (v67 >> ((889) & 63))))
                                                        }
                                                        v2 = (v2 && (((v78 | (^(v78))) != v78) || (3.14159 < float32(((float64(float32(v1)) / 1e10) * v1)))))
                                                    }
                                                    v16 -= v16
                                                    if (v21 >= v21) {
                                                        v21 = (v21 | v21)
                                                        var v85 int = (fi50 ^ fi51)
                                                        _ = v85
                                                    } else {
                                                        v86 := GenS7{F0: v68, F1: v17, F2: 42}
                                                        v16 -= v22
                                                        _ = v86
                                                    }
                                                    var v87 bool = (!((18446744073709551615 <= 0) && v68))
                                                    _ = v87
                                                }
                                            }
                                            _ = v78
                                        default:
                                            switch ((("" + "hello") + ("" + ("hello" + "hello"))) + ("" + ("foo bar" + "a"))) {
                                            case "":
                                                var v88 int32 = -2147483648
                                                v89 := GenS9{F0: fi51}
                                                _ = v88
                                                _ = v89
                                            default:
                                                if len(v20) > 0 {
                                                    v20[0] = (fi50 - 128)
                                                }
                                            }
                                        }
                                        switch fi51 {
                                        case -9223372036854775808:
                                            var v90 string = "a"
                                            _ = v90
                                        case 32767:
                                            var v91 int16 = 42
                                            _ = v91
                                        default:
                                            if (127 != (^(v47))) {
                                                switch (fi50 - -492) {
                                                case 127:
                                                    if len(v20) > 0 {
                                                        v20[0] = 0
                                                    }
                                                }
                                                switch 128 {
                                                case -470:
                                                    if v2 {
                                                        v92 := GenS4{F0: (v47 / ((128) | 1)), F1: 648, F2: "foo bar"}
                                                        v16 -= (v22 << ((v67) & 63))
                                                        v93 := GenS7{F0: v68, F1: -159, F2: 42}
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            v22 += 523
                                                        case "foo bar":
                                                            var v94 float64 = v1
                                                            _ = v94
                                                        default:
                                                            if (!true) {
                                                                v68 = v68
                                                                fmt.Println(v17, v67)
                                                                if true {
                                                                    var v95 uint8 = (v47 ^ 0)
                                                                    if v2 {
                                                                        v2 = (v21 == (^(v21)))
                                                                        var v96 int8 = 0
                                                                        _ = v96
                                                                    } else {
                                                                        var v97 bool = v2
                                                                        fmt.Println(v68, v95, v66)
                                                                        var v98 string = ("foo bar" + ("hello" + "\n\t"))
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                    v2 = v2
                                                                    var v99 uint16 = v21
                                                                    var v100 uint32 = (v22 << ((255) & 63))
                                                                    _ = v100
                                                                    _ = v95
                                                                    _ = v99
                                                                } else {
                                                                    v101 := GenS10{F0: v21}
                                                                    _ = v101
                                                                }
                                                            } else {
                                                                var v102 string = "hello"
                                                                switch v102 {
                                                                case "":
                                                                    v1 = (((3.14159 + v1) * v1) + ((v1 * v1) - ((v1 - ((0.1 - 0.0) - 1e10)) / v1)))
                                                                case "\n\t":
                                                                    v103 := GenS9{F0: (fi51 ^ 4294967295)}
                                                                    _ = v103
                                                                case "hello":
                                                                    var v104 uint8 = (v47 - 28)
                                                                    _ = v104
                                                                }
                                                                _ = v102
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        v105 := GenS8{F0: (v21 < 1)}
                                                        if ((v67 / (((v67 / ((4294967295) | 1))) | 1)) <= (v67 << ((42) & 63))) {
                                                            switch (fi50 - fi50) {
                                                            case 32767:
                                                                if v68 {
                                                                    v106 := v4[-9223372036854775808]
                                                                    v17 = v17
                                                                    _ = v106
                                                                } else {
                                                                    var v107 float32 = ((-2.25 + 0.0) * 0.0)
                                                                    fmt.Println(v47, v22)
                                                                    _ = v107
                                                                }
                                                            case -389:
                                                                if v2 {
                                                                    if (738 > v21) {
                                                                        var v108 uint64 = 42
                                                                        _ = v108
                                                                    } else {
                                                                        v109 := v4[(-(fi51))]
                                                                        fmt.Println(v1, v1)
                                                                        _ = v109
                                                                    }
                                                                }
                                                                if len(v20) > 0 {
                                                                    v20[0] = (fi51 - (fi51 + fi50))
                                                                }
                                                            case -966:
                                                                v17 -= (v17 / ((-793) | 1))
                                                            default:
                                                                v110 := GenS11{F0: v2, F1: 166, F2: v68}
                                                                _ = v110
                                                            }
                                                            _ = len(v20)
                                                            if (((3.14159 * ((1.5 - ((((0.1 / 3.14159) / 1.0) - (-2.25 - 1.5)) / 3.14159)) + -2.25)) - 0.1) <= float32((v1 - v1))) {
                                                                v111 := GenS2{F0: 1, F1: fi51, F2: (v17 + (v17 << (((^(v67))) & 63)))}
                                                                fmt.Println(v17, v16, fi51)
                                                                if v2 {
                                                                    if (1 != 634) {
                                                                        var v112 uint64 = 4294967295
                                                                        _ = v112
                                                                    }
                                                                    var v113 string = ((("foo bar" + ("foo bar" + ("\n\t" + "hello"))) + "hello") + "")
                                                                    if (v21 == v21) {
                                                                        v3[2147483648] = true
                                                                        var v114 uint64 = 0
                                                                        _ = v114
                                                                    } else {
                                                                        v113 = v113
                                                                    }
                                                                    fmt.Println(v22)
                                                                    switch v113 {
                                                                    case "foo bar":
                                                                        v115 := GenS7{F0: v68, F1: v17, F2: 1}
                                                                        _ = v115
                                                                    }
                                                                    _ = v113
                                                                } else {
                                                                    if ((v66 | int16(v16)) >= v66) {
                                                                        fmt.Println(v17)
                                                                        v22 = v16
                                                                        v22 = 0
                                                                    }
                                                                    fmt.Println(v17, v2, v16)
                                                                    var v116 uint64 = uint64(v22)
                                                                    _ = v116
                                                                }
                                                                fmt.Println(fi50)
                                                                var v117 uint16 = 568
                                                                _ = v111
                                                                _ = v117
                                                            } else {
                                                                v118 := GenS11{F0: (!v2), F1: 32767, F2: v2}
                                                                v66 += (-(v66))
                                                                _ = v118
                                                            }
                                                        }
                                                        _ = v105
                                                    } else {
                                                        fmt.Println(v22)
                                                        var v119 string = "foo bar"
                                                        _ = v119
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi51
                            }
                            for fi120 := 0; fi120 < 12; fi120++ {
                                if (v17 < int64(v16)) {
                                    fmt.Println(v47, v21, v22)
                                }
                                fmt.Println(v47, v16)
                                _ = fi120
                            }
                        } else {
                            switch (("a" + (("a" + "foo bar") + (("a" + "hello") + ("a" + "hello")))) + "a") {
                            case "\n\t":
                                fmt.Println(v21)
                            }
                            v1 -= (((v1 / v1) + (v1 / v1)) + (0.1 / ((float64((-2.25 - ((((((0.1 / 1.0) - 1.5) / 1e10) + (-2.25 + 0.1)) + (0.1 - (0.1 / 1.0))) + ((1e10 * 0.1) - -2.25)))) / v1) / v1)))
                            v21 = 65535
                        }
                    }
                    _ = fi50
                }
            } else {
                if (v2 || v2) {
                    v2 = false
                }
                v22 = (v16 * v16)
            }
            fmt.Println(v47, v1, v47)
            var v121 int32 = -870
            _ = v121
        }
        _ = v47
    }
    var v122 uint32 = (v22 & (v22 ^ 2))
    v21 = 2
    v1 -= 0.1
    v16 += (v122 / ((399) | 1))
    v123 := GenS12{F0: (v17 | v17), F1: v17}
    for fi124 := 0; fi124 < 19; fi124++ {
        v17 -= (v17 | v17)
        if v2 {
            v22 = v22
            if false {
                v125 := v3[42]
                _ = v125
            } else {
                switch ("hello" + ("\n\t" + "\n\t")) {
                case "foo bar":
                    if true {
                        v126 := GenS13{F0: v22}
                        _ = v126
                    }
                case "a":
                    fmt.Println(v16)
                }
                switch 429 {
                case 32767:
                    switch (fi124 - fi124) {
                    case 127:
                        v122 = (v16 - v122)
                    case 1:
                        v17 = (v17 >> ((255) & 63))
                    }
                case 2147483647:
                    var v127 int = (fi124 | 4294967295)
                    _ = v127
                }
            }
            for fi128 := 0; fi128 < 20; fi128++ {
                v122 = v16
                _ = fi128
            }
        }
        _ = fi124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v122
    _ = v123
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
