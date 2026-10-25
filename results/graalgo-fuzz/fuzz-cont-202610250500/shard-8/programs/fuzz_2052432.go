package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int8; F1 int64; F2 int32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int16; F1 int; F2 uint16 }
type GenS10 struct { F0 int; F1 uint32 }
func genFn0() string {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{(fi7 & -16), (fi7 >> ((0) & 63))}
        fmt.Println(fi7)
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(fi9, fi9)
            if false {
                v10 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
                v11 := GenS1{F0: 42, F1: 18446744073709551615}
                _ = v10
                _ = v11
            } else {
                fmt.Println(fi7)
                fmt.Println(fi7, fi7)
            }
            switch ("hello" + (("hello" + ("hello" + "a")) + "hello")) {
            case "hello":
                v12 := GenS1{F0: -1, F1: 885}
                _ = v12
            case "a":
                var v13 uint16 = 2
                _ = v13
            default:
                if true {
                    var v14 int8 = 103
                    switch ((((("a" + "foo bar") + "") + ("hello" + ("a" + "hello"))) + "hello") + "\n\t") {
                    case "foo bar":
                        fmt.Println(v14, fi9)
                    }
                    _ = v14
                }
                var v15 int64 = -1
                _ = v15
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 423, F1: (fi16 * fi16)}
        _ = fi16
        _ = v17
    }
    var v18 uint = 0
    fmt.Println(v18)
    switch ("\n\t" + ("a" + ("a" + "hello"))) {
    case "hello":
        v18 = (v18 << (((v18 + (v18 | v18))) & 63))
    case "a":
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(fi19)
            v20 := []int{fi19, 32768, fi19, -567}
            v20 = append(v20, (fi19 >> (((v18 << ((v18) & 63))) & 63)))
            _ = fi19
            _ = v20
        }
    }
    _ = v18
    return "\n\t"
}
func genFn1(p0 int8, p1 string, p2 uint) int16 {
    if false {
        v21 := GenS1{F0: -9223372036854775808, F1: 16}
        for fi22 := 0; fi22 < 8; fi22++ {
            switch "foo bar" {
            case "\n\t":
                v23 := []int{}
                _ = v23
            case "a":
                if (1 < 0) {
                    v24 := []int{1}
                    v25 := GenS1{F0: -1, F1: 115}
                    _ = v24
                    _ = v25
                }
            case "":
                if false {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        p2 -= 179
                        _ = fi26
                    }
                }
            default:
                p0 -= p0
            }
            switch (fi22 | fi22) {
            case 32767:
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(p1)
                    switch "" {
                    case "hello":
                        v28 := GenS1{F0: -9223372036854775808, F1: 154}
                        _ = v28
                    case "foo bar":
                        var v29 uint32 = 4294967295
                        _ = v29
                    case "a":
                        v30 := GenS2{F0: (^(p2)), F1: (fi22 + fi22)}
                        _ = v30
                    }
                    _ = fi27
                }
            case 9223372036854775807:
                switch p1 {
                case "":
                    switch fi22 {
                    case -1:
                        fmt.Println(p1, fi22, fi22)
                    }
                case "\n\t":
                    var v31 int32 = 436
                    _ = v31
                }
            case 127:
                if ((0.0 - -2.25) != 1e10) {
                    if false {
                        v32 := []int{}
                        _ = v32
                    } else {
                        p0 += p0
                    }
                    v33 := GenS3{F0: -96, F1: -59, F2: 0}
                    _ = v33
                }
            }
            v34 := GenS3{F0: p0, F1: 9223372036854775807, F2: -1}
            _ = fi22
            _ = v34
        }
        _ = v21
    }
    p0 = -1
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((1e10 / 0.1) / 1.0)
    var v3 int8 = -1
    var v4 uint8 = 128
    v5 := make(map[uint32]int)
    var v6 int64 = 45
    v5[0] = 1
    v2 = 3.14159
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS2{F0: uint(v4), F1: (fi35 << ((0) & 63))}
        var v37 float64 = v2
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        if (("" > (("a" + "") + ("\n\t" + "hello"))) && ((false && true) && (v3 > (v3 << ((255) & 63))))) {
            var v39 uint = 347
            _ = v39
        } else {
            fmt.Println(v3, fi38)
        }
        _ = fi38
    }
    var v40 float64 = v2
    v41 := []int{}
    switch 1 {
    case 921:
        v4 = 127
    case 337:
        v42 := GenS1{F0: -455, F1: 1}
        _ = v42
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        v3 -= (v3 - v3)
        if false {
            v40 += ((v2 - v40) + v40)
            v44 := GenS4{F0: v3, F1: (v6 % (((v6 ^ 1)) | 1))}
            for fi45 := 0; fi45 < 12; fi45++ {
                switch ("\n\t" + ("a" + "hello")) {
                case "hello":
                    delete(v1, (fi43 + -9223372036854775808))
                case "a":
                    if true {
                        if ((uint(v4) <= 58) || false) {
                            switch fi43 {
                            case 1:
                                v2 -= ((0.0 + v40) - v40)
                            case 0:
                                if false {
                                    var v46 float32 = (1.5 / 1.0)
                                    var v47 float64 = (v40 / v40)
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                    var v48 int8 = v3
                                    delete(v1, fi45)
                                    if false {
                                        if len(v41) >= 2 {
                                            v41 = v41[1:len(v41)]
                                        }
                                    } else {
                                        v6 = v6
                                        var v49 int32 = 0
                                        if (v4 >= 128) {
                                            v46 = ((0.1 * v46) - 0.0)
                                            var v50 int8 = (v48 | v48)
                                            v51 := GenS1{F0: v6, F1: 0}
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v2 -= v40
                                            v52 := GenS4{F0: 127, F1: v6}
                                            if true {
                                                switch (-(fi43)) {
                                                case 894:
                                                    switch ((("hello" + "\n\t") + ("hello" + "foo bar")) + ("foo bar" + "\n\t")) {
                                                    case "\n\t":
                                                        if false {
                                                            if true {
                                                                v48 += (v48 << ((4294967295) & 63))
                                                            }
                                                        } else {
                                                            switch fi45 {
                                                            case 42:
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                    case "hello":
                                                        var v53 uint8 = 0
                                                        _ = v53
                                                    }
                                                case 0:
                                                    switch fi43 {
                                                    case -255:
                                                        var v54 int16 = -32768
                                                        _ = v54
                                                    case 898:
                                                        if len(v41) >= 2 {
                                                            v41 = v41[1:len(v41)]
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v55 int8 = v3
                                                if (0 <= 525) {
                                                    v2 = (0.1 / v2)
                                                    var v56 int16 = -980
                                                    _ = v56
                                                } else {
                                                    v49 = v49
                                                    v40 = v40
                                                    switch 149 {
                                                    case 9223372036854775807:
                                                        v57 := GenS5{F0: 1, F1: 4294967295}
                                                        _ = v57
                                                    case -1:
                                                        var v58 string = "foo bar"
                                                        _ = v58
                                                    default:
                                                        v41 = append(v41, int(v4))
                                                    }
                                                }
                                                v46 = 0.0
                                                _ = v55
                                            }
                                            _ = v52
                                        }
                                        var v59 uint16 = 32767
                                        _ = v49
                                        _ = v59
                                    }
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                } else {
                                    _ = len(v41)
                                    if false {
                                        _ = len(v41)
                                        fmt.Println(v4, v6)
                                        var v60 uint64 = 0
                                        _ = v60
                                    } else {
                                        v61 := GenS6{F0: 338}
                                        var v62 bool = (481 >= 42)
                                        fmt.Println(v2)
                                        v63 := GenS7{F0: ("\n\t" + ("hello" + "\n\t")), F1: uint16(v4)}
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                }
                            default:
                                if true {
                                    switch (fi45 * fi45) {
                                    case 32767:
                                        var v64 bool = (uint16(v4) >= 258)
                                        _ = v64
                                    case -708:
                                        v65 := GenS8{F0: genFn1(127, (("a" + "hello") + "\n\t"), 4294967295), F1: 446}
                                        _ = v65
                                    case 127:
                                        v66 := GenS5{F0: (v6 + int64(fi45)), F1: 110}
                                        _ = v66
                                    }
                                }
                            }
                            v4 += v4
                            fmt.Println(fi43, v40, v6)
                        } else {
                            var v67 float32 = 0.0
                            _ = v67
                        }
                        var v68 int64 = v6
                        _ = v68
                    }
                    switch (fi45 % (((-(fi43))) | 1)) {
                    case 597:
                        if false {
                            var v69 uint64 = 4294967295
                            var v70 float64 = ((v40 / v40) / v40)
                            _ = v69
                            _ = v70
                        }
                    default:
                        v2 += (0.0 / (0.0 + v40))
                    }
                    var v71 int8 = (v3 % ((v3) | 1))
                    _ = v71
                }
                if (385 != 42) {
                    v72 := GenS9{F0: 1, F1: (fi45 << ((18446744073709551615) & 63)), F2: 0}
                    v73 := GenS3{F0: v3, F1: -147, F2: -1}
                    v1[fi43] = fi43
                    _ = v72
                    _ = v73
                }
                v74 := GenS10{F0: 128, F1: 4294967295}
                _ = fi45
                _ = v74
            }
            _ = v44
        } else {
            v4 = v4
            fmt.Println(fi43, v6, v3)
        }
        var v75 int32 = 2147483647
        _ = fi43
        _ = v75
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
