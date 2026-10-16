package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int64; F1 uint32; F2 int8 }
type GenS8 struct { F0 int16; F1 string; F2 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint, p2 uint32) int16 {
    p0 = (!p0)
    var v7 int32 = 0
    var v8 uint64 = 262
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1() int8 {
    switch "foo bar" {
    case "\n\t":
        v9 := GenS1{F0: -1, F1: -974, F2: -1}
        _ = v9
    case "foo bar":
        v10 := []int{-15, -1}
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint8 = 1
            _ = fi11
            _ = v12
        }
    }
    var v13 int16 = genFn0((940 <= 4294967295), 18446744073709551615, 2147483648)
    for fi14 := 0; fi14 < 19; fi14++ {
        if true {
            var v15 bool = (1 > 586)
            v16 := GenS2{F0: 1, F1: 42, F2: 0}
            _ = v15
            _ = v16
        } else {
            switch (fi14 << ((703) & 63)) {
            case 42:
                fmt.Println(v13, fi14)
            case 0:
                if false {
                    v17 := GenS1{F0: 100, F1: -290, F2: 52}
                    if false {
                        if false {
                            v18 := GenS3{F0: -32768, F1: -9223372036854775808}
                            for fi19 := 0; fi19 < 12; fi19++ {
                                var v20 float64 = (float64(-2.25) + 0.0)
                                if (481 >= 0) {
                                    v13 = 32767
                                }
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        } else {
                            v13 = (-(v13))
                            v13 = (v13 ^ (v13 & (v13 & 1)))
                            for fi21 := 0; fi21 < 9; fi21++ {
                                v22 := GenS4{F0: 673, F1: (false && (("hello" + "hello") != ""))}
                                fmt.Println(v13)
                                var v23 int = 32768
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                    v13 += v13
                    _ = v17
                } else {
                    if false {
                        var v24 uint16 = 1
                        if (2 == 18446744073709551615) {
                            if true {
                                v24 = v24
                                var v25 float32 = ((-2.25 / (float32(1e10) - 1.5)) - ((1.5 - 1e10) * (1e10 * ((0.1 * (0.1 * (0.1 + (float32(1.5) - 3.14159)))) + (0.1 * 0.1)))))
                                fmt.Println(fi14, v25)
                                _ = v25
                            } else {
                                fmt.Println(v24, v24, fi14)
                            }
                        } else {
                            v26 := GenS3{F0: 1000, F1: (fi14 << ((1) & 63))}
                            v24 = v24
                            if false {
                                var v27 int32 = -1
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    fmt.Println(v24, fi28)
                                    v24 -= (v24 % (((v24 << ((4294967295) & 63))) | 1))
                                    switch 4294967295 {
                                    case -1:
                                        var v29 uint = 0
                                        _ = v29
                                    case 541:
                                        var v30 string = ""
                                        _ = v30
                                    case 128:
                                        var v31 uint16 = 39
                                        _ = v31
                                    default:
                                        v24 -= 32767
                                    }
                                    _ = fi28
                                }
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    var v33 float64 = (3.14159 / (3.14159 * 1e10))
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v27
                            }
                            _ = v26
                        }
                        switch ("foo bar" + ("hello" + "hello")) {
                        case "":
                            if false {
                                if false {
                                    switch "hello" {
                                    case "":
                                        v24 += 0
                                    case "hello":
                                        v34 := []int{2147483647, int(v24), (fi14 << ((529) & 63)), 127}
                                        _ = v34
                                    case "a":
                                        switch "hello" {
                                        case "":
                                            var v35 int64 = -1
                                            _ = v35
                                        case "foo bar":
                                            for fi36 := 0; fi36 < 13; fi36++ {
                                                var v37 string = "\n\t"
                                                switch 0 {
                                                case 32767:
                                                    if (v37 < "") {
                                                        var v38 bool = (-103 < -128)
                                                        switch v37 {
                                                        case "a":
                                                            v39 := GenS4{F0: 2147483647, F1: (!false)}
                                                            _ = v39
                                                        case "\n\t":
                                                            var v40 int32 = -1
                                                            _ = v40
                                                        case "foo bar":
                                                            fmt.Println(fi14, v37, v37)
                                                        default:
                                                            v41 := GenS4{F0: 951, F1: (869 != 973)}
                                                            _ = v41
                                                        }
                                                        _ = v38
                                                    } else {
                                                        var v42 uint32 = 65535
                                                        var v43 uint = 65535
                                                        switch (fi14 & (fi36 * (^(fi14)))) {
                                                        case 4294967295:
                                                            fmt.Println(v42)
                                                        case -1:
                                                            if ((v43 ^ v43) < (v43 << (((v43 >> ((42) & 63))) & 63))) {
                                                                var v44 int = fi36
                                                                fmt.Println(v42, fi36, v43)
                                                                var v45 bool = false
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                fmt.Println(fi14, v42, v42)
                                                                var v46 float32 = -2.25
                                                                _ = v46
                                                            }
                                                        case -178:
                                                            var v47 uint8 = 0
                                                            var v48 uint16 = (v24 % ((v24) | 1))
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    v24 += (v24 + (v24 % ((2) | 1)))
                                                }
                                                if false {
                                                    var v49 int16 = (v13 >> ((678) & 63))
                                                    var v50 int8 = 1
                                                    var v51 float32 = -2.25
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = fi36
                                                _ = v37
                                            }
                                        default:
                                            v52 := []int{}
                                            _ = v52
                                        }
                                    }
                                } else {
                                    fmt.Println(fi14, v13, fi14)
                                    v53 := GenS5{F0: 0}
                                    var v54 uint = 646
                                    _ = v53
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    fmt.Println(v13, v24)
                                    _ = fi55
                                }
                                if false {
                                    for fi56 := 0; fi56 < 12; fi56++ {
                                        if (!(-112 == 100)) {
                                            fmt.Println(fi14, v24, fi14)
                                            switch "foo bar" {
                                            case "":
                                                if (1 >= 18446744073709551615) {
                                                    var v57 uint = 18446744073709551615
                                                    _ = v57
                                                }
                                            }
                                            v58 := []int{fi56, (fi56 % (((-(fi14))) | 1)), fi56}
                                            _ = v58
                                        } else {
                                            var v59 string = ("hello" + ("" + (("a" + ("a" + "foo bar")) + "hello")))
                                            _ = v59
                                        }
                                        v13 += (-(v13))
                                        _ = fi56
                                    }
                                }
                            } else {
                                var v60 int32 = 0
                                for fi61 := 0; fi61 < 11; fi61++ {
                                    v62 := []int{(fi14 / ((42) | 1)), fi14, fi61}
                                    v13 += v13
                                    _ = fi61
                                    _ = v62
                                }
                                if true {
                                    if ((!false) || true) {
                                        var v63 int64 = 42
                                        _ = v63
                                    }
                                    v13 = (v13 - (v13 % ((v13) | 1)))
                                    switch "" {
                                    case "\n\t":
                                        for fi64 := 0; fi64 < 6; fi64++ {
                                            v65 := []int{int(v13)}
                                            fmt.Println(fi14)
                                            _ = fi64
                                            _ = v65
                                        }
                                    case "foo bar":
                                        var v66 float32 = 1e10
                                        _ = v66
                                    case "":
                                        for fi67 := 0; fi67 < 13; fi67++ {
                                            var v68 uint = 977
                                            v24 = (v24 * (v24 | 360))
                                            _ = fi67
                                            _ = v68
                                        }
                                    default:
                                        v69 := GenS6{F0: 950, F1: v24}
                                        _ = v69
                                    }
                                } else {
                                    for fi70 := 0; fi70 < 5; fi70++ {
                                        v60 += (v60 & (v60 - (v60 - 929)))
                                        var v71 float32 = (float32(1e10) * ((0.0 + 0.1) / 1e10))
                                        if (2 != 1) {
                                            v71 = -2.25
                                        }
                                        var v72 int = (^(fi70))
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                    v24 -= v24
                                    var v73 uint8 = 0
                                    _ = v73
                                }
                                _ = v60
                            }
                        default:
                            v74 := []int{4294967295}
                            _ = v74
                        }
                        _ = v24
                    } else {
                        v75 := []int{(fi14 ^ fi14), fi14, 127}
                        var v76 int32 = -463
                        _ = v75
                        _ = v76
                    }
                    if (true || (1.5 >= float64((0.0 * 1e10)))) {
                        for fi77 := 0; fi77 < 17; fi77++ {
                            v78 := GenS7{F0: 42, F1: 972, F2: -128}
                            if true {
                                if false {
                                    v13 += (v13 << ((0) & 63))
                                    fmt.Println(fi77, fi77)
                                } else {
                                    var v79 uint16 = 42
                                    v80 := GenS8{F0: v13, F1: ((("a" + "hello") + "\n\t") + ("a" + ("hello" + "foo bar"))), F2: 2}
                                    _ = v79
                                    _ = v80
                                }
                                fmt.Println(v13, v13)
                                var v81 string = "foo bar"
                                _ = v81
                            } else {
                                if (false || false) {
                                    v13 = v13
                                }
                                v13 = (v13 + (v13 | (v13 ^ (v13 * -359))))
                            }
                            v13 = genFn0(true, 1, 4294967295)
                            v13 = (^(v13))
                            switch "" {
                            case "a":
                                var v82 uint32 = 499
                                _ = v82
                            }
                            if (!true) {
                                if (true && ((-80 >= fi77) || ((false || (1.5 <= 1.5)) && (false || (true || true))))) {
                                    switch "a" {
                                    case "":
                                        v13 = (v13 | (v13 - (-(v13))))
                                    default:
                                        v83 := GenS2{F0: 1, F1: 42, F2: 229}
                                        _ = v83
                                    }
                                    var v84 uint = 281
                                    _ = v84
                                } else {
                                    switch (fi14 << ((794) & 63)) {
                                    case 390:
                                        v85 := []int{fi14, 2147483647, fi77}
                                        _ = v85
                                    case 4294967295:
                                        v13 = (v13 % ((-1) | 1))
                                    case -55:
                                        v13 = v13
                                    default:
                                        v86 := GenS9{F0: -754}
                                        _ = v86
                                    }
                                }
                                fmt.Println(fi77, v13, v13)
                                v13 += v13
                            } else {
                                fmt.Println(v13, fi14, fi14)
                                v13 += 1000
                            }
                            _ = fi77
                            _ = v78
                        }
                        v13 = (v13 & v13)
                        fmt.Println(fi14, v13)
                    }
                    var v87 int16 = (-(v13))
                    _ = v87
                }
            default:
                switch "foo bar" {
                case "a":
                    v13 -= (v13 - (v13 >> ((1) & 63)))
                case "":
                    var v88 float32 = 0.0
                    _ = v88
                case "hello":
                    if (1 >= 42) {
                        v13 = v13
                        for fi89 := 0; fi89 < 13; fi89++ {
                            var v90 int8 = 108
                            if false {
                                var v91 string = "\n\t"
                                v92 := GenS5{F0: 527}
                                _ = v91
                                _ = v92
                            }
                            if (false || true) {
                                fmt.Println(fi14, v13, fi14)
                                if false {
                                    var v93 int64 = 42
                                    v94 := GenS7{F0: (v93 - 9223372036854775807), F1: 113, F2: 1}
                                    v95 := GenS1{F0: 1, F1: (v13 << ((1) & 63)), F2: v90}
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                }
                                v96 := GenS10{F0: 700}
                                _ = v96
                            }
                            _ = fi89
                            _ = v90
                        }
                        v97 := []int{fi14, fi14, (fi14 % ((128) | 1))}
                        _ = v97
                    }
                }
            }
        }
        _ = fi14
    }
    v13 = 42
    _ = v13
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 string = "\n\t"
    var v3 uint32 = 2147483648
    var v4 uint8 = 1
    var v5 int8 = -128
    v6 := make(map[int]int)
    v98 := []int{32768, 32768, -358}
    if (false || true) {
        v5 = (v5 >> ((973) & 63))
        v1 = v3
        if true {
            var v99 float32 = 1e10
            for fi100 := 0; fi100 < 20; fi100++ {
                v5 += genFn1()
                var v101 int8 = (v5 % ((126) | 1))
                switch "foo bar" {
                case "foo bar":
                    if len(v98) > 0 {
                        v98[0] = fi100
                    }
                }
                _ = fi100
                _ = v101
            }
            _ = v99
        } else {
            v102 := v6[0]
            var v103 uint64 = 0
            _ = v102
            _ = v103
        }
    } else {
        v104 := GenS1{F0: genFn1(), F1: 1, F2: (v5 + v5)}
        if len(v98) >= 2 {
            v98 = v98[1:len(v98)]
        }
        _ = v104
    }
    v1 -= (^(v1))
    if true {
        for fi105 := 0; fi105 < 10; fi105++ {
            fmt.Println(v4, fi105, v4)
            v2 = ""
            _ = len(v98)
            _ = fi105
        }
        switch (v2 + "a") {
        case "a":
            for fi106 := 0; fi106 < 5; fi106++ {
                if (!((42 != 42) && false)) {
                    for fi107 := 0; fi107 < 5; fi107++ {
                        fmt.Println(v5)
                        _ = fi107
                    }
                    v3 += v3
                    v4 = 128
                }
                _ = fi106
            }
        }
        if len(v98) > 0 {
            v98[0] = 2147483647
        }
    }
    if (v2 >= ("\n\t" + v2)) {
        switch 32768 {
        case 42:
            v108 := GenS5{F0: 28}
            _ = v108
        default:
            if (965 > int(v5)) {
                v109 := v6[4294967295]
                _ = v109
            }
        }
        v3 = uint32(v4)
        if true {
            fmt.Println(v5, v5, v2)
            var v110 string = v2
            _ = v110
        } else {
            switch (v2 + v2) {
            case "hello":
                var v111 int64 = 42
                _ = v111
            case "":
                v5 = (-(v5))
            }
            for fi112 := 0; fi112 < 15; fi112++ {
                _ = len(v6)
                if false {
                    if (!((593 == 32767) && false)) {
                        _ = len(v98)
                    } else {
                        if len(v98) >= 2 {
                            v98 = v98[1:len(v98)]
                        }
                    }
                    var v113 string = "hello"
                    if true {
                        var v114 uint = 2
                        for fi115 := 0; fi115 < 1; fi115++ {
                            fmt.Println(v3, fi112, v114)
                            var v116 int16 = genFn0(false, (^(v114)), (v1 ^ 2147483647))
                            _ = fi115
                            _ = v116
                        }
                        _ = v114
                    } else {
                        for fi117 := 0; fi117 < 10; fi117++ {
                            var v118 float32 = 0.1
                            _ = fi117
                            _ = v118
                        }
                        v3 += 4294967295
                    }
                    _ = v113
                }
                _ = fi112
            }
        }
    } else {
        v98 = append(v98, -9223372036854775808)
        for fi119 := 0; fi119 < 15; fi119++ {
            for fi120 := 0; fi120 < 13; fi120++ {
                v2 = v2
                _ = fi120
            }
            if true {
                if (!((false && (uint32(v4) < v1)) && (0 >= v5))) {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v5, fi119)
                    }
                    var v121 int64 = 504
                    _ = v121
                }
                v2 = ("hello" + v2)
                v5 -= 100
            } else {
                for fi122 := 0; fi122 < 20; fi122++ {
                    v3 -= v3
                    v3 = 771
                    _ = fi122
                }
            }
            _ = fi119
        }
    }
    fmt.Println(v1, v4)
    fmt.Println(v5, v1, v5)
    var v123 uint16 = 649
    var v124 uint64 = 42
    for fi125 := 0; fi125 < 16; fi125++ {
        var v126 uint32 = v1
        _ = fi125
        _ = v126
    }
    var v127 bool = (v4 <= v4)
    v5 = genFn1()
    for fi128 := 0; fi128 < 10; fi128++ {
        v127 = (!v127)
        _ = fi128
    }
    var v129 uint32 = (v3 * v3)
    v130 := GenS9{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v127:", v127)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v123
    _ = v124
    _ = v127
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
