package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int64; F2 uint64 }
type GenS7 struct { F0 int16; F1 uint64; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0() int8 {
    if true {
        if false {
            var v4 int = 9223372036854775807
            fmt.Println(v4)
            _ = v4
        } else {
            var v5 uint64 = 42
            v6 := []int{}
            v5 = v5
            _ = v5
            _ = v6
        }
        if true {
            var v7 int8 = 0
            for fi8 := 0; fi8 < 18; fi8++ {
                if true {
                    if false {
                        var v9 uint32 = 909
                        _ = v9
                    }
                    v10 := GenS1{F0: 282, F1: 42}
                    fmt.Println(fi8, v7)
                    _ = v10
                }
                fmt.Println(v7, v7)
                _ = fi8
            }
            _ = v7
        } else {
            var v11 int32 = 0
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS2{F0: -9223372036854775808, F1: 42}
                switch "\n\t" {
                case "hello":
                    if false {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := []int{-729, (fi12 & fi14), (-(fi14)), 42}
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 13; fi16++ {
                            var v17 float32 = 0.0
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        var v18 float32 = 3.14159
                        v19 := GenS2{F0: 47, F1: 255}
                        v20 := []int{}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "a":
                    v21 := []int{}
                    _ = v21
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    }
    var v22 string = "\n\t"
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint32 = 4294967295
        if true {
            v25 := []int{fi23, (fi23 | (-(fi23))), 475, 373}
            if (!(947 > 1)) {
                if (-726 < -1) {
                    if true {
                        var v26 uint64 = 42
                        switch (fi23 % ((fi23) | 1)) {
                        case 0:
                            var v27 uint = 65535
                            _ = v27
                        }
                        v22 = ("a" + v22)
                        _ = v26
                    }
                    v28 := GenS3{F0: -128, F1: 2, F2: true}
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v30 := GenS3{F0: 0, F1: 133, F2: false}
                        v22 = v22
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            } else {
                v31 := GenS3{F0: -128, F1: 255, F2: true}
                if (18446744073709551615 == 407) {
                    switch (v22 + ("hello" + "a")) {
                    case "hello":
                        switch (fi23 & (fi23 >> ((0) & 63))) {
                        case 1:
                            v32 := GenS4{F0: 536, F1: 1}
                            _ = v32
                        case -1:
                            v25 = append(v25, fi23)
                        }
                    }
                    if true {
                        v25 = append(v25, fi23)
                    } else {
                        if (843 > 599) {
                            for fi33 := 0; fi33 < 20; fi33++ {
                                var v34 uint8 = 1
                                var v35 int8 = 42
                                if false {
                                    switch ((((v22 + (((((v22 + (v22 + (((v22 + "hello") + ("\n\t" + v22)) + ""))) + "foo bar") + "foo bar") + v22) + ((v22 + ("a" + "")) + ("\n\t" + "foo bar")))) + v22) + "\n\t") + (v22 + "foo bar")) {
                                    case "foo bar":
                                        var v36 uint16 = 255
                                        _ = v36
                                    case "\n\t":
                                        v22 = (("foo bar" + ((v22 + (v22 + v22)) + v22)) + (("" + v22) + v22))
                                    default:
                                        var v37 float64 = 0.1
                                        _ = v37
                                    }
                                }
                                v22 = "a"
                                fmt.Println(v34)
                                v35 += 1
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        }
                        v38 := GenS4{F0: 1, F1: 2}
                        _ = v38
                    }
                } else {
                    if len(v25) > 0 {
                        v25[0] = fi23
                    }
                    switch v22 {
                    case "\n\t":
                        for fi39 := 0; fi39 < 3; fi39++ {
                            v40 := GenS1{F0: -937, F1: 9223372036854775807}
                            if (true || (true || true)) {
                                v24 = v24
                                fmt.Println(v24, v24, fi39)
                            } else {
                                var v41 uint32 = v24
                                v22 = (((v22 + v22) + (v22 + "foo bar")) + (v22 + ("a" + (((v22 + v22) + v22) + v22))))
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                    case "hello":
                        if true {
                            for fi42 := 0; fi42 < 9; fi42++ {
                                fmt.Println(v22, v24, v22)
                                v25 = append(v25, (fi23 % ((fi42) | 1)))
                                _ = fi42
                            }
                            switch (v22 + ((v22 + v22) + v22)) {
                            case "foo bar":
                                for fi43 := 0; fi43 < 15; fi43++ {
                                    if ((false || true) && true) {
                                        var v44 int = (fi23 ^ -2147483648)
                                        switch v44 {
                                        case 764:
                                            v44 = v44
                                        case 127:
                                            v25 = append(v25, (v44 << ((571) & 63)))
                                        case 32768:
                                            v24 = (v24 / (((v24 << ((42) & 63))) | 1))
                                        }
                                        switch -866 {
                                        case 836:
                                            if (1 < 0) {
                                                v22 = v22
                                                var v45 uint64 = 42
                                                _ = v45
                                            } else {
                                                var v46 string = v22
                                                _ = v46
                                            }
                                        }
                                        _ = v44
                                    } else {
                                        v47 := GenS5{F0: 679}
                                        _ = v47
                                    }
                                    v24 += (v24 & v24)
                                    v48 := GenS2{F0: 911, F1: 692}
                                    _ = fi43
                                    _ = v48
                                }
                            default:
                                var v49 int = fi23
                                _ = v49
                            }
                        } else {
                            _ = len(v25)
                            var v50 float32 = 1e10
                            v51 := GenS6{F0: 128, F1: 9223372036854775807, F2: 0}
                            _ = v50
                            _ = v51
                        }
                    }
                }
                if ((1e10 * 3.14159) != 0.0) {
                    var v52 uint = 394
                    if (2147483647 <= 4294967295) {
                        fmt.Println(fi23, v22)
                        var v53 uint64 = 2
                        _ = v53
                    }
                    switch 2147483647 {
                    case -994:
                        fmt.Println(v52)
                    case 9223372036854775807:
                        switch -2147483648 {
                        case 128:
                            for fi54 := 0; fi54 < 13; fi54++ {
                                fmt.Println(v52, v22)
                                if (!(1 == 255)) {
                                    var v55 uint8 = 255
                                    var v56 int8 = 24
                                    var v57 int = (^(fi54))
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                                if (true || true) {
                                    var v58 float32 = 3.14159
                                    _ = v58
                                } else {
                                    if len(v25) > 0 {
                                        v25[0] = (fi23 + fi23)
                                    }
                                }
                                _ = fi54
                            }
                        case 1:
                            _ = len(v25)
                        case -9223372036854775808:
                            v24 = uint32(v52)
                        default:
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                        }
                    default:
                        switch "a" {
                        case "hello":
                            v52 -= v52
                        case "":
                            v59 := GenS1{F0: -946, F1: 0}
                            _ = v59
                        default:
                            v60 := GenS7{F0: 467, F1: 4294967295, F2: v22}
                            _ = v60
                        }
                    }
                    _ = v52
                } else {
                    var v61 bool = (127 == 1)
                    for fi62 := 0; fi62 < 18; fi62++ {
                        var v63 uint = 2
                        _ = fi62
                        _ = v63
                    }
                    _ = v61
                }
                _ = v31
            }
            v64 := GenS8{F0: 127}
            _ = v25
            _ = v64
        }
        _ = fi23
        _ = v24
    }
    if true {
        if (-36 != 42) {
            if true {
                v22 = v22
            }
            v65 := []int{}
            v22 = "foo bar"
            _ = v65
        } else {
            v66 := GenS8{F0: 127}
            v67 := []int{1, 982, -9223372036854775808}
            switch v22 {
            case "foo bar":
                var v68 uint16 = 69
                _ = v68
            case "a":
                var v69 uint8 = 255
                _ = v69
            default:
                v22 = (((v22 + "hello") + v22) + v22)
            }
            _ = v66
            _ = v67
        }
        v22 = v22
    } else {
        fmt.Println(v22)
        v70 := GenS3{F0: 42, F1: 0, F2: (4294967295 != 4294967295)}
        _ = v70
    }
    v22 = "foo bar"
    _ = v22
    return 42
}
func genFn1() uint32 {
    var v71 uint64 = 42
    fmt.Println(v71)
    switch "a" {
    case "foo bar":
        v71 = (v71 << ((0) & 63))
    case "\n\t":
        var v72 float64 = (0.1 / 3.14159)
        _ = v72
    case "hello":
        if false {
            var v73 int16 = 0
            v71 += v71
            _ = v73
        } else {
            var v74 uint = 970
            _ = v74
        }
    }
    var v75 bool = false
    _ = v71
    _ = v75
    return 208
}

func main() {
    var v1 uint = 255
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    var v76 bool = ((!(false || false)) || true)
    fmt.Println(v1)
    switch 2147483647 {
    case 0:
        if (13 >= 856) {
            v77 := []int{-1}
            switch (("foo bar" + ("a" + "hello")) + "a") {
            case "foo bar":
                v76 = v76
            case "hello":
                v1 = (v1 ^ v1)
            case "a":
                v1 += v1
            default:
                v2 = v2
            }
            v2 += v2
            _ = v77
        }
    }
    fmt.Println(v2, v76)
    v76 = ((!(0 == 1)) || (!false))
    switch "\n\t" {
    case "":
        if false {
            fmt.Println(v2, v2)
            var v78 int16 = 1000
            for fi79 := 0; fi79 < 10; fi79++ {
                v80 := GenS7{F0: (v78 % ((-1) | 1)), F1: 2, F2: "foo bar"}
                for fi81 := 0; fi81 < 14; fi81++ {
                    switch int(v2) {
                    case 824:
                        if (!true) {
                            switch ("\n\t" + "hello") {
                            case "":
                                v76 = false
                            case "\n\t":
                                switch fi79 {
                                case 4294967295:
                                    var v82 int64 = 9223372036854775807
                                    _ = v82
                                case -1:
                                    var v83 uint64 = 0
                                    _ = v83
                                case -685:
                                    v84 := GenS9{F0: (("" + "foo bar") + ""), F1: 1}
                                    _ = v84
                                }
                            case "foo bar":
                                var v85 uint8 = 2
                                _ = v85
                            }
                            var v86 uint64 = 0
                            _ = v86
                        } else {
                            fmt.Println(v78)
                            var v87 int8 = (^(v2))
                            _ = v87
                        }
                    case 75:
                        v2 -= (v2 ^ (-(v2)))
                    }
                    if ((!(743 > v1)) && true) {
                        v88 := GenS10{F0: v1}
                        _ = v88
                    }
                    _ = fi81
                }
                _ = fi79
                _ = v80
            }
            _ = v78
        } else {
            switch "" {
            case "":
                var v89 int8 = 0
                _ = v89
            case "hello":
                v76 = (0.1 < (3.14159 / 1.0))
            }
            v1 -= v1
            for fi90 := 0; fi90 < 16; fi90++ {
                if (v76 || true) {
                    if v76 {
                        v2 = v2
                        v76 = false
                    }
                    for fi91 := 0; fi91 < 19; fi91++ {
                        var v92 int64 = 860
                        fmt.Println(v76, fi91)
                        _ = fi91
                        _ = v92
                    }
                }
                v2 += (v2 / (((-(v2))) | 1))
                for fi93 := 0; fi93 < 19; fi93++ {
                    v94 := GenS5{F0: -698}
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi90, v1)
                    case "hello":
                        var v95 uint32 = 844
                        _ = v95
                    case "foo bar":
                        v3[1] = (fi90 & (-(fi90)))
                        var v96 string = ""
                        var v97 uint8 = 42
                        _ = v96
                        _ = v97
                    }
                    _ = fi93
                    _ = v94
                }
                _ = fi90
            }
        }
    case "hello":
        v2 += (v2 * v2)
    default:
        var v98 uint8 = 0
        _ = v98
    }
    v2 = (v2 & v2)
    var v99 uint16 = 65535
    v76 = (0.1 != -2.25)
    fmt.Println(v2, v76)
    for fi100 := 0; fi100 < 3; fi100++ {
        var v101 uint64 = 639
        if v76 {
            for fi102 := 0; fi102 < 4; fi102++ {
                v103 := GenS6{F0: fi102, F1: 0, F2: v101}
                v104 := GenS1{F0: int16(v99), F1: 42}
                if false {
                    var v105 int16 = -844
                    _ = v105
                } else {
                    fmt.Println(fi102, v101)
                    v76 = v76
                }
                switch (fi100 << ((uint(v99)) & 63)) {
                case -9223372036854775808:
                    if (!false) {
                        v106 := GenS6{F0: (fi102 - (fi100 ^ (-(fi102)))), F1: 1, F2: (^(v101))}
                        _ = v106
                    }
                case 32768:
                    var v107 int = (fi100 / ((fi102) | 1))
                    _ = v107
                default:
                    v101 = 18446744073709551615
                }
                if (int32(v99) >= 2147483647) {
                    v2 += 0
                } else {
                    var v108 float32 = float32(0.0)
                    var v109 int64 = 0
                    _ = v108
                    _ = v109
                }
                _ = fi102
                _ = v103
                _ = v104
            }
        }
        if v76 {
            v2 = v2
            v101 -= (v101 + 18446744073709551615)
        }
        _ = fi100
        _ = v101
    }
    for fi110 := 0; fi110 < 17; fi110++ {
        var v111 bool = (544 > 2)
        if true {
            var v112 uint64 = 0
            for fi113 := 0; fi113 < 13; fi113++ {
                var v114 bool = (-2.25 >= 3.14159)
                _ = fi113
                _ = v114
            }
            _ = v112
        } else {
            var v115 int8 = (^(v2))
            for fi116 := 0; fi116 < 18; fi116++ {
                v117 := GenS9{F0: ("foo bar" + "foo bar"), F1: 42}
                v1 += v1
                _ = fi116
                _ = v117
            }
            _ = v115
        }
        v118 := []int{}
        _ = fi110
        _ = v111
        _ = v118
    }
    v119 := GenS11{F0: 18446744073709551615, F1: 0, F2: 344}
    v120 := []int{9223372036854775807, 32767, 42, 330}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v120:", len(v120), v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v76:", v76)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v76
    _ = v99
}
