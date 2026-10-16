package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int; F1 uint16; F2 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64; F1 int64; F2 int8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0() string {
    switch 1 {
    case -2147483648:
        v6 := []int{-344, -2147483648, 32768, 937}
        _ = v6
    case 0:
        v7 := []int{497, 0}
        _ = v7
    case -780:
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint = 255
            v10 := GenS1{F0: -650, F1: true, F2: 916}
            for fi11 := 0; fi11 < 14; fi11++ {
                switch fi11 {
                case -1:
                    if true {
                        var v12 int64 = 0
                        _ = v12
                    }
                case 891:
                    var v13 uint = 0
                    _ = v13
                case 88:
                    var v14 uint32 = 0
                    _ = v14
                default:
                    v9 += (v9 % (((v9 - 1)) | 1))
                }
                _ = fi11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        if true {
            var v15 float32 = ((1e10 * (-2.25 / -2.25)) * (0.1 / (-2.25 - 1e10)))
            v16 := []int{}
            switch (("hello" + "\n\t") + "a") {
            case "":
                fmt.Println(v15)
            case "a":
                for fi17 := 0; fi17 < 3; fi17++ {
                    fmt.Println(v15, v15)
                    _ = fi17
                }
            case "\n\t":
                for fi18 := 0; fi18 < 16; fi18++ {
                    if (-32768 > -1) {
                        v15 += (3.14159 - 0.1)
                    }
                    _ = fi18
                }
            default:
                v19 := GenS2{F0: 42, F1: "hello"}
                _ = v19
            }
            _ = v15
            _ = v16
        }
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int64 = -1
            _ = fi21
            _ = v22
        }
        var v23 bool = false
        _ = fi20
        _ = v23
    }
    v24 := []int{42}
    _ = v24
    return "\n\t"
}
func genFn1(p0 int16) bool {
    var v25 uint64 = 514
    var v26 uint = 65535
    v27 := []int{}
    switch ("a" + "") {
    case "foo bar":
        v27 = append(v27, -486)
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            v26 -= v26
            if false {
                if ((^(v26)) == (v26 % ((v26) | 1))) {
                    v26 -= v26
                    p0 += p0
                    if (-1 >= 0) {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            if false {
                                _ = len(v27)
                                fmt.Println(fi28, p0, fi28)
                            } else {
                                fmt.Println(fi28, v25)
                                if false {
                                    v30 := GenS1{F0: 2147483647, F1: (((fi28 ^ fi28) == (^(fi28))) || false), F2: 0}
                                    if false {
                                        if len(v27) > 0 {
                                            v27[0] = fi28
                                        }
                                    }
                                    var v31 float32 = (0.0 * -2.25)
                                    _ = v30
                                    _ = v31
                                }
                            }
                            _ = fi29
                        }
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v26 += v26
                            _ = fi32
                        }
                    } else {
                        var v33 int64 = 20
                        _ = v33
                    }
                } else {
                    v34 := GenS2{F0: -47, F1: "a"}
                    _ = v34
                }
                _ = len(v27)
                for fi35 := 0; fi35 < 17; fi35++ {
                    _ = len(v27)
                    if (-2.25 >= 1e10) {
                        v25 = (v25 % (((v25 / ((uint64(v26)) | 1))) | 1))
                    }
                    v36 := GenS1{F0: -1, F1: true, F2: 2147483648}
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi28
        }
    }
    _ = p0
    _ = v25
    _ = v26
    _ = v27
    return false
}

func main() {
    var v1 bool = (true && ((!false) || (150 >= 223)))
    var v2 float64 = ((1e10 + (0.1 * ((-2.25 - float64(0.0)) + -2.25))) * -2.25)
    var v3 int64 = 42
    var v4 int = 32767
    v5 := make(map[uint32]bool)
    v1 = false
    var v37 string = "hello"
    v1 = v1
    fmt.Println(v3)
    fmt.Println(v4, v2, v1)
    v37 = "a"
    if ((v1 && (v37 != (v37 + v37))) && (((!((695 > 1) && false)) || (!(!(-47 <= 0)))) || (0 != 2))) {
        for fi38 := 0; fi38 < 12; fi38++ {
            if v1 {
                fmt.Println(v1)
            }
            var v39 bool = v1
            var v40 string = "\n\t"
            _ = fi38
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 2; fi41++ {
            switch ((v37 + "foo bar") + "") {
            case "foo bar":
                fmt.Println(v2, v2)
            default:
                for fi42 := 0; fi42 < 16; fi42++ {
                    var v43 string = "a"
                    _ = fi42
                    _ = v43
                }
            }
            if (302 == 0) {
                var v44 int16 = 32767
                v1 = v1
                for fi45 := 0; fi45 < 10; fi45++ {
                    fmt.Println(v1)
                    var v46 bool = (v1 && v1)
                    v44 = v44
                    _ = fi45
                    _ = v46
                }
                _ = v44
            } else {
                var v47 uint32 = 2147483648
                switch (v4 ^ 32767) {
                case 32768:
                    var v48 bool = v1
                    _ = v48
                case 1:
                    if v1 {
                        if v1 {
                            var v49 int32 = -1
                            fmt.Println(v49)
                            _ = v49
                        } else {
                            var v50 uint16 = 452
                            for fi51 := 0; fi51 < 10; fi51++ {
                                v1 = (v2 > v2)
                                _ = fi51
                            }
                            _ = v50
                        }
                        v2 = (1.5 + (1e10 - v2))
                    }
                }
                var v52 int8 = -1
                _ = v47
                _ = v52
            }
            _ = fi41
        }
    } else {
        v53 := GenS3{F0: 219, F1: 32767}
        if (v2 == 3.14159) {
            _ = len(v5)
            v54 := []int{v4, (v4 << ((0) & 63)), v4, (-(v4))}
            _ = v54
        } else {
            v3 = 404
        }
        _ = v53
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        var v56 uint64 = 372
        v4 -= 0
        if (v3 != 516) {
            v37 = v37
        } else {
            if (!false) {
                if true {
                    v57 := []int{v4}
                    for fi58 := 0; fi58 < 12; fi58++ {
                        v59 := GenS4{F0: 42, F1: (-382 == -2147483648), F2: 15}
                        if (false && false) {
                            var v60 int64 = -170
                            var v61 uint64 = v56
                            _ = v60
                            _ = v61
                        } else {
                            switch v37 {
                            case "":
                                switch fi58 {
                                case 9223372036854775807:
                                    switch "" {
                                    case "a":
                                        if false {
                                            fmt.Println(v2, fi58)
                                            switch ((("a" + v37) + v37) + ("a" + "a")) {
                                            case "hello":
                                                if false {
                                                    if v1 {
                                                        var v62 uint64 = 42
                                                        if (0.1 > (3.14159 * -2.25)) {
                                                            if v1 {
                                                                var v63 uint64 = v62
                                                                fmt.Println(v37)
                                                                v5[42] = v1
                                                                _ = v63
                                                            } else {
                                                                v4 -= fi58
                                                                if v1 {
                                                                    var v64 uint = 608
                                                                    fmt.Println(fi55)
                                                                    v57 = append(v57, v4)
                                                                    _ = v64
                                                                } else {
                                                                    v56 = v62
                                                                    fmt.Println(fi55, fi58, v37)
                                                                    if (((42 < 183) && (1 < -113)) || false) {
                                                                        fmt.Println(v37, v4)
                                                                    } else {
                                                                        switch v37 {
                                                                        case "":
                                                                            var v65 float32 = 1e10
                                                                            _ = v65
                                                                        }
                                                                        fmt.Println(v4, v56)
                                                                        var v66 int16 = -1
                                                                        _ = v66
                                                                    }
                                                                }
                                                                v5[2] = true
                                                            }
                                                            var v67 uint16 = 255
                                                            v5[42] = v1
                                                            v4 = fi55
                                                            _ = v67
                                                        }
                                                        v37 = ""
                                                        _ = v62
                                                    }
                                                    v68 := GenS1{F0: 17, F1: (true || ((v56 << ((42) & 63)) != (v56 - (v56 << ((4294967295) & 63))))), F2: 2147483647}
                                                    _ = len(v5)
                                                    _ = len(v5)
                                                    _ = v68
                                                } else {
                                                    var v69 int = (fi58 - (fi58 << ((0) & 63)))
                                                    _ = v69
                                                }
                                            case "foo bar":
                                                if true {
                                                    v3 = (^(v3))
                                                    v56 += 42
                                                } else {
                                                    v70 := v5[991]
                                                    v3 = 169
                                                    var v71 uint8 = 128
                                                    _ = v70
                                                    _ = v71
                                                }
                                            case "a":
                                                var v72 float64 = float64((((3.14159 / 1e10) - (float32(v2) * (1.5 + (float32(-2.25) / 1e10)))) * ((((0.0 / (-2.25 - (((0.1 - 3.14159) - ((3.14159 / (float32(0.1) * 0.1)) * 1.5)) - (3.14159 - -2.25)))) * ((((3.14159 + 0.0) / -2.25) * -2.25) / (1e10 - ((0.1 - (float32(v2) / 3.14159)) - -2.25)))) + 1.5) / float32(-2.25))))
                                                _ = v72
                                            }
                                            var v73 uint32 = 332
                                            _ = v73
                                        }
                                    case "":
                                        var v74 int64 = (v3 & (v3 / (((v3 + 0)) | 1)))
                                        _ = v74
                                    default:
                                        var v75 uint64 = 2
                                        v56 = 1
                                        v37 = v37
                                        _ = v75
                                    }
                                case 4294967295:
                                    var v76 int8 = 100
                                    _ = v76
                                case 1:
                                    _ = len(v5)
                                }
                                fmt.Println(v2, v1, v2)
                                if v1 {
                                    var v77 int64 = 0
                                    var v78 int64 = (^(v3))
                                    if false {
                                        switch 128 {
                                        case 3:
                                            switch fi55 {
                                            case 42:
                                                if (1e10 > -2.25) {
                                                    _ = len(v5)
                                                    switch v37 {
                                                    case "\n\t":
                                                        v57 = append(v57, -830)
                                                    default:
                                                        if v1 {
                                                            fmt.Println(v3, v77, fi58)
                                                        } else {
                                                            v3 = -992
                                                            if v1 {
                                                                switch v37 {
                                                                case "a":
                                                                    switch fi58 {
                                                                    case -947:
                                                                        v2 = 0.1
                                                                    case -2147483648:
                                                                        v1 = false
                                                                    case 32768:
                                                                        v79 := GenS5{F0: -19}
                                                                        _ = v79
                                                                    default:
                                                                        switch (fi55 + (v4 >> ((581) & 63))) {
                                                                        case 4294967295:
                                                                            switch (v4 % (((fi58 << ((683) & 63))) | 1)) {
                                                                            case -997:
                                                                                if (472 < 2) {
                                                                                    if v1 {
                                                                                        v80 := GenS6{F0: -338, F1: 149}
                                                                                        v81 := GenS7{F0: -52, F1: 901}
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if len(v57) > 0 {
                                                                                                v57[0] = (^(fi55))
                                                                                            }
                                                                                        default:
                                                                                            v82 := GenS1{F0: -561, F1: (536 != 255), F2: 699}
                                                                                            if len(v57) >= 2 {
                                                                                                v57 = v57[1:len(v57)]
                                                                                            }
                                                                                            _ = v82
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 1:
                                                                                fmt.Println(v3, v78, v4)
                                                                            case 128:
                                                                                fmt.Println(v77, fi58)
                                                                            default:
                                                                                v83 := GenS4{F0: 2147483647, F1: true, F2: 0}
                                                                                _ = v83
                                                                            }
                                                                        case 2147483647:
                                                                            v78 -= -22
                                                                        case 0:
                                                                            _ = len(v5)
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    switch fi55 {
                                                                    case 4294967295:
                                                                        v56 = (v56 >> ((42) & 63))
                                                                    default:
                                                                        switch fi55 {
                                                                        case 32768:
                                                                            v1 = v1
                                                                        case -196:
                                                                            var v84 float64 = 3.14159
                                                                            _ = v84
                                                                        case 127:
                                                                            fmt.Println(v2, v78)
                                                                        default:
                                                                            var v85 uint64 = (v56 + v56)
                                                                            _ = v85
                                                                        }
                                                                    }
                                                                }
                                                                switch fi55 {
                                                                case 994:
                                                                    v57 = append(v57, fi58)
                                                                case 127:
                                                                    if false {
                                                                        v86 := GenS6{F0: -1, F1: 1000}
                                                                        v87 := v5[65535]
                                                                        _ = v86
                                                                        _ = v87
                                                                    } else {
                                                                        if v1 {
                                                                            fmt.Println(v3, v56)
                                                                        }
                                                                        var v88 int = fi58
                                                                        _ = v88
                                                                    }
                                                                case 4294967295:
                                                                    if ((v78 <= v78) || (-567 <= 1)) {
                                                                        v89 := GenS7{F0: 1, F1: 42}
                                                                        _ = v89
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            case -358:
                                                v90 := GenS7{F0: 91, F1: -2147483648}
                                                _ = v90
                                            case 128:
                                                if v1 {
                                                    if v1 {
                                                        if len(v57) >= 2 {
                                                            v57 = v57[1:len(v57)]
                                                        }
                                                    } else {
                                                        var v91 int = 2147483647
                                                        v3 -= v3
                                                        v3 -= (v78 & -9223372036854775808)
                                                        v92 := GenS3{F0: 255, F1: 2}
                                                        if v1 {
                                                            v93 := GenS3{F0: 128, F1: 1}
                                                            if true {
                                                                v78 -= v78
                                                                v94 := GenS5{F0: -128}
                                                                v78 += (v77 ^ 9223372036854775807)
                                                                fmt.Println(v77)
                                                                _ = v94
                                                            }
                                                            if false {
                                                                if (0.0 > ((((3.14159 - 0.0) + float32((0.0 - v2))) * 3.14159) * -2.25)) {
                                                                    v95 := GenS8{F0: v3, F1: 128}
                                                                    var v96 uint = 65535
                                                                    var v97 int8 = 127
                                                                    _ = v95
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                                var v98 int16 = 42
                                                                _ = v98
                                                            } else {
                                                                var v99 int8 = 100
                                                                v56 -= v56
                                                                _ = v99
                                                            }
                                                            v100 := GenS5{F0: 115}
                                                            fmt.Println(v91, v91)
                                                            var v101 uint64 = v56
                                                            _ = v100
                                                            _ = v101
                                                            _ = v93
                                                        } else {
                                                            if len(v57) > 0 {
                                                                v57[0] = v4
                                                            }
                                                            var v102 uint16 = 32767
                                                            _ = v102
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    var v103 uint32 = 0
                                                    _ = v103
                                                }
                                            }
                                        }
                                        switch (v37 + v37) {
                                        case "":
                                            fmt.Println(v56)
                                        case "\n\t":
                                            v56 += (^(v56))
                                        }
                                    }
                                    _ = v77
                                    _ = v78
                                } else {
                                    v104 := GenS7{F0: 1, F1: 2147483647}
                                    var v105 int32 = -892
                                    _ = v104
                                    _ = v105
                                }
                                v3 -= v3
                                _ = len(v57)
                            case "foo bar":
                                if true {
                                    switch (v4 & -1) {
                                    case -9223372036854775808:
                                        if (-2.25 <= 1.5) {
                                            v106 := GenS9{F0: (v4 | v4), F1: 65535, F2: (v56 >> ((1) & 63))}
                                            v107 := GenS8{F0: (v3 << ((51) & 63)), F1: 128}
                                            var v108 uint64 = (v56 + (v56 % (((v56 & 1)) | 1)))
                                            v4 = v4
                                            switch v37 {
                                            case "\n\t":
                                                v1 = v1
                                            default:
                                                fmt.Println(v4, v37, fi55)
                                            }
                                            _ = v106
                                            _ = v107
                                            _ = v108
                                        } else {
                                            switch (-(fi55)) {
                                            case 791:
                                                if (1 > -128) {
                                                    delete(v5, 670)
                                                    var v109 uint16 = 65535
                                                    var v110 float64 = 1e10
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    var v111 uint8 = 1
                                                    _ = v111
                                                }
                                            case 755:
                                                var v112 uint8 = 186
                                                _ = v112
                                            case 930:
                                                if v1 {
                                                    v37 = ((v37 + v37) + "a")
                                                } else {
                                                    if len(v57) > 0 {
                                                        v57[0] = v4
                                                    }
                                                    var v113 uint8 = 2
                                                    switch (v37 + v37) {
                                                    case "foo bar":
                                                        fmt.Println(fi58)
                                                    case "\n\t":
                                                        var v114 bool = v1
                                                        _ = len(v57)
                                                        fmt.Println(v4)
                                                        _ = v114
                                                    }
                                                    _ = v113
                                                }
                                            default:
                                                v37 = (("a" + (v37 + v37)) + v37)
                                            }
                                            var v115 int = (v4 * (fi55 / (((^(fi55))) | 1)))
                                            _ = v115
                                        }
                                    case 396:
                                        var v116 uint8 = 1
                                        _ = v116
                                    default:
                                        if (!v1) {
                                            if (v2 > (3.14159 * (v2 / v2))) {
                                                switch -39 {
                                                case 127:
                                                    if (32767 != -32768) {
                                                        if v1 {
                                                            fmt.Println(v1, v2)
                                                            switch (v37 + "\n\t") {
                                                            case "a":
                                                                var v117 int64 = v3
                                                                _ = v117
                                                            case "foo bar":
                                                                if ((int64(v56) > (v3 & (v3 & v3))) && ((v2 - v2) <= v2)) {
                                                                    v1 = v1
                                                                    if v1 {
                                                                        if (((((v37 + v37) + v37) + "hello") + (("foo bar" + v37) + (v37 + v37))) > ("\n\t" + (((v37 + (v37 + v37)) + "foo bar") + (v37 + v37)))) {
                                                                            v118 := GenS10{F0: fi55}
                                                                            _ = v118
                                                                        } else {
                                                                            if (false || (680 >= 32767)) {
                                                                                v119 := GenS4{F0: 464, F1: ((!true) || (815 == 65535)), F2: -128}
                                                                                v57 = append(v57, -9223372036854775808)
                                                                                _ = v119
                                                                            } else {
                                                                                v1 = v1
                                                                            }
                                                                            if len(v57) > 0 {
                                                                                v57[0] = (v4 * (fi55 >> ((923) & 63)))
                                                                            }
                                                                        }
                                                                        v1 = v1
                                                                        v120 := GenS3{F0: 127, F1: 42}
                                                                        v2 = 3.14159
                                                                        _ = v120
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                v121 := GenS10{F0: (fi55 * (fi58 << ((915) & 63)))}
                                                                v122 := GenS8{F0: v3, F1: 0}
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                var v123 uint16 = 673
                                                                v57 = append(v57, fi55)
                                                                v1 = true
                                                                _ = v123
                                                            }
                                                        }
                                                        var v124 int8 = -128
                                                        v2 -= v2
                                                        _ = v124
                                                    } else {
                                                        v4 = (v4 << ((uint(v56)) & 63))
                                                        _ = len(v5)
                                                        var v125 uint16 = 255
                                                        _ = v125
                                                    }
                                                case 32767:
                                                    v3 = (^(v3))
                                                default:
                                                    if v1 {
                                                        if v1 {
                                                            if (v3 != v3) {
                                                                if (((v37 + "foo bar") + "hello") > (v37 + (v37 + (v37 + v37)))) {
                                                                    if (false || v1) {
                                                                        var v126 uint32 = 0
                                                                        if (((-32768 > -623) && v1) || v1) {
                                                                            fmt.Println(v2)
                                                                            var v127 float32 = 1.5
                                                                            var v128 float32 = 0.1
                                                                            fmt.Println(v4, fi58)
                                                                            var v129 bool = v1
                                                                            _ = v127
                                                                            _ = v128
                                                                            _ = v129
                                                                        }
                                                                        switch (v37 + v37) {
                                                                        case "hello":
                                                                            v56 -= (v56 / ((42) | 1))
                                                                        case "":
                                                                            v1 = (216 >= 73)
                                                                        }
                                                                        _ = v126
                                                                    }
                                                                } else {
                                                                    switch (fi55 % ((-9223372036854775808) | 1)) {
                                                                    case -1:
                                                                        v3 += 42
                                                                        v37 = v37
                                                                    case -485:
                                                                        if v1 {
                                                                            v130 := GenS3{F0: 128, F1: 1}
                                                                            fmt.Println(fi55)
                                                                            delete(v5, 1)
                                                                            _ = v130
                                                                        }
                                                                    }
                                                                }
                                                                v131 := GenS6{F0: -885, F1: -929}
                                                                var v132 float64 = ((float64(1e10) + v2) / 0.1)
                                                                v37 = v37
                                                                var v133 int64 = (v3 - v3)
                                                                _ = v131
                                                                _ = v132
                                                                _ = v133
                                                            }
                                                            v4 += -2147483648
                                                            switch v37 {
                                                            case "":
                                                                if false {
                                                                    v37 = (("" + v37) + v37)
                                                                    if (2 > 9) {
                                                                        if (-1 >= 2147483647) {
                                                                            if false {
                                                                                var v134 uint = 0
                                                                                fmt.Println(v37)
                                                                                var v135 int = -9223372036854775808
                                                                                _ = v134
                                                                                _ = v135
                                                                            }
                                                                        } else {
                                                                            var v136 uint32 = 42
                                                                            var v137 bool = (!v1)
                                                                            v57 = append(v57, v4)
                                                                            var v138 int8 = -128
                                                                            _ = v136
                                                                            _ = v137
                                                                            _ = v138
                                                                        }
                                                                        var v139 int = v4
                                                                        _ = v139
                                                                    }
                                                                } else {
                                                                    var v140 string = "hello"
                                                                    if len(v57) > 0 {
                                                                        v57[0] = 4294967295
                                                                    }
                                                                    _ = v140
                                                                }
                                                            case "\n\t":
                                                                v57 = append(v57, (fi58 % ((fi58) | 1)))
                                                            default:
                                                                switch (v37 + (v37 + ((v37 + v37) + v37))) {
                                                                case "":
                                                                    var v141 uint32 = 1
                                                                    _ = v141
                                                                case "hello":
                                                                    switch (fi55 / ((fi55) | 1)) {
                                                                    case 792:
                                                                        v4 -= v4
                                                                    default:
                                                                        if true {
                                                                            var v142 float32 = 0.0
                                                                            switch (v4 | (^(fi55))) {
                                                                            case 998:
                                                                                fmt.Println(fi55, v56)
                                                                            case 42:
                                                                                switch (fi55 - (-(v4))) {
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v142, v56)
                                                                                default:
                                                                                    fmt.Println(v2, v56, v37)
                                                                                }
                                                                            case -9223372036854775808:
                                                                                v5[42] = v1
                                                                                v143 := GenS11{F0: -520}
                                                                                fmt.Println(v3, v1)
                                                                                _ = v143
                                                                            default:
                                                                                var v144 float64 = (v2 / v2)
                                                                                _ = v144
                                                                            }
                                                                            _ = v142
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if (42 <= -2147483648) {
                                                                        v145 := GenS12{F0: 1}
                                                                        _ = len(v5)
                                                                        _ = v145
                                                                    } else {
                                                                        v4 -= fi58
                                                                        v4 = fi55
                                                                        if ((528 < 42) && false) {
                                                                            fmt.Println(v37, v2, fi58)
                                                                            _ = len(v5)
                                                                        } else {
                                                                            fmt.Println(v4, v37)
                                                                            if len(v57) > 0 {
                                                                                v57[0] = (fi58 | (v4 ^ (v4 >> ((922) & 63))))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                if (!v1) {
                                                                    _ = len(v57)
                                                                } else {
                                                                    v146 := GenS13{F0: v3, F1: v3, F2: 42}
                                                                    v1 = ((false || (127 == uint8(v56))) || true)
                                                                    v4 += (fi55 ^ (v4 % (((fi55 - (fi55 << ((22) & 63)))) | 1)))
                                                                    _ = v146
                                                                }
                                                            }
                                                            _ = len(v5)
                                                            fmt.Println(fi55)
                                                        }
                                                        switch "" {
                                                        case "":
                                                            if len(v57) > 0 {
                                                                v57[0] = fi58
                                                            }
                                                        case "foo bar":
                                                            var v147 uint32 = uint32(v56)
                                                            _ = v147
                                                        case "a":
                                                            var v148 uint8 = 2
                                                            _ = v148
                                                        default:
                                                            var v149 int32 = -639
                                                            _ = v149
                                                        }
                                                    } else {
                                                        var v150 uint64 = (v56 + v56)
                                                        _ = v150
                                                    }
                                                }
                                                var v151 uint64 = (v56 - v56)
                                                v152 := GenS10{F0: (^(fi55))}
                                                var v153 float64 = v2
                                                _ = v151
                                                _ = v152
                                                _ = v153
                                            }
                                            v4 = fi55
                                            v154 := GenS3{F0: 79, F1: uint16(v56)}
                                            var v155 string = "a"
                                            _ = v154
                                            _ = v155
                                        } else {
                                            fmt.Println(v4, v4, v37)
                                            if v1 {
                                                var v156 bool = v1
                                                _ = v156
                                            }
                                            var v157 uint32 = 0
                                            _ = v157
                                        }
                                        v1 = true
                                    }
                                } else {
                                    v158 := GenS4{F0: 4294967295, F1: v1, F2: -72}
                                    if (42 > 1) {
                                        if v1 {
                                            v3 = v3
                                            v37 = (("\n\t" + v37) + "a")
                                            if v1 {
                                                var v159 int16 = 32767
                                                _ = v159
                                            }
                                        }
                                        fmt.Println(v37, v1, v3)
                                        if v1 {
                                            var v160 uint8 = 255
                                            var v161 bool = v1
                                            _ = v160
                                            _ = v161
                                        }
                                    }
                                    var v162 string = v37
                                    v163 := GenS13{F0: (v3 >> ((255) & 63)), F1: (v3 & v3), F2: 100}
                                    _ = v158
                                    _ = v162
                                    _ = v163
                                }
                            }
                        }
                        _ = fi58
                        _ = v59
                    }
                    if false {
                        var v164 int16 = 42
                        v165 := GenS7{F0: (^(v164)), F1: 42}
                        _ = v164
                        _ = v165
                    }
                    _ = v57
                }
                switch (v4 & -2147483648) {
                case -2147483648:
                    v56 = v56
                case 2147483647:
                    delete(v5, 2)
                }
                v166 := GenS8{F0: v3, F1: 66}
                _ = v166
            } else {
                for fi167 := 0; fi167 < 14; fi167++ {
                    switch (fi55 << ((4294967295) & 63)) {
                    case 127:
                        v168 := []int{(v4 << ((461) & 63))}
                        _ = v168
                    case -758:
                        v169 := GenS14{F0: 1}
                        _ = v169
                    case 1:
                        switch 1 {
                        case 127:
                            fmt.Println(fi55, v37, v3)
                        case 875:
                            var v170 uint = 255
                            _ = v170
                        case 32767:
                            switch "\n\t" {
                            case "a":
                                v3 = 42
                            default:
                                switch v37 {
                                case "":
                                    v171 := GenS9{F0: v4, F1: 324, F2: v56}
                                    _ = v171
                                default:
                                    v2 -= v2
                                }
                                fmt.Println(v3, v3)
                                fmt.Println(v37, v37, v2)
                            }
                        default:
                            if (514 > 42) {
                                v1 = (!true)
                                var v172 int64 = (v3 / (((v3 % ((v3) | 1))) | 1))
                                v173 := []int{(fi167 / ((-2147483648) | 1)), v4, 147, fi55}
                                _ = v172
                                _ = v173
                            } else {
                                v1 = v1
                                fmt.Println(v56, v1)
                                v174 := GenS3{F0: 255, F1: 2}
                                _ = v174
                            }
                        }
                    default:
                        v4 += (v4 + (fi167 % ((v4) | 1)))
                    }
                    var v175 int64 = (v3 - v3)
                    _ = fi167
                    _ = v175
                }
            }
            fmt.Println(fi55, v37, v56)
        }
        _ = fi55
        _ = v56
    }
    if v1 {
        v176 := v5[65535]
        v5[107] = (!(!(v1 && true)))
        _ = v176
    } else {
        var v177 uint = 65535
        _ = v177
    }
    var v178 uint64 = 4294967295
    v179 := GenS15{F0: 0, F1: (-(v3)), F2: 1}
    var v180 int = v4
    v37 = v37
    var v181 int = v4
    delete(v5, 2)
    switch (v37 + (((v37 + (v37 + v37)) + "\n\t") + v37)) {
    case "hello":
        switch (v180 ^ 1) {
        case 42:
            switch v37 {
            case "":
                switch "a" {
                case "a":
                    for fi182 := 0; fi182 < 2; fi182++ {
                        v183 := []int{(fi182 & (^(v181))), 343, v180, (v181 * (fi182 / (((fi182 / (((fi182 + (v181 & (v180 / ((fi182) | 1))))) | 1))) | 1)))}
                        _ = fi182
                        _ = v183
                    }
                }
            }
        default:
            v184 := GenS4{F0: 531, F1: ((float64((1.5 - 1e10)) / v2) == v2), F2: 1}
            _ = v184
        }
    case "a":
        fmt.Println(v1, v2, v181)
    case "":
        v37 = (v37 + ((v37 + v37) + (v37 + v37)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v178:", v178)
    fmt.Println("v180:", v180)
    fmt.Println("v181:", v181)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v178
    _ = v179
    _ = v180
    _ = v181
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
