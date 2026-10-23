package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS10 struct { F0 int8; F1 int; F2 string }
type GenS11 struct { F0 uint32; F1 string; F2 int64 }
type GenS12 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS13 struct { F0 bool; F1 int8; F2 bool }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint; F1 int32; F2 uint16 }
type GenS16 struct { F0 string; F1 uint16; F2 uint64 }
type GenS17 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]bool)
    var v4 uint16 = 316
    v5 := GenS1{F0: 1, F1: -22}
    fmt.Println(v4, v2)
    v2 = v2
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint64 = 2
        v8 := []int{2147483647}
        var v9 string = "foo bar"
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float32 = (3.14159 / ((1.5 - ((0.1 * (0.0 * 1e10)) / 1e10)) / 3.14159))
    var v11 int8 = -128
    fmt.Println(v4, v2)
    if (true || (!(1 == uint32(v2)))) {
        v4 = (v4 % ((65535) | 1))
        v10 = v10
        fmt.Println(v4)
    }
    var v12 uint32 = 1
    switch ("\n\t" + "foo bar") {
    case "hello":
        v13 := []int{778, 128}
        _ = v13
    case "a":
        if true {
            if false {
                v14 := GenS2{F0: 1, F1: 255}
                v2 = 0
                var v15 uint64 = 18446744073709551615
                _ = v14
                _ = v15
            } else {
                if (v2 < (v2 + (v2 / ((65535) | 1)))) {
                    v1[32768] = (v12 & v12)
                    v16 := GenS3{F0: (v2 + 18446744073709551615), F1: 9223372036854775807, F2: -1}
                    v2 += (v2 * (v2 + (v2 + 42)))
                    _ = v16
                }
                if (!(float32(3.14159) >= 3.14159)) {
                    delete(v3, v12)
                    for fi17 := 0; fi17 < 4; fi17++ {
                        switch "foo bar" {
                        case "":
                            v11 += 122
                        }
                        fmt.Println(fi17, v4)
                        v4 += v4
                        _ = fi17
                    }
                } else {
                    v1[6] = (v12 & v12)
                }
            }
            v4 -= v4
        } else {
            if true {
                var v18 int8 = (-(v11))
                var v19 uint64 = 2
                v4 = v4
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 5; fi20++ {
                    switch (fi20 >> (((^(v2))) & 63)) {
                    case 752:
                        for fi21 := 0; fi21 < 6; fi21++ {
                            switch fi21 {
                            case 0:
                                switch fi21 {
                                case 4294967295:
                                    if (false || (true && (!(float64(v10) >= -2.25)))) {
                                        var v22 uint8 = 1
                                        fmt.Println(v2, fi21, v22)
                                        v23 := GenS4{F0: v4, F1: 418, F2: -1}
                                        _ = v22
                                        _ = v23
                                    }
                                case 32768:
                                    if true {
                                        v10 = v10
                                    }
                                default:
                                    v24 := GenS5{F0: -2147483648}
                                    _ = v24
                                }
                            case 9223372036854775807:
                                var v25 int64 = -1
                                _ = v25
                            }
                            if ((-(fi21)) <= fi20) {
                                switch ("hello" + ("" + "a")) {
                                case "":
                                    v26 := v3[(v12 % ((v12) | 1))]
                                    _ = v26
                                case "hello":
                                    var v27 string = "\n\t"
                                    _ = v27
                                case "\n\t":
                                    var v28 uint = v2
                                    _ = v28
                                default:
                                    if true {
                                        v29 := GenS4{F0: v4, F1: 0, F2: 9223372036854775807}
                                        if (uint64(v12) <= 721) {
                                            var v30 float64 = 3.14159
                                            _ = v30
                                        } else {
                                            var v31 uint64 = 18446744073709551615
                                            var v32 int = fi21
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v29
                                    } else {
                                        v33 := GenS6{F0: 42}
                                        switch ((("hello" + "\n\t") + "hello") + "a") {
                                        case "a":
                                            v4 = 173
                                        case "hello":
                                            v34 := GenS5{F0: 1}
                                            _ = v34
                                        case "\n\t":
                                            v35 := GenS7{F0: 0, F1: v11}
                                            _ = v35
                                        default:
                                            v36 := []int{fi20, (fi21 ^ (fi21 % (((fi21 * 42)) | 1))), (fi20 % (((fi20 % ((0) | 1))) | 1)), (fi20 >> ((4294967295) & 63))}
                                            _ = v36
                                        }
                                        if (!true) {
                                            switch 4294967295 {
                                            case -9223372036854775808:
                                                switch fi21 {
                                                case -32:
                                                    v37 := []int{(^(fi20)), fi21, fi20}
                                                    _ = v37
                                                case 2147483647:
                                                    if true {
                                                        var v38 uint = v2
                                                        if (2 > uint8(v4)) {
                                                            switch ("hello" + "foo bar") {
                                                            case "":
                                                                delete(v3, v12)
                                                            case "a":
                                                                v39 := GenS5{F0: 1}
                                                                _ = v39
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                            v40 := GenS3{F0: v38, F1: fi20, F2: 1}
                                                            fmt.Println(v38)
                                                            _ = v40
                                                        }
                                                        _ = v38
                                                    } else {
                                                        if (false && (!((true || (-95 > -1)) && false))) {
                                                            switch "" {
                                                            case "":
                                                                v41 := GenS6{F0: -32768}
                                                                _ = v41
                                                            case "hello":
                                                                var v42 int8 = (-(v11))
                                                                _ = v42
                                                            }
                                                            v11 = (v11 / ((v11) | 1))
                                                            if (0 <= 42) {
                                                                var v43 uint8 = 42
                                                                switch ("" + "foo bar") {
                                                                case "foo bar":
                                                                    v44 := GenS4{F0: (v4 & v4), F1: uint64(v2), F2: 1}
                                                                    _ = v44
                                                                case "a":
                                                                    if (("foo bar" == ("" + "\n\t")) || false) {
                                                                        v10 = ((v10 * v10) + 3.14159)
                                                                        if true {
                                                                            if ("foo bar" > "a") {
                                                                                var v45 uint64 = 42
                                                                                v46 := v3[v12]
                                                                                var v47 uint16 = (^(v4))
                                                                                _ = v45
                                                                                _ = v46
                                                                                _ = v47
                                                                            } else {
                                                                                v4 = 0
                                                                            }
                                                                            var v48 uint = v2
                                                                            v49 := GenS8{F0: 42, F1: (v4 ^ (v4 / (((^(v4))) | 1)))}
                                                                            v11 = -128
                                                                            var v50 uint8 = v43
                                                                            _ = v48
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v12, v12)
                                                                    }
                                                                case "hello":
                                                                    var v51 bool = false
                                                                    _ = v51
                                                                }
                                                                _ = v43
                                                            } else {
                                                                var v52 uint64 = 4294967295
                                                                _ = v52
                                                            }
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                if true {
                                                                    v11 = v11
                                                                    var v53 float32 = v10
                                                                    if false {
                                                                        fmt.Println(v4, v12)
                                                                    } else {
                                                                        var v54 int = 605
                                                                        if false {
                                                                            switch ("foo bar" + "") {
                                                                            case "\n\t":
                                                                                var v55 bool = (123 <= 42)
                                                                                _ = v55
                                                                            case "":
                                                                                var v56 int16 = 1000
                                                                                _ = v56
                                                                            case "hello":
                                                                                if (2147483647 == -416) {
                                                                                    var v57 bool = true
                                                                                    v2 -= (^(v2))
                                                                                    var v58 int32 = -2147483648
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v4, v2, v10)
                                                                                var v59 int32 = 1
                                                                                v60 := GenS9{F0: 2, F1: 2, F2: v4}
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                v11 -= -90
                                                                            }
                                                                            v12 -= (v12 - v12)
                                                                        }
                                                                        v61 := []int{v54, (fi21 % ((fi21) | 1))}
                                                                        var v62 uint16 = (v4 + (v4 << (((v2 / ((18446744073709551615) | 1))) & 63)))
                                                                        _ = v54
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    _ = v53
                                                                } else {
                                                                    if true {
                                                                        var v63 uint = uint(v4)
                                                                        _ = v63
                                                                    }
                                                                    v2 = v2
                                                                    delete(v1, fi21)
                                                                    v4 = 2
                                                                }
                                                            case "\n\t":
                                                                v2 = 742
                                                            case "hello":
                                                                var v64 int8 = int8(v12)
                                                                fmt.Println(v11, v2, v2)
                                                                v65 := GenS10{F0: 12, F1: (-(fi20)), F2: "\n\t"}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                        }
                                                        v66 := v1[(^(fi21))]
                                                        _ = v66
                                                    }
                                                default:
                                                    if false {
                                                        v67 := GenS11{F0: 74, F1: "hello", F2: 11}
                                                        if false {
                                                            v68 := []int{}
                                                            fmt.Println(fi20, fi21, fi21)
                                                            _ = v68
                                                        } else {
                                                            switch (((("a" + "a") + "") + "hello") + "hello") {
                                                            case "":
                                                                var v69 uint16 = v4
                                                                _ = v69
                                                            default:
                                                                v11 = 0
                                                            }
                                                            if (v11 < int8(v2)) {
                                                                var v70 int64 = 42
                                                                v12 -= (v12 & 731)
                                                                _ = v70
                                                            }
                                                            fmt.Println(fi21, fi20)
                                                        }
                                                        var v71 int = 42
                                                        _ = v67
                                                        _ = v71
                                                    }
                                                }
                                            case 32767:
                                                if true {
                                                    v72 := GenS10{F0: 42, F1: 128, F2: ("hello" + ((("hello" + ("hello" + "hello")) + ((("hello" + "hello") + "hello") + "\n\t")) + ("" + (("\n\t" + ("a" + "\n\t")) + "foo bar"))))}
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(fi21, v2, fi21)
                                                    default:
                                                        var v73 bool = false
                                                        _ = v73
                                                    }
                                                    v74 := GenS2{F0: -32768, F1: 42}
                                                    _ = v72
                                                    _ = v74
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        switch "" {
                                                        case "":
                                                            v75 := []int{(fi20 & -96), fi21, fi21}
                                                            _ = v75
                                                        case "foo bar":
                                                            if ((354 == -1) || (true && (0 != 2147483647))) {
                                                                v76 := GenS12{F0: 1000, F1: 2, F2: 127}
                                                                fmt.Println(v4, v11)
                                                                var v77 string = (("a" + "\n\t") + ("hello" + "\n\t"))
                                                                _ = v76
                                                                _ = v77
                                                            } else {
                                                                var v78 float64 = ((0.0 - ((1.5 / (0.0 * float64(v10))) * (((float64(1e10) * 3.14159) * 0.1) * 1.5))) + 0.0)
                                                                _ = v78
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v79 uint8 = 70
                                                                _ = v79
                                                            default:
                                                                v80 := GenS8{F0: (^(fi21)), F1: 1}
                                                                _ = v80
                                                            }
                                                            v81 := GenS1{F0: (v4 ^ (v4 - v4)), F1: 1}
                                                            _ = v81
                                                        default:
                                                            v4 = v4
                                                        }
                                                        if (0 < -1) {
                                                            var v82 uint16 = 32767
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v83 := []int{}
                                                                _ = v83
                                                            case "\n\t":
                                                                switch -285 {
                                                                case 128:
                                                                    var v84 uint16 = 545
                                                                    _ = v84
                                                                case -1:
                                                                    _ = len(v1)
                                                                }
                                                            }
                                                            v85 := GenS1{F0: (v82 + (v82 % ((v4) | 1))), F1: -2147483648}
                                                            if false {
                                                                v12 += (v12 % ((v12) | 1))
                                                                v11 = v11
                                                            }
                                                            _ = v82
                                                            _ = v85
                                                        }
                                                        v11 = 0
                                                        v2 = (v2 << ((v2) & 63))
                                                        var v86 float64 = 0.0
                                                        v2 = (v2 ^ v2)
                                                        var v87 uint = v2
                                                        var v88 uint = v87
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    case "a":
                                                        if true {
                                                            v89 := []int{(fi21 ^ (-(fi21))), fi20}
                                                            _ = v89
                                                        }
                                                    }
                                                    v11 = v11
                                                    var v90 int8 = v11
                                                    _ = v90
                                                }
                                            default:
                                                v91 := GenS13{F0: false, F1: (v11 % (((^(v11))) | 1)), F2: true}
                                                _ = v91
                                            }
                                            fmt.Println(v2)
                                            var v92 int32 = -1
                                            _ = v92
                                        } else {
                                            var v93 float64 = 1e10
                                            switch ("a" + ("" + "hello")) {
                                            case "":
                                                v94 := []int{fi21, (fi21 % ((fi21) | 1)), -2147483648, (fi21 / (((^(fi21))) | 1))}
                                                _ = v94
                                            case "hello":
                                                var v95 uint32 = 42
                                                _ = v95
                                            case "a":
                                                fmt.Println(v11, fi20, v93)
                                            }
                                            var v96 int64 = 1
                                            _ = v93
                                            _ = v96
                                        }
                                        _ = v33
                                    }
                                }
                            } else {
                                if false {
                                    v10 = v10
                                    var v97 int = (fi21 - (fi20 << ((v2) & 63)))
                                    switch (fi20 / ((9223372036854775807) | 1)) {
                                    case 361:
                                        var v98 int16 = 32767
                                        _ = v98
                                    case 812:
                                        var v99 int64 = 42
                                        _ = v99
                                    }
                                    _ = v97
                                } else {
                                    v11 = v11
                                    _ = len(v1)
                                    v11 += v11
                                }
                                fmt.Println(v10, v10, v12)
                                v11 = 1
                                v2 -= v2
                                var v100 uint64 = 807
                                _ = v100
                            }
                            v101 := GenS6{F0: 1000}
                            _ = fi21
                            _ = v101
                        }
                    case -796:
                        if (0 != fi20) {
                            var v102 int64 = 773
                            v103 := GenS14{F0: 1000}
                            _ = v102
                            _ = v103
                        }
                    case -9223372036854775808:
                        v3[(v12 * 911)] = ((v12 & v12) <= v12)
                    default:
                        _ = len(v1)
                    }
                    var v104 uint64 = 2
                    v105 := []int{fi20, (fi20 * -2147483648), fi20, (fi20 & (-(fi20)))}
                    _ = fi20
                    _ = v104
                    _ = v105
                }
                for fi106 := 0; fi106 < 10; fi106++ {
                    v11 = v11
                    _ = fi106
                }
                delete(v1, 32768)
            }
            v4 -= v4
        }
    default:
        v2 = v2
    }
    fmt.Println(v2)
    for fi107 := 0; fi107 < 4; fi107++ {
        if true {
            v10 += 0.1
            v108 := v1[int(v2)]
            switch fi107 {
            case -2147483648:
                var v109 uint16 = v4
                _ = v109
            case 42:
                v110 := GenS5{F0: -2147483648}
                _ = v110
            case 222:
                v108 = (v12 - 2147483648)
            default:
                v4 -= (v4 >> ((v2) & 63))
            }
            _ = v108
        } else {
            v11 -= -1
            v12 = (v12 ^ 4294967295)
            var v111 float64 = ((0.1 * ((1e10 + -2.25) * (1.5 + 1.5))) / 3.14159)
            _ = v111
        }
        _ = fi107
    }
    for fi112 := 0; fi112 < 1; fi112++ {
        fmt.Println(v2, v12)
        if (false && true) {
            v113 := GenS4{F0: 2, F1: 42, F2: 356}
            v114 := []int{fi112}
            fmt.Println(v11, v4)
            _ = v113
            _ = v114
        } else {
            if (!(v2 > 755)) {
                v3[v12] = false
                if ((42 == 42) && (v4 < v4)) {
                    var v115 int16 = 1000
                    v116 := GenS6{F0: v115}
                    _ = v115
                    _ = v116
                }
                v117 := GenS15{F0: v2, F1: -1, F2: v4}
                _ = v117
            } else {
                v118 := []int{(fi112 * (-(fi112))), (fi112 ^ -2147483648)}
                if (!(!(v2 <= (v2 >> (((v2 >> ((v2) & 63))) & 63))))) {
                    var v119 int64 = 9223372036854775807
                    v11 -= (v11 | -1)
                    switch fi112 {
                    case 176:
                        if len(v118) > 0 {
                            v118[0] = (fi112 ^ 4294967295)
                        }
                    default:
                        v120 := GenS16{F0: ("foo bar" + "\n\t"), F1: v4, F2: uint64(v2)}
                        _ = v120
                    }
                    _ = v119
                }
                var v121 int16 = 0
                _ = v118
                _ = v121
            }
            var v122 uint64 = 1
            _ = v122
        }
        _ = fi112
    }
    if (false && false) {
        v123 := []int{-905, 42, 1}
        _ = v123
    }
    v124 := GenS17{F0: ("foo bar" + "hello"), F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
