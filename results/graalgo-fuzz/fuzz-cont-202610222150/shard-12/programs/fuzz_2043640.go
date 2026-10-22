package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int8; F1 uint16; F2 uint }
type GenS8 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int32; F1 int16; F2 bool }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 int64; F1 int16; F2 int64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS21 struct { F0 int8 }
func genFn0() uint16 {
    v7 := []int{2147483647, 32767, 9223372036854775807, -1}
    switch "" {
    case "a":
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                if (1e10 <= 1e10) {
                    if (!((!false) && ((-997 != -216) || (!true)))) {
                        var v9 int16 = 42
                        v9 += (v9 >> ((201) & 63))
                        var v10 uint = 42
                        _ = v10
                        _ = v9
                    } else {
                        var v11 uint64 = 1
                        _ = v11
                    }
                    var v12 uint32 = 1
                    var v13 int = (fi8 + 42)
                    _ = v12
                    _ = v13
                }
                var v14 string = ("" + "\n\t")
                _ = v14
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS1{F0: (fi15 * fi15), F1: 42, F2: 1}
                v17 := GenS2{F0: -2147483648, F1: 65535, F2: -712}
                var v18 int16 = 340
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 float64 = 3.14159
            _ = fi8
            _ = v19
        }
    case "\n\t":
        v20 := GenS3{F0: 0, F1: -32768}
        _ = v20
    case "foo bar":
        if true {
            v21 := GenS1{F0: 127, F1: 208, F2: 2}
            var v22 string = "a"
            _ = v21
            _ = v22
        }
    default:
        switch ("hello" + "hello") {
        case "a":
            switch ("" + (("hello" + "a") + "")) {
            case "foo bar":
                var v23 uint8 = 199
                _ = v23
            case "hello":
                v24 := GenS3{F0: 32767, F1: 1000}
                _ = v24
            }
        default:
            var v25 uint16 = 1
            _ = v25
        }
    }
    var v26 bool = (-70 != 502)
    v27 := GenS2{F0: -1, F1: 0, F2: 1}
    _ = v26
    _ = v27
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -464
    var v2 uint16 = 1
    v3 := make(map[string]string)
    v4 := []string{}
    var v5 int64 = -9223372036854775808
    var v6 uint64 = 4294967295
    v2 = (^(v2))
    v5 = 0
    v1 = v1
    v2 += (v2 << ((554) & 63))
    if ((v6 & v6) < 4294967295) {
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 float32 = 1.5
            _ = fi28
            _ = v29
        }
        delete(v3, "a")
    } else {
        var v30 int64 = v5
        v31 := GenS1{F0: int(v6), F1: 0, F2: 42}
        if false {
            for fi32 := 0; fi32 < 18; fi32++ {
                for fi33 := 0; fi33 < 6; fi33++ {
                    if true {
                        v34 := GenS2{F0: 2147483647, F1: 2, F2: -894}
                        if (("" + ("\n\t" + ("foo bar" + ("\n\t" + ("\n\t" + "hello"))))) == "hello") {
                            switch "foo bar" {
                            case "a":
                                switch (fi32 + fi32) {
                                case 127:
                                    var v35 float64 = (0.0 / float64(1.5))
                                    _ = v35
                                case -337:
                                    if true {
                                        if false {
                                            v36 := GenS3{F0: (v2 << ((4294967295) & 63)), F1: 70}
                                            v37 := GenS4{F0: v6, F1: (fi33 * fi33)}
                                            var v38 uint = 18446744073709551615
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                        v39 := GenS5{F0: (!false), F1: 2147483647}
                                        _ = v39
                                    } else {
                                        v3[""] = "a"
                                        if (1.5 < 0.0) {
                                            fmt.Println(v6, fi33)
                                            v4 = append(v4, "")
                                        } else {
                                            v2 = 65535
                                            var v40 uint16 = (v2 / (((^(v2))) | 1))
                                            _ = v40
                                        }
                                    }
                                }
                                v5 += v5
                                if false {
                                    switch fi32 {
                                    case 817:
                                        var v41 uint = 42
                                        var v42 int8 = 127
                                        v42 += 125
                                        _ = v41
                                        _ = v42
                                    case -9223372036854775808:
                                        v43 := GenS1{F0: 9223372036854775807, F1: 1, F2: 193}
                                        _ = v43
                                    default:
                                        v44 := v3["hello"]
                                        _ = v44
                                    }
                                    switch 42 {
                                    case 9223372036854775807:
                                        v2 -= genFn0()
                                        if false {
                                            v5 = (v30 >> ((4294967295) & 63))
                                        } else {
                                            if false {
                                                fmt.Println(v1, v30, v5)
                                                v3[("" + ("\n\t" + "\n\t"))] = ("\n\t" + "")
                                                var v45 int8 = 0
                                                _ = v45
                                            }
                                            var v46 uint32 = 4294967295
                                            _ = v46
                                        }
                                        v2 += (v2 + (v2 | v2))
                                        var v47 uint16 = genFn0()
                                        _ = v47
                                    case -2147483648:
                                        if true {
                                            var v48 float64 = (0.0 + (1.5 + -2.25))
                                            if (!("" > ("hello" + "\n\t"))) {
                                                v49 := GenS6{F0: "a", F1: fi33}
                                                if (0 != v6) {
                                                    fmt.Println(fi33, v2)
                                                } else {
                                                    if true {
                                                        if (((!false) && true) && true) {
                                                            fmt.Println(v2, v1)
                                                            switch (fi32 << ((42) & 63)) {
                                                            case -510:
                                                                switch ("foo bar" + (("foo bar" + "a") + "hello")) {
                                                                case "\n\t":
                                                                    v50 := GenS4{F0: v6, F1: fi33}
                                                                    _ = v50
                                                                default:
                                                                    if true {
                                                                        var v51 int64 = v30
                                                                        _ = v51
                                                                    } else {
                                                                        if true {
                                                                            var v52 uint = 1
                                                                            fmt.Println(v6)
                                                                            var v53 uint = (v52 ^ (v52 % (((v52 ^ (v52 / (((v52 << ((v52) & 63))) | 1)))) | 1)))
                                                                            _ = v52
                                                                            _ = v53
                                                                        } else {
                                                                            v54 := GenS7{F0: -16, F1: 32767, F2: 2}
                                                                            v55 := GenS8{F0: (v2 / (((v2 ^ v2)) | 1)), F1: 255, F2: (v5 * -536)}
                                                                            delete(v3, "")
                                                                            v56 := GenS9{F0: 127}
                                                                            v57 := GenS10{F0: -2147483648}
                                                                            fmt.Println(v1)
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                            default:
                                                                switch fi32 {
                                                                case 127:
                                                                    if ((("hello" + "\n\t") + "a") < ("" + "")) {
                                                                        fmt.Println(fi33, v30, v2)
                                                                        if (!(true || (true && false))) {
                                                                            v58 := GenS8{F0: genFn0(), F1: uint8(v6), F2: -9223372036854775808}
                                                                            var v59 string = ((("foo bar" + (("\n\t" + "foo bar") + "\n\t")) + "") + ("" + ("" + "")))
                                                                            v60 := GenS11{F0: 955, F1: (v2 % ((v2) | 1)), F2: (v6 >> ((255) & 63))}
                                                                            v61 := GenS7{F0: -128, F1: (v2 % (((v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))) | 1)), F2: 2}
                                                                            _ = v58
                                                                            _ = v59
                                                                            _ = v60
                                                                            _ = v61
                                                                        } else {
                                                                            v62 := GenS12{F0: 42, F1: 32767, F2: false}
                                                                            v63 := GenS6{F0: "a", F1: fi33}
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                var v64 int8 = 127
                                                                                _ = v64
                                                                            case "foo bar":
                                                                                switch fi32 {
                                                                                case -476:
                                                                                    var v65 int16 = (^(v1))
                                                                                    _ = v65
                                                                                case -622:
                                                                                    if (!true) {
                                                                                        var v66 bool = true
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v66
                                                                                    }
                                                                                case 775:
                                                                                    v48 = (float64(3.14159) / v48)
                                                                                default:
                                                                                    fmt.Println(v1)
                                                                                    switch fi33 {
                                                                                    case 1:
                                                                                        v6 += (v6 | v6)
                                                                                    case 42:
                                                                                        var v67 float64 = v48
                                                                                        _ = v67
                                                                                    default:
                                                                                        v2 = (v2 | 32767)
                                                                                    }
                                                                                    if false {
                                                                                        var v68 float64 = v48
                                                                                        v2 -= (v2 - v2)
                                                                                        switch (fi33 + fi33) {
                                                                                        case 127:
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    var v69 int32 = -2147483648
                                                                                                    _ = v69
                                                                                                default:
                                                                                                    var v70 int32 = 42
                                                                                                    _ = v70
                                                                                                }
                                                                                                if false {
                                                                                                    if true {
                                                                                                        v68 -= v68
                                                                                                    }
                                                                                                    if (fi32 <= (fi32 % (((^(fi33))) | 1))) {
                                                                                                        if ((genFn0() > (v2 >> ((1) & 63))) || true) {
                                                                                                            if false {
                                                                                                                v2 -= 255
                                                                                                            }
                                                                                                        }
                                                                                                        var v71 float32 = float32(0.0)
                                                                                                        v30 -= (v30 / (((v30 >> ((18446744073709551615) & 63))) | 1))
                                                                                                        v72 := GenS13{F0: ("" + "hello"), F1: 641}
                                                                                                        fmt.Println(fi33, v68)
                                                                                                        _ = v71
                                                                                                        _ = v72
                                                                                                    } else {
                                                                                                        var v73 int64 = (v5 + 42)
                                                                                                        _ = v73
                                                                                                    }
                                                                                                    switch "foo bar" {
                                                                                                    case "\n\t":
                                                                                                        if ((v2 % ((v2) | 1)) != uint16(v6)) {
                                                                                                            var v74 int8 = 100
                                                                                                            v75 := GenS5{F0: (0 >= 68), F1: 707}
                                                                                                            _ = v74
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v76 := GenS14{F0: v5, F1: 1000, F2: -634}
                                                                                                        _ = v76
                                                                                                    default:
                                                                                                        var v77 uint = 0
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    var v78 int8 = 0
                                                                                                    v79 := GenS13{F0: "\n\t", F1: 4294967295}
                                                                                                    var v80 uint64 = v6
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                } else {
                                                                                                    v6 = v6
                                                                                                    fmt.Println(v1, v2)
                                                                                                    var v81 int64 = v5
                                                                                                    _ = v81
                                                                                                }
                                                                                            }
                                                                                        case -103:
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = "\n\t"
                                                                                            }
                                                                                        default:
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                switch fi32 {
                                                                                                case 4294967295:
                                                                                                    var v82 uint32 = 42
                                                                                                    _ = v82
                                                                                                case 32768:
                                                                                                    v6 = v6
                                                                                                default:
                                                                                                    var v83 uint = 255
                                                                                                    _ = v83
                                                                                                }
                                                                                            case "":
                                                                                                v2 = v2
                                                                                            default:
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = "foo bar"
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            v4 = append(v4, ("foo bar" + ("foo bar" + "")))
                                                                                            var v84 uint32 = uint32(v2)
                                                                                            if (!true) {
                                                                                                if false {
                                                                                                    v5 += v30
                                                                                                    if (127 >= 119) {
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                v1 += (v1 & 42)
                                                                                                                if (v30 > (v5 % (((v5 | (-(v5)))) | 1))) {
                                                                                                                    if (!(0 > 42)) {
                                                                                                                        v84 += v84
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = ""
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        var v85 int = fi32
                                                                                                                        _ = v85
                                                                                                                    } else {
                                                                                                                        var v86 int64 = (^(v30))
                                                                                                                        v87 := GenS6{F0: "", F1: 9223372036854775807}
                                                                                                                        _ = v86
                                                                                                                        _ = v87
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            v2 = (v2 * (v2 & (v2 + genFn0())))
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                var v88 float64 = 0.0
                                                                                                                switch fi33 {
                                                                                                                case 0:
                                                                                                                    var v89 uint16 = v2
                                                                                                                    _ = v89
                                                                                                                case 1:
                                                                                                                    v88 = v68
                                                                                                                }
                                                                                                                v90 := v3[("foo bar" + "")]
                                                                                                                _ = v88
                                                                                                                _ = v90
                                                                                                            } else {
                                                                                                                v2 = uint16(v84)
                                                                                                            }
                                                                                                        }
                                                                                                        v91 := GenS12{F0: 1, F1: v1, F2: true}
                                                                                                        _ = v91
                                                                                                    }
                                                                                                } else {
                                                                                                    v3["hello"] = ("foo bar" + ("a" + ""))
                                                                                                    var v92 string = "a"
                                                                                                    v68 = v68
                                                                                                    _ = v92
                                                                                                }
                                                                                            }
                                                                                            _ = v84
                                                                                        }
                                                                                        v93 := GenS14{F0: v5, F1: v1, F2: v5}
                                                                                        _ = v68
                                                                                        _ = v93
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                v94 := GenS15{F0: 2}
                                                                                _ = v94
                                                                            }
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                default:
                                                                    v30 = v30
                                                                }
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        v2 += 42
                                                        if false {
                                                            fmt.Println(v2, v30, v6)
                                                            v95 := GenS10{F0: -2147483648}
                                                            v96 := GenS16{F0: true}
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            if true {
                                                                if (!true) {
                                                                    var v97 bool = false
                                                                    if ("foo bar" <= "\n\t") {
                                                                        var v98 uint32 = 648
                                                                        var v99 int16 = (v1 % ((32767) | 1))
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    _ = v97
                                                                }
                                                                switch (fi33 >> ((2) & 63)) {
                                                                case 2147483647:
                                                                    switch (("" + "foo bar") + (("hello" + "foo bar") + "foo bar")) {
                                                                    case "":
                                                                        v1 = (v1 - (v1 % ((v1) | 1)))
                                                                        fmt.Println(v48, v30)
                                                                        var v100 uint64 = 338
                                                                        v5 = (v30 << ((4294967295) & 63))
                                                                        _ = v100
                                                                    case "hello":
                                                                        v48 = 0.0
                                                                    case "foo bar":
                                                                        v2 = (v2 << ((444) & 63))
                                                                    }
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                            } else {
                                                                if false {
                                                                    fmt.Println(v1, v48, v30)
                                                                } else {
                                                                    if (42 <= (v6 | 18446744073709551615)) {
                                                                        var v101 int64 = (v5 | (v5 << ((0) & 63)))
                                                                        if true {
                                                                            switch fi33 {
                                                                            case 32767:
                                                                                v102 := GenS6{F0: "hello", F1: fi32}
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                        fmt.Println(v2, fi32)
                                                                        v103 := GenS5{F0: ((!false) || false), F1: -552}
                                                                        var v104 uint8 = 250
                                                                        _ = v101
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                    var v105 float64 = 0.1
                                                                    var v106 int = 397
                                                                    _ = v105
                                                                    _ = v106
                                                                }
                                                                var v107 uint = 1
                                                                switch ("hello" + "") {
                                                                case "a":
                                                                    _ = len(v3)
                                                                    switch fi32 {
                                                                    case -239:
                                                                        v108 := GenS4{F0: (v6 ^ v6), F1: (fi33 & fi33)}
                                                                        _ = v108
                                                                    case 2147483647:
                                                                        var v109 int32 = 42
                                                                        _ = v109
                                                                    case 802:
                                                                        switch 0 {
                                                                        case -2147483648:
                                                                            fmt.Println(v30, fi33)
                                                                        case 42:
                                                                            var v110 float64 = v48
                                                                            _ = v110
                                                                        case -368:
                                                                            var v111 int64 = int64(v6)
                                                                            _ = v111
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v112 := GenS17{F0: 1}
                                                                        var v113 int = -2147483648
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        switch (fi32 % ((987) | 1)) {
                                                                        case -2147483648:
                                                                            fmt.Println(v107)
                                                                        case 2147483647:
                                                                            v107 = (^(v107))
                                                                        case -9223372036854775808:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ("hello" + "")
                                                                            }
                                                                        }
                                                                        v114 := GenS18{F0: 865}
                                                                        var v115 uint64 = v6
                                                                        if true {
                                                                            var v116 int = fi33
                                                                            _ = v116
                                                                        }
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                default:
                                                                    var v117 float64 = v48
                                                                    v2 = genFn0()
                                                                    v6 -= (v6 << ((v107) & 63))
                                                                    var v118 uint16 = v2
                                                                    v3[("foo bar" + "foo bar")] = "\n\t"
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                                _ = v107
                                                            }
                                                            var v119 float32 = (1.5 / float32(v48))
                                                            _ = v119
                                                        }
                                                    } else {
                                                        var v120 int16 = (v1 * v1)
                                                        _ = v120
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v121 := GenS19{F0: 42}
                                                    _ = v121
                                                }
                                                if (4294967295 >= 42) {
                                                    var v122 bool = (v2 == v2)
                                                    switch fi32 {
                                                    case 9223372036854775807:
                                                        v123 := GenS20{F0: uint8(v2), F1: uint8(v2), F2: (^(v2))}
                                                        _ = v123
                                                    }
                                                    v2 -= (v2 << ((186) & 63))
                                                    _ = v122
                                                }
                                                _ = v49
                                            }
                                            _ = v48
                                        } else {
                                            var v124 int32 = -464
                                            _ = v124
                                        }
                                    default:
                                        var v125 int = -9223372036854775808
                                        _ = v125
                                    }
                                    v3["foo bar"] = "a"
                                } else {
                                    v2 = (v2 & (v2 ^ genFn0()))
                                }
                            case "\n\t":
                                v30 = v5
                            }
                            v3["\n\t"] = "hello"
                        } else {
                            var v126 int8 = -118
                            _ = v126
                        }
                        v127 := GenS21{F0: int8(v5)}
                        _ = v127
                        _ = v34
                    }
                    _ = fi33
                }
                _ = fi32
            }
            if true {
                switch "foo bar" {
                case "a":
                    var v128 uint8 = 42
                    _ = v128
                case "foo bar":
                    if (4294967295 != -2147483648) {
                        v129 := GenS11{F0: -170, F1: uint16(v6), F2: v6}
                        var v130 string = "\n\t"
                        _ = v129
                        _ = v130
                    } else {
                        _ = len(v4)
                    }
                case "":
                    v131 := GenS12{F0: 2147483647, F1: v1, F2: (false && (199 > (v1 ^ v1)))}
                    _ = v131
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v132 uint32 = 4294967295
                _ = v132
            }
        } else {
            var v133 bool = (true || true)
            _ = v133
        }
        _ = v30
        _ = v31
    }
    v5 += (-(v5))
    _ = len(v3)
    fmt.Println(v6, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
