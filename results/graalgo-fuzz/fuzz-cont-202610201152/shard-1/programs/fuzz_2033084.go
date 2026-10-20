package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 int64; F1 int16; F2 bool }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 uint32; F1 int64; F2 int8 }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint64; F1 uint32 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS23 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 uint8, p1 uint) uint8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int32 = 149
        v6 := []int{}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 416
    fmt.Println(p1, p1)
    v7 += (v7 % (((v7 % ((4294967295) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return uint8(p1)
}
func genFn1(p0 int32, p1 uint32, p2 uint8) bool {
    v8 := GenS1{F0: 342, F1: (p0 & p0)}
    v9 := []int{32767, 1, 128, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    v3 := make(map[string]bool)
    if true {
        switch ((v2 + v2) + (v2 + v2)) {
        case "":
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := GenS2{F0: fi10}
                if (-1 >= int64(fi10)) {
                    var v12 int64 = 42
                    if (2 == 538) {
                        v3[""] = false
                        v2 = ""
                        v13 := GenS3{F0: (-(fi10)), F1: 255}
                        _ = v13
                    }
                    v2 = (((v2 + (((v2 + ((v2 + v2) + ("\n\t" + "hello"))) + "") + v2)) + ("" + v2)) + "")
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        }
    } else {
        fmt.Println(v2)
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS2{F0: fi14}
            v16 := []int{(fi14 | (fi14 ^ (fi14 ^ fi14))), fi14}
            var v17 int64 = 0
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v2 = ""
            _ = fi18
        }
    }
    if false {
        delete(v3, ("foo bar" + (v2 + v2)))
        var v19 uint64 = 0
        v19 += v19
        _ = v19
    }
    switch v2 {
    case "a":
        var v20 bool = true
        _ = v20
    case "hello":
        var v21 bool = false
        _ = v21
    }
    v2 = v2
    if true {
        if false {
            var v22 uint32 = 2
            _ = v22
        } else {
            v23 := GenS2{F0: 0}
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint8 = 199
            _ = fi24
            _ = v25
        }
    }
    switch 0 {
    case 42:
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 uint64 = 4294967295
            v28 := []int{fi26, 42}
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        switch (((((v2 + (v2 + v2)) + "") + (v2 + v2)) + (v2 + v2)) + v2) {
        case "hello":
            var v29 uint32 = 2
            _ = v29
        case "foo bar":
            if ((637 < 42) || (1.5 > 3.14159)) {
                v2 = v2
            }
        case "a":
            v2 = v2
        }
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if (false || ((!true) || false)) {
            var v31 uint = 1
            _ = v31
        }
        var v32 bool = (32767 <= 252)
        _ = fi30
        _ = v32
    }
    var v33 float64 = float64(3.14159)
    v34 := GenS4{F0: 1, F1: "\n\t"}
    for fi35 := 0; fi35 < 20; fi35++ {
        v33 -= v33
        fmt.Println(fi35, v2, v33)
        _ = fi35
    }
    v2 = v2
    var v36 uint64 = 2
    if false {
        for fi37 := 0; fi37 < 1; fi37++ {
            v38 := GenS2{F0: (fi37 + (fi37 * -805))}
            _ = fi37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            switch (v2 + (v2 + v2)) {
            case "foo bar":
                v2 = (v2 + v2)
            default:
                var v41 int64 = 409
                if ((v41 * v41) > v41) {
                    switch fi39 {
                    case -2147483648:
                        var v42 int8 = 58
                        _ = v42
                    default:
                        v36 = (v36 >> ((387) & 63))
                    }
                    var v43 float64 = v33
                    v43 -= (v33 / v43)
                    _ = v43
                } else {
                    var v44 uint8 = 1
                    if true {
                        var v45 float64 = v33
                        switch fi39 {
                        case 1:
                            if true {
                                if false {
                                    v33 += (v45 - v33)
                                    fmt.Println(v2)
                                }
                                v45 = float64(float32(v45))
                                v44 += v44
                                v46 := []int{32768, int(v36), int(v36), fi40}
                                var v47 int32 = -584
                                _ = v46
                                _ = v47
                            }
                        case 32768:
                            var v48 uint64 = v36
                            _ = v48
                        case 4294967295:
                            fmt.Println(v33, v44, v36)
                        default:
                            v44 = genFn0((v44 % (((v44 * 233)) | 1)), 4294967295)
                        }
                        if (!false) {
                            v45 = v33
                        }
                        switch fi40 {
                        case 0:
                            v45 -= (v33 + v33)
                        case 9223372036854775807:
                            switch "a" {
                            case "a":
                                if (0 > v44) {
                                    if true {
                                        var v49 string = ((v2 + v2) + ("a" + v2))
                                        if false {
                                            v50 := []int{fi39}
                                            if true {
                                                v44 -= (^(v44))
                                                if true {
                                                    var v51 int32 = -740
                                                    var v52 uint16 = 1
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v44 += (^(v44))
                                                    v53 := GenS2{F0: fi40}
                                                    switch ((v2 + v2) + v49) {
                                                    case "\n\t":
                                                        v41 = v41
                                                    case "foo bar":
                                                        v36 += v36
                                                    default:
                                                        var v54 float32 = 0.1
                                                        _ = v54
                                                    }
                                                    _ = v53
                                                }
                                                var v55 int16 = -29
                                                _ = v55
                                            } else {
                                                if ((104 > 255) || false) {
                                                    delete(v1, -9223372036854775808)
                                                    _ = len(v50)
                                                } else {
                                                    v50 = append(v50, (fi40 % (((^(fi40))) | 1)))
                                                }
                                                v3[(((v2 + v49) + "hello") + (v49 + ("\n\t" + (((((v2 + "") + v2) + v49) + v49) + v2))))] = false
                                                v41 -= v41
                                                var v56 bool = (0 < 255)
                                                var v57 int16 = -197
                                                _ = v56
                                                _ = v57
                                            }
                                            _ = v50
                                        } else {
                                            if true {
                                                fmt.Println(v41)
                                            } else {
                                                v58 := GenS5{F0: (v36 / ((v36) | 1))}
                                                switch (fi40 >> ((0) & 63)) {
                                                case 32768:
                                                    var v59 int8 = -26
                                                    _ = v59
                                                case -9223372036854775808:
                                                    v36 = v36
                                                default:
                                                    v60 := GenS5{F0: v36}
                                                    _ = v60
                                                }
                                                _ = v58
                                            }
                                            if true {
                                                switch (fi39 / ((2147483647) | 1)) {
                                                case 2147483647:
                                                    if (!((v44 << ((790) & 63)) >= (v44 >> ((65535) & 63)))) {
                                                        var v61 uint = 0
                                                        _ = v61
                                                    } else {
                                                        var v62 uint16 = 0
                                                        if true {
                                                            if true {
                                                                v63 := GenS3{F0: (fi40 | fi40), F1: 263}
                                                                if (true && (!false)) {
                                                                    var v64 uint8 = 127
                                                                    if true {
                                                                        switch v2 {
                                                                        case "\n\t":
                                                                            v65 := GenS2{F0: fi40}
                                                                            _ = v65
                                                                        default:
                                                                            var v66 float64 = -2.25
                                                                            _ = v66
                                                                        }
                                                                    } else {
                                                                        switch v49 {
                                                                        case "hello":
                                                                            v67 := GenS6{F0: (v44 | (v44 >> ((326) & 63)))}
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                var v68 int = -2147483648
                                                                                _ = v68
                                                                            case "a":
                                                                                fmt.Println(v2)
                                                                            case "foo bar":
                                                                                var v69 int32 = 0
                                                                                _ = v69
                                                                            }
                                                                            _ = v67
                                                                        }
                                                                        var v70 bool = ((("hello" + "") + "a") < v2)
                                                                        _ = v70
                                                                    }
                                                                    _ = v64
                                                                } else {
                                                                    v71 := []int{fi40, 0}
                                                                    _ = v71
                                                                }
                                                                v41 += v41
                                                                var v72 string = (v49 + (v2 + v2))
                                                                var v73 float32 = 1e10
                                                                delete(v3, "\n\t")
                                                                _ = v63
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            switch "hello" {
                                                            case "":
                                                                if false {
                                                                    var v74 uint32 = 2147483647
                                                                    if true {
                                                                        v75 := GenS7{F0: v49}
                                                                        v45 = 1e10
                                                                        switch 4294967295 {
                                                                        case 9223372036854775807:
                                                                            switch "a" {
                                                                            case "":
                                                                                v36 = (^(v36))
                                                                            default:
                                                                                if true {
                                                                                    var v76 float32 = (((-2.25 / float32((1e10 * v33))) - ((0.1 / 1.5) - 3.14159)) / 1.0)
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            switch ((((v2 + v49) + v2) + "\n\t") + v2) {
                                                                            case "foo bar":
                                                                                var v77 bool = (uint8(v36) > (v44 << ((65535) & 63)))
                                                                                _ = v77
                                                                            case "":
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                        delete(v3, "a")
                                                                        _ = v75
                                                                    }
                                                                    _ = v74
                                                                } else {
                                                                    v78 := GenS2{F0: (fi39 * (fi40 % ((32767) | 1)))}
                                                                    v1[42] = (32 != int8(v36))
                                                                    _ = v78
                                                                }
                                                            case "hello":
                                                                if (785 >= 2147483648) {
                                                                    v41 -= v41
                                                                }
                                                            case "foo bar":
                                                                switch "" {
                                                                case "\n\t":
                                                                    if false {
                                                                        v49 = v2
                                                                        v49 = "foo bar"
                                                                    } else {
                                                                        v79 := GenS1{F0: (v41 + (v41 * (-(v41)))), F1: 2147483647}
                                                                        v80 := GenS8{F0: genFn0((^(v44)), 4294967295), F1: v2}
                                                                        var v81 uint = 1
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                case "foo bar":
                                                                    switch (fi40 % ((fi40) | 1)) {
                                                                    case 4294967295:
                                                                        if false {
                                                                            v82 := GenS8{F0: (v44 / ((v44) | 1)), F1: (("" + (("" + v2) + ("hello" + v2))) + "\n\t")}
                                                                            fmt.Println(fi39, v2)
                                                                            var v83 int64 = v41
                                                                            v84 := GenS9{F0: v83, F1: -32768, F2: true}
                                                                            v41 = (v41 & (v41 + int64(v62)))
                                                                            var v85 int = fi39
                                                                            _ = v82
                                                                            _ = v83
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                    default:
                                                                        v86 := GenS3{F0: int(v41), F1: (^(v62))}
                                                                        _ = v86
                                                                    }
                                                                }
                                                            default:
                                                                if (344 == 255) {
                                                                    v87 := GenS1{F0: v41, F1: 0}
                                                                    _ = v87
                                                                } else {
                                                                    if true {
                                                                        v36 = v36
                                                                    }
                                                                }
                                                            }
                                                            v45 = float64(0.1)
                                                        } else {
                                                            v88 := GenS6{F0: v44}
                                                            switch (fi39 | (fi39 >> ((2) & 63))) {
                                                            case 0:
                                                                switch fi40 {
                                                                case 127:
                                                                    v33 += (0.0 / (v33 * ((((float64((1e10 + 1.5)) - (0.1 * v45)) * 1.5) * (v33 / float64(0.1))) / v33)))
                                                                case 0:
                                                                    v33 -= 0.0
                                                                }
                                                            default:
                                                                switch (v2 + v2) {
                                                                case "\n\t":
                                                                    v45 -= v45
                                                                default:
                                                                    v45 = v45
                                                                }
                                                            }
                                                            _ = v88
                                                        }
                                                        v89 := GenS10{F0: false, F1: 2147483647}
                                                        v90 := []int{(^(fi40)), (fi39 << ((uint(v44)) & 63)), fi39}
                                                        _ = v62
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                case -2147483648:
                                                    fmt.Println(v49, v45, v2)
                                                }
                                                v2 = ""
                                            }
                                            if true {
                                                var v91 int64 = (^(v41))
                                                _ = v91
                                            } else {
                                                if (v36 < v36) {
                                                    if (!true) {
                                                        if true {
                                                            var v92 uint64 = 2
                                                            _ = v92
                                                        } else {
                                                            if (3.14159 > (3.14159 * -2.25)) {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if true {
                                                                        if true {
                                                                            if (!(!false)) {
                                                                                var v93 int = -945
                                                                                _ = v93
                                                                            }
                                                                        } else {
                                                                            v94 := []int{42, -2147483648, 42}
                                                                            switch 42 {
                                                                            case 127:
                                                                                v95 := GenS1{F0: v41, F1: -373}
                                                                                fmt.Println(v36, v36, fi39)
                                                                                v33 = v33
                                                                                _ = v95
                                                                            default:
                                                                                var v96 int32 = -2147483648
                                                                                _ = v96
                                                                            }
                                                                            switch v49 {
                                                                            case "":
                                                                                v49 = v2
                                                                            }
                                                                            _ = v94
                                                                        }
                                                                    } else {
                                                                        v45 = (3.14159 * 0.0)
                                                                        v36 = v36
                                                                        var v97 uint64 = v36
                                                                        _ = v97
                                                                    }
                                                                case "hello":
                                                                    var v98 uint = 42
                                                                    _ = v98
                                                                case "a":
                                                                    v45 = v45
                                                                }
                                                                v2 = (v2 + "\n\t")
                                                            }
                                                            v45 = 0.0
                                                            v99 := v3[v49]
                                                            _ = v99
                                                        }
                                                        fmt.Println(v49, v36)
                                                    } else {
                                                        if false {
                                                            var v100 uint32 = 65535
                                                            _ = v100
                                                        } else {
                                                            fmt.Println(v36, v36, fi40)
                                                            var v101 int8 = 1
                                                            var v102 uint8 = 137
                                                            if ("foo bar" <= "foo bar") {
                                                                var v103 float32 = 1e10
                                                                v104 := []int{(fi40 ^ (fi39 / (((fi40 + fi39)) | 1))), 735, 0, fi40}
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                v41 = -9223372036854775808
                                                                var v105 int = 9223372036854775807
                                                                v44 = 42
                                                                if (((v49 + "foo bar") + ((((v49 + "hello") + "a") + (v2 + "\n\t")) + v2)) < ("a" + "a")) {
                                                                    var v106 uint8 = (v102 | genFn0(genFn0((v102 % ((uint8(v36)) | 1)), 0), 4294967295))
                                                                    v36 = (v36 * (v36 >> ((65535) & 63)))
                                                                    var v107 int32 = -759
                                                                    _ = v106
                                                                    _ = v107
                                                                } else {
                                                                    v102 = (v102 << ((42) & 63))
                                                                }
                                                                v101 = (^(v101))
                                                                var v108 int8 = -1
                                                                if false {
                                                                    var v109 int = (^(fi40))
                                                                    if false {
                                                                        var v110 uint = 4294967295
                                                                        fmt.Println(v108, fi40, v102)
                                                                        if false {
                                                                            var v111 uint32 = 65535
                                                                            v45 = 0.1
                                                                            v112 := []int{fi40, v109}
                                                                            var v113 string = "hello"
                                                                            _ = v111
                                                                            _ = v112
                                                                            _ = v113
                                                                        }
                                                                        _ = v110
                                                                    }
                                                                    _ = v109
                                                                } else {
                                                                    v41 -= int64(v105)
                                                                }
                                                                delete(v1, (fi40 - (fi40 / ((fi40) | 1))))
                                                                _ = v105
                                                                _ = v108
                                                            }
                                                            switch (v2 + v49) {
                                                            case "hello":
                                                                var v114 uint64 = (^(v36))
                                                                _ = v114
                                                            case "\n\t":
                                                                v45 = 1e10
                                                            case "foo bar":
                                                                switch "" {
                                                                case "hello":
                                                                    var v115 int32 = 42
                                                                    _ = v115
                                                                case "\n\t":
                                                                    fmt.Println(v45)
                                                                    var v116 float64 = 1.5
                                                                    if true {
                                                                        v117 := GenS11{F0: 1, F1: 2147483647, F2: uint16(v102)}
                                                                        switch v2 {
                                                                        case "":
                                                                            if (-1 >= (^(v41))) {
                                                                                if false {
                                                                                    v2 = v2
                                                                                    var v118 float64 = 3.14159
                                                                                    _ = v118
                                                                                }
                                                                                v2 = (v49 + "hello")
                                                                                var v119 bool = false
                                                                                fmt.Println(v116)
                                                                                _ = v119
                                                                            }
                                                                        case "hello":
                                                                            if (3.14159 == 0.1) {
                                                                                var v120 int8 = (-(v101))
                                                                                fmt.Println(v116, v33)
                                                                                var v121 int32 = 0
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                        default:
                                                                            v122 := GenS11{F0: v36, F1: 673, F2: 640}
                                                                            v123 := GenS2{F0: int(v102)}
                                                                            _ = v122
                                                                            _ = v123
                                                                        }
                                                                        if true {
                                                                            v44 = 128
                                                                        } else {
                                                                            v44 += genFn0(v102, 527)
                                                                            var v124 uint64 = (v36 - v36)
                                                                            _ = v124
                                                                        }
                                                                        _ = v117
                                                                    } else {
                                                                        var v125 float32 = ((-2.25 - 1.5) / (3.14159 * (float32(v116) + 0.1)))
                                                                        var v126 int = fi39
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                    _ = v116
                                                                default:
                                                                    fmt.Println(v49, v41, v49)
                                                                    v36 -= v36
                                                                }
                                                            default:
                                                                v127 := GenS3{F0: (-(fi40)), F1: 255}
                                                                _ = v127
                                                            }
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                        v128 := GenS6{F0: v44}
                                                        v49 = v49
                                                        _ = v128
                                                    }
                                                    var v129 uint = 18446744073709551615
                                                    _ = v129
                                                }
                                                v41 = v41
                                                v41 += v41
                                                v36 = (v36 | 4294967295)
                                            }
                                        }
                                        _ = len(v3)
                                        _ = v49
                                    }
                                } else {
                                    if false {
                                        switch (-(fi40)) {
                                        case 9223372036854775807:
                                            var v130 float32 = ((1e10 * (0.0 - ((0.0 + -2.25) / 1.5))) - 1e10)
                                            _ = v130
                                        }
                                        if false {
                                            v131 := []int{fi39, (^(fi39)), -53, fi40}
                                            v131 = append(v131, fi40)
                                            var v132 uint = 142
                                            _ = v131
                                            _ = v132
                                        }
                                        if ((!false) || true) {
                                            v33 = ((float64(3.14159) / 1.0) * 1e10)
                                        }
                                        var v133 float32 = ((0.1 - ((-2.25 / -2.25) * -2.25)) / 1.0)
                                        _ = v133
                                    } else {
                                        v134 := []int{(fi39 - (fi40 << ((302) & 63)))}
                                        if (false || (!(v2 < v2))) {
                                            v45 -= float64(1e10)
                                            if (fi40 >= (fi40 << ((2) & 63))) {
                                                var v135 uint8 = (v44 & v44)
                                                var v136 float32 = 0.0
                                                _ = v135
                                                _ = v136
                                            }
                                            switch (fi40 | fi40) {
                                            case -2147483648:
                                                v45 = v33
                                            case 762:
                                                fmt.Println(v2, v44, fi39)
                                                v137 := v1[9223372036854775807]
                                                _ = v137
                                            }
                                        }
                                        fmt.Println(v33)
                                        _ = v134
                                    }
                                    fmt.Println(v45, v2)
                                    if true {
                                        if true {
                                            if (2 >= 756) {
                                                v41 = (-(v41))
                                                if false {
                                                    v1[fi39] = false
                                                    var v138 int = (fi40 % ((128) | 1))
                                                    _ = v138
                                                }
                                                if (!false) {
                                                    v41 = v41
                                                    v44 = (v44 - (v44 - 42))
                                                    delete(v1, fi39)
                                                } else {
                                                    v139 := []int{fi40, (fi40 + (^(fi39))), (fi40 * -9223372036854775808)}
                                                    if (!true) {
                                                        var v140 uint8 = (v44 - v44)
                                                        v141 := GenS12{F0: 0}
                                                        if (367 <= 2147483648) {
                                                            v3[v2] = true
                                                            _ = len(v1)
                                                        } else {
                                                            if false {
                                                                var v142 string = ((("\n\t" + "a") + ((v2 + (v2 + v2)) + v2)) + "\n\t")
                                                                var v143 uint32 = 2
                                                                v144 := GenS2{F0: -2147483648}
                                                                _ = v142
                                                                _ = v143
                                                                _ = v144
                                                            } else {
                                                                v45 = v33
                                                                var v145 uint8 = (v44 << ((1) & 63))
                                                                switch (fi40 * 4294967295) {
                                                                case -558:
                                                                    if ((v41 == (v41 | -1)) || true) {
                                                                        switch 32767 {
                                                                        case 147:
                                                                            if false {
                                                                                v140 = (^(v140))
                                                                            }
                                                                        default:
                                                                            var v146 float64 = (v45 - v33)
                                                                            _ = v146
                                                                        }
                                                                        var v147 bool = false
                                                                        var v148 int16 = -32768
                                                                        var v149 uint64 = (v36 >> ((4294967295) & 63))
                                                                        v45 -= ((0.1 - v45) / (v45 + v33))
                                                                        _ = v147
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                case -975:
                                                                    _ = len(v3)
                                                                }
                                                                _ = v145
                                                            }
                                                            fmt.Println(v36, v140)
                                                            switch v2 {
                                                            case "a":
                                                                var v150 uint8 = (v140 - 183)
                                                                v41 = 1
                                                                _ = v150
                                                            case "foo bar":
                                                                if ((fi39 < (-(fi40))) && false) {
                                                                    switch fi39 {
                                                                    case 128:
                                                                        if (2147483647 > 1) {
                                                                            switch v2 {
                                                                            case "a":
                                                                                switch (fi40 << ((711) & 63)) {
                                                                                case 128:
                                                                                    v140 -= genFn0(v140, 0)
                                                                                case 4294967295:
                                                                                    switch int(v36) {
                                                                                    case -2147483648:
                                                                                        _ = len(v139)
                                                                                    case 32767:
                                                                                        v140 -= (v44 - 255)
                                                                                    default:
                                                                                        v151 := GenS13{F0: (^(v41)), F1: (fi40 == fi39)}
                                                                                        _ = v151
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                v152 := GenS14{F0: 1, F1: (v41 - (-(v41))), F2: 36}
                                                                                _ = v152
                                                                            }
                                                                            v41 -= v41
                                                                        } else {
                                                                            fmt.Println(v140)
                                                                            v153 := GenS11{F0: v36, F1: 2147483648, F2: 995}
                                                                            _ = v153
                                                                        }
                                                                    case 2147483647:
                                                                        var v154 uint16 = 0
                                                                        switch -1 {
                                                                        case 4294967295:
                                                                            var v155 int = fi39
                                                                            _ = v155
                                                                        case 42:
                                                                            fmt.Println(v140, v2, fi39)
                                                                        case -1:
                                                                            v140 = uint8(v36)
                                                                            delete(v3, (v2 + v2))
                                                                            if ("a" == "a") {
                                                                                v156 := v1[(fi39 & 42)]
                                                                                switch -2147483648 {
                                                                                case -9223372036854775808:
                                                                                    var v157 float64 = ((v33 * v45) * (v45 * v45))
                                                                                    _ = v157
                                                                                default:
                                                                                    switch (fi39 | (fi39 >> ((2) & 63))) {
                                                                                    case 1:
                                                                                        v45 += 0.1
                                                                                    }
                                                                                }
                                                                                _ = v156
                                                                            }
                                                                            v158 := GenS5{F0: v36}
                                                                            delete(v3, v2)
                                                                            v139 = append(v139, fi39)
                                                                            if false {
                                                                                if false {
                                                                                    v159 := GenS15{F0: genFn0(v44, 1), F1: v2}
                                                                                    var v160 uint64 = v36
                                                                                    _ = v159
                                                                                    _ = v160
                                                                                } else {
                                                                                    var v161 uint8 = (v44 - (v140 | 112))
                                                                                    v44 += uint8(v36)
                                                                                    v44 -= (v161 * (v44 * 1))
                                                                                    _ = v161
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    v41 += 935
                                                                                    if true {
                                                                                        v162 := GenS16{F0: 79}
                                                                                        v163 := v1[fi40]
                                                                                        v140 = 197
                                                                                        _ = v162
                                                                                        _ = v163
                                                                                    }
                                                                                } else {
                                                                                    v164 := GenS17{F0: (^(v44))}
                                                                                    var v165 uint16 = v154
                                                                                    var v166 float64 = 1e10
                                                                                    _ = v164
                                                                                    _ = v165
                                                                                    _ = v166
                                                                                }
                                                                                if len(v139) >= 2 {
                                                                                    v139 = v139[1:len(v139)]
                                                                                }
                                                                                var v167 int = (^(fi39))
                                                                                v36 = v36
                                                                                _ = v167
                                                                            }
                                                                            var v168 uint32 = 2
                                                                            _ = v158
                                                                            _ = v168
                                                                        default:
                                                                            if (false && true) {
                                                                                fmt.Println(v33, v154)
                                                                                _ = len(v3)
                                                                            } else {
                                                                                v36 = (v36 >> ((1) & 63))
                                                                            }
                                                                        }
                                                                        if len(v139) >= 2 {
                                                                            v139 = v139[1:len(v139)]
                                                                        }
                                                                        _ = v154
                                                                    default:
                                                                        var v169 string = "a"
                                                                        _ = v169
                                                                    }
                                                                    fmt.Println(fi40)
                                                                    fmt.Println(v41)
                                                                } else {
                                                                    v170 := v3[(v2 + v2)]
                                                                    _ = v170
                                                                }
                                                                var v171 uint16 = 685
                                                                switch v2 {
                                                                case "":
                                                                    var v172 string = (("foo bar" + v2) + v2)
                                                                    _ = v172
                                                                case "a":
                                                                    var v173 int8 = 100
                                                                    _ = v173
                                                                case "hello":
                                                                    var v174 int32 = 0
                                                                    _ = v174
                                                                default:
                                                                    v45 = ((1.5 - v33) - v45)
                                                                }
                                                                _ = v171
                                                            default:
                                                                var v175 string = "\n\t"
                                                                _ = v175
                                                            }
                                                        }
                                                        fmt.Println(v140, fi40)
                                                        var v176 uint = 2
                                                        v44 -= uint8(v176)
                                                        _ = v140
                                                        _ = v141
                                                        _ = v176
                                                    }
                                                    _ = len(v3)
                                                    var v177 int64 = v41
                                                    v178 := GenS18{F0: 0}
                                                    _ = v139
                                                    _ = v177
                                                    _ = v178
                                                }
                                            }
                                            fmt.Println(v45, v41)
                                        }
                                        switch fi40 {
                                        case 127:
                                            if (!((true && false) || true)) {
                                                var v179 uint64 = v36
                                                var v180 uint = 2
                                                v36 -= 4294967295
                                                var v181 float64 = v33
                                                var v182 int64 = 1
                                                _ = v179
                                                _ = v180
                                                _ = v181
                                                _ = v182
                                            } else {
                                                switch v2 {
                                                case "\n\t":
                                                    if ((true && true) || ((v44 * 5) < 84)) {
                                                        var v183 string = (("foo bar" + v2) + ("hello" + v2))
                                                        _ = v183
                                                    } else {
                                                        var v184 uint16 = 0
                                                        v45 = (v45 / (v33 - v45))
                                                        var v185 uint8 = 2
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                case "":
                                                    var v186 uint = 65535
                                                    _ = v186
                                                default:
                                                    var v187 int64 = 0
                                                    _ = v187
                                                }
                                                v188 := []int{fi39}
                                                _ = v188
                                            }
                                        default:
                                            switch fi40 {
                                            case 706:
                                                v189 := []int{fi40}
                                                _ = v189
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                v190 := GenS19{F0: (v36 << ((465) & 63)), F1: 2}
                                _ = v190
                            }
                        case -648:
                            switch fi39 {
                            case 325:
                                if (false && (1.5 <= float32(v45))) {
                                    if false {
                                        v33 += v33
                                    } else {
                                        var v191 uint = 0
                                        _ = v191
                                    }
                                    fmt.Println(v41)
                                    if ((false && true) || true) {
                                        var v192 int16 = 1
                                        if true {
                                            fmt.Println(fi40, v36)
                                            v193 := []int{32768, (fi39 + (fi40 & (fi40 << ((0) & 63)))), int(v192), fi40}
                                            var v194 int = (fi39 / ((113) | 1))
                                            switch (v194 & 54) {
                                            case -257:
                                                v45 = v33
                                            case 9223372036854775807:
                                                v195 := GenS15{F0: (v44 >> ((4294967295) & 63)), F1: v2}
                                                _ = v195
                                            }
                                            v194 += -1
                                            _ = v193
                                            _ = v194
                                        }
                                        switch (fi40 & fi39) {
                                        case 756:
                                            v33 = 3.14159
                                        case -961:
                                            v36 = 4294967295
                                            if true {
                                                switch fi40 {
                                                case 0:
                                                    if ((-2.25 * float32(0.0)) == -2.25) {
                                                        v196 := []int{(^(fi40)), fi40, fi40, 128}
                                                        v197 := GenS3{F0: 32768, F1: 2}
                                                        _ = v196
                                                        _ = v197
                                                    } else {
                                                        fmt.Println(fi40)
                                                        var v198 int8 = -128
                                                        var v199 uint8 = (v44 & (v44 - (v44 >> ((1) & 63))))
                                                        _ = v198
                                                        _ = v199
                                                    }
                                                }
                                                v200 := GenS10{F0: false, F1: 74}
                                                v201 := GenS14{F0: 458, F1: v41, F2: -101}
                                                v192 += 336
                                                v33 = v45
                                                _ = v200
                                                _ = v201
                                            } else {
                                                delete(v1, int(v44))
                                            }
                                        case 4294967295:
                                            var v202 int8 = -7
                                            _ = v202
                                        default:
                                            v2 = "hello"
                                        }
                                        _ = v192
                                    } else {
                                        if (false || (true && (860 != 65535))) {
                                            v203 := GenS14{F0: 156, F1: 0, F2: 100}
                                            _ = v203
                                        }
                                    }
                                }
                                v204 := GenS14{F0: 65535, F1: -1, F2: 0}
                                _ = v204
                            default:
                                v205 := []int{(fi40 % ((fi40) | 1)), (fi39 * fi40), fi40, fi40}
                                _ = v205
                            }
                        }
                        v2 = v2
                        _ = v45
                    } else {
                        v44 += (v44 * v44)
                        fmt.Println(v33, fi39, v2)
                        var v206 int64 = v41
                        switch (v2 + (("\n\t" + "foo bar") + "hello")) {
                        case "\n\t":
                            v207 := GenS20{F0: 32767}
                            _ = v207
                        default:
                            if false {
                                var v208 int8 = 41
                                v209 := GenS2{F0: -236}
                                v1[(fi39 + 857)] = true
                                var v210 int8 = int8(fi40)
                                _ = v208
                                _ = v209
                                _ = v210
                            } else {
                                var v211 uint32 = 543
                                v212 := GenS21{F0: v211}
                                _ = v211
                                _ = v212
                            }
                        }
                        delete(v3, (((v2 + ("hello" + v2)) + (v2 + v2)) + ((v2 + (v2 + ((((((((("a" + (v2 + (((("" + (v2 + "hello")) + (v2 + (v2 + (v2 + "hello")))) + v2) + (v2 + v2)))) + v2) + v2) + v2) + v2) + v2) + "") + v2) + "a"))) + (((v2 + (v2 + (v2 + (v2 + v2)))) + (("" + "a") + (v2 + ""))) + ((v2 + v2) + (v2 + ""))))))
                        _ = v206
                    }
                    var v213 uint32 = 79
                    _ = v213
                    _ = v44
                }
                fmt.Println(v41, v36)
                switch fi40 {
                case 32767:
                    v214 := v1[fi40]
                    v215 := v3[("\n\t" + v2)]
                    _ = v214
                    _ = v215
                case 32768:
                    v33 += v33
                default:
                    if false {
                        v216 := []int{(fi39 >> ((2) & 63))}
                        _ = v216
                    }
                }
                _ = v41
            }
            v33 = v33
            if false {
                if ("a" >= v2) {
                    var v217 float64 = v33
                    if false {
                        if true {
                            v218 := GenS19{F0: 4294967295, F1: 813}
                            if (((true || false) && ((1.5 / 1.0) <= (0.1 * 0.0))) && (42 < 516)) {
                                var v219 float32 = float32(1.5)
                                if (true && (1e10 <= v217)) {
                                    fmt.Println(v2)
                                    v33 = v217
                                }
                                var v220 string = (v2 + v2)
                                _ = v219
                                _ = v220
                            } else {
                                fmt.Println(v217)
                                if false {
                                    var v221 int64 = -1
                                    var v222 int = int(v36)
                                    _ = v221
                                    _ = v222
                                } else {
                                    if true {
                                        switch fi39 {
                                        case -9223372036854775808:
                                            v223 := GenS2{F0: fi39}
                                            v36 -= 2
                                            _ = v223
                                        case 4294967295:
                                            if (17 == 65535) {
                                                v224 := GenS18{F0: 127}
                                                _ = v224
                                            }
                                        case 0:
                                            v36 += (^(v36))
                                        }
                                    }
                                    var v225 int8 = -3
                                    switch (fi40 << ((1) & 63)) {
                                    case 554:
                                        var v226 int64 = -9223372036854775808
                                        _ = v226
                                    default:
                                        v227 := GenS7{F0: "a"}
                                        _ = v227
                                    }
                                    if ("\n\t" > "foo bar") {
                                        var v228 uint16 = 42
                                        _ = v228
                                    } else {
                                        var v229 string = (((v2 + "") + v2) + v2)
                                        if ("\n\t" <= (v2 + ("\n\t" + (v2 + "hello")))) {
                                            if (!(65535 > 4294967295)) {
                                                if false {
                                                    fmt.Println(fi40, v2)
                                                    var v230 uint16 = 32767
                                                    v231 := GenS16{F0: uint16(v36)}
                                                    _ = v230
                                                    _ = v231
                                                } else {
                                                    fmt.Println(v225)
                                                    var v232 float64 = 0.1
                                                    _ = v232
                                                }
                                                v36 -= v36
                                                v229 = (v2 + (v229 + v2))
                                            } else {
                                                v233 := v3[(v229 + v229)]
                                                if v233 {
                                                    v234 := GenS12{F0: v225}
                                                    _ = v234
                                                } else {
                                                    v235 := GenS19{F0: v36, F1: 2147483647}
                                                    _ = v235
                                                }
                                                v225 += (v225 - (v225 ^ (v225 * (v225 >> ((1) & 63)))))
                                                _ = v233
                                            }
                                        }
                                        _ = v229
                                    }
                                    if ("" >= v2) {
                                        fmt.Println(v2)
                                    } else {
                                        var v236 int32 = -1
                                        v236 += v236
                                        var v237 float64 = (3.14159 - v217)
                                        _ = v236
                                        _ = v237
                                    }
                                    _ = v225
                                }
                            }
                            _ = v218
                        }
                        var v238 float64 = 1e10
                        if false {
                            var v239 bool = (2 != 4294967295)
                            _ = v239
                        } else {
                            v240 := GenS22{F0: 42, F1: 42, F2: 42}
                            v241 := v1[(fi39 & (fi39 + 32768))]
                            _ = v240
                            _ = v241
                        }
                        _ = v238
                    }
                    _ = v217
                } else {
                    switch (fi39 * 986) {
                    case 128:
                        if (!true) {
                            fmt.Println(fi39)
                        } else {
                            v242 := []int{(fi40 % ((127) | 1))}
                            _ = v242
                        }
                    case -2147483648:
                        v243 := GenS18{F0: -128}
                        _ = v243
                    case 2147483647:
                        switch fi40 {
                        case 42:
                            v2 = ("a" + ((v2 + v2) + "\n\t"))
                        default:
                            v33 = v33
                        }
                    }
                    var v244 int64 = -9223372036854775808
                    v2 = (v2 + "\n\t")
                    _ = v244
                }
                _ = len(v3)
                var v245 uint = 765
                _ = v245
            }
            fmt.Println(v2, v36, v33)
            _ = fi40
        }
        _ = fi39
    }
    v246 := GenS10{F0: (-2.25 == -2.25), F1: 1}
    var v247 int64 = 1
    if false {
        v2 = v2
        v33 = v33
        v248 := GenS23{F0: v247, F1: 42, F2: "hello"}
        _ = v248
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v247:", v247)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v246
    _ = v247
    _ = v3
    _ = v33
    _ = v34
    _ = v36
}
