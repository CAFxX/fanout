package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 int16; F2 int16 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int16; F2 int8 }
type GenS9 struct { F0 int32; F1 string; F2 int16 }
type GenS10 struct { F0 uint16; F1 uint64; F2 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 uint32, p1 string) uint32 {
    var v7 float64 = -2.25
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: uint8(p0), F1: 901}
    v10 := GenS1{F0: 61, F1: 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 int, p1 int32) string {
    var v11 int16 = 1000
    v11 = v11
    _ = p0
    _ = p1
    _ = v11
    return ("foo bar" + "a")
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{(("" + "hello") + "\n\t"), "", "a"}
    var v3 int8 = -128
    var v4 uint16 = 0
    var v5 int = 127
    var v6 bool = true
    v12 := GenS2{F0: 2, F1: 42}
    if ((true || false) && v6) {
        if (v6 && (1.5 == 0.1)) {
            v13 := GenS3{F0: v5}
            fmt.Println(v4, v3)
            _ = v13
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            switch (v5 ^ (v5 & -887)) {
            case 4294967295:
                for fi15 := 0; fi15 < 11; fi15++ {
                    switch v5 {
                    case 755:
                        v6 = v6
                    case 9223372036854775807:
                        switch "hello" {
                        case "\n\t":
                            if v6 {
                                var v16 float64 = (0.0 - 1.5)
                                v2 = append(v2, ("a" + "hello"))
                                _ = v16
                            }
                        }
                        v17 := GenS2{F0: 255, F1: 1}
                        _ = v17
                    }
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        v5 = fi14
                    case "a":
                        switch "foo bar" {
                        case "a":
                            var v18 int16 = 42
                            _ = v18
                        case "":
                            if false {
                                v19 := GenS4{F0: 4294967295, F1: 944, F2: v6}
                                delete(v1, fi14)
                                if true {
                                    if false {
                                        var v20 int32 = 42
                                        v21 := GenS3{F0: (fi14 * v5)}
                                        switch ("a" + "a") {
                                        case "hello":
                                            v6 = (v4 <= (v4 + (v4 | (v4 * 922))))
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = (("foo bar" + "a") + ("a" + "a"))
                                            }
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            if ((0.1 - 0.0) == 0.0) {
                                                switch v5 {
                                                case 68:
                                                    var v22 bool = ((!v6) || (-2.25 <= 3.14159))
                                                    _ = v22
                                                case 9223372036854775807:
                                                    var v23 uint16 = v4
                                                    _ = v23
                                                case 2147483647:
                                                    v1[2147483647] = 737
                                                default:
                                                    if (252 > 2) {
                                                        v6 = (v6 || v6)
                                                        if len(v2) > 0 {
                                                            v2[0] = "\n\t"
                                                        }
                                                        var v24 int64 = 1
                                                        _ = v24
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = (("\n\t" + "hello") + "hello")
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        var v25 bool = v6
                                        _ = v20
                                        _ = v21
                                        _ = v25
                                    } else {
                                        if (v6 && v6) {
                                            v5 -= 32768
                                            fmt.Println(v5, v6)
                                        }
                                        if false {
                                            if (65535 >= 2) {
                                                v3 = -80
                                            } else {
                                                var v26 uint8 = 127
                                                v2 = append(v2, "a")
                                                var v27 string = "a"
                                                _ = v26
                                                _ = v27
                                            }
                                        } else {
                                            v1[(v5 << ((65535) & 63))] = 65535
                                            var v28 uint = 787
                                            if v6 {
                                                switch (v5 + fi14) {
                                                case 127:
                                                    v3 = v3
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        v3 -= (v3 - (v3 >> ((42) & 63)))
                                                    case "\n\t":
                                                        var v29 int64 = 0
                                                        _ = v29
                                                    }
                                                }
                                            } else {
                                                if (32767 < v4) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            }
                                            v28 -= (v28 - (v28 << ((255) & 63)))
                                            v30 := GenS5{F0: -643, F1: 448, F2: 1}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v31 := GenS1{F0: 1, F1: (v28 << ((v28) & 63))}
                                                _ = v31
                                            case "\n\t":
                                                if v6 {
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(v5)
                                                    case "":
                                                        var v32 uint64 = 4294967295
                                                        _ = v32
                                                    }
                                                    v33 := GenS1{F0: 2, F1: v28}
                                                    v1[(fi14 / (((fi14 + fi15)) | 1))] = 65535
                                                    _ = v33
                                                }
                                            default:
                                                var v34 int64 = -459
                                                _ = v34
                                            }
                                            _ = v28
                                            _ = v30
                                        }
                                    }
                                    v1[9223372036854775807] = 0
                                    v35 := GenS6{F0: 42, F1: 42}
                                    _ = v35
                                }
                                switch -2147483648 {
                                case -64:
                                    fmt.Println(v4)
                                default:
                                    v4 -= v4
                                }
                                var v36 string = ""
                                _ = v19
                                _ = v36
                            } else {
                                if (true && (!((0.0 + -2.25) >= 1.5))) {
                                    fmt.Println(v3)
                                    if v6 {
                                        v2 = append(v2, "\n\t")
                                        switch "" {
                                        case "a":
                                            var v37 uint = 0
                                            _ = v37
                                        case "\n\t":
                                            v6 = true
                                        case "hello":
                                            switch 9223372036854775807 {
                                            case 4294967295:
                                                v38 := GenS7{F0: 348}
                                                _ = v38
                                            case -9223372036854775808:
                                                switch (-(fi15)) {
                                                case 4294967295:
                                                    fmt.Println(v6)
                                                case 128:
                                                    if false {
                                                        var v39 int64 = 42
                                                        _ = v39
                                                    }
                                                default:
                                                    v4 += v4
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        default:
                            var v40 uint64 = 600
                            _ = v40
                        }
                        switch "" {
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                v3 -= (v3 - 42)
                            }
                        }
                        switch (fi15 - fi14) {
                        case 0:
                            var v41 float32 = 1.5
                            v3 -= (-(v3))
                            var v42 string = "foo bar"
                            _ = v41
                            _ = v42
                        default:
                            v43 := GenS5{F0: 32767, F1: -1, F2: 1000}
                            _ = v43
                        }
                    case "\n\t":
                        var v44 uint16 = (^(v4))
                        _ = v44
                    }
                    var v45 uint8 = 127
                    _ = fi15
                    _ = v45
                }
            case 433:
                if (!(v6 || (genFn0(354, ("\n\t" + (("a" + (("foo bar" + "hello") + "a")) + ("foo bar" + "\n\t")))) == 596))) {
                    v5 += v5
                    v4 -= 42
                }
            }
            if ((v6 && (float32(0.1) < (-2.25 - -2.25))) || false) {
                fmt.Println(fi14)
            }
            _ = fi14
        }
        var v46 int64 = 0
        _ = v46
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v47 string = ("" + "")
    switch v47 {
    case "hello":
        v5 = (v5 * (v5 + v5))
    default:
        if v6 {
            fmt.Println(v4, v6)
            switch (v5 - v5) {
            case 42:
                if ((!v6) && (v3 == (v3 + -128))) {
                    delete(v1, (v5 | (v5 ^ (v5 % ((9223372036854775807) | 1)))))
                    var v48 int8 = 127
                    if v6 {
                        var v49 bool = (!(1 <= 255))
                        var v50 bool = (v47 > (v47 + (v47 + v47)))
                        v51 := GenS4{F0: 4294967295, F1: 255, F2: (18446744073709551615 != 0)}
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    _ = v48
                } else {
                    v5 = v5
                    switch (v5 / ((2147483647) | 1)) {
                    case 42:
                        switch v5 {
                        case -2147483648:
                            v52 := GenS8{F0: 236, F1: 1, F2: (v3 - 0)}
                            _ = v52
                        case -909:
                            v53 := GenS9{F0: -1, F1: (v47 + "foo bar"), F2: -1}
                            _ = v53
                        case 9223372036854775807:
                            v47 = ("hello" + "hello")
                        }
                    case -541:
                        for fi54 := 0; fi54 < 17; fi54++ {
                            for fi55 := 0; fi55 < 9; fi55++ {
                                if (778 < 42) {
                                    if (0 == 1) {
                                        v56 := GenS5{F0: 1000, F1: 42, F2: 32767}
                                        v5 += fi54
                                        v5 += (^(fi54))
                                        _ = v56
                                    } else {
                                        v5 -= (-(fi54))
                                        v2 = append(v2, (v47 + v47))
                                    }
                                } else {
                                    if v6 {
                                        v57 := GenS7{F0: 42}
                                        v47 = "a"
                                        _ = v57
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    }
                                    var v58 float64 = -2.25
                                    var v59 uint32 = 0
                                    _ = v58
                                    _ = v59
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    switch fi54 {
                                    case -400:
                                        v6 = false
                                    default:
                                        _ = len(v2)
                                    }
                                case "hello":
                                    var v60 int16 = 32767
                                    _ = v60
                                default:
                                    switch -9223372036854775808 {
                                    case 29:
                                        fmt.Println(v6, v5, v4)
                                    case 128:
                                        if (!v6) {
                                            v4 = (v4 ^ 255)
                                        } else {
                                            fmt.Println(v5, v5)
                                        }
                                    default:
                                        fmt.Println(v47)
                                    }
                                }
                                var v61 float64 = -2.25
                                _ = fi55
                                _ = v61
                            }
                            v47 = "a"
                            _ = fi54
                        }
                    default:
                        if true {
                            v62 := v1[32767]
                            _ = v62
                        } else {
                            fmt.Println(v47)
                        }
                    }
                    var v63 uint = 65535
                    _ = v63
                }
            default:
                v64 := GenS7{F0: 9}
                _ = v64
            }
            v5 = (v5 * v5)
        }
    }
    if (-249 > 592) {
        v1[0] = 277
        v5 += (v5 % ((v5) | 1))
    } else {
        if v6 {
            var v65 int64 = -9223372036854775808
            for fi66 := 0; fi66 < 3; fi66++ {
                switch (v5 ^ fi66) {
                case -9223372036854775808:
                    var v67 float32 = 1.5
                    _ = v67
                case 42:
                    v68 := GenS10{F0: v4, F1: 515, F2: "hello"}
                    _ = v68
                default:
                    for fi69 := 0; fi69 < 4; fi69++ {
                        v47 = ("hello" + "a")
                        fmt.Println(v6, v5, v5)
                        _ = fi69
                    }
                }
                v3 -= 118
                for fi70 := 0; fi70 < 13; fi70++ {
                    var v71 uint = 18446744073709551615
                    fmt.Println(v71, v4)
                    var v72 int16 = 708
                    fmt.Println(v72, v3, v71)
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                _ = fi66
            }
            _ = v65
        }
        if (genFn0(65535, "\n\t") >= 714) {
            switch v5 {
            case 0:
                v73 := GenS11{F0: v4}
                _ = v73
            }
            switch (v5 % ((-173) | 1)) {
            case -2147483648:
                switch 334 {
                case 222:
                    var v74 int8 = 103
                    _ = v74
                case 42:
                    var v75 string = (v47 + ("hello" + v47))
                    _ = v75
                case -1:
                    for fi76 := 0; fi76 < 4; fi76++ {
                        v77 := GenS10{F0: v4, F1: 42, F2: ("a" + (("hello" + "foo bar") + v47))}
                        _ = fi76
                        _ = v77
                    }
                }
            case 16:
                switch 2147483647 {
                case 743:
                    var v78 int32 = 42
                    _ = v78
                case 2147483647:
                    if (!v6) {
                        var v79 uint32 = 42
                        _ = v79
                    }
                }
            case -661:
                var v80 int64 = 42
                _ = v80
            }
            v81 := GenS4{F0: 242, F1: 825, F2: v6}
            _ = v81
        }
    }
    v82 := GenS2{F0: 42, F1: -886}
    var v83 uint32 = 849
    if (42 >= 42) {
        v4 -= (v4 - 65535)
        for fi84 := 0; fi84 < 6; fi84++ {
            v85 := GenS12{F0: (v83 << ((611) & 63)), F1: 1, F2: (v47 >= v47)}
            for fi86 := 0; fi86 < 10; fi86++ {
                fmt.Println(v6)
                _ = fi86
            }
            v1[(v5 - 575)] = 951
            _ = fi84
            _ = v85
        }
    } else {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        var v87 bool = v6
        var v88 uint8 = 93
        _ = v87
        _ = v88
    }
    var v89 uint8 = 219
    fmt.Println(v4)
    v90 := GenS6{F0: int8(v89), F1: -360}
    v89 += (v89 / (((v89 & (v89 & 2))) | 1))
    switch "hello" {
    case "a":
        v6 = true
    case "":
        _ = len(v1)
    }
    var v91 float64 = float64(3.14159)
    v83 += v83
    v83 += genFn0((^(v83)), ("foo bar" + (v47 + v47)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v89
    _ = v90
    _ = v91
}
