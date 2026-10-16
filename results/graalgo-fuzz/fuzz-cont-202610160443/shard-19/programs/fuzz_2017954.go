package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 int }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 bool; F1 bool }
type GenS15 struct { F0 int64; F1 string; F2 uint8 }
type GenS16 struct { F0 int16; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 816
    var v2 bool = ((-72 >= -9223372036854775808) && (!false))
    v3 := make(map[string]string)
    var v4 int64 = 553
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    switch "foo bar" {
    case "\n\t":
        v7 := GenS1{F0: 0}
        _ = v7
    default:
        var v8 int64 = 468
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            switch (fi9 + (-(fi9))) {
            case 128:
                var v10 bool = (255 < 65535)
                _ = v10
            case 32767:
                fmt.Println(fi9, fi9)
            case 2147483647:
                if (v2 && ((v1 | v1) >= (v1 & (v1 + 0)))) {
                    v11 := []int{fi9, 32768}
                    _ = v11
                } else {
                    v2 = v2
                    if v2 {
                        if v2 {
                            var v12 string = "\n\t"
                            var v13 uint = 154
                            for fi14 := 0; fi14 < 16; fi14++ {
                                delete(v5, ((((v12 + ((("hello" + "hello") + "a") + v12)) + v12) + (((("\n\t" + v12) + "hello") + v12) + v12)) + v12))
                                v15 := GenS2{F0: int16(v1), F1: -2147483648}
                                _ = fi14
                                _ = v15
                            }
                            _ = v12
                            _ = v13
                        } else {
                            var v16 uint16 = 2
                            v17 := GenS3{F0: v1, F1: true, F2: fi9}
                            _ = v16
                            _ = v17
                        }
                        var v18 float32 = (0.1 / float32(1e10))
                        _ = v18
                    } else {
                        if (306 > 65535) {
                            v19 := []int{fi9, (fi9 << ((4294967295) & 63)), 4294967295, (fi9 ^ (fi9 + (fi9 & 68)))}
                            if v2 {
                                if ((((-1 >= 0) && (3.14159 <= 1e10)) && true) || (true && v2)) {
                                    if false {
                                        if true {
                                            for fi20 := 0; fi20 < 20; fi20++ {
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                v21 := GenS2{F0: -424, F1: -555}
                                                v4 = int64(fi9)
                                                _ = fi20
                                                _ = v21
                                            }
                                            fmt.Println(v4)
                                            v22 := GenS2{F0: 32767, F1: 2147483647}
                                            _ = v22
                                        }
                                    }
                                }
                            }
                            v23 := GenS4{F0: (v4 / (((v4 / ((0) | 1))) | 1)), F1: 255, F2: 2147483647}
                            _ = v19
                            _ = v23
                        }
                        v2 = (v2 || (v4 != v4))
                    }
                }
            default:
                if false {
                    var v24 int16 = -32768
                    if v2 {
                        var v25 int16 = -631
                        _ = v25
                    }
                    _ = v24
                }
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: int32(v4)}
                if v2 {
                    switch "hello" {
                    case "hello":
                        v2 = (v1 < (v1 / ((v1) | 1)))
                    case "foo bar":
                        if (v2 && (-1 != -1)) {
                            var v28 uint = 18446744073709551615
                            v29 := GenS5{F0: "a"}
                            switch ("hello" + "a") {
                            case "foo bar":
                                v30 := GenS6{F0: 1, F1: 2147483647}
                                _ = v30
                            case "hello":
                                v31 := GenS1{F0: 365}
                                _ = v31
                            case "":
                                if (true || (!false)) {
                                    if v2 {
                                        v32 := GenS3{F0: v1, F1: true, F2: (fi9 & (fi9 - (fi26 * (fi9 >> ((uint(v1)) & 63)))))}
                                        _ = v32
                                    }
                                    if true {
                                        switch "a" {
                                        case "\n\t":
                                            v1 = (v1 / ((306) | 1))
                                        case "foo bar":
                                            var v33 int32 = 393
                                            _ = v33
                                        case "hello":
                                            v28 += (^(v28))
                                        default:
                                            if false {
                                                v1 = (v1 + v1)
                                                var v34 int8 = 127
                                                _ = v34
                                            } else {
                                                v35 := GenS7{F0: -128, F1: 982}
                                                _ = v35
                                            }
                                        }
                                        v36 := GenS8{F0: "a", F1: (v4 << (((v28 / ((65535) | 1))) & 63))}
                                        _ = v36
                                    }
                                }
                            default:
                                v37 := GenS9{F0: 42}
                                _ = v37
                            }
                            v38 := []int{-2147483648}
                            v4 = v4
                            _ = v28
                            _ = v29
                            _ = v38
                        } else {
                            var v39 uint16 = 556
                            fmt.Println(v1, v39, fi9)
                            v39 = (v39 + v39)
                            _ = v39
                        }
                    }
                    fmt.Println(v2, fi26, fi9)
                } else {
                    var v40 uint = 1
                    v41 := GenS6{F0: 32767, F1: (fi26 & fi9)}
                    v1 = v1
                    var v42 int32 = -854
                    switch "" {
                    case "hello":
                        _ = len(v5)
                    }
                    _ = v40
                    _ = v41
                    _ = v42
                }
                _ = fi26
                _ = v27
            }
            v2 = (v2 || v2)
        }
        fmt.Println(v2, v1)
        v1 -= v1
        _ = fi9
    }
    if (v2 || false) {
        v43 := GenS10{F0: 0, F1: -1}
        if v2 {
            var v44 float32 = 1e10
            var v45 uint8 = 18
            _ = v44
            _ = v45
        } else {
            for fi46 := 0; fi46 < 2; fi46++ {
                if (1 != 0) {
                    v47 := []int{}
                    delete(v3, "")
                    v1 -= 1
                    _ = v47
                }
                fmt.Println(v4, fi46, fi46)
                if v2 {
                    switch (fi46 % ((fi46) | 1)) {
                    case 1:
                        v48 := GenS11{F0: v2, F1: 0}
                        _ = v48
                    default:
                        if (898 == 65535) {
                            switch ("hello" + "a") {
                            case "\n\t":
                                var v49 bool = (55 == 3)
                                _ = v49
                            }
                            var v50 int = (fi46 / (((^(fi46))) | 1))
                            switch 128 {
                            case 4294967295:
                                fmt.Println(v50, v1, v50)
                            case 2147483647:
                                if true {
                                    for fi51 := 0; fi51 < 1; fi51++ {
                                        v52 := GenS3{F0: v1, F1: (!true), F2: (fi46 % ((int(v1)) | 1))}
                                        var v53 int16 = 992
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                    switch (^(v50)) {
                                    case -268:
                                        v2 = true
                                    default:
                                        for fi54 := 0; fi54 < 11; fi54++ {
                                            v1 -= 42
                                            fmt.Println(v4, v4, fi46)
                                            v55 := GenS5{F0: "\n\t"}
                                            var v56 int = (fi54 * (v50 / ((128) | 1)))
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                    }
                                } else {
                                    v1 = v1
                                    if v2 {
                                        for fi57 := 0; fi57 < 9; fi57++ {
                                            if (920 == 1) {
                                                v2 = true
                                                var v58 string = "foo bar"
                                                v59 := v3[(v58 + v58)]
                                                v60 := GenS12{F0: 1}
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            } else {
                                                switch (fi46 + -1) {
                                                case 32768:
                                                    if (("a" + ("\n\t" + ("hello" + "a"))) <= "foo bar") {
                                                        var v61 bool = false
                                                        _ = v61
                                                    }
                                                case 228:
                                                    v1 += 1
                                                case 4294967295:
                                                    v1 = 18446744073709551615
                                                default:
                                                    v62 := []int{(fi46 << ((1) & 63)), (fi46 >> ((0) & 63))}
                                                    _ = v62
                                                }
                                                var v63 float32 = 0.1
                                                _ = v63
                                            }
                                            var v64 int = -620
                                            v65 := v6[2]
                                            _ = fi57
                                            _ = v64
                                            _ = v65
                                        }
                                        var v66 int8 = 42
                                        _ = v66
                                    } else {
                                        v2 = ((2147483647 >= 42) || v2)
                                    }
                                }
                            default:
                                _ = len(v5)
                            }
                            _ = v50
                        } else {
                            fmt.Println(v4, v1, v1)
                        }
                    }
                    v67 := []int{fi46}
                    _ = v67
                } else {
                    v2 = (((float32(3.14159) / (1e10 / 1.0)) * (-2.25 - -2.25)) > (-2.25 * 1.5))
                    var v68 int32 = -882
                    _ = v68
                }
                _ = fi46
            }
            if v2 {
                var v69 float32 = 0.0
                var v70 int16 = -1
                _ = v69
                _ = v70
            }
        }
        _ = len(v5)
        _ = v43
    } else {
        switch 102 {
        case 42:
            v71 := GenS13{F0: -9223372036854775808, F1: 128}
            _ = v71
        case 32768:
            if v2 {
                fmt.Println(v2, v1)
            }
        default:
            var v72 uint = 305
            _ = v72
        }
    }
    switch "foo bar" {
    case "":
        v73 := GenS14{F0: v2, F1: (!false)}
        _ = v73
    }
    v2 = v2
    for fi74 := 0; fi74 < 9; fi74++ {
        for fi75 := 0; fi75 < 7; fi75++ {
            var v76 uint = 894
            _ = fi75
            _ = v76
        }
        var v77 float64 = -2.25
        if (70 == -3) {
            var v78 int = (-(fi74))
            switch (v78 | v78) {
            case 127:
                var v79 uint = 60
                _ = v79
            case -332:
                var v80 uint8 = 0
                _ = v80
            case 940:
                v77 = v77
            default:
                switch fi74 {
                case 42:
                    fmt.Println(v78, v4, fi74)
                case -463:
                    var v81 uint8 = 42
                    _ = v81
                case 551:
                    switch "foo bar" {
                    case "":
                        if (((fi74 + v78) <= (v78 % ((-2147483648) | 1))) || ((v2 && v2) || true)) {
                            v82 := GenS10{F0: 994, F1: -2147483648}
                            v77 = ((v77 * v77) / v77)
                            _ = v82
                        } else {
                            v1 = v1
                        }
                    case "foo bar":
                        v83 := []int{v78, (v78 << ((899) & 63)), (fi74 + 4294967295), (fi74 >> ((1) & 63))}
                        _ = v83
                    case "\n\t":
                        if v2 {
                            var v84 uint64 = 975
                            _ = v84
                        }
                    default:
                        v78 = fi74
                    }
                }
            }
            var v85 float64 = ((((v77 * ((v77 + float64(1.5)) - (((0.1 - -2.25) - v77) / float64(float32(-2.25))))) * v77) + (v77 / (v77 - v77))) - (1.5 * float64(1.5)))
            _ = v78
            _ = v85
        } else {
            v3["foo bar"] = "a"
            for fi86 := 0; fi86 < 19; fi86++ {
                v87 := []int{-2147483648, (fi86 & (fi74 | fi86))}
                var v88 float32 = 1.5
                _ = len(v6)
                _ = fi86
                _ = v87
                _ = v88
            }
        }
        _ = fi74
        _ = v77
    }
    var v89 float64 = float64(3.14159)
    v1 = (v1 & 0)
    v90 := GenS15{F0: v4, F1: (("a" + "") + "\n\t"), F2: 2}
    if (-92 > 42) {
        var v91 int64 = 388
        for fi92 := 0; fi92 < 19; fi92++ {
            if (true && ((fi92 | fi92) >= fi92)) {
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v6)
                    case "":
                        if (v2 && v2) {
                            v89 += (float64(((((1.5 / 1.5) / 1.0) - 1.5) - 1.5)) + ((v89 + v89) * (v89 + float64(0.1))))
                        }
                    case "hello":
                        v93 := []int{(fi92 & -2147483648), (fi92 * fi92)}
                        _ = v93
                    default:
                        for fi94 := 0; fi94 < 19; fi94++ {
                            v89 -= -2.25
                            if (100 > -128) {
                                _ = len(v6)
                                switch "a" {
                                case "hello":
                                    var v95 float32 = (float32(v89) * 0.0)
                                    _ = v95
                                case "foo bar":
                                    var v96 bool = (v2 && v2)
                                    _ = v96
                                }
                            } else {
                                if false {
                                    var v97 uint = 255
                                    v89 -= (1.5 * v89)
                                    _ = v97
                                } else {
                                    if v2 {
                                        v91 -= 0
                                        v98 := []int{939, (fi94 % (((fi92 * 1)) | 1))}
                                        if v2 {
                                            var v99 int = (^(fi94))
                                            _ = v99
                                        } else {
                                            fmt.Println(v89, v91)
                                        }
                                        _ = v98
                                    }
                                    delete(v3, "hello")
                                    switch -953 {
                                    case 32767:
                                        v100 := v6[0]
                                        _ = v100
                                    case 0:
                                        v101 := []int{128, 127, fi92}
                                        _ = v101
                                    case 1:
                                        var v102 string = ((((((("\n\t" + (("\n\t" + "\n\t") + "hello")) + "a") + "\n\t") + "foo bar") + "") + (("foo bar" + (("\n\t" + "hello") + "hello")) + "foo bar")) + "\n\t")
                                        _ = v102
                                    default:
                                        fmt.Println(v89, v1, v91)
                                    }
                                }
                                var v103 int8 = -1
                                v104 := GenS7{F0: (^(v103)), F1: 65535}
                                _ = v103
                                _ = v104
                            }
                            if (v2 || v2) {
                                if (true || (!(623 <= 255))) {
                                    var v105 uint16 = 1
                                    fmt.Println(v91, v2)
                                    var v106 int = fi94
                                    _ = v105
                                    _ = v106
                                }
                            } else {
                                var v107 int16 = 32767
                                if v2 {
                                    v89 = 1.5
                                }
                                _ = v107
                            }
                            _ = fi94
                        }
                    }
                    switch "hello" {
                    case "":
                        switch 9223372036854775807 {
                        case 1:
                            var v108 uint32 = 0
                            _ = v108
                        case -9223372036854775808:
                            v91 -= v91
                        }
                    }
                    delete(v3, "hello")
                }
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v89, fi92, v4)
                case "hello":
                    fmt.Println(v89, v2, v2)
                default:
                    var v109 int32 = -158
                    _ = v109
                }
            }
            v110 := GenS16{F0: 11, F1: v1, F2: 18446744073709551615}
            _ = fi92
            _ = v110
        }
        _ = v91
    }
    for fi111 := 0; fi111 < 19; fi111++ {
        var v112 uint32 = 1
        v89 -= v89
        v1 -= v1
        _ = fi111
        _ = v112
    }
    for fi113 := 0; fi113 < 4; fi113++ {
        v2 = true
        var v114 uint64 = v1
        _ = fi113
        _ = v114
    }
    delete(v3, "\n\t")
    v4 = (v4 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
}
