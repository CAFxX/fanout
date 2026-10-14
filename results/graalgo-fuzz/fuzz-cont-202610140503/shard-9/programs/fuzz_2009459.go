package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint8; F2 int32 }
type GenS9 struct { F0 int64; F1 uint; F2 int32 }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8; F1 int }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS15 struct { F0 int; F1 uint }
type GenS16 struct { F0 bool }
func genFn0(p0 string) uint16 {
    if (!((1.5 - 0.1) < -2.25)) {
        if (!true) {
            var v6 uint16 = 255
            v7 := []int{}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 0, F1: 255, F2: 65535}
                if false {
                    v10 := []int{fi8}
                    var v11 uint8 = 31
                    _ = v10
                    _ = v11
                } else {
                    fmt.Println(fi8, fi8)
                    var v12 string = (p0 + (p0 + ("a" + ((p0 + "hello") + (p0 + p0)))))
                    fmt.Println(p0)
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
        }
        if true {
            if false {
                p0 = ""
                p0 = ""
                if (0 != 185) {
                    if (3.14159 != (0.1 * 0.0)) {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            p0 = p0
                            _ = fi13
                        }
                    }
                } else {
                    fmt.Println(p0)
                    v14 := []int{}
                    switch p0 {
                    case "foo bar":
                        p0 = ("\n\t" + p0)
                    }
                    _ = v14
                }
            } else {
                p0 = p0
                switch p0 {
                case "foo bar":
                    p0 = ("hello" + p0)
                default:
                    p0 = "\n\t"
                }
                var v15 uint = 42
                _ = v15
            }
            var v16 uint = 1
            v16 = v16
            _ = v16
        }
    } else {
        v17 := []int{4294967295, 0}
        _ = v17
    }
    v18 := []int{-1, 32767, 883}
    switch 32768 {
    case -915:
        if (-286 > 0) {
            for fi19 := 0; fi19 < 3; fi19++ {
                p0 = ((p0 + (p0 + ("foo bar" + p0))) + p0)
                p0 = ((p0 + p0) + p0)
                var v20 int8 = 42
                _ = fi19
                _ = v20
            }
        } else {
            v18 = append(v18, 127)
            var v21 uint32 = 4294967295
            _ = v21
        }
    case 127:
        if len(v18) > 0 {
            v18[0] = 32767
        }
    default:
        var v22 uint32 = 2147483648
        _ = v22
    }
    _ = p0
    _ = v18
    return 42
}
func genFn1(p0 uint) int64 {
    var v23 int64 = 9223372036854775807
    fmt.Println(p0, p0)
    switch "foo bar" {
    case "hello":
        var v24 uint8 = 42
        _ = v24
    case "":
        var v25 int8 = 0
        _ = v25
    case "foo bar":
        fmt.Println(p0, p0)
    }
    v26 := []int{32768, 0, -9223372036854775808}
    fmt.Println(p0, v23)
    _ = p0
    _ = v23
    _ = v26
    return (v23 + v23)
}

func main() {
    var v1 float64 = (float64(0.0) * 0.0)
    v2 := make(map[uint32]int)
    var v3 int64 = -9223372036854775808
    var v4 int32 = 2147483647
    var v5 int16 = 805
    v1 = 0.0
    var v27 int32 = (v4 ^ v4)
    var v28 uint32 = 4294967295
    if false {
        fmt.Println(v27)
    } else {
        var v29 int32 = (^(v4))
        var v30 int = -527
        if true {
            fmt.Println(v1, v3)
            switch "" {
            case "\n\t":
                v28 -= 1
            case "a":
                switch (v30 % ((-126) | 1)) {
                case 514:
                    fmt.Println(v3, v27, v4)
                case 923:
                    v31 := GenS2{F0: 127}
                    _ = v31
                case 48:
                    switch "\n\t" {
                    case "hello":
                        v32 := GenS2{F0: 0}
                        _ = v32
                    case "\n\t":
                        v33 := GenS3{F0: ("a" + "\n\t"), F1: (v28 >> ((4294967295) & 63))}
                        _ = v33
                    case "foo bar":
                        _ = len(v2)
                    default:
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v35 := GenS4{F0: 1, F1: (fi34 * (fi34 << ((373) & 63)))}
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            }
        } else {
            v29 = (v29 / (((v27 & 344)) | 1))
            switch (v30 - 127) {
            case 234:
                if true {
                    if (176 > 2) {
                        if (762 < 1) {
                            v36 := GenS1{F0: 42, F1: 741, F2: 897}
                            var v37 float32 = ((-2.25 * -2.25) / (3.14159 + (1e10 / (1e10 + 0.0))))
                            _ = v36
                            _ = v37
                        }
                    } else {
                        fmt.Println(v27, v29)
                    }
                    v38 := []int{-326}
                    for fi39 := 0; fi39 < 18; fi39++ {
                        var v40 uint = 65535
                        if len(v38) > 0 {
                            v38[0] = fi39
                        }
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                } else {
                    var v41 int = (^(v30))
                    v1 += -2.25
                    for fi42 := 0; fi42 < 11; fi42++ {
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v44 := v2[0]
                            v3 -= v3
                            v4 = (v29 & (v27 + (v27 / ((int32(v3)) | 1))))
                            _ = fi43
                            _ = v44
                        }
                        var v45 uint64 = 1
                        v41 = 128
                        _ = fi42
                        _ = v45
                    }
                    _ = v41
                }
            case 32767:
                switch "" {
                case "":
                    if (!false) {
                        var v46 string = "a"
                        if false {
                            if false {
                                v28 = 4294967295
                                v3 += 0
                                if true {
                                    v47 := v2[(v28 + v28)]
                                    _ = v47
                                } else {
                                    v3 = (v3 & (-(v3)))
                                    v46 = "hello"
                                    fmt.Println(v28, v5)
                                }
                            }
                            switch 75 {
                            case 32768:
                                var v48 int64 = (v3 | 147)
                                _ = v48
                            case -103:
                                switch (v30 * (v30 + v30)) {
                                case -9223372036854775808:
                                    if true {
                                        var v49 int8 = 1
                                        _ = v49
                                    } else {
                                        switch (v46 + v46) {
                                        case "foo bar":
                                            v50 := GenS5{F0: 42, F1: v4}
                                            _ = v50
                                        case "":
                                            if (v29 >= v27) {
                                                switch (^(v30)) {
                                                case -9223372036854775808:
                                                    var v51 int64 = (v3 + (v3 / ((v3) | 1)))
                                                    _ = v51
                                                case 0:
                                                    if (false || false) {
                                                        switch ("" + ((v46 + ((((v46 + ("" + v46)) + "foo bar") + (("a" + "foo bar") + "foo bar")) + (v46 + "foo bar"))) + ((((v46 + "hello") + ("a" + v46)) + v46) + ("" + "foo bar")))) {
                                                        case "a":
                                                            v30 = (v30 | (v30 & v30))
                                                        default:
                                                            fmt.Println(v5, v4)
                                                        }
                                                    }
                                                case 4294967295:
                                                    v52 := GenS5{F0: -1, F1: v29}
                                                    _ = v52
                                                }
                                                v53 := GenS6{F0: -105, F1: 127}
                                                for fi54 := 0; fi54 < 2; fi54++ {
                                                    v29 += 929
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch v46 {
                                                        case "foo bar":
                                                            if false {
                                                                v4 = (v4 | v27)
                                                            }
                                                        case "hello":
                                                            if ((18446744073709551615 != 100) || (v46 < ("\n\t" + v46))) {
                                                                if true {
                                                                    v55 := GenS2{F0: -45}
                                                                    _ = v55
                                                                }
                                                                switch "\n\t" {
                                                                case "a":
                                                                    fmt.Println(v4, v4)
                                                                }
                                                            } else {
                                                                v56 := GenS7{F0: genFn1(397)}
                                                                v57 := v2[42]
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                        }
                                                    case "":
                                                        v58 := v2[v28]
                                                        _ = v58
                                                    }
                                                    v59 := GenS8{F0: false, F1: 2, F2: 42}
                                                    _ = fi54
                                                    _ = v59
                                                }
                                                _ = v53
                                            } else {
                                                var v60 bool = true
                                                v60 = v60
                                                v61 := []int{2147483647}
                                                _ = v60
                                                _ = v61
                                            }
                                        case "\n\t":
                                            switch (v30 + (v30 + (v30 << ((1) & 63)))) {
                                            case 2147483647:
                                                var v62 float32 = -2.25
                                                _ = v62
                                            case 4294967295:
                                                if true {
                                                    var v63 bool = false
                                                    for fi64 := 0; fi64 < 18; fi64++ {
                                                        v65 := GenS8{F0: v63, F1: 255, F2: (v27 | (v4 + (v4 * (-(v4)))))}
                                                        var v66 uint8 = 14
                                                        _ = fi64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = v63
                                                } else {
                                                    v30 = v30
                                                    for fi67 := 0; fi67 < 12; fi67++ {
                                                        if (0.0 < (1e10 / -2.25)) {
                                                            fmt.Println(v28)
                                                            var v68 int64 = (v3 - -1)
                                                            for fi69 := 0; fi69 < 20; fi69++ {
                                                                fmt.Println(fi67, v30, v28)
                                                                _ = fi69
                                                            }
                                                            _ = v68
                                                        }
                                                        _ = fi67
                                                    }
                                                }
                                            case -689:
                                                v2[1] = 340
                                            }
                                        default:
                                            fmt.Println(v1)
                                        }
                                        var v70 int16 = (v5 << ((255) & 63))
                                        v71 := GenS9{F0: 42, F1: 4294967295, F2: v29}
                                        _ = v70
                                        _ = v71
                                    }
                                case 128:
                                    fmt.Println(v28)
                                default:
                                    if ((((v46 + ("foo bar" + "foo bar")) + v46) + v46) >= "") {
                                        for fi72 := 0; fi72 < 12; fi72++ {
                                            if (false && false) {
                                                if (!(1 > 126)) {
                                                    fmt.Println(v29)
                                                    var v73 bool = (false && (-142 <= fi72))
                                                    _ = v73
                                                } else {
                                                    v74 := []int{(v30 - fi72), (v30 * (^(v30)))}
                                                    v29 -= v29
                                                    switch fi72 {
                                                    case 127:
                                                        for fi75 := 0; fi75 < 19; fi75++ {
                                                            if (1 > 1) {
                                                                var v76 int8 = -11
                                                                _ = v76
                                                            } else {
                                                                v5 = (v5 * (v5 + (v5 - 1)))
                                                                v1 = v1
                                                                delete(v2, (v28 - 65535))
                                                            }
                                                            if (!(false || (true || false))) {
                                                                v77 := GenS10{F0: (v30 << ((1) & 63)), F1: 255}
                                                                if len(v74) > 0 {
                                                                    v74[0] = v30
                                                                }
                                                                v78 := GenS2{F0: 85}
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            v79 := GenS11{F0: 9223372036854775807}
                                                            _ = fi75
                                                            _ = v79
                                                        }
                                                    default:
                                                        var v80 float32 = 1e10
                                                        _ = v80
                                                    }
                                                    _ = v74
                                                }
                                            } else {
                                                for fi81 := 0; fi81 < 17; fi81++ {
                                                    if (472 != (v30 / (((v30 + (fi72 + fi81))) | 1))) {
                                                        var v82 float64 = 0.0
                                                        if false {
                                                            v83 := GenS12{F0: 1, F1: (fi81 - (fi72 >> ((42) & 63)))}
                                                            if (563 <= 255) {
                                                                fmt.Println(v82, fi81)
                                                            }
                                                            _ = v83
                                                        } else {
                                                            var v84 uint32 = (v28 << ((18446744073709551615) & 63))
                                                            v85 := []int{0, (fi72 & int(v84)), 32768}
                                                            v27 -= v27
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        _ = v82
                                                    }
                                                    if true {
                                                        if false {
                                                            var v86 uint32 = v28
                                                            if true {
                                                                switch v46 {
                                                                case "foo bar":
                                                                    var v87 uint32 = (v86 / (((v28 + v86)) | 1))
                                                                    _ = v87
                                                                case "\n\t":
                                                                    v3 += genFn1(347)
                                                                    v86 -= v28
                                                                    var v88 bool = (!false)
                                                                    _ = v88
                                                                }
                                                                var v89 bool = false
                                                                var v90 int = v30
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                var v91 float64 = 1.5
                                                                _ = v91
                                                            }
                                                            v28 = 2
                                                            _ = v86
                                                        }
                                                        var v92 int = (v30 / (((v30 % ((v30) | 1))) | 1))
                                                        _ = v92
                                                    }
                                                    _ = fi81
                                                }
                                            }
                                            switch v46 {
                                            case "":
                                                for fi93 := 0; fi93 < 18; fi93++ {
                                                    if true {
                                                        v28 = (v28 >> ((65535) & 63))
                                                        var v94 uint16 = 42
                                                        v95 := GenS13{F0: 2147483647}
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v96 := v2[(^(v28))]
                                                        v96 = v96
                                                        var v97 int32 = v4
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    _ = fi93
                                                }
                                            }
                                            _ = fi72
                                        }
                                        v98 := GenS11{F0: v30}
                                        _ = v98
                                    }
                                }
                            case -655:
                                v2[(v28 << ((808) & 63))] = (v30 % (((v30 / (((^(v30))) | 1))) | 1))
                            }
                        }
                        _ = v46
                    } else {
                        v28 = v28
                        for fi99 := 0; fi99 < 8; fi99++ {
                            v29 -= v27
                            v100 := GenS5{F0: 100, F1: v4}
                            _ = fi99
                            _ = v100
                        }
                    }
                case "foo bar":
                    if (true || true) {
                        if false {
                            _ = len(v2)
                            var v101 uint8 = 0
                            _ = v101
                        } else {
                            for fi102 := 0; fi102 < 10; fi102++ {
                                v103 := []int{(-(v30)), fi102, 9223372036854775807, 127}
                                fmt.Println(v28, v3)
                                var v104 uint16 = 255
                                _ = fi102
                                _ = v103
                                _ = v104
                            }
                            var v105 uint32 = (v28 + v28)
                            _ = v105
                        }
                        v106 := GenS4{F0: 42, F1: v30}
                        _ = v106
                    } else {
                        fmt.Println(v30)
                    }
                default:
                    v29 -= v4
                }
            }
        }
        _ = v29
        _ = v30
    }
    v1 = ((((v1 * v1) * v1) * 0.0) / v1)
    if true {
        if (true || false) {
            v107 := GenS3{F0: ("\n\t" + "foo bar"), F1: 42}
            switch "a" {
            case "hello":
                for fi108 := 0; fi108 < 15; fi108++ {
                    v28 = v28
                    v3 -= v3
                    _ = fi108
                }
            case "a":
                v109 := v2[v28]
                _ = v109
            }
            var v110 uint16 = 42
            _ = v107
            _ = v110
        }
        v111 := []int{}
        if true {
            var v112 uint64 = 0
            if true {
                v113 := GenS5{F0: -128, F1: (v4 & (v4 / ((v27) | 1)))}
                var v114 uint = 0
                _ = v113
                _ = v114
            } else {
                if (3.14159 != 1.5) {
                    fmt.Println(v27)
                    for fi115 := 0; fi115 < 9; fi115++ {
                        v111 = append(v111, (^(fi115)))
                        _ = fi115
                    }
                    for fi116 := 0; fi116 < 12; fi116++ {
                        v112 = 4294967295
                        if len(v111) >= 2 {
                            v111 = v111[1:len(v111)]
                        }
                        _ = fi116
                    }
                } else {
                    for fi117 := 0; fi117 < 3; fi117++ {
                        for fi118 := 0; fi118 < 7; fi118++ {
                            if false {
                                var v119 uint = 2
                                var v120 string = ""
                                _ = v119
                                _ = v120
                            } else {
                                v5 = -32768
                                var v121 int8 = 127
                                _ = v121
                            }
                            v112 -= (v112 + v112)
                            _ = fi118
                        }
                        _ = fi117
                    }
                    if true {
                        if (float64(3.14159) <= 0.0) {
                            if true {
                                if false {
                                    if false {
                                        fmt.Println(v1, v28)
                                        for fi122 := 0; fi122 < 8; fi122++ {
                                            if true {
                                                v3 = int64(fi122)
                                                var v123 int8 = 0
                                                _ = v123
                                            } else {
                                                if len(v111) > 0 {
                                                    v111[0] = fi122
                                                }
                                                v28 -= 2147483648
                                            }
                                            _ = fi122
                                        }
                                        for fi124 := 0; fi124 < 5; fi124++ {
                                            if len(v111) > 0 {
                                                v111[0] = (fi124 | (^(fi124)))
                                            }
                                            _ = fi124
                                        }
                                    }
                                    var v125 uint64 = v112
                                    v1 = v1
                                    _ = v125
                                }
                                if (int16(v27) > 28) {
                                    var v126 string = "a"
                                    for fi127 := 0; fi127 < 9; fi127++ {
                                        var v128 float64 = v1
                                        if (127 > 0) {
                                            var v129 int64 = v3
                                            _ = v129
                                        } else {
                                            switch (^(fi127)) {
                                            case 133:
                                                if (!false) {
                                                    switch (-(fi127)) {
                                                    case -619:
                                                        switch (fi127 - 706) {
                                                        case 4294967295:
                                                            var v130 uint16 = 444
                                                            _ = v130
                                                        default:
                                                            var v131 int16 = (-(v5))
                                                            _ = v131
                                                        }
                                                    default:
                                                        if (v126 >= v126) {
                                                            fmt.Println(v128)
                                                            v1 = (v128 * ((float64((0.1 / (((((-2.25 + 0.1) - (0.1 * 0.1)) + float32(1e10)) / 1.0) * -2.25))) + (((float64((0.0 / 1.5)) / ((float64(((((1.5 + 3.14159) * (1e10 * -2.25)) / ((3.14159 - -2.25) / (((float32((1.5 + float64((((1e10 + 0.0) * (0.0 * 0.0)) - (((float32(1e10) / -2.25) * ((1.5 + 3.14159) - 0.1)) + (-2.25 * float32(((float64((((float32(v128) / 1.0) * 1.5) - (1e10 + 1e10))) + 1.5) * v1)))))))) * (1e10 + 0.1)) / (-2.25 / float32(1.5))) / 0.1))) + float32(0.0))) + v1) + ((1e10 + 0.0) / (0.1 * v1)))) / 0.1) + v1)) * v128))
                                                        } else {
                                                            v132 := GenS3{F0: "hello", F1: uint32(v112)}
                                                            _ = v132
                                                        }
                                                    }
                                                    var v133 float64 = v128
                                                    _ = v133
                                                }
                                            case 128:
                                                v134 := GenS14{F0: v28, F1: (!true), F2: 127}
                                                _ = v134
                                            case 32768:
                                                for fi135 := 0; fi135 < 14; fi135++ {
                                                    var v136 uint16 = 255
                                                    v137 := GenS1{F0: 42, F1: v136, F2: 272}
                                                    v138 := GenS14{F0: (^(v28)), F1: (-2.25 <= -2.25), F2: 128}
                                                    v126 = v126
                                                    _ = fi135
                                                    _ = v136
                                                    _ = v137
                                                    _ = v138
                                                }
                                            }
                                        }
                                        if (!(((!true) || (!(false && false))) || false)) {
                                            v4 += 2147483647
                                        } else {
                                            if false {
                                                v139 := GenS14{F0: v28, F1: true, F2: 207}
                                                v111 = append(v111, fi127)
                                                v3 = (v3 >> ((4294967295) & 63))
                                                _ = v139
                                            }
                                            v2[(v28 + 2147483648)] = fi127
                                        }
                                        _ = fi127
                                        _ = v128
                                    }
                                    v4 -= v4
                                    _ = v126
                                } else {
                                    delete(v2, (v28 >> ((169) & 63)))
                                    v3 += genFn1(18446744073709551615)
                                    v27 -= (v4 / (((v27 * (v27 ^ v4))) | 1))
                                }
                                v5 -= (v5 >> ((1) & 63))
                            } else {
                                var v140 float32 = 1e10
                                _ = v140
                            }
                        } else {
                            fmt.Println(v3, v112)
                            for fi141 := 0; fi141 < 19; fi141++ {
                                if false {
                                    var v142 int32 = v4
                                    var v143 int8 = 0
                                    var v144 uint32 = v28
                                    _ = v142
                                    _ = v143
                                    _ = v144
                                } else {
                                    switch fi141 {
                                    case 0:
                                        fmt.Println(fi141, v28, v5)
                                    case 42:
                                        v112 = (^(v112))
                                    default:
                                        var v145 int = (^(fi141))
                                        _ = v145
                                    }
                                    v27 = (v4 & v27)
                                    var v146 int32 = -683
                                    _ = v146
                                }
                                fmt.Println(v28)
                                v28 = (v28 + v28)
                                _ = fi141
                            }
                        }
                        fmt.Println(v112, v1, v3)
                        var v147 uint64 = (v112 | (v112 + (^(v112))))
                        _ = v147
                    } else {
                        for fi148 := 0; fi148 < 8; fi148++ {
                            v112 -= v112
                            var v149 int16 = v5
                            _ = fi148
                            _ = v149
                        }
                        if (0.1 < -2.25) {
                            v27 += (v4 * v4)
                        } else {
                            for fi150 := 0; fi150 < 10; fi150++ {
                                switch 32768 {
                                case 42:
                                    if (v112 != v112) {
                                        v5 -= v5
                                        var v151 int16 = -32768
                                        _ = v151
                                    } else {
                                        v152 := v2[v28]
                                        if ((!((^(v112)) <= (^(v112)))) || (9223372036854775807 < (v3 & genFn1(2)))) {
                                            var v153 int32 = int32(v3)
                                            _ = v153
                                        } else {
                                            v154 := GenS1{F0: 1, F1: 42, F2: 887}
                                            v28 = v28
                                            var v155 int = v152
                                            _ = v154
                                            _ = v155
                                        }
                                        var v156 bool = (!((1.5 + 1e10) == 0.1))
                                        _ = v152
                                        _ = v156
                                    }
                                case 257:
                                    var v157 uint64 = 792
                                    _ = v157
                                }
                                if (true && true) {
                                    v158 := GenS3{F0: "\n\t", F1: 4294967295}
                                    for fi159 := 0; fi159 < 4; fi159++ {
                                        fmt.Println(fi150, v5)
                                        _ = fi159
                                    }
                                    _ = v158
                                }
                                _ = fi150
                            }
                            for fi160 := 0; fi160 < 5; fi160++ {
                                v1 = (float64(float32(0.1)) + v1)
                                _ = fi160
                            }
                            fmt.Println(v27, v28, v4)
                        }
                    }
                }
                if (849 == v28) {
                    _ = len(v2)
                } else {
                    v5 += (v5 | -32768)
                }
                v2[65535] = 195
            }
            _ = v112
        }
        _ = v111
    } else {
        fmt.Println(v1, v28, v4)
        var v161 float32 = (1e10 * float32(v1))
        if ("foo bar" != ("a" + "hello")) {
            if true {
                for fi162 := 0; fi162 < 4; fi162++ {
                    switch fi162 {
                    case -563:
                        v161 -= v161
                    }
                    for fi163 := 0; fi163 < 2; fi163++ {
                        if (!(!false)) {
                            v164 := v2[(v28 % ((686) | 1))]
                            fmt.Println(v5)
                            _ = v164
                        }
                        var v165 uint = 547
                        v166 := GenS15{F0: (fi163 + fi162), F1: (v165 + 2)}
                        _ = fi163
                        _ = v165
                        _ = v166
                    }
                    _ = len(v2)
                    _ = fi162
                }
                if (!((!true) && true)) {
                    fmt.Println(v161)
                }
            }
        } else {
            v1 = (-2.25 * (-2.25 + 0.1))
            v3 -= genFn1(2)
            var v167 uint32 = 4294967295
            _ = v167
        }
        _ = v161
    }
    switch "a" {
    case "foo bar":
        if ("" <= "") {
            var v168 int64 = 0
            fmt.Println(v27, v1, v5)
            v169 := GenS13{F0: (^(v28))}
            _ = v168
            _ = v169
        }
    case "\n\t":
        var v170 uint16 = 667
        _ = v170
    case "hello":
        v171 := GenS11{F0: 628}
        _ = v171
    default:
        switch 128 {
        case 32768:
            switch ("hello" + ("a" + ("" + ("\n\t" + "foo bar")))) {
            case "foo bar":
                if (v27 != v4) {
                    v172 := GenS2{F0: 127}
                    v173 := GenS12{F0: 255, F1: -476}
                    _ = v172
                    _ = v173
                }
            case "a":
                v174 := []int{32767, -1}
                _ = v174
            }
        case 128:
            v2[v28] = 569
        }
    }
    _ = len(v2)
    for fi175 := 0; fi175 < 4; fi175++ {
        switch 42 {
        case -465:
            fmt.Println(v4)
        case -105:
            v4 += v4
        default:
            v3 -= (v3 + v3)
        }
        _ = fi175
    }
    v176 := GenS16{F0: ("hello" == "a")}
    v3 += (v3 - v3)
    if false {
        var v177 float32 = 1.5
        var v178 int = -1
        _ = v177
        _ = v178
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v176
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
