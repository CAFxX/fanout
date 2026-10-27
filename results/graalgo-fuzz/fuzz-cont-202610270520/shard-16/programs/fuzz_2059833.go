package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 string }
type GenS5 struct { F0 uint64; F1 string; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 bool; F2 int16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 int64; F2 uint8 }
type GenS13 struct { F0 uint64; F1 int8 }

func main() {
    v1 := []uint32{65535, 285}
    var v2 string = (("" + "hello") + "foo bar")
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 13; fi5++ {
        v2 = (v2 + v2)
        v2 = v2
        v6 := GenS1{F0: 1}
        _ = fi5
        _ = v6
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = v2
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v2)
            v9 := GenS2{F0: 65535}
            _ = fi8
            _ = v9
        }
        if (!(-386 <= -1)) {
            v2 = ("a" + "\n\t")
            v2 = v2
            _ = len(v1)
        }
        _ = fi7
    }
    if (18446744073709551615 == 42) {
        fmt.Println(v2)
    } else {
        var v10 int64 = 565
        if (!false) {
            if true {
                var v11 float32 = 0.0
                _ = v11
            }
            fmt.Println(v10, v10)
            if false {
                if len(v1) > 0 {
                    v1[0] = 843
                }
                v10 -= v10
                var v12 float64 = 1e10
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int32 = -1
                v16 := GenS3{F0: ((3.14159 > 3.14159) || true), F1: 42}
                _ = fi14
                _ = v15
                _ = v16
            }
            if false {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int16 = 1
                    v10 = v10
                    var v19 int8 = 127
                    v19 -= v19
                    var v20 float64 = 0.0
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                if true {
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (1.5 != float32(0.1)) {
                        fmt.Println(v10, v2, fi13)
                        fmt.Println(fi13, fi13)
                    } else {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(v10, fi21)
                            if (!false) {
                                var v22 uint = 116
                                if true {
                                    v23 := GenS4{F0: 1, F1: fi21, F2: ("\n\t" + "hello")}
                                    _ = v23
                                }
                                v24 := GenS4{F0: -1, F1: fi13, F2: ("hello" + "")}
                                _ = v22
                                _ = v24
                            }
                            v25 := GenS5{F0: 18446744073709551615, F1: v2, F2: v2}
                            var v26 string = v2
                            _ = len(v1)
                            _ = fi21
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v2, fi13, fi13)
                        if false {
                            v27 := GenS4{F0: 1, F1: 4294967295, F2: "\n\t"}
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 12; fi28++ {
                                if ((0 == -119) && false) {
                                    var v29 float32 = 0.1
                                    _ = v29
                                }
                                if (684 >= 11) {
                                    if true {
                                        if true {
                                            v10 -= v10
                                        } else {
                                            var v30 uint16 = 652
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            fmt.Println(v2, fi13)
                                            _ = v30
                                        }
                                    }
                                    if (true || ((-116 == 1) && false)) {
                                        if len(v1) > 0 {
                                            v1[0] = 977
                                        }
                                        v31 := GenS3{F0: (-272 != -2147483648), F1: 4294967295}
                                        _ = v31
                                    } else {
                                        fmt.Println(v2, fi28, fi13)
                                        switch fi28 {
                                        case 127:
                                            _ = len(v1)
                                        case 2147483647:
                                            fmt.Println(fi13, fi13)
                                            if false {
                                                switch v2 {
                                                case "\n\t":
                                                    fmt.Println(fi13, v10)
                                                default:
                                                    delete(v3, ((("" + (((("a" + "\n\t") + v2) + v2) + ("foo bar" + ("foo bar" + v2)))) + "a") + v2))
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 817
                                                }
                                            } else {
                                                v32 := GenS6{F0: 42}
                                                if len(v1) > 0 {
                                                    v1[0] = 296
                                                }
                                                _ = v32
                                            }
                                        }
                                    }
                                    v2 = (("" + v2) + "")
                                }
                                v33 := GenS4{F0: 0, F1: -1, F2: v2}
                                _ = fi28
                                _ = v33
                            }
                        }
                    }
                } else {
                    if false {
                        var v34 int64 = (v10 >> ((877) & 63))
                        v35 := GenS7{F0: 255, F1: v2}
                        _ = v34
                        _ = v35
                    }
                    v10 = -9223372036854775808
                }
                v2 = "hello"
            } else {
                var v36 bool = (0 > 2147483647)
                var v37 bool = false
                v38 := GenS7{F0: 18446744073709551615, F1: "foo bar"}
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi13
        }
        _ = v10
    }
    var v39 float64 = (0.1 + (((3.14159 - 3.14159) - -2.25) - -2.25))
    fmt.Println(v39)
    v2 = "\n\t"
    switch "\n\t" {
    case "\n\t":
        switch v2 {
        case "a":
            v39 = 0.1
        case "\n\t":
            v40 := GenS2{F0: 1}
            _ = v40
        case "hello":
            var v41 float64 = (1.5 - ((((1e10 + ((v39 - ((v39 * (v39 / (v39 * (v39 + v39)))) - (1e10 / ((((float64(0.1) / ((v39 * v39) / 1.5)) / (3.14159 * (v39 + 1.5))) / v39) + 1e10)))) + (1.5 + v39))) * (1.5 - float64((1e10 * 0.0)))) * 1e10) * v39))
            _ = v41
        }
    case "foo bar":
        var v42 bool = true
        _ = v42
    default:
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    if true {
        v43 := GenS5{F0: 4294967295, F1: v2, F2: v2}
        v44 := GenS4{F0: -2147483648, F1: 1, F2: v2}
        v39 = 0.0
        _ = v43
        _ = v44
    } else {
        switch 128 {
        case -904:
            if (!true) {
                v3[(v2 + v2)] = 608
                fmt.Println(v2, v2)
            }
        case -350:
            v45 := GenS3{F0: true, F1: 2}
            _ = v45
        default:
            v2 = ((v2 + v2) + "")
        }
    }
    var v46 int32 = 0
    switch 744 {
    case 32768:
        v39 -= v39
    case -9223372036854775808:
        var v47 int64 = 581
        _ = v47
    case 413:
        var v48 uint64 = 88
        _ = v48
    }
    switch 2147483647 {
    case -1:
        v49 := GenS4{F0: (v46 | v46), F1: 792, F2: (v2 + "a")}
        _ = v49
    case 999:
        if true {
            _ = len(v1)
        } else {
            v50 := v3[v2]
            _ = v50
        }
    case 620:
        fmt.Println(v46)
    }
    fmt.Println(v46, v46, v2)
    fmt.Println(v39)
    for fi51 := 0; fi51 < 4; fi51++ {
        v3[(v2 + (v2 + v2))] = 1
        for fi52 := 0; fi52 < 10; fi52++ {
            if (!(477 == 438)) {
                v46 -= (v46 << ((42) & 63))
                v2 = v2
                var v53 uint64 = 1
                _ = v53
            } else {
                if (!true) {
                    var v54 string = v2
                    switch fi51 {
                    case -9223372036854775808:
                        if (false && (639 <= 42)) {
                            v55 := GenS8{F0: (fi51 * fi52)}
                            _ = v55
                        }
                    default:
                        fmt.Println(fi51)
                        var v56 float64 = -2.25
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v56
                    }
                    v46 = (v46 * (v46 >> ((4294967295) & 63)))
                    _ = v54
                } else {
                    var v57 float64 = v39
                    if false {
                        var v58 uint64 = 0
                        _ = v58
                    }
                    v2 = v2
                    v1 = append(v1, 42)
                    v59 := GenS2{F0: 99}
                    _ = v57
                    _ = v59
                }
                if ((v2 + ((v2 + v2) + "foo bar")) == v2) {
                    var v60 int64 = -9223372036854775808
                    _ = v60
                }
            }
            switch -746 {
            case 4294967295:
                switch ("a" + (v2 + v2)) {
                case "foo bar":
                    v39 = (v39 * v39)
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 39
                    }
                case "hello":
                    v2 = ""
                default:
                    v61 := GenS8{F0: fi51}
                    _ = v61
                }
            }
            if true {
                v62 := GenS3{F0: (((!true) && ((v2 + ("foo bar" + v2)) >= ("hello" + "hello"))) && true), F1: 2}
                if (268 < -1) {
                    switch -289 {
                    case 1:
                        if (742 <= 801) {
                            var v63 int64 = 9223372036854775807
                            v3[""] = 4294967295
                            _ = v63
                        }
                    case 42:
                        switch v2 {
                        case "foo bar":
                            if (true || ((v46 << ((255) & 63)) <= (v46 - -2147483648))) {
                                var v64 float64 = v39
                                _ = len(v1)
                                _ = v64
                            } else {
                                fmt.Println(fi52, v2)
                                fmt.Println(v46)
                                v65 := GenS2{F0: 790}
                                _ = v65
                            }
                        case "\n\t":
                            var v66 uint64 = 1
                            _ = v66
                        case "hello":
                            fmt.Println(fi52, fi51, fi52)
                        default:
                            var v67 int = (fi52 ^ fi52)
                            _ = v67
                        }
                    case 4294967295:
                        v68 := GenS9{F0: v46}
                        _ = v68
                    default:
                        fmt.Println(v46, fi51, fi52)
                    }
                    var v69 string = "\n\t"
                    switch -2147483648 {
                    case 128:
                        v2 = ""
                    }
                    _ = v69
                } else {
                    var v70 uint16 = 409
                    fmt.Println(fi52, v70)
                    if true {
                        if (((float32(v39) == (0.0 + 1.5)) && true) || true) {
                            var v71 uint8 = 2
                            var v72 int64 = 734
                            v73 := GenS10{F0: 356, F1: (true || true), F2: -837}
                            _ = v71
                            _ = v72
                            _ = v73
                        } else {
                            var v74 uint = 255
                            _ = v74
                        }
                        _ = len(v1)
                    } else {
                        var v75 uint = 42
                        if true {
                            v46 -= v46
                            if len(v1) > 0 {
                                v1[0] = 305
                            }
                            if true {
                                v70 -= uint16(v75)
                            }
                        } else {
                            if (1 <= 42) {
                                if (((float32(v39) + 0.1) == 1.5) && true) {
                                    v75 = (v75 ^ (v75 | 1))
                                    v76 := GenS11{F0: 1}
                                    _ = v76
                                }
                            }
                        }
                        v77 := GenS12{F0: -444, F1: 1, F2: 255}
                        _ = v75
                        _ = v77
                    }
                    _ = v70
                }
                switch (fi51 * 9223372036854775807) {
                case 9223372036854775807:
                    v2 = v2
                case 127:
                    v39 -= 3.14159
                default:
                    v78 := GenS13{F0: 1, F1: 100}
                    _ = v78
                }
                if false {
                    switch (fi52 * (fi51 >> ((988) & 63))) {
                    case -527:
                        v39 += float64((1.5 - 1e10))
                    default:
                        v79 := GenS11{F0: 9223372036854775807}
                        fmt.Println(v39)
                        _ = v79
                    }
                } else {
                    _ = len(v4)
                    switch 521 {
                    case 42:
                        var v80 uint = 2
                        _ = v80
                    case -2147483648:
                        v2 = ((v2 + v2) + v2)
                    }
                }
                if false {
                    v39 += (v39 - v39)
                    fmt.Println(v2)
                    v39 = v39
                } else {
                    var v81 uint16 = 42
                    var v82 int = (-(fi51))
                    _ = v81
                    _ = v82
                }
                v46 = -2147483648
                _ = v62
            } else {
                v39 -= (0.1 / v39)
            }
            _ = fi52
        }
        for fi83 := 0; fi83 < 1; fi83++ {
            v2 = "foo bar"
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi83
        }
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
}
