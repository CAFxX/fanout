package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 bool; F1 int32; F2 int32 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 int16; F1 string; F2 uint8 }
type GenS14 struct { F0 uint64; F1 int }
func genFn0(p0 string, p1 int32, p2 uint32) uint {
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 3; fi8++ {
        v7 += v7
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 string, p2 int32) int64 {
    var v9 uint32 = 65535
    for fi10 := 0; fi10 < 17; fi10++ {
        if (18446744073709551615 != 42) {
            v9 = v9
            v9 = (v9 ^ 1)
            v11 := GenS1{F0: 0}
            _ = v11
        } else {
            fmt.Println(p0)
            var v12 float64 = 0.0
            switch ("\n\t" + ("hello" + p1)) {
            case "\n\t":
                for fi13 := 0; fi13 < 1; fi13++ {
                    if false {
                        var v14 uint16 = 255
                        p2 -= -2147483648
                        _ = v14
                    } else {
                        var v15 int16 = -1
                        _ = v15
                    }
                    if false {
                        p1 = p1
                        if (-818 > fi13) {
                            fmt.Println(v12, p2, p2)
                            v9 = (v9 & 2147483647)
                        } else {
                            fmt.Println(v12)
                            switch (fi10 - fi10) {
                            case -653:
                                var v16 float64 = (0.0 * v12)
                                _ = v16
                            case -9223372036854775808:
                                switch fi10 {
                                case -446:
                                    if true {
                                        p0 = (p0 + p0)
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v17 int64 = 42
                                            _ = v17
                                        }
                                        fmt.Println(v12, p0)
                                    } else {
                                        switch p1 {
                                        case "":
                                            fmt.Println(v9)
                                        case "hello":
                                            var v18 float32 = 0.1
                                            _ = v18
                                        }
                                        v12 = v12
                                    }
                                case 128:
                                    v19 := []int{fi13, -1, (fi10 / (((fi10 | fi13)) | 1)), (fi13 - fi13)}
                                    _ = v19
                                case 4294967295:
                                    switch int(p2) {
                                    case 42:
                                        var v20 float64 = 0.0
                                        _ = v20
                                    case -917:
                                        v21 := []int{(fi10 - fi10), int(v9), int(v9)}
                                        _ = v21
                                    case -1:
                                        var v22 string = p1
                                        _ = v22
                                    }
                                default:
                                    v23 := GenS1{F0: int8(fi10)}
                                    _ = v23
                                }
                            }
                        }
                    } else {
                        var v24 uint64 = 18446744073709551615
                        _ = v24
                    }
                    _ = fi13
                }
            }
            _ = v12
        }
        fmt.Println(v9, p1)
        p2 = (p2 ^ p2)
        _ = fi10
    }
    v25 := GenS2{F0: -789, F1: -1, F2: 32767}
    v26 := []int{127, 42, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 float64 = (0.1 + 1.5)
    v2 := []uint32{0, 708}
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    var v5 string = "hello"
    var v6 uint16 = 32767
    fmt.Println(v1, v5, v1)
    switch -2147483648 {
    case -2147483648:
        if (2147483647 >= 1) {
            v27 := GenS1{F0: 110}
            for fi28 := 0; fi28 < 5; fi28++ {
                for fi29 := 0; fi29 < 20; fi29++ {
                    if false {
                        switch (v5 + (("hello" + v5) + "hello")) {
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                v30 := GenS2{F0: 9223372036854775807, F1: 971, F2: -926}
                                _ = v30
                            case "a":
                                var v31 uint8 = 1
                                _ = v31
                            case "foo bar":
                                v5 = v5
                                var v32 uint = 1
                                v5 = "foo bar"
                                _ = v32
                            }
                        case "":
                            var v33 int8 = 100
                            _ = v33
                        }
                        v1 = float64((3.14159 * -2.25))
                        var v34 int = fi29
                        _ = v34
                    } else {
                        var v35 uint32 = 758
                        v36 := GenS1{F0: -128}
                        v5 = "hello"
                        _ = v35
                        _ = v36
                    }
                    var v37 int16 = 32767
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    _ = fi29
                    _ = v37
                }
                _ = fi28
            }
            v1 = (0.1 + 0.1)
            _ = v27
        }
    case 2147483647:
        for fi38 := 0; fi38 < 10; fi38++ {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            switch v5 {
            case "\n\t":
                var v39 int8 = -37
                _ = v39
            case "":
                v2 = append(v2, 4294967295)
            case "hello":
                var v40 int8 = -128
                _ = v40
            default:
                if false {
                    if false {
                        switch -194 {
                        case 127:
                            v41 := GenS2{F0: fi38, F1: 435, F2: 1}
                            _ = v41
                        case -1:
                            var v42 int16 = 0
                            _ = v42
                        case -278:
                            if (42 > 42) {
                                var v43 uint16 = v6
                                _ = v43
                            } else {
                                v1 = (float64(1.5) + float64(1.5))
                                var v44 uint64 = 4294967295
                                _ = v44
                            }
                        default:
                            switch v5 {
                            case "":
                                v5 = ((v5 + v5) + v5)
                            case "hello":
                                var v45 float32 = ((float32(v1) + 0.0) + 1e10)
                                _ = v45
                            }
                        }
                    }
                    fmt.Println(v6)
                    v2 = append(v2, 4294967295)
                } else {
                    v46 := GenS3{F0: (v6 | 526), F1: 0}
                    var v47 uint8 = 127
                    fmt.Println(v5)
                    _ = v46
                    _ = v47
                }
            }
            fmt.Println(v5)
            _ = fi38
        }
    case 1:
        var v48 float64 = (v1 + -2.25)
        _ = v48
    }
    if (!false) {
        for fi49 := 0; fi49 < 12; fi49++ {
            v50 := GenS4{F0: true, F1: -664, F2: 242}
            var v51 int = (fi49 >> ((930) & 63))
            switch fi49 {
            case 632:
                for fi52 := 0; fi52 < 14; fi52++ {
                    var v53 uint16 = (v6 | 255)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v54 float64 = -2.25
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            case -2147483648:
                var v55 uint = 893
                _ = v55
            case 127:
                v6 = (v6 * v6)
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        fmt.Println(v6, v6, v6)
        var v56 float32 = 0.1
        _ = v56
    } else {
        v5 = (v5 + v5)
    }
    var v57 float32 = 3.14159
    var v58 uint64 = 0
    v58 -= 18446744073709551615
    var v59 uint32 = 4294967295
    var v60 uint32 = (v59 + (v59 % ((v59) | 1)))
    if len(v2) > 0 {
        v2[0] = 887
    }
    v1 -= 0.1
    if (true || false) {
        if true {
            var v61 int16 = 0
            if true {
                fmt.Println(v60, v6)
                var v62 float64 = 0.1
                fmt.Println(v62, v5, v6)
                _ = v62
            }
            if ((127 == 131) && false) {
                switch -671 {
                case 9223372036854775807:
                    var v63 int32 = -655
                    _ = v63
                case 4294967295:
                    if false {
                        if (-32 != 2147483647) {
                            var v64 float32 = v57
                            _ = v64
                        } else {
                            v61 -= v61
                            var v65 uint8 = 0
                            delete(v3, v59)
                            _ = v65
                        }
                    }
                case -9223372036854775808:
                    var v66 int32 = -2147483648
                    _ = v66
                default:
                    for fi67 := 0; fi67 < 9; fi67++ {
                        v5 = ("a" + (v5 + ("foo bar" + v5)))
                        if false {
                            switch (v5 + (v5 + (v5 + v5))) {
                            case "":
                                if (-1 < 74) {
                                    v68 := GenS5{F0: genFn1(42, "\n\t", -907), F1: v5}
                                    var v69 int8 = 42
                                    if (((-601 != 32767) && false) && true) {
                                        var v70 uint16 = v6
                                        v2 = append(v2, v59)
                                        _ = v70
                                    } else {
                                        switch fi67 {
                                        case 2147483647:
                                            v2 = append(v2, v60)
                                        case 863:
                                            v3[v60] = fi67
                                        default:
                                            var v71 uint = 65535
                                            _ = v71
                                        }
                                        var v72 uint64 = v58
                                        for fi73 := 0; fi73 < 18; fi73++ {
                                            var v74 uint = uint(v60)
                                            var v75 int64 = 531
                                            _ = fi73
                                            _ = v74
                                            _ = v75
                                        }
                                        _ = v72
                                    }
                                    _ = v68
                                    _ = v69
                                }
                            case "\n\t":
                                _ = len(v2)
                            case "hello":
                                v4[v59] = ("a" + (v5 + v5))
                            default:
                                switch fi67 {
                                case 128:
                                    if (false && false) {
                                        fmt.Println(v1, v60)
                                    } else {
                                        if ((v61 + v61) < (v61 % (((^(v61))) | 1))) {
                                            for fi76 := 0; fi76 < 12; fi76++ {
                                                if true {
                                                    v58 += v58
                                                    var v77 string = ""
                                                    _ = v77
                                                }
                                                if (9223372036854775807 >= 1) {
                                                    var v78 uint = 255
                                                    v60 -= (v59 ^ 2147483647)
                                                    if true {
                                                        var v79 string = (v5 + v5)
                                                        switch ((v79 + ((v5 + ("" + v79)) + "\n\t")) + "a") {
                                                        case "foo bar":
                                                            v80 := GenS6{F0: 4, F1: -2147483648}
                                                            _ = v80
                                                        case "\n\t":
                                                            if ((fi67 / ((fi67) | 1)) > (fi67 & (fi76 << (((v78 % (((v78 << (((v78 - 4294967295)) & 63))) | 1))) & 63)))) {
                                                                switch (fi67 - int(v78)) {
                                                                case 1:
                                                                    v81 := v3[uint32(v6)]
                                                                    _ = v81
                                                                case 988:
                                                                    v5 = ((v79 + "a") + "")
                                                                case -5:
                                                                    fmt.Println(v57, v57, v6)
                                                                default:
                                                                    v59 -= v60
                                                                }
                                                                v79 = (v5 + v5)
                                                                fmt.Println(v58, v61)
                                                            } else {
                                                                v82 := GenS7{F0: (!false)}
                                                                _ = v82
                                                            }
                                                        }
                                                        _ = v79
                                                    } else {
                                                        v61 = (v61 - (v61 & (v61 % ((-32768) | 1))))
                                                        v83 := GenS8{F0: 175, F1: -1, F2: v60}
                                                        fmt.Println(fi76, v78, v6)
                                                        v84 := GenS3{F0: 1, F1: (-(v61))}
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    _ = v78
                                                }
                                                var v85 bool = ((v5 + (v5 + "\n\t")) <= "hello")
                                                _ = fi76
                                                _ = v85
                                            }
                                            switch (v5 + "") {
                                            case "":
                                                v59 = (v60 ^ 0)
                                            case "hello":
                                                switch (fi67 - fi67) {
                                                case 9223372036854775807:
                                                    if false {
                                                        v1 += 1e10
                                                        var v86 float64 = v1
                                                        _ = v86
                                                    } else {
                                                        if (uint16(v58) != 2) {
                                                            var v87 uint8 = 42
                                                            if len(v2) > 0 {
                                                                v2[0] = 145
                                                            }
                                                            _ = v87
                                                        } else {
                                                            for fi88 := 0; fi88 < 3; fi88++ {
                                                                v60 += v59
                                                                v5 = v5
                                                                _ = fi88
                                                            }
                                                            switch (fi67 & fi67) {
                                                            case 623:
                                                                var v89 int = (fi67 >> ((903) & 63))
                                                                _ = v89
                                                            }
                                                            switch v5 {
                                                            case "hello":
                                                                _ = len(v4)
                                                            default:
                                                                var v90 float64 = 3.14159
                                                                _ = v90
                                                            }
                                                        }
                                                        v91 := GenS9{F0: (-(fi67))}
                                                        v6 += 13
                                                        _ = v91
                                                    }
                                                }
                                            }
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = (^(v60))
                                            }
                                        }
                                        v92 := GenS10{F0: 128, F1: 65535}
                                        if true {
                                            v93 := GenS11{F0: (("hello" + v5) + "a"), F1: 22}
                                            _ = v93
                                        } else {
                                            v3[56] = 4294967295
                                        }
                                        _ = v92
                                    }
                                }
                            }
                            var v94 float32 = v57
                            switch (v5 + (v5 + "\n\t")) {
                            case "a":
                                v57 -= (v57 / v94)
                            case "foo bar":
                                v59 = 671
                            case "":
                                switch v5 {
                                case "hello":
                                    if false {
                                        for fi95 := 0; fi95 < 8; fi95++ {
                                            var v96 string = v5
                                            var v97 uint64 = v58
                                            var v98 float32 = ((v57 * v94) / v94)
                                            _ = fi95
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                        var v99 uint8 = 0
                                        switch (fi67 >> ((337) & 63)) {
                                        case -724:
                                            v61 -= (v61 % ((-604) | 1))
                                        }
                                        _ = v99
                                    } else {
                                        for fi100 := 0; fi100 < 3; fi100++ {
                                            var v101 int64 = 0
                                            _ = fi100
                                            _ = v101
                                        }
                                        if true {
                                            for fi102 := 0; fi102 < 15; fi102++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v103 := GenS9{F0: 127}
                                                var v104 string = (v5 + v5)
                                                _ = fi102
                                                _ = v103
                                                _ = v104
                                            }
                                        }
                                    }
                                case "":
                                    for fi105 := 0; fi105 < 13; fi105++ {
                                        fmt.Println(v61, v58)
                                        _ = fi105
                                    }
                                case "foo bar":
                                    v57 += v57
                                }
                            default:
                                for fi106 := 0; fi106 < 9; fi106++ {
                                    _ = len(v3)
                                    _ = fi106
                                }
                            }
                            _ = v94
                        }
                        v2 = append(v2, 107)
                        _ = fi67
                    }
                }
                for fi107 := 0; fi107 < 1; fi107++ {
                    var v108 string = "hello"
                    v57 += -2.25
                    for fi109 := 0; fi109 < 15; fi109++ {
                        v110 := GenS11{F0: v5, F1: 0}
                        v111 := GenS12{F0: (v59 ^ v60), F1: 2}
                        v60 = (v60 + (v59 | 42))
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                    _ = fi107
                    _ = v108
                }
            } else {
                v57 = (0.1 / (v57 / (((0.1 * v57) * ((v57 - (-2.25 - v57)) + float32((v1 - (1e10 / (v1 * (float64(v57) * (3.14159 * v1)))))))) + v57)))
                var v112 float32 = v57
                for fi113 := 0; fi113 < 12; fi113++ {
                    var v114 float64 = v1
                    _ = len(v2)
                    v112 = (1.5 * ((1.5 + v57) * v112))
                    _ = fi113
                    _ = v114
                }
                _ = v112
            }
            _ = v61
        } else {
            v2 = append(v2, v60)
            fmt.Println(v60, v1, v1)
            v115 := GenS13{F0: 42, F1: ("foo bar" + ""), F2: 127}
            _ = v115
        }
    } else {
        if (42 <= 0) {
            v116 := GenS14{F0: v58, F1: 127}
            fmt.Println(v6, v57, v57)
            _ = v116
        }
        if false {
            var v117 string = (v5 + "\n\t")
            var v118 int64 = -9223372036854775808
            _ = v117
            _ = v118
        }
    }
    var v119 uint32 = 42
    for fi120 := 0; fi120 < 8; fi120++ {
        var v121 float32 = v57
        v122 := GenS14{F0: 2, F1: (fi120 | fi120)}
        _ = fi120
        _ = v121
        _ = v122
    }
    if len(v2) > 0 {
        v2[0] = v119
    }
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
