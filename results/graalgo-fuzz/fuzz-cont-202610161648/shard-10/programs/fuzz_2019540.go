package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint32; F2 uint }
type GenS11 struct { F0 int16; F1 uint16; F2 string }
type GenS12 struct { F0 int; F1 int16 }
type GenS13 struct { F0 int32; F1 string; F2 int64 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 string; F1 uint32; F2 string }
type GenS16 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int = 0
    var v4 int16 = -1
    v5 := GenS1{F0: true}
    fmt.Println(v4)
    var v6 float32 = ((1e10 - (1e10 + ((((float32((float64(0.0) / 1.5)) + 3.14159) / 0.1) - 1.5) - 1e10))) - 1.5)
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS2{F0: 0, F1: ("\n\t" + "a"), F2: ("foo bar" + ("foo bar" + "foo bar"))}
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            v3 = fi7
            _ = fi10
        }
        _ = fi7
    }
    v2[((("\n\t" + ("" + "")) + "a") + "foo bar")] = 2
    var v11 int32 = -1
    var v12 float64 = float64((v6 * ((1.5 / v6) - v6)))
    switch 988 {
    case 128:
        v13 := []int{v3, (v3 * v3), 32767, (v3 & 231)}
        _ = v13
    default:
        var v14 int16 = (v4 ^ v4)
        _ = v14
    }
    v15 := GenS3{F0: 0, F1: ""}
    v16 := GenS4{F0: 42, F1: 0, F2: 32767}
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            v11 = (v11 | v11)
            if (!false) {
                switch (^(v3)) {
                case 1:
                    fmt.Println(v4, v11, v4)
                default:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v19 := []int{}
                        switch "\n\t" {
                        case "a":
                            v1[2] = 2147483648
                            var v20 uint16 = 255
                            switch (^(fi17)) {
                            case 4294967295:
                                if (0 != 1) {
                                    v11 += v11
                                }
                            case 32767:
                                v12 = v12
                            }
                            _ = v20
                        case "":
                            var v21 uint64 = 0
                            _ = len(v19)
                            _ = v21
                        case "\n\t":
                            var v22 string = (("a" + "\n\t") + "")
                            _ = v22
                        default:
                            v12 = -2.25
                        }
                        v23 := GenS1{F0: (2 <= 728)}
                        _ = fi18
                        _ = v19
                        _ = v23
                    }
                }
                if true {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v6 = 1.5
                        _ = fi24
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        if (v11 != (v11 | v11)) {
                            fmt.Println(v6, v6)
                        }
                        fmt.Println(fi25, v12, fi25)
                        fmt.Println(fi25, v11, fi25)
                        _ = fi25
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        switch "a" {
                        case "a":
                            if (325 > 759) {
                                var v27 uint16 = 65535
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    fmt.Println(v4)
                                default:
                                    var v28 bool = (32767 == (v4 - (v4 - v4)))
                                    switch "a" {
                                    case "hello":
                                        var v29 int8 = 3
                                        _ = v29
                                    case "foo bar":
                                        var v30 bool = (false && v28)
                                        _ = v30
                                    }
                                    v31 := GenS4{F0: 443, F1: 42, F2: v27}
                                    switch -533 {
                                    case -2147483648:
                                        switch ("\n\t" + "") {
                                        case "hello":
                                            var v32 int64 = 767
                                            _ = v32
                                        case "a":
                                            if (!false) {
                                                v33 := GenS1{F0: v28}
                                                v34 := GenS5{F0: 255}
                                                fmt.Println(fi17, v4, v11)
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        switch fi17 {
                                        case 128:
                                            v35 := GenS3{F0: 32767, F1: ""}
                                            _ = v35
                                        default:
                                            var v36 float64 = (v12 - v12)
                                            _ = v36
                                        }
                                    case 4294967295:
                                        var v37 int32 = 150
                                        _ = v37
                                    default:
                                        v4 -= v4
                                    }
                                    _ = v28
                                    _ = v31
                                }
                                _ = v27
                            }
                        }
                        _ = fi26
                    }
                } else {
                    for fi38 := 0; fi38 < 10; fi38++ {
                        switch "\n\t" {
                        case "a":
                            v39 := GenS6{F0: false, F1: 0, F2: 2}
                            _ = v39
                        case "\n\t":
                            if (!("hello" <= "foo bar")) {
                                var v40 float64 = v12
                                var v41 uint8 = 0
                                _ = v40
                                _ = v41
                            }
                        default:
                            v4 = (v4 - (v4 / ((v4) | 1)))
                        }
                        if true {
                            v3 = v3
                            v42 := GenS4{F0: 1, F1: 0, F2: 399}
                            _ = v42
                        }
                        switch "a" {
                        case "":
                            v43 := []int{(fi38 ^ fi38), v3}
                            _ = v43
                        default:
                            var v44 uint32 = 42
                            _ = v44
                        }
                        _ = fi38
                    }
                    if (true && true) {
                        v45 := v1[563]
                        var v46 uint = 20
                        v47 := GenS7{F0: 2, F1: v46}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v48 uint32 = 65535
                        _ = v48
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            var v49 uint = 0
                            _ = v49
                        }
                    }
                }
            } else {
                if (v11 <= v11) {
                    v3 = fi17
                    switch "foo bar" {
                    case "":
                        for fi50 := 0; fi50 < 17; fi50++ {
                            switch fi17 {
                            case 4294967295:
                                v6 += v6
                            }
                            _ = len(v2)
                            switch (fi50 * v3) {
                            case 42:
                                v51 := GenS4{F0: 4294967295, F1: 54, F2: 32767}
                                _ = v51
                            default:
                                fmt.Println(v6, fi50, fi17)
                            }
                            delete(v1, 0)
                            var v52 uint8 = 128
                            if false {
                                v52 = (v52 | 127)
                            } else {
                                v4 += (-(v4))
                                var v53 int64 = -690
                                _ = v53
                            }
                            _ = fi50
                            _ = v52
                        }
                    case "a":
                        if true {
                            var v54 int = v3
                            _ = len(v1)
                            _ = v54
                        }
                    }
                    var v55 uint16 = 463
                    _ = v55
                } else {
                    v56 := GenS8{F0: v11}
                    _ = v56
                }
                if (852 != 186) {
                    v6 = v6
                    v6 += (1.5 - 1.5)
                    var v57 uint16 = 42
                    _ = v57
                } else {
                    fmt.Println(v12, v11)
                }
                for fi58 := 0; fi58 < 2; fi58++ {
                    v59 := v1[610]
                    v4 -= int16(v11)
                    _ = fi58
                    _ = v59
                }
            }
            _ = fi17
        }
    } else {
        v60 := v1[328]
        v4 = v4
        _ = v60
    }
    v11 -= (^(v11))
    if false {
        v12 -= v12
        if false {
            switch "\n\t" {
            case "\n\t":
                v4 = v4
            case "a":
                v1[2147483648] = 1
            }
            switch "\n\t" {
            case "\n\t":
                v6 -= ((v6 + 0.0) * (v6 - v6))
            case "hello":
                var v61 bool = false
                _ = v61
            default:
                if true {
                    _ = len(v1)
                } else {
                    fmt.Println(v11, v6)
                    switch "" {
                    case "hello":
                        switch 32768 {
                        case -1:
                            v62 := v1[1]
                            _ = v62
                        case 128:
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v3, v11)
                            case "a":
                                var v63 uint = 0
                                _ = v63
                            default:
                                v12 -= 0.0
                            }
                        default:
                            if true {
                                v4 += int16(v11)
                                if false {
                                    v64 := GenS9{F0: 2}
                                    _ = v64
                                }
                                switch 42 {
                                case 127:
                                    if false {
                                        v65 := GenS3{F0: 2, F1: "foo bar"}
                                        v66 := GenS5{F0: 0}
                                        _ = v65
                                        _ = v66
                                    }
                                default:
                                    var v67 bool = true
                                    _ = v67
                                }
                            }
                        }
                    case "\n\t":
                        v68 := []int{}
                        _ = v68
                    case "":
                        if (!(1 == 42)) {
                            for fi69 := 0; fi69 < 18; fi69++ {
                                var v70 uint8 = 128
                                v4 = (v4 + v4)
                                _ = fi69
                                _ = v70
                            }
                            fmt.Println(v3, v6, v4)
                        }
                    default:
                        _ = len(v2)
                    }
                    for fi71 := 0; fi71 < 3; fi71++ {
                        for fi72 := 0; fi72 < 5; fi72++ {
                            if true {
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v12)
                                    if (4294967295 != 82) {
                                        if (!(2 >= 127)) {
                                            switch ("a" + "\n\t") {
                                            case "\n\t":
                                                if ((4294967295 >= 4294967295) || ((-9223372036854775808 > -9223372036854775808) || true)) {
                                                    v12 += (v12 / v12)
                                                    fmt.Println(fi72, v4, v11)
                                                    if false {
                                                        v73 := []int{fi71, v3, (fi72 >> ((65535) & 63)), (v3 - -618)}
                                                        v3 += 128
                                                        _ = v73
                                                    }
                                                } else {
                                                    v11 = v11
                                                    var v74 uint16 = 101
                                                    _ = v74
                                                }
                                            }
                                            var v75 int = (fi72 | fi71)
                                            var v76 int16 = v4
                                            _ = v75
                                            _ = v76
                                        }
                                        if false {
                                            v77 := GenS9{F0: 1}
                                            fmt.Println(v4)
                                            _ = v77
                                        }
                                    } else {
                                        if (false || (!false)) {
                                            var v78 bool = (2 != 933)
                                            if v78 {
                                                v79 := GenS10{F0: 42, F1: 876, F2: 255}
                                                fmt.Println(v12)
                                                v11 = (v11 + (v11 / (((v11 + (v11 & v11))) | 1)))
                                                _ = v79
                                            }
                                            switch 128 {
                                            case -5:
                                                v3 = (fi72 >> ((65535) & 63))
                                                var v80 float32 = v6
                                                _ = v80
                                            default:
                                                if ((!(225 == (v4 << ((482) & 63)))) && true) {
                                                    var v81 uint64 = 0
                                                    _ = v81
                                                } else {
                                                    v12 -= 3.14159
                                                    v82 := GenS11{F0: 1, F1: 682, F2: "hello"}
                                                    var v83 uint32 = 54
                                                    _ = v82
                                                    _ = v83
                                                }
                                            }
                                            fmt.Println(v3, v4)
                                            if v78 {
                                                var v84 int32 = v11
                                                v12 = (float64(1e10) + 0.1)
                                                var v85 uint32 = 2147483647
                                                var v86 bool = ((v3 & fi71) > (fi71 | (fi71 & fi71)))
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            _ = v78
                                        }
                                    }
                                default:
                                    v87 := []int{}
                                    _ = v87
                                }
                                var v88 int32 = (v11 - (v11 | (^(v11))))
                                _ = v88
                            }
                            v4 = -69
                            if ((fi71 % (((fi71 * -361)) | 1)) > -2147483648) {
                                v89 := GenS9{F0: 2}
                                _ = v89
                            } else {
                                if (42 >= 1) {
                                    var v90 int32 = v11
                                    var v91 uint32 = 562
                                    _ = v90
                                    _ = v91
                                }
                            }
                            _ = fi72
                        }
                        _ = fi71
                    }
                }
            }
            fmt.Println(v4, v11)
        } else {
            if true {
                v4 += v4
                for fi92 := 0; fi92 < 3; fi92++ {
                    if true {
                        v93 := GenS12{F0: (^(v3)), F1: (-(v4))}
                        _ = v93
                    } else {
                        var v94 int16 = -1
                        fmt.Println(v12, v4)
                        for fi95 := 0; fi95 < 3; fi95++ {
                            if (-1 >= 89) {
                                if (false || (116 > 65535)) {
                                    if true {
                                        fmt.Println(fi95)
                                        switch "hello" {
                                        case "\n\t":
                                            if true {
                                                v96 := []int{9223372036854775807, (-(fi95))}
                                                fmt.Println(v3)
                                                if (42 >= 0) {
                                                    var v97 int32 = 0
                                                    v6 = v6
                                                    v98 := GenS4{F0: 2, F1: 0, F2: 513}
                                                    if true {
                                                        v94 -= (v4 ^ (v4 & (v4 / (((-(v94))) | 1))))
                                                        v99 := GenS13{F0: 1, F1: "foo bar", F2: 347}
                                                        _ = v99
                                                    } else {
                                                        var v100 float32 = v6
                                                        v12 += (v12 + 0.0)
                                                        _ = v100
                                                    }
                                                    delete(v2, "")
                                                    switch "a" {
                                                    case "hello":
                                                        v4 = v4
                                                    case "a":
                                                        var v101 float64 = v12
                                                        _ = v101
                                                    default:
                                                        var v102 int8 = -1
                                                        _ = v102
                                                    }
                                                    v12 -= v12
                                                    v11 = (v97 * 73)
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    v4 -= 0
                                                    fmt.Println(v6)
                                                    var v103 uint64 = 1
                                                    _ = v103
                                                }
                                                var v104 uint64 = 507
                                                _ = v104
                                                _ = v96
                                            }
                                        }
                                        delete(v2, ("hello" + ""))
                                    } else {
                                        var v105 int32 = (v11 + (v11 ^ (v11 << ((1) & 63))))
                                        switch v3 {
                                        case -37:
                                            var v106 int8 = -72
                                            _ = v106
                                        case 127:
                                            if true {
                                                v2["hello"] = 841
                                                switch ("foo bar" + (("" + "a") + "hello")) {
                                                case "foo bar":
                                                    var v107 uint16 = 2
                                                    _ = v107
                                                case "hello":
                                                    v108 := GenS6{F0: false, F1: 2, F2: 494}
                                                    _ = v108
                                                case "\n\t":
                                                    switch (fi92 >> ((569) & 63)) {
                                                    case 128:
                                                        if (false || false) {
                                                            var v109 uint64 = 675
                                                            _ = v109
                                                        }
                                                    case 84:
                                                        v110 := GenS14{F0: 9223372036854775807}
                                                        _ = v110
                                                    }
                                                }
                                            } else {
                                                v11 = (^(v11))
                                            }
                                        }
                                        switch (fi95 << ((4294967295) & 63)) {
                                        case 127:
                                            var v111 int32 = 2147483647
                                            _ = v111
                                        case 197:
                                            v11 = -2147483648
                                        case 32768:
                                            if false {
                                                var v112 float32 = v6
                                                v12 -= v12
                                                _ = v112
                                            }
                                        default:
                                            fmt.Println(v11)
                                        }
                                        _ = v105
                                    }
                                    var v113 float32 = v6
                                    switch 0 {
                                    case -2147483648:
                                        var v114 int = (-(v3))
                                        _ = v114
                                    case -1:
                                        v115 := GenS15{F0: "hello", F1: 2147483647, F2: ("hello" + "")}
                                        _ = v115
                                    case -103:
                                        v2[""] = 4294967295
                                    default:
                                        v116 := []int{(v3 * (fi92 % ((4294967295) | 1)))}
                                        _ = v116
                                    }
                                    _ = v113
                                } else {
                                    switch -466 {
                                    case 32767:
                                        fmt.Println(v11, v12)
                                    case -2147483648:
                                        v4 += -32768
                                    }
                                    var v117 int16 = (v94 ^ v94)
                                    if false {
                                        fmt.Println(v4, v94, fi92)
                                        v117 = (v4 ^ 32767)
                                    }
                                    v12 = float64(v6)
                                    v11 -= (v11 * v11)
                                    _ = v117
                                }
                                if true {
                                    if (true || ((v4 + (v4 / (((v94 - (-(v94)))) | 1))) != v94)) {
                                        v12 = -2.25
                                    } else {
                                        switch 2147483647 {
                                        case 32767:
                                            if (v11 > 992) {
                                                switch 127 {
                                                case 4294967295:
                                                    var v118 int = (-(v3))
                                                    _ = v118
                                                case 127:
                                                    v94 = (v4 + v4)
                                                case -391:
                                                    var v119 uint8 = 42
                                                    _ = v119
                                                }
                                                fmt.Println(v4, fi92, fi95)
                                            }
                                        case 127:
                                            v120 := []int{(fi95 | v3), (fi95 + (v3 * 1)), 32768}
                                            _ = v120
                                        default:
                                            if true {
                                                var v121 uint32 = 1
                                                fmt.Println(v12, v12)
                                                _ = v121
                                            }
                                        }
                                        v94 -= (v94 * (v94 ^ (v94 | v4)))
                                    }
                                    var v122 int16 = v4
                                    switch "a" {
                                    case "foo bar":
                                        v122 = (v122 - v122)
                                    case "a":
                                        fmt.Println(fi92, v4, v12)
                                    case "hello":
                                        fmt.Println(v12, v12)
                                        if false {
                                            if (539 <= 2) {
                                                v12 -= v12
                                            }
                                        }
                                    }
                                    _ = v122
                                }
                            } else {
                                v123 := GenS16{F0: 101, F1: (-(v4))}
                                var v124 int = (fi95 - v3)
                                _ = v123
                                _ = v124
                            }
                            var v125 bool = true
                            _ = fi95
                            _ = v125
                        }
                        _ = v94
                    }
                    _ = fi92
                }
                var v126 uint32 = 42
                _ = v126
            }
            v11 = -2147483648
            var v127 uint32 = 986
            _ = v127
        }
        var v128 float64 = 1e10
        _ = v128
    } else {
        var v129 uint8 = 115
        for fi130 := 0; fi130 < 5; fi130++ {
            v131 := []int{(fi130 - (fi130 * 0)), 42, v3, (^(v3))}
            var v132 uint64 = 42
            v133 := GenS16{F0: 987, F1: v4}
            _ = fi130
            _ = v131
            _ = v132
            _ = v133
        }
        _ = v129
    }
    v6 += 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
