package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int; F1 uint; F2 uint64 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 int16; F1 uint16 }
type GenS12 struct { F0 int16; F1 int16; F2 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 string; F1 uint32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint16 }
func genFn0(p0 int, p1 int32, p2 uint32) int8 {
    var v6 float64 = 3.14159
    v7 := []int{int(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int32 = -813
    v2 := make(map[int]int)
    v3 := []int{-822}
    var v4 uint32 = 2147483648
    v5 := []uint32{2, v4, (v4 >> ((18446744073709551615) & 63)), v4}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int32 = (v1 << ((7) & 63))
        if true {
            var v10 int16 = 707
            switch (-(fi8)) {
            case 4294967295:
                if (v10 > (v10 % (((v10 | (^(v10)))) | 1))) {
                    fmt.Println(v9, v1, v9)
                } else {
                    var v11 float64 = float64(3.14159)
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (true || (!true)) {
                        v13 := v2[(^(fi8))]
                        var v14 int8 = 58
                        if false {
                            fmt.Println(v14, v1, v14)
                            delete(v2, (fi12 % ((-9223372036854775808) | 1)))
                        } else {
                            v15 := GenS1{F0: (v4 >> ((479) & 63))}
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    }
                    v9 -= v9
                    _ = fi12
                }
            }
            fmt.Println(v4, v1, fi8)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v4 -= (^(v4))
    switch -2147483648 {
    case 2147483647:
        if ((3.14159 + (0.0 + 0.1)) == (0.0 * float32(1.5))) {
            var v16 uint64 = 0
            _ = v16
        } else {
            v5 = append(v5, v4)
            switch 2147483647 {
            case 127:
                switch -390 {
                case -999:
                    if (false && ((^(v1)) <= 0)) {
                        switch -2147483648 {
                        case 0:
                            if false {
                                v4 = (v4 / ((v4) | 1))
                                switch (("a" + "\n\t") + "") {
                                case "foo bar":
                                    var v17 int8 = 0
                                    _ = v17
                                case "":
                                    v1 = v1
                                }
                            } else {
                                v4 = 0
                                v1 = (v1 >> ((288) & 63))
                                fmt.Println(v4, v4)
                            }
                        default:
                            switch 32768 {
                            case 0:
                                fmt.Println(v4, v4)
                            }
                        }
                    }
                case 127:
                    switch 9223372036854775807 {
                    case 2147483647:
                        v1 += v1
                    default:
                        if ((0.1 * -2.25) > (((float32(((1.5 - (1.5 - -2.25)) - 1e10)) / (1.5 + (((-2.25 - float32((((0.0 / float64((3.14159 + (((0.0 / 1e10) - (1.5 - 0.1)) - 3.14159)))) * (-2.25 / 1e10)) - ((3.14159 * 0.1) - float64(1e10))))) / float32(1.5)) * (1.5 * (1.5 + 1e10))))) * float32(3.14159)) / -2.25)) {
                            v4 = (v4 - 65535)
                            if (((-2.25 * ((-2.25 / 1.0) * ((0.1 / (float32(-2.25) - 1.5)) / 1e10))) - (1e10 + (-2.25 / 1e10))) > 3.14159) {
                                var v18 int64 = -1
                                v1 = v1
                                _ = v18
                            }
                        }
                    }
                case 32767:
                    switch (("foo bar" + "\n\t") + ("hello" + ("hello" + ""))) {
                    case "a":
                        v1 = -621
                    case "":
                        for fi19 := 0; fi19 < 1; fi19++ {
                            if true {
                                fmt.Println(v1, v4, v4)
                            } else {
                                v20 := GenS2{F0: (fi19 >> ((4294967295) & 63)), F1: 128, F2: v1}
                                _ = v20
                            }
                            v4 = v4
                            var v21 int64 = 9223372036854775807
                            _ = fi19
                            _ = v21
                        }
                    }
                }
            case -2147483648:
                var v22 string = ("foo bar" + "\n\t")
                _ = v22
            default:
                if (true && false) {
                    fmt.Println(v1)
                }
            }
            v4 += 612
        }
    case 4294967295:
        var v23 uint8 = 23
        _ = v23
    case 9223372036854775807:
        var v24 string = "\n\t"
        _ = v24
    default:
        for fi25 := 0; fi25 < 12; fi25++ {
            for fi26 := 0; fi26 < 19; fi26++ {
                fmt.Println(fi25)
                _ = fi26
            }
            v27 := v2[fi25]
            _ = fi25
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int64 = 1
        switch ("\n\t" + "") {
        case "":
            v30 := GenS3{F0: -128}
            _ = v30
        default:
            var v31 int32 = (-(v1))
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    if (9223372036854775807 > 42) {
        fmt.Println(v1)
    } else {
        switch "" {
        case "a":
            fmt.Println(v1)
        case "\n\t":
            if false {
                var v32 float64 = (-2.25 - (1e10 - 1.5))
                v1 += v1
                var v33 int64 = 42
                _ = v32
                _ = v33
            }
        }
        if true {
            var v34 uint64 = 0
            switch ("foo bar" + "foo bar") {
            case "a":
                var v35 int8 = 100
                _ = v35
            default:
                delete(v2, 4294967295)
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                v1 -= -1
                _ = fi36
            }
            _ = v34
        }
        switch -773 {
        case 32767:
            for fi37 := 0; fi37 < 6; fi37++ {
                v4 += (v4 / ((v4) | 1))
                if true {
                    if (157 != 1) {
                        v2[(fi37 << ((42) & 63))] = fi37
                    }
                }
                v38 := GenS4{F0: 0, F1: 42}
                _ = fi37
                _ = v38
            }
        case 2147483647:
            fmt.Println(v1)
        case 128:
            v4 = 4294967295
        default:
            var v39 float32 = (1e10 / 1.0)
            _ = v39
        }
    }
    if ("foo bar" != "a") {
        var v40 int16 = 1000
        var v41 int16 = v40
        v42 := GenS5{F0: 32767, F1: 255, F2: 42}
        _ = v40
        _ = v41
        _ = v42
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        v44 := GenS6{F0: (fi43 << ((65535) & 63)), F1: v4}
        for fi45 := 0; fi45 < 9; fi45++ {
            switch 815 {
            case 32768:
                if (!false) {
                    var v46 float32 = float32(3.14159)
                    fmt.Println(fi45)
                    var v47 int64 = -226
                    _ = v46
                    _ = v47
                } else {
                    switch (fi45 - 128) {
                    case 0:
                        switch "" {
                        case "a":
                            var v48 uint32 = v4
                            _ = v48
                        case "foo bar":
                            if false {
                                _ = len(v2)
                                v49 := GenS1{F0: v4}
                                if ((fi43 ^ (fi45 | fi43)) < (fi43 ^ (^(fi45)))) {
                                    v50 := GenS7{F0: -1}
                                    if false {
                                        var v51 bool = true
                                        switch (("a" + "\n\t") + "\n\t") {
                                        case "hello":
                                            var v52 string = ("" + "a")
                                            _ = v52
                                        case "\n\t":
                                            v4 = 65535
                                        }
                                        v4 -= (v4 % ((v4) | 1))
                                        _ = v51
                                    } else {
                                        var v53 int64 = 9223372036854775807
                                        var v54 uint32 = (v4 >> ((65535) & 63))
                                        var v55 int = (fi43 + (fi45 / ((fi43) | 1)))
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 float32 = 0.0
                                    _ = v50
                                    _ = v56
                                }
                                _ = v49
                            }
                            v57 := GenS8{F0: 337, F1: 251, F2: 1000}
                            _ = v57
                        }
                    default:
                        fmt.Println(fi43)
                    }
                    fmt.Println(fi45)
                }
            case 4294967295:
                if false {
                    if (485 <= 426) {
                        v4 = v4
                        v58 := GenS8{F0: 105, F1: 128, F2: -520}
                        _ = v58
                    } else {
                        var v59 float64 = (-2.25 + 1.5)
                        v1 = (v1 >> ((675) & 63))
                        _ = v59
                    }
                    v60 := GenS9{F0: 428}
                    _ = v60
                } else {
                    v61 := v2[fi43]
                    var v62 float32 = -2.25
                    switch fi43 {
                    case 848:
                        if (255 > 606) {
                            v4 -= v4
                            v63 := GenS2{F0: 0, F1: 0, F2: 0}
                            if (!false) {
                                v64 := GenS8{F0: 4294967295, F1: 255, F2: 1}
                                var v65 uint16 = 138
                                _ = v64
                                _ = v65
                            }
                            _ = v63
                        } else {
                            v66 := GenS10{F0: 127, F1: 255}
                            _ = v66
                        }
                    case -1:
                        var v67 uint64 = 266
                        _ = v67
                    default:
                        switch 32768 {
                        case 128:
                            v68 := GenS11{F0: 42, F1: 2}
                            _ = v68
                        default:
                            switch fi45 {
                            case -2147483648:
                                v62 = v62
                                var v69 uint32 = v4
                                _ = v69
                            case -1:
                                fmt.Println(fi45)
                            }
                        }
                    }
                    _ = v61
                    _ = v62
                }
            default:
                v4 -= (v4 % ((v4) | 1))
            }
            v4 += (v4 >> ((745) & 63))
            _ = fi45
        }
        for fi70 := 0; fi70 < 5; fi70++ {
            switch "\n\t" {
            case "":
                if false {
                    v71 := GenS9{F0: v4}
                    delete(v2, 4294967295)
                    _ = v71
                }
            case "hello":
                fmt.Println(fi70)
            }
            switch fi43 {
            case 32767:
                var v72 bool = ("foo bar" > ("\n\t" + "hello"))
                _ = v72
            default:
                var v73 bool = false
                _ = v73
            }
            if ((v4 >= v4) || ((!(!true)) && false)) {
                var v74 uint64 = 4294967295
                fmt.Println(v74, v74, fi43)
                var v75 int16 = 0
                _ = v74
                _ = v75
            }
            _ = fi70
        }
        _ = fi43
        _ = v44
    }
    if (49 < 325) {
        v4 -= 2147483647
        for fi76 := 0; fi76 < 14; fi76++ {
            if (-1 != -321) {
                v77 := v2[(fi76 ^ fi76)]
                _ = v77
            } else {
                v78 := GenS11{F0: 42, F1: 0}
                if false {
                    switch (fi76 >> ((857) & 63)) {
                    case -9223372036854775808:
                        switch "hello" {
                        case "foo bar":
                            var v79 float32 = (-2.25 + 0.0)
                            _ = v79
                        case "":
                            if true {
                                fmt.Println(v4)
                                for fi80 := 0; fi80 < 1; fi80++ {
                                    if (-941 > 0) {
                                        fmt.Println(v1, fi76)
                                        var v81 int = fi80
                                        var v82 int = -388
                                        _ = v81
                                        _ = v82
                                    }
                                    v83 := GenS8{F0: 2, F1: 2, F2: 0}
                                    _ = fi80
                                    _ = v83
                                }
                            } else {
                                _ = len(v2)
                            }
                        case "\n\t":
                            for fi84 := 0; fi84 < 8; fi84++ {
                                if len(v3) > 0 {
                                    v3[0] = 603
                                }
                                if (42 != 106) {
                                    v1 = (v1 | (v1 - v1))
                                    v3 = append(v3, fi84)
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(fi76, fi84, fi76)
                                    case "foo bar":
                                        var v85 string = ("hello" + "\n\t")
                                        v86 := v2[(fi84 * fi76)]
                                        var v87 uint32 = v4
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                    var v88 uint16 = 42
                                    _ = v88
                                } else {
                                    v89 := GenS8{F0: 202, F1: 127, F2: -27}
                                    v4 += (^(v4))
                                    _ = v89
                                }
                                _ = fi84
                            }
                        default:
                            if true {
                                if len(v5) > 0 {
                                    v5[0] = v4
                                }
                            }
                        }
                    case 128:
                        switch (fi76 ^ (fi76 % ((32768) | 1))) {
                        case 128:
                            v1 = 0
                        case 2147483647:
                            for fi90 := 0; fi90 < 8; fi90++ {
                                v4 = v4
                                v1 -= (^(v1))
                                _ = fi90
                            }
                        default:
                            for fi91 := 0; fi91 < 14; fi91++ {
                                v1 += (^(v1))
                                fmt.Println(fi76)
                                v4 = (v4 ^ 4294967295)
                                _ = fi91
                            }
                        }
                    }
                    for fi92 := 0; fi92 < 7; fi92++ {
                        v93 := GenS10{F0: 65535, F1: 0}
                        v94 := v2[(^(fi76))]
                        v95 := GenS4{F0: 32767, F1: 65535}
                        v96 := GenS8{F0: 1, F1: 21, F2: 935}
                        _ = fi92
                        _ = v93
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                    if (!true) {
                        v97 := v2[fi76]
                        var v98 bool = ((v4 << ((2) & 63)) < v4)
                        _ = v97
                        _ = v98
                    }
                }
                v99 := GenS12{F0: -1, F1: -32768, F2: ("a" + "a")}
                _ = v78
                _ = v99
            }
            _ = len(v2)
            for fi100 := 0; fi100 < 17; fi100++ {
                v1 -= (v1 - (v1 - v1))
                v101 := GenS13{F0: (-129 < 42)}
                v102 := v2[(fi100 + fi100)]
                _ = fi100
                _ = v101
                _ = v102
            }
            _ = fi76
        }
        fmt.Println(v1, v4)
    }
    v4 = 2
    _ = len(v5)
    v4 = (^(v4))
    switch -1 {
    case 0:
        for fi103 := 0; fi103 < 2; fi103++ {
            fmt.Println(v4, v4)
            _ = fi103
        }
    case 32768:
        v104 := GenS14{F0: (((("foo bar" + "\n\t") + "hello") + ("hello" + "\n\t")) + ("foo bar" + "foo bar")), F1: v4}
        _ = v104
    case 128:
        fmt.Println(v1)
    default:
        if (("" + ("foo bar" + ("" + ("foo bar" + "foo bar")))) != "a") {
            v1 = v1
            if true {
                if false {
                    if false {
                        fmt.Println(v1, v4)
                        var v105 uint64 = 1
                        _ = v105
                    }
                    fmt.Println(v1)
                } else {
                    v106 := GenS15{F0: 2147483647}
                    v4 = 2147483647
                    switch 9223372036854775807 {
                    case 32768:
                        switch ("\n\t" + "hello") {
                        case "":
                            var v107 int32 = v1
                            _ = v107
                        case "foo bar":
                            fmt.Println(v1, v4)
                        }
                    }
                    _ = v106
                }
                switch (("hello" + "foo bar") + "") {
                case "foo bar":
                    v108 := GenS16{F0: 117}
                    _ = v108
                case "":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    for fi109 := 0; fi109 < 13; fi109++ {
                        var v110 uint64 = 1
                        _ = fi109
                        _ = v110
                    }
                }
            }
        } else {
            v4 = 0
        }
    }
    v1 += (v1 << ((172) & 63))
    for fi111 := 0; fi111 < 11; fi111++ {
        v3 = append(v3, fi111)
        _ = fi111
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
