package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float32 = 0.0
    var v2 int64 = 1
    v3 := []string{"hello", "\n\t", ""}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 65535
    var v6 float32 = (((v1 * v1) / v1) - v1)
    var v7 float32 = v6
    var v8 uint8 = 128
    v7 -= float32(1e10)
    fmt.Println(v1, v5, v7)
    switch 128 {
    case -61:
        var v9 string = ""
        _ = v9
    case 4294967295:
        var v10 int16 = 32767
        _ = v10
    case 0:
        var v11 int64 = 0
        _ = v11
    default:
        var v12 float64 = (3.14159 - 1.5)
        _ = v12
    }
    switch "foo bar" {
    case "a":
        for fi13 := 0; fi13 < 1; fi13++ {
            v8 = 2
            switch fi13 {
            case 9223372036854775807:
                switch "hello" {
                case "hello":
                    if false {
                        var v14 bool = (0 <= uint16(v8))
                        for fi15 := 0; fi15 < 2; fi15++ {
                            switch "\n\t" {
                            case "foo bar":
                                var v16 uint32 = v5
                                _ = v16
                            }
                            v1 -= (v1 + (v7 + 0.1))
                            _ = fi15
                        }
                        _ = v14
                    }
                case "\n\t":
                    v17 := v4[(v5 | (^(v5)))]
                    _ = v17
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        switch (fi13 / (((fi13 >> ((uint(v8)) & 63))) | 1)) {
                        case 656:
                            for fi18 := 0; fi18 < 15; fi18++ {
                                if true {
                                    v1 += v6
                                }
                                v3 = append(v3, "\n\t")
                                _ = fi18
                            }
                        default:
                            switch -2 {
                            case 1:
                                fmt.Println(fi13)
                            case -811:
                                v4[(v5 | (^(v5)))] = (v5 + v5)
                            case 887:
                                v4[v5] = 42
                            default:
                                _ = len(v3)
                            }
                        }
                    default:
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v8 += (v8 / ((0) | 1))
                            _ = fi19
                        }
                    }
                default:
                    var v20 uint8 = v8
                    _ = v20
                }
            }
            _ = fi13
        }
    case "hello":
        v21 := GenS1{F0: 1}
        _ = v21
    case "\n\t":
        if (299 == (v2 + v2)) {
            v4[(v5 | v5)] = v5
        }
    default:
        if (0 != 841) {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 float32 = (((v6 - -2.25) + -2.25) + float32(0.0))
                    if (!((128 <= fi22) && (!(18446744073709551615 == 42)))) {
                        if (int32(v2) < -437) {
                            var v25 int16 = 0
                            _ = v25
                        }
                    } else {
                        if true {
                            fmt.Println(v5)
                            var v26 float32 = v6
                            _ = v26
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi22
            }
            if ((1.5 < 1.5) && (1e10 < (((float64((v1 / 1.5)) / 1.5) * (0.0 / 3.14159)) - 0.1))) {
                v6 = -2.25
                v27 := GenS2{F0: 1}
                _ = v27
            }
        }
    }
    v28 := GenS1{F0: 323}
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := GenS2{F0: (v2 + (v2 >> ((2) & 63)))}
        v4[1] = v5
        _ = fi29
        _ = v30
    }
    v5 = (v5 ^ 588)
    v31 := GenS3{F0: (v5 - (v5 >> ((0) & 63)))}
    v32 := GenS2{F0: (v2 + 9223372036854775807)}
    v33 := GenS1{F0: 579}
    v34 := v4[4294967295]
    for fi35 := 0; fi35 < 18; fi35++ {
        if false {
            _ = len(v4)
            v2 = 334
            if (!(!false)) {
                var v36 int32 = int32(v34)
                _ = v36
            }
        } else {
            v37 := GenS1{F0: 403}
            if false {
                for fi38 := 0; fi38 < 7; fi38++ {
                    if true {
                        _ = len(v3)
                        v8 = (v8 | 1)
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                if true {
                                    v2 += 42
                                }
                            case "":
                                v39 := GenS3{F0: v5}
                                _ = v39
                            case "hello":
                                var v40 int = 923
                                _ = v40
                            }
                        } else {
                            v41 := GenS2{F0: v2}
                            v5 = (v5 & (v34 << ((18446744073709551615) & 63)))
                            _ = v41
                        }
                    } else {
                        var v42 uint16 = 32767
                        v7 -= 1.5
                        _ = v42
                    }
                    if (false || true) {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    } else {
                        if (!false) {
                            v7 += v1
                            var v43 int = fi38
                            if (!(-338 > 32767)) {
                                var v44 float64 = ((((3.14159 * (3.14159 + 1e10)) - -2.25) - 1e10) / float64(1e10))
                                v5 += 42
                                v43 = fi38
                                _ = v44
                            } else {
                                fmt.Println(v1)
                                if true {
                                    var v45 int64 = 51
                                    if false {
                                        v46 := GenS1{F0: 32767}
                                        var v47 uint64 = 2
                                        fmt.Println(v1)
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v45
                                }
                                var v48 int64 = -1
                                _ = v48
                            }
                            _ = v43
                        }
                        var v49 float32 = (((v1 * (v7 * 0.0)) + (v6 * v7)) / v1)
                        v50 := GenS4{F0: 42, F1: 18446744073709551615}
                        _ = v49
                        _ = v50
                    }
                    var v51 int16 = -867
                    var v52 uint = 1
                    fmt.Println(v52)
                    _ = fi38
                    _ = v51
                    _ = v52
                }
                var v53 float32 = (1.5 * (1.5 / 3.14159))
                _ = v53
            } else {
                var v54 bool = (v8 > v8)
                if (("hello" < "\n\t") && (false && (!false))) {
                    var v55 int32 = 0
                    var v56 uint8 = v8
                    _ = v55
                    _ = v56
                }
                fmt.Println(v8)
                _ = v54
            }
            var v57 uint16 = 105
            _ = v37
            _ = v57
        }
        switch fi35 {
        case -451:
            v1 -= v7
        default:
            for fi58 := 0; fi58 < 14; fi58++ {
                var v59 uint32 = 488
                if false {
                    v7 += v1
                    v6 = (float32(3.14159) * v6)
                    if (149 <= -1) {
                        if (42 > 127) {
                            v59 += v34
                            if ((v2 % ((v2) | 1)) < (^(v2))) {
                                fmt.Println(v2, v59)
                            } else {
                                fmt.Println(v6, fi58, v1)
                                _ = len(v3)
                            }
                            if (65535 == 571) {
                                v60 := GenS5{F0: 42, F1: 42}
                                _ = v60
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v61 := GenS6{F0: 343}
                        switch "\n\t" {
                        case "a":
                            v8 = v8
                        case "hello":
                            var v62 bool = true
                            _ = v62
                        case "foo bar":
                            v63 := GenS4{F0: (v8 + (v8 | (^(v8)))), F1: 42}
                            switch "foo bar" {
                            case "foo bar":
                                v64 := v4[(v59 << ((2) & 63))]
                                _ = v64
                            case "":
                                v65 := GenS7{F0: 2147483648}
                                _ = v65
                            case "\n\t":
                                var v66 uint = 4294967295
                                _ = v66
                            }
                            var v67 uint8 = 2
                            _ = v63
                            _ = v67
                        }
                        var v68 int8 = 42
                        v69 := GenS2{F0: (^(v2))}
                        var v70 uint64 = uint64(v8)
                        _ = v61
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                }
                if true {
                    if true {
                        var v71 uint32 = (v34 | (v59 - 2147483647))
                        v7 -= (1e10 + 1.5)
                        _ = len(v4)
                        _ = v71
                    } else {
                        switch "a" {
                        case "\n\t":
                            var v72 uint32 = v59
                            _ = v72
                        case "hello":
                            if (2 == 1) {
                                var v73 uint16 = 885
                                switch "a" {
                                case "":
                                    v6 = float32(1e10)
                                default:
                                    _ = len(v4)
                                }
                                var v74 int8 = -68
                                v75 := v4[v34]
                                _ = v73
                                _ = v74
                                _ = v75
                            } else {
                                v76 := GenS2{F0: v2}
                                _ = v76
                            }
                        case "":
                            v8 = v8
                        }
                    }
                    v77 := GenS8{F0: (^(v8)), F1: 2147483647}
                    var v78 int16 = -836
                    _ = v77
                    _ = v78
                }
                _ = fi58
                _ = v59
            }
        }
        _ = fi35
    }
    v7 = 1e10
    for fi79 := 0; fi79 < 3; fi79++ {
        for fi80 := 0; fi80 < 19; fi80++ {
            v8 += 1
            if false {
                if (true && (-814 < -1)) {
                    _ = len(v4)
                    var v81 int32 = 432
                    switch (("hello" + ("hello" + ("a" + ("hello" + ("a" + "hello"))))) + ("hello" + "")) {
                    case "\n\t":
                        var v82 int16 = 1
                        _ = v82
                    }
                    _ = v81
                } else {
                    switch fi79 {
                    case 227:
                        fmt.Println(v6, v34)
                    case -198:
                        switch fi80 {
                        case 1:
                            if true {
                                v1 = v1
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 4294967295:
                        v6 = (1.5 + float32(float64(float32(float64(v7)))))
                    default:
                        fmt.Println(v6)
                    }
                    v83 := GenS9{F0: v8, F1: (v8 & (^(v8)))}
                    _ = v83
                }
            }
            var v84 uint16 = 255
            if (!((820 == 65535) && false)) {
                _ = len(v3)
                fmt.Println(fi79, v7, v34)
                v1 -= v7
            }
            var v85 int16 = 1000
            _ = fi80
            _ = v84
            _ = v85
        }
        _ = fi79
    }
    var v86 bool = true
    for fi87 := 0; fi87 < 12; fi87++ {
        v2 = -1
        _ = fi87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v86
}
