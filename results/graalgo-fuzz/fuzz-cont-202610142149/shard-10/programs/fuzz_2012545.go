package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 int; F1 uint8; F2 bool }
type GenS12 struct { F0 uint; F1 uint64; F2 int16 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint64; F1 uint32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint32; F1 int16; F2 int }
type GenS18 struct { F0 uint64; F1 int8; F2 int64 }
type GenS19 struct { F0 uint8; F1 int }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int32; F1 int8 }
type GenS22 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS23 struct { F0 uint32; F1 uint16 }
type GenS24 struct { F0 int32; F1 uint64 }
func genFn0(p0 string, p1 int, p2 int) bool {
    var v5 int = p2
    switch p0 {
    case "":
        switch ("foo bar" + "a") {
        case "hello":
            var v6 string = (p0 + "a")
            _ = v6
        case "foo bar":
            var v7 uint = 42
            _ = v7
        case "":
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(fi8, p0, p1)
                v5 -= fi8
                switch (p2 << ((179) & 63)) {
                case -126:
                    v9 := GenS1{F0: 0, F1: 960}
                    _ = v9
                case -2147483648:
                    v10 := []int{(p2 + 684), (p1 >> ((65535) & 63)), (p1 | (v5 - p1))}
                    _ = v10
                case 127:
                    fmt.Println(fi8, p2, p2)
                }
                _ = fi8
            }
        }
    }
    if ((1 <= -771) && ((1.5 == 1.5) || (819 >= 65535))) {
        p2 = (p2 * (-(p2)))
        switch p1 {
        case 0:
            var v11 int = v5
            _ = v11
        case -1:
            if false {
                v12 := []int{(p1 / (((p2 + p2)) | 1))}
                _ = v12
            }
        case -326:
            fmt.Println(p2, p1, p0)
        }
    }
    var v13 int32 = -429
    switch p0 {
    case "hello":
        v14 := GenS2{F0: -133, F1: 702}
        _ = v14
    case "a":
        v15 := []int{(p2 * -1), -764, (p2 ^ 290)}
        _ = v15
    case "foo bar":
        if true {
            if true {
                if false {
                    var v16 uint64 = 0
                    v17 := GenS1{F0: 128, F1: 1000}
                    var v18 string = "\n\t"
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS3{F0: p1, F1: v13}
                    _ = v19
                }
            }
        } else {
            v13 = (v13 << ((1) & 63))
            v13 = 229
        }
    default:
        p0 = (p0 + (p0 + "\n\t"))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return (255 > 2)
}
func genFn1(p0 uint32, p1 uint64, p2 uint) bool {
    for fi20 := 0; fi20 < 17; fi20++ {
        p2 = p2
        p1 += (p1 * p1)
        var v21 int8 = 42
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint16 = 42
        if (!true) {
            var v24 uint64 = (p1 - (p1 / (((p1 % ((p1) | 1))) | 1)))
            _ = v24
        }
        var v25 float32 = (float32((0.1 / (0.0 - (((-2.25 * (0.1 * (0.0 - 1.5))) - (0.1 * 3.14159)) * 1e10)))) / (1e10 + (-2.25 + ((((0.1 + ((1.5 / -2.25) / 0.1)) / 0.1) - (1e10 + (-2.25 - 3.14159))) + 0.1))))
        _ = fi22
        _ = v23
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := []int{-811, fi26}
        if len(v27) > 0 {
            v27[0] = -1
        }
        _ = fi26
        _ = v27
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 0
    var v3 int16 = -970
    v4 := make(map[uint32]string)
    var v28 string = "a"
    _ = len(v4)
    v29 := GenS4{F0: 2}
    v30 := GenS1{F0: 1, F1: v3}
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint64 = 293
        if (270 == 0) {
            v33 := GenS5{F0: v3, F1: -1, F2: v3}
            if false {
                for fi34 := 0; fi34 < 5; fi34++ {
                    if true {
                        if ((fi31 / (((fi34 / ((128) | 1))) | 1)) <= fi34) {
                            fmt.Println(v32)
                            fmt.Println(fi34, fi31, v32)
                            v35 := GenS1{F0: 127, F1: -32768}
                            _ = v35
                        } else {
                            var v36 float32 = -2.25
                            _ = v36
                        }
                        _ = len(v1)
                    } else {
                        var v37 int = 1
                        if false {
                            v38 := GenS5{F0: -595, F1: 42, F2: v3}
                            var v39 float64 = float64(-2.25)
                            _ = v38
                            _ = v39
                        } else {
                            v40 := GenS5{F0: 32767, F1: 291, F2: (-(v3))}
                            if false {
                                switch v28 {
                                case "":
                                    v41 := GenS2{F0: -1, F1: 42}
                                    _ = v41
                                case "foo bar":
                                    var v42 float64 = ((1.5 - (((0.1 / 0.1) * 1e10) + -2.25)) * 0.1)
                                    if false {
                                        if (!true) {
                                            v2 = (v2 >> ((18446744073709551615) & 63))
                                        }
                                        if (v3 >= v3) {
                                            var v43 bool = false
                                            var v44 uint64 = v32
                                            _ = v43
                                            _ = v44
                                        } else {
                                            fmt.Println(fi34, v3)
                                            if true {
                                                v42 -= v42
                                                var v45 int64 = -287
                                                _ = v45
                                            } else {
                                                v28 = (v28 + (v28 + v28))
                                                v32 = v32
                                                var v46 uint16 = 775
                                                _ = v46
                                            }
                                        }
                                    }
                                    var v47 float64 = v42
                                    _ = v42
                                    _ = v47
                                case "a":
                                    v2 -= 42
                                default:
                                    v37 = (fi31 % (((^(fi34))) | 1))
                                }
                                _ = len(v1)
                            }
                            if false {
                                var v48 bool = (!true)
                                var v49 int64 = 400
                                _ = v48
                                _ = v49
                            } else {
                                switch -597 {
                                case 9223372036854775807:
                                    var v50 float32 = 1.5
                                    _ = v50
                                case 42:
                                    v37 += -815
                                case 0:
                                    if true {
                                        if (!false) {
                                            if false {
                                                _ = len(v4)
                                                switch ((v28 + v28) + v28) {
                                                case "\n\t":
                                                    var v51 bool = (!(!(!true)))
                                                    _ = v51
                                                case "":
                                                    var v52 string = (v28 + v28)
                                                    _ = v52
                                                default:
                                                    v53 := []int{}
                                                    _ = v53
                                                }
                                                v32 += (v32 | (v32 / (((v32 + (v32 << ((4294967295) & 63)))) | 1)))
                                            } else {
                                                v54 := []int{fi31, (v37 / (((fi31 - -292)) | 1))}
                                                v2 = (v2 * v2)
                                                var v55 uint8 = 255
                                                _ = v54
                                                _ = v55
                                            }
                                            if true {
                                                var v56 uint8 = 2
                                                var v57 uint16 = 65535
                                                _ = v56
                                                _ = v57
                                            }
                                            delete(v1, (fi31 % ((760) | 1)))
                                        } else {
                                            if (false && ((!false) && false)) {
                                                v58 := GenS2{F0: -856, F1: 65535}
                                                if false {
                                                    if false {
                                                        fmt.Println(v32, v3)
                                                        if true {
                                                            var v59 float64 = 0.0
                                                            _ = v59
                                                        } else {
                                                            if (v32 < (v32 * v32)) {
                                                                if true {
                                                                    switch (v28 + ("hello" + (v28 + v28))) {
                                                                    case "a":
                                                                        v60 := GenS1{F0: 128, F1: (v3 / (((v3 % ((0) | 1))) | 1))}
                                                                        _ = v60
                                                                    case "":
                                                                        v61 := GenS6{F0: 0}
                                                                        _ = v61
                                                                    default:
                                                                        fmt.Println(v28, v32)
                                                                    }
                                                                    var v62 int8 = (v2 - v2)
                                                                    v63 := GenS4{F0: 2147483648}
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    var v64 uint = 0
                                                                    _ = v64
                                                                }
                                                                if (!((42 >= 0) && ((0 <= 18446744073709551615) && false))) {
                                                                    v65 := GenS5{F0: (v3 - (v3 & (v3 ^ (v3 | 42)))), F1: 0, F2: v3}
                                                                    _ = v65
                                                                }
                                                                var v66 int64 = 52
                                                                _ = v66
                                                            } else {
                                                                switch v37 {
                                                                case -261:
                                                                    if (0.1 <= -2.25) {
                                                                        if true {
                                                                            v67 := GenS2{F0: 1, F1: 2}
                                                                            fmt.Println(v32)
                                                                            _ = v67
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v3, fi34, fi34)
                                                                                switch fi34 {
                                                                                case 32768:
                                                                                    if false {
                                                                                        v68 := []int{128, v37, (fi31 | (-(fi31))), 32767}
                                                                                        v69 := GenS2{F0: 1, F1: 259}
                                                                                        switch (fi31 * (fi34 ^ fi34)) {
                                                                                        case -2147483648:
                                                                                            fmt.Println(v2)
                                                                                        case 9223372036854775807:
                                                                                            if false {
                                                                                                v70 := GenS2{F0: -674, F1: 2}
                                                                                                var v71 int16 = -657
                                                                                                if (true && false) {
                                                                                                    var v72 uint8 = 0
                                                                                                    v73 := v1[int(v2)]
                                                                                                    _ = v72
                                                                                                    _ = v73
                                                                                                } else {
                                                                                                    fmt.Println(v71, fi34)
                                                                                                    v3 -= -1
                                                                                                }
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                        }
                                                                                        v74 := GenS4{F0: 1}
                                                                                        var v75 bool = true
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                    }
                                                                                case 4294967295:
                                                                                    switch (-(fi34)) {
                                                                                    case 127:
                                                                                        var v76 string = ("foo bar" + v28)
                                                                                        _ = v76
                                                                                    case 32767:
                                                                                        v2 = (^(v2))
                                                                                    default:
                                                                                        v77 := GenS7{F0: false}
                                                                                        _ = v77
                                                                                    }
                                                                                case -1:
                                                                                    v78 := GenS2{F0: 2147483647, F1: 0}
                                                                                    _ = v78
                                                                                }
                                                                                var v79 int32 = -2147483648
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v80 := []int{(-(fi34)), 127, (fi31 << ((42) & 63))}
                                                                        v81 := v1[fi31]
                                                                        var v82 uint32 = 133
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                case 9223372036854775807:
                                                                    if (false || ((true || (!true)) || false)) {
                                                                        var v83 int = 4294967295
                                                                        _ = v83
                                                                    } else {
                                                                        switch v37 {
                                                                        case 183:
                                                                            v84 := []int{-708, -1}
                                                                            _ = v84
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        fmt.Println(v32)
                                                                        var v85 int = (v37 << ((1) & 63))
                                                                        _ = v85
                                                                    }
                                                                default:
                                                                    v86 := GenS8{F0: 1}
                                                                    _ = v86
                                                                }
                                                                switch v28 {
                                                                case "":
                                                                    var v87 float64 = float64((1e10 - (3.14159 + 3.14159)))
                                                                    var v88 uint16 = 255
                                                                    _ = v87
                                                                    _ = v88
                                                                case "foo bar":
                                                                    v2 -= v2
                                                                case "hello":
                                                                    if ((float64((1.5 * 0.0)) * 0.1) < (1.5 * 0.1)) {
                                                                        v89 := GenS9{F0: v37, F1: 10}
                                                                        v90 := GenS10{F0: 2, F1: 212}
                                                                        fmt.Println(v2)
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                default:
                                                                    switch 0 {
                                                                    case 782:
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if false {
                                                                                var v91 int8 = -68
                                                                                if true {
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        switch v28 {
                                                                                        case "a":
                                                                                            v92 := GenS2{F0: 2147483647, F1: 42}
                                                                                            _ = v92
                                                                                        case "":
                                                                                            if (!(float64(float32((((float64(-2.25) - -2.25) * (0.0 + 3.14159)) / 1e10))) < 1.5)) {
                                                                                                var v93 float32 = 0.0
                                                                                                var v94 uint16 = 32767
                                                                                                v4[0] = "foo bar"
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v95 uint64 = 4294967295
                                                                                            _ = v95
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            v96 := GenS11{F0: (fi31 % (((fi31 << ((943) & 63))) | 1)), F1: 2, F2: (!(913 >= 485))}
                                                                                            v97 := v1[(fi31 << ((308) & 63))]
                                                                                            var v98 uint64 = v32
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        } else {
                                                                                            v28 = "foo bar"
                                                                                        }
                                                                                    case "a":
                                                                                        if false {
                                                                                            var v99 int8 = v2
                                                                                            _ = v99
                                                                                        } else {
                                                                                            var v100 bool = false
                                                                                            switch (fi34 ^ (-(fi34))) {
                                                                                            case 127:
                                                                                                var v101 uint64 = 46
                                                                                                _ = v101
                                                                                            case 1:
                                                                                                v102 := GenS12{F0: 4294967295, F1: (v32 & (v32 % ((v32) | 1))), F2: v3}
                                                                                                _ = v102
                                                                                            case 762:
                                                                                                if (1.5 > -2.25) {
                                                                                                    var v103 int64 = -1
                                                                                                    _ = v103
                                                                                                } else {
                                                                                                    v4[4294967295] = v28
                                                                                                }
                                                                                            }
                                                                                            v91 += (v2 - v2)
                                                                                            _ = v100
                                                                                        }
                                                                                        v104 := GenS13{F0: 2, F1: v28}
                                                                                        _ = v104
                                                                                    default:
                                                                                        v105 := []int{(-(fi31)), fi34, -1, v37}
                                                                                        _ = v105
                                                                                    }
                                                                                }
                                                                                v106 := GenS9{F0: 0, F1: 2}
                                                                                _ = v106
                                                                                _ = v91
                                                                            }
                                                                        case "hello":
                                                                            if (((v28 + v28) == v28) || true) {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        v107 := GenS9{F0: v37, F1: 42}
                                                                                        v2 = (^(v2))
                                                                                        var v108 int64 = 42
                                                                                        v2 += (v2 - (v2 + (v2 % (((v2 ^ 8)) | 1))))
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    } else {
                                                                                        fmt.Println(v37, v28, v37)
                                                                                        var v109 float64 = (0.1 + 0.0)
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                                var v110 uint = 1
                                                                                if false {
                                                                                    if true {
                                                                                        v111 := GenS14{F0: (v32 << ((uint(v32)) & 63))}
                                                                                        _ = v111
                                                                                    }
                                                                                    switch (v37 >> ((v110) & 63)) {
                                                                                    case 32767:
                                                                                        var v112 int8 = (v2 << (((v110 / ((uint(v32)) | 1))) & 63))
                                                                                        _ = v112
                                                                                    case 42:
                                                                                        var v113 int8 = v2
                                                                                        _ = v113
                                                                                    case -9223372036854775808:
                                                                                        if (v110 > (v110 % (((v110 | v110)) | 1))) {
                                                                                            var v114 string = "foo bar"
                                                                                            if false {
                                                                                                if true {
                                                                                                    switch v28 {
                                                                                                    case "hello":
                                                                                                        if false {
                                                                                                            fmt.Println(fi31, v28)
                                                                                                            if (!(((v28 + ("a" + ((v28 + v114) + v28))) + v28) != "hello")) {
                                                                                                                v115 := GenS15{F0: 2, F1: 0}
                                                                                                                v37 += fi31
                                                                                                                fmt.Println(v114, v114, v2)
                                                                                                                var v116 float64 = 3.14159
                                                                                                                _ = v115
                                                                                                                _ = v116
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        _ = len(v1)
                                                                                                    }
                                                                                                    switch -511 {
                                                                                                    case 128:
                                                                                                        if false {
                                                                                                            var v117 uint8 = 0
                                                                                                            _ = v117
                                                                                                        } else {
                                                                                                            var v118 uint = (v110 << (((^(v110))) & 63))
                                                                                                            delete(v1, (fi31 - fi34))
                                                                                                            _ = v118
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        if false {
                                                                                                            var v119 float64 = (0.1 / (1e10 + 1e10))
                                                                                                            v37 -= (v37 ^ int(v110))
                                                                                                            if true {
                                                                                                                switch v37 {
                                                                                                                case 9223372036854775807:
                                                                                                                    var v120 uint32 = uint32(v32)
                                                                                                                    _ = v120
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        var v121 int64 = 959
                                                                                                                        fmt.Println(v28)
                                                                                                                        v122 := []int{(^(fi34))}
                                                                                                                        _ = v121
                                                                                                                        _ = v122
                                                                                                                    } else {
                                                                                                                        v3 = 333
                                                                                                                    }
                                                                                                                }
                                                                                                                v123 := []int{}
                                                                                                                fmt.Println(v119, v37, v119)
                                                                                                                _ = v123
                                                                                                            }
                                                                                                            v124 := GenS16{F0: (-(v2))}
                                                                                                            switch v114 {
                                                                                                            case "hello":
                                                                                                                if (!true) {
                                                                                                                    v125 := GenS7{F0: false}
                                                                                                                    v119 = 3.14159
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v37 -= (-(fi31))
                                                                                                            }
                                                                                                            _ = v119
                                                                                                            _ = v124
                                                                                                        } else {
                                                                                                            v126 := GenS8{F0: 209}
                                                                                                            fmt.Println(v2, v37)
                                                                                                            v37 += v37
                                                                                                            _ = v126
                                                                                                        }
                                                                                                    }
                                                                                                    v127 := []int{-9223372036854775808}
                                                                                                    v2 = int8(v32)
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    fmt.Println(v2, v114)
                                                                                                    var v128 int32 = 1
                                                                                                    v129 := []int{}
                                                                                                    var v130 int16 = v3
                                                                                                    if len(v129) >= 2 {
                                                                                                        v129 = v129[1:len(v129)]
                                                                                                    }
                                                                                                    if ((!true) && (v2 < v2)) {
                                                                                                        v110 = v110
                                                                                                    }
                                                                                                    _ = v128
                                                                                                    _ = v129
                                                                                                    _ = v130
                                                                                                }
                                                                                                if true {
                                                                                                    v32 = v32
                                                                                                    fmt.Println(v32)
                                                                                                } else {
                                                                                                    v131 := GenS17{F0: 2147483648, F1: int16(v2), F2: v37}
                                                                                                    switch int(v3) {
                                                                                                    case 128:
                                                                                                        var v132 uint32 = 2147483648
                                                                                                        _ = v132
                                                                                                    }
                                                                                                    _ = v131
                                                                                                }
                                                                                                fmt.Println(v110, v37, v2)
                                                                                            } else {
                                                                                                v133 := GenS18{F0: 42, F1: v2, F2: 1}
                                                                                                _ = v133
                                                                                            }
                                                                                            var v134 int64 = 706
                                                                                            _ = v114
                                                                                            _ = v134
                                                                                        }
                                                                                        if (1 == 2147483648) {
                                                                                            var v135 int16 = (v3 >> ((v110) & 63))
                                                                                            var v136 int8 = 127
                                                                                            var v137 int16 = (v135 / ((int16(v32)) | 1))
                                                                                            v2 = 1
                                                                                            var v138 bool = ("foo bar" <= ((v28 + ("\n\t" + ((((((v28 + (v28 + v28)) + "foo bar") + ("\n\t" + v28)) + "") + "a") + v28))) + v28))
                                                                                            v138 = v138
                                                                                            v139 := v1[fi34]
                                                                                            _ = v135
                                                                                            _ = v136
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                            _ = v139
                                                                                        } else {
                                                                                            var v140 int64 = 9223372036854775807
                                                                                            v141 := GenS7{F0: false}
                                                                                            if false {
                                                                                                v28 = (v28 + (v28 + (("a" + v28) + ((v28 + v28) + v28))))
                                                                                            }
                                                                                            _ = v140
                                                                                            _ = v141
                                                                                        }
                                                                                        var v142 uint8 = 122
                                                                                        _ = len(v1)
                                                                                        _ = v142
                                                                                    }
                                                                                    switch v28 {
                                                                                    case "a":
                                                                                        switch ("a" + v28) {
                                                                                        case "":
                                                                                            v143 := GenS19{F0: 174, F1: fi34}
                                                                                            _ = v143
                                                                                        default:
                                                                                            v144 := v4[656]
                                                                                            _ = v144
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v145 int16 = v3
                                                                                        v146 := GenS12{F0: v110, F1: v32, F2: 594}
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                    case "":
                                                                                        v37 -= 32768
                                                                                    }
                                                                                    v3 = (v3 ^ (^(v3)))
                                                                                    if true {
                                                                                        fmt.Println(v28)
                                                                                        var v147 uint64 = 4294967295
                                                                                        v3 -= v3
                                                                                        _ = v147
                                                                                    } else {
                                                                                        v37 += fi31
                                                                                        v1[(-(fi31))] = (v28 + v28)
                                                                                        v32 = (v32 - v32)
                                                                                        v148 := GenS20{F0: -9223372036854775808}
                                                                                        _ = v148
                                                                                    }
                                                                                    if false {
                                                                                        var v149 float32 = ((0.1 / -2.25) - 0.0)
                                                                                        _ = v149
                                                                                    } else {
                                                                                        if false {
                                                                                            if ((v37 * fi34) == (fi34 >> (((v110 + v110)) & 63))) {
                                                                                                fmt.Println(fi34, v2, fi34)
                                                                                            } else {
                                                                                                var v150 float32 = 1e10
                                                                                                var v151 int64 = 85
                                                                                                _ = v150
                                                                                                _ = v151
                                                                                            }
                                                                                            var v152 uint8 = 168
                                                                                            _ = v152
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v110
                                                                            } else {
                                                                                v153 := GenS12{F0: 2, F1: v32, F2: (v3 ^ 1)}
                                                                                _ = v153
                                                                            }
                                                                        case "a":
                                                                            if true {
                                                                                v3 = (v3 >> ((2) & 63))
                                                                                v154 := GenS17{F0: 2, F1: (v3 + v3), F2: 4294967295}
                                                                                v32 = (^(v32))
                                                                                _ = v154
                                                                            } else {
                                                                                if true {
                                                                                    var v155 int = (v37 & fi34)
                                                                                    _ = v155
                                                                                }
                                                                                v156 := []int{(fi34 & int(v2)), fi34, 32767, fi34}
                                                                                fmt.Println(fi34, v32, fi31)
                                                                                _ = v156
                                                                            }
                                                                        }
                                                                    case -282:
                                                                        fmt.Println(fi31)
                                                                    }
                                                                }
                                                                if false {
                                                                    v157 := GenS18{F0: 181, F1: -128, F2: -1}
                                                                    fmt.Println(v3, fi31)
                                                                    _ = v157
                                                                }
                                                                var v158 uint64 = v32
                                                                v28 = v28
                                                                _ = v158
                                                            }
                                                            var v159 uint = 2
                                                            _ = v159
                                                        }
                                                    } else {
                                                        fmt.Println(v28, fi34, v37)
                                                        switch (v28 + v28) {
                                                        case "foo bar":
                                                            var v160 float32 = ((((float32(0.0) / float32(1e10)) - 1.5) / 1.0) * 1e10)
                                                            _ = v160
                                                        case "hello":
                                                            v161 := GenS12{F0: 737, F1: (^(v32)), F2: v3}
                                                            _ = v161
                                                        case "":
                                                            var v162 string = ("\n\t" + v28)
                                                            _ = v162
                                                        default:
                                                            v3 = v3
                                                        }
                                                    }
                                                }
                                                v163 := []int{(fi31 ^ v37)}
                                                _ = v163
                                                _ = v58
                                            } else {
                                                v2 = 42
                                                v3 -= v3
                                                var v164 int8 = -55
                                                fmt.Println(v32)
                                                _ = v164
                                            }
                                            var v165 uint32 = 1
                                            switch -1 {
                                            case -1:
                                                var v166 float32 = 1.5
                                                _ = v166
                                            case 374:
                                                var v167 uint64 = (v32 + 42)
                                                v168 := GenS21{F0: 0, F1: (v2 + v2)}
                                                _ = v167
                                                _ = v168
                                            case 32767:
                                                fmt.Println(v32)
                                            }
                                            _ = v165
                                        }
                                        fmt.Println(v32, fi31)
                                    }
                                default:
                                    v169 := v4[2147483648]
                                    _ = v169
                                }
                                v170 := []int{(fi31 >> ((4294967295) & 63)), fi31, v37, (fi34 | -676)}
                                switch -739 {
                                case 32767:
                                    v32 += v32
                                case -9223372036854775808:
                                    v171 := GenS22{F0: 191, F1: 806, F2: 42}
                                    _ = v171
                                }
                                var v172 int32 = 1
                                v2 -= 42
                                _ = v170
                                _ = v172
                            }
                            _ = v40
                        }
                        v173 := v4[0]
                        _ = v173
                        _ = v37
                    }
                    _ = fi34
                }
                v174 := GenS23{F0: 950, F1: 353}
                var v175 string = "a"
                _ = v174
                _ = v175
            } else {
                v176 := []int{fi31}
                for fi177 := 0; fi177 < 7; fi177++ {
                    var v178 uint32 = 203
                    _ = fi177
                    _ = v178
                }
                _ = v176
            }
            _ = v33
        } else {
            v32 = v32
        }
        var v179 int32 = -2147483648
        _ = fi31
        _ = v179
        _ = v32
    }
    for fi180 := 0; fi180 < 15; fi180++ {
        v3 += 861
        for fi181 := 0; fi181 < 2; fi181++ {
            if ((211 == 127) && (false || true)) {
                var v182 uint8 = 25
                _ = v182
            } else {
                v183 := GenS20{F0: -858}
                v184 := GenS24{F0: 1, F1: 901}
                _ = v183
                _ = v184
            }
            _ = fi181
        }
        _ = fi180
    }
    for fi185 := 0; fi185 < 15; fi185++ {
        if ((float32(((((3.14159 - -2.25) + (1.5 - 1.5)) / (((0.1 / 1.0) / (1.5 + ((((0.1 - 0.0) - float64(-2.25)) / 3.14159) / (3.14159 * (1e10 * float64(-2.25)))))) - (-2.25 / (float64(0.1) + 1.5)))) - -2.25)) <= -2.25) || (1 != v2)) {
            delete(v1, fi185)
            for fi186 := 0; fi186 < 17; fi186++ {
                v2 = v2
                var v187 int32 = 578
                _ = fi186
                _ = v187
            }
            var v188 int8 = v2
            _ = v188
        }
        _ = fi185
    }
    fmt.Println(v3)
    v4[1] = "\n\t"
    var v189 float32 = ((-2.25 + 1e10) + ((0.0 + 3.14159) * ((0.1 - 0.1) * (0.1 * (1.5 / 1.0)))))
    for fi190 := 0; fi190 < 14; fi190++ {
        switch fi190 {
        case -24:
            fmt.Println(v3)
        case 0:
            v191 := []int{127, 32768, fi190}
            _ = v191
        case -237:
            v3 = (-(v3))
        default:
            var v192 string = ("a" + (v28 + (v28 + v28)))
            _ = v192
        }
        _ = fi190
    }
    fmt.Println(v28)
    var v193 uint64 = 659
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v189:", v189)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v189
    _ = v193
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
