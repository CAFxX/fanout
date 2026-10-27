package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int64; F2 string }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint8; F1 int8; F2 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint; F1 uint64 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int64 = -1
        v7 := GenS1{F0: "a", F1: 2}
        v8 := GenS1{F0: ("foo bar" + "hello"), F1: 1}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 bool = true
    var v10 int8 = -128
    var v11 uint32 = 974
    _ = v10
    _ = v11
    _ = v9
    return -2147483648
}

func main() {
    var v1 int = 9223372036854775807
    var v2 string = "foo bar"
    var v3 uint16 = 638
    v4 := make(map[string]int)
    if (false && true) {
        if (!(!(false && (0 >= 193)))) {
            for fi12 := 0; fi12 < 11; fi12++ {
                if false {
                    v3 += (v3 * v3)
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        if true {
                            v14 := GenS1{F0: (v2 + (v2 + (v2 + v2))), F1: v3}
                            var v15 string = (v2 + "hello")
                            var v16 uint = 255
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v17 := []int{fi13}
                            _ = len(v17)
                            if (255 >= 204) {
                                v18 := GenS2{F0: uint64(v3), F1: false, F2: 1}
                                switch v2 {
                                case "a":
                                    switch fi13 {
                                    case 332:
                                        var v19 float32 = float32(((float64(((0.1 - 1e10) * ((0.1 + (0.1 / 1.5)) / 1.0))) / (3.14159 * 0.1)) / (0.1 + -2.25)))
                                        _ = v19
                                    case -136:
                                        switch ((v2 + v2) + v2) {
                                        case "a":
                                            var v20 uint16 = v3
                                            _ = v20
                                        case "\n\t":
                                            _ = len(v17)
                                        }
                                    case -1:
                                        v21 := GenS3{F0: v1, F1: 127}
                                        _ = v21
                                    }
                                }
                                _ = v18
                            }
                            var v22 int = fi13
                            _ = v17
                            _ = v22
                        }
                        if ((true && false) || false) {
                            v3 = 2
                            var v23 float64 = 0.0
                            var v24 int = (fi12 ^ v1)
                            switch (v2 + (v2 + "\n\t")) {
                            case "hello":
                                var v25 uint32 = 126
                                _ = v25
                            case "\n\t":
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            default:
                                switch 281 {
                                case -223:
                                    switch fi12 {
                                    case 42:
                                        v23 += v23
                                    case 1:
                                        v27 := []int{127, fi13}
                                        _ = v27
                                    }
                                case 128:
                                    var v28 int = (^(v24))
                                    _ = v28
                                }
                            }
                            if (128 > 42) {
                                v1 = v1
                                v29 := []int{}
                                v2 = "\n\t"
                                switch (v2 + v2) {
                                case "a":
                                    v30 := GenS2{F0: 2, F1: true, F2: 0}
                                    _ = v30
                                case "hello":
                                    var v31 int64 = 132
                                    _ = v31
                                case "\n\t":
                                    var v32 uint = 65535
                                    _ = v32
                                default:
                                    v3 = v3
                                }
                                v2 = v2
                                v2 = "\n\t"
                                if (1 > 42) {
                                    if (82 == 95) {
                                        var v33 int64 = -9223372036854775808
                                        var v34 bool = (1 != 0)
                                        var v35 uint16 = 1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                } else {
                                    switch (fi13 / ((fi12) | 1)) {
                                    case -1:
                                        fmt.Println(v3, fi13, v3)
                                    case 4294967295:
                                        if false {
                                            switch v2 {
                                            case "foo bar":
                                                fmt.Println(fi12, fi12, v24)
                                            case "":
                                                v36 := GenS1{F0: v2, F1: (v3 << ((42) & 63))}
                                                _ = v36
                                            case "a":
                                                var v37 uint16 = v3
                                                _ = v37
                                            default:
                                                if true {
                                                    if (true || true) {
                                                        var v38 string = "hello"
                                                        var v39 float64 = 3.14159
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                } else {
                                                    var v40 float64 = (3.14159 + 0.1)
                                                    v4[""] = v1
                                                    fmt.Println(v24, v23, fi13)
                                                    _ = v40
                                                }
                                            }
                                            if (42 != 2147483647) {
                                                v24 -= (v1 ^ (^(fi13)))
                                                v23 += float64(0.1)
                                            }
                                            var v41 uint64 = 18446744073709551615
                                            _ = v41
                                        } else {
                                            var v42 float64 = v23
                                            if ((v23 < 0.0) || (42 <= 1)) {
                                                v43 := GenS3{F0: fi12, F1: -128}
                                                _ = v43
                                            }
                                            _ = v42
                                        }
                                    case 0:
                                        v44 := GenS2{F0: 653, F1: false, F2: 0}
                                        _ = v44
                                    default:
                                        v45 := GenS4{F0: 485}
                                        _ = v45
                                    }
                                    var v46 uint32 = 2147483648
                                    v47 := GenS2{F0: 833, F1: (false && (!((121 != 65535) || (3.14159 != (v23 / v23))))), F2: 1}
                                    _ = v46
                                    _ = v47
                                }
                                _ = v29
                            }
                            v2 = "a"
                            _ = v23
                            _ = v24
                        } else {
                            v48 := GenS1{F0: v2, F1: v3}
                            _ = v48
                        }
                        var v49 bool = true
                        _ = fi13
                        _ = v49
                    }
                    v3 = v3
                }
                _ = fi12
            }
            v3 -= 2
        }
    }
    fmt.Println(v3, v2, v2)
    var v50 uint16 = 130
    v1 = -100
    v2 = (v2 + "foo bar")
    v51 := GenS5{F0: 8, F1: -423, F2: (((v2 + v2) + v2) + "foo bar")}
    switch v1 {
    case -2147483648:
        var v52 int32 = 42
        _ = v52
    case 0:
        v4[v2] = (v1 + (^(v1)))
    }
    delete(v4, ("" + "foo bar"))
    for fi53 := 0; fi53 < 18; fi53++ {
        v50 -= 1
        switch (v2 + (v2 + "\n\t")) {
        case "a":
            var v54 uint = 2
            _ = v54
        case "hello":
            v55 := GenS6{F0: (v2 + (v2 + ((v2 + v2) + ((v2 + (v2 + (v2 + ("foo bar" + v2)))) + (("" + (v2 + (v2 + "\n\t"))) + "\n\t"))))), F1: v1}
            _ = v55
        case "foo bar":
            var v56 uint32 = 0
            _ = v56
        }
        if (!(!(((1 != genFn0()) || true) || false))) {
            v1 += (v1 % ((v1) | 1))
            fmt.Println(v2, v50, fi53)
            for fi57 := 0; fi57 < 15; fi57++ {
                switch v2 {
                case "":
                    v3 -= 1
                default:
                    v1 = fi57
                }
                _ = fi57
            }
        }
        _ = fi53
    }
    v50 -= 439
    var v58 uint32 = 915
    v59 := GenS7{F0: uint8(v50), F1: 17, F2: 0}
    v60 := GenS8{F0: v1}
    fmt.Println(v50)
    v61 := []int{v1, (v1 % (((v1 >> ((18446744073709551615) & 63))) | 1)), v1, v1}
    fmt.Println(v3, v1, v1)
    for fi62 := 0; fi62 < 9; fi62++ {
        if true {
            v1 = v1
        } else {
            if (!(true || false)) {
                fmt.Println(v58)
                v3 += v3
                var v63 uint16 = v50
                _ = v63
            }
        }
        if (false || false) {
            v58 = v58
            switch int(v58) {
            case -616:
                var v64 bool = false
                _ = v64
            case -9223372036854775808:
                for fi65 := 0; fi65 < 19; fi65++ {
                    switch "\n\t" {
                    case "hello":
                        v50 = (v50 * v50)
                    case "\n\t":
                        var v66 uint64 = 18446744073709551615
                        _ = v66
                    case "":
                        if true {
                            var v67 int32 = 821
                            _ = v67
                        } else {
                            switch v2 {
                            case "":
                                fmt.Println(v1, v50)
                            case "\n\t":
                                if true {
                                    v68 := v4[("foo bar" + v2)]
                                    v69 := GenS4{F0: 927}
                                    var v70 int64 = 1
                                    var v71 int16 = -622
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                } else {
                                    if (v1 == fi62) {
                                        switch -1 {
                                        case -1:
                                            fmt.Println(v50)
                                        case 709:
                                            if true {
                                                if (314 == 65535) {
                                                    var v72 int = v1
                                                    var v73 uint16 = 2
                                                    _ = v72
                                                    _ = v73
                                                }
                                                var v74 float32 = (1e10 + (float32(((0.0 + (1e10 * (1.5 + -2.25))) + 1e10)) - 1.5))
                                                var v75 uint32 = (v58 - v58)
                                                var v76 uint64 = 0
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                        }
                                    }
                                    v50 = v50
                                }
                            case "hello":
                                if ((v2 + (v2 + v2)) == "\n\t") {
                                    var v77 int32 = -2147483648
                                    _ = v77
                                } else {
                                    var v78 uint8 = 0
                                    v50 = 32767
                                    _ = v78
                                }
                            default:
                                v79 := GenS1{F0: v2, F1: (v3 % ((42) | 1))}
                                _ = v79
                            }
                            var v80 uint16 = 0
                            v81 := GenS5{F0: 255, F1: -1, F2: ((v2 + v2) + "a")}
                            _ = v80
                            _ = v81
                        }
                    default:
                        _ = len(v4)
                    }
                    v82 := GenS6{F0: v2, F1: fi65}
                    switch (v1 - fi62) {
                    case -2147483648:
                        v1 -= fi65
                        v1 = (fi65 + fi65)
                    case 2147483647:
                        fmt.Println(v50, v50, fi65)
                    case 503:
                        switch (-(fi62)) {
                        case -697:
                            var v83 float32 = ((0.1 + 0.0) + (0.1 - ((-2.25 * 0.1) * 0.0)))
                            _ = v83
                        }
                    default:
                        v2 = ((v2 + ("a" + "a")) + ("hello" + "a"))
                    }
                    v58 = 4294967295
                    v84 := GenS5{F0: 93, F1: -9223372036854775808, F2: v2}
                    _ = fi65
                    _ = v82
                    _ = v84
                }
            case 42:
                for fi85 := 0; fi85 < 2; fi85++ {
                    switch (v1 << ((0) & 63)) {
                    case -718:
                        v2 = (("" + v2) + "a")
                    case 9223372036854775807:
                        _ = len(v4)
                    }
                    _ = fi85
                }
            default:
                v58 += (^(v58))
            }
        } else {
            v86 := GenS6{F0: v2, F1: 4294967295}
            switch v2 {
            case "foo bar":
                var v87 uint = 99
                _ = v87
            case "":
                if (true && true) {
                    v1 = fi62
                    v88 := GenS9{F0: 42}
                    _ = v88
                } else {
                    switch v2 {
                    case "a":
                        var v89 int = fi62
                        _ = v89
                    case "\n\t":
                        var v90 uint8 = 58
                        _ = v90
                    case "foo bar":
                        if false {
                            for fi91 := 0; fi91 < 7; fi91++ {
                                v50 += (v50 * uint16(v58))
                                switch v1 {
                                case -599:
                                    switch -9223372036854775808 {
                                    case 1:
                                        var v92 int = fi91
                                        _ = v92
                                    case 152:
                                        var v93 float32 = 0.0
                                        _ = v93
                                    case 926:
                                        v1 = 0
                                    default:
                                        _ = len(v61)
                                    }
                                case -98:
                                    var v94 string = v2
                                    if (true || ((true || (128 == 54)) || false)) {
                                        v94 = "a"
                                        v95 := GenS8{F0: v1}
                                        _ = v95
                                    }
                                    _ = v94
                                case 751:
                                    v96 := GenS10{F0: -2147483648, F1: 65535}
                                    _ = v96
                                default:
                                    v3 -= v50
                                }
                                v3 = v50
                                _ = fi91
                            }
                            v2 = "hello"
                            switch v2 {
                            case "hello":
                                var v97 uint16 = 255
                                _ = v97
                            }
                        } else {
                            var v98 uint = uint(v58)
                            if true {
                                v99 := GenS3{F0: v1, F1: 47}
                                for fi100 := 0; fi100 < 10; fi100++ {
                                    v1 -= fi100
                                    _ = len(v61)
                                    if true {
                                        if (!true) {
                                            var v101 uint = (v98 * v98)
                                            switch fi62 {
                                            case 0:
                                                v98 = v101
                                            default:
                                                v50 -= v3
                                            }
                                            var v102 uint32 = v58
                                            _ = v101
                                            _ = v102
                                        } else {
                                            if (!false) {
                                                var v103 int64 = -35
                                                v104 := GenS5{F0: 2, F1: (^(v103)), F2: "foo bar"}
                                                var v105 uint64 = 18446744073709551615
                                                fmt.Println(v1, v50)
                                                var v106 uint64 = (^(v105))
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                            }
                                            v107 := GenS11{F0: (v98 << ((v98) & 63)), F1: v50, F2: 0}
                                            var v108 int64 = -794
                                            if false {
                                                switch int(v108) {
                                                case 2147483647:
                                                    v109 := GenS11{F0: (v98 * v98), F1: 65535, F2: 42}
                                                    _ = v109
                                                case 0:
                                                    var v110 uint = (v98 % (((v98 << ((0) & 63))) | 1))
                                                    _ = v110
                                                }
                                            } else {
                                                v111 := v4[v2]
                                                v98 += (^(v98))
                                                _ = v111
                                            }
                                            _ = v107
                                            _ = v108
                                        }
                                        var v112 int32 = -1
                                        _ = v112
                                    } else {
                                        if (!(255 == uint8(v3))) {
                                            fmt.Println(fi100, fi62, fi62)
                                            if true {
                                                v113 := GenS12{F0: (^(v98)), F1: 18446744073709551615}
                                                v114 := GenS6{F0: "hello", F1: fi62}
                                                v115 := GenS5{F0: 168, F1: 42, F2: ("hello" + v2)}
                                                var v116 uint8 = 0
                                                _ = v113
                                                _ = v114
                                                _ = v115
                                                _ = v116
                                            }
                                            var v117 uint16 = v50
                                            _ = v117
                                        }
                                        if ((int64(v58) > -9223372036854775808) || ((^(v3)) <= (v50 << (((^(v98))) & 63)))) {
                                            v2 = (((v2 + "hello") + v2) + "hello")
                                        } else {
                                            fmt.Println(v3)
                                            var v118 float32 = (0.0 / 1.5)
                                            v119 := GenS1{F0: (v2 + v2), F1: uint16(v98)}
                                            _ = v118
                                            _ = v119
                                        }
                                        v61 = append(v61, -864)
                                    }
                                    _ = fi100
                                }
                                for fi120 := 0; fi120 < 3; fi120++ {
                                    v58 = uint32(v98)
                                    fmt.Println(v98, v3)
                                    switch (v2 + v2) {
                                    case "hello":
                                        var v121 bool = false
                                        _ = v121
                                    case "":
                                        v4["\n\t"] = 9223372036854775807
                                    case "a":
                                        if true {
                                            if (true && true) {
                                                if (98 == 42) {
                                                    v58 = 65535
                                                    if len(v61) >= 2 {
                                                        v61 = v61[1:len(v61)]
                                                    }
                                                } else {
                                                    switch v2 {
                                                    case "":
                                                        v4[v2] = -311
                                                    default:
                                                        _ = len(v61)
                                                    }
                                                    fmt.Println(v50)
                                                }
                                            }
                                            var v122 uint = v98
                                            _ = v122
                                        }
                                    }
                                    _ = fi120
                                }
                                _ = v99
                            } else {
                                fmt.Println(v58, fi62)
                            }
                            _ = v98
                        }
                    }
                    if len(v61) > 0 {
                        v61[0] = v1
                    }
                }
            case "hello":
                v123 := GenS7{F0: 255, F1: 127, F2: -124}
                _ = v123
            default:
                delete(v4, "")
            }
            _ = v86
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
