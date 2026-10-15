package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 string }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 string; F2 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 uint8; F1 uint32; F2 uint8 }

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := []string{""}
    var v5 float64 = (0.1 + 3.14159)
    for fi6 := 0; fi6 < 3; fi6++ {
        v5 = v5
        _ = fi6
    }
    if (1000 < 32767) {
        v1 = v1
        var v7 int8 = -1
        v1 = "a"
        _ = v7
    }
    v5 = float64((0.1 / -2.25))
    _ = len(v4)
    v8 := GenS1{F0: 573, F1: 32767, F2: ((v1 + v1) + "a")}
    v9 := v3[42]
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 240, F1: v9, F2: "foo bar"}
        v5 = -2.25
        if true {
            switch (v1 + ("foo bar" + "foo bar")) {
            case "":
                if (((42 <= 873) && true) && (100 > -7)) {
                    v9 -= (v9 ^ 32768)
                    switch "hello" {
                    case "\n\t":
                        switch ((v1 + ("" + "a")) + (v1 + ("" + v1))) {
                        case "\n\t":
                            for fi12 := 0; fi12 < 16; fi12++ {
                                if true {
                                    v1 = v1
                                    v9 = fi10
                                    v13 := GenS1{F0: 1000, F1: 32767, F2: v1}
                                    _ = v13
                                } else {
                                    v1 = v1
                                    switch "a" {
                                    case "hello":
                                        v1 = v1
                                    case "":
                                        fmt.Println(v5)
                                    default:
                                        fmt.Println(v1)
                                    }
                                    fmt.Println(v5)
                                }
                                var v14 int8 = 1
                                v9 = fi12
                                fmt.Println(v1, v1, fi12)
                                _ = fi12
                                _ = v14
                            }
                        case "a":
                            for fi15 := 0; fi15 < 2; fi15++ {
                                _ = len(v3)
                                _ = fi15
                            }
                        case "hello":
                            for fi16 := 0; fi16 < 15; fi16++ {
                                var v17 int32 = 42
                                var v18 uint32 = 2
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        default:
                            var v19 uint16 = 0
                            _ = v19
                        }
                    case "":
                        var v20 int64 = -804
                        _ = v20
                    case "a":
                        fmt.Println(fi10)
                    default:
                        v1 = v1
                    }
                    v21 := GenS2{F0: 32767, F1: 23, F2: 121}
                    _ = v21
                } else {
                    v5 = v5
                    _ = len(v4)
                }
            case "foo bar":
                var v22 uint64 = 360
                _ = v22
            case "a":
                var v23 uint = 0
                _ = v23
            default:
                switch "a" {
                case "hello":
                    if false {
                        v24 := GenS2{F0: 255, F1: 255, F2: -128}
                        var v25 uint64 = 335
                        _ = v24
                        _ = v25
                    } else {
                        if false {
                            var v26 uint16 = 511
                            v27 := GenS1{F0: -402, F1: fi10, F2: v1}
                            switch (fi10 / ((v9) | 1)) {
                            case 4294967295:
                                if true {
                                    switch 9223372036854775807 {
                                    case 32767:
                                        if true {
                                            var v28 int16 = 0
                                            v29 := GenS1{F0: v28, F1: (v9 / ((v9) | 1)), F2: v1}
                                            v30 := GenS3{F0: 0, F1: (v28 * -954)}
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                    case 32768:
                                        if (true && true) {
                                            var v31 bool = true
                                            if v31 {
                                                var v32 int8 = 127
                                                for fi33 := 0; fi33 < 16; fi33++ {
                                                    var v34 uint16 = (v26 % (((v26 + (v26 ^ v26))) | 1))
                                                    _ = fi33
                                                    _ = v34
                                                }
                                                _ = v32
                                            }
                                            var v35 string = "a"
                                            _ = v31
                                            _ = v35
                                        } else {
                                            v9 -= (fi10 % ((v9) | 1))
                                            v9 = -2147483648
                                        }
                                    }
                                    switch "" {
                                    case "\n\t":
                                        if (568 == 2) {
                                            var v36 int64 = 392
                                            v37 := GenS4{F0: (v36 & (v36 ^ v36)), F1: true, F2: (v1 + "\n\t")}
                                            _ = v36
                                            _ = v37
                                        }
                                    case "foo bar":
                                        fmt.Println(v26, v9, v1)
                                    case "hello":
                                        v1 = v1
                                    }
                                    var v38 uint32 = 4294967295
                                    _ = v38
                                }
                            case -2147483648:
                                v39 := GenS1{F0: 42, F1: v9, F2: "a"}
                                _ = v39
                            default:
                                v5 = (1.5 / ((0.0 + -2.25) * ((v5 - 0.0) - v5)))
                            }
                            _ = v26
                            _ = v27
                        } else {
                            v40 := GenS5{F0: 255, F1: -128}
                            v41 := GenS6{F0: 4294967295}
                            _ = v40
                            _ = v41
                        }
                        var v42 uint64 = 1
                        _ = v42
                    }
                case "":
                    fmt.Println(v9)
                case "foo bar":
                    var v43 uint8 = 151
                    _ = v43
                default:
                    v9 += 9223372036854775807
                }
            }
            delete(v3, 229)
            for fi44 := 0; fi44 < 2; fi44++ {
                v4 = append(v4, "hello")
                if len(v4) > 0 {
                    v4[0] = v1
                }
                _ = fi44
            }
        }
        _ = fi10
        _ = v11
    }
    if (v5 >= -2.25) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch 1 {
        case 243:
            switch v9 {
            case 9223372036854775807:
                switch v9 {
                case 42:
                    var v45 uint = 2
                    _ = v45
                case -166:
                    var v46 bool = (!(!false))
                    _ = v46
                case 283:
                    v47 := GenS7{F0: 0}
                    _ = v47
                }
            case 1:
                for fi48 := 0; fi48 < 18; fi48++ {
                    var v49 int = v9
                    var v50 uint16 = 65535
                    var v51 uint8 = 158
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
            default:
                v52 := GenS8{F0: -1, F1: v1, F2: 2147483647}
                _ = v52
            }
        case -215:
            switch "" {
            case "a":
                for fi53 := 0; fi53 < 7; fi53++ {
                    var v54 int = (fi53 >> ((2) & 63))
                    switch ("\n\t" + v1) {
                    case "a":
                        v1 = v1
                    case "hello":
                        if false {
                            var v55 uint32 = 2147483647
                            _ = v55
                        } else {
                            for fi56 := 0; fi56 < 20; fi56++ {
                                v5 -= ((v5 + (v5 / (v5 / 0.1))) * 1.5)
                                _ = fi56
                            }
                        }
                    case "foo bar":
                        var v57 bool = (128 != 127)
                        _ = v57
                    default:
                        var v58 int64 = 0
                        _ = v58
                    }
                    _ = fi53
                    _ = v54
                }
            case "\n\t":
                v59 := v3[0]
                _ = v59
            case "":
                if true {
                    fmt.Println(v1, v9)
                    switch (("" + (v1 + v1)) + "hello") {
                    case "":
                        v60 := GenS4{F0: -9223372036854775808, F1: true, F2: (v1 + (v1 + (v1 + "hello")))}
                        _ = v60
                    default:
                        var v61 uint32 = 252
                        _ = v61
                    }
                }
            default:
                v62 := GenS1{F0: -1, F1: (v9 - v9), F2: v1}
                _ = v62
            }
        default:
            if true {
                fmt.Println(v5)
            } else {
                v9 += (v9 | 514)
                v3[165] = v9
                v3[67] = v9
            }
        }
        for fi63 := 0; fi63 < 1; fi63++ {
            v64 := GenS3{F0: 1, F1: 0}
            _ = fi63
            _ = v64
        }
    }
    v5 -= float64(1e10)
    var v65 uint64 = 4294967295
    for fi66 := 0; fi66 < 11; fi66++ {
        v1 = (v1 + (v1 + "\n\t"))
        for fi67 := 0; fi67 < 11; fi67++ {
            v65 += 42
            _ = fi67
        }
        delete(v2, (fi66 << ((0) & 63)))
        _ = fi66
    }
    for fi68 := 0; fi68 < 7; fi68++ {
        v65 = (^(v65))
        for fi69 := 0; fi69 < 13; fi69++ {
            if true {
                v9 += 200
            } else {
                if false {
                    if (542 >= 794) {
                        fmt.Println(v5, v9, v9)
                    } else {
                        if (42 <= 935) {
                            var v70 int16 = 1000
                            var v71 string = v1
                            switch (((v1 + v1) + (v71 + v71)) + v1) {
                            case "\n\t":
                                _ = len(v3)
                            case "":
                                switch (fi69 % ((v9) | 1)) {
                                case 1:
                                    v65 = v65
                                case 127:
                                    var v72 int8 = 127
                                    _ = v72
                                default:
                                    var v73 float32 = float32((v5 / 1e10))
                                    switch ((v1 + v1) + v71) {
                                    case "foo bar":
                                        if ((v73 * (float32(3.14159) * v73)) < (((v73 - 0.1) + 0.1) / v73)) {
                                            switch "" {
                                            case "":
                                                switch v1 {
                                                case "a":
                                                    v65 += v65
                                                }
                                            case "\n\t":
                                                if len(v4) > 0 {
                                                    v4[0] = ((("a" + v71) + (v71 + v71)) + (v71 + v71))
                                                }
                                            case "hello":
                                                _ = len(v3)
                                            }
                                        }
                                    }
                                    _ = v73
                                }
                            }
                            _ = v70
                            _ = v71
                        } else {
                            var v74 int64 = 537
                            v5 = float64((3.14159 - float32(v5)))
                            v4 = append(v4, v1)
                            _ = v74
                        }
                        _ = len(v4)
                    }
                    v75 := GenS9{F0: v65}
                    v65 = (v65 >> ((225) & 63))
                    _ = v75
                } else {
                    if (!false) {
                        v9 -= fi69
                        var v76 uint16 = 211
                        v76 = (v76 * 709)
                        v65 = (v65 | (^(v65)))
                        _ = v76
                    }
                    if false {
                        var v77 string = (v1 + (v1 + "\n\t"))
                        switch -232 {
                        case 669:
                            switch fi68 {
                            case -2147483648:
                                var v78 string = ("" + v77)
                                _ = v78
                            case 128:
                                var v79 uint16 = 2
                                _ = v79
                            }
                        default:
                            var v80 uint = 18446744073709551615
                            _ = v80
                        }
                        _ = v77
                    }
                    v1 = v1
                }
                v1 = v1
            }
            v81 := GenS10{F0: 603, F1: (^(v65))}
            _ = len(v4)
            var v82 uint16 = 42
            _ = fi69
            _ = v81
            _ = v82
        }
        _ = len(v3)
        _ = fi68
    }
    v65 = 18446744073709551615
    if (-245 >= -1) {
        for fi83 := 0; fi83 < 19; fi83++ {
            switch v1 {
            case "\n\t":
                var v84 uint16 = 869
                _ = v84
            case "a":
                v85 := GenS4{F0: -201, F1: true, F2: "foo bar"}
                _ = v85
            case "":
                var v86 uint64 = (v65 % ((v65) | 1))
                _ = v86
            }
            v1 = v1
            _ = fi83
        }
        v87 := v3[65535]
        _ = v87
    } else {
        v9 = v9
        v65 = 1
    }
    v88 := GenS11{F0: 255, F1: 65535, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v8
    _ = v88
    _ = v9
}
