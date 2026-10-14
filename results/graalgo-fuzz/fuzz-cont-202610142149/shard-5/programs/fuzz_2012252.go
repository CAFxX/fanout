package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint64; F2 int8 }
type GenS7 struct { F0 int32; F1 bool; F2 uint16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 int32, p1 int32) bool {
    switch 32768 {
    case 2147483647:
        switch "hello" {
        case "\n\t":
            var v5 int64 = 789
            _ = v5
        default:
            var v6 uint32 = 2147483648
            _ = v6
        }
    }
    var v7 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v7
    return false
}
func genFn1(p0 uint16, p1 uint16, p2 bool) uint16 {
    var v8 uint = 18446744073709551615
    var v9 int32 = 767
    p1 -= p1
    p1 += (p1 << ((v8) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (^(p1))
}

func main() {
    var v1 bool = (555 <= 1)
    var v2 float32 = (3.14159 / (0.1 * (1.5 / 1.5)))
    var v3 uint64 = 42
    v4 := []string{}
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            if (0 != 0) {
                fmt.Println(fi11, v1, v2)
                v3 = (^(v3))
                v1 = v1
                v12 := GenS1{F0: 1}
                _ = v12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS2{F0: 1000}
            var v15 float32 = 1e10
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    var v16 int32 = 2147483647
    var v17 int64 = -1
    if (!v1) {
        fmt.Println(v3)
        if ((18446744073709551615 == 2) && v1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v18 uint8 = 27
            _ = v18
        }
        if v1 {
            var v19 int32 = int32(v3)
            _ = v19
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float64 = 1.5
                if (32767 >= 255) {
                    switch "a" {
                    case "foo bar":
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 string = "hello"
                            v21 = 1.5
                            _ = fi22
                            _ = v23
                        }
                    }
                    v24 := GenS3{F0: 1, F1: 42, F2: 1}
                    _ = v24
                } else {
                    v16 += v16
                }
                var v25 int32 = (-(v16))
                _ = fi20
                _ = v21
                _ = v25
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            fmt.Println(v17)
        }
    }
    switch 128 {
    case 32767:
        if ((v1 || (!(2 == 188))) && v1) {
            for fi26 := 0; fi26 < 3; fi26++ {
                fmt.Println(v2, fi26)
                for fi27 := 0; fi27 < 9; fi27++ {
                    if true {
                        if (!(v1 || ((fi27 * (fi26 * -641)) < 1))) {
                            v28 := GenS4{F0: 42, F1: 100, F2: (v17 % ((v17) | 1))}
                            var v29 uint32 = 0
                            var v30 int32 = (v16 << ((715) & 63))
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            if len(v4) > 0 {
                                v4[0] = ("hello" + "hello")
                            }
                        }
                        v17 -= v17
                    }
                    _ = fi27
                }
                _ = fi26
            }
            var v31 float32 = v2
            v16 += (v16 + (v16 ^ 2147483647))
            _ = v31
        } else {
            switch ("\n\t" + ("hello" + "\n\t")) {
            case "hello":
                for fi32 := 0; fi32 < 3; fi32++ {
                    switch "" {
                    case "\n\t":
                        if (!false) {
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            for fi33 := 0; fi33 < 5; fi33++ {
                                if v1 {
                                    if false {
                                        switch (^(fi33)) {
                                        case 1:
                                            if v1 {
                                                switch ("hello" + "a") {
                                                case "\n\t":
                                                    v16 -= -379
                                                }
                                                if v1 {
                                                    v17 -= v17
                                                    if false {
                                                        v16 += (^(v16))
                                                        if true {
                                                            fmt.Println(v1, v16)
                                                            var v34 bool = v1
                                                            if v34 {
                                                                v35 := GenS3{F0: (^(v3)), F1: -84, F2: 2}
                                                                v17 = (v17 ^ (v17 - v17))
                                                                _ = v35
                                                            } else {
                                                                var v36 string = ""
                                                                _ = v36
                                                            }
                                                            v3 = v3
                                                            _ = v34
                                                        } else {
                                                            v16 = v16
                                                            v37 := GenS1{F0: 555}
                                                            _ = v37
                                                        }
                                                        v38 := GenS3{F0: v3, F1: 42, F2: 1}
                                                        v2 -= 0.0
                                                        _ = v38
                                                    }
                                                    var v39 uint32 = 65535
                                                    fmt.Println(v17)
                                                    _ = v39
                                                }
                                                fmt.Println(v3)
                                            }
                                        default:
                                            v40 := GenS5{F0: 42}
                                            _ = v40
                                        }
                                        if (!(v2 >= v2)) {
                                            switch "foo bar" {
                                            case "a":
                                                v16 = v16
                                            }
                                        } else {
                                            if v1 {
                                                v16 += 42
                                                var v41 uint8 = 124
                                                _ = v41
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    v2 = v2
                                                default:
                                                    if true {
                                                        var v42 int16 = 0
                                                        fmt.Println(fi32)
                                                        var v43 string = "hello"
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                                if v1 {
                                                    if (!(100 <= -56)) {
                                                        if len(v4) > 0 {
                                                            v4[0] = "a"
                                                        }
                                                        v1 = true
                                                        _ = len(v4)
                                                    } else {
                                                        var v44 int8 = 100
                                                        v1 = ((((179 < 42) || v1) && v1) || v1)
                                                        _ = v44
                                                    }
                                                    _ = len(v4)
                                                    var v45 uint32 = 4294967295
                                                    _ = v45
                                                } else {
                                                    if (int16(v16) < 1000) {
                                                        if (v17 < (v17 - (v17 - v17))) {
                                                            switch ("a" + ((("a" + (("a" + "\n\t") + "foo bar")) + "foo bar") + "\n\t")) {
                                                            case "hello":
                                                                v46 := GenS6{F0: 536, F1: (v3 >> ((640) & 63)), F2: -72}
                                                                _ = v46
                                                            case "":
                                                                if v1 {
                                                                    if (("a" + "a") < ("a" + "hello")) {
                                                                        v2 -= (((float32(1e10) + v2) * 1.5) * float32(0.0))
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                } else {
                                                                    switch (fi32 % ((0) | 1)) {
                                                                    case -835:
                                                                        switch 779 {
                                                                        case 713:
                                                                            switch (fi33 - fi32) {
                                                                            case 32768:
                                                                                v1 = (1 == 54)
                                                                            case 187:
                                                                                var v47 uint32 = 1
                                                                                _ = v47
                                                                            }
                                                                        default:
                                                                            v48 := GenS5{F0: 0}
                                                                            _ = v48
                                                                        }
                                                                    case 2147483647:
                                                                        fmt.Println(v16, v16, v3)
                                                                    case 4294967295:
                                                                        fmt.Println(fi33)
                                                                    default:
                                                                        var v49 float64 = 0.0
                                                                        _ = v49
                                                                    }
                                                                    if (42 != 446) {
                                                                        v2 += v2
                                                                        fmt.Println(fi32, v1, fi33)
                                                                        if true {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ""
                                                                            }
                                                                            var v50 uint64 = v3
                                                                            _ = v50
                                                                        }
                                                                    }
                                                                    var v51 float32 = v2
                                                                    var v52 int8 = 100
                                                                    v16 += v16
                                                                    var v53 uint = 255
                                                                    v51 -= (-2.25 / v2)
                                                                    _ = v51
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                            case "\n\t":
                                                                v54 := GenS7{F0: (v16 >> ((2) & 63)), F1: false, F2: 1}
                                                                _ = v54
                                                            default:
                                                                var v55 float64 = 1.5
                                                                _ = v55
                                                            }
                                                            var v56 uint16 = 325
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v57 string = ""
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        switch fi33 {
                                                        case 1:
                                                            v58 := GenS1{F0: 4294967295}
                                                            _ = v58
                                                        }
                                                    }
                                                }
                                            }
                                            v16 += v16
                                            var v59 string = "\n\t"
                                            _ = v59
                                        }
                                        fmt.Println(fi32, fi33, v2)
                                    }
                                }
                                v16 += int32(fi32)
                                _ = fi33
                            }
                            var v60 int8 = 127
                            _ = v60
                        }
                    }
                    v16 = v16
                    _ = fi32
                }
            default:
                v3 = v3
            }
            v16 -= v16
        }
    case -355:
        switch (((("" + ("foo bar" + "foo bar")) + ("foo bar" + "foo bar")) + (("hello" + ("a" + ("a" + ((("\n\t" + "\n\t") + ("\n\t" + "")) + "a")))) + "a")) + ("a" + "")) {
        case "hello":
            if (18 < 0) {
                v17 = v17
            }
        default:
            var v61 int64 = (v17 % ((1) | 1))
            _ = v61
        }
    case 250:
        v62 := GenS4{F0: 0, F1: 1, F2: (v17 ^ v17)}
        _ = v62
    }
    switch -9223372036854775808 {
    case 1:
        if true {
            v1 = v1
        } else {
            fmt.Println(v3, v2, v2)
            var v63 bool = (v1 || (false && (1e10 == (3.14159 / 1.0))))
            v64 := GenS8{F0: 42, F1: 42}
            _ = v63
            _ = v64
        }
    case 343:
        if (false && v1) {
            if v1 {
                if v1 {
                    v17 -= 771
                    v3 -= v3
                }
                for fi65 := 0; fi65 < 4; fi65++ {
                    v1 = (!(2147483647 == 2147483647))
                    var v66 int16 = 42
                    _ = fi65
                    _ = v66
                }
            } else {
                switch 32767 {
                case -9223372036854775808:
                    for fi67 := 0; fi67 < 19; fi67++ {
                        _ = len(v4)
                        var v68 int32 = 42
                        _ = fi67
                        _ = v68
                    }
                case -801:
                    var v69 uint16 = 0
                    _ = v69
                case 1:
                    var v70 bool = (!(!v1))
                    _ = v70
                }
            }
            switch 42 {
            case -9223372036854775808:
                for fi71 := 0; fi71 < 3; fi71++ {
                    v72 := GenS3{F0: (v3 * (v3 * 4294967295)), F1: -76, F2: 42}
                    _ = fi71
                    _ = v72
                }
            }
        } else {
            v17 = -326
            v16 -= 0
        }
    case -2147483648:
        for fi73 := 0; fi73 < 7; fi73++ {
            v17 += v17
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v74 int8 = 45
            _ = fi73
            _ = v74
        }
    }
    if (((((v2 - v2) * 0.1) > 1.5) || v1) && true) {
        _ = len(v4)
        var v75 uint64 = v3
        fmt.Println(v1, v3, v1)
        _ = v75
    }
    if (1 < (v3 << ((18446744073709551615) & 63))) {
        v1 = false
    }
    v76 := GenS9{F0: 1}
    switch "foo bar" {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = "hello"
        }
    case "":
        v77 := GenS6{F0: 42, F1: v3, F2: 100}
        _ = v77
    }
    if v1 {
        switch 42 {
        case -571:
            switch 127 {
            case 128:
                var v78 uint32 = 192
                _ = v78
            case 4294967295:
                switch 206 {
                case -1:
                    for fi79 := 0; fi79 < 18; fi79++ {
                        v4 = append(v4, "\n\t")
                        _ = fi79
                    }
                }
            case 625:
                v80 := GenS10{F0: "foo bar", F1: int32(v17)}
                _ = v80
            default:
                var v81 uint64 = (v3 ^ v3)
                _ = v81
            }
        case 9223372036854775807:
            if len(v4) > 0 {
                v4[0] = ""
            }
        case -1:
            var v82 uint = 65535
            _ = v82
        default:
            v17 -= v17
        }
        v83 := GenS11{F0: v17, F1: v16}
        _ = v83
    } else {
        switch "\n\t" {
        case "hello":
            v84 := GenS10{F0: ((("a" + "a") + "\n\t") + "foo bar"), F1: v16}
            _ = v84
        case "\n\t":
            v3 -= 806
        default:
            var v85 int64 = 42
            _ = v85
        }
    }
    if false {
        if true {
            v17 = (v17 << ((0) & 63))
            for fi86 := 0; fi86 < 19; fi86++ {
                v87 := GenS12{F0: 1, F1: 2147483647}
                v17 = (v17 << ((42) & 63))
                fmt.Println(fi86, v2, v1)
                _ = fi86
                _ = v87
            }
            for fi88 := 0; fi88 < 16; fi88++ {
                if ((v1 && v1) && v1) {
                    var v89 float32 = v2
                    fmt.Println(v3)
                    _ = v89
                }
                switch 128 {
                case 9223372036854775807:
                    for fi90 := 0; fi90 < 17; fi90++ {
                        if len(v4) > 0 {
                            v4[0] = ("\n\t" + "a")
                        }
                        var v91 float32 = (float32((-2.25 + 1.5)) * (float32((float64(float32((3.14159 * 3.14159))) / (((3.14159 * float64(v2)) / 0.1) + (0.1 + float64((3.14159 - (v2 / 3.14159))))))) - v2))
                        _ = fi90
                        _ = v91
                    }
                }
                v2 -= v2
                _ = fi88
            }
        } else {
            fmt.Println(v16, v3)
            var v92 int32 = v16
            switch "a" {
            case "a":
                if v1 {
                    var v93 float32 = 1e10
                    if v1 {
                        if v1 {
                            var v94 int16 = -1
                            var v95 string = ""
                            v93 -= v93
                            _ = v94
                            _ = v95
                        } else {
                            v16 = v16
                            _ = len(v4)
                            var v96 uint8 = 128
                            _ = v96
                        }
                        if v1 {
                            for fi97 := 0; fi97 < 13; fi97++ {
                                v17 += v17
                                for fi98 := 0; fi98 < 2; fi98++ {
                                    v92 += v16
                                    if v1 {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                    } else {
                                        fmt.Println(fi97)
                                        var v99 string = "\n\t"
                                        var v100 uint8 = 235
                                        switch (fi98 | (fi98 << ((2) & 63))) {
                                        case -1:
                                            v101 := GenS9{F0: -1}
                                            _ = v101
                                        case 288:
                                            if len(v4) > 0 {
                                                v4[0] = ""
                                            }
                                            var v102 uint64 = v3
                                            if len(v4) > 0 {
                                                v4[0] = v99
                                            }
                                            _ = v102
                                        case 4294967295:
                                            var v103 float64 = (float64((((v2 * -2.25) * float32((-2.25 - 3.14159))) + v93)) / (float64(1.5) - -2.25))
                                            _ = v103
                                        }
                                        var v104 uint64 = 454
                                        _ = v100
                                        _ = v104
                                        _ = v99
                                    }
                                    _ = fi98
                                }
                                v16 -= 749
                                _ = fi97
                            }
                            fmt.Println(v93, v1, v17)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v105 int = 2147483647
                        _ = v105
                    }
                    _ = v93
                } else {
                    if (!((("\n\t" + ("\n\t" + "a")) >= "\n\t") || false)) {
                        var v106 uint16 = 32767
                        _ = v106
                    }
                    fmt.Println(v1)
                    _ = len(v4)
                }
            default:
                v1 = v1
            }
            _ = v92
        }
        if (255 >= 521) {
            v16 = v16
        }
        var v107 bool = (!((!v1) && (v1 || v1)))
        _ = v107
    }
    fmt.Println(v2, v1)
    v17 += v17
    if v1 {
        v2 += v2
    } else {
        if v1 {
            fmt.Println(v2, v17)
            var v108 uint16 = 255
            _ = v108
        } else {
            for fi109 := 0; fi109 < 17; fi109++ {
                v16 = v16
                _ = fi109
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if (v1 || (!v1)) {
            v110 := GenS13{F0: (v17 << ((439) & 63)), F1: 488, F2: 255}
            _ = v110
        }
        for fi111 := 0; fi111 < 13; fi111++ {
            if (!false) {
                fmt.Println(v1)
            } else {
                v3 = v3
            }
            _ = fi111
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v76
}
