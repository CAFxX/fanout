package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 bool; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 int8; F2 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 uint16 }
func genFn0() int16 {
    if false {
        v5 := GenS1{F0: 0, F1: 127}
        _ = v5
    } else {
        v6 := []int{32767}
        var v7 uint32 = 2
        v7 = (^(v7))
        _ = v6
        _ = v7
    }
    var v8 uint = 2
    var v9 uint = 929
    _ = v8
    _ = v9
    return -488
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 42
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v10 uint8 = 42
    v10 -= (v10 + (v10 | 0))
    if (1 >= 2) {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: 738, F1: v10}
            v2 += (v2 << ((828) & 63))
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v10, v2, fi13)
            if false {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v2 = v2
                    _ = fi14
                }
            } else {
                switch -185 {
                case 400:
                    for fi15 := 0; fi15 < 1; fi15++ {
                        switch ("\n\t" + ("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + "foo bar")))) {
                        case "":
                            fmt.Println(v10, fi15, v2)
                        case "a":
                            v2 -= v2
                        default:
                            fmt.Println(fi15, v10)
                        }
                        var v16 float32 = (0.0 + 1e10)
                        _ = fi15
                        _ = v16
                    }
                case 32767:
                    if true {
                        var v17 int = (fi13 & -373)
                        _ = v17
                    } else {
                        v10 = 169
                        switch fi13 {
                        case 32767:
                            if ((fi13 << ((4294967295) & 63)) != (fi13 / (((fi13 / (((fi13 >> ((255) & 63))) | 1))) | 1))) {
                                var v18 uint16 = 32767
                                fmt.Println(v18)
                                _ = v18
                            }
                        default:
                            v19 := []int{(fi13 >> ((65535) & 63)), 1}
                            _ = v19
                        }
                        v20 := GenS1{F0: 42, F1: v10}
                        _ = v20
                    }
                default:
                    var v21 int32 = -2147483648
                    _ = v21
                }
                var v22 bool = false
                var v23 float64 = (float64((float32(0.1) * (0.0 / 3.14159))) * (-2.25 - 3.14159))
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v10 += (v10 ^ v10)
                var v25 string = ("a" + ("a" + (("hello" + "") + "foo bar")))
                v26 := GenS1{F0: 0, F1: 67}
                if false {
                    v27 := GenS2{F0: true}
                    var v28 uint8 = v10
                    v28 += v10
                    _ = v27
                    _ = v28
                } else {
                    if true {
                        v29 := []int{(fi13 + fi24)}
                        _ = v29
                    }
                    fmt.Println(v2)
                }
                v30 := []int{int(v10)}
                _ = fi24
                _ = v25
                _ = v26
                _ = v30
            }
            _ = fi13
        }
    }
    fmt.Println(v10, v10)
    v31 := []int{-760}
    v10 -= (^(v10))
    var v32 int8 = v2
    v33 := GenS2{F0: (2147483647 == -2147483648)}
    var v34 uint = 0
    if false {
        var v35 int16 = 72
        var v36 int32 = 42
        _ = v35
        _ = v36
    }
    switch 1 {
    case 582:
        v2 = (v2 >> (((v34 | (v34 << ((v34) & 63)))) & 63))
    default:
        v37 := GenS1{F0: int16(v34), F1: uint8(v34)}
        _ = v37
    }
    if (!false) {
        if ((875 >= -816) && false) {
            v38 := GenS2{F0: ((3.14159 + float32(float64(0.1))) < (0.1 * (1.5 / 0.1)))}
            if ((v10 >> (((v34 ^ 2)) & 63)) > (v10 % ((v10) | 1))) {
                if (32767 == 32767) {
                    v34 += (v34 ^ (v34 ^ (v34 / (((v34 * v34)) | 1))))
                    v10 = v10
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if (0.1 < -2.25) {
                            if (!(true || ((false && false) || true))) {
                                for fi40 := 0; fi40 < 8; fi40++ {
                                    var v41 int = fi40
                                    v34 = 0
                                    v42 := GenS3{F0: 42, F1: 51, F2: 0}
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            } else {
                                var v43 uint = 255
                                for fi44 := 0; fi44 < 20; fi44++ {
                                    switch (fi39 | (fi44 | (fi44 >> ((v43) & 63)))) {
                                    case 4294967295:
                                        fmt.Println(v34, v2, fi39)
                                    case 9223372036854775807:
                                        switch fi39 {
                                        case -9223372036854775808:
                                            if (float32((float64(1e10) - (0.1 / 0.1))) < 1e10) {
                                                var v45 int64 = 42
                                                if (true || (true || true)) {
                                                    fmt.Println(v45, v10, v45)
                                                    var v46 int16 = genFn0()
                                                    v47 := GenS3{F0: 32767, F1: 255, F2: 741}
                                                    _ = len(v1)
                                                    var v48 float64 = -2.25
                                                    var v49 float32 = (3.14159 / 3.14159)
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    if len(v31) > 0 {
                                                        v31[0] = -34
                                                    }
                                                    v32 -= (-(v32))
                                                    var v50 int16 = -1
                                                    _ = v50
                                                }
                                                switch "a" {
                                                case "a":
                                                    var v51 uint32 = 42
                                                    _ = v51
                                                case "foo bar":
                                                    var v52 bool = (-2.25 == 0.0)
                                                    _ = v52
                                                case "\n\t":
                                                    if (false || ((((0.0 * (((0.0 - (float64(-2.25) / ((1.5 + float64((3.14159 / float32(3.14159)))) / -2.25))) - (1.5 + (0.0 * 3.14159))) + -2.25)) - float64((1.5 + 0.1))) / (3.14159 - ((float64(3.14159) - ((1.5 + 3.14159) - 3.14159)) - 3.14159))) == (0.0 * 1.5))) {
                                                        v53 := GenS4{F0: -1}
                                                        _ = v53
                                                    }
                                                    fmt.Println(v10, v43, fi44)
                                                default:
                                                    var v54 int64 = v45
                                                    _ = v54
                                                }
                                                _ = v45
                                            } else {
                                                fmt.Println(v2, fi39, v43)
                                                var v55 float32 = 0.1
                                                _ = v55
                                            }
                                        default:
                                            v32 -= (v2 / (((v32 - v32)) | 1))
                                        }
                                    case 42:
                                        v56 := GenS4{F0: v2}
                                        _ = v56
                                    }
                                    var v57 float32 = (float32(0.0) / (0.1 * 1e10))
                                    switch int(v2) {
                                    case -532:
                                        v2 += (-(v32))
                                    case 673:
                                        v58 := v1[int(v34)]
                                        _ = v58
                                    case 1:
                                        if false {
                                            var v59 int16 = 1000
                                            if true {
                                                var v60 uint8 = (v10 - (v10 * 0))
                                                v61 := GenS5{F0: v32, F1: ((3.14159 == v57) && (0 >= -2147483648)), F2: 65535}
                                                var v62 bool = (65535 != v34)
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            }
                                            switch -368 {
                                            case 9223372036854775807:
                                                if (42 <= 128) {
                                                    var v63 uint8 = 2
                                                    if true {
                                                        fmt.Println(fi39)
                                                    } else {
                                                        switch fi44 {
                                                        case 0:
                                                            if (2 != 459) {
                                                                _ = len(v31)
                                                                fmt.Println(v59, v34)
                                                            }
                                                        case 128:
                                                            v64 := GenS4{F0: (v2 / ((-99) | 1))}
                                                            _ = v64
                                                        case -822:
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    if false {
                                                        _ = len(v1)
                                                        var v65 bool = false
                                                        _ = v65
                                                    } else {
                                                        var v66 string = ((("a" + "hello") + "\n\t") + "\n\t")
                                                        _ = v66
                                                    }
                                                    fmt.Println(v10, v32, v32)
                                                    v34 = v34
                                                    _ = v63
                                                } else {
                                                    v67 := v4[656]
                                                    _ = v67
                                                }
                                            case -657:
                                                var v68 uint64 = 393
                                                _ = v68
                                            case -1:
                                                if (1e10 > 1e10) {
                                                    v69 := GenS1{F0: -641, F1: v10}
                                                    switch (-(fi44)) {
                                                    case 32767:
                                                        v70 := GenS3{F0: 817, F1: uint16(v34), F2: 42}
                                                        _ = v70
                                                    case 282:
                                                        v43 -= 497
                                                    default:
                                                        var v71 int64 = 9223372036854775807
                                                        _ = v71
                                                    }
                                                    var v72 int = fi44
                                                    if false {
                                                        v73 := GenS6{F0: 1}
                                                        v72 -= (fi39 & v72)
                                                        v72 = fi44
                                                        _ = v73
                                                    }
                                                    _ = v69
                                                    _ = v72
                                                } else {
                                                    v74 := GenS7{F0: (v43 << (((v34 ^ 116)) & 63)), F1: v32, F2: 42}
                                                    if ((1.5 - 1e10) < -2.25) {
                                                        v75 := GenS8{F0: v32, F1: 42, F2: (v10 % (((v10 << ((0) & 63))) | 1))}
                                                        if (int64(v10) >= -1) {
                                                            v57 -= v57
                                                        } else {
                                                            v10 = v10
                                                        }
                                                        _ = v75
                                                    } else {
                                                        v43 += (v34 / ((882) | 1))
                                                        var v76 float64 = (1.5 / 1.5)
                                                        fmt.Println(fi44, v43)
                                                        v34 = (v34 | (^(v34)))
                                                        v77 := v3[128]
                                                        v43 += v43
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    if false {
                                                        _ = len(v31)
                                                    } else {
                                                        delete(v4, 552)
                                                        var v78 float64 = (3.14159 + ((float64(v57) / 1.0) + -2.25))
                                                        _ = v78
                                                    }
                                                    _ = v74
                                                }
                                            }
                                            v59 = 1
                                            v2 += -16
                                            _ = v59
                                        } else {
                                            v1[fi39] = true
                                            var v79 uint16 = 581
                                            _ = v79
                                        }
                                    }
                                    if (1 <= -733) {
                                        v80 := GenS5{F0: v2, F1: true, F2: 65535}
                                        _ = v80
                                    }
                                    switch fi44 {
                                    case 127:
                                        if len(v31) > 0 {
                                            v31[0] = (fi39 + fi44)
                                        }
                                    case 128:
                                        fmt.Println(fi44, v2, fi44)
                                    case 587:
                                        v2 = 42
                                    }
                                    v10 = (^(v10))
                                    v81 := GenS1{F0: -32768, F1: (^(v10))}
                                    v82 := GenS9{F0: 42, F1: true}
                                    v43 = uint(v10)
                                    var v83 uint8 = (^(v10))
                                    _ = fi44
                                    _ = v57
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                }
                                switch (fi39 ^ fi39) {
                                case 127:
                                    v34 += (v34 & v34)
                                case -2147483648:
                                    for fi84 := 0; fi84 < 10; fi84++ {
                                        v2 += (v32 ^ v32)
                                        v34 -= (^(v34))
                                        if (true && false) {
                                            v85 := GenS10{F0: 2}
                                            if len(v31) > 0 {
                                                v31[0] = fi84
                                            }
                                            _ = v85
                                        } else {
                                            v86 := v4[451]
                                            _ = v86
                                        }
                                        _ = fi84
                                    }
                                }
                                _ = v43
                            }
                            v10 -= (v10 * v10)
                        } else {
                            switch 1 {
                            case 14:
                                var v87 uint = (v34 ^ (v34 ^ 609))
                                _ = v87
                            case 4294967295:
                                fmt.Println(v10, v10, v2)
                            }
                            switch (fi39 & (fi39 | 72)) {
                            case 1:
                                var v88 uint32 = 22
                                _ = v88
                            case 747:
                                v89 := GenS5{F0: v2, F1: (648 > uint32(v34)), F2: 2147483647}
                                _ = v89
                            default:
                                var v90 int = -347
                                _ = v90
                            }
                            var v91 bool = (!true)
                            _ = v91
                        }
                        for fi92 := 0; fi92 < 9; fi92++ {
                            fmt.Println(v2)
                            _ = fi92
                        }
                        _ = fi39
                    }
                }
                for fi93 := 0; fi93 < 8; fi93++ {
                    var v94 uint64 = 1
                    v10 = (v10 << ((v34) & 63))
                    _ = fi93
                    _ = v94
                }
            }
            _ = v38
        } else {
            var v95 string = (("" + "foo bar") + "hello")
            var v96 uint32 = 2147483648
            v97 := GenS2{F0: (766 <= 0)}
            _ = v95
            _ = v96
            _ = v97
        }
    } else {
        if (42 <= 175) {
            v34 = (v34 ^ 65535)
            v2 = v2
            var v98 int16 = 319
            _ = v98
        } else {
            fmt.Println(v2, v10)
            v32 -= v32
        }
        var v99 int32 = 42
        fmt.Println(v2, v10)
        _ = v99
    }
    for fi100 := 0; fi100 < 15; fi100++ {
        for fi101 := 0; fi101 < 18; fi101++ {
            v1[(fi101 & (fi100 + fi100))] = true
            var v102 int8 = 1
            _ = fi101
            _ = v102
        }
        _ = fi100
    }
    v10 -= v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
