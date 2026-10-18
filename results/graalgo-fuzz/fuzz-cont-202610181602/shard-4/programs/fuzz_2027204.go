package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 bool; F1 int64; F2 int8 }
type GenS10 struct { F0 bool; F1 int8; F2 int16 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS13 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int32; F1 string; F2 int }
type GenS16 struct { F0 int16; F1 uint32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint64) int16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: 255}
        switch (fi4 ^ fi4) {
        case 32767:
            if true {
                v6 := []int{fi4}
                var v7 int16 = 721
                _ = v6
                _ = v7
            } else {
                if (true || true) {
                    var v8 float64 = (3.14159 - 3.14159)
                    for fi9 := 0; fi9 < 13; fi9++ {
                        if false {
                            var v10 float32 = (0.0 - (((0.0 + float32((3.14159 * (v8 * (v8 / (v8 / ((float64(((0.0 - (3.14159 - 3.14159)) - 3.14159)) + v8) * v8))))))) / 1.0) * 1e10))
                            v8 = 1e10
                            _ = v10
                        } else {
                            var v11 uint = uint(p0)
                            _ = v11
                        }
                        _ = fi9
                    }
                    switch (fi4 * int(p0)) {
                    case 127:
                        p0 += p0
                    default:
                        p0 = 18446744073709551615
                    }
                    _ = v8
                } else {
                    var v12 bool = (354 < 996)
                    p0 = (p0 ^ 42)
                    _ = v12
                }
                p0 += (p0 >> ((396) & 63))
                v13 := GenS2{F0: 0, F1: 775}
                _ = v13
            }
        case -460:
            p0 = (p0 & (p0 & p0))
        default:
            if ((((-2147483648 < -2147483648) || (true && true)) || true) && (!false)) {
                var v14 float32 = (((((-2.25 - float32((((-2.25 - (0.0 + 0.1)) / (3.14159 * 0.1)) - 0.1))) * (float32((float64(((3.14159 - -2.25) + -2.25)) + 0.0)) + (0.1 / 0.1))) + 1e10) * 3.14159) / float32(0.1))
                _ = v14
            }
        }
        _ = fi4
        _ = v5
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        switch "" {
        case "\n\t":
            p0 = p0
        default:
            v16 := GenS3{F0: (p0 % ((p0) | 1))}
            _ = v16
        }
        var v17 uint32 = 34
        for fi18 := 0; fi18 < 12; fi18++ {
            if (true || false) {
                var v19 int = fi15
                switch fi15 {
                case 0:
                    switch ("hello" + "a") {
                    case "a":
                        switch "\n\t" {
                        case "a":
                            v17 -= (v17 - (v17 & v17))
                            v20 := GenS3{F0: (p0 ^ (p0 - p0))}
                            if true {
                                v21 := GenS1{F0: 255}
                                _ = v21
                            } else {
                                v22 := []int{}
                                v23 := GenS1{F0: 788}
                                if (42 < 2147483647) {
                                    _ = len(v22)
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v24 int8 = 0
                                        _ = v24
                                    case "":
                                        v19 -= (v19 >> ((255) & 63))
                                    case "hello":
                                        var v25 int16 = int16(p0)
                                        _ = v25
                                    default:
                                        fmt.Println(fi15)
                                        v19 = fi15
                                        fmt.Println(fi18, fi15, v17)
                                    }
                                } else {
                                    v17 = (v17 % (((v17 >> ((0) & 63))) | 1))
                                    v22 = append(v22, fi15)
                                }
                                _ = v22
                                _ = v23
                            }
                            var v26 int32 = 297
                            switch "foo bar" {
                            case "hello":
                                v26 = 0
                            }
                            switch "\n\t" {
                            case "hello":
                                v27 := GenS4{F0: -1, F1: "", F2: 65535}
                                _ = v27
                            default:
                                if false {
                                    v19 = (v19 + v19)
                                    fmt.Println(fi18)
                                    var v28 string = "foo bar"
                                    _ = v28
                                }
                            }
                            switch "\n\t" {
                            case "hello":
                                var v29 bool = true
                                _ = v29
                            case "":
                                var v30 int8 = 100
                                _ = v30
                            case "foo bar":
                                v26 += v26
                            default:
                                var v31 uint32 = v17
                                _ = v31
                            }
                            _ = v20
                            _ = v26
                        }
                    case "foo bar":
                        if (128 != 42) {
                            p0 += p0
                            v17 = 672
                        }
                    case "hello":
                        v32 := GenS4{F0: fi18, F1: ("foo bar" + "hello"), F2: 643}
                        _ = v32
                    default:
                        if true {
                            p0 = (^(p0))
                            if (((fi15 % (((^(fi15))) | 1)) <= fi18) && (!(102 > 30))) {
                                switch 128 {
                                case 42:
                                    var v33 float32 = (-2.25 / ((0.0 + 3.14159) * (float32(0.0) + -2.25)))
                                    _ = v33
                                case 0:
                                    v34 := []int{v19}
                                    _ = v34
                                case 127:
                                    v35 := GenS5{F0: 42}
                                    _ = v35
                                }
                                var v36 uint16 = 65535
                                _ = v36
                            }
                            fmt.Println(fi15, v17, v19)
                        } else {
                            v19 = (fi18 - fi18)
                            v37 := GenS1{F0: 42}
                            v38 := []int{128, (v19 ^ (v19 - (fi18 / (((v19 % (((fi18 % ((fi18) | 1))) | 1))) | 1))))}
                            _ = v37
                            _ = v38
                        }
                    }
                case -1:
                    switch fi15 {
                    case 32768:
                        var v39 uint64 = (^(p0))
                        _ = v39
                    case 127:
                        v17 += v17
                    case -2147483648:
                        v40 := GenS6{F0: 1}
                        _ = v40
                    }
                case 42:
                    v17 += 2
                default:
                    v19 = fi15
                }
                var v41 uint16 = 32767
                _ = v19
                _ = v41
            }
            fmt.Println(v17, v17)
            var v42 int = (fi15 - fi18)
            _ = fi18
            _ = v42
        }
        _ = fi15
        _ = v17
    }
    var v43 int16 = 0
    if ((-2.25 < float64(1e10)) && (!(-2.25 > ((0.1 + float32(0.1)) / 1.0)))) {
        var v44 uint32 = 42
        fmt.Println(p0)
        v45 := GenS1{F0: 0}
        _ = v44
        _ = v45
    }
    _ = p0
    _ = v43
    return (v43 * v43)
}
func genFn1(p0 int32, p1 int64) uint32 {
    for fi46 := 0; fi46 < 12; fi46++ {
        v47 := GenS1{F0: 2}
        v48 := GenS7{F0: p1}
        fmt.Println(fi46, fi46, p1)
        _ = fi46
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        v50 := []int{fi49}
        var v51 int = (fi49 * 328)
        p1 = -937
        _ = fi49
        _ = v50
        _ = v51
    }
    if (p0 >= p0) {
        v52 := []int{-1, 766}
        for fi53 := 0; fi53 < 4; fi53++ {
            p1 -= p1
            _ = fi53
        }
        _ = v52
    } else {
        fmt.Println(p0, p1)
        if (false || (-2.25 <= 1e10)) {
            v54 := []int{42, 200, 32767}
            switch "a" {
            case "foo bar":
                var v55 float64 = 1e10
                _ = v55
            }
            if (615 <= 1) {
                var v56 float32 = 1e10
                p0 += p0
                _ = v56
            } else {
                p0 -= 42
                fmt.Println(p0)
            }
            _ = v54
        }
    }
    var v57 uint8 = 35
    _ = p0
    _ = p1
    _ = v57
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    var v3 bool = (-361 < -9223372036854775808)
    v58 := GenS4{F0: 127, F1: "a", F2: 333}
    for fi59 := 0; fi59 < 7; fi59++ {
        v60 := []int{42}
        v61 := GenS8{F0: 181, F1: 2}
        var v62 int8 = 42
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    v3 = true
    v3 = (!v3)
    if false {
        var v63 int32 = 901
        _ = v63
    } else {
        v64 := []int{}
        switch 32768 {
        case -1:
            for fi65 := 0; fi65 < 4; fi65++ {
                if (!v3) {
                    var v66 uint = 1
                    var v67 int32 = 2147483647
                    switch (("a" + "foo bar") + "hello") {
                    case "a":
                        for fi68 := 0; fi68 < 20; fi68++ {
                            if ((0 < -50) && (346 >= int32(fi65))) {
                                delete(v1, (fi65 - int(v66)))
                            } else {
                                switch (-(fi65)) {
                                case 410:
                                    v3 = v3
                                }
                            }
                            _ = fi68
                        }
                    default:
                        v2 = (v2 | v2)
                    }
                    _ = v66
                    _ = v67
                } else {
                    var v69 uint16 = 9
                    _ = v69
                }
                var v70 uint8 = 127
                _ = len(v1)
                _ = fi65
                _ = v70
            }
        case 9223372036854775807:
            v2 = v2
        default:
            v1[-1] = "\n\t"
        }
        switch ("\n\t" + "foo bar") {
        case "a":
            delete(v1, 127)
        }
        _ = v64
    }
    var v71 int64 = 42
    if (((0.1 + 1e10) < 1.5) || false) {
        if (!((v2 ^ 106) <= (v2 % ((v2) | 1)))) {
            if v3 {
                v3 = v3
            } else {
                delete(v1, 32767)
            }
            var v72 int32 = -542
            v2 = (v2 % ((466) | 1))
            _ = v72
        } else {
            v71 = 1
            var v73 int16 = genFn0(42)
            _ = v73
        }
        for fi74 := 0; fi74 < 15; fi74++ {
            v2 = (v2 + 1)
            if (!v3) {
                for fi75 := 0; fi75 < 14; fi75++ {
                    if v3 {
                        fmt.Println(fi74, fi75)
                    } else {
                        v76 := []int{}
                        v77 := v1[fi74]
                        var v78 int = (^(fi74))
                        v79 := GenS5{F0: (v71 % (((-(v71))) | 1))}
                        if v3 {
                            var v80 uint32 = 2
                            v81 := GenS9{F0: true, F1: (v71 ^ (v71 * v71)), F2: 42}
                            _ = v80
                            _ = v81
                        } else {
                            var v82 uint32 = 1
                            if len(v76) > 0 {
                                v76[0] = fi75
                            }
                            v83 := GenS10{F0: false, F1: 89, F2: 0}
                            _ = v82
                            _ = v83
                        }
                        if (0 < int8(fi74)) {
                            v84 := GenS8{F0: 4294967295, F1: 2}
                            _ = v84
                        } else {
                            _ = len(v76)
                            if v3 {
                                v3 = true
                                switch "foo bar" {
                                case "\n\t":
                                    v78 -= v78
                                case "hello":
                                    var v85 uint32 = genFn1(-1, (v71 * (v71 >> ((2) & 63))))
                                    _ = v85
                                case "foo bar":
                                    var v86 int32 = 42
                                    _ = v86
                                default:
                                    var v87 float64 = (0.1 / 1e10)
                                    _ = v87
                                }
                                var v88 float32 = 0.0
                                _ = v88
                            }
                            v78 += (fi75 % ((-2147483648) | 1))
                        }
                        var v89 uint16 = 505
                        _ = v76
                        _ = v77
                        _ = v78
                        _ = v79
                        _ = v89
                    }
                    var v90 string = "hello"
                    fmt.Println(v2, v2)
                    v91 := v1[-9223372036854775808]
                    _ = fi75
                    _ = v90
                    _ = v91
                }
                var v92 int32 = -2147483648
                _ = v92
            }
            _ = fi74
        }
    } else {
        v93 := []int{1, 4294967295}
        var v94 uint32 = 1
        _ = v93
        _ = v94
    }
    if v3 {
        switch 127 {
        case -966:
            v95 := GenS11{F0: 2147483647, F1: 0}
            _ = v95
        }
        delete(v1, -431)
    } else {
        v3 = v3
    }
    if v3 {
        for fi96 := 0; fi96 < 6; fi96++ {
            v2 = v2
            _ = fi96
        }
        v2 -= (v2 | (v2 >> ((255) & 63)))
    } else {
        for fi97 := 0; fi97 < 12; fi97++ {
            var v98 float64 = 1e10
            if v3 {
                switch ("foo bar" + "foo bar") {
                case "":
                    var v99 uint8 = 128
                    _ = v99
                }
                fmt.Println(v2, v98, v3)
            }
            _ = fi97
            _ = v98
        }
    }
    for fi100 := 0; fi100 < 19; fi100++ {
        for fi101 := 0; fi101 < 15; fi101++ {
            if (v2 != v2) {
                if v3 {
                    v3 = (((fi100 >> ((4294967295) & 63)) == (fi101 - (fi101 << ((42) & 63)))) || v3)
                    var v102 float32 = 0.1
                    fmt.Println(fi101)
                    _ = v102
                } else {
                    if false {
                        v3 = v3
                        if (32767 >= 1) {
                            fmt.Println(fi100)
                            v103 := v1[4294967295]
                            _ = v103
                        } else {
                            v104 := GenS12{F0: 42, F1: v2, F2: (v71 | 0)}
                            v2 += 600
                            v2 = v2
                            _ = v104
                        }
                        v71 = v71
                    }
                }
                v105 := GenS13{F0: 551, F1: (v2 ^ (v2 + (v2 | (v2 & (v2 & v2))))), F2: 100}
                if (!v3) {
                    v2 -= v2
                }
                _ = v105
            } else {
                if (v3 && true) {
                    var v106 uint = 65535
                    if (v3 || v3) {
                        v71 -= (v71 & v71)
                        var v107 uint = (^(v106))
                        _ = v107
                    } else {
                        v71 = v71
                        var v108 float32 = 1e10
                        v109 := GenS14{F0: (-(fi101))}
                        _ = v108
                        _ = v109
                    }
                    _ = v106
                } else {
                    _ = len(v1)
                }
                v71 += (^(v71))
            }
            fmt.Println(fi101)
            _ = fi101
        }
        _ = fi100
    }
    v110 := []int{128, 9223372036854775807, -696}
    for fi111 := 0; fi111 < 15; fi111++ {
        for fi112 := 0; fi112 < 17; fi112++ {
            if ((-2.25 / (0.0 + (1e10 + -2.25))) <= 1.5) {
                var v113 int16 = 1000
                _ = v113
            } else {
                v114 := GenS15{F0: int32(v71), F1: ((("" + "") + "a") + "hello"), F2: 128}
                delete(v1, fi111)
                if v3 {
                    if false {
                        v71 -= 9223372036854775807
                        v115 := GenS16{F0: genFn0((v2 ^ 18446744073709551615)), F1: 2147483647}
                        fmt.Println(v71, fi112)
                        var v116 int = fi112
                        _ = v115
                        _ = v116
                    } else {
                        var v117 string = ((("hello" + "hello") + "\n\t") + "foo bar")
                        if (!true) {
                            v118 := GenS6{F0: -2147483648}
                            var v119 int = (fi111 / (((^(fi111))) | 1))
                            _ = v118
                            _ = v119
                        }
                        var v120 int32 = -2147483648
                        _ = v117
                        _ = v120
                    }
                } else {
                    if (!true) {
                        fmt.Println(v2)
                    } else {
                        v3 = (283 != 2)
                        var v121 uint32 = 623
                        _ = v121
                    }
                    v3 = true
                }
                if false {
                    var v122 int32 = 42
                    v71 += v71
                    _ = v122
                }
                _ = v114
            }
            var v123 int64 = (-(v71))
            var v124 int64 = v123
            _ = fi112
            _ = v123
            _ = v124
        }
        _ = fi111
    }
    switch "\n\t" {
    case "hello":
        switch -879 {
        case -447:
            v125 := GenS17{F0: 4}
            _ = v125
        case 32768:
            var v126 int64 = 366
            _ = v126
        case 202:
            v3 = v3
        }
    }
    var v127 int32 = 0
    v2 = 18446744073709551615
    v3 = v3
    switch ("" + "a") {
    case "\n\t":
        for fi128 := 0; fi128 < 10; fi128++ {
            for fi129 := 0; fi129 < 1; fi129++ {
                v130 := GenS15{F0: 0, F1: "\n\t", F2: 806}
                switch "hello" {
                case "foo bar":
                    v127 += (v127 * (v127 << ((65535) & 63)))
                }
                _ = fi129
                _ = v130
            }
            for fi131 := 0; fi131 < 10; fi131++ {
                v71 -= 42
                _ = fi131
            }
            _ = fi128
        }
    default:
        v71 = -143
    }
    v132 := GenS18{F0: 23, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v110
    _ = v127
    _ = v132
    _ = v2
    _ = v3
    _ = v58
    _ = v71
}
