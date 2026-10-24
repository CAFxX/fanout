package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string; F2 int }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int; F1 int16; F2 int64 }
type GenS7 struct { F0 int64; F1 int8; F2 int32 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint) uint16 {
    p0 = (p0 << (((^(p0))) & 63))
    p0 -= 153
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint32 = 2147483648
        v6 += v6
        _ = fi5
        _ = v6
    }
    p0 = p0
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = (v1 + (v1 | v1))
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 6; fi8++ {
        if true {
            v9 := GenS1{F0: "a", F1: "foo bar"}
            _ = v9
        }
        var v10 bool = (42 > 42)
        _ = fi8
        _ = v10
    }
    v11 := GenS1{F0: "a", F1: "\n\t"}
    if true {
        fmt.Println(v7)
    }
    var v12 string = ""
    v13 := GenS1{F0: ("a" + "hello"), F1: "foo bar"}
    if (255 != 1) {
        v14 := GenS1{F0: "", F1: (("hello" + v12) + "hello")}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (2 > 0) {
                fmt.Println(v3)
                if true {
                    v16 := []int{-732}
                    _ = v16
                } else {
                    fmt.Println(v7, v1)
                }
                if false {
                    v17 := GenS1{F0: (v12 + v12), F1: "\n\t"}
                    v18 := []int{}
                    var v19 int64 = 9223372036854775807
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v20 := []int{}
                    v7 -= (v1 % ((genFn0(255)) | 1))
                    v21 := GenS1{F0: v12, F1: "\n\t"}
                    _ = v20
                    _ = v21
                }
            }
            fmt.Println(v3, v1)
            if true {
                v22 := []int{fi15, (fi15 & (fi15 >> ((250) & 63))), -2147483648, (-(fi15))}
                var v23 int16 = -284
                fmt.Println(v2)
                _ = v22
                _ = v23
            }
            _ = fi15
        }
    }
    if false {
        v24 := GenS1{F0: v12, F1: (v12 + "")}
        v25 := []int{4294967295, 499, 2147483647, -2147483648}
        _ = v24
        _ = v25
    } else {
        var v26 bool = false
        var v27 float32 = float32(float64(0.1))
        _ = v26
        _ = v27
    }
    switch "" {
    case "\n\t":
        switch -9223372036854775808 {
        case 127:
            v2 = 42
        case 128:
            switch (v12 + v12) {
            case "":
                v12 = ("" + v12)
            case "foo bar":
                fmt.Println(v1, v7, v2)
            }
        case 32768:
            var v28 float64 = -2.25
            _ = v28
        }
    case "":
        v2 += v2
    case "a":
        fmt.Println(v7)
    default:
        v29 := []int{714, -307}
        _ = v29
    }
    if (false && false) {
        v1 = (v3 >> ((2) & 63))
        var v30 bool = (443 == 0)
        _ = v30
    } else {
        v31 := []int{2147483647, 853, int(v2), -980}
        v32 := GenS2{F0: "hello", F1: v12, F2: 1}
        var v33 bool = false
        _ = v31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 uint64 = v2
        _ = fi34
        _ = v35
    }
    switch 127 {
    case 32767:
        if false {
            if (true || (true && false)) {
                v12 = (v12 + v12)
                fmt.Println(v1, v1)
                var v36 int64 = 0
                _ = v36
            }
        } else {
            if false {
                for fi37 := 0; fi37 < 13; fi37++ {
                    for fi38 := 0; fi38 < 10; fi38++ {
                        v1 -= genFn0(2)
                        v39 := []int{-655}
                        if len(v39) > 0 {
                            v39[0] = (fi37 % (((fi38 << ((0) & 63))) | 1))
                        }
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
                for fi40 := 0; fi40 < 14; fi40++ {
                    v41 := GenS3{F0: (-(fi40))}
                    _ = fi40
                    _ = v41
                }
            } else {
                v42 := GenS3{F0: 128}
                v2 -= v2
                switch 59 {
                case -2147483648:
                    switch 240 {
                    case -473:
                        if (61 >= -93) {
                            switch 64 {
                            case 1:
                                var v43 string = ((("\n\t" + v12) + "foo bar") + ("\n\t" + ((v12 + v12) + "foo bar")))
                                _ = v43
                            case 32767:
                                v3 += v1
                            case 2147483647:
                                switch 9223372036854775807 {
                                case 538:
                                    for fi44 := 0; fi44 < 3; fi44++ {
                                        v45 := GenS4{F0: 42, F1: ((v12 + "a") + ""), F2: 128}
                                        for fi46 := 0; fi46 < 12; fi46++ {
                                            switch v12 {
                                            case "a":
                                                v7 += v1
                                            case "":
                                                if (fi46 != 4294967295) {
                                                    var v47 string = v12
                                                    v1 = v1
                                                    var v48 float64 = 1e10
                                                    switch v12 {
                                                    case "hello":
                                                        if true {
                                                            fmt.Println(v47, v12)
                                                        }
                                                    default:
                                                        v12 = ((v47 + (v47 + v12)) + "foo bar")
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                }
                                            default:
                                                var v49 uint8 = 128
                                                _ = v49
                                            }
                                            v50 := []int{(^(fi46)), fi46}
                                            if false {
                                                switch -9223372036854775808 {
                                                case 482:
                                                    fmt.Println(v1, fi44)
                                                }
                                                if false {
                                                    v50 = append(v50, (fi46 * fi44))
                                                    if true {
                                                        var v51 uint64 = (^(v2))
                                                        switch ((v12 + "") + v12) {
                                                        case "\n\t":
                                                            v2 += (v2 / ((v2) | 1))
                                                        default:
                                                            v52 := GenS3{F0: fi44}
                                                            _ = v52
                                                        }
                                                        v7 -= (v7 / (((v1 + (v1 << ((0) & 63)))) | 1))
                                                        v3 += v7
                                                        _ = v51
                                                    } else {
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(fi46, fi44)
                                                        case "\n\t":
                                                            v1 = 574
                                                        case "hello":
                                                            var v53 uint = 1
                                                            _ = v53
                                                        default:
                                                            var v54 int32 = 1
                                                            _ = v54
                                                        }
                                                        v12 = "a"
                                                        v55 := GenS5{F0: (v3 - genFn0(58)), F1: 38}
                                                        _ = v55
                                                    }
                                                    var v56 int32 = -964
                                                    _ = v56
                                                } else {
                                                    v3 += v7
                                                }
                                                v57 := GenS6{F0: (fi44 | fi46), F1: -127, F2: -1}
                                                switch fi44 {
                                                case -1:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if (!(-31 > 1000)) {
                                                            v58 := GenS3{F0: fi44}
                                                            v2 -= (v2 & 944)
                                                            v12 = "foo bar"
                                                            _ = v58
                                                        } else {
                                                            var v59 int16 = 1
                                                            var v60 int = fi44
                                                            v50 = append(v50, fi46)
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                    if len(v50) >= 2 {
                                                        v50 = v50[1:len(v50)]
                                                    }
                                                    var v61 uint8 = 1
                                                    fmt.Println(fi44, v7)
                                                    v62 := GenS1{F0: "", F1: ((v12 + v12) + ((v12 + "") + v12))}
                                                    _ = v61
                                                    _ = v62
                                                case 32768:
                                                    var v63 int8 = 100
                                                    _ = v63
                                                }
                                                _ = v57
                                            } else {
                                                var v64 uint64 = (v2 | 18446744073709551615)
                                                _ = v64
                                            }
                                            _ = fi46
                                            _ = v50
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                case 1:
                                    v65 := GenS3{F0: -2147483648}
                                    _ = v65
                                case 9223372036854775807:
                                    v2 += 12
                                }
                            }
                            switch int(v1) {
                            case 32768:
                                fmt.Println(v12, v7, v12)
                            case 797:
                                fmt.Println(v12, v7)
                            }
                            for fi66 := 0; fi66 < 20; fi66++ {
                                if false {
                                    v67 := []int{fi66, (-(fi66))}
                                    _ = v67
                                }
                                switch v12 {
                                case "hello":
                                    v2 -= 179
                                default:
                                    if true {
                                        for fi68 := 0; fi68 < 11; fi68++ {
                                            v12 = (v12 + ((("a" + ("" + "a")) + v12) + (v12 + ("foo bar" + (v12 + (v12 + "a"))))))
                                            _ = fi68
                                        }
                                    }
                                }
                                fmt.Println(v3, v1, v3)
                                _ = fi66
                            }
                        } else {
                            var v69 int32 = 1
                            _ = v69
                        }
                    }
                case -682:
                    fmt.Println(v7)
                case -286:
                    v70 := []int{}
                    _ = v70
                }
                _ = v42
            }
        }
    case -559:
        v71 := GenS7{F0: -1, F1: 0, F2: -1}
        _ = v71
    case 42:
        v3 = (v7 << ((57) & 63))
    }
    if (((float64((0.1 + float32((float64(0.1) - 1.5)))) * (((-2.25 * 0.1) * 3.14159) / 0.1)) == (float64((3.14159 / 1.5)) - (((0.1 / (0.0 - 1.5)) * ((0.0 + 1.5) + (0.0 / 1.0))) + -2.25))) || false) {
        v72 := GenS8{F0: false, F1: v2}
        _ = v72
    }
    switch 0 {
    case 1:
        fmt.Println(v12)
    }
    var v73 int8 = 0
    v74 := GenS5{F0: 65535, F1: 7}
    v75 := GenS9{F0: 1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v73
    _ = v74
    _ = v75
}
