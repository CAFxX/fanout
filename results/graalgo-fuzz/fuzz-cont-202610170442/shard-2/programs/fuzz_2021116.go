package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int16; F1 int32; F2 int64 }
type GenS9 struct { F0 uint8; F1 bool; F2 int64 }
type GenS10 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS14 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int8; F1 int8; F2 uint }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 uint8; F1 int8; F2 uint }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 int; F1 int32 }
type GenS25 struct { F0 string; F1 uint16; F2 string }
type GenS26 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8, p1 bool) uint {
    var v6 uint = 428
    if (!true) {
        for fi7 := 0; fi7 < 15; fi7++ {
            if p1 {
                if p1 {
                    var v8 int = fi7
                    if p1 {
                        switch (v8 * v8) {
                        case 127:
                            p0 = (p0 % (((p0 & p0)) | 1))
                        case 32767:
                            p1 = false
                        case 32768:
                            v6 = (v6 << (((v6 - v6)) & 63))
                        default:
                            if true {
                                var v9 uint16 = 42
                                for fi10 := 0; fi10 < 19; fi10++ {
                                    fmt.Println(v8)
                                    fmt.Println(v6)
                                    v9 = v9
                                    _ = fi10
                                }
                                var v11 int8 = 99
                                _ = v11
                                _ = v9
                            }
                        }
                    }
                    _ = v8
                }
                v6 -= (v6 + (v6 >> ((845) & 63)))
                v12 := GenS1{F0: 62, F1: 2147483648}
                _ = v12
            }
            _ = fi7
        }
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(v6, p0, v6)
            _ = fi14
        }
        if p1 {
            var v15 uint16 = 2
            _ = v15
        }
        if p1 {
            v16 := GenS2{F0: "hello", F1: 1}
            for fi17 := 0; fi17 < 19; fi17++ {
                if false {
                    if (v6 <= 18446744073709551615) {
                        v18 := GenS3{F0: 4294967295, F1: v6, F2: 42}
                        _ = v18
                    } else {
                        v19 := []int{-419, fi17, (fi17 >> ((v6) & 63))}
                        _ = v19
                    }
                    v20 := []int{}
                    _ = v20
                }
                v21 := []int{0, 32767, (fi13 >> ((v6) & 63)), (fi13 & (fi13 + (-(fi17))))}
                v22 := GenS3{F0: 4294967295, F1: (v6 / (((v6 % (((v6 * v6)) | 1))) | 1)), F2: 0}
                if (1 == 127) {
                    v23 := GenS1{F0: 121, F1: 4294967295}
                    _ = v23
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v24 bool = (!(28 != 801))
                        _ = v24
                    case "foo bar":
                        if len(v21) > 0 {
                            v21[0] = fi13
                        }
                    default:
                        if p1 {
                            v21 = append(v21, fi17)
                            if (false || (-9223372036854775808 >= 505)) {
                                v25 := GenS4{F0: 0, F1: -32768}
                                var v26 float32 = (3.14159 + 0.1)
                                v21 = append(v21, (-(fi17)))
                                _ = v25
                                _ = v26
                            } else {
                                var v27 string = "a"
                                _ = v27
                            }
                            v28 := GenS2{F0: "", F1: 32768}
                            _ = len(v21)
                            if len(v21) > 0 {
                                v21[0] = fi13
                            }
                            _ = v28
                        } else {
                            p0 -= p0
                        }
                        v29 := GenS5{F0: 1, F1: int(p0)}
                        if (-2.25 >= (3.14159 - 0.0)) {
                            p0 = (p0 >> (((v6 % (((v6 * (v6 % ((v6) | 1)))) | 1))) & 63))
                            if (42 <= 42) {
                                v6 = v6
                                v6 += v6
                                fmt.Println(v6)
                                v30 := GenS5{F0: 0, F1: fi13}
                                _ = v30
                            } else {
                                v31 := GenS4{F0: 37, F1: 0}
                                if false {
                                    fmt.Println(p0)
                                }
                                var v32 int = fi17
                                _ = v31
                                _ = v32
                            }
                            if (!((^(fi17)) <= (fi17 >> ((v6) & 63)))) {
                                fmt.Println(fi17, fi17)
                                if (!(!(p0 < p0))) {
                                    var v33 int = fi17
                                    _ = v33
                                }
                            } else {
                                if p1 {
                                    p1 = true
                                    switch fi17 {
                                    case 0:
                                        var v34 uint8 = 128
                                        _ = v34
                                    case -2147483648:
                                        p1 = (2 <= 0)
                                    default:
                                        v21 = append(v21, (fi17 % ((fi13) | 1)))
                                    }
                                }
                                switch (("" + ("a" + "\n\t")) + ("hello" + "\n\t")) {
                                case "\n\t":
                                    var v35 int8 = (p0 | -1)
                                    _ = v35
                                case "":
                                    var v36 string = (("foo bar" + "foo bar") + ("" + ("" + (("\n\t" + "") + "hello"))))
                                    _ = v36
                                case "hello":
                                    var v37 int64 = 42
                                    _ = v37
                                default:
                                    if p1 {
                                        if ((p1 && p1) && (1 <= 676)) {
                                            var v38 int = (fi13 - 32767)
                                            switch "hello" {
                                            case "\n\t":
                                                switch ("" + "foo bar") {
                                                case "a":
                                                    v39 := GenS6{F0: -2147483648}
                                                    _ = v39
                                                default:
                                                    switch -2147483648 {
                                                    case -125:
                                                        v6 -= 635
                                                    }
                                                    var v40 int8 = p0
                                                    _ = v40
                                                }
                                            default:
                                                var v41 uint = v6
                                                _ = v41
                                            }
                                            if p1 {
                                                v38 -= (fi13 ^ v38)
                                            } else {
                                                v42 := GenS6{F0: int32(v6)}
                                                var v43 uint64 = 4294967295
                                                if ((v6 * (v6 * v6)) == (v6 - v6)) {
                                                    var v44 bool = false
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    _ = v44
                                                } else {
                                                    if len(v21) > 0 {
                                                        v21[0] = 9223372036854775807
                                                    }
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                            switch "a" {
                                            case "hello":
                                                if p1 {
                                                    switch "\n\t" {
                                                    case "":
                                                        v21 = append(v21, fi17)
                                                    }
                                                    v21 = append(v21, fi17)
                                                    v45 := GenS2{F0: "foo bar", F1: 931}
                                                    _ = v45
                                                }
                                            case "":
                                                var v46 uint = 782
                                                _ = v46
                                            }
                                            if (42 >= 1) {
                                                var v47 float32 = 0.1
                                                _ = v47
                                            }
                                            _ = v38
                                        } else {
                                            v48 := GenS2{F0: "hello", F1: 127}
                                            var v49 uint8 = 0
                                            _ = v48
                                            _ = v49
                                        }
                                        v6 += (v6 % ((618) | 1))
                                    } else {
                                        var v50 string = "\n\t"
                                        if (-98 != p0) {
                                            var v51 int64 = 0
                                            _ = v51
                                        }
                                        var v52 uint = v6
                                        var v53 string = (("a" + v50) + "a")
                                        _ = v50
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            }
                            p0 = (p0 | (^(p0)))
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                p1 = ("hello" >= ("\n\t" + "a"))
                            }
                        }
                        _ = v29
                    }
                    var v54 int64 = -1
                    p1 = (0 != 2147483647)
                    _ = v54
                }
                v6 += v6
                _ = fi17
                _ = v21
                _ = v22
            }
            var v55 float64 = (float64(3.14159) - 3.14159)
            _ = v16
            _ = v55
        } else {
            v56 := []int{}
            p1 = (!(!(v6 != (v6 * (v6 / ((18446744073709551615) | 1))))))
            v57 := GenS7{F0: true, F1: 81, F2: 128}
            _ = v56
            _ = v57
        }
        _ = fi13
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        p1 = p1
        _ = fi58
    }
    var v59 uint64 = 0
    _ = p0
    _ = p1
    _ = v59
    _ = v6
    return (v6 << ((v6) & 63))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    v3 := make(map[string]bool)
    var v4 int8 = -128
    v5 := make(map[string]bool)
    v60 := GenS6{F0: -1}
    v4 -= (v4 | (v4 >> ((genFn0((v4 % ((-128) | 1)), (2 != 42))) & 63)))
    var v61 uint8 = 128
    if true {
        for fi62 := 0; fi62 < 18; fi62++ {
            _ = len(v3)
            if (2147483647 >= 0) {
                v61 = (v61 >> ((4294967295) & 63))
                var v63 int16 = 42
                _ = v63
            } else {
                v64 := GenS8{F0: 1000, F1: 1, F2: 814}
                for fi65 := 0; fi65 < 10; fi65++ {
                    v61 -= (v61 << ((0) & 63))
                    _ = fi65
                }
                _ = v64
            }
            _ = fi62
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v61 += (^(v61))
        for fi66 := 0; fi66 < 6; fi66++ {
            if (false && false) {
                switch "a" {
                case "":
                    fmt.Println(v61, v61)
                case "hello":
                    v2 = append(v2, fi66)
                case "a":
                    switch (fi66 | (fi66 / ((-1) | 1))) {
                    case 699:
                        switch "foo bar" {
                        case "":
                            fmt.Println(fi66, v4)
                        case "a":
                            if (false && (100 != 42)) {
                                if (!true) {
                                    v4 = -1
                                    v67 := GenS9{F0: v61, F1: true, F2: -9223372036854775808}
                                    _ = v67
                                }
                                if true {
                                    for fi68 := 0; fi68 < 5; fi68++ {
                                        fmt.Println(v61, v61)
                                        switch fi68 {
                                        case -876:
                                            if (("foo bar" + "\n\t") == "") {
                                                v61 = 1
                                                if (!((v4 + v4) >= -1)) {
                                                    var v69 int32 = -1
                                                    var v70 int = (fi66 | (fi68 >> ((894) & 63)))
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v71 := v5[("\n\t" + "hello")]
                                                if true {
                                                    var v72 int16 = -32768
                                                    _ = v72
                                                }
                                                var v73 uint32 = 321
                                                _ = v71
                                                _ = v73
                                            } else {
                                                switch (fi66 >> ((genFn0((v4 % ((127) | 1)), (9223372036854775807 < 1))) & 63)) {
                                                case 0:
                                                    var v74 uint = 4294967295
                                                    _ = v74
                                                case 42:
                                                    var v75 int = fi66
                                                    _ = v75
                                                case 4294967295:
                                                    var v76 int32 = 0
                                                    _ = v76
                                                default:
                                                    if false {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "":
                                                                v77 := GenS2{F0: "\n\t", F1: fi68}
                                                                _ = v77
                                                            case "a":
                                                                if true {
                                                                    var v78 string = ("a" + "")
                                                                    v61 = v61
                                                                    _ = v78
                                                                }
                                                            }
                                                            var v79 int32 = -784
                                                            if true {
                                                                v79 = -1
                                                            } else {
                                                                v80 := GenS10{F0: 42, F1: 439, F2: v4}
                                                                fmt.Println(v61)
                                                                var v81 bool = (352 > -343)
                                                                v82 := GenS6{F0: -598}
                                                                v83 := GenS11{F0: 685}
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            _ = v79
                                                        }
                                                        v61 -= v61
                                                        if (187 == -2147483648) {
                                                            var v84 uint64 = 18446744073709551615
                                                            var v85 uint16 = 823
                                                            var v86 uint32 = 895
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 32767
                                                        }
                                                    }
                                                }
                                                v61 += v61
                                                v87 := v5[("\n\t" + "foo bar")]
                                                v88 := GenS8{F0: 981, F1: -1, F2: 590}
                                                if false {
                                                    v89 := GenS7{F0: v87, F1: 65535, F2: (v61 - v61)}
                                                    _ = v89
                                                }
                                                _ = v87
                                                _ = v88
                                            }
                                        case -445:
                                            fmt.Println(v61)
                                        case 561:
                                            _ = len(v2)
                                        }
                                        switch "a" {
                                        case "hello":
                                            if (383 > 18446744073709551615) {
                                                if (((0.0 * float32(1e10)) / 3.14159) == -2.25) {
                                                    fmt.Println(v4)
                                                    if true {
                                                        var v90 float32 = (-2.25 + 3.14159)
                                                        _ = v90
                                                    } else {
                                                        if (0.1 != 1e10) {
                                                            if true {
                                                                v91 := GenS12{F0: (!(!false))}
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi66
                                                                }
                                                                _ = v91
                                                            }
                                                            v92 := GenS1{F0: v61, F1: 2147483648}
                                                            v61 = (v61 ^ v61)
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                                fmt.Println(v4)
                                                v4 = v4
                                                var v93 uint32 = 866
                                                _ = v93
                                            } else {
                                                if (-32768 <= 32767) {
                                                    var v94 uint8 = v61
                                                    var v95 bool = (32767 <= -1)
                                                    _ = v94
                                                    _ = v95
                                                } else {
                                                    v96 := GenS8{F0: -1, F1: 785, F2: 9223372036854775807}
                                                    fmt.Println(v4, v61)
                                                    var v97 int8 = v4
                                                    _ = v96
                                                    _ = v97
                                                }
                                                v98 := GenS13{F0: 841, F1: (v61 >> ((8) & 63)), F2: -437}
                                                _ = v98
                                            }
                                        case "\n\t":
                                            switch fi68 {
                                            case 42:
                                                if ((v4 | 0) <= v4) {
                                                    var v99 uint8 = v61
                                                    switch "" {
                                                    case "a":
                                                        v2 = append(v2, 128)
                                                    case "\n\t":
                                                        if ((0.0 / 3.14159) >= (((0.1 * 1.5) * 0.1) + (3.14159 * -2.25))) {
                                                            _ = len(v2)
                                                        }
                                                    default:
                                                        var v100 int16 = 0
                                                        _ = v100
                                                    }
                                                    _ = v99
                                                }
                                            case 2147483647:
                                                fmt.Println(fi66)
                                            }
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                var v101 uint64 = 2
                                                _ = v101
                                            }
                                        }
                                        _ = fi68
                                    }
                                    var v102 int32 = 54
                                    _ = v102
                                } else {
                                    v61 = v61
                                }
                                v5["foo bar"] = (!false)
                            } else {
                                v103 := GenS10{F0: 32767, F1: 42, F2: 93}
                                switch ((("" + "hello") + "foo bar") + "foo bar") {
                                case "foo bar":
                                    var v104 int8 = v4
                                    _ = v104
                                case "\n\t":
                                    v4 = v4
                                }
                                delete(v1, 301)
                                _ = v103
                            }
                        case "hello":
                            for fi105 := 0; fi105 < 1; fi105++ {
                                switch "" {
                                case "hello":
                                    if ((false && ("hello" < "a")) || (1.5 == float64(((1.5 + -2.25) - 1.5)))) {
                                        v4 -= (v4 + (^(v4)))
                                    }
                                case "":
                                    v4 -= v4
                                default:
                                    var v106 uint8 = (v61 ^ (v61 >> ((257) & 63)))
                                    _ = v106
                                }
                                _ = fi105
                            }
                        }
                    default:
                        for fi107 := 0; fi107 < 16; fi107++ {
                            if true {
                                v4 = (v4 / ((14) | 1))
                                v108 := GenS14{F0: 0, F1: 0, F2: 244}
                                _ = v108
                            } else {
                                switch fi66 {
                                case 32768:
                                    if true {
                                        var v109 uint16 = 255
                                        _ = v109
                                    } else {
                                        v61 = (v61 - v61)
                                        v61 = (v61 | (v61 ^ 1))
                                        _ = len(v2)
                                        v4 = v4
                                    }
                                default:
                                    fmt.Println(v61, fi66)
                                }
                                v110 := GenS15{F0: 42}
                                _ = v110
                            }
                            var v111 uint8 = (v61 - (v61 + (^(v61))))
                            delete(v1, 631)
                            _ = fi107
                            _ = v111
                        }
                    }
                }
                for fi112 := 0; fi112 < 3; fi112++ {
                    v113 := GenS16{F0: v4}
                    switch "\n\t" {
                    case "foo bar":
                        var v114 int = (fi66 - fi66)
                        if (42 <= -2147483648) {
                            v61 -= v61
                            v114 += fi112
                        } else {
                            var v115 int16 = -32768
                            _ = v115
                        }
                        if (1 == 1) {
                            v4 = 42
                            fmt.Println(v114, fi66, fi112)
                            var v116 int64 = 9223372036854775807
                            _ = v116
                        } else {
                            var v117 float32 = -2.25
                            v4 += (v4 / (((v4 << ((65535) & 63))) | 1))
                            _ = v117
                        }
                        _ = v114
                    case "hello":
                        v118 := GenS13{F0: 2, F1: v61, F2: 746}
                        _ = v118
                    case "":
                        if (32767 > 32767) {
                            var v119 string = (((("hello" + "a") + "a") + "hello") + "")
                            v61 -= 138
                            v120 := GenS17{F0: true, F1: 4294967295}
                            _ = v119
                            _ = v120
                        } else {
                            var v121 uint32 = 2147483648
                            v2 = append(v2, fi112)
                            var v122 uint = 65535
                            _ = v121
                            _ = v122
                        }
                    }
                    _ = fi112
                    _ = v113
                }
                if len(v2) > 0 {
                    v2[0] = fi66
                }
            } else {
                var v123 uint64 = 235
                fmt.Println(v123, v61)
                _ = v123
            }
            _ = fi66
        }
    }
    var v124 uint8 = v61
    v125 := GenS12{F0: (true || (331 == 2))}
    for fi126 := 0; fi126 < 14; fi126++ {
        v4 += (v4 | 100)
        switch (fi126 % ((fi126) | 1)) {
        case 913:
            switch ("foo bar" + "\n\t") {
            case "a":
                if (-1 <= int32(v124)) {
                    switch (fi126 + (fi126 | 1)) {
                    case 0:
                        var v127 uint = 18446744073709551615
                        _ = v127
                    case -1:
                        v61 += (^(v124))
                    case 127:
                        if true {
                            if (-1 >= 743) {
                                var v128 float64 = 1e10
                                var v129 uint16 = 42
                                if (v129 <= v129) {
                                    v129 = (v129 * uint16(v61))
                                }
                                _ = v128
                                _ = v129
                            }
                            var v130 bool = true
                            _ = v130
                        }
                    }
                    var v131 int16 = -1
                    v61 = (v124 ^ 0)
                    _ = v131
                }
            case "\n\t":
                switch fi126 {
                case 32767:
                    if (true || true) {
                        switch "hello" {
                        case "foo bar":
                            v4 += (v4 % (((v4 << ((42) & 63))) | 1))
                        }
                        fmt.Println(v4, v124)
                        v132 := GenS18{F0: 934}
                        _ = v132
                    } else {
                        var v133 uint16 = uint16(v124)
                        _ = v133
                    }
                case -9223372036854775808:
                    if true {
                        fmt.Println(v4)
                        v134 := GenS19{F0: 127, F1: v4, F2: 1}
                        switch ("" + "hello") {
                        case "\n\t":
                            v135 := GenS15{F0: 240}
                            _ = v135
                        case "hello":
                            v136 := GenS6{F0: -1}
                            _ = v136
                        }
                        _ = v134
                    }
                case 9223372036854775807:
                    v4 -= (v4 << ((18446744073709551615) & 63))
                default:
                    v4 = 63
                }
            case "":
                switch "\n\t" {
                case "foo bar":
                    if (188 >= 1) {
                        for fi137 := 0; fi137 < 4; fi137++ {
                            switch "a" {
                            case "":
                                if (!((528 >= 42) || false)) {
                                    if (!true) {
                                        fmt.Println(v124)
                                        var v138 uint16 = 0
                                        if false {
                                            v139 := GenS20{F0: -184}
                                            _ = v139
                                        } else {
                                            fmt.Println(fi137, v138, v124)
                                            var v140 int64 = 9223372036854775807
                                            var v141 int16 = 1
                                            _ = v140
                                            _ = v141
                                        }
                                        _ = v138
                                    } else {
                                        v142 := GenS21{F0: 0}
                                        _ = v142
                                    }
                                }
                                v143 := GenS22{F0: 2, F1: (v4 >> ((0) & 63)), F2: 119}
                                _ = v143
                            case "foo bar":
                                delete(v1, 738)
                            }
                            v5["\n\t"] = true
                            v61 = (^(v124))
                            _ = fi137
                        }
                    } else {
                        for fi144 := 0; fi144 < 4; fi144++ {
                            if true {
                                if (false || (!(!("foo bar" == "a")))) {
                                    if ("hello" > "\n\t") {
                                        v124 -= (v61 | 128)
                                        var v145 uint = 73
                                        v4 += 9
                                        _ = v145
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v61 = v61
                                }
                                v61 += (v61 << ((65535) & 63))
                                v146 := GenS9{F0: 127, F1: true, F2: -931}
                                if ("\n\t" < "foo bar") {
                                    var v147 uint = genFn0((v4 - 106), false)
                                    _ = v147
                                } else {
                                    if (!true) {
                                        switch ("\n\t" + ("a" + "hello")) {
                                        case "":
                                            var v148 int64 = 0
                                            _ = v148
                                        case "\n\t":
                                            if (2147483648 > 2) {
                                                v124 -= (^(v61))
                                                v149 := GenS9{F0: 180, F1: false, F2: -1}
                                                var v150 int8 = (v4 ^ (v4 + (^(v4))))
                                                _ = v149
                                                _ = v150
                                            } else {
                                                if true {
                                                    var v151 bool = true
                                                    var v152 uint16 = 65535
                                                    _ = v151
                                                    _ = v152
                                                } else {
                                                    var v153 float32 = 0.1
                                                    v124 -= 127
                                                    v61 = v61
                                                    if false {
                                                        if (599 != int32(v124)) {
                                                            if true {
                                                                var v154 uint16 = 2
                                                                var v155 bool = true
                                                                _ = v154
                                                                _ = v155
                                                            }
                                                            v61 = v124
                                                        }
                                                        delete(v3, "\n\t")
                                                    }
                                                    _ = v153
                                                }
                                                v61 += v124
                                            }
                                        }
                                        var v156 int16 = 1000
                                        _ = v156
                                    } else {
                                        v61 = (^(v124))
                                    }
                                    v124 = v61
                                    v157 := v3[(((("a" + "foo bar") + ("hello" + ((("foo bar" + "\n\t") + "a") + "hello"))) + "\n\t") + "a")]
                                    if v157 {
                                        if (fi126 > (fi126 << ((42) & 63))) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v4 = (v4 * (v4 >> ((1) & 63)))
                                        }
                                        v157 = true
                                        v4 -= v4
                                        v61 = v61
                                        if (true && v157) {
                                            v158 := GenS16{F0: v4}
                                            _ = v158
                                        }
                                    }
                                    _ = v157
                                }
                                _ = v146
                            } else {
                                v159 := GenS23{F0: (^(v4))}
                                if (-32768 <= 1) {
                                    v124 -= v124
                                } else {
                                    var v160 uint8 = v61
                                    v161 := GenS8{F0: 1, F1: 2147483647, F2: 0}
                                    v162 := GenS24{F0: fi144, F1: 42}
                                    _ = v160
                                    _ = v161
                                    _ = v162
                                }
                                _ = v159
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi144
                        }
                        for fi163 := 0; fi163 < 15; fi163++ {
                            v124 = v124
                            var v164 int64 = -9223372036854775808
                            _ = fi163
                            _ = v164
                        }
                        if (true || true) {
                            v4 = v4
                            var v165 int64 = -9223372036854775808
                            var v166 float32 = -2.25
                            _ = v165
                            _ = v166
                        } else {
                            var v167 float32 = (1e10 * 0.0)
                            v61 = v124
                            _ = v167
                        }
                    }
                case "hello":
                    v124 -= v61
                case "":
                    v168 := GenS25{F0: ("hello" + "foo bar"), F1: 42, F2: "hello"}
                    _ = v168
                }
            }
        case 496:
            switch "foo bar" {
            case "":
                v61 -= v124
            case "hello":
                v124 -= v61
            }
        case 42:
            fmt.Println(v61)
        }
        _ = fi126
    }
    switch 32768 {
    case 42:
        v4 = 0
    case 2147483647:
        var v169 uint = uint(v61)
        _ = v169
    }
    v124 = (v61 | v61)
    var v170 uint = 2
    var v171 string = ""
    if false {
        v4 = v4
        if false {
            if false {
                var v172 uint8 = (v61 / ((v124) | 1))
                fmt.Println(v124, v4)
                if (128 == 127) {
                    switch v171 {
                    case "foo bar":
                        switch 971 {
                        case 900:
                            switch -23 {
                            case 4294967295:
                                if (((!(0 == 42)) || true) || false) {
                                    if true {
                                        var v173 uint32 = 42
                                        _ = v173
                                    }
                                    var v174 bool = true
                                    if v174 {
                                        switch (v171 + v171) {
                                        case "":
                                            var v175 float32 = 1e10
                                            _ = v175
                                        case "a":
                                            switch 484 {
                                            case -643:
                                                fmt.Println(v61, v124, v4)
                                            case 32767:
                                                v4 = 42
                                            default:
                                                var v176 int = -140
                                                _ = v176
                                            }
                                        case "foo bar":
                                            fmt.Println(v124, v124, v170)
                                        }
                                        switch v171 {
                                        case "hello":
                                            var v177 int = -9223372036854775808
                                            _ = v177
                                        }
                                        v61 = (v61 - (v61 | v172))
                                    } else {
                                        v171 = v171
                                        switch "\n\t" {
                                        case "hello":
                                            v178 := v3["\n\t"]
                                            _ = v178
                                        case "a":
                                            _ = len(v2)
                                        case "\n\t":
                                            v179 := GenS26{F0: uint64(v170), F1: 915}
                                            _ = v179
                                        default:
                                            v61 = (v124 & v61)
                                        }
                                        v174 = (887 > 0)
                                    }
                                    _ = v174
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = int(v61)
                                    }
                                    var v180 int16 = -119
                                    _ = v180
                                }
                            default:
                                var v181 int64 = -1
                                _ = v181
                            }
                        case 962:
                            v124 = v124
                        }
                    }
                    if ((0.0 - (float32(((((3.14159 + -2.25) / (3.14159 + 1.5)) - (0.1 * 1.5)) / 1.0)) + (1.5 - (-2.25 - -2.25)))) > 0.0) {
                        var v182 uint16 = 42
                        for fi183 := 0; fi183 < 17; fi183++ {
                            var v184 int32 = 42
                            _ = fi183
                            _ = v184
                        }
                        _ = v182
                    }
                }
                _ = v172
            } else {
                for fi185 := 0; fi185 < 15; fi185++ {
                    var v186 float32 = 3.14159
                    var v187 int16 = 42
                    _ = fi185
                    _ = v186
                    _ = v187
                }
                var v188 uint32 = 871
                _ = v188
            }
        }
        v124 = v124
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v124
    _ = v125
    _ = v170
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
}
