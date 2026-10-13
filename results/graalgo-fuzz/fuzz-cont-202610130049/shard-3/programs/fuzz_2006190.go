package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 string }

func main() {
    v1 := []uint32{2147483647, 1, 2147483648}
    var v2 float64 = 0.0
    var v3 uint64 = 1
    var v4 uint16 = 32767
    v5 := make(map[int]int)
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    v3 += 4294967295
    var v6 uint64 = (v3 << ((65535) & 63))
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint64 = v6
        v8 = (v6 & (^(v8)))
        if false {
            v4 = (v4 << ((65535) & 63))
        } else {
            v8 = 42
            if false {
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 uint = 255
                    _ = fi9
                    _ = v10
                }
            }
            switch ("" + ("foo bar" + "hello")) {
            case "\n\t":
                v11 := GenS1{F0: 2, F1: 0, F2: -18}
                _ = v11
            case "foo bar":
                var v12 float32 = (-2.25 * (1.5 / 0.1))
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint32 = 4294967295
        var v15 float64 = (-2.25 - (v2 / ((v2 + (v2 * v2)) - v2)))
        v4 -= 788
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = len(v1)
    if true {
        v6 -= v3
        if ((!(((("foo bar" != ("" + "a")) && (((525 >= -2147483648) && (4294967295 < 4294967295)) || false)) && (false && true)) || true)) || false) {
            v6 = (v3 % (((v6 - v3)) | 1))
            switch "hello" {
            case "":
                fmt.Println(v4)
            case "a":
                fmt.Println(v3)
            case "foo bar":
                var v16 uint16 = 894
                _ = v16
            }
            if (true && true) {
                switch 127 {
                case -9223372036854775808:
                    v4 -= v4
                case 0:
                    switch 74 {
                    case 32768:
                        v17 := GenS1{F0: 1, F1: 18446744073709551615, F2: 1}
                        _ = v17
                    }
                case -2147483648:
                    if false {
                        v2 = v2
                        switch -9223372036854775808 {
                        case -1:
                            switch "" {
                            case "a":
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    switch (-(fi18)) {
                                    case 32767:
                                        fmt.Println(v2, v2)
                                    case 0:
                                        v19 := GenS2{F0: "foo bar", F1: (v4 << ((2) & 63)), F2: (fi18 & fi18)}
                                        _ = v19
                                    case 9223372036854775807:
                                        v3 = (v6 + v3)
                                    }
                                    var v20 uint = 0
                                    if true {
                                        v3 = v6
                                        _ = len(v5)
                                    } else {
                                        for fi21 := 0; fi21 < 17; fi21++ {
                                            if (1 != 1) {
                                                var v22 uint8 = 128
                                                v5[fi18] = 32767
                                                _ = v22
                                            } else {
                                                var v23 int64 = -1
                                                v24 := GenS3{F0: -575}
                                                var v25 uint64 = (v3 >> (((v20 | v20)) & 63))
                                                if (((((-2.25 / (float32(v2) + (0.0 / 3.14159))) + 0.0) + 0.1) * 1.5) > (0.0 * 3.14159)) {
                                                    if true {
                                                        v20 = v20
                                                        fmt.Println(v3)
                                                        if len(v1) > 0 {
                                                            v1[0] = 2
                                                        }
                                                    } else {
                                                        var v26 uint8 = 30
                                                        fmt.Println(v26)
                                                        fmt.Println(fi21, fi21)
                                                        _ = v26
                                                    }
                                                }
                                                v2 = ((((1.5 - v2) + (v2 - (v2 * ((v2 + (v2 * 0.0)) * 3.14159)))) / 1.0) / v2)
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = fi21
                                        }
                                        if true {
                                            var v27 uint32 = 0
                                            _ = v27
                                        }
                                        delete(v5, (fi18 * (fi18 - (fi18 - (fi18 | fi18)))))
                                    }
                                    _ = fi18
                                    _ = v20
                                }
                            case "foo bar":
                                v3 = v3
                            }
                        case 0:
                            var v28 float32 = 0.1
                            _ = v28
                        }
                        var v29 int32 = 2147483647
                        _ = v29
                    }
                }
                v30 := GenS4{F0: (-694 <= 0)}
                var v31 string = "hello"
                _ = v30
                _ = v31
            } else {
                fmt.Println(v4)
                fmt.Println(v6, v3)
                for fi32 := 0; fi32 < 10; fi32++ {
                    v33 := GenS2{F0: (("\n\t" + "\n\t") + "a"), F1: 1, F2: (fi32 ^ fi32)}
                    _ = fi32
                    _ = v33
                }
            }
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            if false {
                if (0 == 1) {
                    v3 += 42
                }
                if true {
                    switch "hello" {
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if false {
                            var v35 uint16 = (^(v4))
                            v2 += v2
                            v4 += v4
                            _ = v35
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if ((-2147483648 <= 0) && true) {
                                v3 = 4294967295
                            } else {
                                v36 := GenS4{F0: false}
                                _ = v36
                            }
                            switch ("\n\t" + ("\n\t" + ("\n\t" + (((("foo bar" + (((((("foo bar" + ("a" + "foo bar")) + "a") + "\n\t") + ("" + "a")) + "foo bar") + ("hello" + "a"))) + ("hello" + ("" + (("" + (("\n\t" + "hello") + "hello")) + ("hello" + "foo bar"))))) + "\n\t") + "")))) {
                            case "a":
                                if true {
                                    switch 834 {
                                    case -547:
                                        fmt.Println(fi34)
                                    default:
                                        if true {
                                            for fi37 := 0; fi37 < 8; fi37++ {
                                                v38 := GenS3{F0: 785}
                                                v5[32767] = (^(fi34))
                                                v39 := v5[310]
                                                _ = fi37
                                                _ = v38
                                                _ = v39
                                            }
                                        } else {
                                            var v40 float32 = -2.25
                                            _ = v40
                                        }
                                    }
                                    v2 -= ((float64(1e10) - v2) * 0.0)
                                } else {
                                    if false {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v2 += v2
                                        case "a":
                                            v4 = (v4 - (^(v4)))
                                        case "hello":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            for fi41 := 0; fi41 < 6; fi41++ {
                                                var v42 int8 = 6
                                                v43 := GenS5{F0: 65535, F1: 0}
                                                _ = fi41
                                                _ = v42
                                                _ = v43
                                            }
                                        }
                                        if true {
                                            v6 = (v3 * 0)
                                            v3 = (^(v6))
                                            fmt.Println(v3)
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v3, fi34)
                                        }
                                    }
                                    v3 += v3
                                    switch fi34 {
                                    case 0:
                                        if false {
                                            v44 := GenS6{F0: 2147483647, F1: -128}
                                            _ = v44
                                        }
                                    case -9223372036854775808:
                                        var v45 int8 = int8(v3)
                                        _ = v45
                                    case 474:
                                        for fi46 := 0; fi46 < 11; fi46++ {
                                            v5[(fi34 - 2147483647)] = 32768
                                            v5[fi46] = -1
                                            v47 := GenS6{F0: -506, F1: 1}
                                            _ = fi46
                                            _ = v47
                                        }
                                    default:
                                        v4 -= v4
                                    }
                                }
                            default:
                                v48 := GenS2{F0: "", F1: (v4 + (v4 | uint16(v3))), F2: 9223372036854775807}
                                _ = v48
                            }
                        }
                    }
                    if false {
                        switch (("" + ("foo bar" + "hello")) + ((("hello" + ((("" + ("" + ("\n\t" + ("hello" + "hello")))) + "") + "hello")) + "foo bar") + "hello")) {
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v3 = v6
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v2, v6)
                        case "":
                            fmt.Println(v3)
                        case "a":
                            _ = len(v1)
                        }
                        v49 := GenS7{F0: (v6 <= v3), F1: fi34}
                        if (true && (fi34 >= (fi34 / (((fi34 - (fi34 * fi34))) | 1)))) {
                            fmt.Println(v3, v4)
                            if true {
                                var v50 bool = (2 <= 0)
                                fmt.Println(fi34, v2)
                                if (!(!(4294967295 <= 79))) {
                                    v50 = (799 == 765)
                                    switch "" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = 2
                                        }
                                    case "foo bar":
                                        if (v50 || (1 == -128)) {
                                            v51 := v5[(fi34 ^ 461)]
                                            var v52 uint = 18446744073709551615
                                            var v53 uint = (v52 % (((^(v52))) | 1))
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                    default:
                                        fmt.Println(v6)
                                    }
                                    v2 = v2
                                }
                                _ = v50
                            } else {
                                v54 := GenS5{F0: v4, F1: 4294967295}
                                _ = v54
                            }
                            switch "" {
                            case "foo bar":
                                var v55 int64 = 9223372036854775807
                                _ = v55
                            case "\n\t":
                                v2 -= 0.1
                            case "hello":
                                v56 := GenS2{F0: "\n\t", F1: (v4 * 255), F2: fi34}
                                _ = v56
                            }
                        } else {
                            if (!false) {
                                var v57 float64 = 1e10
                                v2 = 0.1
                                v58 := GenS6{F0: 0, F1: 0}
                                _ = v57
                                _ = v58
                            }
                            v4 = 1
                            v3 += (v3 % ((v3) | 1))
                        }
                        _ = v49
                    }
                    var v59 uint32 = uint32(v6)
                    _ = v59
                } else {
                    if false {
                        v60 := GenS4{F0: (-2147483648 > (fi34 - fi34))}
                        var v61 uint64 = (v6 | v3)
                        if len(v1) > 0 {
                            v1[0] = 450
                        }
                        _ = v60
                        _ = v61
                    }
                }
            }
            _ = fi34
        }
    }
    v62 := GenS8{F0: 0}
    fmt.Println(v2)
    switch 1 {
    case 42:
        if (-9223372036854775808 >= 42) {
            for fi63 := 0; fi63 < 2; fi63++ {
                var v64 uint64 = 640
                if ((v2 < 1.5) || ((true && true) && (!false))) {
                    var v65 int8 = 0
                    _ = v65
                }
                _ = fi63
                _ = v64
            }
        }
    case 128:
        v6 = v6
    case 576:
        if false {
            v6 += (v6 + (v3 - 18446744073709551615))
            var v66 uint16 = v4
            _ = v66
        }
    default:
        if false {
            v67 := GenS9{F0: 1, F1: ("\n\t" + "\n\t")}
            _ = v67
        } else {
            var v68 float64 = float64(-2.25)
            _ = v68
        }
    }
    if ((42 >= -9223372036854775808) || true) {
        var v69 bool = true
        var v70 int32 = 2147483647
        if v69 {
            v71 := GenS2{F0: ("" + "a"), F1: v4, F2: 1}
            _ = v71
        } else {
            var v72 uint16 = v4
            fmt.Println(v69, v2, v6)
            _ = v72
        }
        _ = v69
        _ = v70
    } else {
        var v73 bool = true
        _ = v73
    }
    var v74 bool = false
    if v74 {
        for fi75 := 0; fi75 < 2; fi75++ {
            if v74 {
                var v76 uint32 = 1
                _ = v76
            }
            if true {
                fmt.Println(v74, v4, v3)
                fmt.Println(v3, v3, v2)
                if len(v1) > 0 {
                    v1[0] = 2
                }
            } else {
                fmt.Println(v3)
            }
            _ = fi75
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 701
        }
        for fi77 := 0; fi77 < 12; fi77++ {
            fmt.Println(v3)
            _ = fi77
        }
        switch -2147483648 {
        case -223:
            v78 := GenS4{F0: false}
            _ = v78
        }
    }
    var v79 int8 = -8
    var v80 int64 = 9223372036854775807
    switch 4294967295 {
    case 4294967295:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v81 := GenS5{F0: (v4 >> ((0) & 63)), F1: 4294967295}
        _ = v81
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v74
    _ = v79
    _ = v80
}
