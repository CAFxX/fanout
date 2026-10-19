package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 int16; F1 int64; F2 uint16 }
type GenS9 struct { F0 uint64; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 int8, p2 uint32) int16 {
    var v4 float64 = 1.5
    var v5 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -299
}

func main() {
    v1 := []string{"\n\t"}
    var v2 bool = (true || (458 == -35))
    var v3 int = 32767
    var v6 float32 = 1.5
    v6 += ((v6 * (-2.25 * v6)) + ((float32(-2.25) + 1e10) / v6))
    v6 -= 1e10
    if false {
        fmt.Println(v6, v6)
    }
    var v7 uint64 = 1
    v7 = v7
    fmt.Println(v2, v6)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            switch (fi9 ^ 32767) {
            case -9223372036854775808:
                v6 = v6
            default:
                var v10 uint32 = 2147483648
                var v11 string = (("hello" + ("" + "\n\t")) + "\n\t")
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: -9223372036854775808}
            v3 -= 32768
            var v13 int8 = int8(v3)
            v3 = fi9
            v14 := GenS2{F0: 42}
            var v15 int32 = 42
            v7 += 523
            _ = fi9
            _ = v12
            _ = v13
            _ = v14
            _ = v15
        }
        if ((!false) || v2) {
            v1 = append(v1, "foo bar")
            for fi16 := 0; fi16 < 2; fi16++ {
                if (-43 <= 100) {
                    var v17 bool = ((-1 > -229) || (!(v7 > (^(v7)))))
                    v7 = (v7 % (((v7 ^ (v7 ^ (v7 ^ 736)))) | 1))
                    _ = v17
                }
                v3 -= (v3 << ((41) & 63))
                _ = fi16
            }
            switch "" {
            case "foo bar":
                v3 -= v3
            case "":
                var v18 int = (fi8 + 4294967295)
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS1{F0: 658}
                v21 := GenS3{F0: 2, F1: -128}
                v22 := GenS1{F0: -178}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = (1.5 <= float64(3.14159))
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := GenS1{F0: 0}
                var v25 uint8 = 1
                v3 -= v3
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS3{F0: 0, F1: -81}
            _ = v26
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 uint16 = 0
            _ = len(v1)
            v28 -= 278
            v1 = append(v1, "foo bar")
            _ = fi27
            _ = v28
        }
        _ = fi8
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v6 = (v6 / float32(-2.25))
        for fi30 := 0; fi30 < 2; fi30++ {
            v3 = v3
            if false {
                v31 := GenS3{F0: 1, F1: -68}
                fmt.Println(v6, v2, v7)
                v6 = ((-2.25 + (v6 * -2.25)) + (v6 - 0.0))
                _ = v31
            } else {
                switch (fi29 * fi30) {
                case 1:
                    switch "\n\t" {
                    case "":
                        switch fi30 {
                        case 1:
                            v32 := GenS2{F0: -2147483648}
                            _ = v32
                        }
                    case "a":
                        v3 -= fi29
                    }
                case 4294967295:
                    if (1e10 != (1e10 * -2.25)) {
                        v3 = (fi30 >> ((65535) & 63))
                    } else {
                        fmt.Println(fi30, v7, v7)
                        var v33 uint16 = 2
                        switch "" {
                        case "foo bar":
                            fmt.Println(v7, v33, v3)
                        case "a":
                            v33 -= 65535
                            v2 = (857 > 0)
                            if v2 {
                                var v34 uint16 = v33
                                var v35 int8 = -1
                                switch "" {
                                case "hello":
                                    switch 656 {
                                    case -833:
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            if ((((((float32(1e10) - v6) - 1.5) * v6) < (0.1 + 3.14159)) || false) || (42 <= 493)) {
                                                _ = len(v1)
                                            } else {
                                                v3 = (-(fi30))
                                                v7 = (v7 >> ((4294967295) & 63))
                                                var v36 int = (fi29 % (((^(v3))) | 1))
                                                _ = v36
                                            }
                                        case "hello":
                                            fmt.Println(v3, v3)
                                        case "":
                                            fmt.Println(v33, v3)
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        if v2 {
                                            v34 += v33
                                            fmt.Println(v34)
                                            if v2 {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (1 > 128) {
                                                        if (float32(0.1) == v6) {
                                                            v37 := GenS4{F0: ("" + "foo bar"), F1: 42, F2: int8(v33)}
                                                            _ = v37
                                                        } else {
                                                            v2 = false
                                                            fmt.Println(fi29, v2)
                                                        }
                                                        fmt.Println(v6, v34)
                                                    } else {
                                                        v38 := GenS4{F0: "", F1: -1, F2: (^(v35))}
                                                        _ = v38
                                                    }
                                                default:
                                                    var v39 int16 = 42
                                                    _ = v39
                                                }
                                            } else {
                                                v35 = -1
                                                v1 = append(v1, "hello")
                                                if v2 {
                                                    if v2 {
                                                        v35 -= v35
                                                        _ = len(v1)
                                                    }
                                                    if (239 == -2147483648) {
                                                        var v40 int64 = -1
                                                        if ((fi29 - 32767) > (-(fi30))) {
                                                            fmt.Println(v40)
                                                        }
                                                        _ = v40
                                                    }
                                                    var v41 int16 = -32768
                                                    _ = v41
                                                }
                                                v3 = 435
                                                v7 = 2
                                                v34 = (^(v34))
                                                v42 := GenS5{F0: (!(!v2)), F1: v7}
                                                _ = v42
                                            }
                                        }
                                    case 32767:
                                        var v43 bool = v2
                                        _ = v43
                                    case -787:
                                        if (false || false) {
                                            v2 = v2
                                        }
                                    }
                                case "":
                                    if (!v2) {
                                        switch fi30 {
                                        case 127:
                                            fmt.Println(v35, v6, v2)
                                        case 128:
                                            var v44 int = int(v35)
                                            if (!false) {
                                                switch -1 {
                                                case 32767:
                                                    v6 = 0.1
                                                case 127:
                                                    var v45 uint = 279
                                                    _ = v45
                                                }
                                                var v46 float64 = -2.25
                                                _ = v46
                                            } else {
                                                var v47 bool = false
                                                if (v47 && (("hello" + "") < "\n\t")) {
                                                    v44 += (-(fi29))
                                                } else {
                                                    _ = len(v1)
                                                    switch ("" + ("hello" + "\n\t")) {
                                                    case "\n\t":
                                                        v1 = append(v1, ("foo bar" + "hello"))
                                                    default:
                                                        v35 += 48
                                                    }
                                                    fmt.Println(v3)
                                                }
                                                switch (fi29 / ((v3) | 1)) {
                                                case 0:
                                                    var v48 uint16 = 65535
                                                    _ = v48
                                                case 42:
                                                    var v49 float64 = (1e10 + -2.25)
                                                    _ = v49
                                                case 9223372036854775807:
                                                    v44 += 4294967295
                                                    switch ("\n\t" + ("hello" + "a")) {
                                                    case "foo bar":
                                                        switch "foo bar" {
                                                        case "a":
                                                            v6 = float32(0.1)
                                                        case "hello":
                                                            if false {
                                                                if len(v1) > 0 {
                                                                    v1[0] = ("" + "hello")
                                                                }
                                                                v50 := GenS2{F0: 42}
                                                                v51 := GenS3{F0: (v33 ^ 42), F1: v35}
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case "foo bar":
                                                            _ = len(v1)
                                                        default:
                                                            if (1 <= 0) {
                                                                v52 := GenS3{F0: (^(v34)), F1: v35}
                                                                _ = v52
                                                            } else {
                                                                v7 = v7
                                                            }
                                                        }
                                                    case "a":
                                                        var v53 uint8 = 161
                                                        _ = v53
                                                    case "":
                                                        switch -2147483648 {
                                                        case 766:
                                                            fmt.Println(fi30, v7, v3)
                                                        default:
                                                            if ((!v2) && v2) {
                                                                var v54 int = -1
                                                                switch 32767 {
                                                                case 32767:
                                                                    fmt.Println(v33, v33)
                                                                }
                                                                v55 := GenS3{F0: (v34 >> ((65535) & 63)), F1: (^(v35))}
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        }
                                                    default:
                                                        v1 = append(v1, "hello")
                                                    }
                                                default:
                                                    var v56 int64 = 1
                                                    _ = v56
                                                }
                                                _ = v47
                                            }
                                            var v57 float32 = 1e10
                                            _ = v44
                                            _ = v57
                                        default:
                                            v7 = (v7 | (v7 - (v7 % ((v7) | 1))))
                                        }
                                        v58 := GenS2{F0: 0}
                                        _ = v58
                                    } else {
                                        v3 = fi29
                                    }
                                }
                                _ = v34
                                _ = v35
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                            }
                        case "":
                            v59 := GenS6{F0: v7, F1: 0, F2: -1}
                            _ = v59
                        }
                        _ = v33
                    }
                default:
                    v2 = (10 < 1)
                }
                if (((false && v2) || v2) || v2) {
                    if (0 <= 128) {
                        fmt.Println(v6, v3, v2)
                        v60 := GenS4{F0: ("hello" + "a"), F1: 42, F2: 100}
                        _ = v60
                    } else {
                        switch "a" {
                        case "hello":
                            _ = len(v1)
                        case "a":
                            switch "a" {
                            case "hello":
                                fmt.Println(v6)
                            default:
                                if (!v2) {
                                    if (42 < 0) {
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v61 uint8 = 50
                                            _ = v61
                                        }
                                        var v62 int = fi30
                                        v7 = (v7 / (((v7 * (v7 >> ((103) & 63)))) | 1))
                                        _ = v62
                                    }
                                }
                            }
                        case "foo bar":
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                        }
                        v7 = (v7 & (v7 * v7))
                        fmt.Println(v3, v3)
                        var v63 float32 = float32(0.0)
                        v3 = 494
                        _ = v63
                    }
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = len(v1)
                } else {
                    if v2 {
                        if (255 != 851) {
                            var v64 float64 = 3.14159
                            _ = v64
                        }
                        var v65 int64 = 9223372036854775807
                        _ = v65
                    } else {
                        v6 = v6
                    }
                    var v66 uint64 = 42
                    v66 = v7
                    _ = v66
                }
            }
            _ = fi30
        }
        _ = fi29
    }
    var v67 bool = (232 <= 0)
    if (v7 != 1) {
        var v68 int64 = -9
        _ = v68
    } else {
        if (820 > 0) {
            for fi69 := 0; fi69 < 15; fi69++ {
                var v70 float32 = v6
                if (2147483647 != 42) {
                    for fi71 := 0; fi71 < 18; fi71++ {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        var v72 int = (v3 % ((fi69) | 1))
                        v67 = v67
                        _ = fi71
                        _ = v72
                    }
                    var v73 float64 = 0.1
                    v1 = append(v1, ("a" + (("hello" + "foo bar") + "")))
                    _ = v73
                }
                _ = fi69
                _ = v70
            }
        }
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        for fi75 := 0; fi75 < 3; fi75++ {
            if (!false) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                var v76 string = "hello"
                if len(v1) > 0 {
                    v1[0] = (("hello" + v76) + "foo bar")
                }
                _ = v76
            }
            var v77 uint32 = 1
            switch (v3 & 100) {
            case 772:
                fmt.Println(v3, fi74, v67)
            case 987:
                v7 = 4294967295
            case -381:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v78 uint16 = 2
                var v79 string = ""
                _ = v78
                _ = v79
            }
            _ = fi75
            _ = v77
        }
        if v67 {
            v7 -= v7
        } else {
            v1 = append(v1, "hello")
            if v67 {
                v3 = 127
                switch ((((("\n\t" + "\n\t") + "hello") + ("\n\t" + "")) + "hello") + "hello") {
                case "foo bar":
                    v80 := GenS7{F0: 172, F1: (v3 | fi74)}
                    _ = v80
                }
                v81 := GenS8{F0: -1, F1: 1, F2: 32767}
                _ = v81
            } else {
                switch "a" {
                case "a":
                    var v82 float32 = v6
                    _ = v82
                case "\n\t":
                    switch (fi74 * 128) {
                    case -853:
                        v3 -= 5
                    }
                default:
                    var v83 int16 = 0
                    _ = v83
                }
                var v84 int = v3
                var v85 int = v3
                _ = v84
                _ = v85
            }
        }
        _ = fi74
    }
    v86 := GenS9{F0: (v7 * (^(v7))), F1: (("" + "") + ("foo bar" + "a")), F2: 771}
    fmt.Println(v2)
    for fi87 := 0; fi87 < 3; fi87++ {
        v6 = (v6 - (((-2.25 * float32(1e10)) + v6) / float32((-2.25 * 3.14159))))
        v3 = (v3 | fi87)
        _ = len(v1)
        _ = fi87
    }
    v67 = false
    var v88 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v67
    _ = v7
    _ = v86
    _ = v88
}
