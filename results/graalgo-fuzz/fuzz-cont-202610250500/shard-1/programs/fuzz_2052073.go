package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint32; F1 int8; F2 int8 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 bool; F1 int8; F2 int64 }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 string; F1 int16; F2 int8 }
type GenS13 struct { F0 int16; F1 uint }
type GenS14 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS15 struct { F0 string; F1 uint32 }
type GenS16 struct { F0 bool; F1 int64; F2 uint8 }
type GenS17 struct { F0 int32; F1 int16; F2 int32 }

func main() {
    v1 := []string{}
    v2 := []string{"", "a", ("hello" + "\n\t"), "a"}
    var v3 bool = ("a" != "a")
    var v4 uint16 = 86
    if v3 {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        var v5 float32 = (-2.25 - 1e10)
        var v6 uint32 = 1
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: 32767, F2: 127}
    v2 = append(v2, ("a" + "foo bar"))
    v8 := GenS1{F0: 470, F1: (v4 >> ((65535) & 63)), F2: 127}
    v4 = (v4 >> ((569) & 63))
    if (172 > 2147483648) {
        for fi9 := 0; fi9 < 2; fi9++ {
            if true {
                v10 := GenS1{F0: 2, F1: (v4 ^ v4), F2: 127}
                _ = v10
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(fi11, v3)
                var v12 uint64 = 866
                _ = fi11
                _ = v12
            }
            var v13 int = 9223372036854775807
            _ = fi9
            _ = v13
        }
    } else {
        v4 += (v4 - (^(v4)))
        switch 42 {
        case 32767:
            if false {
                fmt.Println(v3, v3)
                if len(v2) > 0 {
                    v2[0] = ("" + "\n\t")
                }
            } else {
                var v14 string = ("a" + "\n\t")
                _ = v14
            }
        default:
            fmt.Println(v3, v4)
        }
        fmt.Println(v3)
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float64 = (1e10 / 3.14159)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    switch 211 {
    case 32768:
        if v3 {
            var v18 int64 = -9223372036854775808
            for fi19 := 0; fi19 < 6; fi19++ {
                if v3 {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v3 = v3
                        var v21 float32 = (float32(((float64(((3.14159 + 0.1) + 0.1)) * (0.0 * 0.0)) + ((3.14159 / -2.25) + 1.5))) - 0.0)
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS2{F0: 793, F1: 42, F2: 246}
                        v3 = v3
                        v24 := GenS3{F0: 9223372036854775807, F1: 0, F2: fi19}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    if v3 {
                        v18 = (v18 | (v18 ^ 9223372036854775807))
                    } else {
                        v4 = 65535
                    }
                } else {
                    var v25 float32 = 0.0
                    if (("\n\t" < ("hello" + "hello")) || ((v18 & (v18 | (-(v18)))) < (-(v18)))) {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            if true {
                                var v27 uint16 = (v4 & (v4 % ((42) | 1)))
                                if v3 {
                                    fmt.Println(fi19, fi26, v18)
                                    v3 = (4294967295 != 1)
                                }
                                _ = len(v2)
                                v28 := GenS4{F0: int(v18), F1: 2}
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                _ = v27
                                _ = v28
                            }
                            _ = fi26
                        }
                        for fi29 := 0; fi29 < 20; fi29++ {
                            v30 := GenS1{F0: 4294967295, F1: 181, F2: 255}
                            if true {
                                v4 -= (^(v4))
                            } else {
                                v25 += v25
                            }
                            _ = fi29
                            _ = v30
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 int = (fi31 / (((fi19 ^ (-(fi19)))) | 1))
                            var v33 bool = (737 <= 2)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 10; fi34++ {
                            if v3 {
                                var v35 uint64 = 2
                                v35 = 1
                                fmt.Println(v4, v3)
                                v25 += v25
                                _ = v35
                            } else {
                                var v36 int64 = v18
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if (v3 || true) {
                                    var v37 int16 = 1000
                                    var v38 string = ("hello" + ("" + ("hello" + "hello")))
                                    _ = v37
                                    _ = v38
                                }
                                _ = v36
                            }
                            var v39 uint64 = 552
                            if len(v2) > 0 {
                                v2[0] = (("a" + "foo bar") + "hello")
                            }
                            v3 = v3
                            switch "" {
                            case "foo bar":
                                v40 := GenS4{F0: fi34, F1: uint32(v4)}
                                _ = v40
                            default:
                                v18 += v18
                            }
                            v41 := GenS5{F0: 816, F1: 0}
                            if true {
                                var v42 uint32 = 4294967295
                                fmt.Println(fi19, v42)
                                if (4294967295 < 1) {
                                    if (42 > 42) {
                                        fmt.Println(v25, fi19)
                                        if (("\n\t" + "hello") > "") {
                                            fmt.Println(v4)
                                            var v43 string = "foo bar"
                                            switch ((("\n\t" + ((v43 + v43) + ((v43 + (("foo bar" + ("a" + (v43 + ""))) + v43)) + v43))) + v43) + (v43 + v43)) {
                                            case "foo bar":
                                                var v44 float64 = 0.0
                                                _ = v44
                                            case "\n\t":
                                                v4 = 65535
                                            case "hello":
                                                fmt.Println(v43, v18)
                                            }
                                            _ = v43
                                        }
                                        var v45 int8 = 127
                                        switch "" {
                                        case "foo bar":
                                            if v3 {
                                                fmt.Println(v42, v3, v18)
                                                var v46 uint64 = v39
                                                if v3 {
                                                    var v47 bool = (v3 && ((-2.25 * (0.1 / float64(v25))) == 1.5))
                                                    switch ((("" + "a") + ("a" + "\n\t")) + "a") {
                                                    case "foo bar":
                                                        if (!v47) {
                                                            _ = len(v2)
                                                            _ = len(v2)
                                                        }
                                                    case "":
                                                        v42 -= (^(v42))
                                                    default:
                                                        v48 := GenS2{F0: -812, F1: 0, F2: uint8(v4)}
                                                        _ = v48
                                                    }
                                                    v39 = (v46 & v46)
                                                    _ = v47
                                                } else {
                                                    var v49 uint8 = 127
                                                    _ = v49
                                                }
                                                _ = v46
                                            } else {
                                                var v50 int32 = 2147483647
                                                v42 = (v42 * 42)
                                                v4 += 602
                                                _ = v50
                                            }
                                        }
                                        _ = v45
                                    } else {
                                        switch ("\n\t" + ("" + "foo bar")) {
                                        case "foo bar":
                                            v51 := GenS6{F0: (v42 ^ v42), F1: -51, F2: int8(v42)}
                                            _ = v51
                                        case "\n\t":
                                            if false {
                                                v52 := GenS3{F0: -203, F1: v42, F2: fi19}
                                                if true {
                                                    if false {
                                                        if (!(2 >= 128)) {
                                                            var v53 float32 = (1.5 / v25)
                                                            v54 := GenS3{F0: 42, F1: 42, F2: 32767}
                                                            var v55 int64 = int64(v42)
                                                            _ = v53
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v42 -= (v42 ^ v42)
                                                }
                                                _ = v52
                                            }
                                        }
                                        switch (-(fi19)) {
                                        case -412:
                                            _ = len(v1)
                                            v25 = v25
                                            switch "\n\t" {
                                            case "hello":
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                            default:
                                                if (0 <= -1) {
                                                    fmt.Println(v4, v18)
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    v39 = 4294967295
                                                } else {
                                                    switch fi34 {
                                                    case 32767:
                                                        if true {
                                                            var v56 bool = (128 <= 70)
                                                            if (0.1 >= 0.1) {
                                                                if false {
                                                                    if v3 {
                                                                        switch 300 {
                                                                        case -1:
                                                                            v57 := GenS7{F0: uint64(v42), F1: "\n\t"}
                                                                            _ = v57
                                                                        case -17:
                                                                            v58 := GenS4{F0: 128, F1: (v42 + v42)}
                                                                            _ = v58
                                                                        }
                                                                        fmt.Println(v42)
                                                                        var v59 string = ""
                                                                        _ = v59
                                                                    } else {
                                                                        if v56 {
                                                                            v25 = (v25 * v25)
                                                                            if len(v1) > 0 {
                                                                                v1[0] = "hello"
                                                                            }
                                                                        }
                                                                        v60 := GenS3{F0: (fi19 >> ((397) & 63)), F1: 65535, F2: (fi34 & fi19)}
                                                                        var v61 int8 = 0
                                                                        v2 = append(v2, "")
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                v3 = (!(!true))
                                                                var v62 string = "foo bar"
                                                                _ = v62
                                                            } else {
                                                                if true {
                                                                    var v63 int64 = (v18 / ((1) | 1))
                                                                    if v3 {
                                                                        fmt.Println(fi34)
                                                                        var v64 uint32 = 1
                                                                        _ = v64
                                                                    }
                                                                    var v65 uint8 = 219
                                                                    v66 := GenS8{F0: (v42 + v42), F1: (v39 | (v39 | v39))}
                                                                    fmt.Println(v25)
                                                                    fmt.Println(v3)
                                                                    var v67 string = ""
                                                                    _ = v63
                                                                    _ = v65
                                                                    _ = v66
                                                                    _ = v67
                                                                } else {
                                                                    var v68 float32 = v25
                                                                    _ = len(v2)
                                                                    _ = v68
                                                                }
                                                                if (879 != (^(v42))) {
                                                                    v2 = append(v2, "hello")
                                                                    v18 = -603
                                                                } else {
                                                                    v3 = v3
                                                                    if (!(fi34 < (fi34 - 534))) {
                                                                        v69 := GenS4{F0: 32768, F1: v42}
                                                                        var v70 uint8 = uint8(v4)
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                    v39 = (^(v39))
                                                                }
                                                                v42 = v42
                                                            }
                                                            v71 := GenS9{F0: v39, F1: 2147483647}
                                                            _ = v56
                                                            _ = v71
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v2 = append(v2, "")
                                                            case "\n\t":
                                                                switch (fi34 + fi19) {
                                                                case 42:
                                                                    v72 := GenS2{F0: 2147483647, F1: 923, F2: 42}
                                                                    _ = v72
                                                                case 406:
                                                                    var v73 int = fi34
                                                                    _ = v73
                                                                }
                                                            default:
                                                                v74 := GenS10{F0: v3, F1: -50, F2: v18}
                                                                _ = v74
                                                            }
                                                            if v3 {
                                                                v4 -= 749
                                                                if v3 {
                                                                    v42 = 4294967295
                                                                } else {
                                                                    v39 += (^(v39))
                                                                }
                                                            } else {
                                                                v4 = v4
                                                            }
                                                            v18 += v18
                                                        }
                                                    }
                                                    switch fi19 {
                                                    case 32767:
                                                        var v75 bool = v3
                                                        _ = v75
                                                    default:
                                                        fmt.Println(fi34, fi19)
                                                    }
                                                    var v76 int = (fi34 + fi34)
                                                    var v77 string = "hello"
                                                    var v78 uint16 = (v4 % (((v4 + v4)) | 1))
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                                switch fi19 {
                                                case -2147483648:
                                                    v79 := GenS11{F0: (v42 | v42), F1: 65535}
                                                    _ = v79
                                                default:
                                                    if ((v39 - 466) > v39) {
                                                        if v3 {
                                                            var v80 uint8 = 255
                                                            var v81 int16 = 1
                                                            var v82 int = fi19
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                        v83 := GenS12{F0: "", F1: 1, F2: 100}
                                                        _ = v83
                                                    } else {
                                                        var v84 uint32 = (v42 % (((v42 ^ v42)) | 1))
                                                        if v3 {
                                                            v85 := GenS7{F0: 2, F1: "\n\t"}
                                                            v86 := GenS13{F0: -1, F1: 18446744073709551615}
                                                            if len(v1) > 0 {
                                                                v1[0] = ""
                                                            }
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            var v87 uint32 = (^(v42))
                                                            var v88 float32 = 0.1
                                                            fmt.Println(v42)
                                                            if len(v2) > 0 {
                                                                v2[0] = "hello"
                                                            }
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        v84 = (v42 ^ v84)
                                                        _ = v84
                                                    }
                                                }
                                                if true {
                                                    var v89 int8 = 127
                                                    var v90 float64 = (((1e10 - 0.1) * (1.5 * 0.1)) + 1e10)
                                                    v90 = 0.1
                                                    v25 = (v25 / (v25 * (0.0 + 0.0)))
                                                    v42 += (v42 ^ (v42 >> ((2) & 63)))
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "a":
                                                    var v91 uint = 65535
                                                    _ = v91
                                                case "hello":
                                                    v92 := GenS9{F0: v39, F1: 212}
                                                    _ = v92
                                                }
                                            default:
                                                if v3 {
                                                    if v3 {
                                                        _ = len(v2)
                                                    } else {
                                                        fmt.Println(fi19, v25, v3)
                                                    }
                                                }
                                            }
                                            v93 := GenS10{F0: (900 > 1), F1: 0, F2: v18}
                                            fmt.Println(v18, v25)
                                            v94 := GenS14{F0: v4, F1: 300, F2: 85}
                                            _ = v93
                                            _ = v94
                                        case -2147483648:
                                            var v95 float32 = float32(3.14159)
                                            _ = v95
                                        default:
                                            v1 = append(v1, "hello")
                                        }
                                        var v96 uint64 = 4294967295
                                        switch 128 {
                                        case 4294967295:
                                            v97 := GenS15{F0: ("foo bar" + (("foo bar" + (("a" + ("\n\t" + ("" + "hello"))) + "\n\t")) + "")), F1: v42}
                                            _ = v97
                                        default:
                                            v39 = (v39 >> ((584) & 63))
                                        }
                                        var v98 uint64 = 42
                                        _ = v96
                                        _ = v98
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = ("" + (("" + "\n\t") + ""))
                                    }
                                } else {
                                    fmt.Println(v4)
                                }
                                _ = v42
                            }
                            _ = fi34
                            _ = v39
                            _ = v41
                        }
                    }
                    _ = v25
                }
                if len(v1) > 0 {
                    v1[0] = ("a" + "")
                }
                _ = fi19
            }
            _ = v18
        }
    case -2147483648:
        switch "foo bar" {
        case "foo bar":
            if v3 {
                var v99 string = ((((("a" + ("" + "")) + ((("foo bar" + ("" + "hello")) + "\n\t") + "a")) + ("a" + "hello")) + ("a" + "hello")) + ("foo bar" + ((("foo bar" + (("a" + "foo bar") + "a")) + ((("" + "a") + "hello") + (("\n\t" + "hello") + "foo bar"))) + "\n\t")))
                switch v99 {
                case "":
                    fmt.Println(v3, v99)
                case "foo bar":
                    v4 = v4
                case "hello":
                    v100 := GenS5{F0: -2147483648, F1: 1}
                    _ = v100
                }
                _ = v99
            } else {
                v4 = v4
                var v101 string = ("" + "\n\t")
                _ = v101
            }
        case "a":
            fmt.Println(v3, v3)
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        var v102 int16 = -32768
        _ = v102
    }
    var v103 uint32 = 0
    switch "\n\t" {
    case "":
        fmt.Println(v3, v4, v4)
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v4, v4, v4)
        case "hello":
            for fi104 := 0; fi104 < 19; fi104++ {
                if v3 {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v4, fi104, fi104)
                    case "":
                        var v105 uint64 = 334
                        _ = v105
                    }
                } else {
                    v2 = append(v2, "a")
                    if false {
                        v103 = 167
                        var v106 bool = false
                        _ = v106
                    }
                }
                _ = fi104
            }
        case "":
            v103 += (v103 * v103)
        default:
            if len(v2) > 0 {
                v2[0] = ("\n\t" + ("a" + ("a" + ("hello" + "hello"))))
            }
        }
    default:
        if (("hello" + "") < ("hello" + "hello")) {
            for fi107 := 0; fi107 < 12; fi107++ {
                switch (fi107 - fi107) {
                case 725:
                    if (!(fi107 <= 0)) {
                        v108 := GenS11{F0: v103, F1: (^(v103))}
                        v3 = (v3 && false)
                        _ = v108
                    }
                case 42:
                    v3 = v3
                case 128:
                    for fi109 := 0; fi109 < 2; fi109++ {
                        v110 := GenS4{F0: fi109, F1: v103}
                        switch (fi107 | 2147483647) {
                        case 32768:
                            if (1 > -1) {
                                var v111 bool = v3
                                fmt.Println(fi107)
                                _ = v111
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case 26:
                            v103 = (v103 | (v103 * v103))
                        case -2147483648:
                            if true {
                                if v3 {
                                    var v112 int32 = -13
                                    fmt.Println(fi107, v3)
                                    _ = v112
                                } else {
                                    var v113 float32 = (3.14159 / (1e10 * -2.25))
                                    switch "a" {
                                    case "a":
                                        if false {
                                            v114 := GenS16{F0: v3, F1: -1, F2: 255}
                                            _ = v114
                                        }
                                    case "hello":
                                        v113 = v113
                                    case "\n\t":
                                        var v115 int64 = int64(v4)
                                        _ = v115
                                    }
                                    v103 -= v103
                                    v116 := GenS17{F0: 0, F1: 1, F2: 2147483647}
                                    v117 := GenS17{F0: 2147483647, F1: 1000, F2: -245}
                                    _ = v113
                                    _ = v116
                                    _ = v117
                                }
                            }
                        }
                        var v118 uint = 42
                        switch (fi107 + fi107) {
                        case 32767:
                            v119 := GenS3{F0: fi107, F1: v103, F2: (-(fi107))}
                            _ = v119
                        }
                        var v120 string = ""
                        _ = fi109
                        _ = v110
                        _ = v118
                        _ = v120
                    }
                }
                fmt.Println(v4, v3)
                if true {
                    v121 := GenS2{F0: -2147483648, F1: 732, F2: 109}
                    v103 += (v103 - v103)
                    v3 = v3
                    _ = v121
                } else {
                    v122 := GenS8{F0: v103, F1: 42}
                    _ = v122
                }
                _ = fi107
            }
            v103 -= v103
        }
    }
    for fi123 := 0; fi123 < 10; fi123++ {
        _ = len(v2)
        var v124 float64 = (0.0 * 0.1)
        _ = fi123
        _ = v124
    }
    var v125 uint16 = (v4 + 646)
    v125 += 1
    var v126 uint16 = v125
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v103
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
