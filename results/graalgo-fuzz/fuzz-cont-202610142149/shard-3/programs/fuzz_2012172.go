package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint8; F1 uint32; F2 int }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 string; F1 string }
type GenS16 struct { F0 uint64; F1 uint16 }
type GenS17 struct { F0 int8; F1 int16; F2 uint8 }
type GenS18 struct { F0 int8; F1 int32; F2 int64 }
func genFn0(p0 uint8) uint8 {
    v4 := GenS1{F0: (((("a" + (((("a" + "\n\t") + "foo bar") + ("" + "")) + "\n\t")) + (("\n\t" + "a") + "\n\t")) + "hello") + ("foo bar" + ("\n\t" + ("hello" + ""))))}
    v5 := []int{-344}
    var v6 int64 = 635
    v6 = (v6 | -9223372036854775808)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            switch (fi7 / (((fi8 + fi7)) | 1)) {
            case 128:
                var v9 bool = true
                _ = v9
            default:
                var v10 float32 = 3.14159
                _ = v10
            }
            _ = fi8
        }
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint16, p1 uint, p2 uint) uint16 {
    v11 := GenS2{F0: 0, F1: -32768}
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p0
}

func main() {
    var v1 int8 = 127
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v12 int32 = 42
    if false {
        if (true && (-808 > 1)) {
            switch 1 {
            case -9223372036854775808:
                switch -2147483648 {
                case -2147483648:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v14 := GenS1{F0: "foo bar"}
                        if false {
                            v1 -= (v1 / ((int8(v12)) | 1))
                        }
                        _ = fi13
                        _ = v14
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v12, v1)
                    }
                }
            case 128:
                v12 += 34
            case -525:
                v12 -= v12
            }
        }
        var v15 int8 = v1
        v15 = 80
        _ = v15
    } else {
        v12 = (v12 - 0)
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                if false {
                    if (!(genFn0(genFn0(128)) >= 40)) {
                        if false {
                            v18 := []int{fi16, (-(fi17)), 9223372036854775807, fi16}
                            _ = v18
                        }
                        v19 := GenS3{F0: ("\n\t" + "a"), F1: ("hello" + ("" + "hello")), F2: fi16}
                        _ = v19
                    } else {
                        var v20 int16 = -363
                        if (!false) {
                            v12 = v12
                            switch "a" {
                            case "hello":
                                if true {
                                    if (false && (false || ("a" > (((("" + "a") + "hello") + "") + ("" + (("hello" + "\n\t") + ("\n\t" + ""))))))) {
                                        v12 += v12
                                        v21 := []int{42, fi16}
                                        _ = v21
                                    } else {
                                        v22 := v3[2]
                                        v1 = int8(fi16)
                                        _ = v22
                                    }
                                    v23 := v2["a"]
                                    _ = v23
                                } else {
                                    var v24 int64 = 1
                                    switch int(v12) {
                                    case 9223372036854775807:
                                        if true {
                                            fmt.Println(fi16)
                                            if ((v24 * v24) == 1) {
                                                v25 := v2[("\n\t" + "a")]
                                                _ = v25
                                            }
                                            if (1e10 != (0.0 + 1e10)) {
                                                var v26 uint64 = 2
                                                _ = v26
                                            }
                                        }
                                    case 32768:
                                        fmt.Println(fi17, v20)
                                        var v27 uint16 = 350
                                        if true {
                                            v28 := GenS4{F0: 619, F1: v1}
                                            fmt.Println(fi16, v1)
                                            _ = v28
                                        } else {
                                            switch fi16 {
                                            case 2147483647:
                                                if ((fi17 + -1) != fi17) {
                                                    v29 := GenS1{F0: "foo bar"}
                                                    var v30 float64 = (((float64(-2.25) + -2.25) + -2.25) / 1.0)
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                        }
                                        _ = len(v2)
                                        v31 := []int{(^(fi17)), fi16}
                                        var v32 int32 = (-(v12))
                                        v33 := GenS2{F0: -2147483648, F1: (v20 / ((v20) | 1))}
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(fi17, v12, v24)
                                        default:
                                            _ = len(v2)
                                        }
                                        _ = len(v2)
                                        _ = v27
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v24
                                }
                            case "\n\t":
                                fmt.Println(v20)
                            case "foo bar":
                                v34 := GenS5{F0: ((fi17 & 32767) < 128), F1: 0}
                                _ = v34
                            }
                            fmt.Println(fi16, v1)
                        } else {
                            fmt.Println(v12)
                            _ = len(v3)
                            v35 := []int{504}
                            _ = v35
                        }
                        v36 := GenS2{F0: -9223372036854775808, F1: (v20 * v20)}
                        _ = v20
                        _ = v36
                    }
                    switch "hello" {
                    case "\n\t":
                        v1 = v1
                    case "a":
                        var v37 string = ("a" + "\n\t")
                        _ = v37
                    case "":
                        if false {
                            if (2147483648 >= 42) {
                                if false {
                                    fmt.Println(v12, v12)
                                    if (int64(v12) >= 42) {
                                        _ = len(v2)
                                        v38 := GenS3{F0: "foo bar", F1: "a", F2: fi17}
                                        v39 := []int{(fi17 << ((2) & 63))}
                                        _ = v38
                                        _ = v39
                                    }
                                    if (42 <= 1) {
                                        v40 := []int{-9223372036854775808, fi16}
                                        fmt.Println(fi17)
                                        v41 := GenS3{F0: "", F1: "hello", F2: (fi16 << ((4294967295) & 63))}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        fmt.Println(fi17, v12, fi16)
                                        fmt.Println(v12, v1)
                                        v42 := GenS6{F0: 127, F1: 65535, F2: 1}
                                        _ = v42
                                    }
                                } else {
                                    v43 := GenS7{F0: 2, F1: "\n\t"}
                                    if false {
                                        v44 := GenS5{F0: (!(127 >= 42)), F1: 4294967295}
                                        _ = v44
                                    } else {
                                        v1 = 126
                                        v12 += -1
                                        switch ("a" + "a") {
                                        case "":
                                            v45 := GenS8{F0: false}
                                            _ = v45
                                        case "\n\t":
                                            fmt.Println(fi16, fi17, fi17)
                                        case "foo bar":
                                            if (!false) {
                                                v46 := []int{fi17, -9223372036854775808}
                                                _ = v46
                                            } else {
                                                if (2 < 268) {
                                                    var v47 uint8 = 0
                                                    var v48 string = ""
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    fmt.Println(fi16, v1)
                                                    v12 = (v12 - v12)
                                                }
                                                if true {
                                                    var v49 int64 = 0
                                                    _ = v49
                                                }
                                                var v50 string = "foo bar"
                                                _ = v50
                                            }
                                        default:
                                            var v51 int32 = v12
                                            _ = v51
                                        }
                                    }
                                    _ = v43
                                }
                                v52 := GenS9{F0: "", F1: true}
                                v53 := GenS9{F0: (((((("foo bar" + ("\n\t" + ("a" + "hello"))) + "hello") + ("a" + ((("\n\t" + ("\n\t" + "")) + ("" + "hello")) + "\n\t"))) + "hello") + "\n\t") + ""), F1: false}
                                v12 += (v12 + v12)
                                _ = v52
                                _ = v53
                            }
                            var v54 string = ((("\n\t" + "foo bar") + (("foo bar" + ("\n\t" + "\n\t")) + "foo bar")) + "")
                            _ = v54
                        } else {
                            var v55 bool = true
                            _ = v55
                        }
                        v12 -= (-(v12))
                    }
                }
                if (148 < 1) {
                    fmt.Println(fi17, fi17)
                    var v56 int64 = -1
                    v57 := GenS2{F0: 32767, F1: 58}
                    _ = v56
                    _ = v57
                } else {
                    if true {
                        switch fi17 {
                        case -148:
                            if false {
                                var v58 int16 = -1
                                switch "foo bar" {
                                case "hello":
                                    if true {
                                        v59 := GenS6{F0: 2, F1: 394, F2: -748}
                                        _ = v59
                                    }
                                case "":
                                    var v60 float64 = ((float64(-2.25) - 0.0) * float64((3.14159 / 1e10)))
                                    var v61 int = fi17
                                    _ = v60
                                    _ = v61
                                }
                                v1 -= 0
                                _ = v58
                            }
                        case 996:
                            if (true && true) {
                                v12 += (v12 * (-(v12)))
                                var v62 uint32 = 1
                                if true {
                                    fmt.Println(v1, fi17, v62)
                                } else {
                                    v12 = v12
                                    if (1 >= 65535) {
                                        var v63 uint = 655
                                        _ = v63
                                    } else {
                                        fmt.Println(v12)
                                        var v64 uint16 = 2
                                        _ = v64
                                    }
                                    switch 2147483647 {
                                    case 128:
                                        v62 = 2147483647
                                    }
                                }
                                var v65 int64 = 9223372036854775807
                                v66 := GenS10{F0: 255}
                                _ = v62
                                _ = v65
                                _ = v66
                            }
                        case -619:
                            v67 := GenS8{F0: true}
                            _ = v67
                        default:
                            switch fi16 {
                            case 4294967295:
                                if true {
                                    v12 += 706
                                    v68 := []int{(fi17 - (^(fi16)))}
                                    v69 := GenS5{F0: true, F1: 645}
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v70 int32 = v12
                                    var v71 float64 = ((0.1 - float64(float32(0.0))) - (0.1 + (1.5 * 3.14159)))
                                    if (793 == 2147483647) {
                                        v72 := GenS1{F0: ""}
                                        if true {
                                            v70 += v12
                                            v73 := GenS2{F0: -586, F1: 1}
                                            if true {
                                                switch (("foo bar" + "") + "hello") {
                                                case "foo bar":
                                                    fmt.Println(fi17, v1, v1)
                                                case "":
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (1 > 42) {
                                                            delete(v2, "a")
                                                            var v74 int = fi17
                                                            v1 = v1
                                                            _ = v74
                                                        } else {
                                                            fmt.Println(v12, fi17, v12)
                                                            var v75 bool = (((1e10 - v71) + float64((0.1 * float32(v71)))) > float64((0.0 - -2.25)))
                                                            if true {
                                                                var v76 int32 = v70
                                                                _ = v76
                                                            }
                                                            _ = v75
                                                        }
                                                    case "a":
                                                        fmt.Println(v1, fi16)
                                                        v77 := v3[4294967295]
                                                        _ = v77
                                                    }
                                                }
                                                if true {
                                                    switch (fi17 >> ((255) & 63)) {
                                                    case 2147483647:
                                                        v78 := GenS11{F0: 2, F1: 1}
                                                        _ = v78
                                                    }
                                                    _ = len(v2)
                                                    v70 -= (v12 - v12)
                                                    switch fi17 {
                                                    case -355:
                                                        v79 := []int{32767, fi17, (fi17 % ((fi17) | 1)), -467}
                                                        _ = v79
                                                    }
                                                } else {
                                                    var v80 uint16 = 0
                                                    v81 := []int{(fi16 / ((127) | 1)), (fi17 ^ (-(fi16)))}
                                                    v80 = v80
                                                    _ = v80
                                                    _ = v81
                                                }
                                            } else {
                                                v12 = (v70 & (v12 << ((403) & 63)))
                                                switch ("a" + "a") {
                                                case "":
                                                    switch (fi17 * (-(fi17))) {
                                                    case 796:
                                                        var v82 uint8 = 1
                                                        _ = v82
                                                    case 2147483647:
                                                        var v83 float64 = v71
                                                        _ = v83
                                                    case 9223372036854775807:
                                                        if false {
                                                            v70 += (v70 / ((-2147483648) | 1))
                                                            v84 := []int{(-(fi16)), 494, -1, fi17}
                                                            if ((false && false) || (0 == 4294967295)) {
                                                                v70 = -2147483648
                                                                if false {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch ("hello" + "hello") {
                                                                        case "":
                                                                            var v85 string = "foo bar"
                                                                            _ = v85
                                                                        case "foo bar":
                                                                            v1 += (v1 ^ v1)
                                                                        case "\n\t":
                                                                            if ((0 != 434) || false) {
                                                                                v1 -= (v1 / ((v1) | 1))
                                                                                switch fi16 {
                                                                                case 32768:
                                                                                    delete(v3, 772)
                                                                                case 128:
                                                                                    if (false || false) {
                                                                                        v86 := GenS4{F0: 855, F1: (-(v1))}
                                                                                        fmt.Println(fi17, v70)
                                                                                        var v87 uint = 18446744073709551615
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    }
                                                                                case 280:
                                                                                    if true {
                                                                                        var v88 uint = 227
                                                                                        _ = v88
                                                                                    } else {
                                                                                        if false {
                                                                                            v89 := GenS9{F0: "hello", F1: (true || ("hello" >= ("a" + ("a" + "a"))))}
                                                                                            var v90 uint16 = 42
                                                                                            v71 = v71
                                                                                            var v91 uint = 4294967295
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        } else {
                                                                                            v92 := GenS11{F0: 2147483648, F1: 158}
                                                                                            _ = v92
                                                                                        }
                                                                                        v1 = (v1 + (-(v1)))
                                                                                        v93 := GenS8{F0: (!true)}
                                                                                        var v94 int = fi16
                                                                                        if true {
                                                                                            v1 = v1
                                                                                            if true {
                                                                                                v1 = 42
                                                                                                v95 := GenS3{F0: "hello", F1: ("a" + (("a" + "foo bar") + "")), F2: 0}
                                                                                                var v96 int16 = int16(v70)
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            } else {
                                                                                                v12 -= v70
                                                                                                fmt.Println(v94, v12)
                                                                                            }
                                                                                        }
                                                                                        if (!false) {
                                                                                            v1 -= v1
                                                                                        } else {
                                                                                            v71 -= v71
                                                                                            fmt.Println(v70, fi16, v71)
                                                                                        }
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v2["a"] = fi16
                                                                        }
                                                                    case "a":
                                                                        var v97 uint64 = 1
                                                                        _ = v97
                                                                    case "foo bar":
                                                                        v98 := GenS5{F0: (65535 >= 2), F1: 42}
                                                                        var v99 uint8 = 255
                                                                        _ = v98
                                                                        _ = v99
                                                                    default:
                                                                        if (0 >= 0) {
                                                                            var v100 int = fi17
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                    var v101 uint16 = 813
                                                                    _ = len(v2)
                                                                    _ = v101
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(v12)
                                                                    }
                                                                    var v102 int64 = -1
                                                                    v103 := GenS12{F0: 18446744073709551615}
                                                                    _ = v102
                                                                    _ = v103
                                                                }
                                                            }
                                                            _ = v84
                                                        } else {
                                                            if (1000 >= 386) {
                                                                fmt.Println(fi16)
                                                                v70 = int32(fi17)
                                                                if false {
                                                                    v104 := []int{0}
                                                                    v1 = v1
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v105 int8 = -84
                                                                        _ = v105
                                                                    case "":
                                                                        _ = len(v104)
                                                                    case "\n\t":
                                                                        v1 += v1
                                                                        fmt.Println(v71, v1)
                                                                        v106 := GenS13{F0: -325, F1: ("foo bar" + "")}
                                                                        _ = v106
                                                                    default:
                                                                        var v107 int8 = v1
                                                                        _ = v107
                                                                    }
                                                                    _ = v104
                                                                } else {
                                                                    fmt.Println(fi16, fi16, v12)
                                                                }
                                                            } else {
                                                                var v108 bool = (false && ((((v1 ^ -128) == (^(v1))) || ("a" >= "\n\t")) || (true || false)))
                                                                if v108 {
                                                                    v70 -= v12
                                                                    delete(v3, 42)
                                                                }
                                                                var v109 int8 = -65
                                                                var v110 string = ""
                                                                if v108 {
                                                                    v111 := GenS8{F0: false}
                                                                    _ = v111
                                                                } else {
                                                                    var v112 bool = (v108 && true)
                                                                    v113 := GenS14{F0: 9223372036854775807}
                                                                    if v112 {
                                                                        var v114 int64 = 0
                                                                        var v115 uint64 = 436
                                                                        _ = v114
                                                                        _ = v115
                                                                    } else {
                                                                        delete(v2, (v110 + "a"))
                                                                        fmt.Println(v71, fi17, v71)
                                                                    }
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                                if ((fi16 ^ fi16) >= 128) {
                                                                    var v116 int64 = 0
                                                                    _ = v116
                                                                } else {
                                                                    if (float32(3.14159) != -2.25) {
                                                                        v70 -= (v70 % ((v12) | 1))
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                    fmt.Println(v109, v71, v1)
                                                                    var v117 uint32 = 1
                                                                    _ = v117
                                                                }
                                                                _ = v108
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                            v70 += (v70 % (((v12 / (((^(v12))) | 1))) | 1))
                                                            var v118 int = (fi16 | fi16)
                                                            v71 = (v71 + v71)
                                                            if (0 != 127) {
                                                                v119 := GenS13{F0: 0, F1: "\n\t"}
                                                                switch (("" + "\n\t") + "\n\t") {
                                                                case "":
                                                                    if (("foo bar" + "\n\t") > (("\n\t" + ("hello" + "hello")) + "a")) {
                                                                        v71 = float64(-2.25)
                                                                        var v120 string = "\n\t"
                                                                        v121 := v3[2147483647]
                                                                        _ = v120
                                                                        _ = v121
                                                                    } else {
                                                                        switch fi17 {
                                                                        case 2147483647:
                                                                            if (!(!(742 != 1))) {
                                                                                var v122 float64 = (v71 - (0.1 * -2.25))
                                                                                v3[4294967295] = "hello"
                                                                                v123 := GenS15{F0: (("hello" + "") + ("foo bar" + "a")), F1: "a"}
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                        default:
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v124 int16 = 1
                                                                                _ = v124
                                                                            case "hello":
                                                                                v1 -= (v1 / ((v1) | 1))
                                                                            }
                                                                            if (1 <= (v70 ^ (v12 >> ((1) & 63)))) {
                                                                                var v125 int16 = 881
                                                                                var v126 uint16 = 255
                                                                                if ((v1 % ((v1) | 1)) < v1) {
                                                                                    if true {
                                                                                        fmt.Println(v70, fi17, v1)
                                                                                    } else {
                                                                                        var v127 int8 = v1
                                                                                        fmt.Println(v118)
                                                                                        v118 += (v118 + 581)
                                                                                        var v128 int16 = v125
                                                                                        switch int(v128) {
                                                                                        case -2147483648:
                                                                                            var v129 uint16 = 32767
                                                                                            _ = v129
                                                                                        case 1:
                                                                                            if ((true && true) || (false && true)) {
                                                                                                fmt.Println(v127)
                                                                                                v130 := GenS16{F0: 0, F1: v126}
                                                                                                _ = v130
                                                                                            }
                                                                                        }
                                                                                        _ = v127
                                                                                        _ = v128
                                                                                    }
                                                                                }
                                                                                _ = v125
                                                                                _ = v126
                                                                            }
                                                                            var v131 uint8 = 128
                                                                            _ = v131
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v132 := GenS17{F0: (v1 * (v1 ^ v1)), F1: 0, F2: 1}
                                                                    _ = v132
                                                                }
                                                                _ = v119
                                                            } else {
                                                                delete(v2, "hello")
                                                                var v133 float64 = v71
                                                                v1 -= (v1 % ((-4) | 1))
                                                                _ = v133
                                                            }
                                                            _ = v118
                                                        }
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        var v134 int16 = 0
                                                        v12 = (v70 << ((65535) & 63))
                                                        _ = v134
                                                    } else {
                                                        v1 = 69
                                                        _ = len(v2)
                                                    }
                                                case "foo bar":
                                                    v71 = v71
                                                default:
                                                    var v135 uint8 = 196
                                                    _ = v135
                                                }
                                            }
                                            _ = v73
                                        } else {
                                            fmt.Println(fi16, v71, fi16)
                                            fmt.Println(v70, fi16, v71)
                                        }
                                        v136 := GenS17{F0: v1, F1: 42, F2: 1}
                                        _ = v136
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                fmt.Println(v1)
                            case 966:
                                v1 -= v1
                            default:
                                var v137 int8 = (v1 ^ (v1 << ((1) & 63)))
                                _ = v137
                            }
                        }
                        v12 += v12
                        if (!((18446744073709551615 == 2) || (true && false))) {
                            var v138 uint64 = 18446744073709551615
                            _ = v138
                        } else {
                            var v139 float64 = 0.0
                            _ = v139
                        }
                        if true {
                            v1 += v1
                            switch "a" {
                            case "a":
                                if (true && ((-(v1)) > v1)) {
                                    fmt.Println(v1)
                                    fmt.Println(v1, v1)
                                }
                            default:
                                v12 = -2147483648
                            }
                        }
                        v140 := GenS1{F0: ""}
                        v141 := GenS18{F0: v1, F1: v12, F2: -9223372036854775808}
                        v1 = v1
                        v1 = v1
                        var v142 int = int(v1)
                        var v143 uint32 = 4294967295
                        var v144 int16 = 0
                        _ = v140
                        _ = v141
                        _ = v142
                        _ = v143
                        _ = v144
                    } else {
                        v145 := GenS13{F0: -830, F1: ("\n\t" + "")}
                        v1 += v1
                        if true {
                            v146 := GenS8{F0: (("\n\t" + "a") != (((("" + "a") + "a") + ("" + "a")) + "\n\t"))}
                            v1 -= (v1 | 100)
                            switch fi17 {
                            case 32767:
                                var v147 int = fi17
                                _ = v147
                            }
                            _ = v146
                        }
                        _ = v145
                    }
                    var v148 float32 = (1.5 + (0.1 * 1.5))
                    _ = v148
                }
                var v149 uint = 533
                _ = fi17
                _ = v149
            }
            switch fi16 {
            case 1:
                var v150 bool = (65535 <= 42)
                _ = v150
            case 128:
                fmt.Println(v12)
            }
            fmt.Println(v12, v12)
            _ = fi16
        }
        v12 = v12
    }
    var v151 uint = 392
    switch "hello" {
    case "\n\t":
        fmt.Println(v151, v12, v1)
    case "hello":
        var v152 uint16 = 835
        _ = v152
    case "a":
        if ((v12 > 1) && (4294967295 != 42)) {
            var v153 int64 = -970
            v154 := []int{309, 572}
            var v155 bool = false
            _ = v153
            _ = v154
            _ = v155
        }
    }
    if (-2.25 > (0.1 - ((1.5 * 1.5) - (1e10 / -2.25)))) {
        v151 = (v151 << (((v151 * (v151 | 42))) & 63))
    } else {
        v151 = (v151 & (v151 & (^(v151))))
        if true {
            v12 -= v12
        } else {
            v12 = v12
            if (42 <= 42) {
                if (v1 < 126) {
                    v1 += 1
                }
            }
        }
    }
    fmt.Println(v151, v12)
    fmt.Println(v151, v12)
    v151 = v151
    v12 = (v12 % ((-360) | 1))
    var v156 int64 = 9223372036854775807
    switch (("a" + ("hello" + ("\n\t" + "foo bar"))) + "hello") {
    case "\n\t":
        v151 = (v151 - (v151 + v151))
    case "foo bar":
        v157 := GenS11{F0: 42, F1: 618}
        _ = v157
    case "hello":
        var v158 int32 = (v12 - v12)
        _ = v158
    default:
        fmt.Println(v156, v12)
    }
    v159 := v3[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v151:", v151)
    fmt.Println("v156:", v156)
    fmt.Println("v159:", v159)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v151
    _ = v156
    _ = v159
    _ = v2
    _ = v3
}
