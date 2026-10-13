package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int32 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 bool; F1 uint16; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 int16; F1 uint16; F2 int8 }
type GenS13 struct { F0 bool }
func genFn0() string {
    var v7 int8 = 127
    var v8 int8 = (v7 / (((-(v7))) | 1))
    var v9 float64 = 1e10
    var v10 uint64 = 770
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int8 = (v7 | 100)
                _ = fi12
                _ = v13
            }
        }
        if (true && (!(42 <= -32768))) {
            if false {
                v9 -= v9
            }
        } else {
            var v14 int = fi11
            v15 := GenS1{F0: 127, F1: 1, F2: 426}
            if false {
                if false {
                    v16 := GenS2{F0: 2}
                    _ = v16
                }
                var v17 int32 = 1
                if false {
                    v18 := []int{v14, -1, v14, (fi11 * (-(v14)))}
                    _ = v18
                }
                _ = v17
            } else {
                var v19 uint8 = 255
                v20 := []int{(v14 * (v14 >> ((18446744073709551615) & 63))), -41, fi11, v14}
                var v21 int = (-(v14))
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v14
            _ = v15
        }
        var v22 int8 = v7
        _ = fi11
        _ = v22
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (("\n\t" + "hello") + "foo bar")
}
func genFn1(p0 uint8, p1 int64) int {
    v23 := GenS1{F0: p0, F1: 32767, F2: 723}
    if (!(false || false)) {
        switch 0 {
        case 42:
            switch ("a" + ((("foo bar" + ("" + "")) + "foo bar") + "\n\t")) {
            case "a":
                v24 := []int{121, -2147483648, 127}
                _ = v24
            default:
                v25 := []int{-374, 127}
                _ = v25
            }
        case 0:
            v26 := []int{127, 128, 32768}
            _ = v26
        case -2147483648:
            var v27 uint64 = 0
            _ = v27
        default:
            p0 -= p0
        }
        v28 := GenS3{F0: 65535, F1: (^(p0))}
        _ = v28
    } else {
        if true {
            if (!(1 == 255)) {
                var v29 string = ("a" + "a")
                if ((float32(0.0) + (float32(0.0) * 1e10)) >= (1e10 + 1e10)) {
                    fmt.Println(v29, p0)
                } else {
                    var v30 string = "\n\t"
                    v30 = "foo bar"
                    switch (v30 + "") {
                    case "foo bar":
                        v31 := GenS4{F0: 18446744073709551615, F1: 2}
                        _ = v31
                    case "":
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v30 = (("" + v30) + "foo bar")
                            var v33 int16 = 1
                            for fi34 := 0; fi34 < 13; fi34++ {
                                fmt.Println(fi34)
                                switch fi34 {
                                case -2147483648:
                                    switch "hello" {
                                    case "":
                                        if true {
                                            v35 := []int{fi32, fi34}
                                            var v36 int64 = p1
                                            p1 -= (v36 / (((^(p1))) | 1))
                                            _ = v35
                                            _ = v36
                                        }
                                    case "\n\t":
                                        v29 = (("\n\t" + "\n\t") + "foo bar")
                                        var v37 int8 = -121
                                        if false {
                                            if (p0 > p0) {
                                                fmt.Println(fi32, fi32)
                                            }
                                        }
                                        _ = v37
                                    case "a":
                                        var v38 int = (-(fi34))
                                        _ = v38
                                    }
                                    if true {
                                        v30 = v30
                                    }
                                case -421:
                                    switch ("\n\t" + v29) {
                                    case "":
                                        v39 := GenS4{F0: 4294967295, F1: p0}
                                        _ = v39
                                    case "foo bar":
                                        p1 += 9223372036854775807
                                    case "hello":
                                        var v40 uint16 = 288
                                        _ = v40
                                    default:
                                        v41 := []int{127, (fi34 << ((255) & 63))}
                                        _ = v41
                                    }
                                case 128:
                                    var v42 int8 = -1
                                    _ = v42
                                default:
                                    var v43 int = (fi32 << ((1) & 63))
                                    _ = v43
                                }
                                var v44 float64 = float64(float32((0.0 + ((1.5 - 1.5) - float64(0.1)))))
                                _ = fi34
                                _ = v44
                            }
                            _ = fi32
                            _ = v33
                        }
                    case "hello":
                        v45 := []int{-1, 75}
                        _ = v45
                    }
                    _ = v30
                }
                _ = v29
            } else {
                p1 += p1
                var v46 int64 = (-(p1))
                _ = v46
            }
            fmt.Println(p1)
            for fi47 := 0; fi47 < 14; fi47++ {
                var v48 uint16 = 92
                var v49 uint32 = 4294967295
                var v50 string = "foo bar"
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
        }
        if false {
            p1 -= -393
        }
        if ((0 != 42) || (p1 >= 376)) {
            v51 := GenS4{F0: 2, F1: (p0 | p0)}
            v52 := GenS5{F0: 1, F1: 1000, F2: -340}
            _ = v51
            _ = v52
        }
    }
    _ = p0
    _ = p1
    _ = v23
    return -9223372036854775808
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 42
    var v3 int8 = -32
    var v4 int64 = 1
    var v5 uint8 = 1
    var v6 int8 = int8(v4)
    v2 = (v4 << ((141) & 63))
    for fi53 := 0; fi53 < 15; fi53++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi53
    }
    fmt.Println(v2)
    fmt.Println(v4)
    for fi54 := 0; fi54 < 7; fi54++ {
        if false {
            v55 := GenS3{F0: 448, F1: 127}
            v5 += (v5 % ((17) | 1))
            _ = v55
        } else {
            fmt.Println(v5)
            v2 = -798
            switch (("\n\t" + "") + ("a" + "\n\t")) {
            case "a":
                v3 = 101
            case "foo bar":
                var v56 int8 = v6
                _ = v56
            default:
                var v57 int = (fi54 ^ 42)
                _ = v57
            }
        }
        _ = len(v1)
        var v58 uint8 = (^(v5))
        _ = fi54
        _ = v58
    }
    v59 := GenS6{F0: 0, F1: "\n\t"}
    var v60 uint = 242
    v61 := GenS7{F0: (((427 < 4294967295) || (4294967295 == 4294967295)) || (false || false)), F1: 2, F2: -9223372036854775808}
    v62 := GenS7{F0: false, F1: 0, F2: 1}
    v63 := GenS4{F0: 2, F1: (v5 / ((v5) | 1))}
    for fi64 := 0; fi64 < 7; fi64++ {
        var v65 int8 = v6
        for fi66 := 0; fi66 < 9; fi66++ {
            switch int(v5) {
            case -394:
                switch 4294967295 {
                case 42:
                    if false {
                        v2 = 632
                    }
                }
            }
            if false {
                v67 := GenS8{F0: genFn1((v5 * 0), (v2 & v4))}
                if (!((v60 < (v60 ^ 255)) || ((!true) || true))) {
                    if false {
                        if (false || (!(42 < 2))) {
                            var v68 bool = true
                            var v69 int32 = -2147483648
                            var v70 int8 = int8(v60)
                            _ = v68
                            _ = v69
                            _ = v70
                        } else {
                            v71 := GenS4{F0: 4294967295, F1: v5}
                            var v72 int8 = v65
                            _ = v71
                            _ = v72
                        }
                        v2 += 1
                        v73 := GenS8{F0: genFn1((v5 | v5), (v4 * (v4 + 439)))}
                        _ = v73
                    }
                } else {
                    var v74 float64 = 0.0
                    v1 = append(v1, 1)
                    _ = v74
                }
                _ = v67
            }
            v60 = v60
            _ = fi66
        }
        _ = fi64
        _ = v65
    }
    if (v3 >= (v6 >> ((v60) & 63))) {
        if (-712 >= -822) {
            fmt.Println(v60, v3)
            v3 = (v6 - v3)
            v3 += (v6 | 127)
        }
    } else {
        var v75 float32 = 0.0
        v76 := GenS1{F0: v5, F1: 32767, F2: 361}
        v3 = (v6 ^ v3)
        _ = v75
        _ = v76
    }
    v77 := GenS3{F0: (v60 * (v60 + 4294967295)), F1: v5}
    switch "" {
    case "":
        var v78 uint16 = 2
        _ = v78
    case "foo bar":
        v79 := GenS9{F0: 270, F1: 4294967295, F2: 0}
        _ = v79
    case "hello":
        var v80 uint32 = 2
        _ = v80
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi81 := 0; fi81 < 2; fi81++ {
        for fi82 := 0; fi82 < 1; fi82++ {
            var v83 uint32 = 42
            _ = fi82
            _ = v83
        }
        v5 = (v5 | v5)
        for fi84 := 0; fi84 < 17; fi84++ {
            v85 := GenS4{F0: 0, F1: (^(v5))}
            var v86 string = ""
            if (true || false) {
                switch (fi84 % ((273) | 1)) {
                case 42:
                    v3 = (v3 | 42)
                case -1:
                    fmt.Println(fi81, fi81)
                case 2147483647:
                    if false {
                        v87 := GenS8{F0: (fi84 >> ((v60) & 63))}
                        _ = v87
                    }
                default:
                    v88 := GenS8{F0: 4294967295}
                    _ = v88
                }
                v89 := GenS6{F0: 2147483647, F1: ((((v86 + (("\n\t" + v86) + (v86 + (v86 + v86)))) + (v86 + v86)) + v86) + v86)}
                fmt.Println(v60)
                v4 = (v2 + (v2 & (v2 >> (((v60 | (v60 % ((255) | 1)))) & 63))))
                _ = v89
            }
            if (!(!false)) {
                var v90 uint64 = 1
                var v91 int = fi81
                _ = v90
                _ = v91
            } else {
                if ((false || (true || (false && true))) && (120 != 832)) {
                    switch (((v86 + v86) + v86) + ("hello" + v86)) {
                    case "a":
                        fmt.Println(fi84, fi81, fi84)
                    default:
                        v5 += (v5 ^ v5)
                    }
                    v92 := GenS1{F0: (v5 % ((1) | 1)), F1: 65535, F2: -844}
                    _ = v92
                } else {
                    fmt.Println(v2, v6, fi84)
                    v2 += v4
                    if false {
                        v93 := GenS10{F0: v86}
                        switch v86 {
                        case "hello":
                            v86 = "a"
                        case "":
                            v86 = v86
                        default:
                            if true {
                                var v94 int = -9223372036854775808
                                _ = v94
                            } else {
                                v95 := GenS11{F0: v3, F1: 65535}
                                _ = v95
                            }
                        }
                        if ((!true) && false) {
                            if true {
                                v60 = v60
                            } else {
                                v96 := GenS7{F0: ((32767 >= 919) && (!false)), F1: 42, F2: (^(fi81))}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v96
                            }
                            v97 := GenS11{F0: v3, F1: 368}
                            if true {
                                v98 := GenS6{F0: 1, F1: (v86 + v86)}
                                v2 = v2
                                _ = v98
                            } else {
                                v1 = append(v1, 2147483648)
                                if true {
                                    if true {
                                        fmt.Println(v6)
                                        if true {
                                            v4 = v4
                                        }
                                        v99 := GenS10{F0: v86}
                                        _ = v99
                                    }
                                }
                                v86 = v86
                                var v100 uint32 = 170
                                fmt.Println(v2, v3, v60)
                                if ((^(fi84)) <= fi81) {
                                    switch v86 {
                                    case "":
                                        v86 = (((("a" + "a") + "\n\t") + "foo bar") + "\n\t")
                                    }
                                    var v101 string = "foo bar"
                                    _ = v101
                                } else {
                                    if (v86 < (v86 + "\n\t")) {
                                        v102 := GenS10{F0: "\n\t"}
                                        v2 += (v2 >> ((v60) & 63))
                                        if ("foo bar" <= (("foo bar" + v86) + v86)) {
                                            var v103 int64 = v2
                                            var v104 uint64 = uint64(v100)
                                            var v105 uint8 = v5
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                        } else {
                                            v100 = 2147483648
                                            if len(v1) > 0 {
                                                v1[0] = (^(v100))
                                            }
                                            if false {
                                                var v106 int8 = (v6 >> (((v60 * 255)) & 63))
                                                _ = v106
                                            } else {
                                                fmt.Println(fi81, fi81)
                                                v107 := GenS2{F0: 2}
                                                _ = v107
                                            }
                                        }
                                        _ = v102
                                    }
                                    v60 = v60
                                    if false {
                                        v100 = v100
                                        if false {
                                            if false {
                                                if true {
                                                    switch fi84 {
                                                    case -1:
                                                        if true {
                                                            if false {
                                                                v5 = (v5 & (v5 % ((v5) | 1)))
                                                            }
                                                            var v108 uint64 = 2
                                                            var v109 uint8 = v5
                                                            _ = v108
                                                            _ = v109
                                                        } else {
                                                            var v110 int32 = 1
                                                            _ = v110
                                                        }
                                                    default:
                                                        v2 = v4
                                                    }
                                                    v111 := GenS9{F0: 1, F1: 456, F2: -2147483648}
                                                    var v112 int = genFn1(v5, 42)
                                                    if ("hello" <= (v86 + v86)) {
                                                        v113 := GenS7{F0: false, F1: 1, F2: (^(v112))}
                                                        v3 -= (v3 | v6)
                                                        _ = v113
                                                    } else {
                                                        v1 = append(v1, 786)
                                                        v112 += (fi81 & 128)
                                                        fmt.Println(v4, v3, v2)
                                                    }
                                                    v86 = "hello"
                                                    if false {
                                                        v114 := GenS12{F0: 1, F1: 255, F2: (v3 | (v6 ^ 15))}
                                                        var v115 uint32 = v100
                                                        var v116 int = (fi81 & int(v6))
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                    if (!false) {
                                                        v112 += int(v100)
                                                        var v117 uint32 = v100
                                                        if (v2 < (v2 % ((v2) | 1))) {
                                                            v4 -= (v4 ^ -9223372036854775808)
                                                            var v118 float32 = ((1.5 - (((3.14159 * 0.0) - -2.25) / (1e10 + 1.5))) / 1.0)
                                                            _ = v118
                                                        }
                                                        _ = v117
                                                    }
                                                    _ = v111
                                                    _ = v112
                                                } else {
                                                    switch fi84 {
                                                    case 4294967295:
                                                        v60 = 65535
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case 12:
                                                        if len(v1) > 0 {
                                                            v1[0] = v100
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        var v119 uint32 = v100
                                                        _ = v119
                                                    default:
                                                        var v120 uint32 = (v100 - (^(v100)))
                                                        _ = v120
                                                    }
                                                }
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (255 != 32767) {
                                                    v100 += 1
                                                } else {
                                                    v121 := GenS2{F0: 2}
                                                    var v122 uint8 = (v5 | uint8(v100))
                                                    if ((42 == -32768) || true) {
                                                        var v123 uint64 = 42
                                                        switch (v86 + (("a" + "\n\t") + v86)) {
                                                        case "foo bar":
                                                            v124 := GenS10{F0: "foo bar"}
                                                            _ = v124
                                                        default:
                                                            switch 735 {
                                                            case 32768:
                                                                var v125 uint64 = v123
                                                                _ = v125
                                                            case 1:
                                                                _ = len(v1)
                                                            }
                                                        }
                                                        _ = v123
                                                    }
                                                    _ = v121
                                                    _ = v122
                                                }
                                                switch v86 {
                                                case "foo bar":
                                                    fmt.Println(v86)
                                                case "hello":
                                                    v5 = (^(v5))
                                                default:
                                                    if true {
                                                        if false {
                                                            if (42 <= (v5 & (v5 + (v5 + 42)))) {
                                                                v126 := GenS13{F0: (false && false)}
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v126
                                                            } else {
                                                                var v127 uint = v60
                                                                var v128 int16 = -1
                                                                v100 = v100
                                                                var v129 int16 = 1
                                                                _ = v127
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        }
                                                        v4 -= v2
                                                        if (0 < (v60 * 42)) {
                                                            switch fi84 {
                                                            case 755:
                                                                if (255 < 2) {
                                                                    var v130 int16 = -235
                                                                    if false {
                                                                        if true {
                                                                            v6 += 127
                                                                            v3 = -1
                                                                        }
                                                                        var v131 int64 = (v4 | v4)
                                                                        v132 := GenS10{F0: "a"}
                                                                        _ = v131
                                                                        _ = v132
                                                                    }
                                                                    if (81 <= 677) {
                                                                        fmt.Println(v60, v60)
                                                                    } else {
                                                                        var v133 int8 = (v6 ^ (v6 | (v3 / (((v3 / (((v6 % (((v3 % ((int8(fi81)) | 1))) | 1))) | 1))) | 1))))
                                                                        _ = v133
                                                                    }
                                                                    _ = v130
                                                                }
                                                            case 32767:
                                                                if (42 <= uint(v5)) {
                                                                    fmt.Println(v5, v3, v100)
                                                                }
                                                            }
                                                            var v134 uint64 = uint64(v100)
                                                            v100 = v100
                                                            _ = v134
                                                        } else {
                                                            v135 := GenS1{F0: 128, F1: 255, F2: -1}
                                                            var v136 int32 = 2147483647
                                                            var v137 uint8 = v5
                                                            _ = v135
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            v1 = append(v1, (v100 << (((v60 << ((uint(v100)) & 63))) & 63)))
                                            fmt.Println(fi81)
                                        }
                                        switch "\n\t" {
                                        case "":
                                            var v138 uint32 = v100
                                            _ = v138
                                        case "foo bar":
                                            var v139 int16 = -1
                                            _ = v139
                                        default:
                                            switch "hello" {
                                            case "":
                                                var v140 int = 42
                                                _ = v140
                                            default:
                                                var v141 int16 = 1
                                                _ = v141
                                            }
                                        }
                                    }
                                }
                                _ = v100
                            }
                            _ = v97
                        }
                        _ = v93
                    } else {
                        fmt.Println(fi84)
                        switch (v86 + v86) {
                        case "foo bar":
                            fmt.Println(v60, v4, fi81)
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        }
                        var v142 string = (v86 + "\n\t")
                        _ = v142
                    }
                    v2 += -892
                }
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    fmt.Println(v60, v2, v4)
                }
                var v143 uint32 = 0
                _ = v143
            }
            var v144 float32 = 3.14159
            _ = fi84
            _ = v144
            _ = v85
            _ = v86
        }
        _ = fi81
    }
    var v145 bool = ((^(v4)) != 926)
    var v146 string = ("\n\t" + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v145
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v77
}
