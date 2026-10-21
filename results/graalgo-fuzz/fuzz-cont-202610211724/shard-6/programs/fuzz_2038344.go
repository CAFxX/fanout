package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
type GenS6 struct { F0 int; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int8, p1 int8) int8 {
    v7 := GenS1{F0: -758}
    var v8 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 + p0)
}
func genFn1() uint16 {
    v9 := GenS1{F0: 1}
    var v10 uint32 = 1
    v10 -= 24
    _ = v10
    _ = v9
    return 687
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 2
    var v3 int64 = -432
    var v4 int16 = -32768
    var v5 int8 = -52
    var v6 uint8 = v1
    v11 := GenS1{F0: -9223372036854775808}
    var v12 string = "\n\t"
    v13 := GenS1{F0: -505}
    switch v12 {
    case "":
        fmt.Println(v6, v5, v12)
    case "a":
        var v14 uint = 426
        _ = v14
    default:
        if ((v5 > (v5 | v5)) && true) {
            var v15 float64 = -2.25
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 int64 = 0
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: 42, F1: -2147483648, F2: 1}
            _ = v15
            _ = v18
        }
    }
    v19 := []int{9223372036854775807, 32768}
    switch (v12 + v12) {
    case "a":
        for fi20 := 0; fi20 < 7; fi20++ {
            fmt.Println(v12, v1)
            v3 = (v3 & v3)
            switch fi20 {
            case 42:
                v6 = (v1 ^ v6)
            case -854:
                switch -588 {
                case 32768:
                    v5 += -1
                case 9223372036854775807:
                    switch (fi20 * (fi20 / ((int(v5)) | 1))) {
                    case 4294967295:
                        var v21 uint32 = (^(v2))
                        _ = v21
                    case -9223372036854775808:
                        var v22 uint64 = 904
                        _ = v22
                    }
                }
            case 32767:
                switch "" {
                case "hello":
                    switch 1 {
                    case 4294967295:
                        v12 = v12
                    case 127:
                        v23 := GenS3{F0: 0, F1: 65535, F2: 0}
                        _ = v23
                    }
                case "a":
                    if ((!(69 < (v1 << ((2) & 63)))) || (!true)) {
                        if (!true) {
                            var v24 float32 = 0.1
                            _ = v24
                        } else {
                            v2 = 2147483647
                            var v25 int8 = genFn0(genFn0((v5 * v5), genFn0(v5, v5)), (v5 ^ (-(v5))))
                            _ = v25
                        }
                        if (18446744073709551615 == 2) {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                var v27 uint32 = (v2 * 2)
                                switch (v12 + v12) {
                                case "\n\t":
                                    v1 += v6
                                case "":
                                    v5 = int8(v6)
                                }
                                switch 298 {
                                case -878:
                                    var v28 int32 = 778
                                    _ = v28
                                case 1:
                                    v29 := GenS4{F0: v27, F1: (fi20 * 522)}
                                    _ = v29
                                }
                                _ = fi26
                                _ = v27
                            }
                            switch int(v2) {
                            case 0:
                                switch v12 {
                                case "foo bar":
                                    var v30 uint32 = (^(v2))
                                    _ = v30
                                default:
                                    v2 = 0
                                }
                            case 722:
                                switch v12 {
                                case "\n\t":
                                    var v31 string = "a"
                                    _ = v31
                                case "":
                                    v19 = append(v19, (fi20 - 2147483647))
                                }
                            case 32768:
                                v32 := GenS5{F0: genFn0((v5 >> ((18446744073709551615) & 63)), genFn0(v5, (v5 % ((genFn0(genFn0((v5 / (((-(v5))) | 1)), 127), 42)) | 1)))), F1: (-1 <= v5), F2: (v4 << ((255) & 63))}
                                _ = v32
                            }
                            var v33 uint64 = 10
                            _ = v33
                        } else {
                            fmt.Println(v2)
                        }
                    } else {
                        v34 := GenS4{F0: v2, F1: (fi20 % ((fi20) | 1))}
                        v3 = v3
                        switch "\n\t" {
                        case "\n\t":
                            v35 := GenS6{F0: -380, F1: (^(v6)), F2: 2}
                            _ = v35
                        case "hello":
                            if ((-(v4)) >= (v4 + 1000)) {
                                var v36 int8 = (v5 | v5)
                                v6 -= (v1 | (v6 ^ (v1 | v1)))
                                if ((false && false) && (1 != 0)) {
                                    v37 := GenS3{F0: 42, F1: 0, F2: (v1 + (v1 & (^(v1))))}
                                    _ = v37
                                } else {
                                    if len(v19) > 0 {
                                        v19[0] = (fi20 ^ 127)
                                    }
                                }
                                _ = v36
                            }
                        default:
                            v38 := GenS6{F0: -320, F1: v6, F2: 1}
                            _ = v38
                        }
                        _ = v34
                    }
                case "\n\t":
                    if (!true) {
                        var v39 uint8 = v6
                        _ = v39
                    } else {
                        v4 = (-(v4))
                        if true {
                            fmt.Println(fi20, v6)
                            for fi40 := 0; fi40 < 11; fi40++ {
                                v1 = (v1 << ((227) & 63))
                                v3 += v3
                                v41 := GenS7{F0: (^(v2)), F1: 80, F2: fi40}
                                _ = fi40
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 6; fi42++ {
                                var v43 int = fi20
                                if len(v19) > 0 {
                                    v19[0] = fi20
                                }
                                _ = fi42
                                _ = v43
                            }
                        } else {
                            v44 := GenS5{F0: (-(v5)), F1: true, F2: v4}
                            _ = v44
                        }
                    }
                default:
                    v2 += (v2 - 4294967295)
                }
            default:
                var v45 uint8 = 96
                _ = v45
            }
            _ = fi20
        }
    case "hello":
        var v46 uint16 = 65535
        _ = v46
    }
    v47 := GenS5{F0: (-(v5)), F1: false, F2: (v4 << ((0) & 63))}
    if (1 == 934) {
        switch v12 {
        case "\n\t":
            for fi48 := 0; fi48 < 13; fi48++ {
                switch ((("" + v12) + v12) + v12) {
                case "foo bar":
                    for fi49 := 0; fi49 < 16; fi49++ {
                        switch v12 {
                        case "foo bar":
                            v50 := GenS7{F0: 20, F1: v1, F2: fi49}
                            _ = v50
                        case "hello":
                            _ = len(v19)
                        }
                        v3 -= (v3 * (v3 | v3))
                        var v51 int64 = v3
                        _ = fi49
                        _ = v51
                    }
                case "\n\t":
                    v5 = (-(v5))
                }
                v19 = append(v19, -628)
                _ = fi48
            }
        case "a":
            v52 := GenS3{F0: 42, F1: 961, F2: (v1 | (v6 * 42))}
            _ = v52
        default:
            for fi53 := 0; fi53 < 10; fi53++ {
                for fi54 := 0; fi54 < 6; fi54++ {
                    switch v12 {
                    case "hello":
                        v6 = 1
                    default:
                        v19 = append(v19, (-(fi54)))
                    }
                    v3 -= int64(v4)
                    switch (fi53 / (((-(fi54))) | 1)) {
                    case 1:
                        v2 += 42
                    case 4294967295:
                        v55 := GenS8{F0: (^(v1)), F1: 4294967295}
                        _ = v55
                    case -9223372036854775808:
                        if len(v19) > 0 {
                            v19[0] = (fi53 / (((fi53 - (-(fi54)))) | 1))
                        }
                    }
                    _ = fi54
                }
                _ = fi53
            }
        }
        fmt.Println(v6, v2, v5)
        fmt.Println(v5)
    }
    if (false || false) {
        switch v12 {
        case "hello":
            for fi56 := 0; fi56 < 13; fi56++ {
                v57 := GenS4{F0: 65535, F1: 128}
                if true {
                    var v58 bool = false
                    v59 := GenS9{F0: 42, F1: 255}
                    _ = v58
                    _ = v59
                } else {
                    fmt.Println(v3)
                }
                fmt.Println(v2, v6)
                _ = fi56
                _ = v57
            }
        case "a":
            if len(v19) > 0 {
                v19[0] = 0
            }
        case "\n\t":
            var v60 int32 = 2147483647
            _ = v60
        default:
            var v61 int = 32768
            _ = v61
        }
        if (v4 >= v4) {
            var v62 int32 = 0
            var v63 float64 = 1e10
            v64 := GenS4{F0: v2, F1: 128}
            _ = v62
            _ = v63
            _ = v64
        }
        v65 := GenS10{F0: 423}
        _ = v65
    } else {
        for fi66 := 0; fi66 < 11; fi66++ {
            for fi67 := 0; fi67 < 3; fi67++ {
                var v68 int16 = v4
                _ = len(v19)
                _ = fi67
                _ = v68
            }
            _ = len(v19)
            if ((^(fi66)) >= (fi66 % (((fi66 >> ((18446744073709551615) & 63))) | 1))) {
                if len(v19) > 0 {
                    v19[0] = fi66
                }
                var v69 int8 = (v5 + genFn0(-1, (v5 - v5)))
                if (1 == (^(v5))) {
                    var v70 int8 = v5
                    for fi71 := 0; fi71 < 17; fi71++ {
                        v3 = 964
                        if true {
                            v19 = append(v19, (fi71 / ((-849) | 1)))
                            fmt.Println(v5, v1)
                            v72 := GenS5{F0: 118, F1: ((367 <= 489) || true), F2: v4}
                            _ = v72
                        } else {
                            v73 := GenS11{F0: 368, F1: (v2 + (v2 * (v2 << ((0) & 63))))}
                            var v74 uint64 = 526
                            var v75 uint16 = 554
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        if (!true) {
                            v1 = v6
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            case "a":
                                switch int(v6) {
                                case -599:
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                case 9223372036854775807:
                                    v19 = append(v19, 4294967295)
                                }
                            default:
                                v3 -= (v3 + (v3 & -1))
                            }
                        } else {
                            var v76 bool = (false && true)
                            var v77 float32 = 1e10
                            v2 = v2
                            _ = v76
                            _ = v77
                        }
                        if (false && ((1e10 - float64(float32(0.1))) < -2.25)) {
                            var v78 string = (v12 + "a")
                            v69 -= (v5 + (v5 & (v69 - (v70 % ((v5) | 1)))))
                            var v79 int32 = 0
                            _ = v78
                            _ = v79
                        }
                        var v80 int64 = -1
                        _ = fi71
                        _ = v80
                    }
                    _ = len(v19)
                    _ = v70
                }
                _ = v69
            } else {
                var v81 string = (v12 + ("hello" + (v12 + v12)))
                fmt.Println(v4, v2, fi66)
                _ = v81
            }
            _ = fi66
        }
        var v82 string = (v12 + "foo bar")
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
