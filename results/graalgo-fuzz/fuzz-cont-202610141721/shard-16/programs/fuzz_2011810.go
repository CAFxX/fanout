package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 uint; F2 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint8; F1 uint64; F2 bool }
type GenS17 struct { F0 int; F1 string; F2 int8 }
type GenS18 struct { F0 uint; F1 int8; F2 uint }
type GenS19 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS20 struct { F0 uint16; F1 int32; F2 int64 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 string }
type GenS24 struct { F0 int; F1 int16 }
type GenS25 struct { F0 uint16; F1 uint64 }
type GenS26 struct { F0 int16; F1 uint16; F2 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]bool)
    var v3 int64 = -173
    v4 := []string{"hello"}
    v5 := []uint32{v1}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: 112, F1: 1}
            _ = fi7
            _ = v8
        }
        var v9 uint16 = 65535
        if (65535 != 65535) {
            var v10 int64 = 9223372036854775807
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint8 = 2
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        _ = fi6
        _ = v9
    }
    fmt.Println(v1)
    v1 = v1
    var v13 float64 = 0.1
    if false {
        switch 664 {
        case 32768:
            switch 4294967295 {
            case 32767:
                for fi14 := 0; fi14 < 12; fi14++ {
                    if false {
                        var v15 uint64 = 18446744073709551615
                        _ = v15
                    } else {
                        v16 := GenS2{F0: v3, F1: 18446744073709551615, F2: v3}
                        if false {
                            if ((!false) && false) {
                                var v17 float32 = 3.14159
                                _ = v17
                            }
                            v18 := GenS1{F0: 93, F1: 42}
                            _ = v18
                        }
                        _ = v16
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        fmt.Println(v13)
                        _ = fi19
                    }
                    var v20 string = "a"
                    _ = fi14
                    _ = v20
                }
            case 2147483647:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case -2147483648:
            var v21 uint64 = 555
            _ = v21
        default:
            for fi22 := 0; fi22 < 3; fi22++ {
                if (false && true) {
                    fmt.Println(v13)
                    _ = len(v5)
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    v3 = 0
                    v1 += v1
                    v3 -= v3
                    _ = fi23
                }
                _ = fi22
            }
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: 42, F1: -2147483648}
            v13 += v13
            _ = fi24
            _ = v25
        }
    }
    v1 -= 395
    for fi26 := 0; fi26 < 12; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            fmt.Println(v3, v3)
            _ = fi27
        }
        var v28 uint = 1
        _ = fi26
        _ = v28
    }
    v13 += (v13 * -2.25)
    var v29 int16 = 32767
    v30 := GenS3{F0: 1}
    for fi31 := 0; fi31 < 2; fi31++ {
        v13 += v13
        _ = fi31
    }
    v32 := v2["\n\t"]
    if false {
        fmt.Println(v29, v3, v29)
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS1{F0: 65535, F1: -291}
            if v32 {
                if v32 {
                    var v35 int16 = 1000
                    v36 := GenS4{F0: 42, F1: 1}
                    if (!(!true)) {
                        var v37 uint = 18446744073709551615
                        _ = v37
                    } else {
                        if v32 {
                            v38 := GenS5{F0: 472, F1: -1, F2: (v1 & v1)}
                            for fi39 := 0; fi39 < 1; fi39++ {
                                if ((0 <= 1) || (!("hello" != "foo bar"))) {
                                    v35 = 42
                                } else {
                                    v40 := GenS6{F0: (v35 % ((v35) | 1))}
                                    v41 := GenS7{F0: 1, F1: 42}
                                    switch fi33 {
                                    case 9223372036854775807:
                                        if ((2 > 296) || v32) {
                                            fmt.Println(v35, v3, fi39)
                                            if (true && v32) {
                                                v42 := GenS2{F0: 9223372036854775807, F1: 2, F2: (v3 ^ (v3 - (^(v3))))}
                                                switch (-(fi39)) {
                                                case -849:
                                                    fmt.Println(v13)
                                                case -2147483648:
                                                    var v43 string = ("hello" + ("foo bar" + "hello"))
                                                    _ = v43
                                                case 9223372036854775807:
                                                    v29 -= 32767
                                                }
                                                _ = v42
                                            } else {
                                                v44 := v2["\n\t"]
                                                switch "a" {
                                                case "a":
                                                    if true {
                                                        v5 = append(v5, v1)
                                                        v45 := GenS2{F0: (v3 & (v3 - (v3 | (^(v3))))), F1: 54, F2: -1}
                                                        _ = v45
                                                    } else {
                                                        var v46 int16 = v29
                                                        _ = v46
                                                    }
                                                case "\n\t":
                                                    _ = len(v4)
                                                default:
                                                    v1 = (v1 % ((v1) | 1))
                                                }
                                                if ((fi39 + 42) > 4294967295) {
                                                    switch ("foo bar" + ((("" + "") + "") + "")) {
                                                    case "\n\t":
                                                        v47 := GenS2{F0: v3, F1: 0, F2: 9223372036854775807}
                                                        _ = v47
                                                    default:
                                                        v35 = v29
                                                    }
                                                    var v48 int = -9223372036854775808
                                                    _ = v48
                                                } else {
                                                    if v32 {
                                                        v32 = false
                                                        if (2147483647 > 2147483647) {
                                                            fmt.Println(v3)
                                                            v29 = (v35 - int16(v1))
                                                        } else {
                                                            if false {
                                                                _ = len(v5)
                                                            }
                                                            v44 = v44
                                                            var v49 int8 = 0
                                                            _ = v49
                                                        }
                                                        if (0.1 < 3.14159) {
                                                            var v50 uint32 = 65535
                                                            if v32 {
                                                                v51 := GenS4{F0: 2, F1: 0}
                                                                if (255 < 803) {
                                                                    if (!(v32 && false)) {
                                                                        v52 := GenS7{F0: 0, F1: 127}
                                                                        if (1.5 == v13) {
                                                                            if v44 {
                                                                                fmt.Println(v44)
                                                                                v53 := GenS8{F0: (!true)}
                                                                                fmt.Println(v13, v3, v35)
                                                                                _ = v53
                                                                            } else {
                                                                                var v54 uint8 = 255
                                                                                _ = v54
                                                                            }
                                                                            fmt.Println(v29, v3, v29)
                                                                        } else {
                                                                            switch -1 {
                                                                            case -2147483648:
                                                                                v13 -= 0.0
                                                                            case 71:
                                                                                var v55 bool = true
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                        _ = v52
                                                                    }
                                                                } else {
                                                                    var v56 bool = (!((!v32) || v32))
                                                                    var v57 uint16 = 2
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                _ = v51
                                                            } else {
                                                                v58 := GenS4{F0: 42, F1: 1}
                                                                var v59 uint8 = 173
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            v4 = append(v4, "hello")
                                                            _ = v50
                                                        } else {
                                                            if (float32(v13) == float32(-2.25)) {
                                                                if (-2.25 <= -2.25) {
                                                                    fmt.Println(fi39, fi39)
                                                                    fmt.Println(v13)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v1 >> ((359) & 63))
                                                                    }
                                                                    if (1000 == (v35 / ((v35) | 1))) {
                                                                        v60 := GenS2{F0: v3, F1: 4294967295, F2: v3}
                                                                        _ = v60
                                                                    } else {
                                                                        v44 = true
                                                                        if v32 {
                                                                            fmt.Println(fi39, fi33)
                                                                            var v61 uint = 4294967295
                                                                            if true {
                                                                                fmt.Println(v29, v29, v3)
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    _ = len(v2)
                                                                                case "hello":
                                                                                    if ((v61 - 668) >= (v61 - uint(v1))) {
                                                                                        v13 = v13
                                                                                    } else {
                                                                                        v32 = (!(42 < (v29 * (^(v29)))))
                                                                                        if ((v1 << ((v61) & 63)) < v1) {
                                                                                            v61 -= v61
                                                                                            _ = len(v2)
                                                                                            var v62 float64 = 1e10
                                                                                            _ = v62
                                                                                        } else {
                                                                                            if v44 {
                                                                                                v63 := GenS4{F0: 0, F1: -128}
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    v35 = int16(fi33)
                                                                                                default:
                                                                                                    if v32 {
                                                                                                        switch ("a" + ("a" + "a")) {
                                                                                                        case "a":
                                                                                                            _ = len(v4)
                                                                                                        case "hello":
                                                                                                            switch (("a" + "") + "hello") {
                                                                                                            case "hello":
                                                                                                                var v64 int = 42
                                                                                                                _ = v64
                                                                                                            case "a":
                                                                                                                if ((-(v3)) != 9223372036854775807) {
                                                                                                                    v13 -= ((v13 * -2.25) * -2.25)
                                                                                                                    v61 = v61
                                                                                                                } else {
                                                                                                                    v61 += 255
                                                                                                                    var v65 uint32 = (^(v1))
                                                                                                                    _ = v65
                                                                                                                }
                                                                                                            default:
                                                                                                                fmt.Println(v61, fi39, v13)
                                                                                                                v66 := GenS9{F0: "hello"}
                                                                                                                v67 := GenS9{F0: ("" + "\n\t")}
                                                                                                                _ = v66
                                                                                                                _ = v67
                                                                                                            }
                                                                                                            v68 := GenS10{F0: 255, F1: v44, F2: (v1 & v1)}
                                                                                                            var v69 uint64 = uint64(v1)
                                                                                                            var v70 uint8 = 172
                                                                                                            var v71 bool = ((1e10 * (v13 + 0.1)) != v13)
                                                                                                            v71 = (!false)
                                                                                                            _ = v68
                                                                                                            _ = v69
                                                                                                            _ = v70
                                                                                                            _ = v71
                                                                                                        default:
                                                                                                            v72 := GenS1{F0: 111, F1: 1}
                                                                                                            _ = v72
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v73 := GenS11{F0: v1, F1: 127}
                                                                                                _ = v63
                                                                                                _ = v73
                                                                                            }
                                                                                            if (v61 == 255) {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    switch "" {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v13, fi33)
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            var v74 int8 = 100
                                                                                                            v75 := GenS2{F0: -9223372036854775808, F1: 332, F2: v3}
                                                                                                            _ = v74
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    case "":
                                                                                                        v13 = 0.1
                                                                                                    }
                                                                                                    if (((v13 * (-2.25 * 3.14159)) - (0.0 - float64(1e10))) == v13) {
                                                                                                        v76 := GenS5{F0: 42, F1: 1, F2: v1}
                                                                                                        if ((v61 * (v61 & v61)) > (v61 >> ((62) & 63))) {
                                                                                                            fmt.Println(v32, fi39)
                                                                                                            if ("foo bar" < "") {
                                                                                                                if len(v4) >= 2 {
                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                }
                                                                                                                v77 := v2["\n\t"]
                                                                                                                _ = v77
                                                                                                            }
                                                                                                            _ = len(v2)
                                                                                                        } else {
                                                                                                            var v78 float64 = float64((((1.5 + (-2.25 - float32(-2.25))) * -2.25) / (1e10 - 3.14159)))
                                                                                                            switch "" {
                                                                                                            case "a":
                                                                                                                v79 := GenS8{F0: (!(1 >= 833))}
                                                                                                                var v80 uint32 = 4294967295
                                                                                                                fmt.Println(v1, v1)
                                                                                                                var v81 float64 = ((0.1 + (v13 + v78)) + (float64(((1e10 + (-2.25 * (((0.1 / 1e10) + 0.0) * -2.25))) - 1e10)) + 0.0))
                                                                                                                switch (fi39 * -396) {
                                                                                                                case 32768:
                                                                                                                    var v82 uint64 = 62
                                                                                                                    _ = v82
                                                                                                                }
                                                                                                                if v44 {
                                                                                                                    if len(v4) > 0 {
                                                                                                                        v4[0] = "a"
                                                                                                                    }
                                                                                                                    v5 = append(v5, (v80 % (((^(v1))) | 1)))
                                                                                                                    var v83 int32 = 2147483647
                                                                                                                    _ = v83
                                                                                                                } else {
                                                                                                                    v1 = 2
                                                                                                                }
                                                                                                                _ = v79
                                                                                                                _ = v80
                                                                                                                _ = v81
                                                                                                            case "foo bar":
                                                                                                                v61 = v61
                                                                                                            }
                                                                                                            _ = v78
                                                                                                        }
                                                                                                        v1 += (v1 / (((v1 - (v1 & (v1 << (((v61 << ((790) & 63))) & 63))))) | 1))
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        fmt.Println(v32)
                                                                                                        if (v29 > (^(v29))) {
                                                                                                            switch "foo bar" {
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v29, v32)
                                                                                                            case "a":
                                                                                                                var v84 uint16 = 851
                                                                                                                _ = v84
                                                                                                            case "":
                                                                                                                var v85 int = -2147483648
                                                                                                                var v86 uint64 = 1
                                                                                                                _ = v85
                                                                                                                _ = v86
                                                                                                            default:
                                                                                                                var v87 uint = (v61 & (v61 << ((v61) & 63)))
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        }
                                                                                                        switch (^(fi33)) {
                                                                                                        case -9223372036854775808:
                                                                                                            v13 = v13
                                                                                                        case -955:
                                                                                                            switch (-(fi33)) {
                                                                                                            case 235:
                                                                                                                var v88 int32 = 617
                                                                                                                _ = v88
                                                                                                            default:
                                                                                                                v89 := GenS7{F0: 100, F1: 127}
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            var v90 int = -772
                                                                                                            _ = v90
                                                                                                        default:
                                                                                                            fmt.Println(v61, v1)
                                                                                                        }
                                                                                                        var v91 uint = v61
                                                                                                        _ = v91
                                                                                                    }
                                                                                                case "":
                                                                                                    switch "a" {
                                                                                                    case "\n\t":
                                                                                                        switch ("" + ((("\n\t" + ("hello" + "\n\t")) + "") + ("foo bar" + "foo bar"))) {
                                                                                                        case "foo bar":
                                                                                                            if v44 {
                                                                                                                v92 := GenS5{F0: (v1 | v1), F1: 2147483647, F2: uint32(v61)}
                                                                                                                _ = v92
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v13 = (v13 + (1e10 + 0.0))
                                                                                                    }
                                                                                                case "a":
                                                                                                    var v93 uint = v61
                                                                                                    v1 -= v1
                                                                                                    _ = v93
                                                                                                default:
                                                                                                    if ("\n\t" >= ("" + "\n\t")) {
                                                                                                        var v94 uint8 = 128
                                                                                                        var v95 uint8 = v94
                                                                                                        _ = v94
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v96 := GenS12{F0: (fi39 << ((v61) & 63))}
                                                                                                            _ = v96
                                                                                                        }
                                                                                                        var v97 float64 = (v13 / v13)
                                                                                                        if v44 {
                                                                                                            var v98 uint = v61
                                                                                                            _ = v98
                                                                                                        } else {
                                                                                                            v32 = v44
                                                                                                            var v99 uint16 = uint16(v61)
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        _ = v97
                                                                                                    }
                                                                                                }
                                                                                                v2["\n\t"] = v32
                                                                                                v35 += 0
                                                                                            }
                                                                                            v3 = v3
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v100 uint64 = 0
                                                                                    _ = v100
                                                                                }
                                                                            }
                                                                            var v101 uint8 = 42
                                                                            _ = v101
                                                                            _ = v61
                                                                        }
                                                                        if v32 {
                                                                            switch (("foo bar" + ("foo bar" + ("" + "\n\t"))) + "\n\t") {
                                                                            case "hello":
                                                                                v102 := GenS13{F0: v3, F1: 4294967295, F2: 2}
                                                                                _ = v102
                                                                            case "a":
                                                                                var v103 float64 = (v13 * 0.0)
                                                                                _ = v103
                                                                            case "":
                                                                                v32 = v32
                                                                            default:
                                                                                v29 = 1
                                                                            }
                                                                            v104 := GenS8{F0: false}
                                                                            _ = v104
                                                                        }
                                                                    }
                                                                }
                                                                if (fi39 >= fi39) {
                                                                    if v44 {
                                                                        var v105 float64 = v13
                                                                        _ = v105
                                                                    } else {
                                                                        var v106 int = (fi33 * (fi39 ^ (fi33 + fi33)))
                                                                        _ = v106
                                                                    }
                                                                }
                                                                v44 = false
                                                            } else {
                                                                if v44 {
                                                                    v107 := v2[("foo bar" + "\n\t")]
                                                                    _ = v107
                                                                } else {
                                                                    if (v44 || true) {
                                                                        var v108 int16 = 1000
                                                                        if true {
                                                                            var v109 string = "a"
                                                                            switch fi39 {
                                                                            case -1:
                                                                                if ((0.0 == 0.1) || (v3 >= v3)) {
                                                                                    if (!(v32 || v32)) {
                                                                                        var v110 int32 = 1
                                                                                        var v111 bool = v44
                                                                                        v1 += v1
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                } else {
                                                                                    var v112 uint32 = v1
                                                                                    var v113 int = fi33
                                                                                    fmt.Println(v112, v3)
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                }
                                                                            case -295:
                                                                                if v44 {
                                                                                    var v114 uint16 = 0
                                                                                    v108 -= (v108 - v108)
                                                                                    v44 = true
                                                                                    _ = v114
                                                                                } else {
                                                                                    var v115 string = (("a" + v109) + "")
                                                                                    if (0.1 < v13) {
                                                                                        var v116 uint16 = 0
                                                                                        v32 = ((fi33 / ((fi33) | 1)) != (fi39 % ((1) | 1)))
                                                                                        _ = v116
                                                                                    }
                                                                                    if ((-(fi33)) != int(v1)) {
                                                                                        fmt.Println(v35, v44)
                                                                                        v2[(v109 + (v109 + "a"))] = true
                                                                                    }
                                                                                    _ = v115
                                                                                }
                                                                                if (0 > 671) {
                                                                                    var v117 float64 = v13
                                                                                    var v118 float32 = 0.1
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                } else {
                                                                                    switch fi39 {
                                                                                    case -1:
                                                                                        if (v109 >= "a") {
                                                                                            var v119 float64 = -2.25
                                                                                            var v120 string = ((v109 + "foo bar") + ((v109 + ("" + (("\n\t" + v109) + v109))) + v109))
                                                                                            var v121 bool = false
                                                                                            _ = v119
                                                                                            _ = v120
                                                                                            _ = v121
                                                                                        } else {
                                                                                            v13 = v13
                                                                                            fmt.Println(v1)
                                                                                            var v122 int32 = 0
                                                                                            _ = v122
                                                                                        }
                                                                                    case 42:
                                                                                        if (v1 <= v1) {
                                                                                            var v123 int8 = 100
                                                                                            if true {
                                                                                                if (!(("\n\t" + (v109 + v109)) < v109)) {
                                                                                                    var v124 uint8 = 127
                                                                                                    var v125 uint = 1
                                                                                                    v126 := GenS14{F0: v109}
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                    _ = v126
                                                                                                }
                                                                                                v1 += (v1 + 0)
                                                                                            } else {
                                                                                                delete(v2, "foo bar")
                                                                                                var v127 int64 = -9223372036854775808
                                                                                                _ = v127
                                                                                            }
                                                                                            _ = v123
                                                                                        } else {
                                                                                            v35 = 0
                                                                                            v32 = v32
                                                                                            _ = len(v4)
                                                                                        }
                                                                                    default:
                                                                                        switch fi33 {
                                                                                        case 32768:
                                                                                            if v32 {
                                                                                                switch v109 {
                                                                                                case "a":
                                                                                                    if (((1.5 + 0.1) - 0.1) <= (float32(0.1) - (0.1 * 1.5))) {
                                                                                                        if v44 {
                                                                                                            v13 += v13
                                                                                                            fmt.Println(v108)
                                                                                                        }
                                                                                                        switch v109 {
                                                                                                        case "a":
                                                                                                            v128 := GenS2{F0: v3, F1: 4294967295, F2: (v3 >> ((18446744073709551615) & 63))}
                                                                                                            _ = v128
                                                                                                        case "\n\t":
                                                                                                            v3 += 42
                                                                                                        default:
                                                                                                            var v129 uint64 = 1
                                                                                                            _ = v129
                                                                                                        }
                                                                                                        var v130 uint16 = 65535
                                                                                                        _ = v130
                                                                                                    } else {
                                                                                                        switch "foo bar" {
                                                                                                        case "":
                                                                                                            if true {
                                                                                                                delete(v2, (v109 + "\n\t"))
                                                                                                                if (42 != 1) {
                                                                                                                    fmt.Println(fi33, v44)
                                                                                                                    fmt.Println(v108, v108, v13)
                                                                                                                    if true {
                                                                                                                        v131 := GenS9{F0: ("hello" + ("hello" + v109))}
                                                                                                                        _ = v131
                                                                                                                    } else {
                                                                                                                        var v132 int16 = (v35 | v29)
                                                                                                                        _ = v132
                                                                                                                    }
                                                                                                                    var v133 int8 = 1
                                                                                                                    _ = v133
                                                                                                                } else {
                                                                                                                    v134 := GenS15{F0: 42}
                                                                                                                    if v32 {
                                                                                                                        v29 = (v108 % ((1) | 1))
                                                                                                                        switch fi39 {
                                                                                                                        case 886:
                                                                                                                            if (v32 || true) {
                                                                                                                                if v32 {
                                                                                                                                    v35 = (v35 % ((32767) | 1))
                                                                                                                                    v35 -= -32768
                                                                                                                                    switch v109 {
                                                                                                                                    case "":
                                                                                                                                        v29 = (v29 + v35)
                                                                                                                                    case "a":
                                                                                                                                        if (1 <= 668) {
                                                                                                                                            fmt.Println(v44)
                                                                                                                                            v108 -= (-(v35))
                                                                                                                                            var v135 uint64 = 260
                                                                                                                                            _ = v135
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v3)
                                                                                                                                            var v136 int8 = -1
                                                                                                                                            switch fi39 {
                                                                                                                                            case -267:
                                                                                                                                                if (v44 && false) {
                                                                                                                                                    fmt.Println(v32, v29, v44)
                                                                                                                                                    if (863 > 257) {
                                                                                                                                                        if (!(643 <= (fi39 >> ((1) & 63)))) {
                                                                                                                                                            var v137 int8 = 100
                                                                                                                                                            v138 := v2[(v109 + (v109 + v109))]
                                                                                                                                                            _ = v137
                                                                                                                                                            _ = v138
                                                                                                                                                        }
                                                                                                                                                        v139 := GenS4{F0: 42, F1: v136}
                                                                                                                                                        if (fi33 < fi39) {
                                                                                                                                                            var v140 float64 = v13
                                                                                                                                                            var v141 float64 = (-2.25 * (v13 * v140))
                                                                                                                                                            v2[v109] = (float64(-2.25) == ((v13 / 1.0) - v141))
                                                                                                                                                            switch v109 {
                                                                                                                                                            case "":
                                                                                                                                                                switch (fi39 / ((32767) | 1)) {
                                                                                                                                                                case 4294967295:
                                                                                                                                                                    var v142 int = fi33
                                                                                                                                                                    _ = v142
                                                                                                                                                                case 32768:
                                                                                                                                                                    if (v13 <= float64(1e10)) {
                                                                                                                                                                        if (!(v44 || (!v44))) {
                                                                                                                                                                            fmt.Println(v3, v136, v32)
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        if (-738 > 945) {
                                                                                                                                                                            var v143 int64 = (v3 & (v3 * 488))
                                                                                                                                                                            if len(v4) > 0 {
                                                                                                                                                                                v4[0] = "a"
                                                                                                                                                                            }
                                                                                                                                                                            v144 := GenS1{F0: 562, F1: -2147483648}
                                                                                                                                                                            _ = v143
                                                                                                                                                                            _ = v144
                                                                                                                                                                        } else {
                                                                                                                                                                            var v145 uint32 = v1
                                                                                                                                                                            if (v3 == (v3 ^ v3)) {
                                                                                                                                                                                var v146 int64 = (^(v3))
                                                                                                                                                                                if v32 {
                                                                                                                                                                                    fmt.Println(v108, v32)
                                                                                                                                                                                    v136 -= (v136 | v136)
                                                                                                                                                                                } else {
                                                                                                                                                                                    var v147 bool = false
                                                                                                                                                                                    var v148 int8 = v136
                                                                                                                                                                                    _ = v147
                                                                                                                                                                                    _ = v148
                                                                                                                                                                                }
                                                                                                                                                                                _ = v146
                                                                                                                                                                            }
                                                                                                                                                                            v149 := GenS16{F0: 0, F1: 42, F2: (!(v1 <= (v145 >> ((943) & 63))))}
                                                                                                                                                                            if (v32 && false) {
                                                                                                                                                                                v150 := GenS6{F0: -32768}
                                                                                                                                                                                v151 := GenS9{F0: ""}
                                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                                }
                                                                                                                                                                                _ = v150
                                                                                                                                                                                _ = v151
                                                                                                                                                                            } else {
                                                                                                                                                                                v32 = (true && v44)
                                                                                                                                                                            }
                                                                                                                                                                            _ = v145
                                                                                                                                                                            _ = v149
                                                                                                                                                                        }
                                                                                                                                                                        fmt.Println(v29, v13)
                                                                                                                                                                        if v32 {
                                                                                                                                                                            var v152 int16 = (v35 | -32768)
                                                                                                                                                                            _ = v152
                                                                                                                                                                        } else {
                                                                                                                                                                            v109 = ("hello" + (v109 + "foo bar"))
                                                                                                                                                                        }
                                                                                                                                                                        v108 = v35
                                                                                                                                                                        v153 := GenS9{F0: v109}
                                                                                                                                                                        var v154 uint = 0
                                                                                                                                                                        if (1 < int8(v154)) {
                                                                                                                                                                            v155 := GenS14{F0: v109}
                                                                                                                                                                            if v32 {
                                                                                                                                                                                fmt.Println(v32)
                                                                                                                                                                                var v156 int8 = (v136 | v136)
                                                                                                                                                                                v5 = append(v5, 4294967295)
                                                                                                                                                                                _ = v156
                                                                                                                                                                            } else {
                                                                                                                                                                                switch v109 {
                                                                                                                                                                                case "hello":
                                                                                                                                                                                    v44 = false
                                                                                                                                                                                }
                                                                                                                                                                                var v157 int16 = v108
                                                                                                                                                                                _ = v157
                                                                                                                                                                            }
                                                                                                                                                                            v13 = (v141 + v141)
                                                                                                                                                                            _ = v155
                                                                                                                                                                        }
                                                                                                                                                                        _ = v153
                                                                                                                                                                        _ = v154
                                                                                                                                                                    }
                                                                                                                                                                case 807:
                                                                                                                                                                    var v158 uint32 = (v1 ^ (v1 / (((v1 | v1)) | 1)))
                                                                                                                                                                    _ = v158
                                                                                                                                                                default:
                                                                                                                                                                    if (true && false) {
                                                                                                                                                                        fmt.Println(v35)
                                                                                                                                                                        var v159 uint64 = 2
                                                                                                                                                                        v160 := v2["a"]
                                                                                                                                                                        _ = v159
                                                                                                                                                                        _ = v160
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                var v161 float64 = v13
                                                                                                                                                                _ = v161
                                                                                                                                                            }
                                                                                                                                                            var v162 int32 = 379
                                                                                                                                                            _ = v140
                                                                                                                                                            _ = v141
                                                                                                                                                            _ = v162
                                                                                                                                                        } else {
                                                                                                                                                            switch (v109 + (v109 + v109)) {
                                                                                                                                                            case "\n\t":
                                                                                                                                                                v2[v109] = v32
                                                                                                                                                            case "":
                                                                                                                                                                var v163 int64 = (-(v3))
                                                                                                                                                                _ = v163
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v139
                                                                                                                                                    }
                                                                                                                                                    if (113 > 2) {
                                                                                                                                                        v1 = (v1 % (((v1 & (v1 | v1))) | 1))
                                                                                                                                                        v164 := GenS17{F0: 4294967295, F1: ((v109 + ((v109 + ("" + (((v109 + (v109 + v109)) + v109) + v109))) + "a")) + v109), F2: 100}
                                                                                                                                                        if v44 {
                                                                                                                                                            v3 = (v3 / ((v3) | 1))
                                                                                                                                                            switch v109 {
                                                                                                                                                            case "hello":
                                                                                                                                                                if (252 != 4294967295) {
                                                                                                                                                                    var v165 uint8 = 255
                                                                                                                                                                    fmt.Println(fi39, v109)
                                                                                                                                                                    if v44 {
                                                                                                                                                                        if (v32 || ((v1 % ((258) | 1)) != v1)) {
                                                                                                                                                                            fmt.Println(v165, v35, fi39)
                                                                                                                                                                            v13 = 0.0
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        v109 = v109
                                                                                                                                                                        var v166 int16 = -1
                                                                                                                                                                        _ = v166
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v136, v29, v35)
                                                                                                                                                                    _ = v165
                                                                                                                                                                }
                                                                                                                                                            case "foo bar":
                                                                                                                                                                delete(v2, "\n\t")
                                                                                                                                                            case "\n\t":
                                                                                                                                                                if len(v4) > 0 {
                                                                                                                                                                    v4[0] = v109
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                if len(v5) > 0 {
                                                                                                                                                                    v5[0] = 2147483648
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v167 float64 = (v13 + (v13 / 1.0))
                                                                                                                                                            _ = v167
                                                                                                                                                        }
                                                                                                                                                        var v168 float64 = (((float64(0.0) - 1e10) - -2.25) - (((v13 / v13) - (0.0 - 0.0)) - (v13 / v13)))
                                                                                                                                                        v1 = (^(v1))
                                                                                                                                                        _ = v164
                                                                                                                                                        _ = v168
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    v169 := GenS2{F0: (v3 | (v3 + v3)), F1: 18446744073709551615, F2: v3}
                                                                                                                                                    v1 = (v1 / (((v1 - v1)) | 1))
                                                                                                                                                    _ = v169
                                                                                                                                                }
                                                                                                                                            case 4294967295:
                                                                                                                                                v29 -= (-(v29))
                                                                                                                                            case -355:
                                                                                                                                                v170 := GenS18{F0: 1, F1: (v136 * (v136 - (-(v136)))), F2: 4294967295}
                                                                                                                                                _ = v170
                                                                                                                                            }
                                                                                                                                            if v32 {
                                                                                                                                                var v171 uint8 = 127
                                                                                                                                                var v172 int16 = 42
                                                                                                                                                v32 = (!true)
                                                                                                                                                _ = v171
                                                                                                                                                _ = v172
                                                                                                                                            } else {
                                                                                                                                                v2["hello"] = v44
                                                                                                                                            }
                                                                                                                                            _ = v136
                                                                                                                                        }
                                                                                                                                    case "foo bar":
                                                                                                                                        switch -9223372036854775808 {
                                                                                                                                        case 127:
                                                                                                                                            v2[v109] = (v109 >= v109)
                                                                                                                                        case 875:
                                                                                                                                            v173 := GenS18{F0: 65535, F1: 100, F2: 255}
                                                                                                                                            _ = v173
                                                                                                                                        case 498:
                                                                                                                                            if (false || false) {
                                                                                                                                                v32 = (0.0 != -2.25)
                                                                                                                                                if (1 < 18446744073709551615) {
                                                                                                                                                    fmt.Println(v3, fi39)
                                                                                                                                                    var v174 uint32 = 65535
                                                                                                                                                    _ = v174
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                v32 = v32
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            var v175 uint8 = 128
                                                                                                                                            _ = v175
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v176 int = int(v3)
                                                                                                                                _ = v176
                                                                                                                            }
                                                                                                                        }
                                                                                                                        fmt.Println(v109, v35)
                                                                                                                    }
                                                                                                                    _ = v134
                                                                                                                }
                                                                                                                v32 = (v109 <= v109)
                                                                                                                fmt.Println(v13, v35, v109)
                                                                                                                switch fi39 {
                                                                                                                case 9223372036854775807:
                                                                                                                    v2[v109] = v44
                                                                                                                case 895:
                                                                                                                    _ = len(v2)
                                                                                                                default:
                                                                                                                    if v32 {
                                                                                                                        v35 = (^(v108))
                                                                                                                        switch int(v3) {
                                                                                                                        case 32767:
                                                                                                                            var v177 uint8 = 1
                                                                                                                            if v32 {
                                                                                                                                switch 128 {
                                                                                                                                case 127:
                                                                                                                                    v109 = (v109 + v109)
                                                                                                                                case 2147483647:
                                                                                                                                    v178 := GenS3{F0: -32768}
                                                                                                                                    _ = v178
                                                                                                                                case -9223372036854775808:
                                                                                                                                    v179 := GenS4{F0: 927, F1: 127}
                                                                                                                                    _ = v179
                                                                                                                                }
                                                                                                                                fmt.Println(fi39, v177, v35)
                                                                                                                            } else {
                                                                                                                                var v180 int = (fi33 << ((18446744073709551615) & 63))
                                                                                                                                _ = v180
                                                                                                                            }
                                                                                                                            v181 := GenS5{F0: (v1 % (((v1 ^ (v1 | v1))) | 1)), F1: -510, F2: v1}
                                                                                                                            fmt.Println(v32, v3, v35)
                                                                                                                            switch fi33 {
                                                                                                                            case 83:
                                                                                                                                fmt.Println(v109, v108, fi39)
                                                                                                                            default:
                                                                                                                                v4 = append(v4, v109)
                                                                                                                                switch -9223372036854775808 {
                                                                                                                                case 2147483647:
                                                                                                                                    var v182 int64 = int64(fi33)
                                                                                                                                    _ = v182
                                                                                                                                case 32768:
                                                                                                                                    fmt.Println(v177, v44)
                                                                                                                                    if false {
                                                                                                                                        fmt.Println(fi33, v109, v44)
                                                                                                                                    } else {
                                                                                                                                        if len(v4) > 0 {
                                                                                                                                            v4[0] = v109
                                                                                                                                        }
                                                                                                                                        v3 = v3
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    switch (^(fi39)) {
                                                                                                                                    case 32767:
                                                                                                                                        fmt.Println(v3, fi33, v108)
                                                                                                                                    case 1:
                                                                                                                                        var v183 uint16 = 255
                                                                                                                                        _ = v183
                                                                                                                                    default:
                                                                                                                                        var v184 uint8 = 1
                                                                                                                                        _ = v184
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if len(v5) >= 2 {
                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v177
                                                                                                                            _ = v181
                                                                                                                        default:
                                                                                                                            var v185 uint64 = 42
                                                                                                                            _ = v185
                                                                                                                        }
                                                                                                                        if (true || (3.14159 > 3.14159)) {
                                                                                                                            if len(v4) > 0 {
                                                                                                                                v4[0] = v109
                                                                                                                            }
                                                                                                                            v35 = -184
                                                                                                                        }
                                                                                                                        fmt.Println(v3)
                                                                                                                        var v186 uint8 = uint8(v1)
                                                                                                                        _ = v186
                                                                                                                    }
                                                                                                                }
                                                                                                                var v187 float64 = v13
                                                                                                                _ = v187
                                                                                                            }
                                                                                                        case "a":
                                                                                                            var v188 string = v109
                                                                                                            _ = v188
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v108, v108, v32)
                                                                                                            if (!(65535 <= v1)) {
                                                                                                                v5 = append(v5, v1)
                                                                                                                if (0 == 65535) {
                                                                                                                    fmt.Println(v35, v44, v1)
                                                                                                                    var v189 int16 = (v108 ^ (-(v29)))
                                                                                                                    if (int64(v1) != 9223372036854775807) {
                                                                                                                        var v190 float32 = -2.25
                                                                                                                        v191 := GenS19{F0: v1, F1: 650, F2: 127}
                                                                                                                        v29 += v35
                                                                                                                        _ = v190
                                                                                                                        _ = v191
                                                                                                                    } else {
                                                                                                                        v189 += v108
                                                                                                                        if (!(!(true && v32))) {
                                                                                                                            switch v109 {
                                                                                                                            case "hello":
                                                                                                                                if true {
                                                                                                                                    var v192 uint32 = (v1 - v1)
                                                                                                                                    _ = v192
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v189, v44)
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                var v193 uint16 = 2
                                                                                                                                _ = v193
                                                                                                                            }
                                                                                                                            v35 -= v189
                                                                                                                        } else {
                                                                                                                            var v194 bool = (v1 <= (^(v1)))
                                                                                                                            _ = v194
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v189
                                                                                                                }
                                                                                                                v1 = (v1 << ((2) & 63))
                                                                                                            } else {
                                                                                                                var v195 uint16 = 0
                                                                                                                v196 := GenS5{F0: (v1 / (((v1 & v1)) | 1)), F1: 854, F2: 413}
                                                                                                                fmt.Println(v35, v13, v3)
                                                                                                                _ = v195
                                                                                                                _ = v196
                                                                                                            }
                                                                                                            v13 = -2.25
                                                                                                            var v197 bool = true
                                                                                                            _ = v197
                                                                                                        }
                                                                                                        v109 = v109
                                                                                                        if v32 {
                                                                                                            v44 = v32
                                                                                                        } else {
                                                                                                            v198 := GenS11{F0: v1, F1: 127}
                                                                                                            v199 := GenS19{F0: v1, F1: 467, F2: 42}
                                                                                                            _ = v198
                                                                                                            _ = v199
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(fi33, v44)
                                                                                                if (v1 != v1) {
                                                                                                    var v200 int32 = -525
                                                                                                    v108 += (v29 & (-(v108)))
                                                                                                    _ = v200
                                                                                                } else {
                                                                                                    var v201 int32 = 1
                                                                                                    _ = len(v2)
                                                                                                    _ = v201
                                                                                                }
                                                                                            }
                                                                                        case 0:
                                                                                            _ = len(v2)
                                                                                        case 262:
                                                                                            var v202 uint64 = 4294967295
                                                                                            _ = v202
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v35, v108)
                                                                            }
                                                                            switch 128 {
                                                                            case 2147483647:
                                                                                v203 := GenS2{F0: (v3 & 265), F1: 4294967295, F2: 1}
                                                                                _ = v203
                                                                            case -9223372036854775808:
                                                                                v1 -= (v1 ^ v1)
                                                                            }
                                                                            _ = v109
                                                                        }
                                                                        if true {
                                                                            var v204 float64 = float64(3.14159)
                                                                            _ = v204
                                                                        } else {
                                                                            var v205 uint64 = 42
                                                                            _ = v205
                                                                        }
                                                                        _ = v108
                                                                    } else {
                                                                        v32 = true
                                                                    }
                                                                }
                                                                v206 := GenS19{F0: v1, F1: 32767, F2: 42}
                                                                var v207 int64 = -513
                                                                _ = v206
                                                                _ = v207
                                                            }
                                                        }
                                                    } else {
                                                        v13 += v13
                                                        v208 := GenS20{F0: 255, F1: 2147483647, F2: (v3 >> ((0) & 63))}
                                                        _ = v208
                                                    }
                                                }
                                                if (fi33 < fi39) {
                                                    _ = len(v2)
                                                }
                                                _ = v44
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v209 := v2[("hello" + ("a" + "\n\t"))]
                                            var v210 int16 = v29
                                            _ = v209
                                            _ = v210
                                        }
                                        var v211 uint32 = (v1 - (v1 / ((42) | 1)))
                                        fmt.Println(v211, v29, v211)
                                        _ = v211
                                    case 24:
                                        if v32 {
                                            if false {
                                                if false {
                                                    if v32 {
                                                        v32 = false
                                                        switch int(v1) {
                                                        case 32768:
                                                            var v212 int64 = 0
                                                            var v213 int32 = -2147483648
                                                            switch -794 {
                                                            case 268:
                                                                switch (fi39 & fi33) {
                                                                case 900:
                                                                    var v214 float64 = float64((0.0 - 1e10))
                                                                    _ = v214
                                                                case 9223372036854775807:
                                                                    var v215 float64 = -2.25
                                                                    _ = v215
                                                                case 32768:
                                                                    fmt.Println(v1, v32, v32)
                                                                default:
                                                                    var v216 int = (fi39 >> ((810) & 63))
                                                                    _ = v216
                                                                }
                                                            case 32767:
                                                                v217 := GenS21{F0: -263}
                                                                _ = v217
                                                            default:
                                                                v218 := GenS6{F0: v29}
                                                                _ = v218
                                                            }
                                                            _ = v212
                                                            _ = v213
                                                        case 32767:
                                                            if len(v5) > 0 {
                                                                v5[0] = 2147483647
                                                            }
                                                        default:
                                                            var v219 uint = 255
                                                            _ = v219
                                                        }
                                                    }
                                                    var v220 bool = (42 > 4294967295)
                                                    _ = v220
                                                } else {
                                                    if (!v32) {
                                                        delete(v2, "a")
                                                    }
                                                    v29 -= v35
                                                }
                                            } else {
                                                var v221 uint32 = v1
                                                v13 += (v13 * 1e10)
                                                _ = v221
                                            }
                                            if ((v32 && (v32 && (!(uint8(v1) > 72)))) || v32) {
                                                var v222 int32 = 0
                                                v35 = 1
                                                v1 += (v1 << ((255) & 63))
                                                _ = v222
                                            }
                                            v3 = 863
                                        }
                                    case 0:
                                        _ = len(v2)
                                    default:
                                        v35 -= (v29 << ((666) & 63))
                                    }
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi39
                            }
                            if v32 {
                                for fi223 := 0; fi223 < 13; fi223++ {
                                    v224 := GenS22{F0: (!((("a" + "a") + "hello") != ""))}
                                    fmt.Println(v3)
                                    v225 := v2[""]
                                    _ = fi223
                                    _ = v224
                                    _ = v225
                                }
                            } else {
                                fmt.Println(v29)
                                switch ("" + "foo bar") {
                                case "hello":
                                    var v226 int = 127
                                    _ = v226
                                default:
                                    var v227 uint32 = v1
                                    _ = v227
                                }
                            }
                            _ = v38
                        } else {
                            if (v29 < v29) {
                                v29 = v29
                                v29 += (-(v29))
                            }
                        }
                        for fi228 := 0; fi228 < 16; fi228++ {
                            var v229 int64 = v3
                            _ = fi228
                            _ = v229
                        }
                    }
                    _ = v35
                    _ = v36
                }
                if (1 != (v1 - 678)) {
                    var v230 int16 = v29
                    _ = v230
                }
                switch "a" {
                case "":
                    switch ("hello" + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        switch "" {
                        case "\n\t":
                            v2["hello"] = v32
                        default:
                            v4 = append(v4, "a")
                        }
                    default:
                        if v32 {
                            if v32 {
                                switch "hello" {
                                case "a":
                                    for fi231 := 0; fi231 < 1; fi231++ {
                                        if true {
                                            v232 := GenS23{F0: ("a" + (("\n\t" + "foo bar") + "a"))}
                                            if v32 {
                                                if (true && ("" <= ("" + ("hello" + "hello")))) {
                                                    v1 -= (v1 % (((v1 % (((v1 >> ((2) & 63))) | 1))) | 1))
                                                }
                                                v13 -= v13
                                                var v233 int32 = 2147483647
                                                if (fi231 == int(v1)) {
                                                    if false {
                                                        v32 = v32
                                                        v29 = (v29 ^ (-(v29)))
                                                        var v234 uint8 = 2
                                                        _ = v234
                                                    } else {
                                                        v2[(("a" + ("hello" + "")) + "hello")] = (((((("" + ("hello" + ("" + "a"))) + ("foo bar" + ("\n\t" + "hello"))) + (("hello" + "hello") + ((("\n\t" + "a") + "hello") + ("\n\t" + "\n\t")))) + "foo bar") >= "hello") && false)
                                                    }
                                                } else {
                                                    var v235 uint32 = (v1 % ((v1) | 1))
                                                    v32 = ((v235 << ((uint(v235)) & 63)) < (^(v1)))
                                                    v236 := GenS18{F0: 2, F1: 127, F2: 2}
                                                    _ = v235
                                                    _ = v236
                                                }
                                                if false {
                                                    if v32 {
                                                        v237 := GenS5{F0: v1, F1: (^(v233)), F2: v1}
                                                        v2[(("\n\t" + "foo bar") + "hello")] = false
                                                        if (660 != 2) {
                                                            _ = len(v4)
                                                        } else {
                                                            fmt.Println(fi33, v233)
                                                            v238 := v2["\n\t"]
                                                            if ((fi33 * (fi33 - (fi231 ^ (fi33 | fi33)))) < (-(fi231))) {
                                                                var v239 string = "hello"
                                                                var v240 bool = true
                                                                var v241 int32 = (v233 & 1)
                                                                var v242 uint = 18446744073709551615
                                                                if false {
                                                                    if ((!(v238 || ((-2.25 - 1.5) >= float32(v13)))) && ((v242 >> ((255) & 63)) > (v242 >> ((1) & 63)))) {
                                                                        _ = len(v5)
                                                                        v2[v239] = ((v3 * v3) > (v3 - (v3 ^ v3)))
                                                                        var v243 string = v239
                                                                        _ = v243
                                                                    }
                                                                    if false {
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                _ = v239
                                                                _ = v240
                                                                _ = v241
                                                                _ = v242
                                                            } else {
                                                                v29 += v29
                                                                v238 = (42 < 35)
                                                                v29 = 0
                                                                v3 -= int64(v29)
                                                            }
                                                            _ = v238
                                                        }
                                                        _ = v237
                                                    } else {
                                                        v32 = (v32 || (!v32))
                                                    }
                                                    v13 = ((0.1 - v13) + float64(3.14159))
                                                } else {
                                                    if false {
                                                        switch 238 {
                                                        case -2147483648:
                                                            var v244 uint64 = 0
                                                            _ = v244
                                                        default:
                                                            if v32 {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v245 string = ""
                                                                    _ = v245
                                                                case "hello":
                                                                    v1 = 65535
                                                                default:
                                                                    v5 = append(v5, (v1 - 0))
                                                                }
                                                                if true {
                                                                    var v246 string = ("" + "\n\t")
                                                                    var v247 float32 = 3.14159
                                                                    _ = v246
                                                                    _ = v247
                                                                } else {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v1 - v1)
                                                                    }
                                                                }
                                                            } else {
                                                                v248 := GenS13{F0: (v3 & v3), F1: 797, F2: 128}
                                                                var v249 int8 = -1
                                                                if v32 {
                                                                    v250 := GenS24{F0: -1, F1: (v29 % (((-(v29))) | 1))}
                                                                    v5 = append(v5, v1)
                                                                    var v251 uint8 = 0
                                                                    v252 := GenS25{F0: 2, F1: 18446744073709551615}
                                                                    _ = v250
                                                                    _ = v251
                                                                    _ = v252
                                                                }
                                                                _ = v248
                                                                _ = v249
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                        if (false && v32) {
                                                            var v253 uint32 = v1
                                                            _ = v253
                                                        } else {
                                                            var v254 uint16 = 2
                                                            var v255 int16 = (^(v29))
                                                            v255 += -1
                                                            _ = v254
                                                            _ = v255
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi231, v13)
                                                if v32 {
                                                    var v256 int8 = int8(v1)
                                                    _ = v256
                                                }
                                                var v257 float32 = (3.14159 * 0.1)
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v258 := GenS19{F0: (v1 % ((65535) | 1)), F1: 2, F2: -1}
                                                if v32 {
                                                    v233 = (^(v233))
                                                    fmt.Println(v29, fi33, v1)
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch "foo bar" {
                                                        case "a":
                                                            fmt.Println(fi33, fi231, v257)
                                                        case "foo bar":
                                                            if v32 {
                                                                if v32 {
                                                                    v259 := GenS11{F0: v1, F1: 182}
                                                                    _ = v259
                                                                }
                                                            } else {
                                                                fmt.Println(v13)
                                                            }
                                                        case "":
                                                            var v260 uint64 = 0
                                                            _ = v260
                                                        default:
                                                            v3 -= v3
                                                        }
                                                    }
                                                }
                                                if (!(-1 != v233)) {
                                                    v2["hello"] = v32
                                                    switch (fi33 * fi231) {
                                                    case 32768:
                                                        var v261 float32 = v257
                                                        v3 = (-(v3))
                                                        _ = v261
                                                    case -9223372036854775808:
                                                        var v262 uint32 = 751
                                                        _ = v262
                                                    case 873:
                                                        v263 := GenS19{F0: v1, F1: 32767, F2: 42}
                                                        _ = v263
                                                    }
                                                    switch (fi33 - 310) {
                                                    case -2147483648:
                                                        v257 = 1.5
                                                    case 401:
                                                        var v264 int8 = 64
                                                        _ = v264
                                                    }
                                                }
                                                v265 := GenS26{F0: (-(v29)), F1: 0, F2: 42}
                                                _ = v233
                                                _ = v257
                                                _ = v258
                                                _ = v265
                                            }
                                            _ = v232
                                        }
                                        _ = fi231
                                    }
                                default:
                                    var v266 int64 = v3
                                    _ = v266
                                }
                            }
                            v29 = 32767
                        }
                    }
                }
            } else {
                v29 += (v29 - (-(v29)))
            }
            _ = fi33
            _ = v34
        }
        for fi267 := 0; fi267 < 16; fi267++ {
            var v268 int32 = -370
            _ = fi267
            _ = v268
        }
    }
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
