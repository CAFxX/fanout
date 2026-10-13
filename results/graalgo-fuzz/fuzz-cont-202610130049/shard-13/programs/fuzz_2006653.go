package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint16; F2 string }
type GenS7 struct { F0 int16 }

func main() {
    var v1 int = -2147483648
    v2 := []string{(("hello" + "a") + "foo bar")}
    var v3 uint64 = 4294967295
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v4 := GenS1{F0: 0, F1: 0}
    v1 -= 9223372036854775807
    v3 = v3
    if (!false) {
        var v5 int32 = 2147483647
        if false {
            fmt.Println(v5, v5)
            v1 += (-(v1))
        } else {
            fmt.Println(v5, v5)
            switch 42 {
            case 4294967295:
                fmt.Println(v3, v1, v1)
            case -742:
                if true {
                    v5 -= 42
                }
            case -1:
                fmt.Println(v5)
            default:
                if false {
                    for fi6 := 0; fi6 < 13; fi6++ {
                        for fi7 := 0; fi7 < 20; fi7++ {
                            switch (-(fi7)) {
                            case 210:
                                if true {
                                    switch 32767 {
                                    case -1:
                                        switch "" {
                                        case "hello":
                                            switch -869 {
                                            case 0:
                                                fmt.Println(v5, v3)
                                                v2 = append(v2, "")
                                                if false {
                                                    _ = len(v2)
                                                } else {
                                                    v1 -= v1
                                                    v5 = 42
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v8 uint8 = 127
                                                    if true {
                                                        fmt.Println(v5)
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v3 = 2
                                                        case "a":
                                                            var v9 uint32 = 2147483648
                                                            var v10 float32 = (1.5 / -2.25)
                                                            _ = v10
                                                            _ = v9
                                                        }
                                                        fmt.Println(v1)
                                                        v11 := GenS1{F0: v8, F1: 0}
                                                        _ = v11
                                                    } else {
                                                        v12 := GenS2{F0: 2, F1: false, F2: 32767}
                                                        var v13 string = ""
                                                        switch "" {
                                                        case "a":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "\n\t":
                                                            if false {
                                                                v14 := GenS2{F0: 65535, F1: true, F2: 255}
                                                                _ = v14
                                                            } else {
                                                                switch v1 {
                                                                case -1:
                                                                    var v15 int8 = 93
                                                                    _ = v15
                                                                }
                                                            }
                                                        }
                                                        _ = v12
                                                        _ = v13
                                                    }
                                                    _ = v8
                                                }
                                            case -187:
                                                var v16 float64 = 3.14159
                                                var v17 uint8 = 0
                                                v18 := GenS3{F0: "\n\t", F1: 42, F2: fi6}
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                            case 282:
                                                v2 = append(v2, "hello")
                                            default:
                                                var v19 float32 = ((1e10 + 1.5) * 0.0)
                                                _ = v19
                                            }
                                        case "":
                                            if (632 != 787) {
                                                v5 = -865
                                                v3 -= (v3 % (((v3 & (v3 & (v3 << ((4294967295) & 63))))) | 1))
                                            }
                                        case "foo bar":
                                            if (false || true) {
                                                var v20 uint64 = (v3 - v3)
                                                if true {
                                                    var v21 uint16 = 25
                                                    if (-1 >= -9223372036854775808) {
                                                        v2 = append(v2, ("" + ""))
                                                        var v22 int16 = 1
                                                        _ = v22
                                                    }
                                                    fmt.Println(v21, fi7, fi6)
                                                    _ = v21
                                                } else {
                                                    v1 += v1
                                                    switch (v1 - fi7) {
                                                    case 4294967295:
                                                        v23 := GenS3{F0: "hello", F1: 42, F2: 32768}
                                                        _ = v23
                                                    case 1:
                                                        v24 := GenS4{F0: 255, F1: int8(fi6)}
                                                        _ = v24
                                                    case -155:
                                                        v1 += 127
                                                    }
                                                }
                                                _ = v20
                                            } else {
                                                _ = len(v2)
                                            }
                                        default:
                                            v2 = append(v2, ("\n\t" + ("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "a"))))))
                                        }
                                    case 128:
                                        var v25 int = (fi6 >> ((2) & 63))
                                        _ = v25
                                    case -2147483648:
                                        var v26 uint32 = 42
                                        _ = v26
                                    default:
                                        switch "a" {
                                        case "a":
                                            v3 = v3
                                        case "hello":
                                            fmt.Println(v5)
                                        }
                                    }
                                    v3 = v3
                                }
                            case -2147483648:
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            case -9223372036854775808:
                                v1 = (-(fi7))
                            }
                            _ = fi7
                        }
                        if true {
                            v27 := GenS5{F0: "\n\t"}
                            if ((v1 ^ -2147483648) > v1) {
                                var v28 float64 = 1.5
                                v28 = -2.25
                                _ = v28
                            }
                            if ((v5 + v5) <= (v5 * v5)) {
                                switch "" {
                                case "hello":
                                    v5 += (v5 / ((v5) | 1))
                                case "":
                                    var v29 uint32 = 916
                                    _ = v29
                                default:
                                    if (v1 >= fi6) {
                                        for fi30 := 0; fi30 < 13; fi30++ {
                                            switch "" {
                                            case "\n\t":
                                                v31 := GenS6{F0: 2147483648, F1: 255, F2: "a"}
                                                _ = v31
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                var v32 string = "hello"
                                                v1 = (fi6 + (fi6 + (fi30 * (v1 / (((-(v1))) | 1)))))
                                                _ = v32
                                            }
                                            v1 += fi6
                                            var v33 int16 = 32767
                                            _ = fi30
                                            _ = v33
                                        }
                                    } else {
                                        for fi34 := 0; fi34 < 14; fi34++ {
                                            var v35 uint = 349
                                            var v36 uint32 = 2
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                        }
                                        v1 = fi6
                                        var v37 uint64 = 4294967295
                                        _ = v37
                                    }
                                }
                                var v38 string = ((("a" + "hello") + "hello") + "")
                                _ = v38
                            } else {
                                var v39 int16 = -1
                                v2 = append(v2, ((((("\n\t" + ("foo bar" + (("\n\t" + ("a" + "hello")) + ("hello" + (("" + ("hello" + "\n\t")) + "a"))))) + "hello") + "\n\t") + "\n\t") + "hello"))
                                _ = v39
                            }
                            _ = v27
                        } else {
                            if false {
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    if false {
                                        var v41 uint16 = 65535
                                        var v42 uint8 = uint8(v3)
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        var v43 bool = (!(171 <= 932))
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                        _ = v43
                                    }
                                    v44 := GenS4{F0: 85, F1: 100}
                                    v45 := GenS2{F0: 540, F1: (("a" + "hello") == "hello"), F2: 65535}
                                    _ = fi40
                                    _ = v44
                                    _ = v45
                                }
                                if (true || true) {
                                    if (!true) {
                                        v5 = v5
                                    } else {
                                        v46 := GenS3{F0: "hello", F1: -1, F2: v1}
                                        v3 = 18446744073709551615
                                        if false {
                                            if false {
                                                if false {
                                                    switch v1 {
                                                    case 4294967295:
                                                        v1 = -1
                                                    case -9223372036854775808:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v3 += (v3 + v3)
                                                } else {
                                                    if true {
                                                        switch v1 {
                                                        case -705:
                                                            var v47 int32 = (-(v5))
                                                            _ = v47
                                                        case -1:
                                                            var v48 uint16 = 0
                                                            _ = v48
                                                        case 128:
                                                            fmt.Println(v1)
                                                        default:
                                                            for fi49 := 0; fi49 < 7; fi49++ {
                                                                var v50 int16 = -605
                                                                v51 := GenS1{F0: 255, F1: (v50 << ((155) & 63))}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = fi49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        }
                                                    }
                                                    v3 += v3
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            } else {
                                                v1 += fi6
                                                v52 := GenS4{F0: 128, F1: -88}
                                                var v53 uint32 = 65535
                                                _ = v52
                                                _ = v53
                                            }
                                        } else {
                                            if false {
                                                if (904 < 0) {
                                                    _ = len(v2)
                                                    var v54 uint = 454
                                                    _ = v54
                                                }
                                            }
                                        }
                                        _ = v46
                                    }
                                    var v55 uint8 = 127
                                    _ = v55
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            fmt.Println(v3, v5, v5)
                        }
                        if false {
                            for fi56 := 0; fi56 < 10; fi56++ {
                                var v57 float32 = -2.25
                                v2 = append(v2, "\n\t")
                                _ = fi56
                                _ = v57
                            }
                            v58 := GenS4{F0: 130, F1: -1}
                            for fi59 := 0; fi59 < 20; fi59++ {
                                v1 = fi6
                                _ = fi59
                            }
                            _ = v58
                        }
                        _ = fi6
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v5
    } else {
        var v60 int32 = 0
        var v61 float64 = 0.1
        var v62 int64 = -1
        _ = v60
        _ = v61
        _ = v62
    }
    var v63 bool = (!true)
    v3 = (v3 >> ((0) & 63))
    if len(v2) > 0 {
        v2[0] = ("hello" + "a")
    }
    switch v1 {
    case -1:
        v1 = (v1 + v1)
    case -359:
        var v64 uint8 = 42
        _ = v64
    case -163:
        var v65 uint = 4294967295
        _ = v65
    }
    switch (^(v1)) {
    case 9223372036854775807:
        switch (v1 - 474) {
        case 128:
            v66 := GenS5{F0: "a"}
            _ = v66
        case -9223372036854775808:
            fmt.Println(v1, v1, v3)
        case 32767:
            var v67 bool = v63
            _ = v67
        default:
            var v68 int8 = -128
            _ = v68
        }
    default:
        if false {
            if v63 {
                switch (v1 >> ((4294967295) & 63)) {
                case 0:
                    if v63 {
                        for fi69 := 0; fi69 < 14; fi69++ {
                            v63 = v63
                            if (110 < 2147483647) {
                                fmt.Println(v1)
                                switch -9223372036854775808 {
                                case 128:
                                    var v70 string = ""
                                    _ = v70
                                case 9223372036854775807:
                                    var v71 uint16 = 255
                                    _ = v71
                                default:
                                    var v72 uint32 = 199
                                    _ = v72
                                }
                                for fi73 := 0; fi73 < 18; fi73++ {
                                    var v74 float32 = float32((1e10 * ((0.1 + 1e10) * 1.5)))
                                    _ = fi73
                                    _ = v74
                                }
                            }
                            fmt.Println(v1)
                            _ = fi69
                        }
                        if (v63 && v63) {
                            fmt.Println(v1)
                            var v75 uint = 42
                            _ = v75
                        } else {
                            v3 = v3
                            v76 := GenS7{F0: 1}
                            _ = v76
                        }
                    } else {
                        for fi77 := 0; fi77 < 10; fi77++ {
                            if (100 > 1) {
                                var v78 int64 = -236
                                v79 := GenS6{F0: 2, F1: 166, F2: (("a" + "a") + "a")}
                                _ = v78
                                _ = v79
                            }
                            if (!(!false)) {
                                for fi80 := 0; fi80 < 17; fi80++ {
                                    var v81 int = (fi80 % ((fi77) | 1))
                                    v82 := GenS1{F0: 42, F1: 42}
                                    _ = fi80
                                    _ = v81
                                    _ = v82
                                }
                                var v83 float64 = 0.0
                                _ = v83
                            }
                            switch "\n\t" {
                            case "a":
                                for fi84 := 0; fi84 < 9; fi84++ {
                                    var v85 uint64 = (v3 * v3)
                                    fmt.Println(v3, v85, v85)
                                    _ = len(v2)
                                    _ = fi84
                                    _ = v85
                                }
                            case "":
                                v3 += 18446744073709551615
                            default:
                                v86 := GenS1{F0: 127, F1: 18}
                                _ = v86
                            }
                            _ = fi77
                        }
                    }
                case -9223372036854775808:
                    v3 = 42
                case 1:
                    fmt.Println(v3, v1)
                default:
                    fmt.Println(v3, v1, v1)
                }
                _ = len(v2)
                v3 += (v3 / ((42) | 1))
            } else {
                fmt.Println(v3, v1, v63)
                v1 = 32767
                var v87 int32 = -511
                _ = v87
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
