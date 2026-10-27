package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8) uint16 {
    v5 := []int{}
    p0 = (p0 >> ((5) & 63))
    if len(v5) > 0 {
        v5[0] = 1
    }
    if (false && ((4294967295 != 2147483648) || (255 <= 2))) {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 uint64 = 2
            if (false && (!true)) {
                v8 := GenS1{F0: 231}
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        _ = len(v5)
        var v9 float32 = 1e10
        _ = v9
    }
    var v10 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 343
    var v3 int32 = (v2 << ((93) & 63))
    var v4 int = -113
    var v11 bool = true
    fmt.Println(v2, v2, v4)
    if v11 {
        var v12 uint = 2
        _ = v12
    }
    v13 := GenS2{F0: 1}
    switch -2147483648 {
    case 127:
        for fi14 := 0; fi14 < 3; fi14++ {
            if false {
                var v15 uint8 = 255
                _ = v15
            }
            var v16 bool = true
            v17 := []int{v4, v4, (v4 / ((fi14) | 1)), 2147483647}
            _ = fi14
            _ = v16
            _ = v17
        }
    case -9223372036854775808:
        var v18 int64 = 1
        _ = v18
    case 32768:
        v4 -= -9223372036854775808
    default:
        v19 := v1[v4]
        _ = v19
    }
    v20 := GenS3{F0: 4294967295}
    switch v4 {
    case -424:
        v21 := GenS4{F0: v4}
        _ = v21
    case 42:
        if (!(false || v11)) {
            var v22 int = -128
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint16 = 286
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 uint32 = 2
                    v27 := GenS4{F0: (-(v22))}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v11 = (0 <= -914)
                _ = fi23
                _ = v24
            }
            v28 := []int{(v4 / (((-(v4))) | 1)), (v22 >> ((65535) & 63))}
            _ = v22
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                v11 = (-1 == -249)
                var v30 float64 = 1.5
                switch "" {
                case "":
                    if v11 {
                        v30 += ((float64((0.1 / (1e10 + ((float32(float64(float32(1e10))) - 0.1) * float32(v30))))) + float64((1.5 - 0.1))) - v30)
                        if false {
                            switch "\n\t" {
                            case "a":
                                var v31 int32 = v3
                                _ = v31
                            }
                        } else {
                            v32 := []int{fi29, (v4 + (-(v4))), (fi29 >> ((2) & 63)), fi29}
                            if true {
                                var v33 int8 = 1
                                _ = v33
                            }
                            v4 = 796
                            _ = v32
                        }
                    }
                case "hello":
                    var v34 string = ("a" + ((("hello" + ("" + "\n\t")) + ("a" + (("" + ("a" + ("a" + "\n\t"))) + "hello"))) + ""))
                    _ = v34
                }
                _ = fi29
                _ = v30
            }
            v35 := GenS5{F0: 42, F1: (v2 << ((42) & 63))}
            _ = v35
        }
    default:
        if (42 <= v2) {
            switch (("\n\t" + (("hello" + ((("" + ("" + ("\n\t" + "foo bar"))) + "foo bar") + "\n\t")) + (("hello" + "foo bar") + "a"))) + "a") {
            case "hello":
                if v11 {
                    var v36 int64 = -9223372036854775808
                    for fi37 := 0; fi37 < 18; fi37++ {
                        for fi38 := 0; fi38 < 17; fi38++ {
                            switch "" {
                            case "\n\t":
                                var v39 uint32 = 2
                                v40 := GenS1{F0: 32767}
                                _ = v39
                                _ = v40
                            case "":
                                fmt.Println(v3, fi38)
                            }
                            var v41 uint32 = 541
                            _ = fi38
                            _ = v41
                        }
                        var v42 float32 = 1e10
                        var v43 int8 = 42
                        _ = fi37
                        _ = v42
                        _ = v43
                    }
                    fmt.Println(v2)
                    _ = v36
                }
            case "":
                v44 := GenS3{F0: 835}
                _ = v44
            }
        } else {
            fmt.Println(v3)
            var v45 uint = 0
            _ = v45
        }
    }
    if (v11 && v11) {
        switch (v4 * int(v2)) {
        case 32768:
            delete(v1, (-(v4)))
        case 2147483647:
            switch "hello" {
            case "hello":
                v46 := []int{}
                _ = v46
            case "\n\t":
                if ((v4 >> ((2) & 63)) <= v4) {
                    if v11 {
                        v47 := GenS6{F0: 9223372036854775807}
                        if (-62 != 41) {
                            for fi48 := 0; fi48 < 4; fi48++ {
                                var v49 string = ("hello" + "\n\t")
                                if (v49 >= "\n\t") {
                                    fmt.Println(fi48, v49)
                                    if (!v11) {
                                        var v50 int = (fi48 ^ fi48)
                                        _ = v50
                                    }
                                    switch (v4 % ((0) | 1)) {
                                    case 958:
                                        fmt.Println(v49, v49, v3)
                                    default:
                                        v51 := GenS2{F0: 118}
                                        _ = v51
                                    }
                                } else {
                                    var v52 float64 = (float64(((1e10 * 0.0) - float32(-2.25))) * (float64(((((((3.14159 / 1.0) / 0.1) - 1.5) / 0.1) - 0.0) - 0.0)) - -2.25))
                                    if v11 {
                                        if true {
                                            for fi53 := 0; fi53 < 17; fi53++ {
                                                v1[fi53] = (fi48 ^ (fi53 % (((-(v4))) | 1)))
                                                switch (v49 + ((v49 + v49) + v49)) {
                                                case "\n\t":
                                                    if v11 {
                                                        if (!(0.1 < (0.1 * 3.14159))) {
                                                            if (v49 <= ((v49 + v49) + v49)) {
                                                                var v54 int16 = 1000
                                                                var v55 float32 = ((-2.25 * -2.25) * 1.5)
                                                                v54 -= v54
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            v4 = fi48
                                                            v3 = v2
                                                        }
                                                        fmt.Println(v11)
                                                    } else {
                                                        if v11 {
                                                            switch (^(v4)) {
                                                            case 484:
                                                                v56 := GenS5{F0: -535, F1: 1}
                                                                _ = v56
                                                            }
                                                            var v57 uint = 1
                                                            v2 = v2
                                                            _ = v57
                                                        }
                                                        v11 = v11
                                                        v52 += float64(float32(float64((((-2.25 + 0.1) + float32(float64(float32((v52 / -2.25))))) - 0.0))))
                                                    }
                                                }
                                                _ = fi53
                                            }
                                            var v58 string = (v49 + "a")
                                            fmt.Println(v11)
                                            _ = v58
                                        } else {
                                            if (int16(v4) == 535) {
                                                v59 := GenS7{F0: (v4 & fi48), F1: -32768}
                                                var v60 uint8 = 163
                                                var v61 uint32 = 2147483647
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            }
                                            var v62 int8 = -96
                                            _ = v62
                                        }
                                        if true {
                                            fmt.Println(v49, v52)
                                            v4 -= (fi48 - -9223372036854775808)
                                        }
                                    } else {
                                        fmt.Println(v4)
                                        v11 = v11
                                        switch (v49 + v49) {
                                        case "\n\t":
                                            v52 = 0.1
                                        case "a":
                                            for fi63 := 0; fi63 < 19; fi63++ {
                                                switch "hello" {
                                                case "foo bar":
                                                    _ = len(v1)
                                                case "\n\t":
                                                    fmt.Println(v49)
                                                    v3 += (v2 - (v2 | (v2 ^ 2147483647)))
                                                    v64 := GenS3{F0: 291}
                                                    switch v49 {
                                                    case "hello":
                                                        var v65 int32 = (v3 % ((v2) | 1))
                                                        _ = v65
                                                    }
                                                    _ = v64
                                                }
                                                _ = fi63
                                            }
                                        }
                                    }
                                    _ = v52
                                }
                                _ = fi48
                                _ = v49
                            }
                            if false {
                                if v11 {
                                    v2 = (v2 / ((v3) | 1))
                                    for fi66 := 0; fi66 < 8; fi66++ {
                                        v2 = v2
                                        v2 = (v3 ^ 42)
                                        _ = fi66
                                    }
                                    fmt.Println(v2)
                                }
                            }
                        }
                        v67 := []int{v4, v4}
                        _ = v47
                        _ = v67
                    }
                    if v11 {
                        for fi68 := 0; fi68 < 8; fi68++ {
                            var v69 uint64 = 18
                            var v70 uint8 = 127
                            for fi71 := 0; fi71 < 2; fi71++ {
                                fmt.Println(v2, v11, fi68)
                                _ = len(v1)
                                v3 -= (v2 & v2)
                                _ = fi71
                            }
                            _ = fi68
                            _ = v69
                            _ = v70
                        }
                        var v72 int64 = 1
                        _ = v72
                    } else {
                        if (v11 || true) {
                            v11 = (-1 < 9223372036854775807)
                            v4 += -142
                            delete(v1, (v4 - (v4 + (v4 | -975))))
                        } else {
                            v73 := GenS3{F0: 678}
                            v74 := GenS1{F0: -525}
                            _ = v73
                            _ = v74
                        }
                    }
                    if ("a" == "hello") {
                        if (!(1.5 >= 1.5)) {
                            fmt.Println(v2)
                            var v75 uint8 = 200
                            _ = v75
                        } else {
                            if true {
                                for fi76 := 0; fi76 < 4; fi76++ {
                                    v77 := []int{-1, (-(v4)), (fi76 * (^(v4)))}
                                    _ = fi76
                                    _ = v77
                                }
                                switch "" {
                                case "hello":
                                    var v78 float64 = -2.25
                                    _ = v78
                                case "\n\t":
                                    v4 -= (v4 >> ((18446744073709551615) & 63))
                                }
                            }
                            fmt.Println(v11, v4, v2)
                        }
                    } else {
                        v79 := GenS8{F0: -32768}
                        v80 := GenS7{F0: v4, F1: 0}
                        var v81 int = v4
                        _ = v79
                        _ = v80
                        _ = v81
                    }
                }
            }
        default:
            fmt.Println(v2, v3, v4)
        }
    } else {
        v1[v4] = (v4 + (v4 << ((42) & 63)))
    }
    fmt.Println(v2)
    delete(v1, (v4 << ((542) & 63)))
    var v82 uint32 = 2147483648
    var v83 uint8 = 204
    for fi84 := 0; fi84 < 17; fi84++ {
        var v85 float32 = (3.14159 * 1e10)
        switch (fi84 - fi84) {
        case 1:
            v1[9223372036854775807] = v4
        }
        _ = fi84
        _ = v85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v82
    _ = v83
}
