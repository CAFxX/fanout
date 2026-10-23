package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 uint64; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 bool; F1 int; F2 uint32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint16; F1 string }
type GenS15 struct { F0 int32; F1 int16; F2 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 bool }
func genFn0(p0 string) int16 {
    v7 := []int{4294967295, -2147483648, 515, 42}
    p0 = p0
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = p0
    _ = v7
    return -32768
}
func genFn1() int {
    var v8 float32 = (((0.0 + 1.5) + ((1.5 - 3.14159) - 3.14159)) - -2.25)
    v9 := GenS1{F0: 42, F1: true, F2: 947}
    switch 418 {
    case -1:
        var v10 float64 = 0.1
        _ = v10
    case 2147483647:
        if true {
            v11 := GenS2{F0: 4294967295, F1: -2147483648, F2: 340}
            var v12 uint = 65535
            _ = v11
            _ = v12
        }
    case -592:
        var v13 int64 = 42
        _ = v13
    default:
        if ((false && ((!true) && false)) && (0 == 18446744073709551615)) {
            fmt.Println(v8)
        }
    }
    var v14 uint32 = 1
    _ = v14
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{1, uint32(v1), 147, 680}
    var v3 int32 = 42
    var v4 uint = 41
    var v5 string = "foo bar"
    var v6 int64 = -423
    v2 = append(v2, 2)
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 int64 = (-(v6))
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 32768, F1: (127 != 127), F2: 42}
        v19 := GenS3{F0: v5, F1: 52}
        _ = fi15
        _ = v18
        _ = v19
    }
    _ = len(v2)
    var v20 int8 = -128
    var v21 int16 = 1000
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS1{F0: fi22, F1: true, F2: v21}
        _ = fi22
        _ = v23
    }
    switch ("hello" + "hello") {
    case "foo bar":
        var v24 int = int(v21)
        _ = v24
    }
    v6 = (v6 + int64(v20))
    v25 := GenS4{F0: 42}
    v26 := GenS5{F0: (v4 ^ (v4 % ((42) | 1)))}
    if (((0.0 < 0.0) || (128 != -2147483648)) && (-1 <= (^(v21)))) {
        switch v5 {
        case "foo bar":
            if true {
                switch v5 {
                case "\n\t":
                    var v27 uint64 = 603
                    _ = v27
                case "foo bar":
                    v28 := GenS6{F0: v6}
                    _ = v28
                case "a":
                    var v29 int64 = (v6 / ((1) | 1))
                    _ = v29
                }
                v30 := GenS7{F0: ("hello" < v5), F1: 18446744073709551615, F2: v6}
                switch 0 {
                case -1:
                    if (0 < 18446744073709551615) {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            var v32 bool = (true || true)
                            var v33 int16 = genFn0((v5 + "hello"))
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        if true {
                            var v34 uint32 = 0
                            _ = v34
                        }
                    }
                case -2147483648:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 bool = true
                        _ = len(v2)
                        v37 := GenS8{F0: v1}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                case 127:
                    if ((!(v21 > (v21 | genFn0(v5)))) && true) {
                        switch 1 {
                        case 4294967295:
                            v6 -= v6
                        case 134:
                            var v38 uint32 = 4294967295
                            _ = v38
                        default:
                            var v39 int = 1
                            _ = v39
                        }
                        if (false || ((v4 != (v4 | v4)) && (int(v1) <= -459))) {
                            v3 = -1
                            var v40 float32 = float32((1.5 / float64(3.14159)))
                            _ = v40
                        }
                    } else {
                        if (0.1 >= -2.25) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v41 string = (("" + "a") + ("a" + v5))
                            switch "foo bar" {
                            case "":
                                if false {
                                    var v42 int32 = v3
                                    v5 = v41
                                    v2 = append(v2, 4294967295)
                                    _ = v42
                                }
                            case "foo bar":
                                v20 += (v20 ^ 0)
                            case "a":
                                v6 = 1
                            }
                            _ = v41
                        } else {
                            var v43 uint32 = 938
                            _ = v43
                        }
                        v5 = (v5 + (v5 + v5))
                        if (-587 <= -1) {
                            v3 = v3
                            v1 += v1
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((v6 + 445) <= v6) {
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    switch (((v5 + ("\n\t" + v5)) + (v5 + (v5 + v5))) + "hello") {
                                    case "\n\t":
                                        if false {
                                            if ((-9223372036854775808 > int64(v21)) || false) {
                                                for fi45 := 0; fi45 < 7; fi45++ {
                                                    switch "" {
                                                    case "hello":
                                                        var v46 int64 = 9223372036854775807
                                                        _ = v46
                                                    }
                                                    _ = fi45
                                                }
                                                var v47 int32 = (v3 ^ (v3 >> (((v4 & 204)) & 63)))
                                                _ = v47
                                            } else {
                                                var v48 bool = (!(!true))
                                                v3 = (v3 << ((v4) & 63))
                                                _ = v48
                                            }
                                            for fi49 := 0; fi49 < 13; fi49++ {
                                                var v50 uint16 = v1
                                                var v51 int32 = (^(v3))
                                                _ = len(v2)
                                                _ = fi49
                                                _ = v50
                                                _ = v51
                                            }
                                        }
                                    case "":
                                        for fi52 := 0; fi52 < 7; fi52++ {
                                            var v53 int64 = v6
                                            fmt.Println(v1, v20, v3)
                                            _ = fi52
                                            _ = v53
                                        }
                                    }
                                    _ = fi44
                                }
                            }
                        }
                    }
                }
                _ = v30
            }
        case "\n\t":
            for fi54 := 0; fi54 < 13; fi54++ {
                for fi55 := 0; fi55 < 10; fi55++ {
                    if ((180 >= (^(v4))) && true) {
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        v2 = append(v2, 65535)
                        fmt.Println(v3, v21, v21)
                        switch v5 {
                        case "hello":
                            v1 = v1
                        default:
                            if false {
                                fmt.Println(v4, v20, fi54)
                                v1 = v1
                            } else {
                                switch ("foo bar" + v5) {
                                case "foo bar":
                                    v56 := GenS7{F0: false, F1: 970, F2: int64(v1)}
                                    _ = v56
                                case "hello":
                                    v2 = append(v2, 2147483648)
                                case "\n\t":
                                    if false {
                                        var v57 string = "\n\t"
                                        _ = v57
                                    }
                                default:
                                    var v58 int64 = (v6 - 42)
                                    var v59 int8 = (v20 | v20)
                                    _ = v58
                                    _ = v59
                                }
                            }
                        }
                        v60 := GenS9{F0: v21}
                        _ = v60
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            if len(v2) > 0 {
                                v2[0] = 780
                            }
                        } else {
                            v5 = v5
                            v20 = v20
                            _ = len(v2)
                        }
                        var v61 float32 = 0.1
                        _ = v61
                    }
                    _ = fi55
                }
                _ = fi54
            }
        case "a":
            var v62 uint = v4
            _ = v62
        }
        switch -9223372036854775808 {
        case 4294967295:
            switch v5 {
            case "a":
                for fi63 := 0; fi63 < 19; fi63++ {
                    for fi64 := 0; fi64 < 15; fi64++ {
                        v3 -= -553
                        if (1.5 >= float64((float32(0.0) + 0.0))) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 -= 4294967295
                            var v65 uint64 = 415
                            _ = v65
                        } else {
                            v66 := GenS9{F0: v21}
                            switch fi64 {
                            case 128:
                                v67 := GenS9{F0: int16(v3)}
                                v6 = -217
                                var v68 float32 = 1e10
                                _ = v67
                                _ = v68
                            case -868:
                                _ = len(v2)
                            case -9223372036854775808:
                                v69 := GenS10{F0: 1, F1: false}
                                _ = v69
                            }
                            _ = v66
                        }
                        var v70 int = (fi63 - (fi63 | fi64))
                        _ = fi64
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 15; fi71++ {
                        var v72 string = (v5 + "")
                        _ = fi71
                        _ = v72
                    }
                    var v73 uint = v4
                    _ = fi63
                    _ = v73
                }
            }
        default:
            switch 42 {
            case 32768:
                switch v5 {
                case "":
                    for fi74 := 0; fi74 < 11; fi74++ {
                        if false {
                            for fi75 := 0; fi75 < 19; fi75++ {
                                switch genFn1() {
                                case -965:
                                    v2 = append(v2, 144)
                                case 0:
                                    switch fi75 {
                                    case 4294967295:
                                        fmt.Println(v5, v20)
                                    case 329:
                                        var v76 uint64 = 18446744073709551615
                                        _ = v76
                                    case 797:
                                        v3 -= (^(v3))
                                    default:
                                        var v77 int64 = -9223372036854775808
                                        _ = v77
                                    }
                                    if ("foo bar" > v5) {
                                        switch fi75 {
                                        case 585:
                                            var v78 int = fi75
                                            _ = v78
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = 744
                                            }
                                        }
                                        if true {
                                            fmt.Println(v20)
                                            var v79 float64 = -2.25
                                            _ = v79
                                        }
                                        fmt.Println(v6, fi75)
                                    }
                                    v3 = 0
                                case 975:
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                }
                                fmt.Println(v21, v4, v4)
                                _ = fi75
                            }
                        }
                        if (!true) {
                            if (fi74 != genFn1()) {
                                if (4294967295 != 0) {
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        var v81 uint32 = 0
                                        _ = fi80
                                        _ = v81
                                    }
                                    if true {
                                        var v82 int = fi74
                                        _ = v82
                                    } else {
                                        v83 := GenS11{F0: -1, F1: (v4 >> ((v4) & 63))}
                                        _ = v83
                                    }
                                    switch genFn1() {
                                    case -113:
                                        switch (v5 + "foo bar") {
                                        case "":
                                            fmt.Println(v1, v20)
                                        case "foo bar":
                                            if false {
                                                switch (fi74 | (fi74 * fi74)) {
                                                case 32768:
                                                    v84 := GenS12{F0: (!(!true)), F1: fi74, F2: uint32(v1)}
                                                    _ = v84
                                                case 711:
                                                    v85 := GenS13{F0: (v21 * (v21 - 1000))}
                                                    _ = v85
                                                case -956:
                                                    v86 := GenS12{F0: (0 <= 1), F1: int(v3), F2: 0}
                                                    _ = v86
                                                default:
                                                    var v87 uint32 = 1
                                                    _ = v87
                                                }
                                            }
                                        default:
                                            var v88 int = fi74
                                            _ = v88
                                        }
                                    case 128:
                                        fmt.Println(v6, v1, v4)
                                    case 4294967295:
                                        v21 -= v21
                                    }
                                } else {
                                    switch (v5 + "foo bar") {
                                    case "hello":
                                        v21 = v21
                                    case "":
                                        v3 = v3
                                    case "foo bar":
                                        v2 = append(v2, 2)
                                    default:
                                        fmt.Println(v1)
                                    }
                                }
                                switch ((v5 + ("a" + v5)) + v5) {
                                case "\n\t":
                                    if len(v2) > 0 {
                                        v2[0] = 350
                                    }
                                default:
                                    v89 := GenS1{F0: genFn1(), F1: true, F2: (v21 ^ -32768)}
                                    _ = v89
                                }
                                var v90 int = genFn1()
                                _ = v90
                            }
                            v21 = genFn0(v5)
                            for fi91 := 0; fi91 < 15; fi91++ {
                                v2 = append(v2, 17)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi91
                            }
                        } else {
                            if (4294967295 > 42) {
                                v92 := GenS14{F0: v1, F1: ((("a" + (v5 + v5)) + "") + ("foo bar" + "foo bar"))}
                                for fi93 := 0; fi93 < 16; fi93++ {
                                    fmt.Println(v4, v4)
                                    var v94 float32 = 1.5
                                    v1 -= uint16(v4)
                                    _ = fi93
                                    _ = v94
                                }
                                var v95 uint = v4
                                _ = v92
                                _ = v95
                            }
                            for fi96 := 0; fi96 < 5; fi96++ {
                                var v97 uint64 = 471
                                fmt.Println(v6)
                                switch 32767 {
                                case -1:
                                    var v98 int32 = (^(v3))
                                    v99 := GenS15{F0: v3, F1: (v21 * (v21 - v21)), F2: (v21 ^ (v21 | -525))}
                                    v21 += v21
                                    _ = v98
                                    _ = v99
                                case -9223372036854775808:
                                    v2 = append(v2, 2147483648)
                                }
                                fmt.Println(v97, fi74, v4)
                                _ = fi96
                                _ = v97
                            }
                            v100 := GenS6{F0: (v6 + 42)}
                            _ = v100
                        }
                        var v101 string = "hello"
                        _ = fi74
                        _ = v101
                    }
                case "a":
                    if false {
                        if (1e10 <= (0.1 * ((1e10 - -2.25) / -2.25))) {
                            v102 := GenS8{F0: v1}
                            _ = v102
                        }
                        v103 := GenS8{F0: v1}
                        v21 = v21
                        _ = v103
                    }
                default:
                    v104 := GenS16{F0: 4294967295}
                    _ = v104
                }
            case -1:
                var v105 string = v5
                _ = v105
            case 725:
                switch -9223372036854775808 {
                case 0:
                    var v106 uint32 = 2147483647
                    _ = v106
                default:
                    v107 := GenS7{F0: ((v6 ^ -1) != 1), F1: 667, F2: (v6 >> ((v4) & 63))}
                    _ = v107
                }
            }
        }
        for fi108 := 0; fi108 < 11; fi108++ {
            if (!(false || false)) {
                if (v1 >= (v1 * (v1 | v1))) {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    fmt.Println(v4, v20)
                    for fi109 := 0; fi109 < 3; fi109++ {
                        v3 = 150
                        var v110 uint = (v4 + 255)
                        var v111 int64 = v6
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                }
                if false {
                    fmt.Println(v3)
                    var v112 string = (v5 + "hello")
                    _ = v112
                } else {
                    if true {
                        var v113 int16 = v21
                        switch v5 {
                        case "":
                            var v114 bool = (0 < 623)
                            _ = v114
                        }
                        var v115 uint = 658
                        _ = v113
                        _ = v115
                    }
                    _ = len(v2)
                }
                if (!(127 != 255)) {
                    v20 = 1
                } else {
                    for fi116 := 0; fi116 < 5; fi116++ {
                        switch fi116 {
                        case -214:
                            var v117 float32 = (1e10 / -2.25)
                            _ = v117
                        case 13:
                            v20 = v20
                        case -793:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if true {
                            var v118 bool = true
                            fmt.Println(v21)
                            _ = v118
                        }
                        v119 := GenS17{F0: false}
                        v2 = append(v2, 4294967295)
                        _ = fi116
                        _ = v119
                    }
                }
            }
            var v120 int16 = (v21 % ((-701) | 1))
            v20 = 127
            _ = fi108
            _ = v120
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
