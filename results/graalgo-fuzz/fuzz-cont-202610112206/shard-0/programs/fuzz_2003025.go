package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS5 struct { F0 bool; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint; F1 int64; F2 uint }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 int64; F2 int64 }
type GenS16 struct { F0 int64; F1 uint; F2 uint8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint64; F1 uint }
func genFn0(p0 uint16, p1 int64) uint64 {
    var v6 int = 1
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            for fi8 := 0; fi8 < 15; fi8++ {
                v6 = v6
                p1 = 104
                v9 := []int{(v6 % ((fi7) | 1)), 127}
                _ = fi8
                _ = v9
            }
        }
        var v10 int64 = 9223372036854775807
        var v11 uint64 = 4294967295
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: (v6 / (((v6 - v6)) | 1))}
    v13 := GenS1{F0: (v6 | (v6 + (-(v6))))}
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint8 = 117
            fmt.Println(v16, p1, fi15)
            _ = fi15
            _ = v16
        }
        v17 := []int{-1}
        _ = fi14
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v6
    return 2
}

func main() {
    v1 := []string{"foo bar", "foo bar", ("hello" + (("\n\t" + "\n\t") + "\n\t")), ((("" + "hello") + ("a" + "a")) + (("\n\t" + ("a" + "\n\t")) + (("a" + "a") + (("foo bar" + (("" + ("foo bar" + "\n\t")) + "foo bar")) + (("a" + "a") + (("foo bar" + ("foo bar" + ("hello" + (("" + "\n\t") + ("hello" + ((("" + "a") + "") + "\n\t")))))) + (("foo bar" + "\n\t") + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar"))))))))}
    v2 := []uint32{2147483647, 42}
    v3 := make(map[int]string)
    v4 := []uint32{}
    v5 := []int{9223372036854775807, 0, -805}
    var v18 int16 = -1
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 float32 = 0.0
        if false {
            v21 := GenS1{F0: fi19}
            _ = v21
        }
        v3[(fi19 << ((0) & 63))] = ("hello" + "\n\t")
        _ = fi19
        _ = v20
    }
    var v22 int64 = 0
    fmt.Println(v22, v22)
    for fi23 := 0; fi23 < 5; fi23++ {
        if (1e10 <= ((1e10 / (0.0 + (-2.25 / 1e10))) / 0.1)) {
            if ("hello" != "foo bar") {
                switch (fi23 >> ((0) & 63)) {
                case -862:
                    var v24 int16 = (v18 << ((4294967295) & 63))
                    _ = v24
                case -1:
                    v25 := GenS2{F0: "\n\t", F1: -32768, F2: -2147483648}
                    _ = v25
                default:
                    switch -9223372036854775808 {
                    case 19:
                        for fi26 := 0; fi26 < 14; fi26++ {
                            if (v18 < v18) {
                                var v27 uint16 = 42
                                _ = v27
                            } else {
                                switch (fi26 >> ((42) & 63)) {
                                case -1:
                                    v28 := GenS2{F0: "foo bar", F1: v18, F2: 1}
                                    _ = v28
                                case 114:
                                    switch "foo bar" {
                                    case "a":
                                        if false {
                                            switch "a" {
                                            case "hello":
                                                if ((!true) || ("" == "hello")) {
                                                    fmt.Println(v18)
                                                    var v29 int = fi26
                                                    if false {
                                                        var v30 string = ((("foo bar" + ("\n\t" + "foo bar")) + (("\n\t" + ("foo bar" + "")) + (("hello" + "hello") + "hello"))) + ("" + "a"))
                                                        var v31 float64 = 1.5
                                                        fmt.Println(v31, v30, v30)
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    v32 := GenS2{F0: "foo bar", F1: v18, F2: 1}
                                                    if true {
                                                        v33 := GenS2{F0: (("a" + "a") + "a"), F1: -996, F2: 1}
                                                        fmt.Println(v18, v22, v18)
                                                        _ = v33
                                                    }
                                                    _ = v29
                                                    _ = v32
                                                }
                                            default:
                                                v18 = v18
                                            }
                                            v34 := GenS3{F0: (true || false), F1: 873}
                                            if (false && (0 < v18)) {
                                                _ = len(v2)
                                                if true {
                                                    v35 := GenS3{F0: false, F1: 255}
                                                    v36 := GenS2{F0: ("a" + (("foo bar" + ("a" + "\n\t")) + "a")), F1: (v18 - v18), F2: 2147483647}
                                                    switch fi23 {
                                                    case -74:
                                                        var v37 uint32 = 42
                                                        _ = v37
                                                    case 128:
                                                        if ((!true) || (false && true)) {
                                                            v18 += (v18 - v18)
                                                            v4 = append(v4, 367)
                                                        } else {
                                                            var v38 int32 = -2147483648
                                                            if true {
                                                                var v39 uint32 = 2147483647
                                                                _ = v39
                                                            } else {
                                                                v22 = v22
                                                            }
                                                            var v40 int8 = 8
                                                            _ = v38
                                                            _ = v40
                                                        }
                                                    case 416:
                                                        v18 = v18
                                                    default:
                                                        v41 := GenS1{F0: (fi23 << ((42) & 63))}
                                                        _ = v41
                                                    }
                                                    _ = v35
                                                    _ = v36
                                                }
                                            } else {
                                                fmt.Println(v18, fi23)
                                            }
                                            _ = v34
                                        }
                                    case "\n\t":
                                        v42 := GenS3{F0: (!true), F1: 984}
                                        _ = v42
                                    case "foo bar":
                                        if false {
                                            if false {
                                                v22 = (v22 * 42)
                                                if len(v2) > 0 {
                                                    v2[0] = 899
                                                }
                                                v22 -= v22
                                            } else {
                                                var v43 int8 = 0
                                                var v44 int = (fi23 + fi23)
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 int16 = (v18 << ((2) & 63))
                                            switch (-(fi23)) {
                                            case 32768:
                                                fmt.Println(fi23, fi23, v22)
                                            case 252:
                                                if false {
                                                    var v46 string = ("hello" + "\n\t")
                                                    var v47 float64 = (1.5 - ((3.14159 - (1e10 * -2.25)) - (((3.14159 + 1.5) * 1.5) + ((0.1 + ((1e10 - 1e10) + (0.1 / float64(3.14159)))) * 3.14159))))
                                                    var v48 int8 = 127
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    var v49 uint16 = 32767
                                                    v45 -= v18
                                                    fmt.Println(fi23)
                                                    v18 = (v18 ^ (^(v45)))
                                                    _ = v49
                                                }
                                            }
                                            _ = v45
                                        }
                                    }
                                }
                                if true {
                                    switch 128 {
                                    case 607:
                                        if false {
                                            v22 = (v22 + v22)
                                            var v50 uint = 42
                                            _ = v50
                                        }
                                        v18 += (^(v18))
                                        v18 -= 1
                                    case 450:
                                        fmt.Println(fi23, fi26, v22)
                                        v51 := GenS3{F0: (72 > -128), F1: 18446744073709551615}
                                        if true {
                                            v22 = v22
                                        }
                                        if true {
                                            var v52 int = 32768
                                            _ = v52
                                        }
                                        _ = v51
                                    case 1:
                                        fmt.Println(v22, v22)
                                    default:
                                        _ = len(v3)
                                    }
                                    v22 = 1
                                    v53 := v3[-990]
                                    _ = v53
                                } else {
                                    var v54 int = fi26
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    var v55 uint8 = 2
                                    _ = v54
                                    _ = v55
                                }
                                var v56 float64 = 1.5
                                _ = v56
                            }
                            if (v22 >= v22) {
                                var v57 float32 = 0.0
                                _ = v57
                            } else {
                                switch (fi23 << ((18446744073709551615) & 63)) {
                                case 887:
                                    if true {
                                        _ = len(v5)
                                        var v58 uint32 = 751
                                        fmt.Println(v18, v22, fi23)
                                        _ = v58
                                    } else {
                                        v18 -= v18
                                        var v59 uint32 = 239
                                        if (!((fi26 & (fi26 * int(v18))) <= 791)) {
                                            v60 := GenS4{F0: 693, F1: genFn0(32767, (v22 + v22)), F2: 1000}
                                            fmt.Println(fi26, fi23, fi23)
                                            v22 += v22
                                            _ = v60
                                        } else {
                                            v22 -= (v22 | (-(v22)))
                                            if (v59 <= v59) {
                                                var v61 uint32 = v59
                                                var v62 uint = 42
                                                switch "hello" {
                                                case "hello":
                                                    v59 -= v61
                                                case "foo bar":
                                                    var v63 int64 = v22
                                                    _ = v63
                                                case "a":
                                                    fmt.Println(fi23, v18, fi26)
                                                default:
                                                    var v64 uint64 = 981
                                                    _ = v64
                                                }
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                        _ = v59
                                    }
                                }
                                switch "hello" {
                                case "":
                                    v18 += (v18 ^ -1)
                                case "\n\t":
                                    if (216 <= 99) {
                                        fmt.Println(v18)
                                        var v65 uint16 = 65535
                                        if (1 >= 255) {
                                            fmt.Println(v65)
                                            v66 := GenS5{F0: true, F1: true, F2: 1}
                                            switch "foo bar" {
                                            case "a":
                                                if true {
                                                    if (!false) {
                                                        v67 := GenS2{F0: "", F1: -815, F2: -137}
                                                        _ = v67
                                                    } else {
                                                        fmt.Println(v65)
                                                        if len(v4) > 0 {
                                                            v4[0] = 1
                                                        }
                                                        v22 = v22
                                                    }
                                                    switch fi23 {
                                                    case -1:
                                                        v65 -= (^(v65))
                                                    case 128:
                                                        var v68 uint = 2
                                                        _ = v68
                                                    case 242:
                                                        if true {
                                                            switch "\n\t" {
                                                            case "a":
                                                                fmt.Println(v22)
                                                            case "hello":
                                                                fmt.Println(v65, v22)
                                                            default:
                                                                switch "a" {
                                                                case "a":
                                                                    if (-2147483648 != -2147483648) {
                                                                        v65 = 429
                                                                        if true {
                                                                            v69 := GenS4{F0: 65535, F1: genFn0(v65, v22), F2: v18}
                                                                            if len(v1) > 0 {
                                                                                v1[0] = "\n\t"
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                    } else {
                                                                        var v70 uint8 = 24
                                                                        if true {
                                                                            if (42 >= 100) {
                                                                                v71 := GenS4{F0: v65, F1: 18446744073709551615, F2: v18}
                                                                                _ = v71
                                                                            }
                                                                            v18 = -32768
                                                                            v4 = append(v4, 0)
                                                                        } else {
                                                                            if false {
                                                                                v22 += v22
                                                                                if true {
                                                                                    v70 += v70
                                                                                    if ((2 >= 10) && true) {
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 422
                                                                                        }
                                                                                        v18 = v18
                                                                                        if true {
                                                                                            delete(v3, 2147483647)
                                                                                        } else {
                                                                                            fmt.Println(v70, v65)
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            switch (fi23 * fi23) {
                                                                                            case -1:
                                                                                                fmt.Println(fi26, v22)
                                                                                            case 32768:
                                                                                                if true {
                                                                                                    v3[(fi23 ^ (fi26 / ((fi26) | 1)))] = ("foo bar" + ("\n\t" + "foo bar"))
                                                                                                } else {
                                                                                                    switch fi23 {
                                                                                                    case -342:
                                                                                                        v18 -= -533
                                                                                                    case -2147483648:
                                                                                                        switch 4294967295 {
                                                                                                        case 42:
                                                                                                            fmt.Println(v18, v70, v70)
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(fi26, v18)
                                                                                                    var v72 int8 = 0
                                                                                                    _ = v72
                                                                                                }
                                                                                            case -907:
                                                                                                v18 = v18
                                                                                            default:
                                                                                                switch ("foo bar" + "hello") {
                                                                                                case "foo bar":
                                                                                                    var v73 int8 = 100
                                                                                                    _ = v73
                                                                                                case "\n\t":
                                                                                                    v74 := GenS6{F0: (v65 + 579)}
                                                                                                    _ = v74
                                                                                                case "":
                                                                                                    var v75 string = ("a" + ("foo bar" + "a"))
                                                                                                    var v76 bool = true
                                                                                                    _ = v75
                                                                                                    _ = v76
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            switch (("hello" + (("hello" + "foo bar") + ("hello" + ""))) + "hello") {
                                                                                            case "a":
                                                                                                var v77 uint64 = 2
                                                                                                _ = v77
                                                                                            case "":
                                                                                                if (v65 == 32767) {
                                                                                                    v78 := GenS4{F0: v65, F1: 1, F2: (v18 << ((255) & 63))}
                                                                                                    v65 += (v65 / ((42) | 1))
                                                                                                    var v79 string = ""
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                } else {
                                                                                                    fmt.Println(v65, fi23, v22)
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v80 := GenS1{F0: fi23}
                                                                                                _ = v80
                                                                                            }
                                                                                        default:
                                                                                            v22 = 0
                                                                                        }
                                                                                        v18 -= (v18 << ((255) & 63))
                                                                                    }
                                                                                    var v81 int64 = (^(v22))
                                                                                    v22 -= (v81 & 809)
                                                                                    _ = v81
                                                                                }
                                                                                var v82 bool = (42 < 0)
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        var v83 float64 = 1.5
                                                                        _ = v70
                                                                        _ = v83
                                                                    }
                                                                case "":
                                                                    v84 := GenS5{F0: false, F1: false, F2: 1}
                                                                    _ = v84
                                                                case "\n\t":
                                                                    v85 := GenS7{F0: 105, F1: 185}
                                                                    if (-48 == -53) {
                                                                        if true {
                                                                            var v86 int = (fi23 & (fi26 << ((18446744073709551615) & 63)))
                                                                            _ = len(v3)
                                                                            fmt.Println(v18, v86, v86)
                                                                            _ = v86
                                                                        }
                                                                        v87 := GenS7{F0: 42, F1: 0}
                                                                        _ = v87
                                                                    } else {
                                                                        if false {
                                                                            if true {
                                                                                if false {
                                                                                    fmt.Println(fi26, v65)
                                                                                    v88 := GenS8{F0: 255, F1: v22, F2: 255}
                                                                                    var v89 float64 = 1e10
                                                                                    v4 = append(v4, 1)
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                } else {
                                                                                    v90 := GenS6{F0: 1}
                                                                                    if (true && true) {
                                                                                        switch ("a" + "\n\t") {
                                                                                        case "a":
                                                                                            var v91 uint8 = 2
                                                                                            _ = v91
                                                                                        default:
                                                                                            v18 += (v18 ^ v18)
                                                                                        }
                                                                                    }
                                                                                    if ((fi26 - -9223372036854775808) >= (fi23 / ((fi23) | 1))) {
                                                                                        v22 += v22
                                                                                        var v92 int = (-(fi26))
                                                                                        _ = v92
                                                                                    }
                                                                                    _ = v90
                                                                                }
                                                                                switch fi23 {
                                                                                case -865:
                                                                                    if (!((fi26 & fi23) < fi26)) {
                                                                                        v93 := v3[(fi26 | fi26)]
                                                                                        var v94 uint32 = 42
                                                                                        if (fi26 == fi26) {
                                                                                            var v95 string = "foo bar"
                                                                                            var v96 float64 = (float64(1.5) / 1e10)
                                                                                            var v97 int = -389
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        } else {
                                                                                            var v98 bool = false
                                                                                            if false {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 2147483648
                                                                                                }
                                                                                                if (v65 < uint16(v94)) {
                                                                                                    v93 = v93
                                                                                                    v93 = v93
                                                                                                    if v98 {
                                                                                                        var v99 string = "a"
                                                                                                        fmt.Println(v99, v99)
                                                                                                        v100 := GenS9{F0: 2, F1: -1}
                                                                                                        fmt.Println(fi23, fi23, v22)
                                                                                                        v101 := GenS6{F0: v65}
                                                                                                        v102 := GenS9{F0: 127, F1: 42}
                                                                                                        v98 = (v98 || ((v65 >> ((1) & 63)) <= v65))
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                        _ = v102
                                                                                                        _ = v99
                                                                                                    } else {
                                                                                                        var v103 bool = ((!v98) && v98)
                                                                                                        _ = v103
                                                                                                    }
                                                                                                    var v104 int64 = (v22 / ((v22) | 1))
                                                                                                    _ = v104
                                                                                                }
                                                                                                _ = len(v3)
                                                                                                v65 = (v65 & v65)
                                                                                                v105 := GenS10{F0: 42, F1: 56}
                                                                                                if (!true) {
                                                                                                    v93 = ""
                                                                                                } else {
                                                                                                    switch fi23 {
                                                                                                    case -1:
                                                                                                        fmt.Println(v18, v65)
                                                                                                    case 0:
                                                                                                        v4 = append(v4, v94)
                                                                                                    case 127:
                                                                                                        v106 := GenS11{F0: true, F1: 255}
                                                                                                        _ = v106
                                                                                                    default:
                                                                                                        v107 := GenS9{F0: 246, F1: int32(v94)}
                                                                                                        _ = v107
                                                                                                    }
                                                                                                }
                                                                                                var v108 int32 = 1
                                                                                                fmt.Println(v18)
                                                                                                _ = v105
                                                                                                _ = v108
                                                                                            } else {
                                                                                                switch fi26 {
                                                                                                case 361:
                                                                                                    fmt.Println(v98, v22, fi23)
                                                                                                    if v98 {
                                                                                                        v22 = (v22 * v22)
                                                                                                        var v109 bool = (-115 < v22)
                                                                                                        _ = v109
                                                                                                    }
                                                                                                case 32767:
                                                                                                    var v110 int8 = 100
                                                                                                    _ = v110
                                                                                                case 32768:
                                                                                                    var v111 int8 = 0
                                                                                                    _ = v111
                                                                                                }
                                                                                                v93 = (v93 + v93)
                                                                                                var v112 int8 = -128
                                                                                                _ = v112
                                                                                            }
                                                                                            v113 := GenS10{F0: -128, F1: 100}
                                                                                            _ = v113
                                                                                            _ = v98
                                                                                        }
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                default:
                                                                                    var v114 int16 = 42
                                                                                    _ = v114
                                                                                }
                                                                                fmt.Println(v18, v65)
                                                                            }
                                                                            fmt.Println(v65, fi26)
                                                                            v22 = -1
                                                                        } else {
                                                                            var v115 uint64 = 962
                                                                            _ = v115
                                                                        }
                                                                        v65 -= v65
                                                                    }
                                                                    var v116 uint32 = 79
                                                                    v65 += v65
                                                                    v3[fi23] = ""
                                                                    if (true && true) {
                                                                        if false {
                                                                            var v117 uint8 = 127
                                                                            if true {
                                                                                fmt.Println(v117, v22, v116)
                                                                                var v118 int8 = -1
                                                                                _ = v118
                                                                            }
                                                                            var v119 uint32 = 2
                                                                            _ = v117
                                                                            _ = v119
                                                                        } else {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            var v120 int = (^(fi26))
                                                                            _ = v120
                                                                        }
                                                                        fmt.Println(v65, v65)
                                                                    } else {
                                                                        v65 -= 660
                                                                    }
                                                                    _ = v116
                                                                    _ = v85
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        v121 := GenS2{F0: "foo bar", F1: (v18 & v18), F2: 0}
                                                        _ = v121
                                                    }
                                                    switch "a" {
                                                    case "":
                                                        v122 := GenS3{F0: false, F1: 42}
                                                        _ = v122
                                                    }
                                                    v22 = v22
                                                } else {
                                                    if (42 >= (fi26 / ((fi23) | 1))) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        v123 := GenS12{F0: v18, F1: fi23}
                                                        _ = v123
                                                    }
                                                }
                                                delete(v3, fi23)
                                                v18 -= (v18 + v18)
                                                var v124 string = "a"
                                                _ = v124
                                            default:
                                                if ((true || true) && ((fi23 | 127) != (fi23 ^ fi23))) {
                                                    v125 := GenS6{F0: (v65 * 1)}
                                                    fmt.Println(fi23, v22)
                                                    delete(v3, (fi26 + 492))
                                                    switch fi26 {
                                                    case 160:
                                                        var v126 int16 = (^(v18))
                                                        _ = v126
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 2
                                                    }
                                                    _ = v125
                                                } else {
                                                    v127 := GenS4{F0: 0, F1: 433, F2: v18}
                                                    var v128 int64 = (v22 - 1)
                                                    if true {
                                                        v3[(fi23 - fi23)] = ""
                                                    } else {
                                                        if true {
                                                            v128 = (v128 >> ((42) & 63))
                                                            v129 := GenS13{F0: uint32(v65)}
                                                            if ("hello" > "a") {
                                                                v130 := GenS14{F0: (v65 & (v65 ^ v65))}
                                                                v22 = 42
                                                                _ = v130
                                                            } else {
                                                                if false {
                                                                    _ = len(v4)
                                                                    var v131 uint64 = 831
                                                                    var v132 uint8 = 60
                                                                    _ = v131
                                                                    _ = v132
                                                                } else {
                                                                    v133 := GenS2{F0: ("foo bar" + "hello"), F1: (v18 ^ 862), F2: 0}
                                                                    var v134 uint8 = 127
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                v65 -= (^(v65))
                                                            }
                                                            _ = v129
                                                        } else {
                                                            v3[(fi26 << ((255) & 63))] = "hello"
                                                            delete(v3, 32767)
                                                        }
                                                        var v135 uint8 = 255
                                                        switch (fi26 | (fi23 >> ((4294967295) & 63))) {
                                                        case 9223372036854775807:
                                                            if true {
                                                                v136 := GenS15{F0: 42, F1: 143, F2: v128}
                                                                _ = v136
                                                            }
                                                        case 32767:
                                                            v128 -= -9223372036854775808
                                                        }
                                                        var v137 int64 = v128
                                                        _ = v135
                                                        _ = v137
                                                    }
                                                    v65 += v65
                                                    var v138 string = ("" + "hello")
                                                    _ = v127
                                                    _ = v128
                                                    _ = v138
                                                }
                                            }
                                            _ = v66
                                        } else {
                                            if false {
                                                var v139 int16 = (^(v18))
                                                switch "foo bar" {
                                                case "":
                                                    var v140 int = (^(fi23))
                                                    _ = v140
                                                case "foo bar":
                                                    v139 -= v139
                                                }
                                                _ = v139
                                            } else {
                                                v65 = (^(v65))
                                                v141 := v3[-389]
                                                _ = v141
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v65
                                    }
                                case "foo bar":
                                    var v142 uint8 = 127
                                    _ = v142
                                }
                            }
                            var v143 bool = true
                            _ = fi26
                            _ = v143
                        }
                    }
                }
                delete(v3, (-(fi23)))
            }
            switch fi23 {
            case 890:
                v144 := v3[fi23]
                _ = v144
            case -9223372036854775808:
                v4 = append(v4, 65535)
            case -1:
                var v145 bool = (false || (2147483648 >= 65535))
                _ = v145
            default:
                v3[(fi23 - (fi23 | (fi23 * (fi23 << ((1) & 63)))))] = "foo bar"
            }
        } else {
            var v146 int32 = -1
            v18 -= v18
            v147 := GenS10{F0: 42, F1: -55}
            _ = v146
            _ = v147
        }
        _ = len(v3)
        delete(v3, fi23)
        _ = fi23
    }
    if true {
        var v148 uint16 = 255
        if (true && true) {
            v149 := GenS16{F0: -51, F1: 619, F2: 128}
            fmt.Println(v22)
            var v150 uint16 = v148
            _ = v149
            _ = v150
        }
        var v151 float32 = ((1e10 + 1e10) + ((-2.25 + -2.25) / float32((3.14159 - 1.5))))
        _ = v148
        _ = v151
    }
    switch 128 {
    case 128:
        var v152 float32 = ((-2.25 - -2.25) * -2.25)
        _ = v152
    case -2147483648:
        fmt.Println(v22, v22)
    case -851:
        v18 = 970
    default:
        _ = len(v4)
    }
    if true {
        var v153 uint = 183
        v18 = (v18 >> ((v153) & 63))
        _ = len(v1)
        _ = v153
    }
    var v154 uint64 = 263
    v18 = (v18 - v18)
    switch -174 {
    case -9223372036854775808:
        switch 2147483647 {
        case -2147483648:
            for fi155 := 0; fi155 < 17; fi155++ {
                if ((65535 > 2) || ((3.14159 - 3.14159) <= (-2.25 / 0.1))) {
                    v156 := v3[(-(fi155))]
                    var v157 uint16 = 978
                    fmt.Println(v18, v154)
                    _ = v156
                    _ = v157
                }
                _ = fi155
            }
        default:
            fmt.Println(v154)
        }
    case 32767:
        switch 390 {
        case 32767:
            switch "a" {
            case "":
                v158 := GenS14{F0: 2}
                _ = v158
            case "hello":
                if false {
                    v159 := GenS17{F0: true}
                    var v160 string = ""
                    _ = v159
                    _ = v160
                } else {
                    if (!((57 <= 101) || true)) {
                        for fi161 := 0; fi161 < 19; fi161++ {
                            for fi162 := 0; fi162 < 5; fi162++ {
                                v163 := GenS18{F0: 0}
                                _ = fi162
                                _ = v163
                            }
                            v154 = 18446744073709551615
                            v3[(^(fi161))] = ("\n\t" + "hello")
                            _ = fi161
                        }
                        v22 = v22
                        if len(v4) > 0 {
                            v4[0] = 338
                        }
                    }
                }
            case "\n\t":
                for fi164 := 0; fi164 < 15; fi164++ {
                    var v165 float64 = 1e10
                    if true {
                        var v166 uint32 = 2147483648
                        _ = v166
                    }
                    _ = len(v1)
                    _ = fi164
                    _ = v165
                }
            default:
                if ("foo bar" != "") {
                    v167 := GenS10{F0: -128, F1: -38}
                    if (!true) {
                        v5 = append(v5, -1)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        var v168 int = 4294967295
                        _ = v168
                    }
                    _ = v167
                } else {
                    v18 = (v18 % (((v18 % (((v18 >> ((4294967295) & 63))) | 1))) | 1))
                }
            }
        }
    default:
        if false {
            var v169 float64 = 1.5
            if true {
                var v170 uint = 2
                v171 := GenS9{F0: 42, F1: -284}
                delete(v3, -90)
                _ = v170
                _ = v171
            }
            _ = v169
        } else {
            var v172 uint8 = 42
            var v173 int32 = -864
            v154 = 1
            _ = v172
            _ = v173
        }
    }
    var v174 uint32 = 2
    for fi175 := 0; fi175 < 19; fi175++ {
        _ = len(v1)
        var v176 bool = (!(!(255 != 42)))
        switch "" {
        case "foo bar":
            v177 := GenS19{F0: v154, F1: 2}
            _ = v177
        case "":
            fmt.Println(v154, v22)
        }
        _ = fi175
        _ = v176
    }
    for fi178 := 0; fi178 < 19; fi178++ {
        var v179 int64 = (v22 % ((v22) | 1))
        _ = fi178
        _ = v179
    }
    v180 := v3[-624]
    v181 := GenS2{F0: (v180 + (v180 + ("foo bar" + (v180 + ("foo bar" + (("a" + (v180 + ((v180 + v180) + (v180 + "\n\t")))) + v180)))))), F1: 1000, F2: -2147483648}
    if (229 <= uint(v154)) {
        if ((^(v174)) >= v174) {
            var v182 uint = 4294967295
            _ = v182
        }
        if (((1e10 - (float64(1.5) + ((float64(0.0) * 3.14159) + 0.0))) + -2.25) >= -2.25) {
            var v183 uint = 201
            var v184 uint16 = 1
            if (v183 <= (^(v183))) {
                switch 367 {
                case 127:
                    fmt.Println(v183)
                case -412:
                    var v185 uint8 = 245
                    _ = v185
                default:
                    if true {
                        for fi186 := 0; fi186 < 1; fi186++ {
                            for fi187 := 0; fi187 < 1; fi187++ {
                                v184 -= v184
                                _ = fi187
                            }
                            v183 -= (v183 % ((v183) | 1))
                            if true {
                                switch v180 {
                                case "\n\t":
                                    var v188 uint8 = 127
                                    _ = v188
                                case "":
                                    switch fi186 {
                                    case 995:
                                        delete(v3, (fi186 << (((v183 + v183)) & 63)))
                                    case 42:
                                        v189 := v3[(-(fi186))]
                                        _ = v189
                                    case 127:
                                        var v190 int32 = 1
                                        _ = v190
                                    default:
                                        v18 += 1
                                    }
                                default:
                                    var v191 uint32 = (^(v174))
                                    _ = v191
                                }
                            }
                            _ = fi186
                        }
                        switch 2147483647 {
                        case 32768:
                            v183 = (^(v183))
                        case 563:
                            var v192 string = v180
                            _ = v192
                        case -2147483648:
                            var v193 uint8 = 90
                            _ = v193
                        default:
                            if (!false) {
                                v5 = append(v5, 127)
                                if true {
                                    v174 = (v174 * 1)
                                }
                            }
                        }
                        if false {
                            v154 = v154
                        }
                    } else {
                        v194 := GenS3{F0: false, F1: (^(v183))}
                        v174 = 4294967295
                        _ = v194
                    }
                }
            }
            _ = v183
            _ = v184
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v154:", v154)
    fmt.Println("v174:", v174)
    fmt.Println("v18:", v18)
    fmt.Println("v180:", v180)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v154
    _ = v174
    _ = v18
    _ = v180
    _ = v181
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
