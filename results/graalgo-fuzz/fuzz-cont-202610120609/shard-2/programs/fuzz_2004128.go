package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint; F1 uint; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS9 struct { F0 int32; F1 bool; F2 int32 }
type GenS10 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int8, p1 uint32, p2 int16) uint64 {
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 -= p0
        switch ("\n\t" + (("hello" + "a") + "a")) {
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: true}
                if (65535 > 0) {
                    var v9 uint = 255
                    fmt.Println(v9)
                    var v10 float32 = 0.1
                    _ = v10
                    _ = v9
                } else {
                    var v11 float64 = (0.1 - -2.25)
                    p2 = p2
                    var v12 int16 = p2
                    _ = v11
                    _ = v12
                }
                _ = fi7
                _ = v8
            }
        }
        _ = fi6
    }
    fmt.Println(p1, p2)
    var v13 int = 2147483647
    v14 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    var v3 uint32 = 42
    var v4 uint64 = 2
    var v5 uint8 = 1
    var v15 uint16 = 42
    v4 -= v4
    fmt.Println(v3, v4)
    switch ("foo bar" + ("hello" + "")) {
    case "":
        if false {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int32 = 2147483647
                v3 -= v3
                _ = fi16
                _ = v17
            }
        } else {
            switch ("\n\t" + "foo bar") {
            case "a":
                var v18 int16 = 42
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 2; fi19++ {
                    if false {
                        switch ("a" + "") {
                        case "":
                            v4 = uint64(v2)
                        }
                        switch (fi19 % ((-672) | 1)) {
                        case 32768:
                            v4 = uint64(v3)
                        default:
                            v5 = v5
                        }
                        fmt.Println(v4)
                    } else {
                        v20 := []int{(fi19 * (fi19 ^ (fi19 | (fi19 << ((0) & 63))))), fi19, (^(fi19))}
                        _ = v20
                    }
                    fmt.Println(v5, v15, v15)
                    if (1 <= 307) {
                        v21 := GenS2{F0: 4294967295, F1: fi19, F2: 434}
                        _ = v21
                    } else {
                        var v22 float32 = (1e10 / 0.1)
                        var v23 float64 = (((1e10 - ((1e10 * 0.1) + 0.0)) / 1.0) + 0.1)
                        v24 := []int{(fi19 * (fi19 | (fi19 * fi19))), fi19, fi19}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi19
                }
            case "":
                var v25 int64 = -1
                _ = v25
            default:
                var v26 int16 = -1
                _ = v26
            }
            var v27 float32 = -2.25
            v2 = (v15 + (v15 | v2))
            _ = v27
        }
    case "\n\t":
        v5 += 42
    }
    if ((true && (2 > (v5 % ((v5) | 1)))) && (!true)) {
        if true {
            for fi28 := 0; fi28 < 15; fi28++ {
                if ((("foo bar" + ("foo bar" + "\n\t")) + (("" + ("foo bar" + "a")) + "a")) > "a") {
                    v29 := GenS2{F0: 1, F1: (fi28 / ((fi28) | 1)), F2: 2}
                    if false {
                        v30 := GenS1{F0: false}
                        delete(v1, "hello")
                        _ = v30
                    }
                    fmt.Println(fi28, v5, v4)
                    _ = v29
                } else {
                    fmt.Println(v5, v4)
                    fmt.Println(v3)
                }
                _ = fi28
            }
            var v31 int32 = 1
            v3 = v3
            _ = v31
        }
        v32 := GenS3{F0: (v4 * 1), F1: 784}
        if ((!(!false)) && true) {
            v33 := GenS2{F0: 1, F1: 32767, F2: (v3 % ((v3) | 1))}
            var v34 int8 = -1
            _ = v33
            _ = v34
        }
        _ = v32
    } else {
        fmt.Println(v3, v15)
        for fi35 := 0; fi35 < 11; fi35++ {
            delete(v1, ("" + "a"))
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint16 = (v2 | (v15 ^ 370))
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS3{F0: v4, F1: 850}
        _ = v38
    }
    var v39 int16 = 32767
    if (111 >= -143) {
        v1[""] = ((genFn0(42, (v3 - uint32(v4)), v39) <= 4294967295) || (v39 <= v39))
        _ = len(v1)
        v5 = (^(v5))
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if true {
            v39 = 0
        }
        var v41 float64 = 3.14159
        var v42 int16 = (v39 | -32768)
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 string = "\n\t"
    fmt.Println(v15, v3, v5)
    var v44 uint16 = v2
    fmt.Println(v4, v44)
    fmt.Println(v15)
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 uint64 = v4
        switch 59 {
        case 4294967295:
            v46 -= v46
        case 2147483647:
            fmt.Println(v3)
        case 433:
            var v47 int = (fi45 % (((fi45 * 858)) | 1))
            _ = v47
        default:
            fmt.Println(v43, v2)
        }
        v48 := GenS4{F0: 45, F1: 779, F2: (^(v4))}
        _ = fi45
        _ = v46
        _ = v48
    }
    v49 := GenS5{F0: -1, F1: 2}
    v50 := v1[(v43 + ("" + v43))]
    if (!(!v50)) {
        var v51 int = int(v39)
        v52 := []int{(v51 * (v51 & (v51 + (v51 & v51)))), (v51 * int(v3)), 615, v51}
        _ = v51
        _ = v52
    } else {
        switch v43 {
        case "a":
            if (true && true) {
                if (0.1 > 3.14159) {
                    if v50 {
                        var v53 uint32 = (v3 + v3)
                        _ = v53
                    }
                } else {
                    for fi54 := 0; fi54 < 1; fi54++ {
                        switch (fi54 << ((532) & 63)) {
                        case 4294967295:
                            fmt.Println(v4)
                        case 538:
                            if (v39 < v39) {
                                for fi55 := 0; fi55 < 18; fi55++ {
                                    v43 = (("a" + ((v43 + "\n\t") + v43)) + ("" + (v43 + "")))
                                    _ = fi55
                                }
                                if (v50 || true) {
                                    v15 += v44
                                }
                            } else {
                                if v50 {
                                    v56 := GenS6{F0: v50, F1: 1}
                                    var v57 string = v43
                                    var v58 uint32 = v3
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                v44 = 2
                                var v59 int = fi54
                                _ = v59
                            }
                        default:
                            v2 = uint16(v4)
                        }
                        v60 := GenS2{F0: uint(v2), F1: 32767, F2: (v3 & uint32(v2))}
                        if (255 > 0) {
                            v44 += v15
                            if false {
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    if false {
                                        if true {
                                            fmt.Println(v3, fi61)
                                            v62 := GenS6{F0: ((v5 >> ((736) & 63)) < (v5 << ((868) & 63))), F1: 463}
                                            switch (^(fi61)) {
                                            case -158:
                                                var v63 int8 = 86
                                                _ = v63
                                            case 127:
                                                switch v43 {
                                                case "a":
                                                    var v64 float32 = float32((1.5 * -2.25))
                                                    _ = v64
                                                case "hello":
                                                    v65 := GenS5{F0: int64(v4), F1: 0}
                                                    _ = v65
                                                default:
                                                    var v66 int32 = int32(v5)
                                                    _ = v66
                                                }
                                            case 1:
                                                v43 = "hello"
                                            }
                                            _ = v62
                                        } else {
                                            var v67 int64 = 148
                                            if (91 == 100) {
                                                var v68 float32 = 0.1
                                                var v69 float32 = (float32(0.1) - float32(0.1))
                                                v70 := GenS3{F0: genFn0(42, v3, v39), F1: 110}
                                                v71 := v1[v43]
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                            if v50 {
                                                if (!((4294967295 == (v4 >> ((255) & 63))) || ((0.1 >= (1.5 - -2.25)) || (1 < 0)))) {
                                                    v15 = 0
                                                    v72 := GenS7{F0: v50, F1: 18446744073709551615}
                                                    if v50 {
                                                        var v73 uint32 = (v3 + (v3 >> ((uint(v3)) & 63)))
                                                        if (v4 <= uint64(v5)) {
                                                            v74 := []int{}
                                                            _ = v74
                                                        } else {
                                                            v73 = v73
                                                        }
                                                        v39 += 0
                                                        _ = v73
                                                    }
                                                    _ = v72
                                                }
                                            } else {
                                                var v75 float32 = ((3.14159 * 0.1) - 0.1)
                                                _ = v75
                                            }
                                            _ = v67
                                        }
                                        v3 = v3
                                        var v76 uint8 = uint8(v3)
                                        _ = v76
                                    }
                                    if (0 >= 302) {
                                        switch -2147483648 {
                                        case -898:
                                            fmt.Println(v15, fi61)
                                        default:
                                            var v77 int32 = 2147483647
                                            _ = v77
                                        }
                                    }
                                    switch "hello" {
                                    case "hello":
                                        var v78 int16 = 32767
                                        _ = v78
                                    }
                                    var v79 bool = false
                                    var v80 uint16 = (^(v44))
                                    _ = fi61
                                    _ = v79
                                    _ = v80
                                }
                            } else {
                                v81 := []int{(-(fi54)), (fi54 | (fi54 >> ((18446744073709551615) & 63)))}
                                v81 = append(v81, 32768)
                                v82 := GenS8{F0: (v5 % ((v5) | 1)), F1: 127, F2: uint16(v5)}
                                _ = v81
                                _ = v82
                            }
                        }
                        _ = fi54
                        _ = v60
                    }
                    if false {
                        var v83 int = 128
                        _ = v83
                    }
                    var v84 int16 = v39
                    _ = v84
                }
                for fi85 := 0; fi85 < 18; fi85++ {
                    if v50 {
                        if (("hello" <= "a") && (v43 <= ("foo bar" + v43))) {
                            if (1 <= -318) {
                                if (!(0 > -93)) {
                                    v86 := GenS5{F0: 533, F1: uint(v4)}
                                    _ = v86
                                }
                                var v87 uint8 = v5
                                _ = v87
                            }
                            var v88 uint64 = v4
                            fmt.Println(v39, fi85, v43)
                            _ = v88
                        } else {
                            for fi89 := 0; fi89 < 17; fi89++ {
                                v44 -= v2
                                _ = fi89
                            }
                            v90 := GenS9{F0: 1, F1: (1.5 > (0.1 * (float32((1.5 + ((3.14159 / 1.5) / 0.1))) * -2.25))), F2: -183}
                            for fi91 := 0; fi91 < 19; fi91++ {
                                v92 := GenS10{F0: (fi85 % (((fi85 ^ 128)) | 1)), F1: 808, F2: uint(v2)}
                                _ = fi91
                                _ = v92
                            }
                            _ = v90
                        }
                    } else {
                        v44 += uint16(v3)
                        switch ("a" + v43) {
                        case "foo bar":
                            switch 127 {
                            case 686:
                                for fi93 := 0; fi93 < 4; fi93++ {
                                    if true {
                                        v2 = (v44 | v44)
                                        fmt.Println(v4)
                                        var v94 string = v43
                                        v2 += v15
                                        _ = v94
                                    }
                                    v3 = v3
                                    var v95 string = "\n\t"
                                    _ = fi93
                                    _ = v95
                                }
                            }
                        case "a":
                            v39 = (v39 * (v39 | v39))
                        default:
                            if true {
                                v44 -= v44
                                var v96 uint16 = v44
                                _ = v96
                            }
                        }
                    }
                    v2 = (v15 << ((uint(v4)) & 63))
                    for fi97 := 0; fi97 < 17; fi97++ {
                        delete(v1, v43)
                        _ = fi97
                    }
                    _ = fi85
                }
                var v98 float32 = 0.0
                _ = v98
            } else {
                switch (v43 + "hello") {
                case "a":
                    var v99 uint32 = (v3 + v3)
                    _ = v99
                default:
                    if ((v4 & (^(v4))) > (v4 >> ((2) & 63))) {
                        v39 += (v39 >> ((218) & 63))
                        for fi100 := 0; fi100 < 4; fi100++ {
                            for fi101 := 0; fi101 < 15; fi101++ {
                                fmt.Println(v4, v4, v43)
                                if (v5 != v5) {
                                    if v50 {
                                        var v102 string = v43
                                        v5 = 2
                                        _ = v102
                                    }
                                    if (uint8(v3) == (v5 ^ (v5 / ((v5) | 1)))) {
                                        var v103 float64 = (0.1 / -2.25)
                                        var v104 int8 = 42
                                        fmt.Println(v44, v103)
                                        _ = v103
                                        _ = v104
                                    }
                                } else {
                                    v105 := []int{(fi101 % ((534) | 1)), (fi100 % ((fi101) | 1))}
                                    _ = v105
                                }
                                v2 -= (v44 ^ uint16(v4))
                                _ = fi101
                            }
                            var v106 string = (v43 + v43)
                            if (v39 > v39) {
                                var v107 float64 = 3.14159
                                var v108 uint = 42
                                _ = v107
                                _ = v108
                            } else {
                                v109 := []int{32767, fi100, fi100}
                                v110 := GenS2{F0: 942, F1: (fi100 % ((fi100) | 1)), F2: v3}
                                if false {
                                    for fi111 := 0; fi111 < 9; fi111++ {
                                        v5 -= 207
                                        _ = fi111
                                    }
                                    v4 -= genFn0(42, (^(v3)), int16(v44))
                                }
                                _ = v109
                                _ = v110
                            }
                            _ = fi100
                            _ = v106
                        }
                    }
                }
                for fi112 := 0; fi112 < 2; fi112++ {
                    switch v43 {
                    case "a":
                        v3 -= v3
                    case "\n\t":
                        _ = len(v1)
                    case "":
                        var v113 string = ((v43 + (("a" + v43) + (("a" + (v43 + v43)) + "a"))) + v43)
                        _ = v113
                    default:
                        switch (fi112 - int(v5)) {
                        case 4294967295:
                            fmt.Println(v43)
                        case -791:
                            var v114 int16 = 32767
                            _ = v114
                        default:
                            v115 := []int{(^(fi112)), int(v4), (fi112 - (fi112 + (fi112 - fi112))), fi112}
                            _ = v115
                        }
                    }
                    if (!(!(0 < 2147483647))) {
                        for fi116 := 0; fi116 < 15; fi116++ {
                            var v117 int8 = -12
                            _ = fi116
                            _ = v117
                        }
                        if v50 {
                            var v118 float64 = 1e10
                            switch int(v4) {
                            case 128:
                                var v119 float32 = 3.14159
                                _ = v119
                            }
                            v5 = (v5 >> ((255) & 63))
                            _ = v118
                        } else {
                            v120 := []int{(fi112 | (-(fi112))), int(v5)}
                            if false {
                                for fi121 := 0; fi121 < 7; fi121++ {
                                    v43 = (v43 + v43)
                                    fmt.Println(v15, v15, v4)
                                    _ = fi121
                                }
                                var v122 int32 = -1
                                _ = v122
                            }
                            _ = v120
                        }
                        v5 = (v5 & v5)
                    }
                    for fi123 := 0; fi123 < 10; fi123++ {
                        if (!v50) {
                            if ((!true) && v50) {
                                v2 = (^(v15))
                            }
                            fmt.Println(v4, fi112, v44)
                            _ = len(v1)
                        }
                        _ = fi123
                    }
                    _ = fi112
                }
                fmt.Println(v39, v4)
            }
        default:
            v39 -= v39
        }
    }
    var v124 uint64 = (v4 * v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v124
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v49
    _ = v5
    _ = v50
}
