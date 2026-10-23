package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 uint8; F1 int; F2 string }
type GenS8 struct { F0 string; F1 int64; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 int; F2 uint16 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 string; F1 int }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int; F1 uint16 }
type GenS15 struct { F0 string; F1 int64; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = ((0.1 - (0.1 - 0.1)) / 1e10)
    v3 := []int{32768, 32767}
    v4 := make(map[uint32]bool)
    var v5 float64 = ((v2 / v2) * 0.1)
    switch "" {
    case "":
        v6 := GenS1{F0: 1, F1: 32767}
        _ = v6
    case "foo bar":
        if (!((0 > -1) && (42 > 4294967295))) {
            v1[0] = false
            var v7 uint8 = 152
            _ = v7
        }
    case "hello":
        for fi8 := 0; fi8 < 18; fi8++ {
            v5 -= v2
            _ = fi8
        }
    default:
        v9 := GenS2{F0: 807, F1: 32767, F2: 737}
        _ = v9
    }
    if ("hello" < "hello") {
        if len(v3) > 0 {
            v3[0] = 32768
        }
        if (0 >= 255) {
            var v10 string = ""
            _ = len(v4)
            _ = v10
        } else {
            v11 := GenS1{F0: 100, F1: 0}
            _ = v11
        }
        v2 -= v5
    }
    v4[2147483648] = (1 > 0)
    v12 := v1[32767]
    var v13 int64 = -1
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int16 = 42
            switch "foo bar" {
            case "foo bar":
                switch (("a" + ("\n\t" + ("foo bar" + "foo bar"))) + "foo bar") {
                case "hello":
                    var v17 uint8 = 0
                    _ = v17
                }
            default:
                if (!(v12 || (!false))) {
                    v16 += v16
                    switch "\n\t" {
                    case "\n\t":
                        var v18 string = "\n\t"
                        _ = v18
                    case "foo bar":
                        fmt.Println(v2)
                    case "a":
                        if (false && v12) {
                            v19 := v4[372]
                            v5 += float64(3.14159)
                            if (1 >= 127) {
                                var v20 float64 = (0.0 - v2)
                                _ = v20
                            }
                            switch ("hello" + ("a" + "a")) {
                            case "a":
                                if (((v13 ^ (v13 % ((v13) | 1))) != 1) || (!v12)) {
                                    v3 = append(v3, 32767)
                                    v21 := GenS2{F0: 24, F1: 1, F2: (fi15 >> ((42) & 63))}
                                    v12 = true
                                    _ = v21
                                }
                            case "\n\t":
                                v5 -= v2
                            case "hello":
                                if v12 {
                                    var v22 uint = 4294967295
                                    if (!false) {
                                        fmt.Println(v19)
                                    } else {
                                        fmt.Println(v2)
                                        var v23 float32 = 1e10
                                        if true {
                                            v24 := GenS3{F0: 32767}
                                            fmt.Println(v23, v2, v16)
                                            fmt.Println(v12, fi15, fi15)
                                            fmt.Println(fi15)
                                            delete(v4, 0)
                                            _ = v24
                                        } else {
                                            v4[1] = (34 >= 0)
                                            v22 = (v22 << ((2) & 63))
                                            if v12 {
                                                switch (fi14 | (fi14 << ((v22) & 63))) {
                                                case 9223372036854775807:
                                                    v25 := v4[0]
                                                    _ = v25
                                                case -9223372036854775808:
                                                    if v19 {
                                                        if (65535 != 65535) {
                                                            var v26 float32 = float32(v2)
                                                            v22 += 932
                                                            var v27 int32 = 0
                                                            _ = v26
                                                            _ = v27
                                                        } else {
                                                            if v19 {
                                                                v28 := GenS2{F0: 2, F1: 32767, F2: fi15}
                                                                if v12 {
                                                                    var v29 uint = v22
                                                                    _ = v29
                                                                }
                                                                _ = v28
                                                            }
                                                        }
                                                        v5 = (1.5 - v5)
                                                    }
                                                case 289:
                                                    v30 := GenS4{F0: (v22 + v22)}
                                                    _ = v30
                                                }
                                                var v31 int8 = 81
                                                _ = v31
                                            } else {
                                                v19 = true
                                            }
                                        }
                                        _ = v23
                                    }
                                    var v32 float32 = 1e10
                                    _ = v22
                                    _ = v32
                                }
                                var v33 int64 = v13
                                _ = v33
                            default:
                                delete(v4, 350)
                            }
                            if v12 {
                                v34 := GenS2{F0: 0, F1: 1, F2: fi14}
                                v35 := GenS5{F0: 2, F1: 2}
                                _ = v34
                                _ = v35
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v19
                        } else {
                            v36 := v4[2147483647]
                            _ = v36
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if v12 {
                            var v37 uint32 = 2
                            v12 = v12
                            if false {
                                var v38 int = fi14
                                v39 := GenS3{F0: 2}
                                if true {
                                    v1[(fi14 + (-(v38)))] = (0 <= 551)
                                }
                                _ = v38
                                _ = v39
                            }
                            if false {
                                if v12 {
                                    v40 := GenS6{F0: 1, F1: (1.5 >= float32(0.0))}
                                    _ = v40
                                } else {
                                    v5 = 0.1
                                    v41 := v4[(v37 & v37)]
                                    _ = v41
                                }
                                var v42 int32 = 466
                                var v43 int = (fi14 >> ((255) & 63))
                                _ = v42
                                _ = v43
                            } else {
                                switch 1 {
                                case -9223372036854775808:
                                    if v12 {
                                        var v44 uint8 = 115
                                        v2 -= v5
                                        var v45 int64 = -291
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v46 := v1[(^(fi15))]
                                        _ = v46
                                    }
                                case 74:
                                    v16 -= v16
                                case 4294967295:
                                    v3 = append(v3, (fi15 | fi15))
                                }
                                if false {
                                    v13 = (v13 << ((65535) & 63))
                                    if v12 {
                                        switch fi14 {
                                        case 474:
                                            v5 = (v2 * v2)
                                        case 213:
                                            v47 := GenS7{F0: 1, F1: (fi15 << ((795) & 63)), F2: "hello"}
                                            _ = v47
                                        }
                                    }
                                    var v48 uint32 = (v37 << ((0) & 63))
                                    _ = v48
                                }
                                v3 = append(v3, int(v16))
                                var v49 float64 = -2.25
                                var v50 uint16 = 42
                                v51 := GenS8{F0: "\n\t", F1: 1, F2: fi14}
                                if (v37 <= v37) {
                                    v52 := GenS5{F0: v50, F1: 0}
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        if true {
                                            switch "foo bar" {
                                            case "hello":
                                                var v53 uint32 = 2147483647
                                                _ = v53
                                            default:
                                                var v54 int8 = -1
                                                _ = v54
                                            }
                                            if false {
                                                v4[v37] = v12
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(fi15, v16, fi15)
                                                case "foo bar":
                                                    v37 = v37
                                                }
                                                v49 = float64(1e10)
                                            }
                                        } else {
                                            v55 := GenS2{F0: 4294967295, F1: 963, F2: (-(fi14))}
                                            var v56 int8 = -37
                                            _ = len(v4)
                                            _ = v55
                                            _ = v56
                                        }
                                        switch "a" {
                                        case "a":
                                            if len(v3) > 0 {
                                                v3[0] = (fi15 % (((fi14 >> ((42) & 63))) | 1))
                                            }
                                        case "foo bar":
                                            var v57 bool = (768 > 42)
                                            _ = v57
                                        case "\n\t":
                                            fmt.Println(v13, v12, v13)
                                        default:
                                            var v58 uint8 = 199
                                            _ = v58
                                        }
                                    case "a":
                                        v5 = (v2 / (v5 - v49))
                                    case "foo bar":
                                        var v59 uint = 42
                                        _ = v59
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v60 string = ("" + "")
                                        _ = v60
                                    case "":
                                        fmt.Println(fi14)
                                    }
                                    _ = v52
                                }
                                v12 = ((v37 / ((292) | 1)) <= v37)
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            v61 := GenS9{F0: "\n\t"}
                            _ = v37
                            _ = v61
                        } else {
                            v12 = true
                            var v62 uint = 241
                            var v63 uint16 = 2
                            _ = v62
                            _ = v63
                        }
                    default:
                        v64 := GenS2{F0: 2147483648, F1: 2, F2: fi14}
                        _ = v64
                    }
                    var v65 string = "foo bar"
                    v3 = append(v3, (fi14 / ((-95) | 1)))
                    if v12 {
                        switch v65 {
                        case "hello":
                            v16 += v16
                        default:
                            v2 = ((((v5 * v2) / (v2 * ((0.1 + float64(1.5)) - (1e10 + (((float64(((1e10 + 1.5) / (1e10 * float32(((v2 / v5) - float64((-2.25 * float32(((float64(0.1) + v2) + v2))))))))) + (1.5 / v2)) + v2) / (v2 + 0.0)))))) - ((v2 - v2) * v5)) - (v5 * float64((float32(v2) * (3.14159 + ((3.14159 + ((3.14159 - (((float32(v2) * 1.5) / 3.14159) / 0.1)) - (3.14159 / 1.0))) + ((0.1 * float32(v5)) - float32((v5 - 3.14159)))))))))
                        }
                        var v66 int16 = (v16 & -32768)
                        v67 := GenS10{F0: 0, F1: (fi14 - fi14), F2: 42}
                        _ = v66
                        _ = v67
                    } else {
                        if (343 != 681) {
                            if v12 {
                                if (!v12) {
                                    fmt.Println(v5, v13, v12)
                                    fmt.Println(v5)
                                    delete(v4, 4294967295)
                                    v13 = -9223372036854775808
                                } else {
                                    var v68 float64 = 1e10
                                    v5 -= v5
                                    _ = v68
                                }
                            } else {
                                if (2 <= 2) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if v12 {
                                        var v69 uint32 = 999
                                        fmt.Println(v69)
                                        _ = v69
                                    } else {
                                        v70 := v1[128]
                                        _ = v70
                                    }
                                    fmt.Println(v2)
                                } else {
                                    if false {
                                        if (false || (275 == 42)) {
                                            v2 = (0.0 + v2)
                                            v5 = 3.14159
                                        }
                                        v71 := GenS11{F0: (false || ((!v12) && v12)), F1: 0}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v71
                                    } else {
                                        v72 := GenS12{F0: "\n\t", F1: fi15}
                                        _ = v72
                                    }
                                    delete(v1, (fi14 + fi14))
                                }
                            }
                            v73 := GenS3{F0: 0}
                            if (("hello" + v65) < ((v65 + v65) + "hello")) {
                                var v74 uint8 = 174
                                _ = v74
                            } else {
                                switch (fi15 & 632) {
                                case 13:
                                    v75 := GenS13{F0: v13}
                                    _ = v75
                                case 32768:
                                    switch 477 {
                                    case 487:
                                        v13 = v13
                                        if false {
                                            var v76 float64 = ((v5 - v5) - (v5 - float64((0.1 * 0.1))))
                                            _ = v76
                                        } else {
                                            v77 := GenS12{F0: "hello", F1: -2147483648}
                                            _ = v77
                                        }
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v12, v2, v16)
                                        case "":
                                            _ = len(v1)
                                        default:
                                            v13 = v13
                                        }
                                    case -358:
                                        switch fi15 {
                                        case -9223372036854775808:
                                            v2 = v2
                                        case 2147483647:
                                            switch fi15 {
                                            case 1:
                                                v13 -= 0
                                            }
                                        default:
                                            v78 := GenS14{F0: -2147483648, F1: 32767}
                                            _ = v78
                                        }
                                    }
                                }
                                var v79 int8 = 42
                                v13 = (v13 >> ((565) & 63))
                                v80 := GenS15{F0: "foo bar", F1: (v13 << ((42) & 63)), F2: int(v13)}
                                v12 = v12
                                switch v65 {
                                case "hello":
                                    var v81 uint32 = 2147483647
                                    _ = v81
                                case "a":
                                    v13 = v13
                                default:
                                    v82 := GenS12{F0: v65, F1: fi15}
                                    _ = v82
                                }
                                _ = v79
                                _ = v80
                            }
                            _ = v73
                        } else {
                            v3 = append(v3, (fi15 * 128))
                        }
                        switch 32767 {
                        case 128:
                            var v83 int64 = (v13 >> ((65535) & 63))
                            _ = v83
                        case 1:
                            fmt.Println(fi15, v16)
                        }
                        var v84 uint32 = 2147483647
                        _ = v84
                    }
                    _ = v65
                }
            }
            _ = fi15
            _ = v16
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi14
    }
    var v85 uint64 = 1
    fmt.Println(v2, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
