package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint32; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64) int {
    var v6 uint64 = 2
    p0 += (p0 << ((1) & 63))
    fmt.Println(p0)
    v7 := GenS1{F0: 255, F1: 1, F2: 100}
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 uint16 = 334
    for fi8 := 0; fi8 < 8; fi8++ {
        v5 -= 0
        v5 = 65535
        fmt.Println(fi8, v5)
        _ = fi8
    }
    v9 := []int{}
    var v10 bool = true
    var v11 float32 = -2.25
    switch "\n\t" {
    case "\n\t":
        if v10 {
            v5 = (v5 ^ 32767)
            var v12 int64 = 822
            _ = v12
        } else {
            v9 = append(v9, 1)
        }
    case "foo bar":
        var v13 int64 = 1
        _ = v13
    case "":
        if (-9223372036854775808 != -918) {
            for fi14 := 0; fi14 < 13; fi14++ {
                if ((42 == 42) && v10) {
                    v5 += 531
                    v1[(fi14 * 9223372036854775807)] = (-(fi14))
                } else {
                    v10 = (v5 < 0)
                    fmt.Println(fi14, v5, v5)
                    for fi15 := 0; fi15 < 8; fi15++ {
                        switch "hello" {
                        case "hello":
                            var v16 int32 = 116
                            switch int(v16) {
                            case -794:
                                var v17 int8 = -1
                                fmt.Println(fi15, v11, v17)
                                v18 := GenS1{F0: 528, F1: 130, F2: v17}
                                _ = v17
                                _ = v18
                            case 0:
                                v19 := GenS1{F0: 942, F1: 0, F2: 1}
                                _ = v19
                            default:
                                if true {
                                    if (-1 <= 1) {
                                        v20 := GenS1{F0: 42, F1: 42, F2: 100}
                                        if false {
                                            v10 = v10
                                            v5 = (v5 >> ((2) & 63))
                                            fmt.Println(fi14, fi14, v16)
                                            fmt.Println(v5, fi15, fi14)
                                        } else {
                                            v21 := GenS2{F0: 518}
                                            _ = v21
                                        }
                                        fmt.Println(v16, v11, v16)
                                        _ = v20
                                    }
                                } else {
                                    switch (("" + "hello") + "\n\t") {
                                    case "foo bar":
                                        v16 += (v16 * 1)
                                    }
                                    v10 = ((^(v5)) != v5)
                                }
                            }
                            _ = v16
                        case "":
                            fmt.Println(fi14)
                        default:
                            var v22 int32 = -506
                            _ = v22
                        }
                        if (v5 >= 0) {
                            if false {
                                switch ("" + ("\n\t" + "")) {
                                case "hello":
                                    if (v11 >= 0.1) {
                                        fmt.Println(fi14, v11)
                                        switch (fi15 - (fi15 - genFn0(792))) {
                                        case 127:
                                            v23 := GenS2{F0: 65535}
                                            _ = v23
                                        case 32767:
                                            v5 = (v5 | v5)
                                        case -9223372036854775808:
                                            v24 := GenS2{F0: 280}
                                            if (((float32(0.1) / float32(1e10)) * (v11 / v11)) < v11) {
                                                v10 = v10
                                            } else {
                                                v25 := GenS2{F0: 65535}
                                                var v26 float32 = (v11 / (v11 - v11))
                                                if (v10 && v10) {
                                                    if (!(v26 <= v11)) {
                                                        if (0.1 != 0.1) {
                                                            var v27 uint = 220
                                                            var v28 uint = (v27 | 4294967295)
                                                            v29 := GenS3{F0: 212, F1: 127}
                                                            delete(v3, ("a" + "\n\t"))
                                                            v5 = 42
                                                            _ = v27
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        if v10 {
                                                            v30 := GenS4{F0: 1}
                                                            _ = v30
                                                        }
                                                    }
                                                    if true {
                                                        v31 := GenS5{F0: 127, F1: 0, F2: 2147483647}
                                                        if v10 {
                                                            v32 := GenS6{F0: 959}
                                                            switch fi14 {
                                                            case 2147483647:
                                                                fmt.Println(v11, v5, fi15)
                                                            case -1:
                                                                var v33 uint64 = 1
                                                                _ = v33
                                                            case -2147483648:
                                                                v5 += v5
                                                            default:
                                                                _ = len(v9)
                                                            }
                                                            _ = v32
                                                        } else {
                                                            var v34 uint16 = (^(v5))
                                                            var v35 uint = 747
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        if v10 {
                                                            if (!v10) {
                                                                var v36 int64 = 0
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                                switch (fi15 | 387) {
                                                                case 0:
                                                                    var v37 uint64 = 42
                                                                    _ = v37
                                                                case -1:
                                                                    fmt.Println(v11, fi15)
                                                                }
                                                                _ = v36
                                                            }
                                                            if v10 {
                                                                if ((fi14 <= (-(fi15))) && (!(947 == 0))) {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v38 int16 = 1
                                                                        _ = v38
                                                                    case "":
                                                                        v39 := v1[(fi15 / ((fi15) | 1))]
                                                                        _ = v39
                                                                    case "a":
                                                                        if v10 {
                                                                            switch fi14 {
                                                                            case -9223372036854775808:
                                                                                if v10 {
                                                                                    switch fi15 {
                                                                                    case -973:
                                                                                        var v40 int16 = 1000
                                                                                        v26 -= v26
                                                                                        v10 = v10
                                                                                        v10 = (1 > int16(v5))
                                                                                        v26 = float32(0.1)
                                                                                        fmt.Println(fi14, v11)
                                                                                        _ = v40
                                                                                    }
                                                                                    v5 -= (v5 / ((255) | 1))
                                                                                }
                                                                            default:
                                                                                v41 := GenS7{F0: 128, F1: false}
                                                                                _ = v41
                                                                            }
                                                                        }
                                                                        v5 = (^(v5))
                                                                    default:
                                                                        v5 = (v5 % ((v5) | 1))
                                                                    }
                                                                }
                                                                v9 = append(v9, fi15)
                                                                fmt.Println(fi15, v10, v26)
                                                            } else {
                                                                if (v10 || (!((v5 * (v5 * (v5 * v5))) > (v5 >> ((18446744073709551615) & 63))))) {
                                                                    var v42 string = ""
                                                                    v11 += (v11 * v26)
                                                                    v43 := GenS8{F0: 9223372036854775807, F1: -32768}
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            }
                                                        } else {
                                                            v9 = append(v9, 32768)
                                                            v26 -= v26
                                                        }
                                                        fmt.Println(v26)
                                                        fmt.Println(v11, fi14, v5)
                                                        _ = len(v9)
                                                        _ = v31
                                                    }
                                                    if true {
                                                        var v44 int64 = 9223372036854775807
                                                        fmt.Println(fi15, v10)
                                                        _ = v44
                                                    }
                                                } else {
                                                    _ = len(v9)
                                                    v45 := GenS3{F0: 255, F1: -128}
                                                    v46 := GenS7{F0: 127, F1: true}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                _ = v25
                                                _ = v26
                                            }
                                            if (v10 && true) {
                                                if false {
                                                    var v47 string = "\n\t"
                                                    v48 := GenS7{F0: 10, F1: (1.5 == (0.0 + (-2.25 / 1.5)))}
                                                    v10 = (v5 > (v5 % ((0) | 1)))
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    v10 = (!true)
                                                    if true {
                                                        var v49 int16 = 1000
                                                        _ = v49
                                                    } else {
                                                        fmt.Println(fi15, v10)
                                                    }
                                                    v10 = (241 >= 2)
                                                }
                                                v10 = true
                                                var v50 int8 = -128
                                                _ = v50
                                            } else {
                                                if false {
                                                    v10 = (float64(0.1) == 1.5)
                                                    fmt.Println(v10, v10)
                                                    if (1 >= 0) {
                                                        v51 := GenS7{F0: 42, F1: false}
                                                        _ = v51
                                                    }
                                                }
                                                var v52 uint8 = 42
                                                var v53 int64 = 0
                                                v52 = 146
                                                var v54 uint = 0
                                                var v55 int8 = -93
                                                v53 = (v53 ^ (v53 >> (((v54 * (v54 << (((v54 + v54)) & 63)))) & 63)))
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            fmt.Println(fi15)
                                            if ((v10 && (-753 <= -9223372036854775808)) && (!(v10 && (-32768 == 42)))) {
                                                delete(v1, 1)
                                            } else {
                                                fmt.Println(v5, v10, fi15)
                                                _ = len(v9)
                                                fmt.Println(v10, v10, fi14)
                                            }
                                            _ = v24
                                        default:
                                            var v56 int = (fi15 + (fi15 & genFn0(18446744073709551615)))
                                            v57 := GenS1{F0: 42, F1: 2, F2: -99}
                                            _ = v56
                                            _ = v57
                                        }
                                    } else {
                                        v10 = (v10 || ("hello" < (("foo bar" + "a") + "foo bar")))
                                        var v58 int16 = -925
                                        _ = len(v9)
                                        _ = v58
                                    }
                                case "":
                                    var v59 int32 = 2147483647
                                    _ = v59
                                case "\n\t":
                                    if (0 > 42) {
                                        v11 = float32((1.5 * 0.0))
                                        v60 := GenS9{F0: 177, F1: 0, F2: (v5 * 459)}
                                        var v61 int8 = 39
                                        var v62 int64 = 9223372036854775807
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        v63 := GenS10{F0: 554}
                                        v2[130] = true
                                        _ = v63
                                    }
                                default:
                                    var v64 uint64 = 42
                                    _ = v64
                                }
                            } else {
                                v11 = 0.0
                                if len(v9) > 0 {
                                    v9[0] = fi14
                                }
                                fmt.Println(fi14)
                            }
                            var v65 uint8 = 47
                            _ = v65
                        } else {
                            v66 := GenS5{F0: -128, F1: 280, F2: int32(v5)}
                            v2[685] = v10
                            v67 := v2[404]
                            _ = v66
                            _ = v67
                        }
                        if v10 {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v68 := GenS3{F0: 1, F1: 1}
                            if ("\n\t" != "\n\t") {
                                var v69 uint8 = 42
                                delete(v2, 1)
                                fmt.Println(v10, fi14, v10)
                                _ = v69
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    if v10 {
                                        switch (fi15 % ((genFn0(448)) | 1)) {
                                        case 533:
                                            v10 = v10
                                        default:
                                            var v70 uint8 = 2
                                            _ = v70
                                        }
                                    }
                                case "foo bar":
                                    v9 = append(v9, (-(fi15)))
                                    var v71 bool = (v5 < 122)
                                    v71 = (42 < -1)
                                    _ = v71
                                }
                                v10 = (1 != 42)
                            }
                            _ = v68
                        }
                        var v72 int8 = -116
                        v9 = append(v9, (^(fi15)))
                        if v10 {
                            var v73 float64 = -2.25
                            _ = v73
                        } else {
                            var v74 int32 = int32(v72)
                            _ = v74
                        }
                        _ = fi15
                        _ = v72
                    }
                }
                v11 -= v11
                _ = fi14
            }
        } else {
            v10 = v10
            for fi75 := 0; fi75 < 12; fi75++ {
                v10 = (65535 >= 42)
                fmt.Println(v10, v11, fi75)
                if v10 {
                    var v76 string = ("hello" + "\n\t")
                    _ = len(v1)
                    var v77 uint64 = 784
                    _ = v76
                    _ = v77
                }
                _ = fi75
            }
            fmt.Println(v10)
        }
    }
    if (v10 || ("foo bar" != (("\n\t" + "") + ("a" + ("\n\t" + ("\n\t" + ("foo bar" + ""))))))) {
        v11 = v11
        for fi78 := 0; fi78 < 16; fi78++ {
            switch 703 {
            case 127:
                var v79 uint8 = 255
                _ = v79
            default:
                fmt.Println(fi78, v11, v5)
            }
            v10 = (v5 != (^(v5)))
            _ = fi78
        }
    }
    for fi80 := 0; fi80 < 7; fi80++ {
        fmt.Println(v5)
        v4[("" + "foo bar")] = 65535
        fmt.Println(v11)
        _ = fi80
    }
    fmt.Println(v10, v10)
    v5 += 2
    var v81 uint = 530
    switch "" {
    case "hello":
        for fi82 := 0; fi82 < 8; fi82++ {
            v5 += 2
            if len(v9) > 0 {
                v9[0] = (-(fi82))
            }
            _ = fi82
        }
    }
    for fi83 := 0; fi83 < 7; fi83++ {
        if v10 {
            var v84 int16 = 1
            fmt.Println(v81, fi83, v84)
            v5 -= v5
            _ = v84
        } else {
            if v10 {
                var v85 int32 = 42
                v81 = 255
                for fi86 := 0; fi86 < 18; fi86++ {
                    v87 := GenS11{F0: fi86, F1: (v81 >> ((v81) & 63))}
                    if len(v9) > 0 {
                        v9[0] = fi83
                    }
                    _ = fi86
                    _ = v87
                }
                _ = v85
            } else {
                if ((v5 - 1) == (v5 * 0)) {
                    fmt.Println(v10, v11, v11)
                }
                if true {
                    var v88 uint16 = 252
                    if (v10 || false) {
                        v81 += (^(v81))
                        v89 := GenS12{F0: 255, F1: 42}
                        _ = v89
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v88
                } else {
                    fmt.Println(fi83, fi83, v11)
                    v5 += (v5 >> ((v81) & 63))
                }
            }
        }
        var v90 uint8 = 127
        switch fi83 {
        case 882:
            fmt.Println(v10, fi83, v11)
        case 0:
            if false {
                _ = len(v2)
                v91 := GenS7{F0: v90, F1: false}
                for fi92 := 0; fi92 < 1; fi92++ {
                    if false {
                        var v93 int32 = 0
                        v90 = v90
                        _ = v93
                    }
                    switch 128 {
                    case -440:
                        v11 += 0.1
                    case -460:
                        fmt.Println(fi83)
                    }
                    v94 := GenS7{F0: v90, F1: (fi92 <= -17)}
                    _ = fi92
                    _ = v94
                }
                _ = v91
            }
        case 128:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = fi83
        _ = v90
    }
    for fi95 := 0; fi95 < 12; fi95++ {
        var v96 int8 = 127
        _ = fi95
        _ = v96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v9
}
