package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 bool; F1 int32; F2 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 string }
type GenS11 struct { F0 int16; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint32; F1 uint16; F2 int }
func genFn0() int32 {
    var v6 uint32 = 129
    v7 := GenS1{F0: 246, F1: 323, F2: "hello"}
    switch "hello" {
    case "foo bar":
        v6 += (v6 ^ 65535)
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{}
            v6 = (v6 >> ((232) & 63))
            _ = fi8
            _ = v9
        }
    case "a":
        fmt.Println(v6)
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v6 -= v6
        var v11 int = 32767
        v12 := GenS2{F0: false, F1: -2147483648, F2: 255}
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 += 1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -710
    var v3 uint8 = 128
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    v13 := GenS2{F0: (!true), F1: 42, F2: 42}
    switch 2147483647 {
    case -366:
        v14 := GenS2{F0: (("\n\t" + "foo bar") >= "a"), F1: 2147483647, F2: 4294967295}
        _ = v14
    case 32767:
        var v15 uint64 = 2
        _ = v15
    case -65:
        v2 = v2
    }
    v16 := []int{128}
    var v17 uint16 = 421
    switch 407 {
    case 956:
        if false {
            fmt.Println(v3, v3, v17)
        } else {
            v18 := GenS2{F0: true, F1: 52, F2: 1}
            fmt.Println(v3)
            if ((v3 + 42) >= (v3 + 2)) {
                var v19 uint8 = v3
                _ = len(v1)
                _ = v19
            } else {
                v20 := GenS1{F0: v3, F1: v2, F2: "hello"}
                for fi21 := 0; fi21 < 18; fi21++ {
                    if (3.14159 >= (0.0 - -2.25)) {
                        if (0 <= genFn0()) {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        }
                        var v22 float32 = (1e10 / (0.1 - float32(1e10)))
                        v23 := GenS3{F0: 180, F1: (0 > 42), F2: 401}
                        _ = v22
                        _ = v23
                    }
                    v2 = (v2 << ((2) & 63))
                    _ = fi21
                }
                _ = v20
            }
            _ = v18
        }
    default:
        v3 = (v3 - (v3 - v3))
    }
    delete(v5, 9223372036854775807)
    switch "foo bar" {
    case "":
        for fi24 := 0; fi24 < 8; fi24++ {
            v2 -= (v2 / ((v2) | 1))
            _ = fi24
        }
    case "a":
        v2 -= -261
    }
    if (-5 <= 0) {
        switch ("" + ("" + ("a" + "hello"))) {
        case "\n\t":
            var v25 float64 = 0.1
            _ = v25
        }
    } else {
        var v26 uint8 = (v3 >> ((337) & 63))
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v2 = (v2 ^ 0)
        v3 = (v3 << ((113) & 63))
        v16 = append(v16, (fi27 - (fi27 << ((18446744073709551615) & 63))))
        _ = fi27
    }
    var v28 uint = 65535
    v2 -= (v2 ^ (^(v2)))
    if len(v16) > 0 {
        v16[0] = int(v17)
    }
    v3 = (v3 ^ (v3 ^ 127))
    v29 := GenS3{F0: 18446744073709551615, F1: (-9223372036854775808 >= (v2 & (-(v2)))), F2: 1000}
    switch ("\n\t" + "a") {
    case "":
        switch "a" {
        case "":
            v30 := GenS2{F0: true, F1: 1, F2: v28}
            _ = v30
        case "foo bar":
            var v31 uint64 = 648
            _ = v31
        case "a":
            for fi32 := 0; fi32 < 14; fi32++ {
                if ((((("a" + ("\n\t" + "\n\t")) + "foo bar") + "") + "a") == ("\n\t" + (((("" + "a") + "\n\t") + "\n\t") + "hello"))) {
                    v28 += v28
                }
                if (0.0 < 0.0) {
                    if (v2 <= (v2 ^ v2)) {
                        v28 = uint(v3)
                        if (1 > int16(fi32)) {
                            var v33 uint = v28
                            _ = v33
                        } else {
                            var v34 int8 = 62
                            v35 := GenS3{F0: 2, F1: (false && false), F2: -32768}
                            _ = v34
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 3; fi36++ {
                            if true {
                                v2 = 0
                            }
                            _ = fi36
                        }
                    } else {
                        if true {
                            var v37 bool = false
                            if false {
                                v16 = append(v16, 49)
                            } else {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                if (!true) {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(fi32)
                                    case "":
                                        v2 += -9223372036854775808
                                    case "\n\t":
                                        if (!(42 > 1)) {
                                            var v38 string = "hello"
                                            _ = v38
                                        }
                                    }
                                    v1[fi32] = fi32
                                    v2 = 0
                                }
                                v39 := GenS4{F0: -122, F1: 48, F2: -1}
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 11; fi40++ {
                                _ = len(v5)
                                switch "" {
                                case "\n\t":
                                    var v41 int16 = 42
                                    _ = v41
                                }
                                v3 -= v3
                                _ = fi40
                            }
                            _ = v37
                        }
                        if true {
                            switch fi32 {
                            case -895:
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    var v43 int8 = 25
                                    switch "a" {
                                    case "hello":
                                        if (-1 <= 0) {
                                            var v44 string = "foo bar"
                                            v17 = 0
                                            _ = v44
                                        }
                                    default:
                                        var v45 int16 = -948
                                        var v46 float32 = 3.14159
                                        _ = v45
                                        _ = v46
                                    }
                                    v47 := GenS5{F0: true}
                                    var v48 uint64 = 500
                                    v17 -= uint16(v3)
                                    switch 8 {
                                    case 1:
                                        var v49 int8 = (v43 % ((108) | 1))
                                        _ = v49
                                    }
                                    _ = fi42
                                    _ = v43
                                    _ = v47
                                    _ = v48
                                }
                            default:
                                v50 := GenS5{F0: false}
                                _ = v50
                            }
                            var v51 float32 = (((float32(3.14159) - float32(1e10)) / 1.5) - 1.5)
                            _ = v51
                        }
                    }
                    if true {
                        var v52 uint16 = 2
                        v3 += (v3 * 128)
                        if true {
                            fmt.Println(v17)
                        }
                        _ = v52
                    } else {
                        switch fi32 {
                        case 32767:
                            if (fi32 >= fi32) {
                                v17 = v17
                                var v53 int16 = 1
                                switch fi32 {
                                case 32767:
                                    switch int(v28) {
                                    case 1:
                                        v54 := GenS4{F0: 112, F1: (v53 * -520), F2: -128}
                                        _ = v54
                                    case 468:
                                        v17 += 255
                                    case -161:
                                        if (!(1e10 < (1e10 * 1e10))) {
                                            var v55 int = (fi32 << ((0) & 63))
                                            delete(v4, 2147483648)
                                            switch (fi32 + (v55 / ((9223372036854775807) | 1))) {
                                            case 127:
                                                var v56 float64 = 0.0
                                                _ = v56
                                            case -415:
                                                switch "hello" {
                                                case "":
                                                    var v57 int8 = 100
                                                    _ = v57
                                                case "hello":
                                                    if true {
                                                        var v58 uint = v28
                                                        fmt.Println(v28)
                                                        var v59 int = (v55 * v55)
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                case "foo bar":
                                                    for fi60 := 0; fi60 < 13; fi60++ {
                                                        switch (v55 / ((1) | 1)) {
                                                        case -425:
                                                            if (!("foo bar" >= ("a" + "a"))) {
                                                                switch fi60 {
                                                                case -2147483648:
                                                                    var v61 uint16 = v17
                                                                    _ = v61
                                                                case 1:
                                                                    switch (fi60 | (v55 ^ (fi60 & fi60))) {
                                                                    case -687:
                                                                        v62 := GenS3{F0: 4294967295, F1: (v2 <= v2), F2: (v53 >> ((uint(v3)) & 63))}
                                                                        _ = v62
                                                                    }
                                                                }
                                                                var v63 int32 = int32(v28)
                                                                if ((v63 / ((v63) | 1)) > (v63 ^ genFn0())) {
                                                                    v64 := v5[42]
                                                                    fmt.Println(v3)
                                                                    var v65 int32 = 1
                                                                    fmt.Println(v17, v65)
                                                                    fmt.Println(v53, v65, v2)
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    v3 -= v3
                                                                    v66 := GenS6{F0: (fi60 | (^(v55)))}
                                                                    v67 := GenS4{F0: 127, F1: (v53 + (^(v53))), F2: 100}
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                _ = v63
                                                            }
                                                            fmt.Println(v28)
                                                            var v68 uint32 = 159
                                                            _ = v68
                                                        case -9223372036854775808:
                                                            v69 := GenS4{F0: -86, F1: (^(v53)), F2: 15}
                                                            _ = v69
                                                        default:
                                                            v70 := GenS6{F0: (fi32 >> ((v28) & 63))}
                                                            _ = v70
                                                        }
                                                        switch (fi60 >> (((v28 % (((v28 & 106)) | 1))) & 63)) {
                                                        case 128:
                                                            var v71 int8 = -128
                                                            _ = v71
                                                        case 269:
                                                            if ((fi60 << (((v28 | (v28 + (v28 >> ((v28) & 63))))) & 63)) > -9223372036854775808) {
                                                                if (!false) {
                                                                    v55 -= fi32
                                                                    if true {
                                                                        v28 -= (v28 + v28)
                                                                        var v72 float64 = (-2.25 * 3.14159)
                                                                        v73 := GenS1{F0: 184, F1: v2, F2: ("\n\t" + "")}
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        if (!(v17 >= 2)) {
                                                                            _ = len(v1)
                                                                            if (-2147483648 < fi60) {
                                                                                if true {
                                                                                    fmt.Println(v55, v3, v53)
                                                                                    var v74 uint32 = 2147483648
                                                                                    _ = v74
                                                                                }
                                                                                v1[(fi32 & (fi32 ^ (fi60 & -502)))] = (fi32 >> ((4294967295) & 63))
                                                                                _ = len(v16)
                                                                                if false {
                                                                                    var v75 int32 = 2147483647
                                                                                    _ = v75
                                                                                } else {
                                                                                    fmt.Println(fi60, v53, fi32)
                                                                                    v76 := GenS1{F0: 0, F1: (v2 * -9223372036854775808), F2: ""}
                                                                                    _ = v76
                                                                                }
                                                                                if false {
                                                                                    v3 = v3
                                                                                } else {
                                                                                    var v77 bool = false
                                                                                    var v78 float32 = (float32((0.1 - (0.0 + 0.0))) - (0.0 - -2.25))
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                }
                                                                                var v79 bool = ((v17 % (((^(v17))) | 1)) >= v17)
                                                                                v80 := GenS1{F0: v3, F1: v2, F2: "hello"}
                                                                                var v81 int32 = 1
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                fmt.Println(v28, v17)
                                                                            }
                                                                            if true {
                                                                                var v82 float64 = 3.14159
                                                                                if (((v17 & 48) == (v17 * v17)) || ((241 >= (v17 / (((v17 ^ (^(v17)))) | 1))) || ((fi60 >> ((v28) & 63)) > v55))) {
                                                                                    if false {
                                                                                        v82 += 1e10
                                                                                    } else {
                                                                                        var v83 int16 = v53
                                                                                        var v84 uint16 = 820
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                    }
                                                                                }
                                                                                fmt.Println(v55, v53)
                                                                                _ = v82
                                                                            } else {
                                                                                fmt.Println(v55, v55, v28)
                                                                            }
                                                                            switch (fi60 - (fi60 % ((fi32) | 1))) {
                                                                            case 652:
                                                                                var v85 bool = true
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        v28 += v28
                                                                        var v86 int32 = 1
                                                                        _ = v86
                                                                    }
                                                                    if ((v28 + 2) != 255) {
                                                                        if true {
                                                                            switch "" {
                                                                            case "a":
                                                                                fmt.Println(v53, v3, v53)
                                                                            case "\n\t":
                                                                                var v87 string = ("\n\t" + "a")
                                                                                _ = v87
                                                                            }
                                                                            var v88 uint32 = 1
                                                                            _ = v88
                                                                        }
                                                                        v17 = 2
                                                                        switch fi32 {
                                                                        case 128:
                                                                            v53 = v53
                                                                            v53 -= v53
                                                                        case 504:
                                                                            var v89 int64 = v2
                                                                            _ = v89
                                                                        case 32768:
                                                                            _ = len(v16)
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    v90 := GenS7{F0: (v17 <= 2)}
                                                                    v53 = v53
                                                                    if (0 != 294) {
                                                                        if false {
                                                                            if false {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi60, v55)
                                                                                case "":
                                                                                    var v91 uint8 = v3
                                                                                    _ = v91
                                                                                default:
                                                                                    v53 = v53
                                                                                }
                                                                                switch (fi32 << (((v28 / ((2) | 1))) & 63)) {
                                                                                case 414:
                                                                                    v92 := GenS8{F0: 42, F1: (v3 * (v3 << ((v28) & 63)))}
                                                                                    _ = v92
                                                                                default:
                                                                                    if false {
                                                                                        v2 -= (v2 & -9223372036854775808)
                                                                                    } else {
                                                                                        v3 -= v3
                                                                                        if ((4294967295 < v28) || false) {
                                                                                            v53 = v53
                                                                                            var v93 int32 = 1
                                                                                            if (598 >= 4294967295) {
                                                                                                var v94 uint32 = 65535
                                                                                                var v95 int8 = -1
                                                                                                var v96 uint64 = 0
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            }
                                                                                            _ = v93
                                                                                        } else {
                                                                                            _ = len(v1)
                                                                                            var v97 int8 = 67
                                                                                            if len(v16) >= 2 {
                                                                                                v16 = v16[1:len(v16)]
                                                                                            }
                                                                                            var v98 bool = false
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        }
                                                                                        switch (^(fi32)) {
                                                                                        case 9223372036854775807:
                                                                                            var v99 float32 = 3.14159
                                                                                            _ = v99
                                                                                        case 20:
                                                                                            switch -9223372036854775808 {
                                                                                            case 1:
                                                                                                var v100 float64 = (0.1 * 1.5)
                                                                                                _ = v100
                                                                                            case 620:
                                                                                                if len(v16) > 0 {
                                                                                                    v16[0] = v55
                                                                                                }
                                                                                            case 32767:
                                                                                                var v101 int64 = (v2 - 42)
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if (("\n\t" + "") < "a") {
                                                                                    v2 += v2
                                                                                    var v102 int8 = 25
                                                                                    _ = v102
                                                                                } else {
                                                                                    var v103 bool = true
                                                                                    v104 := GenS3{F0: uint64(v28), F1: false, F2: v53}
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                                var v105 uint8 = (^(v3))
                                                                                _ = v105
                                                                            }
                                                                            fmt.Println(fi32, fi32)
                                                                            v106 := v5[42]
                                                                            _ = v106
                                                                        } else {
                                                                            v107 := GenS3{F0: 42, F1: true, F2: v53}
                                                                            switch fi60 {
                                                                            case -48:
                                                                                var v108 int16 = v53
                                                                                _ = v108
                                                                            case 32768:
                                                                                v109 := GenS9{F0: -1, F1: (v3 + (v3 << ((0) & 63)))}
                                                                                _ = v109
                                                                            }
                                                                            var v110 int64 = v2
                                                                            v5[-974] = (1e10 <= (float64(0.0) * float64(float32(float64(0.1)))))
                                                                            var v111 int16 = (^(v53))
                                                                            _ = v107
                                                                            _ = v110
                                                                            _ = v111
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            v17 += (^(v17))
                                                                            if (9223372036854775807 <= 9223372036854775807) {
                                                                                fmt.Println(fi60)
                                                                                v112 := v1[4294967295]
                                                                                _ = v112
                                                                            } else {
                                                                                switch v55 {
                                                                                case 0:
                                                                                    v113 := GenS3{F0: 1, F1: (true || ("hello" < "foo bar")), F2: 0}
                                                                                    _ = v113
                                                                                case -1:
                                                                                    _ = len(v5)
                                                                                case 2147483647:
                                                                                    var v114 int64 = (v2 | v2)
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                        }
                                                                        switch (("\n\t" + "") + ("foo bar" + ("" + ""))) {
                                                                        case "foo bar":
                                                                            v115 := v1[fi32]
                                                                            _ = v115
                                                                        case "a":
                                                                            if true {
                                                                                fmt.Println(v3, v55)
                                                                                fmt.Println(v17)
                                                                                v16 = append(v16, -2147483648)
                                                                                var v116 float32 = 3.14159
                                                                                _ = v116
                                                                            } else {
                                                                                var v117 int8 = -5
                                                                                _ = v117
                                                                            }
                                                                        default:
                                                                            var v118 uint64 = 976
                                                                            _ = v118
                                                                        }
                                                                        v3 = (v3 - (v3 & uint8(v28)))
                                                                    }
                                                                    switch fi60 {
                                                                    case 9223372036854775807:
                                                                        v5[(fi60 / (((-(fi60))) | 1))] = true
                                                                    case -229:
                                                                        switch fi32 {
                                                                        case -1:
                                                                            v119 := GenS2{F0: true, F1: 9, F2: (^(v28))}
                                                                            var v120 int32 = 42
                                                                            if (!(178 == (v28 * 18446744073709551615))) {
                                                                                v28 += (v28 * (v28 << ((605) & 63)))
                                                                            }
                                                                            fmt.Println(v55, fi60)
                                                                            _ = v119
                                                                            _ = v120
                                                                        case -2147483648:
                                                                            var v121 uint64 = 2
                                                                            _ = v121
                                                                        case -879:
                                                                            v3 += v3
                                                                        default:
                                                                            switch fi60 {
                                                                            case 128:
                                                                                v2 = v2
                                                                            default:
                                                                                v122 := GenS2{F0: (0 < -523), F1: 42, F2: v28}
                                                                                _ = v122
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v90
                                                                } else {
                                                                    var v123 int32 = genFn0()
                                                                    _ = v123
                                                                }
                                                            }
                                                        case 4294967295:
                                                            if len(v16) > 0 {
                                                                v16[0] = (v55 & (fi60 >> ((42) & 63)))
                                                            }
                                                        }
                                                        v16 = append(v16, fi32)
                                                        if true {
                                                            delete(v5, -59)
                                                            v2 += v2
                                                            var v124 float64 = (3.14159 + 0.1)
                                                            _ = v124
                                                        }
                                                        _ = fi60
                                                    }
                                                default:
                                                    if len(v16) > 0 {
                                                        v16[0] = (fi32 >> ((v28) & 63))
                                                    }
                                                }
                                            case 782:
                                                var v125 uint32 = uint32(v3)
                                                _ = v125
                                            default:
                                                for fi126 := 0; fi126 < 17; fi126++ {
                                                    v127 := GenS10{F0: 65535, F1: 0, F2: ""}
                                                    v5[fi32] = true
                                                    v28 += v28
                                                    _ = fi126
                                                    _ = v127
                                                }
                                            }
                                            _ = v55
                                        } else {
                                            var v128 int16 = -1
                                            _ = v128
                                        }
                                    }
                                default:
                                    switch 128 {
                                    case 1:
                                        v17 = 647
                                    case -878:
                                        var v129 uint16 = v17
                                        _ = v129
                                    default:
                                        v17 = v17
                                    }
                                }
                                _ = v53
                            }
                        case -2147483648:
                            for fi130 := 0; fi130 < 11; fi130++ {
                                var v131 uint32 = 141
                                v3 += (v3 % (((v3 | (^(v3)))) | 1))
                                if false {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    v28 += (v28 << ((v28) & 63))
                                    fmt.Println(v2, v2)
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        switch fi32 {
                                        case 925:
                                            fmt.Println(fi32, v131, fi32)
                                        case -630:
                                            v2 += (v2 % (((v2 >> ((42) & 63))) | 1))
                                        }
                                    case "hello":
                                        v16 = append(v16, (fi32 / (((fi130 >> (((v28 >> ((uint(v17)) & 63))) & 63))) | 1)))
                                    case "a":
                                        v17 += uint16(v3)
                                    default:
                                        var v132 bool = false
                                        _ = v132
                                    }
                                }
                                _ = fi130
                                _ = v131
                            }
                        case 1:
                            var v133 string = "hello"
                            _ = v133
                        }
                    }
                    switch (fi32 + 32768) {
                    case 519:
                        for fi134 := 0; fi134 < 17; fi134++ {
                            switch fi134 {
                            case -2147483648:
                                var v135 uint32 = 4294967295
                                _ = v135
                            }
                            if (true || false) {
                                if false {
                                    v3 = 44
                                    _ = len(v16)
                                    if true {
                                        if true {
                                            var v136 int8 = 117
                                            if true {
                                                v137 := GenS8{F0: (fi32 - fi32), F1: (v3 | 42)}
                                                v28 -= 0
                                                _ = v137
                                            }
                                            v2 -= v2
                                            _ = v136
                                        }
                                        if false {
                                            var v138 string = ("a" + ("" + ("\n\t" + "foo bar")))
                                            var v139 float32 = float32((3.14159 + (0.0 - 3.14159)))
                                            v3 = (v3 + (v3 << (((v28 ^ v28)) & 63)))
                                            v140 := GenS11{F0: -32768, F1: ((v28 / (((^(v28))) | 1)) > v28), F2: 42}
                                            _ = v138
                                            _ = v139
                                            _ = v140
                                        }
                                        fmt.Println(v3, v28)
                                        v2 += (v2 | -1)
                                        switch fi134 {
                                        case 127:
                                            v2 = -615
                                        case 932:
                                            switch ("hello" + "a") {
                                            case "hello":
                                                switch (fi134 | (fi32 / ((32767) | 1))) {
                                                case 836:
                                                    v141 := GenS8{F0: fi32, F1: (^(v3))}
                                                    v3 = 1
                                                    var v142 int32 = 42
                                                    _ = v141
                                                    _ = v142
                                                }
                                            case "\n\t":
                                                v143 := GenS6{F0: fi134}
                                                _ = v143
                                            default:
                                                v17 = v17
                                            }
                                        }
                                    } else {
                                        var v144 uint = (v28 - (v28 - (v28 / (((v28 | (v28 << ((2) & 63)))) | 1))))
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        _ = len(v1)
                                        _ = v144
                                    }
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v2 = (v2 * int64(v28))
                                    case "hello":
                                        v145 := GenS3{F0: 1, F1: ((false && true) || ((0.1 - ((0.0 - (float64((1e10 + -2.25)) + -2.25)) - (float64(((-2.25 - (1e10 - (0.0 * 0.0))) / 0.1)) - (0.0 / 3.14159)))) >= (float64(0.1) / 1e10))), F2: 1000}
                                        _ = v145
                                    case "":
                                        v146 := GenS6{F0: 4294967295}
                                        _ = v146
                                    default:
                                        var v147 int32 = 0
                                        _ = v147
                                    }
                                    v3 += v3
                                    _ = len(v16)
                                }
                            }
                            var v148 uint64 = 1
                            _ = fi134
                            _ = v148
                        }
                    case -2147483648:
                        if (647 > 32767) {
                            var v149 int32 = 1
                            for fi150 := 0; fi150 < 20; fi150++ {
                                v151 := GenS12{F0: 0, F1: v17, F2: (fi32 * fi32)}
                                _ = fi150
                                _ = v151
                            }
                            _ = v149
                        } else {
                            for fi152 := 0; fi152 < 8; fi152++ {
                                v17 = 1
                                _ = fi152
                            }
                        }
                    default:
                        v17 = 255
                    }
                }
                _ = fi32
            }
        default:
            v28 += v28
        }
    case "foo bar":
        if (!true) {
            delete(v1, 2147483647)
            switch -2147483648 {
            case 181:
                v2 = v2
            case 32767:
                var v153 int8 = 100
                _ = v153
            case 1:
                if true {
                    v2 -= 1
                }
            }
            delete(v4, 2)
        }
    default:
        fmt.Println(v3, v28)
    }
    var v154 uint64 = 963
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v154:", v154)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v154
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
