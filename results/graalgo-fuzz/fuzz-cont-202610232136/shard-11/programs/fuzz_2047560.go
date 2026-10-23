package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 bool; F1 uint; F2 int32 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint; F1 int8; F2 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16; F1 uint }
type GenS13 struct { F0 bool }
func genFn0() int {
    var v5 int16 = 1000
    v5 -= 583
    _ = v5
    return 1
}
func genFn1() bool {
    var v6 int64 = -9223372036854775808
    if ((!false) && (1 >= -128)) {
        v7 := []int{32767}
        fmt.Println(v6)
        var v8 uint = 2
        _ = v7
        _ = v8
    }
    var v9 uint = 0
    var v10 uint = v9
    _ = v10
    _ = v6
    _ = v9
    return false
}

func main() {
    var v1 bool = ((true || false) || (!(!(true || true))))
    var v2 float64 = (-2.25 + (3.14159 * -2.25))
    var v3 uint64 = 851
    var v4 bool = (!(2147483647 > -1))
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{fi11, (fi11 << ((970) & 63)), 4294967295, -390}
        switch fi11 {
        case 2147483647:
            v3 -= (v3 * (^(v3)))
        case 9223372036854775807:
            switch fi11 {
            case 2147483647:
                v12 = append(v12, genFn0())
            case 32768:
                for fi13 := 0; fi13 < 1; fi13++ {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    if true {
                        fmt.Println(v2, v3, v3)
                    }
                    _ = fi13
                }
            }
        case 826:
            var v14 int16 = 1
            _ = v14
        }
        switch "hello" {
        case "a":
            v2 += v2
        case "":
            if v4 {
                var v15 int = 1
                fmt.Println(v2, v4)
                _ = v15
            } else {
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v4)
                    switch (fi16 & fi11) {
                    case -624:
                        v4 = v4
                    case 32767:
                        var v17 int64 = 9223372036854775807
                        _ = v17
                    case -582:
                        var v18 uint16 = 678
                        _ = v18
                    }
                    switch (fi16 | genFn0()) {
                    case 32768:
                        if false {
                            v19 := GenS1{F0: 42, F1: 132, F2: -2147483648}
                            var v20 uint = 65535
                            _ = v19
                            _ = v20
                        } else {
                            v1 = (!((1.5 - v2) != v2))
                            v2 -= v2
                        }
                    }
                    _ = fi16
                }
                for fi21 := 0; fi21 < 12; fi21++ {
                    v4 = v4
                    var v22 uint32 = 4294967295
                    _ = fi21
                    _ = v22
                }
            }
        case "\n\t":
            v23 := GenS2{F0: "foo bar"}
            _ = v23
        default:
            v24 := GenS3{F0: 127}
            _ = v24
        }
        _ = fi11
        _ = v12
    }
    var v25 int32 = 0
    if v4 {
        var v26 int16 = 42
        if (!false) {
            switch (("hello" + (("" + "a") + ("a" + ("" + ("a" + ("" + (((("" + ("" + "")) + "\n\t") + ("\n\t" + "hello")) + ""))))))) + "foo bar") {
            case "":
                if v4 {
                    var v27 bool = (18446744073709551615 <= 0)
                    _ = v27
                } else {
                    if false {
                        v28 := []int{}
                        v29 := GenS3{F0: 113}
                        _ = v28
                        _ = v29
                    }
                    if (v1 || v4) {
                        v1 = (!((!((2 != 2147483647) && ((!false) && (32767 >= 853)))) && (v4 || v1)))
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v31 := GenS1{F0: 593, F1: 42, F2: (fi30 % ((fi30) | 1))}
                            v4 = v4
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        v32 := GenS4{F0: (!(0.0 > 1e10))}
                        var v33 int64 = int64(v3)
                        _ = v32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 1; fi34++ {
                        if (65535 <= 1) {
                            v35 := []int{466, fi34, (-(fi34))}
                            for fi36 := 0; fi36 < 12; fi36++ {
                                fmt.Println(v2, v1, v26)
                                fmt.Println(v1, v1)
                                _ = fi36
                            }
                            for fi37 := 0; fi37 < 6; fi37++ {
                                switch "" {
                                case "foo bar":
                                    var v38 float64 = float64(1e10)
                                    _ = v38
                                case "hello":
                                    if v1 {
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                        v39 := GenS1{F0: 505, F1: 42, F2: (fi34 >> ((897) & 63))}
                                        v40 := GenS5{F0: 42}
                                        if ((v25 | v25) < (v25 + (v25 ^ v25))) {
                                            fmt.Println(fi37)
                                            fmt.Println(v1, fi34, v1)
                                        }
                                        var v41 string = ("" + ("foo bar" + "hello"))
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                case "\n\t":
                                    if ((1.5 * (-2.25 + v2)) == float64(((1e10 * 0.1) / 3.14159))) {
                                        v42 := GenS4{F0: (-1 <= 42)}
                                        switch fi34 {
                                        case 9223372036854775807:
                                            v25 = (-(v25))
                                        }
                                        _ = v42
                                    } else {
                                        if (9223372036854775807 != int64(v25)) {
                                            v25 = (v25 >> ((212) & 63))
                                            if (1 < 255) {
                                                var v43 int32 = (v25 / (((v25 | (-(v25)))) | 1))
                                                v44 := GenS6{F0: ((v25 * 0) > 2147483647), F1: 2}
                                                if (!false) {
                                                    fmt.Println(v1, v43, v3)
                                                }
                                                _ = v43
                                                _ = v44
                                            } else {
                                                switch fi37 {
                                                case 795:
                                                    v4 = v1
                                                case -9223372036854775808:
                                                    var v45 int64 = -362
                                                    _ = v45
                                                case -584:
                                                    switch (fi37 + genFn0()) {
                                                    case -1:
                                                        _ = len(v35)
                                                    case 2147483647:
                                                        if (4294967295 == 42) {
                                                            v35 = append(v35, (^(fi37)))
                                                            v1 = true
                                                            var v46 float32 = (0.1 - 0.0)
                                                            _ = v46
                                                        }
                                                    default:
                                                        switch ("" + ("\n\t" + "\n\t")) {
                                                        case "hello":
                                                            fmt.Println(v1)
                                                        case "":
                                                            v1 = v4
                                                        default:
                                                            switch ("\n\t" + ("hello" + "foo bar")) {
                                                            case "a":
                                                                v47 := GenS6{F0: true, F1: 2}
                                                                _ = v47
                                                            case "hello":
                                                                v48 := GenS7{F0: ((1 == 1) || v4), F1: 18446744073709551615, F2: (^(v25))}
                                                                _ = v48
                                                            default:
                                                                var v49 uint64 = (^(v3))
                                                                _ = v49
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v50 int = fi37
                                                    _ = v50
                                                }
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            v51 := GenS6{F0: (true && (255 >= 0)), F1: 255}
                                            var v52 int32 = 42
                                            _ = v51
                                            _ = v52
                                        }
                                        v1 = ((v26 ^ v26) <= (v26 | 1))
                                        v2 += 1.5
                                        _ = len(v35)
                                        _ = len(v35)
                                    }
                                default:
                                    v3 += v3
                                }
                                fmt.Println(fi34, fi34)
                                v26 += v26
                                _ = fi37
                            }
                            _ = v35
                        } else {
                            v1 = v1
                            v53 := []int{1}
                            _ = v53
                        }
                        _ = fi34
                    }
                }
            case "a":
                v2 = v2
            case "\n\t":
                for fi54 := 0; fi54 < 2; fi54++ {
                    var v55 uint = 255
                    v26 -= (v26 & (v26 >> ((v55) & 63)))
                    var v56 int8 = -1
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            }
            switch 730 {
            case 1:
                if (42 >= v3) {
                    v57 := []int{}
                    for fi58 := 0; fi58 < 10; fi58++ {
                        _ = len(v57)
                        if (249 < 0) {
                            v3 -= (^(v3))
                            if v4 {
                                v3 += 0
                                switch "foo bar" {
                                case "hello":
                                    v59 := GenS2{F0: "hello"}
                                    _ = v59
                                case "foo bar":
                                    var v60 bool = (333 != 42)
                                    _ = v60
                                case "\n\t":
                                    var v61 float64 = -2.25
                                    _ = v61
                                }
                            }
                            var v62 uint32 = 879
                            _ = v62
                        }
                        _ = fi58
                    }
                    for fi63 := 0; fi63 < 11; fi63++ {
                        if v4 {
                            if true {
                                var v64 uint64 = v3
                                if false {
                                    var v65 int8 = -108
                                    v66 := GenS8{F0: 612, F1: 0}
                                    _ = v65
                                    _ = v66
                                } else {
                                    for fi67 := 0; fi67 < 1; fi67++ {
                                        v68 := GenS9{F0: 65535, F1: -87, F2: fi67}
                                        v1 = (v26 > (v26 & v26))
                                        var v69 uint8 = 250
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                    v70 := GenS10{F0: v1}
                                    _ = v70
                                }
                                var v71 int32 = (v25 ^ (^(v25)))
                                _ = v64
                                _ = v71
                            }
                        }
                        if v4 {
                            for fi72 := 0; fi72 < 20; fi72++ {
                                v2 = float64(1e10)
                                var v73 int16 = -1
                                var v74 uint16 = 42
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            fmt.Println(fi63, fi63)
                        }
                        v75 := GenS3{F0: 33}
                        _ = fi63
                        _ = v75
                    }
                    _ = v57
                }
            default:
                if (false && (((v26 >> ((255) & 63)) < (v26 % ((42) | 1))) && true)) {
                    switch 746 {
                    case 0:
                        v26 = 32767
                    default:
                        var v76 uint64 = v3
                        _ = v76
                    }
                    v2 -= v2
                    v26 = 0
                }
            }
            v2 -= float64(1.5)
        }
        switch 0 {
        case 337:
            if ((42 == 127) && (43 != 121)) {
                v77 := GenS11{F0: (662 >= -221)}
                fmt.Println(v2)
                _ = v77
            }
        case 32767:
            var v78 uint16 = 951
            _ = v78
        case -237:
            var v79 string = ("foo bar" + ((("hello" + "") + "\n\t") + "a"))
            _ = v79
        default:
            switch "a" {
            case "a":
                if (!(!(!(-87 < -1)))) {
                    fmt.Println(v25, v2)
                    v25 = v25
                    var v80 int32 = v25
                    _ = v80
                }
            case "\n\t":
                for fi81 := 0; fi81 < 5; fi81++ {
                    var v82 float32 = 0.0
                    _ = fi81
                    _ = v82
                }
            default:
                v83 := GenS9{F0: 267, F1: 0, F2: 4294967295}
                _ = v83
            }
        }
        _ = v26
    } else {
        var v84 uint32 = 4294967295
        var v85 float64 = v2
        _ = v84
        _ = v85
    }
    switch genFn0() {
    case 0:
        switch 128 {
        case 128:
            if false {
                fmt.Println(v2, v3, v2)
                v86 := []int{127, genFn0()}
                switch 32767 {
                case 42:
                    v87 := GenS12{F0: 21, F1: 148}
                    _ = v87
                case -1:
                    if v4 {
                        fmt.Println(v2)
                        var v88 string = "a"
                        var v89 uint32 = 65535
                        _ = v88
                        _ = v89
                    } else {
                        var v90 float64 = 1.5
                        _ = v90
                    }
                case -2147483648:
                    var v91 int32 = -358
                    _ = v91
                default:
                    if false {
                        v2 -= 1.5
                        fmt.Println(v2)
                        switch "hello" {
                        case "a":
                            v3 = v3
                        case "\n\t":
                            v92 := GenS1{F0: 65535, F1: 128, F2: 2147483647}
                            _ = v92
                        default:
                            if len(v86) > 0 {
                                v86[0] = 128
                            }
                        }
                    } else {
                        var v93 int8 = -128
                        if len(v86) > 0 {
                            v86[0] = -856
                        }
                        _ = v93
                    }
                }
                _ = v86
            } else {
                if (12 > int16(v25)) {
                    if v1 {
                        var v94 uint32 = 982
                        var v95 float32 = (0.0 + ((3.14159 * float32(0.0)) + float32(v2)))
                        for fi96 := 0; fi96 < 5; fi96++ {
                            fmt.Println(v3, v1)
                            _ = fi96
                        }
                        _ = v94
                        _ = v95
                    }
                    var v97 int32 = (v25 + v25)
                    fmt.Println(v4, v3)
                    _ = v97
                } else {
                    var v98 float32 = ((-2.25 + 3.14159) / 0.1)
                    v99 := []int{650}
                    _ = v98
                    _ = v99
                }
            }
        case 32768:
            for fi100 := 0; fi100 < 14; fi100++ {
                fmt.Println(fi100, fi100, fi100)
                var v101 uint = 2
                v4 = v4
                _ = fi100
                _ = v101
            }
        }
    case -555:
        v102 := []int{-1}
        _ = v102
    case 435:
        v2 += (float64((-2.25 * (((0.1 + (1e10 - ((3.14159 + float32(1.5)) - 1e10))) - 3.14159) - ((1.5 * 1.5) / 1.0)))) / v2)
    default:
        fmt.Println(v3)
    }
    var v103 uint8 = 167
    var v104 int8 = 18
    v1 = (!v1)
    fmt.Println(v2, v4)
    for fi105 := 0; fi105 < 20; fi105++ {
        if (!v1) {
            var v106 int16 = -855
            v107 := []int{}
            v108 := GenS13{F0: (32767 >= -429)}
            _ = v106
            _ = v107
            _ = v108
        }
        _ = fi105
    }
    switch "a" {
    case "a":
        var v109 int8 = (v104 & 127)
        _ = v109
    }
    var v110 int8 = v104
    switch ("hello" + "\n\t") {
    case "":
        var v111 bool = (98 == 928)
        _ = v111
    case "a":
        v112 := []int{1, 32767, -748, 4294967295}
        _ = v112
    case "\n\t":
        var v113 string = "\n\t"
        _ = v113
    }
    if (!((((^(v103)) >= (v103 * v103)) && (4294967295 <= 2147483647)) && false)) {
        v104 = (v110 + (v104 - v110))
    } else {
        for fi114 := 0; fi114 < 12; fi114++ {
            var v115 int8 = int8(fi114)
            var v116 uint64 = v3
            _ = fi114
            _ = v115
            _ = v116
        }
        var v117 float64 = (v2 * v2)
        var v118 string = (("foo bar" + "foo bar") + "hello")
        _ = v117
        _ = v118
    }
    var v119 int32 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v110:", v110)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v103
    _ = v104
    _ = v110
    _ = v119
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
