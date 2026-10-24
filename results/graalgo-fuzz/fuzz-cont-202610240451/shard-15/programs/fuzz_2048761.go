package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int64; F1 string; F2 int16 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 uint64; F1 uint; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 uint64; F2 uint }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint16; F1 string; F2 uint32 }
type GenS17 struct { F0 uint64; F1 string }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 int16; F1 int16 }
type GenS20 struct { F0 uint16; F1 int64 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int32; F1 uint8 }
type GenS23 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint8, p1 uint, p2 uint) uint64 {
    p1 = p1
    var v5 string = "a"
    v6 := GenS1{F0: 2, F1: (p0 <= (^(p0)))}
    v7 := GenS2{F0: (p0 + p0), F1: -2147483648, F2: 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = 2147483647
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int8) int32 {
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: (p0 << ((65535) & 63)), F1: (217 <= 296)}
        _ = fi10
        _ = v11
    }
    p1 = (^(p1))
    fmt.Println(p0, p0)
    v12 := GenS2{F0: 129, F1: 9223372036854775807, F2: -85}
    _ = p0
    _ = p1
    _ = v12
    return 647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 42
    var v3 int64 = 42
    v4 := []int{}
    v13 := GenS3{F0: genFn1(65535, 0), F1: 4294967295}
    var v14 int = -367
    delete(v1, 2147483647)
    for fi15 := 0; fi15 < 9; fi15++ {
        v3 += (v3 % (((v3 << ((4294967295) & 63))) | 1))
        switch fi15 {
        case 42:
            fmt.Println(v2, v14)
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            v3 += (v3 ^ v3)
            _ = len(v4)
            _ = fi16
        }
        _ = fi15
    }
    if false {
        if true {
            var v17 string = "\n\t"
            var v18 int32 = (v2 ^ (^(v2)))
            var v19 float64 = ((float64(float32(((1.5 / (-2.25 + (1e10 / 1e10))) + -2.25))) / 1.0) - 1.5)
            _ = v17
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 10; fi20++ {
                v3 -= (v3 & (v3 % ((v3) | 1)))
                for fi21 := 0; fi21 < 20; fi21++ {
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        if false {
                            var v22 int16 = 32767
                            if (-2.25 >= 0.1) {
                                v23 := GenS3{F0: (v2 + v2), F1: 4294967295}
                                if (42 >= 127) {
                                    var v24 int8 = 99
                                    v24 = 127
                                    var v25 string = "hello"
                                    _ = v24
                                    _ = v25
                                } else {
                                    if (83 >= 2) {
                                        if (144 != 55) {
                                            v26 := v1[0]
                                            fmt.Println(v14, fi20, v26)
                                            if v26 {
                                                v27 := GenS4{F0: v14, F1: int32(v3), F2: 42}
                                                _ = v27
                                            } else {
                                                v2 = (v2 | v2)
                                                v28 := GenS3{F0: -2147483648, F1: 18446744073709551615}
                                                v29 := v1[2147483647]
                                                _ = v28
                                                _ = v29
                                            }
                                            _ = v26
                                        } else {
                                            switch ("\n\t" + ("foo bar" + ("hello" + ""))) {
                                            case "hello":
                                                fmt.Println(v14, v3, v14)
                                            }
                                            var v30 bool = true
                                            _ = v30
                                        }
                                        var v31 string = ""
                                        _ = v31
                                    } else {
                                        switch (fi20 & 1) {
                                        case -9223372036854775808:
                                            delete(v1, 2147483647)
                                        default:
                                            switch ("a" + "hello") {
                                            case "":
                                                var v32 uint = 1
                                                _ = v32
                                            case "foo bar":
                                                var v33 int32 = 42
                                                _ = v33
                                            case "hello":
                                                v34 := GenS4{F0: (fi20 | (^(fi21))), F1: (v2 ^ (v2 | genFn1(18446744073709551615, 127))), F2: genFn0(177, 1, 361)}
                                                _ = v34
                                            }
                                        }
                                        var v35 uint16 = 65535
                                        var v36 int = (-(v14))
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 uint = 42
                                    _ = v37
                                }
                                _ = v23
                            }
                            _ = v22
                        }
                    case "foo bar":
                        v1[879] = true
                        v38 := GenS5{F0: (((!(1 >= -852)) || true) || (-103 < 1)), F1: 851, F2: (fi21 + -353)}
                        _ = v38
                    case "a":
                        v2 -= -875
                    default:
                        v2 = (v2 & (v2 / (((v2 | (v2 >> ((255) & 63)))) | 1)))
                        if (!(!(32767 > 42))) {
                            var v39 uint16 = 42
                            var v40 string = "foo bar"
                            if false {
                                switch (fi20 + (fi20 ^ v14)) {
                                case 0:
                                    switch ((((v40 + (v40 + "foo bar")) + v40) + "a") + v40) {
                                    case "foo bar":
                                        v39 = (v39 % (((v39 % ((v39) | 1))) | 1))
                                    default:
                                        if true {
                                            var v41 uint8 = 255
                                            var v42 float64 = ((0.0 - 0.1) / 1.5)
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                case 127:
                                    var v43 uint8 = 0
                                    _ = v43
                                case 42:
                                    fmt.Println(fi21, fi20)
                                }
                                v39 -= v39
                            }
                            _ = v39
                            _ = v40
                        }
                        if (("hello" + (("hello" + ("hello" + "")) + (("\n\t" + "hello") + ((("" + ("a" + "")) + "a") + "\n\t")))) >= ("a" + "\n\t")) {
                            if true {
                                v44 := GenS5{F0: (!(2 <= 147)), F1: 65535, F2: 0}
                                _ = v44
                            } else {
                                if (false && (!(false || (!(false || false))))) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (541 > 18446744073709551615) {
                                        fmt.Println(v2)
                                        _ = len(v4)
                                    }
                                }
                                if true {
                                    v14 = (-(v14))
                                    var v45 uint16 = 0
                                    switch "" {
                                    case "a":
                                        if false {
                                            if false {
                                                var v46 int32 = genFn1(2, 0)
                                                v45 = v45
                                                v3 = v3
                                                switch ("a" + ("" + "foo bar")) {
                                                case "hello":
                                                    if (v3 >= v3) {
                                                        switch 0 {
                                                        case 385:
                                                            var v47 int8 = -128
                                                            _ = v47
                                                        case -9223372036854775808:
                                                            var v48 int32 = v46
                                                            _ = v48
                                                        }
                                                    }
                                                default:
                                                    var v49 uint8 = 128
                                                    _ = v49
                                                }
                                                switch "a" {
                                                case "hello":
                                                    switch ((("" + (("" + "\n\t") + ("foo bar" + ""))) + ("" + "foo bar")) + "hello") {
                                                    case "hello":
                                                        switch "a" {
                                                        case "\n\t":
                                                            if ((100 != -128) && ((v46 << ((1) & 63)) >= (v2 * 1))) {
                                                                var v50 int16 = 42
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v45, fi21, v46)
                                                                v51 := GenS6{F0: v45, F1: (fi20 << ((42) & 63))}
                                                                _ = v51
                                                            }
                                                        case "hello":
                                                            switch (("\n\t" + "a") + "a") {
                                                            case "foo bar":
                                                                _ = len(v1)
                                                            default:
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v52 := GenS1{F0: 255, F1: false}
                                                                    _ = v52
                                                                case "a":
                                                                    var v53 uint8 = 255
                                                                    _ = v53
                                                                case "foo bar":
                                                                    var v54 uint = 786
                                                                    if true {
                                                                        if ("foo bar" != "a") {
                                                                            if (!true) {
                                                                                v4 = append(v4, v14)
                                                                            }
                                                                            v14 += (v14 * v14)
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v55 float32 = ((1.5 / (-2.25 * 3.14159)) - (1e10 / ((0.0 + 0.0) + (((1.5 * ((float32(-2.25) / 1e10) * -2.25)) * (0.1 + 3.14159)) / 1.0))))
                                                                                _ = v55
                                                                            case "a":
                                                                                var v56 uint32 = 943
                                                                                _ = v56
                                                                            }
                                                                            var v57 string = ""
                                                                            v58 := GenS3{F0: (v46 % (((v46 >> ((v54) & 63))) | 1)), F1: 570}
                                                                            v59 := GenS1{F0: v54, F1: (!true)}
                                                                            var v60 int64 = v3
                                                                            _ = v57
                                                                            _ = v58
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            v2 += 374
                                                                        }
                                                                        _ = len(v1)
                                                                    }
                                                                    v61 := GenS4{F0: 32768, F1: (^(v46)), F2: 0}
                                                                    _ = v54
                                                                    _ = v61
                                                                default:
                                                                    var v62 uint32 = 815
                                                                    _ = v62
                                                                }
                                                            }
                                                        default:
                                                            switch fi21 {
                                                            case 32767:
                                                                var v63 uint8 = 85
                                                                _ = v63
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v64 int32 = v46
                                                        _ = v64
                                                    case "":
                                                        fmt.Println(fi21, fi21)
                                                    default:
                                                        if (!("" > "")) {
                                                            var v65 uint = 0
                                                            _ = v65
                                                        }
                                                    }
                                                case "":
                                                    v45 = (v45 | v45)
                                                case "\n\t":
                                                    v1[42] = false
                                                default:
                                                    fmt.Println(v3)
                                                }
                                                var v66 bool = true
                                                _ = v46
                                                _ = v66
                                            }
                                            v67 := GenS7{F0: (v3 + (-(v3))), F1: ("hello" + "foo bar"), F2: -928}
                                            _ = v67
                                        }
                                    case "hello":
                                        v1[0] = false
                                    case "\n\t":
                                        var v68 uint8 = 1
                                        var v69 float64 = 3.14159
                                        _ = v68
                                        _ = v69
                                    default:
                                        var v70 float32 = 0.0
                                        _ = v70
                                    }
                                    _ = v45
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v3, fi21)
                                        if false {
                                            switch "" {
                                            case "hello":
                                                fmt.Println(fi20, fi21, v2)
                                            }
                                            var v71 string = ""
                                            var v72 int32 = v2
                                            _ = v71
                                            _ = v72
                                        }
                                        v73 := GenS1{F0: 18446744073709551615, F1: ((1e10 / 1.0) >= 1.5)}
                                        _ = v73
                                    case "a":
                                        var v74 float64 = float64(-2.25)
                                        _ = v74
                                    default:
                                        v14 += fi20
                                    }
                                    var v75 uint64 = 841
                                    _ = v75
                                }
                            }
                        } else {
                            switch (("\n\t" + "foo bar") + "\n\t") {
                            case "":
                                if (-129 < 1) {
                                    v3 = v3
                                }
                            case "foo bar":
                                v76 := v1[112]
                                _ = v76
                            default:
                                v77 := GenS8{F0: 1, F1: ((v2 - v2) <= v2)}
                                _ = v77
                            }
                            v14 = v14
                            fmt.Println(fi21)
                        }
                    }
                    fmt.Println(fi20, v2, fi20)
                    _ = fi21
                }
                var v78 float64 = (1.5 - 0.1)
                _ = fi20
                _ = v78
            }
        }
        v3 = v3
        var v79 int32 = -2147483648
        _ = v79
    }
    v2 -= (v2 * 724)
    switch "a" {
    case "":
        switch "a" {
        case "foo bar":
            var v80 int = -1
            _ = v80
        case "a":
            for fi81 := 0; fi81 < 15; fi81++ {
                v82 := GenS9{F0: fi81, F1: 255}
                _ = len(v4)
                v2 += genFn1(687, 42)
                _ = fi81
                _ = v82
            }
        default:
            switch "foo bar" {
            case "hello":
                if (((v2 << ((2) & 63)) >= v2) && ((1 >= 2147483647) && true)) {
                    if len(v4) > 0 {
                        v4[0] = v14
                    }
                    switch "\n\t" {
                    case "hello":
                        delete(v1, 2147483648)
                    }
                } else {
                    v83 := GenS10{F0: 4294967295, F1: 65535, F2: v14}
                    _ = v83
                }
            case "":
                if false {
                    v84 := GenS11{F0: (v14 | 32768)}
                    v14 = v14
                    _ = v84
                }
            }
        }
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        var v85 int16 = 640
        _ = v85
    }
    if ((((^(v3)) < v3) && true) || true) {
        if (false || true) {
            v3 -= -1
            fmt.Println(v3, v14, v2)
        }
    } else {
        if ("hello" < "hello") {
            fmt.Println(v2, v14)
            if false {
                switch v14 {
                case 9223372036854775807:
                    for fi86 := 0; fi86 < 8; fi86++ {
                        switch "foo bar" {
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = v14
                            }
                        }
                        _ = fi86
                    }
                case -69:
                    var v87 int16 = 1000
                    _ = v87
                default:
                    var v88 bool = ((-(v14)) > v14)
                    _ = v88
                }
                switch (("foo bar" + ("\n\t" + "")) + "") {
                case "":
                    v3 = v3
                case "foo bar":
                    var v89 int = v14
                    _ = v89
                }
                if len(v4) > 0 {
                    v4[0] = 32768
                }
            } else {
                switch (v14 / ((-364) | 1)) {
                case -2147483648:
                    for fi90 := 0; fi90 < 15; fi90++ {
                        if true {
                            if false {
                                _ = len(v4)
                            } else {
                                var v91 uint = 4294967295
                                _ = v91
                            }
                            for fi92 := 0; fi92 < 11; fi92++ {
                                v2 = -2147483648
                                v93 := GenS10{F0: 4294967295, F1: 65535, F2: fi92}
                                _ = fi92
                                _ = v93
                            }
                            if ((v3 | 9223372036854775807) > (^(v3))) {
                                if false {
                                    var v94 int16 = 196
                                    for fi95 := 0; fi95 < 15; fi95++ {
                                        var v96 int32 = v2
                                        _ = fi95
                                        _ = v96
                                    }
                                    switch fi90 {
                                    case 1:
                                        if (true && false) {
                                            var v97 float32 = ((1.5 * (float32((-2.25 + float64(0.1))) * 0.0)) / 1.5)
                                            for fi98 := 0; fi98 < 16; fi98++ {
                                                v99 := GenS8{F0: 1, F1: (65535 < 0)}
                                                var v100 int16 = (-(v94))
                                                v101 := GenS8{F0: 228, F1: true}
                                                _ = fi98
                                                _ = v100
                                                _ = v101
                                                _ = v99
                                            }
                                            var v102 int8 = 115
                                            _ = v102
                                            _ = v97
                                        } else {
                                            var v103 int8 = -128
                                            if (0.1 == 1e10) {
                                                var v104 uint = 2
                                                for fi105 := 0; fi105 < 13; fi105++ {
                                                    if true {
                                                        v103 = v103
                                                        v103 += v103
                                                        if (!(!true)) {
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi90 >> (((v104 / ((696) | 1))) & 63))
                                                            }
                                                            var v106 uint16 = 32767
                                                            _ = v106
                                                        } else {
                                                            var v107 int8 = (^(v103))
                                                            v14 = (^(fi90))
                                                            _ = v107
                                                        }
                                                    } else {
                                                        v104 -= (v104 & (^(v104)))
                                                        fmt.Println(fi90, fi90)
                                                    }
                                                    if (!false) {
                                                        if (65535 < 255) {
                                                            fmt.Println(fi105)
                                                            if (255 < 133) {
                                                                if ((v2 & 2147483647) == (^(v2))) {
                                                                    v14 = fi90
                                                                    var v108 uint8 = 253
                                                                    if len(v4) > 0 {
                                                                        v4[0] = int(v94)
                                                                    }
                                                                    var v109 int = fi105
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                                v110 := v1[0]
                                                                _ = v110
                                                            }
                                                        } else {
                                                            var v111 float64 = 0.1
                                                            if true {
                                                                fmt.Println(v94)
                                                                v111 = v111
                                                                if (!((3.14159 + float32(float64((float32(((v111 - 0.0) / 1.0)) * 0.0)))) == float32(v111))) {
                                                                    _ = len(v4)
                                                                    v3 = (v3 % ((v3) | 1))
                                                                    v14 = 4294967295
                                                                    v112 := v1[2]
                                                                    _ = v112
                                                                }
                                                                switch (v14 >> ((v104) & 63)) {
                                                                case 318:
                                                                    var v113 int32 = (v2 & v2)
                                                                    _ = v113
                                                                case 2147483647:
                                                                    v94 -= (v94 + (^(v94)))
                                                                }
                                                                var v114 int64 = 0
                                                                _ = v114
                                                            } else {
                                                                var v115 int8 = (v103 * (v103 - v103))
                                                                fmt.Println(v3, fi105, v14)
                                                                var v116 float64 = 0.0
                                                                _ = v115
                                                                _ = v116
                                                            }
                                                            _ = v111
                                                        }
                                                        v94 = v94
                                                        fmt.Println(v104, v104, fi90)
                                                    }
                                                    _ = fi105
                                                }
                                                switch ("" + "foo bar") {
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = fi90
                                                    }
                                                case "foo bar":
                                                    delete(v1, 163)
                                                default:
                                                    if true {
                                                        v117 := GenS9{F0: fi90, F1: (^(v104))}
                                                        fmt.Println(v94, v103, v2)
                                                        _ = v117
                                                    } else {
                                                        for fi118 := 0; fi118 < 5; fi118++ {
                                                            var v119 bool = false
                                                            switch -2147483648 {
                                                            case 42:
                                                                if (v119 || false) {
                                                                    v2 = v2
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v120 := v1[uint32(v104)]
                                                                        _ = v120
                                                                    case "":
                                                                        fmt.Println(fi118, v2)
                                                                    default:
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if v119 {
                                                                                v121 := GenS1{F0: v104, F1: (v119 || (3.14159 <= (1.5 * 1.5)))}
                                                                                v122 := GenS12{F0: -40}
                                                                                fmt.Println(fi118)
                                                                                var v123 string = "hello"
                                                                                if v119 {
                                                                                    var v124 uint64 = 2
                                                                                    v14 += (fi90 / (((fi90 * -2147483648)) | 1))
                                                                                    _ = v124
                                                                                } else {
                                                                                    var v125 uint64 = genFn0(1, (v104 / (((^(v104))) | 1)), v104)
                                                                                    if (v119 || v119) {
                                                                                        fmt.Println(fi90)
                                                                                    } else {
                                                                                        var v126 uint16 = 42
                                                                                        _ = v126
                                                                                    }
                                                                                    var v127 string = v123
                                                                                    var v128 uint = v104
                                                                                    v103 -= (v103 | (v103 | (v103 - 100)))
                                                                                    _ = v125
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                                _ = v121
                                                                                _ = v122
                                                                                _ = v123
                                                                            } else {
                                                                                v3 = v3
                                                                            }
                                                                        case "\n\t":
                                                                            var v129 float32 = 3.14159
                                                                            _ = v129
                                                                        default:
                                                                            fmt.Println(v14, fi90)
                                                                        }
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                v2 += 0
                                                            case 9223372036854775807:
                                                                if v119 {
                                                                    fmt.Println(v103, v14)
                                                                }
                                                            default:
                                                                if true {
                                                                    v130 := GenS13{F0: 248, F1: 18446744073709551615, F2: 1}
                                                                    v94 = (v94 ^ v94)
                                                                    var v131 int16 = (v94 >> ((v104) & 63))
                                                                    _ = v130
                                                                    _ = v131
                                                                }
                                                            }
                                                            _ = fi118
                                                            _ = v119
                                                        }
                                                        v103 += v103
                                                    }
                                                }
                                                _ = v104
                                            }
                                            delete(v1, 2147483648)
                                            _ = v103
                                        }
                                    default:
                                        var v132 int8 = 42
                                        _ = v132
                                    }
                                    _ = v94
                                }
                                v133 := GenS9{F0: (v14 & fi90), F1: 159}
                                _ = v133
                            }
                        }
                        v134 := GenS14{F0: 42, F1: 214}
                        fmt.Println(v14, v14)
                        _ = fi90
                        _ = v134
                    }
                case -9223372036854775808:
                    fmt.Println(v3)
                }
                for fi135 := 0; fi135 < 1; fi135++ {
                    v136 := GenS15{F0: (-(v3))}
                    var v137 uint8 = 128
                    v137 = 1
                    _ = fi135
                    _ = v136
                    _ = v137
                }
                v2 = (v2 + 42)
            }
            if false {
                if (-387 == -32768) {
                    v138 := GenS16{F0: 65535, F1: "\n\t", F2: 42}
                    _ = v138
                } else {
                    v139 := GenS16{F0: 361, F1: "hello", F2: 828}
                    _ = v139
                }
                v3 -= (v3 | (-(v3)))
            } else {
                if (((v14 * v14) != 128) || ((!true) && false)) {
                    v14 += v14
                    if false {
                        var v140 uint16 = 45
                        _ = v140
                    }
                } else {
                    v14 = v14
                    var v141 float32 = (3.14159 / (0.0 + 0.1))
                    _ = v141
                }
                v14 += v14
            }
        }
    }
    for fi142 := 0; fi142 < 9; fi142++ {
        v1[2147483648] = false
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v143 int32 = 966
        _ = fi142
        _ = v143
    }
    var v144 uint64 = 376
    var v145 int16 = 42
    if (int32(v3) != v2) {
        if true {
            switch v14 {
            case -9223372036854775808:
                v2 = v2
            case 32767:
                switch "" {
                case "":
                    v4 = append(v4, 931)
                default:
                    for fi146 := 0; fi146 < 9; fi146++ {
                        if (!(((3.14159 + 1e10) == 0.0) && true)) {
                            if false {
                                if (!((0.1 >= float32(3.14159)) && false)) {
                                    if false {
                                        v147 := GenS17{F0: (^(v144)), F1: (("foo bar" + "a") + "foo bar")}
                                        if len(v4) > 0 {
                                            v4[0] = (v14 ^ fi146)
                                        }
                                        _ = v147
                                    }
                                    switch (^(v14)) {
                                    case 0:
                                        fmt.Println(v14)
                                    case 42:
                                        switch (fi146 >> ((4294967295) & 63)) {
                                        case -555:
                                            switch v14 {
                                            case -474:
                                                for fi148 := 0; fi148 < 18; fi148++ {
                                                    var v149 uint8 = 1
                                                    _ = fi148
                                                    _ = v149
                                                }
                                            case 580:
                                                v150 := GenS4{F0: v14, F1: (v2 ^ (v2 << ((1) & 63))), F2: 18446744073709551615}
                                                _ = v150
                                            case -2147483648:
                                                v1[109] = true
                                            }
                                        case 128:
                                            if len(v4) > 0 {
                                                v4[0] = v14
                                            }
                                        default:
                                            for fi151 := 0; fi151 < 1; fi151++ {
                                                var v152 uint32 = 525
                                                var v153 uint16 = 32767
                                                fmt.Println(v152, v145)
                                                if true {
                                                    v154 := v1[v152]
                                                    if v154 {
                                                        if (1 != 1) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v154 = ((0.1 * 1.5) > 1.5)
                                                            case "foo bar":
                                                                v144 += (^(v144))
                                                            case "hello":
                                                                switch (fi146 * v14) {
                                                                case -2147483648:
                                                                    if true {
                                                                        v144 -= 311
                                                                        v145 = int16(v3)
                                                                        v154 = v154
                                                                    }
                                                                case 4294967295:
                                                                    if v154 {
                                                                        if (int8(v3) >= 127) {
                                                                            var v155 float32 = 1e10
                                                                            if v154 {
                                                                                var v156 uint32 = 4294967295
                                                                                _ = v156
                                                                            } else {
                                                                                fmt.Println(v152, v152)
                                                                                v155 = float32((1.5 * -2.25))
                                                                            }
                                                                            fmt.Println(v154)
                                                                            _ = v155
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v154, v14)
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch int(v2) {
                                                                        case -1:
                                                                            var v157 uint = 65535
                                                                            _ = v157
                                                                        case 4294967295:
                                                                            v14 = fi146
                                                                        default:
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                fmt.Println(v154, v3, v14)
                                                                            default:
                                                                                v158 := GenS18{F0: "hello", F1: v153}
                                                                                _ = v158
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                var v159 uint32 = (v152 * (v152 ^ 4294967295))
                                                                _ = v159
                                                            default:
                                                                fmt.Println(fi146)
                                                            }
                                                            var v160 float64 = (1e10 * float64(((-2.25 / 3.14159) - ((0.0 / 1e10) - float32((((-2.25 / 1.0) + 1.5) * 0.1))))))
                                                            v161 := v1[2147483647]
                                                            v3 += (v3 >> ((0) & 63))
                                                            if (7 < -1) {
                                                                if false {
                                                                    var v162 int64 = (v3 << ((18446744073709551615) & 63))
                                                                    _ = len(v4)
                                                                    _ = v162
                                                                }
                                                                fmt.Println(v153, v145)
                                                                var v163 string = "hello"
                                                                v14 = (v14 / (((v14 >> ((18446744073709551615) & 63))) | 1))
                                                                _ = v163
                                                            }
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                        var v164 int = -2147483648
                                                        v165 := GenS18{F0: "", F1: (^(v153))}
                                                        _ = v164
                                                        _ = v165
                                                    }
                                                    _ = v154
                                                } else {
                                                    v14 = v14
                                                }
                                                if true {
                                                    if (!true) {
                                                        if ((v14 << ((2) & 63)) != (fi146 >> ((65535) & 63))) {
                                                            switch (v14 / ((fi151) | 1)) {
                                                            case 32767:
                                                                if ((2 < 13) || true) {
                                                                    var v166 string = ""
                                                                    _ = len(v1)
                                                                    fmt.Println(v14, v145)
                                                                    _ = v166
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v167 int8 = -128
                                                                    v152 += v152
                                                                    _ = v167
                                                                }
                                                            }
                                                            if false {
                                                                v153 = 2
                                                                switch "hello" {
                                                                case "":
                                                                    v168 := GenS5{F0: false, F1: 65535, F2: (fi151 ^ 127)}
                                                                    _ = v168
                                                                case "a":
                                                                    _ = len(v4)
                                                                case "\n\t":
                                                                    var v169 bool = true
                                                                    _ = v169
                                                                default:
                                                                    v153 += v153
                                                                }
                                                            }
                                                            _ = len(v1)
                                                        } else {
                                                            v152 = (v152 ^ v152)
                                                            var v170 int16 = v145
                                                            var v171 bool = (2 == 53)
                                                            _ = v170
                                                            _ = v171
                                                        }
                                                        var v172 bool = (v145 == v145)
                                                        var v173 bool = (186 > 0)
                                                        _ = v172
                                                        _ = v173
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            _ = len(v4)
                                                        }
                                                        fmt.Println(fi146, v152)
                                                        if false {
                                                            fmt.Println(v14, fi151, v14)
                                                            if ((("hello" + ("foo bar" + ((("\n\t" + ("a" + "")) + "a") + "\n\t"))) + (("\n\t" + "a") + "\n\t")) < "hello") {
                                                                if true {
                                                                    v174 := GenS5{F0: true, F1: 255, F2: fi146}
                                                                    if false {
                                                                        if false {
                                                                            var v175 float32 = float32(3.14159)
                                                                            _ = v175
                                                                        }
                                                                        v14 -= (fi151 / ((42) | 1))
                                                                    } else {
                                                                        var v176 bool = ((v153 > v153) || false)
                                                                        _ = v176
                                                                    }
                                                                    fmt.Println(v145, v153, v153)
                                                                    fmt.Println(fi146)
                                                                    _ = v174
                                                                }
                                                            } else {
                                                                var v177 uint = 778
                                                                _ = v177
                                                            }
                                                        } else {
                                                            v152 = 42
                                                            fmt.Println(v153, v144, v14)
                                                            var v178 string = "a"
                                                            _ = v178
                                                        }
                                                        var v179 uint8 = 128
                                                        _ = v179
                                                    }
                                                    v180 := GenS11{F0: (^(fi151))}
                                                    v181 := GenS17{F0: (v144 * (v144 ^ v144)), F1: ((("\n\t" + ("" + ("hello" + "foo bar"))) + ("a" + ("\n\t" + ("a" + "\n\t")))) + "hello")}
                                                    fmt.Println(v144, fi151)
                                                    _ = v180
                                                    _ = v181
                                                }
                                                var v182 int = (v14 - (-(fi151)))
                                                _ = fi151
                                                _ = v152
                                                _ = v153
                                                _ = v182
                                            }
                                        }
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 127
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = v14
                                    }
                                }
                                if false {
                                    if true {
                                        switch "\n\t" {
                                        case "hello":
                                            v145 += (v145 << ((839) & 63))
                                        case "a":
                                            fmt.Println(v145, v2)
                                        default:
                                            for fi183 := 0; fi183 < 18; fi183++ {
                                                v2 = -636
                                                var v184 int32 = (v2 - -2147483648)
                                                switch ((("foo bar" + "foo bar") + "\n\t") + "hello") {
                                                case "\n\t":
                                                    delete(v1, 2147483648)
                                                case "a":
                                                    switch (fi183 & (v14 << ((uint(v144)) & 63))) {
                                                    case 599:
                                                        v2 = (v2 / ((int32(fi183)) | 1))
                                                    }
                                                case "hello":
                                                    switch (v14 & 128) {
                                                    case 0:
                                                        v185 := v1[335]
                                                        _ = v185
                                                    case 9223372036854775807:
                                                        v186 := GenS19{F0: v145, F1: v145}
                                                        _ = v186
                                                    default:
                                                        switch -1 {
                                                        case 127:
                                                            switch "a" {
                                                            case "foo bar":
                                                                if len(v4) > 0 {
                                                                    v4[0] = fi183
                                                                }
                                                            case "hello":
                                                                var v187 int8 = 100
                                                                _ = v187
                                                            case "\n\t":
                                                                v188 := GenS20{F0: 47, F1: 9223372036854775807}
                                                                _ = v188
                                                            default:
                                                                var v189 uint = 435
                                                                _ = v189
                                                            }
                                                        case 1:
                                                            v1[2] = true
                                                        case 9223372036854775807:
                                                            if (65535 > 650) {
                                                                var v190 string = "\n\t"
                                                                var v191 string = "a"
                                                                _ = v190
                                                                _ = v191
                                                            } else {
                                                                fmt.Println(fi183, v145, v3)
                                                            }
                                                        }
                                                        v192 := GenS13{F0: 42, F1: (v144 >> ((782) & 63)), F2: 606}
                                                        _ = v192
                                                    }
                                                default:
                                                    _ = len(v1)
                                                    var v193 float64 = 3.14159
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v193
                                                }
                                                if true {
                                                    fmt.Println(fi183, v3)
                                                } else {
                                                    var v194 bool = (94 >= 127)
                                                    if v194 {
                                                        v14 = fi146
                                                        v184 = v184
                                                    }
                                                    _ = v194
                                                }
                                                _ = fi183
                                                _ = v184
                                            }
                                        }
                                        var v195 float64 = (3.14159 + 3.14159)
                                        _ = v195
                                    } else {
                                        v144 -= (v144 | v144)
                                        v3 -= v3
                                        switch (v14 << ((767) & 63)) {
                                        case -1:
                                            for fi196 := 0; fi196 < 16; fi196++ {
                                                delete(v1, 2147483647)
                                                _ = fi196
                                            }
                                        case 32767:
                                            var v197 int8 = -128
                                            _ = v197
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                    var v198 int32 = (^(v2))
                                    _ = v198
                                } else {
                                    for fi199 := 0; fi199 < 7; fi199++ {
                                        var v200 int64 = v3
                                        _ = fi199
                                        _ = v200
                                    }
                                    if false {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        for fi201 := 0; fi201 < 10; fi201++ {
                                            if true {
                                                var v202 uint32 = 256
                                                _ = v202
                                            }
                                            if false {
                                                fmt.Println(v144)
                                                var v203 int16 = (v145 << ((255) & 63))
                                                _ = v203
                                            }
                                            _ = fi201
                                        }
                                        delete(v1, 206)
                                    } else {
                                        fmt.Println(v14, v144)
                                        v2 = (v2 & genFn1(18446744073709551615, -125))
                                    }
                                    v204 := GenS21{F0: -9223372036854775808}
                                    _ = v204
                                }
                            }
                            if (false && true) {
                                v205 := GenS22{F0: v2, F1: 20}
                                _ = v205
                            } else {
                                if false {
                                    for fi206 := 0; fi206 < 17; fi206++ {
                                        v145 = (v145 ^ v145)
                                        switch ("" + "\n\t") {
                                        case "a":
                                            v14 = fi206
                                        case "foo bar":
                                            var v207 string = ("" + ("hello" + "\n\t"))
                                            _ = v207
                                        case "":
                                            if false {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v208 uint16 = 97
                                                _ = v208
                                            } else {
                                                var v209 uint16 = 2
                                                v3 -= (v3 + (v3 + v3))
                                                _ = v209
                                            }
                                        }
                                        v14 += (v14 / (((v14 % ((fi206) | 1))) | 1))
                                        _ = fi206
                                    }
                                } else {
                                    fmt.Println(fi146)
                                    if (42 > uint(v144)) {
                                        switch ("a" + ("foo bar" + "foo bar")) {
                                        case "foo bar":
                                            v210 := GenS1{F0: 191, F1: (127 < 127)}
                                            _ = v210
                                        }
                                    }
                                }
                            }
                            v211 := GenS23{F0: 42, F1: 65535, F2: 0}
                            _ = v211
                        } else {
                            var v212 int32 = genFn1(255, 127)
                            _ = len(v1)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v212
                        }
                        var v213 int = (fi146 % ((v14) | 1))
                        _ = fi146
                        _ = v213
                    }
                }
            default:
                for fi214 := 0; fi214 < 2; fi214++ {
                    v4 = append(v4, (fi214 / ((128) | 1)))
                    var v215 string = "foo bar"
                    _ = fi214
                    _ = v215
                }
            }
        } else {
            var v216 int = int(v145)
            for fi217 := 0; fi217 < 6; fi217++ {
                v2 = (v2 * (v2 ^ v2))
                _ = fi217
            }
            for fi218 := 0; fi218 < 18; fi218++ {
                v145 = (v145 << ((65535) & 63))
                v219 := GenS19{F0: -528, F1: v145}
                v145 -= v145
                _ = fi218
                _ = v219
            }
            _ = v216
        }
        for fi220 := 0; fi220 < 6; fi220++ {
            var v221 int16 = 42
            _ = fi220
            _ = v221
        }
        switch 9223372036854775807 {
        case -2147483648:
            v144 = 18446744073709551615
        case 42:
            v145 -= (v145 >> ((42) & 63))
        case 32768:
            for fi222 := 0; fi222 < 3; fi222++ {
                var v223 float64 = float64((0.1 * 1e10))
                _ = fi222
                _ = v223
            }
        }
    } else {
        if (v3 <= (v3 + (^(v3)))) {
            if (false || (false && (!false))) {
                var v224 float32 = (((3.14159 * ((1e10 - (((0.1 + 1e10) / (1.5 * 3.14159)) - 1e10)) / (3.14159 * 3.14159))) - (3.14159 * 1e10)) + 1e10)
                _ = v224
            }
        } else {
            var v225 uint32 = 4294967295
            fmt.Println(v225, v3)
            _ = v225
        }
        for fi226 := 0; fi226 < 2; fi226++ {
            if (!(genFn0(2, 42, 0) <= genFn0(2, 2, 114))) {
                v227 := GenS7{F0: (v3 % (((v3 | (v3 + v3))) | 1)), F1: "foo bar", F2: (v145 << ((42) & 63))}
                v144 = genFn0(2, uint(v144), 65535)
                _ = len(v1)
                _ = v227
            } else {
                switch ("hello" + (("" + "\n\t") + "\n\t")) {
                case "foo bar":
                    fmt.Println(v145)
                }
                v228 := GenS24{F0: 4294967295}
                v2 -= v2
                _ = v228
            }
            if false {
                for fi229 := 0; fi229 < 16; fi229++ {
                    v4 = append(v4, int(v2))
                    if (v14 <= (fi226 + (v14 / ((-516) | 1)))) {
                        var v230 int = fi226
                        _ = v230
                    } else {
                        switch "a" {
                        case "a":
                            fmt.Println(v2, fi226, v3)
                        case "\n\t":
                            fmt.Println(v144, v145, v14)
                        case "":
                            v4 = append(v4, (^(fi226)))
                        default:
                            if ((42 < 42) && false) {
                                v144 -= v144
                            }
                        }
                        var v231 bool = (!false)
                        v231 = v231
                        _ = v231
                    }
                    var v232 uint16 = 2
                    v233 := GenS23{F0: 42, F1: 65535, F2: 2147483648}
                    _ = fi229
                    _ = v232
                    _ = v233
                }
                v145 = v145
            }
            _ = fi226
        }
        for fi234 := 0; fi234 < 12; fi234++ {
            v3 = v3
            var v235 float32 = -2.25
            _ = fi234
            _ = v235
        }
    }
    for fi236 := 0; fi236 < 4; fi236++ {
        v237 := GenS25{F0: (18446744073709551615 < 531), F1: ((v14 | v14) < 4294967295), F2: true}
        _ = fi236
        _ = v237
    }
    var v238 uint8 = uint8(v144)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v238:", v238)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v144
    _ = v145
    _ = v2
    _ = v238
    _ = v3
    _ = v4
}
