package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 uint32 }
type GenS11 struct { F0 int32; F1 bool }

func main() {
    var v1 uint64 = 95
    var v2 bool = (3.14159 <= -2.25)
    var v3 uint16 = 1
    var v4 bool = (-128 < 61)
    v5 := []uint32{2147483647}
    v6 := []uint32{42, 4294967295, 42, 65535}
    if v2 {
        v1 -= v1
        var v7 float64 = 0.0
        switch "\n\t" {
        case "a":
            switch 32767 {
            case -1:
                v8 := GenS1{F0: (v3 << ((0) & 63))}
                _ = v8
            case -823:
                if v2 {
                    fmt.Println(v4)
                }
            case 0:
                if len(v6) > 0 {
                    v6[0] = uint32(v3)
                }
            }
        case "hello":
            v9 := GenS2{F0: -1, F1: 2147483648}
            _ = v9
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = 603
            }
        }
        _ = v7
    }
    var v10 string = ("a" + ("hello" + "hello"))
    for fi11 := 0; fi11 < 6; fi11++ {
        v2 = v2
        if len(v6) > 0 {
            v6[0] = 2
        }
        var v12 uint16 = 32767
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int64 = -9223372036854775808
        v15 := GenS2{F0: (v14 / (((^(v14))) | 1)), F1: uint32(v3)}
        v4 = (!(42 > -1))
        _ = fi13
        _ = v14
        _ = v15
    }
    if (true || (42 >= -390)) {
        v6 = append(v6, 2147483647)
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v10)
            _ = fi16
        }
    } else {
        var v17 float32 = 0.1
        if (0 < 653) {
            if true {
                var v18 uint = 4294967295
                v1 -= v1
                _ = v18
            }
        } else {
            var v19 uint8 = 255
            if (false && true) {
                switch int(v3) {
                case 127:
                    if v4 {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v3 = 32767
                            if (!false) {
                                switch v10 {
                                case "foo bar":
                                    var v21 uint8 = (v19 >> ((107) & 63))
                                    _ = v21
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v10 = v10
                            }
                            if ((v1 << ((871) & 63)) < (^(v1))) {
                                v2 = (1.5 <= 0.0)
                                if (v3 <= 804) {
                                    for fi22 := 0; fi22 < 13; fi22++ {
                                        if v2 {
                                            switch "hello" {
                                            case "":
                                                var v23 int8 = 42
                                                _ = v23
                                            case "\n\t":
                                                switch (fi22 - 128) {
                                                case -734:
                                                    var v24 int64 = 42
                                                    _ = v24
                                                case 338:
                                                    v1 -= (v1 + 20)
                                                    if false {
                                                        var v25 bool = v4
                                                        if (v2 && ((!v2) && v4)) {
                                                            v26 := GenS1{F0: v3}
                                                            if (v2 || true) {
                                                                fmt.Println(v10)
                                                            }
                                                            _ = v26
                                                        } else {
                                                            v27 := GenS1{F0: 42}
                                                            _ = v27
                                                        }
                                                        var v28 uint8 = uint8(v1)
                                                        _ = v25
                                                        _ = v28
                                                    }
                                                    var v29 float64 = 1e10
                                                    _ = v29
                                                case 42:
                                                    var v30 int8 = 100
                                                    _ = v30
                                                default:
                                                    v31 := GenS2{F0: -352, F1: 235}
                                                    _ = v31
                                                }
                                            }
                                        } else {
                                            var v32 uint16 = v3
                                            fmt.Println(fi20, v4, v17)
                                            _ = v32
                                        }
                                        v33 := GenS3{F0: 2147483647}
                                        if (42 <= 1) {
                                            var v34 uint8 = v19
                                            _ = v34
                                        } else {
                                            v6 = append(v6, 2147483648)
                                            fmt.Println(v4)
                                        }
                                        _ = fi22
                                        _ = v33
                                    }
                                } else {
                                    if true {
                                        if v2 {
                                            if v2 {
                                                for fi35 := 0; fi35 < 12; fi35++ {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = fi35
                                                }
                                                var v36 int = fi20
                                                var v37 bool = (42 == 42)
                                                _ = v36
                                                _ = v37
                                            }
                                            switch v10 {
                                            case "":
                                                switch ((v10 + "") + "\n\t") {
                                                case "a":
                                                    v38 := GenS1{F0: v3}
                                                    _ = v38
                                                default:
                                                    v3 = (v3 % ((v3) | 1))
                                                }
                                            }
                                        }
                                        v3 += 42
                                        if v2 {
                                            var v39 uint64 = v1
                                            _ = v39
                                        } else {
                                            var v40 float32 = -2.25
                                            fmt.Println(v2)
                                            v41 := GenS1{F0: 255}
                                            _ = v40
                                            _ = v41
                                        }
                                    } else {
                                        v1 -= v1
                                        var v42 bool = v2
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 13; fi43++ {
                                        fmt.Println(v1, fi20, v19)
                                        _ = fi43
                                    }
                                }
                                v44 := GenS4{F0: 1, F1: 65535, F2: 225}
                                _ = v44
                            }
                            _ = fi20
                        }
                    }
                case 32768:
                    switch 1 {
                    case 127:
                        var v45 uint32 = 4294967295
                        _ = v45
                    default:
                        switch -173 {
                        case -636:
                            for fi46 := 0; fi46 < 7; fi46++ {
                                if v4 {
                                    for fi47 := 0; fi47 < 15; fi47++ {
                                        if v2 {
                                            if (!(false || ("\n\t" <= (((v10 + (("" + "a") + v10)) + v10) + v10)))) {
                                                v48 := GenS5{F0: 4294967295}
                                                _ = v48
                                            }
                                        } else {
                                            fmt.Println(v1)
                                            if true {
                                                v3 = 889
                                                v19 -= (v19 | (^(v19)))
                                                v2 = (1 < 0)
                                            }
                                        }
                                        _ = fi47
                                    }
                                    var v49 int64 = 42
                                    _ = len(v5)
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: 0, F1: 1, F2: 4294967295}
                                    switch (-(fi46)) {
                                    case -746:
                                        for fi51 := 0; fi51 < 15; fi51++ {
                                            v52 := GenS6{F0: 1}
                                            var v53 int = (fi51 + -695)
                                            v54 := GenS7{F0: 18446744073709551615}
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                    default:
                                        var v55 uint16 = 0
                                        _ = v55
                                    }
                                    _ = len(v5)
                                    _ = v50
                                }
                                var v56 bool = v2
                                _ = fi46
                                _ = v56
                            }
                        case 661:
                            var v57 uint16 = v3
                            _ = v57
                        default:
                            var v58 int8 = 42
                            _ = v58
                        }
                    }
                default:
                    v2 = true
                }
                var v59 string = v10
                v59 = ("a" + (v59 + (v10 + (v10 + v10))))
                _ = v59
            } else {
                v19 -= (v19 / (((v19 | v19)) | 1))
                v17 = (0.0 * ((float32(-2.25) / 3.14159) / (0.1 / 3.14159)))
                v19 = v19
            }
            _ = v19
        }
        if (!(1 == 42)) {
            v60 := GenS7{F0: (v1 & v1)}
            fmt.Println(v1)
            _ = v60
        }
        _ = v17
    }
    fmt.Println(v3, v3)
    switch 127 {
    case 0:
        v10 = "\n\t"
    case 128:
        v4 = v2
    default:
        if false {
            fmt.Println(v2)
            var v61 int16 = -1
            v62 := GenS8{F0: 2}
            _ = v61
            _ = v62
        } else {
            var v63 uint32 = 166
            v64 := GenS9{F0: 32767}
            for fi65 := 0; fi65 < 17; fi65++ {
                if (617 != 0) {
                    v5 = append(v5, v63)
                }
                v1 += 270
                v1 -= 2
                _ = fi65
            }
            _ = v63
            _ = v64
        }
    }
    for fi66 := 0; fi66 < 17; fi66++ {
        var v67 int8 = -128
        _ = fi66
        _ = v67
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((v3 + (v3 & 147)) == v3) {
        for fi68 := 0; fi68 < 1; fi68++ {
            if true {
                fmt.Println(fi68)
            }
            if true {
                fmt.Println(v1)
            } else {
                v2 = false
                v69 := GenS6{F0: 100}
                var v70 float32 = 1e10
                _ = v69
                _ = v70
            }
            _ = fi68
        }
    } else {
        v10 = v10
        if false {
            for fi71 := 0; fi71 < 1; fi71++ {
                if v2 {
                    fmt.Println(v1, v2, fi71)
                }
                _ = fi71
            }
            v72 := GenS9{F0: -933}
            _ = v72
        }
        if (!v2) {
            v1 += v1
            if false {
                v2 = (0 < 0)
                var v73 int16 = 1000
                _ = v73
            } else {
                v4 = (v2 || (380 < 42))
                switch 128 {
                case -1:
                    if (18446744073709551615 >= 284) {
                        if (86 > 2) {
                            v4 = (1 == 1000)
                        }
                        v5 = append(v5, 1)
                        v10 = v10
                    }
                }
                v74 := GenS10{F0: 155, F1: 2}
                _ = v74
            }
            if (v2 || (42 <= 684)) {
                v75 := GenS6{F0: 1}
                switch -747 {
                case 42:
                    for fi76 := 0; fi76 < 1; fi76++ {
                        var v77 int8 = 127
                        fmt.Println(v2)
                        _ = fi76
                        _ = v77
                    }
                case 4294967295:
                    v78 := GenS7{F0: 119}
                    _ = v78
                case 2147483647:
                    v79 := GenS1{F0: v3}
                    _ = v79
                default:
                    fmt.Println(v1)
                }
                _ = v75
            } else {
                v3 -= (^(v3))
                switch (v10 + v10) {
                case "foo bar":
                    var v80 uint32 = 0
                    _ = v80
                case "\n\t":
                    v4 = ((!v4) || (0 <= 42))
                default:
                    var v81 float32 = 0.1
                    _ = v81
                }
                v3 += v3
            }
        } else {
            for fi82 := 0; fi82 < 7; fi82++ {
                for fi83 := 0; fi83 < 10; fi83++ {
                    var v84 bool = (!true)
                    _ = fi83
                    _ = v84
                }
                switch "\n\t" {
                case "":
                    var v85 int64 = 1
                    _ = v85
                case "a":
                    switch fi82 {
                    case 445:
                        for fi86 := 0; fi86 < 10; fi86++ {
                            switch (v10 + (v10 + "hello")) {
                            case "":
                                if (4294967295 > v1) {
                                    var v87 int8 = -128
                                    _ = v87
                                } else {
                                    var v88 uint64 = 2
                                    v89 := GenS2{F0: -9223372036854775808, F1: 65535}
                                    v90 := GenS6{F0: -128}
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                }
                            case "\n\t":
                                v3 = 42
                            case "foo bar":
                                v1 = 18446744073709551615
                            default:
                                var v91 int = fi82
                                _ = v91
                            }
                            _ = fi86
                        }
                    case 128:
                        if v2 {
                            for fi92 := 0; fi92 < 9; fi92++ {
                                var v93 uint16 = v3
                                fmt.Println(v4)
                                _ = fi92
                                _ = v93
                            }
                            v5 = append(v5, 2)
                        }
                    case 9223372036854775807:
                        for fi94 := 0; fi94 < 9; fi94++ {
                            if true {
                                switch "hello" {
                                case "a":
                                    var v95 int64 = -9223372036854775808
                                    _ = v95
                                case "":
                                    if false {
                                        var v96 float32 = 0.0
                                        if (2 < 255) {
                                            if v4 {
                                                if (2 > v1) {
                                                    if ((1 == 4294967295) && (!(42 > 1))) {
                                                        v97 := GenS9{F0: 4294967295}
                                                        var v98 int64 = -9223372036854775808
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        if (!(v1 >= 71)) {
                                                            switch v10 {
                                                            case "":
                                                                var v99 float64 = 1.5
                                                                _ = v99
                                                            }
                                                        }
                                                    }
                                                    if (!v2) {
                                                        switch (v10 + (v10 + "a")) {
                                                        case "a":
                                                            var v100 uint8 = 128
                                                            _ = v100
                                                        case "\n\t":
                                                            var v101 int8 = -128
                                                            _ = v101
                                                        case "":
                                                            v102 := GenS8{F0: 42}
                                                            _ = v102
                                                        default:
                                                            var v103 uint16 = 32767
                                                            _ = v103
                                                        }
                                                        v4 = v4
                                                        var v104 int8 = -114
                                                        _ = v104
                                                    }
                                                    fmt.Println(v4)
                                                } else {
                                                    var v105 uint64 = (v1 * (v1 - (v1 | v1)))
                                                    fmt.Println(v1, v4, v3)
                                                    _ = v105
                                                }
                                                var v106 uint16 = uint16(v1)
                                                v107 := GenS11{F0: 1, F1: v4}
                                                _ = v106
                                                _ = v107
                                            }
                                            v96 += ((float32((0.1 - 1e10)) / (v96 - float32(((0.1 * (0.1 + 1e10)) + 0.0)))) + v96)
                                            var v108 float64 = float64((((v96 + (float32((1.5 * float64(0.1))) + v96)) - v96) / v96))
                                            _ = v108
                                        } else {
                                            var v109 int8 = -1
                                            v109 -= v109
                                            v110 := GenS5{F0: v1}
                                            _ = v109
                                            _ = v110
                                        }
                                        var v111 uint16 = (v3 + 255)
                                        v112 := GenS5{F0: (v1 >> ((0) & 63))}
                                        _ = v111
                                        _ = v112
                                        _ = v96
                                    } else {
                                        v3 -= (v3 >> ((4294967295) & 63))
                                        fmt.Println(fi82, v3)
                                        v1 += v1
                                    }
                                case "\n\t":
                                    v113 := GenS2{F0: 9223372036854775807, F1: 65535}
                                    _ = v113
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            } else {
                                fmt.Println(v10)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v1 = 18446744073709551615
                            }
                            if false {
                                fmt.Println(v4, v1)
                                v5 = append(v5, 649)
                                v114 := GenS2{F0: -9223372036854775808, F1: 586}
                                _ = v114
                            }
                            _ = fi94
                        }
                    }
                }
                var v115 float64 = ((((((3.14159 / 1.5) / 1.0) / 1.5) - 1e10) - (1e10 * ((-2.25 * 0.1) / 3.14159))) * 0.0)
                _ = fi82
                _ = v115
            }
            var v116 int8 = 42
            _ = v116
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
