package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16; F1 int; F2 int }
type GenS4 struct { F0 int64; F1 bool }

func main() {
    var v1 string = (("" + "a") + "")
    var v2 uint16 = 0
    var v3 uint64 = 90
    var v4 uint16 = (^(v2))
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: -32768, F1: -2147483648, F2: 1000}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int16 = 42
        switch (v1 + ("" + v1)) {
        case "\n\t":
            var v9 uint8 = 0
            _ = v9
        case "a":
            if (((1.5 * 1e10) + float32((3.14159 * (-2.25 - 1.5)))) > float32(1.5)) {
                var v10 float64 = -2.25
                _ = v10
            }
        case "":
            var v11 float32 = (1e10 / 1e10)
            _ = v11
        }
        v2 -= 32767
        _ = fi7
        _ = v8
    }
    if false {
        var v12 uint = 1
        _ = v12
    } else {
        fmt.Println(v2, v3, v3)
    }
    var v13 uint32 = 4294967295
    var v14 uint64 = v3
    v3 -= 42
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float64 = 1e10
            fmt.Println(v14, v4, v1)
            fmt.Println(fi15, v4)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := []int{-240, -9223372036854775808, 157}
    if ((127 == 42) && ((-1 < 211) && true)) {
        if true {
            v4 = (v2 - v2)
            var v19 int64 = -305
            var v20 uint64 = v3
            _ = v19
            _ = v20
        } else {
            _ = len(v5)
            v5[(v13 | uint32(v2))] = 0
        }
        if false {
            var v21 float64 = (((1e10 / 1.5) / 1.5) / ((0.1 * 0.0) + 3.14159))
            _ = v21
        } else {
            var v22 uint32 = 2147483648
            fmt.Println(v3, v22)
            var v23 float32 = 3.14159
            _ = v22
            _ = v23
        }
        v13 = (v13 & v13)
    }
    if (!(0 == 0)) {
        for fi24 := 0; fi24 < 18; fi24++ {
            v2 = v2
            v25 := GenS1{F0: 42, F1: -659, F2: -457}
            v5[v13] = (fi24 & -2147483648)
            _ = fi24
            _ = v25
        }
        switch 42 {
        case 32767:
            var v26 uint64 = (v14 / ((v3) | 1))
            _ = v26
        case 128:
            v18 = append(v18, 870)
        case 265:
            switch 4294967295 {
            case -232:
                if true {
                    var v27 float64 = (((1e10 * 0.0) + -2.25) - (-2.25 * 0.1))
                    var v28 int32 = -1
                    if ((v4 | v4) >= v2) {
                        if (!(((3.14159 + (3.14159 + 0.1)) * -2.25) != 0.0)) {
                            if true {
                                v2 = 0
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    var v30 float32 = -2.25
                                    _ = fi29
                                    _ = v30
                                }
                            }
                            if len(v18) > 0 {
                                v18[0] = 42
                            }
                            switch v1 {
                            case "a":
                                if true {
                                    if true {
                                        var v31 float64 = -2.25
                                        _ = v31
                                    } else {
                                        if false {
                                            if false {
                                                v32 := GenS2{F0: (v3 ^ (v14 << ((4294967295) & 63))), F1: 0}
                                                var v33 uint8 = 64
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        switch -2 {
                                        case 632:
                                            if (v28 < 697) {
                                                v27 = ((v27 - ((((v27 / (v27 - v27)) / (float64(3.14159) - (v27 / v27))) / -2.25) / 1e10)) + ((v27 / v27) / (v27 + float64((0.1 + -2.25)))))
                                            } else {
                                                v34 := GenS2{F0: (v3 << ((255) & 63)), F1: v28}
                                                _ = v34
                                            }
                                        case 2147483647:
                                            v1 = (v1 + "\n\t")
                                        case 424:
                                            fmt.Println(v13, v1, v1)
                                        }
                                        v13 -= v13
                                    }
                                    v35 := GenS2{F0: (v14 / (((v14 | 293)) | 1)), F1: v28}
                                    fmt.Println(v28)
                                    _ = v35
                                } else {
                                    for fi36 := 0; fi36 < 18; fi36++ {
                                        var v37 uint32 = v13
                                        var v38 int8 = 100
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                    if len(v18) > 0 {
                                        v18[0] = -9223372036854775808
                                    }
                                    for fi39 := 0; fi39 < 5; fi39++ {
                                        fmt.Println(v14, v1, v3)
                                        _ = fi39
                                    }
                                }
                            default:
                                switch 2147483647 {
                                case 2147483647:
                                    switch v1 {
                                    case "foo bar":
                                        delete(v5, (^(v13)))
                                    default:
                                        var v40 uint64 = (^(v14))
                                        _ = v40
                                    }
                                case 4294967295:
                                    switch 1 {
                                    case 1:
                                        var v41 uint8 = 0
                                        _ = v41
                                    case -902:
                                        fmt.Println(v4, v13)
                                    case 32768:
                                        var v42 string = "\n\t"
                                        _ = v42
                                    default:
                                        if false {
                                            if (-1 == 127) {
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                var v43 int64 = 917
                                                _ = v43
                                            }
                                        } else {
                                            fmt.Println(v1, v28, v1)
                                        }
                                    }
                                default:
                                    v14 = (v14 - v14)
                                }
                            }
                        }
                        var v44 uint16 = (v2 << ((18446744073709551615) & 63))
                        _ = v44
                    }
                    _ = v27
                    _ = v28
                } else {
                    v3 -= 251
                    _ = len(v5)
                }
            }
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            if false {
                v5[v13] = (fi45 + -1)
                fmt.Println(v14)
                if (-1 == -290) {
                    v3 = (v14 * 1)
                    v13 -= 209
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                }
            }
            var v46 int = fi45
            var v47 int64 = 9223372036854775807
            _ = fi45
            _ = v46
            _ = v47
        }
    } else {
        if (!(75 > uint8(v13))) {
            v48 := v5[uint32(v3)]
            if true {
                v2 -= 0
            }
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        v14 -= (v14 ^ v3)
        var v50 uint64 = (^(v14))
        _ = fi49
        _ = v50
    }
    v13 -= (^(v13))
    for fi51 := 0; fi51 < 2; fi51++ {
        var v52 int = 407
        if false {
            v2 -= (v4 + v2)
            fmt.Println(v4, v14)
            switch "\n\t" {
            case "\n\t":
                var v53 uint32 = (v13 % (((v13 >> ((408) & 63))) | 1))
                _ = v53
            case "hello":
                v14 = 18446744073709551615
            }
        } else {
            switch (v1 + v1) {
            case "hello":
                v2 -= (v4 * v2)
            case "":
                if (-2.25 != (0.1 / 3.14159)) {
                    var v54 uint16 = (v4 ^ v2)
                    v55 := GenS3{F0: -558, F1: fi51, F2: (-(fi51))}
                    var v56 float64 = ((((0.0 * 0.1) - 1.5) + ((1e10 * 0.1) + 1.5)) * -2.25)
                    _ = v54
                    _ = v55
                    _ = v56
                } else {
                    delete(v5, 4294967295)
                    v3 = v3
                }
            }
            if ((false || true) && (true || false)) {
                var v57 int16 = 32767
                v58 := GenS2{F0: (^(v14)), F1: -2147483648}
                _ = v57
                _ = v58
            }
            var v59 int8 = 42
            _ = v59
        }
        v1 = v1
        _ = fi51
        _ = v52
    }
    for fi60 := 0; fi60 < 16; fi60++ {
        v2 = 42
        if (-9223372036854775808 > fi60) {
            if (true || false) {
                if (0 != 1) {
                    if len(v18) > 0 {
                        v18[0] = -1
                    }
                    for fi61 := 0; fi61 < 7; fi61++ {
                        v62 := GenS3{F0: 0, F1: fi61, F2: -1}
                        v18 = append(v18, fi60)
                        _ = fi61
                        _ = v62
                    }
                    v5[(v13 << ((uint(v3)) & 63))] = 701
                } else {
                    var v63 uint64 = (v3 * 628)
                    var v64 int64 = 1
                    _ = v63
                    _ = v64
                }
                if ((v4 % ((v2) | 1)) < v4) {
                    v3 = (^(v14))
                } else {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        switch -980 {
                        case 1:
                            v13 += 1
                        }
                        _ = fi65
                    }
                    if (!((v2 - (^(v4))) == 938)) {
                        fmt.Println(v1, v4)
                    } else {
                        var v66 string = "\n\t"
                        _ = v66
                    }
                    v1 = (("\n\t" + v1) + (v1 + v1))
                }
            } else {
                if true {
                    for fi67 := 0; fi67 < 15; fi67++ {
                        var v68 uint8 = 42
                        v13 -= v13
                        _ = fi67
                        _ = v68
                    }
                    if ((!((fi60 << ((182) & 63)) == fi60)) && false) {
                        switch (-(fi60)) {
                        case 128:
                            var v69 uint = 4294967295
                            _ = v69
                        }
                    } else {
                        v70 := v5[v13]
                        var v71 int8 = -122
                        _ = v70
                        _ = v71
                    }
                    v2 -= uint16(v14)
                }
                var v72 bool = (0.1 < float64(0.0))
                for fi73 := 0; fi73 < 10; fi73++ {
                    _ = len(v18)
                    v3 -= v14
                    v74 := GenS4{F0: 1, F1: v72}
                    v75 := GenS2{F0: 0, F1: 955}
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                _ = v72
            }
        } else {
            var v76 uint64 = (v14 ^ (v14 * v14))
            var v77 uint = 255
            switch fi60 {
            case -698:
                v77 = 1
            }
            _ = v76
            _ = v77
        }
        switch v1 {
        case "hello":
            var v78 uint32 = v13
            _ = v78
        }
        _ = fi60
    }
    var v79 string = (v1 + v1)
    delete(v5, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
