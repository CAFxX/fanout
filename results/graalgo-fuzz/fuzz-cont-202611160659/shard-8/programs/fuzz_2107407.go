package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool; F1 uint8; F2 int32 }

func main() {
    var v1 bool = (-387 <= 32768)
    var v2 uint8 = 2
    var v3 int8 = 100
    var v4 uint16 = 0
    var v5 uint64 = 4294967295
    v6 := GenS1{F0: -938}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2)
    case "foo bar":
        var v7 uint64 = 30
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v3)
            var v10 int = 32767
            if (v1 && true) {
                var v11 int32 = 498
                var v12 uint16 = 1
                _ = v11
                _ = v12
            } else {
                var v13 uint32 = 1
                var v14 int64 = 42
                switch int(v14) {
                case 273:
                    if true {
                        switch "" {
                        case "foo bar":
                            for fi15 := 0; fi15 < 3; fi15++ {
                                switch "\n\t" {
                                case "a":
                                    v4 = v4
                                }
                                v13 = (v13 * uint32(v4))
                                v5 = v5
                                _ = fi15
                            }
                        case "\n\t":
                            if v1 {
                                if (false && true) {
                                    switch fi9 {
                                    case 32768:
                                        v16 := GenS1{F0: 1}
                                        _ = v16
                                    case 2147483647:
                                        switch -857 {
                                        case 127:
                                            var v17 uint64 = 0
                                            _ = v17
                                        case 995:
                                            if (false && v1) {
                                                v13 = (^(v13))
                                                var v18 bool = ((v13 / (((v13 % ((v13) | 1))) | 1)) == (v13 >> ((0) & 63)))
                                                _ = v18
                                            }
                                        default:
                                            var v19 uint = 255
                                            _ = v19
                                        }
                                    default:
                                        if v1 {
                                            var v20 uint16 = v4
                                            switch v10 {
                                            case 32768:
                                                if (("" + "a") == "foo bar") {
                                                    var v21 uint16 = v4
                                                    _ = v21
                                                }
                                            case -9223372036854775808:
                                                var v22 int64 = 9223372036854775807
                                                _ = v22
                                            }
                                            _ = v20
                                        } else {
                                            v23 := GenS1{F0: -69}
                                            var v24 string = ""
                                            if (!true) {
                                                v25 := []int{}
                                                fmt.Println(v10, v5)
                                                _ = v25
                                            }
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                } else {
                                    fmt.Println(v13, v13, v14)
                                    var v26 float32 = 0.0
                                    for fi27 := 0; fi27 < 13; fi27++ {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v28 int32 = -1
                                            _ = v28
                                        case "hello":
                                            var v29 float32 = 1e10
                                            _ = v29
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "foo bar":
                                                v30 := GenS1{F0: -123}
                                                _ = v30
                                            }
                                        default:
                                            v31 := GenS2{F0: 510}
                                            v13 -= v13
                                            _ = v31
                                        }
                                        var v32 int64 = v14
                                        if false {
                                            var v33 uint = 4294967295
                                            _ = v33
                                        }
                                        _ = fi27
                                        _ = v32
                                    }
                                    _ = v26
                                }
                            } else {
                                if ((^(v10)) < (v10 / ((v10) | 1))) {
                                    var v34 uint32 = (v13 | 926)
                                    _ = v34
                                }
                                if false {
                                    var v35 float64 = float64((0.0 - 0.0))
                                    fmt.Println(v1)
                                    _ = v35
                                }
                                var v36 int8 = (v3 ^ (-(v3)))
                                _ = v36
                            }
                        }
                    } else {
                        var v37 int16 = 1
                        for fi38 := 0; fi38 < 18; fi38++ {
                            v10 = 725
                            v1 = true
                            var v39 uint = 700
                            _ = fi38
                            _ = v39
                        }
                        if v1 {
                            var v40 uint = 1
                            _ = v40
                        }
                        _ = v37
                    }
                }
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        for fi41 := 0; fi41 < 17; fi41++ {
            v3 += v3
            _ = fi41
        }
        fmt.Println(v3)
    } else {
        switch 32767 {
        case 0:
            fmt.Println(v5)
        case 1:
            if false {
                v3 = v3
            } else {
                var v42 int32 = -2147483648
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 int8 = v3
                    v1 = true
                    var v45 int8 = (^(v3))
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 15; fi46++ {
                    v47 := []int{(fi46 + -9223372036854775808), (fi46 * (fi46 | -1))}
                    if len(v47) > 0 {
                        v47[0] = (^(fi46))
                    }
                    _ = fi46
                    _ = v47
                }
                _ = v42
            }
        default:
            v5 = (v5 ^ v5)
        }
        v48 := GenS1{F0: -1}
        if v1 {
            if (v1 || (false && (!v1))) {
                for fi49 := 0; fi49 < 20; fi49++ {
                    v50 := GenS1{F0: 581}
                    _ = fi49
                    _ = v50
                }
                switch 9223372036854775807 {
                case 713:
                    switch "hello" {
                    case "\n\t":
                        v51 := GenS2{F0: (^(v4))}
                        _ = v51
                    case "foo bar":
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v53 := GenS2{F0: 2}
                            v4 -= (v4 % ((v4) | 1))
                            v54 := GenS3{F0: ("foo bar" + "hello")}
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    }
                }
            } else {
                v55 := GenS4{F0: 2, F1: (^(v5)), F2: (v3 % (((v3 & 0)) | 1))}
                switch "a" {
                case "hello":
                    for fi56 := 0; fi56 < 13; fi56++ {
                        for fi57 := 0; fi57 < 16; fi57++ {
                            v58 := []int{-9223372036854775808, 60, -9223372036854775808}
                            v2 -= v2
                            _ = fi57
                            _ = v58
                        }
                        _ = fi56
                    }
                default:
                    if v1 {
                        if v1 {
                            if (!(false || v1)) {
                                var v59 uint32 = 65535
                                v60 := []int{}
                                var v61 int8 = (-(v3))
                                _ = v59
                                _ = v60
                                _ = v61
                            } else {
                                if v1 {
                                    v5 += 0
                                    for fi62 := 0; fi62 < 16; fi62++ {
                                        v5 -= v5
                                        v63 := GenS5{F0: v3}
                                        _ = fi62
                                        _ = v63
                                    }
                                    v64 := GenS4{F0: (^(v4)), F1: v5, F2: v3}
                                    _ = v64
                                } else {
                                    if (((v1 && ((v2 & (v2 - (^(v2)))) == v2)) && true) && true) {
                                        if true {
                                            switch ("foo bar" + "\n\t") {
                                            case "\n\t":
                                                v65 := []int{4294967295, -2147483648, -480, 42}
                                                _ = v65
                                            case "foo bar":
                                                var v66 float32 = 1e10
                                                _ = v66
                                            case "hello":
                                                v67 := GenS5{F0: v3}
                                                _ = v67
                                            default:
                                                if (42 < 0) {
                                                    v68 := []int{-9223372036854775808}
                                                    if (v1 || v1) {
                                                        v4 = v4
                                                    }
                                                    _ = v68
                                                } else {
                                                    v69 := []int{42, 474, 2147483647, 2147483647}
                                                    if v1 {
                                                        fmt.Println(v2, v2, v4)
                                                        v70 := GenS2{F0: v4}
                                                        v69 = append(v69, 127)
                                                        _ = v70
                                                    } else {
                                                        if (!v1) {
                                                            var v71 int8 = v3
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v69
                                                }
                                            }
                                        }
                                        v72 := GenS5{F0: v3}
                                        _ = v72
                                    } else {
                                        if false {
                                            var v73 int32 = 2147483647
                                            v73 -= 845
                                            var v74 uint64 = v5
                                            _ = v73
                                            _ = v74
                                        }
                                        var v75 int = 42
                                        _ = v75
                                    }
                                }
                                v1 = (0.1 > float32((0.1 + 0.1)))
                                v4 = v4
                            }
                            for fi76 := 0; fi76 < 20; fi76++ {
                                if ((-747 != -32768) && true) {
                                    switch fi76 {
                                    case 4294967295:
                                        for fi77 := 0; fi77 < 16; fi77++ {
                                            var v78 int = -9223372036854775808
                                            var v79 int64 = 9223372036854775807
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                    case -2147483648:
                                        var v80 float64 = 1.5
                                        _ = v80
                                    }
                                    fmt.Println(v3)
                                    if (true || v1) {
                                        v81 := []int{fi76}
                                        _ = v81
                                    } else {
                                        v82 := GenS6{F0: v3, F1: v2}
                                        _ = v82
                                    }
                                }
                                v1 = false
                                if v1 {
                                    var v83 int8 = 42
                                    var v84 uint64 = v5
                                    if (("a" + "") < "hello") {
                                        switch fi76 {
                                        case 128:
                                            fmt.Println(v1)
                                        case 0:
                                            var v85 float64 = (3.14159 + 1e10)
                                            _ = v85
                                        case 9223372036854775807:
                                            fmt.Println(v4, v83)
                                        default:
                                            if (0 == uint64(v4)) {
                                                for fi86 := 0; fi86 < 14; fi86++ {
                                                    var v87 int32 = 42
                                                    _ = fi86
                                                    _ = v87
                                                }
                                            }
                                        }
                                        if (0.1 <= -2.25) {
                                            v83 += int8(fi76)
                                            v88 := GenS7{F0: 0, F1: 0, F2: (v84 ^ v5)}
                                            switch 32768 {
                                            case 2147483647:
                                                switch ("\n\t" + "foo bar") {
                                                case "\n\t":
                                                    var v89 uint16 = (v4 & v4)
                                                    _ = v89
                                                }
                                            case 118:
                                                fmt.Println(v84, v4, v2)
                                            default:
                                                if v1 {
                                                    switch "a" {
                                                    case "\n\t":
                                                        if ((v2 << ((874) & 63)) == (v2 >> ((1) & 63))) {
                                                            for fi90 := 0; fi90 < 2; fi90++ {
                                                                v2 += (^(v2))
                                                                _ = fi90
                                                            }
                                                        } else {
                                                            for fi91 := 0; fi91 < 1; fi91++ {
                                                                if false {
                                                                    fmt.Println(fi76)
                                                                } else {
                                                                    v92 := GenS8{F0: 42, F1: 42}
                                                                    switch -2147483648 {
                                                                    case 127:
                                                                        fmt.Println(v1, v3, v1)
                                                                    case 1:
                                                                        fmt.Println(v1)
                                                                    }
                                                                    var v93 bool = false
                                                                    v94 := []int{}
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                                fmt.Println(v5)
                                                                var v95 float32 = (0.0 * (0.0 - 3.14159))
                                                                _ = fi91
                                                                _ = v95
                                                            }
                                                            v96 := GenS4{F0: v4, F1: v5, F2: (v83 ^ v3)}
                                                            v97 := []int{(fi76 | -446), fi76}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    case "a":
                                                        var v98 uint64 = (^(v84))
                                                        _ = v98
                                                    }
                                                } else {
                                                    var v99 int32 = 2147483647
                                                    for fi100 := 0; fi100 < 3; fi100++ {
                                                        switch ("\n\t" + "\n\t") {
                                                        case "\n\t":
                                                            var v101 float64 = 1.5
                                                            _ = v101
                                                        case "":
                                                            fmt.Println(v1, fi100)
                                                        default:
                                                            var v102 uint32 = 65535
                                                            _ = v102
                                                        }
                                                        _ = fi100
                                                    }
                                                    _ = v99
                                                }
                                            }
                                            _ = v88
                                        }
                                        if v1 {
                                            if (-1 == 9223372036854775807) {
                                                if v1 {
                                                    var v103 int8 = v3
                                                    if (1e10 >= 3.14159) {
                                                        switch (fi76 * (fi76 ^ (^(fi76)))) {
                                                        case 4294967295:
                                                            v1 = (v2 >= 5)
                                                        case 379:
                                                            var v104 uint16 = v4
                                                            _ = v104
                                                        case 465:
                                                            var v105 int8 = (v103 | v3)
                                                            _ = v105
                                                        }
                                                        switch "hello" {
                                                        case "":
                                                            var v106 string = ("" + ("" + ("" + (("foo bar" + (("a" + "") + ("foo bar" + "foo bar"))) + "foo bar"))))
                                                            _ = v106
                                                        case "hello":
                                                            v107 := GenS2{F0: uint16(v2)}
                                                            _ = v107
                                                        case "\n\t":
                                                            var v108 bool = v1
                                                            _ = v108
                                                        }
                                                    } else {
                                                        v109 := GenS9{F0: 2}
                                                        _ = v109
                                                    }
                                                    if v1 {
                                                        var v110 uint32 = 155
                                                        _ = v110
                                                    }
                                                    _ = v103
                                                }
                                                v111 := GenS10{F0: 2147483648, F1: v2}
                                                _ = v111
                                            } else {
                                                v2 = 128
                                                if (2 != uint(v5)) {
                                                    for fi112 := 0; fi112 < 10; fi112++ {
                                                        v113 := GenS9{F0: 512}
                                                        v2 += v2
                                                        v84 = (^(v5))
                                                        if ((true || (v1 && false)) && (((1.5 / -2.25) * 0.1) != 1.5)) {
                                                            if false {
                                                                v5 += (v5 * uint64(v4))
                                                                if false {
                                                                    fmt.Println(v1)
                                                                    v4 = v4
                                                                    var v114 int32 = -411
                                                                    var v115 uint32 = 42
                                                                    _ = v114
                                                                    _ = v115
                                                                } else {
                                                                    var v116 int16 = 0
                                                                    v117 := GenS8{F0: -5, F1: 2}
                                                                    if ("hello" <= ("foo bar" + (((("hello" + ("a" + "foo bar")) + ("foo bar" + "\n\t")) + "") + ("a" + "foo bar")))) {
                                                                        v4 -= (v4 | (v4 * (v4 << ((728) & 63))))
                                                                        if true {
                                                                            v84 -= (^(v84))
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v118 int16 = v116
                                                                                _ = v118
                                                                            case "":
                                                                                v119 := GenS11{F0: -9223372036854775808}
                                                                                _ = v119
                                                                            }
                                                                            v120 := []int{(fi112 & (fi112 & (-(fi76)))), (fi112 | (fi76 >> ((736) & 63))), (fi112 ^ (fi112 + (-(fi76))))}
                                                                            _ = v120
                                                                        } else {
                                                                            if true {
                                                                                v121 := GenS5{F0: (v83 << ((102) & 63))}
                                                                                v4 += (^(v4))
                                                                                var v122 int8 = 127
                                                                                _ = v121
                                                                                _ = v122
                                                                            }
                                                                            v123 := []int{fi112, fi112, fi112, 1}
                                                                            if v1 {
                                                                                v124 := GenS12{F0: (fi112 - 242), F1: v1}
                                                                                _ = v124
                                                                            }
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                            } else {
                                                                var v125 bool = v1
                                                                v126 := GenS11{F0: 4294967295}
                                                                v1 = (3.14159 < 1.5)
                                                                fmt.Println(v84, fi112, fi112)
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                        }
                                                        _ = fi112
                                                        _ = v113
                                                    }
                                                    var v127 bool = v1
                                                    _ = v127
                                                } else {
                                                    v128 := GenS13{F0: -1}
                                                    v5 = (^(v84))
                                                    _ = v128
                                                }
                                                fmt.Println(v4, fi76)
                                            }
                                            v129 := GenS14{F0: (v1 || (2 <= v2)), F1: (v2 >> ((255) & 63)), F2: 77}
                                            if (!(9223372036854775807 <= -1)) {
                                                if false {
                                                    v1 = (929 == uint32(v2))
                                                    switch ("hello" + ("\n\t" + ("foo bar" + "a"))) {
                                                    case "a":
                                                        var v130 uint16 = v4
                                                        _ = v130
                                                    default:
                                                        for fi131 := 0; fi131 < 12; fi131++ {
                                                            fmt.Println(v3)
                                                            var v132 uint16 = 255
                                                            v5 = 0
                                                            _ = fi131
                                                            _ = v132
                                                        }
                                                    }
                                                } else {
                                                    v133 := GenS14{F0: ((!v1) || (v4 < v4)), F1: (v2 << ((849) & 63)), F2: -1}
                                                    for fi134 := 0; fi134 < 18; fi134++ {
                                                        v3 = v83
                                                        v1 = v1
                                                        _ = fi134
                                                    }
                                                    _ = v133
                                                }
                                                var v135 string = ""
                                                switch 32767 {
                                                case 1:
                                                    v136 := []int{}
                                                    _ = v136
                                                case -9223372036854775808:
                                                    var v137 int16 = -32768
                                                    _ = v137
                                                default:
                                                    v138 := []int{-880, -824, fi76}
                                                    _ = v138
                                                }
                                                _ = v135
                                            }
                                            _ = v129
                                        } else {
                                            var v139 uint = 0
                                            v4 = (v4 << (((v139 - (v139 & v139))) & 63))
                                            v140 := []int{fi76}
                                            _ = v139
                                            _ = v140
                                        }
                                    }
                                    _ = v83
                                    _ = v84
                                }
                                _ = fi76
                            }
                            v141 := []int{-290, -884}
                            _ = v141
                        }
                        if (-2147483648 >= -2147483648) {
                            v142 := GenS7{F0: 0, F1: 2, F2: v5}
                            v143 := GenS6{F0: (v3 & (-(v3))), F1: v2}
                            v144 := []int{}
                            _ = v142
                            _ = v143
                            _ = v144
                        }
                    }
                }
                var v145 uint = 1
                _ = v145
                _ = v55
            }
        } else {
            var v146 uint32 = 170
            v147 := GenS8{F0: -822, F1: v146}
            var v148 string = "\n\t"
            _ = v146
            _ = v147
            _ = v148
        }
        _ = v48
    }
    v4 = v4
    v4 -= 279
    v1 = v1
    switch 774 {
    case -381:
        fmt.Println(v2, v3, v1)
    case 42:
        v149 := []int{-1, 4294967295, 0, -654}
        _ = v149
    default:
        v4 -= 115
    }
    fmt.Println(v5, v2, v2)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
