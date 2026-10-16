package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int64; F1 int32; F2 int16 }
type GenS7 struct { F0 uint; F1 bool; F2 int }
type GenS8 struct { F0 int; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 bool; F1 uint16; F2 int16 }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int; F1 int8 }
type GenS15 struct { F0 int32; F1 uint32 }
type GenS16 struct { F0 int32; F1 uint8; F2 int8 }
func genFn0(p0 string) uint32 {
    var v4 int64 = 42
    switch "hello" {
    case "\n\t":
        var v5 int64 = -9223372036854775808
        _ = v5
    case "hello":
        v6 := GenS1{F0: 32767, F1: 65535}
        _ = v6
    case "":
        v7 := GenS1{F0: 0, F1: 630}
        _ = v7
    }
    v8 := GenS1{F0: 32767, F1: 980}
    if (42 < 4294967295) {
        fmt.Println(p0, p0)
    } else {
        var v9 uint8 = 45
        _ = v9
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    _ = v8
    return 42
}
func genFn1(p0 int) string {
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            switch ("" + ("" + ("hello" + "foo bar"))) {
            case "foo bar":
                var v12 uint8 = 2
                _ = v12
            case "hello":
                fmt.Println(fi11, p0, fi10)
            case "a":
                var v13 int = 1
                _ = v13
            default:
                v14 := GenS1{F0: 32767, F1: 791}
                _ = v14
            }
            _ = fi11
        }
        var v15 uint = 317
        v16 := GenS1{F0: 110, F1: 673}
        _ = fi10
        _ = v15
        _ = v16
    }
    var v17 float64 = ((-2.25 + 3.14159) * ((-2.25 / (3.14159 - 1.5)) - 1e10))
    fmt.Println(p0, v17)
    _ = p0
    _ = v17
    return "a"
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v1 = -2147483648
    var v18 bool = (!(32767 < 2))
    if true {
        if ((v1 == (v1 - (v1 ^ v1))) || (false || (!(-128 >= 0)))) {
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(v18, fi19)
                v20 := GenS2{F0: ("foo bar" + ("foo bar" + (("" + "a") + "foo bar"))), F1: ("" + "hello")}
                _ = fi19
                _ = v20
            }
            v21 := GenS3{F0: 42}
            var v22 uint32 = genFn0("")
            _ = v21
            _ = v22
        }
        if ((!v18) || ((3.14159 * ((3.14159 + 1.5) + 1.5)) >= 0.0)) {
            var v23 float64 = 1.5
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if v18 {
            v25 := GenS4{F0: 115, F1: 255}
            _ = v25
        } else {
            if (120 > 1) {
                v26 := GenS4{F0: -128, F1: 916}
                v1 = 376
                _ = v26
            } else {
                fmt.Println(v1)
                switch "foo bar" {
                case "a":
                    if v18 {
                        var v27 int16 = -32768
                        if ("foo bar" >= ("" + "a")) {
                            if (!(0 == 4294967295)) {
                                delete(v3, genFn0(("a" + "hello")))
                                switch fi24 {
                                case 4294967295:
                                    switch "\n\t" {
                                    case "hello":
                                        switch v1 {
                                        case 9223372036854775807:
                                            v1 -= fi24
                                        case 2147483647:
                                            _ = len(v3)
                                        case 127:
                                            for fi28 := 0; fi28 < 12; fi28++ {
                                                v1 = fi28
                                                if false {
                                                    v18 = v18
                                                    if (true || (!(65535 >= 65535))) {
                                                        if false {
                                                            if (v27 != (-(v27))) {
                                                                v29 := GenS4{F0: -19, F1: 562}
                                                                _ = v29
                                                            }
                                                            switch fi28 {
                                                            case 127:
                                                                v1 += 2147483647
                                                            case 7:
                                                                fmt.Println(v27, v1)
                                                            case 32767:
                                                                delete(v2, 2147483647)
                                                            }
                                                            switch ((("\n\t" + "hello") + "a") + ("hello" + "hello")) {
                                                            case "\n\t":
                                                                fmt.Println(v18)
                                                            case "":
                                                                v1 -= (v1 / ((fi28) | 1))
                                                            }
                                                        } else {
                                                            var v30 uint8 = 113
                                                            fmt.Println(fi24, fi24)
                                                            var v31 int32 = 2147483647
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        var v32 uint16 = 2
                                                        _ = v32
                                                    }
                                                }
                                                _ = fi28
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v1, v27)
                                    case "":
                                        var v33 uint = 409
                                        _ = v33
                                    default:
                                        v18 = v18
                                    }
                                case 0:
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        v35 := GenS1{F0: 213, F1: 65535}
                                        if (!v18) {
                                            var v36 bool = true
                                            _ = v36
                                        }
                                        v18 = v18
                                        _ = fi34
                                        _ = v35
                                    }
                                case 312:
                                    for fi37 := 0; fi37 < 12; fi37++ {
                                        v18 = true
                                        delete(v2, 42)
                                        fmt.Println(v27)
                                        _ = fi37
                                    }
                                }
                                v38 := []int{v1, 0, (-(fi24))}
                                _ = v38
                            } else {
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    fmt.Println(v18, fi24)
                                    v1 = v1
                                    var v40 int = 32767
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            v41 := GenS4{F0: 1, F1: 42}
                            _ = v41
                        }
                        var v42 int64 = -1
                        _ = v27
                        _ = v42
                    } else {
                        v2[932] = v18
                        v43 := []int{v1}
                        _ = v43
                    }
                case "foo bar":
                    var v44 uint8 = 0
                    _ = v44
                case "":
                    if false {
                        if v18 {
                            delete(v3, 65)
                            var v45 float64 = 0.1
                            var v46 uint32 = genFn0(("hello" + "\n\t"))
                            _ = v45
                            _ = v46
                        } else {
                            for fi47 := 0; fi47 < 6; fi47++ {
                                v1 = (v1 >> ((1) & 63))
                                if ((((1.5 + (0.0 * float32(0.1))) / 1.5) + (-2.25 + 3.14159)) <= -2.25) {
                                    fmt.Println(v18, fi24, v18)
                                }
                                _ = fi47
                            }
                        }
                        switch (v1 | (fi24 << ((65535) & 63))) {
                        case 32768:
                            v48 := GenS2{F0: "foo bar", F1: "a"}
                            _ = v48
                        }
                        v1 = (v1 % ((v1) | 1))
                    }
                default:
                    v18 = true
                }
                var v49 uint8 = 210
                _ = v49
            }
            fmt.Println(v18)
            fmt.Println(v1)
        }
        for fi50 := 0; fi50 < 15; fi50++ {
            v51 := GenS1{F0: 583, F1: 2}
            v52 := GenS5{F0: 42, F1: 2147483647}
            var v53 uint16 = 1
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi24
    }
    var v54 float64 = (float64(float32(1e10)) - (1e10 - 1e10))
    var v55 int = v1
    var v56 int = v55
    v55 = (v55 + 0)
    var v57 uint8 = 128
    if true {
        switch 0 {
        case 128:
            fmt.Println(v56, v1, v1)
        }
        v58 := v2[65535]
        _ = v58
    }
    var v59 uint16 = 0
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "a":
            v56 -= -32
        default:
            v60 := GenS6{F0: -969, F1: 42, F2: -1}
            _ = v60
        }
    case "foo bar":
        switch v1 {
        case 32768:
            v61 := v3[1]
            _ = v61
        case 4294967295:
            switch "" {
            case "hello":
                v1 = (-(v55))
            default:
                if (2147483647 < 4294967295) {
                    v62 := GenS7{F0: 697, F1: false, F2: v56}
                    v59 += (v59 << ((18446744073709551615) & 63))
                    if false {
                        var v63 uint8 = (v57 >> ((514) & 63))
                        _ = v63
                    }
                    _ = v62
                }
            }
        case -1:
            if v18 {
                v64 := GenS5{F0: 87, F1: 65535}
                for fi65 := 0; fi65 < 19; fi65++ {
                    if ((v56 << ((991) & 63)) != (v56 >> ((540) & 63))) {
                        for fi66 := 0; fi66 < 8; fi66++ {
                            if (664 < 42) {
                                v54 = (float64(0.1) - v54)
                            }
                            _ = fi66
                        }
                        var v67 string = ""
                        v68 := []int{fi65, (v1 << ((1) & 63)), v55, (fi65 * (v1 - v56))}
                        _ = v67
                        _ = v68
                    }
                    _ = fi65
                }
                _ = v64
            } else {
                for fi69 := 0; fi69 < 16; fi69++ {
                    v70 := GenS8{F0: 32768, F1: -128, F2: (v57 >> ((0) & 63))}
                    v57 += v57
                    v71 := v3[4294967295]
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                v72 := []int{(v55 | 32768), (v56 << ((42) & 63)), -175, (v1 + 968)}
                v73 := GenS8{F0: v1, F1: 20, F2: v57}
                _ = v72
                _ = v73
            }
        }
    case "":
        if (-693 == -1) {
            v74 := GenS9{F0: 548}
            v57 -= (^(v57))
            _ = v74
        } else {
            for fi75 := 0; fi75 < 11; fi75++ {
                v76 := GenS10{F0: (v56 & (v55 / ((fi75) | 1))), F1: -128}
                v77 := GenS11{F0: (1e10 == -2.25), F1: v59, F2: 790}
                for fi78 := 0; fi78 < 10; fi78++ {
                    if (!(!v18)) {
                        v79 := GenS8{F0: (v1 | (fi75 & fi75)), F1: 42, F2: (v57 + v57)}
                        _ = v79
                    }
                    if (("a" + "") == "hello") {
                        var v80 int = (v56 + v1)
                        if v18 {
                            v18 = true
                            if v18 {
                                if false {
                                    v55 = (fi78 >> ((557) & 63))
                                } else {
                                    var v81 int64 = 9223372036854775807
                                    _ = v81
                                }
                                v55 += fi78
                            }
                            fmt.Println(v56)
                            var v82 int32 = 42
                            _ = v82
                        } else {
                            v18 = false
                            switch fi75 {
                            case 32767:
                                if v18 {
                                    var v83 uint64 = 18446744073709551615
                                    _ = v83
                                } else {
                                    if v18 {
                                        if (v18 && (1 > 65535)) {
                                            var v84 float32 = float32(v54)
                                            fmt.Println(v54, v55)
                                            var v85 float64 = ((v54 - (v54 * ((((v54 - (((((v54 + (v54 - (v54 + 1.5))) * (v54 + (v54 * (v54 + 0.1)))) / v54) - v54) + ((v54 + float64(v84)) / v54))) / 3.14159) + v54) - v54))) * ((1.5 / 0.1) * (v54 + (v54 * 1.5))))
                                            _ = v84
                                            _ = v85
                                        }
                                        v86 := GenS4{F0: 100, F1: 65535}
                                        v80 += (fi75 >> ((546) & 63))
                                        _ = v86
                                    } else {
                                        var v87 bool = (124 <= 133)
                                        var v88 uint = 65535
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            v56 = (v56 * (fi78 & (v80 + (v55 << ((v88) & 63)))))
                                        case "foo bar":
                                            v89 := []int{(v80 | (v55 | (v56 / (((v1 >> ((1) & 63))) | 1)))), -319, v55}
                                            _ = v89
                                        case "":
                                            if true {
                                                var v90 int8 = -66
                                                _ = v90
                                            }
                                        default:
                                            if v87 {
                                                _ = len(v2)
                                                if (!v87) {
                                                    v2[719] = v87
                                                }
                                                v54 += (v54 / 0.1)
                                                v91 := GenS10{F0: int(v59), F1: 0}
                                                _ = v91
                                            } else {
                                                if true {
                                                    v92 := GenS4{F0: 1, F1: uint(v57)}
                                                    _ = v92
                                                }
                                                v56 = (-(fi78))
                                                if (((((0.0 * float32(3.14159)) - 1.5) == float32(0.0)) || (!(v1 < v56))) && (971 != -1)) {
                                                    if ((0.0 / (((0.0 + (1.5 * 0.1)) + 0.0) * (0.0 - 3.14159))) > (1.5 + 1.5)) {
                                                        var v93 int8 = 90
                                                        _ = v93
                                                    }
                                                    if v87 {
                                                        v55 = v1
                                                    }
                                                }
                                            }
                                        }
                                        _ = v87
                                        _ = v88
                                    }
                                    switch fi75 {
                                    case 32768:
                                        var v94 uint = 255
                                        _ = v94
                                    default:
                                        if (v56 >= (v80 - fi75)) {
                                            v1 += -559
                                            var v95 uint8 = (^(v57))
                                            _ = v95
                                        }
                                    }
                                }
                                var v96 int32 = -1
                                v54 += (v54 * 1e10)
                                if (v18 || (v57 > 255)) {
                                    v54 += (v54 / v54)
                                    switch (fi78 - 32767) {
                                    case 127:
                                        var v97 int32 = -1
                                        _ = v97
                                    default:
                                        fmt.Println(v96, v57)
                                        switch ("\n\t" + "\n\t") {
                                        case "":
                                            var v98 bool = (!v18)
                                            _ = v98
                                        case "hello":
                                            v18 = v18
                                            var v99 uint = 1
                                            v96 = v96
                                            v56 = 4294967295
                                            var v100 uint64 = 473
                                            v18 = false
                                            _ = v100
                                            _ = v99
                                        case "\n\t":
                                            v1 += 903
                                        default:
                                            var v101 uint8 = (v57 / ((v57) | 1))
                                            _ = v101
                                        }
                                        v1 = fi75
                                    }
                                    v57 = (v57 << ((718) & 63))
                                    switch ("a" + "hello") {
                                    case "hello":
                                        if (!v18) {
                                            switch ((("\n\t" + "foo bar") + "foo bar") + "foo bar") {
                                            case "":
                                                v102 := GenS5{F0: 0, F1: 2}
                                                _ = v102
                                            case "\n\t":
                                                var v103 uint8 = (v57 - 1)
                                                _ = v103
                                            default:
                                                v18 = (!true)
                                                switch 607 {
                                                case 42:
                                                    var v104 uint8 = v57
                                                    _ = v104
                                                case 0:
                                                    v105 := GenS7{F0: 862, F1: (false && ((1 != 18446744073709551615) && (1 <= 42))), F2: (fi75 + v55)}
                                                    switch "hello" {
                                                    case "a":
                                                        v106 := []int{-200, fi78, (fi75 ^ int(v57))}
                                                        switch (fi75 ^ -2147483648) {
                                                        case 2147483647:
                                                            v3[2147483647] = ("a" + "a")
                                                        case 110:
                                                            v96 += 163
                                                        case 0:
                                                            v107 := GenS11{F0: ((18446744073709551615 <= 2) && v18), F1: v59, F2: 0}
                                                            _ = v107
                                                        default:
                                                            switch fi75 {
                                                            case -867:
                                                                fmt.Println(fi78, v1, v1)
                                                            case -550:
                                                                switch v55 {
                                                                case 128:
                                                                    fmt.Println(v59, v54)
                                                                }
                                                                if (!v18) {
                                                                    v108 := GenS10{F0: 42, F1: -128}
                                                                    _ = v108
                                                                } else {
                                                                    switch v80 {
                                                                    case 128:
                                                                        v54 = float64(1e10)
                                                                    case -1:
                                                                        v54 += v54
                                                                    case -338:
                                                                        fmt.Println(v80)
                                                                    default:
                                                                        var v109 uint = 0
                                                                        _ = v109
                                                                    }
                                                                    v106 = append(v106, v1)
                                                                }
                                                                var v110 uint32 = 322
                                                                switch 845 {
                                                                case 127:
                                                                    v111 := GenS10{F0: 128, F1: -1}
                                                                    _ = v111
                                                                case 128:
                                                                    if v18 {
                                                                        var v112 string = "foo bar"
                                                                        v59 = 32767
                                                                        v2[v110] = v18
                                                                        _ = v112
                                                                    }
                                                                }
                                                                if len(v106) > 0 {
                                                                    v106[0] = v1
                                                                }
                                                                _ = v110
                                                            case 32767:
                                                                v113 := GenS12{F0: 127, F1: (v18 || v18)}
                                                                _ = v113
                                                            }
                                                        }
                                                        fmt.Println(fi78)
                                                        _ = v106
                                                    case "foo bar":
                                                        v59 = v59
                                                    default:
                                                        v114 := GenS13{F0: 4294967295}
                                                        _ = v114
                                                    }
                                                    _ = v105
                                                case 1:
                                                    var v115 int = 9223372036854775807
                                                    _ = v115
                                                default:
                                                    v116 := GenS2{F0: "a", F1: ("\n\t" + (((("\n\t" + (("" + "foo bar") + "hello")) + ("hello" + ("hello" + "hello"))) + "a") + "\n\t"))}
                                                    _ = v116
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    var v117 uint32 = 2
                                                    _ = v117
                                                default:
                                                    fmt.Println(fi78)
                                                }
                                                if v18 {
                                                    var v118 uint = 675
                                                    v119 := []int{fi78, (-(fi78)), 263, (v56 * -158)}
                                                    if (3.14159 <= 0.0) {
                                                        v80 = (fi75 - v80)
                                                        var v120 string = "\n\t"
                                                        _ = v120
                                                    }
                                                    _ = v118
                                                    _ = v119
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if (514 == 4294967295) {
                                                            if (v54 < float64(0.0)) {
                                                                fmt.Println(v59, v80, v56)
                                                            } else {
                                                                var v121 int = v55
                                                                _ = v121
                                                            }
                                                            _ = len(v3)
                                                            v80 = v55
                                                        }
                                                    case "":
                                                        v122 := []int{v56, v1, (v80 - v80)}
                                                        var v123 int16 = 808
                                                        var v124 int8 = -17
                                                        _ = v122
                                                        _ = v123
                                                        _ = v124
                                                    case "hello":
                                                        var v125 bool = (!(0.0 > (0.0 * float32(v54))))
                                                        v126 := GenS1{F0: 610, F1: (^(v59))}
                                                        _ = v125
                                                        _ = v126
                                                    default:
                                                        v59 -= (v59 * (v59 >> ((2) & 63)))
                                                    }
                                                    v127 := GenS13{F0: 0}
                                                    if true {
                                                        fmt.Println(v96, fi78)
                                                        v128 := []int{(fi75 << ((42) & 63)), v80, -929, (fi75 >> ((255) & 63))}
                                                        if ((v55 | -1) > 42) {
                                                            v57 += (^(v57))
                                                        } else {
                                                            v129 := v2[65535]
                                                            v130 := GenS12{F0: 4294967295, F1: false}
                                                            v131 := GenS14{F0: int(v96), F1: -1}
                                                            _ = v129
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                        switch fi75 {
                                                        case 32768:
                                                            var v132 uint32 = 65535
                                                            _ = v132
                                                        default:
                                                            if (v57 == 128) {
                                                                var v133 uint8 = 161
                                                                _ = v133
                                                            }
                                                        }
                                                        v134 := GenS6{F0: 42, F1: v96, F2: int16(v96)}
                                                        _ = v128
                                                        _ = v134
                                                    } else {
                                                        var v135 int = 127
                                                        v135 = fi78
                                                        switch -141 {
                                                        case -2147483648:
                                                            fmt.Println(v80, v55, v56)
                                                        case 345:
                                                            var v136 uint8 = v57
                                                            _ = v136
                                                        case -318:
                                                            if (1 >= uint64(v57)) {
                                                                v59 = 65535
                                                                if true {
                                                                    v137 := v3[2147483647]
                                                                    _ = v137
                                                                } else {
                                                                    v59 += (v59 ^ v59)
                                                                    switch (("a" + ((("foo bar" + "a") + ("a" + (("a" + "\n\t") + "hello"))) + "")) + (("a" + (((("a" + "hello") + (("foo bar" + "foo bar") + ("hello" + "a"))) + ("a" + "")) + "hello")) + "\n\t")) {
                                                                    case "hello":
                                                                        v138 := []int{(v55 * -2147483648), (v55 | fi78), (v1 >> ((255) & 63)), v80}
                                                                        _ = v138
                                                                    }
                                                                }
                                                                if ((v54 * (v54 + v54)) < v54) {
                                                                    fmt.Println(v57)
                                                                } else {
                                                                    v3[genFn0(("foo bar" + ("a" + "hello")))] = "foo bar"
                                                                }
                                                            }
                                                        }
                                                        _ = v135
                                                    }
                                                    _ = v127
                                                }
                                                var v139 uint64 = 4294967295
                                                v140 := GenS4{F0: 42, F1: uint(v59)}
                                                if v18 {
                                                    v18 = v18
                                                }
                                                v141 := GenS3{F0: v139}
                                                if false {
                                                    var v142 int64 = -809
                                                    v143 := []int{int(v59), v56}
                                                    var v144 string = "a"
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                }
                                                var v145 uint64 = v139
                                                var v146 uint32 = uint32(v145)
                                                _ = v139
                                                _ = v140
                                                _ = v141
                                                _ = v145
                                                _ = v146
                                            }
                                        } else {
                                            var v147 uint8 = (v57 % (((v57 | v57)) | 1))
                                            _ = v147
                                        }
                                    default:
                                        fmt.Println(v59)
                                    }
                                }
                                v148 := GenS15{F0: (v96 - (^(v96))), F1: 860}
                                _ = v148
                                _ = v96
                            }
                        }
                        var v149 string = ""
                        _ = v149
                        _ = v80
                    }
                    v55 += 624
                    _ = fi78
                }
                _ = fi75
                _ = v76
                _ = v77
            }
        }
    default:
        var v150 uint16 = v59
        _ = v150
    }
    fmt.Println(v56, v18, v18)
    for fi151 := 0; fi151 < 9; fi151++ {
        var v152 float64 = ((0.1 - v54) + 0.0)
        for fi153 := 0; fi153 < 1; fi153++ {
            if v18 {
                fmt.Println(fi153, fi153, v152)
                if v18 {
                    var v154 uint8 = v57
                    if (((-2.25 * v152) * v54) <= 3.14159) {
                        var v155 int64 = 1
                        v156 := GenS16{F0: 1, F1: v57, F2: -1}
                        _ = v155
                        _ = v156
                    } else {
                        v54 = v54
                        switch "a" {
                        case "foo bar":
                            var v157 string = ("\n\t" + ("" + ""))
                            _ = v157
                        case "":
                            v158 := []int{-404}
                            _ = v158
                        }
                        var v159 uint8 = (^(v57))
                        v152 = (float64(float32((((v54 / v54) / 1.5) + ((0.1 + (v152 * v54)) / ((v54 * ((((v54 + (1e10 - v54)) / (v152 - v152)) / v152) / 1e10)) + v54))))) / v152)
                        _ = v159
                    }
                    if v18 {
                        v154 = v154
                    } else {
                        switch "" {
                        case "\n\t":
                            var v160 int8 = -1
                            _ = v160
                        }
                    }
                    v161 := GenS4{F0: 1, F1: 18446744073709551615}
                    _ = v154
                    _ = v161
                } else {
                    v18 = true
                }
                v18 = v18
            }
            switch (^(v55)) {
            case 2147483647:
                delete(v3, 0)
            case -294:
                var v162 float32 = (3.14159 - 0.0)
                _ = v162
            case 840:
                v163 := []int{fi151}
                _ = v163
            default:
                v164 := []int{v56, v1, (^(v55)), 42}
                _ = v164
            }
            v55 -= v55
            _ = fi153
        }
        _ = fi151
        _ = v152
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v59
}
