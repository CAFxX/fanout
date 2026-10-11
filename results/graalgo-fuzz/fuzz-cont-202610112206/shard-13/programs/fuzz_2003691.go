package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 uint64; F1 int16; F2 int64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32 }
func genFn0(p0 int16) int64 {
    var v7 int8 = -1
    var v8 int64 = 42
    p0 += (^(p0))
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return (v8 >> ((65535) & 63))
}

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = 9223372036854775807
    var v3 bool = false
    var v4 bool = v3
    var v5 uint16 = 423
    var v6 float64 = 0.0
    v9 := GenS1{F0: (v1 ^ v1)}
    v4 = (!(v5 > (v5 << ((65535) & 63))))
    var v10 float32 = ((1.5 - ((0.0 / 1e10) / (1.5 + ((3.14159 * (((0.0 + 1.5) / 1.0) + 3.14159)) / 3.14159)))) - (float32(0.0) + (float32(0.1) / 1.5)))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint = 1
        var v13 uint32 = 2147483648
        _ = fi11
        _ = v12
        _ = v13
    }
    switch (("" + "hello") + "\n\t") {
    case "a":
        var v14 float32 = (v10 + 1.5)
        _ = v14
    case "foo bar":
        if ((v5 + v5) == v5) {
            if v4 {
                switch -121 {
                case 1:
                    v15 := []int{42, -9223372036854775808}
                    _ = v15
                case -2147483648:
                    v3 = (!(!v4))
                case -36:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 uint16 = v5
                            _ = fi17
                            _ = v18
                        }
                        var v19 int8 = int8(fi16)
                        _ = fi16
                        _ = v19
                    }
                default:
                    v20 := []int{int(v5)}
                    _ = v20
                }
                var v21 uint32 = 2147483647
                _ = v21
            } else {
                if false {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v1 = v1
                        for fi23 := 0; fi23 < 14; fi23++ {
                            fmt.Println(fi23)
                            var v24 uint64 = 4294967295
                            _ = fi23
                            _ = v24
                        }
                        _ = fi22
                    }
                }
                if (v4 && (-904 > -562)) {
                    if (!(v2 != genFn0(int16(v2)))) {
                        if v4 {
                            if (0 >= 255) {
                                v25 := GenS1{F0: 1}
                                switch (("\n\t" + ("" + "a")) + (("\n\t" + "\n\t") + "a")) {
                                case "a":
                                    v26 := []int{152, 4294967295, int(v1)}
                                    _ = v26
                                case "\n\t":
                                    if false {
                                        var v27 string = ("" + ("\n\t" + "foo bar"))
                                        _ = v27
                                    } else {
                                        for fi28 := 0; fi28 < 9; fi28++ {
                                            switch "hello" {
                                            case "\n\t":
                                                switch "hello" {
                                                case "a":
                                                    if (!(((-(v2)) >= genFn0(-1)) || v3)) {
                                                        for fi29 := 0; fi29 < 8; fi29++ {
                                                            var v30 bool = v4
                                                            var v31 uint = 1
                                                            _ = fi29
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        fmt.Println(fi28)
                                                    } else {
                                                        v32 := GenS1{F0: 42}
                                                        if (65535 == 178) {
                                                            v10 = v10
                                                            for fi33 := 0; fi33 < 13; fi33++ {
                                                                v34 := []int{}
                                                                if false {
                                                                    var v35 int64 = v2
                                                                    _ = v35
                                                                }
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                            var v36 int32 = v1
                                                            _ = v36
                                                        }
                                                        v37 := []int{560, (fi28 >> ((852) & 63))}
                                                        _ = v32
                                                        _ = v37
                                                    }
                                                case "":
                                                    for fi38 := 0; fi38 < 14; fi38++ {
                                                        v6 -= v6
                                                        v39 := []int{42, fi28}
                                                        fmt.Println(fi28, v3, v5)
                                                        _ = fi38
                                                        _ = v39
                                                    }
                                                case "hello":
                                                    v40 := GenS1{F0: -1}
                                                    _ = v40
                                                default:
                                                    var v41 int8 = int8(v1)
                                                    _ = v41
                                                }
                                            case "a":
                                                if (v3 || (2 == 255)) {
                                                    var v42 float64 = v6
                                                    var v43 uint = 255
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    if (-532 <= -917) {
                                                        var v44 int16 = 32767
                                                        _ = v44
                                                    }
                                                }
                                            default:
                                                var v45 string = "hello"
                                                _ = v45
                                            }
                                            _ = fi28
                                        }
                                        v46 := GenS1{F0: v1}
                                        v1 -= v1
                                        _ = v46
                                    }
                                default:
                                    var v47 int16 = 469
                                    _ = v47
                                }
                                if (!v3) {
                                    fmt.Println(v10)
                                } else {
                                    switch 2147483647 {
                                    case -2147483648:
                                        switch 2147483647 {
                                        case -640:
                                            fmt.Println(v4, v3)
                                        case 1:
                                            if ((678 != 4294967295) || (((uint8(v5) < 241) && v4) || (false || false))) {
                                                v5 += v5
                                                v48 := GenS2{F0: 1, F1: v4, F2: 775}
                                                switch "hello" {
                                                case "hello":
                                                    v49 := []int{127}
                                                    _ = v49
                                                case "foo bar":
                                                    if (v3 && true) {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v6, v1, v10)
                                                        case "a":
                                                            switch 2147483647 {
                                                            case -669:
                                                                var v50 float64 = 3.14159
                                                                _ = v50
                                                            }
                                                        case "\n\t":
                                                            for fi51 := 0; fi51 < 8; fi51++ {
                                                                v52 := []int{(-(fi51)), fi51, fi51}
                                                                _ = fi51
                                                                _ = v52
                                                            }
                                                        default:
                                                            var v53 uint = 42
                                                            _ = v53
                                                        }
                                                    } else {
                                                        for fi54 := 0; fi54 < 3; fi54++ {
                                                            v1 += (v1 ^ (v1 ^ (v1 * 317)))
                                                            _ = fi54
                                                        }
                                                        v55 := GenS1{F0: v1}
                                                        _ = v55
                                                    }
                                                case "":
                                                    var v56 uint = 42
                                                    _ = v56
                                                }
                                                _ = v48
                                            } else {
                                                v3 = true
                                                v57 := GenS3{F0: 2147483647, F1: uint64(v5)}
                                                _ = v57
                                            }
                                        case -9223372036854775808:
                                            var v58 float64 = 0.1
                                            _ = v58
                                        }
                                    case 2147483647:
                                        var v59 int16 = 1000
                                        _ = v59
                                    case 0:
                                        v1 = v1
                                    default:
                                        var v60 uint8 = 127
                                        _ = v60
                                    }
                                }
                                _ = v25
                            } else {
                                var v61 uint32 = 2147483648
                                var v62 bool = v3
                                _ = v61
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 12; fi63++ {
                                v64 := GenS4{F0: 2, F1: 42}
                                var v65 string = "foo bar"
                                var v66 uint16 = v5
                                _ = fi63
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                        }
                    } else {
                        v67 := GenS3{F0: 0, F1: 18446744073709551615}
                        var v68 int64 = 9223372036854775807
                        var v69 int32 = v1
                        _ = v67
                        _ = v68
                        _ = v69
                    }
                    switch -9223372036854775808 {
                    case 387:
                        var v70 uint32 = 42
                        _ = v70
                    }
                    v71 := []int{-9223372036854775808, -1, 32768}
                    _ = v71
                }
            }
            var v72 uint8 = 255
            var v73 uint = 2
            _ = v72
            _ = v73
        }
    case "\n\t":
        for fi74 := 0; fi74 < 18; fi74++ {
            v2 = v2
            _ = fi74
        }
    default:
        fmt.Println(v10, v6, v4)
    }
    switch 0 {
    case 2147483647:
        for fi75 := 0; fi75 < 6; fi75++ {
            var v76 uint = 42
            var v77 int32 = (v1 - 2147483647)
            v78 := []int{}
            _ = fi75
            _ = v76
            _ = v77
            _ = v78
        }
    case -1:
        fmt.Println(v2, v4)
    }
    var v79 int32 = (v1 >> ((18446744073709551615) & 63))
    v3 = (2147483647 <= 0)
    switch ("a" + (("foo bar" + "\n\t") + (("\n\t" + ("" + "")) + "foo bar"))) {
    case "foo bar":
        if v4 {
            for fi80 := 0; fi80 < 17; fi80++ {
                v81 := []int{272}
                _ = fi80
                _ = v81
            }
            fmt.Println(v3, v10)
            var v82 int = -680
            _ = v82
        } else {
            v83 := []int{4294967295, -114}
            _ = v83
        }
    case "\n\t":
        for fi84 := 0; fi84 < 1; fi84++ {
            v85 := GenS5{F0: "\n\t", F1: fi84}
            switch ("hello" + ("hello" + ("hello" + ("a" + (("\n\t" + "a") + ((("a" + "foo bar") + "foo bar") + ("hello" + "\n\t"))))))) {
            case "foo bar":
                if true {
                    v86 := GenS5{F0: "\n\t", F1: (fi84 ^ int(v79))}
                    v87 := GenS6{F0: 180}
                    _ = v86
                    _ = v87
                } else {
                    if false {
                        for fi88 := 0; fi88 < 6; fi88++ {
                            if false {
                                var v89 uint32 = 42
                                var v90 uint16 = (^(v5))
                                v1 = -301
                                _ = v89
                                _ = v90
                            } else {
                                v1 -= v79
                                if (!(!(!(!v4)))) {
                                    switch "a" {
                                    case "foo bar":
                                        if v4 {
                                            var v91 int = 4294967295
                                            if (!true) {
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(fi88, v6, v2)
                                                case "\n\t":
                                                    v10 += v10
                                                case "":
                                                    fmt.Println(v1, fi88)
                                                default:
                                                    if (!true) {
                                                        v2 += (v2 | (v2 >> ((255) & 63)))
                                                        if true {
                                                            if (1 <= 2) {
                                                                v79 -= 0
                                                            } else {
                                                                v79 = (v1 / ((2147483647) | 1))
                                                                switch "a" {
                                                                case "hello":
                                                                    if false {
                                                                        v4 = v4
                                                                        v5 += 991
                                                                        v91 = 2147483647
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                v92 := GenS7{F0: 65535, F1: fi88}
                                                                _ = v92
                                                            case "\n\t":
                                                                if true {
                                                                    if (!(v10 < (((v10 / v10) - float32(v6)) - float32(float64((0.0 / float32(v6))))))) {
                                                                        var v93 int16 = 0
                                                                        _ = v93
                                                                    } else {
                                                                        fmt.Println(fi88, v10, v79)
                                                                        var v94 uint8 = 2
                                                                        _ = v94
                                                                    }
                                                                    if v4 {
                                                                        fmt.Println(v1)
                                                                        v95 := GenS1{F0: (v79 + -1)}
                                                                        fmt.Println(v4, v3)
                                                                        _ = v95
                                                                    } else {
                                                                        if true {
                                                                            v3 = v3
                                                                            var v96 uint64 = 150
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    var v97 uint8 = 127
                                                                    _ = v97
                                                                } else {
                                                                    var v98 uint8 = 42
                                                                    _ = v98
                                                                }
                                                                v99 := GenS8{F0: int16(v2)}
                                                                _ = v99
                                                            default:
                                                                if false {
                                                                    if true {
                                                                        var v100 int = fi84
                                                                        v101 := []int{}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    var v102 uint8 = uint8(v5)
                                                                    _ = v102
                                                                } else {
                                                                    v103 := []int{v91}
                                                                    if (v5 >= v5) {
                                                                        v2 = v2
                                                                        v103 = append(v103, fi84)
                                                                    } else {
                                                                        v2 = (v2 + 0)
                                                                        if (!(!(v10 == (0.0 / (v10 / v10))))) {
                                                                            v104 := GenS7{F0: 457, F1: int(v2)}
                                                                            fmt.Println(fi84, v79)
                                                                            var v105 int = -653
                                                                            if len(v103) >= 2 {
                                                                                v103 = v103[1:len(v103)]
                                                                            }
                                                                            fmt.Println(v2)
                                                                            _ = len(v103)
                                                                            _ = v104
                                                                            _ = v105
                                                                        } else {
                                                                            fmt.Println(v1, v5, v2)
                                                                            if len(v103) >= 2 {
                                                                                v103 = v103[1:len(v103)]
                                                                            }
                                                                            var v106 uint32 = 42
                                                                            _ = v106
                                                                        }
                                                                    }
                                                                    _ = v103
                                                                }
                                                            }
                                                            var v107 int16 = -1
                                                            _ = v107
                                                        }
                                                        var v108 float64 = v6
                                                        _ = v108
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    var v109 int16 = 1
                                                    _ = v109
                                                }
                                                v110 := GenS9{F0: 1, F1: v79}
                                                _ = v110
                                            } else {
                                                v91 = (fi84 >> ((1) & 63))
                                                v3 = (v5 != v5)
                                                v1 -= (v1 / ((int32(v5)) | 1))
                                            }
                                            _ = v91
                                        }
                                    }
                                    v111 := GenS3{F0: 65535, F1: 2}
                                    v10 = (v10 + float32(v6))
                                    _ = v111
                                }
                                v112 := GenS10{F0: 18446744073709551615, F1: 1000, F2: (v2 & (v2 / ((genFn0(180)) | 1)))}
                                _ = v112
                            }
                            switch (-(fi84)) {
                            case -2147483648:
                                var v113 float32 = v10
                                _ = v113
                            case 2147483647:
                                var v114 uint32 = 65535
                                _ = v114
                            case 128:
                                switch fi88 {
                                case 195:
                                    v3 = true
                                case 32767:
                                    var v115 uint64 = 477
                                    _ = v115
                                }
                            default:
                                v2 = (v2 ^ genFn0(-32768))
                            }
                            _ = fi88
                        }
                        var v116 bool = v3
                        var v117 uint16 = 32767
                        _ = v116
                        _ = v117
                    } else {
                        v2 -= (^(v2))
                    }
                    var v118 bool = v3
                    var v119 int16 = -32768
                    _ = v118
                    _ = v119
                }
            case "hello":
                v120 := GenS11{F0: 128}
                _ = v120
            case "":
                switch "hello" {
                case "\n\t":
                    var v121 int64 = (v2 | 9223372036854775807)
                    _ = v121
                }
            }
            _ = fi84
            _ = v85
        }
    case "hello":
        fmt.Println(v79)
    default:
        v122 := []int{2147483647, 32768, 127}
        _ = v122
    }
    v123 := GenS9{F0: 1, F1: v79}
    var v124 uint8 = 208
    if false {
        switch 2147483647 {
        case -9223372036854775808:
            switch ((((((("a" + "") + "\n\t") + "a") + "") + "") + "") + "foo bar") {
            case "\n\t":
                for fi125 := 0; fi125 < 9; fi125++ {
                    v124 += v124
                    if true {
                        v126 := []int{2147483647, 127}
                        v127 := GenS3{F0: 2, F1: 4294967295}
                        v128 := GenS8{F0: -1}
                        _ = v126
                        _ = v127
                        _ = v128
                    }
                    for fi129 := 0; fi129 < 1; fi129++ {
                        v2 += (v2 + (v2 & genFn0(442)))
                        v130 := GenS2{F0: (v124 & 109), F1: false, F2: 303}
                        fmt.Println(v6, v6)
                        _ = fi129
                        _ = v130
                    }
                    _ = fi125
                }
            }
        case -298:
            v1 -= (v79 >> ((18446744073709551615) & 63))
        case 9223372036854775807:
            fmt.Println(v1, v5, v1)
        default:
            if (42 > -1) {
                switch -395 {
                case 9223372036854775807:
                    v10 = 1e10
                }
            } else {
                v131 := GenS12{F0: 90}
                fmt.Println(v1, v5, v4)
                if v3 {
                    fmt.Println(v6, v124, v4)
                    v6 = v6
                }
                _ = v131
            }
        }
        var v132 int16 = -32768
        _ = v132
    } else {
        v133 := []int{0, 0}
        fmt.Println(v10)
        fmt.Println(v1)
        _ = v133
    }
    switch ("foo bar" + "") {
    case "a":
        v134 := GenS13{F0: v1}
        _ = v134
    case "hello":
        v4 = (true && (0 <= 32768))
    case "\n\t":
        fmt.Println(v10, v6, v6)
    default:
        v5 -= (^(v5))
    }
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v9
}
