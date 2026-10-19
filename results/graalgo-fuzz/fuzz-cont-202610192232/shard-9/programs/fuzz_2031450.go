package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }
type GenS5 struct { F0 int64; F1 uint8; F2 int16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint8; F1 bool; F2 int16 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8; F1 int64; F2 int8 }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32) int64 {
    v5 := []int{4294967295}
    var v6 int8 = -128
    p0 = 2147483648
    var v7 int64 = 880
    v8 := GenS1{F0: (v7 >> ((353) & 63)), F1: 149}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16) uint32 {
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{-713}
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 649, F1: 1}
        _ = fi9
        _ = v12
    }
    var v13 int = 972
    fmt.Println(p0, v13)
    var v14 uint16 = p0
    _ = p0
    _ = v13
    _ = v14
    return 808
}

func main() {
    var v1 int8 = 42
    var v2 int64 = -725
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]uint32)
    var v15 float64 = 3.14159
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := []int{}
        if false {
            if false {
                v18 := GenS2{F0: 255}
                _ = v18
            }
            v19 := GenS2{F0: uint8(v3)}
            _ = v19
        } else {
            delete(v4, 822)
        }
        _ = fi16
        _ = v17
    }
    v20 := []int{4294967295}
    v21 := GenS1{F0: (v2 * 267), F1: v3}
    switch ("" + ((("\n\t" + "a") + ((("hello" + "") + ("" + ("foo bar" + ""))) + "\n\t")) + "")) {
    case "hello":
        var v22 uint32 = 42
        _ = v22
    case "":
        if (32767 != 42) {
            v23 := GenS3{F0: -9223372036854775808}
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 float64 = v15
                v15 -= v15
                _ = fi24
                _ = v25
            }
            switch ("hello" + "foo bar") {
            case "a":
                if ((0.0 * ((0.1 + 3.14159) + 0.0)) >= -2.25) {
                    switch 32767 {
                    case 9223372036854775807:
                        var v26 float32 = (1e10 + 1.5)
                        _ = v26
                    case 2147483647:
                        var v27 uint64 = (^(v3))
                        _ = v27
                    }
                    v20 = append(v20, -239)
                }
            }
            _ = v23
        }
    case "\n\t":
        var v28 int = 32767
        _ = v28
    }
    var v29 int16 = 0
    fmt.Println(v2, v1)
    if true {
        switch -337 {
        case 1:
            v2 = (v2 >> ((65535) & 63))
        }
        switch ("\n\t" + "\n\t") {
        case "":
            switch "hello" {
            case "foo bar":
                var v30 uint8 = 42
                _ = v30
            case "":
                fmt.Println(v15, v2)
            default:
                if false {
                    var v31 float64 = 1.5
                    _ = v31
                } else {
                    v1 = (v1 * 1)
                    if true {
                        var v32 uint = 945
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 11; fi33++ {
                            var v34 uint32 = 889
                            v2 -= v2
                            v20 = append(v20, 42)
                            _ = fi33
                            _ = v34
                        }
                        var v35 uint = 42
                        for fi36 := 0; fi36 < 10; fi36++ {
                            switch (("\n\t" + (("a" + ("\n\t" + ("" + "\n\t"))) + "a")) + "foo bar") {
                            case "foo bar":
                                v37 := v4[2]
                                _ = v37
                            case "a":
                                fmt.Println(v15)
                            default:
                                if (255 > 0) {
                                    switch fi36 {
                                    case 32767:
                                        switch (fi36 >> ((v35) & 63)) {
                                        case 4294967295:
                                            var v38 bool = ((true || (int64(v1) < v2)) && (!true))
                                            _ = v38
                                        case 805:
                                            v1 += v1
                                        case -9223372036854775808:
                                            switch fi36 {
                                            case 4294967295:
                                                var v39 uint16 = 2
                                                _ = v39
                                            case -1:
                                                if (((-2.25 + 3.14159) * v15) != v15) {
                                                    fmt.Println(v1)
                                                    var v40 uint64 = v3
                                                    _ = v40
                                                }
                                            case -848:
                                                switch "hello" {
                                                case "":
                                                    var v41 bool = ((fi36 - (fi36 >> ((2) & 63))) <= int(v2))
                                                    _ = v41
                                                case "hello":
                                                    switch fi36 {
                                                    case 32767:
                                                        v29 -= v29
                                                    }
                                                case "foo bar":
                                                    var v42 int32 = int32(v29)
                                                    _ = v42
                                                default:
                                                    for fi43 := 0; fi43 < 19; fi43++ {
                                                        v20 = append(v20, -9223372036854775808)
                                                        _ = fi43
                                                    }
                                                }
                                            default:
                                                v29 = (v29 - (v29 % ((v29) | 1)))
                                            }
                                        default:
                                            v35 = (v35 & v35)
                                        }
                                    case 2147483647:
                                        fmt.Println(v29, v29, v1)
                                    case 32768:
                                        if false {
                                            fmt.Println(v2, v15)
                                        } else {
                                            v44 := GenS2{F0: 247}
                                            for fi45 := 0; fi45 < 5; fi45++ {
                                                v20 = append(v20, (fi36 / ((fi45) | 1)))
                                                var v46 bool = ((1 > 0) || (439 <= 1))
                                                _ = fi45
                                                _ = v46
                                            }
                                            for fi47 := 0; fi47 < 3; fi47++ {
                                                v48 := GenS2{F0: 59}
                                                v2 -= (v2 ^ v2)
                                                var v49 uint8 = 34
                                                _ = fi47
                                                _ = v48
                                                _ = v49
                                            }
                                            _ = v44
                                        }
                                    default:
                                        for fi50 := 0; fi50 < 7; fi50++ {
                                            var v51 uint16 = 506
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                    switch 4294967295 {
                                    case 0:
                                        for fi52 := 0; fi52 < 5; fi52++ {
                                            var v53 uint = v35
                                            v29 = (v29 | v29)
                                            _ = fi52
                                            _ = v53
                                        }
                                    case 158:
                                        v54 := GenS4{F0: 0, F1: (v35 & 4294967295), F2: v29}
                                        _ = v54
                                    }
                                    v55 := GenS4{F0: v2, F1: (v35 % (((v35 + v35)) | 1)), F2: v29}
                                    _ = v55
                                } else {
                                    for fi56 := 0; fi56 < 17; fi56++ {
                                        if true {
                                            var v57 int = (fi56 | fi56)
                                            _ = v57
                                        }
                                        switch ("\n\t" + "\n\t") {
                                        case "\n\t":
                                            v29 += (^(v29))
                                        case "":
                                            switch "" {
                                            case "":
                                                var v58 uint16 = 0
                                                _ = v58
                                            default:
                                                switch (^(fi56)) {
                                                case 612:
                                                    fmt.Println(v15, v35, v15)
                                                case 629:
                                                    _ = len(v20)
                                                }
                                            }
                                        }
                                        fmt.Println(v1)
                                        v3 -= v3
                                        _ = fi56
                                    }
                                    var v59 uint16 = 255
                                    if (3.14159 > (3.14159 - float32(1.5))) {
                                        v60 := v4[1]
                                        v61 := GenS5{F0: v2, F1: uint8(v35), F2: (v29 | v29)}
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v59
                                }
                            }
                            v62 := GenS6{F0: (-(v29)), F1: int64(v29)}
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = fi36
                            _ = v62
                        }
                        _ = v35
                    }
                }
            }
        case "hello":
            v1 = v1
        default:
            switch 567 {
            case 57:
                var v63 string = ""
                _ = v63
            case 2147483647:
                var v64 int32 = -592
                _ = v64
            }
        }
    } else {
        switch 4294967295 {
        case 1:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        }
        fmt.Println(v3)
        for fi65 := 0; fi65 < 4; fi65++ {
            v66 := GenS1{F0: v2, F1: v3}
            v15 += ((-2.25 * -2.25) / (v15 * float64(((float32(1.5) / -2.25) + -2.25))))
            var v67 uint32 = genFn1(2)
            _ = fi65
            _ = v66
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        for fi69 := 0; fi69 < 13; fi69++ {
            switch (("\n\t" + "\n\t") + ("\n\t" + "")) {
            case "a":
                _ = len(v20)
            case "hello":
                var v70 int64 = 0
                _ = v70
            case "foo bar":
                var v71 int32 = 42
                _ = v71
            default:
                fmt.Println(v29, v2, v3)
            }
            _ = len(v4)
            v15 += float64(0.0)
            fmt.Println(v3, fi69)
            if true {
                if false {
                    if false {
                        var v72 int32 = 100
                        var v73 uint64 = v3
                        fmt.Println(v72, v15)
                        _ = v72
                        _ = v73
                    } else {
                        v1 = (^(v1))
                    }
                    var v74 int8 = (v1 ^ v1)
                    _ = v74
                }
                v75 := GenS3{F0: 42}
                fmt.Println(v3, v29)
                _ = v75
            }
            var v76 uint64 = v3
            v76 += (^(v76))
            fmt.Println(v1, v1)
            var v77 uint = 695
            if true {
                if (fi68 <= (fi69 - (fi68 & (fi68 >> ((1) & 63))))) {
                    if (2147483648 > 65535) {
                        v76 += v3
                        v78 := GenS1{F0: v2, F1: (v3 * (v76 << ((42) & 63)))}
                        v2 -= genFn0(genFn1(240))
                        _ = v78
                    } else {
                        switch "" {
                        case "foo bar":
                            v3 = v76
                        default:
                            v79 := GenS7{F0: 128, F1: ((v15 + v15) <= -2.25), F2: v29}
                            _ = v79
                        }
                    }
                    var v80 uint64 = v3
                    if (true || false) {
                        if (1 == 76) {
                            var v81 int8 = -1
                            _ = v81
                        }
                    } else {
                        v2 -= (^(v2))
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v4[4294967295] = 2147483647
                    case "foo bar":
                        var v82 int32 = -1
                        _ = v82
                    case "hello":
                        var v83 uint32 = 2
                        _ = v83
                    default:
                        v15 += ((-2.25 * v15) / ((v15 + (v15 / 1.5)) - float64(((-2.25 / 1.5) + -2.25))))
                    }
                    v29 = -565
                    switch (fi68 << ((v77) & 63)) {
                    case 42:
                        var v84 uint32 = 401
                        _ = v84
                    case -9223372036854775808:
                        v85 := GenS8{F0: -2147483648, F1: 42}
                        _ = v85
                    }
                    switch fi68 {
                    case 32767:
                        var v86 float32 = 3.14159
                        _ = v86
                    case 32768:
                        if (!true) {
                            _ = len(v20)
                            v87 := GenS9{F0: int64(v1)}
                            _ = v87
                        } else {
                            v88 := GenS10{F0: 173, F1: genFn0(2147483647), F2: v1}
                            switch (fi69 << (((v77 & v77)) & 63)) {
                            case 128:
                                if (!false) {
                                    v29 = (v29 - v29)
                                    v2 -= 9223372036854775807
                                    v1 -= (v1 ^ v1)
                                }
                            default:
                                var v89 uint16 = 255
                                _ = v89
                            }
                            delete(v4, genFn1(2))
                            _ = v88
                        }
                    case 42:
                        v2 -= v2
                    default:
                        var v90 uint32 = 42
                        _ = v90
                    }
                    _ = v80
                }
                v2 = (v2 - v2)
                v77 -= (v77 * v77)
            } else {
                v1 += (-(v1))
                v91 := GenS5{F0: v2, F1: 1, F2: v29}
                var v92 uint64 = (v76 + (v76 + 4294967295))
                var v93 string = ""
                if true {
                    v1 = (v1 ^ -48)
                }
                _ = v91
                _ = v92
                _ = v93
            }
            var v94 uint16 = 846
            _ = fi69
            _ = v76
            _ = v77
            _ = v94
        }
        for fi95 := 0; fi95 < 12; fi95++ {
            v1 -= int8(v3)
            _ = fi95
        }
        _ = fi68
    }
    if (1 != 111) {
        v3 = v3
        var v96 int64 = v2
        _ = v96
    }
    delete(v4, 836)
    for fi97 := 0; fi97 < 3; fi97++ {
        var v98 uint = 855
        _ = fi97
        _ = v98
    }
    var v99 float32 = 3.14159
    for fi100 := 0; fi100 < 14; fi100++ {
        fmt.Println(v3)
        _ = fi100
    }
    var v101 uint = 65535
    if (!false) {
        if ((v15 / v15) >= 0.0) {
            var v102 bool = true
            switch ("hello" + "foo bar") {
            case "\n\t":
                fmt.Println(v101, v2)
            default:
                v103 := GenS11{F0: v3, F1: v101}
                _ = v103
            }
            if len(v20) > 0 {
                v20[0] = 2147483647
            }
            _ = v102
        } else {
            var v104 bool = false
            switch 128 {
            case -957:
                v1 += v1
            }
            _ = v104
        }
        for fi105 := 0; fi105 < 8; fi105++ {
            v106 := GenS1{F0: 0, F1: 18446744073709551615}
            var v107 int64 = 42
            v101 += v101
            _ = fi105
            _ = v106
            _ = v107
        }
    } else {
        v2 -= genFn0(2147483648)
        for fi108 := 0; fi108 < 6; fi108++ {
            var v109 uint64 = v3
            var v110 int = fi108
            v111 := GenS12{F0: 0, F1: (v2 ^ v2)}
            _ = fi108
            _ = v109
            _ = v110
            _ = v111
        }
    }
    v112 := GenS5{F0: v2, F1: 86, F2: (v29 & (v29 * v29))}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v99:", v99)
    _ = v1
    _ = v101
    _ = v112
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v99
}
