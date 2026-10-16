package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int; F2 uint }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 int64; F1 uint; F2 int64 }
type GenS8 struct { F0 uint32; F1 uint }

func main() {
    var v1 string = "a"
    var v2 int32 = 274
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 int32 = v2
    var v6 int32 = v5
    v5 += 2147483647
    v7 := GenS1{F0: -128}
    fmt.Println(v5)
    var v8 uint16 = 130
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: 866, F1: fi9, F2: 65535}
        if (4294967295 > 589) {
            if false {
                var v11 float32 = (1.5 / 1.0)
                var v12 int64 = -1
                _ = v11
                _ = v12
            } else {
                var v13 int = (fi9 >> ((65535) & 63))
                if (2147483648 >= 0) {
                    switch v1 {
                    case "":
                        var v14 string = (((v1 + v1) + ("" + ("" + (v1 + v1)))) + v1)
                        _ = v14
                    case "a":
                        v8 = v8
                    default:
                        if true {
                            switch v1 {
                            case "a":
                                v15 := GenS2{F0: -1, F1: 32767, F2: 0}
                                _ = v15
                            default:
                                v16 := []int{fi9, (^(v13)), (v13 - fi9), v13}
                                _ = v16
                            }
                        } else {
                            v6 = (v5 | (v5 | (v6 - v5)))
                            v17 := GenS1{F0: -1}
                            v1 = v1
                            _ = v17
                        }
                    }
                } else {
                    if false {
                        if false {
                            if true {
                                if (((-2.25 == 1.5) || (true && true)) && false) {
                                    v18 := v3[42]
                                    v19 := GenS1{F0: 127}
                                    _ = v18
                                    _ = v19
                                }
                                switch (v1 + v1) {
                                case "foo bar":
                                    v20 := GenS2{F0: 425, F1: v13, F2: 4294967295}
                                    _ = v20
                                case "a":
                                    v21 := []int{v13, (v13 & (fi9 % (((v13 << ((18446744073709551615) & 63))) | 1))), (v13 / ((v13) | 1))}
                                    _ = v21
                                case "\n\t":
                                    var v22 float32 = float32(3.14159)
                                    _ = v22
                                default:
                                    delete(v4, (^(v13)))
                                }
                                v6 += (v6 / ((v6) | 1))
                            } else {
                                v5 = v5
                                v13 -= -595
                                fmt.Println(v2)
                            }
                            var v23 int32 = v6
                            var v24 uint16 = (^(v8))
                            _ = v23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 float64 = 0.0
                            _ = fi25
                            _ = v26
                        }
                    }
                    v27 := []int{(^(fi9)), (fi9 - (fi9 << ((42) & 63))), (v13 << ((352) & 63)), -439}
                    if len(v27) > 0 {
                        v27[0] = 777
                    }
                    _ = v27
                }
                var v28 uint32 = 747
                _ = v13
                _ = v28
            }
            fmt.Println(v8)
        }
        _ = fi9
        _ = v10
    }
    v29 := GenS1{F0: -1}
    for fi30 := 0; fi30 < 6; fi30++ {
        v2 += (v5 << ((18446744073709551615) & 63))
        var v31 float32 = 0.0
        switch (v1 + "\n\t") {
        case "hello":
            v8 -= 32767
        default:
            if true {
                var v32 int64 = -335
                _ = v32
            } else {
                if (1 > 65535) {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        v34 := []int{fi33, 128, fi30}
                        _ = fi33
                        _ = v34
                    }
                    v6 += (v2 - (v5 - (v2 / ((1) | 1))))
                } else {
                    _ = len(v4)
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 uint16 = 255
                    fmt.Println(fi35)
                    _ = fi35
                    _ = v36
                }
                v37 := []int{(fi30 ^ (fi30 - (^(fi30)))), fi30, 2147483647}
                _ = v37
            }
        }
        _ = fi30
        _ = v31
    }
    v6 = v6
    if (1.5 == 1.5) {
        if false {
            v38 := GenS1{F0: 0}
            for fi39 := 0; fi39 < 13; fi39++ {
                if (1 != 42) {
                    for fi40 := 0; fi40 < 7; fi40++ {
                        v41 := []int{}
                        v6 = (v6 << ((42) & 63))
                        var v42 uint = 450
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    if (!false) {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    var v44 int = fi43
                                    v8 -= v8
                                    fmt.Println(fi39, v1, v6)
                                    _ = fi43
                                    _ = v44
                                }
                                var v45 int32 = (v6 / ((v6) | 1))
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    var v47 uint16 = v8
                                    _ = fi46
                                    _ = v47
                                }
                                _ = v45
                            } else {
                                var v48 uint16 = (^(v8))
                                v49 := GenS1{F0: -128}
                                _ = v48
                                _ = v49
                            }
                        default:
                            v1 = (v1 + ((v1 + v1) + v1))
                        }
                    }
                    v2 += v2
                } else {
                    v50 := GenS1{F0: 127}
                    v2 = v6
                    v5 -= (v5 - 218)
                    _ = v50
                }
                v6 += 1
                _ = fi39
            }
            _ = v38
        } else {
            if false {
                switch 42 {
                case 424:
                    delete(v3, 2)
                case 1:
                    v8 += v8
                }
                switch 9223372036854775807 {
                case 32768:
                    v6 += v5
                default:
                    var v51 int64 = -9223372036854775808
                    _ = v51
                }
            }
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 uint8 = 127
                fmt.Println(v2, v1, v8)
                _ = fi52
                _ = v53
            }
        }
        switch "foo bar" {
        case "hello":
            switch "" {
            case "foo bar":
                fmt.Println(v5)
            case "a":
                v2 += (v6 / ((int32(v8)) | 1))
            default:
                if true {
                    for fi54 := 0; fi54 < 17; fi54++ {
                        v55 := GenS1{F0: 0}
                        _ = fi54
                        _ = v55
                    }
                    var v56 int8 = 42
                    _ = v56
                } else {
                    if false {
                        switch v1 {
                        case "\n\t":
                            v8 = v8
                        }
                        for fi57 := 0; fi57 < 2; fi57++ {
                            if ((127 != 42) || (!false)) {
                                for fi58 := 0; fi58 < 17; fi58++ {
                                    v2 = v6
                                    v59 := GenS1{F0: 56}
                                    if true {
                                        v8 = 672
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                                if (false && true) {
                                    v5 = (-(v6))
                                    if false {
                                        for fi60 := 0; fi60 < 9; fi60++ {
                                            v4[fi60] = 724
                                            _ = fi60
                                        }
                                        v61 := GenS3{F0: (v1 + v1), F1: false, F2: int(v6)}
                                        if false {
                                            v62 := []int{fi57, fi57}
                                            _ = v62
                                        } else {
                                            var v63 uint = 4294967295
                                            for fi64 := 0; fi64 < 10; fi64++ {
                                                fmt.Println(v1, fi64)
                                                switch 127 {
                                                case 32767:
                                                    v1 = "a"
                                                }
                                                fmt.Println(v63, v2, fi64)
                                                _ = fi64
                                            }
                                            for fi65 := 0; fi65 < 11; fi65++ {
                                                var v66 float32 = 0.1
                                                fmt.Println(v8, v63, fi65)
                                                _ = len(v4)
                                                var v67 float32 = v66
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                            _ = v63
                                        }
                                        _ = v61
                                    }
                                    v5 -= v2
                                }
                            } else {
                                fmt.Println(v2)
                                for fi68 := 0; fi68 < 3; fi68++ {
                                    if (984 != -1) {
                                        switch ("hello" + ("" + v1)) {
                                        case "\n\t":
                                            v69 := GenS1{F0: 85}
                                            _ = v69
                                        default:
                                            if (!(48 > 45)) {
                                                v70 := []int{2147483647, fi68}
                                                _ = v70
                                            } else {
                                                v71 := []int{416}
                                                _ = v71
                                            }
                                        }
                                        v5 -= v2
                                    }
                                    v8 -= v8
                                    _ = fi68
                                }
                                for fi72 := 0; fi72 < 8; fi72++ {
                                    var v73 int8 = 0
                                    _ = fi72
                                    _ = v73
                                }
                            }
                            switch v1 {
                            case "hello":
                                if ((!false) || (771 != 243)) {
                                    switch ((v1 + v1) + ("a" + "")) {
                                    case "hello":
                                        v74 := GenS4{F0: v1}
                                        _ = v74
                                    case "":
                                        fmt.Println(v5)
                                    default:
                                        for fi75 := 0; fi75 < 14; fi75++ {
                                            var v76 string = "hello"
                                            v77 := GenS1{F0: 127}
                                            if (1 == 9223372036854775807) {
                                                v76 = "\n\t"
                                                v78 := GenS5{F0: (fi57 << ((2) & 63))}
                                                _ = v78
                                            } else {
                                                fmt.Println(v5, v5)
                                                fmt.Println(fi75, fi75, fi75)
                                                v6 = 194
                                            }
                                            _ = fi75
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                } else {
                                    switch (fi57 + (fi57 / ((fi57) | 1))) {
                                    case -1:
                                        for fi79 := 0; fi79 < 11; fi79++ {
                                            var v80 int = (fi57 - (fi79 >> ((984) & 63)))
                                            var v81 int = v80
                                            var v82 string = v1
                                            v83 := []int{(v81 << ((2) & 63)), v80}
                                            switch -1 {
                                            case -969:
                                                delete(v3, 2147483647)
                                            case 42:
                                                v6 -= -2147483648
                                            case -9223372036854775808:
                                                var v84 float32 = ((-2.25 + (-2.25 + 1e10)) * 1e10)
                                                _ = v84
                                            }
                                            _ = fi79
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                            _ = v83
                                        }
                                    case -2147483648:
                                        v6 -= (v2 + v2)
                                    default:
                                        for fi85 := 0; fi85 < 2; fi85++ {
                                            var v86 uint64 = 911
                                            fmt.Println(v1)
                                            _ = fi85
                                            _ = v86
                                        }
                                    }
                                    v5 -= -2147483648
                                    v2 -= -2147483648
                                }
                            case "a":
                                v87 := GenS5{F0: 601}
                                _ = v87
                            case "foo bar":
                                fmt.Println(v5)
                            default:
                                for fi88 := 0; fi88 < 20; fi88++ {
                                    var v89 int64 = -1
                                    v90 := []int{2147483647, fi88, -9223372036854775808, fi88}
                                    _ = fi88
                                    _ = v89
                                    _ = v90
                                }
                            }
                            _ = fi57
                        }
                        switch 4294967295 {
                        case 297:
                            v91 := GenS6{F0: 128, F1: 0, F2: (100 > 127)}
                            _ = v91
                        default:
                            fmt.Println(v8, v1)
                        }
                    } else {
                        v2 -= v5
                    }
                    var v92 int16 = 1000
                    var v93 uint16 = v8
                    _ = v92
                    _ = v93
                }
            }
        case "a":
            var v94 float64 = 1e10
            _ = v94
        case "":
            for fi95 := 0; fi95 < 15; fi95++ {
                fmt.Println(fi95)
                fmt.Println(v1, v1)
                var v96 uint = 65535
                _ = fi95
                _ = v96
            }
        default:
            for fi97 := 0; fi97 < 7; fi97++ {
                v8 -= 483
                _ = fi97
            }
        }
    }
    var v98 int8 = 42
    switch 32768 {
    case -920:
        switch 9223372036854775807 {
        case -9223372036854775808:
            switch 915 {
            case 4294967295:
                v8 = 255
            case 0:
                var v99 uint8 = 255
                _ = v99
            }
        case 1:
            var v100 int = 32768
            _ = v100
        }
    case 128:
        for fi101 := 0; fi101 < 3; fi101++ {
            var v102 float64 = -2.25
            for fi103 := 0; fi103 < 11; fi103++ {
                v104 := GenS7{F0: -1, F1: 42, F2: 9223372036854775807}
                v105 := GenS7{F0: 1, F1: 11, F2: 42}
                if true {
                    switch -2147483648 {
                    case 165:
                        if false {
                            v98 -= (v98 / ((100) | 1))
                            switch int(v5) {
                            case 1:
                                v106 := []int{fi101, (fi101 % (((fi101 * (fi103 / ((int(v98)) | 1)))) | 1))}
                                _ = v106
                            case 128:
                                switch -1 {
                                case 0:
                                    var v107 uint32 = 346
                                    _ = v107
                                default:
                                    if false {
                                        var v108 int8 = v98
                                        _ = v108
                                    }
                                }
                            case 211:
                                if (!true) {
                                    v6 = (v2 & v2)
                                    v109 := []int{(fi103 ^ 32767), (fi103 ^ fi101), fi103, 0}
                                    v6 += (v6 | v5)
                                    v110 := v3[0]
                                    _ = v109
                                    _ = v110
                                } else {
                                    v5 -= v5
                                    var v111 bool = (!(v8 == (^(v8))))
                                    if (!v111) {
                                        var v112 uint = 4294967295
                                        v113 := GenS3{F0: "", F1: (true && (v1 <= v1)), F2: (fi103 >> ((v112) & 63))}
                                        _ = v112
                                        _ = v113
                                    } else {
                                        delete(v3, 42)
                                    }
                                    _ = v111
                                }
                            }
                            v114 := []int{fi103, -1}
                            _ = v114
                        } else {
                            var v115 float64 = (0.1 / (v102 * (1e10 / ((float64((3.14159 - (0.1 - 3.14159))) * (3.14159 * ((v102 + v102) + v102))) / v102))))
                            var v116 uint8 = 42
                            switch v1 {
                            case "a":
                                v115 = (1.5 - float64(3.14159))
                            default:
                                fmt.Println(v2, v6, v102)
                            }
                            _ = v115
                            _ = v116
                        }
                    case -145:
                        v117 := GenS3{F0: "foo bar", F1: (721 < v8), F2: fi103}
                        _ = v117
                    default:
                        var v118 uint16 = v8
                        _ = v118
                    }
                } else {
                    v1 = v1
                    if (false && (871 != 0)) {
                        var v119 int64 = 42
                        if (v8 == (v8 >> ((143) & 63))) {
                            v98 = 42
                            v120 := GenS8{F0: 1, F1: 580}
                            v121 := []int{-753, 32767, fi103}
                            _ = v120
                            _ = v121
                        } else {
                            v8 = (v8 << ((699) & 63))
                            var v122 uint8 = 128
                            _ = v122
                        }
                        _ = v119
                    }
                }
                switch v1 {
                case "a":
                    v2 -= v2
                default:
                    v102 = 0.1
                }
                if true {
                    v123 := []int{(fi101 * 9223372036854775807), (fi101 << ((65535) & 63))}
                    _ = v123
                }
                _ = fi103
                _ = v104
                _ = v105
            }
            v102 = ((v102 / v102) - v102)
            _ = fi101
            _ = v102
        }
    default:
        var v124 int64 = 1
        _ = v124
    }
    v6 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v98
}
