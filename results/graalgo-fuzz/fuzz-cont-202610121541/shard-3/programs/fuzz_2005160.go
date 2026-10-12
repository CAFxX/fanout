package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 bool; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 string; F2 uint32 }
type GenS12 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 string, p2 uint16) uint32 {
    fmt.Println(p2, p2, p2)
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            p2 += p2
            _ = fi6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 float64 = (0.1 + 3.14159)
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        var v10 float32 = 1e10
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 16
}
func genFn1(p0 uint8, p1 bool) int32 {
    if true {
        v11 := []int{586, 128, 1, -654}
        _ = v11
    }
    if (!(!p1)) {
        var v12 int8 = 67
        _ = v12
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 992
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 0
    var v3 float64 = 1.5
    var v4 uint = 255
    var v5 float64 = v3
    v13 := GenS1{F0: (("hello" + ("" + "\n\t")) + ("\n\t" + (("\n\t" + (("a" + "\n\t") + "a")) + (("foo bar" + ("hello" + "\n\t")) + "foo bar")))), F1: (v2 + (v2 >> (((v4 / (((v4 ^ (v4 * uint(v2)))) | 1))) & 63))), F2: 32767}
    var v14 int = 32768
    var v15 uint8 = 255
    switch v14 {
    case -2147483648:
        if (true || true) {
            for fi16 := 0; fi16 < 7; fi16++ {
                if (3.14159 >= ((1e10 / (float32(3.14159) - float32(v3))) - 0.0)) {
                    if false {
                        var v17 int = v14
                        v18 := []int{int(v2), 32767, 127, (v14 / (((v14 >> (((v4 / (((v4 * (v4 & (v4 << ((v4) & 63))))) | 1))) & 63))) | 1))}
                        _ = v17
                        _ = v18
                    }
                    switch v14 {
                    case -174:
                        var v19 int64 = 0
                        _ = v19
                    case 4294967295:
                        switch (v14 >> ((v4) & 63)) {
                        case 9223372036854775807:
                            var v20 bool = (1 > 9223372036854775807)
                            _ = v20
                        default:
                            if true {
                                if true {
                                    var v21 float32 = (1e10 / 1.5)
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        v4 += (v4 / (((^(v4))) | 1))
                                        _ = fi22
                                    }
                                    _ = v21
                                }
                            } else {
                                v2 += v2
                            }
                        }
                    default:
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 int16 = 1
                            _ = fi23
                            _ = v24
                        }
                    }
                    if false {
                        v5 += 0.0
                        v4 -= v4
                        v25 := []int{(v14 + (fi16 * v14)), (fi16 % ((v14) | 1))}
                        _ = v25
                    } else {
                        if true {
                            v3 -= 3.14159
                            if true {
                                v26 := []int{32767, (fi16 - fi16), v14}
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v2 += 644
                                    var v28 float64 = v5
                                    if true {
                                        switch 128 {
                                        case 9223372036854775807:
                                            v14 = 1
                                        default:
                                            fmt.Println(v4)
                                        }
                                        v29 := GenS1{F0: "foo bar", F1: (v2 & v2), F2: -32768}
                                        _ = v29
                                    } else {
                                        if (!(!(!true))) {
                                            v4 += v4
                                            v5 = v5
                                            var v30 uint = (v4 + (v4 << (((v4 - (v4 ^ (v4 / ((1) | 1))))) & 63)))
                                            v14 = v14
                                            fmt.Println(v30)
                                            v31 := []int{(fi16 & (fi16 - v14)), (v14 % ((128) | 1))}
                                            _ = v30
                                            _ = v31
                                        } else {
                                            v2 += v2
                                            var v32 int16 = 1000
                                            v2 -= 65535
                                            _ = v32
                                        }
                                        switch -2147483648 {
                                        case 127:
                                            switch "" {
                                            case "":
                                                v3 = ((v28 + (v3 * (-2.25 + (v5 + (1e10 - (float64((float32(3.14159) - 3.14159)) - v5)))))) - v3)
                                            case "foo bar":
                                                v5 -= -2.25
                                            case "a":
                                                switch (v14 % ((-695) | 1)) {
                                                case 32768:
                                                    var v33 uint8 = 197
                                                    _ = v33
                                                default:
                                                    if (v15 < v15) {
                                                        fmt.Println(fi27)
                                                    }
                                                }
                                            }
                                            var v34 int16 = -1
                                            v35 := []int{42, fi27, (-(v14)), -1}
                                            if true {
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v3, fi27)
                                                }
                                                _ = len(v35)
                                                v4 = v4
                                            } else {
                                                v35 = append(v35, (fi16 % ((v14) | 1)))
                                            }
                                            v5 += 0.1
                                            _ = v34
                                            _ = v35
                                        case 128:
                                            if true {
                                                switch (("a" + "\n\t") + "foo bar") {
                                                case "":
                                                    v28 += 1.5
                                                case "\n\t":
                                                    var v36 string = ("\n\t" + ("foo bar" + ("" + "a")))
                                                    _ = v36
                                                }
                                                var v37 float32 = 1.5
                                                v38 := GenS1{F0: ("foo bar" + "foo bar"), F1: (v2 - v2), F2: 0}
                                                v39 := GenS2{F0: 1000}
                                                v40 := GenS3{F0: v4, F1: (((v1 << ((0) & 63)) > (v1 + -164)) || false)}
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                        case -152:
                                            v15 += (v15 + (v15 & v15))
                                        }
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            if false {
                                var v41 float64 = (v3 - (3.14159 * v3))
                                for fi42 := 0; fi42 < 1; fi42++ {
                                    var v43 int = (v14 ^ 976)
                                    v3 = (v3 * (0.0 / v3))
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                        } else {
                            if (!true) {
                                fmt.Println(fi16)
                                var v44 uint8 = 1
                                var v45 int16 = -1
                                _ = v44
                                _ = v45
                            } else {
                                for fi46 := 0; fi46 < 2; fi46++ {
                                    v1 = genFn1((v15 / ((255) | 1)), true)
                                    var v47 int = fi16
                                    _ = fi46
                                    _ = v47
                                }
                            }
                            v48 := []int{}
                            var v49 uint = v4
                            _ = v48
                            _ = v49
                        }
                        for fi50 := 0; fi50 < 6; fi50++ {
                            v51 := GenS4{F0: "hello", F1: 1, F2: v15}
                            if true {
                                if (((3.14159 - (3.14159 - (float32(v5) - -2.25))) - 3.14159) > 1.5) {
                                    fmt.Println(v3, v14, fi50)
                                } else {
                                    v14 = 4294967295
                                    v52 := []int{466, (fi50 * fi16), (fi16 * fi16)}
                                    _ = v52
                                }
                                if true {
                                    v53 := []int{v14, (fi16 % ((-760) | 1)), 0}
                                    if len(v53) >= 2 {
                                        v53 = v53[1:len(v53)]
                                    }
                                    if (1 == 100) {
                                        var v54 float64 = float64(1.5)
                                        _ = v54
                                    } else {
                                        switch fi50 {
                                        case 127:
                                            v55 := GenS5{F0: (v15 << (((^(v4))) & 63))}
                                            _ = v55
                                        case -643:
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi16, v3, v1)
                                            case "\n\t":
                                                if false {
                                                    v2 = (^(v2))
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v1, v1)
                                                    case "":
                                                        v56 := GenS3{F0: 65535, F1: true}
                                                        _ = v56
                                                    default:
                                                        if (false || false) {
                                                            v57 := GenS6{F0: (fi50 >> (((v4 << ((255) & 63))) & 63)), F1: (v4 >> (((v4 + 2)) & 63))}
                                                            var v58 uint8 = (v15 % ((v15) | 1))
                                                            _ = len(v53)
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            var v59 int32 = v1
                                                            v2 = v2
                                                            if false {
                                                                var v60 string = ""
                                                                _ = v60
                                                            } else {
                                                                if (4294967295 == 2) {
                                                                    if (2147483647 < 42) {
                                                                        if len(v53) > 0 {
                                                                            v53[0] = v14
                                                                        }
                                                                        if true {
                                                                            var v61 float64 = ((v3 - v5) / (v5 - (1.5 - v5)))
                                                                            v4 += (v4 * v4)
                                                                            _ = v61
                                                                        } else {
                                                                            v4 = 0
                                                                            _ = len(v53)
                                                                            v14 = fi50
                                                                        }
                                                                        var v62 uint32 = genFn0((v59 | (v1 * v59)), "hello", 32767)
                                                                        if len(v53) > 0 {
                                                                            v53[0] = fi50
                                                                        }
                                                                        _ = v62
                                                                    }
                                                                }
                                                                var v63 uint64 = 384
                                                                switch (fi50 ^ fi16) {
                                                                case 4294967295:
                                                                    if false {
                                                                        if true {
                                                                            v64 := GenS6{F0: (fi50 << ((42) & 63)), F1: v4}
                                                                            v1 -= v1
                                                                            var v65 uint16 = v2
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                        var v66 uint32 = 2
                                                                        v53 = append(v53, int(v2))
                                                                        _ = v66
                                                                    } else {
                                                                        fmt.Println(v59, v1, fi50)
                                                                        if (!((1.5 * float64(((0.1 + (1.5 - 0.0)) - 0.1))) != 1e10)) {
                                                                            v2 += (^(v2))
                                                                        }
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v67 bool = false
                                                                    _ = v67
                                                                case 127:
                                                                    var v68 uint32 = 524
                                                                    _ = v68
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v69 := GenS4{F0: ("\n\t" + "foo bar"), F1: -364, F2: (v15 * 60)}
                                                                    _ = v69
                                                                }
                                                                _ = v63
                                                            }
                                                            v59 = v59
                                                            _ = v59
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        var v70 int = (fi16 + 4294967295)
                                                        _ = v70
                                                    }
                                                }
                                            case "hello":
                                                v3 = (1e10 + (v5 / (v5 + v5)))
                                            }
                                        case 128:
                                            if (1 <= -477) {
                                                v53 = append(v53, (v14 - fi16))
                                            } else {
                                                var v71 uint8 = v15
                                                v3 -= v5
                                                _ = v71
                                            }
                                        default:
                                            v72 := GenS6{F0: int(v1), F1: v4}
                                            _ = v72
                                        }
                                        _ = len(v53)
                                        v3 = v3
                                    }
                                    _ = v53
                                }
                                var v73 int8 = 42
                                v1 += int32(v14)
                                v74 := GenS6{F0: 452, F1: 848}
                                _ = v73
                                _ = v74
                            }
                            var v75 float64 = (v3 * ((((-2.25 + v5) / (v5 * (float64(3.14159) - v5))) + (v3 - (v3 + v3))) + 3.14159))
                            _ = fi50
                            _ = v51
                            _ = v75
                        }
                        for fi76 := 0; fi76 < 18; fi76++ {
                            v1 -= (v1 % ((v1) | 1))
                            if (!(-1 != 45)) {
                                v77 := GenS2{F0: 0}
                                v78 := GenS3{F0: (v4 & (v4 + (v4 + 1))), F1: (!(1 < 34))}
                                var v79 string = ("a" + "hello")
                                _ = v77
                                _ = v78
                                _ = v79
                            } else {
                                v2 = v2
                            }
                            _ = fi76
                        }
                    }
                }
                _ = fi16
            }
        }
    case -915:
        v80 := GenS7{F0: false, F1: 42, F2: 18446744073709551615}
        _ = v80
    case -1:
        v2 += v2
    }
    if (int16(v14) != 1) {
        v3 = v5
    } else {
        if ((-2.25 * 3.14159) > 1.5) {
            fmt.Println(v15, v5)
            if true {
                switch (v14 & (v14 + (v14 | (v14 ^ (v14 ^ (v14 / (((v14 << ((1) & 63))) | 1))))))) {
                case 42:
                    switch 32767 {
                    case 408:
                        switch v14 {
                        case 1:
                            if (186 < 18446744073709551615) {
                                v14 = v14
                                switch "" {
                                case "foo bar":
                                    v81 := []int{(v14 / (((v14 % ((v14) | 1))) | 1)), (v14 % ((32767) | 1))}
                                    _ = v81
                                case "hello":
                                    for fi82 := 0; fi82 < 11; fi82++ {
                                        switch (^(v14)) {
                                        case -1:
                                            if false {
                                                var v83 int8 = -28
                                                var v84 string = "foo bar"
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                        _ = fi82
                                    }
                                default:
                                    v2 = (v2 << (((^(v4))) & 63))
                                }
                                if (v4 <= (v4 | 275)) {
                                    fmt.Println(v15)
                                    v85 := []int{1, (v14 % (((v14 & (v14 + (v14 ^ v14)))) | 1))}
                                    v5 += (float64(-2.25) - (v3 + v3))
                                    _ = v85
                                } else {
                                    switch int(v15) {
                                    case 128:
                                        switch v14 {
                                        case 2147483647:
                                            var v86 int64 = 158
                                            _ = v86
                                        case 4294967295:
                                            v4 = v4
                                        case -765:
                                            v87 := []int{v14, (v14 % ((v14) | 1))}
                                            _ = v87
                                        default:
                                            v88 := GenS3{F0: 42, F1: (1 >= 1)}
                                            _ = v88
                                        }
                                    case 32768:
                                        var v89 float64 = v3
                                        _ = v89
                                    default:
                                        v90 := []int{}
                                        _ = v90
                                    }
                                }
                            }
                        }
                    case -276:
                        v1 -= v1
                    default:
                        if (1.5 > float32(v5)) {
                            if ((!false) || (true || (uint32(v15) < 806))) {
                                fmt.Println(v5, v4)
                                v5 += float64(0.0)
                                for fi91 := 0; fi91 < 17; fi91++ {
                                    v92 := GenS3{F0: (v4 - (v4 % ((v4) | 1))), F1: (32767 != (v2 / ((107) | 1)))}
                                    _ = fi91
                                    _ = v92
                                }
                            } else {
                                switch v14 {
                                case 42:
                                    var v93 float64 = (float64(((-2.25 + (((1.5 + (3.14159 / 1.0)) + 1.5) / float32(float64(1.5)))) / 0.1)) + (float64(0.0) - v3))
                                    _ = v93
                                case -625:
                                    v4 -= 65535
                                case 0:
                                    v1 = v1
                                }
                                for fi94 := 0; fi94 < 3; fi94++ {
                                    var v95 int8 = -128
                                    _ = fi94
                                    _ = v95
                                }
                                v96 := []int{(v14 ^ v14), v14, 0, v14}
                                _ = v96
                            }
                        }
                    }
                case -2147483648:
                    var v97 float64 = (v3 * ((v5 + (v3 * float64(((0.1 / 1.5) / (3.14159 * (1e10 * (1.5 - 0.0))))))) + v5))
                    _ = v97
                default:
                    if (v3 == v3) {
                        var v98 float64 = (-2.25 * ((((v5 + ((0.0 + (float64(1.5) * 3.14159)) * 3.14159)) * ((float64(3.14159) + (-2.25 + v5)) - 1.5)) - v3) * v3))
                        _ = v98
                    }
                }
                v99 := []int{}
                _ = v99
            }
            v3 = ((v5 / v5) * ((1e10 / v5) + v3))
        }
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        v5 = v5
    default:
        fmt.Println(v14)
    }
    if true {
        v5 -= 0.0
        var v100 uint64 = 0
        _ = v100
    } else {
        switch "hello" {
        case "":
            if false {
                var v101 float32 = 1.5
                _ = v101
            } else {
                switch 127 {
                case 128:
                    fmt.Println(v3, v1, v15)
                case 4294967295:
                    var v102 uint32 = 937
                    _ = v102
                default:
                    for fi103 := 0; fi103 < 20; fi103++ {
                        v104 := GenS5{F0: (v15 | v15)}
                        switch 636 {
                        case 199:
                            v14 = 986
                        case -9223372036854775808:
                            v5 = 3.14159
                        case 943:
                            v1 = v1
                        default:
                            var v105 bool = true
                            _ = v105
                        }
                        _ = fi103
                        _ = v104
                    }
                }
            }
        default:
            if true {
                v14 -= (v14 & (v14 | (v14 << ((v4) & 63))))
            }
        }
        switch ("" + ("hello" + "a")) {
        case "\n\t":
            v106 := []int{}
            _ = v106
        case "hello":
            var v107 bool = false
            _ = v107
        }
        v14 = v14
    }
    v108 := []int{(v14 >> (((v4 % ((v4) | 1))) & 63)), (v14 - (^(v14))), v14}
    if len(v108) > 0 {
        v108[0] = v14
    }
    switch (v14 + v14) {
    case 370:
        if ((!false) || true) {
            if true {
                v14 = (v14 / ((v14) | 1))
                v5 += (((float64((0.0 * (-2.25 + (0.0 / (1e10 + -2.25))))) * v5) + v5) + v3)
                v109 := GenS3{F0: 164, F1: (!false)}
                _ = v109
            }
        } else {
            fmt.Println(v14)
            switch (v14 & v14) {
            case -9223372036854775808:
                v110 := GenS3{F0: 4294967295, F1: true}
                _ = v110
            case 128:
                var v111 int64 = 42
                _ = v111
            case -362:
                _ = len(v108)
            default:
                var v112 int32 = int32(v4)
                _ = v112
            }
        }
    case -9223372036854775808:
        fmt.Println(v3, v14, v5)
    default:
        if (v1 >= v1) {
            var v113 int32 = (v1 & (v1 >> ((v4) & 63)))
            var v114 int32 = 2147483647
            for fi115 := 0; fi115 < 15; fi115++ {
                v5 += (0.0 - float64(1e10))
                if ("hello" < (("" + "foo bar") + "\n\t")) {
                    for fi116 := 0; fi116 < 3; fi116++ {
                        var v117 int32 = (v114 - 2147483647)
                        switch "foo bar" {
                        case "hello":
                            v14 = (^(fi116))
                        case "a":
                            fmt.Println(v113, v114)
                        case "":
                            if (("\n\t" + (("a" + ("hello" + "\n\t")) + "a")) <= "foo bar") {
                                var v118 int8 = -128
                                switch ("a" + (("hello" + "foo bar") + "foo bar")) {
                                case "hello":
                                    fmt.Println(v113, v5, v113)
                                case "\n\t":
                                    var v119 uint32 = 2147483647
                                    _ = v119
                                default:
                                    v120 := GenS4{F0: "foo bar", F1: v114, F2: (v15 >> ((v4) & 63))}
                                    _ = v120
                                }
                                fmt.Println(v5, v4)
                                switch "" {
                                case "foo bar":
                                    if false {
                                        if true {
                                            v15 = v15
                                            if len(v108) >= 2 {
                                                v108 = v108[1:len(v108)]
                                            }
                                        }
                                    }
                                default:
                                    if (310 >= -32768) {
                                        fmt.Println(v117, v117, v14)
                                        if false {
                                            fmt.Println(v114, v114, fi115)
                                            if true {
                                                switch (("\n\t" + "a") + "a") {
                                                case "hello":
                                                    var v121 int32 = (v117 % ((genFn1(v15, (fi115 != -333))) | 1))
                                                    _ = v121
                                                case "":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v122 uint64 = 486
                                                        _ = v122
                                                    }
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "":
                                                        v123 := GenS3{F0: 4294967295, F1: true}
                                                        _ = v123
                                                    }
                                                }
                                                var v124 uint64 = 4294967295
                                                _ = v124
                                            }
                                        } else {
                                            if (((false || true) && (false || true)) || false) {
                                                var v125 float32 = 1.5
                                                _ = v125
                                            } else {
                                                v126 := GenS5{F0: 128}
                                                _ = v126
                                            }
                                            if (!((v4 + v4) >= (v4 - (^(v4))))) {
                                                switch v14 {
                                                case 2147483647:
                                                    fmt.Println(v113)
                                                case 127:
                                                    v127 := GenS8{F0: (v4 ^ v4)}
                                                    _ = v127
                                                default:
                                                    switch 49 {
                                                    case 9223372036854775807:
                                                        var v128 uint64 = 230
                                                        _ = v128
                                                    case -9223372036854775808:
                                                        v129 := GenS9{F0: (v15 + (v15 ^ v15)), F1: 56}
                                                        _ = v129
                                                    case 537:
                                                        _ = len(v108)
                                                    }
                                                }
                                                switch ("" + "\n\t") {
                                                case "foo bar":
                                                    var v130 float32 = 1.5
                                                    _ = v130
                                                default:
                                                    var v131 float32 = 0.1
                                                    _ = v131
                                                }
                                                fmt.Println(fi115)
                                            } else {
                                                if (int16(v2) > -1) {
                                                    if len(v108) > 0 {
                                                        v108[0] = (fi116 | 32768)
                                                    }
                                                    if (3.14159 > (float64(((float32(v3) - (-2.25 - 1.5)) / 1.0)) * v3)) {
                                                        if len(v108) >= 2 {
                                                            v108 = v108[1:len(v108)]
                                                        }
                                                    }
                                                }
                                                var v132 int = (fi116 & 566)
                                                _ = v132
                                            }
                                            if (true || true) {
                                                var v133 uint16 = 42
                                                if ((v2 - v133) != (^(v2))) {
                                                    v134 := GenS10{F0: 1}
                                                    _ = v134
                                                }
                                                if (9223372036854775807 != 511) {
                                                    fmt.Println(v14)
                                                    v4 = (^(v4))
                                                    var v135 uint64 = 318
                                                    _ = v135
                                                }
                                                _ = v133
                                            }
                                        }
                                    }
                                }
                                switch (v14 % ((v14) | 1)) {
                                case 9223372036854775807:
                                    if len(v108) > 0 {
                                        v108[0] = (-(fi116))
                                    }
                                case -9223372036854775808:
                                    v4 -= v4
                                default:
                                    v114 += (v1 & (v113 % ((v117) | 1)))
                                }
                                _ = v118
                            }
                        }
                        v5 = (v3 - 0.1)
                        _ = fi116
                        _ = v117
                    }
                } else {
                    v3 = v5
                    var v136 int8 = 127
                    _ = v136
                }
                switch ("\n\t" + ("foo bar" + "hello")) {
                case "foo bar":
                    for fi137 := 0; fi137 < 11; fi137++ {
                        var v138 uint64 = 42
                        v139 := GenS8{F0: 2}
                        v3 = (0.1 - (-2.25 - (0.0 + float64(1e10))))
                        _ = fi137
                        _ = v138
                        _ = v139
                    }
                case "a":
                    v4 += (v4 ^ 65535)
                case "":
                    if (!(v3 <= v5)) {
                        for fi140 := 0; fi140 < 10; fi140++ {
                            fmt.Println(fi115, v114, v1)
                            v108 = append(v108, 42)
                            _ = fi140
                        }
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            switch fi115 {
                            case -47:
                                var v141 string = ""
                                _ = v141
                            case 912:
                                var v142 uint16 = (v2 % ((v2) | 1))
                                _ = v142
                            case 310:
                                v15 -= v15
                            default:
                                _ = len(v108)
                            }
                        case 4294967295:
                            switch "" {
                            case "a":
                                v2 -= (v2 >> ((v4) & 63))
                            }
                        default:
                            fmt.Println(v1)
                        }
                        v2 = (v2 << (((^(v4))) & 63))
                    } else {
                        fmt.Println(v1)
                    }
                default:
                    v4 = v4
                }
                _ = fi115
            }
            _ = v113
            _ = v114
        }
    }
    switch -1 {
    case 127:
        if ((v1 / (((v1 | (v1 << (((v4 * 91)) & 63)))) | 1)) >= 1) {
            if (-9223372036854775808 < 1) {
                for fi143 := 0; fi143 < 11; fi143++ {
                    if true {
                        fmt.Println(v1)
                        if (v15 < v15) {
                            v144 := GenS10{F0: 2147483647}
                            var v145 int32 = 0
                            if true {
                                for fi146 := 0; fi146 < 15; fi146++ {
                                    if false {
                                        if true {
                                            fmt.Println(v3, v4)
                                            fmt.Println(v3, fi143, v1)
                                            if true {
                                                v5 += v3
                                                v108 = append(v108, 32767)
                                            }
                                            var v147 uint8 = (v15 | (^(v15)))
                                            _ = v147
                                        } else {
                                            fmt.Println(v3, v1, fi146)
                                        }
                                        switch (v14 | -49) {
                                        case 85:
                                            fmt.Println(v14, v14, v15)
                                        }
                                        var v148 string = ""
                                        v149 := GenS11{F0: 1, F1: v148, F2: genFn0((v1 * 0), ((v148 + v148) + ""), (^(v2)))}
                                        fmt.Println(v145, fi146, fi146)
                                        _ = v148
                                        _ = v149
                                    } else {
                                        var v150 int32 = int32(v2)
                                        v3 = v5
                                        _ = v150
                                    }
                                    _ = fi146
                                }
                            }
                            _ = v144
                            _ = v145
                        } else {
                            v5 = v5
                            if true {
                                v15 += (v15 * 206)
                                var v151 uint32 = 4294967295
                                _ = len(v108)
                                _ = v151
                            }
                            var v152 uint64 = 2
                            _ = v152
                        }
                    } else {
                        var v153 uint8 = 63
                        _ = v153
                    }
                    var v154 bool = false
                    v3 = 1e10
                    _ = fi143
                    _ = v154
                }
                var v155 int32 = v1
                _ = v155
            } else {
                v156 := GenS11{F0: -1, F1: ("hello" + "a"), F2: 4294967295}
                fmt.Println(v2, v14, v14)
                _ = v156
            }
            v3 -= v3
            switch ("foo bar" + "hello") {
            case "a":
                v157 := GenS9{F0: (^(v15)), F1: 807}
                _ = v157
            case "foo bar":
                for fi158 := 0; fi158 < 17; fi158++ {
                    if len(v108) >= 2 {
                        v108 = v108[1:len(v108)]
                    }
                    if false {
                        fmt.Println(v3, v14, v5)
                        for fi159 := 0; fi159 < 4; fi159++ {
                            v4 = (^(v4))
                            _ = fi159
                        }
                        var v160 int32 = v1
                        _ = v160
                    } else {
                        fmt.Println(v1, v14)
                        var v161 int = v14
                        v2 -= (v2 | v2)
                        _ = v161
                    }
                    fmt.Println(v14, fi158, v2)
                    _ = fi158
                }
            default:
                switch (("foo bar" + "") + "") {
                case "\n\t":
                    v162 := GenS10{F0: 2147483648}
                    _ = v162
                case "foo bar":
                    v163 := GenS8{F0: (v4 - (v4 << ((702) & 63)))}
                    _ = v163
                }
            }
        }
    case -299:
        v164 := GenS1{F0: "foo bar", F1: v2, F2: 818}
        _ = v164
    default:
        v165 := GenS10{F0: 500}
        _ = v165
    }
    v166 := GenS12{F0: 2147483647, F1: 42}
    fmt.Println(v3)
    fmt.Println(v15, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v108
    _ = v13
    _ = v14
    _ = v15
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
