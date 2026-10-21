package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int32; F2 string }
type GenS5 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 bool; F1 int32; F2 int32 }
type GenS10 struct { F0 int16 }
func genFn0() int {
    var v7 int = -1
    v8 := []int{v7, -9223372036854775808, (v7 * -2147483648), v7}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v7
    _ = v8
    return v7
}
func genFn1() bool {
    v9 := []int{-1}
    v10 := GenS1{F0: 32767}
    var v11 int = 0
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            var v13 float64 = 0.1
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v11, fi14, fi14)
            var v15 uint = 65535
            fmt.Println(v11)
            _ = fi14
            _ = v15
        }
        if len(v9) > 0 {
            v9[0] = v11
        }
        _ = fi12
    }
    if (100 < 18446744073709551615) {
        if len(v9) > 0 {
            v9[0] = (v11 + (v11 % ((v11) | 1)))
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return (!(127 == 93))
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]string)
    v3 := []string{("\n\t" + "foo bar")}
    var v4 uint16 = 32767
    var v5 uint8 = 128
    var v6 bool = (-32768 != 0)
    switch (("hello" + "a") + "a") {
    case "foo bar":
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := GenS2{F0: 65535, F1: 357}
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS3{F0: 230}
                var v21 uint64 = 1
                var v22 string = ("a" + "a")
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi16
        }
    case "\n\t":
        if true {
            v1 += 15
        }
    }
    v23 := GenS2{F0: uint32(v4), F1: 592}
    switch "hello" {
    case "hello":
        var v24 uint16 = v4
        _ = v24
    case "\n\t":
        var v25 uint8 = v5
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v1 += 42
        _ = fi26
    }
    _ = len(v2)
    if len(v3) > 0 {
        v3[0] = (("" + ("foo bar" + "a")) + "a")
    }
    v5 -= v5
    v6 = v6
    if (v6 && v6) {
        if true {
            v27 := GenS2{F0: 2147483647, F1: 2}
            v28 := GenS2{F0: 197, F1: 42}
            _ = v27
            _ = v28
        } else {
            if (0 <= -1) {
                fmt.Println(v4, v1, v6)
            } else {
                if (true && (!v6)) {
                    var v29 int = genFn0()
                    _ = v29
                }
            }
            var v30 uint = 42
            _ = v30
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            _ = len(v2)
            v4 -= v4
            _ = fi31
        }
    } else {
        v6 = (128 < (v5 | 198))
        if v6 {
            v6 = v6
            switch 128 {
            case -1:
                if v6 {
                    if false {
                        switch 9223372036854775807 {
                        case -929:
                            var v32 int64 = -450
                            _ = v32
                        case 415:
                            if (uint8(v4) >= (v5 | 151)) {
                                fmt.Println(v6, v5)
                            } else {
                                v5 -= v5
                                if (!(0 == 925)) {
                                    for fi33 := 0; fi33 < 16; fi33++ {
                                        v5 = v5
                                        v34 := GenS4{F0: -1, F1: 0, F2: "foo bar"}
                                        _ = fi33
                                        _ = v34
                                    }
                                } else {
                                    switch 1 {
                                    case 32768:
                                        if v6 {
                                            switch "\n\t" {
                                            case "hello":
                                                v1 = 127
                                            case "a":
                                                var v35 int16 = -1
                                                _ = v35
                                            }
                                            v6 = v6
                                            fmt.Println(v1)
                                        }
                                    case 42:
                                        for fi36 := 0; fi36 < 1; fi36++ {
                                            if (2147483647 != 2147483647) {
                                                if true {
                                                    if false {
                                                        for fi37 := 0; fi37 < 18; fi37++ {
                                                            v38 := GenS2{F0: 2147483647, F1: 2}
                                                            _ = fi37
                                                            _ = v38
                                                        }
                                                    }
                                                    v5 += (v5 ^ v5)
                                                    switch (-(fi36)) {
                                                    case -1:
                                                        var v39 uint32 = 9
                                                        _ = v39
                                                    case 2147483647:
                                                        v40 := GenS1{F0: 32767}
                                                        _ = v40
                                                    }
                                                }
                                                if v6 {
                                                    var v41 uint8 = v5
                                                    _ = v41
                                                } else {
                                                    fmt.Println(v1)
                                                }
                                                for fi42 := 0; fi42 < 10; fi42++ {
                                                    fmt.Println(v6)
                                                    _ = fi42
                                                }
                                            }
                                            for fi43 := 0; fi43 < 17; fi43++ {
                                                if v6 {
                                                    if (float64((((((3.14159 / float32(0.1)) * (-2.25 / -2.25)) - 3.14159) - 0.0) - ((1e10 / ((-2.25 - ((3.14159 * ((0.0 + (0.0 - (((-2.25 - 1.5) / 3.14159) - (1.5 / -2.25)))) * (1e10 + 0.0))) + 3.14159)) * (3.14159 * 0.1))) * -2.25))) <= 3.14159) {
                                                        var v44 uint64 = 782
                                                        v2[2147483647] = ("foo bar" + ("\n\t" + ""))
                                                        _ = v44
                                                    }
                                                    fmt.Println(v6, v6, v5)
                                                } else {
                                                    if false {
                                                        var v45 uint8 = v5
                                                        _ = v45
                                                    }
                                                    var v46 int8 = v1
                                                    v4 = v4
                                                    v46 -= (v1 & (v1 & v1))
                                                    _ = v46
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    if (65535 < 2147483648) {
                                                        v6 = (true && false)
                                                        v6 = (!v6)
                                                    }
                                                }
                                                var v47 int32 = 799
                                                if v6 {
                                                    if ((("\n\t" + "a") + "") != (("a" + "a") + ("\n\t" + (("foo bar" + "hello") + "hello")))) {
                                                        switch -1 {
                                                        case 128:
                                                            v6 = v6
                                                        }
                                                        var v48 bool = ((true || (v6 && true)) || (v5 > (v5 ^ (v5 | 180))))
                                                        var v49 string = "a"
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    var v50 int = 42
                                                    fmt.Println(v5, v5)
                                                    _ = v50
                                                } else {
                                                    v4 = (v4 / ((411) | 1))
                                                }
                                                v2[2147483648] = "foo bar"
                                                fmt.Println(v4, fi43)
                                                _ = fi43
                                                _ = v47
                                            }
                                            var v51 string = "foo bar"
                                            _ = fi36
                                            _ = v51
                                        }
                                    case 1:
                                        switch 42 {
                                        case 131:
                                            fmt.Println(v5, v4)
                                        }
                                    }
                                    switch "hello" {
                                    case "a":
                                        var v52 float32 = (0.0 + (float32(3.14159) * 1e10))
                                        _ = v52
                                    default:
                                        v5 -= v5
                                    }
                                }
                                for fi53 := 0; fi53 < 11; fi53++ {
                                    v4 -= 34
                                    for fi54 := 0; fi54 < 9; fi54++ {
                                        v6 = false
                                        _ = fi54
                                    }
                                    var v55 int = fi53
                                    _ = fi53
                                    _ = v55
                                }
                            }
                        }
                        v5 += v5
                    } else {
                        switch -231 {
                        case 32767:
                            if ((!v6) || true) {
                                fmt.Println(v1, v6)
                            }
                        case 4294967295:
                            fmt.Println(v5, v6)
                        case 0:
                            var v56 int64 = 0
                            _ = v56
                        default:
                            var v57 uint = 18446744073709551615
                            _ = v57
                        }
                        v6 = (((!(!v6)) || (255 > 0)) && true)
                    }
                    v6 = (1 <= 42)
                    v5 += (v5 + (v5 ^ 233))
                } else {
                    var v58 int8 = 127
                    var v59 int = -1
                    _ = v58
                    _ = v59
                }
            case -9223372036854775808:
                if (v5 != v5) {
                    v1 = v1
                    v60 := GenS3{F0: 604}
                    _ = v60
                } else {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        for fi62 := 0; fi62 < 11; fi62++ {
                            v5 += 2
                            if v6 {
                                delete(v2, 1)
                                var v63 float64 = (((0.0 + 1.5) - 1.5) - -2.25)
                                _ = v63
                            }
                            _ = fi62
                        }
                        v64 := GenS5{F0: 587, F1: 564, F2: 4294967295}
                        v1 += (v1 & (v1 ^ v1))
                        _ = fi61
                        _ = v64
                    }
                    v4 -= v4
                    fmt.Println(v4)
                }
            case -375:
                switch "\n\t" {
                case "a":
                    v6 = (v1 <= -75)
                }
            }
            switch 522 {
            case 42:
                for fi65 := 0; fi65 < 7; fi65++ {
                    var v66 int32 = 1
                    v67 := GenS5{F0: (-(v66)), F1: 0, F2: 198}
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            case 9223372036854775807:
                var v68 int8 = (v1 | 127)
                _ = v68
            case 32768:
                v69 := GenS6{F0: -761}
                _ = v69
            }
        } else {
            switch ("" + "foo bar") {
            case "foo bar":
                v70 := GenS2{F0: 2147483648, F1: 2147483647}
                _ = v70
            case "a":
                fmt.Println(v4, v4)
            }
        }
    }
    if v6 {
        v71 := GenS3{F0: 0}
        for fi72 := 0; fi72 < 5; fi72++ {
            v73 := v2[1]
            var v74 uint16 = (v4 ^ (v4 % (((^(v4))) | 1)))
            _ = fi72
            _ = v73
            _ = v74
        }
        _ = v71
    } else {
        v6 = v6
    }
    v75 := GenS7{F0: 2}
    v4 = v4
    if (true || (127 == v5)) {
        fmt.Println(v1, v6, v4)
        switch "a" {
        case "a":
            var v76 uint32 = 2147483647
            _ = v76
        case "hello":
            v2[42] = "\n\t"
        case "foo bar":
            for fi77 := 0; fi77 < 7; fi77++ {
                var v78 int16 = -173
                var v79 uint64 = 0
                _ = fi77
                _ = v78
                _ = v79
            }
        default:
            for fi80 := 0; fi80 < 5; fi80++ {
                if (3.14159 == (float64(float32(((-2.25 * 3.14159) + (1e10 / 1.5)))) - 0.0)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if v6 {
                        var v81 int16 = 1000
                        v4 = v4
                        switch fi80 {
                        case 751:
                            var v82 uint64 = 201
                            _ = v82
                        }
                        _ = v81
                    } else {
                        for fi83 := 0; fi83 < 20; fi83++ {
                            v84 := v2[42]
                            _ = fi83
                            _ = v84
                        }
                    }
                } else {
                    fmt.Println(v5, v4)
                }
                if true {
                    v85 := GenS8{F0: 4294967295, F1: -32768}
                    fmt.Println(v1, fi80, fi80)
                    v4 = (v4 | v4)
                    _ = v85
                } else {
                    var v86 int64 = 9223372036854775807
                    if v6 {
                        v87 := GenS9{F0: (!(1e10 > 0.1)), F1: -134, F2: 0}
                        _ = v87
                    } else {
                        var v88 int8 = v1
                        switch (("foo bar" + ("a" + (("a" + "a") + ""))) + "foo bar") {
                        case "hello":
                            if (1 < -431) {
                                var v89 int64 = (-(v86))
                                v88 = (v1 << ((2) & 63))
                                _ = v89
                            }
                        case "foo bar":
                            v90 := GenS9{F0: (float64(1e10) != -2.25), F1: 1, F2: 0}
                            _ = v90
                        case "a":
                            if false {
                                var v91 int64 = 42
                                var v92 float32 = 3.14159
                                var v93 int8 = (^(v1))
                                _ = v91
                                _ = v92
                                _ = v93
                            } else {
                                fmt.Println(v86, v1)
                                switch "hello" {
                                case "":
                                    var v94 bool = ((fi80 & 0) != (fi80 / ((genFn0()) | 1)))
                                    _ = v94
                                case "foo bar":
                                    for fi95 := 0; fi95 < 12; fi95++ {
                                        fmt.Println(fi80, v5, v1)
                                        _ = fi95
                                    }
                                default:
                                    v96 := GenS10{F0: 719}
                                    _ = v96
                                }
                            }
                        }
                        _ = v88
                    }
                    v1 += int8(v86)
                    _ = v86
                }
                switch ("\n\t" + "hello") {
                case "a":
                    v97 := GenS4{F0: int32(v1), F1: 2147483647, F2: "hello"}
                    _ = v97
                case "hello":
                    switch (^(fi80)) {
                    case 128:
                        if (0 < 1) {
                            fmt.Println(v6, v1, v5)
                            for fi98 := 0; fi98 < 1; fi98++ {
                                switch fi98 {
                                case 2147483647:
                                    if true {
                                        if (v4 != (^(v4))) {
                                            v1 -= (v1 - v1)
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                        }
                                    }
                                }
                                _ = fi98
                            }
                            switch genFn0() {
                            case -9223372036854775808:
                                v5 -= v5
                            default:
                                v6 = true
                            }
                        } else {
                            _ = len(v3)
                            for fi99 := 0; fi99 < 10; fi99++ {
                                switch fi80 {
                                case -1:
                                    if (1.5 != ((0.0 + 3.14159) + float32(1.5))) {
                                        v100 := GenS1{F0: 0}
                                        _ = v100
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = fi99
                            }
                        }
                    case 9223372036854775807:
                        v6 = v6
                    case 343:
                        var v101 uint64 = 1
                        _ = v101
                    default:
                        switch "hello" {
                        case "":
                            var v102 uint32 = 2
                            _ = v102
                        case "hello":
                            var v103 bool = (32767 > 776)
                            _ = v103
                        case "foo bar":
                            if v6 {
                                v4 += 1
                                var v104 float64 = -2.25
                                _ = v104
                            }
                        }
                    }
                }
                _ = fi80
            }
        }
        for fi105 := 0; fi105 < 18; fi105++ {
            var v106 uint32 = 4294967295
            fmt.Println(v4, v6, v1)
            _ = fi105
            _ = v106
        }
    }
    var v107 int8 = (v1 << ((475) & 63))
    for fi108 := 0; fi108 < 11; fi108++ {
        switch fi108 {
        case -9223372036854775808:
            v4 -= (v4 << ((4294967295) & 63))
        case 2147483647:
            for fi109 := 0; fi109 < 7; fi109++ {
                v110 := GenS10{F0: 42}
                v4 += v4
                _ = fi109
                _ = v110
            }
        default:
            var v111 uint32 = 2147483648
            _ = v111
        }
        if (!v6) {
            v1 = 75
        } else {
            var v112 int = fi108
            var v113 float32 = 3.14159
            for fi114 := 0; fi114 < 9; fi114++ {
                v4 -= v4
                _ = len(v2)
                _ = fi114
            }
            _ = v112
            _ = v113
        }
        _ = fi108
    }
    for fi115 := 0; fi115 < 1; fi115++ {
        for fi116 := 0; fi116 < 4; fi116++ {
            delete(v2, 42)
            _ = fi116
        }
        _ = fi115
    }
    switch 0 {
    case 371:
        for fi117 := 0; fi117 < 14; fi117++ {
            v6 = false
            fmt.Println(v5, v107, v1)
            _ = fi117
        }
    case 127:
        _ = len(v2)
    case 623:
        var v118 uint16 = (v4 & 0)
        _ = v118
    }
    v119 := v2[0]
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v107
    _ = v119
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
}
