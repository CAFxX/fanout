package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 string; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 int8; F1 bool }
func genFn0(p0 bool, p1 uint8) uint64 {
    switch 790 {
    case 0:
        v6 := GenS1{F0: uint(p1), F1: 42}
        _ = v6
    }
    p1 = p1
    var v7 string = "\n\t"
    v8 := GenS2{F0: 2147483648}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int32) int {
    switch ("\n\t" + "foo bar") {
    case "hello":
        if (437 == 36) {
            v9 := []int{128}
            var v10 int16 = 1000
            _ = v10
            _ = v9
        }
    case "a":
        v11 := []int{0, 32767, 0, -490}
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        p0 -= (^(p0))
        _ = fi12
    }
    _ = p0
    return 42
}

func main() {
    var v1 string = ("foo bar" + "")
    var v2 uint16 = 0
    var v3 int32 = -2147483648
    var v4 uint64 = 4294967295
    var v5 int16 = 0
    var v13 uint = 2
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := []int{(fi14 + (fi14 << ((v13) & 63))), 2147483647}
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS1{F0: 998, F1: (v2 << (((v13 ^ v13)) & 63))}
            v2 = v2
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v3 -= (-(v3))
        v2 -= (v2 - v2)
        v19 := []int{(fi18 ^ (^(fi18))), fi18}
        _ = fi18
        _ = v19
    }
    if (-128 >= 42) {
        fmt.Println(v1)
        switch genFn1(2147483647) {
        case 0:
            v20 := GenS2{F0: 2}
            _ = v20
        case 42:
            if true {
                for fi21 := 0; fi21 < 20; fi21++ {
                    if ((-2.25 + 0.0) > (1.5 * 1.5)) {
                        v22 := GenS3{F0: 42, F1: false, F2: 42}
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := GenS2{F0: 42}
                        _ = fi23
                        _ = v24
                    }
                    _ = fi21
                }
            }
        case -2147483648:
            switch genFn1(2147483647) {
            case 32768:
                for fi25 := 0; fi25 < 11; fi25++ {
                    v2 -= v2
                    v4 += 92
                    _ = fi25
                }
            }
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            for fi27 := 0; fi27 < 11; fi27++ {
                if true {
                    v3 = v3
                    v3 = v3
                    switch fi26 {
                    case 127:
                        var v28 float32 = (((-2.25 * -2.25) / float32(((0.0 + 1e10) * float64((float32(1.5) + 1.5))))) - (0.0 * 1.5))
                        _ = v28
                    case -1:
                        if false {
                            switch v1 {
                            case "hello":
                                switch v1 {
                                case "hello":
                                    if false {
                                        switch v1 {
                                        case "hello":
                                            v13 -= v13
                                        case "\n\t":
                                            v3 = int32(v2)
                                        case "foo bar":
                                            if true {
                                                var v29 int = fi27
                                                var v30 int16 = v5
                                                switch ((v1 + ((((v1 + ("a" + ("\n\t" + v1))) + (("a" + (v1 + v1)) + v1)) + v1) + "hello")) + v1) {
                                                case "hello":
                                                    switch v29 {
                                                    case 128:
                                                        switch v29 {
                                                        case -2147483648:
                                                            if true {
                                                                var v31 uint64 = v4
                                                                if (v2 >= 0) {
                                                                    switch fi27 {
                                                                    case 1:
                                                                        var v32 uint32 = uint32(v31)
                                                                        _ = v32
                                                                    }
                                                                    var v33 bool = ((v31 + (v4 ^ v4)) == v31)
                                                                    v34 := []int{genFn1(int32(v13)), fi26, fi27, 128}
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                    var v35 int8 = 100
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                    _ = v33
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                _ = v31
                                                            }
                                                        case -873:
                                                            v36 := []int{190, (fi26 | -863)}
                                                            _ = v36
                                                        case 4294967295:
                                                            if (((!(int64(v29) > 42)) && false) && (-1 <= -78)) {
                                                                var v37 float32 = (0.0 / (((float32((((((3.14159 + 3.14159) * (0.1 + -2.25)) + (float64(((1.5 / -2.25) * 0.0)) - 3.14159)) + ((1e10 + (float64(1.5) * 0.0)) * 3.14159)) * 0.1)) / 1e10) / (-2.25 * float32(((-2.25 / 0.1) * float64(0.1))))) - 1e10))
                                                                switch (fi27 - fi26) {
                                                                case 32768:
                                                                    switch v1 {
                                                                    case "hello":
                                                                        if false {
                                                                            var v38 uint = (v13 * (v13 - (v13 ^ v13)))
                                                                            if (((((1.5 - 0.0) / (3.14159 + 1.5)) * 3.14159) / 1.0) < 1e10) {
                                                                                v39 := []int{1, 774, (fi26 + genFn1(v3))}
                                                                                _ = v39
                                                                            }
                                                                            _ = v38
                                                                        }
                                                                    case "a":
                                                                        var v40 float32 = (float32(-2.25) + v37)
                                                                        _ = v40
                                                                    }
                                                                case 128:
                                                                    fmt.Println(fi26, v29)
                                                                case -1:
                                                                    v41 := []int{genFn1(2147483647), genFn1(-562)}
                                                                    _ = v41
                                                                default:
                                                                    fmt.Println(v13)
                                                                }
                                                                _ = v37
                                                            } else {
                                                                v5 -= (-(v30))
                                                            }
                                                        }
                                                    case -2147483648:
                                                        v42 := []int{}
                                                        _ = v42
                                                    case 104:
                                                        fmt.Println(v13, v29)
                                                    }
                                                case "":
                                                    v29 = (fi26 << ((v13) & 63))
                                                    var v43 int32 = 0
                                                    _ = v43
                                                case "foo bar":
                                                    var v44 int16 = 1
                                                    _ = v44
                                                }
                                                if false {
                                                    v45 := GenS4{F0: (v3 - (v3 * 712))}
                                                    _ = v45
                                                }
                                                _ = v29
                                                _ = v30
                                            } else {
                                                var v46 uint32 = 840
                                                v47 := GenS1{F0: (v13 + 1), F1: (v2 % ((uint16(v46)) | 1))}
                                                _ = v46
                                                _ = v47
                                            }
                                        }
                                        fmt.Println(v4, v4, v2)
                                    } else {
                                        switch (fi27 & (fi26 ^ 4294967295)) {
                                        case -1:
                                            switch "" {
                                            case "":
                                                switch (v1 + (v1 + v1)) {
                                                case "\n\t":
                                                    var v48 uint8 = 162
                                                    _ = v48
                                                case "hello":
                                                    fmt.Println(v13, v1)
                                                    v1 = v1
                                                case "foo bar":
                                                    fmt.Println(fi26, fi26)
                                                }
                                            }
                                        case -440:
                                            var v49 int = -1
                                            if false {
                                                var v50 uint64 = v4
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                    }
                                case "":
                                    switch fi27 {
                                    case -9223372036854775808:
                                        var v51 int32 = v3
                                        _ = v51
                                    }
                                case "foo bar":
                                    if (int64(v5) < int64(v3)) {
                                        v52 := GenS2{F0: 65}
                                        _ = v52
                                    } else {
                                        v53 := []int{}
                                        v2 -= 255
                                        v4 += v4
                                        _ = v53
                                    }
                                default:
                                    v13 -= (v13 ^ (v13 + v13))
                                }
                            }
                            switch fi26 {
                            case 0:
                                switch "hello" {
                                case "\n\t":
                                    if (!(false && (true || (3.14159 < 0.1)))) {
                                        if false {
                                            v4 += (v4 + 42)
                                        } else {
                                            var v54 int32 = (-(v3))
                                            var v55 int64 = -752
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(v4)
                                    fmt.Println(v5, v3)
                                case "":
                                    v4 = (v4 | v4)
                                default:
                                    fmt.Println(v4, v2)
                                }
                            default:
                                var v56 uint = v13
                                v57 := GenS1{F0: v13, F1: v2}
                                _ = v56
                                _ = v57
                            }
                            v13 = (v13 - 18446744073709551615)
                        } else {
                            if (100 >= -2) {
                                var v58 float32 = (((3.14159 * (float32(((0.1 + 3.14159) * -2.25)) / 1.0)) / 1.0) / float32((1e10 / -2.25)))
                                var v59 string = "\n\t"
                                _ = v58
                                _ = v59
                            } else {
                                fmt.Println(v4, v2)
                            }
                        }
                    default:
                        var v60 bool = false
                        if (false && (-43 <= v5)) {
                            if (((v1 + "\n\t") != "\n\t") && false) {
                                if (0.1 != 0.1) {
                                    v61 := GenS5{F0: v2}
                                    fmt.Println(v60, fi26)
                                    fmt.Println(v5, v1)
                                    switch ((v1 + (v1 + v1)) + ((v1 + v1) + v1)) {
                                    case "":
                                        fmt.Println(v3, v2, v60)
                                    case "hello":
                                        var v62 uint16 = (v2 - 42)
                                        v13 += (v13 / (((v13 << ((2) & 63))) | 1))
                                        v63 := []int{(^(fi26)), genFn1((v3 % ((v3) | 1))), (fi26 + genFn1(v3)), fi26}
                                        _ = v62
                                        _ = v63
                                    default:
                                        v13 = (v13 % ((v13) | 1))
                                        v64 := GenS1{F0: (v13 / (((^(v13))) | 1)), F1: 65535}
                                        switch v1 {
                                        case "foo bar":
                                            v1 = v1
                                        case "a":
                                            v65 := []int{(fi26 | 32768), fi27, 0, (^(fi27))}
                                            _ = v65
                                        }
                                        _ = v64
                                    }
                                    _ = v61
                                }
                                var v66 int = fi27
                                var v67 uint = v13
                                if v60 {
                                    if v60 {
                                        var v68 int64 = -646
                                        fmt.Println(v2, v5, v13)
                                        var v69 float64 = ((0.1 + (1.5 / -2.25)) + float64(1e10))
                                        fmt.Println(v67, v2, v13)
                                        _ = v68
                                        _ = v69
                                    } else {
                                        var v70 int32 = -824
                                        _ = v70
                                    }
                                    fmt.Println(v67, v13, fi26)
                                } else {
                                    var v71 float64 = 1.5
                                    var v72 uint32 = 42
                                    _ = v71
                                    _ = v72
                                }
                                v73 := GenS6{F0: 1, F1: "hello", F2: 255}
                                v66 = (v66 ^ genFn1(v3))
                                v74 := GenS3{F0: v3, F1: v60, F2: 1}
                                _ = v66
                                _ = v67
                                _ = v73
                                _ = v74
                            }
                            if v60 {
                                fmt.Println(fi26, v5)
                                var v75 uint16 = 53
                                var v76 int16 = (v5 & v5)
                                _ = v75
                                _ = v76
                            }
                        } else {
                            v1 = "a"
                        }
                        v77 := GenS7{F0: v4, F1: (fi26 & 1)}
                        _ = v60
                        _ = v77
                    }
                    var v78 string = ""
                    var v79 int64 = int64(v2)
                    v80 := []int{fi26}
                    var v81 int16 = v5
                    _ = v78
                    _ = v79
                    _ = v80
                    _ = v81
                } else {
                    var v82 uint32 = 938
                    var v83 int16 = v5
                    _ = v82
                    _ = v83
                }
                v84 := GenS8{F0: v5, F1: v2, F2: (v4 / ((uint64(v2)) | 1))}
                v85 := []int{fi26, (fi27 | genFn1((v3 | v3))), fi26}
                _ = fi27
                _ = v84
                _ = v85
            }
            var v86 uint32 = 65535
            _ = fi26
            _ = v86
        }
    }
    fmt.Println(v5)
    if true {
        v4 -= (v4 - v4)
    } else {
        for fi87 := 0; fi87 < 20; fi87++ {
            v88 := GenS9{F0: (^(v4)), F1: 1}
            v89 := GenS10{F0: (^(v4))}
            var v90 int16 = (-(v5))
            _ = fi87
            _ = v88
            _ = v89
            _ = v90
        }
        if (v2 > 0) {
            var v91 int64 = 9223372036854775807
            v13 -= (v13 ^ (v13 << (((v13 >> ((v13) & 63))) & 63)))
            _ = v91
        } else {
            for fi92 := 0; fi92 < 3; fi92++ {
                for fi93 := 0; fi93 < 2; fi93++ {
                    if (2 <= 4294967295) {
                        fmt.Println(v4)
                    } else {
                        if (true && true) {
                            if false {
                                var v94 uint8 = 2
                                var v95 uint64 = (v4 << ((uint(v94)) & 63))
                                v96 := GenS9{F0: 823, F1: (v4 << ((v13) & 63))}
                                switch int(v3) {
                                case 0:
                                    var v97 int16 = (v5 & v5)
                                    _ = v97
                                default:
                                    var v98 uint16 = (v2 & v2)
                                    _ = v98
                                }
                                v99 := []int{(fi92 + (fi93 << (((v13 * v13)) & 63))), fi93}
                                var v100 uint16 = uint16(v4)
                                _ = v100
                                _ = v94
                                _ = v95
                                _ = v96
                                _ = v99
                            }
                        }
                        var v101 int16 = (^(v5))
                        _ = v101
                    }
                    _ = fi93
                }
                v102 := GenS9{F0: v4, F1: genFn0((true && true), 42)}
                _ = fi92
                _ = v102
            }
        }
    }
    v103 := GenS10{F0: (v4 * v4)}
    v5 += (v5 << ((v13) & 63))
    v5 = (-(v5))
    v4 = (v4 ^ (v4 - v4))
    if (!false) {
        v3 = (-(v3))
        var v104 bool = true
        switch -606 {
        case 414:
            switch 483 {
            case 369:
                if v104 {
                    v105 := []int{563, -2147483648}
                    _ = v105
                } else {
                    var v106 string = (v1 + v1)
                    _ = v106
                }
            case 32768:
                v107 := []int{0, 127, -945}
                _ = v107
            case 127:
                for fi108 := 0; fi108 < 11; fi108++ {
                    v109 := []int{int(v4)}
                    for fi110 := 0; fi110 < 12; fi110++ {
                        var v111 uint16 = (^(v2))
                        _ = fi110
                        _ = v111
                    }
                    _ = fi108
                    _ = v109
                }
            default:
                switch (v1 + v1) {
                case "\n\t":
                    var v112 int8 = -1
                    _ = v112
                case "hello":
                    var v113 float32 = (0.0 - ((-2.25 * float32(((0.0 / (3.14159 / 3.14159)) + -2.25))) - float32(1e10)))
                    _ = v113
                default:
                    v104 = true
                }
            }
        default:
            for fi114 := 0; fi114 < 18; fi114++ {
                for fi115 := 0; fi115 < 3; fi115++ {
                    if v104 {
                        v3 += (v3 & (v3 ^ (-(v3))))
                    } else {
                        var v116 uint = v13
                        switch v1 {
                        case "":
                            v117 := GenS11{F0: 23, F1: fi114}
                            _ = v117
                        case "hello":
                            var v118 int16 = v5
                            _ = v118
                        case "\n\t":
                            if v104 {
                                v2 = 2
                            }
                        }
                        v119 := []int{fi115, -944, 42, (fi114 % ((genFn1(v3)) | 1))}
                        _ = v116
                        _ = v119
                    }
                    _ = fi115
                }
                v3 = (v3 & v3)
                for fi120 := 0; fi120 < 10; fi120++ {
                    v121 := GenS4{F0: v3}
                    _ = fi120
                    _ = v121
                }
                _ = fi114
            }
        }
        _ = v104
    }
    for fi122 := 0; fi122 < 13; fi122++ {
        var v123 int16 = v5
        for fi124 := 0; fi124 < 6; fi124++ {
            v125 := GenS12{F0: 85, F1: true}
            switch ("hello" + "\n\t") {
            case "\n\t":
                v126 := []int{(fi122 >> ((v13) & 63)), (fi122 & genFn1((v3 % ((v3) | 1))))}
                _ = v126
            case "":
                fmt.Println(v4, v3)
            default:
                v1 = v1
            }
            v1 = (("hello" + ("\n\t" + v1)) + v1)
            _ = fi124
            _ = v125
        }
        fmt.Println(v123, v3)
        _ = fi122
        _ = v123
    }
    switch (v1 + "foo bar") {
    case "a":
        v127 := []int{}
        _ = v127
    case "hello":
        var v128 int8 = -1
        _ = v128
    default:
        v3 = v3
    }
    v3 = v3
    for fi129 := 0; fi129 < 4; fi129++ {
        fmt.Println(v13)
        v5 = (-(v5))
        _ = fi129
    }
    var v130 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v13
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
