package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int; F1 string; F2 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint; F1 uint64; F2 int32 }

func main() {
    v1 := []int{-2147483648, -660, -512}
    var v2 uint8 = 2
    var v3 int = 0
    var v4 int = (v3 + v3)
    var v5 int16 = -990
    v6 := []int{v3}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch (("foo bar" + "") + "a") {
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint = 42
            _ = len(v6)
            _ = fi7
            _ = v8
        }
    }
    if (42 <= 190) {
        if len(v1) > 0 {
            v1[0] = 0
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v5)
            var v10 int16 = (v5 + v5)
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int16 = (v5 | v5)
            var v13 uint32 = 2147483647
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 uint16 = 915
        _ = v14
    }
    v15 := GenS1{F0: 2147483648, F1: 610}
    switch (v3 * v3) {
    case 2147483647:
        v5 = v5
    case 946:
        var v16 string = "a"
        _ = v16
    case -2147483648:
        if false {
            if len(v6) > 0 {
                v6[0] = (v4 & v4)
            }
            var v17 uint32 = 4294967295
            _ = v17
        } else {
            fmt.Println(v4)
            if true {
                var v18 float64 = (1e10 / 1.0)
                _ = v18
            }
            v3 = v3
        }
    default:
        for fi19 := 0; fi19 < 11; fi19++ {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                _ = len(v1)
                var v20 string = "foo bar"
                var v21 float32 = 0.0
                _ = v20
                _ = v21
            }
            switch 127 {
            case 405:
                var v22 int64 = -1
                _ = v22
            default:
                if false {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v23 string = "a"
                        _ = v23
                    case "\n\t":
                        v3 = (fi19 << ((2) & 63))
                    case "hello":
                        v24 := GenS2{F0: v2, F1: 1, F2: -1}
                        _ = v24
                    default:
                        v25 := GenS3{F0: "a", F1: (^(v2))}
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v4 -= (-(v3))
                        _ = fi26
                    }
                    switch (^(v4)) {
                    case 9223372036854775807:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if true {
                            var v27 int = fi19
                            v3 -= 1
                            _ = v27
                        }
                    }
                } else {
                    fmt.Println(fi19, v4)
                    if true {
                        v4 = (v4 | 508)
                        var v28 float64 = 1.5
                        if len(v1) > 0 {
                            v1[0] = -1
                        }
                        _ = v28
                    } else {
                        v3 += (v4 - (v4 | fi19))
                    }
                    v29 := GenS2{F0: v2, F1: 42, F2: -512}
                    _ = v29
                }
            }
            if (180 == (^(v2))) {
                v3 = fi19
                fmt.Println(v5)
            } else {
                var v30 uint8 = v2
                for fi31 := 0; fi31 < 4; fi31++ {
                    v2 = 128
                    var v32 float64 = 3.14159
                    switch "" {
                    case "":
                        if (true && (32767 > 65535)) {
                            fmt.Println(v4, fi31)
                        }
                    case "a":
                        v2 = v2
                    default:
                        v4 -= -728
                    }
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
            _ = fi19
        }
    }
    var v33 int16 = v5
    if (v2 != v2) {
        v33 = int16(v2)
        var v34 string = "hello"
        _ = v34
    } else {
        var v35 uint16 = 949
        v36 := GenS3{F0: "hello", F1: (v2 / ((2) | 1))}
        var v37 float32 = (1e10 - (1e10 - 0.1))
        _ = v35
        _ = v36
        _ = v37
    }
    v5 -= 172
    var v38 bool = (0 < 1)
    _ = len(v6)
    if false {
        switch ("hello" + "hello") {
        case "":
            v3 += v3
        case "hello":
            v39 := GenS4{F0: 1, F1: 255}
            _ = v39
        }
        if (!((-2.25 - 3.14159) == 1.5)) {
            if v38 {
                v38 = v38
                var v40 uint32 = 42
                for fi41 := 0; fi41 < 7; fi41++ {
                    v4 -= 2147483647
                    _ = fi41
                }
                _ = v40
            } else {
                for fi42 := 0; fi42 < 10; fi42++ {
                    if v38 {
                        for fi43 := 0; fi43 < 12; fi43++ {
                            fmt.Println(v2)
                            switch 42 {
                            case 42:
                                v44 := GenS2{F0: 58, F1: 1, F2: 980}
                                var v45 float64 = 0.0
                                _ = v44
                                _ = v45
                            case 32768:
                                v33 = (v33 >> ((65535) & 63))
                            case 32767:
                                v4 = -2147483648
                            }
                            _ = fi43
                        }
                        var v46 uint64 = 2
                        switch "\n\t" {
                        case "":
                            for fi47 := 0; fi47 < 2; fi47++ {
                                var v48 uint16 = 430
                                _ = fi47
                                _ = v48
                            }
                        }
                        _ = v46
                    }
                    if ("hello" == ("" + "a")) {
                        if (-4 <= 1) {
                            v5 = 0
                            var v49 string = "hello"
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 9; fi50++ {
                                if (686 >= 457) {
                                    if false {
                                        v5 = v5
                                        if ("\n\t" >= "a") {
                                            var v51 uint8 = (^(v2))
                                            v33 = (-(v33))
                                            var v52 string = "\n\t"
                                            _ = v51
                                            _ = v52
                                        } else {
                                            switch fi50 {
                                            case 473:
                                                switch "a" {
                                                case "foo bar":
                                                    v53 := GenS4{F0: 255, F1: 18446744073709551615}
                                                    _ = v53
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "hello":
                                                        switch "" {
                                                        case "hello":
                                                            v54 := GenS5{F0: v38}
                                                            _ = v54
                                                        case "\n\t":
                                                            if v38 {
                                                                var v55 uint64 = 300
                                                                var v56 string = "hello"
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        case "a":
                                                            var v57 int8 = 100
                                                            _ = v57
                                                        }
                                                    default:
                                                        if (0.0 == 0.0) {
                                                            v58 := GenS5{F0: (9223372036854775807 < int64(v33))}
                                                            if (!(v4 == (fi50 % ((fi50) | 1)))) {
                                                                fmt.Println(fi42)
                                                            }
                                                            if (!v38) {
                                                                v59 := GenS1{F0: 182, F1: 2}
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        } else {
                                                            var v60 int16 = v33
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                            default:
                                                if len(v6) > 0 {
                                                    v6[0] = (v4 & 127)
                                                }
                                                fmt.Println(v5, fi50, v5)
                                                v4 += (fi50 & 2147483647)
                                            }
                                            v61 := GenS6{F0: v5}
                                            switch (fi42 - (fi50 | (fi42 % (((v3 % (((^(v3))) | 1))) | 1)))) {
                                            case 9223372036854775807:
                                                var v62 uint32 = 1
                                                _ = v62
                                            case -2147483648:
                                                switch v4 {
                                                case -2147483648:
                                                    v63 := GenS7{F0: -1, F1: -32768}
                                                    _ = v63
                                                case 0:
                                                    v3 = int(v2)
                                                }
                                            default:
                                                v64 := GenS8{F0: -9223372036854775808, F1: "a", F2: v33}
                                                _ = v64
                                            }
                                            var v65 int32 = 42
                                            _ = v61
                                            _ = v65
                                        }
                                    }
                                    v5 = (^(v5))
                                }
                                v66 := GenS6{F0: 1000}
                                _ = fi50
                                _ = v66
                            }
                            v67 := GenS9{F0: v2}
                            _ = v67
                        }
                        if false {
                            if false {
                                v68 := GenS1{F0: 1, F1: 18446744073709551615}
                                for fi69 := 0; fi69 < 9; fi69++ {
                                    fmt.Println(fi42, v3, v2)
                                    var v70 string = ""
                                    v38 = v38
                                    _ = fi69
                                    _ = v70
                                }
                                _ = v68
                            } else {
                                for fi71 := 0; fi71 < 3; fi71++ {
                                    v5 = 0
                                    _ = fi71
                                }
                                v3 = (fi42 - v4)
                            }
                            for fi72 := 0; fi72 < 16; fi72++ {
                                fmt.Println(v5, fi42)
                                _ = fi72
                            }
                        } else {
                            if v38 {
                                var v73 uint = 23
                                for fi74 := 0; fi74 < 3; fi74++ {
                                    var v75 float64 = -2.25
                                    var v76 uint16 = 1
                                    if v38 {
                                        if (42 < 1) {
                                            switch -745 {
                                            case 1:
                                                if (true && v38) {
                                                    v77 := GenS7{F0: 739, F1: 1000}
                                                    switch "" {
                                                    case "a":
                                                        var v78 string = "hello"
                                                        _ = v78
                                                    case "":
                                                        var v79 int8 = 42
                                                        _ = v79
                                                    case "foo bar":
                                                        v80 := GenS6{F0: v33}
                                                        _ = v80
                                                    default:
                                                        if true {
                                                            v75 -= v75
                                                        } else {
                                                            v75 -= v75
                                                            v81 := GenS9{F0: 74}
                                                            _ = v81
                                                        }
                                                    }
                                                    var v82 uint16 = v76
                                                    _ = v77
                                                    _ = v82
                                                } else {
                                                    v76 += v76
                                                }
                                            default:
                                                var v83 uint16 = v76
                                                _ = v83
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = (fi74 * fi74)
                                            }
                                            v38 = (2 <= 127)
                                            v1 = append(v1, (v4 & (fi42 / ((v4) | 1))))
                                            v38 = false
                                            var v84 int32 = -530
                                            v4 = fi74
                                            _ = v84
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v85 int32 = 0
                                                _ = v85
                                            }
                                        }
                                        var v86 uint = v73
                                        if (1 != 4294967295) {
                                            v1 = append(v1, 4294967295)
                                            v3 = 32767
                                        } else {
                                            v87 := GenS2{F0: 42, F1: 1, F2: 0}
                                            _ = v87
                                        }
                                        _ = v86
                                    } else {
                                        var v88 uint = v73
                                        fmt.Println(v4, v88, v73)
                                        _ = v88
                                    }
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                                _ = v73
                            }
                            var v89 uint64 = 856
                            var v90 float64 = 1e10
                            _ = v89
                            _ = v90
                        }
                    }
                    switch "a" {
                    case "":
                        switch fi42 {
                        case -1:
                            for fi91 := 0; fi91 < 8; fi91++ {
                                if ((true || ((v2 >> ((65535) & 63)) >= (^(v2)))) && (42 != 642)) {
                                    _ = len(v1)
                                    var v92 uint16 = 790
                                    v6 = append(v6, (fi42 * v4))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v92
                                }
                                _ = fi91
                            }
                        default:
                            var v93 int64 = -1
                            _ = v93
                        }
                    default:
                        var v94 int64 = 9223372036854775807
                        _ = v94
                    }
                    _ = fi42
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v95 bool = (!(2147483647 >= 198))
                _ = v95
            }
            fmt.Println(v33, v3, v5)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v3 = 409
    } else {
        v96 := GenS8{F0: v4, F1: "\n\t", F2: v33}
        var v97 int32 = -2147483648
        _ = v96
        _ = v97
    }
    if ((!v38) && v38) {
        v98 := GenS5{F0: false}
        v33 = v5
        _ = v98
    } else {
        for fi99 := 0; fi99 < 8; fi99++ {
            for fi100 := 0; fi100 < 17; fi100++ {
                v33 += v33
                var v101 int64 = 0
                _ = fi100
                _ = v101
            }
            for fi102 := 0; fi102 < 6; fi102++ {
                if len(v6) > 0 {
                    v6[0] = (-(fi99))
                }
                switch "" {
                case "a":
                    var v103 uint16 = 1
                    _ = v103
                case "hello":
                    v104 := GenS10{F0: 127}
                    _ = v104
                case "foo bar":
                    v2 += 127
                default:
                    var v105 int8 = 1
                    _ = v105
                }
                _ = fi102
            }
            _ = fi99
        }
    }
    if v38 {
        if (((3.14159 * (float64((0.1 + 1e10)) - 0.1)) / (0.1 - float64(3.14159))) > 3.14159) {
            if (!v38) {
                var v106 uint32 = 65535
                fmt.Println(v5)
                var v107 int64 = 42
                _ = v106
                _ = v107
            }
        } else {
            var v108 uint = 65535
            if v38 {
                fmt.Println(v38)
                for fi109 := 0; fi109 < 3; fi109++ {
                    v5 = v33
                    var v110 uint64 = 18446744073709551615
                    _ = fi109
                    _ = v110
                }
                var v111 uint64 = 52
                _ = v111
            }
            _ = v108
        }
        switch ((("foo bar" + "foo bar") + "hello") + ("" + "")) {
        case "hello":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            v2 += v2
        }
    } else {
        if (!v38) {
            if v38 {
                v4 += (v4 << ((0) & 63))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v112 := GenS11{F0: 255, F1: uint64(v2), F2: -592}
                _ = v112
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
