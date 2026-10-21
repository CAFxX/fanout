package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8) int {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{}
            p0 = -125
            if (591 > 2147483648) {
                var v9 bool = true
                v9 = (0 >= 239)
                p0 = p0
                _ = v9
            } else {
                var v10 uint8 = 1
                _ = v10
            }
            p0 = -128
            v11 := GenS1{F0: 1, F1: false, F2: 843}
            _ = fi7
            _ = v11
            _ = v8
        }
        _ = fi6
    }
    var v12 int16 = 32767
    if false {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint64 = 42
            if false {
                if false {
                    switch fi13 {
                    case 1:
                        for fi15 := 0; fi15 < 9; fi15++ {
                            fmt.Println(v12, v14)
                            switch "foo bar" {
                            case "hello":
                                var v16 int = (fi15 << ((4294967295) & 63))
                                _ = v16
                            default:
                                fmt.Println(fi15, fi15)
                            }
                            v17 := []int{}
                            v14 = 1
                            v18 := GenS1{F0: 65535, F1: (false || false), F2: 32767}
                            _ = fi15
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    var v19 int64 = -1
                    _ = v19
                }
            }
            _ = fi13
            _ = v14
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: 1, F1: true, F2: 2}
        _ = fi20
        _ = v21
    }
    var v22 uint8 = 62
    _ = p0
    _ = v12
    _ = v22
    return 127
}

func main() {
    var v1 int = 128
    var v2 bool = true
    v3 := []int{(v1 | (-(v1))), (v1 << ((312) & 63)), v1}
    var v4 float64 = float64(float32((float64(float32(float64((3.14159 - 1.5)))) - (-2.25 * (float64(0.0) - 0.1)))))
    var v5 uint16 = 0
    switch (("" + ("foo bar" + "hello")) + "") {
    case "a":
        var v23 int64 = -1
        _ = v23
    default:
        fmt.Println(v5, v4, v5)
    }
    v4 += 3.14159
    if v2 {
        if v2 {
            if false {
                var v24 int32 = -67
                if (96 >= 250) {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        switch 4294967295 {
                        case 753:
                            for fi26 := 0; fi26 < 6; fi26++ {
                                var v27 float64 = -2.25
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = fi25
                    }
                } else {
                    v28 := GenS1{F0: 575, F1: v2, F2: (^(v5))}
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v24 = 1
                        var v30 bool = v2
                        var v31 bool = (v2 && v30)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    if (v24 < (v24 >> ((42) & 63))) {
                        fmt.Println(v4, v4)
                    }
                    _ = v28
                }
                _ = v24
            }
            v2 = false
        }
        fmt.Println(v2, v5)
        if len(v3) > 0 {
            v3[0] = v1
        }
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint16 = 255
            v33 = v33
            _ = fi32
            _ = v33
        }
        var v34 uint32 = 65535
        switch "" {
        case "\n\t":
            switch "\n\t" {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v4 -= ((float64((1e10 - float32((v4 / (0.1 * (((-2.25 + v4) / -2.25) + v4)))))) * float64(0.1)) - -2.25)
            }
        }
        _ = v34
    }
    var v35 uint8 = 127
    if v2 {
        for fi36 := 0; fi36 < 17; fi36++ {
            if v2 {
                if ((v4 - v4) == ((v4 * (v4 * v4)) / 3.14159)) {
                    var v37 int16 = 0
                    for fi38 := 0; fi38 < 2; fi38++ {
                        var v39 uint = 1
                        if (v2 || ((v2 || false) || ((((0.1 / 1.5) / ((0.1 * 0.0) + (1e10 + 0.1))) - 1e10) != -2.25))) {
                            switch "" {
                            case "hello":
                                v5 = (v5 % (((v5 << (((^(v39))) & 63))) | 1))
                            case "foo bar":
                                v35 = 42
                            default:
                                var v40 uint16 = (v5 << ((v39) & 63))
                                _ = v40
                            }
                            var v41 uint64 = 2
                            var v42 float64 = (v4 + 0.0)
                            _ = v41
                            _ = v42
                        }
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
            } else {
                fmt.Println(v2)
            }
            var v43 int = (^(v1))
            if v2 {
                var v44 int8 = -1
                v1 += fi36
                _ = v44
            } else {
                for fi45 := 0; fi45 < 13; fi45++ {
                    v2 = v2
                    fmt.Println(v35)
                    switch genFn0(-7) {
                    case -654:
                        v43 += (v1 ^ (v1 % ((-2147483648) | 1)))
                    case 2147483647:
                        fmt.Println(v43, v43, fi36)
                    }
                    if (0 != 42) {
                        v46 := GenS1{F0: uint(v5), F1: ("hello" == ("hello" + "a")), F2: 42}
                        var v47 int16 = 32767
                        _ = v46
                        _ = v47
                    } else {
                        if v2 {
                            v48 := GenS1{F0: 18446744073709551615, F1: (v2 && (v2 && false)), F2: 425}
                            var v49 int8 = 127
                            var v50 float32 = 0.0
                            _ = v48
                            _ = v49
                            _ = v50
                        } else {
                            v1 = genFn0(0)
                            v51 := GenS1{F0: 42, F1: v2, F2: 945}
                            _ = v51
                        }
                        if v2 {
                            var v52 uint = 0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v53 float32 = 1e10
                            _ = v52
                            _ = v53
                        } else {
                            v35 += 0
                        }
                        v2 = v2
                        v4 -= 0.0
                        fmt.Println(fi45, fi45)
                        switch fi36 {
                        case 947:
                            v1 = v43
                        }
                    }
                    _ = fi45
                }
                var v54 int32 = -870
                if (int32(v5) == v54) {
                    v55 := GenS2{F0: 604, F1: 234, F2: ""}
                    if false {
                        if (9223372036854775807 == 23) {
                            var v56 int16 = -32768
                            v57 := GenS3{F0: v56}
                            v1 = fi36
                            _ = v56
                            _ = v57
                        } else {
                            switch ("foo bar" + ((("" + "\n\t") + "foo bar") + "\n\t")) {
                            case "\n\t":
                                _ = len(v3)
                            case "a":
                                v35 += (v35 << ((255) & 63))
                            case "":
                                if ((v5 + (^(v5))) > (v5 | 2)) {
                                    if len(v3) > 0 {
                                        v3[0] = (v43 << ((629) & 63))
                                    }
                                    v43 -= (fi36 & v43)
                                    if (!(255 > 42)) {
                                        var v58 uint = 20
                                        v35 += 36
                                        _ = v58
                                    } else {
                                        _ = len(v3)
                                        v4 = (1.5 - 0.1)
                                    }
                                }
                            default:
                                var v59 uint32 = 1
                                _ = v59
                            }
                        }
                        var v60 float32 = float32(v4)
                        if (0.0 == ((0.1 / 1.0) / v4)) {
                            v1 += (v43 % ((1) | 1))
                            v3 = append(v3, 32767)
                            v61 := GenS2{F0: 390, F1: (v35 - 88), F2: (("hello" + "\n\t") + "\n\t")}
                            _ = v61
                        }
                        _ = v60
                    } else {
                        fmt.Println(v54, v43, v2)
                        v62 := GenS3{F0: 1000}
                        v2 = (!false)
                        _ = v62
                    }
                    _ = v55
                } else {
                    switch (v1 + v1) {
                    case 9223372036854775807:
                        v43 = fi36
                    case 706:
                        var v63 bool = ((v4 * 0.1) > v4)
                        _ = v63
                    case 0:
                        v64 := GenS1{F0: 848, F1: (18446744073709551615 == 42), F2: 2}
                        _ = v64
                    default:
                        v65 := GenS1{F0: 2, F1: (-947 == 1), F2: v5}
                        _ = v65
                    }
                }
                _ = v54
            }
            _ = fi36
            _ = v43
        }
        for fi66 := 0; fi66 < 1; fi66++ {
            fmt.Println(v1, fi66)
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi67 := 0; fi67 < 18; fi67++ {
                    if ((^(v5)) > v5) {
                        fmt.Println(v35, v2)
                    } else {
                        v35 = v35
                        if (-490 > -1) {
                            var v68 uint64 = 0
                            fmt.Println(v68, v2)
                            var v69 int32 = -1
                            _ = v68
                            _ = v69
                        }
                    }
                    v35 = (v35 ^ v35)
                    _ = fi67
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = v1
                }
                for fi70 := 0; fi70 < 14; fi70++ {
                    v5 -= (v5 % ((v5) | 1))
                    _ = fi70
                }
            }
            _ = fi66
        }
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        switch (^(fi71)) {
        case -40:
            for fi72 := 0; fi72 < 10; fi72++ {
                fmt.Println(v5, v4)
                fmt.Println(v4, fi71, v2)
                switch (("hello" + (("\n\t" + "hello") + ("" + ""))) + ("\n\t" + (("hello" + "hello") + ("hello" + "")))) {
                case "":
                    switch (fi71 << ((1) & 63)) {
                    case 4294967295:
                        var v73 float64 = v4
                        _ = v73
                    default:
                        v74 := GenS4{F0: -1}
                        _ = v74
                    }
                case "foo bar":
                    if ((!(0 > -1)) || false) {
                        var v75 int8 = 42
                        _ = v75
                    } else {
                        if v2 {
                            if (1 <= 32) {
                                var v76 int32 = 0
                                var v77 float64 = (1e10 - (((3.14159 / v4) + 0.0) + float64(1e10)))
                                if (1.5 > 1e10) {
                                    v4 -= (v77 / 1e10)
                                }
                                _ = v76
                                _ = v77
                            } else {
                                var v78 uint16 = 42
                                fmt.Println(v4, v5)
                                _ = v78
                            }
                            var v79 float64 = v4
                            _ = v79
                        }
                        var v80 uint = 4294967295
                        _ = v80
                    }
                }
                _ = fi72
            }
        case 0:
            switch "hello" {
            case "\n\t":
                if (-1 <= int32(v5)) {
                    v2 = (true && v2)
                }
            case "a":
                var v81 uint16 = 255
                _ = v81
            }
        default:
            if len(v3) > 0 {
                v3[0] = int(v5)
            }
        }
        var v82 int32 = 42
        _ = fi71
        _ = v82
    }
    var v83 uint32 = 786
    fmt.Println(v83)
    v35 += v35
    v2 = (18446744073709551615 < 619)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v83
}
