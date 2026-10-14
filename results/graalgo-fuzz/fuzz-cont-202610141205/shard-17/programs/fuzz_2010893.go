package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 bool; F1 int64; F2 int16 }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 uint32, p1 string) uint16 {
    if true {
        var v6 uint16 = 235
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{(fi7 & (fi7 + (fi7 % ((227) | 1))))}
            v9 := GenS1{F0: v6}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: (v6 | (v6 | (v6 % ((v6) | 1))))}
        _ = v10
        _ = v6
    } else {
        v11 := []int{2147483647, -938, 340, 32768}
        p1 = ""
        for fi12 := 0; fi12 < 1; fi12++ {
            v11 = append(v11, 32768)
            _ = fi12
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        p1 = p1
        _ = fi13
    }
    if false {
        var v14 uint8 = 128
        _ = v14
    } else {
        p0 = 42
    }
    if false {
        p1 = "a"
        if true {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 float64 = 1.5
                p0 = (p0 * 42)
                _ = fi15
                _ = v16
            }
            var v17 int16 = 845
            _ = v17
        } else {
            var v18 bool = (true || false)
            switch -219 {
            case 42:
                if (755 == 32767) {
                    if false {
                        p0 = (p0 % (((^(p0))) | 1))
                        v18 = v18
                        p1 = "foo bar"
                    }
                } else {
                    p1 = p1
                    v19 := []int{}
                    _ = v19
                }
            case 128:
                v20 := GenS1{F0: 255}
                _ = v20
            case -546:
                v18 = (-128 < -128)
            }
            p1 = "a"
            _ = v18
        }
        p0 -= p0
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint64, p1 int64) uint8 {
    p0 = 0
    if (!(true && (42 != 42))) {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := []int{}
            var v23 uint32 = 20
            _ = fi21
            _ = v22
            _ = v23
        }
        p0 = 42
        v24 := GenS1{F0: 65535}
        _ = v24
    } else {
        if true {
            fmt.Println(p0)
            p0 -= 4294967295
            v25 := []int{597, 32768}
            _ = v25
        } else {
            p0 += p0
        }
    }
    var v26 uint64 = 4294967295
    var v27 uint8 = 1
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return v27
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 uint = 65535
    var v4 string = ""
    v5 := []string{v4}
    v1[1] = 127
    if ((0 != -991) && (!((3.14159 / (-2.25 + 1e10)) <= float32((((0.0 + float64(1.5)) - (float64((float32((0.1 * float64((0.0 - 3.14159)))) * 1.5)) * 3.14159)) / -2.25))))) {
        v3 = v3
        var v28 uint16 = 255
        _ = v28
    }
    switch -531 {
    case 698:
        v29 := GenS2{F0: 707}
        _ = v29
    default:
        v3 = v3
    }
    _ = len(v5)
    var v30 int32 = 42
    var v31 uint8 = 128
    var v32 uint32 = 189
    var v33 uint8 = v31
    if (!true) {
        if false {
            v5 = append(v5, "")
            if false {
                v4 = v4
                var v34 string = v4
                _ = v34
            } else {
                if ((((0 != 42) || false) || false) && false) {
                    if false {
                        v31 = 128
                        v31 -= 14
                    }
                    fmt.Println(v32, v31)
                    for fi35 := 0; fi35 < 6; fi35++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v3, v4, v31)
                        _ = fi35
                    }
                } else {
                    v36 := GenS1{F0: 32767}
                    var v37 uint16 = 42
                    _ = v36
                    _ = v37
                }
                switch 9223372036854775807 {
                case 42:
                    for fi38 := 0; fi38 < 6; fi38++ {
                        fmt.Println(v32, v30)
                        _ = fi38
                    }
                default:
                    switch 807 {
                    case 32768:
                        v30 = 0
                    case -9223372036854775808:
                        v39 := GenS2{F0: 4294967295}
                        _ = v39
                    }
                }
                fmt.Println(v33)
            }
            for fi40 := 0; fi40 < 2; fi40++ {
                v5 = append(v5, v4)
                for fi41 := 0; fi41 < 5; fi41++ {
                    if true {
                        fmt.Println(v30)
                        fmt.Println(fi41)
                        if true {
                            v42 := GenS2{F0: v3}
                            v33 += (v33 >> ((v3) & 63))
                            _ = len(v5)
                            v43 := GenS1{F0: 2}
                            v44 := GenS2{F0: 466}
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        v31 += genFn1(580, 0)
                    }
                    var v45 int16 = 1
                    v46 := v1[(v32 << (((v3 - v3)) & 63))]
                    _ = fi41
                    _ = v45
                    _ = v46
                }
                v47 := v1[87]
                _ = fi40
                _ = v47
            }
        }
    } else {
        if len(v5) > 0 {
            v5[0] = v4
        }
        var v48 uint16 = uint16(v3)
        v49 := GenS2{F0: (v3 - v3)}
        _ = v48
        _ = v49
    }
    v50 := GenS1{F0: 0}
    v51 := GenS3{F0: v30, F1: 127}
    switch 1 {
    case 670:
        switch 0 {
        case -9223372036854775808:
            _ = len(v2)
        case 1:
            var v52 string = "a"
            _ = v52
        case 9223372036854775807:
            for fi53 := 0; fi53 < 6; fi53++ {
                if true {
                    var v54 uint8 = 1
                    for fi55 := 0; fi55 < 1; fi55++ {
                        var v56 int32 = v30
                        var v57 uint32 = v32
                        fmt.Println(v31, v54)
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    var v58 uint16 = 560
                    _ = v54
                    _ = v58
                } else {
                    fmt.Println(v32)
                    if true {
                        v32 += 2
                        var v59 float32 = (((0.0 + (0.0 + (1.5 * -2.25))) - 3.14159) * ((1.5 - -2.25) / 1.5))
                        v60 := GenS1{F0: 435}
                        _ = v59
                        _ = v60
                    }
                }
                if true {
                    if (!((-(v30)) < 1)) {
                        switch (fi53 ^ 32767) {
                        case -425:
                            var v61 uint64 = 862
                            _ = v61
                        case 220:
                            switch v4 {
                            case "foo bar":
                                v62 := GenS3{F0: v30, F1: 93}
                                _ = v62
                            }
                        case 32767:
                            fmt.Println(v4, v31, v32)
                        }
                        if ((-32768 < -32768) && (fi53 != (fi53 / ((fi53) | 1)))) {
                            for fi63 := 0; fi63 < 6; fi63++ {
                                v3 = 4294967295
                                v5 = append(v5, v4)
                                if false {
                                    var v64 bool = (((1.5 * -2.25) + ((0.1 - 0.0) - 1.5)) != -2.25)
                                    _ = v64
                                } else {
                                    if (!(false || false)) {
                                        v3 -= 42
                                        if (-2.25 > ((0.1 - 0.1) * float64(((1.5 + ((1.5 + float32(0.1)) - 0.0)) / 1e10)))) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v33)
                                            case "":
                                                if (812 != 9223372036854775807) {
                                                    v32 -= (^(v32))
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    switch (fi63 - -9223372036854775808) {
                                                    case 32768:
                                                        switch v4 {
                                                        case "":
                                                            if true {
                                                                var v65 uint = (v3 ^ 622)
                                                                if (739 > 42) {
                                                                    var v66 int16 = 32767
                                                                    _ = v66
                                                                }
                                                                _ = v65
                                                            } else {
                                                                fmt.Println(v3, v32, fi63)
                                                                var v67 int32 = (v30 + v30)
                                                                if (!true) {
                                                                    v68 := v2[v32]
                                                                    if len(v5) > 0 {
                                                                        v5[0] = v4
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    v67 = v67
                                                                    v67 = v30
                                                                    if (((v3 | (^(v3))) < v3) || true) {
                                                                        if true {
                                                                            switch v4 {
                                                                            case "\n\t":
                                                                                var v69 int8 = 1
                                                                                _ = v69
                                                                            case "foo bar":
                                                                                delete(v2, (v32 << ((255) & 63)))
                                                                            case "hello":
                                                                                switch fi53 {
                                                                                case 128:
                                                                                    switch (v4 + ("hello" + "a")) {
                                                                                    case "hello":
                                                                                        if (9223372036854775807 >= 42) {
                                                                                            var v70 uint = v3
                                                                                            v31 -= v33
                                                                                            var v71 int = (fi63 + 32768)
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        } else {
                                                                                            if true {
                                                                                                if false {
                                                                                                    var v72 string = ((v4 + "foo bar") + "a")
                                                                                                    _ = v72
                                                                                                } else {
                                                                                                    var v73 float32 = 3.14159
                                                                                                    fmt.Println(v67, v32, v67)
                                                                                                    var v74 int = fi53
                                                                                                    fmt.Println(fi53, v67)
                                                                                                    _ = v73
                                                                                                    _ = v74
                                                                                                }
                                                                                                var v75 uint16 = 0
                                                                                                v30 = int32(v75)
                                                                                                _ = v75
                                                                                            } else {
                                                                                                var v76 int64 = -9223372036854775808
                                                                                                _ = v76
                                                                                            }
                                                                                            var v77 int16 = -495
                                                                                            _ = v77
                                                                                        }
                                                                                    case "a":
                                                                                        v78 := GenS4{F0: 0}
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v32)
                                                                            var v79 uint64 = 841
                                                                            v4 = ((v4 + ((v4 + v4) + "foo bar")) + v4)
                                                                            if (genFn1((v79 >> ((2) & 63)), -9223372036854775808) >= v33) {
                                                                                if true {
                                                                                    if (false && false) {
                                                                                        switch (fi53 + 212) {
                                                                                        case 2147483647:
                                                                                            v31 = (v33 - (^(v31)))
                                                                                        }
                                                                                        v3 = uint(v79)
                                                                                    } else {
                                                                                        var v80 string = (v4 + (v4 + v4))
                                                                                        var v81 float64 = (1e10 + float64(-2.25))
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    }
                                                                                    var v82 int8 = int8(v79)
                                                                                    if true {
                                                                                        v3 = v3
                                                                                        v33 = (^(v33))
                                                                                    }
                                                                                    _ = v82
                                                                                }
                                                                                v83 := GenS1{F0: 65535}
                                                                                _ = v83
                                                                            }
                                                                            v30 = (v30 & (v30 ^ -2147483648))
                                                                            _ = v79
                                                                        }
                                                                        v84 := GenS3{F0: (-(v67)), F1: 0}
                                                                        v33 = genFn1(730, -945)
                                                                        v33 += 255
                                                                        _ = v84
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v85 uint64 = 4294967295
                                                                            _ = v85
                                                                        case "":
                                                                            v33 += genFn1(42, -1)
                                                                        case "foo bar":
                                                                            v67 -= (v67 * (v67 % ((1) | 1)))
                                                                        }
                                                                        switch (^(fi53)) {
                                                                        case -2147483648:
                                                                            var v86 uint16 = 255
                                                                            _ = v86
                                                                        case 32768:
                                                                            var v87 uint32 = (v32 / ((160) | 1))
                                                                            _ = v87
                                                                        case 279:
                                                                            v1[v32] = fi63
                                                                        default:
                                                                            v32 -= v32
                                                                        }
                                                                        v4 = v4
                                                                    }
                                                                }
                                                                _ = v67
                                                            }
                                                        default:
                                                            if (1e10 == 0.0) {
                                                                v31 += v33
                                                                if false {
                                                                    switch (fi63 >> ((v3) & 63)) {
                                                                    case 2147483647:
                                                                        if false {
                                                                            fmt.Println(v4, v30)
                                                                            if false {
                                                                                switch fi63 {
                                                                                case -1:
                                                                                    v4 = "foo bar"
                                                                                default:
                                                                                    if (1.5 > -2.25) {
                                                                                        v33 += (v33 >> (((v3 / (((v3 >> (((v3 * (v3 >> (((v3 + 2)) & 63)))) & 63))) | 1))) & 63))
                                                                                    }
                                                                                    var v88 int8 = int8(v3)
                                                                                    switch 42 {
                                                                                    case -930:
                                                                                        var v89 uint32 = 0
                                                                                        _ = v89
                                                                                    case -2147483648:
                                                                                        v90 := GenS1{F0: 0}
                                                                                        _ = v90
                                                                                    case 4294967295:
                                                                                        switch (v4 + "") {
                                                                                        case "hello":
                                                                                            _ = len(v1)
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                var v91 int8 = (v88 * v88)
                                                                                                v33 = (v33 / ((genFn1(42, 79)) | 1))
                                                                                                _ = v91
                                                                                            }
                                                                                        case "a":
                                                                                            v31 += v33
                                                                                        default:
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    default:
                                                                                        var v92 int = 127
                                                                                        _ = v92
                                                                                    }
                                                                                    _ = len(v5)
                                                                                    _ = v88
                                                                                }
                                                                                fmt.Println(v32)
                                                                            }
                                                                            var v93 float64 = (-2.25 * 1e10)
                                                                            v33 -= v31
                                                                            var v94 string = v4
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                    }
                                                                    var v95 int8 = 5
                                                                    v3 += (^(v3))
                                                                    _ = v95
                                                                }
                                                                switch 522 {
                                                                case 0:
                                                                    var v96 bool = true
                                                                    _ = v96
                                                                case 9223372036854775807:
                                                                    switch -2147483648 {
                                                                    case 917:
                                                                        v97 := GenS1{F0: 255}
                                                                        _ = v97
                                                                    case 602:
                                                                        fmt.Println(fi63, fi63)
                                                                    default:
                                                                        var v98 uint32 = uint32(v3)
                                                                        _ = v98
                                                                    }
                                                                default:
                                                                    if true {
                                                                        v30 = v30
                                                                        v30 -= -2147483648
                                                                        var v99 uint8 = v33
                                                                        _ = v99
                                                                    } else {
                                                                        fmt.Println(v33, v4)
                                                                        delete(v1, (^(v32)))
                                                                        fmt.Println(fi63, v32)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if (v33 < v31) {
                                                            var v100 int = (fi53 | 32767)
                                                            if true {
                                                                if (0.0 != float64(((((0.1 - 0.1) - 1.5) / 1.0) + 1e10))) {
                                                                    if (false || (42 == 64)) {
                                                                        if (!false) {
                                                                            v1[(v32 >> ((v3) & 63))] = 127
                                                                            switch fi53 {
                                                                            case -9223372036854775808:
                                                                                var v101 uint32 = (^(v32))
                                                                                _ = v101
                                                                            case 32767:
                                                                                var v102 uint16 = 0
                                                                                _ = v102
                                                                            }
                                                                            fmt.Println(v100, v33, v33)
                                                                        }
                                                                        var v103 int8 = -1
                                                                        if false {
                                                                            v4 = v4
                                                                            var v104 uint16 = genFn0(v32, "hello")
                                                                            v3 += v3
                                                                            _ = v104
                                                                        }
                                                                        _ = v103
                                                                    } else {
                                                                        var v105 bool = (uint16(v33) == 1)
                                                                        _ = v105
                                                                    }
                                                                    fmt.Println(v32)
                                                                }
                                                                v2[v32] = false
                                                            }
                                                            _ = v100
                                                        } else {
                                                            v106 := GenS1{F0: 1}
                                                            fmt.Println(v4)
                                                            v107 := GenS1{F0: 415}
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    }
                                                }
                                                var v108 int64 = 1
                                                switch (fi63 + (fi63 | (fi63 * (fi53 - (fi53 * fi53))))) {
                                                case 4294967295:
                                                    var v109 int8 = -80
                                                    _ = v109
                                                }
                                                _ = v108
                                            default:
                                                v31 = (v33 + (v33 ^ (v31 ^ 42)))
                                            }
                                            switch "foo bar" {
                                            case "":
                                                v110 := GenS4{F0: -1}
                                                _ = v110
                                            case "hello":
                                                var v111 uint64 = 42
                                                _ = v111
                                            case "\n\t":
                                                var v112 uint8 = (^(v33))
                                                _ = v112
                                            }
                                            if (((v30 >> ((v3) & 63)) != v30) && (true && true)) {
                                                switch (fi53 >> ((v3) & 63)) {
                                                case -9223372036854775808:
                                                    if (true || false) {
                                                        var v113 int16 = 0
                                                        v31 = v33
                                                        if true {
                                                            v3 += (v3 + 18446744073709551615)
                                                        } else {
                                                            _ = len(v5)
                                                            delete(v2, 1)
                                                            var v114 int16 = (v113 | v113)
                                                            _ = v114
                                                        }
                                                        _ = v113
                                                    }
                                                default:
                                                    var v115 int8 = int8(v3)
                                                    _ = v115
                                                }
                                                if (!(v4 == "foo bar")) {
                                                    v116 := v2[(^(v32))]
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v116
                                                } else {
                                                    v30 += (^(v30))
                                                }
                                                if (!(1 != 483)) {
                                                    var v117 uint64 = 0
                                                    _ = v117
                                                } else {
                                                    switch ("foo bar" + v4) {
                                                    case "\n\t":
                                                        fmt.Println(fi63, v4, v31)
                                                    case "a":
                                                        var v118 bool = false
                                                        _ = v118
                                                    case "hello":
                                                        v33 = v33
                                                    default:
                                                        v119 := GenS5{F0: (fi53 / ((fi63) | 1)), F1: 0}
                                                        _ = v119
                                                    }
                                                    var v120 bool = (v4 >= ("hello" + v4))
                                                    delete(v1, v32)
                                                    _ = v120
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v121 float32 = (3.14159 + (0.1 + ((1e10 * 3.14159) / ((-2.25 - 0.1) / float32(3.14159)))))
                                                fmt.Println(v31)
                                                v122 := GenS1{F0: 1}
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                        v123 := GenS2{F0: v3}
                                        _ = v123
                                    } else {
                                        switch ((v4 + (v4 + v4)) + "foo bar") {
                                        case "":
                                            _ = len(v5)
                                            var v124 uint64 = 4294967295
                                            _ = v124
                                        case "foo bar":
                                            if (false && false) {
                                                v32 += v32
                                                var v125 int64 = 0
                                                v30 = (^(v30))
                                                _ = v125
                                            }
                                        default:
                                            v3 -= 2
                                        }
                                        _ = len(v5)
                                    }
                                    v31 -= (v31 * v31)
                                }
                                var v126 string = ""
                                fmt.Println(v30, v4, v30)
                                _ = len(v2)
                                _ = fi63
                                _ = v126
                            }
                            for fi127 := 0; fi127 < 3; fi127++ {
                                v128 := GenS6{F0: 1}
                                _ = fi127
                                _ = v128
                            }
                        } else {
                            for fi129 := 0; fi129 < 15; fi129++ {
                                v2[(v32 >> ((v3) & 63))] = (!(v30 == v30))
                                _ = fi129
                            }
                        }
                    } else {
                        var v130 string = ("foo bar" + "foo bar")
                        _ = v130
                    }
                    v30 += 770
                    if (((255 >= 1) || false) && true) {
                        var v131 string = ((("\n\t" + ((v4 + "a") + v4)) + "hello") + v4)
                        if true {
                            v32 = v32
                            for fi132 := 0; fi132 < 13; fi132++ {
                                v31 -= (^(v31))
                                var v133 uint64 = 0
                                switch (fi53 + fi53) {
                                case -505:
                                    v134 := GenS5{F0: fi53, F1: int8(v33)}
                                    _ = v134
                                case -1:
                                    fmt.Println(v4)
                                case 32767:
                                    fmt.Println(v3)
                                default:
                                    v33 -= (^(v31))
                                }
                                _ = fi132
                                _ = v133
                            }
                            v1[2147483647] = fi53
                        }
                        _ = v131
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        for fi135 := 0; fi135 < 5; fi135++ {
                            v136 := GenS1{F0: 2}
                            v137 := v1[(v32 ^ 223)]
                            _ = fi135
                            _ = v136
                            _ = v137
                        }
                    case "a":
                        v32 -= 65535
                    case "hello":
                        v138 := GenS7{F0: v3, F1: (!false)}
                        _ = v138
                    }
                    var v139 uint16 = 255
                    delete(v1, uint32(v139))
                    _ = v139
                }
                _ = fi53
            }
        default:
            v31 = genFn1(2, -557)
        }
    }
    v4 = v4
    if false {
        switch "" {
        case "foo bar":
            if (v32 < uint32(v3)) {
                v140 := GenS4{F0: -84}
                v141 := GenS8{F0: false, F1: 42, F2: 1}
                v33 = v33
                _ = v140
                _ = v141
            }
        case "\n\t":
            for fi142 := 0; fi142 < 6; fi142++ {
                if false {
                    var v143 int16 = 0
                    v144 := GenS6{F0: 100}
                    for fi145 := 0; fi145 < 20; fi145++ {
                        var v146 float64 = 1.5
                        _ = fi145
                        _ = v146
                    }
                    _ = v143
                    _ = v144
                } else {
                    for fi147 := 0; fi147 < 7; fi147++ {
                        if false {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if false {
                                var v148 uint64 = 21
                                switch (fi142 & (-(fi147))) {
                                case 1:
                                    _ = len(v1)
                                }
                                _ = v148
                            }
                            if false {
                                var v149 float32 = (((float32(1.5) * 0.0) * 0.0) + -2.25)
                                _ = v149
                            }
                            var v150 uint64 = 42
                            v33 = (v33 ^ (v31 * genFn1(42, 9223372036854775807)))
                            _ = v150
                        } else {
                            var v151 float32 = ((3.14159 + -2.25) * float32(-2.25))
                            v151 = (-2.25 * v151)
                            var v152 string = v4
                            _ = v151
                            _ = v152
                        }
                        v3 = (v3 >> ((255) & 63))
                        _ = fi147
                    }
                    if (0 <= 59) {
                        _ = len(v5)
                        v32 = (v32 % ((v32) | 1))
                        if false {
                            v3 -= (v3 * (v3 - (v3 >> ((v3) & 63))))
                            v31 += (v33 % ((v31) | 1))
                        } else {
                            var v153 int32 = v30
                            for fi154 := 0; fi154 < 1; fi154++ {
                                if (9223372036854775807 > -9223372036854775808) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if ((fi142 | 58) > fi142) {
                                        switch 32767 {
                                        case 1:
                                            v155 := v2[v32]
                                            _ = v155
                                        case 42:
                                            fmt.Println(fi142)
                                        default:
                                            fmt.Println(v3)
                                        }
                                        var v156 int16 = 331
                                        v157 := GenS9{F0: v3, F1: (0.1 > 1e10)}
                                        _ = v156
                                        _ = v157
                                    }
                                } else {
                                    var v158 uint32 = 2
                                    _ = v158
                                }
                                var v159 int32 = 19
                                var v160 uint16 = uint16(v3)
                                _ = fi154
                                _ = v159
                                _ = v160
                            }
                            v4 = ("" + (v4 + ((v4 + "foo bar") + "foo bar")))
                            _ = v153
                        }
                    }
                    v33 -= genFn1(2, 9223372036854775807)
                }
                v161 := GenS10{F0: (v32 ^ (v32 | (^(v32)))), F1: (v32 / ((v32) | 1)), F2: v3}
                var v162 int32 = (v30 & v30)
                _ = fi142
                _ = v161
                _ = v162
            }
        case "hello":
            fmt.Println(v33)
        default:
            v33 += (v33 * 71)
        }
        var v163 int = 32767
        _ = v163
    }
    v3 = v3
    if ((v30 == v30) || false) {
        var v164 uint16 = 2
        var v165 uint64 = 1
        _ = v164
        _ = v165
    } else {
        v3 = (v3 - (v3 ^ (v3 + v3)))
        v3 += (v3 & v3)
        _ = len(v1)
    }
    v166 := GenS3{F0: 0, F1: 0}
    v32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v166
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
