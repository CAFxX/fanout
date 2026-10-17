package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS7 struct { F0 int32; F1 int64; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 string }

func main() {
    v1 := []uint32{1}
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 int16 = 351
    var v5 int8 = 20
    _ = len(v1)
    v5 += v5
    if (((!true) && true) && ((0.0 + 1.5) < 1.5)) {
        v6 := GenS1{F0: 4294967295}
        _ = v6
    }
    switch 0 {
    case 0:
        v4 = v4
    default:
        v7 := GenS1{F0: 388}
        _ = v7
    }
    v8 := v3[65535]
    v8 = (-(v8))
    v9 := v2["hello"]
    if false {
        delete(v3, 910)
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            switch (v8 / ((32767) | 1)) {
            case 0:
                fmt.Println(fi10, v4)
            case 32768:
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 uint32 = 0
                    v4 = (-(v4))
                    _ = fi11
                    _ = v12
                }
            case 945:
                var v13 int32 = 42
                _ = v13
            default:
                v14 := GenS1{F0: 255}
                _ = v14
            }
            v15 := GenS1{F0: 65535}
            for fi16 := 0; fi16 < 4; fi16++ {
                switch 9223372036854775807 {
                case 408:
                    v5 += 127
                case 32768:
                    fmt.Println(fi16)
                case 2147483647:
                    v8 += (v9 / ((v8) | 1))
                default:
                    _ = len(v1)
                }
                if (3.14159 >= ((-2.25 - 0.0) * (float32(1e10) / -2.25))) {
                    v17 := GenS1{F0: 309}
                    var v18 uint = 463
                    if (42 == 2) {
                        if len(v1) > 0 {
                            v1[0] = 193
                        }
                    }
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
            _ = fi10
            _ = v15
        }
        if true {
            switch 32768 {
            case 42:
                var v19 float64 = (float64((float32((0.1 - -2.25)) + 3.14159)) / 1.0)
                _ = v19
            default:
                var v20 uint8 = 255
                _ = v20
            }
            _ = len(v1)
            v4 += v4
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v2[""] = (^(fi21))
            _ = fi21
        }
    }
    if true {
        switch "\n\t" {
        case "":
            if true {
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch (^(fi22)) {
                    case -376:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            switch ("a" + "a") {
                            case "":
                                fmt.Println(v5)
                            default:
                                v8 = v8
                            }
                            _ = fi23
                        }
                    case 128:
                        v9 += v9
                    case 4294967295:
                        var v24 int64 = -1
                        _ = v24
                    }
                    switch ((("a" + "\n\t") + "foo bar") + "") {
                    case "a":
                        v4 = 401
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            v9 += (v8 ^ (v8 | (-(fi22))))
                        }
                    case "":
                        v25 := GenS2{F0: 18446744073709551615}
                        _ = v25
                    default:
                        fmt.Println(v4)
                    }
                    _ = fi22
                }
            } else {
                switch v8 {
                case 0:
                    var v26 uint32 = 676
                    _ = v26
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi27 := 0; fi27 < 7; fi27++ {
                    if (true && false) {
                        v8 = (-(v9))
                    }
                    _ = fi27
                }
            }
        case "foo bar":
            for fi28 := 0; fi28 < 1; fi28++ {
                switch v8 {
                case -30:
                    v29 := GenS1{F0: 2}
                    _ = v29
                default:
                    fmt.Println(v8, v5)
                }
                for fi30 := 0; fi30 < 1; fi30++ {
                    if true {
                        v4 += 1000
                        switch ((("\n\t" + ("hello" + ("foo bar" + "foo bar"))) + "a") + "") {
                        case "":
                            v31 := GenS1{F0: 4294967295}
                            _ = v31
                        default:
                            var v32 uint64 = 609
                            _ = v32
                        }
                        v1 = append(v1, 69)
                    } else {
                        switch "\n\t" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "hello":
                            fmt.Println(v4, v5)
                        case "foo bar":
                            v2["hello"] = (v9 / ((32768) | 1))
                        }
                        v5 -= (v5 | (v5 & (v5 ^ (-(v5)))))
                    }
                    var v33 float32 = 1e10
                    _ = fi30
                    _ = v33
                }
                _ = fi28
            }
        case "a":
            v8 += (^(v9))
        }
        var v34 int = v8
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := v3[206]
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    if ((0 <= 483) && false) {
        v37 := GenS1{F0: 0}
        var v38 uint = 42
        _ = v37
        _ = v38
    } else {
        switch (v8 % ((32768) | 1)) {
        case -1:
            switch "foo bar" {
            case "\n\t":
                var v39 float32 = (1.5 * (float32((0.0 / 3.14159)) + (1.5 - -2.25)))
                _ = v39
            case "hello":
                for fi40 := 0; fi40 < 12; fi40++ {
                    switch "foo bar" {
                    case "":
                        v4 -= (v4 >> ((109) & 63))
                    case "foo bar":
                        _ = len(v3)
                    }
                    fmt.Println(v9, v4)
                    _ = fi40
                }
            case "":
                if (127 >= 128) {
                    fmt.Println(v4, v8, v8)
                    for fi41 := 0; fi41 < 8; fi41++ {
                        v42 := GenS3{F0: 656, F1: 42, F2: 697}
                        _ = fi41
                        _ = v42
                    }
                    v5 = 100
                } else {
                    switch (v9 * (v8 * v8)) {
                    case -9223372036854775808:
                        var v43 uint = 667
                        _ = v43
                    case 201:
                        for fi44 := 0; fi44 < 2; fi44++ {
                            v45 := GenS2{F0: 4294967295}
                            for fi46 := 0; fi46 < 1; fi46++ {
                                switch -650 {
                                case 9223372036854775807:
                                    var v47 int = fi44
                                    _ = v47
                                case 42:
                                    var v48 uint32 = 477
                                    _ = v48
                                default:
                                    var v49 float32 = 0.1
                                    _ = v49
                                }
                                if ((float32((-2.25 * 1.5)) >= float32((0.0 + -2.25))) && false) {
                                    var v50 int32 = -280
                                    v51 := GenS4{F0: 2, F1: -9223372036854775808, F2: 0}
                                    if true {
                                        if (2 >= 255) {
                                            switch ("a" + "") {
                                            case "":
                                                v4 = v4
                                            }
                                        }
                                        if true {
                                            if (42 <= 65535) {
                                                var v52 uint = 65535
                                                _ = v52
                                            } else {
                                                v53 := GenS4{F0: 128, F1: -970, F2: 42}
                                                v54 := GenS1{F0: 65535}
                                                var v55 string = ((((((((("foo bar" + "foo bar") + "") + "") + "foo bar") + "hello") + "a") + "hello") + ("hello" + ("" + ("" + "")))) + "foo bar")
                                                if (42 == 271) {
                                                    var v56 int8 = v5
                                                    _ = v56
                                                } else {
                                                    if false {
                                                        v57 := v3[2147483648]
                                                        _ = v57
                                                    }
                                                    v58 := GenS3{F0: 509, F1: 2, F2: 1}
                                                    v50 = -2147483648
                                                    _ = v58
                                                }
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            v59 := GenS2{F0: 18446744073709551615}
                                            _ = v59
                                        } else {
                                            v60 := GenS3{F0: 2147483648, F1: 42, F2: 606}
                                            var v61 string = "foo bar"
                                            v62 := v2[v61]
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                } else {
                                    v63 := GenS2{F0: 22}
                                    _ = v63
                                }
                                if ((1 < 0) && true) {
                                    v8 -= v9
                                    var v64 uint64 = 18446744073709551615
                                    _ = v64
                                }
                                _ = fi46
                            }
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    if false {
                                        v1 = append(v1, 2147483648)
                                        var v66 string = "a"
                                        _ = v66
                                    } else {
                                        switch ("foo bar" + ("" + "hello")) {
                                        case "foo bar":
                                            v5 -= v5
                                        case "hello":
                                            if true {
                                                fmt.Println(v9)
                                            } else {
                                                switch (^(fi65)) {
                                                case 128:
                                                    var v67 int8 = (v5 / ((v5) | 1))
                                                    _ = v67
                                                case -183:
                                                    delete(v3, 1)
                                                default:
                                                    if ((true && (!(float64(float32(-2.25)) < 3.14159))) && ((!true) || true)) {
                                                        if (42 == 42) {
                                                            fmt.Println(v9, v9, v8)
                                                            _ = len(v1)
                                                            var v68 int = (v9 & (fi44 - -560))
                                                            fmt.Println(v5, v5, v5)
                                                            var v69 uint = 42
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    } else {
                                                        v9 -= (fi44 >> ((0) & 63))
                                                    }
                                                }
                                            }
                                        default:
                                            v1 = append(v1, 0)
                                        }
                                        switch "a" {
                                        case "":
                                            if (107 >= 9223372036854775807) {
                                                v70 := GenS5{F0: 0}
                                                v8 += 760
                                                switch "foo bar" {
                                                case "hello":
                                                    v71 := GenS6{F0: 809, F1: 128, F2: -938}
                                                    _ = v71
                                                case "":
                                                    if true {
                                                        if (2 == 2147483648) {
                                                            _ = len(v1)
                                                            v8 += (fi44 ^ (^(fi65)))
                                                        } else {
                                                            if true {
                                                                v3[908] = (v8 + (fi65 << ((21) & 63)))
                                                                var v72 int16 = v4
                                                                _ = v72
                                                            } else {
                                                                if true {
                                                                    v73 := GenS6{F0: 42, F1: 86, F2: (v4 * v4)}
                                                                    fmt.Println(fi44, fi44)
                                                                    var v74 uint32 = 200
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                            }
                                                            var v75 int32 = -599
                                                            switch fi65 {
                                                            case 2147483647:
                                                                var v76 int64 = 9223372036854775807
                                                                _ = v76
                                                            case -434:
                                                                var v77 uint = 572
                                                                _ = v77
                                                            }
                                                            _ = v75
                                                        }
                                                        v8 -= 2147483647
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v9 += 32768
                                                    case "":
                                                        v8 += (fi65 * (fi44 * (fi65 / (((-(v9))) | 1))))
                                                    case "foo bar":
                                                        v1 = append(v1, 1)
                                                    default:
                                                        var v78 int16 = v4
                                                        _ = v78
                                                    }
                                                }
                                                _ = v70
                                            }
                                        case "\n\t":
                                            delete(v3, 166)
                                        case "foo bar":
                                            _ = len(v1)
                                        }
                                        v9 = 32767
                                        var v79 uint32 = 2
                                        v5 += (v5 - (^(v5)))
                                        _ = v79
                                    }
                                    v9 -= (v9 & -891)
                                    _ = fi65
                                }
                            case "":
                                v8 += -228
                            }
                            _ = fi44
                            _ = v45
                        }
                    default:
                        if (172 == 36) {
                            switch "a" {
                            case "":
                                var v80 int64 = -9223372036854775808
                                _ = v80
                            case "\n\t":
                                for fi81 := 0; fi81 < 3; fi81++ {
                                    var v82 string = ""
                                    _ = fi81
                                    _ = v82
                                }
                            case "a":
                                switch "hello" {
                                case "":
                                    if false {
                                        for fi83 := 0; fi83 < 6; fi83++ {
                                            v84 := GenS4{F0: 114, F1: 42, F2: 2}
                                            _ = len(v1)
                                            v85 := GenS7{F0: 522, F1: -1, F2: v4}
                                            _ = fi83
                                            _ = v84
                                            _ = v85
                                        }
                                    } else {
                                        fmt.Println(v9, v8)
                                        var v86 uint = 181
                                        _ = v86
                                    }
                                case "hello":
                                    if ((!(0.1 > (1.5 + 3.14159))) && true) {
                                        switch (^(v9)) {
                                        case -164:
                                            v4 += 42
                                        case -9223372036854775808:
                                            if true {
                                                for fi87 := 0; fi87 < 9; fi87++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = 2
                                                    }
                                                    if true {
                                                        fmt.Println(v4, v4)
                                                        var v88 uint8 = 194
                                                        v8 = 42
                                                        _ = v88
                                                    } else {
                                                        var v89 int32 = 42
                                                        _ = v89
                                                    }
                                                    delete(v3, 4294967295)
                                                    _ = fi87
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 1
                                                }
                                            } else {
                                                var v90 int16 = v4
                                                v90 += -601
                                                v90 = v90
                                                _ = v90
                                            }
                                        case 32767:
                                            var v91 uint64 = 1
                                            _ = v91
                                        }
                                    }
                                }
                            }
                            var v92 int16 = 1
                            _ = v92
                        } else {
                            switch v8 {
                            case 32768:
                                fmt.Println(v8, v5, v4)
                            default:
                                var v93 float64 = ((3.14159 / float64((1.5 / -2.25))) - (1.5 * float64(1e10)))
                                _ = v93
                            }
                            switch v8 {
                            case -323:
                                v94 := v3[2147483647]
                                _ = v94
                            case -218:
                                v95 := GenS8{F0: 202}
                                _ = v95
                            case -987:
                                switch 90 {
                                case 28:
                                    if ((v8 >> ((530) & 63)) <= (v9 >> ((1) & 63))) {
                                        if false {
                                            var v96 float64 = -2.25
                                            v8 = v9
                                            switch ("" + (("hello" + "") + "a")) {
                                            case "hello":
                                                switch 562 {
                                                case 127:
                                                    v9 = (v8 % ((v8) | 1))
                                                }
                                            case "\n\t":
                                                var v97 int16 = (-(v4))
                                                _ = v97
                                            case "a":
                                                switch v9 {
                                                case 32768:
                                                    _ = len(v1)
                                                case 9223372036854775807:
                                                    v4 = v4
                                                }
                                            default:
                                                var v98 int = v8
                                                _ = v98
                                            }
                                            _ = v96
                                        }
                                        for fi99 := 0; fi99 < 12; fi99++ {
                                            v8 += 1
                                            switch (fi99 - v8) {
                                            case 793:
                                                switch (((("foo bar" + ("\n\t" + "\n\t")) + ("a" + "a")) + ("\n\t" + "hello")) + "\n\t") {
                                                case "\n\t":
                                                    for fi100 := 0; fi100 < 4; fi100++ {
                                                        if (true && true) {
                                                            v5 = 1
                                                            var v101 int32 = 0
                                                            v102 := GenS6{F0: 2, F1: 1, F2: -543}
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                        if (!(1 < 131)) {
                                                            if (2 <= 18446744073709551615) {
                                                                fmt.Println(v4, v5, fi99)
                                                                fmt.Println(fi99, fi99, v4)
                                                            } else {
                                                                v103 := v2["foo bar"]
                                                                var v104 int16 = v4
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            var v105 int32 = -1
                                                            _ = v105
                                                        } else {
                                                            v106 := GenS3{F0: 910, F1: 0, F2: 4294967295}
                                                            fmt.Println(fi99, v4)
                                                            _ = v106
                                                        }
                                                        var v107 int = (fi99 | -9223372036854775808)
                                                        _ = fi100
                                                        _ = v107
                                                    }
                                                default:
                                                    if true {
                                                        var v108 int64 = -9223372036854775808
                                                        _ = v108
                                                    } else {
                                                        v109 := GenS3{F0: 4294967295, F1: 506, F2: 42}
                                                        v1 = append(v1, 1)
                                                        for fi110 := 0; fi110 < 8; fi110++ {
                                                            if false {
                                                                v111 := GenS9{F0: 2147483647, F1: 128}
                                                                if (("\n\t" + "") > "a") {
                                                                    fmt.Println(v5)
                                                                } else {
                                                                    v112 := GenS3{F0: 2147483647, F1: 42, F2: 2}
                                                                    _ = v112
                                                                }
                                                                _ = v111
                                                            } else {
                                                                delete(v2, "a")
                                                                switch ((("hello" + "hello") + "a") + "a") {
                                                                case "foo bar":
                                                                    v8 = (fi110 * 9223372036854775807)
                                                                case "":
                                                                    v113 := GenS10{F0: -576, F1: (v5 | v5)}
                                                                    _ = v113
                                                                case "\n\t":
                                                                    v4 += (v4 * 42)
                                                                default:
                                                                    var v114 uint16 = 32767
                                                                    _ = v114
                                                                }
                                                                v4 = (^(v4))
                                                                v4 = (v4 >> ((776) & 63))
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v1 = append(v1, 1)
                                                                fmt.Println(v5, v5)
                                                            }
                                                            _ = fi110
                                                        }
                                                        _ = v109
                                                    }
                                                }
                                            case -9223372036854775808:
                                                if false {
                                                    v4 += v4
                                                    v115 := GenS3{F0: 42, F1: 65535, F2: 110}
                                                    if true {
                                                        v116 := GenS1{F0: 808}
                                                        _ = len(v3)
                                                        _ = v116
                                                    }
                                                    _ = v115
                                                } else {
                                                    v1 = append(v1, 60)
                                                    v117 := GenS11{F0: ""}
                                                    var v118 float64 = (0.0 - ((1e10 + 3.14159) * float64(1.5)))
                                                    _ = v117
                                                    _ = v118
                                                }
                                            }
                                            _ = fi99
                                        }
                                        var v119 float64 = ((1e10 / -2.25) * ((-2.25 / float64((0.0 + -2.25))) / float64(1e10)))
                                        _ = v119
                                    }
                                case 2147483647:
                                    v8 = -1
                                }
                            }
                        }
                    }
                }
            default:
                v9 = v8
            }
        case -781:
            v4 += v4
        case 4294967295:
            for fi120 := 0; fi120 < 19; fi120++ {
                var v121 uint = 65535
                _ = fi120
                _ = v121
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
