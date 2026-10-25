package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int }
func genFn0(p0 uint, p1 int8, p2 uint32) int64 {
    var v5 int64 = 9223372036854775807
    var v6 float32 = -2.25
    v5 += v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 65535
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 2
    fmt.Println(v4, v2, v2)
    switch -1 {
    case 128:
        fmt.Println(v1)
    }
    if false {
        if true {
            v7 := GenS1{F0: v2, F1: 42}
            v8 := v3[2]
            switch -9223372036854775808 {
            case 128:
                var v9 uint = 1
                _ = v9
            default:
                if false {
                    var v10 float64 = float64(1e10)
                    switch int(v1) {
                    case 4294967295:
                        switch "hello" {
                        case "hello":
                            fmt.Println(v10)
                        }
                    case 916:
                        fmt.Println(v8, v2)
                    case 32767:
                        if (true && false) {
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    v10 = v10
                                    if (42 > 4294967295) {
                                        var v11 uint = 27
                                        v3[v8] = (v8 | v8)
                                        if ((^(v8)) != v8) {
                                            v12 := GenS1{F0: 430, F1: v11}
                                            v11 -= v11
                                            fmt.Println(v4)
                                            _ = v12
                                        }
                                        _ = v11
                                    }
                                } else {
                                    for fi13 := 0; fi13 < 8; fi13++ {
                                        if (1 >= v1) {
                                            fmt.Println(v8, v10)
                                        } else {
                                            v14 := []int{(^(fi13)), 1, (^(fi13)), 9223372036854775807}
                                            v15 := GenS1{F0: (v2 - 2), F1: 152}
                                            for fi16 := 0; fi16 < 17; fi16++ {
                                                v8 -= (^(v8))
                                                v17 := GenS1{F0: v4, F1: 2}
                                                v18 := v3[v8]
                                                _ = fi16
                                                _ = v17
                                                _ = v18
                                            }
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = fi13
                                    }
                                }
                            case "hello":
                                for fi19 := 0; fi19 < 13; fi19++ {
                                    v2 = (v4 ^ 434)
                                    var v20 float64 = 1.5
                                    v21 := []int{}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            default:
                                v22 := GenS2{F0: uint64(v2)}
                                _ = v22
                            }
                        }
                    }
                    var v23 int16 = -714
                    _ = v10
                    _ = v23
                }
            }
            _ = v7
            _ = v8
        }
    } else {
        var v24 int = 32767
        _ = v24
    }
    fmt.Println(v1)
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := GenS3{F0: genFn0(2, -128, 0)}
        _ = fi25
        _ = v26
    }
    switch 32768 {
    case 9223372036854775807:
        v4 -= 42
    case 2147483647:
        switch "" {
        case "\n\t":
            for fi27 := 0; fi27 < 2; fi27++ {
                switch fi27 {
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        var v28 int8 = -57
                        _ = v28
                    case "\n\t":
                        var v29 int16 = 1000
                        _ = v29
                    default:
                        var v30 float32 = 1.5
                        _ = v30
                    }
                case -529:
                    var v31 float64 = (1e10 + 1.5)
                    _ = v31
                }
                _ = fi27
            }
        }
    case 128:
        if (-1 > -211) {
            for fi32 := 0; fi32 < 6; fi32++ {
                _ = len(v3)
                var v33 uint16 = (^(v2))
                if (!(255 > 65535)) {
                    if (genFn0(255, -128, 1) >= genFn0(42, 127, uint32(v33))) {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            v1 -= 42
                            fmt.Println(v2)
                            var v35 float64 = 3.14159
                            if (394 <= 2) {
                                v36 := GenS4{F0: (0 == 97), F1: -99, F2: 625}
                                _ = v36
                            }
                            _ = fi34
                            _ = v35
                        }
                    }
                } else {
                    v1 += 2147483647
                    switch ("" + (("a" + "") + (("" + "\n\t") + "a"))) {
                    case "a":
                        v37 := GenS1{F0: (^(v2)), F1: 65535}
                        _ = v37
                    case "foo bar":
                        fmt.Println(v33)
                    case "hello":
                        for fi38 := 0; fi38 < 11; fi38++ {
                            var v39 uint64 = 1
                            var v40 int32 = v1
                            v41 := GenS5{F0: (fi38 ^ (fi32 ^ fi32))}
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    default:
                        fmt.Println(v4)
                    }
                }
                _ = fi32
                _ = v33
            }
            if true {
                for fi42 := 0; fi42 < 10; fi42++ {
                    if true {
                        v43 := []int{fi42, (fi42 - fi42), 2147483647}
                        for fi44 := 0; fi44 < 17; fi44++ {
                            var v45 float32 = 0.1
                            var v46 int16 = int16(v1)
                            var v47 uint8 = 255
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 19; fi48++ {
                            v4 = (v2 & 863)
                            _ = fi48
                        }
                        _ = v43
                    }
                    if (true && true) {
                        v4 = v2
                        v49 := GenS1{F0: (v2 % ((2) | 1)), F1: 65535}
                        _ = v49
                    }
                    _ = fi42
                }
                for fi50 := 0; fi50 < 1; fi50++ {
                    switch "" {
                    case "":
                        v51 := GenS4{F0: true, F1: 42, F2: 2147483647}
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 10; fi52++ {
                            var v53 int = (fi50 & (fi50 >> ((1) & 63)))
                            _ = fi52
                            _ = v53
                        }
                    }
                    v54 := GenS6{F0: v2}
                    if false {
                        v55 := []int{fi50}
                        _ = v55
                    } else {
                        if ((!false) && false) {
                            var v56 uint16 = v2
                            for fi57 := 0; fi57 < 5; fi57++ {
                                v58 := GenS4{F0: (-308 < -951), F1: -1, F2: 42}
                                fmt.Println(v1, fi50, fi50)
                                if (0 >= 42) {
                                    var v59 float32 = (3.14159 / 1.5)
                                    v60 := []int{(fi57 / ((fi57) | 1)), (fi57 / ((fi57) | 1))}
                                    v56 += 1
                                    switch fi57 {
                                    case 0:
                                        if (-2147483648 == (v1 - v1)) {
                                            v59 = float32((0.0 - -2.25))
                                        }
                                    case 1:
                                        delete(v3, 547)
                                    case 32767:
                                        _ = len(v3)
                                    }
                                    var v61 int = (^(fi57))
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                v4 = v2
                                _ = fi57
                                _ = v58
                            }
                            v62 := []int{(fi50 >> ((2) & 63))}
                            _ = v56
                            _ = v62
                        } else {
                            for fi63 := 0; fi63 < 14; fi63++ {
                                var v64 uint32 = 2
                                var v65 int16 = -579
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            var v66 int8 = 127
                            v67 := v3[4294967295]
                            _ = v66
                            _ = v67
                        }
                        v2 -= (v2 | 2)
                    }
                    _ = fi50
                    _ = v54
                }
                if (-644 != 1) {
                    fmt.Println(v4, v1)
                    var v68 int = 9223372036854775807
                    _ = v68
                } else {
                    fmt.Println(v2, v4, v4)
                }
            }
        } else {
            v69 := GenS7{F0: 819, F1: 0}
            v4 -= v2
            _ = v69
        }
    }
    if (1 == 0) {
        fmt.Println(v4, v4)
        switch "hello" {
        case "a":
            v2 -= 122
        case "":
            for fi70 := 0; fi70 < 9; fi70++ {
                var v71 uint16 = 937
                v72 := []int{}
                if (!(-1 <= 22)) {
                    v73 := GenS8{F0: 100, F1: ("" + ("" + "\n\t"))}
                    for fi74 := 0; fi74 < 8; fi74++ {
                        if ((true && false) && true) {
                            v2 = v4
                            v2 += v2
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v71)
                            case "":
                                v75 := GenS9{F0: -72}
                                _ = v75
                            }
                        }
                        _ = fi74
                    }
                    _ = v73
                } else {
                    for fi76 := 0; fi76 < 17; fi76++ {
                        v1 -= 0
                        _ = fi76
                    }
                    v2 = (v4 / ((v71) | 1))
                    if true {
                        v71 += v4
                    } else {
                        if true {
                            v2 = v71
                            v77 := GenS10{F0: 113, F1: 42}
                            _ = v77
                        } else {
                            if ((0.0 <= float64(1e10)) || (("" + "") > "a")) {
                                var v78 uint8 = 6
                                _ = v78
                            }
                            if ((((((-2.25 - ((float64(((((((((1e10 + 3.14159) * (float32(1.5) * (float32(-2.25) / 3.14159))) + -2.25) * 0.0) * 1.5) - (1e10 * 1e10)) - float32((((1e10 / 1.0) - ((0.0 - ((1e10 * ((0.1 * -2.25) * 3.14159)) - -2.25)) / -2.25)) / float64(1e10)))) - 3.14159)) * (3.14159 - 1.5)) - 1e10)) - 0.1) - (1e10 - (float64((((3.14159 - -2.25) - float32(0.0)) - (0.0 * 1.5))) - 1e10))) / (0.1 + (0.1 + (0.1 - float64((1e10 * 1.5)))))) > ((float64(0.0) * float64((-2.25 / ((0.0 * ((0.1 / 1.0) / 3.14159)) - (1.5 - -2.25))))) * -2.25)) && (!true)) {
                                switch "" {
                                case "hello":
                                    for fi79 := 0; fi79 < 3; fi79++ {
                                        var v80 bool = true
                                        var v81 string = "foo bar"
                                        _ = fi79
                                        _ = v80
                                        _ = v81
                                    }
                                case "a":
                                    fmt.Println(v4, fi70)
                                case "":
                                    var v82 uint8 = 0
                                    _ = v82
                                }
                                for fi83 := 0; fi83 < 8; fi83++ {
                                    var v84 float64 = (1e10 / (3.14159 + 3.14159))
                                    _ = fi83
                                    _ = v84
                                }
                                if (42 > -5) {
                                    v85 := GenS11{F0: "\n\t"}
                                    fmt.Println(v4)
                                    _ = v85
                                } else {
                                    _ = len(v72)
                                    switch "\n\t" {
                                    case "a":
                                        if (0 > (v2 % (((v4 % ((65535) | 1))) | 1))) {
                                            delete(v3, 216)
                                        } else {
                                            if len(v72) >= 2 {
                                                v72 = v72[1:len(v72)]
                                            }
                                            v4 -= (^(v71))
                                        }
                                    case "foo bar":
                                        if (!false) {
                                            fmt.Println(v1)
                                        }
                                    }
                                }
                            } else {
                                v4 -= (v2 ^ 992)
                                switch fi70 {
                                case 609:
                                    for fi86 := 0; fi86 < 19; fi86++ {
                                        v87 := GenS6{F0: (^(v4))}
                                        _ = fi86
                                        _ = v87
                                    }
                                case 42:
                                    if false {
                                        v4 += v2
                                        var v88 float64 = (3.14159 + (1.5 + 3.14159))
                                        _ = v88
                                    } else {
                                        v89 := GenS12{F0: (^(v2))}
                                        if true {
                                            v90 := GenS7{F0: (^(v2)), F1: -1}
                                            var v91 int32 = v1
                                            _ = v90
                                            _ = v91
                                        }
                                        var v92 int = (fi70 >> ((42) & 63))
                                        _ = v89
                                        _ = v92
                                    }
                                case 9223372036854775807:
                                    v1 -= 42
                                default:
                                    v4 = v4
                                }
                            }
                        }
                        switch "hello" {
                        case "foo bar":
                            v2 = (v2 + 42)
                        }
                        if (255 != 43) {
                            if (true || (0 != 438)) {
                                v71 = v2
                                var v93 uint32 = 645
                                _ = v93
                            } else {
                                v71 = v4
                                if len(v72) >= 2 {
                                    v72 = v72[1:len(v72)]
                                }
                                fmt.Println(v2, v2)
                            }
                            v4 += (v4 ^ v4)
                            v2 = (^(v4))
                        }
                    }
                }
                _ = fi70
                _ = v71
                _ = v72
            }
        }
    } else {
        if false {
            v2 = (v4 + (v2 ^ (v2 * 2)))
        }
        v94 := GenS6{F0: (v4 * 32767)}
        _ = v94
    }
    v4 = 42
    fmt.Println(v4)
    for fi95 := 0; fi95 < 1; fi95++ {
        var v96 int8 = 127
        var v97 int16 = 0
        v98 := GenS13{F0: fi95}
        _ = fi95
        _ = v96
        _ = v97
        _ = v98
    }
    var v99 string = ""
    var v100 uint = 700
    switch ((v99 + ("" + ("\n\t" + v99))) + v99) {
    case "foo bar":
        for fi101 := 0; fi101 < 3; fi101++ {
            var v102 uint32 = 4294967295
            v103 := []int{fi101, (fi101 - fi101), fi101}
            v104 := GenS12{F0: v2}
            _ = fi101
            _ = v102
            _ = v103
            _ = v104
        }
    case "hello":
        v105 := GenS11{F0: "a"}
        _ = v105
    }
    v106 := GenS12{F0: v2}
    if false {
        for fi107 := 0; fi107 < 17; fi107++ {
            if false {
                v108 := []int{}
                switch fi107 {
                case 42:
                    fmt.Println(v2, v2)
                case -88:
                    if true {
                        fmt.Println(v2, v100, fi107)
                        v4 += (v4 % (((v2 >> ((v100) & 63))) | 1))
                        fmt.Println(v1, v4)
                    }
                case -1:
                    var v109 uint32 = uint32(v2)
                    _ = v109
                }
                _ = v108
            } else {
                var v110 int = (fi107 >> (((v100 >> (((^(v100))) & 63))) & 63))
                switch 52 {
                case 9223372036854775807:
                    var v111 string = (("" + v99) + (v99 + "foo bar"))
                    _ = v111
                case -1:
                    fmt.Println(v110, v100)
                case 620:
                    v100 -= 1
                default:
                    switch ((v99 + v99) + "a") {
                    case "hello":
                        fmt.Println(v100, v2, v2)
                    case "":
                        if (797 >= 2) {
                            v112 := GenS9{F0: 100}
                            var v113 float64 = 1.5
                            _ = v112
                            _ = v113
                        } else {
                            var v114 bool = (4294967295 < 18446744073709551615)
                            _ = v114
                        }
                    default:
                        switch (^(fi107)) {
                        case 307:
                            fmt.Println(v2, v1, v99)
                        default:
                            v99 = "a"
                        }
                    }
                }
                _ = len(v3)
                _ = v110
            }
            fmt.Println(fi107, v100, v2)
            v2 -= v4
            _ = fi107
        }
        var v115 int = -134
        for fi116 := 0; fi116 < 11; fi116++ {
            var v117 uint64 = 181
            _ = fi116
            _ = v117
        }
        _ = v115
    } else {
        _ = len(v3)
        v2 = 691
    }
    v3[65535] = 2147483647
    fmt.Println(v100, v4)
    switch "\n\t" {
    case "\n\t":
        v118 := GenS8{F0: int8(v1), F1: "foo bar"}
        _ = v118
    default:
        var v119 float64 = -2.25
        _ = v119
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v99
}
