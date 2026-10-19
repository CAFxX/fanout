package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 int16; F1 string; F2 int8 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 int; F1 bool; F2 int }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 string; F1 int64; F2 int64 }
func genFn0() uint32 {
    v5 := []int{}
    var v6 uint32 = 1
    fmt.Println(v6)
    v7 := GenS1{F0: 0, F1: 0, F2: 127}
    var v8 int64 = -9223372036854775808
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (v6 << ((4294967295) & 63))
}

func main() {
    var v1 bool = true
    var v2 int32 = 42
    var v3 int16 = 90
    v4 := []uint32{2147483647, 2147483647, 82, 723}
    for fi9 := 0; fi9 < 16; fi9++ {
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = fi9
    }
    var v10 int8 = 0
    v11 := GenS2{F0: 880}
    if v1 {
        if true {
            v12 := GenS2{F0: 95}
            _ = v12
        } else {
            if (int16(v10) == v3) {
                var v13 string = "a"
                for fi14 := 0; fi14 < 5; fi14++ {
                    if v1 {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            if (float64((float32(((0.0 * (3.14159 - (3.14159 / 0.1))) * ((1e10 * (3.14159 / (-2.25 * 0.1))) * 0.1))) - -2.25)) >= (((-2.25 + (3.14159 - 0.0)) * 3.14159) + (1.5 * (1.5 * (float64(0.1) - 1e10))))) {
                                switch v13 {
                                case "hello":
                                    v10 = v10
                                default:
                                    v10 += (v10 % (((v10 / (((v10 >> ((1) & 63))) | 1))) | 1))
                                }
                                var v16 uint32 = 822
                                v17 := GenS2{F0: 1}
                                _ = v16
                                _ = v17
                            }
                            v1 = (v2 <= (^(v2)))
                            _ = fi15
                        }
                        switch "a" {
                        case "a":
                            switch "" {
                            case "foo bar":
                                _ = len(v4)
                            default:
                                for fi18 := 0; fi18 < 2; fi18++ {
                                    switch fi14 {
                                    case 804:
                                        v3 = -32768
                                    case 32767:
                                        var v19 int = (fi14 / ((-2147483648) | 1))
                                        _ = v19
                                    }
                                    v20 := GenS3{F0: 648, F1: 0, F2: ((fi14 - (fi14 << ((65535) & 63))) == (fi18 + -142))}
                                    v3 = (v3 ^ (v3 % (((v3 + v3)) | 1)))
                                    v10 -= (v10 | v10)
                                    v21 := GenS4{F0: (-(v3)), F1: (v13 + v13), F2: (-(v10))}
                                    _ = fi18
                                    _ = v20
                                    _ = v21
                                }
                            }
                        case "":
                            for fi22 := 0; fi22 < 19; fi22++ {
                                switch "a" {
                                case "":
                                    var v23 uint64 = 0
                                    _ = v23
                                case "a":
                                    v24 := GenS3{F0: 0, F1: 2, F2: (!v1)}
                                    _ = v24
                                case "\n\t":
                                    fmt.Println(v10, v2)
                                default:
                                    var v25 bool = (255 <= 4294967295)
                                    _ = v25
                                }
                                fmt.Println(fi14, v3, v3)
                                var v26 uint8 = 128
                                _ = fi22
                                _ = v26
                            }
                        default:
                            if (966 <= 18446744073709551615) {
                                v27 := GenS5{F0: (fi14 | 32768), F1: 689}
                                _ = v27
                            }
                        }
                        fmt.Println(v10, v2, v3)
                    }
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch 9223372036854775807 {
                        case 32768:
                            if (((!(!(v3 == (v3 * 1000)))) || ((255 >= 4294967295) && false)) || true) {
                                if len(v4) > 0 {
                                    v4[0] = 2
                                }
                            } else {
                                var v28 uint8 = 25
                                var v29 int64 = 1
                                v4 = append(v4, 65535)
                                _ = v28
                                _ = v29
                            }
                        default:
                            var v30 uint16 = 2
                            _ = v30
                        }
                    }
                    v31 := GenS4{F0: (v3 + v3), F1: v13, F2: 0}
                    _ = fi14
                    _ = v31
                }
                for fi32 := 0; fi32 < 18; fi32++ {
                    if ((^(v3)) <= 42) {
                        var v33 uint8 = 0
                        v13 = "hello"
                        _ = v33
                    }
                    _ = fi32
                }
                _ = v13
            } else {
                if true {
                    v3 += (v3 ^ v3)
                } else {
                    if v1 {
                        var v34 int64 = -9223372036854775808
                        var v35 int = 32768
                        v34 = (v34 % (((^(v34))) | 1))
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int = 2147483647
                        for fi37 := 0; fi37 < 19; fi37++ {
                            switch fi37 {
                            case 32767:
                                _ = len(v4)
                            case 127:
                                for fi38 := 0; fi38 < 4; fi38++ {
                                    if v1 {
                                        if len(v4) > 0 {
                                            v4[0] = 42
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            if v1 {
                                                v39 := GenS6{F0: (fi37 / ((v36) | 1)), F1: 42}
                                                fmt.Println(fi38)
                                                var v40 string = "hello"
                                                var v41 int64 = 1
                                                v36 = fi37
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            } else {
                                                var v42 string = "\n\t"
                                                _ = v42
                                            }
                                        case "":
                                            var v43 int = v36
                                            _ = v43
                                        default:
                                            v4 = append(v4, 2)
                                        }
                                        v4 = append(v4, 1)
                                    } else {
                                        v10 -= (v10 & 1)
                                        if v1 {
                                            fmt.Println(v2, v2, v3)
                                            if (4294967295 == 65535) {
                                                if true {
                                                    var v44 float64 = 1.5
                                                    v2 += 2147483647
                                                    _ = v44
                                                } else {
                                                    var v45 int32 = 2147483647
                                                    _ = len(v4)
                                                    v46 := GenS3{F0: 4294967295, F1: 42, F2: v1}
                                                    _ = v45
                                                    _ = v46
                                                }
                                            }
                                            var v47 int = (fi38 / ((fi37) | 1))
                                            _ = v47
                                        }
                                        v48 := GenS2{F0: 255}
                                        v1 = (fi38 >= (fi38 / ((fi37) | 1)))
                                        _ = v48
                                    }
                                    var v49 uint16 = 32767
                                    _ = fi38
                                    _ = v49
                                }
                            }
                            _ = fi37
                        }
                        fmt.Println(v3)
                        _ = v36
                    }
                }
            }
            var v50 int = 32767
            var v51 uint8 = 42
            _ = v50
            _ = v51
        }
        if v1 {
            if v1 {
                v1 = v1
            } else {
                switch "a" {
                case "":
                    v52 := GenS2{F0: 288}
                    _ = v52
                }
            }
            v4 = append(v4, 42)
            var v53 uint = 65535
            _ = v53
        } else {
            if v1 {
                v54 := GenS7{F0: false, F1: 0}
                var v55 uint = 42
                v55 += 918
                _ = v54
                _ = v55
            } else {
                if (3.14159 < 3.14159) {
                    if v1 {
                        fmt.Println(v3)
                        if false {
                            var v56 int8 = (-(v10))
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 3; fi57++ {
                                v2 = v2
                                _ = fi57
                            }
                        }
                    } else {
                        if v1 {
                            v10 = v10
                            if false {
                                if (v1 || true) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                fmt.Println(v1, v2, v2)
                            } else {
                                v10 -= v10
                                v3 += v3
                            }
                            for fi58 := 0; fi58 < 10; fi58++ {
                                v2 -= v2
                                if v1 {
                                    switch "" {
                                    case "a":
                                        v59 := GenS4{F0: v3, F1: "\n\t", F2: -128}
                                        _ = v59
                                    case "\n\t":
                                        switch (fi58 >> ((640) & 63)) {
                                        case -743:
                                            var v60 uint8 = 1
                                            _ = v60
                                        case 856:
                                            for fi61 := 0; fi61 < 13; fi61++ {
                                                var v62 int = 9223372036854775807
                                                if v1 {
                                                    _ = len(v4)
                                                    v63 := GenS8{F0: 1, F1: "foo bar"}
                                                    _ = v63
                                                }
                                                _ = fi61
                                                _ = v62
                                            }
                                        }
                                    default:
                                        var v64 uint = 18446744073709551615
                                        _ = v64
                                    }
                                }
                                _ = fi58
                            }
                        }
                    }
                    var v65 bool = (2 != 18446744073709551615)
                    _ = v65
                } else {
                    if len(v4) > 0 {
                        v4[0] = 61
                    }
                }
                if true {
                    for fi66 := 0; fi66 < 8; fi66++ {
                        v1 = v1
                        if v1 {
                            fmt.Println(v1, v3, v3)
                            fmt.Println(v1)
                            v2 = 2147483647
                        } else {
                            switch "a" {
                            case "":
                                v67 := GenS8{F0: 1, F1: ("foo bar" + ("\n\t" + "\n\t"))}
                                _ = v67
                            case "hello":
                                v3 -= 1000
                            }
                            v1 = (444 != 1)
                        }
                        _ = fi66
                    }
                }
            }
            for fi68 := 0; fi68 < 11; fi68++ {
                fmt.Println(v2)
                v2 += (-(v2))
                fmt.Println(fi68, v2)
                _ = fi68
            }
            for fi69 := 0; fi69 < 13; fi69++ {
                v70 := GenS9{F0: fi69, F1: v1, F2: (fi69 + 118)}
                _ = fi69
                _ = v70
            }
        }
    }
    for fi71 := 0; fi71 < 19; fi71++ {
        if (42 != 0) {
            if false {
                if v1 {
                    var v72 float32 = (0.1 * (((1e10 - 3.14159) - ((0.0 - 0.1) - 1.5)) / ((float32(0.1) + 0.1) - -2.25)))
                    _ = v72
                }
                v1 = false
                var v73 int = fi71
                _ = v73
            } else {
                fmt.Println(fi71)
                if (!((255 == 242) || false)) {
                    v4 = append(v4, 65535)
                    v4 = append(v4, 4294967295)
                }
                if v1 {
                    if v1 {
                        v74 := GenS1{F0: (v10 + (v10 >> ((42) & 63))), F1: 1, F2: 2}
                        var v75 string = ""
                        if v1 {
                            v4 = append(v4, 2147483647)
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            for fi76 := 0; fi76 < 19; fi76++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi76
                            }
                        }
                        _ = v74
                        _ = v75
                    }
                    v1 = v1
                    v1 = v1
                }
            }
            v1 = (!(!true))
            if false {
                var v77 int32 = -124
                switch ("a" + "") {
                case "a":
                    v10 -= -7
                case "hello":
                    switch 32768 {
                    case 32767:
                        var v78 uint = 18446744073709551615
                        _ = v78
                    }
                default:
                    fmt.Println(v10)
                }
                var v79 int64 = -1
                _ = v77
                _ = v79
            } else {
                var v80 uint16 = 65535
                var v81 uint16 = v80
                _ = len(v4)
                _ = v80
                _ = v81
            }
        } else {
            v82 := GenS10{F0: 255, F1: 352}
            _ = v82
        }
        v10 = (v10 / (((v10 >> ((42) & 63))) | 1))
        _ = fi71
    }
    var v83 int64 = 633
    var v84 int32 = v2
    var v85 bool = (128 <= 67)
    v4 = append(v4, 937)
    fmt.Println(v1)
    var v86 uint64 = 2
    fmt.Println(v3)
    if v85 {
        v87 := GenS11{F0: "", F1: (v83 / ((v83) | 1)), F2: -9223372036854775808}
        if true {
            for fi88 := 0; fi88 < 18; fi88++ {
                _ = len(v4)
                var v89 uint16 = 2
                _ = fi88
                _ = v89
            }
        } else {
            switch int(v84) {
            case -2147483648:
                _ = len(v4)
            case -9223372036854775808:
                v90 := GenS9{F0: -2147483648, F1: false, F2: 2147483647}
                _ = v90
            case -458:
                if (176 >= 255) {
                    v86 = (v86 >> ((0) & 63))
                    for fi91 := 0; fi91 < 15; fi91++ {
                        var v92 int = fi91
                        var v93 float32 = float32(0.1)
                        for fi94 := 0; fi94 < 11; fi94++ {
                            _ = len(v4)
                            var v95 uint = 4294967295
                            v86 = v86
                            var v96 int64 = v83
                            _ = fi94
                            _ = v95
                            _ = v96
                        }
                        _ = fi91
                        _ = v92
                        _ = v93
                    }
                } else {
                    var v97 uint8 = 196
                    for fi98 := 0; fi98 < 17; fi98++ {
                        for fi99 := 0; fi99 < 18; fi99++ {
                            var v100 float32 = 0.1
                            var v101 uint = 495
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                        v2 += v84
                        _ = fi98
                    }
                    _ = v97
                }
            default:
                v3 = 5
            }
        }
        v10 -= v10
        _ = v87
    } else {
        fmt.Println(v2, v83)
        v2 = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
