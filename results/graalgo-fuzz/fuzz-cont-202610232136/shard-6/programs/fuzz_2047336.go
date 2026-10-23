package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint }
func genFn0(p0 int32) uint16 {
    v7 := []int{4294967295, 774, 4294967295, 317}
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 = p0
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: -128}
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 127}
            _ = fi8
            _ = v11
        }
        switch "a" {
        case "":
            p0 = 1
        case "a":
            v12 := GenS1{F0: -128}
            _ = v12
        default:
            switch (("" + "hello") + (("\n\t" + "foo bar") + "hello")) {
            case "foo bar":
                if true {
                    switch 128 {
                    case 2147483647:
                        v7 = append(v7, 9223372036854775807)
                    default:
                        var v13 uint = 951
                        _ = v13
                    }
                    p0 = 44
                }
            case "hello":
                if (2 == 295) {
                    var v14 int = 32768
                    var v15 uint = 820
                    v15 -= (v15 | (v15 | (^(v15))))
                    _ = v14
                    _ = v15
                }
            default:
                p0 += -1
            }
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            switch "hello" {
            case "":
                if (1 >= -67) {
                    switch (fi16 | fi16) {
                    case 1:
                        p0 = 0
                    default:
                        switch ("a" + "a") {
                        case "foo bar":
                            var v17 float32 = float32((1.5 - float64(0.0)))
                            _ = v17
                        case "\n\t":
                            var v18 uint8 = 170
                            _ = v18
                        default:
                            p0 = p0
                        }
                    }
                    if true {
                        if false {
                            switch fi16 {
                            case 461:
                                var v19 uint32 = 2147483647
                                _ = v19
                            default:
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    v21 := GenS1{F0: -128}
                                    var v22 string = "a"
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                            var v23 uint64 = 42
                            _ = v23
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                }
            case "\n\t":
                p0 -= (p0 - (p0 | 0))
            case "foo bar":
                _ = len(v7)
            default:
                if len(v7) > 0 {
                    v7[0] = -432
                }
            }
            p0 = p0
            var v24 string = "\n\t"
            _ = fi16
            _ = v24
        }
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int64) int32 {
    var v25 bool = false
    if (4294967295 <= 42) {
        if false {
            var v26 uint32 = 538
            fmt.Println(v25)
            _ = v26
        } else {
            switch ("foo bar" + ("hello" + "hello")) {
            case "hello":
                p0 = 9223372036854775807
            }
            var v27 uint32 = 2147483647
            if false {
                p0 = -654
            }
            _ = v27
        }
        v25 = (v25 || v25)
        var v28 float64 = 1.5
        _ = v28
    }
    p0 += (p0 >> ((18446744073709551615) & 63))
    v29 := []int{42}
    _ = p0
    _ = v25
    _ = v29
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 0
    var v3 int32 = -2147483648
    v4 := make(map[int]string)
    var v5 int8 = 127
    var v6 float64 = (-2.25 + 0.0)
    for fi30 := 0; fi30 < 15; fi30++ {
        v1[-9223372036854775808] = 4294967295
        _ = fi30
    }
    var v31 float32 = 3.14159
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int = (fi32 >> ((842) & 63))
        v34 := []int{}
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if (!(0 == 1)) {
            var v36 uint32 = 65535
            _ = v36
        } else {
            _ = len(v4)
            fmt.Println(v31, v5)
            fmt.Println(v2)
        }
        v37 := []int{}
        _ = fi35
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        fmt.Println(fi38)
        v39 := GenS2{F0: v5}
        _ = fi38
        _ = v39
    }
    if true {
        fmt.Println(v5, v5)
        if true {
            v2 = (v2 * (v2 >> (((v2 / ((v2) | 1))) & 63)))
            switch -280 {
            case 2147483647:
                fmt.Println(v3, v3, v3)
            default:
                fmt.Println(v5)
            }
        } else {
            _ = len(v1)
            v40 := GenS3{F0: (v2 / ((65535) | 1)), F1: -128}
            _ = v40
        }
    }
    v41 := v4[32767]
    v3 -= (v3 << (((^(v2))) & 63))
    var v42 uint32 = 65535
    for fi43 := 0; fi43 < 3; fi43++ {
        switch (v41 + v41) {
        case "\n\t":
            if ((v5 ^ v5) <= v5) {
                var v44 uint8 = 1
                var v45 string = v41
                switch fi43 {
                case -9223372036854775808:
                    var v46 int64 = -927
                    _ = v46
                default:
                    v47 := []int{128, (^(fi43)), fi43}
                    _ = v47
                }
                _ = v44
                _ = v45
            } else {
                if false {
                    switch 32767 {
                    case 1:
                        v48 := []int{(fi43 | (fi43 ^ 32768)), fi43}
                        _ = v48
                    default:
                        fmt.Println(v31)
                    }
                } else {
                    delete(v4, fi43)
                    v49 := []int{(fi43 | (fi43 ^ (fi43 / ((-724) | 1)))), 128}
                    fmt.Println(v31, v31)
                    _ = v49
                }
                v50 := GenS1{F0: v5}
                v51 := GenS4{F0: v42, F1: -1}
                _ = v50
                _ = v51
            }
        case "a":
            v52 := GenS4{F0: v42, F1: 187}
            _ = v52
        }
        _ = fi43
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        switch v41 {
        case "a":
            if false {
                v54 := GenS3{F0: (v2 * v2), F1: (v5 / (((v5 >> ((v2) & 63))) | 1))}
                _ = v54
            }
        case "\n\t":
            v55 := []int{(fi53 ^ (fi53 + fi53)), (fi53 + fi53), fi53}
            _ = v55
        case "":
            var v56 uint32 = (v42 + (v42 * 82))
            _ = v56
        }
        if false {
            fmt.Println(v42)
            if (v41 == ("hello" + v41)) {
                var v57 float64 = 1e10
                v58 := GenS1{F0: v5}
                _ = v57
                _ = v58
            }
        } else {
            if (((0 < 18446744073709551615) || false) || (false && true)) {
                if ((^(v42)) != 2147483647) {
                    v41 = "foo bar"
                    switch (v41 + "foo bar") {
                    case "foo bar":
                        var v59 uint64 = 2
                        _ = v59
                    case "\n\t":
                        v60 := []int{443, -1, fi53, int(v5)}
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 9; fi61++ {
                        if true {
                            var v62 int32 = 618
                            v63 := []int{(fi53 - 127), fi53, fi61}
                            _ = v62
                            _ = v63
                        } else {
                            v64 := []int{int(v2)}
                            var v65 uint64 = 107
                            if (!false) {
                                v65 = v65
                            } else {
                                if (v6 > v6) {
                                    var v66 uint = v2
                                    _ = v66
                                }
                                v64 = append(v64, fi61)
                                v4[-9223372036854775808] = v41
                            }
                            var v67 bool = ((v2 <= v2) && true)
                            _ = v64
                            _ = v65
                            _ = v67
                        }
                        v3 -= v3
                        var v68 int8 = v5
                        v69 := GenS5{F0: v42}
                        v70 := GenS6{F0: 127, F1: 2147483647}
                        _ = fi61
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                }
                for fi71 := 0; fi71 < 8; fi71++ {
                    var v72 uint = v2
                    var v73 float64 = v6
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            }
            var v74 int16 = int16(v2)
            v5 += v5
            _ = v74
        }
        _ = len(v1)
        _ = fi53
    }
    if (float32(3.14159) != -2.25) {
        var v75 int8 = 59
        switch v41 {
        case "hello":
            var v76 uint16 = 244
            _ = v76
        }
        var v77 uint64 = 4294967295
        _ = v75
        _ = v77
    } else {
        switch "a" {
        case "foo bar":
            switch ("foo bar" + (((("\n\t" + "") + (v41 + v41)) + v41) + (((v41 + ("hello" + (("a" + v41) + (v41 + v41)))) + v41) + v41))) {
            case "a":
                for fi78 := 0; fi78 < 8; fi78++ {
                    v79 := GenS1{F0: (v5 * (-(v5)))}
                    _ = fi78
                    _ = v79
                }
            case "hello":
                v80 := GenS7{F0: 0, F1: v2}
                _ = v80
            }
        }
    }
    var v81 uint8 = 2
    var v82 float64 = v6
    for fi83 := 0; fi83 < 16; fi83++ {
        fmt.Println(v42, v42)
        var v84 int64 = int64(v5)
        _ = fi83
        _ = v84
    }
    var v85 int64 = 365
    v86 := GenS1{F0: (v5 * v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v85
    _ = v86
}
