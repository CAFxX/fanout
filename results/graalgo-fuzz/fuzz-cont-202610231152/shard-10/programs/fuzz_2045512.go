package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int; F2 int32 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]bool)
    var v3 float64 = ((1.5 + 0.0) - (1e10 - 0.1))
    v4 := []int{32768, 4294967295}
    v5 := GenS1{F0: true, F1: true}
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(fi6, v1)
        var v7 int16 = 921
        v7 = 1000
        _ = fi6
        _ = v7
    }
    var v8 int16 = 1000
    v1 -= 2
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            if (-723 > 2147483647) {
                if true {
                    _ = len(v4)
                    v4 = append(v4, fi10)
                    switch fi9 {
                    case 163:
                        var v11 uint32 = v1
                        _ = v11
                    case -40:
                        fmt.Println(v8)
                    }
                    var v12 string = ("\n\t" + ("" + ("a" + "\n\t")))
                    _ = v12
                } else {
                    var v13 uint16 = 1
                    v14 := GenS1{F0: (-539 <= 87), F1: (1 < 42)}
                    switch (("\n\t" + "hello") + "foo bar") {
                    case "hello":
                        fmt.Println(v3, v8)
                        v8 += 570
                        if false {
                            var v15 uint = 514
                            if (!(482 <= fi10)) {
                                v1 += 624
                                var v16 uint = v15
                                if (0 < -674) {
                                    switch "a" {
                                    case "\n\t":
                                        var v17 uint16 = (v13 & 2)
                                        _ = v17
                                    case "":
                                        v15 = (v15 - 936)
                                    }
                                    v16 = v15
                                } else {
                                    var v18 int = fi10
                                    v19 := GenS1{F0: (("" + "hello") == ("" + "foo bar")), F1: (((0.1 / (0.1 / 1.0)) - 0.0) != 3.14159)}
                                    _ = v18
                                    _ = v19
                                }
                                _ = v16
                            } else {
                                switch fi9 {
                                case 0:
                                    if ((771 <= (^(v15))) && true) {
                                        var v20 uint64 = 42
                                        _ = v20
                                    } else {
                                        if (true && true) {
                                            v21 := GenS2{F0: 235, F1: -1, F2: true}
                                            _ = v21
                                        } else {
                                            var v22 float32 = 0.0
                                            v8 += 0
                                            _ = v22
                                        }
                                        v4 = append(v4, (fi9 + fi10))
                                        if (float32(((0.1 - v3) + v3)) != float32(v3)) {
                                            fmt.Println(fi10)
                                            v15 = 696
                                            var v23 uint8 = 128
                                            _ = v23
                                        } else {
                                            v2[fi9] = (!true)
                                            if (!(!true)) {
                                                v3 += (v3 + 3.14159)
                                                v24 := GenS3{F0: 672, F1: (fi9 >= (^(fi9))), F2: 128}
                                                if false {
                                                    var v25 float32 = float32(((1e10 / ((float64(1e10) - 0.0) - 0.0)) * 0.0))
                                                    var v26 int64 = 0
                                                    if true {
                                                        if false {
                                                            v26 = -9223372036854775808
                                                            if false {
                                                                switch fi10 {
                                                                case 127:
                                                                    var v27 int32 = -80
                                                                    _ = v27
                                                                default:
                                                                    if (((!false) || (!true)) || (51 < 5)) {
                                                                        if ((^(v1)) != 807) {
                                                                            v26 += v26
                                                                            if false {
                                                                                v15 = 42
                                                                                switch ("a" + "") {
                                                                                case "hello":
                                                                                    if false {
                                                                                        var v28 int16 = v8
                                                                                        _ = v28
                                                                                    } else {
                                                                                        var v29 uint16 = (v13 >> ((v15) & 63))
                                                                                        v15 -= v15
                                                                                        v29 -= v29
                                                                                        _ = v29
                                                                                    }
                                                                                    v26 = (v26 ^ 1)
                                                                                    fmt.Println(v26)
                                                                                    var v30 int64 = v26
                                                                                    v13 = (v13 / ((0) | 1))
                                                                                    _ = v30
                                                                                case "foo bar":
                                                                                    var v31 int64 = v26
                                                                                    _ = v31
                                                                                }
                                                                                v32 := GenS3{F0: 516, F1: ((^(v13)) != v13), F2: 127}
                                                                                _ = v32
                                                                            } else {
                                                                                _ = len(v4)
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v2[fi10] = (!true)
                                                                                    default:
                                                                                        v3 += v3
                                                                                    }
                                                                                    v26 = (v26 >> ((v15) & 63))
                                                                                    v26 -= v26
                                                                                    _ = len(v4)
                                                                                default:
                                                                                    v26 = (^(v26))
                                                                                    v15 += (v15 << (((^(v15))) & 63))
                                                                                    var v33 uint16 = (v13 - v13)
                                                                                    _ = v33
                                                                                }
                                                                                if (2 >= 198) {
                                                                                    if ((false || (-1 < 86)) && true) {
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            var v34 int8 = 1
                                                                                            var v35 int16 = v8
                                                                                            _ = v34
                                                                                            _ = v35
                                                                                        case "\n\t":
                                                                                            v36 := GenS4{F0: 390, F1: fi9, F2: -348}
                                                                                            _ = v36
                                                                                        case "":
                                                                                            fmt.Println(v25, v8)
                                                                                        default:
                                                                                            v15 -= v15
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v37 uint = (^(v15))
                                                                                        _ = v37
                                                                                    case "a":
                                                                                        fmt.Println(fi10, v13, v25)
                                                                                    case "":
                                                                                        var v38 uint32 = (^(v1))
                                                                                        _ = v38
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v39 uint32 = v1
                                                                        if false {
                                                                            if (false && (!true)) {
                                                                                var v40 uint32 = v1
                                                                                v25 = -2.25
                                                                                var v41 float64 = (float64(v25) / 3.14159)
                                                                                v26 -= v26
                                                                                var v42 uint32 = v1
                                                                                _ = v40
                                                                                _ = v41
                                                                                _ = v42
                                                                            } else {
                                                                                var v43 int8 = -95
                                                                                v43 = -128
                                                                                v44 := GenS2{F0: 2, F1: -121, F2: true}
                                                                                _ = v43
                                                                                _ = v44
                                                                            }
                                                                            var v45 int16 = v8
                                                                            _ = v45
                                                                        }
                                                                        _ = v39
                                                                    }
                                                                }
                                                                v26 = 42
                                                            }
                                                            v26 = v26
                                                            switch fi10 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v13)
                                                            }
                                                        } else {
                                                            fmt.Println(fi10)
                                                        }
                                                        if true {
                                                            v4 = append(v4, (fi10 << (((v15 + v15)) & 63)))
                                                        } else {
                                                            v8 += v8
                                                            var v46 float64 = 0.1
                                                            v26 = 409
                                                            _ = v46
                                                        }
                                                        switch 32768 {
                                                        case 32767:
                                                            switch (fi10 ^ (fi10 ^ (-(fi9)))) {
                                                            case 127:
                                                                v15 = v15
                                                            case -9223372036854775808:
                                                                var v47 uint = 255
                                                                _ = v47
                                                            }
                                                        case 86:
                                                            v2[(^(fi10))] = (v13 >= (v13 - (v13 & v13)))
                                                        case -9223372036854775808:
                                                            v26 = int64(v8)
                                                        }
                                                        fmt.Println(v13, v26)
                                                    } else {
                                                        v2[(^(fi9))] = ((v15 != (^(v15))) || (false && (v3 <= v3)))
                                                    }
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v8 = v8
                                                    v8 = -32768
                                                    fmt.Println(v8)
                                                }
                                                _ = v24
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                case 128:
                                    v8 = v8
                                case -9223372036854775808:
                                    v13 -= v13
                                default:
                                    if (!(((v13 % (((v13 ^ (^(v13)))) | 1)) > v13) && ((!(fi9 <= (fi9 - fi9))) && true))) {
                                        var v48 uint32 = 1
                                        _ = v48
                                    }
                                }
                            }
                            fmt.Println(v13, v13, fi9)
                            _ = v15
                        } else {
                            var v49 uint = 18446744073709551615
                            v13 += v13
                            v50 := GenS1{F0: (-828 == 2147483647), F1: false}
                            _ = v49
                            _ = v50
                        }
                    default:
                        v2[-9223372036854775808] = true
                    }
                    switch "" {
                    case "a":
                        v3 += (v3 / v3)
                    }
                    var v51 int16 = v8
                    _ = v13
                    _ = v14
                    _ = v51
                }
                fmt.Println(v3, v8)
                var v52 int64 = -465
                v52 -= (v52 | 1)
                if true {
                    var v53 uint = 885
                    v53 = 42
                    _ = v53
                }
                var v54 int16 = -769
                _ = v52
                _ = v54
            } else {
                if (float64((1e10 / 1.0)) < (float64(0.0) * (1e10 - 3.14159))) {
                    v55 := GenS4{F0: 1, F1: fi9, F2: 2147483647}
                    if (!(true || ((v8 / ((v8) | 1)) == (v8 / (((v8 - (v8 * -32768))) | 1))))) {
                        var v56 float64 = -2.25
                        var v57 uint16 = 790
                        v58 := GenS3{F0: (v1 + v1), F1: true, F2: 255}
                        _ = v56
                        _ = v57
                        _ = v58
                    } else {
                        if false {
                            var v59 float32 = 0.1
                            fmt.Println(v3, fi9, v1)
                            if false {
                                var v60 int32 = -459
                                _ = v60
                            }
                            fmt.Println(v59, fi10)
                            fmt.Println(fi10, v1, v1)
                            var v61 uint64 = 2
                            _ = v59
                            _ = v61
                        } else {
                            var v62 float32 = 0.0
                            var v63 int32 = 42
                            _ = v62
                            _ = v63
                        }
                        if (789 <= 4294967295) {
                            if (2 >= 589) {
                                v1 = v1
                            } else {
                                v64 := v2[fi10]
                                var v65 int16 = (v8 / ((v8) | 1))
                                _ = v64
                                _ = v65
                            }
                            if true {
                                fmt.Println(fi9)
                                var v66 float32 = 1e10
                                _ = v66
                            }
                            var v67 int = -923
                            _ = v67
                        } else {
                            fmt.Println(fi10, v8)
                            v68 := GenS4{F0: 2, F1: (-(fi9)), F2: -1}
                            v69 := GenS5{F0: (true && (!false)), F1: 711}
                            fmt.Println(v1, v8)
                            _ = v68
                            _ = v69
                        }
                        delete(v2, fi10)
                        var v70 uint32 = v1
                        _ = v70
                    }
                    _ = v55
                } else {
                    var v71 int = 265
                    fmt.Println(v3, fi9)
                    _ = v71
                }
                v3 -= v3
                v72 := GenS2{F0: 42, F1: -2147483648, F2: true}
                _ = v72
            }
            fmt.Println(v3, fi9, fi10)
            _ = fi10
        }
        _ = fi9
    }
    v3 += 0.0
    v8 = v8
    if (!false) {
        v73 := GenS6{F0: 255, F1: 18446744073709551615}
        _ = v73
    } else {
        var v74 float64 = (v3 + ((1e10 + v3) + (3.14159 - float64(0.1))))
        if (9223372036854775807 > 42) {
            var v75 int16 = v8
            var v76 int64 = 1
            v4 = append(v4, 2147483647)
            _ = v75
            _ = v76
        }
        _ = v74
    }
    v1 = (^(v1))
    v3 += 1.5
    _ = len(v4)
    v77 := GenS6{F0: 18, F1: 1}
    v1 += (v1 >> ((731) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v8
}
