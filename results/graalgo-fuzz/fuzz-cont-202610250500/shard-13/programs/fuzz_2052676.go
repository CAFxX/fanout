package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }
type GenS4 struct { F0 bool; F1 int; F2 bool }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint8; F1 bool; F2 int16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int; F2 uint64 }
type GenS10 struct { F0 int32; F1 int; F2 string }
type GenS11 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0(p0 uint64) uint16 {
    if (-51 >= 32768) {
        for fi4 := 0; fi4 < 11; fi4++ {
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 uint16 = 1
                v7 := GenS1{F0: "", F1: 127}
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
        fmt.Println(p0)
    } else {
        v8 := GenS1{F0: "\n\t", F1: -1}
        var v9 uint8 = 1
        if false {
            for fi10 := 0; fi10 < 1; fi10++ {
                p0 = 1
                v9 = v9
                fmt.Println(fi10, v9, v9)
                _ = fi10
            }
            v9 = v9
            if ((-515 < 9223372036854775807) && (!(235 != 255))) {
                v11 := []int{-2147483648, 32767, 706}
                v9 = 127
                _ = v11
            }
        }
        _ = v8
        _ = v9
    }
    if (111 <= 804) {
        for fi12 := 0; fi12 < 17; fi12++ {
            p0 = (p0 << ((385) & 63))
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int32 = 42
                _ = fi13
                _ = v14
            }
            p0 = p0
            _ = fi12
        }
    } else {
        v15 := GenS2{F0: 2147483647}
        if true {
            var v16 uint64 = (p0 * (p0 << ((4294967295) & 63)))
            _ = v16
        } else {
            v17 := GenS3{F0: 969, F1: 107, F2: -1}
            _ = v17
        }
        _ = v15
    }
    v18 := []int{0, 896, -860, 0}
    fmt.Println(p0)
    _ = p0
    _ = v18
    return uint16(p0)
}
func genFn1(p0 uint64, p1 int) uint8 {
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int8 = -1
        v21 := GenS1{F0: ("foo bar" + "\n\t"), F1: v20}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint32 = 65535
    if true {
        if false {
            var v23 uint = 42
            _ = v23
        }
    } else {
        var v24 int64 = -1
        v25 := []int{}
        _ = v24
        _ = v25
    }
    fmt.Println(v22, v22, p1)
    _ = p0
    _ = p1
    _ = v22
    return 2
}

func main() {
    var v1 uint = 1
    v2 := []uint32{2147483647}
    var v3 uint = 1
    fmt.Println(v1, v1)
    switch 2147483647 {
    case -869:
        var v26 uint8 = 0
        _ = v26
    }
    if false {
        var v27 uint16 = 65535
        _ = v27
    }
    v3 = v1
    v1 += 255
    fmt.Println(v3, v1)
    fmt.Println(v3)
    if (0 > genFn0(1)) {
        switch ("" + "\n\t") {
        case "foo bar":
            var v28 uint32 = 278
            _ = v28
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            if false {
                _ = len(v2)
                if (3.14159 <= (0.0 - 1e10)) {
                    if (2 == 0) {
                        var v30 string = ""
                        v1 = (v3 >> ((255) & 63))
                        for fi31 := 0; fi31 < 16; fi31++ {
                            fmt.Println(v3)
                            v32 := GenS3{F0: (v3 << ((4294967295) & 63)), F1: 100, F2: -9223372036854775808}
                            if true {
                                if true {
                                    var v33 float32 = (0.0 * (((1.5 * 3.14159) / 3.14159) + 1.5))
                                    _ = v33
                                }
                            }
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    var v34 int16 = 1000
                    v3 = v1
                    _ = v34
                }
                var v35 int = (fi29 << ((v3) & 63))
                _ = v35
            } else {
                for fi36 := 0; fi36 < 8; fi36++ {
                    if false {
                        var v37 uint32 = 76
                        _ = v37
                    } else {
                        fmt.Println(fi36, fi29)
                        if false {
                            v1 = 4294967295
                        }
                        if (fi36 == (fi29 ^ fi29)) {
                            if len(v2) > 0 {
                                v2[0] = 211
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            if (((false && true) || ((v3 ^ v3) >= (v1 ^ (v3 + 18446744073709551615)))) && (true || true)) {
                                fmt.Println(v1, fi29)
                                v38 := GenS3{F0: v3, F1: 100, F2: int64(fi29)}
                                _ = v38
                            } else {
                                var v39 int = (fi29 - fi36)
                                _ = v39
                            }
                        }
                    }
                    _ = fi36
                }
                v1 += v3
                v3 -= v1
            }
            v40 := GenS3{F0: 255, F1: 95, F2: 9223372036854775807}
            _ = fi29
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 20; fi41++ {
            for fi42 := 0; fi42 < 10; fi42++ {
                fmt.Println(fi42)
                if (9223372036854775807 < 9223372036854775807) {
                    switch "" {
                    case "hello":
                        var v43 int8 = -128
                        _ = v43
                    case "":
                        v3 = 42
                    }
                    var v44 uint64 = 2
                    v45 := GenS2{F0: 114}
                    _ = v44
                    _ = v45
                } else {
                    _ = len(v2)
                    fmt.Println(v3)
                    v46 := GenS3{F0: 4294967295, F1: -1, F2: 9223372036854775807}
                    _ = v46
                }
                _ = fi42
            }
            v1 = v1
            _ = fi41
        }
        var v47 uint8 = 128
        v48 := GenS4{F0: true, F1: 9223372036854775807, F2: false}
        _ = v47
        _ = v48
    }
    var v49 bool = (1e10 > (0.1 / 3.14159))
    v50 := GenS1{F0: "", F1: -128}
    v49 = v49
    fmt.Println(v49, v1)
    v51 := GenS5{F0: 1, F1: 4294967295, F2: v3}
    switch "" {
    case "":
        if len(v2) > 0 {
            v2[0] = 0
        }
    case "\n\t":
        if v49 {
            switch "" {
            case "foo bar":
                var v52 int32 = -1
                _ = v52
            case "hello":
                var v53 uint = v3
                _ = v53
            case "\n\t":
                for fi54 := 0; fi54 < 8; fi54++ {
                    var v55 int = fi54
                    for fi56 := 0; fi56 < 11; fi56++ {
                        v1 = v1
                        var v57 uint8 = 1
                        _ = fi56
                        _ = v57
                    }
                    _ = fi54
                    _ = v55
                }
            default:
                for fi58 := 0; fi58 < 6; fi58++ {
                    switch "" {
                    case "a":
                        var v59 string = ("hello" + ("a" + "\n\t"))
                        _ = v59
                    case "foo bar":
                        fmt.Println(v1, v49)
                    }
                    var v60 int8 = 1
                    v60 -= v60
                    _ = fi58
                    _ = v60
                }
            }
            if v49 {
                v1 += (v1 << (((v3 * v3)) & 63))
            } else {
                var v61 int = 0
                if (96 == 462) {
                    v62 := GenS6{F0: 144, F1: false, F2: -232}
                    v1 = 18446744073709551615
                    _ = v62
                } else {
                    if (!false) {
                        _ = len(v2)
                        var v63 string = "\n\t"
                        v49 = v49
                        _ = v63
                    }
                    v3 -= 0
                    v64 := GenS7{F0: 127, F1: 4294967295}
                    _ = v64
                }
                switch (v61 * v61) {
                case 76:
                    if v49 {
                        if (42 >= 1) {
                            for fi65 := 0; fi65 < 15; fi65++ {
                                v66 := GenS2{F0: 1}
                                _ = fi65
                                _ = v66
                            }
                        } else {
                            v67 := GenS4{F0: (0 <= 514), F1: (v61 - v61), F2: false}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v67
                        }
                    } else {
                        for fi68 := 0; fi68 < 11; fi68++ {
                            v1 = v3
                            _ = fi68
                        }
                        if v49 {
                            v61 = (v61 ^ (-(v61)))
                            switch ("\n\t" + "") {
                            case "":
                                v3 -= (v1 >> (((^(v1))) & 63))
                            case "foo bar":
                                var v69 uint = (^(v3))
                                _ = v69
                            case "\n\t":
                                for fi70 := 0; fi70 < 4; fi70++ {
                                    fmt.Println(v61, v61)
                                    _ = fi70
                                }
                            }
                        } else {
                            v1 += (v3 & v3)
                            v71 := GenS8{F0: v1}
                            var v72 float64 = -2.25
                            _ = v71
                            _ = v72
                        }
                        var v73 float64 = 3.14159
                        _ = v73
                    }
                case 149:
                    switch "" {
                    case "a":
                        v49 = ((!(uint16(v3) > 0)) || v49)
                    case "hello":
                        v74 := GenS3{F0: v1, F1: 1, F2: 42}
                        _ = v74
                    case "\n\t":
                        if v49 {
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        } else {
                            for fi75 := 0; fi75 < 13; fi75++ {
                                var v76 int64 = 9223372036854775807
                                switch "" {
                                case "\n\t":
                                    switch 4294967295 {
                                    case -2147483648:
                                        v61 = (fi75 & -1)
                                    case 0:
                                        if true {
                                            v77 := GenS9{F0: 417, F1: (fi75 + v61), F2: 1}
                                            _ = v77
                                        }
                                    case 776:
                                        v76 += (v76 << ((v1) & 63))
                                    }
                                case "hello":
                                    var v78 int64 = (v76 - v76)
                                    _ = v78
                                case "a":
                                    for fi79 := 0; fi79 < 15; fi79++ {
                                        if v49 {
                                            var v80 int32 = 2147483647
                                            var v81 int = (fi75 + 128)
                                            switch (v81 - (fi79 >> ((4294967295) & 63))) {
                                            case 32768:
                                                v49 = false
                                            case 0:
                                                var v82 int64 = v76
                                                _ = v82
                                            case 42:
                                                var v83 uint32 = 0
                                                var v84 int64 = 347
                                                v85 := GenS7{F0: 64, F1: 4294967295}
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            default:
                                                fmt.Println(v61, v3)
                                                var v86 string = "foo bar"
                                                _ = v86
                                            }
                                            _ = v80
                                            _ = v81
                                        }
                                        if (-2.25 >= ((3.14159 * 1.5) - ((float64((1e10 - ((-2.25 * 3.14159) - -2.25))) - ((1.5 + 1.5) * ((0.0 * 3.14159) * (float64((float32(float64(0.1)) - (float32(1e10) + -2.25))) * float64(1.5))))) * 1e10))) {
                                            if len(v2) > 0 {
                                                v2[0] = 905
                                            }
                                            _ = len(v2)
                                            if v49 {
                                                if (!(((("hello" + "\n\t") + "\n\t") + ("" + ("" + ("" + "")))) >= "a")) {
                                                    _ = len(v2)
                                                    if v49 {
                                                        _ = len(v2)
                                                        if v49 {
                                                            var v87 float32 = (0.1 * 1.5)
                                                            _ = v87
                                                        } else {
                                                            if (1 >= 669) {
                                                                v49 = (0 == 253)
                                                                if len(v2) > 0 {
                                                                    v2[0] = 1
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = uint32(v3)
                                                                }
                                                            }
                                                            fmt.Println(v61, v3)
                                                        }
                                                    } else {
                                                        var v88 int32 = 128
                                                        var v89 int = 475
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                }
                                                if (-2.25 != ((1.5 * 0.0) * (((0.1 * ((-2.25 - -2.25) + 1e10)) - -2.25) + (3.14159 * ((-2.25 * float64(((((0.0 / -2.25) / 1e10) + 3.14159) * 0.0))) * -2.25))))) {
                                                    if (v76 <= 42) {
                                                        var v90 bool = ((v76 + (v76 >> (((v3 + v1)) & 63))) <= (v76 * v76))
                                                        v61 += fi75
                                                        v61 += (fi79 % ((fi79) | 1))
                                                        _ = v90
                                                    }
                                                    if (127 >= 2) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v1 += 620
                                                        case "a":
                                                            var v91 uint64 = 4294967295
                                                            _ = v91
                                                        case "":
                                                            v61 += 0
                                                        }
                                                        if (true || (-128 <= -1)) {
                                                            var v92 float32 = float32(1.5)
                                                            _ = v92
                                                        } else {
                                                            var v93 float64 = 3.14159
                                                            var v94 uint8 = 158
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    } else {
                                                        if (61 >= 66) {
                                                            if false {
                                                                if len(v2) > 0 {
                                                                    v2[0] = 1
                                                                }
                                                            }
                                                            v76 -= 127
                                                            var v95 int8 = 42
                                                            var v96 int = (fi75 % ((-2147483648) | 1))
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 418
                                                        }
                                                    }
                                                } else {
                                                    if v49 {
                                                        fmt.Println(fi79, v61)
                                                    } else {
                                                        v3 -= v3
                                                    }
                                                    v97 := GenS10{F0: 385, F1: fi75, F2: ""}
                                                    _ = v97
                                                }
                                                v98 := GenS5{F0: v76, F1: 1, F2: v1}
                                                _ = v98
                                            }
                                        } else {
                                            switch v61 {
                                            case 42:
                                                v99 := GenS5{F0: v76, F1: 2147483647, F2: (v3 | v1)}
                                                _ = v99
                                            case -457:
                                                if v49 {
                                                    var v100 float32 = (3.14159 * ((0.0 / 0.1) * 1e10))
                                                    fmt.Println(v3, fi79, v76)
                                                    if (!v49) {
                                                        switch ("\n\t" + "foo bar") {
                                                        case "foo bar":
                                                            var v101 float32 = v100
                                                            _ = v101
                                                        }
                                                        fmt.Println(fi79, v1, v61)
                                                        if (-28 != int16(v76)) {
                                                            v49 = v49
                                                            v2 = append(v2, uint32(v1))
                                                        }
                                                    } else {
                                                        v102 := GenS11{F0: 1, F1: -1, F2: 100}
                                                        var v103 uint8 = 242
                                                        if v49 {
                                                            var v104 uint8 = 2
                                                            var v105 uint = v3
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    v61 = (fi75 - v61)
                                                    _ = v100
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = 951
                                                    }
                                                    var v106 int64 = int64(v61)
                                                    var v107 float32 = -2.25
                                                    fmt.Println(v76, v1, fi79)
                                                    _ = v106
                                                    _ = v107
                                                }
                                            case -9223372036854775808:
                                                var v108 int8 = 1
                                                _ = v108
                                            }
                                            fmt.Println(v61, v76)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = fi79
                                    }
                                }
                                switch (v61 / (((fi75 + (^(fi75)))) | 1)) {
                                case 128:
                                    _ = len(v2)
                                case 2147483647:
                                    fmt.Println(v76, v49, fi75)
                                default:
                                    v1 = (v3 * v3)
                                }
                                _ = fi75
                                _ = v76
                            }
                        }
                    default:
                        _ = len(v2)
                    }
                }
                _ = v61
            }
            v3 = 653
        } else {
            v49 = ((0.1 / 1.5) < (1e10 * 1.5))
            v3 = (v1 >> ((v1) & 63))
            var v109 float64 = -2.25
            _ = v109
        }
    case "foo bar":
        switch "a" {
        case "foo bar":
            v110 := GenS8{F0: (v3 / ((65535) | 1))}
            _ = v110
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v51
}
