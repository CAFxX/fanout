package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint16; F2 int32 }
type GenS7 struct { F0 int; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 int64; F1 uint32; F2 uint }
func genFn0() int32 {
    var v6 int32 = 2147483647
    fmt.Println(v6)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{995, fi7, fi7, (-(fi7))}
        v8 = append(v8, (fi7 & fi7))
        if true {
            if false {
                _ = len(v8)
                if (false && ("" != "foo bar")) {
                    if true {
                        v9 := GenS1{F0: 608, F1: 2147483647}
                        v10 := GenS1{F0: 4294967295, F1: 2147483648}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int16 = -1
                        fmt.Println(v11)
                        v12 := GenS2{F0: 2, F1: 1, F2: -1}
                        _ = v11
                        _ = v12
                    }
                    if (2 != 2) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        for fi13 := 0; fi13 < 1; fi13++ {
                            var v14 uint32 = 1
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    switch "" {
                    case "foo bar":
                        if false {
                            v15 := GenS1{F0: 4294967295, F1: 42}
                            _ = v15
                        } else {
                            fmt.Println(fi7)
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v17 := GenS3{F0: 42}
                                _ = fi16
                                _ = v17
                            }
                            v6 += (v6 << ((255) & 63))
                        }
                    case "hello":
                        fmt.Println(fi7, fi7)
                    default:
                        var v18 float64 = (-2.25 + (-2.25 * 1e10))
                        _ = v18
                    }
                } else {
                    v6 = (v6 % ((v6) | 1))
                    if len(v8) > 0 {
                        v8[0] = -345
                    }
                    switch (-(fi7)) {
                    case -753:
                        switch "foo bar" {
                        case "a":
                            for fi19 := 0; fi19 < 17; fi19++ {
                                var v20 int32 = (v6 + (v6 / ((-1) | 1)))
                                var v21 int64 = -950
                                fmt.Println(v20, v21)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        case "hello":
                            for fi22 := 0; fi22 < 8; fi22++ {
                                v6 = (v6 & (v6 | v6))
                                _ = fi22
                            }
                        default:
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v6 -= 0
                                switch "" {
                                case "foo bar":
                                    fmt.Println(fi7)
                                case "\n\t":
                                    v6 -= 1
                                }
                                var v24 uint = 65535
                                _ = fi23
                                _ = v24
                            }
                        }
                    case 42:
                        var v25 bool = false
                        _ = v25
                    case 586:
                        var v26 int64 = 666
                        _ = v26
                    }
                }
                var v27 float64 = 0.0
                _ = v27
            } else {
                if len(v8) > 0 {
                    v8[0] = fi7
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    if (42 >= 65535) {
                        fmt.Println(fi7)
                    }
                    fmt.Println(v6)
                    var v29 float64 = 0.0
                    var v30 uint8 = 255
                    if (!(961 < 32767)) {
                        fmt.Println(fi28)
                    } else {
                        var v31 uint8 = 127
                        v30 -= v31
                        v29 = v29
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        }
        _ = fi7
        _ = v8
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        fmt.Println(fi32, v6)
        _ = fi32
    }
    var v33 int16 = -32768
    _ = v33
    _ = v6
    return v6
}
func genFn1() uint {
    v34 := GenS3{F0: -1}
    if false {
        v35 := GenS4{F0: 18446744073709551615}
        _ = v35
    } else {
        var v36 uint8 = 10
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 uint = 935
            var v40 bool = (!(v39 > (v39 << ((542) & 63))))
            switch (("" + (("" + "hello") + "a")) + ("" + "foo bar")) {
            case "foo bar":
                if (2 != 0) {
                    v39 -= 1
                    if (232 >= 1) {
                        fmt.Println(v40, v40)
                    } else {
                        var v41 string = "a"
                        _ = v41
                    }
                    v42 := GenS5{F0: fi37}
                    _ = v42
                } else {
                    fmt.Println(fi37, v40)
                    v43 := GenS5{F0: (fi37 >> (((^(v39))) & 63))}
                    _ = v43
                }
            case "\n\t":
                v44 := GenS3{F0: -929}
                switch fi37 {
                case 42:
                    fmt.Println(v40)
                case 32767:
                    var v45 uint = (v39 & v39)
                    _ = v45
                case 127:
                    var v46 int = fi38
                    _ = v46
                default:
                    v47 := GenS1{F0: 42, F1: 1}
                    _ = v47
                }
                var v48 uint16 = 32767
                _ = v44
                _ = v48
            case "":
                if (v40 && false) {
                    if v40 {
                        v49 := []int{fi38, fi38}
                        _ = v49
                    } else {
                        if (!(9223372036854775807 >= 1)) {
                            var v50 int32 = 2147483647
                            if v40 {
                                if v40 {
                                    fmt.Println(v39)
                                    if (!(!(v39 == (v39 & v39)))) {
                                        v40 = ((^(v39)) >= v39)
                                        if false {
                                            v40 = ((0.0 - 0.0) > 3.14159)
                                        } else {
                                            var v51 uint8 = 223
                                            var v52 int8 = 100
                                            _ = v51
                                            _ = v52
                                        }
                                        if v40 {
                                            var v53 uint16 = 42
                                            v54 := GenS2{F0: 42, F1: (v39 | v39), F2: 0}
                                            _ = v53
                                            _ = v54
                                        } else {
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v40)
                                            }
                                            var v55 bool = (true || v40)
                                            if false {
                                                v56 := []int{(^(fi37)), fi38, (fi37 * fi37)}
                                                _ = v56
                                            } else {
                                                fmt.Println(v55, fi38)
                                            }
                                            _ = v55
                                        }
                                    }
                                    v57 := GenS1{F0: 2, F1: 1}
                                    _ = v57
                                } else {
                                    v58 := GenS6{F0: 1, F1: 508, F2: v50}
                                    var v59 uint = v39
                                    _ = v58
                                    _ = v59
                                }
                                if (2147483648 < 0) {
                                    var v60 int16 = 1
                                    if v40 {
                                        v60 = v60
                                    } else {
                                        v60 -= (v60 >> (((v39 | (v39 / ((v39) | 1)))) & 63))
                                        fmt.Println(fi37, v50, v50)
                                    }
                                    v40 = (104 == 42)
                                    _ = v60
                                }
                                v50 += (-(v50))
                                var v61 uint64 = 42
                                v62 := []int{fi38}
                                _ = v61
                                _ = v62
                            }
                            var v63 uint64 = 4294967295
                            _ = v50
                            _ = v63
                        } else {
                            v64 := GenS5{F0: fi37}
                            _ = v64
                        }
                    }
                    v65 := []int{fi38}
                    _ = v65
                } else {
                    v39 = (v39 >> ((v39) & 63))
                }
            default:
                v66 := []int{fi38}
                _ = v66
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        for fi67 := 0; fi67 < 8; fi67++ {
            v68 := GenS6{F0: 1, F1: 213, F2: int32(fi37)}
            switch (fi37 * (-(fi37))) {
            case 32768:
                v69 := GenS2{F0: 127, F1: 2, F2: 1000}
                _ = v69
            case -2147483648:
                if (-447 != 42) {
                    var v70 bool = ("foo bar" < ("foo bar" + "hello"))
                    _ = v70
                } else {
                    var v71 int64 = 1
                    var v72 int32 = -591
                    _ = v71
                    _ = v72
                }
            }
            _ = fi67
            _ = v68
        }
        var v73 uint8 = 127
        _ = fi37
        _ = v73
    }
    _ = v34
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]uint32)
    var v3 float32 = (0.1 / (0.1 + 0.0))
    v4 := []string{("foo bar" + "a"), "foo bar", "hello", "\n\t"}
    v5 := []string{(("hello" + (("\n\t" + ("\n\t" + "foo bar")) + (("\n\t" + "foo bar") + ("a" + "a")))) + ""), "foo bar", "a"}
    v1 -= 219
    fmt.Println(v3)
    if (-591 != -627) {
        var v74 uint32 = 520
        if (((((0.0 * 3.14159) / 1.0) - 1.5) / float64((1e10 / -2.25))) == 1e10) {
            v75 := GenS3{F0: 1000}
            for fi76 := 0; fi76 < 18; fi76++ {
                v77 := GenS2{F0: v1, F1: 905, F2: -258}
                var v78 float64 = -2.25
                _ = fi76
                _ = v77
                _ = v78
            }
            v79 := GenS3{F0: 0}
            _ = v75
            _ = v79
        } else {
            _ = len(v4)
            for fi80 := 0; fi80 < 2; fi80++ {
                var v81 float64 = ((3.14159 + -2.25) - -2.25)
                fmt.Println(fi80, v1)
                _ = fi80
                _ = v81
            }
            if true {
                if true {
                    if true {
                        fmt.Println(v3)
                        v82 := GenS3{F0: 960}
                        _ = v82
                    } else {
                        var v83 uint32 = 42
                        switch int(v1) {
                        case 4294967295:
                            switch 617 {
                            case 4294967295:
                                v84 := GenS2{F0: (v1 * (^(v1))), F1: 42, F2: -32768}
                                _ = v84
                            case -259:
                                v85 := GenS4{F0: 677}
                                _ = v85
                            case 32768:
                                fmt.Println(v83, v3)
                            default:
                                _ = len(v4)
                            }
                        default:
                            v1 -= v1
                        }
                        v86 := GenS7{F0: 127, F1: v1, F2: 40}
                        _ = v83
                        _ = v86
                    }
                } else {
                    v87 := v2[("a" + "\n\t")]
                    if false {
                        v87 = v87
                    } else {
                        if (!(42 > 32767)) {
                            switch 4294967295 {
                            case 25:
                                switch -354 {
                                case 32767:
                                    v88 := GenS8{F0: 18446744073709551615}
                                    _ = v88
                                case 4294967295:
                                    for fi89 := 0; fi89 < 6; fi89++ {
                                        var v90 string = ("a" + "\n\t")
                                        switch (fi89 * (-(fi89))) {
                                        case 9223372036854775807:
                                            v74 -= (v74 << ((uint(v1)) & 63))
                                        case -9223372036854775808:
                                            for fi91 := 0; fi91 < 1; fi91++ {
                                                if (true || (!false)) {
                                                    var v92 uint8 = (v1 << ((2) & 63))
                                                    var v93 uint32 = 2147483647
                                                    switch (fi91 % ((fi91) | 1)) {
                                                    case -42:
                                                        var v94 int32 = 42
                                                        _ = v94
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    v87 = v74
                                                    if (!true) {
                                                        v90 = (v90 + v90)
                                                        v95 := GenS6{F0: 42, F1: 32767, F2: 0}
                                                        v96 := GenS9{F0: 0}
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    v90 = v90
                                                }
                                                _ = fi91
                                            }
                                        default:
                                            for fi97 := 0; fi97 < 1; fi97++ {
                                                v1 += v1
                                                var v98 uint8 = 0
                                                _ = fi97
                                                _ = v98
                                            }
                                        }
                                        var v99 uint = 1
                                        _ = fi89
                                        _ = v90
                                        _ = v99
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                            case 32768:
                                v100 := GenS9{F0: 9223372036854775807}
                                _ = v100
                            default:
                                if false {
                                    switch ("hello" + ("" + ("foo bar" + "foo bar"))) {
                                    case "\n\t":
                                        for fi101 := 0; fi101 < 11; fi101++ {
                                            var v102 int = (fi101 - fi101)
                                            _ = fi101
                                            _ = v102
                                        }
                                    case "hello":
                                        for fi103 := 0; fi103 < 20; fi103++ {
                                            if false {
                                                var v104 float32 = 3.14159
                                                v1 -= v1
                                                switch ("" + "foo bar") {
                                                case "hello":
                                                    delete(v2, "")
                                                }
                                                _ = v104
                                            } else {
                                                v3 -= (1e10 + (v3 * 0.1))
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v3 = v3
                                            _ = fi103
                                        }
                                    default:
                                        v74 -= v87
                                    }
                                }
                            }
                            var v105 int16 = -1
                            _ = v105
                        } else {
                            switch ("foo bar" + "hello") {
                            case "a":
                                var v106 int64 = 0
                                _ = v106
                            case "hello":
                                fmt.Println(v74, v87, v3)
                            default:
                                v107 := GenS10{F0: -86, F1: 128}
                                _ = v107
                            }
                            switch "hello" {
                            case "hello":
                                switch "" {
                                case "":
                                    v1 = (v1 >> ((287) & 63))
                                default:
                                    fmt.Println(v3, v74, v1)
                                }
                            case "foo bar":
                                v108 := GenS2{F0: v1, F1: 4294967295, F2: -116}
                                _ = v108
                            case "\n\t":
                                if true {
                                    for fi109 := 0; fi109 < 6; fi109++ {
                                        if len(v5) > 0 {
                                            v5[0] = (("a" + "foo bar") + (("foo bar" + "foo bar") + "hello"))
                                        }
                                        _ = fi109
                                    }
                                    var v110 uint = genFn1()
                                    v111 := GenS1{F0: v74, F1: (v74 / ((2) | 1))}
                                    _ = v110
                                    _ = v111
                                } else {
                                    v3 = (((v3 - float32(float64((v3 / (3.14159 - v3))))) + (v3 * (1e10 + ((v3 / (3.14159 + v3)) / (1e10 / (v3 * float32(-2.25))))))) * v3)
                                    v112 := GenS2{F0: v1, F1: 18446744073709551615, F2: 1000}
                                    for fi113 := 0; fi113 < 17; fi113++ {
                                        v1 -= 42
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v114 int32 = -311
                                        _ = fi113
                                        _ = v114
                                    }
                                    _ = v112
                                }
                            default:
                                switch -433 {
                                case 2147483647:
                                    v115 := GenS6{F0: -128, F1: 0, F2: -50}
                                    _ = v115
                                case 710:
                                    v74 -= (^(v74))
                                default:
                                    var v116 uint8 = v1
                                    _ = v116
                                }
                            }
                            v3 += v3
                        }
                    }
                    _ = v87
                }
            }
        }
        delete(v2, "\n\t")
        _ = v74
    }
    v3 += v3
    var v117 int = 32768
    v2["a"] = 256
    if true {
        v118 := GenS2{F0: v1, F1: 155, F2: -32768}
        switch v117 {
        case 1:
            var v119 float64 = -2.25
            _ = v119
        case 0:
            switch (v117 % (((v117 >> ((185) & 63))) | 1)) {
            case 735:
                v1 -= (^(v1))
            case 42:
                if true {
                    switch ("a" + "foo bar") {
                    case "hello":
                        fmt.Println(v1, v1)
                    default:
                        var v120 int = v117
                        _ = v120
                    }
                }
            case 589:
                fmt.Println(v1, v1)
            }
        case 881:
            v3 -= 0.0
        }
        var v121 bool = true
        _ = v118
        _ = v121
    }
    switch ("a" + (("a" + "\n\t") + ("hello" + (("hello" + "") + "foo bar")))) {
    case "\n\t":
        switch v117 {
        case 883:
            v3 -= v3
        case 2147483647:
            var v122 bool = (true && (!(!true)))
            _ = v122
        case -983:
            var v123 int64 = 9223372036854775807
            _ = v123
        }
    case "hello":
        v124 := GenS5{F0: 1}
        _ = v124
    default:
        for fi125 := 0; fi125 < 2; fi125++ {
            fmt.Println(v117, v117, v1)
            _ = len(v5)
            _ = fi125
        }
    }
    var v126 uint8 = (v1 >> ((255) & 63))
    var v127 uint8 = (v126 >> ((4294967295) & 63))
    var v128 int16 = 0
    for fi129 := 0; fi129 < 12; fi129++ {
        v130 := GenS11{F0: -580, F1: 4294967295, F2: 65535}
        v131 := GenS6{F0: 42, F1: 255, F2: genFn0()}
        fmt.Println(v1, v126, v126)
        _ = fi129
        _ = v130
        _ = v131
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v117
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
