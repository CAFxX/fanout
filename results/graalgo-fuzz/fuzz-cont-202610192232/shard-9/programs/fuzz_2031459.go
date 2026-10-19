package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint; F2 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 int8; F2 string }
func genFn0() uint64 {
    var v6 uint16 = 791
    v7 := GenS1{F0: 0}
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint32, p1 string, p2 int8) uint {
    p0 += p0
    var v8 int8 = (p2 % ((p2) | 1))
    p2 += (v8 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 160
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 42
    var v3 uint16 = 32767
    var v4 int64 = 9223372036854775807
    var v5 uint8 = 1
    var v9 uint = 65535
    if (1 != 233) {
        for fi10 := 0; fi10 < 6; fi10++ {
            v4 = 1
            _ = fi10
        }
    } else {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            if true {
                v5 = (v5 >> (((v9 | 4294967295)) & 63))
                if true {
                    v3 = 65535
                    var v11 uint = v9
                    _ = v11
                }
                v3 = v3
            }
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                v2 += (v2 + v2)
                _ = fi12
            }
        }
    }
    switch "hello" {
    case "":
        v3 -= 1
    case "a":
        var v13 int8 = 0
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: -908}
            if ((1000 < 42) || (!((0.0 / float64(float32(1.5))) < 1.5))) {
                fmt.Println(v5, v3)
                fmt.Println(v2)
            }
            _ = fi14
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 42}
        fmt.Println(v4, v2)
        switch "\n\t" {
        case "":
            v17 := GenS1{F0: -1}
            _ = v17
        case "foo bar":
            var v18 int16 = -1
            _ = v18
        default:
            v3 += (v3 ^ v3)
        }
        _ = v16
    }
    v4 += v4
    switch 578 {
    case 478:
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                if (v5 == v5) {
                    switch "\n\t" {
                    case "hello":
                        v1[1] = "foo bar"
                    case "\n\t":
                        v21 := GenS1{F0: 32767}
                        _ = v21
                    case "a":
                        v3 = (v3 << ((v9) & 63))
                    }
                    v2 -= 965
                    if (!(v3 < 2)) {
                        if true {
                            v1[4294967295] = ("hello" + "\n\t")
                            var v22 string = ""
                            v2 = (v2 >> (((v9 & genFn1(93, v22, 0))) & 63))
                            _ = v22
                        }
                    }
                    var v23 int64 = v4
                    switch "hello" {
                    case "a":
                        v24 := GenS1{F0: -186}
                        _ = v24
                    default:
                        switch -391 {
                        case -9223372036854775808:
                            if ((false || true) && false) {
                                switch fi19 {
                                case 127:
                                    v1[(fi19 % ((fi19) | 1))] = "a"
                                }
                            }
                        case 441:
                            fmt.Println(v9)
                        case 439:
                            v2 = 2147483648
                        default:
                            var v25 float64 = 1e10
                            _ = v25
                        }
                    }
                    v26 := []int{fi20, (^(fi20)), (fi19 ^ int(v9))}
                    _ = v23
                    _ = v26
                } else {
                    fmt.Println(v5, v9)
                    if true {
                        switch ("\n\t" + "") {
                        case "a":
                            fmt.Println(v3)
                        case "\n\t":
                            var v27 int32 = -2147483648
                            _ = v27
                        case "":
                            var v28 uint8 = (v5 >> ((v9) & 63))
                            _ = v28
                        }
                        if ((100 <= 127) && false) {
                            if (false && ((("" + ("hello" + "")) + ("foo bar" + (("" + "foo bar") + ("" + ("foo bar" + ("" + "")))))) == ("foo bar" + ("hello" + "a")))) {
                                fmt.Println(v5, v9, v9)
                                var v29 float64 = 3.14159
                                v29 += v29
                                _ = v29
                            }
                            delete(v1, fi20)
                            switch (("" + "a") + "") {
                            case "foo bar":
                                if false {
                                    v30 := GenS1{F0: 1}
                                    fmt.Println(v4, v9, v9)
                                    v1[(fi20 << ((v9) & 63))] = ("foo bar" + ((("" + ("\n\t" + "a")) + "foo bar") + ("foo bar" + ("hello" + "\n\t"))))
                                    _ = v30
                                } else {
                                    v9 -= (^(v9))
                                    v3 -= (v3 + v3)
                                    v3 = uint16(v2)
                                }
                            case "\n\t":
                                var v31 float32 = 1e10
                                _ = v31
                            }
                            var v32 string = "foo bar"
                            fmt.Println(v3, fi19)
                            _ = v32
                        }
                        v3 += v3
                        delete(v1, (fi20 + fi19))
                        if false {
                            v9 = v9
                            var v33 uint16 = (v3 - v3)
                            _ = v33
                        }
                        v4 = v4
                    }
                    delete(v1, (fi19 ^ (fi19 ^ fi19)))
                }
                fmt.Println(v5)
                fmt.Println(v5, v3, v5)
                _ = fi20
            }
            v34 := GenS1{F0: int16(fi19)}
            _ = fi19
            _ = v34
        }
    case -108:
        v1[-771] = ""
    }
    v5 -= (v5 % ((180) | 1))
    v3 += 42
    v9 = uint(v3)
    if false {
        if (0 < 121) {
            var v35 int16 = -32768
            var v36 uint32 = 2147483647
            _ = v35
            _ = v36
        } else {
            switch -960 {
            case 1:
                v9 += v9
            case -331:
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch "a" {
                    case "hello":
                        if true {
                            var v38 float32 = -2.25
                            fmt.Println(v5)
                            _ = v38
                        }
                    case "\n\t":
                        var v39 float32 = float32((0.1 - 1.5))
                        _ = v39
                    case "foo bar":
                        for fi40 := 0; fi40 < 14; fi40++ {
                            if false {
                                v2 = v2
                            } else {
                                var v41 int8 = -1
                                if true {
                                    if false {
                                        if (!((v9 % ((genFn1(v2, ("\n\t" + ""), (v41 & (v41 << ((genFn1(223, ("\n\t" + "hello"), (v41 | -14))) & 63))))) | 1)) > 228)) {
                                            if true {
                                                switch "\n\t" {
                                                case "a":
                                                    v42 := GenS1{F0: 1}
                                                    switch "" {
                                                    case "a":
                                                        var v43 float64 = 3.14159
                                                        _ = v43
                                                    default:
                                                        var v44 int16 = -32768
                                                        _ = v44
                                                    }
                                                    _ = v42
                                                case "hello":
                                                    _ = len(v1)
                                                case "":
                                                    v3 = uint16(v2)
                                                }
                                            }
                                        } else {
                                            v45 := GenS1{F0: 660}
                                            var v46 uint64 = 18446744073709551615
                                            v47 := GenS1{F0: 1000}
                                            if true {
                                                fmt.Println(v3, v41, v41)
                                                var v48 uint32 = 0
                                                switch 4294967295 {
                                                case 496:
                                                    fmt.Println(v3)
                                                default:
                                                    var v49 int64 = 1
                                                    _ = v49
                                                }
                                                switch 137 {
                                                case -775:
                                                    var v50 uint16 = 290
                                                    _ = v50
                                                case 32767:
                                                    v1[(-(fi37))] = (("" + ("" + "")) + ("a" + ("" + "foo bar")))
                                                case -927:
                                                    switch "a" {
                                                    case "hello":
                                                        v46 -= (v46 - (^(v46)))
                                                    case "a":
                                                        if false {
                                                            v41 += v41
                                                            fmt.Println(fi40, v41)
                                                            var v51 int64 = (^(v4))
                                                            _ = v51
                                                        } else {
                                                            switch (fi40 - fi40) {
                                                            case 127:
                                                                if false {
                                                                    var v52 int64 = (v4 * v4)
                                                                    var v53 float64 = 0.0
                                                                    v54 := GenS2{F0: -32768, F1: 32767, F2: v2}
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                            }
                                                            v9 += uint(v48)
                                                        }
                                                        switch (("foo bar" + "\n\t") + "hello") {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch "" {
                                                                case "foo bar":
                                                                    v3 = 65535
                                                                case "\n\t":
                                                                    var v55 uint64 = uint64(v5)
                                                                    _ = v55
                                                                case "a":
                                                                    switch int(v3) {
                                                                    case 4294967295:
                                                                        switch (^(fi40)) {
                                                                        case 32768:
                                                                            var v56 uint8 = 2
                                                                            _ = v56
                                                                        case 128:
                                                                            v2 -= (v48 << ((v9) & 63))
                                                                        case 112:
                                                                            var v57 float64 = (0.0 + (0.0 / 1.0))
                                                                            _ = v57
                                                                        }
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v5)
                                                                    case 9223372036854775807:
                                                                        v9 += v9
                                                                        fmt.Println(v48, v5, v2)
                                                                        v58 := GenS3{F0: 1}
                                                                        _ = v58
                                                                    }
                                                                }
                                                            case "a":
                                                                switch 617 {
                                                                case 128:
                                                                    var v59 string = ("" + "")
                                                                    _ = v59
                                                                case -979:
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        _ = len(v1)
                                                                    case "\n\t":
                                                                        v3 = (^(v3))
                                                                    case "hello":
                                                                        if (true && false) {
                                                                            switch fi40 {
                                                                            case 4294967295:
                                                                                v4 -= v4
                                                                            case 985:
                                                                                if (4294967295 < genFn0()) {
                                                                                    var v60 uint32 = (v48 - 4294967295)
                                                                                    v61 := GenS4{F0: (v4 & v4), F1: 42}
                                                                                    fmt.Println(v46, v60, v9)
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                var v62 uint64 = 42
                                                                                _ = v62
                                                                            case "hello":
                                                                                v63 := []int{(fi37 + (^(fi37))), fi40, (fi40 / ((1) | 1)), (-(fi37))}
                                                                                _ = v63
                                                                            default:
                                                                                v46 = v46
                                                                            }
                                                                            fmt.Println(v9)
                                                                        } else {
                                                                            switch 607 {
                                                                            case -2147483648:
                                                                                var v64 int32 = -2147483648
                                                                                v65 := GenS4{F0: (v4 << ((65535) & 63)), F1: 1000}
                                                                                var v66 uint = uint(v46)
                                                                                _ = v64
                                                                                _ = v65
                                                                                _ = v66
                                                                            default:
                                                                                fmt.Println(v4, fi40)
                                                                            }
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                v67 := GenS1{F0: 1}
                                                                                _ = v67
                                                                            default:
                                                                                v68 := []int{(fi40 << ((946) & 63)), int(v3), -865}
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        v9 = (v9 << ((v9) & 63))
                                                                    }
                                                                case 32768:
                                                                    var v69 uint32 = 0
                                                                    _ = v69
                                                                default:
                                                                    if true {
                                                                        if true {
                                                                            fmt.Println(fi40, v41)
                                                                        } else {
                                                                            fmt.Println(v2, v9, fi40)
                                                                        }
                                                                    } else {
                                                                        v4 = 1
                                                                        v9 = (v9 << ((730) & 63))
                                                                        v70 := GenS5{F0: (v5 & 0)}
                                                                        v71 := v1[9223372036854775807]
                                                                        v72 := GenS5{F0: v5}
                                                                        _ = v70
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                }
                                                            case "hello":
                                                                v46 = (v46 + 2)
                                                            default:
                                                                var v73 int8 = (-(v41))
                                                                _ = v73
                                                            }
                                                        case "\n\t":
                                                            v1[fi40] = ""
                                                        case "hello":
                                                            fmt.Println(v9, fi40, v2)
                                                        default:
                                                            var v74 uint64 = v46
                                                            _ = v74
                                                        }
                                                    case "foo bar":
                                                        if (42 < 42) {
                                                            var v75 int16 = int16(v4)
                                                            v76 := GenS6{F0: (v9 << ((18446744073709551615) & 63)), F1: (v9 - 18446744073709551615), F2: (^(v2))}
                                                            v3 = (^(v3))
                                                            if (true || false) {
                                                                var v77 float32 = 3.14159
                                                                switch ("foo bar" + "hello") {
                                                                case "hello":
                                                                    v78 := GenS4{F0: v4, F1: -1}
                                                                    _ = v78
                                                                case "\n\t":
                                                                    v77 += v77
                                                                }
                                                                _ = v77
                                                            } else {
                                                                v9 += (v9 * genFn1((v48 >> ((v9) & 63)), (("a" + "foo bar") + "a"), (-(v41))))
                                                            }
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            v4 -= v4
                                                            var v79 uint64 = genFn0()
                                                            var v80 uint64 = 42
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                    default:
                                                        var v81 uint32 = v48
                                                        _ = v81
                                                    }
                                                }
                                                _ = v48
                                            } else {
                                                if (v41 > (v41 + (v41 + v41))) {
                                                    if (false || false) {
                                                        var v82 float32 = float32((-2.25 * float64(-2.25)))
                                                        _ = v82
                                                    }
                                                    v46 = 152
                                                    if (-1 < v41) {
                                                        var v83 string = "hello"
                                                        switch ("" + "foo bar") {
                                                        case "a":
                                                            switch (fi40 * fi37) {
                                                            case 494:
                                                                var v84 bool = (!false)
                                                                _ = v84
                                                            }
                                                        case "hello":
                                                            var v85 string = v83
                                                            _ = v85
                                                        }
                                                        var v86 string = ((v83 + ("" + "a")) + (v83 + "hello"))
                                                        var v87 uint32 = (v2 << ((v9) & 63))
                                                        _ = v83
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                }
                                            }
                                            v4 -= (v4 + int64(fi37))
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        if (!((!(true && (32767 < -32768))) || (-2.25 <= (float32(1.5) * 3.14159)))) {
                                            if false {
                                                var v88 int32 = 338
                                                if (!true) {
                                                    switch "" {
                                                    case "":
                                                        var v89 uint32 = uint32(v3)
                                                        _ = v89
                                                    case "foo bar":
                                                        v90 := GenS7{F0: v4}
                                                        _ = v90
                                                    }
                                                    v2 = v2
                                                } else {
                                                    if true {
                                                        var v91 int16 = -371
                                                        if false {
                                                            if false {
                                                                v2 -= v2
                                                            }
                                                        } else {
                                                            v88 = (v88 - (v88 % (((v88 | (^(v88)))) | 1)))
                                                            v92 := []int{int(v88), fi40}
                                                            var v93 int64 = v4
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                        _ = v91
                                                    }
                                                    var v94 int = 32768
                                                    _ = v94
                                                }
                                                _ = v88
                                            }
                                        }
                                    } else {
                                        var v95 int16 = 1000
                                        _ = v95
                                    }
                                    var v96 uint64 = 0
                                    v3 -= v3
                                    _ = v96
                                } else {
                                    var v97 uint8 = 2
                                    _ = v97
                                }
                                var v98 int32 = -876
                                _ = v41
                                _ = v98
                            }
                            var v99 uint64 = 0
                            v2 = (v2 * v2)
                            _ = fi40
                            _ = v99
                        }
                    default:
                        v4 = (v4 << ((v9) & 63))
                    }
                    v100 := []int{(fi37 * (fi37 * fi37))}
                    _ = fi37
                    _ = v100
                }
            case 0:
                var v101 uint8 = v5
                _ = v101
            default:
                var v102 int64 = (v4 / ((v4) | 1))
                _ = v102
            }
        }
    } else {
        if (true || false) {
            var v103 int32 = 1
            if ((3.14159 <= (-2.25 - float64(1.5))) || false) {
                var v104 float64 = (3.14159 / 1e10)
                v4 -= 246
                _ = v104
            } else {
                var v105 bool = (!false)
                var v106 int32 = (v103 & -537)
                v107 := GenS5{F0: 101}
                _ = v105
                _ = v106
                _ = v107
            }
            v2 = v2
            _ = v103
        }
    }
    v9 = v9
    if false {
        fmt.Println(v3, v4, v9)
    }
    if false {
        var v108 string = ("a" + "hello")
        _ = v108
    }
    if true {
        switch -2147483648 {
        case 1:
            var v109 int32 = 2147483647
            _ = v109
        case 127:
            switch "" {
            case "":
                fmt.Println(v9, v9)
            default:
                for fi110 := 0; fi110 < 2; fi110++ {
                    var v111 int16 = 332
                    switch (fi110 / ((int(v2)) | 1)) {
                    case 0:
                        var v112 int64 = v4
                        _ = v112
                    case 9223372036854775807:
                        var v113 int8 = 1
                        _ = v113
                    default:
                        for fi114 := 0; fi114 < 13; fi114++ {
                            fmt.Println(v2)
                            _ = fi114
                        }
                    }
                    _ = fi110
                    _ = v111
                }
            }
        }
        for fi115 := 0; fi115 < 19; fi115++ {
            v116 := GenS8{F0: (("foo bar" + ((("foo bar" + "foo bar") + "a") + "")) + ("hello" + ("hello" + "foo bar"))), F1: 0, F2: ("a" + "")}
            _ = fi115
            _ = v116
        }
    } else {
        v2 += v2
        if false {
            var v117 int16 = 42
            v118 := GenS7{F0: v4}
            if (0.1 != (-2.25 + (3.14159 / 1.5))) {
                v119 := GenS2{F0: v117, F1: v117, F2: v2}
                v1[-9223372036854775808] = (("" + ("" + "\n\t")) + "\n\t")
                v2 = 162
                _ = v119
            }
            _ = v117
            _ = v118
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
