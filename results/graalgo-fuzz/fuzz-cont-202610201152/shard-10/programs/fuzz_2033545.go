package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int32; F1 string; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS16 struct { F0 uint32; F1 uint32 }
type GenS17 struct { F0 int16; F1 bool }
type GenS18 struct { F0 int; F1 bool }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int; F1 int64; F2 bool }
type GenS21 struct { F0 bool; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 uint64) bool {
    var v6 uint16 = 42
    p1 -= (p1 << ((2) & 63))
    var v7 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 != (p0 << ((255) & 63)))
}
func genFn1(p0 int, p1 int64) int {
    var v8 uint64 = 4294967295
    p0 = -497
    var v9 uint = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 0
    var v3 float32 = ((3.14159 / -2.25) + float32(-2.25))
    v4 := make(map[uint32]int)
    var v5 uint32 = v1
    switch 1 {
    case 824:
        v1 = v1
    default:
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    var v11 int8 = 42
    v2 += v2
    v3 = (3.14159 / 1.0)
    v11 -= 100
    if (((v2 + 9223372036854775807) <= v2) || (-156 == 32768)) {
        var v12 uint32 = (v5 ^ (v5 / ((42) | 1)))
        v13 := v4[v12]
        if true {
            var v14 uint16 = 65535
            var v15 uint8 = 128
            _ = v14
            _ = v15
        } else {
            if (-32768 < 505) {
                switch "foo bar" {
                case "hello":
                    v5 = v5
                case "a":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := GenS1{F0: 4294967295, F1: "foo bar"}
                        _ = fi16
                        _ = v17
                    }
                default:
                    v12 = (v12 % (((v1 * v12)) | 1))
                }
                var v18 uint8 = 119
                _ = v18
            }
        }
        _ = v12
        _ = v13
    }
    switch genFn1(genFn1(2147483647, (v2 & (v2 >> ((4294967295) & 63)))), (v2 % ((v2) | 1))) {
    case 9223372036854775807:
        switch genFn1(787, (v2 / ((9223372036854775807) | 1))) {
        case 42:
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint16 = 716
                for fi21 := 0; fi21 < 11; fi21++ {
                    delete(v4, v5)
                    if false {
                        v22 := GenS1{F0: (fi19 % ((127) | 1)), F1: "\n\t"}
                        var v23 uint8 = 42
                        v24 := GenS2{F0: 375}
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        switch (("\n\t" + "a") + ("a" + "a")) {
                        case "foo bar":
                            v25 := GenS2{F0: 42}
                            _ = v25
                        }
                        if false {
                            var v26 bool = false
                            var v27 int = fi21
                            v11 -= (^(v11))
                            _ = v26
                            _ = v27
                        }
                        v20 -= 1
                    }
                    v28 := GenS2{F0: -561}
                    _ = fi21
                    _ = v28
                }
                _ = fi19
                _ = v20
            }
        case -359:
            v5 -= v1
        default:
            v29 := GenS3{F0: 32768, F1: 1, F2: 965}
            _ = v29
        }
    default:
        _ = len(v4)
    }
    delete(v4, v5)
    for fi30 := 0; fi30 < 5; fi30++ {
        if false {
            v4[v1] = (fi30 * -2147483648)
        }
        _ = fi30
    }
    v2 = v2
    var v31 uint16 = 475
    if (864 <= 4294967295) {
        v3 += 0.1
    }
    if (!true) {
        v1 -= (^(v1))
    }
    if (!true) {
        if false {
            v3 -= v3
            for fi32 := 0; fi32 < 13; fi32++ {
                v33 := GenS2{F0: -1}
                switch "foo bar" {
                case "foo bar":
                    var v34 int16 = 996
                    _ = v34
                default:
                    v35 := GenS4{F0: 1, F1: 2}
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
        }
        v4[(v1 << ((309) & 63))] = -745
        switch ("hello" + "\n\t") {
        case "hello":
            v36 := []int{9223372036854775807}
            _ = v36
        }
    } else {
        v37 := GenS5{F0: 797, F1: ("a" + "\n\t"), F2: (^(v11))}
        for fi38 := 0; fi38 < 8; fi38++ {
            if true {
                v39 := []int{fi38, -2147483648, fi38}
                if (v11 != (-(v11))) {
                    if false {
                        for fi40 := 0; fi40 < 7; fi40++ {
                            switch fi38 {
                            case 128:
                                switch (fi40 >> ((4294967295) & 63)) {
                                case 103:
                                    fmt.Println(v3, v31)
                                case -2147483648:
                                    if (!(1.5 < ((1e10 / 1.0) * (0.1 - (-2.25 * 1e10))))) {
                                        if false {
                                            fmt.Println(v3, fi40)
                                            switch (fi40 & fi40) {
                                            case 815:
                                                _ = len(v4)
                                            case 128:
                                                v3 = v3
                                            }
                                        }
                                    } else {
                                        if (2 == 2) {
                                            if ((-32768 > 81) && false) {
                                                var v41 int32 = -732
                                                _ = v41
                                            } else {
                                                v2 = (v2 ^ -1)
                                                if (!true) {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "hello":
                                                        fmt.Println(fi38, fi40)
                                                    default:
                                                        fmt.Println(v2, v3)
                                                    }
                                                    if true {
                                                        if true {
                                                            v42 := GenS6{F0: 94, F1: (v5 & v1), F2: 127}
                                                            if ((4294967295 != v5) && false) {
                                                                if len(v39) >= 2 {
                                                                    v39 = v39[1:len(v39)]
                                                                }
                                                            }
                                                            _ = len(v4)
                                                            if (0 != 574) {
                                                                v39 = append(v39, fi38)
                                                                v1 -= 2
                                                                v2 = v2
                                                            } else {
                                                                v43 := GenS2{F0: 1}
                                                                var v44 uint8 = 42
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            _ = v42
                                                        }
                                                        _ = len(v39)
                                                        switch ("hello" + "") {
                                                        case "foo bar":
                                                            v45 := GenS7{F0: (788 <= 263), F1: 2}
                                                            _ = v45
                                                        case "a":
                                                            v3 = 1.5
                                                        case "hello":
                                                            if (!true) {
                                                                v46 := GenS8{F0: 18446744073709551615}
                                                                _ = v46
                                                            }
                                                        }
                                                        var v47 float64 = 3.14159
                                                        v47 += (v47 + v47)
                                                        v48 := GenS4{F0: 42, F1: 552}
                                                        v31 += v31
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        var v49 int16 = -32768
                                                        switch "a" {
                                                        case "\n\t":
                                                            if len(v39) >= 2 {
                                                                v39 = v39[1:len(v39)]
                                                            }
                                                        case "a":
                                                            v3 = v3
                                                        case "":
                                                            var v50 int32 = 2147483647
                                                            _ = v50
                                                        }
                                                        _ = v49
                                                    }
                                                    v51 := GenS6{F0: 42, F1: v1, F2: 195}
                                                    var v52 float64 = 1e10
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v53 := v4[65535]
                                                    var v54 int8 = v11
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v5 = (v1 + (v5 & v5))
                                                v3 = 0.0
                                            }
                                            fmt.Println(v5)
                                            v39 = append(v39, (fi38 ^ fi38))
                                            var v55 int8 = (v11 >> ((4294967295) & 63))
                                            _ = v55
                                        } else {
                                            switch "hello" {
                                            case "":
                                                switch fi38 {
                                                case 17:
                                                    if (-1 == 32767) {
                                                        if (!(false || false)) {
                                                            if true {
                                                                v56 := GenS3{F0: (fi38 * fi38), F1: (v31 * v31), F2: 1}
                                                                v31 += 32767
                                                                _ = v56
                                                            }
                                                            v57 := GenS9{F0: 65535, F1: fi40}
                                                            _ = v57
                                                        }
                                                        if len(v39) > 0 {
                                                            v39[0] = (fi40 - genFn1(fi38, (v2 & -792)))
                                                        }
                                                        if len(v39) >= 2 {
                                                            v39 = v39[1:len(v39)]
                                                        }
                                                    } else {
                                                        fmt.Println(v11)
                                                        fmt.Println(v5, fi38, v11)
                                                        switch 32768 {
                                                        case 32768:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                _ = len(v39)
                                                            }
                                                        }
                                                        if (!(((65535 != uint(v1)) || (4294967295 >= 4294967295)) || ((42 == 2) || false))) {
                                                            fmt.Println(v1, v5)
                                                            var v58 float32 = float32(3.14159)
                                                            switch 1 {
                                                            case 597:
                                                                var v59 int = (fi40 >> ((2) & 63))
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        } else {
                                                            v11 = v11
                                                        }
                                                    }
                                                case 0:
                                                    v60 := v4[v5]
                                                    _ = v60
                                                }
                                            default:
                                                if false {
                                                    v31 = v31
                                                    if false {
                                                        v5 += v1
                                                        switch (fi40 * fi38) {
                                                        case 1:
                                                            v3 -= v3
                                                        }
                                                        v31 = v31
                                                    }
                                                    var v61 float64 = 0.1
                                                    _ = v61
                                                } else {
                                                    if len(v39) >= 2 {
                                                        v39 = v39[1:len(v39)]
                                                    }
                                                }
                                            }
                                            _ = len(v4)
                                        }
                                    }
                                case -891:
                                    var v62 float64 = 1e10
                                    _ = v62
                                default:
                                    var v63 string = "\n\t"
                                    _ = v63
                                }
                            default:
                                _ = len(v4)
                            }
                            _ = fi40
                        }
                        var v64 int = (-(fi38))
                        _ = v64
                    } else {
                        fmt.Println(v5, v5)
                    }
                    for fi65 := 0; fi65 < 16; fi65++ {
                        if true {
                            var v66 int32 = 42
                            var v67 uint64 = 151
                            switch "a" {
                            case "":
                                if false {
                                    v68 := GenS1{F0: genFn1(genFn1((^(fi65)), v2), (^(v2))), F1: (("" + ("hello" + "hello")) + ("\n\t" + "\n\t"))}
                                    v69 := GenS10{F0: 2}
                                    _ = v68
                                    _ = v69
                                }
                            default:
                                if true {
                                    var v70 int = (-(fi38))
                                    var v71 int16 = -32768
                                    if (42 != 42) {
                                        fmt.Println(v66, fi65, fi38)
                                        fmt.Println(v70)
                                        v72 := GenS6{F0: 128, F1: v1, F2: 1}
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                            }
                            if ((!true) && (!false)) {
                                v73 := GenS6{F0: 0, F1: (v5 << ((109) & 63)), F2: uint8(v1)}
                                v11 += -73
                                var v74 bool = (v31 > v31)
                                switch "\n\t" {
                                case "hello":
                                    if v74 {
                                        if v74 {
                                            if len(v39) > 0 {
                                                v39[0] = fi65
                                            }
                                            v75 := GenS5{F0: (v66 + int32(v67)), F1: "foo bar", F2: v11}
                                            _ = v75
                                        } else {
                                            v2 = v2
                                        }
                                    } else {
                                        v3 -= (v3 + 1e10)
                                        if (!v74) {
                                            if (1e10 >= 3.14159) {
                                                delete(v4, v5)
                                            }
                                        }
                                    }
                                case "foo bar":
                                    switch (fi38 / ((-516) | 1)) {
                                    case 42:
                                        fmt.Println(v31, fi65)
                                    case 0:
                                        var v76 int8 = (-(v11))
                                        _ = v76
                                    default:
                                        v77 := GenS10{F0: 42}
                                        _ = v77
                                    }
                                case "a":
                                    v78 := GenS7{F0: false, F1: 42}
                                    _ = v78
                                default:
                                    if v74 {
                                        v79 := GenS3{F0: (fi65 + (fi38 << ((65535) & 63))), F1: (^(v31)), F2: (v67 + 0)}
                                        _ = len(v4)
                                        _ = v79
                                    }
                                }
                                _ = v73
                                _ = v74
                            }
                            _ = v66
                            _ = v67
                        } else {
                            if (-34 < (-(v11))) {
                                if len(v39) > 0 {
                                    v39[0] = fi38
                                }
                                var v80 int16 = -1
                                _ = v80
                            }
                            if (v5 <= uint32(v31)) {
                                v81 := v4[42]
                                _ = v81
                            }
                        }
                        fmt.Println(v31, v2, v2)
                        switch "" {
                        case "a":
                            var v82 string = "\n\t"
                            _ = v82
                        default:
                            var v83 float64 = 1.5
                            _ = v83
                        }
                        _ = fi65
                    }
                }
                _ = v39
            } else {
                v84 := GenS11{F0: (v11 ^ (^(v11))), F1: (v31 >> ((uint(v5)) & 63)), F2: 2}
                _ = v84
            }
            var v85 string = "hello"
            _ = fi38
            _ = v85
        }
        _ = v37
    }
    v86 := GenS12{F0: (("foo bar" + "\n\t") + "foo bar"), F1: 42}
    for fi87 := 0; fi87 < 20; fi87++ {
        v31 += v31
        if true {
            if true {
                if (false && ((uint8(v5) > 127) && false)) {
                    for fi88 := 0; fi88 < 8; fi88++ {
                        switch -1 {
                        case 32768:
                            v89 := GenS4{F0: 119, F1: 43}
                            var v90 int16 = 32767
                            v5 = v1
                            _ = v89
                            _ = v90
                        case 9223372036854775807:
                            v31 = v31
                        case -223:
                            if false {
                                v2 -= v2
                                delete(v4, 218)
                                if ((!(!(!(v11 != -128)))) || true) {
                                    v91 := GenS13{F0: 1}
                                    var v92 bool = false
                                    _ = v91
                                    _ = v92
                                }
                            }
                        }
                        v93 := GenS8{F0: 4294967295}
                        _ = fi88
                        _ = v93
                    }
                    v5 = 777
                } else {
                    if ((v1 + v5) >= 2147483648) {
                        var v94 uint16 = v31
                        v2 += -79
                        fmt.Println(v5, v1, fi87)
                        _ = v94
                    } else {
                        switch genFn1((fi87 >> ((255) & 63)), v2) {
                        case -2147483648:
                            fmt.Println(v2, v31)
                        case 2147483647:
                            for fi95 := 0; fi95 < 4; fi95++ {
                                v96 := GenS14{F0: 1}
                                switch "hello" {
                                case "\n\t":
                                    v97 := GenS9{F0: uint32(v31), F1: -2147483648}
                                    _ = v97
                                case "foo bar":
                                    v98 := []int{genFn1(fi95, (v2 / ((int64(fi95)) | 1))), genFn1((fi95 % ((9223372036854775807) | 1)), (v2 & v2))}
                                    _ = v98
                                case "":
                                    v99 := GenS14{F0: 1}
                                    _ = v99
                                }
                                v100 := GenS15{F0: uint8(v31), F1: v1, F2: 70}
                                _ = fi95
                                _ = v100
                                _ = v96
                            }
                        default:
                            v101 := v4[v5]
                            _ = v101
                        }
                        switch genFn1((fi87 >> ((0) & 63)), v2) {
                        case -841:
                            switch ("foo bar" + "hello") {
                            case "a":
                                var v102 uint16 = v31
                                _ = v102
                            default:
                                switch (("hello" + ((("" + "a") + (("hello" + (("a" + (("hello" + "foo bar") + "foo bar")) + "hello")) + "\n\t")) + "a")) + "foo bar") {
                                case "\n\t":
                                    v5 -= (v1 ^ v5)
                                case "hello":
                                    for fi103 := 0; fi103 < 17; fi103++ {
                                        var v104 bool = false
                                        var v105 uint8 = 125
                                        _ = fi103
                                        _ = v104
                                        _ = v105
                                    }
                                case "":
                                    v106 := []int{-907}
                                    _ = v106
                                default:
                                    v107 := []int{128, (fi87 ^ -497), (^(fi87)), genFn1(fi87, (v2 >> ((42) & 63)))}
                                    _ = v107
                                }
                            }
                        case 1:
                            if true {
                                var v108 uint32 = 2147483648
                                _ = v108
                            } else {
                                v109 := GenS16{F0: (v1 - 2147483648), F1: v5}
                                v110 := GenS17{F0: 32767, F1: false}
                                for fi111 := 0; fi111 < 7; fi111++ {
                                    v112 := []int{-2147483648, (fi111 % ((fi87) | 1))}
                                    switch 32768 {
                                    case -9223372036854775808:
                                        switch (("" + ("" + "foo bar")) + ("" + ("" + "\n\t"))) {
                                        case "":
                                            if len(v112) > 0 {
                                                v112[0] = genFn1(-875, v2)
                                            }
                                        case "hello":
                                            v113 := GenS18{F0: (^(fi87)), F1: (4294967295 <= v1)}
                                            _ = v113
                                        }
                                    }
                                    if false {
                                        if (0.0 <= ((1.5 + (0.1 * 0.1)) - 1.5)) {
                                            switch 127 {
                                            case 128:
                                                v114 := v4[(v5 * (^(v1)))]
                                                _ = v114
                                            }
                                            if true {
                                                var v115 float64 = 0.0
                                                _ = v115
                                            }
                                        }
                                    } else {
                                        var v116 uint8 = 255
                                        var v117 uint32 = v1
                                        fmt.Println(v5, v2)
                                        _ = v116
                                        _ = v117
                                    }
                                    _ = fi111
                                    _ = v112
                                }
                                _ = v109
                                _ = v110
                            }
                        }
                    }
                    delete(v4, 4294967295)
                }
            }
            for fi118 := 0; fi118 < 1; fi118++ {
                var v119 float64 = (1e10 - (((0.1 + (-2.25 + (0.0 / 1e10))) - 3.14159) * (1.5 * -2.25)))
                var v120 int32 = 2147483647
                _ = fi118
                _ = v119
                _ = v120
            }
        }
        v31 -= (^(v31))
        _ = fi87
    }
    var v121 int64 = v2
    if true {
        if (true && ("" <= "hello")) {
            v5 += v1
            v121 = v121
            var v122 uint8 = 255
            _ = v122
        }
        delete(v4, v1)
    } else {
        for fi123 := 0; fi123 < 12; fi123++ {
            for fi124 := 0; fi124 < 6; fi124++ {
                var v125 string = "foo bar"
                v126 := GenS19{F0: (fi123 >> ((4294967295) & 63))}
                switch v125 {
                case "a":
                    v1 = 390
                case "foo bar":
                    if ((fi123 << ((255) & 63)) == 32767) {
                        v31 += (v31 % (((v31 & 347)) | 1))
                        if false {
                            v31 += v31
                            fmt.Println(v11)
                            v121 = v2
                        }
                        if true {
                            if (!false) {
                                fmt.Println(v11, v31)
                                v127 := GenS20{F0: fi124, F1: v2, F2: false}
                                fmt.Println(fi124, v5)
                                _ = v127
                            }
                            var v128 int32 = -2147483648
                            _ = v128
                        } else {
                            if (v5 != (v5 & 4294967295)) {
                                if true {
                                    v31 -= (v31 ^ v31)
                                    v129 := []int{(fi123 + fi124), 32768}
                                    _ = v129
                                }
                                v130 := GenS1{F0: (fi124 << ((2) & 63)), F1: v125}
                                v3 += v3
                                _ = v130
                            } else {
                                var v131 bool = true
                                _ = v131
                            }
                            v132 := []int{fi124, fi123, 2147483647}
                            var v133 int = genFn1((fi124 * -2147483648), (v121 | (^(v121))))
                            if len(v132) >= 2 {
                                v132 = v132[1:len(v132)]
                            }
                            v134 := GenS21{F0: false, F1: 999}
                            _ = v132
                            _ = v133
                            _ = v134
                        }
                        v11 += 1
                    }
                }
                _ = fi124
                _ = v125
                _ = v126
            }
            _ = fi123
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v121
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v86
}
