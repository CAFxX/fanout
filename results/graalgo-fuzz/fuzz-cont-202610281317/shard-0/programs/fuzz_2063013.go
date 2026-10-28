package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int32; F2 uint8 }
type GenS9 struct { F0 string; F1 int }
type GenS10 struct { F0 int; F1 int32; F2 uint32 }
type GenS11 struct { F0 uint16; F1 int16; F2 bool }
type GenS12 struct { F0 int; F1 uint; F2 int32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint; F1 uint16 }
type GenS15 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int16, p1 uint32, p2 int8) int16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 = append(v7, (fi8 / ((fi8) | 1)))
        _ = fi8
    }
    p0 = p0
    switch 4294967295 {
    case 128:
        if false {
            p2 = p2
        }
    case 4294967295:
        for fi9 := 0; fi9 < 11; fi9++ {
            switch "" {
            case "a":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "foo bar":
                var v10 int = 32767
                _ = v10
            default:
                var v11 string = "\n\t"
                _ = v11
            }
            fmt.Println(fi9)
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: -32768}
                p0 += 1000
                p2 = p2
                var v14 string = "hello"
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 699
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 int32 = -152
    v6 := make(map[uint32]string)
    v15 := GenS2{F0: 127, F1: ("\n\t" + ("a" + ((("foo bar" + "\n\t") + "foo bar") + "\n\t")))}
    var v16 int16 = int16(v1)
    v17 := GenS3{F0: 42}
    if ((2147483647 > 4294967295) && false) {
        v5 += v5
    }
    switch 127 {
    case 32768:
        var v18 float32 = float32((1e10 - (-2.25 + 1e10)))
        _ = v18
    case 9223372036854775807:
        v5 += -2147483648
    }
    if true {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 float64 = 3.14159
            if (128 != 0) {
                var v21 string = "\n\t"
                _ = v21
            } else {
                var v22 bool = ((((!(!true)) && true) || false) && false)
                if (!v22) {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        _ = len(v4)
                        if ((float32(float64(0.0)) / 0.1) == (-2.25 / 1.0)) {
                            var v24 int32 = -2147483648
                            switch "hello" {
                            case "a":
                                v1 = (v1 - (v1 | (v1 * v1)))
                            case "":
                                var v25 string = "a"
                                _ = v25
                            }
                            _ = v24
                        } else {
                            v20 = 0.0
                            var v26 uint8 = 1
                            var v27 bool = ((-(v16)) != v16)
                            _ = v26
                            _ = v27
                        }
                        v5 -= v5
                        _ = fi23
                    }
                }
                v16 = v16
                _ = v22
            }
            fmt.Println(v5, fi19)
            _ = fi19
            _ = v20
        }
    } else {
        var v28 uint = 790
        _ = v28
    }
    switch 953 {
    case 9223372036854775807:
        var v29 uint64 = 42
        _ = v29
    case 449:
        if (982 >= (^(v1))) {
            v30 := GenS2{F0: 52, F1: "hello"}
            var v31 int16 = v16
            v32 := GenS3{F0: 127}
            _ = v30
            _ = v31
            _ = v32
        } else {
            v16 -= (v16 / ((32767) | 1))
            switch ("\n\t" + ((("a" + ("\n\t" + ("a" + "foo bar"))) + "foo bar") + "")) {
            case "foo bar":
                switch "\n\t" {
                case "":
                    fmt.Println(v16)
                case "\n\t":
                    if (3.14159 != 1e10) {
                        v33 := v2[-211]
                        _ = v33
                    } else {
                        fmt.Println(v16, v16, v16)
                        var v34 string = "a"
                        _ = v34
                    }
                case "hello":
                    var v35 string = "\n\t"
                    _ = v35
                }
            case "":
                v36 := GenS4{F0: false, F1: v1, F2: 2}
                _ = v36
            case "hello":
                if false {
                    if true {
                        for fi37 := 0; fi37 < 19; fi37++ {
                            v38 := GenS5{F0: (213 != 42)}
                            _ = fi37
                            _ = v38
                        }
                        v39 := GenS1{F0: 180}
                        _ = v39
                    } else {
                        if (!(true || true)) {
                            if false {
                                switch "\n\t" {
                                case "\n\t":
                                    if true {
                                        for fi40 := 0; fi40 < 16; fi40++ {
                                            switch "" {
                                            case "hello":
                                                for fi41 := 0; fi41 < 11; fi41++ {
                                                    v5 += (v5 >> ((2) & 63))
                                                    var v42 int64 = -347
                                                    _ = fi41
                                                    _ = v42
                                                }
                                            case "a":
                                                if (((3.14159 + 0.1) + 1.5) == 3.14159) {
                                                    var v43 uint32 = 2147483648
                                                    var v44 bool = ((^(v43)) != 2147483647)
                                                    v1 = (^(v1))
                                                    _ = v43
                                                    _ = v44
                                                }
                                            case "foo bar":
                                                var v45 int = (fi40 / ((fi40) | 1))
                                                _ = v45
                                            default:
                                                var v46 float32 = 3.14159
                                                _ = v46
                                            }
                                            if (!(false && (32768 == 121))) {
                                                var v47 int = fi40
                                                for fi48 := 0; fi48 < 17; fi48++ {
                                                    v47 -= fi48
                                                    if false {
                                                        fmt.Println(fi48, v16, v47)
                                                        var v49 int16 = (v16 | (v16 << ((0) & 63)))
                                                        fmt.Println(v16, v1, v1)
                                                        v47 = (v47 - (-(v47)))
                                                        if false {
                                                            v50 := GenS3{F0: 127}
                                                            v5 = (v5 >> ((42) & 63))
                                                            _ = v50
                                                        } else {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                switch (^(fi48)) {
                                                                case 32767:
                                                                    if (v1 > v1) {
                                                                        v5 = (v5 / (((v5 ^ v5)) | 1))
                                                                        v5 = (v5 - v5)
                                                                        var v51 int = fi48
                                                                        _ = v51
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        if true {
                                                                            v5 -= (v5 / (((^(v5))) | 1))
                                                                            fmt.Println(v1)
                                                                        }
                                                                    }
                                                                case 0:
                                                                    fmt.Println(v1, fi40)
                                                                case 32768:
                                                                    var v52 int8 = 21
                                                                    _ = v52
                                                                }
                                                            case "foo bar":
                                                                var v53 int16 = v16
                                                                _ = v53
                                                            default:
                                                                var v54 int8 = 42
                                                                _ = v54
                                                            }
                                                            delete(v3, fi40)
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            var v55 uint8 = 255
                                                            _ = v55
                                                        case "a":
                                                            fmt.Println(v47, v1)
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            var v56 int16 = v49
                                                            _ = v56
                                                        case "\n\t":
                                                            v57 := []int{}
                                                            v58 := GenS5{F0: true}
                                                            v47 = -2147483648
                                                            if true {
                                                                if (0 <= 855) {
                                                                    if ((923 == 0) || (true || (0 > 255))) {
                                                                        v47 = (v47 / (((fi48 >> ((0) & 63))) | 1))
                                                                        switch (((("foo bar" + "") + "") + "") + "foo bar") {
                                                                        case "a":
                                                                            if true {
                                                                                var v59 float64 = (((0.0 * 1e10) + 0.0) / -2.25)
                                                                                fmt.Println(fi40)
                                                                                _ = v59
                                                                            } else {
                                                                                if true {
                                                                                    if (v16 == v16) {
                                                                                        switch 32767 {
                                                                                        case -655:
                                                                                            if true {
                                                                                                switch 32767 {
                                                                                                case 4294967295:
                                                                                                    v60 := GenS6{F0: -258, F1: "hello"}
                                                                                                    _ = v60
                                                                                                default:
                                                                                                    var v61 uint64 = (v1 ^ (v1 + (v1 << ((1) & 63))))
                                                                                                    _ = v61
                                                                                                }
                                                                                                if (32767 > 1) {
                                                                                                    v49 += 1
                                                                                                    if false {
                                                                                                        fmt.Println(v16, v16, v47)
                                                                                                        fmt.Println(v16)
                                                                                                        if (!(true && ((!true) && true))) {
                                                                                                            var v62 int32 = int32(v16)
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            v63 := GenS7{F0: (v5 * 2147483647)}
                                                                                                            var v64 uint32 = 2147483647
                                                                                                            v2[1] = true
                                                                                                            _ = v63
                                                                                                            _ = v64
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((v1 % ((42) | 1)) >= v1) {
                                                                                                            switch fi40 {
                                                                                                            case 719:
                                                                                                                if false {
                                                                                                                    delete(v3, fi40)
                                                                                                                }
                                                                                                            case 912:
                                                                                                                switch (("a" + "a") + "a") {
                                                                                                                case "":
                                                                                                                    v65 := GenS1{F0: v49}
                                                                                                                    _ = v65
                                                                                                                case "\n\t":
                                                                                                                    var v66 uint = 42
                                                                                                                    _ = v66
                                                                                                                default:
                                                                                                                    v67 := GenS8{F0: "foo bar", F1: 1, F2: 42}
                                                                                                                    _ = v67
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v47)
                                                                                                            delete(v6, 622)
                                                                                                        } else {
                                                                                                            switch (("" + ("a" + "a")) + "hello") {
                                                                                                            case "foo bar":
                                                                                                                var v68 float32 = 1e10
                                                                                                                _ = v68
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    var v69 uint8 = 128
                                                                                                                    v70 := GenS8{F0: "foo bar", F1: v5, F2: (^(v69))}
                                                                                                                    _ = v69
                                                                                                                    _ = v70
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        var v71 int = fi48
                                                                                                                        v72 := GenS9{F0: (("" + "\n\t") + "hello"), F1: v71}
                                                                                                                        fmt.Println(v16)
                                                                                                                        _ = v71
                                                                                                                        _ = v72
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            var v73 int32 = -2147483648
                                                                                                                            v16 = v49
                                                                                                                            delete(v3, 32768)
                                                                                                                            _ = v73
                                                                                                                        } else {
                                                                                                                            var v74 float64 = -2.25
                                                                                                                            v75 := GenS10{F0: -610, F1: v5, F2: 65535}
                                                                                                                            v76 := GenS10{F0: (^(v47)), F1: (-(v5)), F2: 254}
                                                                                                                            _ = v74
                                                                                                                            _ = v75
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        var v77 int32 = v5
                                                                                                                        _ = v77
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v47, v49)
                                                                                                    }
                                                                                                    fmt.Println(v49, v5)
                                                                                                    v16 += (v16 << ((743) & 63))
                                                                                                    _ = len(v57)
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v47)
                                                                                            }
                                                                                        default:
                                                                                            switch (v47 / ((-521) | 1)) {
                                                                                            case 32767:
                                                                                                switch 32768 {
                                                                                                case -186:
                                                                                                    if (9223372036854775807 <= 9223372036854775807) {
                                                                                                        if (3.14159 > (3.14159 * (-2.25 / 1.5))) {
                                                                                                            if len(v57) >= 2 {
                                                                                                                v57 = v57[1:len(v57)]
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        _ = len(v6)
                                                                                                        v5 = v5
                                                                                                        var v78 int = v47
                                                                                                        _ = v78
                                                                                                    }
                                                                                                case 128:
                                                                                                    v79 := GenS11{F0: 710, F1: (v49 & 42), F2: (false || true)}
                                                                                                    _ = v79
                                                                                                }
                                                                                            default:
                                                                                                var v80 int64 = -419
                                                                                                _ = v80
                                                                                            }
                                                                                        }
                                                                                        v81 := GenS12{F0: -2147483648, F1: 962, F2: v5}
                                                                                        _ = v81
                                                                                    } else {
                                                                                        v1 -= v1
                                                                                        _ = len(v57)
                                                                                    }
                                                                                    v82 := GenS4{F0: ((v16 ^ (-(v49))) == (v49 + (-(v49)))), F1: (v1 + 4294967295), F2: 2147483647}
                                                                                    _ = v82
                                                                                } else {
                                                                                    var v83 uint64 = (v1 % ((541) | 1))
                                                                                    v5 = (-(v5))
                                                                                    v6[4294967295] = ""
                                                                                    fmt.Println(v49)
                                                                                    _ = v83
                                                                                }
                                                                                v5 = (v5 >> ((1) & 63))
                                                                                v49 += (v16 ^ (-(v49)))
                                                                                v84 := GenS13{F0: 255}
                                                                                _ = v84
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v85 int64 = 9223372036854775807
                                                                    v16 -= (v49 + int16(fi48))
                                                                    _ = v85
                                                                }
                                                                v6[499] = "hello"
                                                            }
                                                            v86 := GenS7{F0: v5}
                                                            fmt.Println(fi40, v5)
                                                            _ = v57
                                                            _ = v58
                                                            _ = v86
                                                        default:
                                                            v47 = v47
                                                        }
                                                        v47 = fi48
                                                        fmt.Println(v16)
                                                        _ = v49
                                                    } else {
                                                        var v87 uint = 868
                                                        _ = v87
                                                    }
                                                    _ = fi48
                                                }
                                                if true {
                                                    _ = len(v6)
                                                    var v88 int32 = 545
                                                    v89 := GenS2{F0: 127, F1: "\n\t"}
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    v90 := []int{2, v47, (v47 >> ((399) & 63)), 2147483647}
                                                    v5 = (v5 >> ((42) & 63))
                                                    fmt.Println(v16, v47)
                                                    _ = v90
                                                }
                                                _ = v47
                                            }
                                            _ = fi40
                                        }
                                        v1 -= v1
                                        fmt.Println(v16)
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v1 -= 2
                                        case "\n\t":
                                            var v91 string = "foo bar"
                                            _ = v91
                                        case "":
                                            var v92 int = 32767
                                            _ = v92
                                        default:
                                            v93 := GenS10{F0: 1, F1: -2147483648, F2: 864}
                                            _ = v93
                                        }
                                        for fi94 := 0; fi94 < 2; fi94++ {
                                            var v95 int8 = 127
                                            _ = fi94
                                            _ = v95
                                        }
                                        var v96 bool = false
                                        _ = v96
                                    }
                                case "hello":
                                    v97 := GenS5{F0: true}
                                    _ = v97
                                }
                                fmt.Println(v5)
                                delete(v6, 2)
                            } else {
                                for fi98 := 0; fi98 < 20; fi98++ {
                                    if true {
                                        v99 := []int{(^(fi98)), int(v16), -9223372036854775808}
                                        for fi100 := 0; fi100 < 5; fi100++ {
                                            var v101 int32 = (v5 * (v5 - (v5 | (-(v5)))))
                                            _ = fi100
                                            _ = v101
                                        }
                                        _ = v99
                                    }
                                    var v102 int32 = (v5 | v5)
                                    v1 -= 883
                                    _ = fi98
                                    _ = v102
                                }
                                if false {
                                    var v103 uint16 = 2
                                    var v104 uint32 = 865
                                    for fi105 := 0; fi105 < 16; fi105++ {
                                        if false {
                                            fmt.Println(v104)
                                            v1 -= v1
                                        }
                                        switch "foo bar" {
                                        case "":
                                            switch ("foo bar" + "") {
                                            case "foo bar":
                                                v5 = v5
                                            }
                                        case "foo bar":
                                            if false {
                                                v104 = uint32(v1)
                                            }
                                        case "a":
                                            var v106 int8 = 100
                                            _ = v106
                                        default:
                                            for fi107 := 0; fi107 < 10; fi107++ {
                                                _ = len(v2)
                                                _ = fi107
                                            }
                                        }
                                        _ = fi105
                                    }
                                    _ = v103
                                    _ = v104
                                } else {
                                    var v108 uint64 = v1
                                    var v109 uint32 = 0
                                    _ = v108
                                    _ = v109
                                }
                            }
                        }
                        v1 += 542
                    }
                    v5 = (v5 >> ((1) & 63))
                    v5 -= (v5 ^ -2147483648)
                } else {
                    switch -912 {
                    case 635:
                        fmt.Println(v1, v16)
                    case 32767:
                        var v110 int32 = (v5 ^ (v5 ^ v5))
                        _ = v110
                    case 838:
                        fmt.Println(v1, v5)
                    }
                    delete(v4, 840)
                    _ = len(v3)
                }
            }
        }
    }
    for fi111 := 0; fi111 < 12; fi111++ {
        for fi112 := 0; fi112 < 11; fi112++ {
            if true {
                v1 += v1
            }
            fmt.Println(v1)
            _ = fi112
        }
        _ = fi111
    }
    for fi113 := 0; fi113 < 6; fi113++ {
        if ((float64((-2.25 + 0.0)) / 1.0) <= (3.14159 + 0.1)) {
            v16 += v16
            v16 -= (-(v16))
            if false {
                fmt.Println(v1)
                v1 += (v1 >> ((2) & 63))
                var v114 float64 = 0.1
                _ = v114
            }
        }
        if true {
            if ((v5 ^ (v5 & (v5 >> ((341) & 63)))) <= 1) {
                fmt.Println(v5, fi113)
                var v115 int = fi113
                v116 := GenS1{F0: (^(v16))}
                _ = v115
                _ = v116
            } else {
                switch 0 {
                case 32767:
                    v117 := GenS11{F0: 42, F1: v16, F2: true}
                    _ = v117
                }
                var v118 int32 = v5
                _ = v118
            }
            for fi119 := 0; fi119 < 4; fi119++ {
                v5 = v5
                _ = fi119
            }
        } else {
            fmt.Println(v16, fi113)
        }
        _ = fi113
    }
    if (!(!(32768 <= -2147483648))) {
        v120 := GenS14{F0: uint(v1), F1: 202}
        var v121 uint64 = v1
        var v122 int8 = 127
        _ = v120
        _ = v121
        _ = v122
    } else {
        v123 := GenS9{F0: "hello", F1: 721}
        v124 := GenS15{F0: 65535, F1: 0}
        var v125 uint64 = (v1 % (((v1 + (v1 % ((v1) | 1)))) | 1))
        _ = v123
        _ = v124
        _ = v125
    }
    if (true && ((v1 == v1) && (false || true))) {
        var v126 string = "\n\t"
        v127 := GenS10{F0: 127, F1: v5, F2: 65535}
        _ = v126
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
