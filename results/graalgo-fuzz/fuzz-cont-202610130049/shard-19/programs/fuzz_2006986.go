package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 string; F2 uint64 }
type GenS5 struct { F0 int64; F1 int16; F2 int64 }
type GenS6 struct { F0 string; F1 string; F2 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int8; F1 string; F2 uint32 }
type GenS13 struct { F0 string; F1 string }
func genFn0(p0 int16, p1 int16) bool {
    p1 += (p0 + (p1 | 565))
    p0 += -1
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch ("" + "") {
            case "\n\t":
                if (!(119 >= 255)) {
                    v7 := []int{}
                    var v8 int32 = 913
                    _ = v7
                    _ = v8
                }
            case "foo bar":
                fmt.Println(fi5)
            }
            switch "" {
            case "":
                if true {
                    var v9 int64 = 0
                    v10 := GenS1{F0: 1}
                    switch "" {
                    case "foo bar":
                        var v11 float64 = 0.1
                        _ = v11
                    case "hello":
                        if (!false) {
                            var v12 int8 = 11
                            v9 += v9
                            var v13 int8 = v12
                            if (!false) {
                                v14 := GenS2{F0: 42, F1: "hello"}
                                _ = v14
                            }
                            _ = v12
                            _ = v13
                        } else {
                            p1 = (p0 + p1)
                            fmt.Println(fi6, p0)
                            var v15 int64 = v9
                            if true {
                                v9 -= -238
                            }
                            _ = v15
                        }
                    case "\n\t":
                        var v16 int8 = 42
                        _ = v16
                    }
                    _ = v10
                    _ = v9
                } else {
                    var v17 uint64 = 18446744073709551615
                    _ = v17
                }
            case "\n\t":
                var v18 uint64 = 832
                _ = v18
            case "foo bar":
                v19 := []int{(fi6 ^ (fi5 + fi5))}
                _ = v19
            }
            _ = fi6
        }
        _ = fi5
    }
    v20 := []int{-548, -2147483648, 127}
    _ = p0
    _ = p1
    _ = v20
    return (false || (true || (0 < 1)))
}
func genFn1() bool {
    var v21 float64 = 1e10
    switch "a" {
    case "":
        switch "" {
        case "hello":
            v22 := []int{-784, 0}
            _ = v22
        case "":
            for fi23 := 0; fi23 < 17; fi23++ {
                v21 = v21
                v21 = v21
                _ = fi23
            }
        case "a":
            v21 += ((0.0 / (v21 - v21)) - float64(1.5))
        default:
            v24 := []int{127, -632}
            _ = v24
        }
    }
    switch "" {
    case "hello":
        var v25 uint8 = 42
        _ = v25
    case "foo bar":
        if true {
            var v26 int = 128
            _ = v26
        }
    case "\n\t":
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 bool = ((601 != 808) && (!(-2.25 != float32(float64(float32((((float64(((0.0 / -2.25) * -2.25)) * (((float64(0.0) - float64(1.5)) - v21) + 0.0)) * v21) * (v21 - v21))))))))
            var v29 string = ("" + "a")
            if (fi27 == (fi27 % (((fi27 - fi27)) | 1))) {
                v30 := GenS2{F0: 1, F1: v29}
                v21 = -2.25
                var v31 float64 = v21
                _ = v30
                _ = v31
            } else {
                v29 = v29
            }
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v32 := []int{}
    v33 := GenS3{F0: 541}
    _ = v21
    _ = v32
    _ = v33
    return (!true)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    v3 := []int{128}
    var v4 int64 = 1
    if false {
        if true {
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := GenS4{F0: 703, F1: "foo bar", F2: 4294967295}
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 7; fi36++ {
                fmt.Println(fi36, fi36)
                _ = fi36
            }
            if true {
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 float32 = 0.1
                    _ = fi37
                    _ = v38
                }
                switch ("" + "\n\t") {
                case "foo bar":
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v40 := GenS5{F0: v4, F1: 1, F2: (v4 & (^(v4)))}
                        v41 := GenS6{F0: "foo bar", F1: ("foo bar" + "hello"), F2: 66}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                case "a":
                    v4 += v4
                case "\n\t":
                    v42 := GenS2{F0: 4294967295, F1: "a"}
                    _ = v42
                }
                var v43 uint32 = 2147483648
                _ = v43
            }
        } else {
            if (((0.1 + ((1e10 / ((1e10 / 1.5) / 0.1)) + float32(0.1))) / -2.25) != ((1.5 * 1.5) + 1e10)) {
                v44 := GenS5{F0: 42, F1: 122, F2: v4}
                var v45 uint64 = 18446744073709551615
                _ = v44
                _ = v45
            } else {
                var v46 uint64 = 1
                fmt.Println(v46, v4)
                v4 = (v4 | (v4 + -744))
                _ = v46
            }
            var v47 string = "\n\t"
            if ((v4 != (v4 >> ((425) & 63))) || false) {
                for fi48 := 0; fi48 < 12; fi48++ {
                    if (false || (1 >= 2147483647)) {
                        for fi49 := 0; fi49 < 1; fi49++ {
                            var v50 int = (fi49 - (^(fi48)))
                            _ = fi49
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 19; fi51++ {
                            if true {
                                if (v47 < v47) {
                                    v52 := GenS2{F0: 507, F1: (v47 + (v47 + ""))}
                                    _ = v52
                                } else {
                                    v4 = v4
                                    v4 -= (v4 | (v4 / (((v4 - v4)) | 1)))
                                    var v53 bool = true
                                    _ = v53
                                }
                                v47 = v47
                            } else {
                                if (!true) {
                                    switch (v47 + v47) {
                                    case "":
                                        switch ("foo bar" + v47) {
                                        case "hello":
                                            var v54 uint32 = 0
                                            _ = v54
                                        case "\n\t":
                                            v55 := v1[991]
                                            _ = v55
                                        default:
                                            v56 := GenS2{F0: 42, F1: v47}
                                            _ = v56
                                        }
                                    case "hello":
                                        fmt.Println(fi48)
                                    default:
                                        v4 += 0
                                    }
                                    _ = len(v2)
                                }
                                v57 := v1[2]
                                _ = v57
                            }
                            _ = fi51
                        }
                    } else {
                        if true {
                            for fi58 := 0; fi58 < 2; fi58++ {
                                var v59 string = ("foo bar" + ("foo bar" + ((((v47 + v47) + (v47 + "hello")) + ((v47 + "\n\t") + ((v47 + "a") + (v47 + v47)))) + "a")))
                                v60 := GenS3{F0: 0}
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            v3 = append(v3, (fi48 | (-(fi48))))
                        } else {
                            for fi61 := 0; fi61 < 16; fi61++ {
                                fmt.Println(v47, v47)
                                fmt.Println(fi48, fi61, v47)
                                var v62 uint16 = 224
                                _ = fi61
                                _ = v62
                            }
                        }
                    }
                    var v63 float32 = 1e10
                    _ = fi48
                    _ = v63
                }
                for fi64 := 0; fi64 < 3; fi64++ {
                    var v65 float32 = 1.5
                    _ = fi64
                    _ = v65
                }
                switch "a" {
                case "hello":
                    v4 = v4
                case "a":
                    if false {
                        v66 := GenS7{F0: 128}
                        switch v47 {
                        case "\n\t":
                            v67 := GenS1{F0: 0}
                            _ = v67
                        case "foo bar":
                            v47 = ((v47 + (((v47 + v47) + v47) + "foo bar")) + (v47 + v47))
                        case "a":
                            v68 := GenS4{F0: 1, F1: v47, F2: 42}
                            _ = v68
                        default:
                            for fi69 := 0; fi69 < 17; fi69++ {
                                var v70 int16 = 0
                                v70 = -1
                                _ = fi69
                                _ = v70
                            }
                        }
                        _ = v66
                    }
                default:
                    for fi71 := 0; fi71 < 14; fi71++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi71)
                        if false {
                            fmt.Println(v4, fi71, v47)
                            switch fi71 {
                            case 0:
                                v72 := GenS4{F0: int16(v4), F1: v47, F2: 4294967295}
                                _ = v72
                            default:
                                v73 := GenS8{F0: 0, F1: 816, F2: 42}
                                _ = v73
                            }
                        }
                        _ = fi71
                    }
                }
            } else {
                if false {
                    v74 := GenS2{F0: 601, F1: ((("" + "\n\t") + v47) + v47)}
                    if true {
                        if false {
                            v4 += (v4 | v4)
                            v2 = append(v2, 2147483647)
                        } else {
                            v47 = "foo bar"
                        }
                        v4 = v4
                        var v75 uint64 = 42
                        _ = v75
                    } else {
                        v47 = v47
                    }
                    for fi76 := 0; fi76 < 14; fi76++ {
                        v47 = ("\n\t" + "foo bar")
                        _ = fi76
                    }
                    _ = v74
                }
                v77 := GenS7{F0: 110}
                if false {
                    for fi78 := 0; fi78 < 20; fi78++ {
                        v79 := GenS6{F0: v47, F1: v47, F2: 1}
                        for fi80 := 0; fi80 < 16; fi80++ {
                            fmt.Println(v4, fi78, v47)
                            var v81 bool = false
                            v82 := GenS1{F0: 4294967295}
                            var v83 float64 = 3.14159
                            _ = fi80
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                        _ = fi78
                        _ = v79
                    }
                } else {
                    for fi84 := 0; fi84 < 15; fi84++ {
                        var v85 int32 = 196
                        _ = fi84
                        _ = v85
                    }
                    if len(v3) > 0 {
                        v3[0] = 69
                    }
                    if true {
                        for fi86 := 0; fi86 < 17; fi86++ {
                            var v87 uint32 = 1
                            v4 = v4
                            _ = fi86
                            _ = v87
                        }
                        for fi88 := 0; fi88 < 6; fi88++ {
                            var v89 float64 = 3.14159
                            v90 := GenS3{F0: 474}
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                        v4 = (v4 | v4)
                    } else {
                        v47 = (v47 + (v47 + "\n\t"))
                    }
                }
                _ = v77
            }
            _ = v47
        }
        v91 := GenS9{F0: 255}
        v2 = append(v2, 4294967295)
        _ = v91
    } else {
        v92 := GenS8{F0: 2, F1: 318, F2: -2147483648}
        _ = v92
    }
    for fi93 := 0; fi93 < 4; fi93++ {
        var v94 int8 = -10
        _ = fi93
        _ = v94
    }
    v4 = -9223372036854775808
    var v95 uint16 = 0
    v4 = v4
    v96 := GenS10{F0: 127}
    fmt.Println(v4)
    v4 = (v4 << ((42) & 63))
    v1[2147483648] = 65535
    if (true && (float32((((((1e10 / ((float64(float32(1e10)) + 0.1) - 3.14159)) - ((0.1 * 0.1) * (float64(-2.25) - 1.5))) * (float64(float32(1e10)) * -2.25)) / float64(float32(1e10))) + 0.1)) != 0.0)) {
        v95 = v95
        var v97 float64 = (0.0 * (1.5 + 3.14159))
        _ = v97
    } else {
        var v98 uint64 = uint64(v95)
        _ = v98
    }
    v4 = (v4 | -819)
    if false {
        if true {
            v99 := GenS6{F0: "hello", F1: ("hello" + "\n\t"), F2: 246}
            var v100 uint = 718
            if (127 == 98) {
                for fi101 := 0; fi101 < 16; fi101++ {
                    switch int(v100) {
                    case -1:
                        switch "" {
                        case "\n\t":
                            var v102 string = ("a" + "")
                            _ = v102
                        case "foo bar":
                            v95 += v95
                        default:
                            _ = len(v2)
                        }
                    case 1:
                        var v103 uint32 = 2147483648
                        _ = v103
                    }
                    if (-32768 <= 0) {
                        v95 = (v95 - (v95 | 256))
                        v104 := GenS11{F0: (v4 | (v4 % ((v4) | 1))), F1: -52}
                        _ = v104
                    } else {
                        var v105 float32 = 1e10
                        var v106 int32 = -2147483648
                        switch fi101 {
                        case 1:
                            for fi107 := 0; fi107 < 2; fi107++ {
                                var v108 uint16 = (v95 % ((v95) | 1))
                                _ = fi107
                                _ = v108
                            }
                        default:
                            v100 += 755
                        }
                        _ = v105
                        _ = v106
                    }
                    for fi109 := 0; fi109 < 19; fi109++ {
                        v4 += (v4 + (v4 >> ((v100) & 63)))
                        if true {
                            v95 += v95
                            var v110 uint = (v100 / (((v100 | v100)) | 1))
                            v4 += v4
                            _ = v110
                        }
                        _ = fi109
                    }
                    _ = fi101
                }
                switch 1 {
                case -111:
                    for fi111 := 0; fi111 < 10; fi111++ {
                        for fi112 := 0; fi112 < 3; fi112++ {
                            fmt.Println(fi111)
                            fmt.Println(fi112, fi112, fi111)
                            _ = fi112
                        }
                        var v113 uint = (v100 | 4294967295)
                        for fi114 := 0; fi114 < 5; fi114++ {
                            var v115 uint = (v100 >> ((v113) & 63))
                            if false {
                                v116 := GenS12{F0: -128, F1: "foo bar", F2: 744}
                                var v117 bool = (440 >= 42)
                                _ = v116
                                _ = v117
                            } else {
                                var v118 string = ("a" + ("" + "a"))
                                _ = v118
                            }
                            _ = fi114
                            _ = v115
                        }
                        _ = fi111
                        _ = v113
                    }
                default:
                    v95 = v95
                }
            }
            _ = v100
            _ = v99
        } else {
            if false {
                for fi119 := 0; fi119 < 7; fi119++ {
                    v95 += v95
                    v2 = append(v2, int(v4))
                    _ = fi119
                }
                var v120 string = ("hello" + "a")
                _ = v120
            } else {
                var v121 float64 = (3.14159 * float64(0.0))
                v3 = append(v3, 128)
                var v122 uint = 65535
                _ = v121
                _ = v122
            }
            if true {
                var v123 float64 = -2.25
                var v124 bool = (-821 >= 42)
                v124 = true
                _ = v123
                _ = v124
            } else {
                var v125 uint16 = 42
                _ = v125
            }
        }
        if false {
            var v126 uint32 = 42
            _ = v126
        }
    }
    for fi127 := 0; fi127 < 2; fi127++ {
        if (335 < 389) {
            if (42 > 2) {
                var v128 int8 = 42
                var v129 uint = 287
                _ = v128
                _ = v129
            } else {
                v95 -= (v95 << ((42) & 63))
                for fi130 := 0; fi130 < 9; fi130++ {
                    if (!(!false)) {
                        v131 := GenS6{F0: "hello", F1: "a", F2: 128}
                        _ = v131
                    } else {
                        v132 := GenS3{F0: 0}
                        var v133 int16 = -1
                        v95 -= 2
                        _ = v132
                        _ = v133
                    }
                    v4 = (v4 >> ((528) & 63))
                    _ = fi130
                }
            }
            v134 := v1[13]
            fmt.Println(v134, v4, v4)
            _ = v134
        } else {
            switch "a" {
            case "hello":
                for fi135 := 0; fi135 < 12; fi135++ {
                    v4 += v4
                    if len(v2) > 0 {
                        v2[0] = fi135
                    }
                    fmt.Println(fi135, fi127, v95)
                    _ = fi135
                }
            case "":
                v95 += (v95 + (v95 + v95))
            case "a":
                v136 := GenS9{F0: 844}
                _ = v136
            default:
                var v137 float32 = 3.14159
                _ = v137
            }
            var v138 float64 = 3.14159
            _ = v138
        }
        _ = fi127
    }
    v139 := GenS13{F0: "a", F1: "foo bar"}
    var v140 uint = 830
    v4 += (v4 >> ((v140) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v139
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v96
}
