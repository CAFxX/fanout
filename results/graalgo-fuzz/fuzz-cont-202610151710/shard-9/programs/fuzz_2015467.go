package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int8; F1 int; F2 uint64 }
type GenS17 struct { F0 int16; F1 int64; F2 uint16 }
type GenS18 struct { F0 uint32; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int16 = 1
    if ((true && true) || ("foo bar" != "")) {
        for fi4 := 0; fi4 < 2; fi4++ {
            fmt.Println(fi4)
            switch fi4 {
            case 549:
                if (-2.25 == -2.25) {
                    fmt.Println(fi4, fi4)
                    var v5 int16 = (v3 ^ (^(v3)))
                    _ = v5
                } else {
                    if false {
                        for fi6 := 0; fi6 < 1; fi6++ {
                            var v7 uint8 = 128
                            fmt.Println(v3)
                            _ = fi6
                            _ = v7
                        }
                        for fi8 := 0; fi8 < 15; fi8++ {
                            v3 = 32767
                            _ = fi8
                        }
                    } else {
                        var v9 float64 = (1e10 - 0.0)
                        var v10 int32 = 42
                        _ = v10
                        _ = v9
                    }
                    switch (fi4 | (^(fi4))) {
                    case 1:
                        if true {
                            v3 -= int16(fi4)
                        } else {
                            fmt.Println(v3)
                        }
                    case -2147483648:
                        var v11 int32 = 0
                        _ = v11
                    case 0:
                        v3 = 1000
                    default:
                        var v12 uint32 = 65535
                        _ = v12
                    }
                }
            case 32768:
                switch 1 {
                case 9223372036854775807:
                    v3 -= v3
                case 32768:
                    if false {
                        var v13 int8 = 127
                        v14 := GenS1{F0: ("foo bar" + "a"), F1: -377, F2: 65535}
                        _ = v13
                        _ = v14
                    }
                default:
                    if ((!true) || (-1 < 0)) {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v3 += int16(fi4)
                            _ = fi15
                        }
                    } else {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            var v17 int = -1
                            if (fi16 <= fi16) {
                                switch fi4 {
                                case 42:
                                    var v18 int8 = 127
                                    _ = v18
                                case -910:
                                    fmt.Println(fi16, fi4)
                                }
                                var v19 uint = 715
                                _ = v19
                            }
                            var v20 int = fi4
                            _ = fi16
                            _ = v17
                            _ = v20
                        }
                        var v21 bool = false
                        _ = v21
                    }
                }
            case 32767:
                switch (fi4 | (fi4 * 32768)) {
                case 127:
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v3 += (v3 - (v3 - v3))
                        fmt.Println(fi4, v3, fi4)
                        v23 := GenS1{F0: ((("" + "a") + ((("\n\t" + "a") + "a") + "\n\t")) + "a"), F1: 9223372036854775807, F2: 0}
                        fmt.Println(fi22, v3)
                        var v24 uint32 = 0
                        v24 = v24
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case 999:
                    var v25 int32 = 42
                    _ = v25
                case 9223372036854775807:
                    if (65535 > 502) {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            if false {
                                if (!(!false)) {
                                    fmt.Println(v3, fi4, fi26)
                                    v3 -= v3
                                    v3 += v3
                                }
                                fmt.Println(fi4, fi4, fi4)
                            }
                            _ = fi26
                        }
                    } else {
                        if ("foo bar" == "a") {
                            for fi27 := 0; fi27 < 15; fi27++ {
                                if ((true || (("hello" < ((("\n\t" + ("hello" + ("a" + "a"))) + "") + "a")) && (!true))) || false) {
                                    v3 = (v3 % ((32767) | 1))
                                    var v28 uint = 929
                                    _ = v28
                                }
                                _ = fi27
                            }
                        }
                        switch fi4 {
                        case -53:
                            for fi29 := 0; fi29 < 13; fi29++ {
                                switch "foo bar" {
                                case "\n\t":
                                    if false {
                                        v1[2147483647] = 4294967295
                                        fmt.Println(v3, fi29, fi4)
                                    }
                                }
                                if (((0 != -783) && false) || false) {
                                    v3 += (v3 >> ((864) & 63))
                                    switch "hello" {
                                    case "foo bar":
                                        v30 := GenS1{F0: (("\n\t" + "hello") + ((("hello" + ("foo bar" + "a")) + (("a" + "") + "\n\t")) + "a")), F1: 9223372036854775807, F2: 1}
                                        _ = v30
                                    case "":
                                        v3 = (v3 * (v3 << ((422) & 63)))
                                    case "hello":
                                        v3 = v3
                                    default:
                                        fmt.Println(v3, fi29, v3)
                                    }
                                    if (1 == -1) {
                                        fmt.Println(v3, fi4)
                                        if true {
                                            var v31 float64 = float64(3.14159)
                                            var v32 string = ("hello" + "hello")
                                            _ = v31
                                            _ = v32
                                        }
                                    } else {
                                        v3 = v3
                                        var v33 bool = false
                                        v3 = (v3 | int16(fi4))
                                        _ = v33
                                    }
                                    switch 9223372036854775807 {
                                    case 32767:
                                        if true {
                                            v3 = (v3 ^ -1)
                                            var v34 int = int(v3)
                                            switch 191 {
                                            case -9223372036854775808:
                                                fmt.Println(v34, fi4)
                                                v1[(fi4 | (fi29 | fi29))] = 4294967295
                                                var v35 int32 = 2147483647
                                                _ = v35
                                            case 0:
                                                if (fi29 > 42) {
                                                    v36 := GenS1{F0: ("hello" + ("foo bar" + ("a" + "\n\t"))), F1: 1, F2: 255}
                                                    v3 -= v3
                                                    if false {
                                                        if ((v3 ^ v3) == (v3 / ((v3) | 1))) {
                                                            v34 = (fi4 + 4294967295)
                                                        } else {
                                                            v34 = v34
                                                        }
                                                        v37 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: 42}
                                                        var v38 int64 = -1
                                                        switch (fi29 / (((v34 * (fi29 % ((0) | 1)))) | 1)) {
                                                        case 2147483647:
                                                            switch -9223372036854775808 {
                                                            case 2147483647:
                                                                v3 = 1
                                                                var v39 int16 = 32767
                                                                v3 += 630
                                                                _ = v39
                                                            case -612:
                                                                v38 -= (v38 * v38)
                                                            case 127:
                                                                v40 := GenS1{F0: "hello", F1: (v38 / (((v38 & (v38 & (v38 - v38)))) | 1)), F2: 4294967295}
                                                                _ = v40
                                                            }
                                                        case 0:
                                                            switch -336 {
                                                            case -986:
                                                                v41 := GenS1{F0: ("hello" + ("\n\t" + "foo bar")), F1: -9223372036854775808, F2: 209}
                                                                _ = v41
                                                            case 42:
                                                                v42 := GenS1{F0: (("foo bar" + "a") + "\n\t"), F1: v38, F2: 372}
                                                                _ = v42
                                                            case -2147483648:
                                                                v34 -= (fi29 & v34)
                                                            default:
                                                                if false {
                                                                    _ = len(v1)
                                                                    var v43 int8 = 14
                                                                    _ = len(v2)
                                                                    _ = v43
                                                                }
                                                            }
                                                        case -2147483648:
                                                            var v44 uint16 = 42
                                                            _ = v44
                                                        }
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    _ = v36
                                                }
                                            case 128:
                                                if (3.14159 > 1.5) {
                                                    var v45 uint32 = 2
                                                    var v46 int16 = v3
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    var v47 uint64 = 2
                                                    v48 := GenS1{F0: ("\n\t" + "a"), F1: -1, F2: 18446744073709551615}
                                                    v34 = (v34 | -1)
                                                    var v49 bool = false
                                                    v50 := []int{(^(v34))}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                            default:
                                                v51 := []int{}
                                                _ = v51
                                            }
                                            _ = v34
                                        } else {
                                            var v52 int16 = (^(v3))
                                            var v53 uint = 676
                                            v54 := GenS2{F0: 181, F1: v52}
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            if (856 < 798) {
                                                if (-128 >= -1) {
                                                    v3 = -75
                                                    if false {
                                                        fmt.Println(fi4)
                                                    }
                                                    v3 -= (v3 & v3)
                                                } else {
                                                    v3 = -314
                                                    if (381 == 65535) {
                                                        var v55 uint64 = 1
                                                        var v56 uint16 = 422
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        if false {
                                                            v3 = (v3 & 1)
                                                            v57 := GenS2{F0: 65535, F1: 1000}
                                                            var v58 uint16 = 65535
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            fmt.Println(v3, v3, fi4)
                                                            fmt.Println(fi4, v3)
                                                            v3 -= (v3 << ((349) & 63))
                                                        }
                                                        fmt.Println(v3, fi4, fi29)
                                                    }
                                                    var v59 int = (fi29 * 689)
                                                    var v60 uint16 = 32767
                                                    if true {
                                                        v61 := GenS3{F0: 100, F1: -1, F2: "\n\t"}
                                                        v59 = v59
                                                        v62 := GenS4{F0: 127, F1: 1}
                                                        switch "hello" {
                                                        case "hello":
                                                            if true {
                                                                v3 = (v3 >> ((4294967295) & 63))
                                                                v63 := v1[(^(v59))]
                                                                _ = v63
                                                            }
                                                        }
                                                        if (true && (846 <= -1)) {
                                                            v64 := []int{-562, (fi4 << ((1) & 63)), v59}
                                                            if (v60 > (v60 + (^(v60)))) {
                                                                v65 := GenS4{F0: -40, F1: 626}
                                                                v3 -= 42
                                                                fmt.Println(v60, v60)
                                                                v60 = (v60 >> ((4294967295) & 63))
                                                                _ = v65
                                                            }
                                                            var v66 int8 = 1
                                                            _ = v64
                                                            _ = v66
                                                        }
                                                        v60 = 597
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    var v67 uint8 = 128
                                                    _ = v59
                                                    _ = v60
                                                    _ = v67
                                                }
                                            } else {
                                                if false {
                                                    switch 531 {
                                                    case 20:
                                                        v68 := []int{fi4, 22, fi29, 0}
                                                        _ = v68
                                                    }
                                                } else {
                                                    var v69 uint64 = 18446744073709551615
                                                    _ = v69
                                                }
                                            }
                                        case "foo bar":
                                            v3 -= 42
                                        }
                                        v70 := GenS2{F0: 42, F1: v3}
                                        _ = v70
                                    default:
                                        var v71 int32 = 2147483647
                                        _ = v71
                                    }
                                    _ = len(v1)
                                    v72 := GenS5{F0: 118}
                                    var v73 uint8 = 127
                                    _ = v72
                                    _ = v73
                                } else {
                                    switch ("a" + "") {
                                    case "hello":
                                        var v74 uint = 737
                                        _ = v74
                                    case "foo bar":
                                        v3 = (v3 + -641)
                                    case "":
                                        v3 += (^(v3))
                                    default:
                                        if false {
                                            switch int(v3) {
                                            case -450:
                                                v75 := []int{}
                                                _ = v75
                                            case -2147483648:
                                                v76 := v1[(fi4 * fi4)]
                                                if (0 == 109) {
                                                    var v77 uint = 1
                                                    v3 -= 1000
                                                    _ = v77
                                                } else {
                                                    var v78 int16 = (-(v3))
                                                    _ = v78
                                                }
                                                var v79 uint16 = 0
                                                _ = v76
                                                _ = v79
                                            case 42:
                                                switch "hello" {
                                                case "a":
                                                    switch 390 {
                                                    case 128:
                                                        var v80 float64 = (0.1 / 3.14159)
                                                        _ = v80
                                                    }
                                                default:
                                                    v3 += 678
                                                }
                                            default:
                                                fmt.Println(fi29, fi29)
                                            }
                                            if (!(!true)) {
                                                fmt.Println(fi4)
                                            } else {
                                                var v81 uint8 = 255
                                                _ = v81
                                            }
                                        }
                                    }
                                }
                                var v82 float64 = 0.1
                                var v83 int16 = 0
                                v84 := []int{}
                                v3 -= 1
                                _ = fi29
                                _ = v82
                                _ = v83
                                _ = v84
                            }
                        case 42:
                            v3 = 1000
                        case 0:
                            for fi85 := 0; fi85 < 12; fi85++ {
                                switch (fi85 % ((2147483647) | 1)) {
                                case 4294967295:
                                    v3 = (v3 / ((1000) | 1))
                                }
                                delete(v2, 2147483648)
                                _ = fi85
                            }
                        default:
                            v86 := GenS1{F0: "", F1: 9223372036854775807, F2: 465}
                            _ = v86
                        }
                    }
                }
            default:
                for fi87 := 0; fi87 < 9; fi87++ {
                    v88 := GenS6{F0: 127, F1: 2147483647}
                    _ = fi87
                    _ = v88
                }
            }
            _ = fi4
        }
    }
    if true {
        v3 -= 1
        v89 := GenS7{F0: false, F1: -128}
        for fi90 := 0; fi90 < 5; fi90++ {
            var v91 uint = 1
            var v92 uint8 = 129
            _ = fi90
            _ = v91
            _ = v92
        }
        _ = v89
    } else {
        for fi93 := 0; fi93 < 17; fi93++ {
            v3 = -1
            v3 -= (-(v3))
            _ = fi93
        }
        var v94 int = 9223372036854775807
        var v95 uint32 = 4294967295
        _ = v94
        _ = v95
    }
    fmt.Println(v3)
    v3 += (v3 | v3)
    var v96 bool = true
    for fi97 := 0; fi97 < 20; fi97++ {
        fmt.Println(v96)
        for fi98 := 0; fi98 < 17; fi98++ {
            fmt.Println(v3, fi97, v3)
            switch (fi97 & fi97) {
            case 128:
                var v99 uint64 = 1
                _ = v99
            }
            _ = fi98
        }
        switch (fi97 / ((fi97) | 1)) {
        case -1:
            delete(v1, (fi97 / ((fi97) | 1)))
        case 128:
            v100 := GenS8{F0: 0}
            _ = v100
        }
        _ = fi97
    }
    v3 = v3
    v101 := GenS9{F0: 34}
    if true {
        for fi102 := 0; fi102 < 4; fi102++ {
            v103 := GenS1{F0: ("hello" + ("hello" + "\n\t")), F1: 1, F2: 169}
            _ = fi102
            _ = v103
        }
    } else {
        v104 := GenS10{F0: 1, F1: 212}
        var v105 float32 = (((0.0 / 3.14159) - 0.1) + -2.25)
        for fi106 := 0; fi106 < 1; fi106++ {
            v107 := GenS11{F0: v3}
            switch "\n\t" {
            case "\n\t":
                v96 = (42 <= 141)
            case "foo bar":
                if v96 {
                    v108 := GenS12{F0: ((v3 < 1000) && v96)}
                    _ = v108
                } else {
                    v109 := GenS13{F0: 255, F1: (-(v3))}
                    v3 -= (v3 >> ((18446744073709551615) & 63))
                    _ = v109
                }
            }
            v110 := []int{-2147483648}
            _ = fi106
            _ = v107
            _ = v110
        }
        _ = v104
        _ = v105
    }
    switch -716 {
    case -2147483648:
        for fi111 := 0; fi111 < 16; fi111++ {
            var v112 int32 = 293
            var v113 uint32 = 844
            v3 = 591
            _ = fi111
            _ = v112
            _ = v113
        }
    case 656:
        v3 -= v3
    case -1:
        switch "\n\t" {
        case "\n\t":
            for fi114 := 0; fi114 < 8; fi114++ {
                if (2 == 0) {
                    if true {
                        v115 := GenS4{F0: 127, F1: 18446744073709551615}
                        _ = v115
                    }
                    if false {
                        fmt.Println(fi114, v3, v3)
                        fmt.Println(v96)
                        v116 := []int{fi114}
                        _ = v116
                    }
                    for fi117 := 0; fi117 < 5; fi117++ {
                        if v96 {
                            if v96 {
                                if v96 {
                                    switch "a" {
                                    case "\n\t":
                                        if (("" + "\n\t") < ("\n\t" + "a")) {
                                            var v118 uint64 = 891
                                            fmt.Println(v96, v118)
                                            v3 += (v3 - (v3 * (v3 * v3)))
                                            _ = v118
                                        } else {
                                            switch (fi117 + (fi114 & 4294967295)) {
                                            case 128:
                                                v96 = v96
                                            }
                                            v119 := GenS10{F0: 42, F1: 127}
                                            _ = v119
                                        }
                                    case "foo bar":
                                        switch 2147483647 {
                                        case 1:
                                            fmt.Println(v3, v96)
                                        }
                                    }
                                    v120 := GenS14{F0: 673, F1: fi114}
                                    v121 := GenS15{F0: 24}
                                    _ = v120
                                    _ = v121
                                } else {
                                    v122 := GenS15{F0: 35}
                                    switch ("hello" + ("foo bar" + ("foo bar" + "hello"))) {
                                    case "a":
                                        var v123 int = 4294967295
                                        _ = v123
                                    case "foo bar":
                                        v124 := GenS16{F0: -1, F1: -275, F2: 214}
                                        _ = v124
                                    case "hello":
                                        v125 := GenS15{F0: 221}
                                        _ = v125
                                    default:
                                        v96 = v96
                                    }
                                    _ = v122
                                }
                                var v126 uint32 = 4294967295
                                v126 -= v126
                                _ = v126
                            }
                            fmt.Println(v96, v3)
                            switch "hello" {
                            case "\n\t":
                                if true {
                                    var v127 uint64 = 0
                                    _ = v127
                                } else {
                                    if v96 {
                                        fmt.Println(v3)
                                        var v128 uint16 = 42
                                        _ = v128
                                    } else {
                                        var v129 float32 = float32(0.0)
                                        switch "a" {
                                        case "a":
                                            if (!v96) {
                                                var v130 int = fi114
                                                var v131 int16 = v3
                                                _ = v130
                                                _ = v131
                                            } else {
                                                var v132 int32 = 0
                                                _ = v132
                                            }
                                        case "foo bar":
                                            v3 += v3
                                        case "hello":
                                            if (4294967295 > 127) {
                                                fmt.Println(fi117, fi117)
                                                var v133 uint = 65535
                                                v134 := GenS10{F0: 0, F1: 127}
                                                _ = v133
                                                _ = v134
                                            }
                                        }
                                        _ = v129
                                    }
                                }
                            }
                        } else {
                            var v135 float64 = -2.25
                            _ = v135
                        }
                        _ = fi117
                    }
                }
                switch ("hello" + "a") {
                case "hello":
                    var v136 uint16 = 32767
                    _ = v136
                case "":
                    v96 = (2 != 2)
                }
                _ = len(v1)
                _ = fi114
            }
        }
    default:
        fmt.Println(v3)
    }
    v96 = (((3.14159 - 0.0) + float32((0.1 * 1e10))) != (1e10 * (3.14159 - (1.5 / 1e10))))
    v96 = true
    var v137 uint = 65535
    v138 := GenS17{F0: 1, F1: 42, F2: 32767}
    switch 4294967295 {
    case 32768:
        if true {
            fmt.Println(v3, v137, v3)
            if (!v96) {
                for fi139 := 0; fi139 < 18; fi139++ {
                    v96 = ((220 <= 4294967295) && false)
                    v96 = (747 > 2)
                    _ = fi139
                }
                var v140 uint = 255
                v141 := GenS18{F0: 992, F1: 68, F2: v3}
                _ = v140
                _ = v141
            } else {
                v142 := GenS6{F0: -46, F1: -112}
                var v143 int16 = (v3 >> ((v137) & 63))
                v144 := GenS17{F0: (v3 << ((65535) & 63)), F1: -750, F2: 32767}
                _ = v142
                _ = v143
                _ = v144
            }
        }
    default:
        v3 += (v3 << (((v137 ^ (v137 & (v137 << ((v137) & 63))))) & 63))
    }
    var v145 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v137:", v137)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v101
    _ = v137
    _ = v138
    _ = v145
    _ = v2
    _ = v3
    _ = v96
}
