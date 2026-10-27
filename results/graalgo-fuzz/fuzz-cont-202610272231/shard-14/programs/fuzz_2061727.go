package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 bool; F2 int16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint; F1 int32; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 uint64; F1 int32; F2 int64 }
type GenS12 struct { F0 int; F1 int8; F2 int16 }
type GenS13 struct { F0 int64; F1 uint8 }
type GenS14 struct { F0 bool; F1 string; F2 uint8 }
type GenS15 struct { F0 bool; F1 int8; F2 string }
func genFn0(p0 bool, p1 string) uint8 {
    v6 := GenS1{F0: 255, F1: -9223372036854775808, F2: 324}
    p0 = ((3.14159 <= -2.25) || p0)
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1() uint {
    var v7 int8 = 100
    v7 = (v7 + (v7 | -128))
    _ = v7
    return 65535
}

func main() {
    v1 := []string{"", (("foo bar" + ("a" + ("hello" + "hello"))) + "\n\t"), "foo bar"}
    v2 := make(map[string]uint32)
    var v3 float32 = (3.14159 - 0.1)
    var v4 uint64 = 18446744073709551615
    var v5 uint = 0
    switch ("a" + "a") {
    case "":
        if true {
            _ = len(v2)
        } else {
            var v8 uint = genFn1()
            v9 := GenS2{F0: 1, F1: 127}
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 int8 = 1
                v12 := GenS3{F0: -1, F1: true}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
    case "foo bar":
        var v13 int8 = 100
        _ = v13
    }
    delete(v2, "\n\t")
    switch 42 {
    case -947:
        if (-226 != 1) {
            v3 -= v3
        }
    case 32768:
        var v14 int16 = 32767
        _ = v14
    }
    v4 += 42
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v5, v15)
        _ = len(v2)
        _ = fi16
    }
    switch "foo bar" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            for fi17 := 0; fi17 < 13; fi17++ {
                fmt.Println(v3)
                if ("" != "foo bar") {
                    if (false || false) {
                        var v18 string = ""
                        _ = len(v1)
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            v15 -= v15
                            if true {
                                fmt.Println(v4, v15)
                                var v20 uint8 = genFn0(false, "hello")
                                _ = v20
                            } else {
                                var v21 uint16 = v15
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v22 uint = v5
                                _ = v21
                                _ = v22
                            }
                            if (v3 < 0.1) {
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(fi19, fi19)
                                case "a":
                                    v23 := v2["hello"]
                                    v24 := GenS3{F0: 163, F1: (((65535 == 2147483647) || (0 == 1)) && (127 != 2))}
                                    if true {
                                        var v25 float32 = 0.0
                                        _ = v25
                                    } else {
                                        fmt.Println(fi17)
                                        var v26 int = int(v23)
                                        var v27 uint64 = v4
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v23
                                    _ = v24
                                default:
                                    if (!((58 != 2147483647) && true)) {
                                        v5 = (v5 ^ (v5 ^ 478))
                                        v5 = genFn1()
                                        var v28 int32 = 2147483647
                                        switch -9223372036854775808 {
                                        case 2147483647:
                                            fmt.Println(fi19, fi19)
                                        }
                                        var v29 uint64 = v4
                                        _ = v28
                                        _ = v29
                                    }
                                    v30 := GenS1{F0: 2, F1: fi19, F2: 0}
                                    _ = v30
                                }
                                var v31 uint16 = v15
                                fmt.Println(v15, v4, fi17)
                                v32 := GenS4{F0: true}
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(v15, v3, v3)
                                var v33 bool = false
                                _ = v33
                            }
                            _ = fi19
                        }
                        var v34 uint64 = (^(v4))
                        _ = v34
                    }
                    var v35 uint64 = (^(v4))
                    _ = v35
                } else {
                    switch "" {
                    case "a":
                        switch (fi17 << ((v5) & 63)) {
                        case 79:
                            switch ("\n\t" + ((("hello" + "foo bar") + "\n\t") + "hello")) {
                            case "foo bar":
                                switch "\n\t" {
                                case "foo bar":
                                    var v36 string = "foo bar"
                                    _ = v36
                                default:
                                    var v37 float64 = 0.1
                                    _ = v37
                                }
                            case "":
                                v4 -= v4
                            case "\n\t":
                                v38 := GenS3{F0: -1, F1: true}
                                _ = v38
                            }
                        case 592:
                            fmt.Println(fi17)
                        }
                    default:
                        fmt.Println(v3, v4, v4)
                    }
                    if true {
                        for fi39 := 0; fi39 < 9; fi39++ {
                            _ = len(v2)
                            if (v4 > 128) {
                                v40 := GenS5{F0: (v5 >> (((v5 % (((v5 >> ((v5) & 63))) | 1))) & 63)), F1: false, F2: -947}
                                _ = v40
                            } else {
                                if true {
                                    if true {
                                        switch (fi39 + (-(fi17))) {
                                        case -742:
                                            var v41 uint8 = 127
                                            _ = v41
                                        case -261:
                                            _ = len(v1)
                                        }
                                        v15 -= (v15 << (((v5 & v5)) & 63))
                                    } else {
                                        if (-1 <= 2147483647) {
                                            fmt.Println(v15, fi39)
                                            fmt.Println(v5)
                                            if ("foo bar" > "") {
                                                if len(v1) > 0 {
                                                    v1[0] = ""
                                                }
                                            }
                                        }
                                        v4 = v4
                                    }
                                    v4 -= (v4 ^ 0)
                                }
                            }
                            v3 += v3
                            _ = fi39
                        }
                        v4 += 42
                        v2[("\n\t" + "hello")] = 939
                    } else {
                        if (true && (!false)) {
                            for fi42 := 0; fi42 < 10; fi42++ {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    fmt.Println(fi17, fi42)
                                    v43 := GenS4{F0: (!true)}
                                    if true {
                                        v5 += 18446744073709551615
                                    }
                                    v44 := GenS6{F0: 13, F1: fi42}
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi42
                            }
                        }
                    }
                }
                fmt.Println(v4, v4)
                _ = fi17
            }
        case "foo bar":
            for fi45 := 0; fi45 < 4; fi45++ {
                fmt.Println(fi45)
                if (float32(float64(v3)) > v3) {
                    switch fi45 {
                    case 1:
                        var v46 uint64 = (v4 * v4)
                        _ = v46
                    case 32768:
                        v47 := GenS3{F0: 0, F1: true}
                        _ = v47
                    case 42:
                        v15 -= 2
                    }
                } else {
                    delete(v2, "")
                    for fi48 := 0; fi48 < 2; fi48++ {
                        if (1000 > 1000) {
                            v49 := GenS4{F0: (!true)}
                            _ = v49
                        } else {
                            switch fi48 {
                            case 32767:
                                fmt.Println(fi48, v3)
                            default:
                                switch (fi45 >> (((^(v5))) & 63)) {
                                case -9223372036854775808:
                                    v5 = v5
                                    var v50 bool = false
                                    var v51 uint8 = genFn0(v50, ("\n\t" + "a"))
                                    _ = v50
                                    _ = v51
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v5 -= (v5 + (v5 | 65535))
                        v3 -= ((((v3 / 0.1) + v3) + (v3 * 0.1)) - (v3 * float32(1.5)))
                        _ = fi48
                    }
                    v3 = v3
                }
                if (251 == -418) {
                    switch "hello" {
                    case "\n\t":
                        for fi52 := 0; fi52 < 20; fi52++ {
                            var v53 bool = false
                            v4 -= v4
                            _ = fi52
                            _ = v53
                        }
                    default:
                        v54 := GenS6{F0: 1, F1: 42}
                        _ = v54
                    }
                    for fi55 := 0; fi55 < 1; fi55++ {
                        var v56 int64 = 42
                        if false {
                            if true {
                                v56 = 1
                            } else {
                                _ = len(v1)
                            }
                            v57 := GenS7{F0: genFn1(), F1: -2147483648, F2: 1000}
                            _ = v57
                        }
                        v58 := GenS4{F0: (!((!(fi55 <= -444)) && false))}
                        _ = fi55
                        _ = v56
                        _ = v58
                    }
                    delete(v2, ("hello" + "a"))
                }
                _ = fi45
            }
        }
    }
    v4 += 859
    _ = len(v1)
    switch (("hello" + "foo bar") + ("" + "")) {
    case "a":
        v4 += v4
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        if true {
            if true {
                if (!false) {
                    var v60 int16 = int16(v15)
                    _ = v60
                }
                var v61 int8 = -1
                v61 = (v61 - (v61 << (((v5 | 18446744073709551615)) & 63)))
                _ = v61
            }
        } else {
            for fi62 := 0; fi62 < 5; fi62++ {
                if false {
                    v4 += 42
                    if ((fi62 ^ -2147483648) >= fi59) {
                        switch (fi62 >> ((v5) & 63)) {
                        case 42:
                            fmt.Println(fi62, fi59, fi62)
                        case 634:
                            if true {
                                v63 := GenS8{F0: v15, F1: 0, F2: fi59}
                                _ = v63
                            } else {
                                v3 += 3.14159
                                if (!false) {
                                    var v64 int16 = -32768
                                    v65 := v2[("" + (((("hello" + "") + "\n\t") + "hello") + ""))]
                                    _ = v64
                                    _ = v65
                                }
                                var v66 uint = v5
                                _ = v66
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        if (42 > 127) {
                            switch ("" + "\n\t") {
                            case "":
                                v1 = append(v1, "\n\t")
                            case "\n\t":
                                v4 -= v4
                            case "a":
                                var v67 int16 = 1000
                                _ = v67
                            default:
                                delete(v2, "hello")
                            }
                            v15 += uint16(v4)
                        } else {
                            fmt.Println(v3, v5)
                            if false {
                                v3 = v3
                                v68 := GenS9{F0: -9223372036854775808}
                                _ = v68
                            }
                            var v69 uint8 = 1
                            _ = v69
                        }
                        switch "" {
                        case "a":
                            switch "a" {
                            case "hello":
                                var v70 int8 = 0
                                v4 += v4
                                v71 := GenS2{F0: 2, F1: 59}
                                var v72 uint = uint(v15)
                                v73 := GenS4{F0: (!(!true))}
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                            case "foo bar":
                                var v74 int8 = 125
                                _ = v74
                            case "a":
                                v1 = append(v1, "")
                            }
                        case "":
                            var v75 uint32 = 2147483648
                            _ = v75
                        default:
                            var v76 uint64 = 2
                            _ = v76
                        }
                        switch "hello" {
                        case "a":
                            var v77 int32 = 42
                            _ = v77
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "":
                            switch (fi62 % ((640) | 1)) {
                            case -490:
                                v4 = v4
                                v78 := GenS8{F0: v15, F1: 0, F2: (fi59 | fi62)}
                                v4 = v4
                                _ = v78
                            case 42:
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            default:
                                var v79 uint = (v5 << ((genFn1()) & 63))
                                _ = v79
                            }
                        default:
                            v5 += (v5 | 18446744073709551615)
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v3 = v3
                    if false {
                        switch (fi59 % ((int(v15)) | 1)) {
                        case -738:
                            v4 = v4
                        }
                    } else {
                        if (928 >= 0) {
                            if (!((0.1 * -2.25) > ((1e10 / -2.25) / 0.1))) {
                                v4 += (^(v4))
                                var v80 int16 = 1
                                var v81 float64 = float64(v3)
                                _ = v80
                                _ = v81
                            } else {
                                switch (("" + ("foo bar" + "")) + "") {
                                case "a":
                                    var v82 float64 = (-2.25 + float64(v3))
                                    _ = v82
                                case "hello":
                                    if (!true) {
                                        fmt.Println(v15, v4, v15)
                                        v5 -= (v5 + 0)
                                        _ = len(v1)
                                    } else {
                                        var v83 string = ""
                                        var v84 float32 = (v3 * v3)
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                fmt.Println(fi62)
                                v15 = (v15 << ((255) & 63))
                            }
                            v4 += (v4 * (v4 & v4))
                            v5 = (v5 / (((^(v5))) | 1))
                        }
                        v15 = 255
                        var v85 int = int(v5)
                        _ = v85
                    }
                } else {
                    var v86 uint8 = 74
                    v15 = (v15 & v15)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v86
                }
                v1 = append(v1, "a")
                _ = fi62
            }
            _ = len(v2)
        }
        var v87 uint32 = 749
        var v88 int8 = 127
        _ = fi59
        _ = v87
        _ = v88
    }
    v89 := GenS10{F0: 42, F1: 0}
    var v90 float32 = v3
    switch -634 {
    case 0:
        var v91 uint16 = (^(v15))
        _ = v91
    case 32767:
        var v92 int = -9223372036854775808
        _ = v92
    case -705:
        for fi93 := 0; fi93 < 13; fi93++ {
            var v94 bool = (uint(v4) != v5)
            var v95 uint32 = 480
            _ = fi93
            _ = v94
            _ = v95
        }
    default:
        if (v5 > (v5 % ((v5) | 1))) {
            var v96 uint8 = uint8(v5)
            for fi97 := 0; fi97 < 3; fi97++ {
                var v98 uint64 = v4
                v3 -= v90
                var v99 uint64 = v98
                _ = fi97
                _ = v98
                _ = v99
            }
            _ = v96
        } else {
            switch "\n\t" {
            case "\n\t":
                var v100 float32 = 1e10
                _ = v100
            default:
                v3 = v3
            }
        }
    }
    for fi101 := 0; fi101 < 4; fi101++ {
        var v102 int8 = 42
        switch "" {
        case "a":
            switch -2147483648 {
            case 238:
                var v103 float32 = -2.25
                _ = v103
            }
        case "":
            if false {
                for fi104 := 0; fi104 < 16; fi104++ {
                    v102 += (v102 + v102)
                    if ((2 < 798) || false) {
                        if false {
                            fmt.Println(v102, v5, v4)
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                        }
                        v15 = 1
                    } else {
                        v105 := v2["\n\t"]
                        _ = v105
                    }
                    _ = fi104
                }
                var v106 int = (fi101 % ((2147483647) | 1))
                var v107 int8 = -78
                _ = v106
                _ = v107
            } else {
                var v108 int64 = 42
                v90 = 1.5
                if false {
                    var v109 int8 = 0
                    for fi110 := 0; fi110 < 12; fi110++ {
                        if (-1 == 42) {
                            v1 = append(v1, ("\n\t" + "hello"))
                            if false {
                                switch fi101 {
                                case -2147483648:
                                    v111 := GenS4{F0: true}
                                    _ = v111
                                case 128:
                                    switch (^(fi110)) {
                                    case 32767:
                                        if true {
                                            if (false || (!false)) {
                                                var v112 float32 = (float32(((3.14159 / 0.1) / ((1e10 * -2.25) - 1e10))) / (-2.25 + (0.0 - (v90 * 1e10))))
                                                if true {
                                                    if true {
                                                        fmt.Println(v112, v15)
                                                        var v113 float64 = 1e10
                                                        v114 := GenS4{F0: true}
                                                        _ = v113
                                                        _ = v114
                                                    } else {
                                                        v2["\n\t"] = 281
                                                        var v115 uint32 = 2
                                                        _ = v115
                                                    }
                                                }
                                                v116 := GenS7{F0: v5, F1: 148, F2: 1000}
                                                fmt.Println(v4, v112)
                                                _ = v112
                                                _ = v116
                                            }
                                            v102 = v109
                                            if true {
                                                fmt.Println(fi110)
                                                if false {
                                                    v15 += v15
                                                } else {
                                                    v2[("\n\t" + ("hello" + ("\n\t" + "hello")))] = uint32(v4)
                                                    var v117 bool = (!true)
                                                    v4 = 42
                                                    _ = v117
                                                }
                                            }
                                            if true {
                                                v90 -= (v90 * 1e10)
                                                if len(v1) > 0 {
                                                    v1[0] = (((("" + "") + "a") + (("foo bar" + "") + "")) + "foo bar")
                                                }
                                            } else {
                                                v118 := GenS10{F0: -2147483648, F1: -86}
                                                _ = v118
                                            }
                                            v102 = -128
                                        }
                                    default:
                                        fmt.Println(v109, v5)
                                        v109 = v109
                                        var v119 int8 = v102
                                        var v120 float32 = (v3 / 1.0)
                                        _ = v119
                                        _ = v120
                                    }
                                case 2147483647:
                                    var v121 int64 = -9223372036854775808
                                    switch fi101 {
                                    case 15:
                                        switch (^(fi110)) {
                                        case 948:
                                            if false {
                                                v122 := GenS11{F0: (^(v4)), F1: -1, F2: (v108 + 0)}
                                                var v123 uint = (v5 * (v5 + v5))
                                                if (4294967295 <= 65535) {
                                                    var v124 int32 = int32(v4)
                                                    _ = v124
                                                }
                                                _ = v122
                                                _ = v123
                                            } else {
                                                v109 -= -128
                                                _ = len(v1)
                                            }
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                    fmt.Println(v3, v3)
                                    _ = v121
                                }
                                var v125 uint16 = (^(v15))
                                v90 -= v90
                                _ = v125
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            case "foo bar":
                                v126 := GenS6{F0: 1, F1: (fi101 / (((fi110 << ((v5) & 63))) | 1))}
                                _ = v126
                            case "":
                                var v127 string = "a"
                                _ = v127
                            default:
                                var v128 float64 = float64(v90)
                                _ = v128
                            }
                        } else {
                            fmt.Println(v3, v5, v4)
                            switch (fi101 * 0) {
                            case 430:
                                switch fi101 {
                                case -231:
                                    v129 := GenS9{F0: -1}
                                    _ = v129
                                case 151:
                                    v5 += genFn1()
                                }
                            case 0:
                                v15 -= v15
                            case 743:
                                fmt.Println(v102)
                            default:
                                switch (^(fi101)) {
                                case 9223372036854775807:
                                    v102 = int8(v5)
                                case 128:
                                    if false {
                                        var v130 uint8 = genFn0(false, "hello")
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v130
                                    }
                                case -2147483648:
                                    if true {
                                        if true {
                                            fmt.Println(v90, v4)
                                            if (127 > 127) {
                                                _ = len(v1)
                                                switch (-(fi101)) {
                                                case -1:
                                                    var v131 int32 = 42
                                                    _ = v131
                                                case 32768:
                                                    fmt.Println(v5, v108, v5)
                                                case 2147483647:
                                                    var v132 int16 = 429
                                                    _ = v132
                                                }
                                            } else {
                                                v102 = -92
                                                fmt.Println(fi101, fi101, v109)
                                                var v133 uint8 = 0
                                                _ = v133
                                            }
                                        } else {
                                            var v134 uint = (v5 + v5)
                                            delete(v2, "hello")
                                            var v135 float64 = 1.5
                                            if ((v15 % ((582) | 1)) != v15) {
                                                var v136 uint = v5
                                                if ((true || false) && (("a" + "") > "hello")) {
                                                    switch -2147483648 {
                                                    case 127:
                                                        v137 := GenS12{F0: fi101, F1: (v102 << (((v5 / (((^(v134))) | 1))) & 63)), F2: 1000}
                                                        _ = v137
                                                    default:
                                                        var v138 uint16 = v15
                                                        _ = v138
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    var v139 uint16 = (v15 + v15)
                                                    if true {
                                                        v140 := GenS13{F0: v108, F1: 128}
                                                        _ = v140
                                                    }
                                                    v139 += (^(v139))
                                                    _ = v139
                                                }
                                                _ = v136
                                            }
                                            _ = v134
                                            _ = v135
                                        }
                                        v15 = (v15 | (v15 & (v15 + v15)))
                                    } else {
                                        v141 := GenS14{F0: (946 < 109), F1: "", F2: genFn0((false && false), ("" + "a"))}
                                        if ((0 != 4294967295) || ((v15 << ((v5) & 63)) > 79)) {
                                            fmt.Println(v4, fi110, v90)
                                            var v142 float64 = 1e10
                                            v108 = (v108 ^ v108)
                                            _ = v142
                                        }
                                        var v143 int64 = v108
                                        _ = v141
                                        _ = v143
                                    }
                                }
                            }
                            switch "a" {
                            case "a":
                                v3 = 1e10
                                var v144 float32 = v3
                                _ = v144
                            case "":
                                var v145 uint32 = 4294967295
                                _ = v145
                            default:
                                v108 += (v108 * v108)
                            }
                            v4 -= v4
                        }
                        _ = fi110
                    }
                    _ = len(v2)
                    _ = v109
                }
                _ = v108
            }
        case "hello":
            v146 := GenS15{F0: false, F1: v102, F2: "a"}
            _ = v146
        }
        _ = fi101
        _ = v102
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v90
}
