package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 string = "hello"
    var v4 int32 = 0
    if true {
        var v5 uint32 = 4294967295
        if (58 < 1) {
            switch -1 {
            case 1:
                v5 = (v5 ^ v5)
            default:
                v6 := []int{-2147483648, 42, -648, -589}
                _ = v6
            }
            if (((0.0 - (-2.25 + ((((-2.25 / 1.5) + 0.0) * ((-2.25 * 0.1) / (float64((float32((3.14159 / -2.25)) - 1.5)) / 1e10))) - 0.1))) != -2.25) && (v5 > v5)) {
                v3 = ""
                fmt.Println(v3)
            }
            var v7 uint32 = 0
            _ = v7
        } else {
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int32 = v4
                v10 := []int{300, 32767, 9223372036854775807, (-(fi8))}
                if false {
                    switch (fi8 | 158) {
                    case -318:
                        if len(v10) > 0 {
                            v10[0] = 128
                        }
                    case -9223372036854775808:
                        _ = len(v2)
                    case 127:
                        v11 := GenS1{F0: 2, F1: 42}
                        _ = v11
                    default:
                        if (false || true) {
                            v12 := GenS1{F0: 604, F1: 0}
                            _ = v12
                        }
                    }
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v14 := GenS2{F0: v5, F1: -128, F2: 98}
                        var v15 uint8 = 127
                        v16 := GenS3{F0: 1}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 uint64 = 42
                    _ = v17
                } else {
                    switch fi8 {
                    case 9223372036854775807:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            var v19 float32 = (-2.25 - 0.0)
                            _ = fi18
                            _ = v19
                        }
                    case 33:
                        delete(v2, fi8)
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS2{F0: (v5 >> ((2) & 63)), F1: -128, F2: -77}
                v22 := GenS2{F0: (^(v5)), F1: 127, F2: -128}
                switch 2147483647 {
                case 42:
                    v5 -= v5
                case 0:
                    var v23 uint32 = (v5 * (v5 + v5))
                    _ = v23
                case 127:
                    switch v3 {
                    case "foo bar":
                        switch v3 {
                        case "foo bar":
                            if false {
                                var v24 int64 = 9223372036854775807
                                _ = v24
                            } else {
                                var v25 uint32 = 2
                                switch 580 {
                                case 127:
                                    fmt.Println(v4)
                                case -2147483648:
                                    var v26 int64 = 0
                                    _ = v26
                                case 590:
                                    var v27 uint16 = 861
                                    _ = v27
                                }
                                _ = v25
                            }
                        case "a":
                            switch fi20 {
                            case -2147483648:
                                if (0 > -113) {
                                    v28 := GenS4{F0: 2, F1: (v4 & 4)}
                                    v5 -= v5
                                    _ = v28
                                } else {
                                    fmt.Println(v3)
                                    switch (("\n\t" + v3) + "hello") {
                                    case "\n\t":
                                        if false {
                                            if true {
                                                v4 -= -714
                                                var v29 int64 = 868
                                                _ = v29
                                            }
                                        } else {
                                            v30 := []int{32768, fi20, (fi20 * (fi20 >> ((uint(v5)) & 63))), 128}
                                            v3 = (v3 + v3)
                                            _ = v30
                                        }
                                    case "foo bar":
                                        v31 := GenS5{F0: ("\n\t" + v3), F1: 589}
                                        _ = v31
                                    case "":
                                        var v32 int = (^(fi20))
                                        _ = v32
                                    }
                                }
                            case 32768:
                                v33 := GenS6{F0: 716}
                                _ = v33
                            default:
                                if (1 == 0) {
                                    if (v4 < (v4 ^ (v4 & (v4 / ((-2147483648) | 1))))) {
                                        fmt.Println(v5, v3)
                                    } else {
                                        var v34 float64 = float64((-2.25 + 1e10))
                                        fmt.Println(v34, v3, v4)
                                        var v35 uint64 = 2
                                        _ = v34
                                        _ = v35
                                    }
                                    for fi36 := 0; fi36 < 3; fi36++ {
                                        var v37 float32 = 3.14159
                                        _ = fi36
                                        _ = v37
                                    }
                                } else {
                                    for fi38 := 0; fi38 < 20; fi38++ {
                                        fmt.Println(v4, fi20)
                                        var v39 float64 = ((1.5 * (-2.25 * 1e10)) / (-2.25 / (((((0.1 * -2.25) * (-2.25 + 3.14159)) - 0.0) + -2.25) / 3.14159)))
                                        v39 = ((1e10 + v39) + (v39 / v39))
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            }
                        default:
                            var v40 uint64 = 4294967295
                            _ = v40
                        }
                    default:
                        if false {
                            switch (fi20 >> ((970) & 63)) {
                            case 2147483647:
                                switch (fi20 & 0) {
                                case 0:
                                    if (false || true) {
                                        var v41 int32 = (v4 / ((v4) | 1))
                                        _ = v41
                                    } else {
                                        v42 := []int{int(v5), fi20}
                                        _ = v42
                                    }
                                }
                            case -729:
                                v3 = "\n\t"
                            default:
                                if false {
                                    v5 -= (^(v5))
                                    var v43 int8 = -3
                                    _ = v43
                                } else {
                                    switch -432 {
                                    case 32767:
                                        for fi44 := 0; fi44 < 9; fi44++ {
                                            if (3.14159 <= (1.5 + (float32(0.0) * 1.5))) {
                                                v2[(fi20 << ((195) & 63))] = "foo bar"
                                            } else {
                                                _ = len(v1)
                                                v45 := GenS7{F0: (v3 + (("\n\t" + (v3 + "")) + v3))}
                                                v2[519] = "foo bar"
                                                if true {
                                                    fmt.Println(fi44, fi20, fi20)
                                                    delete(v1, ("hello" + ""))
                                                    var v46 int8 = 1
                                                    _ = v46
                                                } else {
                                                    switch fi44 {
                                                    case 32768:
                                                        v47 := GenS7{F0: ((v3 + v3) + "foo bar")}
                                                        _ = v47
                                                    }
                                                    v3 = "a"
                                                    fmt.Println(fi20)
                                                }
                                                v5 -= 932
                                                _ = v45
                                            }
                                            _ = fi44
                                        }
                                    }
                                }
                            }
                            v3 = ("" + v3)
                            v48 := GenS4{F0: 1, F1: (v4 & -915)}
                            _ = v48
                        } else {
                            var v49 int64 = 42
                            switch "" {
                            case "":
                                var v50 uint32 = (v5 | v5)
                                _ = v50
                            case "a":
                                v5 = (v5 >> ((255) & 63))
                            default:
                                if false {
                                    var v51 string = v3
                                    _ = v51
                                }
                            }
                            for fi52 := 0; fi52 < 10; fi52++ {
                                v53 := GenS8{F0: (-168 == int(v4)), F1: v5}
                                _ = fi52
                                _ = v53
                            }
                            _ = v49
                        }
                    }
                default:
                    for fi54 := 0; fi54 < 19; fi54++ {
                        var v55 int64 = 9223372036854775807
                        v56 := GenS9{F0: 2147483647, F1: fi54}
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        v1[((("" + "hello") + v3) + (v3 + v3))] = v3
        _ = v5
    }
    var v57 float64 = (float64(float32(((0.0 - -2.25) / 1.0))) + 1e10)
    switch 127 {
    case 127:
        v57 = (((float64(1.5) * v57) / 0.1) + v57)
    case 32768:
        switch 4294967295 {
        case 57:
            v58 := GenS6{F0: 814}
            _ = v58
        case 127:
            v59 := GenS10{F0: v4}
            _ = v59
        }
    case -9223372036854775808:
        v60 := []int{-1, 2147483647, 9223372036854775807, int(v4)}
        _ = v60
    default:
        var v61 uint32 = 291
        _ = v61
    }
    v57 = (v57 + (v57 - 3.14159))
    v4 = (v4 << ((699) & 63))
    fmt.Println(v3, v4, v57)
    var v62 uint32 = 42
    switch v3 {
    case "a":
        if false {
            v63 := GenS1{F0: 1, F1: 834}
            _ = v63
        }
    case "":
        for fi64 := 0; fi64 < 13; fi64++ {
            v65 := GenS9{F0: 2147483647, F1: -24}
            _ = fi64
            _ = v65
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v62
}
