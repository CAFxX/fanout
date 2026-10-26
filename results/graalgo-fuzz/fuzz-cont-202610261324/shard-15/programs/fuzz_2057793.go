package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int64; F1 bool; F2 int32 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 uint16; F1 int8; F2 int64 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 bool; F1 int16; F2 int }
func genFn0(p0 int8, p1 uint32, p2 int16) int8 {
    fmt.Println(p0)
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            p2 = int16(v4)
            v7 := []int{(fi6 * 32768), (fi6 * fi5), fi6, fi6}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi6
            _ = v7
        }
        v4 -= (v4 & 128)
        _ = fi5
    }
    switch "hello" {
    case "":
        if ("foo bar" >= "hello") {
            if false {
                if true {
                    var v8 uint8 = 255
                    _ = v8
                }
                fmt.Println(p2)
                var v9 int16 = 304
                _ = v9
            }
            p1 = p1
        } else {
            v10 := GenS1{F0: "\n\t", F1: 27, F2: 1}
            var v11 float32 = 1.5
            _ = v10
            _ = v11
        }
    default:
        var v12 uint16 = 2
        _ = v12
    }
    p0 -= (p0 * p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}
func genFn1(p0 int, p1 uint32, p2 uint32) int32 {
    v13 := GenS1{F0: ("a" + ((("hello" + "a") + "") + "foo bar")), F1: 127, F2: 2}
    switch "foo bar" {
    case "a":
        v14 := GenS1{F0: (("a" + "\n\t") + "\n\t"), F1: 0, F2: 18446744073709551615}
        _ = v14
    }
    v15 := []int{-385}
    v16 := GenS1{F0: ("hello" + "a"), F1: -1, F2: 2}
    var v17 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := []string{"foo bar", "a", "foo bar"}
    var v2 string = ""
    v3 := make(map[string]uint32)
    if false {
        for fi18 := 0; fi18 < 11; fi18++ {
            switch (fi18 | fi18) {
            case 128:
                var v19 uint16 = 42
                _ = v19
            default:
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 int16 = 106
                    _ = fi20
                    _ = v21
                }
            }
            if (false && (925 >= 2)) {
                v1 = append(v1, (v2 + v2))
                v2 = (v2 + "a")
                if len(v1) > 0 {
                    v1[0] = (v2 + v2)
                }
            } else {
                v22 := GenS1{F0: v2, F1: 100, F2: 608}
                v23 := v3[(v2 + (v2 + "hello"))]
                fmt.Println(fi18)
                _ = v22
                _ = v23
            }
            _ = fi18
        }
        if true {
            var v24 uint32 = 0
            _ = v24
        }
        switch ("foo bar" + (v2 + (((v2 + "foo bar") + v2) + ((v2 + "\n\t") + "a")))) {
        case "":
            if len(v1) > 0 {
                v1[0] = v2
            }
        case "a":
            fmt.Println(v2)
        }
    } else {
        switch 32768 {
        case 90:
            if ((!true) && (-499 >= -337)) {
                switch ((("" + v2) + "foo bar") + ("foo bar" + v2)) {
                case "":
                    var v25 int64 = -1
                    _ = v25
                default:
                    var v26 bool = (9223372036854775807 != 226)
                    _ = v26
                }
                if true {
                    v2 = "a"
                }
                v27 := GenS1{F0: v2, F1: -1, F2: 409}
                _ = v27
            } else {
                delete(v3, "foo bar")
                for fi28 := 0; fi28 < 12; fi28++ {
                    fmt.Println(fi28)
                    switch (fi28 + fi28) {
                    case 0:
                        for fi29 := 0; fi29 < 1; fi29++ {
                            var v30 uint16 = 832
                            var v31 uint64 = 42
                            switch ("a" + (v2 + v2)) {
                            case "foo bar":
                                v31 = (v31 & v31)
                            case "":
                                if (-32768 <= -1) {
                                    if ((false && (1 >= 42)) && true) {
                                        var v32 uint64 = v31
                                        _ = v32
                                    } else {
                                        fmt.Println(v2, v30, v31)
                                        var v33 uint32 = 0
                                        v30 -= (v30 & v30)
                                        _ = v33
                                    }
                                } else {
                                    if (((3.14159 + 1.5) * 0.0) < float32(0.1)) {
                                        if (253 >= 172) {
                                            switch (-(fi29)) {
                                            case 1:
                                                v30 = (v30 / (((v30 & 1)) | 1))
                                            case 32768:
                                                switch (fi28 - fi29) {
                                                case 127:
                                                    v31 = uint64(v30)
                                                }
                                            case 2147483647:
                                                var v34 uint32 = 1
                                                _ = v34
                                            default:
                                                v30 -= (^(v30))
                                            }
                                        }
                                    } else {
                                        if false {
                                            _ = len(v1)
                                            v35 := GenS2{F0: 42}
                                            switch fi28 {
                                            case -726:
                                                switch fi29 {
                                                case 898:
                                                    var v36 int32 = genFn1(-212, 552, 1)
                                                    if ((0.0 * (((float64(((-2.25 * 1e10) + 1.5)) * 1e10) - -2.25) * 1e10)) < 0.1) {
                                                        v36 -= v36
                                                        v37 := GenS3{F0: "", F1: 42}
                                                        if false {
                                                            fmt.Println(v30, fi28, v2)
                                                            var v38 int16 = 0
                                                            var v39 int = fi28
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v37
                                                    }
                                                    _ = v36
                                                }
                                                if false {
                                                    v30 += v30
                                                    v40 := v3[(v2 + "\n\t")]
                                                    fmt.Println(fi29)
                                                    _ = v40
                                                }
                                                delete(v3, ("hello" + ("" + v2)))
                                                if true {
                                                    v41 := GenS2{F0: -1}
                                                    v42 := GenS4{F0: 1, F1: (!(!((v30 >> ((0) & 63)) > v30))), F2: 800}
                                                    _ = v41
                                                    _ = v42
                                                }
                                                fmt.Println(fi29, v31, v2)
                                            case 2147483647:
                                                v31 = (v31 * (v31 - v31))
                                            }
                                            v43 := GenS4{F0: -450, F1: true, F2: -1}
                                            _ = v35
                                            _ = v43
                                        } else {
                                            var v44 int32 = -1
                                            if (!false) {
                                                _ = len(v3)
                                            }
                                            v45 := GenS2{F0: 9223372036854775807}
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                }
                                switch v2 {
                                case "a":
                                    v30 = (v30 & v30)
                                case "\n\t":
                                    if true {
                                        fmt.Println(v2, fi28, v30)
                                    } else {
                                        fmt.Println(v31, fi28, fi28)
                                        v30 = v30
                                        var v46 int = fi29
                                        _ = v46
                                    }
                                }
                            case "a":
                                var v47 uint8 = 79
                                _ = v47
                            default:
                                v48 := GenS5{F0: 0, F1: genFn1(9223372036854775807, 42, 1)}
                                _ = v48
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                    v3[((v2 + v2) + v2)] = 0
                    _ = fi28
                }
                for fi49 := 0; fi49 < 16; fi49++ {
                    if (1 <= 921) {
                        v50 := GenS6{F0: 488, F1: 1, F2: 42}
                        var v51 uint = 0
                        _ = v50
                        _ = v51
                    } else {
                        if false {
                            switch v2 {
                            case "hello":
                                v1 = append(v1, v2)
                            }
                            var v52 float32 = 3.14159
                            _ = v52
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi53 := 0; fi53 < 16; fi53++ {
                        _ = len(v1)
                        _ = fi53
                    }
                    _ = fi49
                }
            }
        case 0:
            var v54 uint = 42
            _ = v54
        default:
            v55 := GenS6{F0: 42, F1: -99, F2: -709}
            _ = v55
        }
        v56 := GenS7{F0: v2, F1: -32768}
        _ = v56
    }
    v2 = ("hello" + "foo bar")
    v2 = "hello"
    v2 = v2
    var v57 int16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 5; fi59++ {
            var v60 float64 = (-2.25 * 1e10)
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v57)
    v1 = append(v1, v2)
    switch 2147483647 {
    case 128:
        switch 0 {
        case 797:
            v2 = (v2 + "")
        case -52:
            switch (v2 + "") {
            case "hello":
                for fi61 := 0; fi61 < 10; fi61++ {
                    var v62 uint8 = 127
                    _ = fi61
                    _ = v62
                }
            case "":
                delete(v3, v2)
            case "foo bar":
                switch ("foo bar" + (v2 + ((v2 + v2) + (v2 + "a")))) {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    if false {
                        for fi63 := 0; fi63 < 20; fi63++ {
                            v64 := GenS6{F0: 255, F1: 1, F2: int64(v57)}
                            v2 = (v2 + v2)
                            v1 = append(v1, v2)
                            _ = fi63
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 13; fi65++ {
                            v2 = "hello"
                            v57 -= 524
                            var v66 uint = 202
                            _ = fi65
                            _ = v66
                        }
                        delete(v3, "\n\t")
                    } else {
                        v2 = v2
                        switch 4294967295 {
                        case -853:
                            if true {
                                v67 := GenS3{F0: (v2 + "foo bar"), F1: 127}
                                _ = v67
                            }
                        case 329:
                            v2 = v2
                        case 2147483647:
                            fmt.Println(v2, v57)
                        default:
                            v3[v2] = 243
                        }
                        if false {
                            _ = len(v3)
                        }
                    }
                default:
                    for fi68 := 0; fi68 < 6; fi68++ {
                        fmt.Println(v57)
                        var v69 float32 = (3.14159 / 1.5)
                        var v70 string = v2
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                }
            default:
                _ = len(v3)
            }
        case -288:
            v71 := GenS8{F0: true, F1: (v57 - (v57 ^ (v57 << ((253) & 63)))), F2: -167}
            _ = v71
        default:
            var v72 uint32 = 584
            _ = v72
        }
    }
    v2 = v2
    fmt.Println(v57, v2)
    var v73 bool = (!false)
    for fi74 := 0; fi74 < 3; fi74++ {
        v73 = true
        _ = fi74
    }
    var v75 uint64 = 4294967295
    for fi76 := 0; fi76 < 5; fi76++ {
        v75 -= 912
        if (false && false) {
            var v77 uint32 = 165
            var v78 bool = true
            _ = v77
            _ = v78
        } else {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v79 := GenS4{F0: 944, F1: false, F2: -531}
                var v80 int = (fi76 ^ 941)
                v1 = append(v1, (v2 + "\n\t"))
                _ = v79
                _ = v80
            }
            var v81 int8 = 100
            _ = v81
        }
        _ = fi76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v73
    _ = v75
}
