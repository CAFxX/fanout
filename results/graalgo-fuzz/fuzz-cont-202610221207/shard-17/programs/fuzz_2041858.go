package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int; F2 uint8 }
type GenS7 struct { F0 int64; F1 string; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint32; F1 uint; F2 int8 }

func main() {
    v1 := []uint32{42, 0}
    v2 := []uint32{993}
    var v3 float32 = 1e10
    var v4 uint16 = 897
    var v5 int16 = 0
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        if ((v5 % ((v5) | 1)) == (v5 ^ v5)) {
            v8 := GenS1{F0: 1}
            var v9 int8 = 1
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    if (true || (v5 < v5)) {
        if false {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            for fi10 := 0; fi10 < 8; fi10++ {
                v4 = v4
                _ = fi10
            }
        } else {
            var v11 string = "\n\t"
            v12 := GenS1{F0: -1}
            v5 += v5
            _ = v11
            _ = v12
        }
        fmt.Println(v5, v3)
    } else {
        fmt.Println(v5)
        switch -9223372036854775808 {
        case -9223372036854775808:
            switch "foo bar" {
            case "a":
                v4 = (v4 << ((18446744073709551615) & 63))
            case "foo bar":
                v3 -= v3
            }
        case -1:
            v3 -= float32(0.0)
        case 128:
            for fi13 := 0; fi13 < 19; fi13++ {
                v5 += 1000
                v14 := GenS1{F0: 367}
                _ = fi13
                _ = v14
            }
        default:
            if false {
                var v15 float64 = (((((3.14159 / 1.5) * 0.1) / ((0.1 + 0.1) - (1.5 + 3.14159))) * float64(-2.25)) / -2.25)
                v5 += (-(v5))
                v4 = v4
                _ = v15
            } else {
                v4 = (v4 ^ v4)
                if true {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if (!false) {
                            switch 4294967295 {
                            case 32767:
                                fmt.Println(v4, fi16)
                            case -2147483648:
                                v17 := GenS1{F0: 977}
                                _ = v17
                            case -675:
                                v4 += 42
                            default:
                                v3 -= (float32(1e10) - 0.0)
                            }
                            var v18 float32 = v3
                            _ = v18
                        }
                        _ = fi16
                    }
                }
            }
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS1{F0: 1}
            var v21 uint8 = 2
            var v22 float32 = ((v3 + 0.0) / (3.14159 + (((v3 * (1e10 + (0.0 * 0.1))) / v3) + v3)))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v23 := GenS2{F0: 15, F1: 100}
    v3 = -2.25
    if (true && false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        _ = len(v1)
        v3 = v3
        for fi24 := 0; fi24 < 4; fi24++ {
            if false {
                var v25 uint = 65535
                var v26 string = ("" + "hello")
                for fi27 := 0; fi27 < 12; fi27++ {
                    v28 := GenS2{F0: 925, F1: 2}
                    var v29 uint8 = 127
                    fmt.Println(fi24, fi24, v26)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v25
                _ = v26
            } else {
                v4 -= (^(v4))
            }
            if ((-2.25 / v3) > (v3 / (-2.25 + (v3 + ((v3 - v3) - (v3 + 0.1)))))) {
                switch fi24 {
                case -831:
                    v30 := v6[65535]
                    _ = v30
                case -1:
                    fmt.Println(v5)
                case 9223372036854775807:
                    var v31 float64 = 0.1
                    _ = v31
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 uint64 = 309
                    v5 -= (v5 >> ((171) & 63))
                    _ = fi32
                    _ = v33
                }
                for fi34 := 0; fi34 < 13; fi34++ {
                    if false {
                        v35 := GenS3{F0: 1000, F1: "", F2: 1}
                        _ = len(v6)
                        _ = v35
                    } else {
                        if false {
                            switch ("\n\t" + "a") {
                            case "a":
                                v5 += v5
                            case "foo bar":
                                switch ("a" + "hello") {
                                case "foo bar":
                                    fmt.Println(fi34, v4, v3)
                                }
                            case "\n\t":
                                switch fi24 {
                                case 4294967295:
                                    fmt.Println(fi24, v3, fi34)
                                case 32767:
                                    v36 := GenS4{F0: 42, F1: "\n\t"}
                                    _ = v36
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        v37 := GenS2{F0: 0, F1: 52}
                                        _ = v37
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            v5 = v5
                                        case "foo bar":
                                            fmt.Println(v3, fi24)
                                        case "a":
                                            var v38 bool = false
                                            _ = v38
                                        }
                                    }
                                }
                            default:
                                switch fi34 {
                                case -271:
                                    v4 += 255
                                case 2147483647:
                                    var v39 int64 = int64(fi24)
                                    _ = v39
                                case 32767:
                                    if false {
                                        if (0 == -709) {
                                            v40 := GenS2{F0: 1, F1: 117}
                                            _ = v40
                                        }
                                        if (true && ((42 <= 255) || (!(v3 <= v3)))) {
                                            fmt.Println(v5, v4)
                                            switch fi34 {
                                            case -329:
                                                delete(v6, 629)
                                            case 137:
                                                v41 := GenS3{F0: v5, F1: "\n\t", F2: 173}
                                                _ = v41
                                            case -356:
                                                if (-655 == -1) {
                                                    v3 = v3
                                                    if false {
                                                        v3 += 0.1
                                                        var v42 uint32 = 489
                                                        _ = v42
                                                    } else {
                                                        if ((((float32(0.0) / v3) - v3) <= float32(3.14159)) || true) {
                                                            _ = len(v6)
                                                            v43 := GenS2{F0: 236, F1: 127}
                                                            var v44 bool = false
                                                            var v45 uint = 24
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = 1
                                                    }
                                                    if (0 <= 209) {
                                                        fmt.Println(v3, v3)
                                                        v46 := GenS5{F0: 0}
                                                        if len(v1) > 0 {
                                                            v1[0] = 753
                                                        }
                                                        v3 = 3.14159
                                                        v47 := GenS4{F0: 127, F1: "hello"}
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                }
                                            }
                                        }
                                        fmt.Println(fi24)
                                    } else {
                                        var v48 bool = (1.5 <= (0.0 + (0.1 / 0.1)))
                                        fmt.Println(v48, v48, v48)
                                        var v49 string = (("hello" + "a") + "foo bar")
                                        v48 = true
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    switch fi24 {
                                    case 322:
                                        v3 -= 0.0
                                    case -231:
                                        v50 := GenS4{F0: 73, F1: "hello"}
                                        _ = v50
                                    case 127:
                                        switch "\n\t" {
                                        case "a":
                                            switch fi24 {
                                            case 743:
                                                v5 = v5
                                            case 32767:
                                                if false {
                                                    v3 -= v3
                                                    fmt.Println(v4, v3)
                                                    var v51 uint8 = 152
                                                    _ = v51
                                                }
                                            }
                                        case "":
                                            v52 := GenS3{F0: (v5 >> ((369) & 63)), F1: "\n\t", F2: 0}
                                            _ = v52
                                        }
                                    default:
                                        v4 = v4
                                    }
                                }
                            }
                            v5 -= 227
                            fmt.Println(v3)
                        } else {
                            v4 += 234
                            switch fi34 {
                            case 127:
                                v3 += 1e10
                            case 0:
                                var v53 int64 = 9223372036854775807
                                _ = v53
                            default:
                                if (255 <= 42) {
                                    fmt.Println(fi34, fi34)
                                }
                            }
                        }
                        v3 += (float32(-2.25) * 3.14159)
                    }
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        if true {
                            switch (fi24 * -9223372036854775808) {
                            case 127:
                                if true {
                                    switch 62 {
                                    case 1:
                                        v54 := GenS6{F0: 760, F1: (fi34 & int(v5)), F2: 125}
                                        _ = v54
                                    default:
                                        var v55 uint8 = 128
                                        _ = v55
                                    }
                                } else {
                                    _ = len(v2)
                                    v56 := GenS7{F0: 1, F1: (("" + "") + "hello"), F2: 1}
                                    _ = v56
                                }
                            case 128:
                                v3 += v3
                            }
                        } else {
                            v3 -= 1e10
                            var v57 int = (fi24 & (fi24 + -501))
                            _ = v57
                        }
                    case "":
                        v5 = v5
                    default:
                        if (537 >= 360) {
                            var v58 int8 = 100
                            _ = v58
                        } else {
                            if ((1e10 / -2.25) != (3.14159 - 0.0)) {
                                v3 -= 1.5
                                v59 := GenS3{F0: (v5 % ((v5) | 1)), F1: "hello", F2: 255}
                                var v60 uint32 = 693
                                _ = v59
                                _ = v60
                            }
                            if ((!true) || (true && (true && true))) {
                                if (!true) {
                                    v61 := GenS6{F0: 4294967295, F1: -2147483648, F2: 42}
                                    v2 = append(v2, 551)
                                    _ = v61
                                } else {
                                    v4 = (^(v4))
                                    var v62 uint64 = 351
                                    _ = v62
                                }
                                var v63 uint64 = 18446744073709551615
                                _ = v63
                            }
                        }
                    }
                    if (true || (!(true || true))) {
                        v64 := GenS8{F0: (v4 << ((649) & 63)), F1: 1, F2: 1}
                        if ((float64(0.1) / 3.14159) == 1e10) {
                            v65 := GenS9{F0: ""}
                            _ = v65
                        }
                        _ = v64
                    }
                    _ = fi34
                }
            }
            v66 := v6[243]
            _ = fi24
            _ = v66
        }
    }
    delete(v6, 63)
    var v67 int8 = -114
    v68 := GenS10{F0: 0, F1: 255, F2: int8(v5)}
    var v69 float32 = 0.0
    switch 4294967295 {
    case 744:
        v5 -= (-(v5))
    case 2147483647:
        fmt.Println(v3, v3, v69)
    case 32767:
        for fi70 := 0; fi70 < 12; fi70++ {
            v69 = float32(((float64((v3 + 0.0)) / 0.1) + ((-2.25 + 1e10) * -2.25)))
            _ = fi70
        }
    default:
        for fi71 := 0; fi71 < 4; fi71++ {
            fmt.Println(v67, fi71)
            for fi72 := 0; fi72 < 14; fi72++ {
                v6[0] = 4294967295
                v73 := GenS8{F0: v4, F1: 72, F2: uint32(v4)}
                var v74 int8 = (^(v67))
                delete(v6, uint32(v4))
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = fi71
        }
    }
    for fi75 := 0; fi75 < 12; fi75++ {
        var v76 int16 = (v5 + v5)
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 18; fi77++ {
        var v78 float64 = 0.0
        _ = fi77
        _ = v78
    }
    fmt.Println(v69, v69, v67)
    v4 = (v4 * v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
}
