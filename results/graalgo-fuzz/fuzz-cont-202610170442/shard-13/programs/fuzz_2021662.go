package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint16; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int; F1 uint64; F2 uint }
type GenS15 struct { F0 uint32; F1 bool; F2 uint }
type GenS16 struct { F0 int16; F1 bool; F2 int }
type GenS17 struct { F0 uint; F1 uint; F2 bool }
type GenS18 struct { F0 uint64; F1 uint8; F2 int }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int32; F1 int8; F2 int8 }
type GenS21 struct { F0 uint64; F1 uint8; F2 uint }
type GenS22 struct { F0 int32; F1 int32 }
type GenS23 struct { F0 int16; F1 int; F2 int8 }
type GenS24 struct { F0 string; F1 uint; F2 uint64 }
type GenS25 struct { F0 string; F1 int8; F2 uint16 }
type GenS26 struct { F0 uint; F1 int16 }
type GenS27 struct { F0 uint32 }
type GenS28 struct { F0 int8; F1 bool; F2 int32 }
type GenS29 struct { F0 int; F1 uint32 }
type GenS30 struct { F0 uint8; F1 int64; F2 int }
func genFn0() int16 {
    v4 := []int{}
    var v5 int16 = -450
    if false {
        v6 := GenS1{F0: (true || true), F1: 2147483647, F2: v5}
        fmt.Println(v5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v6
    }
    var v7 float64 = 3.14159
    _ = v4
    _ = v5
    _ = v7
    return (-(v5))
}
func genFn1(p0 uint16) uint64 {
    p0 += 0
    v8 := []int{0, 7}
    p0 -= 42
    _ = p0
    _ = v8
    return 63
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -332
    v3 := []string{}
    v2 += (^(v2))
    v2 = genFn0()
    var v9 int32 = -1
    for fi10 := 0; fi10 < 4; fi10++ {
        _ = len(v3)
        fmt.Println(v9, v2, v9)
        var v11 int32 = (v9 << ((767) & 63))
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: v9, F1: v9, F2: ((true && false) && false)}
    var v13 uint32 = 1
    v2 += v2
    fmt.Println(v13, v9, v9)
    if false {
        v2 -= (^(v2))
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint64 = 0
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        var v17 float32 = 1.5
        _ = v17
    }
    var v18 int32 = -1
    v9 += 42
    switch 42 {
    case 32768:
        var v19 bool = (true && ((756 < 42) || true))
        _ = v19
    default:
        v20 := GenS1{F0: true, F1: 42, F2: -60}
        _ = v20
    }
    var v21 string = ("\n\t" + "")
    delete(v1, (v13 ^ v13))
    if true {
        if false {
            v13 -= v13
        }
        v22 := GenS2{F0: (v9 >> ((18446744073709551615) & 63)), F1: (v18 ^ (v9 | (v18 | 2147483647))), F2: (true || false)}
        if ((255 > 0) && ((666 <= 627) && (!false))) {
            fmt.Println(v18, v13, v9)
            if false {
                v23 := v1[v13]
                v9 = v18
                for fi24 := 0; fi24 < 20; fi24++ {
                    if true {
                        if (v9 <= int32(v13)) {
                            _ = len(v3)
                            for fi25 := 0; fi25 < 20; fi25++ {
                                v26 := GenS3{F0: (("foo bar" + (v21 + v21)) + "hello"), F1: 42, F2: 42}
                                v1[(v13 ^ (v23 * v23))] = v13
                                v21 = "hello"
                                var v27 int8 = 100
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            v9 += v18
                        } else {
                            for fi28 := 0; fi28 < 16; fi28++ {
                                v29 := GenS3{F0: v21, F1: 0, F2: 916}
                                var v30 int64 = -9223372036854775808
                                v30 -= (v30 - v30)
                                var v31 uint = 980
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 16; fi32++ {
                                var v33 int = (fi32 + 2147483647)
                                if false {
                                    var v34 uint64 = 56
                                    fmt.Println(v33)
                                    _ = v34
                                } else {
                                    fmt.Println(v9)
                                    fmt.Println(v21)
                                    v35 := GenS3{F0: v21, F1: 2, F2: 65535}
                                    _ = v35
                                }
                                v23 = v23
                                _ = fi32
                                _ = v33
                            }
                        }
                        if false {
                            var v36 int64 = 198
                            fmt.Println(v13, v2, v36)
                            var v37 int16 = (-(v2))
                            _ = v36
                            _ = v37
                        } else {
                            var v38 int64 = 1
                            for fi39 := 0; fi39 < 9; fi39++ {
                                var v40 uint = 65535
                                if (true || false) {
                                    switch (v21 + (v21 + ("a" + (v21 + "hello")))) {
                                    case "foo bar":
                                        _ = len(v3)
                                    default:
                                        v40 = v40
                                    }
                                } else {
                                    switch fi24 {
                                    case -318:
                                        var v41 string = v21
                                        _ = v41
                                    case 32768:
                                        if (int(v23) <= (-(fi24))) {
                                            v21 = v21
                                            switch v21 {
                                            case "hello":
                                                v18 -= v9
                                            case "foo bar":
                                                var v42 uint64 = 0
                                                var v43 uint32 = v23
                                                v44 := GenS2{F0: v18, F1: -1, F2: false}
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            default:
                                                var v45 float32 = -2.25
                                                _ = v45
                                            }
                                        }
                                    case 32767:
                                        if (true || true) {
                                            v40 = v40
                                            if false {
                                                v1[v23] = 65535
                                                var v46 float32 = 3.14159
                                                v47 := GenS2{F0: v9, F1: v9, F2: (true || (("hello" + "a") != (((("" + (v21 + "hello")) + v21) + ("a" + v21)) + (v21 + ((v21 + v21) + "a")))))}
                                                _ = v46
                                                _ = v47
                                            } else {
                                                if false {
                                                    if (22 == (v40 ^ (v40 & 4294967295))) {
                                                        if true {
                                                            v23 = (^(v23))
                                                            v3 = append(v3, v21)
                                                            var v48 int32 = 0
                                                            v49 := GenS4{F0: false, F1: v38, F2: 128}
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            v18 = (v9 & v9)
                                                            v50 := v1[v13]
                                                            v51 := v1[v13]
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 uint64 = 42
                                                        fmt.Println(v40)
                                                        delete(v1, 0)
                                                        _ = v52
                                                    }
                                                    var v53 int = 42
                                                    switch ("foo bar" + ((v21 + "\n\t") + v21)) {
                                                    case "a":
                                                        switch fi24 {
                                                        case -9223372036854775808:
                                                            v13 -= 550
                                                        case 32768:
                                                            var v54 int8 = 100
                                                            _ = v54
                                                        case -146:
                                                            v40 = (v40 + v40)
                                                            switch (^(fi24)) {
                                                            case 4294967295:
                                                                v23 -= (v13 * (v13 ^ (v23 ^ v13)))
                                                            case 855:
                                                                v55 := GenS5{F0: (!(!(!(!true)))), F1: uint64(v40), F2: 418}
                                                                _ = v55
                                                            default:
                                                                v23 -= 65535
                                                            }
                                                            var v56 uint64 = 106
                                                            _ = v56
                                                        }
                                                        v57 := GenS3{F0: ("hello" + ""), F1: 65535, F2: 65535}
                                                        if (true && (v21 >= v21)) {
                                                            v38 = int64(v9)
                                                            v3 = append(v3, ((v21 + "") + v21))
                                                            var v58 float32 = ((1e10 - 3.14159) - 1e10)
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    }
                                                    var v59 uint = 255
                                                    v18 -= v9
                                                    var v60 string = "\n\t"
                                                    v2 -= (v2 / (((-(v2))) | 1))
                                                    if (!(0.0 < 3.14159)) {
                                                        var v61 int8 = 28
                                                        if false {
                                                            if true {
                                                                switch v60 {
                                                                case "a":
                                                                    v60 = v21
                                                                case "":
                                                                    v62 := GenS3{F0: ("hello" + "a"), F1: 893, F2: 0}
                                                                    _ = v62
                                                                case "\n\t":
                                                                    switch (v53 >> ((v40) & 63)) {
                                                                    case -1:
                                                                        v18 -= v9
                                                                    }
                                                                    v21 = "\n\t"
                                                                }
                                                                switch (((v60 + "hello") + v60) + (v60 + (v21 + ("\n\t" + v21)))) {
                                                                case "foo bar":
                                                                    if (true || (!false)) {
                                                                        v1[(v13 + 4294967295)] = v23
                                                                        switch v60 {
                                                                        case "\n\t":
                                                                            switch (v21 + "a") {
                                                                            case "\n\t":
                                                                                v60 = v60
                                                                            }
                                                                        case "foo bar":
                                                                            var v63 string = "\n\t"
                                                                            _ = v63
                                                                        default:
                                                                            v64 := GenS5{F0: ((true || ((!(!true)) && true)) || (!false)), F1: 1, F2: fi24}
                                                                            _ = v64
                                                                        }
                                                                        if ((v2 | 0) == int16(v40)) {
                                                                            if false {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            v60 = v60
                                                                        } else {
                                                                            var v65 int16 = 42
                                                                            _ = v65
                                                                        }
                                                                    } else {
                                                                        v66 := GenS6{F0: v59}
                                                                        _ = v66
                                                                    }
                                                                case "a":
                                                                    v67 := GenS3{F0: "\n\t", F1: 255, F2: 1}
                                                                    _ = v67
                                                                case "hello":
                                                                    if true {
                                                                        v21 = (((v60 + v60) + "hello") + (v60 + v60))
                                                                        if true {
                                                                            v18 = v9
                                                                            if false {
                                                                                _ = len(v1)
                                                                                if (true || (1 != 1)) {
                                                                                    switch v21 {
                                                                                    case "hello":
                                                                                        var v68 uint8 = 255
                                                                                        _ = v68
                                                                                    default:
                                                                                        var v69 uint8 = 2
                                                                                        _ = v69
                                                                                    }
                                                                                }
                                                                                switch v21 {
                                                                                case "":
                                                                                    var v70 uint16 = 2
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                            if ((v61 % ((v61) | 1)) >= int8(v38)) {
                                                                                if (3.14159 == 3.14159) {
                                                                                    _ = len(v1)
                                                                                    v71 := GenS2{F0: v18, F1: v18, F2: false}
                                                                                    _ = v71
                                                                                }
                                                                                if true {
                                                                                    switch (v60 + ("" + v21)) {
                                                                                    case "":
                                                                                        v72 := GenS3{F0: (v21 + v21), F1: 2, F2: 1}
                                                                                        _ = v72
                                                                                    }
                                                                                }
                                                                                if (v9 <= v18) {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        switch (v53 >> (((v40 & v40)) & 63)) {
                                                                                        case 32768:
                                                                                            var v73 uint8 = uint8(v59)
                                                                                            _ = v73
                                                                                        case 9223372036854775807:
                                                                                            v2 = v2
                                                                                        case 0:
                                                                                            v74 := GenS6{F0: v40}
                                                                                            _ = v74
                                                                                        default:
                                                                                            v75 := GenS7{F0: (v2 << ((v59) & 63)), F1: 806, F2: v23}
                                                                                            _ = v75
                                                                                        }
                                                                                    case "":
                                                                                        var v76 uint32 = 4294967295
                                                                                        var v77 string = v60
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    default:
                                                                                        v78 := GenS5{F0: false, F1: 902, F2: -2147483648}
                                                                                        _ = v78
                                                                                    }
                                                                                    fmt.Println(v9, v13, fi24)
                                                                                    v21 = v60
                                                                                    if true {
                                                                                        switch (fi39 / (((fi24 | fi39)) | 1)) {
                                                                                        case 1:
                                                                                            v18 = 2147483647
                                                                                        case -1:
                                                                                            var v79 int64 = v38
                                                                                            _ = v79
                                                                                        default:
                                                                                            switch ((v60 + v21) + (v21 + "foo bar")) {
                                                                                            case "foo bar":
                                                                                                switch 4294967295 {
                                                                                                case 42:
                                                                                                    fmt.Println(v40)
                                                                                                case 127:
                                                                                                    v18 = (v9 * -1)
                                                                                                case 0:
                                                                                                    if true {
                                                                                                        v38 = -1
                                                                                                        switch (fi24 >> ((308) & 63)) {
                                                                                                        case 32768:
                                                                                                            v80 := GenS4{F0: (((v59 >> (((v59 % ((v59) | 1))) & 63)) > v40) || false), F1: v38, F2: 213}
                                                                                                            _ = v80
                                                                                                        case 128:
                                                                                                            var v81 int8 = v61
                                                                                                            _ = v81
                                                                                                        case -459:
                                                                                                            var v82 string = (("" + (v60 + v60)) + v21)
                                                                                                            _ = v82
                                                                                                        default:
                                                                                                            switch (v60 + ((("\n\t" + (v60 + "hello")) + v60) + "foo bar")) {
                                                                                                            case "\n\t":
                                                                                                                v1[v23] = (^(v23))
                                                                                                            case "hello":
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        switch v60 {
                                                                                                                        case "hello":
                                                                                                                            v83 := GenS7{F0: v2, F1: 18446744073709551615, F2: v13}
                                                                                                                            _ = v83
                                                                                                                        case "foo bar":
                                                                                                                            v84 := GenS8{F0: 42, F1: 49}
                                                                                                                            _ = v84
                                                                                                                        case "a":
                                                                                                                            v85 := GenS9{F0: (v9 + v9), F1: 32767}
                                                                                                                            _ = v85
                                                                                                                        }
                                                                                                                        v61 -= v61
                                                                                                                        var v86 int64 = (-(v38))
                                                                                                                        _ = v86
                                                                                                                    }
                                                                                                                    var v87 int8 = v61
                                                                                                                    switch "hello" {
                                                                                                                    case "\n\t":
                                                                                                                        v88 := GenS9{F0: 0, F1: 42}
                                                                                                                        _ = v88
                                                                                                                    case "a":
                                                                                                                        delete(v1, (^(v23)))
                                                                                                                    default:
                                                                                                                        fmt.Println(v60, v9)
                                                                                                                        switch fi24 {
                                                                                                                        case 128:
                                                                                                                            v40 += 18446744073709551615
                                                                                                                        default:
                                                                                                                            v89 := v1[v23]
                                                                                                                            _ = v89
                                                                                                                        }
                                                                                                                        var v90 int = 42
                                                                                                                        _ = v90
                                                                                                                    }
                                                                                                                    _ = v87
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v91 bool = false
                                                                                                                _ = v91
                                                                                                            default:
                                                                                                                v38 += v38
                                                                                                            }
                                                                                                        }
                                                                                                        var v92 uint32 = (v13 + 0)
                                                                                                        _ = v92
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v23 = (v23 << ((v40) & 63))
                                                                                            }
                                                                                        }
                                                                                        v93 := GenS9{F0: v9, F1: 32767}
                                                                                        _ = v93
                                                                                    } else {
                                                                                        var v94 float64 = 3.14159
                                                                                        _ = v94
                                                                                    }
                                                                                } else {
                                                                                    switch -1 {
                                                                                    case 568:
                                                                                        switch (v53 + (fi39 / ((fi39) | 1))) {
                                                                                        case -328:
                                                                                            switch (v53 / ((fi39) | 1)) {
                                                                                            case 25:
                                                                                                fmt.Println(v38)
                                                                                            case 9223372036854775807:
                                                                                                v95 := GenS5{F0: ((^(v59)) < (^(v40))), F1: 42, F2: fi24}
                                                                                                _ = v95
                                                                                            case 907:
                                                                                                if ((22 >= 27) || (0 != 42)) {
                                                                                                    var v96 int = fi39
                                                                                                    v96 += (v96 + (^(fi39)))
                                                                                                    var v97 float32 = -2.25
                                                                                                    _ = v96
                                                                                                    _ = v97
                                                                                                } else {
                                                                                                    var v98 uint64 = 4294967295
                                                                                                    _ = v98
                                                                                                }
                                                                                            }
                                                                                        case 2147483647:
                                                                                            var v99 uint16 = 2
                                                                                            _ = v99
                                                                                        case 32767:
                                                                                            var v100 uint32 = 65535
                                                                                            _ = v100
                                                                                        default:
                                                                                            var v101 bool = (true && true)
                                                                                            _ = v101
                                                                                        }
                                                                                    }
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    var v102 int16 = (v2 / (((-(v2))) | 1))
                                                                                    _ = v102
                                                                                }
                                                                            } else {
                                                                                v103 := GenS2{F0: v9, F1: (v18 >> (((v40 | (v59 / ((1) | 1)))) & 63)), F2: false}
                                                                                v61 -= -128
                                                                                var v104 string = "\n\t"
                                                                                v105 := GenS10{F0: 1}
                                                                                _ = v103
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                        } else {
                                                                            switch v60 {
                                                                            case "":
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        if (549 != 42) {
                                                                                            v3 = append(v3, v21)
                                                                                        }
                                                                                        if false {
                                                                                            var v106 uint = v40
                                                                                            if true {
                                                                                                var v107 int64 = (v38 + 0)
                                                                                                _ = v107
                                                                                            } else {
                                                                                                var v108 bool = (true && ((false && (!(!true))) && (v40 > v106)))
                                                                                                switch (v21 + v21) {
                                                                                                case "a":
                                                                                                    var v109 int16 = v2
                                                                                                    _ = v109
                                                                                                }
                                                                                                delete(v1, v23)
                                                                                                if v108 {
                                                                                                    if v108 {
                                                                                                        v2 += genFn0()
                                                                                                        v60 = (v21 + v60)
                                                                                                        if v108 {
                                                                                                            v40 = (v106 & (v106 % ((4294967295) | 1)))
                                                                                                            v53 = fi24
                                                                                                        } else {
                                                                                                            v110 := GenS1{F0: (129 < 127), F1: -2147483648, F2: (v2 * (v2 * -1))}
                                                                                                            fmt.Println(v38, v108, v18)
                                                                                                            _ = v110
                                                                                                        }
                                                                                                    }
                                                                                                    var v111 int64 = (v38 & v38)
                                                                                                    var v112 int = v53
                                                                                                    _ = v111
                                                                                                    _ = v112
                                                                                                } else {
                                                                                                    var v113 int64 = int64(v59)
                                                                                                    if (int64(v2) != v113) {
                                                                                                        v18 += v9
                                                                                                        v18 = (v18 % (((v9 + (v18 >> ((v59) & 63)))) | 1))
                                                                                                    } else {
                                                                                                        switch (v21 + v60) {
                                                                                                        case "a":
                                                                                                            v1[v13] = v13
                                                                                                        case "":
                                                                                                            v114 := GenS11{F0: -32768}
                                                                                                            _ = v114
                                                                                                        case "\n\t":
                                                                                                            v23 = (v23 % (((v23 | (^(v13)))) | 1))
                                                                                                        default:
                                                                                                            v115 := GenS12{F0: v21, F1: v2}
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        if v108 {
                                                                                                            var v116 uint64 = 42
                                                                                                            _ = v116
                                                                                                        }
                                                                                                        v117 := GenS1{F0: false, F1: v18, F2: genFn0()}
                                                                                                        _ = v117
                                                                                                    }
                                                                                                    _ = len(v1)
                                                                                                    _ = v113
                                                                                                }
                                                                                                _ = v108
                                                                                            }
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = "hello"
                                                                                            }
                                                                                            var v118 float64 = float64(1.5)
                                                                                            v119 := v1[(v23 - v13)]
                                                                                            _ = v106
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        }
                                                                                        v120 := GenS6{F0: (v59 ^ v40)}
                                                                                        _ = v120
                                                                                    }
                                                                                case "a":
                                                                                    if (!true) {
                                                                                        if false {
                                                                                            v121 := GenS6{F0: v59}
                                                                                            fmt.Println(v18, v60, v61)
                                                                                            fmt.Println(v38, v53, v38)
                                                                                            _ = v121
                                                                                        } else {
                                                                                            v18 -= v18
                                                                                            var v122 int8 = (v61 >> ((v40) & 63))
                                                                                            fmt.Println(v9, fi24, v21)
                                                                                            _ = v122
                                                                                        }
                                                                                        v123 := GenS13{F0: 0, F1: 710}
                                                                                        v3 = append(v3, v21)
                                                                                        _ = v123
                                                                                    } else {
                                                                                        v3 = append(v3, v60)
                                                                                        v2 = v2
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v124 uint8 = 255
                                                                                    _ = v124
                                                                                default:
                                                                                    fmt.Println(v38, v53, v18)
                                                                                }
                                                                            case "\n\t":
                                                                                _ = len(v1)
                                                                            case "foo bar":
                                                                                v2 = (v2 << (((v40 & 255)) & 63))
                                                                                v38 = 612
                                                                                var v125 uint = v40
                                                                                var v126 int16 = v2
                                                                                v126 = v2
                                                                                switch fi39 {
                                                                                case 4294967295:
                                                                                    _ = len(v1)
                                                                                case -9223372036854775808:
                                                                                    v127 := v1[(v23 ^ v23)]
                                                                                    _ = v127
                                                                                case 9223372036854775807:
                                                                                    var v128 int64 = 298
                                                                                    _ = v128
                                                                                }
                                                                                var v129 uint16 = 65535
                                                                                var v130 uint8 = 33
                                                                                var v131 float64 = (1e10 / -2.25)
                                                                                var v132 int64 = -1
                                                                                _ = v125
                                                                                _ = v126
                                                                                _ = v129
                                                                                _ = v130
                                                                                _ = v131
                                                                                _ = v132
                                                                            }
                                                                            delete(v1, uint32(v40))
                                                                            var v133 int16 = v2
                                                                            _ = v133
                                                                        }
                                                                        v134 := GenS14{F0: (v53 % (((fi39 + -139)) | 1)), F1: 4294967295, F2: 1}
                                                                        _ = v134
                                                                    }
                                                                default:
                                                                    var v135 int32 = v9
                                                                    var v136 float32 = 0.0
                                                                    _ = v135
                                                                    _ = v136
                                                                }
                                                                v1[v23] = v23
                                                            }
                                                            v137 := GenS6{F0: 2}
                                                            switch int(v59) {
                                                            case 4294967295:
                                                                if false {
                                                                    v138 := GenS15{F0: (^(v13)), F1: (!true), F2: v40}
                                                                    var v139 int8 = v61
                                                                    var v140 uint16 = 32767
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                } else {
                                                                    v141 := GenS13{F0: 1, F1: (v13 * (v13 % ((v13) | 1)))}
                                                                    v21 = (v21 + v21)
                                                                    _ = v141
                                                                }
                                                            case 0:
                                                                v40 = (v40 - (v40 / ((65535) | 1)))
                                                            case 309:
                                                                var v142 bool = false
                                                                _ = v142
                                                            }
                                                            _ = v137
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v143 := GenS8{F0: (v2 >> ((18446744073709551615) & 63)), F1: 1}
                                                            _ = v143
                                                        }
                                                        var v144 bool = false
                                                        var v145 float32 = (((0.1 / (1.5 / 3.14159)) + float32(((0.0 - 1e10) - (0.0 + 0.0)))) / 1.0)
                                                        fmt.Println(v60, v13, v144)
                                                        switch fi24 {
                                                        case 685:
                                                            var v146 int64 = 9223372036854775807
                                                            _ = v146
                                                        case 9223372036854775807:
                                                            switch v21 {
                                                            case "hello":
                                                                var v147 int32 = v18
                                                                _ = v147
                                                            case "foo bar":
                                                                fmt.Println(v9, fi24)
                                                            default:
                                                                if (v144 && (!v144)) {
                                                                    fmt.Println(v53)
                                                                }
                                                            }
                                                        case 333:
                                                            var v148 int8 = (-(v61))
                                                            _ = v148
                                                        default:
                                                            _ = len(v1)
                                                            switch v21 {
                                                            case "foo bar":
                                                                fmt.Println(v13, v144, v18)
                                                            case "a":
                                                                v149 := GenS16{F0: (v2 - v2), F1: v144, F2: (fi24 & (fi24 | fi39))}
                                                                _ = v149
                                                            case "\n\t":
                                                                v9 += (v9 + (^(v9)))
                                                            }
                                                            var v150 int = (^(fi39))
                                                            v2 -= v2
                                                            if v144 {
                                                                _ = len(v3)
                                                                if (1e10 > (-2.25 - 1.5)) {
                                                                    var v151 bool = (v144 || true)
                                                                    switch 2147483647 {
                                                                    case -9223372036854775808:
                                                                        v151 = (!true)
                                                                    case 2147483647:
                                                                        switch fi39 {
                                                                        case 2147483647:
                                                                            var v152 int64 = (v38 * v38)
                                                                            _ = v152
                                                                        case -2147483648:
                                                                            if false {
                                                                                _ = len(v1)
                                                                                fmt.Println(v9)
                                                                                fmt.Println(v145, v18)
                                                                            } else {
                                                                                var v153 float32 = v145
                                                                                v154 := GenS6{F0: (v40 | (v59 << (((v59 & v40)) & 63)))}
                                                                                v1[v13] = 1
                                                                                _ = v153
                                                                                _ = v154
                                                                            }
                                                                        default:
                                                                            if (v21 <= ("hello" + (v60 + v60))) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                fmt.Println(v53)
                                                                            } else {
                                                                                var v155 bool = (!(!(42 != 127)))
                                                                                _ = len(v1)
                                                                                _ = v155
                                                                            }
                                                                        }
                                                                    default:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    _ = v151
                                                                }
                                                            } else {
                                                                var v156 int = (-(fi24))
                                                                if v144 {
                                                                    if (2 > 4294967295) {
                                                                        fmt.Println(v40, fi39)
                                                                        v21 = ""
                                                                        v157 := GenS17{F0: v59, F1: v40, F2: (103 == 42)}
                                                                        _ = v157
                                                                    }
                                                                    if ((v13 != v13) && (v144 || (!((1e10 - (float64(3.14159) / ((float64((((float32(float64(3.14159)) / 1e10) / 1.0) * v145)) / (-2.25 / ((0.0 + float64(float32(1e10))) / 0.1))) - (0.0 * 0.1)))) > -2.25)))) {
                                                                        v145 = ((v145 / -2.25) - v145)
                                                                        _ = len(v1)
                                                                        switch 2147483647 {
                                                                        case 9223372036854775807:
                                                                            v156 = 2147483647
                                                                        case 687:
                                                                            v158 := v1[v23]
                                                                            _ = v158
                                                                        case 127:
                                                                            v150 += (v156 ^ (fi39 & 0))
                                                                        }
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v21
                                                                    }
                                                                }
                                                                _ = v156
                                                            }
                                                            v159 := GenS16{F0: (-(v2)), F1: true, F2: 804}
                                                            var v160 bool = v144
                                                            v38 = (v38 / ((int64(v9)) | 1))
                                                            _ = v150
                                                            _ = v159
                                                            _ = v160
                                                        }
                                                        var v161 uint32 = (v23 & 4294967295)
                                                        _ = v144
                                                        _ = v145
                                                        _ = v161
                                                        _ = v61
                                                    }
                                                    _ = v53
                                                    _ = v59
                                                    _ = v60
                                                }
                                                v21 = "a"
                                                v21 = "foo bar"
                                            }
                                            var v162 bool = false
                                            _ = v162
                                        }
                                    default:
                                        var v163 int8 = -1
                                        _ = v163
                                    }
                                    var v164 int32 = (v9 / (((v9 * (v9 + (v9 + (v18 + (v18 / (((v18 - (v9 ^ int32(v40)))) | 1))))))) | 1))
                                    _ = v164
                                }
                                v165 := v1[v13]
                                var v166 float32 = 1.5
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                var v167 float32 = ((3.14159 + v166) * float32(float64((v166 * (0.1 + 0.0)))))
                                var v168 uint16 = 621
                                if (true && false) {
                                    var v169 string = ((v21 + v21) + v21)
                                    _ = v169
                                } else {
                                    v170 := GenS18{F0: 416, F1: 17, F2: -398}
                                    var v171 float32 = ((float32(0.0) - (v167 + (float32(3.14159) - v167))) * (v167 - (v167 - (0.0 * (3.14159 - ((v167 / float32((((1.5 / (0.0 - -2.25)) - ((0.1 + (-2.25 / (0.1 - 1e10))) * (0.1 * 0.0))) / -2.25))) / float32(((0.0 - float64(3.14159)) / 1.5))))))))
                                    v9 += v9
                                    _ = v170
                                    _ = v171
                                }
                                if (!(!false)) {
                                    v166 -= float32(((0.1 / (((1e10 / -2.25) / 1e10) + (3.14159 + ((0.1 / (((0.1 - 3.14159) * 0.1) * ((-2.25 + ((-2.25 + (float64(v167) / 3.14159)) / (((3.14159 - (-2.25 / ((0.0 / -2.25) - 1e10))) * 0.1) / (0.0 + -2.25)))) + ((((3.14159 + ((0.0 - 0.0) * 0.0)) + 1.5) + -2.25) - 1.5)))) / 1.0)))) + (float64(v167) - 1e10)))
                                }
                                v166 -= v166
                                _ = fi39
                                _ = v165
                                _ = v166
                                _ = v167
                                _ = v168
                                _ = v40
                            }
                            _ = v38
                        }
                    } else {
                        fmt.Println(v23)
                        if len(v3) > 0 {
                            v3[0] = v21
                        }
                        fmt.Println(v23, v2)
                    }
                    switch v21 {
                    case "hello":
                        for fi172 := 0; fi172 < 11; fi172++ {
                            v18 = (v18 & int32(fi24))
                            v21 = ("foo bar" + "\n\t")
                            _ = fi172
                        }
                    case "\n\t":
                        for fi173 := 0; fi173 < 2; fi173++ {
                            if false {
                                if (!((65535 >= uint16(v13)) || ((v2 ^ (v2 + 1000)) > (v2 + (v2 * (-(v2))))))) {
                                    v23 = (^(v13))
                                    fmt.Println(fi173)
                                } else {
                                    if false {
                                        switch ((v21 + "a") + v21) {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "foo bar":
                                            v9 += (v9 * (^(v18)))
                                        case "a":
                                            v3 = append(v3, "foo bar")
                                        default:
                                            var v174 int16 = (v2 % ((v2) | 1))
                                            _ = v174
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    switch v21 {
                                    case "\n\t":
                                        if ((fi24 / (((fi173 % (((fi24 * -726)) | 1))) | 1)) <= -128) {
                                            v9 -= (v18 + v9)
                                        } else {
                                            switch v21 {
                                            case "foo bar":
                                                v23 = (^(v23))
                                            case "a":
                                                v175 := GenS19{F0: (v18 - v18)}
                                                _ = v175
                                            }
                                            if false {
                                                if (false && ((65535 == 175) && false)) {
                                                    v2 = genFn0()
                                                    var v176 int64 = 1
                                                    v177 := v1[2147483647]
                                                    _ = v176
                                                    _ = v177
                                                }
                                                switch 613 {
                                                case 4294967295:
                                                    if (!(!true)) {
                                                        var v178 uint64 = genFn1(478)
                                                        _ = v178
                                                    } else {
                                                        v23 -= v23
                                                    }
                                                case 0:
                                                    v179 := GenS4{F0: true, F1: -233, F2: 1}
                                                    _ = v179
                                                case -215:
                                                    var v180 int8 = -1
                                                    _ = v180
                                                default:
                                                    switch fi173 {
                                                    case -2147483648:
                                                        v181 := GenS19{F0: v9}
                                                        var v182 uint32 = (v23 - (v23 - v23))
                                                        v23 += v23
                                                        _ = v181
                                                        _ = v182
                                                    case -959:
                                                        v183 := GenS6{F0: 490}
                                                        _ = v183
                                                    }
                                                }
                                                if true {
                                                    switch (fi173 ^ 2147483647) {
                                                    case 9223372036854775807:
                                                        v184 := GenS17{F0: 65535, F1: 0, F2: (919 > 18446744073709551615)}
                                                        _ = v184
                                                    case 32768:
                                                        v185 := GenS16{F0: -1, F1: (!false), F2: 32767}
                                                        _ = v185
                                                    }
                                                } else {
                                                    v186 := GenS20{F0: 1, F1: 28, F2: -23}
                                                    var v187 uint16 = 392
                                                    v23 = 2147483648
                                                    delete(v1, (^(v23)))
                                                    _ = v186
                                                    _ = v187
                                                }
                                                var v188 uint = 225
                                                v1[(v23 >> (((v188 >> ((uint(v23)) & 63))) & 63))] = 681
                                                _ = v188
                                            } else {
                                                v189 := GenS6{F0: 2}
                                                v23 = 1
                                                _ = v189
                                            }
                                        }
                                    case "foo bar":
                                        v23 = (v13 | (v13 ^ v13))
                                    case "hello":
                                        if (false || ((v9 != (-(v18))) && true)) {
                                            if ((((-(v2)) < (v2 % (((-(v2))) | 1))) || true) || true) {
                                                v2 = genFn0()
                                                if true {
                                                    var v190 int = fi173
                                                    v23 += 2
                                                    if true {
                                                        switch fi24 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v2)
                                                        case 4294967295:
                                                            v13 -= (v13 * 0)
                                                        case 127:
                                                            switch v21 {
                                                            case "":
                                                                if (false && (v13 != 669)) {
                                                                    v9 -= -2147483648
                                                                    if (false || ((!((v190 - fi173) >= (fi173 ^ v190))) && true)) {
                                                                        delete(v1, 2)
                                                                        switch fi173 {
                                                                        case 899:
                                                                            if ("\n\t" >= (v21 + (((v21 + v21) + ("a" + "")) + ""))) {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    switch v190 {
                                                                                    case -2147483648:
                                                                                        v191 := GenS7{F0: (v2 * (v2 | -662)), F1: 42, F2: 4294967295}
                                                                                        _ = v191
                                                                                    case 128:
                                                                                        var v192 uint8 = 63
                                                                                        _ = v192
                                                                                    case -1:
                                                                                        switch ((v21 + ("\n\t" + (v21 + ((v21 + v21) + ((v21 + "foo bar") + "\n\t"))))) + (v21 + v21)) {
                                                                                        case "hello":
                                                                                            v13 += (^(v13))
                                                                                        case "\n\t":
                                                                                            if ("\n\t" <= "hello") {
                                                                                                v9 = (v9 + v9)
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = ("a" + (("foo bar" + ((("foo bar" + (v21 + ("hello" + v21))) + v21) + "")) + ("a" + v21)))
                                                                                                }
                                                                                                var v193 float64 = float64(float32((1e10 / (3.14159 * 0.1))))
                                                                                                _ = v193
                                                                                            }
                                                                                        case "":
                                                                                            v18 = (v9 - (v9 + -2147483648))
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        v21 = ("foo bar" + v21)
                                                                                        v194 := GenS21{F0: 403, F1: 42, F2: 255}
                                                                                        v195 := GenS12{F0: (v21 + "foo bar"), F1: genFn0()}
                                                                                        _ = v194
                                                                                        _ = v195
                                                                                    } else {
                                                                                        v18 = (^(v9))
                                                                                        if true {
                                                                                            var v196 int64 = 771
                                                                                            switch (fi173 ^ 243) {
                                                                                            case -2147483648:
                                                                                                switch (v190 - fi173) {
                                                                                                case -992:
                                                                                                    switch (^(v190)) {
                                                                                                    case -2147483648:
                                                                                                        var v197 uint16 = 32767
                                                                                                        _ = v197
                                                                                                    default:
                                                                                                        var v198 int32 = (^(v18))
                                                                                                        _ = v198
                                                                                                    }
                                                                                                case -459:
                                                                                                    var v199 bool = (false && false)
                                                                                                    _ = v199
                                                                                                }
                                                                                            case 32767:
                                                                                                v3 = append(v3, v21)
                                                                                            case 582:
                                                                                                var v200 uint16 = 0
                                                                                                _ = v200
                                                                                            }
                                                                                            _ = v196
                                                                                        }
                                                                                        v201 := GenS22{F0: 0, F1: v18}
                                                                                        _ = v201
                                                                                    }
                                                                                }
                                                                                var v202 uint8 = 1
                                                                                _ = v202
                                                                            }
                                                                        default:
                                                                            v18 -= -1
                                                                            var v203 int32 = -2147483648
                                                                            fmt.Println(v203, fi173)
                                                                            fmt.Println(v13, fi173, v203)
                                                                            v21 = "hello"
                                                                            _ = v203
                                                                        }
                                                                        v18 += 2147483647
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            fmt.Println(v18)
                                                                        case "hello":
                                                                            v2 = (v2 | 953)
                                                                        default:
                                                                            if (41 >= 127) {
                                                                                v204 := GenS23{F0: (v2 - v2), F1: -392, F2: -128}
                                                                                switch (-(fi24)) {
                                                                                case 607:
                                                                                    v2 = v2
                                                                                default:
                                                                                    var v205 uint64 = 902
                                                                                    _ = v205
                                                                                }
                                                                                _ = len(v3)
                                                                                var v206 float64 = (((0.0 / (1.5 * ((float64(1.5) + 1e10) * 3.14159))) + 0.0) + float64(0.0))
                                                                                v13 = 4294967295
                                                                                _ = v204
                                                                                _ = v206
                                                                            }
                                                                            switch "a" {
                                                                            case "hello":
                                                                                var v207 int64 = 42
                                                                                _ = v207
                                                                            case "foo bar":
                                                                                var v208 uint8 = 103
                                                                                _ = v208
                                                                            case "a":
                                                                                _ = len(v1)
                                                                            default:
                                                                                if (98 > 121) {
                                                                                    v23 = 458
                                                                                    v209 := GenS4{F0: (((!false) || (true || true)) || true), F1: 0, F2: 2}
                                                                                    _ = v209
                                                                                } else {
                                                                                    v210 := GenS16{F0: v2, F1: (2 == 18446744073709551615), F2: fi173}
                                                                                    v9 += v9
                                                                                    var v211 float64 = (float64((0.1 / 1.0)) * (0.1 - ((-2.25 * (0.0 / -2.25)) / 1e10)))
                                                                                    _ = v210
                                                                                    _ = v211
                                                                                }
                                                                            }
                                                                            v212 := GenS4{F0: (v9 >= (v18 / ((0) | 1))), F1: 746, F2: 42}
                                                                            _ = v212
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v213 int = fi24
                                                                    if (false && true) {
                                                                        var v214 uint32 = 65535
                                                                        if true {
                                                                            v215 := GenS16{F0: v2, F1: true, F2: 42}
                                                                            v18 = 42
                                                                            var v216 int16 = (v2 >> ((65535) & 63))
                                                                            _ = v215
                                                                            _ = v216
                                                                        } else {
                                                                            var v217 uint16 = 255
                                                                            var v218 uint16 = (v217 ^ v217)
                                                                            _ = v217
                                                                            _ = v218
                                                                        }
                                                                        v219 := GenS24{F0: ("foo bar" + "a"), F1: 42, F2: 42}
                                                                        _ = v214
                                                                        _ = v219
                                                                    } else {
                                                                        if false {
                                                                            switch 127 {
                                                                            case -392:
                                                                                v190 = (v190 << ((0) & 63))
                                                                            case 32768:
                                                                                v3 = append(v3, v21)
                                                                            }
                                                                            if ((!(!(v13 != (^(v13))))) && true) {
                                                                                v190 = (v190 + (v213 >> ((65535) & 63)))
                                                                            } else {
                                                                                var v220 int8 = 1
                                                                                if true {
                                                                                    switch 42 {
                                                                                    case -1:
                                                                                        fmt.Println(v220, v190)
                                                                                    default:
                                                                                        var v221 uint64 = genFn1(523)
                                                                                        _ = v221
                                                                                    }
                                                                                    fmt.Println(v18, fi173)
                                                                                }
                                                                                _ = v220
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v213
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v18, v18)
                                                        }
                                                        switch ((((((v21 + "") + (v21 + v21)) + "") + v21) + v21) + "") {
                                                        case "foo bar":
                                                            var v222 int = fi173
                                                            _ = v222
                                                        }
                                                        v21 = v21
                                                    }
                                                    v3 = append(v3, v21)
                                                    _ = v190
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (v21 + ((((v21 + v21) + v21) + "a") + v21))
                                                }
                                            }
                                            v223 := GenS25{F0: ((("foo bar" + v21) + v21) + "foo bar"), F1: 100, F2: 65535}
                                            var v224 string = ((v21 + ((v21 + v21) + v21)) + ("a" + (("hello" + (v21 + (v21 + v21))) + v21)))
                                            _ = v223
                                            _ = v224
                                        }
                                    default:
                                        var v225 int16 = genFn0()
                                        _ = v225
                                    }
                                    if (!false) {
                                        var v226 int = fi24
                                        _ = v226
                                    } else {
                                        fmt.Println(v9, v23, fi24)
                                        var v227 float32 = (float32(1.5) + 0.0)
                                        var v228 int16 = v2
                                        _ = v227
                                        _ = v228
                                    }
                                }
                                switch (fi24 | fi173) {
                                case -136:
                                    switch ((v21 + v21) + v21) {
                                    case "":
                                        var v229 float32 = 1.5
                                        _ = v229
                                    case "foo bar":
                                        v230 := GenS26{F0: 42, F1: -1}
                                        _ = v230
                                    case "a":
                                        v231 := GenS27{F0: v23}
                                        _ = v231
                                    default:
                                        v21 = v21
                                    }
                                case 42:
                                    if (3.14159 < 0.0) {
                                        v21 = "a"
                                    } else {
                                        var v232 bool = (!(genFn0() >= v2))
                                        fmt.Println(v232, fi173)
                                        switch "\n\t" {
                                        case "a":
                                            v232 = ((v13 << ((471) & 63)) < (v13 - v13))
                                        }
                                        _ = v232
                                    }
                                default:
                                    if (!true) {
                                        v233 := GenS28{F0: -93, F1: (true || false), F2: (v9 * v18)}
                                        if false {
                                            v2 = (v2 / (((^(v2))) | 1))
                                            var v234 float32 = float32(0.0)
                                            v18 += (-(v9))
                                            _ = v234
                                        }
                                        _ = v233
                                    } else {
                                        v235 := v1[v13]
                                        _ = v235
                                    }
                                }
                            } else {
                                var v236 uint32 = v13
                                fmt.Println(v23)
                                fmt.Println(v13)
                                var v237 int16 = (v2 / ((v2) | 1))
                                var v238 string = "hello"
                                v239 := GenS11{F0: (v2 - v2)}
                                v236 += v13
                                _ = v236
                                _ = v237
                                _ = v238
                                _ = v239
                            }
                            var v240 int = fi24
                            v18 = -1
                            _ = fi173
                            _ = v240
                        }
                    case "foo bar":
                        if true {
                            if true {
                                switch fi24 {
                                case 9223372036854775807:
                                    var v241 string = (v21 + "\n\t")
                                    _ = v241
                                case 127:
                                    fmt.Println(fi24, v2)
                                case -824:
                                    v242 := GenS29{F0: 9223372036854775807, F1: (v23 / ((754) | 1))}
                                    _ = v242
                                default:
                                    for fi243 := 0; fi243 < 7; fi243++ {
                                        switch (v21 + "\n\t") {
                                        case "":
                                            v1[873] = 4294967295
                                        case "hello":
                                            fmt.Println(fi243)
                                        default:
                                            var v244 float32 = -2.25
                                            _ = v244
                                        }
                                        if (128 == 253) {
                                            var v245 uint64 = 42
                                            var v246 string = v21
                                            v18 = (v9 * (-(v9)))
                                            _ = v245
                                            _ = v246
                                        }
                                        _ = fi243
                                    }
                                }
                            } else {
                                if true {
                                    fmt.Println(v18, v18)
                                    v9 -= (v9 | v9)
                                } else {
                                    v9 = (v9 & (v9 << ((958) & 63)))
                                }
                                if len(v3) > 0 {
                                    v3[0] = (("foo bar" + (("" + v21) + v21)) + v21)
                                }
                                v247 := GenS13{F0: 318, F1: v23}
                                _ = v247
                            }
                            v248 := GenS24{F0: v21, F1: 18446744073709551615, F2: uint64(v13)}
                            fmt.Println(v2)
                            _ = v248
                        } else {
                            v249 := GenS30{F0: 0, F1: 906, F2: (-(fi24))}
                            _ = v249
                        }
                    }
                    if false {
                        v21 = v21
                    }
                    _ = fi24
                }
                _ = v23
            }
        }
        _ = v22
    }
    v18 += (v9 + (v18 - (^(v18))))
    var v250 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v250:", v250)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v250
    _ = v3
    _ = v9
}
