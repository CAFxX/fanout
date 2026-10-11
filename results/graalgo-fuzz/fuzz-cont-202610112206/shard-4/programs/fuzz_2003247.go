package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int; F1 string; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int16; F2 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint; F1 uint; F2 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool; F1 uint32 }
func genFn0(p0 int16) int64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 int = 1
    _ = p0
    _ = v6
    return -688
}
func genFn1() uint64 {
    var v7 bool = false
    v8 := GenS1{F0: 83}
    v9 := []int{-658, -785}
    v10 := GenS1{F0: 127}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 bool = true
    var v2 int64 = 0
    v3 := make(map[string]int)
    var v4 int8 = 0
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 = v4
        var v12 float32 = 1e10
        v4 += v4
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 128
    for fi14 := 0; fi14 < 10; fi14++ {
        if (true && false) {
            if (!((0.1 + float64(0.0)) != -2.25)) {
                v15 := []int{}
                if v1 {
                    _ = len(v15)
                    v16 := GenS1{F0: (v13 - (v13 / ((v13) | 1)))}
                    _ = v16
                } else {
                    v3[""] = int(v4)
                }
                fmt.Println(v1, v2, v4)
                _ = v15
            }
            v1 = (!(499 == 65535))
            if ((0 < (v2 - 397)) || false) {
                v1 = (!(!false))
                v2 += genFn0(-882)
                v17 := v3["foo bar"]
                _ = v17
            }
        }
        _ = fi14
    }
    if (v1 && (("\n\t" + "a") >= (("a" + ("\n\t" + "a")) + ("hello" + "")))) {
        if (!((!(0 > 48)) && (128 <= (^(v13))))) {
            if false {
                v1 = (2147483648 == 681)
                v18 := GenS1{F0: (v13 << ((42) & 63))}
                _ = v18
            } else {
                switch -390 {
                case 32767:
                    _ = len(v3)
                case 4294967295:
                    var v19 uint32 = 65535
                    _ = v19
                default:
                    if true {
                        var v20 uint32 = 2147483647
                        _ = v20
                    }
                }
            }
            var v21 float32 = ((-2.25 * (1e10 * float32(float64(-2.25)))) - 1e10)
            _ = v21
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                switch fi22 {
                case 4294967295:
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        fmt.Println(v1, v13)
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            _ = len(v3)
                        }
                    case "a":
                        var v23 string = (("a" + "foo bar") + "")
                        _ = v23
                    }
                case 32768:
                    var v24 float32 = 1e10
                    _ = v24
                case 188:
                    fmt.Println(fi22)
                default:
                    fmt.Println(v1, v13, v2)
                }
                delete(v3, "")
                _ = fi22
            }
            fmt.Println(v4)
        }
        if v1 {
            var v25 float32 = 0.1
            var v26 uint16 = 32767
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS1{F0: v13}
                _ = fi27
                _ = v28
            }
            _ = v25
            _ = v26
        } else {
            v29 := GenS2{F0: (^(v2)), F1: 2147483647, F2: -1}
            var v30 uint = 0
            _ = v29
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 20; fi31++ {
            v13 = (v13 - 42)
            var v32 string = (("hello" + "") + ((("a" + (("foo bar" + "") + "\n\t")) + "foo bar") + "\n\t"))
            _ = fi31
            _ = v32
        }
    }
    var v33 int = 659
    v1 = true
    var v34 uint64 = 970
    v35 := GenS3{F0: 708, F1: v2}
    fmt.Println(v34, v33, v2)
    if (float64((0.0 - (float32(1e10) * float32(1e10)))) <= -2.25) {
        fmt.Println(v1)
        _ = len(v3)
        switch 32767 {
        case 775:
            v36 := GenS4{F0: 2147483647, F1: 2, F2: (v4 & (^(v4)))}
            _ = v36
        case 128:
            var v37 uint = 158
            _ = v37
        case 32768:
            fmt.Println(v13, v2)
        }
    }
    if v1 {
        if v1 {
            switch "a" {
            case "foo bar":
                if true {
                    v34 -= v34
                }
            default:
                for fi38 := 0; fi38 < 7; fi38++ {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        if (!v1) {
                            v33 += (fi38 << ((1) & 63))
                            switch (^(fi39)) {
                            case 1:
                                v40 := []int{(v33 | v33), -2147483648, (v33 ^ 32768), (fi38 - -1)}
                                _ = v40
                            }
                            v41 := GenS5{F0: ("foo bar" + ""), F1: 2147483647}
                            _ = v41
                        }
                        var v42 bool = (v33 != v33)
                        fmt.Println(v42)
                        switch "hello" {
                        case "foo bar":
                            var v43 bool = ((false || false) && (int16(v33) >= 0))
                            var v44 uint64 = genFn1()
                            v34 = (v44 + 2)
                            _ = v43
                            _ = v44
                        case "":
                            v4 -= 35
                        }
                        _ = fi39
                        _ = v42
                    }
                    if (v1 && (4294967295 != genFn1())) {
                        _ = len(v3)
                        if v1 {
                            var v45 float64 = 0.0
                            switch ("a" + "a") {
                            case "":
                                var v46 string = ("a" + ("foo bar" + ""))
                                _ = v46
                            case "foo bar":
                                fmt.Println(v1)
                            case "\n\t":
                                if v1 {
                                    v47 := GenS4{F0: 2147483647, F1: 53, F2: v4}
                                    var v48 int8 = 42
                                    switch "foo bar" {
                                    case "hello":
                                        switch v33 {
                                        case 2147483647:
                                            var v49 float64 = (v45 - (v45 - (0.1 * (v45 - 1e10))))
                                            _ = v49
                                        case 32767:
                                            v50 := []int{-378}
                                            _ = v50
                                        default:
                                            for fi51 := 0; fi51 < 4; fi51++ {
                                                var v52 uint8 = v13
                                                v52 += v13
                                                _ = fi51
                                                _ = v52
                                            }
                                        }
                                    }
                                    _ = v47
                                    _ = v48
                                }
                            }
                            v2 = 42
                            _ = v45
                        } else {
                            if (true || v1) {
                                var v53 uint64 = v34
                                _ = v53
                            }
                        }
                        for fi54 := 0; fi54 < 15; fi54++ {
                            if ((1e10 + 0.0) != (0.0 * 1e10)) {
                                v34 = 18446744073709551615
                                v13 = (v13 & v13)
                                var v55 uint32 = 1
                                _ = v55
                            }
                            if false {
                                if v1 {
                                    v13 -= v13
                                    v56 := []int{}
                                    _ = v56
                                } else {
                                    v57 := []int{int(v4), fi38, (v33 >> ((42) & 63)), (fi38 - (fi38 / (((fi54 ^ -2147483648)) | 1)))}
                                    v58 := GenS6{F0: (v13 | (v13 % ((42) | 1))), F1: 0}
                                    _ = v57
                                    _ = v58
                                }
                                if true {
                                    delete(v3, "")
                                } else {
                                    fmt.Println(fi54, v33, v34)
                                    fmt.Println(v13)
                                    v59 := GenS7{F0: v4, F1: (^(v13))}
                                    _ = v59
                                }
                            }
                            if v1 {
                                fmt.Println(v1)
                                v60 := GenS4{F0: 42, F1: (v13 % (((v13 % (((v13 & (v13 + v13))) | 1))) | 1)), F2: 127}
                                var v61 int = (fi38 - v33)
                                var v62 uint8 = 127
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            _ = fi54
                        }
                    }
                    if false {
                        v4 += v4
                        if v1 {
                            fmt.Println(fi38)
                            v13 = (v13 | 0)
                        } else {
                            switch "" {
                            case "hello":
                                if v1 {
                                    var v63 int32 = 42
                                    _ = v63
                                }
                            case "":
                                for fi64 := 0; fi64 < 15; fi64++ {
                                    fmt.Println(fi38, fi64)
                                    v65 := GenS6{F0: (v13 + (v13 * 2)), F1: 352}
                                    _ = fi64
                                    _ = v65
                                }
                            }
                            switch 168 {
                            case -2147483648:
                                for fi66 := 0; fi66 < 2; fi66++ {
                                    delete(v3, "")
                                    _ = fi66
                                }
                            case 42:
                                if v1 {
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        v4 -= v4
                                        v1 = v1
                                        _ = fi67
                                    }
                                    v34 += genFn1()
                                    var v68 int16 = -1
                                    _ = v68
                                }
                            default:
                                v69 := GenS7{F0: 0, F1: v13}
                                _ = v69
                            }
                            fmt.Println(v13, v34)
                        }
                        _ = len(v3)
                    }
                    _ = fi38
                }
            }
            v70 := []int{v33, (v33 & 127), (v33 * -647)}
            _ = v70
        } else {
            v71 := GenS8{F0: (v33 + v33), F1: "\n\t", F2: 1}
            switch "\n\t" {
            case "a":
                v13 = (v13 + v13)
            case "hello":
                if (!true) {
                    var v72 bool = v1
                    v2 = (-(v2))
                    _ = v72
                }
            case "foo bar":
                if true {
                    v4 = v4
                    if ((v34 + v34) > genFn1()) {
                        v13 -= 0
                    } else {
                        v73 := v3["\n\t"]
                        _ = v73
                    }
                } else {
                    for fi74 := 0; fi74 < 10; fi74++ {
                        fmt.Println(fi74, v13)
                        v4 -= 58
                        v75 := GenS6{F0: 127, F1: 197}
                        _ = fi74
                        _ = v75
                    }
                    v76 := []int{int(v4)}
                    for fi77 := 0; fi77 < 4; fi77++ {
                        v3[("hello" + "hello")] = (v33 / (((fi77 - fi77)) | 1))
                        for fi78 := 0; fi78 < 12; fi78++ {
                            if (v1 && ((431 < -122) && (65535 < 2147483648))) {
                                switch (fi78 & (v33 ^ 1)) {
                                case 0:
                                    if (!(!false)) {
                                        v1 = v1
                                        v34 = v34
                                    }
                                case 2147483647:
                                    v1 = ((true || (("hello" + "") <= "a")) && v1)
                                case 9223372036854775807:
                                    v79 := GenS9{F0: genFn0(1)}
                                    _ = v79
                                }
                            } else {
                                fmt.Println(v33)
                                if (!true) {
                                    var v80 float32 = 0.0
                                    v33 += (v33 + 32768)
                                    v81 := GenS7{F0: 127, F1: uint8(v34)}
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 string = "a"
                                    _ = v82
                                }
                                if len(v76) > 0 {
                                    v76[0] = v33
                                }
                                var v83 int32 = 0
                                var v84 uint8 = (^(v13))
                                _ = v83
                                _ = v84
                            }
                            switch ("a" + "") {
                            case "hello":
                                if (v1 || v1) {
                                    v85 := GenS3{F0: -1, F1: v2}
                                    var v86 uint = 65535
                                    _ = v85
                                    _ = v86
                                }
                            default:
                                switch (-(fi78)) {
                                case -766:
                                    v34 += v34
                                    var v87 float64 = float64((-2.25 - (((-2.25 / 1e10) / 1.5) + 1.5)))
                                    _ = v87
                                }
                            }
                            var v88 bool = (-1 > int32(v2))
                            _ = fi78
                            _ = v88
                        }
                        _ = fi77
                    }
                    _ = v76
                }
            }
            var v89 int64 = (v2 + genFn0(1000))
            _ = v71
            _ = v89
        }
        v1 = true
    } else {
        if (v1 && v1) {
            var v90 bool = true
            var v91 uint8 = v13
            v92 := GenS10{F0: 886, F1: -1, F2: 0}
            _ = v90
            _ = v91
            _ = v92
        } else {
            v33 = (v33 << ((2) & 63))
            var v93 uint16 = 65535
            _ = v93
        }
        var v94 float32 = float32(((float64(((3.14159 / 1.5) - (-2.25 * -2.25))) + 0.0) - 0.1))
        for fi95 := 0; fi95 < 10; fi95++ {
            v96 := []int{-358, (fi95 | (fi95 << ((148) & 63))), (fi95 >> ((1) & 63))}
            v97 := GenS11{F0: v34}
            _ = fi95
            _ = v96
            _ = v97
        }
        _ = v94
    }
    if v1 {
        fmt.Println(v2, v2, v2)
        switch (^(v33)) {
        case 0:
            v1 = false
        case 2147483647:
            var v98 int32 = 0
            _ = v98
        case -9223372036854775808:
            v33 = (v33 + 127)
        }
        fmt.Println(v4, v1, v34)
    }
    var v99 int8 = (v4 & (v4 << ((215) & 63)))
    fmt.Println(v34, v1, v99)
    if v1 {
        v100 := []int{(v33 * (v33 & (v33 | v33)))}
        switch int(v99) {
        case 0:
            switch "\n\t" {
            case "":
                var v101 int = 128
                _ = v101
            case "hello":
                v102 := v3[("" + "\n\t")]
                _ = v102
            case "a":
                v103 := GenS12{F0: 634, F1: 65535, F2: ("" + ("" + "a"))}
                _ = v103
            }
        default:
            for fi104 := 0; fi104 < 4; fi104++ {
                v100 = append(v100, -372)
                var v105 uint8 = (v13 / ((1) | 1))
                _ = fi104
                _ = v105
            }
        }
        _ = v100
    }
    v99 = (v4 >> ((2) & 63))
    if (1.5 > 0.1) {
        v34 = (v34 & (v34 + 42))
    } else {
        var v106 uint32 = 42
        switch (v33 >> ((255) & 63)) {
        case 98:
            switch ("a" + "a") {
            case "":
                for fi107 := 0; fi107 < 10; fi107++ {
                    v34 += 18446744073709551615
                    for fi108 := 0; fi108 < 15; fi108++ {
                        _ = len(v3)
                        var v109 uint64 = (v34 >> ((1) & 63))
                        if (8 >= 774) {
                            v110 := []int{}
                            _ = v110
                        } else {
                            switch fi107 {
                            case -9223372036854775808:
                                v1 = v1
                            case -2147483648:
                                switch (^(fi108)) {
                                case -462:
                                    var v111 float64 = 1.5
                                    _ = v111
                                case 127:
                                    v112 := []int{(fi107 & (v33 ^ v33))}
                                    var v113 int = 731
                                    var v114 uint16 = 1
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                case -1:
                                    var v115 uint64 = (v109 << ((42) & 63))
                                    _ = v115
                                default:
                                    v116 := GenS9{F0: (-(v2))}
                                    switch "foo bar" {
                                    case "":
                                        var v117 uint16 = 1
                                        _ = v117
                                    case "\n\t":
                                        if (42 < int16(v34)) {
                                            v99 -= 100
                                            v118 := []int{}
                                            _ = v118
                                        } else {
                                            switch "a" {
                                            case "":
                                                v119 := []int{}
                                                var v120 uint16 = 0
                                                if true {
                                                    switch ((("hello" + "foo bar") + "foo bar") + "\n\t") {
                                                    case "hello":
                                                        if ((v106 + v106) != (v106 % ((v106) | 1))) {
                                                            var v121 uint = 255
                                                            _ = v121
                                                        } else {
                                                            if (false || true) {
                                                                var v122 int64 = (v2 * (-(v2)))
                                                                v123 := GenS13{F0: (v13 - v13)}
                                                                if len(v119) > 0 {
                                                                    v119[0] = (v33 | (fi108 >> ((521) & 63)))
                                                                }
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                            fmt.Println(v4)
                                                        }
                                                    case "\n\t":
                                                        v106 -= (v106 | (v106 + 2147483648))
                                                    default:
                                                        var v124 uint64 = (v34 - v34)
                                                        _ = v124
                                                    }
                                                    v119 = append(v119, (fi107 & fi108))
                                                    fmt.Println(v2)
                                                }
                                                _ = v119
                                                _ = v120
                                            case "\n\t":
                                                switch (("a" + ("a" + "")) + "foo bar") {
                                                case "hello":
                                                    switch "a" {
                                                    case "hello":
                                                        var v125 float64 = (1e10 + (1e10 - 1e10))
                                                        _ = v125
                                                    }
                                                    v126 := GenS7{F0: 100, F1: v13}
                                                    _ = v126
                                                case "\n\t":
                                                    var v127 uint64 = v109
                                                    _ = v127
                                                case "a":
                                                    v34 = (v109 >> ((0) & 63))
                                                }
                                            }
                                            var v128 int16 = 0
                                            if ((^(v128)) != v128) {
                                                switch 127 {
                                                case 586:
                                                    if false {
                                                        v3[""] = v33
                                                        v129 := v3["\n\t"]
                                                        if (v1 && v1) {
                                                            v3["hello"] = (v33 * (fi108 / ((-9223372036854775808) | 1)))
                                                            v13 = (v13 >> ((810) & 63))
                                                        }
                                                        _ = v129
                                                    }
                                                case 2147483647:
                                                    _ = len(v3)
                                                }
                                            }
                                            _ = v128
                                        }
                                    default:
                                        switch "hello" {
                                        case "":
                                            v130 := GenS11{F0: genFn1()}
                                            _ = v130
                                        case "hello":
                                            switch fi107 {
                                            case -1:
                                                fmt.Println(v34)
                                            case 127:
                                                if v1 {
                                                    switch fi107 {
                                                    case 32767:
                                                        var v131 uint32 = v106
                                                        _ = v131
                                                    case 4294967295:
                                                        v3[""] = 32767
                                                    case 2147483647:
                                                        v3["a"] = (fi108 + 32768)
                                                    default:
                                                        var v132 int64 = (v2 % ((252) | 1))
                                                        _ = v132
                                                    }
                                                }
                                            default:
                                                v1 = (true || v1)
                                            }
                                        case "\n\t":
                                            if true {
                                                var v133 float32 = -2.25
                                                v134 := []int{1}
                                                _ = v133
                                                _ = v134
                                            }
                                            var v135 uint8 = (v13 & 0)
                                            v136 := GenS9{F0: (v2 & (v2 + genFn0(169)))}
                                            _ = v135
                                            _ = v136
                                        default:
                                            fmt.Println(v1)
                                        }
                                    }
                                    delete(v3, "\n\t")
                                    _ = v116
                                }
                            case -1:
                                v137 := GenS14{F0: 42}
                                _ = v137
                            }
                            if (v99 != 127) {
                                switch ((("" + "foo bar") + "\n\t") + "hello") {
                                case "hello":
                                    if (-1 != 517) {
                                        v2 = -1
                                        switch fi108 {
                                        case -9223372036854775808:
                                            fmt.Println(fi108, v34, v1)
                                        case -460:
                                            if (((fi107 >> ((18446744073709551615) & 63)) > fi107) && v1) {
                                                fmt.Println(fi107)
                                                if (v1 && true) {
                                                    var v138 bool = true
                                                    switch v33 {
                                                    case -2147483648:
                                                        v3[""] = fi108
                                                    }
                                                    _ = v138
                                                }
                                            }
                                        case 480:
                                            v139 := GenS12{F0: 42, F1: 18446744073709551615, F2: "hello"}
                                            _ = v139
                                        }
                                    } else {
                                        v140 := []int{(fi107 >> ((370) & 63)), (v33 >> ((255) & 63)), (fi108 ^ fi107), -2147483648}
                                        var v141 uint = 523
                                        v34 = uint64(v106)
                                        _ = v140
                                        _ = v141
                                    }
                                case "foo bar":
                                    delete(v3, "")
                                }
                                switch "a" {
                                case "a":
                                    var v142 uint = 42
                                    _ = v142
                                }
                                var v143 int8 = v99
                                _ = v143
                            } else {
                                switch (v33 * (v33 ^ fi108)) {
                                case 0:
                                    v144 := GenS15{F0: 1}
                                    _ = v144
                                default:
                                    v145 := GenS14{F0: 0}
                                    _ = v145
                                }
                                var v146 bool = (42 > 0)
                                if (255 == 2) {
                                    v147 := GenS16{F0: (!(!((v106 + (v106 << ((uint(v106)) & 63))) <= v106))), F1: (^(v106))}
                                    v106 -= 528
                                    _ = v147
                                }
                                _ = v146
                            }
                        }
                        _ = fi108
                        _ = v109
                    }
                    _ = fi107
                }
            case "a":
                v148 := GenS14{F0: 65535}
                _ = v148
            case "hello":
                v149 := v3["a"]
                _ = v149
            default:
                fmt.Println(v2, v33, v13)
            }
        }
        _ = v106
    }
    v150 := GenS16{F0: true, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v13
    _ = v150
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v99
}
