package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 uint8; F1 bool; F2 uint64 }
func genFn0(p0 string, p1 uint32, p2 uint32) bool {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int16 = 32767
        fmt.Println(p2)
        _ = fi5
        _ = v6
    }
    v7 := []int{-2147483648, 128}
    fmt.Println(p1, p0, p2)
    var v8 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (-55 <= 127)
}

func main() {
    v1 := []uint32{4294967295}
    var v2 string = ""
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483647
    var v9 float64 = 0.1
    v4 -= (v4 - (v4 & 363))
    _ = len(v3)
    switch "a" {
    case "hello":
        v2 = ("" + v2)
    case "\n\t":
        fmt.Println(v2, v9, v9)
    default:
        fmt.Println(v9, v9, v4)
    }
    delete(v3, v4)
    v2 = v2
    var v10 float64 = (v9 * (v9 - float64(1.5)))
    var v11 int16 = 652
    fmt.Println(v10, v11)
    for fi12 := 0; fi12 < 1; fi12++ {
        v9 += 1e10
        v4 -= v4
        for fi13 := 0; fi13 < 9; fi13++ {
            v10 -= 1.5
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v2, v10)
    for fi14 := 0; fi14 < 5; fi14++ {
        if (false && false) {
            v11 += (v11 >> ((902) & 63))
        } else {
            v15 := v3[(v4 / ((280) | 1))]
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 int = (fi14 ^ fi14)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((true || false) && (1 <= 32767)) {
                    var v18 int16 = (-(v11))
                    _ = v18
                } else {
                    switch -133 {
                    case 896:
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    case 823:
                        if (true && true) {
                            switch (fi16 / (((^(v17))) | 1)) {
                            case 9223372036854775807:
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (v4 / ((v4) | 1))
                                    }
                                    delete(v3, (v4 << ((65535) & 63)))
                                    if true {
                                        fmt.Println(fi16, fi16, fi16)
                                        if true {
                                            var v20 float64 = -2.25
                                            v21 := GenS1{F0: false}
                                            _ = v20
                                            _ = v21
                                        }
                                    } else {
                                        var v22 uint8 = 33
                                        switch (v15 + v2) {
                                        case "\n\t":
                                            var v23 int = v17
                                            _ = v23
                                        case "a":
                                            if (((v22 - 128) == (v22 >> ((18446744073709551615) & 63))) || true) {
                                                fmt.Println(v11)
                                            } else {
                                                v10 = v10
                                                var v24 int64 = 0
                                                _ = v24
                                            }
                                            if (!(v11 >= v11)) {
                                                v25 := GenS2{F0: 1, F1: (v4 << ((4294967295) & 63)), F2: true}
                                                v15 = ""
                                                switch (fi16 << ((uint(v22)) & 63)) {
                                                case 42:
                                                    if len(v1) > 0 {
                                                        v1[0] = 65535
                                                    }
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v25
                                            } else {
                                                v2 = "a"
                                                v3[(v4 | 2)] = v15
                                                switch (^(fi16)) {
                                                case -2147483648:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    if true {
                                                        v26 := GenS3{F0: 0, F1: 219, F2: (v22 | v22)}
                                                        _ = v26
                                                    }
                                                }
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (v4 * v4)
                                        }
                                        _ = v22
                                    }
                                } else {
                                    fmt.Println(fi16, v9, fi16)
                                    v11 = (^(v11))
                                }
                            case 32767:
                                v15 = "hello"
                            case 128:
                                fmt.Println(v4, fi14, v17)
                            default:
                                if true {
                                    v27 := GenS2{F0: 374, F1: v4, F2: (-44 <= -9223372036854775808)}
                                    _ = v27
                                } else {
                                    var v28 uint32 = (^(v4))
                                    if false {
                                        switch fi16 {
                                        case -9223372036854775808:
                                            v3[v28] = "hello"
                                        case 32767:
                                            v11 -= -1
                                        case 128:
                                            fmt.Println(fi14, v28, v28)
                                            if len(v1) > 0 {
                                                v1[0] = v28
                                            }
                                            if (9223372036854775807 != 47) {
                                                var v29 float64 = v10
                                                _ = v29
                                            } else {
                                                delete(v3, v4)
                                            }
                                            var v30 bool = true
                                            _ = v30
                                        default:
                                            var v31 int32 = -2147483648
                                            _ = v31
                                        }
                                    } else {
                                        v11 += -32768
                                    }
                                    switch (fi14 % ((503) | 1)) {
                                    case -505:
                                        v9 = v9
                                    case -2147483648:
                                        fmt.Println(v15)
                                    case -529:
                                        var v32 uint = 574
                                        _ = v32
                                    default:
                                        v28 -= (v28 | v4)
                                    }
                                    if true {
                                        v33 := GenS4{F0: 393, F1: 179}
                                        _ = v33
                                    }
                                    _ = v28
                                }
                            }
                            if (v11 != v11) {
                                var v34 uint16 = 0
                                _ = v34
                            }
                            var v35 int32 = 42
                            _ = v35
                        }
                    case -179:
                        fmt.Println(v10)
                    }
                    switch v15 {
                    case "":
                        v9 += -2.25
                    case "foo bar":
                        if false {
                            v36 := GenS5{F0: 1, F1: false}
                            var v37 float64 = v10
                            var v38 string = "a"
                            fmt.Println(fi14)
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            var v39 string = v15
                            _ = v39
                        }
                        v11 = (v11 - (-(v11)))
                        if true {
                            v2 = "hello"
                        }
                        switch (fi14 ^ v17) {
                        case 128:
                            v40 := GenS6{F0: 2, F1: ("foo bar" == (v2 + (("hello" + v2) + v2))), F2: (-(v11))}
                            _ = v40
                        case 32767:
                            var v41 uint = 4294967295
                            _ = v41
                        case 2147483647:
                            v2 = (v15 + v2)
                        default:
                            var v42 uint = 277
                            _ = v42
                        }
                        switch v15 {
                        case "foo bar":
                            switch (fi16 << ((255) & 63)) {
                            case -487:
                                v9 = float64(0.1)
                            case -9223372036854775808:
                                v43 := v3[127]
                                _ = v43
                            default:
                                v17 = v17
                            }
                        case "hello":
                            v15 = v2
                        }
                    }
                    v1 = append(v1, 209)
                    if (515 >= 678) {
                        v44 := GenS7{F0: -1, F1: v17}
                        switch "foo bar" {
                        case "foo bar":
                            switch (("\n\t" + v15) + v15) {
                            case "a":
                                if ((v17 - 128) == (^(fi16))) {
                                    switch v15 {
                                    case "\n\t":
                                        v2 = ((v15 + (v15 + "a")) + v2)
                                    case "hello":
                                        if true {
                                            var v45 uint = 1
                                            var v46 float32 = 1e10
                                            if false {
                                                switch v2 {
                                                case "a":
                                                    var v47 uint = 2
                                                    _ = v47
                                                case "\n\t":
                                                    if true {
                                                        fmt.Println(v11, v45)
                                                    } else {
                                                        v9 += v10
                                                        fmt.Println(v10)
                                                        v46 += 3.14159
                                                    }
                                                }
                                            } else {
                                                var v48 int = (fi14 | fi14)
                                                _ = v48
                                            }
                                            _ = v45
                                            _ = v46
                                        } else {
                                            switch v17 {
                                            case -453:
                                                var v49 uint8 = 113
                                                _ = v49
                                            case 942:
                                                v2 = v2
                                            default:
                                                delete(v3, v4)
                                            }
                                        }
                                    }
                                }
                            case "":
                                var v50 uint = 65535
                                delete(v3, v4)
                                var v51 uint16 = 32767
                                _ = v50
                                _ = v51
                            default:
                                if true {
                                    var v52 string = v2
                                    _ = v52
                                } else {
                                    v11 = v11
                                }
                            }
                        case "hello":
                            v53 := GenS4{F0: 289, F1: 2}
                            _ = v53
                        case "":
                            fmt.Println(v10, v11, v17)
                        default:
                            v54 := GenS2{F0: 32767, F1: v4, F2: (403 > 340)}
                            _ = v54
                        }
                        if false {
                            var v55 bool = false
                            switch fi16 {
                            case 4294967295:
                                if v55 {
                                    var v56 float32 = -2.25
                                    v55 = (572 > 42)
                                    _ = v56
                                } else {
                                    switch (fi16 / ((v17) | 1)) {
                                    case 0:
                                        var v57 uint8 = 2
                                        _ = v57
                                    case 32768:
                                        fmt.Println(v4, v9)
                                    default:
                                        var v58 uint32 = (v4 + v4)
                                        _ = v58
                                    }
                                    fmt.Println(fi14)
                                    v11 += v11
                                    var v59 uint = 0
                                    _ = v59
                                }
                            default:
                                switch ((((v2 + v2) + ("a" + (v2 + (v2 + ("hello" + (v2 + "\n\t")))))) + "hello") + (("" + v2) + v15)) {
                                case "foo bar":
                                    v60 := GenS3{F0: uint16(v4), F1: 677, F2: 128}
                                    _ = v60
                                case "":
                                    if ((v10 + ((((1e10 / ((v9 + -2.25) - (0.0 - 1.5))) * v10) / (v9 * (3.14159 * v9))) / (v9 - v10))) != (float64(-2.25) - v9)) {
                                        fmt.Println(v10, v9)
                                        v10 += -2.25
                                        v61 := GenS8{F0: ("" + "hello"), F1: 42}
                                        _ = v61
                                    }
                                case "hello":
                                    v17 += v17
                                }
                            }
                            _ = v55
                        }
                        _ = v44
                    }
                    if false {
                        _ = len(v3)
                    }
                    var v62 uint8 = 0
                    if false {
                        if (1.5 >= (v10 - (v9 * v10))) {
                            _ = len(v3)
                            v15 = v2
                            v10 = v10
                        }
                    } else {
                        var v63 float64 = v9
                        _ = v63
                    }
                    _ = v62
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        if false {
            fmt.Println(v9, v9)
            v64 := GenS9{F0: 42, F1: ("hello" >= "hello"), F2: 954}
            _ = v64
        }
        _ = fi14
    }
    if (39 >= 175) {
        for fi65 := 0; fi65 < 17; fi65++ {
            for fi66 := 0; fi66 < 2; fi66++ {
                v2 = "hello"
                _ = fi66
            }
            for fi67 := 0; fi67 < 3; fi67++ {
                v68 := GenS6{F0: 1, F1: true, F2: (v11 & v11)}
                v11 += int16(v4)
                _ = fi67
                _ = v68
            }
            var v69 uint8 = 196
            _ = fi65
            _ = v69
        }
        var v70 int64 = -523
        _ = v70
    }
    _ = len(v1)
    for fi71 := 0; fi71 < 5; fi71++ {
        v11 = (v11 + v11)
        _ = fi71
    }
    for fi72 := 0; fi72 < 7; fi72++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!(-34 > -185)) {
            fmt.Println(v10, v4, v9)
            fmt.Println(v2)
        } else {
            fmt.Println(v10, v4, v4)
        }
        _ = fi72
    }
    var v73 uint32 = 4294967295
    for fi74 := 0; fi74 < 13; fi74++ {
        _ = len(v3)
        _ = fi74
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v9
}
