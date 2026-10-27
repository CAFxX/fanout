package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }
type GenS4 struct { F0 int64; F1 int16; F2 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ""
    var v3 uint = 0
    v4 := []int{-1}
    v5 := make(map[string]int)
    if (false && true) {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 65535, F1: 42}
            var v8 float64 = (0.1 + 1e10)
            _ = fi6
            _ = v7
            _ = v8
        }
        switch 32768 {
        case -2147483648:
            for fi9 := 0; fi9 < 9; fi9++ {
                if (413 > 2) {
                    v10 := GenS1{F0: 32767, F1: 127}
                    var v11 int = 425
                    _ = v10
                    _ = v11
                } else {
                    var v12 int16 = 315
                    _ = len(v4)
                    v13 := GenS1{F0: 65535, F1: 17}
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        default:
            v2 = v2
        }
    }
    v2 = "foo bar"
    if false {
        v2 = "a"
        if true {
            var v14 uint8 = 127
            var v15 uint = (v3 << (((v3 / ((uint(v14)) | 1))) & 63))
            switch -9223372036854775808 {
            case 970:
                for fi16 := 0; fi16 < 18; fi16++ {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v15 += (v3 & 2)
                        switch "hello" {
                        case "a":
                            var v18 uint64 = 0
                            _ = v18
                        case "hello":
                            var v19 bool = false
                            _ = v19
                        case "\n\t":
                            if (false || false) {
                                v15 = v3
                                var v20 int16 = int16(v14)
                                _ = v20
                            }
                            var v21 uint64 = 1
                            _ = v21
                        default:
                            v2 = "hello"
                        }
                        _ = fi17
                    }
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v14 -= (v14 * 63)
                        var v23 float64 = -2.25
                        switch v2 {
                        case "":
                            v15 += v3
                        case "foo bar":
                            switch v2 {
                            case "":
                                v2 = "a"
                            default:
                                switch "hello" {
                                case "hello":
                                    fmt.Println(fi22, fi22, v14)
                                case "\n\t":
                                    fmt.Println(v23, fi22, v14)
                                case "":
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(fi22, v15)
                                    }
                                default:
                                    v24 := v5[v2]
                                    _ = v24
                                }
                            }
                        default:
                            v25 := GenS1{F0: 42, F1: 42}
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = fi16
                }
            case 0:
                switch (((v2 + ("\n\t" + ((v2 + (v2 + v2)) + (("hello" + (v2 + (v2 + (((v2 + ((v2 + "\n\t") + "foo bar")) + ("\n\t" + "a")) + v2)))) + (v2 + "foo bar"))))) + ("a" + (v2 + (v2 + ((v2 + (v2 + "hello")) + v2))))) + "a") {
                case "foo bar":
                    if (829 < -32768) {
                        if false {
                            switch v2 {
                            case "hello":
                                var v26 bool = (!true)
                                _ = v26
                            case "\n\t":
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    var v28 uint16 = 315
                                    for fi29 := 0; fi29 < 2; fi29++ {
                                        var v30 uint64 = 18446744073709551615
                                        var v31 int8 = 63
                                        var v32 uint8 = uint8(v30)
                                        v15 += v3
                                        var v33 bool = (!false)
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                            case "a":
                                var v34 uint64 = 18446744073709551615
                                _ = v34
                            }
                            if true {
                                v2 = "foo bar"
                            }
                            var v35 int64 = 9223372036854775807
                            _ = v35
                        } else {
                            for fi36 := 0; fi36 < 6; fi36++ {
                                var v37 uint32 = 1
                                _ = fi36
                                _ = v37
                            }
                        }
                        if false {
                            v4 = append(v4, 32767)
                            var v38 int32 = -328
                            _ = v38
                        } else {
                            v4 = append(v4, 32768)
                            for fi39 := 0; fi39 < 1; fi39++ {
                                delete(v1, "foo bar")
                                fmt.Println(v15, v15, fi39)
                                v15 += v3
                                _ = fi39
                            }
                        }
                    } else {
                        var v40 string = v2
                        var v41 int32 = 42
                        var v42 int64 = -9223372036854775808
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                case "hello":
                    var v43 string = v2
                    _ = v43
                case "\n\t":
                    if false {
                        var v44 int32 = 2147483647
                        if (((354 == -9223372036854775808) && false) && (!false)) {
                            v45 := GenS1{F0: 2, F1: 39}
                            _ = len(v5)
                            _ = v45
                        } else {
                            for fi46 := 0; fi46 < 7; fi46++ {
                                v3 = (v15 % (((^(v15))) | 1))
                                v14 += (v14 % (((v14 * (v14 ^ uint8(v3)))) | 1))
                                _ = fi46
                            }
                            var v47 uint8 = v14
                            switch v2 {
                            case "":
                                if true {
                                    var v48 uint16 = 191
                                    for fi49 := 0; fi49 < 10; fi49++ {
                                        var v50 int8 = -128
                                        v51 := v5["foo bar"]
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v48
                                }
                            }
                            _ = v47
                        }
                        for fi52 := 0; fi52 < 4; fi52++ {
                            fmt.Println(v15)
                            var v53 float64 = -2.25
                            _ = fi52
                            _ = v53
                        }
                        _ = v44
                    }
                }
            default:
                if (2 < 857) {
                    switch 32767 {
                    case 9223372036854775807:
                        if false {
                            for fi54 := 0; fi54 < 8; fi54++ {
                                v3 -= (v3 & (v3 / ((v15) | 1)))
                                if false {
                                    switch v2 {
                                    case "foo bar":
                                        v3 += 18446744073709551615
                                    case "\n\t":
                                        if false {
                                            for fi55 := 0; fi55 < 1; fi55++ {
                                                v4 = append(v4, (fi55 | fi54))
                                                _ = fi55
                                            }
                                        } else {
                                            v56 := GenS2{F0: 255, F1: -87}
                                            _ = v56
                                        }
                                    case "a":
                                        var v57 bool = false
                                        _ = v57
                                    default:
                                        v2 = v2
                                    }
                                    fmt.Println(fi54, v15)
                                    v58 := GenS3{F0: 4294967295, F1: 127, F2: (v15 ^ (v3 * v3))}
                                    _ = v58
                                }
                                if false {
                                    for fi59 := 0; fi59 < 3; fi59++ {
                                        v60 := GenS4{F0: 42, F1: -431, F2: -743}
                                        v14 = 128
                                        _ = fi59
                                        _ = v60
                                    }
                                    v3 -= v15
                                } else {
                                    v2 = (((v2 + v2) + ("hello" + "a")) + (v2 + ""))
                                }
                                _ = fi54
                            }
                            v3 = 2
                        } else {
                            var v61 int8 = -105
                            _ = v61
                        }
                    }
                } else {
                    var v62 int8 = 1
                    var v63 bool = (true || ((!false) || false))
                    v64 := GenS5{F0: 2, F1: 660}
                    _ = v62
                    _ = v63
                    _ = v64
                }
            }
            _ = v14
            _ = v15
        }
        for fi65 := 0; fi65 < 14; fi65++ {
            var v66 bool = (!false)
            _ = fi65
            _ = v66
        }
    }
    switch -9223372036854775808 {
    case -1:
        var v67 float64 = float64(0.1)
        _ = v67
    case 32767:
        var v68 uint8 = 0
        _ = v68
    case -285:
        var v69 float32 = 0.0
        _ = v69
    default:
        v2 = (v2 + v2)
    }
    v70 := v1[(v2 + (("\n\t" + "foo bar") + v2))]
    v2 = v2
    for fi71 := 0; fi71 < 13; fi71++ {
        if (!false) {
            var v72 uint32 = (v70 | 1)
            var v73 uint32 = v70
            _ = v72
            _ = v73
        } else {
            v74 := GenS4{F0: 0, F1: -1, F2: -556}
            if false {
                v4 = append(v4, fi71)
            }
            _ = v74
        }
        if false {
            v3 += (v3 + v3)
            var v75 int64 = 9223372036854775807
            v76 := GenS1{F0: 2, F1: 100}
            _ = v75
            _ = v76
        }
        var v77 bool = false
        _ = fi71
        _ = v77
    }
    var v78 uint32 = (v70 | (v70 ^ 833))
    v2 = v2
    var v79 float32 = -2.25
    fmt.Println(v78, v3)
    v80 := GenS6{F0: 0, F1: 146}
    v81 := GenS6{F0: 1, F1: 190}
    v79 -= ((v79 * 1.5) / 1.0)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        for fi82 := 0; fi82 < 3; fi82++ {
            v83 := GenS6{F0: (-(fi82)), F1: 127}
            _ = fi82
            _ = v83
        }
    } else {
        v84 := GenS6{F0: 128, F1: 2}
        var v85 uint32 = uint32(v3)
        _ = v84
        _ = v85
    }
    switch v2 {
    case "a":
        if false {
            if true {
                v2 = v2
                for fi86 := 0; fi86 < 13; fi86++ {
                    for fi87 := 0; fi87 < 18; fi87++ {
                        if true {
                            var v88 uint64 = 42
                            switch (((v2 + v2) + v2) + (((v2 + (((v2 + "foo bar") + "foo bar") + ("\n\t" + v2))) + (v2 + v2)) + ((v2 + (v2 + (v2 + v2))) + "\n\t"))) {
                            case "foo bar":
                                v89 := GenS7{F0: v78}
                                _ = v89
                            default:
                                switch (fi87 - fi86) {
                                case 42:
                                    v90 := GenS6{F0: (fi87 | 32767), F1: 128}
                                    _ = v90
                                case 127:
                                    var v91 uint8 = 42
                                    _ = v91
                                case 4294967295:
                                    if (579 != -335) {
                                        var v92 uint32 = 2147483648
                                        _ = v92
                                    } else {
                                        var v93 int64 = -9223372036854775808
                                        if true {
                                            var v94 float32 = (1.5 * (((v79 / ((1e10 / (v79 - (float32(float64(((float32(3.14159) + (v79 + v79)) - (v79 - ((float32(1.5) * ((v79 / (v79 * v79)) - v79)) - float32((0.0 - (0.0 - 1e10)))))))) / 1.0))) / v79)) / 1e10) * (v79 + 1e10)))
                                            _ = v94
                                        } else {
                                            v79 = (-2.25 + v79)
                                        }
                                        _ = v93
                                    }
                                }
                            }
                            fmt.Println(fi86, v88)
                            _ = v88
                        }
                        if false {
                            v95 := GenS8{F0: (fi87 * (fi86 - fi86)), F1: (fi87 + fi86)}
                            _ = v95
                        }
                        fmt.Println(v79)
                        _ = fi87
                    }
                    _ = fi86
                }
            }
        } else {
            v78 = (v70 >> (((v3 ^ v3)) & 63))
        }
    case "\n\t":
        delete(v1, "\n\t")
    case "foo bar":
        for fi96 := 0; fi96 < 2; fi96++ {
            v70 -= v70
            _ = fi96
        }
    }
    v79 += ((v79 / (v79 * ((0.0 + ((float32(((-2.25 * (-2.25 - (0.0 + 0.0))) / ((1.5 * ((((1e10 / 1.5) + -2.25) + -2.25) / (1e10 - ((float64(v79) - 0.1) - -2.25)))) / -2.25))) / (1e10 * v79)) + v79)) - ((v79 * float32(3.14159)) + 1.5)))) - 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
