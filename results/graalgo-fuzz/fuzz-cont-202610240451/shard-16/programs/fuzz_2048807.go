package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 int; F2 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 uint }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 uint8 = 19
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 float64 = 3.14159
    v5 := make(map[int]uint32)
    v4 = v4
    fmt.Println(v1)
    v1 = 117
    if (32768 < 32768) {
        var v6 uint = 203
        _ = v6
    } else {
        v7 := GenS1{F0: 128, F1: ("hello" + "")}
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: ("hello" + ("foo bar" + "foo bar"))}
    var v9 int32 = 1
    switch ("" + "foo bar") {
    case "\n\t":
        if (false && ((v1 + v1) == 1)) {
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 int64 = -1
                delete(v2, 2147483647)
                _ = fi10
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := []int{}
                _ = fi12
                _ = v13
            }
            if true {
                v14 := []int{9223372036854775807, 2147483647, 245, -2147483648}
                _ = v14
            }
        }
    case "a":
        v15 := []int{}
        _ = v15
    default:
        switch 860 {
        case 128:
            fmt.Println(v4, v9, v4)
        case 9223372036854775807:
            for fi16 := 0; fi16 < 13; fi16++ {
                if true {
                    v17 := GenS1{F0: v1, F1: "\n\t"}
                    var v18 uint16 = 659
                    _ = v17
                    _ = v18
                }
                var v19 uint8 = v1
                var v20 uint8 = (v1 / ((127) | 1))
                _ = fi16
                _ = v19
                _ = v20
            }
        }
    }
    switch 60 {
    case 2147483647:
        if (false || (0 < 1)) {
            delete(v3, 32768)
            v21 := GenS2{F0: 1, F1: v9, F2: ""}
            for fi22 := 0; fi22 < 6; fi22++ {
                v4 = (float64(-2.25) * (3.14159 / v4))
                v1 = (v1 * v1)
                _ = fi22
            }
            _ = v21
        } else {
            v9 -= v9
            v4 += 0.0
        }
    case 42:
        v23 := GenS3{F0: 758}
        _ = v23
    case 127:
        var v24 uint16 = 986
        _ = v24
    default:
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := v2[545]
                _ = fi26
                _ = v27
            }
            if (!false) {
                v1 = v1
                fmt.Println(fi25, fi25)
                var v28 string = "a"
                _ = v28
            }
            v29 := v2[uint32(v1)]
            _ = fi25
            _ = v29
        }
    }
    if false {
        if true {
            for fi30 := 0; fi30 < 11; fi30++ {
                fmt.Println(v9, v9, v4)
                var v31 float64 = v4
                v32 := v3[(fi30 + (fi30 << ((1) & 63)))]
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 uint64 = 195
            v34 := []int{42, -1, 4294967295, int(v1)}
            _ = v33
            _ = v34
        }
    } else {
        if false {
            v4 = v4
        } else {
            switch -2147483648 {
            case 805:
                var v35 int64 = 240
                _ = v35
            case 42:
                switch "a" {
                case "foo bar":
                    for fi36 := 0; fi36 < 5; fi36++ {
                        var v37 bool = true
                        v38 := []int{fi36, (fi36 ^ fi36)}
                        v38 = append(v38, (-(fi36)))
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                case "":
                    v39 := GenS1{F0: (v1 ^ v1), F1: ("\n\t" + ("a" + "hello"))}
                    _ = v39
                case "hello":
                    v4 += v4
                default:
                    v40 := GenS3{F0: 375}
                    _ = v40
                }
            case 188:
                for fi41 := 0; fi41 < 13; fi41++ {
                    v5[(fi41 << ((1) & 63))] = 115
                    if ((v1 / (((v1 - v1)) | 1)) >= 43) {
                        v9 -= (v9 | (v9 % ((v9) | 1)))
                        v4 -= 1.5
                    } else {
                        v42 := GenS2{F0: 1, F1: v9, F2: "foo bar"}
                        v4 = float64(-2.25)
                        if (true && true) {
                            var v43 int32 = (v9 / ((v9) | 1))
                            _ = v43
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                switch "hello" {
                                case "":
                                    switch fi41 {
                                    case -470:
                                        if false {
                                            v44 := GenS4{F0: fi41}
                                            v45 := v3[(fi41 | fi41)]
                                            v4 = 1.5
                                            _ = v44
                                            _ = v45
                                        } else {
                                            v1 -= v1
                                            var v46 float64 = v4
                                            var v47 uint32 = 2
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                case "\n\t":
                                    for fi48 := 0; fi48 < 9; fi48++ {
                                        var v49 uint = 98
                                        var v50 uint64 = 327
                                        v50 = 1
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                case "foo bar":
                                    var v51 uint64 = 1
                                    _ = v51
                                default:
                                    if false {
                                        var v52 bool = false
                                        switch "\n\t" {
                                        case "hello":
                                            v53 := GenS5{F0: 1, F1: 656, F2: 1}
                                            _ = v53
                                        default:
                                            if (18446744073709551615 < 18446744073709551615) {
                                                for fi54 := 0; fi54 < 9; fi54++ {
                                                    v4 = v4
                                                    v1 = v1
                                                    if v52 {
                                                        switch fi41 {
                                                        case -927:
                                                            fmt.Println(v52, v4)
                                                        case 2147483647:
                                                            v55 := v2[0]
                                                            _ = v55
                                                        }
                                                        if (v4 == 1e10) {
                                                            if true {
                                                                fmt.Println(v52, fi41, fi54)
                                                            } else {
                                                                switch (fi54 << ((629) & 63)) {
                                                                case 0:
                                                                    v56 := GenS4{F0: fi41}
                                                                    _ = v56
                                                                case 4294967295:
                                                                    if ((!v52) || v52) {
                                                                        fmt.Println(v1, fi54)
                                                                    } else {
                                                                        if true {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                v9 = v9
                                                                            }
                                                                        } else {
                                                                            var v57 int32 = v9
                                                                            v58 := GenS4{F0: fi54}
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                case -862:
                                                                    fmt.Println(fi54, v52, fi54)
                                                                }
                                                                fmt.Println(v52)
                                                                v9 += (v9 / ((v9) | 1))
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v59 := GenS6{F0: 461}
                                                                _ = v59
                                                            case "\n\t":
                                                                fmt.Println(v52)
                                                            case "":
                                                                v9 -= -434
                                                            default:
                                                                var v60 int64 = -1
                                                                _ = v60
                                                            }
                                                            var v61 int32 = 1
                                                            _ = v61
                                                        }
                                                    }
                                                    var v62 int = fi54
                                                    _ = len(v2)
                                                    _ = fi54
                                                    _ = v62
                                                }
                                                if v52 {
                                                    v63 := GenS4{F0: (fi41 >> ((255) & 63))}
                                                    _ = v63
                                                }
                                            }
                                        }
                                        _ = v52
                                    } else {
                                        if (v1 >= v1) {
                                            fmt.Println(v9)
                                            for fi64 := 0; fi64 < 11; fi64++ {
                                                if (100 > 1) {
                                                    switch "hello" {
                                                    case "a":
                                                        v4 = 3.14159
                                                    }
                                                    v4 -= ((1e10 + 3.14159) - (1e10 * (3.14159 / (((v4 + (v4 - v4)) + (1e10 / v4)) * 1e10))))
                                                    switch (fi41 | (fi64 - 33)) {
                                                    case -812:
                                                        var v65 int = fi41
                                                        _ = v65
                                                    case 1:
                                                        if (false && (42 == 77)) {
                                                            v66 := GenS6{F0: 2147483647}
                                                            v1 += (v1 * v1)
                                                            v4 = float64(1e10)
                                                            _ = v66
                                                        } else {
                                                            fmt.Println(fi41, v4)
                                                            var v67 uint32 = 65535
                                                            v68 := v3[32768]
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    case 2147483647:
                                                        v9 = -1
                                                        _ = len(v5)
                                                        v1 = v1
                                                        if (2 == 834) {
                                                            if false {
                                                                if (!(false || false)) {
                                                                    if (((v1 - v1) < v1) && (((v1 | v1) > v1) || (false && true))) {
                                                                        fmt.Println(v4, v4)
                                                                        var v69 int32 = 0
                                                                        v70 := []int{42, 32768, fi41}
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                    v71 := v3[fi64]
                                                                    _ = v71
                                                                }
                                                                v4 += v4
                                                                if false {
                                                                    var v72 int32 = (v9 << ((65535) & 63))
                                                                    if false {
                                                                        var v73 uint = 484
                                                                        fmt.Println(v1, fi41)
                                                                        _ = v73
                                                                    }
                                                                    _ = v72
                                                                } else {
                                                                    delete(v2, 593)
                                                                    var v74 int64 = -649
                                                                    var v75 int8 = int8(v9)
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                            }
                                                        }
                                                        v1 = v1
                                                    }
                                                }
                                                v76 := GenS4{F0: (fi64 - (fi64 >> ((65535) & 63)))}
                                                _ = fi64
                                                _ = v76
                                            }
                                        } else {
                                            v77 := GenS3{F0: 999}
                                            v1 -= (v1 * 165)
                                            _ = v77
                                        }
                                        _ = len(v5)
                                    }
                                }
                            case "\n\t":
                                v9 = v9
                            case "":
                                for fi78 := 0; fi78 < 17; fi78++ {
                                    var v79 int8 = 3
                                    v4 += (0.1 * v4)
                                    fmt.Println(v79, fi41)
                                    _ = fi78
                                    _ = v79
                                }
                            }
                        }
                        _ = v42
                    }
                    _ = fi41
                }
            default:
                for fi80 := 0; fi80 < 5; fi80++ {
                    for fi81 := 0; fi81 < 15; fi81++ {
                        fmt.Println(fi80, fi81)
                        _ = fi81
                    }
                    switch fi80 {
                    case -2147483648:
                        switch "" {
                        case "foo bar":
                            var v82 int = -778
                            _ = v82
                        default:
                            var v83 int = (fi80 | fi80)
                            _ = v83
                        }
                    case 302:
                        var v84 int8 = 127
                        _ = v84
                    case -31:
                        var v85 float64 = v4
                        _ = v85
                    }
                    v1 = (v1 * 19)
                    _ = fi80
                }
            }
            v86 := GenS1{F0: 128, F1: "\n\t"}
            switch -1 {
            case -873:
                v9 = (v9 + 1)
            }
            _ = v86
        }
        v4 = v4
    }
    if true {
        switch -423 {
        case 2147483647:
            if true {
                v1 = 1
                v4 += v4
                v9 -= v9
            } else {
                for fi87 := 0; fi87 < 13; fi87++ {
                    var v88 int = 128
                    v89 := GenS6{F0: (^(v9))}
                    if true {
                        for fi90 := 0; fi90 < 10; fi90++ {
                            v91 := []int{(v88 * 32767), fi90}
                            if len(v91) > 0 {
                                v91[0] = fi90
                            }
                            delete(v3, fi90)
                            _ = fi90
                            _ = v91
                        }
                        switch ("a" + "hello") {
                        case "\n\t":
                            v92 := GenS5{F0: 0, F1: 153, F2: 42}
                            _ = v92
                        case "":
                            for fi93 := 0; fi93 < 18; fi93++ {
                                fmt.Println(v1, v9, v9)
                                var v94 uint32 = 65535
                                v4 = v4
                                _ = fi93
                                _ = v94
                            }
                        default:
                            for fi95 := 0; fi95 < 19; fi95++ {
                                var v96 uint32 = 501
                                v96 = v96
                                v88 = v88
                                _ = fi95
                                _ = v96
                            }
                        }
                    }
                    _ = fi87
                    _ = v88
                    _ = v89
                }
            }
        case 42:
            var v97 int32 = 957
            _ = v97
        case 988:
            for fi98 := 0; fi98 < 20; fi98++ {
                v3[fi98] = 2
                _ = fi98
            }
        }
        var v99 uint64 = 411
        _ = v99
    } else {
        if true {
            v100 := []int{-842, 32768}
            for fi101 := 0; fi101 < 20; fi101++ {
                fmt.Println(v1)
                _ = len(v2)
                var v102 string = "\n\t"
                _ = fi101
                _ = v102
            }
            if false {
                fmt.Println(v1)
            } else {
                var v103 uint64 = 0
                for fi104 := 0; fi104 < 2; fi104++ {
                    for fi105 := 0; fi105 < 19; fi105++ {
                        delete(v2, 0)
                        var v106 uint64 = 1
                        var v107 float32 = 3.14159
                        var v108 int8 = 1
                        _ = fi105
                        _ = v106
                        _ = v107
                        _ = v108
                    }
                    _ = fi104
                }
                v109 := GenS2{F0: -1, F1: v9, F2: ""}
                _ = v103
                _ = v109
            }
            _ = v100
        }
        switch 2147483647 {
        case 127:
            for fi110 := 0; fi110 < 10; fi110++ {
                _ = len(v2)
                if true {
                    for fi111 := 0; fi111 < 8; fi111++ {
                        delete(v2, 65535)
                        var v112 int16 = 42
                        if true {
                            v1 = (v1 * v1)
                            v5[(fi110 * -2147483648)] = 4294967295
                        } else {
                            var v113 float64 = v4
                            var v114 uint8 = v1
                            v114 -= (v1 / ((1) | 1))
                            _ = v113
                            _ = v114
                        }
                        _ = fi111
                        _ = v112
                    }
                    fmt.Println(v1, v9, v1)
                    var v115 float64 = (v4 + 1.5)
                    _ = v115
                }
                var v116 uint64 = 1
                _ = fi110
                _ = v116
            }
        case 2147483647:
            var v117 int16 = -32768
            _ = v117
        case 236:
            if false {
                v118 := GenS5{F0: 18446744073709551615, F1: 4, F2: 1}
                v119 := GenS2{F0: 9223372036854775807, F1: (v9 - v9), F2: ("\n\t" + "a")}
                _ = v118
                _ = v119
            } else {
                var v120 bool = false
                var v121 int8 = 42
                if true {
                    v3[2147483647] = 1
                    var v122 float64 = v4
                    v123 := GenS7{F0: 0}
                    _ = v122
                    _ = v123
                } else {
                    v124 := []int{42, 1, -2147483648, 127}
                    v121 += (v121 % ((127) | 1))
                    _ = v124
                }
                _ = v120
                _ = v121
            }
        default:
            if true {
                v125 := GenS8{F0: v9}
                _ = v125
            } else {
                var v126 int8 = 60
                var v127 bool = false
                if v127 {
                    v128 := GenS9{F0: 32767, F1: 1, F2: true}
                    _ = v128
                } else {
                    v4 += v4
                    v126 -= 0
                    var v129 uint32 = 872
                    _ = v129
                }
                _ = v126
                _ = v127
            }
        }
        switch "hello" {
        case "hello":
            v130 := GenS4{F0: 791}
            _ = v130
        case "a":
            v131 := GenS10{F0: 42}
            _ = v131
        default:
            switch 0 {
            case 127:
                if false {
                    if true {
                        v1 = v1
                        if (!(-1 < 892)) {
                            v4 = v4
                            fmt.Println(v4, v9)
                        } else {
                            var v132 bool = (!true)
                            _ = v132
                        }
                        switch "foo bar" {
                        case "hello":
                            for fi133 := 0; fi133 < 9; fi133++ {
                                v134 := GenS10{F0: 598}
                                fmt.Println(fi133, v4, v1)
                                if (0.0 <= float32(1e10)) {
                                    v135 := GenS11{F0: 1, F1: 1}
                                    switch "foo bar" {
                                    case "":
                                        var v136 uint64 = 554
                                        _ = v136
                                    case "a":
                                        v137 := GenS1{F0: (v1 + v1), F1: "foo bar"}
                                        _ = v137
                                    case "hello":
                                        v138 := GenS10{F0: 255}
                                        _ = v138
                                    default:
                                        var v139 uint16 = 65535
                                        _ = v139
                                    }
                                    fmt.Println(fi133, fi133, fi133)
                                    _ = v135
                                } else {
                                    var v140 uint16 = 245
                                    _ = len(v5)
                                    _ = v140
                                }
                                _ = fi133
                                _ = v134
                            }
                        case "foo bar":
                            switch "" {
                            case "a":
                                fmt.Println(v4, v4)
                            default:
                                v141 := v5[42]
                                _ = v141
                            }
                        }
                    }
                } else {
                    v142 := GenS12{F0: v1}
                    v143 := []int{252}
                    if true {
                        if true {
                            switch 127 {
                            case 2147483647:
                                v9 -= 0
                            case -335:
                                v4 = ((((1e10 - (v4 - 0.1)) + 0.1) * (float64(0.0) + v4)) / v4)
                            case 778:
                                v1 = (v1 * (v1 << ((257) & 63)))
                            }
                        }
                        var v144 string = "foo bar"
                        v145 := GenS7{F0: 254}
                        _ = v144
                        _ = v145
                    }
                    _ = v142
                    _ = v143
                }
            default:
                switch (("\n\t" + ("foo bar" + "a")) + "a") {
                case "":
                    v4 = (v4 - float64(0.0))
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
