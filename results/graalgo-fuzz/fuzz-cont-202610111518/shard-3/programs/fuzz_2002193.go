package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 int32; F1 int; F2 uint32 }
func genFn0(p0 uint32, p1 int16) uint32 {
    v4 := GenS1{F0: (^(p0)), F1: 1}
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int8 = 0
            p0 = 65535
            if (2147483647 >= -823) {
                var v8 float64 = float64((0.1 + 0.0))
                _ = v8
            }
            v9 := GenS2{F0: "a", F1: p1}
            var v10 uint64 = 18446744073709551615
            v11 := []int{(fi6 * fi6)}
            var v12 uint32 = p0
            _ = fi6
            _ = v10
            _ = v11
            _ = v12
            _ = v7
            _ = v9
        }
        p1 -= -32768
        fmt.Println(p0)
        _ = fi5
    }
    if (65535 == 255) {
        var v13 float64 = 1.5
        switch 128 {
        case -715:
            v14 := GenS2{F0: "foo bar", F1: (^(p1))}
            _ = v14
        case 0:
            v15 := []int{2147483647, 2147483647, 32767}
            _ = v15
        default:
            var v16 string = "hello"
            _ = v16
        }
        v13 = 0.1
        _ = v13
    }
    var v17 string = (("a" + "\n\t") + "a")
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return (^(p0))
}
func genFn1(p0 uint8, p1 int32) int {
    var v18 float64 = 0.1
    var v19 uint8 = (p0 / (((p0 & p0)) | 1))
    fmt.Println(v19, p0)
    fmt.Println(v19, v19, p0)
    if true {
        if true {
            var v20 int8 = 1
            switch 42 {
            case -808:
                v20 = 0
            case 9223372036854775807:
                for fi21 := 0; fi21 < 14; fi21++ {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 float64 = v18
                        fmt.Println(v23, v20, fi22)
                        var v24 bool = false
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS2{F0: "", F1: 844}
                    for fi26 := 0; fi26 < 9; fi26++ {
                        switch fi26 {
                        case -457:
                            if (true && true) {
                                var v27 bool = false
                                if false {
                                    v27 = ((((("" + "hello") + "\n\t") + "") + (("a" + "\n\t") + ("foo bar" + "foo bar"))) < ("a" + "\n\t"))
                                    v28 := []int{(^(fi26)), fi21, 414}
                                    _ = v28
                                } else {
                                    if (605 == 18446744073709551615) {
                                        v19 = (v19 ^ (v19 ^ (^(v19))))
                                    }
                                    v29 := []int{32768, int(v20), (-(fi26)), (fi26 * fi26)}
                                    p0 = v19
                                    if v27 {
                                        if (!("" < "\n\t")) {
                                            if len(v29) > 0 {
                                                v29[0] = (fi21 % ((fi26) | 1))
                                            }
                                            v30 := GenS1{F0: 2147483648, F1: v19}
                                            v31 := GenS2{F0: ("foo bar" + "foo bar"), F1: -32768}
                                            _ = v30
                                            _ = v31
                                        } else {
                                            p0 += v19
                                            p1 = 0
                                            if v27 {
                                                v29 = append(v29, -9223372036854775808)
                                            }
                                            fmt.Println(v20, v20)
                                        }
                                    } else {
                                        var v32 float64 = (((v18 / (v18 - float64(1e10))) - 0.1) * (v18 + float64(3.14159)))
                                        p1 -= p1
                                        p0 -= 127
                                        var v33 uint8 = 42
                                        var v34 int8 = (v20 * -123)
                                        switch 4294967295 {
                                        case -9223372036854775808:
                                            v27 = (494 > 42)
                                        }
                                        v18 += v32
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    if (1 < 1) {
                                        _ = len(v29)
                                    }
                                    _ = v29
                                }
                                v35 := GenS2{F0: "\n\t", F1: 32767}
                                if false {
                                    p0 = (^(p0))
                                    v36 := GenS3{F0: 866, F1: -612}
                                    _ = v36
                                } else {
                                    if false {
                                        v37 := GenS4{F0: 2147483648, F1: 0}
                                        _ = v37
                                    }
                                    fmt.Println(v20)
                                }
                                _ = v27
                                _ = v35
                            }
                        case -52:
                            v38 := GenS5{F0: 804, F1: fi26}
                            _ = v38
                        case 9223372036854775807:
                            v18 -= float64(1.5)
                            if (false || (!true)) {
                                v39 := GenS4{F0: 1, F1: genFn0(62, -886)}
                                v40 := GenS5{F0: 18446744073709551615, F1: int(p1)}
                                _ = v39
                                _ = v40
                            } else {
                                v41 := GenS6{F0: ("" + ((("a" + ("foo bar" + ("hello" + (("a" + "foo bar") + "hello")))) + "") + "a"))}
                                if true {
                                    v42 := GenS5{F0: 65535, F1: (fi26 | (fi26 << ((438) & 63)))}
                                    if false {
                                        fmt.Println(v20)
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v43 string = "hello"
                                            _ = v43
                                        case "hello":
                                            v19 += p0
                                        case "a":
                                            v44 := GenS1{F0: 132, F1: p0}
                                            _ = v44
                                        }
                                        var v45 int32 = p1
                                        _ = v45
                                    }
                                    if (true && (42 >= 184)) {
                                        v46 := GenS6{F0: ("" + "a")}
                                        if (((((v20 & 65) != 100) && true) || (true && true)) || false) {
                                            switch (("hello" + ("foo bar" + "a")) + "a") {
                                            case "foo bar":
                                                v47 := []int{fi26, (fi21 - fi26)}
                                                _ = v47
                                            case "\n\t":
                                                var v48 string = ((("foo bar" + ("hello" + "a")) + "foo bar") + "")
                                                _ = v48
                                            default:
                                                if false {
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(v19, v19)
                                                    case "":
                                                        if ((-895 < 531) && true) {
                                                            if (325 <= genFn0(genFn0(0, -1), 730)) {
                                                                var v49 string = ("a" + ("\n\t" + "a"))
                                                                if ("foo bar" < v49) {
                                                                    fmt.Println(v49, v19, v49)
                                                                    var v50 int64 = 9223372036854775807
                                                                    _ = v50
                                                                }
                                                                var v51 int8 = (v20 + -31)
                                                                _ = v49
                                                                _ = v51
                                                            } else {
                                                                p0 -= (^(v19))
                                                                v20 = v20
                                                                if (false || (-124 < 625)) {
                                                                    if ((!false) || false) {
                                                                        switch fi26 {
                                                                        case -379:
                                                                            fmt.Println(v19)
                                                                        }
                                                                        v18 -= v18
                                                                        v52 := GenS7{F0: 444, F1: 32767}
                                                                        _ = v52
                                                                    }
                                                                }
                                                            }
                                                            var v53 int32 = (p1 % (((p1 ^ (^(p1)))) | 1))
                                                            _ = v53
                                                        }
                                                    default:
                                                        fmt.Println(fi26, v19)
                                                    }
                                                }
                                            }
                                        } else {
                                            var v54 string = "a"
                                            switch (fi26 * (-(fi21))) {
                                            case -355:
                                                v55 := GenS7{F0: 383, F1: 493}
                                                _ = v55
                                            case 731:
                                                switch fi21 {
                                                case 1:
                                                    var v56 float64 = ((((v18 - ((1e10 - v18) / (v18 / (((float64(0.1) / v18) - v18) + (v18 * float64(1.5)))))) / (0.0 / v18)) - (v18 + ((v18 / (float64(0.1) - float64(0.0))) * (0.1 - (0.1 / float64(-2.25)))))) * v18)
                                                    _ = v56
                                                case 578:
                                                    v57 := []int{fi26, fi26}
                                                    _ = v57
                                                case 0:
                                                    v54 = v54
                                                default:
                                                    var v58 int64 = 42
                                                    _ = v58
                                                }
                                            default:
                                                switch v54 {
                                                case "a":
                                                    var v59 uint = 255
                                                    _ = v59
                                                case "":
                                                    v60 := []int{fi21, -894}
                                                    _ = v60
                                                case "\n\t":
                                                    v61 := []int{(^(fi21)), 213, fi26}
                                                    _ = v61
                                                default:
                                                    switch 32767 {
                                                    case 2147483647:
                                                        v62 := GenS1{F0: 15, F1: v19}
                                                        _ = v62
                                                    }
                                                }
                                            }
                                            fmt.Println(v54, v19, v19)
                                            _ = v54
                                        }
                                        if (3.14159 >= 1e10) {
                                            var v63 int8 = 100
                                            v19 = (^(v19))
                                            v18 += v18
                                            _ = v63
                                        } else {
                                            v18 = 3.14159
                                            switch ((("" + "foo bar") + "foo bar") + "a") {
                                            case "":
                                                var v64 string = ""
                                                _ = v64
                                            case "hello":
                                                v65 := []int{}
                                                _ = v65
                                            case "foo bar":
                                                var v66 float32 = 3.14159
                                                _ = v66
                                            default:
                                                var v67 int = 551
                                                _ = v67
                                            }
                                            v68 := GenS2{F0: ("foo bar" + "a"), F1: 32767}
                                            var v69 bool = ((1 != 65535) || ((541 == genFn0(1, 1000)) || ("a" == "hello")))
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = v46
                                    } else {
                                        v70 := GenS2{F0: ("a" + "hello"), F1: -32768}
                                        v71 := []int{350, (fi26 * -449), fi21}
                                        fmt.Println(fi21)
                                        if false {
                                            fmt.Println(fi26, v20)
                                            v72 := GenS8{F0: 0, F1: 1}
                                            fmt.Println(p1, p1, p0)
                                            _ = v72
                                        }
                                        switch 127 {
                                        case -2147483648:
                                            var v73 uint64 = 1
                                            _ = v73
                                        default:
                                            fmt.Println(v18)
                                        }
                                        v74 := GenS7{F0: 133, F1: 2}
                                        v75 := GenS3{F0: 2, F1: p1}
                                        _ = v70
                                        _ = v71
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = v42
                                } else {
                                    switch "hello" {
                                    case "":
                                        v76 := GenS9{F0: (-(v20))}
                                        _ = v76
                                    case "\n\t":
                                        if ((((v18 + 0.0) / float64(-2.25)) >= (v18 * ((float64(-2.25) + 3.14159) - v18))) && (!(0 < 18446744073709551615))) {
                                            var v77 int64 = -388
                                            v78 := []int{}
                                            _ = v77
                                            _ = v78
                                        } else {
                                            var v79 int32 = 0
                                            fmt.Println(fi26, v20)
                                            p0 += (p0 + p0)
                                            if (!false) {
                                                var v80 int64 = -887
                                                _ = v80
                                            } else {
                                                v81 := []int{-677}
                                                if (((-2.25 + 0.1) + 3.14159) > ((0.0 - 1e10) - -2.25)) {
                                                    var v82 string = ("a" + ("\n\t" + (("" + "foo bar") + ("foo bar" + "\n\t"))))
                                                    _ = v82
                                                } else {
                                                    if (true && (-145 == fi21)) {
                                                        if len(v81) > 0 {
                                                            v81[0] = 9223372036854775807
                                                        }
                                                        v83 := GenS1{F0: 2147483647, F1: v19}
                                                        _ = v83
                                                    }
                                                    if ((65535 == 952) || (25 >= 2)) {
                                                        if (false && (2 > 313)) {
                                                            fmt.Println(v79, fi26, p0)
                                                            fmt.Println(v19, fi21, v19)
                                                            p1 -= v79
                                                        }
                                                        fmt.Println(v20)
                                                    }
                                                }
                                                if (!(uint64(v19) == 42)) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v84 int8 = -1
                                                        _ = v84
                                                    case "\n\t":
                                                        v85 := GenS4{F0: 42, F1: genFn0(4294967295, 1)}
                                                        _ = v85
                                                    default:
                                                        v81 = append(v81, (fi21 - (fi21 / ((-250) | 1))))
                                                    }
                                                    v20 += (^(v20))
                                                    v86 := GenS10{F0: 42, F1: 0, F2: 871}
                                                    _ = v86
                                                }
                                                _ = v81
                                            }
                                            var v87 uint = 4294967295
                                            _ = v79
                                            _ = v87
                                        }
                                    }
                                }
                                var v88 string = "hello"
                                _ = v41
                                _ = v88
                            }
                            fmt.Println(v18, v18)
                        default:
                            fmt.Println(p0, p0, v18)
                        }
                        v19 -= (^(v19))
                        if (356 == 2) {
                            v18 = 1.5
                        }
                        if (!(-9223372036854775808 <= 1)) {
                            if true {
                                fmt.Println(fi21)
                                var v89 int8 = v20
                                v90 := []int{}
                                _ = v89
                                _ = v90
                            } else {
                                v19 = 1
                            }
                            v91 := []int{(fi26 / (((^(fi21))) | 1)), (^(fi21)), fi26, (fi26 << ((338) & 63))}
                            _ = v91
                        } else {
                            fmt.Println(fi21, fi21)
                        }
                        switch fi21 {
                        case -360:
                            fmt.Println(v19)
                        default:
                            v92 := GenS9{F0: 127}
                            _ = v92
                        }
                        _ = fi26
                    }
                    _ = fi21
                    _ = v25
                }
            case 143:
                if (!false) {
                    for fi93 := 0; fi93 < 2; fi93++ {
                        var v94 uint16 = 65535
                        _ = fi93
                        _ = v94
                    }
                    switch "a" {
                    case "":
                        for fi95 := 0; fi95 < 11; fi95++ {
                            v18 += v18
                            _ = fi95
                        }
                    case "foo bar":
                        v20 = v20
                    case "\n\t":
                        v20 += v20
                    }
                }
            }
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 0
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    var v3 string = v1
    v96 := []int{1, 4294967295, 2147483647}
    var v97 string = v1
    delete(v2, v1)
    var v98 int = 4294967295
    for fi99 := 0; fi99 < 13; fi99++ {
        switch fi99 {
        case -485:
            if (!(15 < 649)) {
                v3 = (v3 + v3)
                for fi100 := 0; fi100 < 14; fi100++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v101 float32 = (((-2.25 + 1.5) + ((0.0 * ((-2.25 - float32(0.0)) + (0.0 - float32(-2.25)))) - float32(0.1))) * 1.5)
                        _ = v101
                    case "":
                        if (!true) {
                            if len(v96) > 0 {
                                v96[0] = (v98 * genFn1(0, 0))
                            }
                            v102 := GenS1{F0: genFn0(42, -1), F1: 0}
                            _ = v102
                        } else {
                            if (42 != 236) {
                                v103 := GenS11{F0: (!false)}
                                switch ((("\n\t" + ((((v97 + ("hello" + v97)) + (v1 + (("a" + (("" + v3) + v97)) + "\n\t"))) + (v3 + v97)) + "hello")) + v97) + v3) {
                                case "":
                                    fmt.Println(fi100, fi100, fi100)
                                case "foo bar":
                                    var v104 int8 = 0
                                    var v105 uint16 = 2
                                    v96 = append(v96, fi99)
                                    var v106 int32 = -181
                                    switch ("hello" + v1) {
                                    case "a":
                                        if true {
                                            _ = len(v96)
                                        }
                                    case "\n\t":
                                        switch (v98 + (fi99 * v98)) {
                                        case -9223372036854775808:
                                            if true {
                                                v1 = v1
                                                var v107 float32 = (0.0 + 3.14159)
                                                _ = v107
                                            } else {
                                                v105 = (^(v105))
                                                v108 := GenS1{F0: genFn0(42, -32768), F1: 42}
                                                _ = v108
                                            }
                                        case 32768:
                                            v98 += (-(fi100))
                                        case 128:
                                            switch "" {
                                            case "foo bar":
                                                v109 := GenS12{F0: 32767, F1: genFn1(uint8(v105), v106)}
                                                _ = v109
                                            default:
                                                var v110 uint16 = (v105 ^ 765)
                                                _ = v110
                                            }
                                        }
                                    case "foo bar":
                                        var v111 int64 = -427
                                        _ = v111
                                    }
                                    fmt.Println(fi99)
                                    v3 = (("foo bar" + v97) + "")
                                    fmt.Println(v3, fi99)
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                case "\n\t":
                                    _ = len(v96)
                                }
                                _ = v103
                            } else {
                                v96 = append(v96, (-(fi99)))
                                v1 = "hello"
                            }
                        }
                    default:
                        v1 = (v97 + "hello")
                    }
                    delete(v2, (v1 + ""))
                    var v112 int64 = 9223372036854775807
                    var v113 uint32 = genFn0(273, -1)
                    fmt.Println(fi99)
                    var v114 float32 = 1.5
                    v115 := GenS1{F0: (v113 % ((2) | 1)), F1: 0}
                    _ = fi100
                    _ = v112
                    _ = v113
                    _ = v114
                    _ = v115
                }
            } else {
                switch (v1 + v1) {
                case "\n\t":
                    var v116 float32 = (float32(-2.25) + 1.5)
                    _ = v116
                case "foo bar":
                    for fi117 := 0; fi117 < 15; fi117++ {
                        v98 -= (fi99 << ((439) & 63))
                        var v118 uint8 = 255
                        _ = fi117
                        _ = v118
                    }
                }
                v119 := GenS1{F0: 2147483648, F1: 128}
                _ = v119
            }
        case -9223372036854775808:
            var v120 string = "hello"
            _ = v120
        case 1:
            switch (v98 % ((genFn1(27, 2147483647)) | 1)) {
            case 32767:
                var v121 uint8 = 127
                _ = v121
            }
        default:
            var v122 uint64 = 18446744073709551615
            _ = v122
        }
        delete(v2, "a")
        v98 -= genFn1(127, 2147483647)
        _ = fi99
    }
    var v123 int16 = 821
    if (127 != 2) {
        fmt.Println(v98, v3, v98)
        switch (v98 % ((1) | 1)) {
        case 128:
            var v124 uint32 = 2
            _ = v124
        }
    }
    if ((!true) || false) {
        v125 := GenS10{F0: 65535, F1: 438, F2: v123}
        v126 := GenS13{F0: 42, F1: genFn1(2, -722), F2: 65535}
        _ = v125
        _ = v126
    } else {
        fmt.Println(v98)
        switch v3 {
        case "hello":
            var v127 bool = (true && false)
            _ = v127
        }
        var v128 float32 = (1e10 / 1.0)
        _ = v128
    }
    v129 := GenS6{F0: ("\n\t" + "hello")}
    if false {
        var v130 uint64 = 241
        switch (v98 + (v98 + (v98 ^ (-(v98))))) {
        case -9223372036854775808:
            var v131 int = (v98 << ((42) & 63))
            _ = v131
        case 127:
            if ("\n\t" <= "hello") {
                if (42 >= 65535) {
                    var v132 bool = (true && (true && (!(42 >= 255))))
                    _ = v132
                }
            }
        }
        fmt.Println(v98)
        _ = v130
    }
    if false {
        var v133 int64 = 9223372036854775807
        for fi134 := 0; fi134 < 17; fi134++ {
            delete(v2, (v3 + v1))
            for fi135 := 0; fi135 < 17; fi135++ {
                _ = len(v2)
                _ = fi135
            }
            v136 := v2["hello"]
            _ = fi134
            _ = v136
        }
        _ = v133
    }
    var v137 int = v98
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v137:", v137)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v96:", len(v96), v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v123
    _ = v129
    _ = v137
    _ = v2
    _ = v3
    _ = v96
    _ = v97
    _ = v98
}
