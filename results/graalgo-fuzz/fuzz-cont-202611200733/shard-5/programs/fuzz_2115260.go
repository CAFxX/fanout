package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint64; F1 int; F2 uint8 }
type GenS8 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int8, p1 int16) uint64 {
    switch (("hello" + (("foo bar" + "") + "")) + "hello") {
    case "":
        var v7 bool = true
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                p0 = (p0 >> ((1) & 63))
                _ = fi9
            }
            var v10 uint32 = 65535
            _ = fi8
            _ = v10
        }
    case "\n\t":
        switch ("" + ("" + ("\n\t" + "a"))) {
        case "\n\t":
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := []int{}
                p0 -= 42
                _ = fi11
                _ = v12
            }
        case "hello":
            v13 := []int{9223372036854775807, 127, -230}
            _ = v13
        case "foo bar":
            switch "" {
            case "\n\t":
                p0 = p0
            default:
                v14 := []int{4294967295}
                _ = v14
            }
        }
    default:
        p0 += 127
    }
    switch 32767 {
    case 32767:
        v15 := []int{42, 0, 32767}
        _ = v15
    default:
        v16 := GenS1{F0: (0 <= 948)}
        _ = v16
    }
    switch "a" {
    case "hello":
        var v17 string = "foo bar"
        _ = v17
    case "a":
        p0 -= (p0 | (p0 * -128))
    default:
        fmt.Println(p0, p1)
    }
    if (((p0 | p0) == p0) || false) {
        if false {
            if true {
                var v18 float64 = 1e10
                _ = v18
            } else {
                var v19 int32 = 1
                _ = v19
            }
            var v20 string = ""
            _ = v20
        } else {
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 int = (fi21 | (fi21 >> ((262) & 63)))
                v23 := GenS1{F0: false}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            if (float32((float64(3.14159) / float64(1e10))) >= -2.25) {
                p1 = p1
            }
            switch "" {
            case "foo bar":
                v25 := []int{}
                _ = v25
            case "\n\t":
                v26 := GenS1{F0: ((false || ((!false) && (0.0 >= (0.0 * ((3.14159 * 1.5) + ((0.1 / 3.14159) - 3.14159)))))) || (1.5 >= (((1e10 / 1.0) / 1.0) - 0.0)))}
                _ = v26
            case "hello":
                fmt.Println(p0)
            }
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(p0, fi27, fi27)
                _ = fi27
            }
            _ = fi24
        }
        switch 9223372036854775807 {
        case -1:
            v28 := GenS2{F0: 42}
            _ = v28
        case -119:
            if (!false) {
                p1 += (^(p1))
                var v29 int = 32768
                v30 := GenS2{F0: v29}
                _ = v29
                _ = v30
            }
        case 912:
            v31 := []int{128, 9223372036854775807, 32767, int(p0)}
            _ = v31
        }
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int64 = 0
    var v2 float32 = (0.0 / -2.25)
    var v3 uint16 = 65535
    var v4 int32 = 2147483647
    v5 := make(map[uint32]int)
    var v6 int64 = (v1 - -9223372036854775808)
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 string = "foo bar"
        _ = fi32
        _ = v33
    }
    var v34 uint64 = 690
    v3 += (v3 / (((v3 | v3)) | 1))
    fmt.Println(v1, v4)
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint = 4294967295
        v37 := GenS3{F0: ("hello" + "a")}
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 int32 = v4
            if (1 <= uint32(v34)) {
                var v40 uint8 = 127
                var v41 float32 = (((((3.14159 / -2.25) / float32((0.1 - 1e10))) - (v2 / 3.14159)) + 3.14159) * float32((3.14159 - float64(v2))))
                var v42 float64 = (0.1 + (0.0 / (1.5 - 1e10)))
                _ = v40
                _ = v41
                _ = v42
            } else {
                v34 += 4294967295
                var v43 uint8 = 2
                _ = v43
            }
            v5[uint32(v36)] = fi35
            _ = fi38
            _ = v39
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v44 uint8 = 241
    v45 := []int{-386, 377, 2147483647, -909}
    for fi46 := 0; fi46 < 11; fi46++ {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v47 := v5[2]
            _ = v47
        default:
            _ = len(v5)
        }
        var v48 int16 = 1000
        fmt.Println(v2)
        _ = fi46
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        switch "\n\t" {
        case "a":
            v4 += v4
        case "hello":
            var v50 uint = 447
            _ = v50
        case "foo bar":
            var v51 uint32 = 277
            _ = v51
        default:
            var v52 string = ("hello" + "foo bar")
            _ = v52
        }
        if true {
            fmt.Println(v4, v2, v44)
            v3 += 42
        }
        _ = fi49
    }
    switch -9223372036854775808 {
    case 4294967295:
        if (1 < (v44 ^ v44)) {
            for fi53 := 0; fi53 < 15; fi53++ {
                v2 += 0.1
                _ = fi53
            }
            v54 := GenS1{F0: true}
            var v55 bool = false
            _ = v54
            _ = v55
        }
    case -847:
        var v56 int16 = -285
        _ = v56
    }
    v3 -= v3
    v2 = v2
    var v57 uint16 = (v3 % ((v3) | 1))
    if (true || false) {
        if true {
            v6 -= (v1 % ((v6) | 1))
            var v58 int8 = 23
            _ = v58
        } else {
            for fi59 := 0; fi59 < 8; fi59++ {
                v44 -= 1
                _ = len(v45)
                var v60 float64 = float64(v2)
                _ = fi59
                _ = v60
            }
            v57 += 2
        }
        v57 = 271
    } else {
        if false {
            if true {
                if false {
                    var v61 int8 = -1
                    switch 1 {
                    case -2147483648:
                        for fi62 := 0; fi62 < 5; fi62++ {
                            for fi63 := 0; fi63 < 20; fi63++ {
                                var v64 uint32 = 65535
                                var v65 int64 = -1
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            v66 := v5[65535]
                            v6 -= v1
                            _ = fi62
                            _ = v66
                        }
                    case -54:
                        v57 += uint16(v44)
                    }
                    _ = v61
                } else {
                    v5[4294967295] = 9223372036854775807
                }
                var v67 int16 = -1
                v1 += int64(v44)
                _ = v67
            } else {
                if (!false) {
                    for fi68 := 0; fi68 < 13; fi68++ {
                        switch (("\n\t" + (("hello" + "foo bar") + ("foo bar" + (("foo bar" + "foo bar") + "")))) + "a") {
                        case "foo bar":
                            if false {
                                v4 -= (v4 % (((^(v4))) | 1))
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    v70 := GenS3{F0: "\n\t"}
                                    _ = fi69
                                    _ = v70
                                }
                                _ = len(v5)
                            } else {
                                v34 -= (v34 - v34)
                                v71 := v5[4294967295]
                                _ = v71
                            }
                        case "":
                            switch fi68 {
                            case 2147483647:
                                fmt.Println(v57, fi68)
                            default:
                                var v72 string = ""
                                _ = v72
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "":
                                for fi73 := 0; fi73 < 16; fi73++ {
                                    v74 := GenS1{F0: true}
                                    switch "\n\t" {
                                    case "a":
                                        v3 = (v57 / (((v57 << ((65535) & 63))) | 1))
                                    default:
                                        switch "" {
                                        case "hello":
                                            switch "hello" {
                                            case "\n\t":
                                                if false {
                                                    var v75 uint = 42
                                                    var v76 uint = v75
                                                    var v77 uint16 = v3
                                                    fmt.Println(v6, v76, v57)
                                                    v6 -= (v1 - (v1 + 9223372036854775807))
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                            }
                                        case "\n\t":
                                            var v78 float32 = (0.1 * v2)
                                            _ = v78
                                        case "foo bar":
                                            v5[589] = (fi68 & fi73)
                                        default:
                                            v79 := GenS4{F0: (v57 / (((v57 << ((0) & 63))) | 1))}
                                            _ = v79
                                        }
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                            default:
                                fmt.Println(v57)
                            }
                        }
                        _ = fi68
                    }
                    v4 = int32(v57)
                } else {
                    var v80 bool = ((!(v57 == 0)) && false)
                    var v81 int32 = (v4 >> ((444) & 63))
                    for fi82 := 0; fi82 < 8; fi82++ {
                        if (!v80) {
                            v83 := GenS1{F0: (32767 < -156)}
                            _ = v83
                        } else {
                            for fi84 := 0; fi84 < 4; fi84++ {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                if (v44 <= v44) {
                                    v44 = 128
                                }
                                _ = fi84
                            }
                            var v85 int64 = v6
                            _ = v85
                        }
                        v86 := GenS3{F0: "a"}
                        _ = fi82
                        _ = v86
                    }
                    _ = v80
                    _ = v81
                }
                var v87 int16 = -906
                _ = v87
            }
        } else {
            switch 2147483647 {
            case 9223372036854775807:
                if len(v45) >= 2 {
                    v45 = v45[1:len(v45)]
                }
            default:
                v88 := GenS3{F0: "hello"}
                _ = v88
            }
            v1 -= v6
        }
        if false {
            var v89 int8 = 1
            v90 := GenS5{F0: v89, F1: -32768}
            _ = v89
            _ = v90
        } else {
            v4 = v4
        }
    }
    switch "" {
    case "a":
        if (v6 > v1) {
            for fi91 := 0; fi91 < 1; fi91++ {
                fmt.Println(v2, v57)
                var v92 int = fi91
                delete(v5, 4294967295)
                _ = fi91
                _ = v92
            }
        } else {
            var v93 uint8 = 255
            for fi94 := 0; fi94 < 12; fi94++ {
                for fi95 := 0; fi95 < 10; fi95++ {
                    v2 += ((v2 + float32(0.0)) / v2)
                    _ = fi95
                }
                v44 = v93
                for fi96 := 0; fi96 < 8; fi96++ {
                    var v97 string = "\n\t"
                    if (!false) {
                        v98 := GenS6{F0: 100, F1: -128}
                        _ = v98
                    } else {
                        switch v97 {
                        case "foo bar":
                            v6 += -1
                        case "a":
                            if false {
                                if (v44 > v93) {
                                    if (v57 <= (v3 | 42)) {
                                        if len(v45) > 0 {
                                            v45[0] = 32767
                                        }
                                    } else {
                                        v93 -= (v93 * (^(v44)))
                                        var v99 float64 = ((float64(v2) * 3.14159) / ((1.5 / ((1.5 + 0.1) * (1.5 / float64(v2)))) - (1.5 + float64(((-2.25 * v2) - v2)))))
                                        var v100 bool = false
                                        _ = v100
                                        _ = v99
                                    }
                                    if false {
                                        v6 += (^(v1))
                                        fmt.Println(v44, v6, v2)
                                    }
                                } else {
                                    v57 = 1
                                    var v101 uint8 = 2
                                    if (52 > 255) {
                                        if true {
                                            _ = len(v45)
                                            var v102 float64 = 1.5
                                            v44 -= (^(v101))
                                            _ = v102
                                        } else {
                                            if false {
                                                var v103 int16 = 42
                                                var v104 int32 = 0
                                                _ = v103
                                                _ = v104
                                            } else {
                                                v105 := GenS3{F0: "hello"}
                                                _ = v105
                                            }
                                        }
                                    }
                                    _ = v101
                                }
                                if true {
                                    delete(v5, 664)
                                    v44 = (v93 * (v44 + 94))
                                    fmt.Println(v44)
                                } else {
                                    v2 = float32(1.5)
                                    var v106 int8 = -107
                                    _ = v106
                                }
                                v44 += v44
                            } else {
                                switch int(v4) {
                                case 4294967295:
                                    v107 := v5[0]
                                    switch (v107 & v107) {
                                    case 0:
                                        if (v97 != v97) {
                                            if true {
                                                v34 += (^(v34))
                                                fmt.Println(fi94)
                                                v97 = v97
                                            } else {
                                                fmt.Println(v4, v44)
                                                v45 = append(v45, (v107 / (((-(fi96))) | 1)))
                                            }
                                            v44 = v93
                                            if ((-(v4)) >= v4) {
                                                if (uint32(v93) != 437) {
                                                    var v108 int = -525
                                                    v5[4294967295] = -77
                                                    v45 = append(v45, -1)
                                                    _ = v108
                                                }
                                                if true {
                                                    switch v97 {
                                                    case "a":
                                                        var v109 string = v97
                                                        _ = v109
                                                    case "\n\t":
                                                        v110 := GenS4{F0: v57}
                                                        _ = v110
                                                    case "foo bar":
                                                        v111 := GenS2{F0: fi96}
                                                        _ = v111
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        var v112 int = fi96
                                                        _ = v112
                                                    case "":
                                                        var v113 int8 = 127
                                                        _ = v113
                                                    default:
                                                        fmt.Println(v6)
                                                    }
                                                    var v114 uint16 = (v57 * (v57 / (((v3 >> ((1) & 63))) | 1)))
                                                    if len(v45) >= 2 {
                                                        v45 = v45[1:len(v45)]
                                                    }
                                                    v1 = (-(v6))
                                                    _ = v114
                                                } else {
                                                    if (v107 > fi96) {
                                                        var v115 uint32 = 280
                                                        v116 := GenS3{F0: "a"}
                                                        _ = v115
                                                        _ = v116
                                                    } else {
                                                        var v117 string = ((((("hello" + "\n\t") + (v97 + ("a" + ((("foo bar" + ((((v97 + v97) + v97) + v97) + (v97 + v97))) + v97) + "")))) + "") + "\n\t") + v97)
                                                        switch "hello" {
                                                        case "hello":
                                                            var v118 uint32 = 2147483647
                                                            _ = v118
                                                        case "\n\t":
                                                            fmt.Println(v97)
                                                        }
                                                        _ = v117
                                                    }
                                                    v119 := GenS7{F0: (v34 ^ 908), F1: fi96, F2: 2}
                                                    _ = v119
                                                }
                                                var v120 int = int(v4)
                                                _ = v120
                                            }
                                        } else {
                                            fmt.Println(v57, v4)
                                            fmt.Println(v44)
                                            v44 = 126
                                            v5[65535] = fi94
                                            var v121 int8 = 100
                                            _ = v121
                                        }
                                        var v122 float32 = 1.5
                                        if (!((350 == 786) && (!(!true)))) {
                                            fmt.Println(v44, v34)
                                            switch v97 {
                                            case "a":
                                                switch (v97 + v97) {
                                                case "":
                                                    v45 = append(v45, v107)
                                                case "hello":
                                                    fmt.Println(v4)
                                                default:
                                                    v123 := GenS4{F0: uint16(v93)}
                                                    _ = v123
                                                }
                                            case "hello":
                                                v122 = v2
                                            }
                                        }
                                        _ = v122
                                    }
                                    _ = len(v5)
                                    _ = v107
                                case -1:
                                    delete(v5, 2)
                                }
                                v2 -= v2
                                v124 := GenS8{F0: 0, F1: "", F2: (v6 ^ v1)}
                                _ = v124
                            }
                        case "hello":
                            var v125 float64 = 1e10
                            _ = v125
                        }
                    }
                    _ = fi96
                    _ = v97
                }
                _ = fi94
            }
            _ = v93
        }
    }
    switch "foo bar" {
    case "a":
        var v126 float64 = 3.14159
        _ = v126
    case "hello":
        v2 += v2
    case "":
        var v127 int32 = (v4 + (v4 | v4))
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v57
    _ = v6
}
