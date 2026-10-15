package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint32; F2 int32 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 int64; F2 uint64 }
type GenS11 struct { F0 int8; F1 uint; F2 uint }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0(p0 int64) uint32 {
    v5 := []int{171}
    p0 = p0
    _ = len(v5)
    p0 -= p0
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 uint16) uint8 {
    var v6 uint32 = 0
    v7 := []int{9223372036854775807}
    v7 = append(v7, 128)
    var v8 float64 = 0.1
    switch ("\n\t" + "a") {
    case "":
        var v9 string = (("foo bar" + (("" + ("" + ("hello" + "hello"))) + "foo bar")) + "a")
        _ = v9
    case "hello":
        v10 := GenS1{F0: 0, F1: true}
        _ = v10
    default:
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int16 = 811
            v13 := GenS1{F0: (fi11 << ((18446744073709551615) & 63)), F1: (-127 >= 1)}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 0
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = 520
    v1 = v1
    var v14 int = 317
    switch "a" {
    case "a":
        fmt.Println(v3)
    case "foo bar":
        switch (v14 & (v14 ^ 396)) {
        case 0:
            if false {
                switch v14 {
                case -2147483648:
                    v15 := GenS1{F0: -818, F1: false}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 11; fi16++ {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if true {
                                if true {
                                    var v18 uint8 = 128
                                    var v19 uint64 = (v3 >> ((uint(v4)) & 63))
                                    fmt.Println(v1, v18, v14)
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(v4, v14)
                                    v14 = (v14 & 1)
                                }
                            }
                            _ = fi17
                        }
                        switch "a" {
                        case "":
                            v14 = (v14 | int(v1))
                        case "hello":
                            var v20 int32 = 0
                            _ = v20
                        case "foo bar":
                            switch (("foo bar" + "hello") + "a") {
                            case "hello":
                                var v21 int32 = 42
                                _ = v21
                            case "":
                                var v22 uint64 = 591
                                _ = v22
                            default:
                                var v23 uint = 4294967295
                                _ = v23
                            }
                        default:
                            if (((-2.25 * (0.1 / (3.14159 - ((1.5 / 1e10) / (-2.25 / ((0.0 + 1.5) - (0.0 + -2.25))))))) >= 3.14159) || false) {
                                v24 := []int{}
                                var v25 int8 = 42
                                _ = v24
                                _ = v25
                            } else {
                                v2 = (v2 % (((v2 * -813)) | 1))
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    v2 = -571
                                    if true {
                                        v4 += v4
                                        if (false || ((v2 << ((2) & 63)) == v2)) {
                                            v14 = fi26
                                        }
                                    }
                                    if (false || true) {
                                        var v27 bool = true
                                        switch ("" + (("foo bar" + "foo bar") + "")) {
                                        case "foo bar":
                                            v28 := []int{fi26, v14}
                                            _ = v28
                                        case "":
                                            if false {
                                                switch fi26 {
                                                case 42:
                                                    v4 += v4
                                                    var v29 bool = v27
                                                    if false {
                                                        v30 := []int{}
                                                        v4 += v4
                                                        v2 -= -32768
                                                        _ = v30
                                                    }
                                                    _ = v29
                                                default:
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(fi16, v14, v14)
                                                    case "":
                                                        v27 = (!(!true))
                                                    case "foo bar":
                                                        v14 += fi26
                                                    default:
                                                        v31 := GenS1{F0: fi16, F1: (0.1 == float64((3.14159 - (float32(float64(0.1)) - (float32(((1.5 / 0.1) / (3.14159 - 1.5))) - 0.0)))))}
                                                        _ = v31
                                                    }
                                                }
                                            }
                                            var v32 uint = 2
                                            v32 += (v32 & v32)
                                            v33 := GenS2{F0: 0}
                                            switch "foo bar" {
                                            case "a":
                                                var v34 uint64 = uint64(v32)
                                                _ = v34
                                            case "\n\t":
                                                v35 := GenS3{F0: 0, F1: 0, F2: 127}
                                                _ = v35
                                            }
                                            _ = v32
                                            _ = v33
                                        case "a":
                                            v4 += genFn0(v1)
                                        }
                                        _ = v27
                                    } else {
                                        if (1 < -2147483648) {
                                            var v36 uint8 = 0
                                            _ = v36
                                        } else {
                                            v3 += (v3 & (v3 * v3))
                                            v37 := GenS4{F0: 404, F1: "a"}
                                            _ = v37
                                        }
                                    }
                                    _ = fi26
                                }
                                v38 := []int{4294967295, fi16, -398, 2147483647}
                                _ = v38
                            }
                        }
                        fmt.Println(fi16, v14)
                        _ = fi16
                    }
                }
                v14 += (v14 % ((128) | 1))
                v39 := GenS5{F0: "\n\t", F1: genFn1(255)}
                _ = v39
            }
        case 4294967295:
            if (true || true) {
                var v40 uint64 = 42
                for fi41 := 0; fi41 < 16; fi41++ {
                    var v42 float32 = 3.14159
                    v43 := GenS3{F0: 0, F1: (v3 % ((v3) | 1)), F2: -117}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            } else {
                var v44 bool = (!(v3 >= v3))
                switch ("foo bar" + "") {
                case "foo bar":
                    for fi45 := 0; fi45 < 19; fi45++ {
                        if false {
                            fmt.Println(v44, v2, v4)
                        } else {
                            for fi46 := 0; fi46 < 2; fi46++ {
                                v47 := GenS4{F0: 2, F1: ""}
                                switch (fi46 >> ((1) & 63)) {
                                case 42:
                                    var v48 bool = (128 != 127)
                                    _ = v48
                                case 128:
                                    v49 := GenS6{F0: "a"}
                                    _ = v49
                                case 127:
                                    if false {
                                        var v50 int8 = 109
                                        _ = v50
                                    }
                                    switch (-(fi46)) {
                                    case 32767:
                                        v14 = (fi46 | (fi46 / (((fi45 % ((fi46) | 1))) | 1)))
                                    case 128:
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if false {
                                                        v51 := []int{v14}
                                                        var v52 int = fi46
                                                        v4 -= genFn0(int64(v4))
                                                        if len(v51) > 0 {
                                                            v51[0] = 42
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                case "hello":
                                                    v53 := GenS5{F0: ("a" + "foo bar"), F1: 42}
                                                    _ = v53
                                                }
                                            case "foo bar":
                                                v54 := GenS1{F0: -9223372036854775808, F1: (!(genFn0(v1) > (v4 * (v4 >> ((144) & 63)))))}
                                                _ = v54
                                            }
                                            if v44 {
                                                v44 = v44
                                            } else {
                                                v14 = (v14 - v14)
                                                fmt.Println(v3)
                                                v55 := []int{}
                                                _ = v55
                                            }
                                            if v44 {
                                                v14 -= 1
                                                v56 := []int{(fi45 / ((-9223372036854775808) | 1))}
                                                var v57 float32 = 1.5
                                                _ = v56
                                                _ = v57
                                            } else {
                                                var v58 int = (v14 >> ((uint(v3)) & 63))
                                                _ = v58
                                            }
                                        } else {
                                            v1 += (^(v1))
                                        }
                                    case -701:
                                        v59 := GenS4{F0: 65535, F1: "hello"}
                                        _ = v59
                                    default:
                                        var v60 int8 = -1
                                        fmt.Println(v1, v3)
                                        fmt.Println(v14, v60, v1)
                                        if true {
                                            if (!v44) {
                                                var v61 int32 = -945
                                                if true {
                                                    fmt.Println(fi45)
                                                    fmt.Println(fi45, fi46)
                                                }
                                                _ = v61
                                            }
                                            fmt.Println(fi45)
                                            if (!("a" == ("hello" + (("\n\t" + "\n\t") + ("hello" + ("" + "foo bar")))))) {
                                                var v62 uint16 = 65535
                                                _ = v62
                                            } else {
                                                var v63 uint = 985
                                                var v64 int64 = (v1 % (((v1 - int64(v2))) | 1))
                                                _ = v63
                                                _ = v64
                                            }
                                        } else {
                                            var v65 int16 = (v2 << ((18446744073709551615) & 63))
                                            fmt.Println(v1, fi46, v60)
                                            v4 = v4
                                            _ = v65
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            v66 := GenS2{F0: 18446744073709551615}
                                            _ = v66
                                        }
                                        _ = v60
                                    }
                                default:
                                    v1 = (v1 << ((4294967295) & 63))
                                }
                                v67 := GenS1{F0: v14, F1: false}
                                v44 = ((!((2147483647 <= -641) || true)) || false)
                                _ = fi46
                                _ = v47
                                _ = v67
                            }
                            v2 = (v2 & (-(v2)))
                        }
                        v68 := GenS1{F0: (v14 << ((1) & 63)), F1: v44}
                        for fi69 := 0; fi69 < 13; fi69++ {
                            v70 := []int{128}
                            _ = fi69
                            _ = v70
                        }
                        _ = fi45
                        _ = v68
                    }
                default:
                    v3 += v3
                }
                _ = v44
            }
        case 500:
            for fi71 := 0; fi71 < 10; fi71++ {
                switch "\n\t" {
                case "foo bar":
                    if true {
                        v72 := GenS7{F0: v2, F1: genFn0((v1 - 9223372036854775807)), F2: 771}
                        switch (fi71 + 127) {
                        case 42:
                            v2 += 0
                        }
                        var v73 int = v14
                        _ = v72
                        _ = v73
                    }
                case "a":
                    v14 += fi71
                case "":
                    var v74 int8 = int8(v2)
                    _ = v74
                }
                for fi75 := 0; fi75 < 16; fi75++ {
                    if false {
                        v1 -= (v1 - 0)
                        v3 = (^(v3))
                        v76 := []int{}
                        _ = v76
                    } else {
                        var v77 int32 = -718
                        v2 -= (v2 ^ (v2 & v2))
                        var v78 int64 = v1
                        v79 := GenS8{F0: v1, F1: 2}
                        v3 += 18
                        var v80 int8 = 1
                        _ = v77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    var v81 int16 = -1
                    _ = fi75
                    _ = v81
                }
                for fi82 := 0; fi82 < 13; fi82++ {
                    v1 = (v1 / ((v1) | 1))
                    _ = fi82
                }
                _ = fi71
            }
        default:
            if (!(("\n\t" + "") >= "hello")) {
                v83 := []int{-9223372036854775808, 1}
                if (!false) {
                    fmt.Println(v3, v3)
                }
                if len(v83) >= 2 {
                    v83 = v83[1:len(v83)]
                }
                _ = v83
            }
        }
    case "":
        v84 := GenS2{F0: (v3 << ((42) & 63))}
        _ = v84
    }
    switch ("a" + (("\n\t" + "hello") + "\n\t")) {
    case "hello":
        var v85 int64 = -9223372036854775808
        _ = v85
    case "":
        for fi86 := 0; fi86 < 6; fi86++ {
            v87 := []int{128, (fi86 ^ v14), v14}
            if len(v87) >= 2 {
                v87 = v87[1:len(v87)]
            }
            for fi88 := 0; fi88 < 7; fi88++ {
                _ = len(v87)
                if (!true) {
                    if false {
                        fmt.Println(v2, fi88)
                        if true {
                            v89 := GenS4{F0: 32767, F1: "foo bar"}
                            _ = v89
                        }
                        v90 := GenS1{F0: (-(fi86)), F1: ((v4 & genFn0((v1 << ((123) & 63)))) < (^(v4)))}
                        _ = v90
                    } else {
                        if len(v87) > 0 {
                            v87[0] = fi88
                        }
                        var v91 uint = uint(v4)
                        _ = v91
                    }
                    var v92 uint8 = 0
                    _ = v92
                }
                switch fi86 {
                case -1:
                    v4 = (v4 % ((v4) | 1))
                case 32768:
                    if len(v87) >= 2 {
                        v87 = v87[1:len(v87)]
                    }
                default:
                    if len(v87) > 0 {
                        v87[0] = v14
                    }
                }
                _ = fi88
            }
            _ = fi86
            _ = v87
        }
    }
    fmt.Println(v4, v3, v4)
    if ("" < ("foo bar" + "")) {
        var v93 uint32 = (v4 | v4)
        v94 := GenS9{F0: (v1 % ((v1) | 1))}
        for fi95 := 0; fi95 < 2; fi95++ {
            v96 := []int{}
            v97 := GenS1{F0: 385, F1: false}
            _ = fi95
            _ = v96
            _ = v97
        }
        _ = v93
        _ = v94
    }
    v2 = (v2 - v2)
    if (-1 >= v1) {
        if true {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v98 := GenS1{F0: 127, F1: (uint(v3) >= 65535)}
                _ = v98
            case "hello":
                for fi99 := 0; fi99 < 11; fi99++ {
                    for fi100 := 0; fi100 < 1; fi100++ {
                        var v101 int16 = (v2 + (-(v2)))
                        _ = fi100
                        _ = v101
                    }
                    var v102 float32 = (((0.1 - (1.5 * (0.0 / 1.0))) + 1.5) / (1.5 + (0.0 / 1e10)))
                    _ = fi99
                    _ = v102
                }
            default:
                if false {
                    v3 += v3
                    for fi103 := 0; fi103 < 20; fi103++ {
                        v2 -= (v2 ^ (v2 & v2))
                        v4 = 65535
                        _ = fi103
                    }
                    v104 := []int{v14}
                    _ = v104
                } else {
                    v105 := []int{}
                    var v106 string = ("hello" + "hello")
                    _ = v105
                    _ = v106
                }
            }
        } else {
            switch v14 {
            case 1:
                for fi107 := 0; fi107 < 16; fi107++ {
                    for fi108 := 0; fi108 < 17; fi108++ {
                        var v109 uint8 = 255
                        v110 := GenS10{F0: 1, F1: (v1 * (v1 & v1)), F2: 797}
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                    v111 := GenS8{F0: (v1 >> ((103) & 63)), F1: genFn0(v1)}
                    _ = fi107
                    _ = v111
                }
            default:
                v14 -= (v14 << ((596) & 63))
            }
            if (2 > 69) {
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v3)
                case "foo bar":
                    for fi112 := 0; fi112 < 2; fi112++ {
                        var v113 int16 = int16(v1)
                        switch (fi112 + v14) {
                        case 9223372036854775807:
                            for fi114 := 0; fi114 < 3; fi114++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v3 += 2
                                default:
                                    v115 := GenS8{F0: 9223372036854775807, F1: v4}
                                    _ = v115
                                }
                                var v116 int16 = 32767
                                if (!false) {
                                    v117 := []int{(v14 * 9223372036854775807), (fi114 << ((42) & 63)), (fi112 ^ 367)}
                                    var v118 int8 = -59
                                    _ = v117
                                    _ = v118
                                }
                                _ = fi114
                                _ = v116
                            }
                        case -513:
                            switch (fi112 << ((2) & 63)) {
                            case 0:
                                v14 += v14
                            case 4294967295:
                                var v119 int = v14
                                _ = v119
                            case 795:
                                var v120 float32 = ((1e10 - ((0.0 - 3.14159) - (0.0 * 0.0))) / 1e10)
                                _ = v120
                            }
                        case -2147483648:
                            if ((^(v2)) >= (v2 * -721)) {
                                v113 = v2
                            } else {
                                var v121 uint16 = 32767
                                for fi122 := 0; fi122 < 3; fi122++ {
                                    fmt.Println(fi122)
                                    var v123 int16 = -1
                                    if ((3.14159 / float32(-2.25)) != (1e10 - (-2.25 * (0.1 + 1.5)))) {
                                        if (!(!true)) {
                                            if (!true) {
                                                var v124 int8 = 100
                                                v113 = int16(v3)
                                                _ = v124
                                            }
                                            var v125 int8 = 100
                                            v123 = v2
                                            _ = v125
                                        }
                                        v123 += (v2 & v2)
                                    } else {
                                        v126 := GenS6{F0: "\n\t"}
                                        var v127 uint8 = 42
                                        _ = v126
                                        _ = v127
                                    }
                                    _ = fi122
                                    _ = v123
                                }
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        fmt.Println(v121)
                                        fmt.Println(v113)
                                    } else {
                                        for fi128 := 0; fi128 < 15; fi128++ {
                                            if true {
                                                switch (v14 - (fi112 << ((65535) & 63))) {
                                                case 0:
                                                    v3 = (^(v3))
                                                case 32768:
                                                    if (fi112 == (fi128 << ((uint(v121)) & 63))) {
                                                        v4 = 14
                                                        v129 := []int{32767, (fi128 & fi128), -9223372036854775808}
                                                        v130 := GenS8{F0: (v1 | v1), F1: v4}
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                default:
                                                    v2 = (v113 + -67)
                                                }
                                            }
                                            fmt.Println(fi128, v2, v4)
                                            _ = fi128
                                        }
                                        var v131 uint32 = (v4 & v4)
                                        _ = v131
                                    }
                                default:
                                    if false {
                                        fmt.Println(v4, v4, v113)
                                    }
                                }
                                _ = v121
                            }
                        default:
                            if (!(v4 == v4)) {
                                v1 = (v1 ^ -723)
                                v3 -= v3
                                v4 += (v4 / (((^(v4))) | 1))
                            } else {
                                v132 := []int{-84, 2147483647, (v14 | (fi112 / ((4294967295) | 1)))}
                                if len(v132) > 0 {
                                    v132[0] = (fi112 / ((v14) | 1))
                                }
                                _ = v132
                            }
                        }
                        _ = fi112
                        _ = v113
                    }
                case "hello":
                    v133 := []int{v14, 32767}
                    _ = v133
                default:
                    var v134 uint64 = (v3 / ((v3) | 1))
                    _ = v134
                }
                for fi135 := 0; fi135 < 14; fi135++ {
                    v136 := GenS11{F0: 121, F1: 2, F2: 2}
                    _ = fi135
                    _ = v136
                }
                v137 := GenS8{F0: (v1 - 193), F1: (v4 & (v4 + v4))}
                _ = v137
            }
            var v138 uint = 462
            _ = v138
        }
        v139 := GenS12{F0: (^(v1)), F1: "hello"}
        _ = v139
    }
    v2 = (v2 * v2)
    v140 := GenS13{F0: (^(v4)), F1: 18446744073709551615, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v140
    _ = v2
    _ = v3
    _ = v4
}
