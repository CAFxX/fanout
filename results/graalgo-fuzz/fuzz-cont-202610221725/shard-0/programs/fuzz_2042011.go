package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int32; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 uint16; F2 int16 }
type GenS13 struct { F0 uint32; F1 bool; F2 int64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 string; F1 int; F2 uint32 }
type GenS17 struct { F0 bool; F1 int8 }
func genFn0(p0 string, p1 bool, p2 string) uint8 {
    var v7 uint8 = 109
    var v8 uint16 = 0
    switch (p0 + "\n\t") {
    case "foo bar":
        if (!((-2.25 * 1e10) != -2.25)) {
            for fi9 := 0; fi9 < 13; fi9++ {
                for fi10 := 0; fi10 < 19; fi10++ {
                    if false {
                        v11 := []int{}
                        var v12 int8 = 100
                        switch "hello" {
                        case "\n\t":
                            switch p2 {
                            case "\n\t":
                                fmt.Println(fi9, p1, fi9)
                            }
                            fmt.Println(v12, fi9)
                            v13 := GenS1{F0: (v7 - (v7 - 1))}
                            _ = len(v11)
                            p0 = ""
                            fmt.Println(fi10, v12, p0)
                            var v14 int8 = int8(v8)
                            v12 = -128
                            _ = v13
                            _ = v14
                        case "":
                            v15 := GenS2{F0: 15}
                            _ = v15
                        case "a":
                            switch ("hello" + (p2 + ("" + ("foo bar" + p0)))) {
                            case "foo bar":
                                v12 = v12
                            }
                        }
                        _ = v11
                        _ = v12
                    } else {
                        var v16 int16 = 372
                        v16 = v16
                        _ = v16
                    }
                    p2 = ("\n\t" + p0)
                    _ = fi10
                }
                _ = fi9
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := []int{879}
                var v19 float32 = (1e10 + 0.1)
                var v20 string = p2
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        } else {
            var v21 int32 = -508
            v22 := GenS2{F0: 1}
            fmt.Println(p2)
            _ = v21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 158
}
func genFn1() bool {
    var v23 uint32 = 439
    v24 := []int{1, 4294967295}
    var v25 uint16 = 0
    if false {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int32 = 0
            var v28 bool = false
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            if false {
                v30 := GenS3{F0: -1, F1: 1}
                if true {
                    v31 := GenS1{F0: 128}
                    fmt.Println(v23, v25, v25)
                    _ = v31
                } else {
                    if false {
                        fmt.Println(v23, fi29)
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v23, v23, v25)
                        case "a":
                            fmt.Println(v23, fi29, v23)
                        default:
                            fmt.Println(fi29, v23, fi29)
                        }
                    } else {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if (159 < uint(v25)) {
                                v33 := GenS2{F0: 2147483647}
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                if (-1 != 100) {
                                    if false {
                                        fmt.Println(v25, v23)
                                    } else {
                                        _ = len(v24)
                                        if false {
                                            if len(v24) > 0 {
                                                v24[0] = (fi32 | fi32)
                                            }
                                            fmt.Println(fi29)
                                            v25 = (v25 & v25)
                                        }
                                    }
                                    fmt.Println(v25, v25)
                                }
                                _ = v33
                            }
                            if (((v25 >= 255) && true) || true) {
                                v34 := GenS4{F0: ""}
                                var v35 float32 = 0.1
                                var v36 int32 = 227
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            switch (fi29 & fi32) {
                            case -429:
                                _ = len(v24)
                            default:
                                var v37 int32 = int32(v25)
                                _ = v37
                            }
                            var v38 bool = true
                            v23 = v23
                            _ = fi32
                            _ = v38
                        }
                        var v39 bool = (13 < 255)
                        _ = v39
                    }
                    var v40 uint8 = uint8(v25)
                    _ = v40
                }
                _ = v30
            } else {
                var v41 string = "\n\t"
                var v42 uint16 = uint16(v23)
                _ = v41
                _ = v42
            }
            _ = fi29
        }
    }
    var v43 int16 = 42
    _ = v23
    _ = v24
    _ = v25
    _ = v43
    return (1.5 < 3.14159)
}

func main() {
    var v1 uint16 = 833
    var v2 int = -9223372036854775808
    var v3 uint16 = (v1 % ((v1) | 1))
    v4 := []uint32{2147483647}
    var v5 uint64 = 926
    v6 := []string{"hello", "a", ("a" + ("" + "a"))}
    if true {
        v44 := GenS5{F0: 42, F1: 0, F2: (v1 * v1)}
        fmt.Println(v2, v5)
        _ = v44
    }
    switch (("a" + ("\n\t" + (("\n\t" + ("a" + "foo bar")) + "hello"))) + "hello") {
    case "foo bar":
        v2 -= v2
    case "":
        v5 -= v5
    case "hello":
        v2 = (-(v2))
    }
    v3 = (v3 << ((65535) & 63))
    var v45 uint64 = 178
    var v46 float32 = (1.5 / (3.14159 - 1.5))
    if true {
        for fi47 := 0; fi47 < 8; fi47++ {
            v46 = v46
            var v48 uint32 = 2
            _ = fi47
            _ = v48
        }
        v45 -= v45
        for fi49 := 0; fi49 < 10; fi49++ {
            for fi50 := 0; fi50 < 5; fi50++ {
                var v51 uint64 = v5
                fmt.Println(v3, fi49)
                var v52 int32 = 1
                _ = fi50
                _ = v51
                _ = v52
            }
            v3 = (v3 >> ((255) & 63))
            switch v2 {
            case 0:
                if (v2 < fi49) {
                    var v53 int16 = 42
                    _ = v53
                } else {
                    v54 := GenS3{F0: 1, F1: 0}
                    _ = v54
                }
            case 312:
                v45 = 42
            default:
                if false {
                    switch 32768 {
                    case 9223372036854775807:
                        switch 2147483647 {
                        case 4294967295:
                            v46 = v46
                        case -975:
                            if len(v6) > 0 {
                                v6[0] = ""
                            }
                        case -117:
                            if false {
                                var v55 uint8 = 127
                                _ = v55
                            }
                        default:
                            v56 := GenS4{F0: "foo bar"}
                            _ = v56
                        }
                    default:
                        if false {
                            var v57 float32 = v46
                            _ = v57
                        } else {
                            v58 := GenS1{F0: genFn0("a", true, "a")}
                            fmt.Println(v46, v45, v5)
                            v59 := GenS5{F0: 1, F1: 0, F2: (v1 >> ((689) & 63))}
                            _ = v58
                            _ = v59
                        }
                    }
                    for fi60 := 0; fi60 < 9; fi60++ {
                        switch "foo bar" {
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v1)
                        default:
                            var v61 float64 = (0.0 + 1.5)
                            _ = v61
                        }
                        v62 := GenS1{F0: 42}
                        fmt.Println(v2, v5, v46)
                        _ = fi60
                        _ = v62
                    }
                    var v63 int = 128
                    _ = v63
                } else {
                    v64 := GenS6{F0: 127, F1: 9223372036854775807}
                    if (((true || (801 > -928)) && false) || (!(true && (114 < 222)))) {
                        if ((!false) && true) {
                            var v65 bool = false
                            fmt.Println(fi49)
                            fmt.Println(v1, v45, v65)
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 8; fi66++ {
                            v5 = (v45 + 0)
                            _ = fi66
                        }
                    }
                    _ = v64
                }
            }
            _ = fi49
        }
    } else {
        if (("hello" + "hello") >= "\n\t") {
            v3 = v3
        }
        if (false && true) {
            switch "foo bar" {
            case "a":
                fmt.Println(v3, v5)
            }
            if false {
                v46 -= v46
                switch ("foo bar" + ("" + (("\n\t" + "a") + "foo bar"))) {
                case "":
                    v67 := GenS7{F0: 1, F1: 0, F2: (v5 | (v45 | 687))}
                    _ = v67
                case "\n\t":
                    v5 = v5
                case "hello":
                    v46 -= v46
                default:
                    v1 += (^(v1))
                }
                v5 -= (v45 % ((0) | 1))
            }
            v45 = 4294967295
        }
        v3 -= 32767
    }
    v2 -= (^(v2))
    var v68 int = (v2 & v2)
    fmt.Println(v68, v68)
    fmt.Println(v1)
    fmt.Println(v45, v46, v68)
    v68 = v2
    v5 = v45
    switch (v2 ^ (v2 | 487)) {
    case 32768:
        v3 = v1
    case 42:
        fmt.Println(v1)
    case 127:
        for fi69 := 0; fi69 < 6; fi69++ {
            if false {
                if (-128 <= -1) {
                    switch 127 {
                    case -807:
                        for fi70 := 0; fi70 < 5; fi70++ {
                            v71 := GenS4{F0: ""}
                            v72 := GenS8{F0: 32767, F1: (^(v3))}
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                    default:
                        if (!true) {
                            v73 := GenS7{F0: 244, F1: 431, F2: (v5 / ((v5) | 1))}
                            _ = v73
                        } else {
                            var v74 int = fi69
                            switch (fi69 & -1) {
                            case 390:
                                var v75 int8 = 0
                                _ = v75
                            case 127:
                                var v76 int8 = -128
                                _ = v76
                            case 2147483647:
                                for fi77 := 0; fi77 < 9; fi77++ {
                                    v1 = v3
                                    var v78 int16 = -32768
                                    var v79 int = fi77
                                    var v80 int64 = int64(v45)
                                    if (((v46 * v46) * v46) >= v46) {
                                        switch ((("foo bar" + "a") + "") + "hello") {
                                        case "":
                                            if (65535 < 255) {
                                                fmt.Println(v3)
                                                if false {
                                                    v81 := GenS3{F0: -1, F1: v78}
                                                    if true {
                                                        var v82 bool = true
                                                        _ = v82
                                                    }
                                                    _ = v81
                                                } else {
                                                    v1 = (v3 + v3)
                                                    var v83 int = v79
                                                    var v84 int16 = int16(v45)
                                                    _ = v83
                                                    _ = v84
                                                }
                                                var v85 int16 = (-(v78))
                                                _ = v85
                                            } else {
                                                var v86 int = 32767
                                                switch v68 {
                                                case 42:
                                                    fmt.Println(fi77)
                                                case 2147483647:
                                                    if len(v6) > 0 {
                                                        v6[0] = "foo bar"
                                                    }
                                                default:
                                                    var v87 uint16 = v1
                                                    _ = v87
                                                }
                                                v74 -= -9223372036854775808
                                                _ = v86
                                            }
                                        case "\n\t":
                                            var v88 uint64 = v45
                                            _ = v88
                                        case "foo bar":
                                            switch ("foo bar" + "a") {
                                            case "a":
                                                v89 := GenS5{F0: -116, F1: -273, F2: (v1 % ((v1) | 1))}
                                                _ = v89
                                            case "\n\t":
                                                v90 := GenS9{F0: v45, F1: (fi69 ^ v79)}
                                                _ = v90
                                            case "foo bar":
                                                if ((v46 * float32(1e10)) > 1.5) {
                                                    v46 = (v46 - v46)
                                                    switch ("hello" + ("hello" + "foo bar")) {
                                                    case "hello":
                                                        if (v45 == v45) {
                                                            var v91 uint16 = (v3 % (((^(v3))) | 1))
                                                            _ = len(v4)
                                                            var v92 uint8 = 127
                                                            if len(v6) > 0 {
                                                                v6[0] = "hello"
                                                            }
                                                            _ = v91
                                                            _ = v92
                                                        } else {
                                                            if false {
                                                                var v93 int16 = (^(v78))
                                                                if len(v4) > 0 {
                                                                    v4[0] = 865
                                                                }
                                                                if true {
                                                                    var v94 bool = (float64(-2.25) < ((((-2.25 - ((((0.1 + ((0.1 + 0.0) / (0.0 + float64(((((v46 / v46) * v46) + (((-2.25 - (0.0 + v46)) + v46) + 0.0)) + v46))))) * (float64((v46 / v46)) - 1.5)) / 1e10) / (1e10 + (3.14159 / -2.25)))) * ((1.5 + 1.5) * 1e10)) * 1.5) - 1e10))
                                                                    var v95 int32 = 2147483647
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                v93 -= v78
                                                                _ = v93
                                                            } else {
                                                                v68 += v68
                                                            }
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v79)
                                                    case "a":
                                                        var v96 int = 9223372036854775807
                                                        v4 = append(v4, 2)
                                                        v97 := GenS10{F0: fi69, F1: 4294967295}
                                                        if true {
                                                            var v98 uint = 468
                                                            if (!true) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = ("foo bar" + "hello")
                                                                }
                                                                var v99 bool = (false && true)
                                                                var v100 int8 = -1
                                                                _ = v100
                                                                _ = v99
                                                            } else {
                                                                var v101 float64 = 3.14159
                                                                v79 = 42
                                                                _ = v101
                                                            }
                                                            _ = v98
                                                        }
                                                        if false {
                                                            if false {
                                                                v5 -= (v45 - 420)
                                                                var v102 uint16 = 2
                                                                _ = v102
                                                            }
                                                            var v103 float64 = 3.14159
                                                            if true {
                                                                _ = len(v6)
                                                                if true {
                                                                    v104 := GenS11{F0: (^(v96))}
                                                                    v105 := GenS12{F0: int(v78), F1: v3, F2: -32768}
                                                                    _ = v104
                                                                    _ = v105
                                                                }
                                                                v6 = append(v6, ("a" + (("" + ((("hello" + ("" + ("hello" + ""))) + "a") + "foo bar")) + ("\n\t" + "foo bar"))))
                                                            } else {
                                                                var v106 int8 = -128
                                                                switch (v96 + int(v5)) {
                                                                case 32767:
                                                                    v107 := GenS10{F0: int(v80), F1: 1}
                                                                    _ = v107
                                                                case 9223372036854775807:
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v108 float32 = v46
                                                                        _ = v108
                                                                    case "\n\t":
                                                                        var v109 int8 = v106
                                                                        _ = v109
                                                                    case "foo bar":
                                                                        if (((v78 <= (v78 << ((18446744073709551615) & 63))) || (v5 > v5)) || (v78 <= v78)) {
                                                                            v45 += (v45 | 0)
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v110 uint16 = (v3 | v3)
                                                                                _ = v110
                                                                            case "a":
                                                                                var v111 int = 32768
                                                                                v112 := GenS7{F0: 42, F1: 255, F2: (v45 | v45)}
                                                                                _ = v111
                                                                                _ = v112
                                                                            default:
                                                                                v6 = append(v6, "a")
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 4294967295
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    var v113 string = "foo bar"
                                                                    _ = v113
                                                                }
                                                                fmt.Println(v79)
                                                                _ = v106
                                                            }
                                                            _ = v103
                                                        } else {
                                                            var v114 int32 = 0
                                                            _ = v114
                                                        }
                                                        var v115 string = ("" + "foo bar")
                                                        _ = v115
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    if true {
                                                        v116 := GenS11{F0: 127}
                                                        fmt.Println(fi77, v78, v3)
                                                        _ = v116
                                                    }
                                                } else {
                                                    var v117 uint = 42
                                                    _ = v117
                                                }
                                            }
                                        }
                                    } else {
                                        var v118 int8 = 42
                                        var v119 bool = true
                                        _ = v118
                                        _ = v119
                                    }
                                    if (true || true) {
                                        _ = len(v4)
                                        fmt.Println(v79, v68)
                                        if true {
                                            var v120 uint16 = 796
                                            v5 -= 18446744073709551615
                                            var v121 uint16 = 110
                                            _ = v120
                                            _ = v121
                                        } else {
                                            if true {
                                                v122 := GenS2{F0: 42}
                                                _ = v122
                                            } else {
                                                v2 += v79
                                                v123 := GenS13{F0: 365, F1: (255 > 2), F2: int64(v45)}
                                                _ = v123
                                            }
                                            v124 := GenS2{F0: 0}
                                            _ = v124
                                        }
                                    } else {
                                        v74 -= v68
                                        v6 = append(v6, (("a" + "\n\t") + ("\n\t" + ("hello" + ("foo bar" + "\n\t")))))
                                        if true {
                                            _ = len(v6)
                                            v125 := GenS14{F0: 172}
                                            _ = len(v6)
                                            var v126 int = v79
                                            v78 = v78
                                            _ = v125
                                            _ = v126
                                        }
                                    }
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                            default:
                                v74 -= -709
                            }
                            var v127 string = "foo bar"
                            _ = v127
                            _ = v74
                        }
                    }
                    if ((!true) && ("\n\t" <= "")) {
                        var v128 uint64 = uint64(v1)
                        _ = v128
                    } else {
                        v4 = append(v4, 148)
                        for fi129 := 0; fi129 < 12; fi129++ {
                            if true {
                                fmt.Println(v68, fi69)
                                v6 = append(v6, "foo bar")
                                v2 = (^(v68))
                                switch (fi69 & (v2 << ((18446744073709551615) & 63))) {
                                case 9223372036854775807:
                                    fmt.Println(v1, v68)
                                case 0:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    var v130 float32 = v46
                                    _ = v130
                                }
                            }
                            fmt.Println(fi69, fi69)
                            _ = fi129
                        }
                    }
                    for fi131 := 0; fi131 < 3; fi131++ {
                        v6 = append(v6, "hello")
                        var v132 int32 = 42
                        _ = fi131
                        _ = v132
                    }
                }
            }
            _ = fi69
        }
    default:
        v46 = (v46 - (v46 - v46))
    }
    if (-1 < (v2 ^ v68)) {
        for fi133 := 0; fi133 < 13; fi133++ {
            var v134 int16 = -1
            switch fi133 {
            case -9223372036854775808:
                var v135 bool = (119 == -82)
                _ = v135
            }
            _ = fi133
            _ = v134
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if ((v45 & 2) > (v5 + v5)) {
            if len(v6) > 0 {
                v6[0] = (((("\n\t" + ("\n\t" + "")) + "foo bar") + "") + "hello")
            }
        }
        if true {
            v136 := GenS15{F0: 941}
            for fi137 := 0; fi137 < 15; fi137++ {
                for fi138 := 0; fi138 < 4; fi138++ {
                    if true {
                        if len(v4) > 0 {
                            v4[0] = uint32(v45)
                        }
                    } else {
                        if (v3 != 2) {
                            fmt.Println(v3, v46, fi137)
                            if ((false && true) || (!false)) {
                                fmt.Println(fi137, fi138)
                                fmt.Println(v46, fi137, v2)
                                if (255 != 128) {
                                    switch 9223372036854775807 {
                                    case -1:
                                        if len(v6) > 0 {
                                            v6[0] = (("\n\t" + "\n\t") + "foo bar")
                                        }
                                    case -884:
                                        v139 := GenS11{F0: (v2 << ((2) & 63))}
                                        _ = v139
                                    default:
                                        if (((true && false) || (0 < -32768)) && true) {
                                            var v140 int16 = -395
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v141 float32 = -2.25
                                            var v142 uint32 = 42
                                            _ = v140
                                            _ = v141
                                            _ = v142
                                        }
                                    }
                                    v45 = 1
                                    if len(v6) > 0 {
                                        v6[0] = ("a" + ("hello" + "a"))
                                    }
                                } else {
                                    v68 += (-(v68))
                                }
                                v5 -= uint64(v3)
                                v143 := GenS3{F0: 42, F1: 1}
                                _ = v143
                            } else {
                                v144 := GenS12{F0: -9223372036854775808, F1: (v1 % ((uint16(v5)) | 1)), F2: 1000}
                                if false {
                                    v145 := GenS16{F0: "hello", F1: (v2 >> ((65535) & 63)), F2: 817}
                                    _ = v145
                                }
                                v68 -= fi137
                                v5 -= 83
                                v68 = (v68 | v2)
                                _ = v144
                            }
                            var v146 int8 = -1
                            _ = v146
                        } else {
                            v5 = v5
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = ("\n\t" + ("hello" + "foo bar"))
                                }
                            } else {
                                v147 := GenS2{F0: 2147483647}
                                _ = v147
                            }
                            var v148 uint8 = 1
                            _ = v148
                        }
                        v4 = append(v4, 627)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v3 += 255
                    var v149 float32 = (v46 - 1e10)
                    fmt.Println(v46, fi138, fi137)
                    fmt.Println(v149)
                    _ = fi138
                    _ = v149
                }
                _ = fi137
            }
            v150 := GenS17{F0: (v5 < v5), F1: 100}
            _ = v136
            _ = v150
        } else {
            v151 := GenS4{F0: ""}
            v46 -= float32(float64(v46))
            _ = v151
        }
    }
    v68 = v2
    switch 128 {
    case 32767:
        v2 += 127
    case -9223372036854775808:
        if true {
            v68 += (v2 | (v68 * (v68 >> ((2) & 63))))
        } else {
            _ = len(v6)
        }
    case 1:
        var v152 uint32 = 65535
        _ = v152
    }
    if false {
        for fi153 := 0; fi153 < 14; fi153++ {
            v154 := GenS7{F0: 2147483647, F1: 2, F2: 753}
            if false {
                var v155 uint32 = 12
                _ = v155
            } else {
                var v156 int16 = 367
                _ = v156
            }
            _ = fi153
            _ = v154
        }
        for fi157 := 0; fi157 < 14; fi157++ {
            v5 += (v5 << ((255) & 63))
            _ = fi157
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v68
}
