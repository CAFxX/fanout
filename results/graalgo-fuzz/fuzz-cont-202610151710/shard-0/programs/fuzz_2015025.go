package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 int16; F2 int }
type GenS12 struct { F0 string; F1 int16; F2 int16 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16; F1 uint32 }
type GenS15 struct { F0 uint64; F1 bool }
type GenS16 struct { F0 uint8; F1 uint8 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 bool; F1 int64; F2 uint8 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint32; F1 int16; F2 int16 }
type GenS21 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS22 struct { F0 string; F1 string }
type GenS23 struct { F0 bool; F1 bool; F2 string }
type GenS24 struct { F0 int; F1 uint32; F2 int64 }
type GenS25 struct { F0 int; F1 uint }
type GenS26 struct { F0 bool; F1 uint64 }
func genFn0(p0 bool, p1 uint32) bool {
    switch -382 {
    case -2147483648:
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := []int{(fi5 ^ 0), -549, (fi5 ^ int(p1)), fi5}
            v7 := GenS1{F0: 128, F1: 42}
            v8 := GenS1{F0: 0, F1: 0}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    case -954:
        p1 -= (^(p1))
    default:
        v9 := []int{}
        _ = v9
    }
    fmt.Println(p1)
    var v10 uint64 = 18446744073709551615
    var v11 int8 = -54
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p0
}
func genFn1(p0 int16, p1 int8, p2 int) uint8 {
    for fi12 := 0; fi12 < 9; fi12++ {
        if true {
            var v13 bool = false
            _ = v13
        } else {
            v14 := GenS1{F0: 1, F1: 0}
            v15 := GenS2{F0: -1, F1: 483}
            p1 += p1
            _ = v14
            _ = v15
        }
        fmt.Println(fi12, p2)
        _ = fi12
    }
    switch ("hello" + ("foo bar" + "foo bar")) {
    case "a":
        var v16 int = (-(p2))
        _ = v16
    case "foo bar":
        if ((p0 % (((p0 * int16(p2))) | 1)) == p0) {
            v17 := GenS1{F0: 127, F1: 1}
            _ = v17
        } else {
            fmt.Println(p2, p2)
        }
    case "\n\t":
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: 42, F1: 2}
            var v20 int64 = -935
            var v21 int32 = -646
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    default:
        var v22 int32 = 2147483647
        _ = v22
    }
    var v23 int32 = int32(p1)
    if false {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS3{F0: 127}
            p2 += -880
            v26 := GenS3{F0: 42}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 196
}

func main() {
    v1 := []uint32{}
    var v2 int8 = -128
    var v3 string = "a"
    v4 := make(map[string]uint32)
    v3 = v3
    switch -2147483648 {
    case 127:
        for fi27 := 0; fi27 < 3; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                v3 = v3
                v2 += v2
                _ = fi28
            }
            if (!false) {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS4{F0: (65535 >= 0), F1: v3, F2: genFn1(1, (v2 >> ((42) & 63)), 32768)}
                    _ = fi29
                    _ = v30
                }
            } else {
                fmt.Println(v3)
                v3 = ((((v3 + "\n\t") + "a") + "a") + (v3 + "a"))
            }
            _ = fi27
        }
    case -2147483648:
        v31 := GenS5{F0: (1 < 42), F1: 0}
        _ = v31
    default:
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint8 = 255
            v33 = v33
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        if (-1 < -32768) {
            if (false && (false && true)) {
                if false {
                    v2 += (v2 & v2)
                    if true {
                        var v35 uint = 1
                        _ = v35
                    }
                } else {
                    var v36 uint = 651
                    var v37 string = ("" + v3)
                    var v38 uint64 = 42
                    _ = v36
                    _ = v37
                    _ = v38
                }
                v39 := GenS4{F0: true, F1: ("foo bar" + "a"), F2: 0}
                _ = v39
            } else {
                v40 := GenS6{F0: 0}
                var v41 int64 = 907
                fmt.Println(v3, v41, fi34)
                _ = v40
                _ = v41
            }
        }
        _ = fi34
    }
    if true {
        var v42 uint = 476
        var v43 uint = v42
        v44 := GenS2{F0: 136, F1: (v42 / ((v43) | 1))}
        _ = v42
        _ = v43
        _ = v44
    }
    var v45 int8 = -128
    v46 := v4[v3]
    v46 = v46
    if true {
        for fi47 := 0; fi47 < 10; fi47++ {
            var v48 int64 = 42
            var v49 uint16 = 65535
            delete(v4, (v3 + v3))
            _ = fi47
            _ = v48
            _ = v49
        }
        if true {
            if false {
                var v50 uint8 = 41
                for fi51 := 0; fi51 < 12; fi51++ {
                    var v52 uint8 = v50
                    for fi53 := 0; fi53 < 5; fi53++ {
                        var v54 float64 = 0.1
                        if ("" < (v3 + ((v3 + (v3 + "")) + v3))) {
                            if (!true) {
                                var v55 uint32 = 42
                                switch ("\n\t" + "hello") {
                                case "":
                                    delete(v4, v3)
                                case "foo bar":
                                    switch 42 {
                                    case -283:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 803:
                                        var v56 bool = false
                                        _ = v56
                                    }
                                case "\n\t":
                                    var v57 int8 = v2
                                    _ = v57
                                default:
                                    v2 = (v45 | v45)
                                    var v58 int = -245
                                    v52 += v50
                                    _ = v58
                                }
                                if false {
                                    v3 = ((v3 + "a") + v3)
                                }
                                v4[v3] = v55
                                v59 := GenS4{F0: false, F1: v3, F2: uint8(v55)}
                                _ = v55
                                _ = v59
                            } else {
                                switch ((v3 + "\n\t") + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v2, v52)
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "hello":
                                    v60 := GenS1{F0: genFn1(-32768, (v2 >> ((2) & 63)), (-(fi53))), F1: 2}
                                    _ = v60
                                default:
                                    var v61 bool = ((-1 >= v2) || false)
                                    _ = v61
                                }
                                fmt.Println(v54, v52, fi53)
                            }
                            v50 += (v50 << ((994) & 63))
                        } else {
                            switch v3 {
                            case "":
                                v62 := GenS7{F0: -1, F1: -1}
                                _ = v62
                            case "hello":
                                v54 = float64(1.5)
                            case "foo bar":
                                var v63 int = fi53
                                _ = v63
                            }
                            switch "a" {
                            case "\n\t":
                                var v64 bool = (177 > 0)
                                _ = v64
                            case "":
                                if true {
                                    switch int(v46) {
                                    case -2147483648:
                                        fmt.Println(v2, v46)
                                    case 127:
                                        switch fi53 {
                                        case -2147483648:
                                            v65 := GenS4{F0: (255 == 65535), F1: (v3 + ("\n\t" + v3)), F2: (v50 * genFn1(-1, (v2 * (v2 & v2)), (fi51 % ((796) | 1))))}
                                            v66 := GenS8{F0: -2147483648, F1: 1000}
                                            _ = v65
                                            _ = v66
                                        case 1:
                                            var v67 int = fi53
                                            _ = v67
                                        default:
                                            v50 = genFn1(42, v2, (fi53 / ((fi53) | 1)))
                                        }
                                    case -564:
                                        if (42 > 2) {
                                            fmt.Println(v46)
                                        } else {
                                            v45 = v45
                                            var v68 uint16 = 27
                                            fmt.Println(v68)
                                            if len(v1) > 0 {
                                                v1[0] = (v46 ^ (v46 | v46))
                                            }
                                            v68 = (v68 & uint16(v46))
                                            if (true && ((fi51 * fi53) < 436)) {
                                                fmt.Println(fi51, v68)
                                                var v69 uint16 = v68
                                                if (false || true) {
                                                    if false {
                                                        v70 := GenS4{F0: true, F1: "foo bar", F2: 0}
                                                        v50 += v52
                                                        v71 := GenS8{F0: int32(v52), F1: 0}
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                                var v72 int = (-(fi53))
                                                v45 -= v45
                                                _ = v69
                                                _ = v72
                                            } else {
                                                if false {
                                                    fmt.Println(v52, v2, v2)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v73 := GenS8{F0: -1, F1: 42}
                                                    _ = v73
                                                }
                                                var v74 uint = uint(v46)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if false {
                                                        v74 = v74
                                                        v46 -= (v46 | uint32(v68))
                                                        if false {
                                                            var v75 float32 = (1e10 * float32((v54 - float64(-2.25))))
                                                            fmt.Println(v75, v50, v50)
                                                            _ = v75
                                                        }
                                                        if false {
                                                            fmt.Println(v2, v2, v52)
                                                            switch ((v3 + ("hello" + (((((v3 + v3) + (v3 + "a")) + (v3 + ((v3 + ("hello" + v3)) + v3))) + (v3 + v3)) + ("a" + "\n\t")))) + "hello") {
                                                            case "hello":
                                                                if false {
                                                                    if ((v3 + v3) > v3) {
                                                                        fmt.Println(v46, v50, v68)
                                                                        v76 := GenS4{F0: ((v50 << ((v74) & 63)) >= (^(v50))), F1: v3, F2: 128}
                                                                        _ = v76
                                                                    } else {
                                                                        var v77 uint64 = 269
                                                                        var v78 uint64 = (v77 + 793)
                                                                        if true {
                                                                            v79 := GenS9{F0: (v3 + "")}
                                                                            var v80 float64 = v54
                                                                            v2 = 0
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                    v81 := GenS10{F0: 112}
                                                                    var v82 uint32 = v46
                                                                    var v83 int32 = -611
                                                                    _ = v81
                                                                    _ = v82
                                                                    _ = v83
                                                                } else {
                                                                    v84 := GenS11{F0: (v46 - v46), F1: 1, F2: fi51}
                                                                    _ = v84
                                                                }
                                                            default:
                                                                var v85 string = "a"
                                                                _ = v85
                                                            }
                                                            v86 := GenS11{F0: (v46 / (((v46 << ((v74) & 63))) | 1)), F1: 1, F2: int(v52)}
                                                            _ = v86
                                                        }
                                                    }
                                                case "a":
                                                    var v87 int64 = -9223372036854775808
                                                    _ = v87
                                                case "hello":
                                                    fmt.Println(v74, v74, fi51)
                                                default:
                                                    if false {
                                                        var v88 uint8 = uint8(v46)
                                                        if true {
                                                            if (!true) {
                                                                if (!true) {
                                                                    var v89 uint8 = genFn1(32767, (v2 - (v2 + (^(v2)))), fi51)
                                                                    _ = v89
                                                                } else {
                                                                    switch (fi51 >> ((4294967295) & 63)) {
                                                                    case 347:
                                                                        v46 -= uint32(v74)
                                                                    case -2147483648:
                                                                        fmt.Println(v3, fi51, v54)
                                                                    case 32768:
                                                                        fmt.Println(v45, v3, v46)
                                                                    default:
                                                                        fmt.Println(fi51)
                                                                    }
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        switch fi51 {
                                                                        case 127:
                                                                            v50 -= uint8(v46)
                                                                        case 2147483647:
                                                                            if ((v2 | (-(v45))) != (v2 % ((v45) | 1))) {
                                                                                if (!(true || ((false && true) && false))) {
                                                                                    if (!false) {
                                                                                        v90 := GenS12{F0: v3, F1: 42, F2: -174}
                                                                                        _ = v90
                                                                                    }
                                                                                    if (((v54 <= 1e10) && true) || true) {
                                                                                        switch -2147483648 {
                                                                                        case 817:
                                                                                            switch v3 {
                                                                                            case "foo bar":
                                                                                                v4["hello"] = 157
                                                                                            }
                                                                                        case 42:
                                                                                            var v91 float64 = (v54 - ((((1e10 - (0.1 / v54)) / ((v54 - v54) + (v54 - v54))) / 3.14159) - 0.0))
                                                                                            _ = v91
                                                                                        default:
                                                                                            v4[v3] = (^(v46))
                                                                                        }
                                                                                        var v92 uint8 = 2
                                                                                        _ = v92
                                                                                    }
                                                                                    if false {
                                                                                        v93 := GenS13{F0: false}
                                                                                        if (501 != 42) {
                                                                                            switch (fi53 % ((fi53) | 1)) {
                                                                                            case -9223372036854775808:
                                                                                                v94 := GenS13{F0: ((v68 << (((v74 * (^(v74)))) & 63)) > 65535)}
                                                                                                _ = v94
                                                                                            case 67:
                                                                                                var v95 uint8 = (v52 & (v50 >> ((4294967295) & 63)))
                                                                                                _ = v95
                                                                                            case 32768:
                                                                                                v96 := GenS14{F0: v68, F1: 2147483648}
                                                                                                _ = v96
                                                                                            default:
                                                                                                v2 = (v2 & v2)
                                                                                            }
                                                                                            var v97 uint8 = (v50 ^ (v88 & 255))
                                                                                            var v98 string = ("hello" + v3)
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        } else {
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                var v99 bool = (202 != 1)
                                                                                                _ = v99
                                                                                            case "hello":
                                                                                                v100 := GenS15{F0: 18446744073709551615, F1: true}
                                                                                                _ = v100
                                                                                            case "":
                                                                                                var v101 uint32 = 1
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    fmt.Println(v46, fi53)
                                                                                                default:
                                                                                                    _ = len(v1)
                                                                                                }
                                                                                                if (42 >= v45) {
                                                                                                    _ = len(v1)
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = v46
                                                                                                    }
                                                                                                    switch (fi51 | fi53) {
                                                                                                    case 0:
                                                                                                        if (0 >= -966) {
                                                                                                            var v102 uint64 = 2
                                                                                                            var v103 string = ((v3 + v3) + (v3 + ("a" + ("hello" + ("\n\t" + v3)))))
                                                                                                            var v104 uint8 = v52
                                                                                                            v105 := GenS16{F0: (v104 >> ((669) & 63)), F1: v104}
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                            _ = v102
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                            _ = v105
                                                                                                        }
                                                                                                    case -901:
                                                                                                        var v106 string = v3
                                                                                                        _ = v106
                                                                                                    case 9223372036854775807:
                                                                                                        v107 := GenS13{F0: false}
                                                                                                        _ = v107
                                                                                                    default:
                                                                                                        fmt.Println(v50, v46)
                                                                                                    }
                                                                                                }
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                        _ = v93
                                                                                    }
                                                                                } else {
                                                                                    v54 += (v54 / v54)
                                                                                    var v108 uint64 = 908
                                                                                    _ = v108
                                                                                }
                                                                                v74 = (v74 * v74)
                                                                                if (false || (42 <= (v68 % (((v68 ^ 2)) | 1)))) {
                                                                                    switch v3 {
                                                                                    case "a":
                                                                                        fmt.Println(v50, v2)
                                                                                        var v109 int16 = -32768
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v109
                                                                                    case "foo bar":
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            if (152 <= (v46 + v46)) {
                                                                                                delete(v4, "hello")
                                                                                                v88 -= (v88 % (((v50 * genFn1(int16(v45), -51, fi53))) | 1))
                                                                                                var v110 int32 = 1
                                                                                                switch 42 {
                                                                                                case -359:
                                                                                                    if false {
                                                                                                        var v111 uint32 = uint32(v68)
                                                                                                        _ = v111
                                                                                                    }
                                                                                                case 4294967295:
                                                                                                    v112 := GenS17{F0: 42}
                                                                                                    _ = v112
                                                                                                case 676:
                                                                                                    if true {
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        var v113 float64 = (0.1 - ((v54 + ((float64(1e10) - v54) * float64((1.5 / 1.5)))) * (v54 / (float64(0.1) + 1e10))))
                                                                                                        fmt.Println(v46, v74, v88)
                                                                                                        var v114 int8 = v2
                                                                                                        if true {
                                                                                                            var v115 float32 = ((3.14159 + (1.5 + float32(v54))) / float32((v54 * -2.25)))
                                                                                                            switch v3 {
                                                                                                            case "foo bar":
                                                                                                                if true {
                                                                                                                    fmt.Println(v2)
                                                                                                                    fmt.Println(v3)
                                                                                                                    v116 := GenS2{F0: 101, F1: (v74 % ((v74) | 1))}
                                                                                                                    _ = v116
                                                                                                                } else {
                                                                                                                    var v117 float64 = (0.0 / v113)
                                                                                                                    _ = len(v1)
                                                                                                                    var v118 int32 = 0
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        _ = v113
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        if ((-(v110)) != (v110 & v110)) {
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = v46
                                                                                                            }
                                                                                                            v52 -= v88
                                                                                                        } else {
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = v46
                                                                                                            }
                                                                                                            var v119 string = v3
                                                                                                            var v120 uint = (v74 | 1)
                                                                                                            v121 := GenS7{F0: 32768, F1: 0}
                                                                                                            _ = v119
                                                                                                            _ = v120
                                                                                                            _ = v121
                                                                                                        }
                                                                                                        switch (-(fi53)) {
                                                                                                        case -790:
                                                                                                            var v122 int64 = -9223372036854775808
                                                                                                            _ = v122
                                                                                                        default:
                                                                                                            if false {
                                                                                                                fmt.Println(v74)
                                                                                                                fmt.Println(v52, v74, v110)
                                                                                                            }
                                                                                                        }
                                                                                                        switch int(v2) {
                                                                                                        case 32768:
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                if ((v46 ^ (v46 >> (((v74 + 1)) & 63))) != (^(v46))) {
                                                                                                                    var v123 int32 = v110
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                            case "":
                                                                                                                if (0 <= 0) {
                                                                                                                    var v124 uint8 = (v88 ^ (v50 >> (((v74 % (((v74 & (v74 | (^(v74))))) | 1))) & 63)))
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                v68 = (v68 * v68)
                                                                                                            default:
                                                                                                                v46 -= 0
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v46 += (v46 << (((v74 ^ (^(v74)))) & 63))
                                                                                                }
                                                                                                v46 -= 1
                                                                                                _ = v110
                                                                                            }
                                                                                            var v125 float32 = ((3.14159 - float32(v54)) / -2.25)
                                                                                            _ = v125
                                                                                        }
                                                                                    default:
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                    var v126 int32 = 2147483647
                                                                                    v54 -= ((1.5 / ((1.5 + v54) + (3.14159 + v54))) + (float64((3.14159 / 0.1)) * ((v54 * v54) - float64(0.0))))
                                                                                    _ = v126
                                                                                }
                                                                                if true {
                                                                                    v127 := GenS13{F0: ((v74 % ((v74) | 1)) != (v74 ^ v74))}
                                                                                    v68 += (^(v68))
                                                                                    var v128 float64 = 3.14159
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                            } else {
                                                                                v129 := GenS18{F0: true, F1: -1, F2: (v88 / (((v88 << ((v74) & 63))) | 1))}
                                                                                var v130 uint8 = v52
                                                                                v131 := GenS19{F0: false}
                                                                                if (true && false) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (^(v46))
                                                                                    }
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v46
                                                                                    }
                                                                                }
                                                                                _ = v129
                                                                                _ = v130
                                                                                _ = v131
                                                                            }
                                                                        default:
                                                                            fmt.Println(v50, v46, v74)
                                                                        }
                                                                    case "a":
                                                                        if (false || false) {
                                                                            var v132 float32 = (0.1 / float32(v54))
                                                                            v50 -= 255
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 1
                                                                            }
                                                                            _ = v132
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    v133 := GenS4{F0: (127 == uint8(v68)), F1: v3, F2: (v50 >> ((255) & 63))}
                                                                                    fmt.Println(v68, v52)
                                                                                    v134 := GenS11{F0: (v46 % ((821) | 1)), F1: -7, F2: (fi53 & (^(fi53)))}
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                } else {
                                                                                    var v135 uint16 = uint16(v74)
                                                                                    _ = v135
                                                                                }
                                                                            case "hello":
                                                                                var v136 int64 = 379
                                                                                _ = v136
                                                                            }
                                                                            fmt.Println(fi53, v50)
                                                                            v45 -= v2
                                                                            if true {
                                                                                var v137 int16 = 0
                                                                                _ = v137
                                                                            }
                                                                            v1 = append(v1, v46)
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            if (v3 <= v3) {
                                                                                v138 := GenS11{F0: 168, F1: -32768, F2: 382}
                                                                                _ = v138
                                                                            }
                                                                            var v139 string = v3
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v3 = ((((((v139 + "\n\t") + (((((((((v3 + v3) + ("\n\t" + (v139 + (v3 + "foo bar")))) + v139) + (("foo bar" + v139) + (v139 + (("hello" + "") + (v3 + "hello"))))) + "a") + v3) + v3) + ("hello" + ("" + "a"))) + v139)) + "a") + v139) + v139) + ((v3 + "\n\t") + v139))
                                                                            case "a":
                                                                                if (!true) {
                                                                                    var v140 uint = uint(v46)
                                                                                    var v141 int8 = 24
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                }
                                                                            case "\n\t":
                                                                                if (42 == -2147483648) {
                                                                                    v88 = genFn1(-999, v2, (^(fi53)))
                                                                                } else {
                                                                                    var v142 int64 = 868
                                                                                    v54 += 1.5
                                                                                    var v143 int8 = (^(v2))
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                }
                                                                            }
                                                                            _ = v139
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            if (false || false) {
                                                                                v2 = v45
                                                                                switch "" {
                                                                                case "":
                                                                                    v144 := GenS20{F0: 854, F1: 0, F2: -32768}
                                                                                    _ = v144
                                                                                case "a":
                                                                                    var v145 uint32 = v46
                                                                                    _ = v145
                                                                                case "\n\t":
                                                                                    switch v3 {
                                                                                    case "":
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                default:
                                                                                    v146 := GenS21{F0: 136, F1: 0, F2: v46}
                                                                                    _ = v146
                                                                                }
                                                                            } else {
                                                                                if (int32(v74) == 2147483647) {
                                                                                    v147 := GenS14{F0: (v68 & v68), F1: v46}
                                                                                    v45 = v2
                                                                                    _ = v147
                                                                                } else {
                                                                                    fmt.Println(v68, fi51)
                                                                                }
                                                                            }
                                                                            v68 += 70
                                                                            fmt.Println(v50)
                                                                        }
                                                                    default:
                                                                        var v148 bool = ((v46 == 65535) || false)
                                                                        _ = v148
                                                                    }
                                                                    var v149 int = 2147483647
                                                                    _ = v149
                                                                }
                                                                var v150 int16 = 505
                                                                v151 := GenS22{F0: (("\n\t" + (v3 + "hello")) + v3), F1: v3}
                                                                switch v3 {
                                                                case "hello":
                                                                    var v152 uint16 = v68
                                                                    _ = v152
                                                                case "":
                                                                    var v153 int32 = -1
                                                                    _ = v153
                                                                case "a":
                                                                    switch (v3 + "\n\t") {
                                                                    case "hello":
                                                                        fmt.Println(v74)
                                                                    case "":
                                                                        v154 := GenS17{F0: 11}
                                                                        _ = v154
                                                                    case "foo bar":
                                                                        switch fi53 {
                                                                        case 42:
                                                                            var v155 float32 = 1e10
                                                                            _ = v155
                                                                        case 128:
                                                                            fmt.Println(v54, v45, fi51)
                                                                        case 1:
                                                                            fmt.Println(v88, v68)
                                                                        default:
                                                                            fmt.Println(v3)
                                                                            v156 := GenS23{F0: true, F1: ((v3 >= (v3 + v3)) || true), F2: "\n\t"}
                                                                            fmt.Println(fi53, v150)
                                                                            _ = v156
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            delete(v4, ((v3 + (((v3 + (v3 + (v3 + (v3 + "\n\t")))) + (v3 + v3)) + "foo bar")) + v3))
                                                                            if true {
                                                                                switch 42 {
                                                                                case -2147483648:
                                                                                    fmt.Println(v45)
                                                                                case 0:
                                                                                    v3 = v3
                                                                                case 32768:
                                                                                    var v157 uint32 = 409
                                                                                    _ = v157
                                                                                default:
                                                                                    var v158 bool = (-30 == (fi53 * fi53))
                                                                                    _ = v158
                                                                                }
                                                                            }
                                                                            if (142 > 0) {
                                                                                v150 -= v150
                                                                            } else {
                                                                                if true {
                                                                                    v150 = v150
                                                                                } else {
                                                                                    var v159 int = (fi51 % ((-883) | 1))
                                                                                    _ = v159
                                                                                }
                                                                                v50 = genFn1((v150 ^ 42), v45, (^(fi53)))
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            var v160 uint8 = v52
                                                                            _ = v160
                                                                        }
                                                                    }
                                                                }
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                            v74 = uint(v68)
                                                        } else {
                                                            v161 := GenS14{F0: 32767, F1: (^(v46))}
                                                            v46 = 306
                                                            _ = v161
                                                        }
                                                        _ = v88
                                                    }
                                                }
                                                v45 -= (v45 - 0)
                                                _ = v74
                                            }
                                            v50 -= v50
                                            if len(v1) > 0 {
                                                v1[0] = (^(v46))
                                            }
                                            _ = v68
                                        }
                                    }
                                    var v162 uint8 = 2
                                    var v163 float32 = 3.14159
                                    v164 := GenS2{F0: -1, F1: uint(v50)}
                                    _ = v162
                                    _ = v163
                                    _ = v164
                                } else {
                                    var v165 bool = ((!(-1 == 42)) || false)
                                    _ = len(v4)
                                    var v166 bool = (131 <= 42)
                                    v46 += (v46 ^ v46)
                                    fmt.Println(v45, v45)
                                    _ = v165
                                    _ = v166
                                }
                            default:
                                var v167 int64 = 9223372036854775807
                                _ = v167
                            }
                        }
                        var v168 int16 = -498
                        var v169 bool = (!(true || false))
                        _ = fi53
                        _ = v168
                        _ = v169
                        _ = v54
                    }
                    _ = fi51
                    _ = v52
                }
                if ((-453 != 32768) || true) {
                    var v170 uint16 = 2
                    _ = v170
                }
                _ = v50
            } else {
                var v171 uint16 = 255
                _ = v171
            }
            v45 = v45
        } else {
            switch v3 {
            case "a":
                if (1 != 0) {
                    if false {
                        switch 1 {
                        case -824:
                            fmt.Println(v3, v2, v3)
                        case 0:
                            var v172 float32 = (0.1 / float32((1e10 - 3.14159)))
                            _ = v172
                        default:
                            if (!false) {
                                fmt.Println(v46, v3, v45)
                                var v173 string = v3
                                if true {
                                    v45 += v45
                                    v174 := GenS12{F0: v173, F1: 0, F2: int16(v46)}
                                    _ = v174
                                } else {
                                    var v175 bool = false
                                    for fi176 := 0; fi176 < 4; fi176++ {
                                        var v177 int16 = 1000
                                        for fi178 := 0; fi178 < 7; fi178++ {
                                            v173 = v3
                                            v179 := v4["foo bar"]
                                            var v180 string = ""
                                            _ = fi178
                                            _ = v179
                                            _ = v180
                                        }
                                        switch fi176 {
                                        case 4294967295:
                                            for fi181 := 0; fi181 < 6; fi181++ {
                                                fmt.Println(v46)
                                                fmt.Println(v46)
                                                if false {
                                                    var v182 float64 = (3.14159 * 0.1)
                                                    _ = v182
                                                } else {
                                                    fmt.Println(v45, v173, v3)
                                                    var v183 int = (^(fi181))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if (!v175) {
                                                        var v184 int16 = v177
                                                        if v175 {
                                                            v184 = (v184 >> ((514) & 63))
                                                            if v175 {
                                                                v185 := GenS24{F0: 9223372036854775807, F1: (v46 * 1), F2: -1}
                                                                v186 := v4["a"]
                                                                _ = v185
                                                                _ = v186
                                                            } else {
                                                                if true {
                                                                    delete(v4, v173)
                                                                    var v187 int32 = -1
                                                                    v4[(v173 + "\n\t")] = 1
                                                                    _ = v187
                                                                }
                                                                var v188 int = fi176
                                                                v175 = (v175 && v175)
                                                                _ = v188
                                                            }
                                                            var v189 int8 = (v2 / (((v2 / ((v45) | 1))) | 1))
                                                            var v190 int = fi176
                                                            _ = v189
                                                            _ = v190
                                                        }
                                                        v3 = v3
                                                        _ = v184
                                                    } else {
                                                        v3 = "foo bar"
                                                        fmt.Println(v175)
                                                    }
                                                    fmt.Println(v46, v183, v45)
                                                    _ = v183
                                                }
                                                _ = fi181
                                            }
                                        case 73:
                                            delete(v4, "")
                                        case 127:
                                            var v191 int32 = -1
                                            _ = v191
                                        }
                                        _ = fi176
                                        _ = v177
                                    }
                                    _ = v175
                                }
                                _ = v173
                            } else {
                                v3 = v3
                            }
                        }
                    } else {
                        fmt.Println(v45)
                    }
                    fmt.Println(v45, v46)
                    var v192 int = 0
                    _ = v192
                }
            case "foo bar":
                var v193 uint16 = 916
                _ = v193
            default:
                v194 := GenS3{F0: 1}
                _ = v194
            }
            switch 351 {
            case 127:
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        switch 0 {
                        case 9223372036854775807:
                            v3 = "\n\t"
                        case 32767:
                            fmt.Println(v3)
                        case 127:
                            var v195 uint = 879
                            _ = v195
                        default:
                            for fi196 := 0; fi196 < 4; fi196++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v197 := GenS25{F0: fi196, F1: 65535}
                                    _ = v197
                                case "":
                                    for fi198 := 0; fi198 < 11; fi198++ {
                                        switch fi198 {
                                        case -793:
                                            var v199 int8 = v2
                                            v200 := GenS4{F0: true, F1: v3, F2: 209}
                                            if true {
                                                v199 += (v199 & v199)
                                            } else {
                                                v45 = v45
                                                v2 += 57
                                                v201 := GenS22{F0: v3, F1: v3}
                                                v202 := GenS17{F0: -125}
                                                _ = v201
                                                _ = v202
                                            }
                                            _ = v199
                                            _ = v200
                                        case 4294967295:
                                            v203 := GenS26{F0: (false || (int8(v46) != 1)), F1: 42}
                                            _ = v203
                                        case 2147483647:
                                            if true {
                                                var v204 int8 = (v2 ^ 0)
                                                switch 128 {
                                                case 2147483647:
                                                    var v205 int64 = 1
                                                    _ = v205
                                                case 514:
                                                    _ = len(v1)
                                                case -133:
                                                    var v206 uint32 = (v46 * (v46 << ((763) & 63)))
                                                    _ = v206
                                                default:
                                                    switch v3 {
                                                    case "a":
                                                        var v207 int8 = 127
                                                        _ = v207
                                                    case "foo bar":
                                                        v208 := GenS21{F0: 0, F1: 728, F2: (v46 << ((150) & 63))}
                                                        _ = v208
                                                    case "\n\t":
                                                        var v209 int16 = -1
                                                        _ = v209
                                                    }
                                                }
                                                _ = v204
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = 1
                                        }
                                        v45 += (v2 | v2)
                                        v2 -= -46
                                        v46 = 0
                                        _ = fi198
                                    }
                                }
                                delete(v4, v3)
                                v3 = v3
                                _ = fi196
                            }
                        }
                    case "foo bar":
                        var v210 int32 = 2147483647
                        _ = v210
                    case "a":
                        v45 = (v45 & 42)
                    }
                }
            }
        }
        v45 = v45
    } else {
        v46 = (v46 + v46)
        if (true || true) {
            v46 = 0
            var v211 int32 = -2147483648
            _ = v211
        }
    }
    for fi212 := 0; fi212 < 9; fi212++ {
        v1 = append(v1, (v46 ^ v46))
        v3 = "hello"
        _ = fi212
    }
    fmt.Println(v2, v45, v46)
    var v213 float64 = -2.25
    if true {
        for fi214 := 0; fi214 < 12; fi214++ {
            if ((v46 - (v46 & v46)) <= v46) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v215 int16 = 821
                _ = v215
            }
            _ = len(v1)
            v3 = v3
            _ = fi214
        }
        var v216 string = v3
        _ = v216
    } else {
        fmt.Println(v213, v45)
    }
    var v217 uint16 = 42
    v2 = v2
    var v218 uint = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v213:", v213)
    fmt.Println("v217:", v217)
    fmt.Println("v218:", v218)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v213
    _ = v217
    _ = v218
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
