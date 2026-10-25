package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int; F2 bool }
type GenS6 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0() uint32 {
    var v4 int8 = 55
    v5 := GenS1{F0: -1, F1: 9223372036854775807}
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint8, p1 bool, p2 int32) string {
    v6 := []int{127, 623, 1, 42}
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!p1) {
            switch 127 {
            case -9223372036854775808:
                switch (fi7 / (((fi7 << ((1) & 63))) | 1)) {
                case 9223372036854775807:
                    switch "\n\t" {
                    case "\n\t":
                        for fi8 := 0; fi8 < 10; fi8++ {
                            p1 = p1
                            p2 = p2
                            _ = fi8
                        }
                    case "":
                        var v9 float32 = float32(0.0)
                        _ = v9
                    default:
                        v10 := GenS1{F0: fi7, F1: 1}
                        _ = v10
                    }
                case -1:
                    if len(v6) > 0 {
                        v6[0] = -200
                    }
                }
            case 9223372036854775807:
                v6 = append(v6, fi7)
            }
            if p1 {
                p1 = p1
                switch "\n\t" {
                case "a":
                    v11 := GenS1{F0: (-(fi7)), F1: -60}
                    _ = v11
                case "foo bar":
                    var v12 string = "hello"
                    _ = v12
                case "\n\t":
                    switch ("foo bar" + ("" + ("hello" + ("foo bar" + "a")))) {
                    case "a":
                        if p1 {
                            fmt.Println(p2, p0, fi7)
                        }
                    case "\n\t":
                        fmt.Println(p2, p2, p0)
                    case "":
                        for fi13 := 0; fi13 < 11; fi13++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v14 := GenS2{F0: 128}
                            if (505 >= uint16(p0)) {
                                p2 = p2
                                p2 = 1
                                v15 := GenS1{F0: 0, F1: 42}
                                _ = v15
                            } else {
                                var v16 bool = p1
                                var v17 uint8 = (p0 ^ (p0 ^ p0))
                                _ = v16
                                _ = v17
                            }
                            _ = fi13
                            _ = v14
                        }
                    default:
                        if (false || p1) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            p1 = p1
                        }
                    }
                default:
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS2{F0: (^(p0))}
                        _ = fi18
                        _ = v19
                    }
                }
            } else {
                var v20 uint32 = 2
                var v21 int64 = -694
                if (p1 || (-128 < 100)) {
                    var v22 float64 = 0.0
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS3{F0: (v21 << ((42) & 63)), F1: (-(v21))}
                        v21 -= v21
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi23
                        _ = v24
                    }
                    if p1 {
                        var v25 bool = true
                        fmt.Println(v25, p1, v25)
                        if p1 {
                            fmt.Println(v21, v21, v20)
                            fmt.Println(v22)
                        } else {
                            v22 -= float64(3.14159)
                        }
                        _ = v25
                    }
                    _ = v22
                } else {
                    var v26 bool = ((255 < 111) && (1000 != -32768))
                    fmt.Println(v26, p1)
                    fmt.Println(p0, v21)
                    _ = v26
                }
                _ = v20
                _ = v21
            }
        } else {
            p2 += (p2 - (p2 - (p2 >> ((461) & 63))))
            fmt.Println(p1, p2)
        }
        _ = fi7
    }
    p1 = ((p1 || (!(-1 >= -988))) && (p1 || (578 <= 0)))
    p2 -= (-(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}

func main() {
    var v1 bool = (0 <= 42)
    var v2 float64 = (1.5 * 3.14159)
    var v3 int = 32767
    v27 := GenS4{F0: 1}
    for fi28 := 0; fi28 < 16; fi28++ {
        switch (fi28 >> ((65535) & 63)) {
        case 32767:
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 uint8 = 255
                v1 = true
                _ = fi29
                _ = v30
            }
        case -829:
            v31 := GenS4{F0: 18446744073709551615}
            _ = v31
        case 0:
            for fi32 := 0; fi32 < 15; fi32++ {
                v2 += (((v2 * v2) / float64((-2.25 / 1.0))) * (float64(0.0) / (v2 + v2)))
                _ = fi32
            }
        }
        switch ((("a" + "a") + ((("hello" + "") + "hello") + "\n\t")) + "foo bar") {
        case "foo bar":
            if (!false) {
                if (65535 <= 65535) {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := []int{-9223372036854775808, (fi28 % ((337) | 1))}
                        v35 := GenS4{F0: 891}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    v36 := []int{v3, (v3 - v3), v3}
                    switch (v3 - (fi28 / (((^(fi28))) | 1))) {
                    case 32767:
                        var v37 bool = v1
                        _ = v37
                    case 2147483647:
                        var v38 uint = 18446744073709551615
                        _ = v38
                    case 740:
                        if (255 == 1) {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                if v1 {
                                    v40 := GenS5{F0: 127, F1: fi28, F2: (-96 < 100)}
                                    _ = v40
                                }
                                if v1 {
                                    v2 += v2
                                    switch (v3 * v3) {
                                    case 562:
                                        v2 -= 1e10
                                    case 32768:
                                        var v41 int = (fi39 | (fi39 * fi39))
                                        _ = v41
                                    case 32767:
                                        v2 = v2
                                    default:
                                        v3 += (fi28 | (v3 & (-(v3))))
                                    }
                                }
                                _ = fi39
                            }
                            v36 = append(v36, (v3 / ((fi28) | 1)))
                        } else {
                            v3 = (v3 | (v3 << ((2) & 63)))
                            var v42 float32 = 0.1
                            _ = v42
                        }
                    default:
                        for fi43 := 0; fi43 < 19; fi43++ {
                            var v44 bool = (42 < 0)
                            var v45 int16 = 1000
                            if v44 {
                                switch ("foo bar" + "") {
                                case "\n\t":
                                    var v46 int64 = -889
                                    v47 := GenS5{F0: 0, F1: v3, F2: true}
                                    v1 = v1
                                    _ = v46
                                    _ = v47
                                case "hello":
                                    fmt.Println(v45, v3, v3)
                                case "":
                                    if (206 > 0) {
                                        var v48 int8 = 1
                                        if len(v36) > 0 {
                                            v36[0] = -942
                                        }
                                        _ = v48
                                    } else {
                                        v49 := GenS2{F0: 0}
                                        if (!((fi28 ^ 982) > fi28)) {
                                            v2 = 0.1
                                            if v1 {
                                                var v50 int32 = 0
                                                var v51 int32 = -559
                                                _ = v50
                                                _ = v51
                                            } else {
                                                switch ("" + (("a" + "foo bar") + "foo bar")) {
                                                case "\n\t":
                                                    var v52 float32 = 1.5
                                                    _ = v52
                                                case "hello":
                                                    v53 := GenS4{F0: 65535}
                                                    _ = v53
                                                default:
                                                    var v54 int = (fi28 * 32767)
                                                    var v55 uint64 = 42
                                                    _ = v54
                                                    _ = v55
                                                }
                                                v56 := GenS5{F0: -1, F1: (fi28 ^ fi43), F2: false}
                                                var v57 string = "hello"
                                                _ = v56
                                                _ = v57
                                            }
                                            v3 = 128
                                        } else {
                                            v3 = (fi43 << ((1) & 63))
                                            fmt.Println(v44, fi43)
                                            var v58 uint16 = 859
                                            _ = v58
                                        }
                                        v44 = v44
                                        fmt.Println(v1)
                                        v59 := GenS5{F0: 0, F1: fi43, F2: (v45 <= v45)}
                                        _ = v49
                                        _ = v59
                                    }
                                default:
                                    switch (-(v3)) {
                                    case -2147483648:
                                        if (!((v3 % ((v3) | 1)) > (^(fi28)))) {
                                            switch v3 {
                                            case 42:
                                                v60 := GenS5{F0: 1, F1: fi43, F2: (!false)}
                                                _ = v60
                                            default:
                                                if len(v36) > 0 {
                                                    v36[0] = 9223372036854775807
                                                }
                                            }
                                            if (1e10 >= v2) {
                                                v45 -= 692
                                                if false {
                                                    if v1 {
                                                        var v61 int64 = -1
                                                        _ = len(v36)
                                                        _ = v61
                                                    }
                                                    v2 -= 1.5
                                                }
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v1, fi28)
                                            case "hello":
                                                if (!((-2.25 + v2) <= v2)) {
                                                    switch 51 {
                                                    case 127:
                                                        v3 = (v3 | fi28)
                                                    case -1:
                                                        fmt.Println(v2)
                                                    }
                                                    switch ("\n\t" + ("a" + "hello")) {
                                                    case "foo bar":
                                                        v36 = append(v36, v3)
                                                    default:
                                                        v2 = 1.5
                                                    }
                                                    if v1 {
                                                        v36 = append(v36, (fi43 + -2147483648))
                                                        v3 += fi28
                                                        v62 := GenS4{F0: 1}
                                                        _ = v62
                                                    } else {
                                                        if (!(fi43 == fi43)) {
                                                            var v63 uint = 4294967295
                                                            _ = v63
                                                        }
                                                        switch fi28 {
                                                        case 32768:
                                                            switch 127 {
                                                            case -2147483648:
                                                                var v64 uint32 = 572
                                                                _ = v64
                                                            }
                                                        }
                                                        var v65 int16 = (v45 >> ((270) & 63))
                                                        _ = v65
                                                    }
                                                    var v66 int64 = 1
                                                    var v67 bool = (!(!(!v1)))
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    v68 := GenS1{F0: fi43, F1: 1}
                                                    _ = v68
                                                }
                                            case "foo bar":
                                                v69 := GenS4{F0: 2}
                                                _ = v69
                                            default:
                                                _ = len(v36)
                                            }
                                            if len(v36) >= 2 {
                                                v36 = v36[1:len(v36)]
                                            }
                                            if v44 {
                                                _ = len(v36)
                                                fmt.Println(v45, v3, v45)
                                                if (!(false || ((!(!(65535 > 2))) || v1))) {
                                                    fmt.Println(v45, v3, v44)
                                                }
                                                if false {
                                                    v2 += 0.1
                                                    if true {
                                                        v70 := GenS3{F0: -1, F1: 565}
                                                        _ = v70
                                                    }
                                                    var v71 string = ""
                                                    _ = v71
                                                }
                                            }
                                            if (!(v1 && ((fi43 + (fi43 << ((2) & 63))) >= 32767))) {
                                                if (("foo bar" + ("hello" + ("hello" + (("hello" + ("foo bar" + "\n\t")) + "foo bar")))) != "\n\t") {
                                                    v72 := GenS1{F0: 2147483647, F1: 867}
                                                    v45 -= v45
                                                    var v73 int8 = 1
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v74 := GenS2{F0: 0}
                                                _ = v74
                                            }
                                            switch "hello" {
                                            case "":
                                                v1 = ((((((-2.25 / (v2 + v2)) + (((-2.25 - (3.14159 * (v2 - ((3.14159 + (float64(1e10) / v2)) + (v2 - v2))))) + 0.0) + float64((-2.25 + float32(3.14159))))) - v2) - (v2 - 0.1)) > 0.0) && false)
                                            }
                                            if (v1 || v44) {
                                                v3 = 127
                                                v3 -= -403
                                                v45 -= v45
                                                v45 -= (v45 - v45)
                                                if (v44 && false) {
                                                    v44 = (!v44)
                                                    fmt.Println(fi28, v2, v45)
                                                    _ = len(v36)
                                                    v44 = v44
                                                }
                                                v36 = append(v36, (v3 * fi28))
                                                var v75 int16 = 1000
                                                _ = v75
                                            }
                                        } else {
                                            var v76 float32 = 1e10
                                            v3 = 4294967295
                                            v45 += (v45 / ((v45) | 1))
                                            _ = v76
                                        }
                                    case 128:
                                        var v77 uint16 = 1
                                        _ = v77
                                    }
                                }
                                v78 := GenS1{F0: (fi43 % ((fi28) | 1)), F1: 9223372036854775807}
                                _ = v78
                            }
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = v36
                }
            }
        default:
            var v79 int64 = -9223372036854775808
            _ = v79
        }
        _ = fi28
    }
    v80 := GenS3{F0: -691, F1: 0}
    v3 = 4294967295
    var v81 bool = (71 >= 9223372036854775807)
    v82 := []int{v3, (v3 / ((9223372036854775807) | 1)), 127}
    if (false || false) {
        v83 := GenS3{F0: -1, F1: -1}
        for fi84 := 0; fi84 < 14; fi84++ {
            v85 := GenS4{F0: 110}
            _ = fi84
            _ = v85
        }
        fmt.Println(v1, v81)
        _ = v83
    }
    fmt.Println(v3, v81, v1)
    v2 = v2
    fmt.Println(v81)
    switch "a" {
    case "foo bar":
        v2 -= ((3.14159 / (1e10 / 1e10)) - -2.25)
    case "\n\t":
        if (0 <= 187) {
            var v86 uint8 = 0
            _ = v86
        } else {
            v82 = append(v82, 9223372036854775807)
            switch (v3 << ((2) & 63)) {
            case 2147483647:
                for fi87 := 0; fi87 < 1; fi87++ {
                    v81 = false
                    _ = fi87
                }
            case 32768:
                var v88 float64 = (float64(1e10) / v2)
                _ = v88
            case 9223372036854775807:
                switch "\n\t" {
                case "hello":
                    v81 = (int32(v3) >= 782)
                }
            }
            for fi89 := 0; fi89 < 7; fi89++ {
                for fi90 := 0; fi90 < 14; fi90++ {
                    var v91 uint32 = 0
                    var v92 int16 = 42
                    _ = fi90
                    _ = v91
                    _ = v92
                }
                _ = fi89
            }
        }
    default:
        var v93 uint64 = 459
        _ = v93
    }
    if (!(808 >= 32767)) {
        _ = len(v82)
    }
    v3 = v3
    if v81 {
        var v94 uint32 = 762
        _ = v94
    }
    for fi95 := 0; fi95 < 5; fi95++ {
        if true {
            var v96 uint32 = genFn0()
            for fi97 := 0; fi97 < 20; fi97++ {
                if (!v1) {
                    v81 = v81
                    var v98 uint32 = v96
                    fmt.Println(v98, v81)
                    _ = len(v82)
                    _ = v98
                }
                v82 = append(v82, (v3 / ((-9223372036854775808) | 1)))
                _ = fi97
            }
            _ = v96
        } else {
            switch "hello" {
            case "":
                var v99 int32 = -1
                _ = v99
            default:
                v3 = (v3 * (^(v3)))
            }
            switch ("foo bar" + "") {
            case "":
                v82 = append(v82, v3)
            case "a":
                var v100 uint64 = 2
                _ = v100
            default:
                switch 647 {
                case 1:
                    v101 := GenS6{F0: 2, F1: 1000, F2: 255}
                    _ = v101
                case 4294967295:
                    var v102 uint32 = 0
                    _ = v102
                default:
                    v103 := GenS7{F0: 488, F1: -9223372036854775808, F2: -128}
                    _ = v103
                }
            }
        }
        v104 := GenS7{F0: 42, F1: 625, F2: -120}
        _ = fi95
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v80
    _ = v81
    _ = v82
}
