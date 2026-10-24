package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int8; F1 int64; F2 int8 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint; F1 int64; F2 uint32 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 int32; F1 int8; F2 int64 }
type GenS13 struct { F0 string; F1 uint8 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int32; F1 int; F2 int8 }
type GenS18 struct { F0 uint; F1 uint16; F2 uint8 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        var v7 uint8 = 234
        _ = fi6
        _ = v7
    }
    var v8 bool = (true || false)
    var v9 int64 = -9223372036854775808
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int32, p1 uint8) int {
    v10 := GenS1{F0: false, F1: 255, F2: p1}
    if (!true) {
        v11 := []int{42}
        p0 = (-(p0))
        _ = v11
    }
    if true {
        fmt.Println(p1, p0)
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                if ("\n\t" <= ((((("\n\t" + "a") + ("" + "hello")) + "a") + ("a" + "hello")) + "\n\t")) {
                    v14 := GenS2{F0: 2}
                    fmt.Println(p0, fi12)
                    switch ("hello" + "") {
                    case "hello":
                        switch ("\n\t" + ("a" + "hello")) {
                        case "foo bar":
                            v15 := GenS3{F0: -1}
                            _ = v15
                        case "":
                            var v16 float64 = (-2.25 + float64((-2.25 * 3.14159)))
                            v17 := []int{fi13, (-(fi13))}
                            _ = v16
                            _ = v17
                        case "a":
                            fmt.Println(p0, fi13)
                        default:
                            var v18 float64 = (((0.0 * 1e10) + 0.0) - 1e10)
                            _ = v18
                        }
                    case "a":
                        if (true && (p1 < p1)) {
                            if (2 != 966) {
                                var v19 int64 = -9223372036854775808
                                v20 := []int{fi13, 9223372036854775807}
                                _ = v19
                                _ = v20
                            } else {
                                p0 = (p0 ^ p0)
                                v21 := []int{fi13, (fi13 & fi12)}
                                _ = v21
                            }
                            if (!(p0 <= p0)) {
                                v22 := []int{(fi13 >> ((860) & 63))}
                                _ = v22
                            }
                            v23 := []int{fi13, (^(fi12)), (fi12 << ((0) & 63))}
                            _ = v23
                        }
                    case "":
                        if true {
                            if (!(751 >= 956)) {
                                if true {
                                    p1 += p1
                                    var v24 uint32 = 2
                                    fmt.Println(p1)
                                    _ = v24
                                }
                                fmt.Println(p0, fi12, p1)
                                var v25 int64 = 42
                                var v26 int8 = -75
                                _ = v25
                                _ = v26
                            }
                        } else {
                            v27 := []int{fi12, (fi13 | fi12), 32767}
                            v28 := GenS4{F0: p0, F1: -96, F2: (("" + ("a" + "foo bar")) + "foo bar")}
                            var v29 uint16 = 0
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    default:
                        if true {
                            fmt.Println(fi12, p0, fi12)
                            var v30 int64 = 528
                            v31 := GenS5{F0: (449 >= 0)}
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint64 = 42
                            switch fi13 {
                            case 4294967295:
                                var v33 int64 = -1
                                _ = v33
                            case 32768:
                                switch "foo bar" {
                                case "":
                                    if (((0.0 / 3.14159) > (-2.25 - 1.5)) && (280 >= -1)) {
                                        switch "a" {
                                        case "":
                                            v34 := []int{fi12, fi13, (fi12 / ((fi13) | 1))}
                                            _ = v34
                                        case "foo bar":
                                            var v35 int16 = -1
                                            _ = v35
                                        }
                                    }
                                default:
                                    if false {
                                        v32 += v32
                                        p0 = (p0 % ((p0) | 1))
                                        v36 := []int{(fi13 % (((^(fi12))) | 1))}
                                        _ = v36
                                    }
                                }
                            }
                            v37 := []int{}
                            fmt.Println(fi13)
                            if len(v37) > 0 {
                                v37[0] = 419
                            }
                            _ = v32
                            _ = v37
                        }
                    }
                    if (uint(p1) == 42) {
                        var v38 uint = 516
                        p0 -= (p0 | p0)
                        _ = v38
                    }
                    v39 := []int{}
                    if false {
                        fmt.Println(p1, p0, p0)
                        var v40 uint64 = 4294967295
                        _ = v40
                    } else {
                        if false {
                            _ = len(v39)
                            p0 += (p0 / (((p0 << ((255) & 63))) | 1))
                            fmt.Println(fi13, p1, fi13)
                        } else {
                            p0 = p0
                            if (true || false) {
                                p0 = p0
                                p0 -= p0
                                v41 := GenS4{F0: -2147483648, F1: -128, F2: ("hello" + "")}
                                _ = v41
                            }
                            var v42 int = 32768
                            v43 := GenS2{F0: 42}
                            _ = v42
                            _ = v43
                        }
                        p1 = p1
                    }
                    _ = v14
                    _ = v39
                }
                v44 := []int{-2147483648}
                fmt.Println(fi13, fi12)
                p1 = (p1 >> ((uint(p1)) & 63))
                _ = fi13
                _ = v44
            }
            switch (fi12 * (^(fi12))) {
            case 424:
                var v45 uint16 = 679
                _ = v45
            case 151:
                p1 += 36
            case 0:
                if (false || (32767 < -1)) {
                    fmt.Println(p0, p0)
                    p1 += (p1 ^ 127)
                    v46 := []int{0}
                    _ = v46
                }
            }
            _ = fi12
        }
        for fi47 := 0; fi47 < 19; fi47++ {
            switch "foo bar" {
            case "hello":
                var v48 uint16 = 327
                _ = v48
            }
            _ = fi47
        }
    }
    var v49 int = -429
    v50 := GenS2{F0: 238}
    _ = p0
    _ = p1
    _ = v10
    _ = v49
    _ = v50
    return (-(v49))
}

func main() {
    var v1 float32 = (1.5 * 0.0)
    var v2 uint32 = 808
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v51 := GenS1{F0: false, F1: 0, F2: genFn0()}
    if (!false) {
        fmt.Println(v1)
        v52 := []int{-2147483648, 32768, 0, 1}
        _ = v52
    }
    fmt.Println(v2)
    v53 := []int{0, -2147483648, 4294967295, 809}
    v3 += v3
    v54 := GenS5{F0: (65535 <= 65535)}
    if len(v53) > 0 {
        v53[0] = int(v3)
    }
    var v55 int32 = -1
    var v56 int64 = 1
    if false {
        for fi57 := 0; fi57 < 14; fi57++ {
            var v58 uint64 = 30
            _ = fi57
            _ = v58
        }
    } else {
        if false {
            fmt.Println(v56, v2)
            v1 = 0.0
            for fi59 := 0; fi59 < 9; fi59++ {
                v56 += (-(v56))
                _ = fi59
            }
        } else {
            for fi60 := 0; fi60 < 20; fi60++ {
                var v61 int16 = 1
                _ = fi60
                _ = v61
            }
            if (genFn1(v55, 42) == 127) {
                var v62 int32 = -50
                switch "" {
                case "a":
                    for fi63 := 0; fi63 < 3; fi63++ {
                        if false {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                v3 -= v3
                                v65 := GenS6{F0: true, F1: 18446744073709551615, F2: 2}
                                _ = fi64
                                _ = v65
                            }
                            if (true && (v2 >= 2147483647)) {
                                v1 += v1
                                fmt.Println(v56, fi63)
                            } else {
                                for fi66 := 0; fi66 < 14; fi66++ {
                                    v1 += ((-2.25 + (((((((1e10 - v1) - v1) * v1) / (v1 * v1)) / v1) - (-2.25 + v1)) - 3.14159)) / (v1 * 1e10))
                                    fmt.Println(v62, v1, fi66)
                                    if false {
                                        switch fi63 {
                                        case 567:
                                            if len(v53) >= 2 {
                                                v53 = v53[1:len(v53)]
                                            }
                                        default:
                                            var v67 bool = ((144 != 255) || ((-2.25 < 0.0) && (127 <= 42)))
                                            var v68 float32 = v1
                                            v69 := GenS5{F0: v67}
                                            var v70 string = "a"
                                            var v71 float64 = 0.0
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        }
                                        v72 := GenS4{F0: 42, F1: 42, F2: "a"}
                                        _ = v72
                                    }
                                    _ = fi66
                                }
                            }
                        } else {
                            if (!false) {
                                v53 = append(v53, int(v3))
                                for fi73 := 0; fi73 < 12; fi73++ {
                                    fmt.Println(v1)
                                    fmt.Println(fi73)
                                    v3 -= (v3 + (v3 * v3))
                                    if ((fi63 % ((fi73) | 1)) < fi63) {
                                        v1 -= -2.25
                                    }
                                    _ = fi73
                                }
                            } else {
                                var v74 float64 = 0.1
                                fmt.Println(v56, v55, v55)
                                fmt.Println(v3)
                                _ = v74
                            }
                        }
                        if false {
                            v3 += v3
                        }
                        switch ("" + "foo bar") {
                        case "":
                            fmt.Println(v56, fi63, v2)
                        case "hello":
                            if (1 < 2) {
                                if true {
                                    switch genFn1(v55, 0) {
                                    case 313:
                                        v3 -= (v3 ^ (v3 ^ (v3 % ((v3) | 1))))
                                    case -1:
                                        switch "hello" {
                                        case "":
                                            if (v1 != (float32(float64(float32(1.5))) * float32(0.0))) {
                                                for fi75 := 0; fi75 < 9; fi75++ {
                                                    v3 += -975
                                                    v76 := GenS7{F0: 0, F1: (v56 - (-(v56))), F2: 100}
                                                    var v77 int8 = 100
                                                    if (!true) {
                                                        var v78 uint32 = (v2 - 2147483648)
                                                        if len(v53) > 0 {
                                                            v53[0] = (fi63 ^ -626)
                                                        }
                                                        v56 -= int64(v78)
                                                        _ = v78
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            v3 -= v3
                                                        case "foo bar":
                                                            if false {
                                                                var v79 uint32 = (v2 >> ((4294967295) & 63))
                                                                var v80 int16 = int16(v77)
                                                                var v81 float32 = ((v1 * v1) * (0.0 * v1))
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            switch (fi63 + fi75) {
                                                            case -268:
                                                                v82 := GenS8{F0: 330, F1: v77}
                                                                _ = v82
                                                            default:
                                                                if (3 <= 65535) {
                                                                    v83 := v5[(v2 % (((v2 - 65535)) | 1))]
                                                                    _ = v83
                                                                }
                                                            }
                                                            if (true || true) {
                                                                var v84 int = (fi63 % ((genFn1(int32(v2), genFn0())) | 1))
                                                                var v85 uint64 = 42
                                                                v62 += (-(v62))
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                v86 := GenS8{F0: 234, F1: v77}
                                                                v1 = float32(float64((-2.25 + (v1 * (v1 / (v1 - v1))))))
                                                                _ = v86
                                                            }
                                                            if len(v53) >= 2 {
                                                                v53 = v53[1:len(v53)]
                                                            }
                                                            delete(v5, v2)
                                                        case "hello":
                                                            var v87 string = ""
                                                            _ = v87
                                                        }
                                                    }
                                                    _ = fi75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                fmt.Println(v1)
                                            }
                                        default:
                                            var v88 float32 = (v1 * 3.14159)
                                            _ = v88
                                        }
                                    default:
                                        var v89 uint8 = 135
                                        _ = v89
                                    }
                                } else {
                                    v56 -= 42
                                }
                                fmt.Println(v55, v2, v55)
                            } else {
                                fmt.Println(v2, v62, v55)
                                v3 -= (v3 >> ((662) & 63))
                                for fi90 := 0; fi90 < 17; fi90++ {
                                    v53 = append(v53, -666)
                                    switch "\n\t" {
                                    case "foo bar":
                                        v91 := GenS9{F0: 1, F1: v56, F2: 2}
                                        _ = v91
                                    case "\n\t":
                                        var v92 float32 = (v1 + (float32(1e10) * (v1 / (v1 / 1.0))))
                                        _ = v92
                                    }
                                    v55 = v62
                                    _ = fi90
                                }
                            }
                        case "a":
                            for fi93 := 0; fi93 < 8; fi93++ {
                                switch (fi93 ^ 1) {
                                case 0:
                                    var v94 uint64 = 163
                                    _ = v94
                                }
                                if false {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v56 = v56
                                    case "":
                                        if (!((v55 - (v62 ^ -1)) <= (v55 - 42))) {
                                            switch ((("hello" + "") + "hello") + "foo bar") {
                                            case "foo bar":
                                                v2 = v2
                                            case "":
                                                var v95 float32 = v1
                                                v62 = v55
                                                var v96 string = ("\n\t" + "foo bar")
                                                _ = v95
                                                _ = v96
                                            case "hello":
                                                if false {
                                                    _ = len(v53)
                                                    var v97 int16 = (^(v3))
                                                    _ = v97
                                                } else {
                                                    fmt.Println(v56, fi63)
                                                    if false {
                                                        if false {
                                                            fmt.Println(v55)
                                                        }
                                                        v55 -= 857
                                                        var v98 uint32 = (v2 >> ((2) & 63))
                                                        _ = v98
                                                    } else {
                                                        if (false || true) {
                                                            v3 += 42
                                                            v56 += (^(v56))
                                                            v99 := GenS4{F0: 42, F1: -128, F2: "\n\t"}
                                                            _ = v99
                                                        }
                                                    }
                                                    switch (fi93 ^ 32767) {
                                                    case 42:
                                                        v62 = -1
                                                        v55 = (v55 / ((42) | 1))
                                                    }
                                                }
                                            }
                                        } else {
                                            v53 = append(v53, fi93)
                                            v3 += v3
                                            var v100 int8 = 100
                                            var v101 uint8 = 42
                                            if false {
                                                if (62 >= (v2 | v2)) {
                                                    switch fi63 {
                                                    case 32768:
                                                        if true {
                                                            switch (fi93 % (((fi63 | -595)) | 1)) {
                                                            case 0:
                                                                if false {
                                                                    v2 = (v2 - 570)
                                                                    v62 += (v62 & v55)
                                                                } else {
                                                                    if false {
                                                                        v102 := GenS8{F0: 1, F1: (v100 / (((v100 % ((127) | 1))) | 1))}
                                                                        _ = v102
                                                                    }
                                                                    if false {
                                                                        var v103 uint16 = 554
                                                                        v104 := GenS8{F0: v103, F1: v100}
                                                                        switch (fi93 >> ((563) & 63)) {
                                                                        case -9223372036854775808:
                                                                            var v105 float32 = 0.1
                                                                            _ = v105
                                                                        case 127:
                                                                            var v106 bool = true
                                                                            _ = v106
                                                                        default:
                                                                            v107 := GenS10{F0: (v100 % ((v100) | 1)), F1: fi93}
                                                                            _ = v107
                                                                        }
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                }
                                                            case 32767:
                                                                v108 := GenS11{F0: true, F1: (v3 / (((-(v3))) | 1))}
                                                                _ = v108
                                                            default:
                                                                v109 := GenS12{F0: v55, F1: -128, F2: (v56 % ((465) | 1))}
                                                                _ = v109
                                                            }
                                                            v1 = ((1e10 * (float32((-2.25 - (3.14159 + 0.0))) - 0.0)) / (v1 * 1e10))
                                                        } else {
                                                            v3 += v3
                                                            switch -305 {
                                                            case 9223372036854775807:
                                                                v1 += 1e10
                                                                var v110 int64 = int64(v2)
                                                                v4["foo bar"] = v2
                                                                _ = v110
                                                            case -1:
                                                                fmt.Println(v3, v3, v2)
                                                                var v111 uint8 = (v101 | 255)
                                                                _ = v111
                                                            case 2147483647:
                                                                v55 += (v55 + (v62 & 792))
                                                            default:
                                                                v3 += (v3 - 987)
                                                            }
                                                        }
                                                    case 42:
                                                        v112 := GenS13{F0: "hello", F1: (v101 & uint8(v2))}
                                                        _ = v112
                                                    }
                                                } else {
                                                    var v113 uint8 = v101
                                                    v5[(^(v2))] = uint32(v101)
                                                    v114 := GenS3{F0: v100}
                                                    _ = v113
                                                    _ = v114
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if false {
                                                        v3 += 1
                                                        var v115 uint = 4294967295
                                                        _ = v115
                                                    }
                                                case "foo bar":
                                                    var v116 int = (-(fi93))
                                                    _ = v116
                                                case "a":
                                                    if (!true) {
                                                        v117 := v5[215]
                                                        _ = v117
                                                    } else {
                                                        if (uint(v101) > 546) {
                                                            switch fi63 {
                                                            case 0:
                                                                v118 := GenS2{F0: 42}
                                                                _ = v118
                                                            }
                                                            if false {
                                                                fmt.Println(v1, v1, v100)
                                                                var v119 uint64 = 574
                                                                _ = v119
                                                            } else {
                                                                v100 -= v100
                                                                v53 = append(v53, genFn1((v55 - -1), (^(v101))))
                                                            }
                                                        }
                                                        _ = len(v4)
                                                    }
                                                default:
                                                    var v120 float64 = ((0.1 + -2.25) + (0.0 - 1e10))
                                                    _ = v120
                                                }
                                                v121 := GenS9{F0: 42, F1: (v56 * (v56 << ((18446744073709551615) & 63))), F2: 721}
                                                _ = v121
                                            }
                                            _ = v100
                                            _ = v101
                                        }
                                    }
                                }
                                var v122 string = ("hello" + ("foo bar" + "\n\t"))
                                _ = fi93
                                _ = v122
                            }
                        default:
                            fmt.Println(v2, v55, v62)
                        }
                        _ = fi63
                    }
                case "\n\t":
                    fmt.Println(v3, v1, v2)
                default:
                    for fi123 := 0; fi123 < 1; fi123++ {
                        var v124 int = int(v62)
                        v2 = v2
                        v55 = (v55 - v62)
                        _ = fi123
                        _ = v124
                    }
                }
                if false {
                    if (("\n\t" + ("" + "foo bar")) >= "") {
                        if true {
                            v125 := GenS14{F0: genFn0(), F1: (v3 * (v3 * v3))}
                            switch 0 {
                            case 127:
                                var v126 int32 = (v62 - v62)
                                _ = v126
                            case -2147483648:
                                fmt.Println(v55)
                            case 0:
                                var v127 int16 = int16(v62)
                                _ = v127
                            }
                            _ = v125
                        }
                        if true {
                            if (135 > 32767) {
                                switch "foo bar" {
                                case "foo bar":
                                    var v128 int64 = (v56 >> ((1) & 63))
                                    _ = v128
                                case "":
                                    if (float32(0.0) < 0.0) {
                                        switch -1 {
                                        case 172:
                                            v129 := GenS13{F0: ("" + "a"), F1: 2}
                                            _ = v129
                                        case 128:
                                            v3 = -1
                                        }
                                        var v130 int8 = 95
                                        _ = v130
                                    }
                                case "a":
                                    v131 := GenS15{F0: v3}
                                    _ = v131
                                }
                                for fi132 := 0; fi132 < 4; fi132++ {
                                    if true {
                                        fmt.Println(v55, v2, v56)
                                    }
                                    v133 := GenS14{F0: 151, F1: -45}
                                    _ = fi132
                                    _ = v133
                                }
                                fmt.Println(v62, v55, v55)
                            } else {
                                v2 = 2147483648
                                var v134 float32 = 0.1
                                v55 -= v55
                                _ = v134
                            }
                            v2 -= v2
                        } else {
                            switch ((("\n\t" + "hello") + ("" + ("" + "hello"))) + "") {
                            case "\n\t":
                                var v135 uint32 = v2
                                _ = v135
                            case "hello":
                                if false {
                                    if false {
                                        var v136 float32 = (v1 - float32((1e10 - float64((0.1 - (v1 / 1.0))))))
                                        var v137 uint32 = v2
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi138 := 0; fi138 < 4; fi138++ {
                                                if true {
                                                    if (!(((false || false) && (42 == 127)) || (true || false))) {
                                                        fmt.Println(v2)
                                                        v1 += (0.0 * (float32((0.0 - (0.1 + 1e10))) + ((((float32(1.5) * v1) + (v136 + v136)) - float32(3.14159)) * 0.0)))
                                                        var v139 int = genFn1((v55 * v62), 42)
                                                        _ = v139
                                                    }
                                                } else {
                                                    if (true && (v1 >= v1)) {
                                                        if len(v53) >= 2 {
                                                            v53 = v53[1:len(v53)]
                                                        }
                                                        fmt.Println(v55, v3, v137)
                                                        if len(v53) > 0 {
                                                            v53[0] = (fi138 / ((fi138) | 1))
                                                        }
                                                    }
                                                    if true {
                                                        for fi140 := 0; fi140 < 13; fi140++ {
                                                            if ((true && (v55 >= -2147483648)) && true) {
                                                                if true {
                                                                    var v141 uint32 = 65535
                                                                    v2 += (v137 ^ 206)
                                                                    delete(v4, "\n\t")
                                                                    switch fi140 {
                                                                    case 354:
                                                                        switch (fi138 << ((65535) & 63)) {
                                                                        case -527:
                                                                            if (!(4294967295 < 0)) {
                                                                                if (1 <= -83) {
                                                                                    var v142 string = "hello"
                                                                                    var v143 string = v142
                                                                                    var v144 int = (fi140 ^ genFn1((v62 | (v62 << ((2) & 63))), 255))
                                                                                    fmt.Println(v1, fi140)
                                                                                    v136 = (((v1 / v1) + float32(1.5)) + v1)
                                                                                    var v145 uint32 = v141
                                                                                    var v146 int32 = -2147483648
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                    _ = v144
                                                                                    _ = v145
                                                                                    _ = v146
                                                                                }
                                                                                switch -1 {
                                                                                case 127:
                                                                                    v56 += int64(v137)
                                                                                    fmt.Println(v55, v3, v3)
                                                                                    delete(v5, 78)
                                                                                case 32767:
                                                                                    fmt.Println(v3, v3)
                                                                                }
                                                                                var v147 int32 = v55
                                                                                _ = v147
                                                                            } else {
                                                                                var v148 int64 = -773
                                                                                v149 := v4["hello"]
                                                                                _ = v148
                                                                                _ = v149
                                                                            }
                                                                        case 28:
                                                                            if false {
                                                                                var v150 float64 = 0.1
                                                                                _ = v150
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                _ = len(v4)
                                                                                var v151 int16 = (v3 << ((42) & 63))
                                                                                _ = v151
                                                                            }
                                                                        }
                                                                    case -2147483648:
                                                                        v3 = (v3 >> ((255) & 63))
                                                                    case -1:
                                                                        if false {
                                                                            v137 -= v141
                                                                        } else {
                                                                            v152 := GenS16{F0: (v56 * -1)}
                                                                            if (v56 == v56) {
                                                                                if (0 <= 42) {
                                                                                    v2 = (v141 >> ((1) & 63))
                                                                                } else {
                                                                                    var v153 bool = false
                                                                                    _ = len(v53)
                                                                                    if (false || false) {
                                                                                        fmt.Println(v56, v153)
                                                                                        if len(v53) >= 2 {
                                                                                            v53 = v53[1:len(v53)]
                                                                                        }
                                                                                    }
                                                                                    _ = v153
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    v53 = append(v53, fi138)
                                                                                default:
                                                                                    v154 := GenS1{F0: (!false), F1: 0, F2: 128}
                                                                                    _ = v154
                                                                                }
                                                                            }
                                                                            _ = v152
                                                                        }
                                                                    }
                                                                    _ = v141
                                                                }
                                                                v4["a"] = 1
                                                                v155 := GenS17{F0: (-(v62)), F1: (fi140 % ((-9) | 1)), F2: 42}
                                                                _ = v155
                                                            } else {
                                                                v55 = (v55 >> ((525) & 63))
                                                                fmt.Println(v56)
                                                                if true {
                                                                    v137 = 144
                                                                } else {
                                                                    v137 -= v2
                                                                    if len(v53) >= 2 {
                                                                        v53 = v53[1:len(v53)]
                                                                    }
                                                                    v136 = 0.0
                                                                    var v156 uint16 = 255
                                                                    _ = v156
                                                                }
                                                            }
                                                            _ = fi140
                                                        }
                                                    } else {
                                                        for fi157 := 0; fi157 < 6; fi157++ {
                                                            v158 := GenS18{F0: 18446744073709551615, F1: 42, F2: 0}
                                                            var v159 int64 = (-(v56))
                                                            _ = fi157
                                                            _ = v158
                                                            _ = v159
                                                        }
                                                        v53 = append(v53, fi138)
                                                    }
                                                }
                                                var v160 uint32 = 1
                                                _ = fi138
                                                _ = v160
                                            }
                                        }
                                        _ = v136
                                        _ = v137
                                    }
                                    var v161 int16 = v3
                                    _ = v161
                                } else {
                                    v3 += 425
                                }
                            case "a":
                                v55 = v55
                            }
                        }
                        var v162 uint16 = 640
                        _ = v162
                    }
                    if (!(347 >= 25)) {
                        var v163 uint8 = 255
                        v55 = (v62 << ((974) & 63))
                        switch 32768 {
                        case 0:
                            var v164 int32 = (^(v62))
                            _ = v164
                        default:
                            if false {
                                if false {
                                    v165 := v4[""]
                                    v166 := GenS18{F0: 65535, F1: 42, F2: 120}
                                    _ = v165
                                    _ = v166
                                }
                            } else {
                                switch -90 {
                                case -557:
                                    var v167 uint = 1
                                    _ = v167
                                case 446:
                                    v2 -= v2
                                case 0:
                                    if (!true) {
                                        v163 += (v163 << ((18446744073709551615) & 63))
                                        switch -9223372036854775808 {
                                        case -510:
                                            fmt.Println(v1)
                                        case -9223372036854775808:
                                            fmt.Println(v56, v62)
                                        }
                                    } else {
                                        fmt.Println(v56, v3)
                                    }
                                default:
                                    v62 = (v62 >> ((18446744073709551615) & 63))
                                }
                                switch int(v163) {
                                case 42:
                                    var v168 uint32 = 4294967295
                                    _ = v168
                                }
                            }
                        }
                        _ = v163
                    } else {
                        var v169 float32 = 1.5
                        v62 += (v55 % ((2147483647) | 1))
                        v169 -= v169
                        _ = v169
                    }
                } else {
                    v1 += (((((v1 / (0.1 / -2.25)) - float32(0.1)) * v1) / v1) - v1)
                    if len(v53) > 0 {
                        v53[0] = genFn1((v62 | 211), 42)
                    }
                }
                _ = v62
            } else {
                var v170 float32 = v1
                var v171 float32 = (((v1 - v170) - (v170 - (1.5 / (0.0 - (3.14159 / (3.14159 + 1.5)))))) - 1e10)
                _ = v170
                _ = v171
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
