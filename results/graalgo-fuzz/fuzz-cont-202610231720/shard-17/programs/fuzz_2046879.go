package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int16; F2 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS10 struct { F0 int64; F1 int16; F2 int8 }
type GenS11 struct { F0 uint16; F1 uint8; F2 uint }
type GenS12 struct { F0 int8; F1 string; F2 int16 }
type GenS13 struct { F0 uint8; F1 int; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("foo bar" + (("\n\t" + (("\n\t" + "") + ("" + "a"))) + ""))
    v3 := []string{}
    v4 := []string{(v2 + ((v2 + "foo bar") + v2))}
    v5 := make(map[int]string)
    var v6 float64 = 0.0
    v2 = ""
    v7 := GenS1{F0: (1 == 18446744073709551615), F1: 0}
    switch v2 {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = v2
        }
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            v2 = (v2 + v2)
            switch fi8 {
            case -397:
                var v9 int16 = -1
                _ = v9
            case 128:
                for fi10 := 0; fi10 < 19; fi10++ {
                    switch v2 {
                    case "a":
                        if ((5 <= 0) && ((v2 < v2) && ((float32(((v6 + 0.1) - (v6 + v6))) + ((((((((1e10 / 1.5) + 1.5) + 1e10) * 3.14159) + 3.14159) * 0.1) - (float32((1e10 / v6)) - float32(float64(1e10)))) / 3.14159)) >= 1e10))) {
                            var v11 int32 = 0
                            v11 = -2147483648
                            var v12 bool = true
                            if ((42 != 42) && v12) {
                                if (float32(v6) != 3.14159) {
                                    delete(v1, "foo bar")
                                } else {
                                    if v12 {
                                        var v13 int32 = v11
                                        _ = v13
                                    } else {
                                        var v14 bool = v12
                                        v14 = v12
                                        if (1 <= 42) {
                                            var v15 int64 = 1
                                            v11 += (v11 >> ((18446744073709551615) & 63))
                                            if v12 {
                                                v1[(v2 + "a")] = 2147483648
                                                v14 = (v11 > v11)
                                            } else {
                                                v16 := GenS1{F0: (((float32(v6) - 1e10) == -2.25) && (!(!(255 <= 127)))), F1: 130}
                                                var v17 int64 = (v15 ^ v15)
                                                _ = v16
                                                _ = v17
                                            }
                                            _ = v15
                                        }
                                        _ = v14
                                    }
                                }
                            } else {
                                if v12 {
                                    var v18 float32 = -2.25
                                    if v12 {
                                        switch v2 {
                                        case "a":
                                            switch fi10 {
                                            case 9223372036854775807:
                                                v18 -= ((v18 + (3.14159 * (1.5 + (v18 / (v18 / (((0.0 - v18) / v18) * v18)))))) * v18)
                                                v19 := GenS1{F0: (false && (2 > 900)), F1: 805}
                                                _ = v19
                                            case 4294967295:
                                                var v20 int64 = -1
                                                _ = v20
                                            }
                                        case "hello":
                                            if (2 >= 0) {
                                                v21 := GenS2{F0: 2, F1: ("\n\t" + v2)}
                                                _ = v21
                                            }
                                            v11 = -1
                                        case "\n\t":
                                            v18 = v18
                                        }
                                        switch fi8 {
                                        case -2147483648:
                                            if (-202 != (fi8 | (-(fi10)))) {
                                                v22 := v1[(v2 + v2)]
                                                v22 = (v22 - v22)
                                                _ = v22
                                            }
                                        case 4294967295:
                                            v23 := v1["hello"]
                                            _ = v23
                                        case -985:
                                            if v12 {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if v12 {
                                                    var v24 uint8 = 42
                                                    v25 := GenS2{F0: (v24 * v24), F1: ("a" + v2)}
                                                    v26 := GenS3{F0: v2, F1: v2, F2: 885}
                                                    _ = v24
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    var v27 int16 = -1
                                                    v12 = (((255 > 2) && v12) && v12)
                                                    if v12 {
                                                        v18 = v18
                                                    } else {
                                                        v6 += (v6 + v6)
                                                    }
                                                    _ = v27
                                                }
                                                v12 = v12
                                                var v28 uint32 = 1
                                                var v29 bool = (v6 == (v6 / v6))
                                                if (1000 < 1000) {
                                                    fmt.Println(v2, v29)
                                                    v30 := GenS4{F0: (v2 + ("" + "\n\t"))}
                                                    v31 := GenS4{F0: ((v2 + v2) + v2)}
                                                    var v32 string = "a"
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 string = (((v2 + "\n\t") + (((v2 + v2) + v2) + (((v2 + "hello") + v2) + v2))) + ("\n\t" + "foo bar"))
                                                    _ = v33
                                                }
                                                v28 = 1
                                                var v34 int8 = 127
                                                _ = v28
                                                _ = v29
                                                _ = v34
                                            } else {
                                                v35 := GenS4{F0: v2}
                                                v12 = (401 != 2)
                                                v36 := GenS5{F0: 32767}
                                                v37 := GenS5{F0: -751}
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                        v18 += (0.0 * 1.5)
                                    }
                                    v2 = "\n\t"
                                    _ = v18
                                } else {
                                    var v38 int64 = -1
                                    v6 = v6
                                    _ = v38
                                }
                                fmt.Println(v6)
                            }
                            _ = v11
                            _ = v12
                        } else {
                            v39 := GenS6{F0: 127}
                            fmt.Println(fi10, fi8)
                            v40 := GenS7{F0: 1, F1: 926, F2: 1}
                            _ = v39
                            _ = v40
                        }
                    case "foo bar":
                        v41 := GenS6{F0: 42}
                        _ = v41
                    case "\n\t":
                        if (true && (1e10 < 1.5)) {
                            v6 = (v6 / v6)
                            v6 = 0.1
                            var v42 int8 = -128
                            _ = v42
                        } else {
                            var v43 uint8 = 1
                            switch ((((v2 + v2) + ((v2 + (v2 + "")) + ("foo bar" + "hello"))) + v2) + v2) {
                            case "\n\t":
                                v44 := GenS8{F0: 366}
                                _ = v44
                            case "foo bar":
                                switch ("foo bar" + v2) {
                                case "foo bar":
                                    switch fi10 {
                                    case -1:
                                        v45 := GenS8{F0: 645}
                                        _ = v45
                                    case -187:
                                        _ = len(v1)
                                    }
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = ((v2 + ("a" + v2)) + v2)
                                    }
                                }
                            case "":
                                v46 := GenS9{F0: 42, F1: 2, F2: 2}
                                v43 -= v43
                                _ = v46
                            default:
                                v6 = float64(-2.25)
                            }
                            var v47 int64 = -9223372036854775808
                            _ = v43
                            _ = v47
                        }
                        v2 = "hello"
                        if (float32(v6) == float32(v6)) {
                            if (127 == 121) {
                                fmt.Println(v6, fi8, v6)
                                if true {
                                    v48 := v1[(v2 + "\n\t")]
                                    if false {
                                        var v49 int32 = 42
                                        v50 := GenS5{F0: -58}
                                        fmt.Println(v2)
                                        var v51 uint64 = 1
                                        switch (fi8 & (fi10 ^ fi8)) {
                                        case 42:
                                            fmt.Println(v2, v49, v51)
                                        }
                                        _ = len(v5)
                                        v52 := v5[(fi10 / (((-(fi10))) | 1))]
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    } else {
                                        switch (fi8 * fi8) {
                                        case -154:
                                            fmt.Println(v6)
                                        default:
                                            _ = len(v3)
                                        }
                                        v53 := GenS10{F0: 0, F1: 32767, F2: 42}
                                        fmt.Println(v2)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            var v54 float64 = ((v6 / v6) * (v6 * v6))
                                            if false {
                                                var v55 uint8 = 255
                                                _ = v55
                                            }
                                            fmt.Println(fi8, v6)
                                            _ = v54
                                        } else {
                                            v48 = v48
                                            switch (fi10 << ((4294967295) & 63)) {
                                            case -1:
                                                switch v2 {
                                                case "":
                                                    if false {
                                                        var v56 int = (fi10 + 1)
                                                        delete(v1, v2)
                                                        v57 := GenS4{F0: ""}
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                case "\n\t":
                                                    v2 = (v2 + (("hello" + v2) + v2))
                                                }
                                            case 127:
                                                var v58 float32 = 3.14159
                                                _ = v58
                                            case -778:
                                                v59 := GenS1{F0: true, F1: 2}
                                                _ = v59
                                            default:
                                                v48 += v48
                                            }
                                        }
                                        _ = v53
                                    }
                                    _ = v48
                                } else {
                                    if (!(128 > 0)) {
                                        v60 := GenS7{F0: 9223372036854775807, F1: 932, F2: 24}
                                        var v61 uint8 = 255
                                        _ = v60
                                        _ = v61
                                    }
                                    v2 = v2
                                }
                                var v62 float32 = ((1.5 * (float32((1.5 + 1e10)) - 3.14159)) - 1.5)
                                _ = v62
                            }
                            var v63 float64 = v6
                            _ = v63
                        }
                    default:
                        if true {
                            _ = len(v4)
                        }
                        v64 := GenS11{F0: 237, F1: 2, F2: 1}
                        _ = v64
                    }
                    var v65 float32 = 1e10
                    _ = fi10
                    _ = v65
                }
            }
            delete(v1, v2)
            _ = fi8
        }
    }
    var v66 int64 = -1
    var v67 uint32 = 2147483647
    if ((("" + (v2 + v2)) + "a") > (v2 + "a")) {
        var v68 string = v2
        _ = v68
    }
    var v69 int64 = 9223372036854775807
    for fi70 := 0; fi70 < 14; fi70++ {
        for fi71 := 0; fi71 < 9; fi71++ {
            v72 := GenS12{F0: 1, F1: v2, F2: 0}
            v73 := GenS13{F0: 0, F1: (fi70 / (((fi70 + (fi71 >> ((255) & 63)))) | 1)), F2: (v69 % (((v69 / ((v66) | 1))) | 1))}
            _ = fi71
            _ = v72
            _ = v73
        }
        var v74 int16 = 321
        _ = fi70
        _ = v74
    }
    var v75 int = 128
    for fi76 := 0; fi76 < 11; fi76++ {
        var v77 int64 = v66
        var v78 bool = true
        _ = fi76
        _ = v77
        _ = v78
    }
    var v79 int8 = -112
    v75 -= v75
    if true {
        for fi80 := 0; fi80 < 3; fi80++ {
            v67 += 2147483647
            _ = fi80
        }
        _ = len(v3)
        fmt.Println(v69)
    } else {
        if (!(42 >= uint8(v67))) {
            var v81 bool = (-1 <= -2147483648)
            _ = len(v4)
            _ = len(v4)
            _ = v81
        } else {
            var v82 float32 = (1.5 / 3.14159)
            _ = v82
        }
        v6 = v6
    }
    fmt.Println(v6, v6, v69)
    v83 := GenS3{F0: v2, F1: v2, F2: 77}
    for fi84 := 0; fi84 < 4; fi84++ {
        fmt.Println(v66, v75, v67)
        fmt.Println(fi84, fi84)
        _ = fi84
    }
    switch v2 {
    case "hello":
        var v85 uint16 = 32767
        _ = v85
    case "\n\t":
        v86 := GenS3{F0: v2, F1: ("a" + "a"), F2: 350}
        _ = v86
    }
    v87 := GenS7{F0: 1, F1: 1000, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v75:", v75)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v69
    _ = v7
    _ = v75
    _ = v79
    _ = v83
    _ = v87
}
