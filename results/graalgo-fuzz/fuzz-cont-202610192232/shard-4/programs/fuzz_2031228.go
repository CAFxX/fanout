package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int16; F2 int }
type GenS4 struct { F0 string; F1 uint16; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16; F1 bool; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint32 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint64; F1 uint64 }
type GenS13 struct { F0 int8; F1 uint16 }
type GenS14 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int; F1 string }
type GenS18 struct { F0 int; F1 uint }
type GenS19 struct { F0 string }

func main() {
    var v1 uint64 = 251
    v2 := make(map[string]bool)
    var v3 bool = (-1 <= 82)
    v4 := []int{-1}
    if v3 {
        v5 := v2[("foo bar" + "")]
        var v6 uint64 = (v1 - 830)
        if false {
            if true {
                fmt.Println(v3, v1)
            }
        }
        _ = v5
        _ = v6
    } else {
        var v7 uint8 = 2
        _ = len(v4)
        _ = v7
    }
    if v3 {
        if v3 {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 42, F1: -1, F2: 11}
                v10 := GenS1{F0: 67, F1: 42, F2: 1}
                v3 = (!v3)
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS2{F0: 4294967295, F1: v1, F2: v3}
            switch "\n\t" {
            case "":
                if (-620 == 42) {
                    switch 32767 {
                    case 2147483647:
                        var v12 int32 = 0
                        _ = v12
                    case 32768:
                        var v13 float64 = 0.0
                        _ = v13
                    }
                    if (v3 && v3) {
                        v4 = append(v4, -902)
                        for fi14 := 0; fi14 < 4; fi14++ {
                            fmt.Println(v3, fi14, v3)
                            _ = fi14
                        }
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 int32 = 0
                            if ("\n\t" <= "\n\t") {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v1 += 4294967295
                        var v17 float64 = (((3.14159 + 1.5) + 1.5) * float64((3.14159 - float32(((((-2.25 * (float64(0.0) - 0.1)) * (0.1 + float64(float32(1.5)))) + 0.1) * 0.0)))))
                        switch 128 {
                        case 9223372036854775807:
                            fmt.Println(v3)
                        }
                        _ = v17
                    }
                    if (!(!v3)) {
                        var v18 uint8 = 128
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 int16 = 894
                            _ = fi19
                            _ = v20
                        }
                        v3 = false
                        v3 = v3
                    }
                }
            case "\n\t":
                var v21 int8 = 100
                _ = v21
            }
            _ = v11
        } else {
            switch "foo bar" {
            case "\n\t":
                var v22 int64 = -1
                _ = v22
            case "":
                switch 0 {
                case 598:
                    v1 = v1
                case 481:
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 uint8 = 0
                        _ = fi23
                        _ = v24
                    }
                case -2147483648:
                    v25 := GenS2{F0: 389, F1: (v1 >> ((2) & 63)), F2: (!true)}
                    _ = v25
                }
            }
            if (252 != 65535) {
                if v3 {
                    v4 = append(v4, 428)
                    v26 := GenS2{F0: 816, F1: v1, F2: (false && (-1000 > 1))}
                    _ = v26
                } else {
                    v27 := GenS2{F0: 255, F1: (^(v1)), F2: (-1 >= -1)}
                    _ = v27
                }
                fmt.Println(v3)
            } else {
                fmt.Println(v3)
            }
        }
        var v28 uint64 = v1
        _ = v28
    } else {
        if v3 {
            switch (("\n\t" + "a") + "a") {
            case "foo bar":
                if v3 {
                    if v3 {
                        var v29 string = "\n\t"
                        switch "a" {
                        case "":
                            if (v3 || true) {
                                switch "foo bar" {
                                case "":
                                    v30 := GenS3{F0: 2, F1: 240, F2: 9223372036854775807}
                                    _ = v30
                                case "hello":
                                    v29 = "foo bar"
                                }
                            }
                        }
                        if false {
                            v3 = v3
                            var v31 int32 = 0
                            _ = v31
                        } else {
                            switch 9223372036854775807 {
                            case 128:
                                for fi32 := 0; fi32 < 2; fi32++ {
                                    v3 = (!(((fi32 - fi32) >= (fi32 << ((0) & 63))) || (true && false)))
                                    v33 := GenS2{F0: 457, F1: 322, F2: ((((v29 + (v29 + v29)) + "hello") + v29) > "\n\t")}
                                    _ = fi32
                                    _ = v33
                                }
                            case 127:
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    if v3 {
                                        v3 = false
                                        v35 := GenS4{F0: v29, F1: 65535, F2: 42}
                                        _ = v35
                                    } else {
                                        for fi36 := 0; fi36 < 7; fi36++ {
                                            var v37 float64 = 0.1
                                            _ = fi36
                                            _ = v37
                                        }
                                        v1 += v1
                                    }
                                    _ = fi34
                                }
                            default:
                                var v38 int8 = 0
                                _ = v38
                            }
                            if (!true) {
                                v3 = (0.1 <= 1e10)
                            } else {
                                if ((0.1 != (1e10 * float64((((1e10 * 0.1) + 3.14159) + (((3.14159 + 0.0) + 0.0) * 3.14159))))) && (!v3)) {
                                    switch -702 {
                                    case 42:
                                        v39 := GenS5{F0: 42}
                                        _ = v39
                                    case -2147483648:
                                        var v40 uint16 = 32767
                                        _ = v40
                                    default:
                                        switch 157 {
                                        case 32767:
                                            switch "" {
                                            case "\n\t":
                                                for fi41 := 0; fi41 < 11; fi41++ {
                                                    fmt.Println(v1, v29, v29)
                                                    if v3 {
                                                        if len(v4) > 0 {
                                                            v4[0] = 300
                                                        }
                                                        v42 := GenS2{F0: 401, F1: v1, F2: true}
                                                        v29 = v29
                                                        _ = v42
                                                    }
                                                    _ = fi41
                                                }
                                            case "":
                                                if v3 {
                                                    v29 = "\n\t"
                                                } else {
                                                    if ((true || false) && ((-40 <= -30) && (!(308 != 1)))) {
                                                        v1 = v1
                                                        v43 := GenS2{F0: 42, F1: 763, F2: ((2147483647 > -2147483648) && false)}
                                                        _ = v43
                                                    }
                                                    var v44 uint32 = uint32(v1)
                                                    fmt.Println(v3, v44, v3)
                                                    _ = v44
                                                }
                                            case "foo bar":
                                                if false {
                                                    if false {
                                                        switch 32767 {
                                                        case -1:
                                                            var v45 int16 = -477
                                                            _ = v45
                                                        case 2147483647:
                                                            if false {
                                                                for fi46 := 0; fi46 < 20; fi46++ {
                                                                    var v47 int = (fi46 + (fi46 - 731))
                                                                    v1 = v1
                                                                    _ = fi46
                                                                    _ = v47
                                                                }
                                                                var v48 int32 = 2147483647
                                                                _ = v48
                                                            } else {
                                                                if len(v4) > 0 {
                                                                    v4[0] = 42
                                                                }
                                                                v2[v29] = (v1 == v1)
                                                                v49 := GenS5{F0: -2147483648}
                                                                _ = v49
                                                            }
                                                        case -479:
                                                            var v50 uint8 = 0
                                                            _ = v50
                                                        }
                                                    } else {
                                                        v29 = (("" + (v29 + v29)) + v29)
                                                        if v3 {
                                                            var v51 uint8 = 133
                                                            _ = v51
                                                        } else {
                                                            v4 = append(v4, -1)
                                                        }
                                                    }
                                                    if true {
                                                        if ((true && (!v3)) && (-645 != 796)) {
                                                            if false {
                                                                if true {
                                                                    fmt.Println(v29, v1, v1)
                                                                    if v3 {
                                                                        for fi52 := 0; fi52 < 10; fi52++ {
                                                                            for fi53 := 0; fi53 < 11; fi53++ {
                                                                                v29 = "foo bar"
                                                                                v1 -= v1
                                                                                _ = fi53
                                                                            }
                                                                            _ = fi52
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    if v3 {
                                                                        if (((1 > 9223372036854775807) || (true && v3)) || v3) {
                                                                            var v54 uint8 = 128
                                                                            _ = v54
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        if (!v3) {
                                                                            for fi55 := 0; fi55 < 2; fi55++ {
                                                                                v56 := GenS1{F0: 70, F1: -1, F2: 128}
                                                                                if (102 == 2) {
                                                                                    for fi57 := 0; fi57 < 6; fi57++ {
                                                                                        v58 := GenS6{F0: v29, F1: 42}
                                                                                        v59 := GenS7{F0: 1, F1: v3, F2: 0}
                                                                                        _ = fi57
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                    if v3 {
                                                                                        fmt.Println(v1, fi55, v29)
                                                                                    }
                                                                                }
                                                                                var v60 int = (fi55 - -1)
                                                                                _ = fi55
                                                                                _ = v56
                                                                                _ = v60
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v3, v29, v3)
                                                                        }
                                                                        if v3 {
                                                                            var v61 string = ((("hello" + (v29 + v29)) + (("a" + (v29 + v29)) + "\n\t")) + v29)
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v62 uint16 = 533
                                                                    var v63 float32 = 3.14159
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                for fi64 := 0; fi64 < 2; fi64++ {
                                                                    v65 := GenS7{F0: 2, F1: (42 < 4294967295), F2: 0}
                                                                    _ = fi64
                                                                    _ = v65
                                                                }
                                                            }
                                                            v66 := GenS8{F0: 943}
                                                            v67 := v2[v29]
                                                            _ = v66
                                                            _ = v67
                                                        } else {
                                                            var v68 int16 = -32768
                                                            _ = len(v4)
                                                            _ = v68
                                                        }
                                                        v29 = (v29 + (v29 + v29))
                                                        v3 = v3
                                                    }
                                                } else {
                                                    for fi69 := 0; fi69 < 5; fi69++ {
                                                        for fi70 := 0; fi70 < 1; fi70++ {
                                                            v29 = v29
                                                            _ = fi70
                                                        }
                                                        if true {
                                                            v1 -= (v1 / (((v1 - v1)) | 1))
                                                            v71 := GenS9{F0: v29, F1: 127}
                                                            _ = v71
                                                        }
                                                        _ = fi69
                                                    }
                                                }
                                            }
                                        case -219:
                                            if v3 {
                                                for fi72 := 0; fi72 < 14; fi72++ {
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi72 & (fi72 / ((fi72) | 1)))
                                                    }
                                                    _ = fi72
                                                }
                                            } else {
                                                if v3 {
                                                    switch (v29 + "\n\t") {
                                                    case "":
                                                        if v3 {
                                                            fmt.Println(v29, v1, v1)
                                                            switch v29 {
                                                            case "foo bar":
                                                                var v73 string = (v29 + "\n\t")
                                                                _ = v73
                                                            case "a":
                                                                for fi74 := 0; fi74 < 1; fi74++ {
                                                                    if (212 != int32(fi74)) {
                                                                        v1 += v1
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (fi74 * (fi74 & fi74))
                                                                        }
                                                                        v29 = v29
                                                                    } else {
                                                                        fmt.Println(v1, v3, v3)
                                                                        _ = len(v2)
                                                                    }
                                                                    delete(v2, ((v29 + "a") + v29))
                                                                    _ = fi74
                                                                }
                                                            default:
                                                                v75 := GenS10{F0: 127, F1: 817}
                                                                _ = v75
                                                            }
                                                            for fi76 := 0; fi76 < 2; fi76++ {
                                                                for fi77 := 0; fi77 < 19; fi77++ {
                                                                    v4 = append(v4, (fi77 | fi77))
                                                                    _ = fi77
                                                                }
                                                                _ = fi76
                                                            }
                                                        } else {
                                                            v78 := GenS8{F0: 1}
                                                            _ = v78
                                                        }
                                                    case "foo bar":
                                                        for fi79 := 0; fi79 < 12; fi79++ {
                                                            v80 := GenS3{F0: 42, F1: 356, F2: (fi79 + fi79)}
                                                            for fi81 := 0; fi81 < 3; fi81++ {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch fi81 {
                                                                    case 9223372036854775807:
                                                                        delete(v2, v29)
                                                                    case 4294967295:
                                                                        v82 := GenS11{F0: (768 >= -1), F1: 2, F2: 42}
                                                                        _ = v82
                                                                    }
                                                                case "a":
                                                                    v1 += (^(v1))
                                                                }
                                                                if v3 {
                                                                    v3 = v3
                                                                }
                                                                v3 = v3
                                                                _ = fi81
                                                            }
                                                            _ = fi79
                                                            _ = v80
                                                        }
                                                    case "\n\t":
                                                        var v83 float64 = 3.14159
                                                        _ = v83
                                                    }
                                                    for fi84 := 0; fi84 < 12; fi84++ {
                                                        v1 = v1
                                                        fmt.Println(v29)
                                                        _ = fi84
                                                    }
                                                    v85 := GenS12{F0: (v1 | (v1 * v1)), F1: v1}
                                                    _ = v85
                                                }
                                                var v86 int16 = -32768
                                                v87 := v2[v29]
                                                _ = v86
                                                _ = v87
                                            }
                                        case 337:
                                            var v88 uint16 = 42
                                            _ = v88
                                        default:
                                            _ = len(v2)
                                        }
                                    }
                                    var v89 uint = 249
                                    _ = v89
                                }
                                delete(v2, ((("foo bar" + "hello") + v29) + ""))
                            }
                            v3 = v3
                        }
                        _ = v29
                    } else {
                        v3 = v3
                        v90 := GenS13{F0: 100, F1: 2}
                        _ = v90
                    }
                } else {
                    var v91 float32 = (-2.25 * 1e10)
                    v92 := GenS3{F0: 65535, F1: 1, F2: 422}
                    if (2 == 2) {
                        v93 := GenS8{F0: 46}
                        _ = v93
                    } else {
                        for fi94 := 0; fi94 < 19; fi94++ {
                            v95 := GenS9{F0: "\n\t", F1: -128}
                            _ = fi94
                            _ = v95
                        }
                        if (!v3) {
                            for fi96 := 0; fi96 < 1; fi96++ {
                                v97 := GenS14{F0: 65535, F1: -114, F2: 2147483648}
                                var v98 int64 = -1
                                for fi99 := 0; fi99 < 10; fi99++ {
                                    if (v3 && (((((v98 & 9223372036854775807) <= int64(fi96)) || v3) && (-359 != -845)) || v3)) {
                                        if (!((fi99 << ((2) & 63)) <= fi99)) {
                                            var v100 uint16 = 0
                                            v101 := GenS15{F0: 4294967295}
                                            _ = v100
                                            _ = v101
                                        } else {
                                            v102 := v2[("" + "hello")]
                                            v98 = (^(v98))
                                            _ = v102
                                        }
                                        var v103 float64 = ((3.14159 * 0.0) / 1.0)
                                        var v104 float32 = v91
                                        _ = v103
                                        _ = v104
                                    }
                                    _ = fi99
                                }
                                _ = fi96
                                _ = v97
                                _ = v98
                            }
                            if (true && ((v3 || v3) || v3)) {
                                v1 = v1
                            } else {
                                if v3 {
                                    var v105 uint64 = (v1 * (^(v1)))
                                    _ = v105
                                }
                                v106 := GenS5{F0: -2147483648}
                                _ = v106
                            }
                        }
                        v91 -= (((3.14159 + float32(((3.14159 + 0.1) * 0.1))) / v91) * -2.25)
                    }
                    _ = v91
                    _ = v92
                }
            case "\n\t":
                if v3 {
                    switch 2147483647 {
                    case -501:
                        for fi107 := 0; fi107 < 11; fi107++ {
                            var v108 int = -2147483648
                            fmt.Println(fi107, v1, v1)
                            var v109 int64 = 42
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                    case -524:
                        var v110 uint32 = 794
                        _ = v110
                    }
                } else {
                    fmt.Println(v3)
                }
            case "a":
                v111 := GenS11{F0: v3, F1: 1, F2: 967}
                _ = v111
            default:
                var v112 float32 = 1.5
                _ = v112
            }
        } else {
            v1 -= (^(v1))
        }
        var v113 uint32 = 65535
        _ = v113
    }
    for fi114 := 0; fi114 < 1; fi114++ {
        v1 = 42
        _ = fi114
    }
    var v115 int16 = 42
    switch "\n\t" {
    case "\n\t":
        if v3 {
            var v116 uint = 18446744073709551615
            fmt.Println(v115, v1, v3)
            v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
            _ = v116
        }
    default:
        v4 = append(v4, -9223372036854775808)
    }
    switch int(v115) {
    case -9223372036854775808:
        var v117 int = -2147483648
        _ = v117
    case 32768:
        _ = len(v4)
    case 127:
        var v118 int = -2147483648
        _ = v118
    }
    if (("\n\t" + "\n\t") != ((("a" + "foo bar") + "foo bar") + "hello")) {
        v119 := v2["\n\t"]
        v120 := v2[("" + "a")]
        switch 128 {
        case -604:
            var v121 int64 = 441
            _ = v121
        }
        _ = v119
        _ = v120
    } else {
        if v3 {
            v122 := GenS4{F0: "", F1: 0, F2: 100}
            switch 127 {
            case 433:
                var v123 uint64 = 0
                _ = v123
            }
            _ = v122
        } else {
            v3 = (0 <= 98)
            v1 -= 42
        }
        for fi124 := 0; fi124 < 1; fi124++ {
            v115 -= (v115 * 0)
            var v125 float64 = (0.1 + 0.1)
            for fi126 := 0; fi126 < 7; fi126++ {
                if v3 {
                    _ = len(v2)
                    if (((1.5 * 1.5) >= 1e10) || true) {
                        v127 := v2[("\n\t" + ("foo bar" + ""))]
                        if (-128 != 96) {
                            v128 := GenS16{F0: (v1 - 42)}
                            _ = v128
                        }
                        var v129 float32 = 3.14159
                        _ = v127
                        _ = v129
                    }
                }
                v130 := GenS17{F0: (fi126 & 2147483647), F1: ""}
                _ = fi126
                _ = v130
            }
            _ = fi124
            _ = v125
        }
        var v131 uint32 = 1
        _ = v131
    }
    var v132 int8 = -128
    var v133 uint64 = 0
    v134 := GenS18{F0: 1, F1: 255}
    v3 = (812 < -1)
    var v135 uint16 = 58
    if (0.1 > (0.1 + 1e10)) {
        v136 := GenS19{F0: ((((("" + ("a" + "foo bar")) + ("a" + "a")) + ("a" + "\n\t")) + (("a" + "") + ("hello" + "a"))) + "foo bar")}
        v137 := GenS1{F0: (^(v132)), F1: 9223372036854775807, F2: 36}
        v138 := GenS19{F0: ""}
        _ = v136
        _ = v137
        _ = v138
    }
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v135:", v135)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v115
    _ = v132
    _ = v133
    _ = v134
    _ = v135
    _ = v2
    _ = v3
    _ = v4
}
