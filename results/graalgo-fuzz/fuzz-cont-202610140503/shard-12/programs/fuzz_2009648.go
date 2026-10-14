package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 int16; F1 uint8; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 uint64; F1 uint; F2 int8 }

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 int8 = -69
    var v4 uint = 49
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int64 = 1
        _ = fi5
        _ = v6
    }
    v4 -= (v4 % ((65535) | 1))
    v7 := GenS1{F0: -942, F1: -9223372036854775808, F2: false}
    for fi8 := 0; fi8 < 13; fi8++ {
        switch (fi8 & -1) {
        case 1:
            switch fi8 {
            case 379:
                fmt.Println(fi8, v3, fi8)
            case 226:
                for fi9 := 0; fi9 < 16; fi9++ {
                    switch "a" {
                    case "foo bar":
                        v4 = v4
                    case "":
                        if (!false) {
                            v10 := v2[((("foo bar" + "hello") + ("hello" + ("foo bar" + "a"))) + "foo bar")]
                            _ = len(v2)
                            v10 = (v10 * v10)
                            _ = v10
                        } else {
                            var v11 uint32 = 42
                            v1 = append(v1, fi8)
                            v3 = v3
                            _ = v11
                        }
                    case "hello":
                        if ((fi8 & (fi8 & 128)) == -757) {
                            switch (fi8 & (fi8 >> ((255) & 63))) {
                            case -2147483648:
                                v4 = 265
                            case 42:
                                v2[""] = 2147483648
                            }
                            var v12 uint64 = 2
                            var v13 int16 = 42
                            _ = v12
                            _ = v13
                        }
                    default:
                        switch (("a" + ("foo bar" + "hello")) + "\n\t") {
                        case "":
                            fmt.Println(v3)
                        case "foo bar":
                            v3 = 0
                        default:
                            v14 := GenS1{F0: -1, F1: 786, F2: false}
                            _ = v14
                        }
                    }
                    _ = fi9
                }
            case 0:
                if (!false) {
                    if true {
                        if false {
                            for fi15 := 0; fi15 < 8; fi15++ {
                                var v16 uint8 = 39
                                _ = fi15
                                _ = v16
                            }
                            switch "\n\t" {
                            case "":
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 uint16 = 2
                                    var v19 int16 = 1
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            case "hello":
                                fmt.Println(v4)
                            default:
                                var v20 float64 = 0.0
                                _ = v20
                            }
                        } else {
                            _ = len(v2)
                        }
                        v3 -= v3
                    }
                    var v21 float64 = 3.14159
                    if (2 != 498) {
                        if (v21 > v21) {
                            if false {
                                delete(v2, ("a" + "foo bar"))
                            }
                        }
                        v21 -= v21
                    }
                    _ = v21
                } else {
                    switch "" {
                    case "\n\t":
                        v4 = (v4 >> ((65535) & 63))
                    case "":
                        var v22 int32 = -1
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        fmt.Println(fi23, v4, fi23)
                        _ = fi23
                    }
                    if ((!(42 >= -32768)) || true) {
                        for fi24 := 0; fi24 < 6; fi24++ {
                            var v25 int8 = (v3 * (v3 ^ 0))
                            _ = fi24
                            _ = v25
                        }
                        _ = len(v1)
                        v3 = v3
                    }
                }
            default:
                fmt.Println(v4, v4)
            }
        case -404:
            var v26 uint32 = 42
            _ = v26
        case 0:
            v27 := GenS1{F0: -2147483648, F1: 42, F2: (0 < 32767)}
            _ = v27
        }
        fmt.Println(v3, v3, v4)
        v28 := GenS2{F0: v3, F1: 152, F2: 253}
        _ = fi8
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        switch (fi29 % ((fi29) | 1)) {
        case -18:
            _ = len(v2)
        case -73:
            for fi30 := 0; fi30 < 6; fi30++ {
                v31 := GenS3{F0: 1, F1: 42}
                v4 = (^(v4))
                _ = fi30
                _ = v31
            }
        case 938:
            fmt.Println(fi29)
        default:
            _ = len(v2)
        }
        if len(v1) > 0 {
            v1[0] = (fi29 / ((fi29) | 1))
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint32 = 921
        if (false && false) {
            fmt.Println(v4, v3)
            var v34 string = (("foo bar" + "a") + ("\n\t" + "a"))
            if (42 == 42) {
                for fi35 := 0; fi35 < 9; fi35++ {
                    v36 := GenS2{F0: v3, F1: 921, F2: 42}
                    var v37 bool = false
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v34 = (("a" + v34) + "")
                var v38 uint = v4
                _ = v38
            } else {
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 uint16 = 32767
                    v34 = v34
                    _ = fi39
                    _ = v40
                }
                var v41 int = fi32
                _ = v41
            }
            _ = v34
        } else {
            v33 += (v33 ^ v33)
            fmt.Println(v4)
            var v42 int32 = int32(fi32)
            _ = v42
        }
        var v43 int = -384
        _ = fi32
        _ = v33
        _ = v43
    }
    if ((!((0.0 * 0.1) < float32((1.5 * -2.25)))) || (!(false && ((false && (true || (false && true))) && ((0.1 / 1.0) < 0.0))))) {
        v4 -= v4
    } else {
        v4 = v4
    }
    fmt.Println(v3)
    switch -408 {
    case -904:
        v4 = 4294967295
    }
    v44 := GenS2{F0: -1, F1: 346, F2: 255}
    var v45 uint16 = 2
    for fi46 := 0; fi46 < 16; fi46++ {
        if (false || (fi46 >= fi46)) {
            switch (("hello" + "foo bar") + ("" + "\n\t")) {
            case "foo bar":
                if (42 > 2) {
                    switch ("a" + (("foo bar" + "hello") + "a")) {
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            var v47 int64 = -826
                            _ = v47
                        case "\n\t":
                            for fi48 := 0; fi48 < 12; fi48++ {
                                fmt.Println(v45, v45)
                                _ = fi48
                            }
                        }
                    case "\n\t":
                        if true {
                            if true {
                                var v49 uint16 = (v45 - 1)
                                v3 -= 107
                                v50 := GenS3{F0: 127, F1: 769}
                                _ = v49
                                _ = v50
                            } else {
                                if false {
                                    switch "" {
                                    case "foo bar":
                                        if (!((1e10 - 0.0) >= 0.0)) {
                                            v51 := v2[(("foo bar" + ("a" + ("\n\t" + "a"))) + "foo bar")]
                                            for fi52 := 0; fi52 < 11; fi52++ {
                                                v45 -= (v45 & v45)
                                                var v53 uint8 = 0
                                                _ = fi52
                                                _ = v53
                                            }
                                            _ = v51
                                        } else {
                                            v1 = append(v1, -2147483648)
                                        }
                                    }
                                    v54 := GenS4{F0: 2, F1: (v45 * 620), F2: ("a" + ("a" + "a"))}
                                    var v55 uint64 = 4294967295
                                    _ = v54
                                    _ = v55
                                } else {
                                    v4 = (v4 + v4)
                                    for fi56 := 0; fi56 < 12; fi56++ {
                                        fmt.Println(fi46, v3)
                                        switch (fi46 >> ((21) & 63)) {
                                        case -482:
                                            _ = len(v2)
                                        case -1:
                                            if (1 >= 0) {
                                                v57 := GenS1{F0: 863, F1: -9223372036854775808, F2: (-2147483648 == int32(fi56))}
                                                var v58 int32 = -1
                                                _ = v57
                                                _ = v58
                                            }
                                        default:
                                            var v59 int16 = 42
                                            _ = v59
                                        }
                                        _ = fi56
                                    }
                                    var v60 int32 = 1
                                    _ = v60
                                }
                            }
                            v61 := GenS5{F0: ""}
                            fmt.Println(v4)
                            _ = v61
                        }
                    }
                    if false {
                        switch (fi46 + fi46) {
                        case 42:
                            v62 := GenS2{F0: 0, F1: 0, F2: (v45 | 499)}
                            _ = v62
                        }
                        v1 = append(v1, (fi46 + (fi46 ^ (fi46 ^ (fi46 >> (((v4 % ((v4) | 1))) & 63))))))
                        switch "\n\t" {
                        case "foo bar":
                            switch (-(fi46)) {
                            case 2147483647:
                                fmt.Println(v4, v4, fi46)
                            default:
                                var v63 uint32 = 0
                                _ = v63
                            }
                        default:
                            var v64 int = 127
                            _ = v64
                        }
                    } else {
                        var v65 bool = (!(!false))
                        fmt.Println(v3, v4)
                        _ = len(v1)
                        _ = v65
                    }
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v2)
                    default:
                        v4 += (v4 + 242)
                    }
                } else {
                    var v66 int64 = 42
                    _ = v66
                }
            }
            fmt.Println(fi46, v3)
        }
        v4 = v4
        v67 := GenS6{F0: 2147483648, F1: 1, F2: v3}
        _ = fi46
        _ = v67
    }
    v3 = -66
    v68 := GenS4{F0: 1, F1: v45, F2: (("hello" + "foo bar") + "")}
    if false {
        v69 := GenS7{F0: -57, F1: 164, F2: 2147483647}
        switch 128 {
        case -2147483648:
            v70 := GenS8{F0: 0}
            _ = v70
        default:
            fmt.Println(v3, v3, v3)
        }
        switch -761 {
        case 32768:
            v71 := GenS9{F0: -9223372036854775808, F1: 2}
            _ = v71
        case 194:
            for fi72 := 0; fi72 < 2; fi72++ {
                for fi73 := 0; fi73 < 18; fi73++ {
                    switch fi72 {
                    case 1:
                        if true {
                            if (-566 >= 582) {
                                v2[("foo bar" + (("foo bar" + "a") + "foo bar"))] = 2147483648
                                v45 = v45
                                v45 = 510
                            }
                        }
                    case 32768:
                        var v74 bool = (!true)
                        _ = v74
                    case 128:
                        if false {
                            fmt.Println(v4, v3)
                            v75 := GenS5{F0: "foo bar"}
                            _ = v75
                        } else {
                            v3 -= (-(v3))
                            if (true || true) {
                                var v76 uint32 = 75
                                fmt.Println(v76, v4)
                                _ = v76
                            } else {
                                var v77 int8 = v3
                                var v78 uint64 = 691
                                _ = v77
                                _ = v78
                            }
                            v79 := GenS4{F0: 4294967295, F1: (v45 + (v45 >> ((v4) & 63))), F2: "foo bar"}
                            _ = v79
                        }
                    default:
                        v45 = 463
                    }
                    fmt.Println(fi72, v45)
                    _ = fi73
                }
                var v80 string = ("" + "foo bar")
                if (false && (29 >= 0)) {
                    delete(v2, (v80 + "a"))
                }
                _ = fi72
                _ = v80
            }
        default:
            for fi81 := 0; fi81 < 13; fi81++ {
                v4 -= v4
                v82 := GenS5{F0: "foo bar"}
                var v83 float64 = (((0.0 + ((3.14159 - 1.5) / 1.0)) + -2.25) + 0.1)
                _ = fi81
                _ = v82
                _ = v83
            }
        }
        _ = v69
    } else {
        if false {
            var v84 uint64 = 1
            if false {
                v84 = (v84 - 1)
                var v85 int32 = 701
                var v86 int = 32768
                _ = v85
                _ = v86
            } else {
                v3 = (-(v3))
                switch "" {
                case "a":
                    v87 := GenS10{F0: 42, F1: (v4 * uint(v45)), F2: (v3 >> ((v4) & 63))}
                    _ = v87
                case "hello":
                    v45 += v45
                case "\n\t":
                    fmt.Println(v84, v3)
                default:
                    _ = len(v1)
                }
                v3 += -76
            }
            _ = v84
        }
    }
    v4 = (v4 & (^(v4)))
    v88 := GenS3{F0: 127, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v68
    _ = v7
    _ = v88
}
