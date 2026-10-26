package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int16; F1 string; F2 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0() int8 {
    var v7 bool = ((255 < 128) || (!(211 <= 128)))
    if v7 {
        v8 := []int{-385, 563, -596}
        var v9 float64 = float64(3.14159)
        _ = v8
        _ = v9
    }
    v10 := []int{2147483647}
    _ = v10
    _ = v7
    return 42
}
func genFn1() int32 {
    v11 := []int{2147483647, 42, 9223372036854775807}
    if ((((0.1 + 0.0) * float32(3.14159)) - float32(1.5)) < 3.14159) {
        var v12 string = ("" + "hello")
        var v13 float64 = (float64((0.1 - 1.5)) - -2.25)
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: (!true), F1: "a", F2: -2147483648}
        _ = len(v11)
        _ = v14
    }
    var v15 uint8 = 18
    _ = v11
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{199, 4294967295, 42, 0}
    v2 := make(map[int]string)
    var v3 uint8 = 2
    var v4 uint = 255
    v5 := make(map[int]int)
    v6 := []string{"\n\t", ("hello" + ("a" + "foo bar")), ("foo bar" + "foo bar"), ("\n\t" + "hello")}
    v16 := GenS1{F0: (108 <= 65535), F1: "foo bar", F2: 42}
    switch 2147483647 {
    case 127:
        v17 := GenS2{F0: -1, F1: -1}
        _ = v17
    }
    if (!true) {
        var v18 uint32 = 4294967295
        if true {
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint = 65535
                fmt.Println(v18, v3)
                var v21 uint32 = v18
                _ = fi19
                _ = v20
                _ = v21
            }
            fmt.Println(v4, v18)
        }
        var v22 int16 = 1
        _ = v18
        _ = v22
    }
    v23 := GenS3{F0: -1}
    if (65535 >= 2) {
        if false {
            for fi24 := 0; fi24 < 9; fi24++ {
                _ = len(v5)
                _ = fi24
            }
            var v25 uint32 = 42
            _ = v25
        } else {
            var v26 uint32 = 26
            _ = v26
        }
    } else {
        if (true || true) {
            switch 0 {
            case 42:
                if (0.0 < float64(((1e10 * 1e10) - 3.14159))) {
                    switch "" {
                    case "\n\t":
                        delete(v5, 32767)
                    case "foo bar":
                        if false {
                            v4 += (^(v4))
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = "hello"
                                }
                                if (false || false) {
                                    v3 += (v3 & (v3 * 128))
                                    var v27 uint64 = 349
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        var v29 uint32 = 0
                                        switch 32767 {
                                        case -1:
                                            if (0 >= 0) {
                                                for fi30 := 0; fi30 < 5; fi30++ {
                                                    switch (^(fi30)) {
                                                    case -2147483648:
                                                        v3 -= v3
                                                    default:
                                                        var v31 uint = uint(v3)
                                                        _ = v31
                                                    }
                                                    switch -9223372036854775808 {
                                                    case -1:
                                                        var v32 uint64 = v27
                                                        _ = v32
                                                    case 42:
                                                        var v33 string = "hello"
                                                        _ = v33
                                                    case 2147483647:
                                                        var v34 uint8 = (v3 | 5)
                                                        _ = v34
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        fmt.Println(fi30)
                                                    default:
                                                        v35 := GenS3{F0: genFn0()}
                                                        _ = v35
                                                    }
                                                    _ = fi30
                                                }
                                                var v36 int = fi28
                                                for fi37 := 0; fi37 < 6; fi37++ {
                                                    v38 := GenS4{F0: 127, F1: (fi37 - (fi37 + fi37))}
                                                    _ = fi37
                                                    _ = v38
                                                }
                                                _ = v36
                                            }
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if true {
                                                v29 = v29
                                                v3 -= v3
                                                var v39 uint32 = (^(v29))
                                                _ = v39
                                            }
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                    _ = v27
                                } else {
                                    if (!(263 <= 42)) {
                                        var v40 int64 = 42
                                        if (!false) {
                                            var v41 int64 = v40
                                            _ = v41
                                        }
                                        _ = v40
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if false {
                                                if false {
                                                    fmt.Println(v3, v3)
                                                } else {
                                                    v42 := v2[-1]
                                                    _ = v42
                                                }
                                            }
                                        case "foo bar":
                                            switch "hello" {
                                            case "hello":
                                                switch "hello" {
                                                case "foo bar":
                                                    var v43 uint32 = 664
                                                    _ = v43
                                                }
                                            }
                                        default:
                                            for fi44 := 0; fi44 < 1; fi44++ {
                                                v45 := GenS5{F0: 1, F1: "foo bar", F2: 2}
                                                _ = fi44
                                                _ = v45
                                            }
                                        }
                                    }
                                    var v46 int = 32768
                                    v4 += (v4 * (v4 - v4))
                                    _ = v46
                                }
                                for fi47 := 0; fi47 < 4; fi47++ {
                                    switch (-(fi47)) {
                                    case -148:
                                        fmt.Println(v4, v3, v4)
                                    case -1:
                                        v4 = 255
                                    case 429:
                                        for fi48 := 0; fi48 < 11; fi48++ {
                                            v3 = 128
                                            var v49 float32 = (3.14159 + -2.25)
                                            _ = fi48
                                            _ = v49
                                        }
                                    default:
                                        _ = len(v1)
                                    }
                                    v50 := GenS5{F0: 736, F1: "\n\t", F2: 1}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi47
                                    _ = v50
                                }
                            } else {
                                if true {
                                    v4 += 0
                                } else {
                                    var v51 bool = (true && false)
                                    _ = v51
                                }
                            }
                            if (413 >= 9223372036854775807) {
                                v52 := GenS1{F0: false, F1: "", F2: 39}
                                v4 += (v4 / (((v4 << ((v4) & 63))) | 1))
                                _ = v52
                            } else {
                                switch -2147483648 {
                                case -1:
                                    v4 += (v4 % ((734) | 1))
                                case 9223372036854775807:
                                    for fi53 := 0; fi53 < 14; fi53++ {
                                        fmt.Println(v3)
                                        _ = fi53
                                    }
                                }
                                if false {
                                    for fi54 := 0; fi54 < 6; fi54++ {
                                        _ = len(v2)
                                        _ = fi54
                                    }
                                }
                            }
                        } else {
                            var v55 int8 = 127
                            switch "hello" {
                            case "hello":
                                v56 := GenS6{F0: v3}
                                _ = v56
                            case "\n\t":
                                if (341 > 42) {
                                    var v57 int64 = -795
                                    switch "foo bar" {
                                    case "":
                                        v2[354] = "hello"
                                    default:
                                        var v58 uint64 = 42
                                        _ = v58
                                    }
                                    _ = v57
                                } else {
                                    v55 = (v55 / (((v55 | v55)) | 1))
                                }
                            case "foo bar":
                                var v59 uint8 = (v3 & (v3 % ((v3) | 1)))
                                _ = v59
                            default:
                                v60 := GenS5{F0: -274, F1: "", F2: 758}
                                _ = v60
                            }
                            _ = v55
                        }
                    case "hello":
                        var v61 uint64 = 4294967295
                        _ = v61
                    default:
                        fmt.Println(v3, v3)
                    }
                    for fi62 := 0; fi62 < 13; fi62++ {
                        v63 := GenS7{F0: (v4 >> ((uint(v3)) & 63)), F1: 0}
                        v4 += v4
                        fmt.Println(v4, v4, fi62)
                        _ = fi62
                        _ = v63
                    }
                    if (0 >= 426) {
                        for fi64 := 0; fi64 < 12; fi64++ {
                            v3 = 128
                            v65 := GenS8{F0: 65535, F1: 461}
                            _ = fi64
                            _ = v65
                        }
                    } else {
                        fmt.Println(v3, v4)
                        fmt.Println(v4, v3)
                    }
                } else {
                    v3 = 42
                }
            case 1:
                v3 -= 173
            case -4:
                for fi66 := 0; fi66 < 4; fi66++ {
                    v3 = (v3 / ((v3) | 1))
                    _ = fi66
                }
            }
        }
    }
    if true {
        v4 = v4
        fmt.Println(v4)
        var v67 float32 = 0.0
        _ = v67
    }
    for fi68 := 0; fi68 < 1; fi68++ {
        v69 := GenS1{F0: (!(!true)), F1: (("\n\t" + ("hello" + "foo bar")) + "a"), F2: 0}
        v70 := GenS4{F0: 0, F1: fi68}
        _ = fi68
        _ = v69
        _ = v70
    }
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    var v71 int8 = 0
    var v72 uint16 = 503
    if (0 != 187) {
        v4 = (^(v4))
        if (280 == 0) {
            switch 127 {
            case 2147483647:
                for fi73 := 0; fi73 < 3; fi73++ {
                    v4 = 495
                    switch (("" + "hello") + (("hello" + "a") + "")) {
                    case "foo bar":
                        for fi74 := 0; fi74 < 19; fi74++ {
                            var v75 uint8 = (v3 | v3)
                            switch ((("" + "") + "a") + (((("foo bar" + "foo bar") + "hello") + ("" + (("" + ("hello" + ("\n\t" + ("foo bar" + ("hello" + "a"))))) + "\n\t"))) + "a")) {
                            case "hello":
                                fmt.Println(v75)
                            case "\n\t":
                                if true {
                                    v3 -= (^(v75))
                                    switch ("foo bar" + "hello") {
                                    case "":
                                        var v76 int64 = 42
                                        _ = v76
                                    case "foo bar":
                                        fmt.Println(v3, v75)
                                    default:
                                        var v77 float64 = (1.5 - float64(3.14159))
                                        _ = v77
                                    }
                                    v71 -= v71
                                }
                            case "":
                                var v78 int = (fi74 - (fi73 / (((fi74 - (^(fi73)))) | 1)))
                                _ = v78
                            default:
                                v75 = uint8(v4)
                            }
                            _ = fi74
                            _ = v75
                        }
                    }
                    v72 = v72
                    _ = fi73
                }
            default:
                switch 32767 {
                case 464:
                    var v79 string = "hello"
                    _ = v79
                case 32767:
                    if true {
                        var v80 bool = (((float32(-2.25) + (float32(1.5) - (float32(0.1) - 1e10))) - -2.25) != float32(3.14159))
                        for fi81 := 0; fi81 < 10; fi81++ {
                            if (v80 && v80) {
                                v71 = v71
                                if false {
                                    v82 := GenS9{F0: -1}
                                    if false {
                                        v83 := GenS9{F0: (^(v71))}
                                        v1 = append(v1, 664)
                                        _ = v83
                                    } else {
                                        v84 := GenS4{F0: (^(v71)), F1: fi81}
                                        var v85 float32 = 3.14159
                                        var v86 bool = (0 != 913)
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    }
                                    _ = v82
                                }
                                delete(v2, (fi81 >> (((v4 + (v4 % ((v4) | 1)))) & 63)))
                            } else {
                                var v87 uint16 = v72
                                var v88 string = "foo bar"
                                for fi89 := 0; fi89 < 3; fi89++ {
                                    if true {
                                        v87 -= 176
                                        switch v88 {
                                        case "foo bar":
                                            fmt.Println(v87, v72, v88)
                                        case "\n\t":
                                            switch (v88 + (v88 + ("" + (v88 + v88)))) {
                                            case "a":
                                                var v90 uint16 = (v72 & (^(v72)))
                                                _ = v90
                                            }
                                        case "":
                                            v3 += (v3 & v3)
                                        }
                                    } else {
                                        if false {
                                            var v91 uint32 = 2147483647
                                            v87 = 42
                                            _ = v91
                                        } else {
                                            if (fi89 <= (fi81 | int(v87))) {
                                                v92 := GenS8{F0: 2147483648, F1: -2147483648}
                                                _ = v92
                                            }
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        switch 533 {
                                        case -998:
                                            v93 := GenS8{F0: 2, F1: int32(v4)}
                                            _ = v93
                                        case -1:
                                            if (!true) {
                                                v87 += 1
                                            }
                                        case 128:
                                            fmt.Println(fi89, fi81)
                                        default:
                                            switch fi81 {
                                            case -9223372036854775808:
                                                switch "" {
                                                case "":
                                                    _ = len(v2)
                                                default:
                                                    delete(v2, (fi81 - (fi81 | 42)))
                                                }
                                            case 0:
                                                switch (fi89 + (-(fi89))) {
                                                case -9223372036854775808:
                                                    fmt.Println(fi81)
                                                case 127:
                                                    var v94 float32 = 3.14159
                                                    v95 := GenS10{F0: -628, F1: v87, F2: 255}
                                                    _ = v94
                                                    _ = v95
                                                case -133:
                                                    if true {
                                                        switch -283 {
                                                        case 32768:
                                                            switch -9223372036854775808 {
                                                            case -106:
                                                                v88 = "foo bar"
                                                            case 774:
                                                                var v96 float32 = (((-2.25 + 1.5) + 0.1) / ((1e10 - (1.5 + (3.14159 / 1.5))) + (1e10 - (0.1 / 3.14159))))
                                                                _ = v96
                                                            case -821:
                                                                fmt.Println(v3, v4, v80)
                                                            }
                                                        }
                                                    } else {
                                                        var v97 int8 = (v71 - v71)
                                                        if (((-1 >= 0) || v80) || (-1 <= 42)) {
                                                            fmt.Println(v87, v4, v87)
                                                        }
                                                        _ = v97
                                                    }
                                                }
                                            }
                                        }
                                        var v98 float32 = (0.1 * 0.1)
                                        _ = v98
                                    }
                                    v4 = 2
                                    _ = fi89
                                }
                                _ = v87
                                _ = v88
                            }
                            v72 = 2
                            _ = fi81
                        }
                        v71 -= 100
                        _ = v80
                    } else {
                        var v99 uint32 = 646
                        _ = v99
                    }
                case 9223372036854775807:
                    v72 = v72
                default:
                    fmt.Println(v71)
                }
            }
            fmt.Println(v71)
            v72 = 42
        }
        v71 += v71
    } else {
        v4 += v4
    }
    v71 += (v71 * v71)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
