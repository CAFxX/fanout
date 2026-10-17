package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint16; F2 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int32; F1 uint32; F2 int }
type GenS7 struct { F0 uint8; F1 int8; F2 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 int16; F2 int32 }
func genFn0(p0 int16) bool {
    var v4 uint8 = 42
    v4 -= (v4 | (v4 * 128))
    v5 := []int{}
    if (2147483648 == 0) {
        v4 -= v4
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint8 = 255
            if true {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v7 = 127
                    if ((true || ("a" < ("" + ("\n\t" + ("hello" + ("" + "")))))) || ((!false) || (127 == -1))) {
                        if (fi6 <= (fi6 << ((4294967295) & 63))) {
                            var v9 string = ""
                            _ = v9
                        }
                        var v10 uint64 = 42
                        var v11 uint16 = 959
                        _ = v10
                        _ = v11
                    } else {
                        if false {
                            if true {
                                v12 := GenS1{F0: v7, F1: 0, F2: 2}
                                _ = v12
                            } else {
                                v4 = 175
                                var v13 uint8 = v4
                                var v14 int64 = 9223372036854775807
                                _ = v13
                                _ = v14
                            }
                            var v15 float64 = 0.1
                            var v16 float64 = v15
                            var v17 int = fi8
                            if true {
                                switch "" {
                                case "foo bar":
                                    v18 := GenS2{F0: (v4 | (v7 & 1)), F1: v17, F2: (-(v17))}
                                    _ = v18
                                case "hello":
                                    var v19 int = (v17 ^ -1)
                                    _ = v19
                                case "\n\t":
                                    fmt.Println(fi8)
                                    switch "hello" {
                                    case "a":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "":
                                        if (2 < 2) {
                                            v4 += (v4 * v7)
                                        } else {
                                            v20 := GenS3{F0: "hello"}
                                            var v21 uint = 42
                                            switch v17 {
                                            case 1:
                                                fmt.Println(v15, v4)
                                            case -9223372036854775808:
                                                if (!false) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    var v22 uint64 = 0
                                                    switch "hello" {
                                                    case "":
                                                        var v23 string = "a"
                                                        if ((1.5 * 3.14159) == 3.14159) {
                                                            fmt.Println(v22, v16)
                                                            v24 := GenS3{F0: "\n\t"}
                                                            v25 := GenS4{F0: 821, F1: uint16(v21), F2: v21}
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            p0 = (-(p0))
                                                            v26 := GenS3{F0: (((v23 + (("\n\t" + "foo bar") + ("" + "\n\t"))) + v23) + v23)}
                                                            v27 := GenS5{F0: int32(v21), F1: (v23 + v23)}
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        var v28 int16 = -913
                                                        _ = v23
                                                        _ = v28
                                                    case "foo bar":
                                                        var v29 string = "\n\t"
                                                        _ = v29
                                                    case "\n\t":
                                                        if (!(fi8 == (v17 % (((^(fi6))) | 1)))) {
                                                            v16 = 3.14159
                                                        } else {
                                                            var v30 int16 = (^(p0))
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v31 int16 = p0
                                                                _ = v31
                                                            case "":
                                                                v21 = v21
                                                            default:
                                                                var v32 uint32 = 192
                                                                _ = v32
                                                            }
                                                            _ = v30
                                                        }
                                                    default:
                                                        if len(v5) > 0 {
                                                            v5[0] = (fi8 << ((720) & 63))
                                                        }
                                                    }
                                                    v7 += 90
                                                    _ = v22
                                                }
                                            case 943:
                                                switch "hello" {
                                                case "":
                                                    v5 = append(v5, (fi6 ^ (^(fi8))))
                                                default:
                                                    v4 += (v4 + 207)
                                                }
                                                v15 = v15
                                            }
                                            _ = v20
                                            _ = v21
                                        }
                                    case "hello":
                                        if (1 == 1) {
                                            var v33 int = fi6
                                            if true {
                                                v16 = 0.1
                                            } else {
                                                v34 := GenS1{F0: 42, F1: 2147483647, F2: 4294967295}
                                                _ = v34
                                            }
                                            fmt.Println(v15, v33)
                                            _ = v33
                                        } else {
                                            v4 = (^(v4))
                                        }
                                    }
                                    if true {
                                        v17 = (^(v17))
                                        v5 = append(v5, (fi8 - v17))
                                    }
                                }
                                v35 := GenS2{F0: 61, F1: fi8, F2: v17}
                                if len(v5) > 0 {
                                    v5[0] = (v17 << ((2) & 63))
                                }
                                _ = v35
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        switch "" {
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                switch (("foo bar" + "\n\t") + "hello") {
                                case "\n\t":
                                    v36 := GenS6{F0: 42, F1: 4294967295, F2: (fi6 - fi6)}
                                    _ = v36
                                default:
                                    v4 -= (v4 << ((0) & 63))
                                }
                            case "hello":
                                var v37 float32 = 0.1
                                var v38 uint64 = 2
                                v7 = v7
                                _ = v37
                                _ = v38
                            case "":
                                v7 += 57
                            }
                        case "":
                            switch (("\n\t" + ("\n\t" + "a")) + ("\n\t" + "hello")) {
                            case "":
                                var v39 int64 = 42
                                _ = v39
                            case "a":
                                switch int(v7) {
                                case -9223372036854775808:
                                    var v40 int32 = -2147483648
                                    _ = v40
                                default:
                                    fmt.Println(fi6, v7, v7)
                                }
                            default:
                                fmt.Println(fi8)
                            }
                        }
                        if true {
                            v7 += 198
                        } else {
                            fmt.Println(fi8, fi8, fi6)
                        }
                    }
                    v4 = (v4 * (v7 + 127))
                    var v41 int = fi8
                    _ = fi8
                    _ = v41
                }
            } else {
                switch "hello" {
                case "a":
                    v42 := GenS6{F0: 0, F1: 42, F2: fi6}
                    _ = v42
                case "foo bar":
                    var v43 float64 = ((0.0 + 3.14159) * (1e10 + 3.14159))
                    _ = v43
                case "hello":
                    v44 := GenS4{F0: 0, F1: 412, F2: 441}
                    _ = v44
                default:
                    p0 = int16(fi6)
                }
                fmt.Println(v4)
            }
            _ = fi6
            _ = v7
        }
        p0 = p0
    } else {
        var v45 int8 = 0
        _ = v45
    }
    _ = len(v5)
    _ = p0
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]bool)
    var v3 uint = 0
    switch 1 {
    case 128:
        for fi46 := 0; fi46 < 15; fi46++ {
            var v47 bool = true
            switch "hello" {
            case "foo bar":
                fmt.Println(v1, fi46)
            case "\n\t":
                fmt.Println(fi46, v47, v47)
            case "":
                v48 := v2["\n\t"]
                _ = v48
            default:
                for fi49 := 0; fi49 < 12; fi49++ {
                    v50 := GenS7{F0: 0, F1: 37, F2: -684}
                    var v51 float64 = (1.5 * 1e10)
                    var v52 uint = v3
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
            }
            v53 := v2[(v1 + v1)]
            _ = fi46
            _ = v47
            _ = v53
        }
    case 111:
        v3 = (v3 / ((v3) | 1))
    }
    v3 -= (v3 ^ (v3 >> (((v3 & 255)) & 63)))
    var v54 uint64 = 2
    v55 := GenS8{F0: 2}
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 int64 = -9223372036854775808
        _ = fi56
        _ = v57
    }
    v58 := GenS1{F0: 255, F1: 65535, F2: (v3 >> ((1) & 63))}
    switch "a" {
    case "foo bar":
        var v59 bool = true
        _ = v59
    default:
        var v60 uint16 = 0
        _ = v60
    }
    if true {
        for fi61 := 0; fi61 < 8; fi61++ {
            v3 = v3
            v54 -= 0
            _ = fi61
        }
    }
    v62 := GenS2{F0: 38, F1: -1, F2: 42}
    v2[(v1 + (v1 + "a"))] = false
    var v63 string = (v1 + "a")
    v64 := GenS8{F0: 255}
    switch 369 {
    case 1:
        v65 := GenS9{F0: -768, F1: (v3 * (v3 / ((v3) | 1)))}
        _ = v65
    case -750:
        var v66 int = 0
        _ = v66
    case 875:
        v67 := GenS6{F0: 2147483647, F1: 2, F2: 0}
        _ = v67
    default:
        if false {
            fmt.Println(v63, v1, v1)
            if (!false) {
                var v68 uint = 0
                if false {
                    var v69 bool = (true && false)
                    var v70 bool = v69
                    if v69 {
                        v54 -= (v54 + 42)
                        v71 := GenS10{F0: -9223372036854775808}
                        _ = v71
                    }
                    _ = v69
                    _ = v70
                } else {
                    if (v1 < ((v63 + v1) + (v1 + ""))) {
                        v72 := []int{}
                        var v73 float64 = 3.14159
                        _ = v72
                        _ = v73
                    }
                    for fi74 := 0; fi74 < 4; fi74++ {
                        v63 = v63
                        for fi75 := 0; fi75 < 13; fi75++ {
                            if (42 > 32767) {
                                v76 := v2[v1]
                                _ = v76
                            } else {
                                v77 := GenS9{F0: 801, F1: v3}
                                _ = v77
                            }
                            _ = fi75
                        }
                        var v78 bool = true
                        _ = fi74
                        _ = v78
                    }
                }
                _ = v68
            }
        } else {
            v79 := []int{-9223372036854775808, 732, 887, 343}
            delete(v2, v63)
            _ = v79
        }
    }
    v3 += uint(v54)
    v80 := GenS11{F0: 869, F1: 1, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v58
    _ = v62
    _ = v63
    _ = v64
    _ = v80
}
