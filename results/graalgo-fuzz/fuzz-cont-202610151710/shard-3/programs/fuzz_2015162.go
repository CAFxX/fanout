package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 int; F2 uint16 }
type GenS12 struct { F0 uint16; F1 int8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16; F1 uint16; F2 int }
type GenS15 struct { F0 int32; F1 bool }
type GenS16 struct { F0 uint; F1 uint }
type GenS17 struct { F0 bool; F1 uint8 }
type GenS18 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 uint8, p2 uint16) uint8 {
    v5 := GenS1{F0: -219, F1: false, F2: 0}
    switch "a" {
    case "hello":
        var v6 uint32 = 1
        _ = v6
    case "":
        for fi7 := 0; fi7 < 10; fi7++ {
            p2 += p2
            switch "" {
            case "\n\t":
                var v8 uint16 = (p2 + p2)
                _ = v8
            }
            p2 -= p2
            _ = fi7
        }
    case "\n\t":
        var v9 float32 = -2.25
        _ = v9
    default:
        v10 := []int{575, 2147483647, 32768, -1}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1(p0 int8) int16 {
    var v11 string = "\n\t"
    switch v11 {
    case "a":
        var v12 uint8 = 210
        _ = v12
    case "foo bar":
        if (-84 == 0) {
            fmt.Println(p0, p0)
        } else {
            if true {
                v11 = v11
                p0 = 42
                v13 := []int{}
                _ = v13
            } else {
                v14 := GenS2{F0: 255}
                if (3.14159 != 1e10) {
                    var v15 int = 299
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v15 -= v15
                        v17 := GenS3{F0: 2, F1: p0, F2: -1}
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v19 := GenS4{F0: 0, F1: 2147483647}
                            fmt.Println(p0, fi18, fi18)
                            if true {
                                var v20 int64 = 578
                                _ = v20
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                    if (!(!true)) {
                        v21 := []int{}
                        var v22 int8 = p0
                        v23 := GenS1{F0: v15, F1: false, F2: v22}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v15
                } else {
                    switch v11 {
                    case "":
                        if ("foo bar" < v11) {
                            switch v11 {
                            case "\n\t":
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    v11 = v11
                                    var v25 uint8 = 0
                                    _ = fi24
                                    _ = v25
                                }
                            case "":
                                v11 = v11
                            default:
                                var v26 int32 = -2147483648
                                _ = v26
                            }
                            var v27 int64 = 0
                            _ = v27
                        } else {
                            p0 += p0
                        }
                    default:
                        v28 := GenS4{F0: -1, F1: -2147483648}
                        _ = v28
                    }
                    v11 = v11
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v30 := []int{fi29, 137, (^(fi29)), fi29}
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    switch (("a" + "a") + v11) {
                    case "hello":
                        v32 := GenS2{F0: 255}
                        _ = v32
                    default:
                        p0 = p0
                    }
                    var v33 int16 = int16(fi31)
                    _ = fi31
                    _ = v33
                }
                _ = v14
            }
            v34 := []int{-2147483648, 190}
            _ = v34
        }
    case "":
        p0 -= -68
    }
    var v35 uint32 = 2
    _ = p0
    _ = v11
    _ = v35
    return 294
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint16 = 32767
    var v4 uint = 289
    for fi36 := 0; fi36 < 10; fi36++ {
        v4 += 4294967295
        _ = fi36
    }
    v1[65535] = 786
    v37 := []int{4294967295, 128}
    delete(v1, 2)
    _ = len(v1)
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 int32 = 0
        v3 -= v3
        _ = len(v37)
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: 0, F1: 42}
    var v41 uint = 18446744073709551615
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 bool = true
        v44 := GenS3{F0: v4, F1: -1, F2: genFn1(-73)}
        if v43 {
            if v43 {
                if (4294967295 >= 0) {
                    var v45 uint16 = v3
                    _ = v45
                } else {
                    switch (("" + (("" + ("hello" + ("foo bar" + "\n\t"))) + ("\n\t" + "foo bar"))) + "a") {
                    case "hello":
                        var v46 int8 = -63
                        _ = v46
                    case "\n\t":
                        if false {
                            if v43 {
                                switch (("\n\t" + "a") + (("a" + "\n\t") + "hello")) {
                                case "hello":
                                    if len(v37) > 0 {
                                        v37[0] = fi42
                                    }
                                case "":
                                    v41 += v41
                                case "foo bar":
                                    v41 += v4
                                }
                                for fi47 := 0; fi47 < 16; fi47++ {
                                    switch "a" {
                                    case "a":
                                        var v48 int8 = 20
                                        _ = v48
                                    case "hello":
                                        v43 = (668 < 18446744073709551615)
                                    case "":
                                        fmt.Println(v41, fi47)
                                    default:
                                        v49 := GenS3{F0: v4, F1: 42, F2: -660}
                                        _ = v49
                                    }
                                    _ = fi47
                                }
                            }
                            v41 = (v4 + (v41 - (v4 | v4)))
                        }
                    default:
                        fmt.Println(v43, fi42)
                    }
                    var v50 int64 = 42
                    _ = v50
                }
            } else {
                switch -281 {
                case 0:
                    switch "hello" {
                    case "foo bar":
                        for fi51 := 0; fi51 < 5; fi51++ {
                            if (((((float64((3.14159 + 1e10)) + ((0.0 * 1e10) / 3.14159)) * (-2.25 + 0.0)) * float64(1e10)) / 0.1) != 3.14159) {
                                var v52 int8 = -1
                                v52 -= -122
                                _ = v52
                            } else {
                                _ = len(v1)
                                var v53 uint16 = (v3 * v3)
                                v54 := GenS5{F0: 835}
                                var v55 int16 = -1
                                if false {
                                    switch ("\n\t" + (("" + ((("foo bar" + "a") + "a") + "\n\t")) + ("" + ""))) {
                                    case "hello":
                                        var v56 uint64 = 169
                                        _ = v56
                                    default:
                                        if true {
                                            v55 -= 173
                                            var v57 float32 = 0.1
                                            _ = v57
                                        } else {
                                            var v58 bool = v43
                                            if (((("foo bar" + "a") + "a") + (("a" + "foo bar") + "foo bar")) > ("" + ("" + ("\n\t" + "a")))) {
                                                v59 := GenS2{F0: (v3 + 245)}
                                                if false {
                                                    var v60 string = (("\n\t" + ("a" + "")) + "a")
                                                    switch fi42 {
                                                    case -413:
                                                        var v61 int = 42
                                                        _ = v61
                                                    }
                                                    v2[-642] = 32768
                                                    v53 += v3
                                                    _ = v60
                                                } else {
                                                    var v62 uint32 = 1
                                                    _ = v62
                                                }
                                                _ = v59
                                            } else {
                                                var v63 uint32 = 536
                                                _ = v63
                                            }
                                            var v64 int64 = -9223372036854775808
                                            _ = v58
                                            _ = v64
                                        }
                                    }
                                    v41 = 4294967295
                                    v3 = (v3 >> ((v4) & 63))
                                } else {
                                    v65 := GenS1{F0: fi42, F1: v43, F2: 1}
                                    _ = v65
                                }
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            v43 = (v43 || v43)
                            if (v43 || ((v43 || (268 != 210)) && (1 != 42))) {
                                v41 += (v41 % ((798) | 1))
                                if v43 {
                                    var v66 uint32 = uint32(v4)
                                    var v67 uint8 = uint8(v3)
                                    v68 := GenS6{F0: 32767}
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                            } else {
                                switch (-(fi42)) {
                                case 32767:
                                    var v69 int16 = -271
                                    _ = v69
                                case 0:
                                    v70 := GenS6{F0: 0}
                                    _ = v70
                                }
                            }
                            _ = fi51
                        }
                    case "a":
                        var v71 int32 = -437
                        _ = v71
                    }
                case 1:
                    if len(v37) > 0 {
                        v37[0] = fi42
                    }
                default:
                    fmt.Println(v4, fi42)
                }
                var v72 uint64 = 1
                _ = v72
            }
            switch (^(fi42)) {
            case -9223372036854775808:
                v73 := GenS6{F0: (^(v3))}
                _ = v73
            default:
                v41 += v4
            }
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    switch "a" {
    case "":
        v74 := GenS7{F0: -231, F1: 255}
        _ = v74
    case "foo bar":
        switch 233 {
        case 930:
            v3 = (v3 ^ (v3 | (v3 << ((v4) & 63))))
        case 0:
            if false {
                v75 := GenS2{F0: v3}
                fmt.Println(v41, v4)
                for fi76 := 0; fi76 < 9; fi76++ {
                    var v77 float32 = 0.0
                    switch "a" {
                    case "\n\t":
                        v41 = (^(v4))
                    case "foo bar":
                        if len(v37) > 0 {
                            v37[0] = fi76
                        }
                    case "hello":
                        v3 += (v3 << ((v41) & 63))
                    default:
                        v78 := GenS8{F0: 1}
                        _ = v78
                    }
                    _ = fi76
                    _ = v77
                }
                _ = v75
            }
        default:
            if (32767 <= 32767) {
                var v79 int = -9223372036854775808
                v3 += (v3 & (^(v3)))
                fmt.Println(v4, v41)
                _ = v79
            }
        }
    }
    switch 45 {
    case 32767:
        for fi80 := 0; fi80 < 5; fi80++ {
            v41 = uint(v3)
            var v81 float32 = float32(float64(0.0))
            if true {
                switch "" {
                case "hello":
                    if ((^(v41)) == (v41 / ((v4) | 1))) {
                        fmt.Println(v41, v81, v3)
                        var v82 bool = (!false)
                        if false {
                            var v83 uint = 826
                            _ = v83
                        } else {
                            if (42 > 0) {
                                var v84 int8 = 42
                                fmt.Println(v84)
                                _ = v84
                            }
                            switch fi80 {
                            case 32768:
                                var v85 string = "foo bar"
                                _ = v85
                            case -9223372036854775808:
                                if v82 {
                                    if len(v37) > 0 {
                                        v37[0] = fi80
                                    }
                                    var v86 int64 = 1
                                    for fi87 := 0; fi87 < 2; fi87++ {
                                        v41 += v41
                                        if true {
                                            var v88 int8 = 100
                                            v89 := GenS9{F0: -2147483648, F1: -2147483648}
                                            v90 := GenS10{F0: (^(v3))}
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                        } else {
                                            var v91 uint32 = 42
                                            v4 = uint(v91)
                                            v3 = 388
                                            _ = v91
                                        }
                                        var v92 float64 = 1.5
                                        if v82 {
                                            v82 = v82
                                            fmt.Println(v3)
                                            var v93 uint16 = (v3 * (v3 % ((42) | 1)))
                                            switch "a" {
                                            case "\n\t":
                                                v93 -= (^(v3))
                                            case "":
                                                switch "foo bar" {
                                                case "":
                                                    v94 := GenS11{F0: v4, F1: 2147483647, F2: (v3 >> ((53) & 63))}
                                                    _ = v94
                                                default:
                                                    v82 = v82
                                                }
                                                v95 := GenS3{F0: 255, F1: 100, F2: 0}
                                                v96 := GenS6{F0: (v93 * (v93 + 765))}
                                                v81 = (float32((1e10 * ((0.0 * (float64((((1.5 * float32((float64(((v81 / (v81 * ((v81 - (0.0 - ((float32((0.0 - (0.0 + v92))) / (v81 - float32((v92 * v92)))) - v81))) + ((((((v81 / v81) / 1.5) + (((float32(v92) * 1.5) - v81) * 3.14159)) / (float32((v92 + -2.25)) - (0.0 - (float32(v92) * ((0.1 * 1e10) - (float32(v92) * (float32(v92) - 3.14159))))))) * 1.5) / v81)))) + -2.25)) + ((v92 * (v92 + v92)) - (float64(float32(3.14159)) - float64((1.5 * (((float32(v92) * v81) - v81) / float32((v92 + (v92 - v92))))))))))) / 0.1) + (v81 / v81))) - -2.25)) - v92))) / float32((float64(v81) / 1e10)))
                                                _ = v95
                                                _ = v96
                                            case "hello":
                                                v97 := GenS12{F0: v93, F1: 0}
                                                _ = v97
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v41 = v41
                                                case "":
                                                    v98 := GenS13{F0: fi87}
                                                    switch fi87 {
                                                    case 127:
                                                        v99 := GenS1{F0: 2147483647, F1: v82, F2: 42}
                                                        _ = v99
                                                    case 32768:
                                                        if len(v37) > 0 {
                                                            v37[0] = fi80
                                                        }
                                                    default:
                                                        v82 = true
                                                    }
                                                    if len(v37) > 0 {
                                                        v37[0] = -1
                                                    }
                                                    _ = v98
                                                }
                                            }
                                            fmt.Println(v41)
                                            _ = v93
                                        } else {
                                            var v100 uint64 = 882
                                            var v101 int16 = genFn1(-128)
                                            _ = v100
                                            _ = v101
                                        }
                                        _ = len(v37)
                                        _ = fi87
                                        _ = v92
                                    }
                                    _ = v86
                                } else {
                                    if false {
                                        fmt.Println(v41, fi80)
                                        var v102 uint64 = 1
                                        _ = v102
                                    }
                                    var v103 uint8 = 128
                                    v41 += (v41 | v4)
                                    _ = v103
                                }
                            case 128:
                                var v104 int16 = 42
                                _ = v104
                            default:
                                var v105 string = (("a" + "") + "hello")
                                _ = v105
                            }
                        }
                        _ = v82
                    } else {
                        v106 := GenS12{F0: (v3 * (^(v3))), F1: 1}
                        var v107 float32 = float32(-2.25)
                        _ = v106
                        _ = v107
                    }
                case "\n\t":
                    var v108 bool = false
                    _ = v108
                case "":
                    var v109 int32 = 1
                    _ = v109
                default:
                    v41 = (v41 % (((v41 | (^(v41)))) | 1))
                }
                v110 := GenS14{F0: (v3 / ((32767) | 1)), F1: v3, F2: fi80}
                _ = v110
            } else {
                if false {
                    v81 += ((((float32(float64(v81)) / (v81 * v81)) - float32((float64((v81 - -2.25)) / (3.14159 + float64((1e10 - 0.1)))))) - v81) + (v81 * (-2.25 - ((v81 / float32(float64((v81 * v81)))) + float32((1.5 / 1e10))))))
                }
                switch "foo bar" {
                case "":
                    var v111 uint = uint(v3)
                    _ = v111
                }
            }
            _ = fi80
            _ = v81
        }
    case 4294967295:
        v41 = v4
    case 32768:
        v112 := GenS11{F0: (v41 << ((4294967295) & 63)), F1: 1, F2: 2}
        _ = v112
    default:
        if true {
            var v113 int16 = genFn1(0)
            _ = v113
        } else {
            switch 9223372036854775807 {
            case 4294967295:
                if true {
                    v4 -= (v4 & 109)
                    switch 2147483647 {
                    case -2147483648:
                        v114 := GenS12{F0: (v3 << ((v4) & 63)), F1: 1}
                        _ = v114
                    case 32767:
                        v37 = append(v37, 769)
                    case 148:
                        for fi115 := 0; fi115 < 12; fi115++ {
                            switch 594 {
                            case 2147483647:
                                for fi116 := 0; fi116 < 1; fi116++ {
                                    var v117 bool = false
                                    var v118 int64 = int64(fi115)
                                    _ = fi116
                                    _ = v117
                                    _ = v118
                                }
                            default:
                                fmt.Println(v4, fi115, v4)
                            }
                            _ = fi115
                        }
                    }
                    v4 = uint(v3)
                } else {
                    switch "" {
                    case "a":
                        for fi119 := 0; fi119 < 13; fi119++ {
                            var v120 bool = (!true)
                            _ = fi119
                            _ = v120
                        }
                    default:
                        v4 += (v41 >> (((v4 | v41)) & 63))
                    }
                }
            case -230:
                if false {
                    if (-2147483648 >= 0) {
                        switch 557 {
                        case 4294967295:
                            if true {
                                v2[-368] = -176
                                switch ("hello" + "") {
                                case "":
                                    v41 = (v4 & (v41 % ((v4) | 1)))
                                case "hello":
                                    v121 := GenS15{F0: 809, F1: (602 > 1)}
                                    _ = v121
                                case "foo bar":
                                    fmt.Println(v3, v41)
                                }
                                if (false && false) {
                                    v4 -= v4
                                } else {
                                    v122 := GenS9{F0: 42, F1: -367}
                                    fmt.Println(v3, v4, v4)
                                    _ = v122
                                }
                            }
                        }
                        var v123 uint32 = 1
                        _ = v123
                    }
                    v124 := GenS11{F0: v4, F1: 350, F2: (^(v3))}
                    v125 := GenS10{F0: (^(v3))}
                    _ = v124
                    _ = v125
                } else {
                    switch ("" + "") {
                    case "a":
                        v41 -= v41
                    case "hello":
                        var v126 int64 = -9223372036854775808
                        _ = v126
                    default:
                        var v127 int = -2147483648
                        _ = v127
                    }
                }
            case 366:
                v4 = (v41 ^ v4)
            }
            v128 := GenS3{F0: v4, F1: -1, F2: -1}
            _ = v128
        }
    }
    switch 42 {
    case 0:
        v129 := GenS2{F0: (v3 ^ (^(v3)))}
        _ = v129
    case 32767:
        if false {
            if (1000 < 32767) {
                v130 := GenS2{F0: 32767}
                v131 := GenS16{F0: v41, F1: 0}
                _ = v130
                _ = v131
            }
            for fi132 := 0; fi132 < 19; fi132++ {
                if true {
                    var v133 bool = (!true)
                    switch ((("hello" + "a") + "") + "\n\t") {
                    case "foo bar":
                        var v134 string = "a"
                        _ = v134
                    case "\n\t":
                        for fi135 := 0; fi135 < 4; fi135++ {
                            var v136 int16 = 42
                            _ = fi135
                            _ = v136
                        }
                    default:
                        fmt.Println(v133)
                    }
                    var v137 uint8 = 202
                    _ = v133
                    _ = v137
                } else {
                    if false {
                        v3 -= v3
                    }
                }
                _ = fi132
            }
        }
    }
    var v138 uint8 = 128
    var v139 uint32 = 486
    if (false && true) {
        var v140 int = 754
        if true {
            fmt.Println(v4)
        } else {
            if true {
                if (!false) {
                    delete(v1, v139)
                }
            } else {
                fmt.Println(v41, v41)
            }
            fmt.Println(v139, v139)
        }
        _ = v140
    } else {
        v141 := GenS5{F0: genFn1(42)}
        _ = v141
    }
    fmt.Println(v138, v41, v4)
    if true {
        if false {
            for fi142 := 0; fi142 < 2; fi142++ {
                switch "hello" {
                case "hello":
                    switch (("\n\t" + (("hello" + (("hello" + "hello") + "a")) + "foo bar")) + ("" + ("a" + ("a" + "foo bar")))) {
                    case "hello":
                        if ((!(-32768 > 424)) && false) {
                            v143 := GenS17{F0: false, F1: v138}
                            var v144 string = "a"
                            var v145 uint16 = v3
                            _ = v143
                            _ = v144
                            _ = v145
                        } else {
                            for fi146 := 0; fi146 < 13; fi146++ {
                                var v147 uint = (^(v4))
                                v148 := GenS4{F0: 1000, F1: -2147483648}
                                _ = fi146
                                _ = v147
                                _ = v148
                            }
                        }
                    case "":
                        v139 = (v139 + 2147483647)
                    case "\n\t":
                        fmt.Println(v138)
                    }
                default:
                    v1[(v139 - v139)] = (v139 / ((v139) | 1))
                }
                fmt.Println(v139)
                _ = fi142
            }
        }
        if ("\n\t" == "a") {
            v149 := GenS18{F0: -2147483648, F1: (v3 + v3), F2: -1}
            v3 = v3
            _ = v149
        }
    }
    var v150 float32 = float32((0.0 - 0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v150:", v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v138
    _ = v139
    _ = v150
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
}
