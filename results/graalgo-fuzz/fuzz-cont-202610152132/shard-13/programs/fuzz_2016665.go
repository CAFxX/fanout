package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 bool }
func genFn0(p0 int32) uint64 {
    var v4 uint64 = 0
    switch "" {
    case "a":
        p0 = (p0 / (((p0 % ((42) | 1))) | 1))
    case "foo bar":
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 uint16 = 65535
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := []int{(^(fi5)), (fi5 + (-(fi7))), fi7}
                if len(v8) > 0 {
                    v8[0] = fi7
                }
                _ = fi7
                _ = v8
            }
            v4 += v4
            _ = fi5
            _ = v6
        }
    case "hello":
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int64 = 42
            p0 -= (p0 ^ (p0 | p0))
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        switch (("" + "foo bar") + ("" + (("a" + (("hello" + ("foo bar" + "a")) + ("\n\t" + "a"))) + ""))) {
        case "a":
            var v12 int64 = 0
            _ = v12
        default:
            var v13 uint16 = 255
            _ = v13
        }
    }
    switch 32768 {
    case -527:
        for fi14 := 0; fi14 < 17; fi14++ {
            switch "a" {
            case "foo bar":
                v15 := []int{(-(fi14))}
                _ = v15
            default:
                fmt.Println(p0, v4)
            }
            if true {
                p0 = (p0 | (p0 - (p0 % (((-(p0))) | 1))))
            }
            if true {
                v16 := []int{}
                _ = v16
            } else {
                switch fi14 {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 9; fi17++ {
                        v18 := GenS1{F0: 32767, F1: 1, F2: 18446744073709551615}
                        if false {
                            v4 += (v4 | (v4 & v4))
                        }
                        _ = fi17
                        _ = v18
                    }
                case -2147483648:
                    if (65535 != 1) {
                        p0 -= p0
                    }
                case -1:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v4, fi19)
                        case "foo bar":
                            if (false || ((true || ((!true) && (false && (919 != 65)))) && true)) {
                                var v20 int = fi19
                                _ = v20
                            } else {
                                fmt.Println(p0, v4)
                                var v21 int16 = 1
                                switch 268 {
                                case 2147483647:
                                    if true {
                                        if ((fi19 - fi19) != fi19) {
                                            var v22 int16 = 42
                                            _ = v22
                                        } else {
                                            switch fi19 {
                                            case 2147483647:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v23 := GenS2{F0: (p0 % (((p0 + int32(fi14))) | 1))}
                                                        _ = v23
                                                    case "hello":
                                                        if false {
                                                            v24 := GenS3{F0: 2, F1: 42}
                                                            var v25 int16 = 32767
                                                            v25 = 42
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                    case "":
                                                        v26 := []int{fi19, -2147483648}
                                                        _ = v26
                                                    default:
                                                        var v27 int64 = -1
                                                        _ = v27
                                                    }
                                                    v28 := []int{(fi19 << ((18446744073709551615) & 63)), 9223372036854775807, (fi19 >> ((548) & 63))}
                                                    var v29 float64 = (0.0 / 3.14159)
                                                    _ = v28
                                                    _ = v29
                                                case "a":
                                                    v4 += v4
                                                case "\n\t":
                                                    switch (fi14 & fi19) {
                                                    case 2147483647:
                                                        var v30 bool = false
                                                        _ = v30
                                                    case 128:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v31 string = ("\n\t" + "")
                                                            _ = v31
                                                        case "":
                                                            if true {
                                                                v32 := GenS4{F0: -1, F1: 1}
                                                                v21 -= (v21 * v21)
                                                                _ = v32
                                                            } else {
                                                                switch fi14 {
                                                                case 127:
                                                                    if (!(1 <= 42)) {
                                                                        if true {
                                                                            v33 := GenS3{F0: 590, F1: v21}
                                                                            v4 = v4
                                                                            v4 = (^(v4))
                                                                            _ = v33
                                                                        } else {
                                                                            if false {
                                                                                var v34 int16 = -32768
                                                                                _ = v34
                                                                            }
                                                                            v35 := GenS1{F0: 32767, F1: (^(p0)), F2: 42}
                                                                            _ = v35
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v4, p0, v4)
                                                                            v36 := GenS5{F0: (p0 << ((339) & 63)), F1: 0, F2: 128}
                                                                            var v37 uint32 = 168
                                                                            _ = v36
                                                                            _ = v37
                                                                        } else {
                                                                            v38 := GenS1{F0: 255, F1: (p0 << ((255) & 63)), F2: 425}
                                                                            _ = v38
                                                                        }
                                                                    }
                                                                case 128:
                                                                    var v39 float32 = ((-2.25 * 1.5) - -2.25)
                                                                    _ = v39
                                                                case -613:
                                                                    var v40 int = (fi14 << ((2) & 63))
                                                                    _ = v40
                                                                default:
                                                                    if true {
                                                                        fmt.Println(p0, v21)
                                                                        v41 := GenS4{F0: 0, F1: -32768}
                                                                        _ = v41
                                                                    }
                                                                }
                                                                if (40 <= 42) {
                                                                    p0 += (^(p0))
                                                                }
                                                                var v42 uint32 = 2147483648
                                                                if true {
                                                                    switch fi14 {
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            v43 := GenS1{F0: 446, F1: 1, F2: 827}
                                                                            switch ((("\n\t" + "") + "") + "\n\t") {
                                                                            case "\n\t":
                                                                                v44 := GenS2{F0: p0}
                                                                                v21 += int16(fi14)
                                                                                p0 -= (p0 + 683)
                                                                                _ = v44
                                                                            default:
                                                                                v45 := []int{(fi19 << ((4294967295) & 63))}
                                                                                _ = v45
                                                                            }
                                                                            fmt.Println(v21, fi19, v4)
                                                                            _ = v43
                                                                        }
                                                                        v21 += 42
                                                                        p0 -= -2147483648
                                                                    case 127:
                                                                        var v46 bool = (42 > -128)
                                                                        _ = v46
                                                                    }
                                                                    var v47 uint8 = 255
                                                                    _ = v47
                                                                }
                                                                _ = v42
                                                            }
                                                            v48 := GenS6{F0: 255, F1: (fi14 - (fi19 % ((fi14) | 1))), F2: 42}
                                                            v21 -= (^(v21))
                                                            _ = v48
                                                        }
                                                    default:
                                                        fmt.Println(v4, fi19, p0)
                                                    }
                                                default:
                                                    if false {
                                                        v49 := GenS6{F0: 128, F1: fi14, F2: (v21 % ((v21) | 1))}
                                                        v4 = (v4 << ((42) & 63))
                                                        v21 -= v21
                                                        _ = v49
                                                    } else {
                                                        var v50 uint64 = v4
                                                        _ = v50
                                                    }
                                                }
                                            case 32767:
                                                v51 := GenS7{F0: 42}
                                                _ = v51
                                            }
                                        }
                                    } else {
                                        var v52 string = "hello"
                                        _ = v52
                                    }
                                default:
                                    v53 := []int{691}
                                    v21 -= 42
                                    _ = v53
                                }
                                _ = v21
                            }
                        case "a":
                            switch "a" {
                            case "hello":
                                var v54 uint64 = (v4 % ((2) | 1))
                                _ = v54
                            case "a":
                                v4 = v4
                            default:
                                fmt.Println(p0)
                            }
                        default:
                            var v55 float64 = 0.1
                            _ = v55
                        }
                        var v56 int = (fi14 >> ((2) & 63))
                        _ = fi19
                        _ = v56
                    }
                }
                if (1000 > -411) {
                    p0 = p0
                    v57 := GenS5{F0: p0, F1: 1, F2: 128}
                    p0 = p0
                    _ = v57
                }
                if true {
                    p0 = (p0 ^ (p0 / ((p0) | 1)))
                    fmt.Println(fi14, v4, p0)
                    fmt.Println(v4)
                } else {
                    v58 := []int{(^(fi14)), (fi14 - 142), (fi14 | fi14), fi14}
                    _ = v58
                }
            }
            _ = fi14
        }
    case -1:
        v59 := []int{-9223372036854775808, 42, 1}
        _ = v59
    case -159:
        fmt.Println(p0, v4)
    }
    if (128 != 42) {
        v4 -= 2
        fmt.Println(v4, v4)
        for fi60 := 0; fi60 < 13; fi60++ {
            var v61 int64 = 9223372036854775807
            for fi62 := 0; fi62 < 7; fi62++ {
                if ((!((0.1 / (3.14159 + (0.0 + float32(3.14159)))) == -2.25)) && true) {
                    fmt.Println(fi62)
                    v63 := GenS2{F0: -2147483648}
                    fmt.Println(v61, v61)
                    var v64 float32 = float32(0.1)
                    switch "foo bar" {
                    case "hello":
                        v64 = v64
                    case "a":
                        var v65 int64 = -1
                        var v66 bool = false
                        v65 += v65
                        _ = v65
                        _ = v66
                    case "foo bar":
                        var v67 float32 = float32(3.14159)
                        _ = v67
                    }
                    _ = v63
                    _ = v64
                }
                var v68 uint16 = 564
                _ = fi62
                _ = v68
            }
            v4 = v4
            _ = fi60
            _ = v61
        }
    } else {
        v4 -= 4294967295
    }
    v69 := []int{4294967295, 127, 4294967295, 888}
    _ = p0
    _ = v4
    _ = v69
    return (v4 % ((v4) | 1))
}
func genFn1(p0 uint32, p1 bool, p2 int) int64 {
    if p1 {
        var v70 uint32 = 726
        var v71 uint = 4294967295
        p2 += -2147483648
        _ = v70
        _ = v71
    } else {
        p1 = (255 >= 296)
    }
    for fi72 := 0; fi72 < 10; fi72++ {
        var v73 uint64 = genFn0(-1)
        _ = fi72
        _ = v73
    }
    if false {
        fmt.Println(p1, p2, p2)
    }
    v74 := []int{p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v74
    return 198
}

func main() {
    v1 := []uint32{4294967295, 2, 124, 42}
    var v2 uint64 = 4294967295
    var v3 int = -751
    if true {
        v2 = (v2 >> ((255) & 63))
        v3 += 197
        v3 += (v3 | 951)
    } else {
        for fi75 := 0; fi75 < 6; fi75++ {
            for fi76 := 0; fi76 < 11; fi76++ {
                v3 += (v3 - fi76)
                v2 = (v2 + (v2 >> ((18446744073709551615) & 63)))
                var v77 uint8 = 255
                _ = fi76
                _ = v77
            }
            v2 = v2
            v2 = v2
            _ = fi75
        }
    }
    _ = len(v1)
    var v78 int8 = -1
    switch int(v78) {
    case 32767:
        for fi79 := 0; fi79 < 15; fi79++ {
            var v80 float32 = 0.1
            v2 = 4294967295
            _ = fi79
            _ = v80
        }
    default:
        v81 := GenS1{F0: 282, F1: 42, F2: 18446744073709551615}
        _ = v81
    }
    if true {
        var v82 bool = true
        _ = v82
    } else {
        v1 = append(v1, 42)
        var v83 string = "a"
        _ = v83
    }
    v2 -= v2
    v84 := GenS8{F0: 100}
    _ = len(v1)
    if (!(!true)) {
        var v85 string = ("hello" + "a")
        if (0.1 == 0.0) {
            v86 := GenS2{F0: -1}
            if (true && (!(true || true))) {
                var v87 uint16 = 2
                _ = v87
            }
            _ = v86
        } else {
            for fi88 := 0; fi88 < 19; fi88++ {
                v2 -= v2
                fmt.Println(v78)
                switch v85 {
                case "foo bar":
                    var v89 float32 = (0.1 - (1e10 - 0.0))
                    _ = v89
                case "\n\t":
                    if (!(203 != 4294967295)) {
                        for fi90 := 0; fi90 < 13; fi90++ {
                            v85 = ((v85 + ((("" + v85) + v85) + (("" + "\n\t") + v85))) + "\n\t")
                            v1 = append(v1, 2)
                            _ = fi90
                        }
                        for fi91 := 0; fi91 < 14; fi91++ {
                            var v92 string = v85
                            _ = fi91
                            _ = v92
                        }
                    } else {
                        for fi93 := 0; fi93 < 20; fi93++ {
                            v2 = genFn0(-1)
                            var v94 uint16 = 32767
                            _ = fi93
                            _ = v94
                        }
                        if true {
                            v85 = ""
                            if false {
                                var v95 uint32 = 1
                                _ = v95
                            }
                        }
                        v3 = (v3 & fi88)
                    }
                default:
                    if false {
                        if (true || (fi88 > (fi88 >> ((255) & 63)))) {
                            v3 = 182
                            fmt.Println(v2, v3, v3)
                            for fi96 := 0; fi96 < 18; fi96++ {
                                switch ("a" + ("" + v85)) {
                                case "":
                                    v97 := GenS9{F0: 2, F1: ""}
                                    _ = v97
                                case "foo bar":
                                    v98 := GenS10{F0: 32767, F1: v2}
                                    _ = v98
                                case "hello":
                                    var v99 uint64 = genFn0(2147483647)
                                    _ = v99
                                default:
                                    v78 = 119
                                }
                                _ = fi96
                            }
                        } else {
                            var v100 uint64 = 378
                            v101 := GenS11{F0: 0}
                            _ = v100
                            _ = v101
                        }
                        v3 += v3
                        for fi102 := 0; fi102 < 9; fi102++ {
                            var v103 string = v85
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v1 = append(v1, 878)
                            v104 := GenS12{F0: true}
                            switch "hello" {
                            case "foo bar":
                                switch v103 {
                                case "a":
                                    fmt.Println(v3)
                                case "hello":
                                    v1 = append(v1, 2147483648)
                                case "":
                                    v85 = v103
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                }
                            case "a":
                                if (!((3.14159 * 3.14159) == (((float32(0.1) / 1.5) / 1.5) - 1e10))) {
                                    _ = len(v1)
                                    v2 -= (v2 ^ (v2 - v2))
                                    var v105 uint16 = 28
                                    _ = v105
                                } else {
                                    v103 = v103
                                    fmt.Println(fi102)
                                }
                            case "\n\t":
                                var v106 bool = (!true)
                                _ = v106
                            default:
                                _ = len(v1)
                            }
                            v107 := GenS8{F0: (v78 >> ((0) & 63))}
                            _ = fi102
                            _ = v103
                            _ = v104
                            _ = v107
                        }
                    } else {
                        var v108 int64 = 9223372036854775807
                        v108 -= (^(v108))
                        v3 = fi88
                        _ = v108
                    }
                }
                _ = fi88
            }
            fmt.Println(v2, v85)
        }
        _ = v85
    } else {
        fmt.Println(v2, v78)
    }
    switch v3 {
    case 2147483647:
        for fi109 := 0; fi109 < 12; fi109++ {
            fmt.Println(v2, fi109)
            _ = fi109
        }
    case -487:
        v110 := GenS4{F0: 127, F1: 1000}
        _ = v110
    case 398:
        var v111 float32 = (float32(1e10) * ((float32(0.1) + 3.14159) + 3.14159))
        _ = v111
    }
    if (true && (!false)) {
        var v112 int = (v3 ^ -1)
        switch "foo bar" {
        case "\n\t":
            if (!false) {
                for fi113 := 0; fi113 < 14; fi113++ {
                    v112 += fi113
                    fmt.Println(v2, v112)
                    _ = fi113
                }
                v114 := GenS6{F0: 42, F1: 0, F2: -32768}
                for fi115 := 0; fi115 < 5; fi115++ {
                    switch (fi115 << ((18446744073709551615) & 63)) {
                    case 147:
                        fmt.Println(v3, v3, v78)
                    case 2147483647:
                        var v116 uint = 434
                        _ = v116
                    case 32767:
                        for fi117 := 0; fi117 < 7; fi117++ {
                            fmt.Println(v112, v78, v78)
                            v3 = v112
                            var v118 uint32 = 1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v119 int8 = v78
                            v112 += v112
                            var v120 uint64 = v2
                            if (false || ((v112 ^ (fi115 << ((65535) & 63))) > 9223372036854775807)) {
                                v120 = (^(v120))
                                switch v3 {
                                case 32768:
                                    v121 := GenS12{F0: ((false && false) && ((!false) || ("hello" == "\n\t")))}
                                    _ = v121
                                case -1:
                                    var v122 uint16 = 42
                                    _ = v122
                                default:
                                    v119 -= v78
                                }
                            }
                            var v123 int32 = -968
                            _ = fi117
                            _ = v118
                            _ = v119
                            _ = v120
                            _ = v123
                        }
                    }
                    if false {
                        for fi124 := 0; fi124 < 20; fi124++ {
                            v125 := GenS12{F0: ((!false) || (219 < -868))}
                            _ = fi124
                            _ = v125
                        }
                    }
                    _ = fi115
                }
                _ = v114
            } else {
                v2 = (v2 * genFn0(int32(v78)))
                for fi126 := 0; fi126 < 12; fi126++ {
                    fmt.Println(v3, v3, v112)
                    switch v3 {
                    case 4294967295:
                        var v127 string = (("\n\t" + "foo bar") + ("" + ""))
                        _ = v127
                    case 346:
                        var v128 int16 = -32768
                        _ = v128
                    }
                    _ = fi126
                }
            }
        case "hello":
            v78 -= v78
        }
        v112 += v112
        _ = v112
    } else {
        var v129 uint8 = 140
        if true {
            fmt.Println(v78, v129, v3)
            v130 := GenS1{F0: 0, F1: -661, F2: 699}
            _ = v130
        }
        _ = v129
    }
    var v131 string = ("" + "foo bar")
    var v132 float32 = 0.1
    v3 = (v3 << ((42) & 63))
    switch v3 {
    case -2147483648:
        var v133 int64 = -1
        _ = v133
    case 1:
        v131 = ("a" + (v131 + v131))
    default:
        if (-72 != v78) {
            fmt.Println(v78, v132, v3)
            var v134 uint64 = v2
            fmt.Println(v2)
            _ = v134
        }
    }
    switch "foo bar" {
    case "a":
        var v135 int32 = 0
        _ = v135
    case "":
        var v136 uint32 = 65535
        _ = v136
    default:
        var v137 uint = 171
        _ = v137
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v131
    _ = v132
    _ = v2
    _ = v3
    _ = v78
    _ = v84
}
