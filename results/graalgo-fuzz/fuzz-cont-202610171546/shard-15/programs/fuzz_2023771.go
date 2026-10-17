package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint; F2 int16 }
type GenS8 struct { F0 uint16; F1 string; F2 uint64 }
type GenS9 struct { F0 uint16; F1 uint32; F2 bool }
type GenS10 struct { F0 int64; F1 int16; F2 int64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 int16, p1 int8, p2 bool) int8 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float32 = float32(((-2.25 + 0.0) * ((-2.25 * (0.0 * 0.0)) / (0.1 * 0.1))))
        v8 := GenS1{F0: (fi6 / ((fi6) | 1))}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = 42
    v10 := GenS1{F0: -743}
    switch 32767 {
    case -9223372036854775808:
        for fi11 := 0; fi11 < 10; fi11++ {
            if true {
                var v12 uint32 = 65535
                _ = v12
            }
            fmt.Println(p0)
            fmt.Println(p1)
            _ = fi11
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}

func main() {
    var v1 uint8 = 179
    var v2 uint64 = 42
    var v3 uint8 = (v1 | 127)
    v4 := []string{"", ("\n\t" + (("a" + ("hello" + "\n\t")) + ""))}
    var v5 uint = 353
    var v13 uint32 = 2
    fmt.Println(v13, v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            switch -1 {
            case -2147483648:
                var v16 int = 32768
                _ = v16
            case 138:
                v3 = v3
                if false {
                    var v17 uint16 = 32767
                    _ = v17
                } else {
                    var v18 int8 = 42
                    _ = v18
                }
            case 0:
                _ = len(v4)
            }
            switch fi14 {
            case 79:
                if ((^(fi15)) != 9223372036854775807) {
                    switch -1 {
                    case -652:
                        v5 = v5
                        if (v2 > (v2 | 212)) {
                            v4 = append(v4, "foo bar")
                            switch (fi14 / ((32768) | 1)) {
                            case -1:
                                v1 = v1
                            case 9223372036854775807:
                                var v19 uint64 = v2
                                _ = v19
                            case 0:
                                var v20 int32 = -2147483648
                                _ = v20
                            default:
                                var v21 int32 = -2147483648
                                _ = v21
                            }
                        } else {
                            var v22 int32 = -59
                            var v23 uint8 = v3
                            var v24 int = (fi14 << (((v5 - (v5 / ((1) | 1)))) & 63))
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    default:
                        v5 += v5
                    }
                    _ = len(v4)
                    switch ("a" + "hello") {
                    case "hello":
                        if (1.5 <= 1.5) {
                            v2 = v2
                            var v25 string = ("" + "")
                            _ = v25
                        }
                    case "\n\t":
                        v4 = append(v4, ("" + ""))
                    case "":
                        switch fi15 {
                        case 862:
                            v13 += (^(v13))
                        case 443:
                            v5 += v5
                        case -2147483648:
                            switch fi15 {
                            case 32768:
                                v1 = v1
                            case -670:
                                v3 = v3
                            default:
                                switch -160 {
                                case 1:
                                    var v26 uint16 = 255
                                    _ = v26
                                case 4294967295:
                                    if true {
                                        fmt.Println(fi14, v13, v3)
                                        v13 -= v13
                                        v1 -= uint8(v5)
                                    } else {
                                        if (247 < -903) {
                                            var v27 int8 = -1
                                            _ = v27
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = ("a" + (("foo bar" + (("hello" + "\n\t") + ("hello" + "\n\t"))) + "hello"))
                                            }
                                            v5 = v5
                                            v3 = 42
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = ("a" + ("\n\t" + "foo bar"))
                                                }
                                                v3 += 70
                                            }
                                            v2 += (v2 & v2)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            fmt.Println(v13, v13)
                                        }
                                    }
                                case -1:
                                    v1 = (v1 | 0)
                                default:
                                    v4 = append(v4, "hello")
                                }
                            }
                        default:
                            switch "a" {
                            case "hello":
                                v3 += 127
                            default:
                                switch "foo bar" {
                                case "a":
                                    var v28 int8 = 1
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    if ((float32(0.0) <= 3.14159) && false) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (!true) {
                                                _ = len(v4)
                                                v29 := GenS2{F0: 0, F1: (v3 % ((v3) | 1)), F2: 0}
                                                _ = v29
                                            } else {
                                                if (!true) {
                                                    var v30 int8 = (v28 % (((v28 / (((v28 | (v28 * v28))) | 1))) | 1))
                                                    v31 := GenS3{F0: 0}
                                                    var v32 float32 = (3.14159 / 1.5)
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                }
                                            }
                                        case "hello":
                                            var v33 uint32 = 2
                                            _ = v33
                                        case "\n\t":
                                            v3 += 128
                                        default:
                                            var v34 int = fi15
                                            _ = v34
                                        }
                                        v2 += (v2 % (((v2 - 273)) | 1))
                                        switch (("" + "\n\t") + "foo bar") {
                                        case "\n\t":
                                            v35 := GenS4{F0: 160}
                                            v36 := GenS5{F0: 871}
                                            var v37 int16 = 42
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                    _ = v28
                                }
                            }
                            var v38 string = ""
                            _ = v38
                        }
                    default:
                        fmt.Println(v5, fi14)
                        v2 = 791
                    }
                }
            case 1:
                if true {
                    v5 = (v5 ^ 18446744073709551615)
                    if false {
                        v39 := GenS6{F0: true, F1: -1}
                        var v40 uint = (v5 << ((uint(v2)) & 63))
                        _ = v39
                        _ = v40
                    } else {
                        v2 += uint64(v3)
                        v41 := GenS4{F0: 105}
                        v42 := GenS7{F0: 61, F1: (v5 * (v5 & (v5 >> (((v5 - v5)) & 63)))), F2: 121}
                        _ = v41
                        _ = v42
                    }
                    var v43 float64 = (1.5 - float64(-2.25))
                    _ = v43
                } else {
                    v4 = append(v4, "foo bar")
                    v3 -= uint8(v5)
                    var v44 uint32 = 65535
                    _ = v44
                }
            case 41:
                v3 -= uint8(v13)
            default:
                var v45 uint16 = 1
                _ = v45
            }
            _ = fi15
        }
        fmt.Println(v3, v13, fi14)
        _ = fi14
    }
    v1 += v1
    var v46 string = "hello"
    if false {
        var v47 int = 488
        _ = v47
    } else {
        v48 := GenS3{F0: 678}
        for fi49 := 0; fi49 < 10; fi49++ {
            switch fi49 {
            case -9223372036854775808:
                switch (v46 + v46) {
                case "\n\t":
                    v4 = append(v4, v46)
                case "hello":
                    for fi50 := 0; fi50 < 18; fi50++ {
                        switch 58 {
                        case 32768:
                            var v51 uint8 = 128
                            _ = v51
                        case 733:
                            v52 := GenS8{F0: 1, F1: v46, F2: 42}
                            _ = v52
                        case -2147483648:
                            if len(v4) > 0 {
                                v4[0] = ((v46 + "foo bar") + v46)
                            }
                        }
                        _ = fi50
                    }
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    v53 := GenS9{F0: 0, F1: (v13 * 2147483648), F2: (false && (!(3.14159 >= (float64(1.5) + 1e10))))}
                    _ = v53
                }
            }
            var v54 int = fi49
            switch 4294967295 {
            case -805:
                var v55 uint64 = v2
                _ = v55
            default:
                var v56 int16 = 42
                _ = v56
            }
            _ = fi49
            _ = v54
        }
        var v57 int64 = int64(v13)
        _ = v48
        _ = v57
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        _ = len(v4)
        _ = fi58
    }
    var v59 uint32 = (v13 * 4294967295)
    v4 = append(v4, (v46 + "hello"))
    switch v46 {
    case "":
        var v60 int16 = 32767
        _ = v60
    case "hello":
        v61 := GenS2{F0: -826, F1: 217, F2: 261}
        _ = v61
    case "a":
        if ((!(255 >= 32767)) && (true || false)) {
            if true {
                v5 -= (^(v5))
                if true {
                    v3 -= (v3 << (((v5 << (((^(v5))) & 63))) & 63))
                }
                if (false || false) {
                    v2 = 4294967295
                    v46 = ("\n\t" + (v46 + ((("a" + v46) + (v46 + v46)) + v46)))
                }
            }
            v62 := GenS10{F0: 9223372036854775807, F1: 0, F2: 1}
            v63 := GenS11{F0: -553}
            _ = v62
            _ = v63
        }
    }
    if true {
        fmt.Println(v3, v1, v13)
        v59 += 65535
    } else {
        v2 = v2
        v64 := GenS9{F0: 805, F1: v59, F2: false}
        _ = len(v4)
        _ = v64
    }
    if true {
        if true {
            v46 = "foo bar"
        }
    }
    for fi65 := 0; fi65 < 16; fi65++ {
        v66 := GenS12{F0: 32767, F1: 9223372036854775807}
        v4 = append(v4, (v46 + v46))
        _ = fi65
        _ = v66
    }
    var v67 int16 = 1000
    if (!(("a" == v46) || (((v46 + v46) + (("" + v46) + "hello")) == ("foo bar" + v46)))) {
        if true {
            v2 = (^(v2))
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 += (v2 >> ((480) & 63))
    }
    if (4294967295 < -2147483648) {
        fmt.Println(v13)
    } else {
        switch "" {
        case "a":
            if (((3.14159 + 1e10) / (3.14159 - 1.5)) < (1e10 / 1.0)) {
                v68 := GenS13{F0: -1}
                _ = v68
            } else {
                for fi69 := 0; fi69 < 17; fi69++ {
                    var v70 uint = (v5 * v5)
                    var v71 uint8 = v3
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
        case "":
            v4 = append(v4, v46)
        case "foo bar":
            v67 = v67
        default:
            switch "\n\t" {
            case "hello":
                v59 += v59
            default:
                v1 -= v3
            }
        }
        v4 = append(v4, v46)
        v72 := GenS14{F0: 1, F1: "hello", F2: (v59 | (v13 / ((447) | 1)))}
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v59
    _ = v67
}
