package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int8; F1 int8; F2 int16 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint; F1 string; F2 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool; F1 int }
type GenS11 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS12 struct { F0 bool; F1 int32; F2 int16 }
type GenS13 struct { F0 uint64; F1 uint }
type GenS14 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64, p1 string) bool {
    fmt.Println(p1, p0)
    var v5 int = -483
    var v6 int32 = 1
    v7 := []int{v5}
    if (65535 >= 65535) {
        switch v5 {
        case 9223372036854775807:
            v8 := GenS1{F0: 42, F1: p0, F2: 1}
            _ = v8
        case -2147483648:
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint64 = p0
                _ = fi9
                _ = v10
            }
        }
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (true || ((1 == (-(v6))) || true)) {
                v12 := GenS2{F0: 0, F1: 255, F2: 783}
                _ = len(v7)
                _ = v12
            } else {
                if (970 >= v6) {
                    v6 = v6
                    if false {
                        if false {
                            switch 127 {
                            case -844:
                                for fi13 := 0; fi13 < 11; fi13++ {
                                    v14 := GenS3{F0: 46, F1: 1, F2: false}
                                    var v15 bool = true
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                        } else {
                            _ = len(v7)
                            switch p1 {
                            case "a":
                                if true {
                                    fmt.Println(p1)
                                    fmt.Println(v5, v5, v6)
                                } else {
                                    v16 := GenS4{F0: p0, F1: fi11}
                                    _ = v16
                                }
                            default:
                                var v17 int32 = v6
                                _ = v17
                            }
                        }
                        var v18 uint16 = 255
                        v18 -= 2
                        _ = v18
                    } else {
                        p1 = "hello"
                    }
                    fmt.Println(v6)
                } else {
                    fmt.Println(p1)
                }
            }
            v5 = fi11
            _ = fi11
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            p1 = ""
            _ = fi19
        }
        fmt.Println(v5, v6)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int, p1 bool) uint64 {
    p1 = false
    v20 := GenS4{F0: 1, F1: (p0 & p0)}
    var v21 int16 = 32767
    if (-1 == 0) {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int32 = -2147483648
            if (!((v21 >> ((2) & 63)) < (v21 & 32767))) {
                var v24 float32 = (3.14159 * ((float32(float64((float32(float64(1.5)) + -2.25))) - float32(float64(0.1))) / (3.14159 / ((0.1 / 1.0) - ((0.0 / 1.0) - 1.5)))))
                var v25 int = p0
                fmt.Println(fi22, p1, fi22)
                _ = v24
                _ = v25
            } else {
                switch (^(fi22)) {
                case -2147483648:
                    v23 = 42
                case 4294967295:
                    v26 := []int{p0}
                    _ = v26
                case -1:
                    for fi27 := 0; fi27 < 15; fi27++ {
                        p0 += fi27
                        _ = fi27
                    }
                default:
                    if (!false) {
                        fmt.Println(p1, p1)
                    } else {
                        var v28 int64 = 9223372036854775807
                        _ = v28
                    }
                }
                fmt.Println(v23, v21, p0)
                if p1 {
                    switch p0 {
                    case 4294967295:
                        v29 := []int{(p0 - (fi22 / ((626) | 1))), fi22, p0}
                        _ = v29
                    case -1:
                        var v30 float64 = 0.1
                        _ = v30
                    case 32768:
                        if (p1 && false) {
                            if p1 {
                                v31 := []int{}
                                v23 -= (v23 % (((v23 | v23)) | 1))
                                var v32 uint = 255
                                _ = v31
                                _ = v32
                            } else {
                                var v33 int = fi22
                                _ = v33
                            }
                            var v34 bool = false
                            p0 = fi22
                            _ = v34
                        } else {
                            var v35 bool = p1
                            v21 = v21
                            _ = v35
                        }
                    }
                    p1 = true
                }
            }
            if true {
                if p1 {
                    p0 += (fi22 - (p0 & p0))
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    switch "a" {
                    case "":
                        p1 = p1
                    case "a":
                        var v37 uint64 = 0
                        _ = v37
                    case "foo bar":
                        v38 := GenS4{F0: 1, F1: 0}
                        _ = v38
                    }
                    if (p1 && ((!p1) || p1)) {
                        v21 = v21
                    } else {
                        v39 := GenS2{F0: 127, F1: 42, F2: -1}
                        v40 := []int{(p0 + -843)}
                        var v41 int16 = 32767
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    var v42 uint = 42
                    _ = fi36
                    _ = v42
                }
            } else {
                if (("foo bar" < ("foo bar" + ("hello" + "a"))) && p1) {
                    fmt.Println(p1)
                    v43 := []int{(fi22 & fi22)}
                    var v44 uint = 255
                    _ = v43
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        v46 := []int{(fi45 << ((18446744073709551615) & 63)), (p0 - fi22), (p0 + 127), fi22}
                        _ = fi45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 19; fi47++ {
                        switch fi47 {
                        case 32767:
                            v21 = -648
                        case 1:
                            var v48 uint64 = 174
                            _ = v48
                        case -655:
                            v23 = (v23 / (((v23 + 1)) | 1))
                        }
                        p1 = p1
                        _ = fi47
                    }
                }
                v49 := GenS2{F0: 2, F1: 0, F2: v21}
                _ = v49
            }
            _ = fi22
            _ = v23
        }
        for fi50 := 0; fi50 < 1; fi50++ {
            switch (^(fi50)) {
            case 4294967295:
                v51 := GenS4{F0: 42, F1: (fi50 << ((0) & 63))}
                _ = v51
            case -3:
                v52 := GenS5{F0: false, F1: (42 > 100)}
                _ = v52
            }
            fmt.Println(p1, p0, v21)
            _ = fi50
        }
        v53 := GenS5{F0: p1, F1: (65535 < 0)}
        _ = v53
    }
    if false {
        var v54 int8 = 100
        _ = v54
    } else {
        if (!(p1 && true)) {
            v55 := GenS6{F0: 54, F1: 100, F2: (v21 >> ((0) & 63))}
            if false {
                v56 := GenS3{F0: 1, F1: 4294967295, F2: p1}
                var v57 uint16 = 255
                if ("a" <= "\n\t") {
                    fmt.Println(v57, v57, p1)
                }
                _ = v56
                _ = v57
            } else {
                var v58 bool = p1
                _ = v58
            }
            if p1 {
                if (-1 == -1) {
                    v21 = (^(v21))
                    if (743 > (v21 & 42)) {
                        fmt.Println(p1, p1)
                    } else {
                        v21 -= (v21 / (((v21 | 0)) | 1))
                        if ((v21 / ((-1) | 1)) <= (v21 >> ((65535) & 63))) {
                            if p1 {
                                for fi59 := 0; fi59 < 8; fi59++ {
                                    if p1 {
                                        v60 := GenS1{F0: 140, F1: 59, F2: 0}
                                        var v61 uint64 = 772
                                        for fi62 := 0; fi62 < 19; fi62++ {
                                            fmt.Println(fi59)
                                            _ = fi62
                                        }
                                        _ = v60
                                        _ = v61
                                    }
                                    var v63 int64 = -1
                                    v63 = (^(v63))
                                    _ = fi59
                                    _ = v63
                                }
                                v21 = 32767
                                for fi64 := 0; fi64 < 14; fi64++ {
                                    for fi65 := 0; fi65 < 20; fi65++ {
                                        v66 := []int{127, (fi65 - -1), (fi65 & (fi64 >> ((255) & 63)))}
                                        _ = fi65
                                        _ = v66
                                    }
                                    _ = fi64
                                }
                            }
                            v67 := GenS7{F0: 42, F1: 0}
                            v68 := []int{p0, 0, (p0 ^ (p0 ^ 853))}
                            _ = v67
                            _ = v68
                        } else {
                            if p1 {
                                var v69 uint8 = 127
                                var v70 string = ""
                                var v71 float64 = ((1.5 + (1e10 * (-2.25 * 3.14159))) * float64(1e10))
                                _ = v69
                                _ = v70
                                _ = v71
                            } else {
                                p1 = p1
                                var v72 int8 = -81
                                for fi73 := 0; fi73 < 5; fi73++ {
                                    v72 = (v72 - (v72 << ((2) & 63)))
                                    _ = fi73
                                }
                                _ = v72
                            }
                        }
                        for fi74 := 0; fi74 < 2; fi74++ {
                            fmt.Println(fi74, p0, v21)
                            var v75 int32 = 2147483647
                            if (1 <= 0) {
                                v76 := GenS8{F0: 804, F1: "a", F2: 1}
                                _ = v76
                            }
                            _ = fi74
                            _ = v75
                        }
                    }
                } else {
                    var v77 float32 = (0.1 * 1e10)
                    p1 = (!p1)
                    var v78 string = "foo bar"
                    _ = v77
                    _ = v78
                }
                switch (p0 - p0) {
                case 127:
                    fmt.Println(v21)
                default:
                    if false {
                        v21 = 0
                        v79 := []int{p0}
                        for fi80 := 0; fi80 < 6; fi80++ {
                            v81 := GenS9{F0: 128}
                            switch ("hello" + ("" + (("foo bar" + "a") + (("" + ((((("hello" + "foo bar") + "\n\t") + "") + ("\n\t" + "\n\t")) + "hello")) + ("\n\t" + "\n\t"))))) {
                            case "":
                                p0 = (p0 * 42)
                            default:
                                var v82 float64 = 3.14159
                                _ = v82
                            }
                            _ = fi80
                            _ = v81
                        }
                        _ = v79
                    } else {
                        v83 := []int{-1}
                        fmt.Println(p1, p0, p1)
                        p1 = false
                        _ = v83
                    }
                }
                var v84 float32 = -2.25
                _ = v84
            } else {
                if false {
                    var v85 uint32 = 2147483647
                    _ = v85
                }
            }
            _ = v55
        }
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 945
}

func main() {
    var v1 float32 = 0.1
    var v2 int = 514
    v3 := []int{(^(v2))}
    var v4 float32 = float32(0.1)
    switch ("a" + "") {
    case "hello":
        for fi86 := 0; fi86 < 1; fi86++ {
            var v87 int = -91
            _ = fi86
            _ = v87
        }
    case "foo bar":
        v1 = v4
    case "\n\t":
        v88 := GenS10{F0: (v4 < 3.14159), F1: (v2 ^ (-(v2)))}
        _ = v88
    default:
        switch (v2 * (v2 << ((2) & 63))) {
        case 9223372036854775807:
            v89 := GenS11{F0: 309, F1: 1, F2: -1}
            _ = v89
        case -223:
            var v90 bool = false
            _ = v90
        case 128:
            _ = len(v3)
        }
    }
    var v91 uint16 = 65535
    fmt.Println(v2, v4, v1)
    var v92 uint8 = 169
    if true {
        v93 := GenS11{F0: 18446744073709551615, F1: v92, F2: -456}
        _ = v93
    } else {
        if (v92 >= v92) {
            for fi94 := 0; fi94 < 7; fi94++ {
                v95 := GenS12{F0: true, F1: 2147483647, F2: 570}
                _ = fi94
                _ = v95
            }
            switch -1 {
            case -9223372036854775808:
                var v96 float32 = v4
                _ = v96
            case 9223372036854775807:
                v2 += 127
            default:
                var v97 uint16 = (v91 - (v91 - (v91 * 2)))
                _ = v97
            }
        } else {
            var v98 float64 = float64(v1)
            fmt.Println(v1)
            if true {
                v4 = (0.1 / (v4 + 3.14159))
            } else {
                switch ((("" + "hello") + ("hello" + (("hello" + ("hello" + ("hello" + "\n\t"))) + "foo bar"))) + ("foo bar" + "foo bar")) {
                case "hello":
                    for fi99 := 0; fi99 < 19; fi99++ {
                        v91 += v91
                        _ = fi99
                    }
                case "a":
                    for fi100 := 0; fi100 < 16; fi100++ {
                        fmt.Println(v92, v98, v1)
                        v2 = 127
                        _ = fi100
                    }
                case "":
                    if false {
                        for fi101 := 0; fi101 < 18; fi101++ {
                            if (false && false) {
                                if (v1 <= v4) {
                                    switch (("" + (("" + "\n\t") + ("a" + ("a" + ("" + "\n\t"))))) + ("\n\t" + (((("\n\t" + (("hello" + "\n\t") + "hello")) + ((("foo bar" + "a") + "\n\t") + "foo bar")) + (((("\n\t" + "foo bar") + "a") + "a") + "hello")) + "hello"))) {
                                    case "":
                                        fmt.Println(v2, v1)
                                    case "foo bar":
                                        var v102 int = fi101
                                        _ = v102
                                    }
                                }
                                for fi103 := 0; fi103 < 14; fi103++ {
                                    var v104 uint16 = (^(v91))
                                    if (genFn1((fi101 / (((fi103 << ((42) & 63))) | 1)), false) != genFn1(fi101, (!(460 == 828)))) {
                                        if len(v3) > 0 {
                                            v3[0] = fi101
                                        }
                                    } else {
                                        if (1 <= 393) {
                                            v105 := GenS8{F0: 4294967295, F1: ("foo bar" + "a"), F2: 0}
                                            _ = v105
                                        } else {
                                            var v106 float32 = v4
                                            _ = v106
                                        }
                                        v107 := GenS5{F0: false, F1: false}
                                        if true {
                                            if true {
                                                var v108 uint32 = 260
                                                _ = v108
                                            } else {
                                                fmt.Println(v92)
                                                v92 = v92
                                                switch "foo bar" {
                                                case "a":
                                                    fmt.Println(v104, v2, v4)
                                                case "\n\t":
                                                    fmt.Println(v104)
                                                default:
                                                    v98 = v98
                                                }
                                            }
                                            v3 = append(v3, -9223372036854775808)
                                            var v109 int32 = -89
                                            _ = v109
                                        } else {
                                            v110 := GenS13{F0: genFn1(-9223372036854775808, true), F1: 876}
                                            var v111 float32 = 1.5
                                            var v112 int64 = 0
                                            var v113 uint8 = 0
                                            _ = v110
                                            _ = v111
                                            _ = v112
                                            _ = v113
                                        }
                                        _ = v107
                                    }
                                    v4 = ((float32(0.0) / 1.0) / 0.1)
                                    _ = fi103
                                    _ = v104
                                }
                                for fi114 := 0; fi114 < 4; fi114++ {
                                    fmt.Println(v98)
                                    v3 = append(v3, (fi101 - fi101))
                                    v3 = append(v3, -2147483648)
                                    _ = fi114
                                }
                            }
                            var v115 float32 = v4
                            v116 := GenS4{F0: 42, F1: fi101}
                            _ = fi101
                            _ = v115
                            _ = v116
                        }
                        fmt.Println(v1, v4, v2)
                        v3 = append(v3, (v2 ^ -2147483648))
                    } else {
                        switch v2 {
                        case 128:
                            if false {
                                if (!((1.5 * (v1 / 1.0)) >= (v1 / 1.0))) {
                                    v4 -= (3.14159 - 3.14159)
                                    if false {
                                        var v117 uint32 = 65535
                                        _ = v117
                                    }
                                }
                                _ = len(v3)
                                v2 = (v2 / ((9223372036854775807) | 1))
                            } else {
                                v118 := GenS12{F0: false, F1: 2147483647, F2: 42}
                                fmt.Println(v1, v4, v4)
                                _ = v118
                            }
                        case -237:
                            if false {
                                switch "\n\t" {
                                case "":
                                    v119 := GenS9{F0: (v92 % ((162) | 1))}
                                    _ = v119
                                case "hello":
                                    var v120 uint16 = 32767
                                    _ = v120
                                default:
                                    v3 = append(v3, (v2 << ((2) & 63)))
                                }
                                if true {
                                    if false {
                                        fmt.Println(v2, v2, v92)
                                        var v121 string = "a"
                                        v122 := GenS7{F0: -1, F1: 4294967295}
                                        _ = v121
                                        _ = v122
                                    } else {
                                        v92 = 174
                                    }
                                }
                                v91 += (v91 & v91)
                            }
                        case 32768:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            var v123 float64 = 3.14159
                            _ = v123
                        }
                    }
                default:
                    for fi124 := 0; fi124 < 7; fi124++ {
                        switch (v2 * v2) {
                        case 32768:
                            switch (fi124 | (fi124 / ((int(v91)) | 1))) {
                            case 32767:
                                var v125 uint16 = (v91 ^ v91)
                                _ = v125
                            case -2147483648:
                                if true {
                                    var v126 uint = 0
                                    fmt.Println(v92, v4)
                                    _ = v126
                                }
                            case 607:
                                if false {
                                    _ = len(v3)
                                    for fi127 := 0; fi127 < 5; fi127++ {
                                        var v128 int32 = -1
                                        switch -9223372036854775808 {
                                        case 2147483647:
                                            v92 = v92
                                        case 42:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            var v129 uint8 = 233
                                            _ = v129
                                        }
                                        v130 := GenS14{F0: v128, F1: -128}
                                        if false {
                                            var v131 int8 = 127
                                            if true {
                                                var v132 int8 = v131
                                                if false {
                                                    v133 := GenS6{F0: (v132 ^ 127), F1: (-(v132)), F2: 32767}
                                                    v3 = append(v3, (fi124 >> ((451) & 63)))
                                                    _ = v133
                                                }
                                                var v134 int8 = v132
                                                var v135 uint8 = uint8(v91)
                                                fmt.Println(v131)
                                                _ = v132
                                                _ = v134
                                                _ = v135
                                            }
                                            _ = v131
                                        }
                                        _ = fi127
                                        _ = v128
                                        _ = v130
                                    }
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case -2147483648:
                            switch int(v92) {
                            case 9223372036854775807:
                                v4 += v1
                            case 1:
                                var v136 int64 = -89
                                _ = v136
                            default:
                                fmt.Println(v91)
                            }
                        default:
                            v2 = -2147483648
                        }
                        _ = fi124
                    }
                }
                var v137 int32 = -1
                v92 -= 1
                _ = v137
            }
            _ = v98
        }
    }
    for fi138 := 0; fi138 < 7; fi138++ {
        var v139 uint8 = 50
        v3 = append(v3, 9223372036854775807)
        _ = fi138
        _ = v139
    }
    var v140 uint64 = 1
    v91 -= v91
    v141 := GenS12{F0: (0 >= -909), F1: 0, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v91
    _ = v92
}
