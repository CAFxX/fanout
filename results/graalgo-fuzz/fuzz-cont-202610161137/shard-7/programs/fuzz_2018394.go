package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string; F2 uint64 }
type GenS5 struct { F0 bool; F1 int16; F2 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int64; F1 uint64; F2 int8 }
func genFn0() uint16 {
    var v5 string = ("" + "hello")
    v5 = v5
    fmt.Println(v5)
    if true {
        v5 = ((v5 + "a") + v5)
        switch (("foo bar" + v5) + v5) {
        case "\n\t":
            v5 = v5
        case "a":
            v5 = (v5 + v5)
        case "foo bar":
            switch 228 {
            case 93:
                for fi6 := 0; fi6 < 19; fi6++ {
                    var v7 uint16 = 2
                    _ = fi6
                    _ = v7
                }
            }
        }
    }
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1e10
    v3 := []uint32{448, 444, 42}
    var v4 string = ("foo bar" + "\n\t")
    v4 = ((v4 + (("foo bar" + v4) + (v4 + v4))) + "\n\t")
    if (684 == 65535) {
        var v8 bool = true
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = v4
            v11 := v1[1]
            v12 := GenS1{F0: -2147483648}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 uint32 = 2147483647
        _ = v13
        _ = v8
    } else {
        if false {
            fmt.Println(v2, v2)
            switch -823 {
            case 1:
                v3 = append(v3, 2147483648)
            default:
                fmt.Println(v4, v2)
            }
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS1{F0: 2147483647}
                var v16 uint = 0
                _ = fi14
                _ = v15
                _ = v16
            }
            if true {
                delete(v1, 2)
                v17 := GenS2{F0: 42}
                _ = v17
            } else {
                if (!(v4 > v4)) {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 int16 = 1
                        fmt.Println(v2, v4)
                        fmt.Println(fi18, v4)
                        _ = fi18
                        _ = v19
                    }
                } else {
                    fmt.Println(v4, v2)
                }
            }
        }
        v2 -= (0.1 + (v2 - v2))
    }
    v4 = v4
    if (false && ((!(3.14159 < 1e10)) && false)) {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint64 = 897
            _ = fi20
            _ = v21
        }
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        var v22 int8 = 0
        _ = v22
    }
    if (!false) {
        var v23 uint64 = 16
        var v24 bool = ((332 <= 1000) && true)
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 18446744073709551615}
    v4 = "foo bar"
    if false {
        fmt.Println(v4, v4)
    } else {
        fmt.Println(v2)
        v26 := v1[4294967295]
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = v26
    }
    switch ((v4 + v4) + (v4 + (v4 + "hello"))) {
    case "hello":
        if false {
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(fi27, fi27)
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 uint16 = genFn0()
                    v30 := GenS1{F0: -2147483648}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                if true {
                    var v31 uint = 265
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 -= (float64(0.0) / v2)
                    _ = v31
                } else {
                    if (!(!true)) {
                        switch (-(fi27)) {
                        case 758:
                            for fi32 := 0; fi32 < 16; fi32++ {
                                switch (fi32 * fi27) {
                                case 128:
                                    fmt.Println(v2)
                                    var v33 bool = ((!true) && (((v4 < (v4 + v4)) || ((255 != 4294967295) || (-2147483648 == -2147483648))) && (18446744073709551615 < 0)))
                                    switch fi27 {
                                    case 127:
                                        if v33 {
                                            v2 += v2
                                            var v34 int16 = -819
                                            switch v4 {
                                            case "hello":
                                                v35 := GenS2{F0: 18446744073709551615}
                                                _ = v35
                                            case "foo bar":
                                                fmt.Println(v4)
                                            case "":
                                                _ = len(v3)
                                            }
                                            _ = v34
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "":
                                            switch v4 {
                                            case "":
                                                if (!v33) {
                                                    v36 := GenS2{F0: 38}
                                                    v37 := GenS3{F0: fi27, F1: 546}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            default:
                                                switch (v4 + ("foo bar" + v4)) {
                                                case "foo bar":
                                                    v4 = "\n\t"
                                                case "\n\t":
                                                    switch v4 {
                                                    case "a":
                                                        if v33 {
                                                            fmt.Println(v33, fi32)
                                                            switch (fi32 & (fi32 / ((-986) | 1))) {
                                                            case 42:
                                                                var v38 int = (fi27 * fi27)
                                                                _ = v38
                                                            case 9223372036854775807:
                                                                if true {
                                                                    var v39 float64 = 1.5
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v39 = ((float64((1.5 - 0.0)) + float64(((0.1 / (1e10 + (3.14159 / 0.1))) - 3.14159))) + v39)
                                                                    _ = v39
                                                                }
                                                            default:
                                                                switch 4294967295 {
                                                                case 808:
                                                                    v40 := GenS4{F0: 2, F1: ("hello" + "hello"), F2: 492}
                                                                    _ = v40
                                                                case -2147483648:
                                                                    if v33 {
                                                                        if v33 {
                                                                            fmt.Println(v2)
                                                                            if (2 != 42) {
                                                                                v33 = (62 != 0)
                                                                                fmt.Println(fi27)
                                                                            } else {
                                                                                var v41 string = v4
                                                                                var v42 float64 = (-2.25 + float64(3.14159))
                                                                                fmt.Println(fi32, v2)
                                                                                v4 = "\n\t"
                                                                                _ = v41
                                                                                _ = v42
                                                                            }
                                                                        }
                                                                        var v43 bool = true
                                                                        _ = v43
                                                                    }
                                                                case 42:
                                                                    v44 := GenS4{F0: 4294967295, F1: (v4 + ((((v4 + v4) + (v4 + v4)) + ((v4 + v4) + "a")) + "foo bar")), F2: 4294967295}
                                                                    _ = v44
                                                                }
                                                            }
                                                            if (v33 && ((v2 / (v2 * (1e10 / v2))) < (((0.0 * v2) / float64(((((float32(float64(float32(v2))) * 1e10) / ((3.14159 / -2.25) + float32(-2.25))) - (1e10 - float32(0.1))) / ((3.14159 + (3.14159 * 3.14159)) - (0.0 * 3.14159))))) * 1.5))) {
                                                                v33 = v33
                                                                if v33 {
                                                                    if v33 {
                                                                        if (32767 < (fi32 >> ((65535) & 63))) {
                                                                            v45 := GenS5{F0: (!v33), F1: 32767, F2: 1}
                                                                            v33 = (22 <= 255)
                                                                            _ = v45
                                                                        }
                                                                        var v46 uint32 = 0
                                                                        _ = v46
                                                                    }
                                                                    if (42 != 0) {
                                                                        v33 = (!(!((fi27 + (fi27 - fi32)) >= fi27)))
                                                                        if (v33 && v33) {
                                                                            v47 := v1[482]
                                                                            fmt.Println(fi32, fi27)
                                                                            var v48 int32 = int32(fi32)
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            if v33 {
                                                                                if (298 > -152) {
                                                                                    v49 := GenS3{F0: 815, F1: 4294967295}
                                                                                    if v33 {
                                                                                        var v50 uint32 = 1
                                                                                        fmt.Println(v2, v2)
                                                                                        _ = v50
                                                                                    }
                                                                                    if (false || ((fi32 & (fi27 << ((1) & 63))) <= (fi32 % (((fi27 >> ((122) & 63))) | 1)))) {
                                                                                        v51 := GenS5{F0: (42 < 128), F1: 42, F2: fi32}
                                                                                        switch 4294967295 {
                                                                                        case 274:
                                                                                            v33 = v33
                                                                                        default:
                                                                                            v4 = ("foo bar" + (("foo bar" + v4) + v4))
                                                                                        }
                                                                                        v1[1] = v33
                                                                                        delete(v1, 2)
                                                                                        v52 := GenS5{F0: v33, F1: 1000, F2: fi27}
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                    }
                                                                                    _ = v49
                                                                                } else {
                                                                                    var v53 int = fi32
                                                                                    switch v4 {
                                                                                    case "foo bar":
                                                                                        switch (fi32 * (fi32 ^ (fi32 >> ((460) & 63)))) {
                                                                                        case 42:
                                                                                            var v54 float64 = (float64((1.5 / 0.1)) + 0.0)
                                                                                            _ = v54
                                                                                        default:
                                                                                            _ = len(v3)
                                                                                        }
                                                                                    }
                                                                                    _ = v53
                                                                                }
                                                                                v4 = ""
                                                                            }
                                                                        }
                                                                        v3 = append(v3, 2147483648)
                                                                    } else {
                                                                        _ = len(v1)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 2147483647
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v55 uint8 = 217
                                                                    _ = v55
                                                                }
                                                            } else {
                                                                var v56 uint8 = 42
                                                                fmt.Println(v2)
                                                                var v57 int64 = 0
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if (!(v4 != v4)) {
                                                            fmt.Println(v2)
                                                            v58 := GenS2{F0: 0}
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v59 int = (^(fi27))
                                                                _ = v59
                                                            case "hello":
                                                                v60 := GenS4{F0: 4294967295, F1: v4, F2: 18446744073709551615}
                                                                _ = v60
                                                            case "":
                                                                v2 -= float64(1e10)
                                                            }
                                                            var v61 uint16 = 32767
                                                            _ = v58
                                                            _ = v61
                                                        } else {
                                                            fmt.Println(v4, v33, fi32)
                                                        }
                                                    }
                                                case "a":
                                                    if (((false || (v33 && (-448 <= -331))) || ((1 > 2) && (42 == -1))) || ((42 != 1) || (!v33))) {
                                                        v62 := GenS5{F0: (v33 && (169 == 0)), F1: -15, F2: (-(fi27))}
                                                        fmt.Println(v33)
                                                        _ = v62
                                                    } else {
                                                        v63 := GenS4{F0: 163, F1: (((v4 + (v4 + "a")) + v4) + v4), F2: 42}
                                                        v64 := GenS6{F0: 42}
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                                if v33 {
                                                    v65 := GenS1{F0: -287}
                                                    _ = v65
                                                } else {
                                                    v33 = (2 != 42)
                                                    v4 = ("" + ("" + "a"))
                                                    v66 := GenS7{F0: 2, F1: v33}
                                                    _ = v66
                                                }
                                            }
                                        }
                                    }
                                    v33 = (32767 < 32767)
                                    _ = v33
                                case 32768:
                                    fmt.Println(fi27, fi32, fi27)
                                case 4294967295:
                                    if true {
                                        v1[2] = false
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                        v2 += (-2.25 + v2)
                                    }
                                default:
                                    switch 4294967295 {
                                    case 503:
                                        var v67 int64 = 0
                                        _ = v67
                                    default:
                                        v4 = "hello"
                                    }
                                }
                                if true {
                                    if ((42 != -1) || false) {
                                        _ = len(v3)
                                    } else {
                                        v4 = v4
                                        v68 := GenS2{F0: 995}
                                        if true {
                                            var v69 uint64 = 42
                                            _ = v69
                                        }
                                        _ = v68
                                    }
                                    var v70 bool = (!false)
                                    fmt.Println(fi27, fi32, fi32)
                                    _ = v70
                                } else {
                                    switch (fi32 & 2147483647) {
                                    case 9223372036854775807:
                                        v4 = v4
                                    case 2147483647:
                                        fmt.Println(v2, v2)
                                    case 42:
                                        switch v4 {
                                        case "\n\t":
                                            if ((!false) && true) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v4 = (v4 + (v4 + ((v4 + v4) + v4)))
                                                }
                                                _ = len(v1)
                                                var v71 int16 = -32768
                                                _ = v71
                                            } else {
                                                v72 := GenS8{F0: -448, F1: 2}
                                                v73 := GenS7{F0: 2, F1: false}
                                                var v74 uint64 = 4294967295
                                                v2 -= 3.14159
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                        default:
                                            if false {
                                                v75 := GenS2{F0: 18446744073709551615}
                                                var v76 uint32 = 2147483647
                                                switch (^(fi27)) {
                                                case 32767:
                                                    v76 = 2147483647
                                                default:
                                                    var v77 bool = true
                                                    _ = v77
                                                }
                                                v2 = v2
                                                _ = v75
                                                _ = v76
                                            }
                                        }
                                    default:
                                        var v78 int = 127
                                        _ = v78
                                    }
                                    if true {
                                        if (true && false) {
                                            fmt.Println(fi27, fi32, fi32)
                                            var v79 int16 = -32768
                                            _ = v79
                                        } else {
                                            v4 = "hello"
                                            var v80 uint8 = 1
                                            _ = v80
                                        }
                                    } else {
                                        v81 := GenS6{F0: 167}
                                        _ = v81
                                    }
                                }
                                v82 := GenS2{F0: 1}
                                _ = fi32
                                _ = v82
                            }
                        }
                        v1[2] = (((!(0 > -1)) && false) || (!(65535 <= 1)))
                        switch "foo bar" {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 882
                            }
                        case "a":
                            v83 := GenS6{F0: int64(fi27)}
                            _ = v83
                        case "":
                            if len(v3) > 0 {
                                v3[0] = 2147483648
                            }
                        }
                    }
                    fmt.Println(fi27)
                }
                _ = fi27
            }
        }
    case "\n\t":
        fmt.Println(v2)
    case "":
        v3 = append(v3, 4294967295)
    default:
        _ = len(v3)
    }
    fmt.Println(v4, v4)
    v4 = v4
    if true {
        v4 = ((v4 + "hello") + v4)
        for fi84 := 0; fi84 < 18; fi84++ {
            var v85 uint8 = 2
            _ = fi84
            _ = v85
        }
    } else {
        for fi86 := 0; fi86 < 12; fi86++ {
            var v87 bool = true
            fmt.Println(v2)
            _ = fi86
            _ = v87
        }
        v4 = v4
    }
    v88 := v1[2]
    var v89 int64 = 890
    for fi90 := 0; fi90 < 11; fi90++ {
        switch ("a" + "a") {
        case "a":
            var v91 uint8 = 2
            _ = v91
        case "":
            delete(v1, 2147483647)
        }
        if (v88 || v88) {
            for fi92 := 0; fi92 < 15; fi92++ {
                if true {
                    var v93 uint8 = 0
                    _ = v93
                }
                switch fi92 {
                case 128:
                    v2 = v2
                }
                _ = fi92
            }
        } else {
            if v88 {
                if (!v88) {
                    if (!true) {
                        v94 := GenS1{F0: 963}
                        _ = v94
                    }
                    if v88 {
                        switch -1 {
                        case -2147483648:
                            v88 = v88
                        case 1:
                            v89 -= v89
                        case -9223372036854775808:
                            var v95 uint = 0
                            _ = v95
                        }
                        var v96 uint32 = 2
                        for fi97 := 0; fi97 < 5; fi97++ {
                            switch v4 {
                            case "foo bar":
                                var v98 uint8 = 0
                                _ = v98
                            case "hello":
                                v96 -= (v96 + 0)
                            default:
                                fmt.Println(fi90, fi97)
                            }
                            var v99 uint16 = genFn0()
                            v100 := GenS9{F0: 42, F1: 2, F2: 17}
                            _ = fi97
                            _ = v100
                            _ = v99
                        }
                        _ = v96
                    }
                    fmt.Println(fi90, v2, v2)
                }
            }
        }
        _ = fi90
    }
    var v101 uint8 = 163
    var v102 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v101
    _ = v102
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v88
    _ = v89
}
