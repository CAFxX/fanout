package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 uint64; F2 int32 }
type GenS9 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int8; F1 int32; F2 int }
type GenS16 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 string = ("foo bar" + "hello")
    v5 := []int{}
    v6 := GenS1{F0: 1}
    v2[-1] = "foo bar"
    v7 := GenS1{F0: 613}
    switch v4 {
    case "hello":
        if true {
            if true {
                var v8 float64 = 3.14159
                fmt.Println(v4, v4)
                switch -141 {
                case 4294967295:
                    var v9 float32 = (1.5 + 0.1)
                    _ = v9
                case 0:
                    switch v4 {
                    case "hello":
                        v3["\n\t"] = true
                    case "foo bar":
                        v10 := GenS2{F0: 42, F1: 32768, F2: 2}
                        _ = v10
                    default:
                        v8 += (v8 + (3.14159 - v8))
                    }
                case -9223372036854775808:
                    if false {
                        v8 = v8
                        switch 0 {
                        case 0:
                            if ((-1 <= 0) && false) {
                                var v11 float32 = 0.0
                                switch v4 {
                                case "\n\t":
                                    fmt.Println(v4)
                                case "a":
                                    v12 := GenS1{F0: 255}
                                    _ = v12
                                default:
                                    v11 += -2.25
                                }
                                for fi13 := 0; fi13 < 20; fi13++ {
                                    fmt.Println(v11, v11, v4)
                                    v14 := v3["hello"]
                                    _ = fi13
                                    _ = v14
                                }
                                _ = v11
                            } else {
                                _ = len(v2)
                            }
                        case 32768:
                            for fi15 := 0; fi15 < 9; fi15++ {
                                for fi16 := 0; fi16 < 2; fi16++ {
                                    v17 := v3[("foo bar" + v4)]
                                    v4 = ""
                                    if v17 {
                                        if (4294967295 < 0) {
                                            switch (v4 + "foo bar") {
                                            case "\n\t":
                                                v4 = "a"
                                            default:
                                                var v18 float64 = 3.14159
                                                _ = v18
                                            }
                                            v17 = ((fi15 + (fi15 & -981)) == fi16)
                                            v19 := GenS2{F0: 42, F1: fi16, F2: 2}
                                            _ = v19
                                        } else {
                                            if (0 < 4294967295) {
                                                if v17 {
                                                    v20 := GenS3{F0: 2, F1: 18446744073709551615}
                                                    var v21 uint8 = 1
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v17, fi15)
                                                    }
                                                }
                                                if (fi15 > fi16) {
                                                    delete(v2, (-(fi15)))
                                                }
                                                v4 = v4
                                                var v22 int32 = 1
                                                if (42 < 832) {
                                                    v22 += v22
                                                } else {
                                                    fmt.Println(fi16, v4, fi16)
                                                    var v23 int = fi16
                                                    if v17 {
                                                        switch v4 {
                                                        case "a":
                                                            v5 = append(v5, fi16)
                                                        default:
                                                            v24 := GenS2{F0: 42, F1: 9223372036854775807, F2: 0}
                                                            _ = v24
                                                        }
                                                        v25 := GenS4{F0: 18446744073709551615}
                                                        var v26 uint = 2
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    _ = v23
                                                }
                                                _ = v22
                                            }
                                            if v17 {
                                                v17 = v17
                                                v1[2147483648] = 2147483647
                                                delete(v2, fi16)
                                                switch (-(fi16)) {
                                                case 42:
                                                    var v27 int64 = -1
                                                    _ = v27
                                                }
                                                var v28 uint64 = 196
                                                _ = v28
                                            } else {
                                                var v29 uint = 42
                                                if (32767 == -928) {
                                                    v30 := GenS5{F0: v17, F1: uint8(v29)}
                                                    var v31 float64 = (((v8 * (float64((float32(3.14159) + (1e10 - (float32((float64((1.5 - 3.14159)) * (-2.25 + (1.5 * 0.0)))) * (float32(v8) * (float32((3.14159 - -2.25)) - (0.1 * (0.1 / 1.5)))))))) - float64(1.5))) + v8) - 3.14159)
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v32 := GenS3{F0: 1, F1: 0}
                                                    if v17 {
                                                        v33 := GenS6{F0: 32768}
                                                        _ = v33
                                                    } else {
                                                        fmt.Println(v29)
                                                        fmt.Println(v4, v8, fi15)
                                                        v17 = (585 >= 65535)
                                                        if (v17 || true) {
                                                            v34 := GenS3{F0: 2, F1: 1}
                                                            switch (fi15 >> ((v29) & 63)) {
                                                            case 0:
                                                                if v17 {
                                                                    fmt.Println(v4, fi16, fi16)
                                                                    var v35 float32 = 3.14159
                                                                    _ = v35
                                                                }
                                                            case 953:
                                                                if (!false) {
                                                                    v36 := GenS5{F0: true, F1: 42}
                                                                    _ = v36
                                                                }
                                                            case 9223372036854775807:
                                                                var v37 int16 = 1000
                                                                _ = v37
                                                            }
                                                            var v38 uint = (v29 & 2)
                                                            _ = v34
                                                            _ = v38
                                                        }
                                                        var v39 int64 = 9223372036854775807
                                                        _ = v39
                                                    }
                                                    v29 = 255
                                                    _ = v32
                                                }
                                                v17 = (!(!v17))
                                                var v40 int64 = 9223372036854775807
                                                _ = v29
                                                _ = v40
                                            }
                                            v8 = ((((v8 / (((3.14159 / v8) + v8) + 0.0)) - v8) * (1.5 + ((1.5 / ((v8 / (v8 * ((v8 * float64(float32(((((v8 / v8) * 0.1) + 0.0) * 1e10)))) / -2.25))) + v8)) - (v8 * (v8 - ((v8 * v8) / (v8 * ((3.14159 * (v8 * 0.0)) + v8)))))))) - v8)
                                        }
                                    } else {
                                        if (true || v17) {
                                            v17 = v17
                                            switch (((v4 + v4) + (v4 + "hello")) + v4) {
                                            case "\n\t":
                                                fmt.Println(v8, v8, fi16)
                                            }
                                            v4 = "hello"
                                        } else {
                                            v4 = v4
                                            _ = len(v5)
                                            _ = len(v5)
                                        }
                                        v41 := GenS7{F0: 555}
                                        _ = v41
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                                var v42 bool = (36 < 1)
                                _ = len(v1)
                                _ = fi15
                                _ = v42
                            }
                        case -1:
                            fmt.Println(v4, v4)
                        }
                    } else {
                        if (v4 < ("foo bar" + "a")) {
                            v43 := GenS1{F0: 65535}
                            var v44 uint32 = 1
                            v8 = v8
                            _ = v43
                            _ = v44
                        } else {
                            if (true && false) {
                                fmt.Println(v4)
                                _ = len(v5)
                                var v45 uint32 = 2147483648
                                _ = v45
                            } else {
                                v4 = v4
                                var v46 bool = true
                                if len(v5) > 0 {
                                    v5[0] = 891
                                }
                                _ = v46
                            }
                        }
                    }
                default:
                    for fi47 := 0; fi47 < 14; fi47++ {
                        if false {
                            if false {
                                v4 = v4
                                for fi48 := 0; fi48 < 9; fi48++ {
                                    var v49 uint32 = 2147483648
                                    fmt.Println(v8, v8)
                                    _ = fi48
                                    _ = v49
                                }
                            } else {
                                v50 := GenS1{F0: 32767}
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 15; fi51++ {
                                if false {
                                    v4 = v4
                                } else {
                                    switch (v4 + "") {
                                    case "hello":
                                        _ = len(v5)
                                    case "a":
                                        var v52 int64 = 0
                                        _ = v52
                                    case "foo bar":
                                        v53 := GenS8{F0: "foo bar", F1: 916, F2: 2147483647}
                                        _ = v53
                                    }
                                    var v54 uint = 0
                                    _ = v54
                                }
                                switch (fi51 ^ 4294967295) {
                                case 127:
                                    var v55 float32 = 1e10
                                    _ = v55
                                case 832:
                                    var v56 int64 = 1
                                    _ = v56
                                }
                                _ = fi51
                            }
                        }
                        _ = fi47
                    }
                }
                _ = v8
            }
            if true {
                if true {
                    v57 := GenS9{F0: 1, F1: 2, F2: 42}
                    _ = v57
                } else {
                    switch v4 {
                    case "\n\t":
                        if false {
                            var v58 uint8 = 2
                            if true {
                                var v59 uint16 = 1
                                v59 = (v59 & v59)
                                _ = v59
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                                delete(v1, 135)
                            }
                            _ = v58
                        }
                    default:
                        v4 = v4
                    }
                }
                v60 := GenS10{F0: true}
                switch (v4 + "\n\t") {
                case "hello":
                    var v61 bool = (!false)
                    _ = v61
                }
                _ = v60
            } else {
                if ((0.0 / float32((float64(0.1) / (3.14159 / 1.0)))) >= ((1e10 + 0.1) + 1e10)) {
                    switch v4 {
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    switch 128 {
                                    case 42:
                                        v4 = v4
                                    case -175:
                                        for fi62 := 0; fi62 < 16; fi62++ {
                                            v4 = v4
                                            var v63 int32 = -372
                                            _ = fi62
                                            _ = v63
                                        }
                                    case 127:
                                        if false {
                                            v64 := GenS6{F0: 2147483647}
                                            fmt.Println(v4)
                                            _ = v64
                                        } else {
                                            _ = len(v5)
                                        }
                                    default:
                                        var v65 uint = 0
                                        _ = v65
                                    }
                                    if (!(true || true)) {
                                        var v66 uint = 873
                                        _ = v66
                                    } else {
                                        v4 = v4
                                        for fi67 := 0; fi67 < 13; fi67++ {
                                            v68 := GenS11{F0: -931, F1: 255, F2: 751}
                                            for fi69 := 0; fi69 < 3; fi69++ {
                                                var v70 int32 = -9
                                                _ = fi69
                                                _ = v70
                                            }
                                            _ = fi67
                                            _ = v68
                                        }
                                    }
                                }
                            case "a":
                                v71 := GenS12{F0: 42}
                                _ = v71
                            case "foo bar":
                                if true {
                                    v4 = "foo bar"
                                    if false {
                                        var v72 string = "\n\t"
                                        var v73 uint = 18446744073709551615
                                        _ = v72
                                        _ = v73
                                    } else {
                                        var v74 uint = 18446744073709551615
                                        fmt.Println(v74, v74)
                                        _ = v74
                                    }
                                    switch 1 {
                                    case 32768:
                                        var v75 int8 = 127
                                        _ = v75
                                    case -221:
                                        fmt.Println(v4)
                                    case -9223372036854775808:
                                        v76 := GenS9{F0: 1, F1: 541, F2: 42}
                                        _ = v76
                                    default:
                                        switch 42 {
                                        case 344:
                                            if (true && true) {
                                                for fi77 := 0; fi77 < 3; fi77++ {
                                                    v4 = ((v4 + (("hello" + "\n\t") + (("foo bar" + "") + v4))) + "")
                                                    v78 := v2[fi77]
                                                    for fi79 := 0; fi79 < 16; fi79++ {
                                                        v78 = (("a" + ((v4 + ((v4 + (v4 + v4)) + (v4 + ((v78 + v78) + v4)))) + "")) + "hello")
                                                        v78 = (v4 + v78)
                                                        _ = fi79
                                                    }
                                                    _ = fi77
                                                    _ = v78
                                                }
                                                v4 = ((v4 + v4) + (((v4 + ((("" + ("foo bar" + v4)) + (v4 + v4)) + (v4 + (v4 + v4)))) + ("a" + v4)) + "hello"))
                                                if (false || false) {
                                                    delete(v3, "")
                                                } else {
                                                    switch -458 {
                                                    case 9223372036854775807:
                                                        var v80 uint8 = 161
                                                        _ = v80
                                                    case 0:
                                                        v4 = ((((v4 + v4) + v4) + v4) + (("a" + v4) + "\n\t"))
                                                    default:
                                                        var v81 bool = (966 >= 1)
                                                        _ = v81
                                                    }
                                                }
                                            } else {
                                                var v82 int = -9223372036854775808
                                                v4 = "foo bar"
                                                _ = v82
                                            }
                                        case -1:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    }
                                }
                            default:
                                if false {
                                    v3[v4] = true
                                }
                            }
                        default:
                            v4 = (v4 + (((v4 + v4) + v4) + (("\n\t" + v4) + "\n\t")))
                        }
                    }
                    for fi83 := 0; fi83 < 1; fi83++ {
                        var v84 int8 = 70
                        v85 := GenS13{F0: 9223372036854775807}
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    if true {
                        v4 = v4
                        v4 = v4
                        var v86 uint8 = 128
                        _ = v86
                    }
                }
                switch ("foo bar" + "") {
                case "a":
                    v87 := GenS6{F0: 9223372036854775807}
                    _ = v87
                case "\n\t":
                    switch v4 {
                    case "foo bar":
                        fmt.Println(v4)
                    case "":
                        for fi88 := 0; fi88 < 10; fi88++ {
                            if (false || true) {
                                v89 := GenS2{F0: 408, F1: 32768, F2: 4294967295}
                                var v90 int8 = -1
                                fmt.Println(v90, v4)
                                _ = v89
                                _ = v90
                            } else {
                                fmt.Println(fi88, fi88)
                                var v91 string = v4
                                for fi92 := 0; fi92 < 1; fi92++ {
                                    v93 := GenS14{F0: (fi92 / (((fi88 | (fi88 >> ((768) & 63)))) | 1))}
                                    v94 := GenS15{F0: 42, F1: 586, F2: (^(fi88))}
                                    var v95 bool = true
                                    v4 = ("hello" + (v4 + "foo bar"))
                                    _ = fi92
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                }
                                _ = v91
                            }
                            switch (fi88 ^ 1) {
                            case -1:
                                var v96 int64 = 791
                                _ = v96
                            }
                            _ = fi88
                        }
                    case "hello":
                        fmt.Println(v4)
                    default:
                        for fi97 := 0; fi97 < 15; fi97++ {
                            if (42 > 42) {
                                _ = len(v5)
                            } else {
                                for fi98 := 0; fi98 < 18; fi98++ {
                                    _ = len(v5)
                                    var v99 int16 = int16(fi97)
                                    _ = fi98
                                    _ = v99
                                }
                                v4 = v4
                                if true {
                                    v1[42] = (fi97 - 127)
                                    for fi100 := 0; fi100 < 4; fi100++ {
                                        if len(v5) > 0 {
                                            v5[0] = fi100
                                        }
                                        _ = fi100
                                    }
                                    v4 = v4
                                } else {
                                    v4 = "foo bar"
                                }
                            }
                            var v101 int8 = 1
                            _ = len(v3)
                            _ = fi97
                            _ = v101
                        }
                    }
                case "":
                    fmt.Println(v4)
                default:
                    _ = len(v5)
                }
                for fi102 := 0; fi102 < 13; fi102++ {
                    fmt.Println(v4, fi102)
                    if len(v5) > 0 {
                        v5[0] = (fi102 << ((272) & 63))
                    }
                    _ = fi102
                }
            }
        }
    case "a":
        fmt.Println(v4)
    case "\n\t":
        for fi103 := 0; fi103 < 14; fi103++ {
            if true {
                if ((float32((1e10 + float64(0.0))) / 0.1) <= 1e10) {
                    for fi104 := 0; fi104 < 12; fi104++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v105 int32 = 2147483647
                        var v106 int = (fi104 * (fi103 | fi103))
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                } else {
                    var v107 bool = (!(-1 >= -1))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v107
                }
                if (1 < 2) {
                    fmt.Println(fi103)
                    for fi108 := 0; fi108 < 9; fi108++ {
                        v4 = v4
                        _ = fi108
                    }
                }
            }
            _ = fi103
        }
    }
    var v109 bool = false
    for fi110 := 0; fi110 < 7; fi110++ {
        if ((4294967295 == 483) && (!((v4 + v4) != "foo bar"))) {
            var v111 int8 = 42
            var v112 uint8 = 42
            _ = v111
            _ = v112
        }
        _ = fi110
    }
    var v113 bool = false
    switch v4 {
    case "foo bar":
        var v114 float64 = float64((1.5 - 1.5))
        _ = v114
    case "\n\t":
        if (0 < 42) {
            fmt.Println(v109, v4, v113)
        } else {
            switch v4 {
            case "":
                v5 = append(v5, 32767)
            }
            switch 9223372036854775807 {
            case 751:
                var v115 bool = v109
                _ = v115
            case 548:
                if (!(v109 && (!true))) {
                    v113 = v109
                    fmt.Println(v113)
                } else {
                    v4 = (((v4 + v4) + "a") + ((v4 + v4) + v4))
                    switch 32768 {
                    case -1:
                        if (2 == 15) {
                            _ = len(v5)
                        } else {
                            for fi116 := 0; fi116 < 12; fi116++ {
                                fmt.Println(v113)
                                _ = len(v3)
                                if (!v109) {
                                    var v117 uint8 = 0
                                    v109 = v109
                                    switch v4 {
                                    case "":
                                        v118 := GenS4{F0: 42}
                                        _ = v118
                                    default:
                                        for fi119 := 0; fi119 < 16; fi119++ {
                                            v120 := GenS16{F0: 587}
                                            _ = fi119
                                            _ = v120
                                        }
                                    }
                                    _ = v117
                                }
                                _ = fi116
                            }
                        }
                    case 32767:
                        for fi121 := 0; fi121 < 5; fi121++ {
                            _ = len(v5)
                            _ = fi121
                        }
                    }
                }
            case 32767:
                if v109 {
                    var v122 int = 9223372036854775807
                    _ = v122
                }
            }
            var v123 float32 = 0.1
            _ = v123
        }
    default:
        var v124 float32 = (0.1 + (float32((1.5 * 0.1)) * 0.1))
        _ = v124
    }
    var v125 int16 = -1
    var v126 int = 32768
    v2[(v126 - -456)] = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v113:", v113)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v109
    _ = v113
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
