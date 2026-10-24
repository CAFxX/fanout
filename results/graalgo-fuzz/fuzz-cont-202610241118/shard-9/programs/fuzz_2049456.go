package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
func genFn0() int16 {
    var v7 uint8 = 255
    v8 := GenS1{F0: 262}
    v9 := GenS2{F0: -1}
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = 0
    v2 := []string{"foo bar", ("a" + "a"), ("\n\t" + "\n\t")}
    v3 := make(map[uint32]uint32)
    var v4 float32 = (((1.5 + (((0.0 + float32((float64(0.1) + ((float64(0.1) * -2.25) / 0.1)))) * (float32(float64(0.0)) / 1.5)) - 1.5)) / 1.5) + 1e10)
    v5 := make(map[int]uint32)
    v6 := []int{-217, 127}
    v10 := v3[2147483647]
    var v11 uint64 = 0
    if (true && false) {
        if (!false) {
            var v12 int16 = 1000
            _ = v12
        }
        if (-1 <= 1) {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 string = "\n\t"
                v15 := GenS1{F0: 2147483648}
                var v16 uint64 = v11
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 int16 = 42
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int64 = 252
                var v20 float32 = v4
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS2{F0: 0}
                    if (((v1 * (v1 ^ -2147483648)) < int32(v11)) || ((fi18 << ((uint(v10)) & 63)) == (fi21 + fi21))) {
                        var v23 int64 = v19
                        _ = v23
                    } else {
                        var v24 float64 = (1e10 + 1e10)
                        var v25 int16 = v17
                        if ((v17 - genFn0()) > v25) {
                            v4 = (v4 * (1e10 / ((v4 / v4) * (v4 / v4))))
                            v26 := GenS3{F0: 2, F1: 491}
                            _ = v26
                        }
                        _ = v24
                        _ = v25
                    }
                    if (42 > 42) {
                        switch "" {
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                var v27 float64 = (float64(v4) + (-2.25 / 1.5))
                                _ = v27
                            }
                        default:
                            v10 = 563
                        }
                        v1 -= (^(v1))
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
        var v28 bool = true
        _ = v28
    } else {
        v10 = (v10 * 1)
        switch -2147483648 {
        case 1:
            switch 32768 {
            case -342:
                switch 9223372036854775807 {
                case 0:
                    v10 = 2147483647
                case 1:
                    for fi29 := 0; fi29 < 4; fi29++ {
                        var v30 uint8 = 90
                        switch (^(fi29)) {
                        case 42:
                            switch (((("" + "") + (("" + ("\n\t" + ("hello" + ("\n\t" + "hello")))) + ((("a" + "foo bar") + "a") + ("" + "hello")))) + "\n\t") + ("\n\t" + "foo bar")) {
                            case "\n\t":
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    v31 := GenS2{F0: 0}
                                    _ = v31
                                case "\n\t":
                                    for fi32 := 0; fi32 < 5; fi32++ {
                                        var v33 uint64 = v11
                                        if ((true || true) || (!false)) {
                                            _ = len(v3)
                                            var v34 int32 = v1
                                            var v35 bool = (!false)
                                            _ = v34
                                            _ = v35
                                        } else {
                                            var v36 uint32 = (v10 * v10)
                                            v37 := v3[(v10 & (v10 / ((v36) | 1)))]
                                            _ = v36
                                            _ = v37
                                        }
                                        if true {
                                            delete(v5, fi32)
                                        } else {
                                            _ = len(v6)
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                case "":
                                    fmt.Println(fi29, v11, fi29)
                                default:
                                    if true {
                                        if false {
                                            v38 := GenS4{F0: 0, F1: 42, F2: 42}
                                            v2 = append(v2, "\n\t")
                                            _ = v38
                                        } else {
                                            v10 -= 2
                                            if (true || (uint(v10) != 338)) {
                                                var v39 int = (fi29 * fi29)
                                                v11 = (v11 ^ (^(v11)))
                                                var v40 uint16 = 414
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        v41 := GenS1{F0: 935}
                                        var v42 uint64 = v11
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 int32 = (v1 * 108)
                                        var v44 float64 = ((float64(1.5) * (1.5 / (1e10 / ((-2.25 * (((0.0 - 0.0) / 1.5) + 1e10)) * ((1.5 - -2.25) * (-2.25 - float64((v4 - 3.14159)))))))) * float64(v4))
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            case "":
                                var v45 uint8 = (v30 - v30)
                                _ = v45
                            }
                        }
                        for fi46 := 0; fi46 < 16; fi46++ {
                            v10 = v10
                            if ((255 > 1) && true) {
                                v10 -= v10
                                if (v11 >= uint64(v30)) {
                                    v4 += v4
                                }
                                var v47 int64 = -1
                                _ = v47
                            } else {
                                var v48 int = 478
                                if (-127 >= 100) {
                                    switch v48 {
                                    case 127:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        if (2 > 1) {
                                            var v49 string = "a"
                                            _ = v49
                                        } else {
                                            v5[(v48 >> ((4294967295) & 63))] = v10
                                            v50 := v3[(v10 % ((v10) | 1))]
                                            v4 -= v4
                                            _ = v50
                                        }
                                    }
                                    var v51 bool = false
                                    v52 := v5[fi29]
                                    _ = v51
                                    _ = v52
                                }
                                if len(v6) > 0 {
                                    v6[0] = -839
                                }
                                _ = v48
                            }
                            if false {
                                if (((true || true) && (((float32(3.14159) + v4) - 3.14159) == v4)) || false) {
                                    v53 := GenS3{F0: uint8(v11), F1: (v10 | (^(v10)))}
                                    _ = v53
                                }
                                if false {
                                    switch "foo bar" {
                                    case "":
                                        _ = len(v2)
                                    case "foo bar":
                                        fmt.Println(fi29, v1, fi46)
                                    case "a":
                                        var v54 bool = false
                                        _ = v54
                                    default:
                                        _ = len(v2)
                                    }
                                } else {
                                    var v55 bool = (!true)
                                    _ = v55
                                }
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = fi46
                                }
                                fmt.Println(v1, v4)
                                fmt.Println(v30, v10, v11)
                            }
                            if true {
                                v2 = append(v2, "foo bar")
                                _ = len(v5)
                            } else {
                                switch (fi46 << ((517) & 63)) {
                                case 4294967295:
                                    fmt.Println(fi29, fi46)
                                case -2147483648:
                                    v56 := GenS2{F0: 9223372036854775807}
                                    _ = v56
                                }
                                var v57 uint64 = v11
                                switch (fi46 ^ fi46) {
                                case -2147483648:
                                    v4 -= (1e10 * v4)
                                case -111:
                                    if (!true) {
                                        if (-128 <= int8(v57)) {
                                            fmt.Println(v10, v11, fi46)
                                            v58 := GenS2{F0: 1}
                                            _ = v58
                                        }
                                    } else {
                                        v10 += v10
                                        v59 := GenS1{F0: v10}
                                        _ = v59
                                    }
                                case 2147483647:
                                    if true {
                                        v60 := GenS5{F0: -273, F1: (v30 * v30)}
                                        _ = v60
                                    }
                                }
                                _ = v57
                            }
                            _ = fi46
                        }
                        _ = fi29
                        _ = v30
                    }
                }
            }
        case 954:
            for fi61 := 0; fi61 < 1; fi61++ {
                v62 := GenS6{F0: v10, F1: 2}
                v63 := GenS7{F0: -124}
                _ = fi61
                _ = v62
                _ = v63
            }
        case 42:
            switch ("a" + ("\n\t" + (("foo bar" + "") + "foo bar"))) {
            case "":
                v64 := GenS1{F0: v10}
                _ = v64
            default:
                var v65 int64 = -1
                _ = v65
            }
        }
    }
    _ = len(v5)
    switch -482 {
    case 78:
        if (0 >= 65535) {
            _ = len(v2)
            v1 = (^(v1))
            var v66 uint8 = 42
            _ = v66
        }
    case 9223372036854775807:
        v10 -= (v10 >> ((0) & 63))
    case 128:
        fmt.Println(v1)
    default:
        for fi67 := 0; fi67 < 7; fi67++ {
            fmt.Println(v10, v10, fi67)
            _ = fi67
        }
    }
    if false {
        if true {
            _ = len(v3)
            var v68 uint = 312
            v68 = uint(v10)
            _ = v68
        }
        v69 := GenS1{F0: v10}
        _ = v69
    }
    _ = len(v2)
    switch 32767 {
    case 32767:
        switch ("" + ((("a" + "\n\t") + "foo bar") + "a")) {
        case "foo bar":
            v10 = (v10 / ((v10) | 1))
        }
    }
    if (true || true) {
        v5[4294967295] = v10
    } else {
        if len(v6) > 0 {
            v6[0] = -1
        }
        var v70 int = 128
        _ = v70
    }
    if (true || (true || false)) {
        v11 += v11
        if (255 >= 255) {
            if len(v2) > 0 {
                v2[0] = ("a" + "hello")
            }
            for fi71 := 0; fi71 < 20; fi71++ {
                v11 -= (v11 * v11)
                v72 := GenS8{F0: v1}
                _ = fi71
                _ = v72
            }
        } else {
            fmt.Println(v4, v11)
            switch (("a" + ("a" + "")) + "foo bar") {
            case "a":
                fmt.Println(v11)
            case "\n\t":
                delete(v5, 4294967295)
            case "":
                switch -2147483648 {
                case 7:
                    var v73 int16 = genFn0()
                    _ = v73
                }
            default:
                fmt.Println(v4, v11, v1)
            }
            v74 := GenS4{F0: 244, F1: 100, F2: int32(v10)}
            _ = v74
        }
        var v75 int64 = -62
        _ = v75
    } else {
        if true {
            for fi76 := 0; fi76 < 14; fi76++ {
                for fi77 := 0; fi77 < 14; fi77++ {
                    fmt.Println(fi77, fi76, v4)
                    _ = fi77
                }
                if (1 < 249) {
                    v4 -= 1e10
                    var v78 float64 = 0.1
                    v3[v10] = (v10 * v10)
                    _ = v78
                } else {
                    if true {
                        fmt.Println(v10, v1, v1)
                        var v79 float32 = -2.25
                        _ = v79
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch "foo bar" {
                    case "foo bar":
                        for fi80 := 0; fi80 < 10; fi80++ {
                            var v81 int16 = 182
                            fmt.Println(v1, v11, v11)
                            var v82 uint32 = (^(v10))
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                    }
                }
                v10 = v10
                _ = fi76
            }
        }
        var v83 int = -704
        var v84 int64 = 1
        _ = v83
        _ = v84
    }
    var v85 uint8 = 56
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
