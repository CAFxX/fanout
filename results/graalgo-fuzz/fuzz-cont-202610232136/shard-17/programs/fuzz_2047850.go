package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint; F2 int8 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint32; F2 int16 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint64, p2 uint16) int8 {
    fmt.Println(p2, p2)
    v4 := GenS1{F0: -836}
    if (1 == 4294967295) {
        var v5 bool = false
        v6 := GenS2{F0: (true && (127 >= 1)), F1: (p0 * (p0 - (^(p0)))), F2: 2147483647}
        _ = v5
        _ = v6
    } else {
        v7 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: -128}
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS4{F0: 556, F1: p0, F2: 32767}
        var v10 float32 = 0.1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{429}
    var v2 int32 = -2147483648
    var v3 uint64 = 2
    var v12 uint8 = 185
    var v13 float64 = (0.1 + 0.1)
    if false {
        var v14 int8 = 127
        v3 = v3
        _ = v14
    } else {
        v3 += v3
        _ = len(v1)
        if true {
            switch 42 {
            case -9223372036854775808:
                for fi15 := 0; fi15 < 11; fi15++ {
                    fmt.Println(fi15, v13, fi15)
                    v16 := GenS5{F0: 42, F1: 1}
                    v13 = 3.14159
                    _ = fi15
                    _ = v16
                }
            case 0:
                v17 := GenS5{F0: (v12 * v12), F1: (-(v2))}
                _ = v17
            case 16:
                switch -150 {
                case 578:
                    switch -315 {
                    case 42:
                        if (-1 > 9223372036854775807) {
                            v18 := GenS2{F0: false, F1: (v12 / (((v12 % (((v12 * (v12 ^ 128))) | 1))) | 1)), F2: 2147483647}
                            var v19 int32 = -1
                            _ = v18
                            _ = v19
                        }
                    case 128:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            switch (fi20 | fi20) {
                            case 999:
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            case -989:
                                for fi21 := 0; fi21 < 11; fi21++ {
                                    var v22 int16 = 1000
                                    _ = fi21
                                    _ = v22
                                }
                            case 128:
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    _ = len(v1)
                                    v2 = -1
                                    v2 += v2
                                    _ = fi23
                                }
                            default:
                                var v24 uint64 = (v3 * (^(v3)))
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v26 := GenS1{F0: 0}
                                var v27 float64 = (v13 * float64(1e10))
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            _ = fi20
                        }
                    case 127:
                        for fi28 := 0; fi28 < 5; fi28++ {
                            var v29 bool = (!false)
                            _ = fi28
                            _ = v29
                        }
                    default:
                        for fi30 := 0; fi30 < 17; fi30++ {
                            switch (("foo bar" + "\n\t") + "hello") {
                            case "":
                                for fi31 := 0; fi31 < 15; fi31++ {
                                    v32 := GenS6{F0: (((fi31 * fi30) == fi30) || (1000 == 0)), F1: (!false)}
                                    switch "foo bar" {
                                    case "hello":
                                        var v33 float32 = 1.5
                                        if false {
                                            if (42 < 2147483648) {
                                                v33 = v33
                                                if (v2 <= v2) {
                                                    v34 := GenS7{F0: (v2 - (v2 * (v2 - v2)))}
                                                    _ = v34
                                                } else {
                                                    var v35 float32 = float32(v13)
                                                    var v36 uint64 = (v3 & v3)
                                                    var v37 int8 = 58
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            } else {
                                                var v38 uint64 = (v3 | uint64(v12))
                                                _ = v38
                                            }
                                            var v39 uint = 206
                                            _ = v39
                                        } else {
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                            }
                                            var v40 uint8 = v12
                                            v41 := GenS7{F0: (v2 * (v2 << ((1) & 63)))}
                                            var v42 int16 = 32767
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = v33
                                    }
                                    v12 += v12
                                    v2 -= v2
                                    v13 = -2.25
                                    _ = fi31
                                    _ = v32
                                }
                            case "foo bar":
                                v12 = v12
                            default:
                                for fi43 := 0; fi43 < 13; fi43++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v12 = 210
                                    _ = fi43
                                }
                            }
                            _ = fi30
                        }
                    }
                case 1:
                    v3 += v3
                }
            }
            v12 = (^(v12))
        }
    }
    v44 := GenS4{F0: 42, F1: v12, F2: -161}
    if false {
        if (!false) {
            switch -144 {
            case -2147483648:
                var v45 int8 = 77
                _ = v45
            case -272:
                v46 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: -10}
                _ = v46
            default:
                var v47 int32 = (v2 ^ v2)
                _ = v47
            }
            switch "\n\t" {
            case "hello":
                var v48 uint64 = (v3 / ((2) | 1))
                _ = v48
            case "foo bar":
                if false {
                    var v49 int64 = -678
                    switch "hello" {
                    case "":
                        for fi50 := 0; fi50 < 7; fi50++ {
                            var v51 int16 = -951
                            fmt.Println(v49, v51)
                            switch "foo bar" {
                            case "":
                                if ((v3 + (v3 * v3)) >= (^(v3))) {
                                    v1 = append(v1, 48)
                                } else {
                                    v52 := GenS8{F0: fi50, F1: 2147483648, F2: v51}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v52
                                }
                            default:
                                var v53 int16 = (v51 & (v51 | 32767))
                                _ = v53
                            }
                            _ = fi50
                            _ = v51
                        }
                    case "foo bar":
                        var v54 uint64 = v3
                        _ = v54
                    case "hello":
                        v12 += (v12 ^ uint8(v3))
                    }
                    if (0 < 65535) {
                        var v55 float32 = 0.0
                        var v56 uint = 1
                        _ = v55
                        _ = v56
                    }
                    _ = v49
                } else {
                    var v57 float64 = (0.1 + (((v13 * float64((((-2.25 + ((float32(1.5) + 1e10) + float32(v13))) + ((0.1 - (0.0 + ((-2.25 - 3.14159) / (1e10 / ((1.5 + 0.0) * ((3.14159 * 0.1) + 3.14159)))))) * 1.5)) / (float32(0.0) - -2.25)))) - 0.0) * (v13 + v13)))
                    v58 := GenS1{F0: (v2 + v2)}
                    v59 := GenS9{F0: 1, F1: -488}
                    _ = v57
                    _ = v58
                    _ = v59
                }
            case "":
                v60 := GenS4{F0: 32767, F1: v12, F2: 32767}
                _ = v60
            default:
                fmt.Println(v3)
            }
            for fi61 := 0; fi61 < 9; fi61++ {
                for fi62 := 0; fi62 < 20; fi62++ {
                    fmt.Println(fi62, v2)
                    v63 := GenS10{F0: (v2 >> ((4294967295) & 63)), F1: 65535}
                    _ = fi62
                    _ = v63
                }
                var v64 float64 = (v13 * v13)
                for fi65 := 0; fi65 < 4; fi65++ {
                    fmt.Println(v13, v3)
                    if true {
                        v3 += 1
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        var v66 int16 = 1
                        _ = v66
                    }
                    v2 = v2
                    _ = fi65
                }
                _ = fi61
                _ = v64
            }
        } else {
            for fi67 := 0; fi67 < 6; fi67++ {
                if false {
                    fmt.Println(fi67, v2, v2)
                } else {
                    var v68 string = "\n\t"
                    _ = v68
                }
                _ = fi67
            }
            var v69 bool = (!true)
            v2 += (v2 / ((0) | 1))
            _ = v69
        }
    }
    var v70 string = ""
    var v71 float32 = (-2.25 - 0.0)
    switch 127 {
    case 2147483647:
        switch "a" {
        case "hello":
            v12 = uint8(v3)
        case "foo bar":
            var v72 uint = 65535
            _ = v72
        default:
            var v73 int8 = 42
            _ = v73
        }
    case 4294967295:
        switch 2147483647 {
        case 1:
            for fi74 := 0; fi74 < 18; fi74++ {
                if true {
                    for fi75 := 0; fi75 < 19; fi75++ {
                        var v76 uint16 = 32767
                        _ = fi75
                        _ = v76
                    }
                    var v77 int32 = -509
                    switch fi74 {
                    case 5:
                        switch ("hello" + "\n\t") {
                        case "hello":
                            v2 = v77
                        case "a":
                            switch v70 {
                            case "":
                                v78 := GenS11{F0: (v2 - 2147483647)}
                                _ = v78
                            }
                        case "\n\t":
                            var v79 uint16 = 32767
                            _ = v79
                        default:
                            if false {
                                var v80 int32 = (^(v77))
                                var v81 int16 = 0
                                v82 := GenS2{F0: true, F1: (v12 | (v12 & (^(v12)))), F2: 2147483647}
                                _ = v80
                                _ = v81
                                _ = v82
                            } else {
                                var v83 bool = true
                                var v84 bool = true
                                var v85 float32 = v71
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                        }
                    case 128:
                        _ = len(v1)
                    case 1:
                        switch (v70 + (v70 + "foo bar")) {
                        case "\n\t":
                            var v86 int32 = int32(fi74)
                            _ = v86
                        }
                    default:
                        switch (fi74 << ((2) & 63)) {
                        case 800:
                            fmt.Println(v70)
                        case -719:
                            if (-2.25 != 0.1) {
                                var v87 uint = uint(v12)
                                _ = v87
                            } else {
                                v3 = 873
                            }
                        case -792:
                            for fi88 := 0; fi88 < 18; fi88++ {
                                switch fi74 {
                                case 2147483647:
                                    var v89 int32 = (^(v77))
                                    _ = v89
                                }
                                _ = fi88
                            }
                        default:
                            var v90 int8 = 36
                            _ = v90
                        }
                    }
                    _ = v77
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if ((v2 & (v2 + v2)) <= v2) {
                    fmt.Println(v71, v71, fi74)
                }
                _ = fi74
            }
        case 248:
            if false {
                var v91 int64 = 42
                v13 += v13
                _ = v91
            }
        case 128:
            v1 = append(v1, 65535)
        default:
            if false {
                v13 -= (1.5 - v13)
            }
        }
    case 1:
        var v92 float32 = v71
        _ = v92
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v93 int16 = 42
    var v94 uint32 = 1
    v12 += (v12 % ((v12) | 1))
    v13 = (float64(v71) - v13)
    v1 = append(v1, v94)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v44
    _ = v70
    _ = v71
    _ = v93
    _ = v94
}
