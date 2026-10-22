package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
type GenS5 struct { F0 uint32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint, p1 string, p2 uint16) uint16 {
    var v7 uint32 = 1
    if (-32768 != -1) {
        var v8 uint64 = 42
        _ = v8
    } else {
        v7 = 0
    }
    var v9 int8 = -1
    p1 = p1
    p1 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return p2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 2147483647
    v3 := []uint32{65535, 258, 1}
    v4 := make(map[uint32]string)
    var v5 uint8 = 2
    var v6 uint8 = (v5 << ((18446744073709551615) & 63))
    v6 -= 127
    for fi10 := 0; fi10 < 18; fi10++ {
        delete(v4, 839)
        for fi11 := 0; fi11 < 9; fi11++ {
            switch fi11 {
            case -495:
                if false {
                    switch (-(fi11)) {
                    case -198:
                        switch "\n\t" {
                        case "foo bar":
                            var v12 int64 = 42
                            _ = v12
                        default:
                            if false {
                                v13 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                                _ = v13
                            }
                        }
                        var v14 uint = 878
                        v2 -= fi10
                        v15 := v1[42]
                        _ = v14
                        _ = v15
                    case 42:
                        if true {
                            switch v2 {
                            case -872:
                                if false {
                                    fmt.Println(v5, v5)
                                }
                            }
                            fmt.Println(fi10, v6, fi11)
                            v16 := GenS2{F0: 1, F1: 2, F2: 42}
                            _ = v16
                        } else {
                            switch (("foo bar" + ("hello" + ("\n\t" + "foo bar"))) + "hello") {
                            case "":
                                switch (fi11 | (v2 | (fi11 ^ fi10))) {
                                case 42:
                                    v5 -= 1
                                case 9223372036854775807:
                                    var v17 bool = true
                                    _ = v17
                                default:
                                    var v18 int32 = 628
                                    _ = v18
                                }
                            case "foo bar":
                                v2 = 958
                            case "hello":
                                fmt.Println(fi11, fi11, fi11)
                            }
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v19 int = fi10
                    _ = v19
                } else {
                    v20 := GenS2{F0: 2, F1: 422, F2: 1}
                    _ = v20
                }
            case 1:
                v3 = append(v3, 433)
            default:
                v5 += v6
            }
            _ = fi11
        }
        _ = fi10
    }
    v21 := GenS1{F0: 1, F1: 18446744073709551615}
    v22 := GenS1{F0: -9223372036854775808, F1: 65535}
    var v23 float64 = 0.1
    switch -220 {
    case 351:
        if (!(42 > -1)) {
            v24 := GenS1{F0: 93, F1: 1}
            _ = len(v3)
            _ = v24
        }
    case 128:
        var v25 int32 = -1
        _ = v25
    }
    switch "" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    switch "a" {
    case "hello":
        var v26 int8 = int8(v2)
        _ = v26
    case "foo bar":
        v1[1] = (v2 ^ v2)
    }
    v5 = v6
    switch (v2 & (v2 * 774)) {
    case -299:
        switch (("hello" + ("a" + "")) + "foo bar") {
        case "foo bar":
            var v27 uint64 = 291
            _ = v27
        case "a":
            v5 += (^(v6))
        }
    case 384:
        switch v2 {
        case 2147483647:
            var v28 int16 = 184
            _ = v28
        default:
            var v29 int32 = 1
            _ = v29
        }
    default:
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int8 = 0
            if true {
                v32 := GenS3{F0: 9223372036854775807, F1: (v5 / ((v5) | 1)), F2: 18446744073709551615}
                v5 = (v5 + v6)
                _ = v32
            }
            var v33 float32 = 0.0
            _ = fi30
            _ = v31
            _ = v33
        }
    }
    fmt.Println(v6)
    if false {
        if false {
            fmt.Println(v23)
            for fi34 := 0; fi34 < 9; fi34++ {
                v35 := GenS4{F0: -78, F1: "\n\t", F2: 1}
                var v36 uint = 4294967295
                var v37 int32 = 0
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 20; fi38++ {
                switch "hello" {
                case "":
                    var v39 bool = false
                    _ = v39
                case "\n\t":
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 int32 = -1
                        var v42 int64 = 1
                        v5 -= v6
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                case "hello":
                    v23 = ((v23 / (((v23 / 1e10) * v23) / v23)) - ((((float64(0.0) + v23) + 1.5) / v23) / (((float64((0.0 * (1.5 * (((((0.1 / 1.0) - ((0.1 - (1.5 * 1.5)) / 1e10)) - (0.0 / 0.1)) - 3.14159) + 1.5)))) + float64(0.1)) + ((v23 - v23) * float64(float32(float64(((-2.25 / 1.0) * float32(v23))))))) - v23)))
                }
                switch ("hello" + "\n\t") {
                case "hello":
                    switch (fi38 << ((242) & 63)) {
                    case 979:
                        var v43 float32 = ((-2.25 * (3.14159 * 3.14159)) - (0.1 + (1.5 * -2.25)))
                        _ = v43
                    case 2147483647:
                        v5 = (v6 % (((^(v6))) | 1))
                    case 1:
                        for fi44 := 0; fi44 < 16; fi44++ {
                            fmt.Println(v23)
                            _ = fi44
                        }
                    }
                case "":
                    if (!(!true)) {
                        v45 := GenS2{F0: 2, F1: 255, F2: 727}
                        v46 := GenS2{F0: 952, F1: 0, F2: 2}
                        if true {
                            switch v2 {
                            case 177:
                                fmt.Println(v6, fi38, fi38)
                            }
                            var v47 int64 = 9223372036854775807
                            if (0 < -32768) {
                                var v48 uint8 = 200
                                switch ("foo bar" + ("" + "a")) {
                                case "":
                                    v49 := GenS1{F0: v47, F1: 920}
                                    _ = v49
                                case "a":
                                    for fi50 := 0; fi50 < 13; fi50++ {
                                        if (!(1 >= 42)) {
                                            switch "" {
                                            case "\n\t":
                                                if (false && (!(1e10 < (float64(float32((v23 + 1.5))) / 1.5)))) {
                                                    var v51 int16 = -1
                                                    if (true && true) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    _ = v51
                                                }
                                            case "hello":
                                                switch (v2 / ((2147483647) | 1)) {
                                                case 42:
                                                    var v52 float64 = ((3.14159 + -2.25) * float64((1e10 / float32(v23))))
                                                    _ = v52
                                                case 782:
                                                    v53 := GenS3{F0: 1, F1: (v6 & (v6 - v48)), F2: 18446744073709551615}
                                                    _ = v53
                                                case 0:
                                                    if (!false) {
                                                        switch -1 {
                                                        case -544:
                                                            if (1000 != 0) {
                                                                fmt.Println(v48)
                                                            } else {
                                                                v5 = v5
                                                                fmt.Println(v5, v6, fi38)
                                                            }
                                                        case 42:
                                                            v23 = float64(0.0)
                                                        default:
                                                            if (0 == 65535) {
                                                                v48 += 255
                                                                switch "a" {
                                                                case "foo bar":
                                                                    _ = len(v4)
                                                                case "hello":
                                                                    v2 -= fi38
                                                                }
                                                            } else {
                                                                fmt.Println(v2, v47, v5)
                                                                v48 += (v5 | (v6 | v48))
                                                            }
                                                        }
                                                        _ = len(v3)
                                                        var v54 float32 = 0.0
                                                        _ = v54
                                                    }
                                                default:
                                                    switch (("hello" + "\n\t") + "a") {
                                                    case "a":
                                                        fmt.Println(v47, v6, fi50)
                                                    case "foo bar":
                                                        if true {
                                                            v3 = append(v3, 419)
                                                            v6 = v48
                                                        }
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi50
                                    }
                                default:
                                    var v55 int64 = -536
                                    _ = v55
                                }
                                v2 = fi38
                                _ = v48
                            } else {
                                var v56 uint16 = 1
                                if false {
                                    v57 := GenS2{F0: v56, F1: 696, F2: 2}
                                    if (2 != 1) {
                                        if ((v47 & -1) != v47) {
                                            v47 += v47
                                            var v58 float64 = float64(1e10)
                                            fmt.Println(v2, v47, v6)
                                            _ = v58
                                        } else {
                                            var v59 uint64 = 4294967295
                                            _ = v59
                                        }
                                        var v60 int64 = (v47 | (v47 >> ((42) & 63)))
                                        _ = v60
                                    }
                                    if true {
                                        if (!(-110 <= 1)) {
                                            switch "a" {
                                            case "a":
                                                for fi61 := 0; fi61 < 4; fi61++ {
                                                    v62 := GenS4{F0: 1000, F1: "hello", F2: v47}
                                                    v47 = v47
                                                    _ = fi61
                                                    _ = v62
                                                }
                                            case "foo bar":
                                                switch (v2 % ((fi38) | 1)) {
                                                case 810:
                                                    fmt.Println(fi38, v5, v56)
                                                }
                                            case "hello":
                                                for fi63 := 0; fi63 < 2; fi63++ {
                                                    if true {
                                                        v47 += (v47 * 0)
                                                        v56 -= (v56 + (v56 + v56))
                                                        var v64 uint16 = v56
                                                        _ = v64
                                                    }
                                                    switch fi63 {
                                                    case 127:
                                                        var v65 bool = false
                                                        _ = v65
                                                    case -240:
                                                        v56 -= genFn0(65535, ("a" + "\n\t"), v56)
                                                    case -2147483648:
                                                        v6 = v6
                                                    default:
                                                        fmt.Println(fi63, v6)
                                                        if true {
                                                            v2 -= (fi63 / (((-(fi38))) | 1))
                                                            if false {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v2 -= 625
                                                                case "":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 1
                                                                    }
                                                                default:
                                                                    v5 += v6
                                                                }
                                                                if true {
                                                                    fmt.Println(v6)
                                                                }
                                                                _ = len(v3)
                                                            } else {
                                                                var v66 int16 = -32768
                                                                _ = v66
                                                            }
                                                        } else {
                                                            if (false && false) {
                                                                var v67 uint = 4294967295
                                                                _ = v67
                                                            } else {
                                                                delete(v1, 168)
                                                                v68 := GenS3{F0: v47, F1: 0, F2: 1}
                                                                _ = len(v4)
                                                                if true {
                                                                    if (-9223372036854775808 < -9223372036854775808) {
                                                                        v69 := GenS2{F0: (v56 | v56), F1: uint(v5), F2: 4294967295}
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        if (0.1 <= 0.0) {
                                                                            if (v23 < 0.1) {
                                                                                v70 := v1[817]
                                                                                fmt.Println(v5)
                                                                                if (1 < 1) {
                                                                                    if false {
                                                                                        if (42 > 42) {
                                                                                            if (!(32767 == 578)) {
                                                                                                fmt.Println(v23, v70)
                                                                                                v71 := GenS3{F0: v47, F1: (v6 + (v5 % (((v5 + (^(v5)))) | 1))), F2: 18446744073709551615}
                                                                                                _ = v71
                                                                                            }
                                                                                            var v72 int16 = 0
                                                                                            _ = v72
                                                                                        } else {
                                                                                            var v73 int = fi38
                                                                                            if (false || true) {
                                                                                                var v74 float64 = ((float64(-2.25) / v23) + (v23 / v23))
                                                                                                v75 := GenS1{F0: v47, F1: 42}
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            } else {
                                                                                                v6 += (v6 & (v6 % (((^(v5))) | 1)))
                                                                                            }
                                                                                            if false {
                                                                                                if false {
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        v70 += 4294967295
                                                                                                    case "foo bar":
                                                                                                        if (!(0.0 != 3.14159)) {
                                                                                                            var v76 string = ("foo bar" + ("foo bar" + "foo bar"))
                                                                                                            v6 -= 151
                                                                                                            v77 := GenS5{F0: 970, F1: -9223372036854775808, F2: 1000}
                                                                                                            _ = v76
                                                                                                            _ = v77
                                                                                                        } else {
                                                                                                            var v78 uint16 = 255
                                                                                                            var v79 int64 = (v47 >> ((1) & 63))
                                                                                                            var v80 string = "foo bar"
                                                                                                            v80 = (("" + "") + v80)
                                                                                                            _ = v78
                                                                                                            _ = v79
                                                                                                            _ = v80
                                                                                                        }
                                                                                                    default:
                                                                                                        delete(v4, 924)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v70
                                                                            }
                                                                            switch (v2 | (^(fi38))) {
                                                                            case 761:
                                                                                v5 = v6
                                                                            case 4294967295:
                                                                                fmt.Println(v6)
                                                                            default:
                                                                                if (2 < 18446744073709551615) {
                                                                                    if false {
                                                                                        if true {
                                                                                            var v81 uint8 = (v5 * (v5 / ((255) | 1)))
                                                                                            var v82 uint16 = (v56 << ((2) & 63))
                                                                                            var v83 uint64 = 994
                                                                                            _ = v81
                                                                                            _ = v82
                                                                                            _ = v83
                                                                                        }
                                                                                        var v84 int = (fi63 | v2)
                                                                                        var v85 int = (v2 / ((-9223372036854775808) | 1))
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 int64 = (-(v47))
                                                                                    if false {
                                                                                        v2 = 2147483647
                                                                                        v2 += fi63
                                                                                    } else {
                                                                                        fmt.Println(v23, v6)
                                                                                        v2 -= fi63
                                                                                        v3 = append(v3, 2147483647)
                                                                                        fmt.Println(v23, fi63, v56)
                                                                                        if false {
                                                                                            v47 += (^(v86))
                                                                                            switch fi63 {
                                                                                            case 127:
                                                                                                var v87 int16 = -553
                                                                                                _ = v87
                                                                                            }
                                                                                            fmt.Println(v86)
                                                                                        } else {
                                                                                            var v88 int = 32768
                                                                                            v89 := v1[4294967295]
                                                                                            var v90 int16 = 720
                                                                                            _ = len(v4)
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                        }
                                                                                    }
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                            if (-1 > int32(v6)) {
                                                                                if (float32(((float64(((1e10 + 3.14159) - ((float32(v23) + 3.14159) / 3.14159))) / (float64(0.1) / (v23 * v23))) / v23)) == 1e10) {
                                                                                    fmt.Println(v56, v5)
                                                                                }
                                                                                fmt.Println(v23, v23, v5)
                                                                                fmt.Println(v47, v47)
                                                                            } else {
                                                                                v91 := v1[uint32(v6)]
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    }
                                                                    var v92 uint = 2
                                                                    var v93 uint = v92
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    var v94 uint8 = (v5 | v6)
                                                                    var v95 int8 = int8(v47)
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                _ = v68
                                                            }
                                                            var v96 uint = 2
                                                            var v97 float32 = 1e10
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        v47 = int64(v56)
                                                    }
                                                    _ = fi63
                                                }
                                            default:
                                                var v98 float64 = (float64(0.1) - ((0.0 - float64((3.14159 / (float32(v23) + 0.1)))) - v23))
                                                _ = v98
                                            }
                                            for fi99 := 0; fi99 < 13; fi99++ {
                                                if (float64(1.5) != 3.14159) {
                                                    var v100 uint = 230
                                                    if ((false || false) && true) {
                                                        if false {
                                                            var v101 int64 = 42
                                                            _ = v101
                                                        } else {
                                                            var v102 int64 = 937
                                                            _ = v102
                                                        }
                                                        v103 := GenS6{F0: 4294967295}
                                                        if true {
                                                            if len(v3) > 0 {
                                                                v3[0] = 42
                                                            }
                                                            fmt.Println(v100, v23)
                                                        } else {
                                                            if true {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    switch (fi38 >> ((v100) & 63)) {
                                                                    case 32768:
                                                                        v1[230] = 0
                                                                    case 0:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 4294967295
                                                                        }
                                                                    default:
                                                                        v2 -= int(v6)
                                                                    }
                                                                case "a":
                                                                    if (v5 >= v6) {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            switch (fi99 << ((v100) & 63)) {
                                                                            case 0:
                                                                                if (711 < 468) {
                                                                                    switch fi38 {
                                                                                    case 32767:
                                                                                        _ = len(v1)
                                                                                    default:
                                                                                        var v104 uint64 = 2
                                                                                        _ = v104
                                                                                    }
                                                                                    var v105 int = (^(fi99))
                                                                                    v106 := GenS2{F0: (v56 >> ((42) & 63)), F1: 18446744073709551615, F2: 18446744073709551615}
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                            case 1:
                                                                                var v107 uint32 = 516
                                                                                _ = v107
                                                                            case 128:
                                                                                v108 := GenS5{F0: 1, F1: (v47 + -9223372036854775808), F2: -1}
                                                                                _ = v108
                                                                            }
                                                                        default:
                                                                            switch "" {
                                                                            case "hello":
                                                                                switch v2 {
                                                                                case 0:
                                                                                    v2 += fi38
                                                                                case 9223372036854775807:
                                                                                    v109 := GenS6{F0: 2147483648}
                                                                                    _ = v109
                                                                                case -108:
                                                                                    v3 = append(v3, 2)
                                                                                default:
                                                                                    var v110 uint32 = 2147483647
                                                                                    _ = v110
                                                                                }
                                                                            case "foo bar":
                                                                                switch (fi38 | (fi38 ^ (fi99 | fi38))) {
                                                                                case 2147483647:
                                                                                    var v111 uint = 42
                                                                                    _ = v111
                                                                                case -88:
                                                                                    v112 := v4[19]
                                                                                    _ = v112
                                                                                case -1:
                                                                                    v23 = v23
                                                                                }
                                                                            case "":
                                                                                if (v47 != 0) {
                                                                                    if ((!false) && ((v100 / (((v100 / (((v100 % (((v100 & (^(v100)))) | 1))) | 1))) | 1)) == uint(v56))) {
                                                                                        v100 += 327
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v23 += v23
                                                                                    case "foo bar":
                                                                                        if (680 != fi99) {
                                                                                            v3 = append(v3, uint32(v100))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v103
                                                    } else {
                                                        v5 = v6
                                                    }
                                                    _ = v100
                                                } else {
                                                    if true {
                                                        v23 = v23
                                                    } else {
                                                        v5 -= 0
                                                    }
                                                    var v113 float32 = (3.14159 / 3.14159)
                                                    _ = v113
                                                }
                                                var v114 int64 = v47
                                                _ = fi99
                                                _ = v114
                                            }
                                        }
                                        fmt.Println(v47, v5, v5)
                                        fmt.Println(v23)
                                    }
                                    _ = v57
                                }
                                _ = v56
                            }
                            _ = v47
                        } else {
                            var v115 uint32 = 251
                            for fi116 := 0; fi116 < 3; fi116++ {
                                var v117 uint64 = 489
                                _ = fi116
                                _ = v117
                            }
                            _ = v115
                        }
                        _ = v45
                        _ = v46
                    }
                }
                _ = fi38
            }
        } else {
            fmt.Println(v23, v23, v6)
            for fi118 := 0; fi118 < 16; fi118++ {
                v4[0] = ("foo bar" + "\n\t")
                switch (v2 >> ((42) & 63)) {
                case 912:
                    v5 = v6
                case 148:
                    v119 := GenS5{F0: 42, F1: 9223372036854775807, F2: 246}
                    _ = v119
                }
                _ = fi118
            }
            switch v2 {
            case 127:
                var v120 uint64 = 680
                _ = v120
            default:
                var v121 string = ("foo bar" + (("hello" + "") + ("" + "a")))
                _ = v121
            }
        }
    }
    var v122 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v122
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
