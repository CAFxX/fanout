package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int64; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 bool; F2 int32 }
type GenS8 struct { F0 string; F1 string; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS10 struct { F0 int; F1 string; F2 bool }
func genFn0(p0 uint64) uint {
    var v6 uint = 1
    var v7 uint = v6
    var v8 int64 = -1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v7 - (v6 >> ((uint(p0)) & 63)))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = float64((0.1 + (((float32(0.0) * float32(-2.25)) - ((0.0 + (3.14159 - 1e10)) + (1e10 - 3.14159))) + 0.1)))
    v3 := make(map[string]uint32)
    var v4 uint = 106
    v5 := []int{-1, 2147483647, 127}
    v9 := GenS1{F0: 32, F1: 32768}
    v4 = v4
    if (38 < 2) {
        var v10 uint16 = 42
        v11 := GenS1{F0: -128, F1: 598}
        _ = v10
        _ = v11
    } else {
        fmt.Println(v4, v2)
        if (uint16(v4) <= 42) {
            var v12 bool = true
            _ = v12
        } else {
            fmt.Println(v2, v2)
            switch "" {
            case "\n\t":
                v13 := GenS2{F0: "a", F1: 188, F2: (v4 <= v4)}
                _ = v13
            case "hello":
                var v14 uint = v4
                _ = v14
            case "a":
                var v15 int8 = -128
                _ = v15
            }
        }
    }
    if false {
        delete(v3, "foo bar")
    }
    v2 = 3.14159
    v16 := v3[""]
    fmt.Println(v4, v2, v4)
    _ = len(v3)
    switch "\n\t" {
    case "":
        if ((float32(v2) * (1.5 * 0.1)) != 3.14159) {
            if ((false && false) || false) {
                var v17 float32 = 0.1
                var v18 int = -981
                var v19 int8 = -1
                _ = v17
                _ = v18
                _ = v19
            }
            if true {
                if (32767 == 0) {
                    if (false || false) {
                        if ((!false) || (v16 < 2147483647)) {
                            switch "\n\t" {
                            case "":
                                v16 = v16
                            case "hello":
                                v2 = v2
                            case "\n\t":
                                v4 += 255
                            }
                        }
                        var v20 float32 = 0.0
                        _ = v20
                    } else {
                        var v21 int8 = 1
                        if ((!true) || (1.5 != 1.5)) {
                            fmt.Println(v4, v21)
                            v22 := GenS3{F0: 1, F1: v21, F2: 0}
                            _ = v22
                        }
                        _ = v21
                    }
                    for fi23 := 0; fi23 < 3; fi23++ {
                        if ((fi23 + fi23) >= 128) {
                            switch (fi23 * 32767) {
                            case -100:
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    if (false && false) {
                                        switch fi23 {
                                        case 0:
                                            v25 := GenS2{F0: "", F1: 1, F2: true}
                                            switch fi24 {
                                            case 9223372036854775807:
                                                switch (fi24 + 1) {
                                                case 188:
                                                    var v26 int32 = 0
                                                    fmt.Println(v16)
                                                    _ = v26
                                                case -1:
                                                    if (true && false) {
                                                        v4 = (v4 * (v4 | v4))
                                                        v27 := GenS4{F0: 49}
                                                        v1[0] = (^(fi24))
                                                        _ = v27
                                                    }
                                                }
                                            case -118:
                                                fmt.Println(fi24, fi23)
                                            }
                                            v4 -= genFn0(1)
                                            _ = v25
                                        default:
                                            var v28 int16 = -424
                                            _ = v28
                                        }
                                        switch "" {
                                        case "hello":
                                            fmt.Println(fi23, v4, v4)
                                        case "":
                                            v5 = append(v5, (fi24 * (fi23 % ((-645) | 1))))
                                        case "\n\t":
                                            switch fi23 {
                                            case -113:
                                                var v29 int8 = -52
                                                _ = v29
                                            case 128:
                                                fmt.Println(v4)
                                            }
                                        }
                                    } else {
                                        var v30 float32 = 0.1
                                        _ = v30
                                    }
                                    _ = fi24
                                }
                            case 1:
                                var v31 float32 = 0.1
                                _ = v31
                            case 819:
                                v16 = (v16 + v16)
                            default:
                                v16 += v16
                            }
                            v3[""] = 2
                        } else {
                            v32 := GenS5{F0: 66}
                            v33 := GenS5{F0: 0}
                            fmt.Println(v4)
                            _ = v32
                            _ = v33
                        }
                        _ = fi23
                    }
                }
                switch "\n\t" {
                case "foo bar":
                    if true {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            fmt.Println(v4)
                            v4 = v4
                            _ = fi34
                        }
                    } else {
                        v4 += genFn0(1)
                    }
                default:
                    if (-1 <= 0) {
                        v4 = 118
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            } else {
                var v35 bool = true
                _ = v35
            }
        }
    case "\n\t":
        v36 := GenS6{F0: v16}
        _ = v36
    case "hello":
        switch "a" {
        case "foo bar":
            for fi37 := 0; fi37 < 9; fi37++ {
                v2 += v2
                v16 -= v16
                if true {
                    if true {
                        v4 = v4
                        if (true || true) {
                            _ = len(v5)
                            if (0.1 <= 0.0) {
                                v38 := GenS6{F0: 606}
                                switch "foo bar" {
                                case "hello":
                                    var v39 float64 = v2
                                    _ = v39
                                case "foo bar":
                                    fmt.Println(fi37, fi37, v4)
                                }
                                _ = v38
                            } else {
                                var v40 int8 = 31
                                for fi41 := 0; fi41 < 12; fi41++ {
                                    v42 := GenS6{F0: v16}
                                    var v43 int32 = 42
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    var v45 int32 = -99
                                    _ = fi44
                                    _ = v45
                                }
                                _ = v40
                            }
                            if (!((false || false) && true)) {
                                v46 := GenS4{F0: 1}
                                var v47 float32 = 1e10
                                _ = v46
                                _ = v47
                            }
                        } else {
                            fmt.Println(v4)
                        }
                        v16 -= (v16 & v16)
                    } else {
                        for fi48 := 0; fi48 < 8; fi48++ {
                            v49 := GenS3{F0: 712, F1: -128, F2: 747}
                            v50 := GenS7{F0: v4, F1: ((-86 < 0) && false), F2: 42}
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                    }
                    fmt.Println(fi37)
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 uint8 = 7
                        v2 += float64(0.1)
                        v52 -= 67
                        v52 -= (v52 / ((v52) | 1))
                        _ = fi51
                        _ = v52
                    }
                }
                _ = fi37
            }
        case "hello":
            v2 = 0.1
        case "a":
            v53 := v3["hello"]
            _ = v53
        }
    }
    fmt.Println(v4, v2)
    for fi54 := 0; fi54 < 6; fi54++ {
        _ = len(v3)
        switch "a" {
        case "a":
            if (9223372036854775807 < 9223372036854775807) {
                for fi55 := 0; fi55 < 18; fi55++ {
                    v56 := GenS6{F0: v16}
                    if ("hello" == "foo bar") {
                        if ((-2147483648 == -619) || true) {
                            switch "\n\t" {
                            case "foo bar":
                                var v57 float64 = float64(-2.25)
                                _ = v57
                            case "hello":
                                v2 -= float64(1e10)
                            case "":
                                var v58 float64 = (3.14159 + ((v2 / (v2 - v2)) * float64(1.5)))
                                _ = v58
                            }
                            var v59 uint32 = 550
                            v2 -= v2
                            _ = v59
                        } else {
                            var v60 uint16 = 65535
                            _ = v60
                        }
                        v61 := v3["foo bar"]
                        _ = v61
                    }
                    v62 := GenS3{F0: 978, F1: -128, F2: 18446744073709551615}
                    _ = fi55
                    _ = v56
                    _ = v62
                }
                for fi63 := 0; fi63 < 11; fi63++ {
                    v64 := GenS6{F0: (v16 >> (((^(v4))) & 63))}
                    _ = fi63
                    _ = v64
                }
                if false {
                    if (2 != 127) {
                        if false {
                            fmt.Println(fi54)
                            v65 := v1[836]
                            fmt.Println(v4)
                            _ = v65
                        } else {
                            v2 = v2
                        }
                        v4 -= (v4 % ((uint(v16)) | 1))
                        switch "" {
                        case "":
                            switch 2147483647 {
                            case 23:
                                var v66 uint32 = v16
                                _ = v66
                            }
                        case "\n\t":
                            if (304 < -32768) {
                                v16 = (v16 << ((2) & 63))
                                var v67 float32 = 1.5
                                _ = v67
                            }
                        default:
                            var v68 int64 = -9223372036854775808
                            _ = v68
                        }
                    } else {
                        v69 := GenS1{F0: 42, F1: 857}
                        v16 = (^(v16))
                        if (!(false && (!(255 < 255)))) {
                            var v70 uint32 = uint32(v4)
                            for fi71 := 0; fi71 < 2; fi71++ {
                                if true {
                                    _ = len(v1)
                                }
                                _ = fi71
                            }
                            _ = v70
                        } else {
                            _ = len(v1)
                            v72 := GenS7{F0: uint(v16), F1: false, F2: -82}
                            _ = v72
                        }
                        _ = v69
                    }
                    fmt.Println(v16)
                    var v73 float64 = v2
                    _ = v73
                }
            } else {
                for fi74 := 0; fi74 < 2; fi74++ {
                    v16 = (v16 | v16)
                    if false {
                        switch fi54 {
                        case 32767:
                            var v75 int = (fi54 % (((fi74 >> ((v4) & 63))) | 1))
                            _ = v75
                        case -297:
                            v2 = 0.0
                        case 128:
                            v2 -= float64(float32(float64(0.1)))
                        }
                        fmt.Println(fi54, v2)
                    } else {
                        v16 += v16
                        v76 := GenS6{F0: (v16 >> ((genFn0(4294967295)) & 63))}
                        switch 4294967295 {
                        case -125:
                            _ = len(v5)
                        }
                        var v77 string = "foo bar"
                        switch (^(fi54)) {
                        case 126:
                            var v78 int32 = 881
                            _ = v78
                        }
                        _ = v76
                        _ = v77
                    }
                    var v79 int64 = -696
                    _ = fi74
                    _ = v79
                }
                v16 = 65535
                v16 = 4294967295
            }
        case "":
            v5 = append(v5, (fi54 % ((fi54) | 1)))
        default:
            var v80 bool = false
            _ = v80
        }
        _ = fi54
    }
    v81 := GenS3{F0: 42, F1: -128, F2: 18446744073709551615}
    var v82 int64 = 0
    for fi83 := 0; fi83 < 13; fi83++ {
        fmt.Println(v16, v82)
        var v84 float32 = (0.1 / ((((3.14159 - 1.5) / ((float32(v2) - (((((1.5 / 3.14159) / 1.5) - 1.5) - 1.5) - 1e10)) / (3.14159 * 1e10))) * 3.14159) / 1.5))
        for fi85 := 0; fi85 < 7; fi85++ {
            v86 := GenS7{F0: (v4 % ((0) | 1)), F1: (-32768 != 907), F2: -1}
            var v87 float64 = 0.1
            delete(v1, uint32(v4))
            _ = fi85
            _ = v86
            _ = v87
        }
        _ = fi83
        _ = v84
    }
    for fi88 := 0; fi88 < 14; fi88++ {
        switch "a" {
        case "":
            for fi89 := 0; fi89 < 6; fi89++ {
                var v90 float64 = v2
                if true {
                    v1[(v16 ^ (v16 / ((v16) | 1)))] = fi89
                    var v91 bool = ((v82 ^ (^(v82))) > (v82 & v82))
                    var v92 uint64 = 1
                    _ = v91
                    _ = v92
                } else {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            fmt.Println(v2, v82)
                        default:
                            if (false || (255 <= (v4 - 2))) {
                                if true {
                                    v82 += (v82 % ((v82) | 1))
                                    fmt.Println(v90, v82)
                                    v93 := GenS4{F0: 557}
                                    v4 -= (v4 % ((v4) | 1))
                                    _ = v93
                                }
                                v2 -= v2
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(fi88)
                    case -1:
                        v94 := GenS8{F0: "foo bar", F1: "a", F2: 0}
                        _ = v94
                    default:
                        if true {
                            if (!false) {
                                var v95 uint = 299
                                v95 = genFn0(42)
                                if true {
                                    if (int8(v82) > 42) {
                                        var v96 int64 = 4
                                        fmt.Println(v82, fi89)
                                        var v97 uint32 = (^(v16))
                                        v4 = v4
                                        var v98 string = "hello"
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    } else {
                                        var v99 int32 = 42
                                        var v100 float32 = (1.5 - 0.1)
                                        var v101 int64 = (v82 * (-(v82)))
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    }
                                    v2 += (v90 + float64(0.1))
                                }
                                _ = v95
                            } else {
                                var v102 uint16 = 1
                                v103 := GenS9{F0: 42, F1: 128, F2: (v102 + v102)}
                                _ = v102
                                _ = v103
                            }
                            v4 = 4294967295
                            var v104 float32 = 3.14159
                            _ = len(v5)
                            v105 := GenS1{F0: 0, F1: fi88}
                            v106 := GenS10{F0: (fi89 * (fi89 + fi89)), F1: "", F2: (((float64(v104) + ((((3.14159 + v90) - (v90 / 1.5)) * float64(0.0)) * 0.0)) - v90) > 1.5)}
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                    }
                    var v107 int8 = 52
                    v90 -= v2
                    _ = v107
                }
                var v108 uint32 = (v16 * v16)
                var v109 uint8 = 0
                if false {
                    fmt.Println(v90)
                    var v110 int32 = 42
                    v111 := GenS10{F0: -906, F1: "hello", F2: (!true)}
                    _ = v110
                    _ = v111
                }
                switch "foo bar" {
                case "":
                    v16 -= v108
                }
                _ = fi89
                _ = v108
                _ = v109
                _ = v90
            }
        case "hello":
            v2 -= (v2 - v2)
        }
        var v112 int64 = v82
        _ = fi88
        _ = v112
    }
    var v113 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v113
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v9
}
