package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int8, p1 int16, p2 uint) string {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 += (p0 | 0)
        _ = fi6
    }
    var v7 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "a"
}

func main() {
    var v1 uint8 = 219
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2147483647, 2147483648, 2147483647}
    v5 := make(map[uint32]uint32)
    fmt.Println(v1)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = v1
        fmt.Println(v9, v1, fi8)
        _ = fi8
        _ = v9
    }
    if (-881 < -305) {
        v1 = v1
        _ = len(v4)
        v10 := GenS1{F0: 265}
        _ = v10
    }
    if true {
        var v11 bool = true
        if (!v11) {
            switch "a" {
            case "a":
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            case "\n\t":
                _ = len(v5)
            case "hello":
                switch ("" + "") {
                case "":
                    var v12 string = "a"
                    _ = v12
                default:
                    if false {
                        v1 += (v1 - v1)
                        switch -2147483648 {
                        case 2147483647:
                            var v13 int64 = 42
                            _ = v13
                        default:
                            for fi14 := 0; fi14 < 1; fi14++ {
                                fmt.Println(v11)
                                for fi15 := 0; fi15 < 9; fi15++ {
                                    var v16 string = "hello"
                                    switch (v16 + v16) {
                                    case "":
                                        fmt.Println(v11, fi15)
                                    }
                                    switch (fi14 * fi15) {
                                    case -457:
                                        v17 := GenS2{F0: 2, F1: (v16 + v16), F2: ""}
                                        _ = v17
                                    case 127:
                                        var v18 int8 = 100
                                        _ = v18
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                                v11 = true
                                _ = fi14
                            }
                        }
                        switch 32767 {
                        case 32768:
                            if (!(!v11)) {
                                if v11 {
                                    var v19 int64 = -1
                                    _ = v19
                                } else {
                                    switch ("" + "") {
                                    case "\n\t":
                                        for fi20 := 0; fi20 < 6; fi20++ {
                                            v21 := GenS3{F0: 1, F1: -2147483648}
                                            var v22 string = "hello"
                                            for fi23 := 0; fi23 < 16; fi23++ {
                                                var v24 uint16 = 511
                                                _ = fi23
                                                _ = v24
                                            }
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                    case "":
                                        var v25 int32 = -245
                                        _ = v25
                                    case "a":
                                        fmt.Println(v1, v1)
                                    default:
                                        v11 = (!(float32(0.1) <= 1e10))
                                    }
                                    var v26 int64 = 1
                                    _ = v26
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v1 = (v1 ^ (v1 % ((v1) | 1)))
                            }
                        default:
                            for fi27 := 0; fi27 < 2; fi27++ {
                                v11 = false
                                v11 = (1 <= 646)
                                _ = fi27
                            }
                        }
                    } else {
                        v1 = (v1 + v1)
                        v1 += 127
                    }
                }
            }
            v28 := GenS2{F0: 42, F1: "a", F2: ("hello" + "foo bar")}
            v11 = ((1e10 * 1.5) == 0.1)
            _ = v28
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                for fi30 := 0; fi30 < 20; fi30++ {
                    v31 := GenS1{F0: 42}
                    if v11 {
                        var v32 int64 = -955
                        _ = v32
                    } else {
                        v11 = v11
                        var v33 uint32 = 886
                        _ = v33
                    }
                    switch "a" {
                    case "foo bar":
                        if false {
                            _ = len(v4)
                            _ = len(v3)
                            var v34 string = ("foo bar" + "hello")
                            var v35 string = v34
                            _ = v34
                            _ = v35
                        } else {
                            v1 += (^(v1))
                            v36 := GenS3{F0: 42, F1: -1}
                            v37 := GenS2{F0: 42, F1: "", F2: ("a" + (("" + "\n\t") + "hello"))}
                            _ = v36
                            _ = v37
                        }
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        v1 += (v1 ^ v1)
                    default:
                        if (((1e10 * -2.25) * 0.0) == 1.5) {
                            v11 = v11
                        } else {
                            fmt.Println(fi30)
                            v11 = v11
                        }
                    }
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            fmt.Println(v11, v11)
            if (!v11) {
                v1 = v1
                if v11 {
                    var v38 uint16 = 372
                    _ = v38
                } else {
                    v4 = append(v4, 786)
                    if (4294967295 >= 18446744073709551615) {
                        var v39 float64 = (1.5 - -2.25)
                        var v40 uint = 0
                        v41 := GenS2{F0: 878, F1: ("" + ""), F2: "foo bar"}
                        _ = v39
                        _ = v40
                        _ = v41
                    } else {
                        v42 := GenS4{F0: (42 == 322), F1: v11, F2: 0}
                        v11 = (42 != -1)
                        v11 = (-104 > -128)
                        _ = v42
                    }
                }
            } else {
                if (0 == 42) {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        delete(v2, 75)
                        if ("a" == "") {
                            var v44 uint = 2
                            _ = len(v2)
                            _ = v44
                        } else {
                            var v45 int16 = 0
                            switch "foo bar" {
                            case "hello":
                                _ = len(v4)
                            }
                            _ = v45
                        }
                        fmt.Println(v11, fi43, v11)
                        _ = fi43
                    }
                } else {
                    var v46 float32 = (1.5 * ((((1.5 * 3.14159) + -2.25) / 1.0) + 3.14159))
                    _ = v46
                }
                for fi47 := 0; fi47 < 6; fi47++ {
                    _ = len(v5)
                    _ = fi47
                }
                v48 := GenS2{F0: 0, F1: (("foo bar" + "foo bar") + (("\n\t" + "") + ("hello" + "\n\t"))), F2: "hello"}
                _ = v48
            }
        }
        for fi49 := 0; fi49 < 3; fi49++ {
            fmt.Println(fi49, fi49)
            _ = fi49
        }
        _ = v11
    }
    if ((("" + "hello") + "a") >= "a") {
        v4 = append(v4, 2147483647)
        var v50 bool = ((0.0 <= 0.1) && true)
        _ = v50
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        var v52 uint16 = 137
        if true {
            var v53 int32 = 1
            var v54 uint64 = 102
            _ = v53
            _ = v54
        }
        if true {
            for fi55 := 0; fi55 < 13; fi55++ {
                fmt.Println(v52, v1)
                fmt.Println(v52)
                _ = fi55
            }
            fmt.Println(fi51, v1)
            if len(v4) > 0 {
                v4[0] = 0
            }
        } else {
            delete(v3, 524)
            var v56 int32 = 2147483647
            _ = v56
        }
        _ = fi51
        _ = v52
    }
    v1 = (v1 >> ((2) & 63))
    v57 := GenS3{F0: 122, F1: 42}
    v1 -= (v1 << ((42) & 63))
    if (!(42 > 648)) {
        var v58 int32 = 1
        if (18 < 1) {
            var v59 bool = true
            var v60 uint8 = (v1 % (((v1 % ((v1) | 1))) | 1))
            _ = v59
            _ = v60
        }
        switch -128 {
        case 4294967295:
            v58 -= (-(v58))
        case 32767:
            fmt.Println(v1, v58)
        case 510:
            if (false || (-32768 >= 42)) {
                if (42 > 2) {
                    switch 9223372036854775807 {
                    case -418:
                        v1 = v1
                    case 9223372036854775807:
                        v61 := GenS5{F0: 255, F1: 1000}
                        _ = v61
                    case 1:
                        var v62 uint16 = 65535
                        _ = v62
                    }
                    var v63 string = "foo bar"
                    fmt.Println(v58)
                    _ = v63
                } else {
                    v1 = (v1 << ((255) & 63))
                    v64 := GenS6{F0: 42}
                    _ = v64
                }
            } else {
                v65 := GenS4{F0: ((0 < 0) || false), F1: (v58 < 0), F2: 9223372036854775807}
                delete(v5, 830)
                for fi66 := 0; fi66 < 19; fi66++ {
                    switch 1 {
                    case -9223372036854775808:
                        switch (^(fi66)) {
                        case 196:
                            var v67 string = "foo bar"
                            _ = v67
                        }
                    case 4294967295:
                        if (59 != 2) {
                            for fi68 := 0; fi68 < 13; fi68++ {
                                switch "hello" {
                                case "":
                                    fmt.Println(fi66, v1)
                                default:
                                    switch (((("foo bar" + ("foo bar" + ("" + "foo bar"))) + "a") + ("a" + "hello")) + (((("a" + "a") + "\n\t") + "hello") + (("" + "") + "foo bar"))) {
                                    case "foo bar":
                                        switch int(v58) {
                                        case 1:
                                            var v69 uint32 = 43
                                            _ = v69
                                        case -455:
                                            v58 += v58
                                        case 302:
                                            var v70 uint16 = 2
                                            if len(v4) > 0 {
                                                v4[0] = 2147483648
                                            }
                                            _ = v70
                                        default:
                                            if false {
                                                switch (fi68 << ((2) & 63)) {
                                                case 127:
                                                    switch fi66 {
                                                    case 42:
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v71 int8 = 100
                                                            _ = v71
                                                        default:
                                                            var v72 string = "\n\t"
                                                            var v73 float32 = 0.1
                                                            v74 := GenS7{F0: 255, F1: 127}
                                                            _ = v72
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                    default:
                                                        v75 := GenS8{F0: 0}
                                                        _ = v75
                                                    }
                                                case 42:
                                                    if false {
                                                        if false {
                                                            if ((3.14159 * (0.0 + 0.0)) > (1e10 + float64((1.5 / 1.0)))) {
                                                                v58 = v58
                                                                v76 := GenS6{F0: 16}
                                                                v58 = (v58 + v58)
                                                                _ = v76
                                                            } else {
                                                                fmt.Println(fi68, fi66, fi68)
                                                                if ((true || false) && ("a" <= "")) {
                                                                    if (!(!true)) {
                                                                        v58 = 42
                                                                        var v77 bool = true
                                                                        var v78 uint8 = 127
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                } else {
                                                                    var v79 int = fi68
                                                                    var v80 int64 = 9223372036854775807
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                var v81 int8 = 42
                                                                _ = v81
                                                            }
                                                            v1 = (v1 % (((v1 << ((42) & 63))) | 1))
                                                            var v82 uint8 = v1
                                                            _ = v82
                                                        } else {
                                                            fmt.Println(fi66, v58)
                                                            var v83 uint = 1
                                                            _ = v83
                                                        }
                                                        _ = len(v4)
                                                        switch (fi66 - (fi68 | fi68)) {
                                                        case 1:
                                                            v84 := v3[4294967295]
                                                            _ = v84
                                                        case -1:
                                                            var v85 int64 = -1
                                                            _ = v85
                                                        case 4294967295:
                                                            var v86 uint8 = (v1 * v1)
                                                            _ = v86
                                                        }
                                                    }
                                                    var v87 int16 = -1
                                                    switch (fi66 & -529) {
                                                    case 258:
                                                        v58 -= 1
                                                    case 32768:
                                                        v1 -= v1
                                                    default:
                                                        v1 += v1
                                                        fmt.Println(v1, v87, fi66)
                                                    }
                                                    _ = v87
                                                case 372:
                                                    var v88 float32 = 1.5
                                                    _ = v88
                                                }
                                                var v89 uint32 = 4294967295
                                                _ = v89
                                            } else {
                                                v90 := GenS3{F0: -1, F1: v58}
                                                _ = v90
                                            }
                                        }
                                    case "":
                                        switch ("foo bar" + "") {
                                        case "hello":
                                            v1 += (^(v1))
                                        case "foo bar":
                                            v1 = v1
                                        case "a":
                                            if false {
                                                var v91 float32 = (0.1 + float32((1e10 + -2.25)))
                                                _ = v91
                                            }
                                        default:
                                            if true {
                                                v1 = 42
                                                v1 = (v1 + (^(v1)))
                                                var v92 uint16 = 1
                                                _ = v92
                                            } else {
                                                var v93 uint8 = (v1 ^ (v1 << ((2) & 63)))
                                                v94 := GenS5{F0: 65535, F1: -728}
                                                var v95 float64 = 1e10
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                            }
                                        }
                                    default:
                                        v96 := GenS2{F0: 4294967295, F1: ("" + (("a" + "a") + ("foo bar" + ""))), F2: "foo bar"}
                                        _ = v96
                                    }
                                    var v97 uint = 4294967295
                                    _ = v97
                                }
                                _ = fi68
                            }
                        }
                    case -2147483648:
                        var v98 int64 = 1
                        _ = v98
                    }
                    if (0 == -128) {
                        for fi99 := 0; fi99 < 7; fi99++ {
                            v58 = (v58 | (v58 + v58))
                            var v100 string = ("" + "a")
                            fmt.Println(v100)
                            fmt.Println(v58)
                            _ = fi99
                            _ = v100
                        }
                    } else {
                        switch ("\n\t" + ("foo bar" + "a")) {
                        case "a":
                            v58 += 874
                        }
                    }
                    _ = fi66
                }
                _ = v65
            }
        }
        _ = v58
    }
    fmt.Println(v1)
    for fi101 := 0; fi101 < 13; fi101++ {
        if false {
            var v102 int = fi101
            _ = v102
        }
        v1 = (v1 ^ v1)
        fmt.Println(v1, fi101)
        _ = fi101
    }
    if false {
        for fi103 := 0; fi103 < 3; fi103++ {
            for fi104 := 0; fi104 < 16; fi104++ {
                v1 -= (v1 << ((347) & 63))
                _ = len(v5)
                _ = fi104
            }
            delete(v5, 2147483647)
            _ = fi103
        }
        fmt.Println(v1)
    } else {
        for fi105 := 0; fi105 < 7; fi105++ {
            for fi106 := 0; fi106 < 5; fi106++ {
                switch "" {
                case "hello":
                    v3[42] = 2
                case "a":
                    fmt.Println(fi105)
                case "foo bar":
                    fmt.Println(fi106, v1)
                }
                _ = fi106
            }
            v1 = v1
            v1 = v1
            _ = fi105
        }
        fmt.Println(v1)
        var v107 int8 = 0
        _ = v107
    }
    if false {
        v4 = append(v4, 1)
        v1 = v1
    }
    v1 = (v1 + v1)
    switch 9223372036854775807 {
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 2147483647:
        if false {
            var v108 uint8 = (v1 >> ((2) & 63))
            v109 := GenS6{F0: -1}
            _ = v108
            _ = v109
        }
    default:
        var v110 uint16 = 42
        _ = v110
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
