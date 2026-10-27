package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint; F2 uint }
type GenS8 struct { F0 int32; F1 uint64; F2 int8 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 uint16; F1 bool; F2 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 int64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint8; F1 int32; F2 string }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int32; F1 bool; F2 uint8 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int8; F1 int; F2 string }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32) uint32 {
    if true {
        for fi4 := 0; fi4 < 10; fi4++ {
            if (1 < -1) {
                v5 := GenS1{F0: 461}
                fmt.Println(p0, fi4)
                if ((int8(p0) >= -128) && (32768 > (^(fi4)))) {
                    v6 := GenS2{F0: 566, F1: 0}
                    _ = v6
                }
                _ = v5
            } else {
                if true {
                    v7 := GenS2{F0: 42, F1: (p0 << ((4294967295) & 63))}
                    _ = v7
                }
                v8 := GenS2{F0: 32767, F1: (^(p0))}
                for fi9 := 0; fi9 < 15; fi9++ {
                    if (fi4 > 790) {
                        var v10 float64 = 3.14159
                        v11 := GenS1{F0: 2}
                        var v12 string = ""
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    p0 += 2147483648
                    v13 := GenS1{F0: 42}
                    _ = fi9
                    _ = v13
                }
                _ = v8
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 int32 = 0
                v15 = v15
                v16 := GenS2{F0: -596, F1: (p0 - (^(p0)))}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi4
        }
        v17 := []int{32768, -2147483648, int(p0)}
        if (-9223372036854775808 >= -65) {
            p0 += 0
            fmt.Println(p0)
        }
        _ = v17
    } else {
        p0 -= (^(p0))
        v18 := []int{-1, -1, -421, 32768}
        _ = v18
    }
    if ((true || (32768 <= 811)) || false) {
        p0 -= 660
    }
    _ = p0
    return (p0 / ((62) | 1))
}
func genFn1(p0 uint16, p1 int, p2 bool) int8 {
    switch "\n\t" {
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            switch (fi19 ^ fi19) {
            case -87:
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (fi20 < 2147483647) {
                        var v21 int16 = 1000
                        var v22 float32 = 0.0
                        _ = v21
                        _ = v22
                    } else {
                        p0 -= 133
                        var v23 uint = 1
                        v24 := []int{}
                        var v25 int = fi20
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    p2 = true
                    fmt.Println(p0, p2)
                    _ = fi20
                }
            case 1:
                v26 := []int{(p1 & p1)}
                _ = v26
            }
            fmt.Println(p0)
            if p2 {
                if (127 > 0) {
                    var v27 int64 = -9223372036854775808
                    _ = v27
                } else {
                    switch 32767 {
                    case -333:
                        fmt.Println(p2, p1, fi19)
                    case 9223372036854775807:
                        fmt.Println(p1, p2, p2)
                    case 32767:
                        var v28 int64 = 9223372036854775807
                        _ = v28
                    default:
                        v29 := GenS2{F0: 42, F1: 42}
                        _ = v29
                    }
                    switch (("hello" + "hello") + ((("\n\t" + "\n\t") + "a") + "a")) {
                    case "":
                        var v30 float64 = float64(((1.5 + 0.1) * (((1e10 + 1.5) * 3.14159) / float32((float64(0.1) - float64(1.5))))))
                        _ = v30
                    case "hello":
                        var v31 float32 = 1.5
                        _ = v31
                    case "\n\t":
                        fmt.Println(fi19, p1)
                    default:
                        p0 -= 0
                    }
                    v32 := []int{p1}
                    _ = v32
                }
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := []int{fi33, p1, (fi33 % ((p1) | 1))}
                    v35 := GenS1{F0: 169}
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 int32 = 1
                _ = v36
            }
            _ = fi19
        }
    case "foo bar":
        switch (("" + "hello") + "hello") {
        case "hello":
            var v37 int16 = 1
            _ = v37
        case "":
            v38 := []int{(p1 | 1)}
            _ = v38
        case "a":
            p1 += (p1 | (p1 / ((p1) | 1)))
        }
    case "":
        var v39 int16 = 42
        _ = v39
    default:
        var v40 uint16 = (p0 & 1)
        _ = v40
    }
    p2 = (-832 == 1)
    p2 = false
    var v41 bool = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    return -128
}

func main() {
    var v1 uint16 = 369
    var v2 uint8 = 2
    var v3 bool = (0 >= 503)
    fmt.Println(v3, v3)
    fmt.Println(v3)
    v42 := GenS3{F0: 526, F1: 821, F2: 127}
    if v3 {
        var v43 uint16 = v1
        v43 += v1
        v44 := []int{32768, 9223372036854775807}
        _ = v43
        _ = v44
    }
    fmt.Println(v3)
    v45 := GenS3{F0: -1, F1: 1, F2: 42}
    var v46 int16 = 1
    if v3 {
        switch 32768 {
        case 32768:
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 int16 = (v46 % (((v46 / (((v46 % ((42) | 1))) | 1))) | 1))
                var v49 int8 = 50
                v50 := GenS2{F0: v48, F1: 265}
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
        case 259:
            v51 := []int{9223372036854775807}
            _ = v51
        }
    } else {
        v46 += (^(v46))
    }
    if ((("\n\t" + ("" + "foo bar")) + "a") < (((("\n\t" + "hello") + "foo bar") + "a") + "hello")) {
        if v3 {
            var v52 int32 = 1
            for fi53 := 0; fi53 < 3; fi53++ {
                v54 := GenS4{F0: v2, F1: 746}
                _ = fi53
                _ = v54
            }
            v52 = (v52 % ((v52) | 1))
            _ = v52
        } else {
            var v55 int16 = v46
            _ = v55
        }
        if (false && v3) {
            switch "a" {
            case "hello":
                if (v3 || v3) {
                    for fi56 := 0; fi56 < 10; fi56++ {
                        v57 := []int{}
                        switch ("hello" + "foo bar") {
                        case "":
                            switch fi56 {
                            case 32768:
                                v58 := GenS4{F0: (v2 - v2), F1: 2}
                                _ = v58
                            case -860:
                                var v59 uint32 = 65535
                                _ = v59
                            case 9223372036854775807:
                                v60 := GenS5{F0: "foo bar", F1: uint32(v1), F2: v3}
                                _ = v60
                            }
                        case "foo bar":
                            if false {
                                if v3 {
                                    v46 = v46
                                }
                                var v61 int16 = v46
                                if v3 {
                                    if len(v57) > 0 {
                                        v57[0] = (fi56 >> ((4294967295) & 63))
                                    }
                                    if len(v57) >= 2 {
                                        v57 = v57[1:len(v57)]
                                    }
                                    switch fi56 {
                                    case 128:
                                        for fi62 := 0; fi62 < 7; fi62++ {
                                            if v3 {
                                                v57 = append(v57, -46)
                                                var v63 float64 = 0.1
                                                _ = v63
                                            }
                                            if true {
                                                v64 := GenS6{F0: v61}
                                                if (("hello" + "") > "") {
                                                    if (!v3) {
                                                        if v3 {
                                                            var v65 uint32 = 150
                                                            var v66 int8 = 64
                                                            switch fi56 {
                                                            case -2147483648:
                                                                v65 = 2147483647
                                                            case -653:
                                                                var v67 int64 = -133
                                                                _ = v67
                                                            default:
                                                                v46 += v46
                                                            }
                                                            v66 -= genFn1((v1 % ((2) | 1)), (fi62 ^ fi62), (v2 == v2))
                                                            var v68 uint32 = (v65 % (((v65 - genFn0(1))) | 1))
                                                            _ = v65
                                                            _ = v66
                                                            _ = v68
                                                        }
                                                        v69 := GenS2{F0: int16(v1), F1: 65535}
                                                        if v3 {
                                                            v70 := GenS3{F0: 1, F1: 0, F2: -128}
                                                            var v71 int64 = 9223372036854775807
                                                            var v72 uint16 = v1
                                                            v2 += v2
                                                            if true {
                                                                switch (fi56 - 420) {
                                                                case -9223372036854775808:
                                                                    fmt.Println(v3, v3, v1)
                                                                case -743:
                                                                    _ = len(v57)
                                                                case 116:
                                                                    v73 := GenS3{F0: -2147483648, F1: (v71 + 359), F2: 1}
                                                                    _ = v73
                                                                }
                                                                _ = len(v57)
                                                            }
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            v74 := GenS7{F0: "", F1: 4294967295, F2: 238}
                                                            v3 = true
                                                            v1 -= v1
                                                            _ = v74
                                                        }
                                                        _ = v69
                                                    } else {
                                                        fmt.Println(v3)
                                                        if v3 {
                                                            if false {
                                                                v75 := GenS1{F0: 37}
                                                                fmt.Println(v3, v2, v1)
                                                                _ = v75
                                                            } else {
                                                                var v76 uint16 = v1
                                                                v46 += 22
                                                                _ = v76
                                                            }
                                                        } else {
                                                            v1 = (v1 / ((v1) | 1))
                                                            v2 -= (v2 >> ((100) & 63))
                                                            fmt.Println(v1, v46, fi56)
                                                            if len(v57) >= 2 {
                                                                v57 = v57[1:len(v57)]
                                                            }
                                                        }
                                                        if (0 == 742) {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                if ((0 == int32(v2)) || v3) {
                                                                    fmt.Println(v1)
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch 580 {
                                                                        case 576:
                                                                            var v77 int = 0
                                                                            _ = v77
                                                                        case 237:
                                                                            switch (("" + "\n\t") + ("\n\t" + "")) {
                                                                            case "a":
                                                                                var v78 string = ("" + ("\n\t" + "a"))
                                                                                _ = v78
                                                                            case "\n\t":
                                                                                v57 = append(v57, (^(fi56)))
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        var v79 int16 = (v61 + 42)
                                                                        _ = v79
                                                                    default:
                                                                        switch ("foo bar" + "\n\t") {
                                                                        case "foo bar":
                                                                            if (((v1 % ((v1) | 1)) != (v1 - v1)) || (0 != genFn0(2147483648))) {
                                                                                switch 32767 {
                                                                                case 127:
                                                                                    if (2 < 2) {
                                                                                        if (931 <= 2) {
                                                                                            if v3 {
                                                                                                var v80 uint = 65535
                                                                                                v81 := GenS8{F0: 1, F1: 2, F2: -18}
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v2)
                                                                                                case "a":
                                                                                                    var v82 float64 = 1.5
                                                                                                    var v83 int8 = 100
                                                                                                    fmt.Println(v82, v3)
                                                                                                    v57 = append(v57, 42)
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                case "":
                                                                                                    var v84 uint64 = 18446744073709551615
                                                                                                    _ = v84
                                                                                                }
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            }
                                                                                            v85 := GenS8{F0: 42, F1: 2, F2: 100}
                                                                                            fmt.Println(fi62)
                                                                                            _ = v85
                                                                                        } else {
                                                                                            switch fi62 {
                                                                                            case -9:
                                                                                                if false {
                                                                                                    var v86 float32 = ((-2.25 / (1.5 + 1e10)) * (((0.1 * 0.0) - (-2.25 / 0.1)) + 0.1))
                                                                                                    _ = v86
                                                                                                }
                                                                                            case 675:
                                                                                                var v87 int8 = 42
                                                                                                _ = v87
                                                                                            default:
                                                                                                _ = len(v57)
                                                                                            }
                                                                                        }
                                                                                        if (v1 <= (v1 ^ v1)) {
                                                                                            v3 = v3
                                                                                            if (v3 || (!(v3 && false))) {
                                                                                                if true {
                                                                                                    if v3 {
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            v88 := GenS9{F0: 42, F1: 127}
                                                                                                            _ = v88
                                                                                                        case "":
                                                                                                            switch "\n\t" {
                                                                                                            case "foo bar":
                                                                                                                if (552 >= 395) {
                                                                                                                    switch "hello" {
                                                                                                                    case "a":
                                                                                                                        var v89 bool = (v3 || (18446744073709551615 < 18446744073709551615))
                                                                                                                        v90 := GenS8{F0: -1, F1: 15, F2: -25}
                                                                                                                        _ = v89
                                                                                                                        _ = v90
                                                                                                                    case "hello":
                                                                                                                        switch ("" + "a") {
                                                                                                                        case "foo bar":
                                                                                                                            if v3 {
                                                                                                                                var v91 float64 = 0.0
                                                                                                                                var v92 float32 = 0.0
                                                                                                                                var v93 bool = true
                                                                                                                                _ = v91
                                                                                                                                _ = v92
                                                                                                                                _ = v93
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            var v94 float64 = 0.1
                                                                                                                            _ = v94
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        switch 0 {
                                                                                                                        case 127:
                                                                                                                            var v95 int64 = 1
                                                                                                                            _ = v95
                                                                                                                        case -2147483648:
                                                                                                                            if len(v57) >= 2 {
                                                                                                                                v57 = v57[1:len(v57)]
                                                                                                                            }
                                                                                                                        case 9223372036854775807:
                                                                                                                            fmt.Println(v3, v1, v3)
                                                                                                                        default:
                                                                                                                            if (-9223372036854775808 >= 9223372036854775807) {
                                                                                                                                var v96 float64 = (float64(0.0) - 3.14159)
                                                                                                                                _ = v96
                                                                                                                            } else {
                                                                                                                                var v97 float32 = 1.5
                                                                                                                                _ = len(v57)
                                                                                                                                v98 := GenS5{F0: "foo bar", F1: 4294967295, F2: true}
                                                                                                                                var v99 int32 = 42
                                                                                                                                switch "foo bar" {
                                                                                                                                case "a":
                                                                                                                                    var v100 int8 = 0
                                                                                                                                    _ = v100
                                                                                                                                }
                                                                                                                                _ = v97
                                                                                                                                _ = v98
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v2 -= v2
                                                                                                                        fmt.Println(v61, v3)
                                                                                                                        if (37 != genFn1(2, -2147483648, (-16 <= -219))) {
                                                                                                                            if len(v57) > 0 {
                                                                                                                                v57[0] = 2147483647
                                                                                                                            }
                                                                                                                            fmt.Println(v1, v1, v61)
                                                                                                                            v3 = ((!v3) || ((v1 % (((v1 % ((v1) | 1))) | 1)) <= 0))
                                                                                                                            v1 += (v1 | 412)
                                                                                                                            var v101 uint32 = 42
                                                                                                                            _ = v101
                                                                                                                        }
                                                                                                                        if ((1.5 <= ((1e10 - -2.25) - 1e10)) && (((!(312 == genFn0(303))) && (0.0 < (0.0 * (0.1 / (0.1 + 1e10))))) || ((1.5 >= 1.5) && (false && ((("foo bar" + "\n\t") > "") || false))))) {
                                                                                                                            var v102 uint16 = v1
                                                                                                                            var v103 float32 = 0.1
                                                                                                                            if len(v57) > 0 {
                                                                                                                                v57[0] = fi56
                                                                                                                            }
                                                                                                                            _ = v102
                                                                                                                            _ = v103
                                                                                                                        } else {
                                                                                                                            if true {
                                                                                                                                if false {
                                                                                                                                    v1 -= (v1 << ((793) & 63))
                                                                                                                                }
                                                                                                                                var v104 int = fi62
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v2, v2)
                                                                                                                                    var v105 int32 = 1
                                                                                                                                    var v106 int = v104
                                                                                                                                    switch (fi56 / (((v104 << ((958) & 63))) | 1)) {
                                                                                                                                    case 32767:
                                                                                                                                        var v107 int = (v104 * v106)
                                                                                                                                        _ = v107
                                                                                                                                    case -1:
                                                                                                                                        var v108 string = "\n\t"
                                                                                                                                        _ = v108
                                                                                                                                    case 271:
                                                                                                                                        var v109 uint32 = 65535
                                                                                                                                        _ = v109
                                                                                                                                    default:
                                                                                                                                        if false {
                                                                                                                                            v105 = (-(v105))
                                                                                                                                            var v110 bool = v3
                                                                                                                                            _ = v110
                                                                                                                                        } else {
                                                                                                                                            if false {
                                                                                                                                                fmt.Println(v1)
                                                                                                                                                v111 := GenS2{F0: 0, F1: uint32(v2)}
                                                                                                                                                _ = v111
                                                                                                                                            }
                                                                                                                                            var v112 uint = 65535
                                                                                                                                            _ = v112
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = len(v57)
                                                                                                                                    fmt.Println(v106)
                                                                                                                                    _ = v105
                                                                                                                                    _ = v106
                                                                                                                                }
                                                                                                                                _ = v104
                                                                                                                            } else {
                                                                                                                                if true {
                                                                                                                                    var v113 int8 = -128
                                                                                                                                    if v3 {
                                                                                                                                        var v114 uint = 2
                                                                                                                                        switch fi62 {
                                                                                                                                        case 32768:
                                                                                                                                            v46 += v61
                                                                                                                                        default:
                                                                                                                                            if true {
                                                                                                                                                v113 = 1
                                                                                                                                                v57 = append(v57, fi62)
                                                                                                                                                switch fi56 {
                                                                                                                                                case 32768:
                                                                                                                                                    if len(v57) > 0 {
                                                                                                                                                        v57[0] = -9223372036854775808
                                                                                                                                                    }
                                                                                                                                                case 128:
                                                                                                                                                    var v115 uint32 = 863
                                                                                                                                                    _ = v115
                                                                                                                                                case 536:
                                                                                                                                                    v46 -= 119
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = v114
                                                                                                                                    }
                                                                                                                                    switch "" {
                                                                                                                                    case "":
                                                                                                                                        switch ("hello" + ("hello" + ((("hello" + "hello") + ("hello" + "foo bar")) + "foo bar"))) {
                                                                                                                                        case "a":
                                                                                                                                            if (!(v1 >= v1)) {
                                                                                                                                                var v116 int64 = -1
                                                                                                                                                switch "hello" {
                                                                                                                                                case "hello":
                                                                                                                                                    v61 -= v61
                                                                                                                                                case "a":
                                                                                                                                                    v3 = v3
                                                                                                                                                }
                                                                                                                                                _ = v116
                                                                                                                                            }
                                                                                                                                        case "":
                                                                                                                                            v117 := GenS10{F0: (^(v1)), F1: true, F2: -2147483648}
                                                                                                                                            _ = v117
                                                                                                                                        case "foo bar":
                                                                                                                                            v118 := GenS11{F0: 32767}
                                                                                                                                            _ = v118
                                                                                                                                        default:
                                                                                                                                            _ = len(v57)
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        switch -374 {
                                                                                                                                        case 663:
                                                                                                                                            v119 := GenS12{F0: -1, F1: 1}
                                                                                                                                            _ = v119
                                                                                                                                        case -2147483648:
                                                                                                                                            v120 := GenS3{F0: -382, F1: -620, F2: (v113 * v113)}
                                                                                                                                            var v121 uint = 358
                                                                                                                                            var v122 int16 = -32768
                                                                                                                                            _ = v120
                                                                                                                                            _ = v121
                                                                                                                                            _ = v122
                                                                                                                                        case 32768:
                                                                                                                                            if len(v57) > 0 {
                                                                                                                                                v57[0] = (fi62 & (fi62 & (fi62 - (fi62 >> ((65535) & 63)))))
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            var v123 float32 = 3.14159
                                                                                                                                            _ = v123
                                                                                                                                        }
                                                                                                                                    case "foo bar":
                                                                                                                                        if false {
                                                                                                                                            v124 := GenS13{F0: 42}
                                                                                                                                            v125 := GenS1{F0: 0}
                                                                                                                                            fmt.Println(fi62, v2, v2)
                                                                                                                                            _ = v124
                                                                                                                                            _ = v125
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        if len(v57) >= 2 {
                                                                                                                                            v57 = v57[1:len(v57)]
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v113
                                                                                                                                } else {
                                                                                                                                    var v126 uint32 = 337
                                                                                                                                    _ = v126
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v127 := GenS4{F0: v2, F1: 534}
                                                                                                                            v128 := GenS4{F0: (v2 + (^(v2))), F1: 1}
                                                                                                                            _ = v127
                                                                                                                            _ = v128
                                                                                                                        }
                                                                                                                        var v129 uint8 = (v2 >> ((2) & 63))
                                                                                                                        _ = v129
                                                                                                                    default:
                                                                                                                        if (1.5 > -2.25) {
                                                                                                                            var v130 int8 = 127
                                                                                                                            switch ((("" + "\n\t") + "hello") + ("hello" + "foo bar")) {
                                                                                                                            case "\n\t":
                                                                                                                                v3 = (uint(v2) == 18446744073709551615)
                                                                                                                            }
                                                                                                                            if (42 != 2) {
                                                                                                                                v130 = 42
                                                                                                                                var v131 int64 = -9223372036854775808
                                                                                                                                if len(v57) > 0 {
                                                                                                                                    v57[0] = fi62
                                                                                                                                }
                                                                                                                                _ = v131
                                                                                                                            }
                                                                                                                            _ = v130
                                                                                                                        } else {
                                                                                                                            var v132 int8 = 0
                                                                                                                            _ = v132
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v133 := GenS14{F0: (^(v2)), F1: 1, F2: (("" + (("foo bar" + ("hello" + "a")) + "\n\t")) + ("a" + "\n\t"))}
                                                                                                                    fmt.Println(v3, v1)
                                                                                                                    v134 := GenS15{F0: v1, F1: (-(v61))}
                                                                                                                    fmt.Println(v2)
                                                                                                                    var v135 uint8 = v2
                                                                                                                    _ = v133
                                                                                                                    _ = v134
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                var v136 uint8 = v2
                                                                                                                _ = v136
                                                                                                            default:
                                                                                                                if len(v57) > 0 {
                                                                                                                    v57[0] = fi56
                                                                                                                }
                                                                                                            }
                                                                                                            if (((42 > 1) && v3) || v3) {
                                                                                                                fmt.Println(v46, fi56, v46)
                                                                                                                var v137 bool = (int(v46) >= -2147483648)
                                                                                                                if v3 {
                                                                                                                    if v3 {
                                                                                                                        var v138 uint8 = 1
                                                                                                                        _ = v138
                                                                                                                    } else {
                                                                                                                        v61 = v46
                                                                                                                        switch "hello" {
                                                                                                                        case "a":
                                                                                                                            v61 = -491
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v139 uint32 = genFn0(65535)
                                                                                                                    _ = v139
                                                                                                                } else {
                                                                                                                    _ = len(v57)
                                                                                                                }
                                                                                                                _ = v137
                                                                                                            }
                                                                                                            v140 := GenS7{F0: ("" + "\n\t"), F1: 1, F2: 42}
                                                                                                            _ = v140
                                                                                                        case "hello":
                                                                                                            fmt.Println(v1, fi56)
                                                                                                        default:
                                                                                                            v46 = v46
                                                                                                            if len(v57) > 0 {
                                                                                                                v57[0] = fi62
                                                                                                            }
                                                                                                        }
                                                                                                        _ = len(v57)
                                                                                                        fmt.Println(fi56)
                                                                                                        v2 -= 127
                                                                                                    } else {
                                                                                                        v1 = v1
                                                                                                        fmt.Println(v2, v61, v61)
                                                                                                        v46 = (v46 << ((845) & 63))
                                                                                                    }
                                                                                                    var v141 int = (fi62 >> ((255) & 63))
                                                                                                    _ = v141
                                                                                                }
                                                                                                var v142 int64 = 0
                                                                                                _ = v142
                                                                                            } else {
                                                                                                v143 := GenS2{F0: v61, F1: 238}
                                                                                                if len(v57) > 0 {
                                                                                                    v57[0] = fi62
                                                                                                }
                                                                                                if (v3 && false) {
                                                                                                    switch (fi56 & (fi56 / (((fi56 & fi62)) | 1))) {
                                                                                                    case 1:
                                                                                                        v1 += 32767
                                                                                                    default:
                                                                                                        var v144 int32 = -2147483648
                                                                                                        _ = v144
                                                                                                    }
                                                                                                    if (!v3) {
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            v145 := GenS16{F0: 2147483647, F1: v3, F2: v2}
                                                                                                            _ = v145
                                                                                                        default:
                                                                                                            v146 := GenS11{F0: v1}
                                                                                                            _ = v146
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    if len(v57) >= 2 {
                                                                                                        v57 = v57[1:len(v57)]
                                                                                                    }
                                                                                                    v3 = (!(631 <= 18446744073709551615))
                                                                                                    var v147 uint16 = 32767
                                                                                                    _ = v147
                                                                                                }
                                                                                                _ = v143
                                                                                            }
                                                                                        } else {
                                                                                            var v148 int32 = 293
                                                                                            var v149 string = "\n\t"
                                                                                            v149 = (v149 + v149)
                                                                                            _ = v148
                                                                                            _ = v149
                                                                                        }
                                                                                        v150 := GenS17{F0: 2147483647}
                                                                                        _ = v150
                                                                                    }
                                                                                case 128:
                                                                                    fmt.Println(fi62)
                                                                                default:
                                                                                    switch (("\n\t" + (("\n\t" + ("\n\t" + "")) + "\n\t")) + "") {
                                                                                    case "hello":
                                                                                        v1 = 255
                                                                                    case "":
                                                                                        var v151 int16 = int16(fi56)
                                                                                        _ = v151
                                                                                    default:
                                                                                        v46 += (v46 + (v46 / (((v61 + (v61 | -32768))) | 1)))
                                                                                    }
                                                                                }
                                                                                _ = len(v57)
                                                                                switch (fi56 | (fi56 - -2147483648)) {
                                                                                case -1:
                                                                                    v1 = (v1 << ((200) & 63))
                                                                                case -605:
                                                                                    if ((uint16(v2) > 444) && v3) {
                                                                                        var v152 uint = 507
                                                                                        var v153 int = fi62
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                    }
                                                                                default:
                                                                                    v154 := GenS5{F0: "\n\t", F1: 0, F2: (-32768 < (v46 * int16(fi62)))}
                                                                                    _ = v154
                                                                                }
                                                                                v3 = (v3 && v3)
                                                                            }
                                                                            fmt.Println(v2, v3, fi56)
                                                                        case "\n\t":
                                                                            if (v46 >= v46) {
                                                                                v155 := GenS18{F0: true}
                                                                                _ = v155
                                                                            } else {
                                                                                v156 := GenS9{F0: -2147483648, F1: 0}
                                                                                switch (fi62 / (((fi62 & fi56)) | 1)) {
                                                                                case -9223372036854775808:
                                                                                    if v3 {
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            switch 1 {
                                                                                            case 32768:
                                                                                                if (810 <= 309) {
                                                                                                    fmt.Println(v3)
                                                                                                    var v157 string = ("" + (("a" + ("a" + "foo bar")) + "foo bar"))
                                                                                                    v2 -= 151
                                                                                                    if v3 {
                                                                                                        v157 = "foo bar"
                                                                                                    }
                                                                                                    _ = v157
                                                                                                } else {
                                                                                                    if len(v57) >= 2 {
                                                                                                        v57 = v57[1:len(v57)]
                                                                                                    }
                                                                                                    if (v3 || false) {
                                                                                                        var v158 uint32 = 234
                                                                                                        _ = v158
                                                                                                    } else {
                                                                                                        switch fi56 {
                                                                                                        case -405:
                                                                                                            var v159 uint16 = (v1 << ((868) & 63))
                                                                                                            _ = v159
                                                                                                        case 42:
                                                                                                            fmt.Println(v46)
                                                                                                        default:
                                                                                                            v160 := GenS3{F0: 489, F1: 812, F2: 126}
                                                                                                            _ = v160
                                                                                                        }
                                                                                                    }
                                                                                                    var v161 int = fi62
                                                                                                    _ = v161
                                                                                                }
                                                                                            case 298:
                                                                                                if true {
                                                                                                    v1 = v1
                                                                                                }
                                                                                            default:
                                                                                                switch (fi62 & (fi62 + fi62)) {
                                                                                                case 9223372036854775807:
                                                                                                    v1 -= (^(v1))
                                                                                                case 2147483647:
                                                                                                    fmt.Println(v46)
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            v46 -= (v61 / ((1000) | 1))
                                                                                        case "\n\t":
                                                                                            var v162 int8 = 127
                                                                                            _ = v162
                                                                                        }
                                                                                        fmt.Println(fi56, v61)
                                                                                        if (1.5 != 3.14159) {
                                                                                            var v163 int = -1
                                                                                            v164 := GenS14{F0: (v2 / ((uint8(v1)) | 1)), F1: 86, F2: "hello"}
                                                                                            _ = len(v57)
                                                                                            _ = v163
                                                                                            _ = v164
                                                                                        }
                                                                                    } else {
                                                                                        var v165 uint = 824
                                                                                        v1 = (v1 / ((v1) | 1))
                                                                                        v166 := GenS19{F0: 20, F1: fi62, F2: ""}
                                                                                        _ = v165
                                                                                        _ = v166
                                                                                    }
                                                                                case 4294967295:
                                                                                    var v167 int32 = 2147483647
                                                                                    _ = v167
                                                                                case 127:
                                                                                    v168 := GenS20{F0: int32(v1)}
                                                                                    _ = v168
                                                                                }
                                                                                _ = v156
                                                                            }
                                                                        }
                                                                    }
                                                                    var v169 int8 = -128
                                                                    _ = v169
                                                                } else {
                                                                    var v170 int = 128
                                                                    _ = v170
                                                                }
                                                            }
                                                            var v171 uint16 = 1
                                                            _ = v171
                                                        }
                                                    }
                                                    var v172 int = fi56
                                                    _ = v172
                                                } else {
                                                    v61 = v46
                                                }
                                                if v3 {
                                                    var v173 uint16 = (v1 & v1)
                                                    var v174 int16 = 1000
                                                    _ = v173
                                                    _ = v174
                                                }
                                                _ = v64
                                            } else {
                                                var v175 uint32 = 2147483647
                                                _ = v175
                                            }
                                            v176 := GenS21{F0: 1, F1: genFn1((v1 / ((uint16(v2)) | 1)), fi56, ((true && (18446744073709551615 < 1)) || v3))}
                                            _ = fi62
                                            _ = v176
                                        }
                                    case -422:
                                        v1 -= v1
                                    default:
                                        _ = len(v57)
                                    }
                                } else {
                                    v61 += -1
                                }
                                _ = v61
                            }
                        case "a":
                            fmt.Println(v1, v3, v1)
                        default:
                            fmt.Println(v2, v3)
                        }
                        _ = fi56
                        _ = v57
                    }
                }
            case "":
                if (-2.25 > 3.14159) {
                    v177 := []int{-624, 9223372036854775807}
                    _ = v177
                } else {
                    var v178 uint8 = (v2 | (v2 << ((372) & 63)))
                    v178 = (v178 >> ((2) & 63))
                    _ = v178
                }
            case "\n\t":
                v179 := []int{9223372036854775807, -515, 127, 128}
                _ = v179
            }
            v180 := []int{4294967295, 9223372036854775807, -9223372036854775808}
            _ = v180
        }
        for fi181 := 0; fi181 < 2; fi181++ {
            switch ("hello" + "a") {
            case "foo bar":
                v3 = ((-2.25 < 0.1) || ("a" <= ""))
            }
            if v3 {
                v1 = v1
                v3 = v3
            } else {
                if ((-2.25 == 1.5) && (false && (2 < 1))) {
                    var v182 uint8 = (v2 << ((4294967295) & 63))
                    v183 := []int{0}
                    _ = v182
                    _ = v183
                }
                v184 := []int{}
                _ = v184
            }
            _ = fi181
        }
    }
    v2 += (v2 * 239)
    switch 127 {
    case 128:
        var v185 float32 = ((3.14159 * (0.0 * 1.5)) - 1.5)
        _ = v185
    case 32767:
        v2 += (v2 / (((v2 >> ((2) & 63))) | 1))
    }
    if v3 {
        if false {
            for fi186 := 0; fi186 < 18; fi186++ {
                for fi187 := 0; fi187 < 19; fi187++ {
                    var v188 int8 = -41
                    _ = fi187
                    _ = v188
                }
                if v3 {
                    for fi189 := 0; fi189 < 20; fi189++ {
                        var v190 int64 = 9223372036854775807
                        _ = fi189
                        _ = v190
                    }
                    if true {
                        var v191 int = fi186
                        v46 += v46
                        _ = v191
                    }
                    v3 = (4294967295 != 18446744073709551615)
                } else {
                    for fi192 := 0; fi192 < 18; fi192++ {
                        v2 -= (^(v2))
                        _ = fi192
                    }
                }
                _ = fi186
            }
        } else {
            v193 := []int{127, -9223372036854775808}
            _ = v193
        }
        var v194 int16 = v46
        _ = v194
    }
    if v3 {
        var v195 float64 = 0.1
        _ = v195
    } else {
        v196 := GenS9{F0: 42, F1: -128}
        var v197 uint8 = v2
        v197 += (v197 << ((50) & 63))
        _ = v196
        _ = v197
    }
    for fi198 := 0; fi198 < 11; fi198++ {
        v199 := []int{}
        _ = fi198
        _ = v199
    }
    for fi200 := 0; fi200 < 4; fi200++ {
        v2 -= (^(v2))
        v3 = (1000 >= -32768)
        var v201 float32 = 0.0
        _ = fi200
        _ = v201
    }
    if (false || (-558 > 127)) {
        for fi202 := 0; fi202 < 14; fi202++ {
            fmt.Println(fi202, v1, v46)
            v203 := GenS12{F0: 1, F1: 9223372036854775807}
            _ = fi202
            _ = v203
        }
        if (-997 < 80) {
            var v204 string = "hello"
            v204 = (v204 + v204)
            _ = v204
        } else {
            if (!false) {
                if (1 != 0) {
                    switch "\n\t" {
                    case "a":
                        v1 -= (v1 + v1)
                    }
                } else {
                    v205 := []int{-1, 128, int(v46), 127}
                    v206 := GenS12{F0: -1, F1: 1}
                    var v207 int16 = v46
                    _ = v205
                    _ = v206
                    _ = v207
                }
                v208 := []int{127, -9223372036854775808}
                _ = v208
            }
            for fi209 := 0; fi209 < 5; fi209++ {
                v210 := []int{}
                for fi211 := 0; fi211 < 20; fi211++ {
                    v3 = true
                    switch "" {
                    case "hello":
                        var v212 float32 = (0.1 - (0.1 / 1e10))
                        _ = v212
                    default:
                        switch "" {
                        case "foo bar":
                            var v213 int64 = 0
                            _ = v213
                        case "a":
                            fmt.Println(fi211)
                        case "hello":
                            if len(v210) >= 2 {
                                v210 = v210[1:len(v210)]
                            }
                            if (((0.0 * -2.25) == ((float64(0.0) / (float64(1.5) + 1.5)) * 0.1)) && (0.0 >= 1e10)) {
                                var v214 int8 = -63
                                var v215 int32 = 1
                                _ = v214
                                _ = v215
                            }
                        default:
                            var v216 int8 = -83
                            _ = v216
                        }
                    }
                    _ = fi211
                }
                fmt.Println(v3)
                _ = fi209
                _ = v210
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v45
    _ = v46
}
