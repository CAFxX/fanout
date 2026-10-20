package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int; F1 uint8; F2 uint8 }
type GenS8 struct { F0 bool; F1 string; F2 uint16 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS14 struct { F0 uint8; F1 bool; F2 bool }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string; F1 uint64 }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int; F1 int64 }
type GenS20 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 int16; F1 uint8 }
type GenS24 struct { F0 string; F1 uint; F2 uint32 }
type GenS25 struct { F0 uint; F1 uint64 }
type GenS26 struct { F0 int8; F1 uint }
func genFn0(p0 uint8) int64 {
    var v7 uint64 = 630
    for fi8 := 0; fi8 < 1; fi8++ {
        switch int(p0) {
        case 9223372036854775807:
            p0 = p0
        default:
            for fi9 := 0; fi9 < 13; fi9++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi9, p0)
                }
                _ = fi9
            }
        }
        _ = fi8
    }
    v10 := GenS1{F0: 42}
    if (false || (true && true)) {
        fmt.Println(v7, v7)
    } else {
        switch ("hello" + ("hello" + "\n\t")) {
        case "hello":
            var v11 uint16 = 255
            _ = v11
        case "foo bar":
            for fi12 := 0; fi12 < 14; fi12++ {
                switch fi12 {
                case 9223372036854775807:
                    var v13 float32 = 3.14159
                    _ = v13
                case -518:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        switch "a" {
                        case "foo bar":
                            switch (fi12 << ((42) & 63)) {
                            case 127:
                                v15 := GenS1{F0: fi12}
                                fmt.Println(fi12, fi14, p0)
                                switch "hello" {
                                case "hello":
                                    var v16 uint16 = 392
                                    v16 = (v16 - (^(v16)))
                                    switch -1 {
                                    case 0:
                                        var v17 bool = true
                                        _ = v17
                                    case 1:
                                        v18 := []int{fi14, 32768}
                                        if (9223372036854775807 <= -1) {
                                            if ((fi14 - int(v7)) >= fi12) {
                                                v16 = (v16 + 255)
                                                p0 -= (p0 / (((p0 + 228)) | 1))
                                                v19 := GenS1{F0: int(v16)}
                                                v7 = (v7 + 2)
                                                var v20 uint = 4294967295
                                                _ = v19
                                                _ = v20
                                            }
                                            var v21 int = int(v7)
                                            _ = v21
                                        } else {
                                            v7 -= (v7 % (((^(v7))) | 1))
                                            var v22 int32 = -1
                                            var v23 uint = 678
                                            _ = v22
                                            _ = v23
                                        }
                                        if len(v18) > 0 {
                                            v18[0] = -1
                                        }
                                        _ = v18
                                    case -9223372036854775808:
                                        var v24 bool = false
                                        _ = v24
                                    }
                                    _ = v16
                                }
                                _ = v15
                            case 128:
                                var v25 int64 = -9223372036854775808
                                _ = v25
                            }
                            var v26 int8 = -62
                            v27 := []int{fi12, -96, fi14}
                            _ = v26
                            _ = v27
                        }
                        v28 := []int{130, (fi12 & 327), fi14, (fi14 >> ((4294967295) & 63))}
                        _ = fi14
                        _ = v28
                    }
                default:
                    for fi29 := 0; fi29 < 8; fi29++ {
                        if true {
                            v30 := []int{(fi12 + 128), fi29}
                            v7 = v7
                            fmt.Println(p0, fi29)
                            v31 := GenS2{F0: true, F1: 127, F2: (v7 | v7)}
                            _ = v30
                            _ = v31
                        } else {
                            var v32 string = ("\n\t" + ("a" + ("hello" + (("hello" + ("a" + ("hello" + "foo bar"))) + ""))))
                            var v33 int16 = -1
                            _ = v32
                            _ = v33
                        }
                        p0 = p0
                        var v34 int = fi29
                        _ = fi29
                        _ = v34
                    }
                }
                v7 -= (v7 >> ((18446744073709551615) & 63))
                _ = fi12
            }
        default:
            v35 := []int{-937, 712, 1}
            _ = v35
        }
        fmt.Println(v7)
        for fi36 := 0; fi36 < 1; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                fmt.Println(fi37, p0)
                var v38 uint32 = 65535
                if (((9223372036854775807 < 255) || ((false && (621 != 4294967295)) || (18446744073709551615 > v7))) && true) {
                    var v39 uint32 = v38
                    if true {
                        var v40 int8 = 42
                        _ = v40
                    }
                    v41 := []int{(fi37 + fi37), fi36, (fi36 >> ((778) & 63))}
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    if ("\n\t" < "\n\t") {
                        if len(v41) > 0 {
                            v41[0] = 4294967295
                        }
                    } else {
                        var v42 uint8 = p0
                        if false {
                            if len(v41) >= 2 {
                                v41 = v41[1:len(v41)]
                            }
                        } else {
                            var v43 bool = (("hello" + "foo bar") >= "foo bar")
                            v44 := GenS3{F0: -1}
                            v39 += v39
                            v45 := GenS2{F0: (((v39 * 2147483647) == 585) && (1.5 < 0.0)), F1: (^(p0)), F2: 0}
                            v46 := GenS4{F0: -90}
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = v42
                    }
                    v47 := GenS1{F0: fi37}
                    _ = v39
                    _ = v41
                    _ = v47
                }
                _ = fi37
                _ = v38
            }
            v48 := []int{fi36}
            _ = fi36
            _ = v48
        }
    }
    if (-1 <= -9223372036854775808) {
        if (!true) {
            var v49 string = ""
            for fi50 := 0; fi50 < 18; fi50++ {
                for fi51 := 0; fi51 < 20; fi51++ {
                    if true {
                        var v52 uint8 = p0
                        _ = v52
                    }
                    if true {
                        var v53 uint8 = p0
                        p0 = p0
                        v7 -= (^(v7))
                        _ = v53
                    }
                    _ = fi51
                }
                for fi54 := 0; fi54 < 4; fi54++ {
                    p0 += p0
                    _ = fi54
                }
                var v55 int16 = 1000
                _ = fi50
                _ = v55
            }
            v49 = v49
            _ = v49
        } else {
            v56 := GenS1{F0: 32768}
            if ("foo bar" <= "\n\t") {
                for fi57 := 0; fi57 < 4; fi57++ {
                    if false {
                        var v58 float32 = ((0.0 + (1e10 - 0.0)) * (1e10 + 3.14159))
                        v7 = 0
                        for fi59 := 0; fi59 < 20; fi59++ {
                            v7 -= 1
                            _ = fi59
                        }
                        _ = v58
                    }
                    _ = fi57
                }
                v60 := GenS5{F0: -2147483648}
                _ = v60
            } else {
                v61 := []int{int(v7), 963, -2147483648, 32767}
                _ = v61
            }
            v62 := GenS6{F0: 0, F1: 2}
            _ = v56
            _ = v62
        }
        for fi63 := 0; fi63 < 4; fi63++ {
            if (true && (!("hello" > "hello"))) {
                for fi64 := 0; fi64 < 19; fi64++ {
                    var v65 uint64 = (v7 >> ((42) & 63))
                    _ = fi64
                    _ = v65
                }
                p0 = 143
            }
            switch fi63 {
            case 32768:
                var v66 int32 = 246
                _ = v66
            default:
                var v67 string = ("hello" + "\n\t")
                _ = v67
            }
            switch "a" {
            case "hello":
                if true {
                    fmt.Println(v7, v7)
                    v7 = 42
                }
            case "foo bar":
                if true {
                    v68 := GenS7{F0: fi63, F1: p0, F2: (p0 ^ p0)}
                    v69 := GenS8{F0: false, F1: "hello", F2: 65535}
                    _ = v68
                    _ = v69
                }
            case "\n\t":
                var v70 uint16 = 1
                _ = v70
            }
            _ = fi63
        }
        var v71 int32 = 0
        _ = v71
    }
    _ = p0
    _ = v10
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{834, 2, 4294967295, 42}
    var v2 float64 = 0.1
    var v3 string = ("\n\t" + "foo bar")
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    var v6 uint32 = 2147483647
    v3 = v3
    v4[v6] = ((v2 * v2) != 0.0)
    for fi72 := 0; fi72 < 11; fi72++ {
        for fi73 := 0; fi73 < 18; fi73++ {
            if true {
                v6 = 641
                var v74 int = (fi73 | -1)
                if false {
                    switch v3 {
                    case "foo bar":
                        switch v3 {
                        case "hello":
                            v75 := GenS8{F0: (!((true || (v2 <= (((v2 + 0.0) / 1.5) / (v2 / (((v2 / v2) + v2) / v2))))) && (((!(!(false || (float32(0.1) != float32(v2))))) || false) || false))), F1: v3, F2: 448}
                            _ = v75
                        case "a":
                            fmt.Println(v6, fi73, fi72)
                        }
                    case "\n\t":
                        v76 := GenS2{F0: (!(!false)), F1: 229, F2: 1}
                        _ = v76
                    }
                    if (1e10 == v2) {
                        delete(v4, v6)
                        v74 -= fi73
                        var v77 int32 = -1
                        var v78 uint32 = v6
                        v79 := GenS5{F0: fi73}
                        _ = v77
                        _ = v78
                        _ = v79
                    } else {
                        var v80 uint64 = 2
                        if (18446744073709551615 >= v80) {
                            v81 := GenS6{F0: 1, F1: 0}
                            v6 = 65535
                            fmt.Println(fi73, v80)
                            _ = v81
                        }
                        _ = v80
                    }
                    if false {
                        var v82 float32 = (1e10 * (1e10 + 1.5))
                        _ = v82
                    }
                }
                _ = v74
            } else {
                var v83 string = ""
                if (1 < uint(v6)) {
                    if (!false) {
                        if (-128 >= -126) {
                            v84 := GenS4{F0: -1}
                            _ = v84
                        } else {
                            switch (v83 + v3) {
                            case "\n\t":
                                if (940 != 18446744073709551615) {
                                    v85 := GenS9{F0: -9223372036854775808, F1: 0}
                                    if (-1 >= -2147483648) {
                                        v83 = v3
                                        _ = len(v1)
                                        v86 := GenS10{F0: -2147483648, F1: false}
                                        var v87 uint8 = 255
                                        fmt.Println(fi72, fi73)
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v2 = v2
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch 214 {
                                        case 4294967295:
                                            switch fi72 {
                                            case -901:
                                                switch fi72 {
                                                case 141:
                                                    v1 = append(v1, (^(v6)))
                                                case 128:
                                                    if true {
                                                        v88 := GenS11{F0: 20, F1: v6}
                                                        _ = v88
                                                    }
                                                default:
                                                    var v89 int64 = 0
                                                    _ = v89
                                                }
                                            default:
                                                fmt.Println(v6, fi73)
                                            }
                                        default:
                                            v2 = v2
                                        }
                                        if false {
                                            if (0 < -72) {
                                                if ((0.0 <= 0.1) || false) {
                                                    _ = len(v1)
                                                    v83 = v3
                                                } else {
                                                    var v90 bool = true
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v90
                                                }
                                                if (-469 == -2147483648) {
                                                    fmt.Println(v3, v2)
                                                } else {
                                                    var v91 bool = true
                                                    switch ("" + (("foo bar" + (v83 + "\n\t")) + "hello")) {
                                                    case "":
                                                        switch v3 {
                                                        case "":
                                                            if (genFn0(198) <= -9223372036854775808) {
                                                                v91 = v91
                                                                v91 = (!v91)
                                                                var v92 uint8 = 1
                                                                switch "" {
                                                                case "hello":
                                                                    if v91 {
                                                                        var v93 uint16 = 697
                                                                        v92 -= v92
                                                                        _ = v93
                                                                    } else {
                                                                        fmt.Println(v91, v92)
                                                                        var v94 float64 = (v2 * (v2 - v2))
                                                                        _ = v94
                                                                    }
                                                                case "":
                                                                    var v95 uint16 = 65535
                                                                    _ = v95
                                                                default:
                                                                    v96 := GenS7{F0: -154, F1: 57, F2: 0}
                                                                    _ = v96
                                                                }
                                                                _ = v92
                                                            }
                                                        case "\n\t":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "foo bar":
                                                            v3 = ((v3 + v83) + ("" + v3))
                                                        }
                                                    case "foo bar":
                                                        switch fi72 {
                                                        case 32767:
                                                            fmt.Println(v3)
                                                        case 0:
                                                            if v91 {
                                                                v97 := GenS7{F0: fi72, F1: 160, F2: 209}
                                                                if (1 != 114) {
                                                                    var v98 uint32 = (^(v6))
                                                                    var v99 int64 = 832
                                                                    if v91 {
                                                                        if v91 {
                                                                            var v100 float64 = (3.14159 - (1.5 / float64(1.5)))
                                                                            _ = v100
                                                                        }
                                                                        switch fi73 {
                                                                        case 768:
                                                                            var v101 uint16 = 65535
                                                                            _ = v101
                                                                        case 4294967295:
                                                                            switch (fi72 | (fi72 - fi73)) {
                                                                            case 42:
                                                                                var v102 uint8 = 0
                                                                                _ = v102
                                                                            case 4294967295:
                                                                                v83 = (v3 + "foo bar")
                                                                            default:
                                                                                var v103 float64 = v2
                                                                                _ = v103
                                                                            }
                                                                        case -157:
                                                                            v5[(fi72 ^ fi73)] = true
                                                                            fmt.Println(v3, v3)
                                                                        }
                                                                        v1 = append(v1, 0)
                                                                    } else {
                                                                        fmt.Println(v99, v91)
                                                                        fmt.Println(v91)
                                                                    }
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                                switch (v83 + "a") {
                                                                case "":
                                                                    v83 = v83
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        switch ((((v83 + v3) + "") + "hello") + (v83 + v83)) {
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                if v91 {
                                                                                    delete(v4, 42)
                                                                                    v83 = v3
                                                                                } else {
                                                                                    switch 2147483647 {
                                                                                    case 9223372036854775807:
                                                                                        v3 = ((v3 + "hello") + (v3 + v83))
                                                                                        var v104 int8 = 96
                                                                                        fmt.Println(fi72, v104, v2)
                                                                                        _ = v104
                                                                                    case 32768:
                                                                                        fmt.Println(v3, v3)
                                                                                    case 573:
                                                                                        var v105 string = "\n\t"
                                                                                        var v106 bool = false
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = v6
                                                                                        }
                                                                                        if (-322 <= 2147483647) {
                                                                                            fmt.Println(v106, v105)
                                                                                        } else {
                                                                                            switch ("\n\t" + "a") {
                                                                                            case "":
                                                                                                if v106 {
                                                                                                    fmt.Println(fi72, v91, v83)
                                                                                                    switch fi73 {
                                                                                                    case 9223372036854775807:
                                                                                                        v107 := GenS4{F0: 0}
                                                                                                        _ = v107
                                                                                                    case 852:
                                                                                                        v108 := GenS9{F0: 595, F1: 42}
                                                                                                        _ = v108
                                                                                                    }
                                                                                                    v109 := GenS12{F0: ("\n\t" + ((v105 + (v105 + v3)) + (v83 + v3))), F1: ""}
                                                                                                    _ = v109
                                                                                                } else {
                                                                                                    switch v3 {
                                                                                                    case "a":
                                                                                                        v110 := v5[fi72]
                                                                                                        _ = v110
                                                                                                    case "foo bar":
                                                                                                        if true {
                                                                                                            fmt.Println(v106)
                                                                                                            delete(v4, v6)
                                                                                                        } else {
                                                                                                            switch (fi73 >> ((65535) & 63)) {
                                                                                                            case -1:
                                                                                                                v5[(fi73 & fi72)] = (-1 > 1)
                                                                                                            }
                                                                                                            switch (("\n\t" + v105) + "hello") {
                                                                                                            case "hello":
                                                                                                                v6 += 65535
                                                                                                            case "foo bar":
                                                                                                                if (v6 != (v6 / ((1) | 1))) {
                                                                                                                    v105 = ((v105 + "a") + v83)
                                                                                                                } else {
                                                                                                                    var v111 float32 = (0.0 / 1e10)
                                                                                                                    _ = v111
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v112 uint64 = 2
                                                                                                                _ = v112
                                                                                                            default:
                                                                                                                var v113 float32 = 0.1
                                                                                                                _ = v113
                                                                                                            }
                                                                                                            fmt.Println(v2)
                                                                                                            if true {
                                                                                                                if (-100 > 1) {
                                                                                                                    v1 = append(v1, v6)
                                                                                                                    v2 -= v2
                                                                                                                }
                                                                                                                var v114 bool = v106
                                                                                                                v115 := GenS5{F0: (-(fi72))}
                                                                                                                _ = v114
                                                                                                                _ = v115
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    v116 := GenS6{F0: -1, F1: v6}
                                                                                                                    if v106 {
                                                                                                                        var v117 bool = v106
                                                                                                                        switch v3 {
                                                                                                                        case "":
                                                                                                                            var v118 uint8 = 42
                                                                                                                            _ = v118
                                                                                                                        case "foo bar":
                                                                                                                            v3 = v3
                                                                                                                        case "a":
                                                                                                                            fmt.Println(v2, v117)
                                                                                                                        }
                                                                                                                        v119 := GenS13{F0: 42, F1: 18446744073709551615, F2: 801}
                                                                                                                        _ = v117
                                                                                                                        _ = v119
                                                                                                                    }
                                                                                                                    switch (("hello" + (v3 + (v3 + (v105 + v3)))) + ("hello" + v3)) {
                                                                                                                    case "a":
                                                                                                                        switch fi72 {
                                                                                                                        case -2147483648:
                                                                                                                            v120 := GenS14{F0: 2, F1: v91, F2: false}
                                                                                                                            _ = v120
                                                                                                                        case 107:
                                                                                                                            var v121 uint8 = 133
                                                                                                                            _ = v121
                                                                                                                        default:
                                                                                                                            if len(v1) > 0 {
                                                                                                                                v1[0] = (v6 - v6)
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v122 := GenS4{F0: -44}
                                                                                                                        _ = v122
                                                                                                                    }
                                                                                                                    _ = v116
                                                                                                                } else {
                                                                                                                    var v123 int16 = 85
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                                v4[(v6 / (((v6 / ((v6) | 1))) | 1))] = false
                                                                                                            }
                                                                                                            v91 = true
                                                                                                        }
                                                                                                    case "":
                                                                                                        v124 := v5[fi73]
                                                                                                        if (47 != -1) {
                                                                                                            v125 := GenS13{F0: 1, F1: 96, F2: v6}
                                                                                                            v1 = append(v1, v6)
                                                                                                            _ = v125
                                                                                                        }
                                                                                                        v126 := GenS3{F0: -128}
                                                                                                        var v127 int8 = 1
                                                                                                        var v128 float32 = 0.1
                                                                                                        _ = v124
                                                                                                        _ = v126
                                                                                                        _ = v127
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    if (127 > -128) {
                                                                                                        _ = len(v1)
                                                                                                        var v129 uint8 = 255
                                                                                                        _ = v129
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v105, v106)
                                                                                            }
                                                                                            fmt.Println(v91, v6)
                                                                                            var v130 float64 = v2
                                                                                            _ = v130
                                                                                        }
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    }
                                                                                    v2 -= float64((-2.25 - float32(((v2 + (((float64((((float32((float64(1e10) / 1.5)) / (1e10 - (0.0 + 0.0))) - (1e10 - (1e10 * -2.25))) / 1.0)) / (v2 * v2)) + v2) / 1.0)) * 1e10))))
                                                                                    if true {
                                                                                        fmt.Println(v83)
                                                                                        var v131 float64 = ((-2.25 - v2) / v2)
                                                                                        var v132 int32 = 0
                                                                                        _ = v131
                                                                                        _ = v132
                                                                                    } else {
                                                                                        switch fi72 {
                                                                                        case 32768:
                                                                                            if true {
                                                                                                if v91 {
                                                                                                    fmt.Println(v3, v91, v2)
                                                                                                    v91 = true
                                                                                                    var v133 int8 = -22
                                                                                                    _ = v133
                                                                                                }
                                                                                            } else {
                                                                                                var v134 uint64 = 2
                                                                                                v135 := GenS4{F0: 71}
                                                                                                switch fi73 {
                                                                                                case 32768:
                                                                                                    v136 := GenS11{F0: v134, F1: v6}
                                                                                                    _ = v136
                                                                                                }
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            }
                                                                                        default:
                                                                                            _ = len(v5)
                                                                                        }
                                                                                        var v137 uint32 = (v6 & (v6 - 37))
                                                                                        _ = v137
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v6 += v6
                                                                        }
                                                                    case "hello":
                                                                        fmt.Println(v2, fi72)
                                                                    case "":
                                                                        v138 := GenS15{F0: v6}
                                                                        _ = v138
                                                                    default:
                                                                        switch v83 {
                                                                        case "":
                                                                            if true {
                                                                                v139 := GenS16{F0: v83, F1: 42}
                                                                                v140 := v5[(fi72 & 128)]
                                                                                switch fi72 {
                                                                                case -778:
                                                                                    v140 = (!(v3 < "a"))
                                                                                default:
                                                                                    v141 := GenS17{F0: -120, F1: 42}
                                                                                    _ = v141
                                                                                }
                                                                                _ = v139
                                                                                _ = v140
                                                                            }
                                                                        case "a":
                                                                            v142 := GenS1{F0: (fi73 >> ((0) & 63))}
                                                                            fmt.Println(v2, v2, fi72)
                                                                            v6 += v6
                                                                            if ((float64((3.14159 * 1e10)) - v2) != 3.14159) {
                                                                                v143 := GenS18{F0: 65535}
                                                                                v1 = append(v1, (v6 / (((v6 >> ((394) & 63))) | 1)))
                                                                                _ = v143
                                                                            }
                                                                            _ = v142
                                                                        }
                                                                    }
                                                                    if (v91 && ((127 <= 115) && v91)) {
                                                                        v83 = (v3 + (v83 + "hello"))
                                                                        var v144 int16 = -1
                                                                        _ = v144
                                                                    }
                                                                default:
                                                                    v83 = ((v3 + v3) + (("foo bar" + "\n\t") + "\n\t"))
                                                                }
                                                                var v145 int = (fi73 / (((^(fi73))) | 1))
                                                                var v146 int32 = 42
                                                                _ = v145
                                                                _ = v146
                                                                _ = v97
                                                            }
                                                        case 779:
                                                            v147 := GenS1{F0: 32768}
                                                            _ = v147
                                                        default:
                                                            var v148 float32 = -2.25
                                                            _ = v148
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(fi72, v83, fi73)
                                                        if (203 <= 164) {
                                                            v6 -= v6
                                                            if (127 > 119) {
                                                                v149 := GenS7{F0: (fi73 / (((fi72 >> ((65535) & 63))) | 1)), F1: 255, F2: 1}
                                                                _ = v149
                                                            } else {
                                                                switch v83 {
                                                                case "\n\t":
                                                                    v91 = ((!((0 <= -1) || (!false))) && (-9223372036854775808 < 9223372036854775807))
                                                                case "foo bar":
                                                                    if true {
                                                                        var v150 int32 = -689
                                                                        v3 = v83
                                                                        _ = v150
                                                                    } else {
                                                                        fmt.Println(v91, v83, v6)
                                                                        v151 := GenS19{F0: (fi73 << ((145) & 63)), F1: 917}
                                                                        _ = v151
                                                                    }
                                                                }
                                                                v5[186] = (!v91)
                                                                fmt.Println(v91, v2, v6)
                                                                if v91 {
                                                                    var v152 float32 = -2.25
                                                                    _ = len(v4)
                                                                    _ = v152
                                                                } else {
                                                                    v3 = (("a" + (("a" + ("" + v83)) + ((v3 + "foo bar") + v83))) + "foo bar")
                                                                    v83 = v3
                                                                    var v153 uint = 42
                                                                    v3 = v83
                                                                    _ = v153
                                                                }
                                                                v6 = (v6 | 4294967295)
                                                            }
                                                            var v154 int64 = 0
                                                            _ = v154
                                                        } else {
                                                            var v155 uint8 = 128
                                                            _ = v155
                                                        }
                                                        var v156 uint32 = v6
                                                        _ = v156
                                                    default:
                                                        v2 -= v2
                                                    }
                                                    _ = v91
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                if (32767 <= 2) {
                                                    var v157 int16 = 32767
                                                    var v158 float64 = (v2 - float64(float32(v2)))
                                                    _ = v157
                                                    _ = v158
                                                }
                                                v6 += (^(v6))
                                            }
                                            fmt.Println(v2)
                                            var v159 float64 = v2
                                            v160 := GenS20{F0: 42, F1: 100, F2: 0}
                                            _ = v159
                                            _ = v160
                                        } else {
                                            v5[2147483647] = ((v3 + ((("foo bar" + ("\n\t" + "")) + (("" + ("\n\t" + v3)) + (v3 + ("\n\t" + v3)))) + v3)) > "a")
                                        }
                                        if (!true) {
                                            v6 += (v6 & v6)
                                            switch v3 {
                                            case "\n\t":
                                                v2 -= 0.1
                                                v161 := GenS8{F0: ((423 != 0) && true), F1: "\n\t", F2: 322}
                                                if (true && true) {
                                                    v6 += 2147483648
                                                }
                                                _ = v161
                                            }
                                        } else {
                                            v162 := GenS21{F0: 1}
                                            v163 := GenS1{F0: fi73}
                                            _ = v162
                                            _ = v163
                                        }
                                    }
                                    _ = v85
                                }
                            case "":
                                var v164 uint = 255
                                _ = v164
                            }
                            var v165 uint32 = 622
                            var v166 uint = 18446744073709551615
                            _ = v165
                            _ = v166
                        }
                        fmt.Println(fi73, fi72, v83)
                    }
                } else {
                    switch ((v3 + ((v3 + v83) + (("foo bar" + "\n\t") + (v83 + v83)))) + "a") {
                    case "a":
                        var v167 int = (fi72 / (((fi72 + fi72)) | 1))
                        _ = v167
                    case "foo bar":
                        if (float32((v2 / v2)) != -2.25) {
                            if (false && false) {
                                v2 += (v2 * 1e10)
                                fmt.Println(v83, v6)
                                var v168 int32 = 1
                                _ = v168
                            } else {
                                v2 = v2
                                var v169 uint16 = 65535
                                _ = v169
                            }
                            if (-697 != -666) {
                                var v170 uint16 = 0
                                _ = v170
                            } else {
                                fmt.Println(v2)
                                v171 := GenS13{F0: 1, F1: 1, F2: 167}
                                fmt.Println(fi73)
                                _ = v171
                            }
                            var v172 uint = 646
                            _ = v172
                        }
                    default:
                        if true {
                            var v173 int16 = -299
                            _ = v173
                        }
                    }
                    var v174 int64 = genFn0(127)
                    fmt.Println(fi73, v3)
                    _ = v174
                }
                v3 = "foo bar"
                _ = v83
            }
            var v175 uint = 255
            v176 := GenS22{F0: 32767}
            _ = fi73
            _ = v175
            _ = v176
        }
        _ = fi72
    }
    v3 = (v3 + "")
    if (1.5 == float32(3.14159)) {
        var v177 bool = (2 < 42)
        v178 := GenS23{F0: -1, F1: 0}
        _ = v177
        _ = v178
    } else {
        v179 := GenS12{F0: (v3 + ""), F1: v3}
        _ = v179
    }
    if false {
        v6 += 2147483648
        if false {
            var v180 int8 = -128
            for fi181 := 0; fi181 < 8; fi181++ {
                v6 = v6
                _ = fi181
            }
            var v182 uint8 = 0
            _ = v180
            _ = v182
        } else {
            v183 := GenS10{F0: -544, F1: false}
            var v184 int64 = 42
            switch ("" + v3) {
            case "":
                fmt.Println(v184, v2)
            default:
                switch 578 {
                case -230:
                    for fi185 := 0; fi185 < 14; fi185++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi186 := 0; fi186 < 10; fi186++ {
                            switch fi185 {
                            case 128:
                                var v187 int16 = -57
                                _ = v187
                            case -2147483648:
                                var v188 int32 = -898
                                _ = v188
                            }
                            _ = fi186
                        }
                        _ = fi185
                    }
                case 9223372036854775807:
                    switch v3 {
                    case "foo bar":
                        v4[(v6 % (((v6 << ((65535) & 63))) | 1))] = true
                    case "hello":
                        var v189 string = (("a" + v3) + (("\n\t" + v3) + v3))
                        _ = v189
                    }
                case -1:
                    for fi190 := 0; fi190 < 1; fi190++ {
                        for fi191 := 0; fi191 < 5; fi191++ {
                            delete(v5, fi190)
                            var v192 int8 = 1
                            switch v3 {
                            case "":
                                var v193 int8 = 0
                                _ = v193
                            case "\n\t":
                                v192 += v192
                            default:
                                switch "" {
                                case "":
                                    var v194 uint16 = 65535
                                    _ = v194
                                case "a":
                                    var v195 int8 = v192
                                    _ = v195
                                case "foo bar":
                                    v6 = v6
                                default:
                                    fmt.Println(v6, fi191, v192)
                                }
                            }
                            if ((genFn0(17) >= v184) && true) {
                                var v196 uint32 = (v6 >> ((325) & 63))
                                if false {
                                    var v197 uint64 = 855
                                    if false {
                                        fmt.Println(v6, v196)
                                        if ((255 != 1) || true) {
                                            fmt.Println(v3, fi191, v184)
                                            var v198 float32 = -2.25
                                            _ = v198
                                        } else {
                                            fmt.Println(v196, v184, fi190)
                                        }
                                    } else {
                                        var v199 string = (("" + ("hello" + v3)) + v3)
                                        if len(v1) > 0 {
                                            v1[0] = v196
                                        }
                                        _ = v199
                                    }
                                    var v200 uint64 = v197
                                    _ = v197
                                    _ = v200
                                }
                                if (4294967295 > 743) {
                                    switch v3 {
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = v6
                                        }
                                    case "\n\t":
                                        switch "hello" {
                                        case "a":
                                            v201 := GenS2{F0: false, F1: 175, F2: 547}
                                            _ = v201
                                        case "foo bar":
                                            switch fi190 {
                                            case 509:
                                                var v202 uint8 = 8
                                                _ = v202
                                            }
                                        default:
                                            v203 := GenS12{F0: (v3 + "a"), F1: (("foo bar" + (((v3 + (("" + (v3 + v3)) + "\n\t")) + v3) + "\n\t")) + v3)}
                                            _ = v203
                                        }
                                    default:
                                        _ = len(v5)
                                    }
                                } else {
                                    var v204 uint32 = v6
                                    _ = v204
                                }
                                _ = v196
                            } else {
                                var v205 int = fi191
                                if (!true) {
                                    v192 += (v192 / (((v192 & v192)) | 1))
                                } else {
                                    var v206 int16 = 0
                                    v207 := GenS24{F0: "hello", F1: 255, F2: (v6 << ((763) & 63))}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v206
                                    _ = v207
                                }
                                _ = v205
                            }
                            _ = fi191
                            _ = v192
                        }
                        _ = fi190
                    }
                default:
                    fmt.Println(v3, v2, v6)
                }
            }
            _ = v183
            _ = v184
        }
        for fi208 := 0; fi208 < 9; fi208++ {
            for fi209 := 0; fi209 < 12; fi209++ {
                fmt.Println(fi208)
                _ = len(v4)
                _ = fi209
            }
            if true {
                v4[v6] = (32767 == 663)
                v210 := GenS25{F0: 2, F1: 2}
                switch v3 {
                case "hello":
                    v2 = ((v2 - ((v2 - (v2 * 0.1)) * (v2 - (float64(-2.25) * (v2 + (v2 - 0.0)))))) + v2)
                case "":
                    v211 := v4[2147483647]
                    _ = v211
                default:
                    v212 := GenS26{F0: 0, F1: 564}
                    _ = v212
                }
                _ = v210
            }
            _ = fi208
        }
    }
    v6 += v6
    delete(v5, -9223372036854775808)
    var v213 uint8 = 85
    var v214 bool = true
    v3 = v3
    v215 := GenS3{F0: -92}
    var v216 uint16 = 2
    var v217 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v213:", v213)
    fmt.Println("v214:", v214)
    fmt.Println("v216:", v216)
    fmt.Println("v217:", v217)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v213
    _ = v214
    _ = v215
    _ = v216
    _ = v217
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
