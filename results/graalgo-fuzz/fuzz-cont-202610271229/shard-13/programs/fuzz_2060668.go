package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 bool; F1 int64; F2 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 bool; F2 string }
type GenS9 struct { F0 uint }
func genFn0() int {
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            var v7 int64 = 873
            var v8 int16 = -1
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 65535}
            v5 = 127
            if (false && false) {
                for fi10 := 0; fi10 < 20; fi10++ {
                    switch (fi6 % ((9223372036854775807) | 1)) {
                    case -9223372036854775808:
                        v11 := []int{}
                        _ = v11
                    case 82:
                        fmt.Println(v5)
                    default:
                        var v12 int16 = -32768
                        _ = v12
                    }
                    v13 := []int{(fi10 >> ((1) & 63)), (fi6 & fi10)}
                    fmt.Println(fi10, fi6, v5)
                    _ = fi10
                    _ = v13
                }
                fmt.Println(v5, v5)
            } else {
                var v14 uint = 0
                var v15 bool = (65535 != 42)
                _ = v14
                _ = v15
            }
            _ = v9
        }
        if false {
            v16 := []int{fi6}
            v5 -= v5
            v17 := GenS1{F0: 638}
            _ = v16
            _ = v17
        }
        _ = fi6
    }
    if (18446744073709551615 != 2) {
        if (-767 < 2147483647) {
            switch "hello" {
            case "a":
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := []int{32768}
                    _ = fi18
                    _ = v19
                }
            case "foo bar":
                var v20 bool = (true && (!false))
                _ = v20
            }
            fmt.Println(v5)
            switch 127 {
            case -2147483648:
                fmt.Println(v5)
            }
        } else {
            v5 -= 0
        }
    }
    if true {
        switch ("a" + ("hello" + "\n\t")) {
        case "\n\t":
            var v21 float32 = (3.14159 * ((1e10 - 0.0) / 3.14159))
            _ = v21
        case "":
            var v22 int64 = 0
            _ = v22
        case "hello":
            var v23 int32 = 138
            _ = v23
        }
        var v24 uint16 = 2
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 int8 = 42
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        fmt.Println(v5)
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS1{F0: 72}
            if (!true) {
                if false {
                    fmt.Println(fi27, fi27)
                    var v29 bool = true
                    _ = v29
                } else {
                    var v30 int64 = 493
                    _ = v30
                }
                var v31 int8 = -1
                var v32 float64 = 0.0
                _ = v31
                _ = v32
            }
            if true {
                v5 = 1
            }
            _ = fi27
            _ = v28
        }
        v5 = (v5 ^ v5)
    }
    fmt.Println(v5)
    _ = v5
    return 881
}

func main() {
    var v1 uint16 = 622
    var v2 uint8 = 1
    var v3 uint64 = 18446744073709551615
    var v4 float32 = (3.14159 / 3.14159)
    v33 := []int{-281, 9223372036854775807, -664, 128}
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    if (42 <= 32767) {
        fmt.Println(v1)
    } else {
        if (979 > 103) {
            switch 127 {
            case -1:
                fmt.Println(v1, v1, v1)
            case 32767:
                switch ("" + "foo bar") {
                case "":
                    var v34 int = 4294967295
                    _ = v34
                case "\n\t":
                    if true {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            var v36 uint = 65535
                            _ = fi35
                            _ = v36
                        }
                        var v37 int32 = 0
                        _ = v37
                    } else {
                        var v38 uint16 = 255
                        var v39 int16 = 1000
                        _ = len(v33)
                        _ = v38
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 4; fi40++ {
                        v3 -= (^(v3))
                        v3 += v3
                        if false {
                            for fi41 := 0; fi41 < 4; fi41++ {
                                v42 := GenS2{F0: -2147483648}
                                _ = fi41
                                _ = v42
                            }
                        } else {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                v1 += v1
                            default:
                                var v43 uint64 = (v3 ^ v3)
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 14; fi44++ {
                                switch ("\n\t" + ("hello" + "a")) {
                                case "hello":
                                    var v45 int = (fi44 | (fi44 >> ((42) & 63)))
                                    _ = v45
                                case "foo bar":
                                    if false {
                                        if (!false) {
                                            var v46 string = (("a" + "") + "\n\t")
                                            switch (v46 + "a") {
                                            case "foo bar":
                                                var v47 int64 = 372
                                                _ = v47
                                            case "":
                                                v48 := GenS3{F0: 640, F1: "hello"}
                                                _ = v48
                                            case "hello":
                                                v46 = v46
                                            default:
                                                fmt.Println(v4, fi44, fi44)
                                            }
                                            _ = len(v33)
                                            _ = v46
                                        } else {
                                            var v49 bool = true
                                            _ = v49
                                        }
                                        if (!true) {
                                            var v50 float64 = float64(0.1)
                                            _ = v50
                                        } else {
                                            var v51 bool = false
                                            switch fi44 {
                                            case 32767:
                                                if len(v33) > 0 {
                                                    v33[0] = (-(fi40))
                                                }
                                            }
                                            var v52 int64 = 1
                                            if ((!(!true)) || v51) {
                                                v52 += (v52 % ((-451) | 1))
                                            } else {
                                                _ = len(v33)
                                            }
                                            if (!(!(!(!v51)))) {
                                                if v51 {
                                                    v3 += 0
                                                    v53 := GenS2{F0: 0}
                                                    _ = v53
                                                } else {
                                                    switch int(v52) {
                                                    case 127:
                                                        fmt.Println(v51)
                                                        v4 = (1e10 + v4)
                                                        var v54 int64 = v52
                                                        _ = v54
                                                    case 9223372036854775807:
                                                        if true {
                                                            var v55 uint32 = 455
                                                            switch 492 {
                                                            case -766:
                                                                if len(v33) >= 2 {
                                                                    v33 = v33[1:len(v33)]
                                                                }
                                                            }
                                                            _ = v55
                                                        } else {
                                                            var v56 float64 = 3.14159
                                                            _ = v56
                                                        }
                                                    case 20:
                                                        if ((v3 - v3) == (^(v3))) {
                                                            v33 = append(v33, (fi44 * fi44))
                                                            fmt.Println(v1, v2, v51)
                                                        }
                                                    }
                                                    fmt.Println(fi40)
                                                }
                                                _ = len(v33)
                                                if (false || (fi40 != fi40)) {
                                                    v52 += (v52 + (^(v52)))
                                                    v2 += v2
                                                    v52 += (v52 & v52)
                                                    var v57 uint32 = 2147483648
                                                    _ = v57
                                                } else {
                                                    fmt.Println(v4, v1)
                                                }
                                            }
                                            v4 = v4
                                            _ = v51
                                            _ = v52
                                        }
                                    } else {
                                        if len(v33) >= 2 {
                                            v33 = v33[1:len(v33)]
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (true && false) {
                                                var v58 int16 = 1000
                                                _ = v58
                                            }
                                        case "a":
                                            switch fi44 {
                                            case 691:
                                                v4 -= 0.1
                                            case 32767:
                                                var v59 uint32 = 2147483647
                                                _ = v59
                                            }
                                        case "foo bar":
                                            var v60 uint8 = (v2 * v2)
                                            _ = v60
                                        }
                                    }
                                }
                                _ = fi44
                            }
                        }
                        _ = fi40
                    }
                }
            case -165:
                var v61 uint16 = (v1 << ((766) & 63))
                _ = v61
            }
            switch 374 {
            case 0:
                for fi62 := 0; fi62 < 20; fi62++ {
                    v3 -= v3
                    _ = fi62
                }
            case 1:
                switch -339 {
                case -160:
                    if false {
                        var v63 string = "a"
                        _ = v63
                    } else {
                        if (0.0 > 1.5) {
                            switch "foo bar" {
                            case "":
                                for fi64 := 0; fi64 < 18; fi64++ {
                                    var v65 uint8 = (v2 % ((v2) | 1))
                                    _ = fi64
                                    _ = v65
                                }
                            case "foo bar":
                                var v66 int64 = -9223372036854775808
                                _ = v66
                            case "a":
                                var v67 int32 = -2147483648
                                _ = v67
                            default:
                                switch "a" {
                                case "":
                                    var v68 uint16 = v1
                                    _ = v68
                                case "a":
                                    v4 = v4
                                }
                            }
                            switch 9223372036854775807 {
                            case -102:
                                fmt.Println(v2, v1)
                            }
                            var v69 float64 = -2.25
                            _ = v69
                        }
                    }
                default:
                    var v70 float64 = (-2.25 - 0.0)
                    _ = v70
                }
            case 293:
                var v71 string = ("a" + "foo bar")
                _ = v71
            default:
                for fi72 := 0; fi72 < 18; fi72++ {
                    fmt.Println(v2, fi72, v2)
                    _ = fi72
                }
            }
            if true {
                switch 2147483647 {
                case 0:
                    v73 := GenS4{F0: 42, F1: 538, F2: 32767}
                    _ = v73
                default:
                    v3 = v3
                }
                v74 := GenS5{F0: (v3 & v3), F1: 42}
                _ = v74
            }
        } else {
            if len(v33) > 0 {
                v33[0] = 204
            }
            var v75 uint = 42
            _ = v75
        }
        fmt.Println(v4, v2, v1)
    }
    switch ("hello" + "") {
    case "":
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    case "foo bar":
        for fi76 := 0; fi76 < 20; fi76++ {
            switch ("" + "\n\t") {
            case "hello":
                v77 := GenS5{F0: (v3 >> ((857) & 63)), F1: -1}
                _ = v77
            case "\n\t":
                if (v1 != uint16(v3)) {
                    switch (fi76 << ((629) & 63)) {
                    case -882:
                        if (true && (!false)) {
                            v78 := GenS2{F0: 618}
                            v79 := GenS5{F0: (v3 << ((0) & 63)), F1: -1}
                            var v80 int16 = 42
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            if false {
                                var v81 int32 = -362
                                for fi82 := 0; fi82 < 15; fi82++ {
                                    var v83 int64 = 0
                                    var v84 uint64 = (v3 * v3)
                                    if false {
                                        v85 := GenS6{F0: (!((fi76 ^ (^(fi82))) < fi76)), F1: (v83 / ((1) | 1)), F2: int32(fi82)}
                                        _ = v85
                                    }
                                    var v86 uint8 = (v2 << ((286) & 63))
                                    v84 = (^(v3))
                                    _ = fi82
                                    _ = v83
                                    _ = v84
                                    _ = v86
                                }
                                _ = v81
                            } else {
                                var v87 bool = true
                                if ((!(-2147483648 <= int32(v2))) || (!(!((-128 <= 42) || (!(false || v87)))))) {
                                    var v88 float64 = (float64(v4) * 1.5)
                                    var v89 float32 = v4
                                    for fi90 := 0; fi90 < 13; fi90++ {
                                        if (930 > 65535) {
                                            var v91 uint8 = (v2 - (^(v2)))
                                            _ = len(v33)
                                            _ = v91
                                        } else {
                                            if ((float64(-2.25) < v88) && ((((v1 - (v1 | 65535)) > (v1 ^ v1)) && (v2 < v2)) && (v1 != v1))) {
                                                v33 = append(v33, (fi76 / (((^(fi76))) | 1)))
                                                v92 := GenS1{F0: 42}
                                                var v93 int = 0
                                                _ = v92
                                                _ = v93
                                            }
                                        }
                                        fmt.Println(v2)
                                        _ = fi90
                                    }
                                    _ = v88
                                    _ = v89
                                }
                                v87 = (32768 > -2147483648)
                                _ = v87
                            }
                            v4 -= (0.1 / 3.14159)
                            fmt.Println(v4, v2)
                        }
                    default:
                        if ((1 >= 1) && true) {
                            var v94 float64 = 1.5
                            v95 := GenS5{F0: 1, F1: int16(v1)}
                            _ = v94
                            _ = v95
                        }
                    }
                    if (-508 <= -9223372036854775808) {
                        if true {
                            var v96 int8 = 98
                            _ = v96
                        } else {
                            var v97 float64 = 0.1
                            _ = v97
                        }
                        if (v2 != 128) {
                            for fi98 := 0; fi98 < 18; fi98++ {
                                switch genFn0() {
                                case -1:
                                    var v99 uint32 = 42
                                    _ = v99
                                }
                                _ = fi98
                            }
                            for fi100 := 0; fi100 < 8; fi100++ {
                                var v101 int64 = -368
                                v101 = v101
                                v102 := GenS4{F0: 0, F1: -2147483648, F2: -580}
                                _ = fi100
                                _ = v101
                                _ = v102
                            }
                            var v103 int = (fi76 ^ (fi76 ^ (^(fi76))))
                            _ = v103
                        }
                        for fi104 := 0; fi104 < 12; fi104++ {
                            if len(v33) > 0 {
                                v33[0] = fi104
                            }
                            _ = fi104
                        }
                    } else {
                        v105 := GenS5{F0: v3, F1: 1}
                        _ = v105
                    }
                    if len(v33) > 0 {
                        v33[0] = (fi76 << ((1) & 63))
                    }
                } else {
                    var v106 uint64 = (v3 ^ 2)
                    _ = v106
                }
            }
            _ = fi76
        }
    default:
        switch 1 {
        case 451:
            _ = len(v33)
        default:
            var v107 int32 = -2147483648
            _ = v107
        }
    }
    v108 := GenS2{F0: -2147483648}
    var v109 uint16 = (v1 / ((0) | 1))
    for fi110 := 0; fi110 < 19; fi110++ {
        switch "hello" {
        case "a":
            if (!(1 <= -942)) {
                var v111 int8 = -1
                v2 -= (v2 >> ((0) & 63))
                _ = v111
            } else {
                for fi112 := 0; fi112 < 12; fi112++ {
                    v113 := GenS7{F0: -1}
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    _ = fi112
                    _ = v113
                }
                if false {
                    v114 := GenS8{F0: 127, F1: (1.5 == (float64((0.1 - 1e10)) - (float64(3.14159) - 3.14159))), F2: ""}
                    var v115 int64 = -9223372036854775808
                    v116 := GenS9{F0: 4294967295}
                    _ = v114
                    _ = v115
                    _ = v116
                }
            }
        case "\n\t":
            v109 -= 32767
        case "hello":
            switch fi110 {
            case 127:
                for fi117 := 0; fi117 < 3; fi117++ {
                    var v118 string = "hello"
                    v118 = "hello"
                    _ = fi117
                    _ = v118
                }
            case -145:
                fmt.Println(v3, v4)
            case 32768:
                switch fi110 {
                case 0:
                    var v119 float32 = v4
                    _ = v119
                case 966:
                    if (1 < 65535) {
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                    }
                }
            }
        }
        _ = fi110
    }
    var v120 bool = true
    switch 0 {
    case 9223372036854775807:
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    default:
        if v120 {
            var v121 int32 = int32(v109)
            _ = v121
        }
    }
    v120 = v120
    switch "a" {
    case "":
        var v122 uint64 = (v3 & v3)
        _ = v122
    case "a":
        switch 128 {
        case 1:
            fmt.Println(v1, v2)
        case -2147483648:
            v3 += 2
        case 4294967295:
            for fi123 := 0; fi123 < 4; fi123++ {
                for fi124 := 0; fi124 < 4; fi124++ {
                    fmt.Println(v2, v2)
                    switch "\n\t" {
                    case "a":
                        if len(v33) > 0 {
                            v33[0] = genFn0()
                        }
                    case "\n\t":
                        var v125 uint16 = v109
                        _ = v125
                    case "hello":
                        v126 := GenS7{F0: -128}
                        _ = v126
                    }
                    _ = fi124
                }
                v4 -= (v4 * float32(((3.14159 - float64((v4 - v4))) / ((-2.25 * 0.1) / (1.5 - (0.0 / -2.25))))))
                _ = fi123
            }
        default:
            v2 += 109
        }
    }
    fmt.Println(v2)
    v127 := GenS5{F0: (v3 >> ((1) & 63)), F1: -1}
    v128 := GenS5{F0: v3, F1: 42}
    v3 += v3
    var v129 float64 = -2.25
    for fi130 := 0; fi130 < 16; fi130++ {
        var v131 float64 = 0.1
        var v132 string = ("foo bar" + "a")
        _ = fi130
        _ = v131
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v120:", v120)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v109
    _ = v120
    _ = v127
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
