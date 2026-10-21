package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint64; F1 string; F2 uint32 }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int16 }
func genFn0() int64 {
    switch -167 {
    case 42:
        switch "\n\t" {
        case "":
            var v6 uint8 = 200
            _ = v6
        case "hello":
            var v7 int = -2147483648
            _ = v7
        case "a":
            var v8 float32 = 3.14159
            _ = v8
        default:
            switch 0 {
            case 9223372036854775807:
                var v9 bool = false
                _ = v9
            default:
                var v10 uint32 = 2
                _ = v10
            }
        }
    case -2147483648:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint = 0
            var v13 int32 = -511
            _ = fi11
            _ = v12
            _ = v13
        }
    case 9223372036854775807:
        var v14 bool = true
        _ = v14
    }
    var v15 int8 = -1
    if true {
        if (133 == 4294967295) {
            if (32767 != 235) {
                v16 := GenS1{F0: ("hello" + "\n\t"), F1: 168}
                v17 := []int{2147483647, -1, 9223372036854775807}
                if false {
                    var v18 uint32 = 2147483647
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v18 += (v18 / ((v18) | 1))
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v21 := GenS1{F0: "foo bar", F1: 128}
                            v15 = (-(v15))
                            _ = fi20
                            _ = v21
                        }
                        _ = len(v17)
                        _ = fi19
                    }
                    switch "hello" {
                    case "":
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v15 += (v15 & v15)
                            var v23 uint64 = 638
                            for fi24 := 0; fi24 < 14; fi24++ {
                                switch fi24 {
                                case -9223372036854775808:
                                    v23 = (v23 >> ((2) & 63))
                                }
                                v15 = v15
                                _ = fi24
                            }
                            _ = fi22
                            _ = v23
                        }
                    default:
                        if (!(-554 < -1)) {
                            v15 -= 100
                            for fi25 := 0; fi25 < 17; fi25++ {
                                switch fi25 {
                                case 32767:
                                    switch "hello" {
                                    case "foo bar":
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            v26 := GenS2{F0: 1}
                                            _ = v26
                                        }
                                    case "a":
                                        if true {
                                            v18 = (v18 << ((262) & 63))
                                            var v27 int64 = 9223372036854775807
                                            _ = v27
                                        }
                                    case "\n\t":
                                        v28 := GenS1{F0: ("a" + ""), F1: 0}
                                        _ = v28
                                    }
                                case 127:
                                    if true {
                                        var v29 uint32 = 42
                                        v17 = append(v17, fi25)
                                        _ = v29
                                    } else {
                                        if false {
                                            if (false && (2 <= 80)) {
                                                if (true && (661 >= 269)) {
                                                    v30 := GenS3{F0: -1}
                                                    if false {
                                                        var v31 int64 = 42
                                                        for fi32 := 0; fi32 < 16; fi32++ {
                                                            fmt.Println(v15, v15)
                                                            fmt.Println(fi32, v31, v31)
                                                            v33 := GenS3{F0: 457}
                                                            _ = fi32
                                                            _ = v33
                                                        }
                                                        _ = v31
                                                    }
                                                    v34 := GenS2{F0: 42}
                                                    _ = v30
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: 1}
                                                _ = v35
                                            } else {
                                                switch 496 {
                                                case 127:
                                                    if false {
                                                        v15 -= (v15 >> ((277) & 63))
                                                        if len(v17) >= 2 {
                                                            v17 = v17[1:len(v17)]
                                                        }
                                                    } else {
                                                        v36 := GenS4{F0: 1}
                                                        _ = v36
                                                    }
                                                case -9223372036854775808:
                                                    v15 -= int8(fi25)
                                                }
                                            }
                                            var v37 int8 = v15
                                            _ = v37
                                        }
                                    }
                                default:
                                    v18 += v18
                                }
                                v38 := GenS1{F0: ("\n\t" + ""), F1: 1}
                                _ = fi25
                                _ = v38
                            }
                            if len(v17) > 0 {
                                v17[0] = -786
                            }
                        }
                    }
                    _ = v18
                } else {
                    v15 += (v15 << ((2) & 63))
                }
                _ = v16
                _ = v17
            }
            v15 -= (v15 ^ (v15 | (v15 >> ((2) & 63))))
        }
        if true {
            var v39 bool = (1 <= -9223372036854775808)
            v15 = (v15 | v15)
            if v39 {
                var v40 string = "foo bar"
                v41 := []int{-205, 42}
                _ = v40
                _ = v41
            }
            _ = v39
        }
        v15 = v15
    }
    var v42 uint8 = 2
    _ = v15
    _ = v42
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 string, p2 int32) uint8 {
    v43 := []int{-781, 1}
    if false {
        var v44 string = (p1 + "")
        _ = v44
    }
    v43 = append(v43, 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 2
    v3 := make(map[int]string)
    v4 := []uint32{4294967295, 470, 973, 2147483647}
    var v5 float32 = 0.1
    v5 += 1e10
    v2 += v2
    var v45 float64 = ((float64(v5) - -2.25) * 0.1)
    v46 := v3[32767]
    if (440 == 0) {
        for fi47 := 0; fi47 < 13; fi47++ {
            switch 127 {
            case -420:
                fmt.Println(v5, v2, v5)
            case 128:
                if true {
                    v48 := GenS5{F0: 42}
                    _ = v48
                }
            }
            v45 = v45
            _ = fi47
        }
        var v49 uint16 = 426
        if false {
            v46 = (v46 + v46)
        } else {
            switch 32768 {
            case 32767:
                v50 := GenS6{F0: (false || false)}
                _ = v50
            case 267:
                if (32767 < -718) {
                    if ((-333 >= -1) || false) {
                        if (!true) {
                            if true {
                                if true {
                                    for fi51 := 0; fi51 < 16; fi51++ {
                                        fmt.Println(v45, v45)
                                        v4 = append(v4, 305)
                                        fmt.Println(fi51)
                                        _ = fi51
                                    }
                                }
                                var v52 int32 = 917
                                _ = v52
                            }
                            var v53 uint32 = 604
                            for fi54 := 0; fi54 < 16; fi54++ {
                                v55 := GenS5{F0: (v2 + (v2 / (((v2 << ((0) & 63))) | 1)))}
                                v3[32768] = v46
                                _ = fi54
                                _ = v55
                            }
                            _ = v53
                        } else {
                            var v56 string = ("" + (v46 + v46))
                            var v57 float32 = (v5 - (0.1 / (3.14159 / v5)))
                            _ = v56
                            _ = v57
                        }
                        var v58 int32 = 2147483647
                        _ = v58
                    }
                    v4 = append(v4, 42)
                }
            case 127:
                v49 = v49
            default:
                for fi59 := 0; fi59 < 2; fi59++ {
                    for fi60 := 0; fi60 < 18; fi60++ {
                        if (!(!(!(false || ((true || false) || true))))) {
                            var v61 uint8 = genFn1(4294967295, v46, 1)
                            v62 := GenS7{F0: 42, F1: 315}
                            if true {
                                v63 := GenS4{F0: -388}
                                if len(v4) > 0 {
                                    v4[0] = 262
                                }
                                v49 = 42
                                var v64 float64 = (float64(float32(1.5)) + v45)
                                _ = v63
                                _ = v64
                            }
                            _ = v61
                            _ = v62
                        } else {
                            v46 = v46
                            fmt.Println(v49, v46, v45)
                            switch v46 {
                            case "\n\t":
                                switch v46 {
                                case "hello":
                                    var v65 int = (fi60 ^ (fi59 - 640))
                                    _ = v65
                                case "a":
                                    var v66 uint = 289
                                    _ = v66
                                case "foo bar":
                                    switch v46 {
                                    case "\n\t":
                                        fmt.Println(fi60, v49)
                                    }
                                default:
                                    _ = len(v4)
                                }
                            case "hello":
                                v67 := GenS8{F0: 0, F1: ("foo bar" + ((("hello" + ("a" + (v46 + (v46 + v46)))) + ((("foo bar" + v46) + (("\n\t" + v46) + v46)) + ((v46 + "") + (v46 + (v46 + ""))))) + "a")), F2: 200}
                                _ = v67
                            default:
                                v46 = (((((("hello" + v46) + v46) + "foo bar") + v46) + v46) + v46)
                            }
                            _ = len(v4)
                        }
                        if false {
                            v45 += v45
                        } else {
                            v68 := GenS2{F0: -892}
                            v5 += -2.25
                            switch v46 {
                            case "\n\t":
                                switch (v46 + v46) {
                                case "a":
                                    if false {
                                        v69 := GenS9{F0: 1, F1: v46}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        delete(v3, fi60)
                                        _ = v69
                                    }
                                default:
                                    _ = len(v4)
                                }
                            case "":
                                if (!((!(true && (-11 > 1))) || (false || true))) {
                                    v5 = v5
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = len(v4)
                                } else {
                                    switch "hello" {
                                    case "":
                                        if (uint64(v49) == 4294967295) {
                                            fmt.Println(v46, fi59)
                                            v2 += v2
                                            var v70 uint8 = (v2 << ((255) & 63))
                                            _ = v70
                                        } else {
                                            var v71 int16 = 0
                                            _ = v71
                                        }
                                    case "a":
                                        v72 := GenS10{F0: v2, F1: uint16(v2)}
                                        _ = v72
                                    case "\n\t":
                                        var v73 int32 = -326
                                        _ = v73
                                    default:
                                        v74 := GenS3{F0: 9223372036854775807}
                                        _ = v74
                                    }
                                }
                            case "foo bar":
                                _ = len(v4)
                            default:
                                v5 = (((v5 + (1.5 / v5)) - v5) / (v5 / v5))
                            }
                            _ = v68
                        }
                        _ = fi60
                    }
                    var v75 bool = false
                    v75 = (2 >= uint32(v2))
                    _ = fi59
                    _ = v75
                }
            }
            var v76 uint64 = 2
            _ = v76
        }
        _ = v49
    } else {
        var v77 uint8 = (v2 % ((genFn1(2, "hello", 42)) | 1))
        v46 = ("" + v46)
        _ = v77
    }
    for fi78 := 0; fi78 < 11; fi78++ {
        var v79 int32 = 2147483647
        for fi80 := 0; fi80 < 14; fi80++ {
            v1[(fi78 >> ((1) & 63))] = 42
            fmt.Println(v2, fi80)
            v2 -= genFn1(4294967295, ((((((((v46 + (("foo bar" + v46) + v46)) + ("\n\t" + v46)) + "\n\t") + v46) + (v46 + ("foo bar" + v46))) + v46) + v46) + "foo bar"), (v79 % ((v79) | 1)))
            _ = fi80
        }
        _ = fi78
        _ = v79
    }
    if false {
        switch 4294967295 {
        case 128:
            for fi81 := 0; fi81 < 2; fi81++ {
                v2 = v2
                _ = fi81
            }
        }
        var v82 string = v46
        if (!(!(!false))) {
            for fi83 := 0; fi83 < 17; fi83++ {
                var v84 int32 = 42
                _ = fi83
                _ = v84
            }
            v85 := GenS3{F0: 9223372036854775807}
            _ = v85
        } else {
            var v86 bool = ((true && (0 < 9223372036854775807)) || false)
            fmt.Println(v5, v82, v46)
            _ = v86
        }
        _ = v82
    } else {
        if true {
            fmt.Println(v2)
        }
        v46 = (v46 + (v46 + (v46 + (v46 + ("" + ("a" + "foo bar"))))))
        for fi87 := 0; fi87 < 20; fi87++ {
            var v88 int = fi87
            switch v46 {
            case "\n\t":
                for fi89 := 0; fi89 < 13; fi89++ {
                    v2 = v2
                    fmt.Println(fi87, v46)
                    _ = fi89
                }
            case "a":
                var v90 uint64 = 2
                _ = v90
            }
            _ = fi87
            _ = v88
        }
    }
    for fi91 := 0; fi91 < 8; fi91++ {
        v92 := GenS11{F0: -32768}
        v2 += (v2 + (^(v2)))
        _ = fi91
        _ = v92
    }
    if (645 >= 42) {
        switch ("foo bar" + "") {
        case "":
            var v93 int = 0
            _ = v93
        case "\n\t":
            var v94 uint16 = 0
            _ = v94
        case "a":
            fmt.Println(v2)
        }
    }
    switch (((v46 + (v46 + (("hello" + "foo bar") + v46))) + ("\n\t" + "")) + "\n\t") {
    case "hello":
        v46 = v46
    case "a":
        if (100 <= 42) {
            var v95 int16 = -1
            v96 := GenS9{F0: 0, F1: v46}
            v2 = v2
            _ = v95
            _ = v96
        }
    case "\n\t":
        v46 = "hello"
    }
    var v97 float32 = float32(v45)
    switch "" {
    case "a":
        v4 = append(v4, 1)
    case "hello":
        v97 = 1.5
    }
    fmt.Println(v46, v97, v46)
    var v98 float32 = v97
    fmt.Println(v5, v46, v97)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v97
    _ = v98
}
