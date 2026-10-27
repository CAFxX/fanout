package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int32; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int64; F2 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int16 }

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{1, 2147483648, 357}
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 4; fi4++ {
        v2 = append(v2, 0)
        _ = fi4
    }
    fmt.Println(v1)
    for fi5 := 0; fi5 < 3; fi5++ {
        delete(v3, 953)
        v1 += -1
        v6 := GenS1{F0: true}
        _ = fi5
        _ = v6
    }
    v1 += 0
    delete(v3, 1)
    _ = len(v2)
    _ = len(v3)
    var v7 int8 = 127
    v7 += (-(v7))
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int8 = -1
        if false {
            fmt.Println(fi8, v9, v7)
            if (2 > 255) {
                fmt.Println(v9, v1, fi8)
            } else {
                v3[916] = "a"
                if ((42 != 0) && true) {
                    switch "foo bar" {
                    case "":
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v10 := GenS2{F0: -1, F1: 4294967295}
                            _ = v10
                        }
                    case "a":
                        switch (fi8 >> ((2) & 63)) {
                        case 9223372036854775807:
                            v11 := v3[2147483648]
                            _ = v11
                        case -302:
                            fmt.Println(v9, v9, fi8)
                        case 4294967295:
                            switch "hello" {
                            case "":
                                var v12 uint16 = 0
                                _ = v12
                            case "foo bar":
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    if false {
                                        _ = len(v3)
                                        switch (((("" + ("" + "hello")) + "foo bar") + ("a" + "hello")) + "a") {
                                        case "":
                                            var v14 uint8 = 4
                                            _ = v14
                                        default:
                                            fmt.Println(fi8, v7)
                                        }
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        v7 = (^(v9))
                                    case "foo bar":
                                        v15 := GenS2{F0: -111, F1: 18446744073709551615}
                                        _ = v15
                                    case "\n\t":
                                        v7 += (-(v7))
                                        var v16 uint8 = 42
                                        fmt.Println(fi8)
                                        _ = v16
                                    default:
                                        v17 := GenS3{F0: 127}
                                        _ = v17
                                    }
                                    switch 9223372036854775807 {
                                    case 128:
                                        v1 -= v1
                                    case 1:
                                        v18 := GenS2{F0: -1, F1: 2}
                                        v19 := GenS2{F0: 615, F1: 0}
                                        v7 += v7
                                        var v20 uint64 = 1
                                        switch "" {
                                        case "":
                                            _ = len(v3)
                                        case "\n\t":
                                            fmt.Println(v7, fi13, v20)
                                        default:
                                            v20 = (v20 * v20)
                                        }
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    case 42:
                                        v2 = append(v2, 65535)
                                        switch ("hello" + ("hello" + "\n\t")) {
                                        case "foo bar":
                                            var v21 bool = (18446744073709551615 >= 0)
                                            _ = v21
                                        case "hello":
                                            var v22 int16 = 42
                                            _ = v22
                                        case "a":
                                            var v23 uint64 = 2
                                            _ = v23
                                        default:
                                            var v24 uint32 = 4294967295
                                            _ = v24
                                        }
                                    default:
                                        v9 = v7
                                    }
                                    v25 := GenS4{F0: 0}
                                    fmt.Println(v1, v1)
                                    _ = fi13
                                    _ = v25
                                }
                            default:
                                switch "a" {
                                case "":
                                    _ = len(v2)
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                }
                            }
                        default:
                            var v26 int64 = 356
                            _ = v26
                        }
                    case "hello":
                        v1 -= (v1 & 1)
                    default:
                        v9 = (v9 + v9)
                    }
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v28 := GenS5{F0: (-(v1)), F1: -2147483648, F2: -1}
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(fi8, fi8)
                } else {
                    _ = len(v2)
                }
                if true {
                    var v29 float32 = 3.14159
                    v9 -= (v7 | (v7 & v7))
                    _ = v29
                }
            }
            fmt.Println(v7, fi8, v1)
        } else {
            v9 = v9
            switch fi8 {
            case 225:
                fmt.Println(fi8, v1)
            case -2147483648:
                for fi30 := 0; fi30 < 7; fi30++ {
                    var v31 int = (fi30 * (fi30 + int(v1)))
                    _ = fi30
                    _ = v31
                }
            }
            switch (-(fi8)) {
            case 32768:
                switch (^(fi8)) {
                case 0:
                    if false {
                        var v32 bool = (((((1e10 + -2.25) + 3.14159) / 1.5) - (1e10 + (0.0 + (0.1 * 1e10)))) > 1e10)
                        switch "foo bar" {
                        case "foo bar":
                            if v32 {
                                v32 = (565 != 506)
                            } else {
                                var v33 float64 = (3.14159 * 0.0)
                                var v34 int16 = 42
                                _ = v33
                                _ = v34
                            }
                        case "a":
                            _ = len(v2)
                        default:
                            v7 = (v7 * v9)
                        }
                        _ = v32
                    }
                case 32767:
                    if ((-(v1)) != v1) {
                        v2 = append(v2, 65535)
                    }
                case -1:
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v36 := GenS6{F0: 1, F1: 255}
                        _ = fi35
                        _ = v36
                    }
                }
            case 0:
                v37 := GenS7{F0: true, F1: 617}
                _ = v37
            default:
                v7 -= v9
            }
        }
        for fi38 := 0; fi38 < 13; fi38++ {
            switch "hello" {
            case "foo bar":
                var v39 int32 = v1
                _ = v39
            case "":
                if ((!(fi38 == (fi8 * (fi8 | (^(fi38)))))) && (1e10 <= (0.1 / 1.5))) {
                    var v40 uint = 18446744073709551615
                    switch (fi8 ^ fi38) {
                    case 128:
                        var v41 uint16 = 42
                        _ = v41
                    case 32768:
                        var v42 uint16 = 255
                        _ = v42
                    case 1:
                        var v43 int32 = (v1 - 425)
                        _ = v43
                    default:
                        if (!(true && (("\n\t" + "hello") <= ""))) {
                            if true {
                                switch (fi38 - int(v7)) {
                                case 2147483647:
                                    var v44 uint64 = 200
                                    _ = v44
                                case 0:
                                    var v45 uint8 = 127
                                    _ = v45
                                default:
                                    switch "hello" {
                                    case "":
                                        if (0 != 429) {
                                            if (float32(0.0) <= (0.1 - 0.0)) {
                                                var v46 uint8 = 102
                                                _ = v46
                                            } else {
                                                v3[4294967295] = ""
                                                _ = len(v2)
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                var v47 int64 = 78
                                                _ = v47
                                            default:
                                                var v48 uint16 = 390
                                                _ = v48
                                            }
                                        }
                                    case "foo bar":
                                        v49 := GenS3{F0: -756}
                                        _ = v49
                                    default:
                                        var v50 bool = (100 <= v9)
                                        v9 -= (v9 & (v9 & -89))
                                        _ = v50
                                    }
                                }
                                _ = len(v3)
                                v51 := v3[2147483648]
                                _ = v51
                            }
                            if (3.14159 == -2.25) {
                                v1 += 479
                                v52 := GenS7{F0: (!false), F1: 0}
                                var v53 uint8 = 9
                                _ = v52
                                _ = v53
                            }
                            v7 -= (v9 | -128)
                        } else {
                            var v54 int16 = 909
                            v55 := GenS8{F0: (fi8 << ((42) & 63))}
                            var v56 uint8 = 1
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            switch 32768 {
                            case -21:
                                if true {
                                    if (v1 != (v1 >> (((v40 | (v40 ^ v40))) & 63))) {
                                        v54 = v54
                                    } else {
                                        if false {
                                            v54 = 32767
                                            switch (fi38 + fi8) {
                                            case 4294967295:
                                                var v57 float64 = 0.1
                                                _ = v57
                                            case 602:
                                                v58 := GenS9{F0: ((("a" + "a") + ("foo bar" + "foo bar")) + ("hello" + ("hello" + (((("" + "") + ("" + "hello")) + ("a" + ("a" + ""))) + ("a" + "")))))}
                                                _ = v58
                                            }
                                            if ((!(!false)) && (!(!(v1 > 42)))) {
                                                switch (fi8 % ((9223372036854775807) | 1)) {
                                                case 9223372036854775807:
                                                    v56 += (v56 ^ (v56 - 42))
                                                default:
                                                    v2 = append(v2, 1)
                                                }
                                                if ((fi38 ^ int(v1)) > -2147483648) {
                                                    if ((^(v54)) > (v54 % ((840) | 1))) {
                                                        var v59 uint8 = (^(v56))
                                                        var v60 int8 = (v9 ^ v7)
                                                        v61 := GenS4{F0: 2}
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        var v62 int64 = 42
                                                        v1 = v1
                                                        v62 = v62
                                                        var v63 int16 = (v54 & (v54 >> ((v40) & 63)))
                                                        v64 := GenS10{F0: true, F1: (-(v62)), F2: -119}
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        var v65 int = fi8
                                                        _ = v65
                                                    default:
                                                        v66 := GenS11{F0: -50}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v66
                                                    }
                                                    var v67 float32 = (3.14159 + (float32(1e10) - 0.1))
                                                    _ = v67
                                                }
                                            }
                                        }
                                        v56 = v56
                                    }
                                } else {
                                    v1 = 2147483647
                                    v68 := GenS12{F0: true}
                                    v69 := GenS13{F0: v56, F1: 4294967295, F2: 948}
                                    _ = v68
                                    _ = v69
                                }
                            case 962:
                                switch (^(fi38)) {
                                case -118:
                                    v70 := GenS14{F0: v1, F1: int64(v1)}
                                    _ = v70
                                case -961:
                                    v1 += int32(v7)
                                default:
                                    var v71 uint = (v40 % ((255) | 1))
                                    _ = v71
                                }
                            case 2147483647:
                                _ = len(v3)
                            default:
                                v72 := GenS8{F0: 32767}
                                _ = v72
                            }
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                    }
                    var v73 uint64 = 447
                    _ = v40
                    _ = v73
                }
            case "hello":
                var v74 int8 = v7
                _ = v74
            }
            v7 += v7
            _ = fi38
        }
        _ = fi8
        _ = v9
    }
    if (!true) {
        v75 := GenS4{F0: 899}
        if (651 != 4294967295) {
            for fi76 := 0; fi76 < 18; fi76++ {
                if false {
                    var v77 int8 = v7
                    _ = v77
                }
                var v78 float64 = float64(0.1)
                fmt.Println(v7)
                _ = fi76
                _ = v78
            }
            v79 := GenS5{F0: 2147483647, F1: 1, F2: 159}
            _ = v79
        } else {
            for fi80 := 0; fi80 < 12; fi80++ {
                var v81 int32 = (v1 << ((2) & 63))
                v1 += v81
                _ = len(v2)
                _ = fi80
                _ = v81
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v75
    }
    var v82 uint8 = 42
    v1 = 2147483647
    v83 := GenS12{F0: true}
    v84 := GenS15{F0: 226, F1: "a"}
    if (-2147483648 > v1) {
        for fi85 := 0; fi85 < 10; fi85++ {
            if (-594 >= (-(v1))) {
                v86 := GenS1{F0: false}
                v87 := GenS16{F0: 762}
                _ = v86
                _ = v87
            } else {
                v82 -= (v82 + (v82 * (v82 * v82)))
                var v88 int16 = 0
                for fi89 := 0; fi89 < 17; fi89++ {
                    v82 = v82
                    v7 = v7
                    v90 := GenS7{F0: (65535 == 0), F1: 2}
                    _ = fi89
                    _ = v90
                }
                _ = v88
            }
            _ = fi85
        }
        v82 -= (^(v82))
        v7 += v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v82
    _ = v83
    _ = v84
}
