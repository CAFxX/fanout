package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int; F1 bool; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 uint64; F2 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int32 = 0
    var v2 bool = false
    var v3 int32 = (v1 ^ 2147483647)
    var v4 uint32 = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: uint16(v4), F1: ("foo bar" + (("a" + "hello") + "a")), F2: 2}
        var v7 bool = (true || (-128 == 42))
        _ = fi5
        _ = v6
        _ = v7
    }
    switch -2147483648 {
    case 42:
        v2 = ((!v2) && false)
    case 402:
        if (1e10 < ((float32(float64(3.14159)) + (float32((1.5 - 1.5)) * (1.5 - 0.0))) / 1.0)) {
            switch ("" + "a") {
            case "hello":
                fmt.Println(v2)
            case "a":
                var v8 int = -470
                _ = v8
            default:
                for fi9 := 0; fi9 < 16; fi9++ {
                    if true {
                        v10 := []int{9223372036854775807}
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v1 = (v3 | 42)
                            _ = fi11
                        }
                        _ = v10
                    } else {
                        if (false && true) {
                            fmt.Println(v3, v4)
                            v1 += 1
                        }
                        var v12 uint = 268
                        if (7 <= 42) {
                            if (!(v2 || v2)) {
                                v4 += (v4 + (^(v4)))
                                var v13 uint = v12
                                _ = v13
                            }
                            var v14 int32 = (v1 + 0)
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                if false {
                                    v16 := []int{}
                                    var v17 string = "\n\t"
                                    var v18 float32 = 1e10
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    fmt.Println(v4)
                                }
                                if (100 < -128) {
                                    switch 1 {
                                    case -1:
                                        if v2 {
                                            var v19 uint = v12
                                            v3 = (^(v3))
                                            var v20 uint32 = uint32(v12)
                                            _ = v19
                                            _ = v20
                                        } else {
                                            var v21 string = "\n\t"
                                            if (!false) {
                                                var v22 int32 = (v3 | (v1 << ((1) & 63)))
                                                if (5 <= 2) {
                                                    var v23 uint8 = 1
                                                    _ = v23
                                                } else {
                                                    v24 := GenS2{F0: -9223372036854775808, F1: (v4 - (v4 + v4)), F2: ((-2.25 + 0.1) < (3.14159 + (0.0 - 0.1)))}
                                                    var v25 uint8 = 2
                                                    _ = v24
                                                    _ = v25
                                                }
                                                fmt.Println(v22, v4, v2)
                                                v26 := GenS1{F0: 0, F1: v21, F2: 472}
                                                _ = v22
                                                _ = v26
                                            }
                                            _ = v21
                                        }
                                    default:
                                        v4 -= (^(v4))
                                    }
                                    v27 := GenS2{F0: 1, F1: (^(v4)), F2: (((0.0 / 0.1) - 1.5) != float32((float64(((0.1 - (0.0 / 0.1)) - float32((3.14159 + ((-2.25 + 0.1) - (3.14159 * -2.25)))))) - 3.14159)))}
                                    v3 += 2147483647
                                    _ = v27
                                } else {
                                    fmt.Println(v4, v3, v3)
                                    v28 := GenS3{F0: -487, F1: 0, F2: 0}
                                    var v29 int64 = 9223372036854775807
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi15
                            }
                            for fi30 := 0; fi30 < 10; fi30++ {
                                fmt.Println(v2, v4, v2)
                                switch "a" {
                                case "":
                                    v31 := []int{fi9, -112, 763}
                                    fmt.Println(fi9, v4, v1)
                                    _ = v31
                                }
                                if (!((false && (-2.25 < (3.14159 - (float64(0.1) / 1.0)))) || (0 > -1))) {
                                    v4 = (v4 + v4)
                                }
                                switch (fi30 >> ((uint(v4)) & 63)) {
                                case 32767:
                                    var v32 uint8 = uint8(v4)
                                    _ = v32
                                case 127:
                                    v33 := []int{32767, -616, fi30, (fi9 << ((v12) & 63))}
                                    _ = v33
                                }
                                v34 := []int{(fi30 << (((^(v12))) & 63)), (fi9 & (fi9 % ((128) | 1))), fi9}
                                v1 -= (v1 ^ v1)
                                _ = fi30
                                _ = v34
                            }
                        }
                        _ = v12
                    }
                    _ = fi9
                }
            }
            v35 := GenS1{F0: 32767, F1: ("" + ""), F2: 18446744073709551615}
            _ = v35
        } else {
            if (!v2) {
                if (2 < 255) {
                    switch "hello" {
                    case "\n\t":
                        v36 := GenS1{F0: 2, F1: ("a" + "\n\t"), F2: uint64(v4)}
                        _ = v36
                    case "foo bar":
                        var v37 uint8 = 255
                        _ = v37
                    default:
                        v4 = v4
                    }
                    v38 := []int{1, -9223372036854775808, 9223372036854775807, 42}
                    v39 := GenS4{F0: 0, F1: v4}
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        v41 := GenS3{F0: 0, F1: 9223372036854775807, F2: 1}
                        _ = fi40
                        _ = v41
                    }
                }
            }
        }
    case 1:
        if v2 {
            switch -9223372036854775808 {
            case -9223372036854775808:
                v42 := GenS5{F0: -9223372036854775808, F1: v2, F2: 9223372036854775807}
                _ = v42
            }
            v4 += v4
        }
    default:
        if v2 {
            var v43 uint32 = (^(v4))
            for fi44 := 0; fi44 < 6; fi44++ {
                v45 := GenS6{F0: 1, F1: v43}
                if v2 {
                    for fi46 := 0; fi46 < 19; fi46++ {
                        var v47 uint8 = 186
                        fmt.Println(v1, v2)
                        var v48 int8 = 1
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    v49 := GenS5{F0: fi44, F1: false, F2: 791}
                    _ = v49
                } else {
                    if ((v1 < -512) && v2) {
                        if (!v2) {
                            for fi50 := 0; fi50 < 20; fi50++ {
                                var v51 int32 = v1
                                var v52 int8 = -128
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        }
                        for fi53 := 0; fi53 < 4; fi53++ {
                            fmt.Println(v4, fi53, v3)
                            v54 := []int{(^(fi44)), (fi53 | (-(fi53)))}
                            if len(v54) > 0 {
                                v54[0] = (fi44 & (fi53 & fi53))
                            }
                            _ = fi53
                            _ = v54
                        }
                        v55 := GenS7{F0: fi44, F1: -3}
                        _ = v55
                    }
                }
                v56 := GenS6{F0: 42, F1: 42}
                _ = fi44
                _ = v45
                _ = v56
            }
            fmt.Println(v3)
            _ = v43
        } else {
            v57 := []int{2147483647, -170, 42, -1}
            _ = v57
        }
    }
    if v2 {
        v3 = -1
    } else {
        for fi58 := 0; fi58 < 16; fi58++ {
            if v2 {
                fmt.Println(v4, v4)
            }
            v3 = (v3 * -1)
            _ = fi58
        }
        v59 := GenS8{F0: 158}
        _ = v59
    }
    var v60 int32 = 261
    fmt.Println(v60, v60, v1)
    if ((v3 << ((216) & 63)) >= v1) {
        fmt.Println(v3, v3)
        var v61 int8 = -128
        _ = v61
    }
    var v62 string = "hello"
    switch (v62 + ("hello" + v62)) {
    case "hello":
        switch -820 {
        case 904:
            if false {
                if v2 {
                    fmt.Println(v1)
                    v63 := GenS9{F0: 0, F1: 0, F2: 1000}
                    var v64 int = 850
                    _ = v63
                    _ = v64
                } else {
                    if ((3.14159 >= (0.0 - 0.0)) || (217 != 42)) {
                        v65 := GenS10{F0: 42}
                        v66 := []int{-1}
                        v3 -= v1
                        _ = v65
                        _ = v66
                    } else {
                        v67 := []int{2147483647, 127}
                        switch (v62 + v62) {
                        case "a":
                            fmt.Println(v1, v1)
                        }
                        _ = v67
                    }
                }
                v68 := GenS11{F0: 2147483647}
                _ = v68
            } else {
                fmt.Println(v4)
            }
        default:
            v3 += -1
        }
    case "\n\t":
        for fi69 := 0; fi69 < 9; fi69++ {
            v70 := []int{}
            var v71 string = v62
            var v72 int64 = -1
            _ = fi69
            _ = v70
            _ = v71
            _ = v72
        }
    default:
        if false {
            for fi73 := 0; fi73 < 15; fi73++ {
                if (142 >= 81) {
                    switch ("a" + "\n\t") {
                    case "a":
                        v60 -= 42
                    case "\n\t":
                        var v74 int32 = (v3 << ((255) & 63))
                        _ = v74
                    default:
                        if (-1 < (-(fi73))) {
                            v75 := []int{-48, fi73, (fi73 * fi73), (fi73 ^ (fi73 | 32768))}
                            _ = len(v75)
                            v4 = 97
                            _ = v75
                        } else {
                            switch -195 {
                            case -1:
                                fmt.Println(v60, v60)
                            }
                        }
                    }
                }
                for fi76 := 0; fi76 < 12; fi76++ {
                    var v77 int32 = 0
                    fmt.Println(v60, v2)
                    v78 := []int{fi73}
                    var v79 int = 0
                    _ = fi76
                    _ = v77
                    _ = v78
                    _ = v79
                }
                _ = fi73
            }
            v4 = 595
        } else {
            switch 127 {
            case -428:
                v80 := []int{}
                _ = v80
            default:
                v81 := GenS12{F0: (v60 / ((-1) | 1)), F1: 145}
                _ = v81
            }
            v82 := []int{2147483647}
            _ = v82
        }
    }
    if (v2 || (!v2)) {
        v83 := []int{-1, 193}
        _ = v83
    } else {
        if v2 {
            var v84 int16 = 814
            fmt.Println(v62, v1, v1)
            _ = v84
        } else {
            fmt.Println(v1, v1, v4)
            v60 = (^(v1))
            for fi85 := 0; fi85 < 12; fi85++ {
                var v86 float32 = -2.25
                v62 = v62
                _ = fi85
                _ = v86
            }
        }
    }
    v1 = (v3 << ((724) & 63))
    switch -820 {
    case -108:
        fmt.Println(v2, v62)
    case 714:
        var v87 uint32 = 42
        _ = v87
    default:
        var v88 uint64 = 446
        _ = v88
    }
    v60 = (v1 << ((332) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v62
}
