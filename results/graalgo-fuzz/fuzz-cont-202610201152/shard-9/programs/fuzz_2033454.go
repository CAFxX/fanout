package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint32; F1 bool; F2 uint }
type GenS8 struct { F0 string; F1 int32; F2 int32 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint; F1 uint64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint; F1 uint }
type GenS15 struct { F0 string; F1 uint8; F2 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint64; F1 int8 }
type GenS19 struct { F0 uint8; F1 string; F2 uint64 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 bool; F1 uint16 }
func genFn0(p0 int, p1 uint16) int {
    var v4 int8 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        p1 = p1
        if (32768 < p0) {
            v6 := []int{fi5, -9223372036854775808}
            _ = v6
        } else {
            v7 := GenS1{F0: p1}
            var v8 uint16 = 255
            _ = v7
            _ = v8
        }
        var v9 float32 = 0.1
        _ = fi5
        _ = v9
    }
    v4 = v4
    var v10 int8 = v4
    fmt.Println(v4, v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return (p0 * p0)
}

func main() {
    var v1 int16 = -202
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    delete(v2, 128)
    for fi11 := 0; fi11 < 12; fi11++ {
        delete(v2, (-(fi11)))
        _ = fi11
    }
    switch "hello" {
    case "":
        if (true && ("hello" >= ("" + ((("" + "hello") + "a") + "foo bar")))) {
            if ("" >= ("\n\t" + "a")) {
                v1 -= v1
            }
            v12 := []int{42, genFn0(int(v1), 42)}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                v3 += ((v3 * float32(((float64(float32(1e10)) + 0.0) - 3.14159))) - v3)
                switch fi13 {
                case 128:
                    switch "a" {
                    case "a":
                        fmt.Println(v3, v3, fi13)
                    default:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v15 := []int{}
                            _ = fi14
                            _ = v15
                        }
                    }
                case 0:
                    if false {
                        v3 = (v3 / float32(3.14159))
                        var v16 float32 = (v3 * v3)
                        _ = v16
                    } else {
                        if ((36 <= -128) || true) {
                            if (!true) {
                                switch ("" + "\n\t") {
                                case "":
                                    switch fi13 {
                                    case 70:
                                        v1 -= int16(fi13)
                                    case 32768:
                                        var v17 int8 = -88
                                        _ = v17
                                    default:
                                        var v18 uint8 = 29
                                        _ = v18
                                    }
                                case "\n\t":
                                    if true {
                                        switch fi13 {
                                        case 32767:
                                            var v19 uint = 42
                                            _ = v19
                                        default:
                                            for fi20 := 0; fi20 < 15; fi20++ {
                                                var v21 uint16 = 65535
                                                switch (fi20 % (((fi13 / ((fi20) | 1))) | 1)) {
                                                case -1:
                                                    if (18446744073709551615 == 4294967295) {
                                                        v22 := GenS1{F0: v21}
                                                        var v23 float64 = (0.1 + ((float64((1.5 / v3)) - ((3.14159 * 3.14159) - (0.0 - (float64(3.14159) / 3.14159)))) + ((((0.0 * 1e10) * -2.25) * 1e10) / 3.14159)))
                                                        v24 := GenS2{F0: (v1 - v1), F1: 2, F2: 1}
                                                        _ = v22
                                                        _ = v23
                                                        _ = v24
                                                    } else {
                                                        var v25 int8 = 0
                                                        switch fi20 {
                                                        case 216:
                                                            v26 := v2[genFn0(int(v21), (v21 - v21))]
                                                            _ = v26
                                                        case 42:
                                                            if ((!false) && ((!false) || (true || (4294967295 == 42)))) {
                                                                switch 32767 {
                                                                case 4294967295:
                                                                    var v27 float32 = v3
                                                                    _ = v27
                                                                case 42:
                                                                    var v28 string = "hello"
                                                                    _ = v28
                                                                case -226:
                                                                    fmt.Println(v25)
                                                                }
                                                                v25 -= -1
                                                                v29 := GenS1{F0: (v21 - (^(v21)))}
                                                                _ = v29
                                                            } else {
                                                                if (!(v21 < 110)) {
                                                                    delete(v2, (fi20 | 32767))
                                                                    switch 42 {
                                                                    case 42:
                                                                        v30 := GenS3{F0: 18446744073709551615}
                                                                        _ = v30
                                                                    case -605:
                                                                        v31 := GenS1{F0: 255}
                                                                        _ = v31
                                                                    case 128:
                                                                        var v32 float32 = (float32((1.5 * 0.1)) + float32(1.5))
                                                                        _ = v32
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                    var v33 uint32 = 2147483647
                                                                    _ = v33
                                                                }
                                                                delete(v2, 32767)
                                                                var v34 int32 = 42
                                                                if (!true) {
                                                                    if false {
                                                                        var v35 uint8 = 2
                                                                        var v36 uint64 = 554
                                                                        _ = v35
                                                                        _ = v36
                                                                    } else {
                                                                        var v37 bool = ((!((!false) && false)) && true)
                                                                        _ = v37
                                                                    }
                                                                }
                                                                v21 = (^(v21))
                                                                v38 := GenS2{F0: (v1 + -32768), F1: 0, F2: 571}
                                                                _ = v34
                                                                _ = v38
                                                            }
                                                        case 1:
                                                            v39 := []int{fi13, fi20, fi20, fi20}
                                                            _ = v39
                                                        default:
                                                            v3 = (float32(1e10) * ((v3 - float32(-2.25)) + (v3 * float32((1.5 / ((((1.5 - (float64(0.0) / 1e10)) + float64(float32((1.5 * (0.1 * (0.0 + (0.1 - 0.0))))))) + ((((3.14159 / (-2.25 - ((0.0 * (float64((((v3 / float32(1.5)) + ((float32(3.14159) * (3.14159 * v3)) * ((3.14159 + -2.25) / (v3 + 3.14159)))) * v3)) / 1e10)) + float64(1e10)))) - 1.5) / (0.0 * ((((1e10 * -2.25) / float64(v3)) + 1e10) - (1e10 - (-2.25 + -2.25))))) * 0.0)) / float64(v3)))))))
                                                        }
                                                        _ = v25
                                                    }
                                                }
                                                _ = fi20
                                                _ = v21
                                            }
                                        }
                                        switch (fi13 % ((fi13) | 1)) {
                                        case -9223372036854775808:
                                            var v40 int16 = v1
                                            _ = v40
                                        case 32768:
                                            for fi41 := 0; fi41 < 9; fi41++ {
                                                if true {
                                                    v42 := GenS4{F0: 0}
                                                    _ = v42
                                                }
                                                if false {
                                                    if false {
                                                        fmt.Println(fi41, v3, fi41)
                                                        v3 -= -2.25
                                                        v3 = float32(((0.1 * (((1.5 * ((3.14159 / 1.0) + (0.0 * -2.25))) - 3.14159) + -2.25)) - float64(3.14159)))
                                                    } else {
                                                        if (!("" <= "hello")) {
                                                            var v43 uint8 = 128
                                                            _ = v43
                                                        }
                                                        v44 := GenS1{F0: 65535}
                                                        _ = v44
                                                    }
                                                    v1 += 1
                                                } else {
                                                    switch (("" + ("hello" + "foo bar")) + "\n\t") {
                                                    case "":
                                                        v1 = -97
                                                    }
                                                    if false {
                                                        if false {
                                                            if false {
                                                                v45 := GenS4{F0: genFn0(-9223372036854775808, 32767)}
                                                                _ = v45
                                                            } else {
                                                                v46 := GenS1{F0: 42}
                                                                if true {
                                                                    if (-2147483648 != -883) {
                                                                        if ((true && false) || ((18446744073709551615 == 2) || (!false))) {
                                                                            var v47 int64 = -9223372036854775808
                                                                            v47 = (v47 - (v47 / ((v47) | 1)))
                                                                            if true {
                                                                                switch ("a" + "a") {
                                                                                case "a":
                                                                                    if true {
                                                                                        v48 := []int{(fi41 + (^(fi13))), fi13, (^(fi41))}
                                                                                        _ = v48
                                                                                    }
                                                                                case "\n\t":
                                                                                    v49 := GenS5{F0: (v47 >> ((255) & 63))}
                                                                                    _ = v49
                                                                                case "":
                                                                                    if (!false) {
                                                                                        v3 = 1.5
                                                                                        switch (("hello" + ("" + "\n\t")) + "a") {
                                                                                        case "foo bar":
                                                                                            var v50 uint8 = 80
                                                                                            _ = v50
                                                                                        case "hello":
                                                                                            v47 = (-(v47))
                                                                                        case "":
                                                                                            switch genFn0(genFn0(fi41, 32767), 255) {
                                                                                            case -2147483648:
                                                                                                v51 := GenS6{F0: 42, F1: 4294967295, F2: -1}
                                                                                                _ = v51
                                                                                            }
                                                                                            var v52 string = ("\n\t" + "\n\t")
                                                                                            _ = v52
                                                                                        default:
                                                                                            if false {
                                                                                                var v53 uint = 0
                                                                                                if (-2.25 > 0.0) {
                                                                                                    v47 = (^(v47))
                                                                                                } else {
                                                                                                    v54 := []int{fi13, (fi41 ^ fi41), fi13, int(v53)}
                                                                                                    var v55 int16 = (v1 | (v1 << ((42) & 63)))
                                                                                                    var v56 bool = (42 <= 550)
                                                                                                    _ = v54
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                }
                                                                                                _ = v53
                                                                                            } else {
                                                                                                v1 = (v1 * 42)
                                                                                                v3 += -2.25
                                                                                                fmt.Println(v1)
                                                                                            }
                                                                                        }
                                                                                        switch 982 {
                                                                                        case 42:
                                                                                            v57 := GenS4{F0: (fi13 << ((4294967295) & 63))}
                                                                                            _ = v57
                                                                                        default:
                                                                                            v1 -= v1
                                                                                        }
                                                                                    } else {
                                                                                        v58 := GenS7{F0: 0, F1: (false || (!(!false))), F2: 18446744073709551615}
                                                                                        var v59 int = genFn0(32767, 65535)
                                                                                        v60 := GenS8{F0: "\n\t", F1: -1, F2: -2147483648}
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                }
                                                                                var v61 string = ("hello" + "")
                                                                                _ = v61
                                                                            }
                                                                            v62 := []int{(fi13 | fi41), fi41, genFn0(fi41, 973), 127}
                                                                            _ = v47
                                                                            _ = v62
                                                                        } else {
                                                                            if (2 == 36) {
                                                                                var v63 string = "foo bar"
                                                                                _ = v63
                                                                            }
                                                                            if false {
                                                                                var v64 int = fi41
                                                                                _ = v64
                                                                            } else {
                                                                                v3 -= (float32((float64(0.1) * ((3.14159 - 3.14159) * (3.14159 + ((0.0 + ((((0.0 + 3.14159) + 0.1) / -2.25) + 0.0)) - 0.1))))) / v3)
                                                                                if false {
                                                                                    v3 -= 1.5
                                                                                }
                                                                                v1 = (v1 ^ -32768)
                                                                            }
                                                                            v65 := v2[genFn0(128, 255)]
                                                                            _ = v65
                                                                        }
                                                                    } else {
                                                                        var v66 bool = false
                                                                        _ = v66
                                                                    }
                                                                    v67 := GenS3{F0: 182}
                                                                    fmt.Println(fi13, v3)
                                                                    _ = v67
                                                                } else {
                                                                    var v68 uint32 = 65535
                                                                    v69 := GenS9{F0: false, F1: 127}
                                                                    var v70 uint32 = v68
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                _ = v46
                                                            }
                                                            var v71 uint16 = 651
                                                            v72 := v2[fi41]
                                                            var v73 int = fi13
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    }
                                                }
                                                v74 := GenS10{F0: 128}
                                                _ = fi41
                                                _ = v74
                                            }
                                        default:
                                            v75 := []int{genFn0(32767, 255), fi13, fi13}
                                            _ = v75
                                        }
                                    }
                                default:
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        var v77 string = (("\n\t" + "\n\t") + "\n\t")
                                        _ = fi76
                                        _ = v77
                                    }
                                }
                                if false {
                                    for fi78 := 0; fi78 < 20; fi78++ {
                                        v79 := GenS11{F0: -326}
                                        _ = fi78
                                        _ = v79
                                    }
                                    for fi80 := 0; fi80 < 20; fi80++ {
                                        v2[(fi80 / ((genFn0(fi80, 153)) | 1))] = fi80
                                        v81 := GenS12{F0: 65535, F1: 4294967295}
                                        v82 := []int{genFn0(-2147483648, 0), fi80, (fi80 * (fi13 * (fi80 / ((fi80) | 1))))}
                                        if (42 <= 1) {
                                            var v83 string = "a"
                                            _ = v83
                                        } else {
                                            var v84 uint = 0
                                            v85 := GenS13{F0: 461}
                                            var v86 int32 = 0
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        }
                                        v1 += 1000
                                        v3 -= (v3 - v3)
                                        if (42 >= -2147483648) {
                                            v87 := GenS5{F0: -9223372036854775808}
                                            v1 = (-(v1))
                                            if (!false) {
                                                switch ("foo bar" + "a") {
                                                case "a":
                                                    if (32767 <= 708) {
                                                        if len(v82) > 0 {
                                                            v82[0] = fi13
                                                        }
                                                    }
                                                case "\n\t":
                                                    v88 := GenS14{F0: 4294967295, F1: 42}
                                                    _ = v88
                                                }
                                                if (42 == 162) {
                                                    fmt.Println(fi13)
                                                } else {
                                                    v89 := GenS15{F0: "hello", F1: 46, F2: fi80}
                                                    v1 = (v1 & v1)
                                                    v3 = (v3 - (v3 - v3))
                                                    _ = v89
                                                }
                                            }
                                            v90 := v2[0]
                                            _ = v87
                                            _ = v90
                                        } else {
                                            switch 32768 {
                                            case 302:
                                                fmt.Println(v3, fi80)
                                            case -1:
                                                switch fi80 {
                                                case 4294967295:
                                                    var v91 uint = 148
                                                    _ = v91
                                                case -1:
                                                    v3 = float32((1.5 * -2.25))
                                                default:
                                                    switch 128 {
                                                    case 32768:
                                                        if (!false) {
                                                            if false {
                                                                var v92 uint64 = 18446744073709551615
                                                                var v93 int64 = 966
                                                                _ = v92
                                                                _ = v93
                                                            } else {
                                                                v82 = append(v82, (fi80 | 150))
                                                                v3 += v3
                                                                v1 += -1
                                                            }
                                                            if (!(int64(v1) > 908)) {
                                                                _ = len(v2)
                                                            } else {
                                                                fmt.Println(v3, v1, fi13)
                                                            }
                                                            v3 = v3
                                                        }
                                                    case 2147483647:
                                                        delete(v2, 1)
                                                    case 0:
                                                        v1 += (v1 & -1)
                                                    default:
                                                        var v94 uint8 = 238
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            var v95 uint32 = 2147483648
                                            _ = v95
                                        }
                                        switch "hello" {
                                        case "hello":
                                            v96 := GenS5{F0: 0}
                                            _ = v96
                                        case "a":
                                            var v97 uint64 = 272
                                            _ = v97
                                        }
                                        if false {
                                            v82 = append(v82, (fi80 ^ fi80))
                                            var v98 uint64 = 4294967295
                                            var v99 float32 = v3
                                            v3 -= ((v3 + 3.14159) + (v3 * (((float32(3.14159) + 0.1) / v3) + v99)))
                                            v99 = 0.0
                                            _ = v98
                                            _ = v99
                                        }
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                    }
                                    fmt.Println(v1, v3, fi13)
                                }
                            } else {
                                for fi100 := 0; fi100 < 18; fi100++ {
                                    var v101 int = (fi13 & (fi13 << ((18446744073709551615) & 63)))
                                    _ = fi100
                                    _ = v101
                                }
                                if false {
                                    v3 = 1e10
                                    v1 = v1
                                } else {
                                    v3 = 1e10
                                    fmt.Println(fi13, v1)
                                    switch "foo bar" {
                                    case "a":
                                        v102 := GenS13{F0: 9223372036854775807}
                                        _ = v102
                                    case "":
                                        switch "" {
                                        case "\n\t":
                                            var v103 uint16 = 842
                                            _ = v103
                                        case "foo bar":
                                            for fi104 := 0; fi104 < 16; fi104++ {
                                                v3 = v3
                                                var v105 uint64 = 237
                                                v106 := GenS16{F0: -244}
                                                _ = fi104
                                                _ = v105
                                                _ = v106
                                            }
                                        case "a":
                                            var v107 int = (fi13 >> ((2) & 63))
                                            _ = v107
                                        }
                                    case "hello":
                                        v1 += 0
                                    default:
                                        if (((42 <= v1) || false) || true) {
                                            v1 -= (v1 & (v1 | -1))
                                            v3 = 0.1
                                        } else {
                                            fmt.Println(v3, v3, fi13)
                                        }
                                    }
                                }
                                v108 := []int{fi13}
                                _ = v108
                            }
                            fmt.Println(fi13, v1)
                            v109 := v2[fi13]
                            _ = v109
                        }
                        for fi110 := 0; fi110 < 9; fi110++ {
                            v3 += 1e10
                            switch "\n\t" {
                            case "\n\t":
                                v111 := GenS17{F0: 100}
                                _ = v111
                            case "hello":
                                var v112 int = (fi110 << ((64) & 63))
                                _ = v112
                            case "":
                                var v113 int8 = 42
                                _ = v113
                            default:
                                v1 = v1
                            }
                            _ = fi110
                        }
                        v3 = v3
                    }
                }
                _ = fi13
            }
        }
    case "foo bar":
        v1 = 0
    }
    if true {
        var v114 float32 = 0.1
        for fi115 := 0; fi115 < 16; fi115++ {
            v114 = float32(float64(v3))
            _ = fi115
        }
        for fi116 := 0; fi116 < 2; fi116++ {
            var v117 uint16 = 667
            if (!((fi116 ^ (^(fi116))) > fi116)) {
                v117 = v117
                switch "a" {
                case "":
                    v1 = v1
                case "a":
                    v3 += float32(1.5)
                case "foo bar":
                    for fi118 := 0; fi118 < 15; fi118++ {
                        var v119 uint32 = 4294967295
                        v120 := GenS10{F0: 42}
                        v2[(-(fi118))] = (fi118 - (fi118 / ((fi116) | 1)))
                        _ = fi118
                        _ = v119
                        _ = v120
                    }
                default:
                    var v121 uint64 = 382
                    _ = v121
                }
                v122 := GenS18{F0: 716, F1: 66}
                _ = v122
            }
            _ = fi116
            _ = v117
        }
        _ = v114
    } else {
        v1 = 1
        for fi123 := 0; fi123 < 13; fi123++ {
            fmt.Println(v1, v1)
            if true {
                if (1 <= 421) {
                    var v124 string = "foo bar"
                    v125 := GenS19{F0: 127, F1: v124, F2: 4294967295}
                    if false {
                        delete(v2, fi123)
                        switch genFn0(fi123, 48) {
                        case -9223372036854775808:
                            var v126 bool = (false || true)
                            _ = v126
                        case 127:
                            v1 = v1
                        case 9223372036854775807:
                            switch ((v124 + v124) + v124) {
                            case "a":
                                v124 = "hello"
                            }
                        default:
                            switch fi123 {
                            case 2147483647:
                                v1 -= v1
                            case 128:
                                v3 += (v3 / v3)
                            default:
                                if true {
                                    v127 := GenS1{F0: 65535}
                                    v2[(fi123 | (fi123 - (fi123 & fi123)))] = (fi123 / ((genFn0((-(fi123)), 1)) | 1))
                                    _ = v127
                                } else {
                                    v128 := GenS20{F0: 36}
                                    var v129 float32 = (v3 / 3.14159)
                                    v3 += v3
                                    _ = v128
                                    _ = v129
                                }
                            }
                        }
                    }
                    _ = v124
                    _ = v125
                } else {
                    var v130 int = 128
                    fmt.Println(v3)
                    _ = v130
                }
                switch "hello" {
                case "foo bar":
                    v1 -= (v1 - -1)
                case "\n\t":
                    v1 = v1
                case "a":
                    if false {
                        v131 := GenS15{F0: "a", F1: 2, F2: fi123}
                        v132 := GenS21{F0: (1e10 != 3.14159), F1: 1}
                        var v133 uint16 = 32767
                        _ = v131
                        _ = v132
                        _ = v133
                    }
                }
            } else {
                v3 += v3
                for fi134 := 0; fi134 < 19; fi134++ {
                    v1 = -343
                    v3 -= v3
                    _ = fi134
                }
            }
            _ = fi123
        }
    }
    var v135 int = 32767
    for fi136 := 0; fi136 < 18; fi136++ {
        if true {
            var v137 uint16 = 2
            fmt.Println(v3)
            _ = v137
        }
        _ = fi136
    }
    if true {
        v138 := GenS7{F0: 1, F1: false, F2: 1}
        var v139 uint = 1
        _ = v138
        _ = v139
    } else {
        v140 := []int{}
        _ = v140
    }
    v1 -= v1
    var v141 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v141:", v141)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v135
    _ = v141
    _ = v2
    _ = v3
}
