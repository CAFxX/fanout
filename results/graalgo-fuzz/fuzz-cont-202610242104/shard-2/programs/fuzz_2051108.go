package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 string; F2 string }
type GenS8 struct { F0 int; F1 uint64; F2 int32 }
type GenS9 struct { F0 int8; F1 int16; F2 int8 }
type GenS10 struct { F0 int16; F1 bool }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 int8; F1 uint64; F2 int8 }
type GenS13 struct { F0 uint; F1 int; F2 int16 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch fi5 {
        case -72:
            if true {
                v6 := []int{fi5, fi5, 128}
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                if true {
                    var v7 uint = 255
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v9 := GenS1{F0: 42}
                        _ = fi8
                        _ = v9
                    }
                    if (558 >= 469) {
                        v7 += (v7 * v7)
                        var v10 int = (fi5 | -2147483648)
                        _ = v10
                    } else {
                        switch "" {
                        case "hello":
                            var v11 uint8 = 1
                            _ = v11
                        }
                        if (255 <= 42) {
                            v7 += v7
                            fmt.Println(fi5, v7)
                        } else {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v7 += (v7 ^ v7)
                                _ = fi12
                            }
                            var v13 float32 = float32((((((1e10 - 1.5) * 0.1) / 1.0) / (1e10 - float64(1.5))) - (0.0 * ((float64(0.1) + 0.1) * float64(0.0)))))
                            for fi14 := 0; fi14 < 10; fi14++ {
                                v15 := GenS2{F0: 863, F1: -32768, F2: -186}
                                if (false || false) {
                                    var v16 uint = 65535
                                    v17 := GenS1{F0: 255}
                                    _ = v16
                                    _ = v17
                                } else {
                                    _ = len(v6)
                                    fmt.Println(fi5)
                                    if true {
                                        if (true || (!(18446744073709551615 != 4294967295))) {
                                            if ("hello" == "a") {
                                                if len(v6) > 0 {
                                                    v6[0] = 440
                                                }
                                                switch "a" {
                                                case "":
                                                    var v18 uint32 = 2
                                                    if ((float64(v13) <= float64(v13)) || false) {
                                                        var v19 uint64 = 18446744073709551615
                                                        v6 = append(v6, -813)
                                                        var v20 uint = 728
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        v7 += 2
                                                    }
                                                    var v21 float32 = v13
                                                    _ = v18
                                                    _ = v21
                                                case "a":
                                                    switch 311 {
                                                    case -1:
                                                        v7 = v7
                                                    case 128:
                                                        if false {
                                                            v22 := GenS3{F0: 42}
                                                            fmt.Println(fi14)
                                                            fmt.Println(fi14, fi5)
                                                            v7 = v7
                                                            _ = v22
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v23 string = ("a" + "hello")
                                                    _ = v23
                                                default:
                                                    var v24 int = (fi5 * 127)
                                                    _ = v24
                                                }
                                                v25 := GenS4{F0: 0, F1: -1}
                                                _ = v25
                                            } else {
                                                switch fi5 {
                                                case -365:
                                                    var v26 float32 = (v13 + v13)
                                                    _ = v26
                                                case 732:
                                                    v7 -= (v7 >> ((v7) & 63))
                                                case 32768:
                                                    v7 = v7
                                                default:
                                                    var v27 int8 = 87
                                                    _ = v27
                                                }
                                                fmt.Println(fi5, v7, fi14)
                                                v13 += float32(1.5)
                                                switch ("\n\t" + "") {
                                                case "a":
                                                    var v28 int16 = -365
                                                    _ = v28
                                                case "":
                                                    var v29 uint = (v7 + (v7 / (((v7 | v7)) | 1)))
                                                    _ = v29
                                                }
                                            }
                                            if (65535 != 2) {
                                                if false {
                                                    var v30 string = ("" + "")
                                                    fmt.Println(fi5, v30, fi5)
                                                    _ = v30
                                                } else {
                                                    v31 := GenS5{F0: ("" + "hello"), F1: 0}
                                                    _ = v31
                                                }
                                            }
                                            fmt.Println(v7, v13)
                                        }
                                        var v32 uint32 = 2
                                        if true {
                                            var v33 int16 = 1
                                            v34 := GenS2{F0: 1, F1: (v33 & (v33 / ((v33) | 1))), F2: 0}
                                            var v35 int64 = 9223372036854775807
                                            fmt.Println(v13)
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v32
                                    } else {
                                        fmt.Println(v13, v13)
                                        var v36 float64 = (0.0 * (-2.25 + -2.25))
                                        v36 -= float64(v13)
                                        var v37 int16 = 32767
                                        v38 := GenS3{F0: 42}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                v13 = (v13 * (1.5 - v13))
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                    }
                    _ = v7
                }
                _ = v6
            } else {
                if false {
                    v39 := GenS2{F0: 42, F1: -980, F2: -2147483648}
                    v40 := []int{-2147483648, fi5, (fi5 / (((-(fi5))) | 1)), fi5}
                    switch -9223372036854775808 {
                    case -897:
                        var v41 uint8 = 76
                        _ = v41
                    }
                    _ = v39
                    _ = v40
                } else {
                    if (100 >= -89) {
                        v42 := GenS3{F0: -9223372036854775808}
                        v43 := GenS6{F0: 211}
                        var v44 int16 = -805
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        v45 := GenS7{F0: 1, F1: "\n\t", F2: (("hello" + ("foo bar" + "foo bar")) + "foo bar")}
                        _ = v45
                    }
                    var v46 uint32 = 65535
                    _ = v46
                }
                var v47 int16 = -1
                _ = v47
            }
        default:
            var v48 uint32 = 0
            _ = v48
        }
        fmt.Println(fi5)
        v49 := GenS3{F0: 9223372036854775807}
        _ = fi5
        _ = v49
    }
    v50 := GenS3{F0: -9223372036854775808}
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS6{F0: 1}
        v53 := []int{(^(fi51))}
        _ = fi51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        var v55 float32 = (0.0 / ((1e10 * float32(1e10)) + 3.14159))
        var v56 int32 = 466
        _ = fi54
        _ = v55
        _ = v56
    }
    if (!false) {
        v57 := []int{}
        if (float64(1.5) < (float64(float32(1e10)) * (0.0 * (1.5 + (3.14159 + 0.0))))) {
            for fi58 := 0; fi58 < 7; fi58++ {
                for fi59 := 0; fi59 < 13; fi59++ {
                    v60 := GenS4{F0: 2147483647, F1: 1}
                    if (false || true) {
                        var v61 uint16 = 255
                        switch ("hello" + "hello") {
                        case "":
                            switch fi59 {
                            case -628:
                                fmt.Println(fi59)
                            case 32768:
                                var v62 int16 = -603
                                _ = v62
                            default:
                                v57 = append(v57, (fi59 - fi58))
                                if len(v57) > 0 {
                                    v57[0] = 4294967295
                                }
                                _ = len(v57)
                                switch 4294967295 {
                                case 2147483647:
                                    switch (fi59 + 167) {
                                    case 2147483647:
                                        v61 += v61
                                    case -9223372036854775808:
                                        if true {
                                            var v63 int64 = 42
                                            _ = v63
                                        }
                                    case -953:
                                        _ = len(v57)
                                    }
                                }
                            }
                        }
                        _ = v61
                    } else {
                        var v64 int8 = 100
                        var v65 int32 = 982
                        var v66 int16 = 42
                        fmt.Println(v65)
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    if (!true) {
                        fmt.Println(fi59)
                    }
                    _ = fi59
                    _ = v60
                }
                _ = fi58
            }
            if (-1 != 0) {
                var v67 uint64 = 42
                _ = v67
            }
            for fi68 := 0; fi68 < 19; fi68++ {
                var v69 uint64 = 2
                _ = fi68
                _ = v69
            }
        }
        var v70 int = -259
        _ = v57
        _ = v70
    } else {
        var v71 int32 = 2147483647
        var v72 uint32 = 4294967295
        _ = v71
        _ = v72
    }
    _ = v50
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 uint8 = 146
    if true {
        v1 = 1
        var v73 float32 = (((float32(((0.0 / 3.14159) * 0.1)) - 1.5) + ((float32((-2.25 + (1.5 + 3.14159))) - -2.25) + float32(((-2.25 - 0.0) - -2.25)))) + float32(-2.25))
        _ = v73
    } else {
        for fi74 := 0; fi74 < 19; fi74++ {
            v1 += v1
            switch 128 {
            case 42:
                v4 = (^(v4))
            }
            if true {
                v4 = v4
                for fi75 := 0; fi75 < 5; fi75++ {
                    switch "a" {
                    case "\n\t":
                        var v76 float64 = (float64((float32(3.14159) - float32((1e10 / (3.14159 + 1e10))))) + (3.14159 * -2.25))
                        _ = v76
                    case "hello":
                        switch (fi74 >> ((238) & 63)) {
                        case 1:
                            delete(v2, (v1 ^ v1))
                        case -510:
                            var v77 bool = (!(-677 == -9223372036854775808))
                            _ = v77
                        case 2147483647:
                            switch fi74 {
                            case 128:
                                v78 := GenS8{F0: (fi74 / (((^(fi74))) | 1)), F1: 0, F2: -2147483648}
                                _ = v78
                            }
                        }
                        v79 := GenS6{F0: 42}
                        if (!(42 <= (v1 - (v1 - 2147483647)))) {
                            v1 = v1
                            var v80 int32 = 1
                            _ = v80
                        }
                        _ = v79
                    case "":
                        var v81 uint16 = 255
                        _ = v81
                    default:
                        fmt.Println(fi75, v4)
                    }
                    _ = fi75
                }
            }
            _ = fi74
        }
        fmt.Println(v4, v4)
        v1 += v1
    }
    switch 0 {
    case 42:
        v1 += v1
    case -308:
        fmt.Println(v1, v1)
    }
    v4 -= (v4 & 127)
    v1 += (v1 | v1)
    v1 = v1
    v82 := GenS9{F0: -1, F1: -32768, F2: 42}
    if (1 > 42) {
        var v83 uint32 = (v1 & (^(v1)))
        switch -189 {
        case -1:
            var v84 int = 0
            _ = v84
        }
        var v85 bool = false
        _ = v83
        _ = v85
    } else {
        v86 := []int{128, -2147483648, 745}
        _ = v86
    }
    for fi87 := 0; fi87 < 11; fi87++ {
        var v88 string = "hello"
        fmt.Println(fi87, fi87, v4)
        switch 9223372036854775807 {
        case -230:
            var v89 float64 = float64(1e10)
            _ = v89
        default:
            var v90 int32 = 1
            _ = v90
        }
        _ = fi87
        _ = v88
    }
    v1 -= v1
    v91 := []int{-719, -1, -2147483648}
    if false {
        if true {
            v92 := GenS10{F0: 428, F1: ("hello" == (("foo bar" + (("a" + "hello") + "hello")) + "a"))}
            _ = v92
        } else {
            switch (((("a" + "hello") + ("\n\t" + ("hello" + "\n\t"))) + "") + "a") {
            case "hello":
                v1 += v1
            case "foo bar":
                v4 = v4
            case "":
                var v93 int = 32767
                _ = v93
            }
        }
    } else {
        delete(v3, 134)
    }
    var v94 string = "a"
    if true {
        v95 := GenS2{F0: -2147483648, F1: 1, F2: -907}
        _ = v95
    } else {
        var v96 uint64 = 42
        v97 := GenS5{F0: v94, F1: -1}
        _ = v96
        _ = v97
    }
    switch "" {
    case "\n\t":
        var v98 float64 = (float64((0.0 / (3.14159 * float32(1e10)))) + (3.14159 + 3.14159))
        _ = v98
    }
    switch "" {
    case "\n\t":
        if len(v91) > 0 {
            v91[0] = 4294967295
        }
    default:
        v99 := GenS9{F0: 127, F1: 466, F2: 125}
        _ = v99
    }
    switch (v94 + v94) {
    case "hello":
        for fi100 := 0; fi100 < 17; fi100++ {
            for fi101 := 0; fi101 < 9; fi101++ {
                v102 := GenS11{F0: 42, F1: 42}
                _ = fi101
                _ = v102
            }
            var v103 uint8 = (v4 / ((42) | 1))
            _ = fi100
            _ = v103
        }
    case "":
        var v104 float64 = 1e10
        _ = v104
    case "a":
        if ((1 != 952) || (!true)) {
            var v105 uint8 = v4
            if false {
                v106 := GenS12{F0: -57, F1: 1, F2: 40}
                _ = len(v91)
                fmt.Println(v4, v4)
                _ = v106
            }
            for fi107 := 0; fi107 < 1; fi107++ {
                var v108 uint16 = 767
                v94 = (((v94 + v94) + v94) + "\n\t")
                _ = fi107
                _ = v108
            }
            _ = v105
        } else {
            var v109 int = -1
            var v110 int64 = 9223372036854775807
            var v111 bool = (v1 == 2147483647)
            _ = v109
            _ = v110
            _ = v111
        }
    default:
        switch 467 {
        case 128:
            for fi112 := 0; fi112 < 6; fi112++ {
                switch v94 {
                case "hello":
                    fmt.Println(v1)
                case "":
                    if false {
                        switch fi112 {
                        case -2147483648:
                            if len(v91) >= 2 {
                                v91 = v91[1:len(v91)]
                            }
                        case 923:
                            v94 = "\n\t"
                        }
                        var v113 int8 = 42
                        _ = v113
                    }
                case "a":
                    if (((v4 % ((v4) | 1)) >= v4) || true) {
                        v114 := GenS10{F0: 136, F1: (!(true || false))}
                        var v115 string = v94
                        _ = v114
                        _ = v115
                    }
                }
                var v116 float32 = -2.25
                v116 = 3.14159
                _ = fi112
                _ = v116
            }
        case 950:
            v117 := GenS13{F0: 424, F1: 150, F2: 1}
            _ = v117
        default:
            var v118 int16 = -1
            _ = v118
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v91
    _ = v94
}
