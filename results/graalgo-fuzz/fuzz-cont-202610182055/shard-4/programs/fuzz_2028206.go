package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 string; F1 bool; F2 string }
type GenS6 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 uint64; F1 int; F2 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32 }

func main() {
    var v1 uint64 = 931
    var v2 uint64 = v1
    v3 := make(map[string]string)
    v4 := []int{42, 0, 127, 128}
    v2 = (^(v1))
    fmt.Println(v1)
    if (990 != 626) {
        v2 = (v1 + (v2 % ((v2) | 1)))
    } else {
        if true {
            for fi5 := 0; fi5 < 5; fi5++ {
                fmt.Println(v2)
                var v6 float64 = -2.25
                v1 += 97
                _ = fi5
                _ = v6
            }
        }
    }
    var v7 float64 = float64(-2.25)
    v8 := GenS1{F0: 128}
    if true {
        var v9 uint = 257
        if true {
            v10 := GenS2{F0: false, F1: -9223372036854775808, F2: -9223372036854775808}
            v2 = v1
            _ = v10
        }
        v11 := GenS3{F0: 127}
        _ = v11
        _ = v9
    }
    if true {
        if (9223372036854775807 == 9223372036854775807) {
            var v12 uint64 = v1
            if len(v4) > 0 {
                v4[0] = 42
            }
            var v13 bool = (false || false)
            _ = v12
            _ = v13
        } else {
            if (true && false) {
                var v14 uint32 = 939
                switch 42 {
                case 208:
                    var v15 uint64 = v2
                    _ = v15
                case 0:
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 int64 = -1
                        if false {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                fmt.Println(fi18, v7)
                                _ = fi18
                            }
                        } else {
                            v19 := v3["a"]
                            v20 := GenS4{F0: uint32(v2), F1: (false && (v19 < (((v19 + v19) + v19) + "hello")))}
                            v21 := GenS4{F0: v14, F1: (1 < -2147483648)}
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        _ = fi16
                        _ = v17
                    }
                case -2147483648:
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 float64 = (v7 * (float64(((0.1 * (((3.14159 - 3.14159) / 3.14159) / (0.1 + float32(v7)))) - (3.14159 * float32(v7)))) - (v7 / v7)))
                        switch -559 {
                        case -183:
                            for fi24 := 0; fi24 < 13; fi24++ {
                                v25 := GenS4{F0: v14, F1: false}
                                if false {
                                    v7 += (v23 + (float64(3.14159) + 0.0))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v26 := GenS3{F0: 2}
                                    _ = v26
                                } else {
                                    fmt.Println(v7)
                                    switch fi24 {
                                    case 1:
                                        v27 := GenS5{F0: "", F1: (65535 <= 702), F2: ("" + ("foo bar" + ""))}
                                        _ = v27
                                    case -776:
                                        v28 := GenS2{F0: (v1 < uint64(v14)), F1: 9223372036854775807, F2: -9223372036854775808}
                                        _ = v28
                                    }
                                    v23 -= v7
                                }
                                _ = fi24
                                _ = v25
                            }
                        case 32767:
                            switch fi22 {
                            case 32768:
                                v29 := GenS6{F0: 65535, F1: 108, F2: 1}
                                _ = v29
                            case -1:
                                if true {
                                    if (false && (false && (!(1000 == 1000)))) {
                                        for fi30 := 0; fi30 < 2; fi30++ {
                                            switch "a" {
                                            case "hello":
                                                var v31 int8 = 127
                                                _ = v31
                                            case "foo bar":
                                                var v32 uint = 89
                                                _ = v32
                                            }
                                            v33 := GenS7{F0: 0, F1: 1}
                                            _ = fi30
                                            _ = v33
                                        }
                                        var v34 uint8 = 173
                                        _ = v34
                                    } else {
                                        for fi35 := 0; fi35 < 13; fi35++ {
                                            v3["hello"] = ""
                                            v36 := GenS8{F0: v2, F1: v2}
                                            _ = fi35
                                            _ = v36
                                        }
                                    }
                                } else {
                                    fmt.Println(v2)
                                    v37 := GenS3{F0: -128}
                                    _ = v37
                                }
                            case 32767:
                                if (float64(1.5) != v7) {
                                    if len(v4) > 0 {
                                        v4[0] = (fi22 + fi22)
                                    }
                                } else {
                                    var v38 int8 = -1
                                    v4 = append(v4, 0)
                                    _ = v38
                                }
                            default:
                                var v39 uint8 = 255
                                _ = v39
                            }
                        case 2147483647:
                            switch "" {
                            case "":
                                v40 := GenS9{F0: (v2 << ((2) & 63)), F1: (fi22 * fi22), F2: -1}
                                _ = v40
                            case "foo bar":
                                v14 -= 2147483648
                            case "hello":
                                v23 += float64(((-2.25 / 0.1) + (((1e10 + float32(((v7 / (1.5 + -2.25)) - v7))) * 1e10) + -2.25)))
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                default:
                    v41 := GenS7{F0: 0, F1: 127}
                    _ = v41
                }
                v1 += (^(v2))
                _ = v14
            }
            var v42 uint = 440
            v43 := GenS10{F0: 9223372036854775807}
            _ = v42
            _ = v43
        }
        if true {
            v1 += v2
            if (2 != 0) {
                var v44 uint64 = 42
                if true {
                    if ((18446744073709551615 <= 42) || false) {
                        var v45 uint8 = 243
                        _ = v45
                    } else {
                        if (-1 <= -66) {
                            fmt.Println(v7, v2, v7)
                            v2 += 4294967295
                            fmt.Println(v2, v1, v7)
                        }
                        var v46 uint16 = 0
                        _ = v46
                    }
                    switch "a" {
                    case "a":
                        v47 := GenS2{F0: (3.14159 != 3.14159), F1: -354, F2: -9223372036854775808}
                        _ = v47
                    case "hello":
                        var v48 float64 = v7
                        _ = v48
                    case "\n\t":
                        for fi49 := 0; fi49 < 1; fi49++ {
                            for fi50 := 0; fi50 < 14; fi50++ {
                                v51 := GenS2{F0: (127 >= 127), F1: 42, F2: 1}
                                var v52 int32 = 555
                                switch "foo bar" {
                                case "\n\t":
                                    var v53 int8 = 42
                                    _ = v53
                                case "hello":
                                    v54 := GenS11{F0: 42}
                                    _ = v54
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v4 = append(v4, -802)
                            v55 := GenS7{F0: int8(v2), F1: 0}
                            _ = fi49
                            _ = v55
                        }
                    default:
                        fmt.Println(v2, v1)
                    }
                } else {
                    switch 32767 {
                    case 127:
                        v44 += v1
                    case 288:
                        for fi56 := 0; fi56 < 5; fi56++ {
                            if true {
                                for fi57 := 0; fi57 < 13; fi57++ {
                                    var v58 uint8 = 222
                                    _ = fi57
                                    _ = v58
                                }
                            } else {
                                v3["a"] = "foo bar"
                                fmt.Println(v1)
                            }
                            _ = fi56
                        }
                    }
                }
                _ = v44
            }
            var v59 int32 = -2147483648
            _ = v59
        }
    } else {
        v4 = append(v4, 4294967295)
        for fi60 := 0; fi60 < 12; fi60++ {
            var v61 bool = true
            _ = fi60
            _ = v61
        }
        var v62 int32 = 414
        _ = v62
    }
    v2 -= v2
    for fi63 := 0; fi63 < 6; fi63++ {
        switch (fi63 ^ fi63) {
        case 127:
            fmt.Println(v1)
        case -245:
            if ("\n\t" < ("a" + (("foo bar" + "a") + "foo bar"))) {
                switch (fi63 / ((fi63) | 1)) {
                case -75:
                    switch (fi63 - fi63) {
                    case 2147483647:
                        for fi64 := 0; fi64 < 8; fi64++ {
                            switch fi63 {
                            case 32768:
                                v65 := GenS5{F0: "hello", F1: false, F2: "\n\t"}
                                _ = v65
                            case 2147483647:
                                v66 := GenS8{F0: v1, F1: v1}
                                _ = v66
                            case 9223372036854775807:
                                v67 := GenS2{F0: (18446744073709551615 <= 0), F1: -463, F2: -1}
                                _ = v67
                            default:
                                switch "a" {
                                case "a":
                                    var v68 int = 4294967295
                                    _ = v68
                                default:
                                    var v69 uint8 = 246
                                    _ = v69
                                }
                            }
                            switch 4294967295 {
                            case 781:
                                if true {
                                    v70 := v3["a"]
                                    var v71 float32 = 0.1
                                    _ = v70
                                    _ = v71
                                }
                            case 32767:
                                _ = len(v4)
                            case -1:
                                var v72 float64 = (v7 / v7)
                                _ = v72
                            }
                            _ = fi64
                        }
                    default:
                        var v73 float32 = (-2.25 - -2.25)
                        _ = v73
                    }
                case -9223372036854775808:
                    v74 := GenS1{F0: fi63}
                    _ = v74
                case 128:
                    v2 += v1
                default:
                    fmt.Println(v2, fi63, v1)
                }
            }
        case 32768:
            var v75 uint16 = 1
            _ = v75
        default:
            v76 := GenS12{F0: 255}
            _ = v76
        }
        for fi77 := 0; fi77 < 11; fi77++ {
            v78 := GenS13{F0: 42}
            _ = fi77
            _ = v78
        }
        v2 = v1
        _ = fi63
    }
    if (!false) {
        switch ((("hello" + "hello") + "") + "") {
        case "":
            v2 += (v1 / ((v2) | 1))
        }
        v4 = append(v4, 127)
        v1 -= 879
    } else {
        v7 = float64(-2.25)
        switch 128 {
        case 832:
            _ = len(v3)
        default:
            if false {
                delete(v3, "hello")
                for fi79 := 0; fi79 < 9; fi79++ {
                    v7 += -2.25
                    _ = fi79
                }
            }
        }
    }
    if (168 >= 128) {
        fmt.Println(v1)
        if true {
            v7 = (v7 * float64(3.14159))
            var v80 float32 = (((((0.1 - 1.5) * 1.5) / 1.0) / 1e10) / (0.0 + (1e10 / 0.1)))
            v1 += (v1 + v2)
            _ = v80
        } else {
            v1 += 24
            v7 = v7
        }
    } else {
        for fi81 := 0; fi81 < 10; fi81++ {
            for fi82 := 0; fi82 < 18; fi82++ {
                v2 = v1
                _ = fi82
            }
            v4 = append(v4, (fi81 & (fi81 | (fi81 % (((fi81 + fi81)) | 1)))))
            v2 += (v1 << ((4294967295) & 63))
            _ = fi81
        }
        for fi83 := 0; fi83 < 12; fi83++ {
            v2 -= v1
            switch (("\n\t" + "foo bar") + "a") {
            case "hello":
                switch fi83 {
                case 128:
                    for fi84 := 0; fi84 < 17; fi84++ {
                        v4 = append(v4, (fi84 << ((2) & 63)))
                        v7 = 0.0
                        var v85 uint = 953
                        _ = fi84
                        _ = v85
                    }
                default:
                    v7 = float64(float32(v7))
                }
            }
            _ = len(v4)
            _ = fi83
        }
        v7 = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
