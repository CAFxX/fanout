package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 string; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8; F1 bool; F2 uint }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 int8 }
func genFn0(p0 int16) int16 {
    var v5 uint = 825
    p0 = p0
    _ = p0
    _ = v5
    return p0
}
func genFn1() int16 {
    if (3.14159 <= (-2.25 + float32((0.0 / 1.0)))) {
        var v6 int8 = 127
        var v7 int8 = v6
        var v8 string = (("hello" + "foo bar") + "\n\t")
        _ = v6
        _ = v7
        _ = v8
    } else {
        if (true || (42 >= -604)) {
            v9 := GenS1{F0: false, F1: 18446744073709551615, F2: 498}
            if (true || (!false)) {
                switch 9223372036854775807 {
                case 1:
                    v10 := []int{32768, 281, 32768}
                    _ = v10
                case 32768:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        fmt.Println(fi11)
                        for fi12 := 0; fi12 < 17; fi12++ {
                            var v13 int16 = -1
                            _ = fi12
                            _ = v13
                        }
                        fmt.Println(fi11)
                        _ = fi11
                    }
                case 42:
                    if (false && true) {
                        if false {
                            var v14 uint16 = 42
                            v14 = (v14 % ((v14) | 1))
                            _ = v14
                        }
                        var v15 float32 = float32((((1.5 * 1e10) - 1.5) + ((1.5 - 0.0) / 1.0)))
                        var v16 int64 = 42
                        _ = v15
                        _ = v16
                    }
                default:
                    var v17 int8 = 116
                    _ = v17
                }
                v18 := GenS1{F0: (0.0 > 0.0), F1: 18446744073709551615, F2: 729}
                if (!true) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 string = "foo bar"
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS2{F0: 2147483647, F1: true, F2: 238}
                    _ = v21
                }
                _ = v18
            }
            _ = v9
        }
    }
    if ((!(!(-1 < 221))) || true) {
        var v22 float64 = -2.25
        if true {
            v23 := []int{540}
            _ = v23
        }
        fmt.Println(v22)
        _ = v22
    } else {
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS1{F0: false, F1: 4294967295, F2: 4294967295}
            var v26 string = ("foo bar" + "hello")
            _ = fi24
            _ = v25
            _ = v26
        }
        if (!true) {
            if (false && (1 != 32767)) {
                if (false || false) {
                    var v27 int16 = 32767
                    v27 = 32767
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 uint = 0
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                } else {
                    if (-128 != 1) {
                        if true {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                var v31 uint32 = 65535
                                _ = fi30
                                _ = v31
                            }
                            if (1 == 65535) {
                                var v32 int16 = 32767
                                v33 := []int{32768}
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                }
                var v34 int32 = -2147483648
                _ = v34
            }
            var v35 float64 = 3.14159
            _ = v35
        }
        v36 := []int{0}
        _ = v36
    }
    v37 := GenS2{F0: 771, F1: ((3.14159 != (((1.5 / (float32(float64(0.0)) - 3.14159)) + (1.5 - float32(1.5))) * (3.14159 - 1e10))) || true), F2: 255}
    for fi38 := 0; fi38 < 18; fi38++ {
        if (2147483647 <= fi38) {
            for fi39 := 0; fi39 < 1; fi39++ {
                var v40 uint8 = 213
                var v41 uint64 = 42
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        v42 := []int{fi38}
        _ = fi38
        _ = v42
    }
    v43 := GenS1{F0: false, F1: 2, F2: 2147483648}
    _ = v37
    _ = v43
    return genFn0(-670)
}

func main() {
    var v1 uint8 = 1
    var v2 int32 = -1
    var v3 uint64 = 42
    var v4 uint32 = 65535
    v44 := []int{1, -9223372036854775808, 2147483647}
    for fi45 := 0; fi45 < 20; fi45++ {
        if false {
            switch ("foo bar" + "foo bar") {
            case "":
                v2 += v2
            case "hello":
                if false {
                    var v46 int8 = 100
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        var v47 int8 = v46
                        _ = v47
                    }
                    _ = v46
                }
            }
        }
        fmt.Println(v1, v1)
        switch ("a" + "a") {
        case "":
            if (!((true || (!true)) && (-9223372036854775808 >= 0))) {
                for fi48 := 0; fi48 < 14; fi48++ {
                    var v49 int16 = -32768
                    v50 := GenS3{F0: -104, F1: (v3 * (v3 % ((v3) | 1)))}
                    switch ("\n\t" + ("\n\t" + (((("hello" + "hello") + "\n\t") + "\n\t") + "\n\t"))) {
                    case "a":
                        if true {
                            v51 := GenS2{F0: v4, F1: true, F2: 944}
                            fmt.Println(fi45, fi48, fi45)
                            if false {
                                v4 = (v4 * v4)
                                var v52 float64 = (0.0 + (0.0 * (0.1 + 3.14159)))
                                v1 = (v1 | (v1 * (^(v1))))
                                _ = v52
                            }
                            _ = v51
                        } else {
                            v53 := GenS1{F0: false, F1: v3, F2: v4}
                            _ = v53
                        }
                    case "foo bar":
                        if (-1 <= int32(fi48)) {
                            if ((v3 << ((617) & 63)) != 18446744073709551615) {
                                if true {
                                    if (false && (255 >= 255)) {
                                        v44 = append(v44, -791)
                                        var v54 int64 = 0
                                        if false {
                                            var v55 int = fi45
                                            _ = v55
                                        }
                                        _ = v54
                                    } else {
                                        v4 -= (v4 << ((uint(v1)) & 63))
                                        if len(v44) > 0 {
                                            v44[0] = (fi45 * int(v49))
                                        }
                                        var v56 int16 = int16(v1)
                                        _ = v56
                                    }
                                    if (true && (true && ((false && true) && (((v1 + v1) < (v1 / ((v1) | 1))) && ((^(v3)) < (v3 % (((v3 * 813)) | 1))))))) {
                                        if true {
                                            v57 := GenS1{F0: ("\n\t" < "foo bar"), F1: v3, F2: (^(v4))}
                                            _ = v57
                                        }
                                        v3 = (^(v3))
                                        if ((v49 > (^(v49))) && true) {
                                            switch fi48 {
                                            case 128:
                                                v58 := GenS4{F0: 122, F1: 1, F2: (v3 + (^(v3)))}
                                                _ = v58
                                            case 32768:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if (0.1 == (-2.25 - 1e10)) {
                                                        if len(v44) >= 2 {
                                                            v44 = v44[1:len(v44)]
                                                        }
                                                        if len(v44) > 0 {
                                                            v44[0] = fi45
                                                        }
                                                    } else {
                                                        v4 += 2147483648
                                                        v59 := GenS4{F0: uint16(v3), F1: 42, F2: 1}
                                                        v60 := GenS5{F0: 881, F1: 466}
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                case "hello":
                                                    if true {
                                                        if (!(true && true)) {
                                                            v2 -= 224
                                                            var v61 float32 = 3.14159
                                                            if true {
                                                                if len(v44) >= 2 {
                                                                    v44 = v44[1:len(v44)]
                                                                }
                                                                v1 = (v1 ^ v1)
                                                                if ("foo bar" >= "a") {
                                                                    v62 := GenS5{F0: 0, F1: 0}
                                                                    fmt.Println(fi45, v61, v2)
                                                                    _ = v62
                                                                }
                                                            }
                                                            fmt.Println(fi45, v2, v49)
                                                            fmt.Println(fi48, v3, v1)
                                                            _ = v61
                                                        }
                                                        v1 = v1
                                                        if (true && false) {
                                                            v63 := GenS4{F0: 2, F1: 2, F2: v3}
                                                            if ((v1 - v1) >= (^(v1))) {
                                                                if false {
                                                                    v4 = (v4 & v4)
                                                                } else {
                                                                    var v64 string = ("foo bar" + "hello")
                                                                    _ = v64
                                                                }
                                                                switch (fi48 % ((fi48) | 1)) {
                                                                case 9223372036854775807:
                                                                    v44 = append(v44, -1)
                                                                default:
                                                                    v65 := GenS3{F0: int8(v1), F1: 18446744073709551615}
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                v66 := GenS1{F0: (-79 >= 0), F1: 479, F2: v4}
                                                                fmt.Println(v4)
                                                                v67 := GenS4{F0: 259, F1: 882, F2: v3}
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                            _ = v63
                                                        }
                                                        var v68 int32 = 1
                                                        v69 := GenS4{F0: 255, F1: 0, F2: (v3 | (v3 | v3))}
                                                        v70 := GenS2{F0: v4, F1: false, F2: 88}
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        v4 -= uint32(v1)
                                                        v4 += v4
                                                    }
                                                    var v71 int = (fi45 | (-(fi48)))
                                                    if ((1 <= 31) && true) {
                                                        if true {
                                                            if len(v44) > 0 {
                                                                v44[0] = (fi48 | (fi45 % ((-549) | 1)))
                                                            }
                                                        } else {
                                                            v72 := GenS6{F0: "a", F1: v4}
                                                            v73 := GenS7{F0: ("hello" + ""), F1: 2147483647, F2: (v3 ^ 912)}
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        v1 = v1
                                                        fmt.Println(v71, v1)
                                                    }
                                                    _ = v71
                                                case "\n\t":
                                                    fmt.Println(v2, fi45, fi48)
                                                    v74 := GenS5{F0: 0, F1: -9223372036854775808}
                                                    _ = v74
                                                default:
                                                    if len(v44) > 0 {
                                                        v44[0] = fi48
                                                    }
                                                }
                                            case 42:
                                                var v75 float32 = 0.0
                                                _ = v75
                                            }
                                        } else {
                                            if len(v44) > 0 {
                                                v44[0] = int(v3)
                                            }
                                            var v76 int32 = 2147483647
                                            v77 := GenS3{F0: 5, F1: (v3 / ((32) | 1))}
                                            v4 += (v4 & (v4 - 65535))
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    v49 += (^(v49))
                                } else {
                                    if false {
                                        v3 -= v3
                                        switch fi45 {
                                        case -9223372036854775808:
                                            fmt.Println(v49, v4, v4)
                                        }
                                        var v78 uint16 = 896
                                        v79 := GenS3{F0: -9, F1: uint64(v1)}
                                        var v80 uint32 = uint32(v3)
                                        v81 := GenS5{F0: 42, F1: 42}
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v82 int8 = 100
                                            _ = v82
                                        case "hello":
                                            if ((^(fi45)) >= fi45) {
                                                var v83 uint64 = (^(v3))
                                                _ = v83
                                            }
                                            if len(v44) >= 2 {
                                                v44 = v44[1:len(v44)]
                                            }
                                            var v84 uint64 = (v3 / ((v3) | 1))
                                            _ = v84
                                        case "\n\t":
                                            switch "hello" {
                                            case "":
                                                var v85 uint16 = 641
                                                _ = v85
                                            case "foo bar":
                                                fmt.Println(fi48, v2, v1)
                                            case "hello":
                                                _ = len(v44)
                                            }
                                        default:
                                            if true {
                                                v1 = 2
                                                switch int(v1) {
                                                case 9223372036854775807:
                                                    if len(v44) >= 2 {
                                                        v44 = v44[1:len(v44)]
                                                    }
                                                }
                                                var v86 bool = false
                                                _ = v86
                                            } else {
                                                if true {
                                                    fmt.Println(v1)
                                                    var v87 string = ("" + "")
                                                    if (1.5 > float64((float32(0.1) - ((((3.14159 * 0.0) + 1e10) + (float32(-2.25) * ((3.14159 * (3.14159 * 1e10)) / 0.1))) * 0.0)))) {
                                                        switch (v87 + (v87 + v87)) {
                                                        case "foo bar":
                                                            switch (("hello" + v87) + "\n\t") {
                                                            case "foo bar":
                                                                if true {
                                                                    if true {
                                                                        switch (-(fi45)) {
                                                                        case 42:
                                                                            switch 344 {
                                                                            case -385:
                                                                                if false {
                                                                                    var v88 float32 = 1e10
                                                                                    v89 := GenS2{F0: 745, F1: false, F2: 0}
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                            case 4294967295:
                                                                                v4 += v4
                                                                            default:
                                                                                if (("a" + v87) > v87) {
                                                                                    v90 := GenS2{F0: v4, F1: false, F2: 510}
                                                                                    _ = v90
                                                                                }
                                                                            }
                                                                        case -2147483648:
                                                                            v91 := GenS4{F0: 255, F1: 9223372036854775807, F2: 2}
                                                                            _ = v91
                                                                        case -484:
                                                                            var v92 bool = (false || false)
                                                                            _ = v92
                                                                        }
                                                                        if (32767 != 421) {
                                                                            var v93 uint32 = (v4 << ((346) & 63))
                                                                            var v94 bool = true
                                                                            var v95 uint16 = 297
                                                                            _ = v93
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                        v96 := GenS8{F0: 100}
                                                                        _ = v96
                                                                    } else {
                                                                        v97 := GenS1{F0: true, F1: (v3 | uint64(v4)), F2: (v4 | v4)}
                                                                        _ = v97
                                                                    }
                                                                    fmt.Println(fi45, v1)
                                                                    _ = len(v44)
                                                                }
                                                            case "hello":
                                                                v3 += 42
                                                            default:
                                                                fmt.Println(v4, v2)
                                                            }
                                                        case "hello":
                                                            var v98 int64 = 9223372036854775807
                                                            _ = v98
                                                        case "a":
                                                            v99 := GenS9{F0: 42, F1: -128}
                                                            _ = v99
                                                        default:
                                                            fmt.Println(v87)
                                                        }
                                                        v44 = append(v44, 42)
                                                        _ = len(v44)
                                                    }
                                                    v1 += v1
                                                    v3 -= v3
                                                    switch v87 {
                                                    case "a":
                                                        v87 = v87
                                                    default:
                                                        v4 = (^(v4))
                                                    }
                                                    var v100 uint32 = 4294967295
                                                    var v101 int64 = 398
                                                    v87 = v87
                                                    if len(v44) > 0 {
                                                        v44[0] = 784
                                                    }
                                                    _ = v100
                                                    _ = v101
                                                    _ = v87
                                                }
                                                switch fi48 {
                                                case -1:
                                                    v102 := GenS10{F0: 411}
                                                    _ = v102
                                                case -741:
                                                    if (v2 <= v2) {
                                                        var v103 bool = (!(v2 == v2))
                                                        if ((v49 % ((genFn1()) | 1)) < -1) {
                                                            fmt.Println(v3)
                                                            var v104 float32 = 1.5
                                                            _ = v104
                                                        } else {
                                                            var v105 int64 = 794
                                                            switch (("hello" + ("\n\t" + ("" + "hello"))) + "hello") {
                                                            case "":
                                                                if false {
                                                                    v106 := GenS11{F0: 100}
                                                                    var v107 float32 = 1e10
                                                                    var v108 float32 = 1e10
                                                                    v2 = (v2 - v2)
                                                                    _ = v106
                                                                    _ = v107
                                                                    _ = v108
                                                                } else {
                                                                    v109 := GenS12{F0: 832, F1: (v105 * (v105 / (((-(v105))) | 1))), F2: 1}
                                                                    var v110 int8 = -42
                                                                    if true {
                                                                        v49 += (v49 >> ((65535) & 63))
                                                                        var v111 uint16 = 1
                                                                        fmt.Println(v4)
                                                                        _ = v111
                                                                    }
                                                                    fmt.Println(v105, fi48)
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                            case "a":
                                                                v112 := GenS13{F0: (v4 - (^(v4)))}
                                                                _ = v112
                                                            }
                                                            v113 := GenS10{F0: 0}
                                                            fmt.Println(fi48, fi48, fi48)
                                                            switch "hello" {
                                                            case "hello":
                                                                if len(v44) > 0 {
                                                                    v44[0] = fi45
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(fi45, v3)
                                                            case "":
                                                                v103 = v103
                                                            }
                                                            _ = v105
                                                            _ = v113
                                                        }
                                                        v1 += v1
                                                        _ = v103
                                                    } else {
                                                        v44 = append(v44, fi48)
                                                        v4 = 4294967295
                                                        v2 -= v2
                                                    }
                                                }
                                                switch int(v2) {
                                                case 1:
                                                    var v114 string = ("foo bar" + ("\n\t" + ("foo bar" + "a")))
                                                    _ = v114
                                                case 32768:
                                                    switch "" {
                                                    case "":
                                                        if false {
                                                            switch "hello" {
                                                            case "hello":
                                                                if len(v44) > 0 {
                                                                    v44[0] = fi48
                                                                }
                                                            case "\n\t":
                                                                v2 = (v2 % ((v2) | 1))
                                                            case "foo bar":
                                                                if true {
                                                                    var v115 uint32 = (v4 ^ 1)
                                                                    _ = v115
                                                                }
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                            v3 = v3
                                                        } else {
                                                            v1 = (v1 << ((392) & 63))
                                                            switch (fi45 / ((fi45) | 1)) {
                                                            case -965:
                                                                switch "foo bar" {
                                                                case "":
                                                                    if (!true) {
                                                                        switch (("hello" + ("hello" + "")) + "foo bar") {
                                                                        case "":
                                                                            var v116 int8 = 127
                                                                            _ = v116
                                                                        case "hello":
                                                                            if true {
                                                                                switch -43 {
                                                                                case -2147483648:
                                                                                    v117 := GenS14{F0: false, F1: int64(v3)}
                                                                                    _ = v117
                                                                                case 57:
                                                                                    v4 += (v4 * 4294967295)
                                                                                case 0:
                                                                                    if len(v44) > 0 {
                                                                                        v44[0] = fi45
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    if (v49 < (v49 >> ((1) & 63))) {
                                                                                        var v118 int = -595
                                                                                        fmt.Println(fi45, v49)
                                                                                        _ = v118
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v3 -= (v3 ^ (v3 | 42))
                                                                                }
                                                                            } else {
                                                                                v1 += 22
                                                                            }
                                                                        case "foo bar":
                                                                            v119 := GenS3{F0: 42, F1: v3}
                                                                            _ = v119
                                                                        default:
                                                                            if len(v44) >= 2 {
                                                                                v44 = v44[1:len(v44)]
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case 128:
                                                                v49 += v49
                                                            default:
                                                                v120 := GenS15{F0: v3}
                                                                _ = v120
                                                            }
                                                            v2 += (v2 - v2)
                                                            var v121 uint = 4294967295
                                                            var v122 float64 = 1e10
                                                            var v123 bool = (255 != 65535)
                                                            _ = v121
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v2, v1, v49)
                                                    case "a":
                                                        v3 -= (v3 | 42)
                                                    default:
                                                        switch 4294967295 {
                                                        case 42:
                                                            var v124 int16 = genFn1()
                                                            if len(v44) > 0 {
                                                                v44[0] = (^(fi45))
                                                            }
                                                            if true {
                                                                v125 := GenS16{F0: 67, F1: false, F2: 1}
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v126 uint64 = (v3 | uint64(v4))
                                                                    _ = v126
                                                                case "":
                                                                    if (!true) {
                                                                        if len(v44) >= 2 {
                                                                            v44 = v44[1:len(v44)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v125
                                                            }
                                                            _ = v124
                                                        default:
                                                            v127 := GenS7{F0: ("\n\t" + "hello"), F1: (v4 + 260), F2: v3}
                                                            _ = v127
                                                        }
                                                    }
                                                case -2147483648:
                                                    fmt.Println(v49, v1)
                                                default:
                                                    if (750 != (v2 % (((v2 ^ v2)) | 1))) {
                                                        v3 = v3
                                                        v128 := GenS15{F0: (v3 ^ v3)}
                                                        if true {
                                                            if false {
                                                                v3 += v3
                                                                if len(v44) >= 2 {
                                                                    v44 = v44[1:len(v44)]
                                                                }
                                                                var v129 int32 = (v2 | (v2 + 874))
                                                                _ = v129
                                                            }
                                                            var v130 bool = false
                                                            _ = v130
                                                        }
                                                        _ = v128
                                                    } else {
                                                        var v131 uint64 = 446
                                                        v4 += v4
                                                        var v132 bool = false
                                                        _ = v131
                                                        _ = v132
                                                    }
                                                    switch fi45 {
                                                    case -276:
                                                        fmt.Println(v49, fi48)
                                                    case -1:
                                                        if true {
                                                            if true {
                                                                v3 += 42
                                                            } else {
                                                                var v133 float64 = 1e10
                                                                if false {
                                                                    if len(v44) > 0 {
                                                                        v44[0] = fi48
                                                                    }
                                                                    fmt.Println(v1, v4, v49)
                                                                }
                                                                fmt.Println(v3)
                                                                _ = v133
                                                            }
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(fi48, v3, v2)
                                                            case "hello":
                                                                switch (^(fi48)) {
                                                                case 4294967295:
                                                                    switch int(v2) {
                                                                    case -5:
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            switch fi48 {
                                                                            case 2147483647:
                                                                                v49 -= v49
                                                                            case -788:
                                                                                v2 = (v2 % ((v2) | 1))
                                                                            case 253:
                                                                                v3 -= (v3 / ((uint64(v4)) | 1))
                                                                            default:
                                                                                v3 = (v3 + (v3 + 981))
                                                                            }
                                                                        case "\n\t":
                                                                            var v134 uint64 = 42
                                                                            _ = v134
                                                                        case "":
                                                                            v135 := GenS17{F0: 127}
                                                                            if len(v44) >= 2 {
                                                                                v44 = v44[1:len(v44)]
                                                                            }
                                                                            fmt.Println(v1)
                                                                            _ = v135
                                                                        default:
                                                                            v3 = 778
                                                                        }
                                                                    case -2147483648:
                                                                        _ = len(v44)
                                                                    default:
                                                                        if (genFn1() != genFn0((v49 >> ((2) & 63)))) {
                                                                            var v136 uint8 = 128
                                                                            if len(v44) > 0 {
                                                                                v44[0] = fi45
                                                                            }
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                var v137 int16 = v49
                                                                                _ = v137
                                                                            case "":
                                                                                v49 = 0
                                                                            case "hello":
                                                                                if false {
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        var v138 float32 = 1.5
                                                                                        _ = v138
                                                                                    default:
                                                                                        v3 = (v3 & (v3 * (v3 - (v3 * (^(v3))))))
                                                                                    }
                                                                                    v49 += (v49 | (v49 / (((v49 & v49)) | 1)))
                                                                                    if true {
                                                                                        fmt.Println(fi45)
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if ((0.0 > ((0.0 + (1.5 + -2.25)) - (((float32((1.5 - 1e10)) * float32((0.0 - ((float64(3.14159) * 0.0) + 0.0)))) - 1e10) + 0.1))) && false) {
                                                                                    var v139 float32 = (((-2.25 / float32(0.1)) / 0.1) * 3.14159)
                                                                                    switch fi45 {
                                                                                    case 129:
                                                                                        v140 := GenS9{F0: v136, F1: 100}
                                                                                        _ = v140
                                                                                    case 4294967295:
                                                                                        v141 := GenS18{F0: 1}
                                                                                        _ = v141
                                                                                    case -9223372036854775808:
                                                                                        v136 += v136
                                                                                    default:
                                                                                        if (2 > 2) {
                                                                                            v1 = 42
                                                                                            v1 = v136
                                                                                            v44 = append(v44, int(v136))
                                                                                        }
                                                                                        if false {
                                                                                            v4 = (v4 / ((v4) | 1))
                                                                                            v3 = v3
                                                                                        } else {
                                                                                            switch ("" + (("\n\t" + (("hello" + ("hello" + "\n\t")) + "")) + (((((("\n\t" + "a") + "foo bar") + "a") + (("\n\t" + "hello") + "")) + "") + "hello"))) {
                                                                                            case "foo bar":
                                                                                                var v142 int32 = 2147483647
                                                                                                _ = v142
                                                                                            case "\n\t":
                                                                                                v4 = v4
                                                                                            case "":
                                                                                                v4 += v4
                                                                                            default:
                                                                                                v139 = 0.0
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v49 = (^(v49))
                                                                                    switch fi48 {
                                                                                    case 950:
                                                                                        v143 := GenS19{F0: "foo bar"}
                                                                                        _ = v143
                                                                                    case 128:
                                                                                        fmt.Println(v136, v136, v3)
                                                                                    }
                                                                                    _ = v139
                                                                                } else {
                                                                                    var v144 int64 = -970
                                                                                    _ = v144
                                                                                }
                                                                            }
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                default:
                                                                    if true {
                                                                        if (fi45 != (fi48 / ((fi45) | 1))) {
                                                                            if (v3 >= v3) {
                                                                                if (32767 == 416) {
                                                                                    switch (fi45 - (fi48 ^ (fi48 << ((uint(v1)) & 63)))) {
                                                                                    case 541:
                                                                                        var v145 bool = (65535 != 255)
                                                                                        _ = v145
                                                                                    case 681:
                                                                                        v146 := GenS6{F0: "hello", F1: 2}
                                                                                        _ = v146
                                                                                    case 363:
                                                                                        var v147 string = "hello"
                                                                                        _ = v147
                                                                                    default:
                                                                                        var v148 int = 4294967295
                                                                                        _ = v148
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        var v149 bool = (fi48 == (fi45 * 32768))
                                                                                        _ = v149
                                                                                    } else {
                                                                                        v1 += v1
                                                                                        var v150 float32 = (0.1 - ((0.0 / (float32(1.5) / 1e10)) + 1.5))
                                                                                        _ = v150
                                                                                    }
                                                                                    var v151 string = "a"
                                                                                    fmt.Println(fi48, v151)
                                                                                    if (false && true) {
                                                                                        var v152 int64 = -890
                                                                                        var v153 int8 = -7
                                                                                        v154 := GenS20{F0: int(v2)}
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                        _ = v154
                                                                                    } else {
                                                                                        var v155 int8 = 0
                                                                                        _ = v155
                                                                                    }
                                                                                    if false {
                                                                                        v2 = 42
                                                                                        v49 -= v49
                                                                                        var v156 bool = (float64(1e10) <= (1e10 * float64(float32(3.14159))))
                                                                                        _ = v156
                                                                                    } else {
                                                                                        var v157 uint8 = (v1 * 0)
                                                                                        var v158 int32 = int32(v49)
                                                                                        _ = v157
                                                                                        _ = v158
                                                                                    }
                                                                                    _ = v151
                                                                                }
                                                                            }
                                                                            v3 -= (v3 << ((0) & 63))
                                                                        }
                                                                    } else {
                                                                        var v159 int16 = (v49 / (((v49 % ((v49) | 1))) | 1))
                                                                        var v160 uint32 = (v4 * (v4 % ((v4) | 1)))
                                                                        _ = v159
                                                                        _ = v160
                                                                    }
                                                                }
                                                            default:
                                                                if len(v44) >= 2 {
                                                                    v44 = v44[1:len(v44)]
                                                                }
                                                            }
                                                            var v161 int = 350
                                                            _ = v161
                                                        } else {
                                                            fmt.Println(v49)
                                                            var v162 int32 = 0
                                                            var v163 int32 = v2
                                                            v1 = 60
                                                            _ = v162
                                                            _ = v163
                                                        }
                                                    }
                                                    switch (-(fi45)) {
                                                    case -2147483648:
                                                        switch (-(fi48)) {
                                                        case 2147483647:
                                                            v1 = v1
                                                        case 32767:
                                                            if len(v44) > 0 {
                                                                v44[0] = 4294967295
                                                            }
                                                        case 4294967295:
                                                            var v164 uint64 = 4294967295
                                                            _ = v164
                                                        default:
                                                            v4 = (v4 - (v4 + v4))
                                                        }
                                                    case -9223372036854775808:
                                                        v44 = append(v44, fi48)
                                                    case 32768:
                                                        if len(v44) > 0 {
                                                            v44[0] = 128
                                                        }
                                                        v2 = 1
                                                        var v165 uint64 = v3
                                                        _ = v165
                                                    }
                                                }
                                                fmt.Println(fi45, v2, v4)
                                            }
                                        }
                                        var v166 int8 = 127
                                        _ = v166
                                    }
                                }
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                            }
                            v3 -= (v3 | v3)
                        }
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                switch "hello" {
                case "hello":
                    var v167 float32 = 1e10
                    _ = v167
                case "\n\t":
                    var v168 int64 = 42
                    _ = v168
                case "":
                    if false {
                        v2 += (v2 * v2)
                        v1 = v1
                    } else {
                        switch fi45 {
                        case 128:
                            for fi169 := 0; fi169 < 19; fi169++ {
                                var v170 float32 = (0.0 - 0.1)
                                v3 -= v3
                                _ = fi169
                                _ = v170
                            }
                        case -346:
                            v4 = 2147483647
                        case -60:
                            fmt.Println(v4, v3)
                        }
                        var v171 int64 = 0
                        _ = v171
                    }
                }
            }
        case "a":
            var v172 int8 = 1
            _ = v172
        case "foo bar":
            switch "" {
            case "foo bar":
                switch fi45 {
                case -9223372036854775808:
                    v173 := GenS13{F0: (^(v4))}
                    _ = v173
                case -2147483648:
                    if len(v44) > 0 {
                        v44[0] = 284
                    }
                }
            case "":
                switch "a" {
                case "a":
                    v174 := GenS14{F0: (true && (176 == 563)), F1: 42}
                    _ = v174
                case "":
                    v1 += uint8(v3)
                default:
                    if len(v44) > 0 {
                        v44[0] = (^(fi45))
                    }
                }
            }
        default:
            var v175 bool = (!(false || true))
            _ = v175
        }
        _ = fi45
    }
    var v176 int32 = v2
    v1 += v1
    v177 := GenS17{F0: -62}
    v2 += v176
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    var v178 uint32 = 2147483648
    for fi179 := 0; fi179 < 1; fi179++ {
        var v180 int64 = -850
        switch (fi179 << ((674) & 63)) {
        case -1:
            for fi181 := 0; fi181 < 7; fi181++ {
                if ((42 > (v3 % (((v3 << ((255) & 63))) | 1))) && true) {
                    var v182 uint32 = v4
                    if true {
                        v183 := GenS15{F0: 2}
                        v176 = 1
                        v44 = append(v44, -1)
                        _ = v183
                    } else {
                        var v184 uint16 = 32767
                        v2 = v176
                        _ = v184
                    }
                    if true {
                        switch fi179 {
                        case 2147483647:
                            if false {
                                var v185 uint64 = v3
                                _ = v185
                            }
                        case 9223372036854775807:
                            if ("hello" < "\n\t") {
                                var v186 int16 = genFn1()
                                switch 4294967295 {
                                case -9223372036854775808:
                                    var v187 int32 = -989
                                    _ = v187
                                default:
                                    fmt.Println(v4, v180, v176)
                                }
                                _ = len(v44)
                                _ = v186
                            }
                        default:
                            fmt.Println(v3, fi181, v1)
                        }
                    }
                    _ = v182
                }
                var v188 uint64 = 4294967295
                var v189 bool = (v4 <= (v4 - v4))
                _ = fi181
                _ = v188
                _ = v189
            }
        case 42:
            if ("hello" != "hello") {
                for fi190 := 0; fi190 < 10; fi190++ {
                    fmt.Println(v4, v4)
                    _ = fi190
                }
            }
        }
        for fi191 := 0; fi191 < 1; fi191++ {
            if (v180 >= 1) {
                var v192 bool = false
                v180 += v180
                v3 += (v3 >> ((255) & 63))
                _ = v192
            }
            fmt.Println(v178, v4)
            _ = fi191
        }
        _ = fi179
        _ = v180
    }
    switch 0 {
    case 32767:
        v193 := GenS10{F0: 860}
        _ = v193
    case 0:
        v176 += (^(v2))
    default:
        if false {
            if (0 < 9223372036854775807) {
                var v194 int64 = -638
                _ = v194
            }
            var v195 bool = (!true)
            switch "hello" {
            case "foo bar":
                v178 = 435
            case "hello":
                if ((v195 && (!(true && (4294967295 < v3)))) || ((-258 >= 9223372036854775807) && (false && ((42 <= 359) && false)))) {
                    fmt.Println(v4, v195)
                    v196 := GenS20{F0: -64}
                    _ = v196
                }
            case "":
                if v195 {
                    var v197 uint8 = v1
                    _ = v197
                } else {
                    var v198 int16 = -131
                    fmt.Println(v2, v198)
                    var v199 float64 = 1e10
                    _ = v198
                    _ = v199
                }
            default:
                if len(v44) > 0 {
                    v44[0] = -607
                }
            }
            _ = v195
        }
    }
    v1 = 118
    if false {
        var v200 uint32 = (v4 - v4)
        _ = v200
    }
    if true {
        for fi201 := 0; fi201 < 19; fi201++ {
            switch "hello" {
            case "foo bar":
                v202 := GenS6{F0: ("\n\t" + "a"), F1: v4}
                _ = v202
            case "":
                if (42 == 0) {
                    var v203 float32 = 0.1
                    _ = v203
                } else {
                    if (32767 > 773) {
                        v204 := GenS6{F0: "hello", F1: 42}
                        v4 = 4294967295
                        v178 += (^(v4))
                        _ = v204
                    } else {
                        var v205 float64 = 3.14159
                        var v206 float64 = v205
                        fmt.Println(v178, v205)
                        _ = v205
                        _ = v206
                    }
                    var v207 uint16 = 0
                    _ = v207
                }
            default:
                if (!(3.14159 > (0.1 - ((((-2.25 * 3.14159) / 1.5) - 1.5) * (((1e10 - (float64(((-2.25 + ((3.14159 - 1e10) + (1.5 + 0.1))) * float32(1.5))) - ((0.0 - 1e10) - -2.25))) * ((1.5 - (1.5 + 1.5)) + 1e10)) - ((-2.25 + 3.14159) - 1.5)))))) {
                    if (("hello" + ("\n\t" + "hello")) == "") {
                        v2 += 0
                        for fi208 := 0; fi208 < 5; fi208++ {
                            v1 -= uint8(v178)
                            _ = fi208
                        }
                    } else {
                        for fi209 := 0; fi209 < 13; fi209++ {
                            var v210 uint = 203
                            v1 = (v1 + (^(v1)))
                            switch "foo bar" {
                            case "\n\t":
                                var v211 uint16 = 0
                                _ = v211
                            }
                            _ = fi209
                            _ = v210
                        }
                        if (genFn1() != 1000) {
                            var v212 float64 = float64(((0.0 - (0.0 + -2.25)) - -2.25))
                            fmt.Println(v178, v1)
                            if true {
                                for fi213 := 0; fi213 < 9; fi213++ {
                                    fmt.Println(v2)
                                    if false {
                                        var v214 string = ("hello" + "\n\t")
                                        v1 = 210
                                        _ = v214
                                    }
                                    switch fi213 {
                                    case 9223372036854775807:
                                        fmt.Println(v4, v4)
                                    default:
                                        if false {
                                            if true {
                                                if len(v44) > 0 {
                                                    v44[0] = (fi201 + fi201)
                                                }
                                                var v215 uint64 = v3
                                                v176 = v2
                                                _ = v215
                                            } else {
                                                v216 := GenS8{F0: -1}
                                                v178 = (v4 ^ 2147483647)
                                                var v217 int16 = 681
                                                _ = v216
                                                _ = v217
                                            }
                                        }
                                    }
                                    var v218 string = "\n\t"
                                    if true {
                                        if false {
                                            switch v218 {
                                            case "\n\t":
                                                var v219 string = (v218 + ((v218 + v218) + (v218 + ((((v218 + "a") + (v218 + v218)) + "hello") + v218))))
                                                _ = v219
                                            default:
                                                v220 := GenS8{F0: 100}
                                                _ = v220
                                            }
                                            var v221 int64 = -1
                                            var v222 float64 = (0.0 + v212)
                                            _ = v221
                                            _ = v222
                                        }
                                    } else {
                                        var v223 int16 = 1000
                                        fmt.Println(v3)
                                        _ = v223
                                    }
                                    var v224 int16 = -719
                                    v225 := GenS21{F0: 18446744073709551615}
                                    _ = fi213
                                    _ = v218
                                    _ = v224
                                    _ = v225
                                }
                                var v226 int64 = 9223372036854775807
                                _ = v226
                            }
                            _ = v212
                        }
                        for fi227 := 0; fi227 < 18; fi227++ {
                            if (true || (-19 > genFn0(42))) {
                                switch fi227 {
                                case -1:
                                    v228 := GenS3{F0: -128, F1: v3}
                                    _ = v228
                                case 178:
                                    var v229 uint32 = (^(v178))
                                    _ = v229
                                default:
                                    switch (fi227 << ((4294967295) & 63)) {
                                    case 9223372036854775807:
                                        v230 := GenS7{F0: ("a" + "a"), F1: v4, F2: 298}
                                        _ = v230
                                    case 829:
                                        v2 = v176
                                    case 32768:
                                        v231 := GenS22{F0: false}
                                        _ = v231
                                    }
                                }
                            }
                            if true {
                                v232 := GenS8{F0: 23}
                                _ = v232
                            }
                            _ = fi227
                        }
                    }
                    if false {
                        var v233 uint64 = 1
                        switch "" {
                        case "":
                            if len(v44) > 0 {
                                v44[0] = (fi201 & (fi201 ^ fi201))
                            }
                        }
                        v176 = v176
                        _ = v233
                    } else {
                        if false {
                            v3 -= v3
                            v234 := GenS14{F0: false, F1: -1}
                            var v235 uint16 = 699
                            _ = v234
                            _ = v235
                        }
                        v3 = 18446744073709551615
                    }
                } else {
                    v3 -= uint64(v1)
                    v3 += v3
                    v178 = (v4 & (v4 ^ (^(v4))))
                }
            }
            fmt.Println(v176, v2, v2)
            v236 := GenS23{F0: 0}
            _ = fi201
            _ = v236
        }
        for fi237 := 0; fi237 < 9; fi237++ {
            v238 := GenS24{F0: (uint16(v1) == 763)}
            _ = fi237
            _ = v238
        }
    }
    v176 = (v176 & 1)
    v2 = (v176 << ((362) & 63))
    v239 := GenS25{F0: -1}
    var v240 int16 = genFn0(-1)
    if len(v44) > 0 {
        v44[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("v178:", v178)
    fmt.Println("v2:", v2)
    fmt.Println("v240:", v240)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v176
    _ = v177
    _ = v178
    _ = v2
    _ = v239
    _ = v240
    _ = v3
    _ = v4
    _ = v44
}
