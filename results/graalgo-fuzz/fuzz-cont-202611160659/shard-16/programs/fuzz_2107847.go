package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint64; F2 int8 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint8; F1 uint32; F2 uint }
type GenS12 struct { F0 uint8; F1 bool; F2 uint }
type GenS13 struct { F0 uint; F1 bool }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int8; F1 uint32; F2 uint }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int; F1 uint8 }
func genFn0() int64 {
    if true {
        var v5 float64 = 0.1
        _ = v5
    }
    var v6 uint = 65535
    switch -218 {
    case 32768:
        v7 := []int{57, -138, -915, 42}
        _ = v7
    default:
        v6 = 288
    }
    var v8 uint64 = 329
    _ = v6
    _ = v8
    return -316
}
func genFn1(p0 uint8, p1 uint64, p2 int16) uint64 {
    switch (("" + "foo bar") + "") {
    case "\n\t":
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int8 = -41
                switch "\n\t" {
                case "foo bar":
                    if (-2.25 <= 1e10) {
                        var v12 uint32 = 42
                        _ = v12
                    }
                case "\n\t":
                    switch (fi10 | fi9) {
                    case 14:
                        v13 := GenS1{F0: 526}
                        _ = v13
                    case -2147483648:
                        v11 -= (v11 * (v11 >> ((4294967295) & 63)))
                    }
                default:
                    var v14 uint16 = 32767
                    _ = v14
                }
                v15 := GenS2{F0: (fi9 * (fi10 * (^(fi9))))}
                if (65535 >= 634) {
                    v16 := GenS3{F0: (p1 * p1), F1: (^(v11))}
                    _ = v16
                }
                var v17 uint = 0
                _ = fi10
                _ = v11
                _ = v15
                _ = v17
            }
            switch (fi9 * (fi9 / (((fi9 + fi9)) | 1))) {
            case 32768:
                v18 := GenS4{F0: genFn0()}
                _ = v18
            }
            p1 = (p1 ^ (p1 % ((42) | 1)))
            _ = fi9
        }
    case "":
        var v19 int32 = -2147483648
        _ = v19
    case "hello":
        v20 := GenS1{F0: 2147483647}
        _ = v20
    default:
        var v21 bool = ((true && true) || (float32(1.5) < float32(((0.1 - (float64(1e10) + 0.1)) - (0.1 * 1e10)))))
        _ = v21
    }
    var v22 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return p1
}

func main() {
    v1 := []string{"a", (((("hello" + "") + "a") + "\n\t") + ("a" + ("a" + "foo bar"))), "\n\t", ("a" + "")}
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v23 uint32 = 1
    delete(v3, -170)
    for fi24 := 0; fi24 < 16; fi24++ {
        switch "a" {
        case "a":
            var v25 int = 4294967295
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 5; fi26++ {
                v27 := GenS5{F0: 1}
                var v28 uint32 = 40
                v29 := GenS4{F0: genFn0()}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
        default:
            v30 := GenS6{F0: (42 > 1), F1: 400, F2: 1}
            _ = v30
        }
        v31 := GenS3{F0: 0, F1: 0}
        v32 := GenS7{F0: -1, F1: 7}
        _ = fi24
        _ = v31
        _ = v32
    }
    switch 848 {
    case -405:
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 bool = (1e10 < float32((3.14159 / 1.0)))
            var v35 string = "foo bar"
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v23 += v23
    for fi36 := 0; fi36 < 12; fi36++ {
        v23 -= 4294967295
        _ = fi36
    }
    if (13 == -128) {
        fmt.Println(v2, v2)
        v1 = append(v1, "")
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := GenS8{F0: 2, F1: false}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 string = "hello"
        _ = fi39
        _ = v40
    }
    if false {
        var v41 float32 = 1.5
        _ = v41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        fmt.Println(v23, v2, v2)
        v1 = append(v1, ("foo bar" + "a"))
        for fi43 := 0; fi43 < 19; fi43++ {
            fmt.Println(fi42, v2)
            v2 = 3.14159
            _ = fi43
        }
        _ = fi42
    }
    switch -7 {
    case 653:
        for fi44 := 0; fi44 < 19; fi44++ {
            if (1 >= 164) {
                if ((0.0 * 1.5) <= float64(((0.0 * v2) / v2))) {
                    switch "a" {
                    case "a":
                        switch (fi44 % (((fi44 | -1)) | 1)) {
                        case 460:
                            fmt.Println(fi44, v23)
                        case -2147483648:
                            v23 = (v23 >> ((4294967295) & 63))
                        case 32768:
                            for fi45 := 0; fi45 < 8; fi45++ {
                                v23 = v23
                                var v46 int = -462
                                _ = fi45
                                _ = v46
                            }
                        default:
                            v47 := v4[("" + "foo bar")]
                            _ = v47
                        }
                    case "hello":
                        for fi48 := 0; fi48 < 14; fi48++ {
                            if (1 != 32767) {
                                fmt.Println(fi48, fi44, fi48)
                                v2 -= 0.1
                                v23 -= v23
                            }
                            if ((!true) && (0.0 < 3.14159)) {
                                if len(v1) > 0 {
                                    v1[0] = (("foo bar" + ("a" + "")) + ("a" + "hello"))
                                }
                                v1 = append(v1, "hello")
                            } else {
                                var v49 uint8 = 2
                                v50 := GenS7{F0: 0, F1: 1}
                                v51 := GenS9{F0: 2}
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            if true {
                                var v52 int32 = 374
                                _ = v52
                            }
                            _ = fi48
                        }
                    case "":
                        fmt.Println(fi44, fi44, v23)
                    }
                    v23 -= 2
                    var v53 int = (fi44 | 9223372036854775807)
                    _ = v53
                } else {
                    v3[32767] = true
                }
                var v54 int32 = -2147483648
                _ = v54
            }
            v2 = v2
            _ = fi44
        }
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        v56 := GenS6{F0: false, F1: 2, F2: -66}
        if false {
            v23 = v23
            if true {
                for fi57 := 0; fi57 < 9; fi57++ {
                    _ = len(v1)
                    v58 := GenS9{F0: 42}
                    if (624 <= 258) {
                        if false {
                            var v59 int16 = 0
                            v4["\n\t"] = "\n\t"
                            _ = v59
                        } else {
                            var v60 float32 = 1.5
                            if ((0.0 * v60) >= (v2 * ((v2 - v2) * 0.0))) {
                                var v61 int = fi57
                                var v62 int16 = 1000
                                var v63 int64 = -191
                                _ = v61
                                _ = v62
                                _ = v63
                            } else {
                                var v64 string = "hello"
                                _ = v64
                            }
                            var v65 uint16 = 42
                            if true {
                                if true {
                                    v60 = float32(1e10)
                                    var v66 int8 = -106
                                    _ = v66
                                } else {
                                    v65 = (v65 ^ (v65 / ((v65) | 1)))
                                    v67 := GenS6{F0: true, F1: genFn1(127, 225, 32767), F2: 100}
                                    _ = v67
                                }
                            } else {
                                var v68 int8 = 62
                                if (v23 > 994) {
                                    var v69 uint16 = (^(v65))
                                    switch ((("foo bar" + "hello") + "a") + "hello") {
                                    case "\n\t":
                                        if true {
                                            if false {
                                                switch "a" {
                                                case "\n\t":
                                                    v2 = 1e10
                                                    fmt.Println(v2, fi55)
                                                    fmt.Println(v65, v65)
                                                case "hello":
                                                    v2 += v2
                                                }
                                                if ((312 == 2) && (65535 <= (v65 / ((v69) | 1)))) {
                                                    v2 = (v60 * 1e10)
                                                    _ = len(v3)
                                                }
                                                switch "hello" {
                                                case "a":
                                                    if false {
                                                        var v70 uint16 = (v65 % (((^(v69))) | 1))
                                                        var v71 uint8 = 42
                                                        v60 += 0.0
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                    v72 := GenS10{F0: 18446744073709551615, F1: (24 >= 128)}
                                                    _ = v72
                                                case "":
                                                    fmt.Println(v69, v69, v68)
                                                default:
                                                    var v73 float64 = 0.0
                                                    _ = v73
                                                }
                                            }
                                            v60 += ((v2 / 1.0) * (1.5 + v2))
                                            if true {
                                                var v74 uint64 = 237
                                                fmt.Println(v68, v65, v69)
                                                _ = v74
                                            } else {
                                                if (false && ((v69 ^ v69) >= 46)) {
                                                    v75 := v4[("\n\t" + ("" + "foo bar"))]
                                                    switch fi57 {
                                                    case 2147483647:
                                                        v23 += (v23 | (^(v23)))
                                                    case 32768:
                                                        v76 := GenS11{F0: 212, F1: v23, F2: 1}
                                                        _ = v76
                                                    case -720:
                                                        v75 = v75
                                                    }
                                                    var v77 uint16 = v69
                                                    var v78 string = v75
                                                    _ = v75
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        switch (fi55 | fi57) {
                                                        case 917:
                                                            switch "a" {
                                                            case "hello":
                                                                v23 = 42
                                                            case "":
                                                                v79 := GenS1{F0: v23}
                                                                _ = v79
                                                            default:
                                                                if false {
                                                                    var v80 int32 = 552
                                                                    _ = v80
                                                                } else {
                                                                    v60 += v2
                                                                    var v81 uint64 = 18446744073709551615
                                                                    if true {
                                                                        fmt.Println(v2, v69, v69)
                                                                        fmt.Println(v81, fi55)
                                                                    }
                                                                    _ = v81
                                                                }
                                                            }
                                                        case 4294967295:
                                                            switch (fi57 / ((fi55) | 1)) {
                                                            case 42:
                                                                v82 := v3[(fi55 & (fi55 & fi57))]
                                                                _ = v82
                                                            case 32768:
                                                                if ((-1 <= -9223372036854775808) || false) {
                                                                    if (false && false) {
                                                                        v68 = v68
                                                                    }
                                                                    var v83 float64 = (-2.25 + ((((0.1 / (-2.25 / 1.0)) * float64(float32(float64(v2)))) / (3.14159 * 1e10)) - (1.5 - 3.14159)))
                                                                    v1 = append(v1, "hello")
                                                                    _ = v83
                                                                }
                                                            case 1:
                                                                if (255 == 0) {
                                                                    v60 += v60
                                                                    v60 = v60
                                                                } else {
                                                                    if true {
                                                                        v84 := GenS7{F0: 1, F1: genFn1(255, genFn1(1, 42, 42), 0)}
                                                                        _ = v84
                                                                    }
                                                                    fmt.Println(v68)
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            fmt.Println(v68, fi57, fi57)
                                                                            var v85 int32 = 42
                                                                            _ = v85
                                                                        } else {
                                                                            var v86 uint = 860
                                                                            _ = v86
                                                                        }
                                                                    case "a":
                                                                        switch fi57 {
                                                                        case -2147483648:
                                                                            _ = len(v1)
                                                                            v69 = (v65 | (v69 << ((65535) & 63)))
                                                                        }
                                                                    case "":
                                                                        var v87 int = fi57
                                                                        _ = v87
                                                                    }
                                                                }
                                                                v69 = 255
                                                                switch "a" {
                                                                case "\n\t":
                                                                    delete(v4, "\n\t")
                                                                default:
                                                                    var v88 int32 = 2147483647
                                                                    _ = v88
                                                                }
                                                            }
                                                        case -379:
                                                            if (!(-32768 != 0)) {
                                                                if (42 < 0) {
                                                                    _ = len(v4)
                                                                    _ = len(v3)
                                                                    if true {
                                                                        v68 = (v68 & (^(v68)))
                                                                    }
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v89 string = ""
                                                                        _ = v89
                                                                    case "":
                                                                        var v90 bool = (v69 >= 1)
                                                                        _ = v90
                                                                    default:
                                                                        if false {
                                                                            v68 += -89
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    }
                                                                    var v91 uint32 = v23
                                                                    var v92 float64 = 0.1
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                v93 := GenS12{F0: 2, F1: (v65 < (^(v69))), F2: 42}
                                                                v94 := GenS13{F0: 2, F1: false}
                                                                _ = v93
                                                                _ = v94
                                                            } else {
                                                                if (!(!((255 == 1) || (true && true)))) {
                                                                    if false {
                                                                        v65 += 412
                                                                        var v95 uint32 = v23
                                                                        _ = v95
                                                                    } else {
                                                                        switch ("\n\t" + ("hello" + "")) {
                                                                        case "a":
                                                                            v96 := GenS14{F0: 744, F1: fi57}
                                                                            fmt.Println(v60)
                                                                            _ = v96
                                                                        case "foo bar":
                                                                            if (-2.25 == (((((0.1 / 0.1) / (0.0 - (3.14159 / (1e10 - 3.14159)))) - 0.1) - (1.5 - 1e10)) * (0.0 * -2.25))) {
                                                                                switch (("foo bar" + "hello") + "") {
                                                                                case "foo bar":
                                                                                    v97 := GenS15{F0: fi55}
                                                                                    _ = v97
                                                                                case "\n\t":
                                                                                    var v98 float64 = float64(0.0)
                                                                                    _ = v98
                                                                                default:
                                                                                    v99 := v4["a"]
                                                                                    _ = v99
                                                                                }
                                                                                v2 -= 3.14159
                                                                            } else {
                                                                                var v100 int64 = -9223372036854775808
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                        switch ("a" + "hello") {
                                                                        case "":
                                                                            var v101 bool = (930 > 188)
                                                                            _ = v101
                                                                        case "\n\t":
                                                                            v2 = (float32(-2.25) * (0.0 + v60))
                                                                        default:
                                                                            var v102 int16 = 1000
                                                                            _ = v102
                                                                        }
                                                                        switch fi55 {
                                                                        case 809:
                                                                            v65 = (^(v69))
                                                                        default:
                                                                            v23 += (v23 & v23)
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v103 uint8 = 255
                                                                    fmt.Println(v2, v60)
                                                                    _ = v103
                                                                }
                                                                switch ("a" + "") {
                                                                case "":
                                                                    v104 := GenS16{F0: v68, F1: v23, F2: 255}
                                                                    _ = v104
                                                                case "\n\t":
                                                                    if true {
                                                                        fmt.Println(fi57, v65, fi57)
                                                                        var v105 float64 = (-2.25 - (-2.25 - 0.0))
                                                                        var v106 uint16 = 42
                                                                        _ = v105
                                                                        _ = v106
                                                                    } else {
                                                                        if true {
                                                                            v65 = 0
                                                                            _ = len(v4)
                                                                            var v107 bool = (false && false)
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v108 int8 = (v68 % ((100) | 1))
                                                                    _ = v108
                                                                }
                                                                if (true || (true || (fi57 < fi57))) {
                                                                    var v109 bool = false
                                                                    v110 := GenS16{F0: (v68 >> ((4294967295) & 63)), F1: (v23 * (v23 + v23)), F2: 65535}
                                                                    v65 -= (v65 / ((914) | 1))
                                                                    _ = v109
                                                                    _ = v110
                                                                } else {
                                                                    var v111 uint8 = 127
                                                                    _ = v111
                                                                }
                                                            }
                                                        default:
                                                            if ((!((^(v69)) >= v65)) || (0 <= uint64(v69))) {
                                                                var v112 uint16 = (^(v69))
                                                                fmt.Println(v68)
                                                                v1 = append(v1, ("hello" + ((("\n\t" + "hello") + "") + ("" + "foo bar"))))
                                                                _ = v112
                                                            }
                                                        }
                                                    }
                                                    var v113 int32 = 1
                                                    _ = v113
                                                }
                                                if false {
                                                    v114 := GenS17{F0: "\n\t"}
                                                    if (3.14159 >= 1.5) {
                                                        if ((fi55 / ((fi57) | 1)) <= fi55) {
                                                            v60 += 1e10
                                                            v60 += 3.14159
                                                        }
                                                        v69 -= v69
                                                        v115 := GenS18{F0: fi57, F1: 228}
                                                        _ = v115
                                                    }
                                                    v68 += (v68 >> ((705) & 63))
                                                    _ = v114
                                                }
                                                v116 := GenS17{F0: ("hello" + "\n\t")}
                                                _ = v116
                                            }
                                        }
                                    case "":
                                        var v117 uint64 = 42
                                        _ = v117
                                    case "hello":
                                        var v118 float64 = 0.0
                                        _ = v118
                                    default:
                                        _ = len(v4)
                                    }
                                    var v119 uint32 = v23
                                    v65 = (v69 + (v69 / ((v69) | 1)))
                                    var v120 int = 1
                                    var v121 float32 = float32((float64(v2) + float64(v60)))
                                    var v122 uint8 = 41
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                    _ = v122
                                    _ = v69
                                } else {
                                    if false {
                                        var v123 uint = 873
                                        var v124 uint32 = 1
                                        v124 += v124
                                        _ = v123
                                        _ = v124
                                    } else {
                                        var v125 float32 = (v60 + v60)
                                        _ = v125
                                    }
                                }
                                var v126 uint = 860
                                var v127 uint32 = (^(v23))
                                _ = v126
                                _ = v127
                                _ = v68
                            }
                            fmt.Println(v23, fi57, v65)
                            _ = v60
                            _ = v65
                        }
                        v2 -= ((0.0 + ((v2 - (v2 + v2)) + float32((1e10 * 0.1)))) - float32((1e10 / 3.14159)))
                        var v128 uint = 18446744073709551615
                        var v129 int16 = 1
                        v23 += 483
                        v130 := GenS13{F0: 453, F1: (!(!(341 != -1)))}
                        _ = v128
                        _ = v129
                        _ = v130
                    } else {
                        delete(v4, "hello")
                    }
                    _ = fi57
                    _ = v58
                }
                for fi131 := 0; fi131 < 16; fi131++ {
                    if ((!(859 == 1)) && true) {
                        var v132 float32 = (v2 - (v2 + v2))
                        _ = v132
                    }
                    _ = fi131
                }
            }
        } else {
            if false {
                var v133 uint64 = 4294967295
                if true {
                    fmt.Println(v23)
                    switch "a" {
                    case "a":
                        v2 = v2
                    case "\n\t":
                        v23 -= v23
                    }
                }
                _ = v133
            }
        }
        _ = fi55
        _ = v56
    }
    v134 := GenS3{F0: 18446744073709551615, F1: -109}
    var v135 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v134
    _ = v135
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
