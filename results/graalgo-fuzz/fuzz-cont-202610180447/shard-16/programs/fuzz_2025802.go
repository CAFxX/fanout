package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint; F1 int; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }

func main() {
    var v1 float32 = 0.0
    var v2 int64 = -9223372036854775808
    v3 := []uint32{4294967295, 44, 840, 0}
    v4 := []string{}
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "foo bar" {
        case "a":
            if (!((^(fi6)) >= (fi6 ^ (^(fi6))))) {
                if (-128 <= -128) {
                    v5 += 42
                } else {
                    for fi7 := 0; fi7 < 18; fi7++ {
                        v5 += v5
                        _ = fi7
                    }
                }
            } else {
                fmt.Println(fi6, v2)
            }
        }
        if false {
            v8 := GenS1{F0: "", F1: 255, F2: 1}
            v9 := GenS1{F0: "hello", F1: 985, F2: 0}
            v10 := GenS1{F0: "foo bar", F1: 42, F2: 18446744073709551615}
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if (1 > 255) {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 float64 = (float64(((v1 * v1) + 1.5)) * 1e10)
            var v13 float64 = (v12 / (v12 + 3.14159))
            fmt.Println(v1)
            _ = fi11
            _ = v12
            _ = v13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((true || (v2 == 9223372036854775807)) && (22 < 42)) {
            v1 = float32(0.0)
            v14 := GenS1{F0: "", F1: 1, F2: 1}
            _ = v14
        }
    }
    var v15 bool = (true || true)
    var v16 bool = v15
    v2 += v2
    for fi17 := 0; fi17 < 20; fi17++ {
        _ = len(v3)
        _ = fi17
    }
    if (16 < 100) {
        var v18 uint8 = 1
        _ = v18
    } else {
        switch 82 {
        case -943:
            switch ("foo bar" + ("a" + "\n\t")) {
            case "hello":
                v3 = append(v3, 226)
            case "foo bar":
                var v19 uint = 892
                _ = v19
            case "a":
                v20 := GenS2{F0: 42}
                _ = v20
            }
        case 1:
            v16 = false
        case 214:
            switch 42 {
            case 42:
                for fi21 := 0; fi21 < 18; fi21++ {
                    v16 = (31 >= -128)
                    _ = fi21
                }
            case 1:
                fmt.Println(v16)
            case -829:
                for fi22 := 0; fi22 < 3; fi22++ {
                    if true {
                        var v23 int16 = -32768
                        _ = v23
                    }
                    fmt.Println(v2, v5)
                    _ = fi22
                }
            }
        }
        if (v15 && v15) {
            v24 := GenS3{F0: 120, F1: (v5 / ((0) | 1)), F2: 236}
            fmt.Println(v15, v5, v2)
            _ = v24
        }
    }
    if v16 {
        v25 := GenS4{F0: "a", F1: 2, F2: 2}
        var v26 float64 = 0.0
        _ = v25
        _ = v26
    }
    v15 = v15
    fmt.Println(v1, v5)
    v1 += v1
    if (32767 > -531) {
        v27 := GenS3{F0: 0, F1: -394, F2: 154}
        switch 0 {
        case -904:
            switch -852 {
            case 32768:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -917:
                if (true && v16) {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v29 := GenS3{F0: 127, F1: (v5 / ((v5) | 1)), F2: 1}
                        switch "" {
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        case "":
                            v4 = append(v4, "hello")
                        case "hello":
                            if (v5 == -1) {
                                var v30 int16 = 1000
                                if v15 {
                                    if ((760 != 4294967295) || true) {
                                        v30 -= 32767
                                        switch "a" {
                                        case "":
                                            for fi31 := 0; fi31 < 3; fi31++ {
                                                if v16 {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = len(v3)
                                                    fmt.Println(v1, v15, v1)
                                                } else {
                                                    var v32 int16 = (v30 * v30)
                                                    _ = v32
                                                }
                                                _ = fi31
                                            }
                                        case "hello":
                                            var v33 int = (^(fi28))
                                            _ = v33
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (42 <= 18446744073709551615) {
                                                    var v34 uint64 = 0
                                                    v16 = (!(803 < 2))
                                                    _ = v34
                                                } else {
                                                    var v35 uint64 = 18446744073709551615
                                                    v15 = ((0 != 65535) || ((v16 || v15) || false))
                                                    if v16 {
                                                        v36 := GenS1{F0: "a", F1: 255, F2: v35}
                                                        switch "" {
                                                        case "\n\t":
                                                            fmt.Println(v1, v2, v15)
                                                        case "foo bar":
                                                            for fi37 := 0; fi37 < 8; fi37++ {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch (("" + ((("foo bar" + "\n\t") + "a") + "")) + "a") {
                                                                    case "foo bar":
                                                                        fmt.Println(v30, v15, v1)
                                                                    case "hello":
                                                                        v38 := GenS1{F0: "\n\t", F1: 2, F2: v35}
                                                                        _ = v38
                                                                    case "a":
                                                                        v3 = append(v3, 4294967295)
                                                                    }
                                                                case "\n\t":
                                                                    v2 -= (^(v2))
                                                                case "":
                                                                    v5 = v5
                                                                default:
                                                                    v39 := GenS2{F0: 2}
                                                                    _ = v39
                                                                }
                                                                _ = fi37
                                                            }
                                                        default:
                                                            v2 -= v2
                                                        }
                                                        for fi40 := 0; fi40 < 12; fi40++ {
                                                            var v41 uint16 = 0
                                                            var v42 bool = v16
                                                            _ = fi40
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = v36
                                                    } else {
                                                        fmt.Println(v5, v15)
                                                    }
                                                    _ = v35
                                                }
                                            }
                                        default:
                                            fmt.Println(v2, v30)
                                        }
                                        if v16 {
                                            for fi43 := 0; fi43 < 17; fi43++ {
                                                if (!v15) {
                                                    var v44 int64 = (v2 << ((407) & 63))
                                                    _ = v44
                                                } else {
                                                    var v45 int = fi43
                                                    var v46 float64 = float64((((v1 + ((((((v1 + float32(0.0)) - v1) - v1) + v1) * (v1 / (((float32(-2.25) * 1.5) / v1) * (v1 / v1)))) - (v1 - v1))) - v1) / (-2.25 * v1)))
                                                    v4 = append(v4, ("" + "foo bar"))
                                                    fmt.Println(v2, v5)
                                                    var v47 int = v45
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                                v30 = (v30 + 1000)
                                                v48 := GenS3{F0: int8(v5), F1: (v5 & v5), F2: 46}
                                                if (0 != (fi43 & 9223372036854775807)) {
                                                    if false {
                                                        _ = len(v4)
                                                    } else {
                                                        v49 := GenS2{F0: 58}
                                                        fmt.Println(v2, v2, v16)
                                                        switch (fi28 % ((fi28) | 1)) {
                                                        case 38:
                                                            v15 = v15
                                                        default:
                                                            v15 = (false || (true && v16))
                                                        }
                                                        _ = v49
                                                    }
                                                    v2 -= v2
                                                    v50 := GenS2{F0: 42}
                                                    _ = v50
                                                } else {
                                                    v4 = append(v4, ("\n\t" + "foo bar"))
                                                    switch (fi43 >> ((0) & 63)) {
                                                    case 80:
                                                        v51 := GenS2{F0: 65535}
                                                        _ = v51
                                                    case 9223372036854775807:
                                                        var v52 uint64 = 1
                                                        _ = v52
                                                    default:
                                                        v15 = (v15 || v15)
                                                    }
                                                    fmt.Println(v1)
                                                }
                                                _ = fi43
                                                _ = v48
                                            }
                                            var v53 uint64 = 4294967295
                                            v2 = (v2 & v2)
                                            _ = v53
                                        }
                                    } else {
                                        v54 := GenS5{F0: 602, F1: (fi28 - (-(fi28))), F2: 239}
                                        if true {
                                            v5 = v5
                                        } else {
                                            v55 := GenS4{F0: "\n\t", F1: 4294967295, F2: 2}
                                            _ = v55
                                        }
                                        fmt.Println(v30, v5, v2)
                                        _ = v54
                                    }
                                    v56 := GenS6{F0: 1}
                                    _ = v56
                                } else {
                                    switch fi28 {
                                    case -9223372036854775808:
                                        if len(v4) > 0 {
                                            v4[0] = "\n\t"
                                        }
                                    case 469:
                                        switch (fi28 | (fi28 - (fi28 << ((448) & 63)))) {
                                        case 429:
                                            var v57 float64 = (1e10 - 1e10)
                                            _ = v57
                                        default:
                                            v16 = (v16 || v15)
                                        }
                                    case 47:
                                        var v58 int32 = (^(v5))
                                        _ = v58
                                    }
                                    var v59 float64 = ((((3.14159 + float64((float32(-2.25) - 1.5))) - (1e10 * 0.1)) * ((-2.25 * (1.5 / 1.0)) - 3.14159)) / 1.5)
                                    _ = v59
                                }
                                _ = v30
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                    switch 128 {
                    case 62:
                        fmt.Println(v2)
                    case 2147483647:
                        v2 -= (^(v2))
                    default:
                        switch "foo bar" {
                        case "hello":
                            v2 += 457
                        case "foo bar":
                            for fi60 := 0; fi60 < 12; fi60++ {
                                var v61 float64 = 1e10
                                _ = fi60
                                _ = v61
                            }
                        default:
                            fmt.Println(v2, v15, v5)
                        }
                    }
                    if ((v2 << ((265) & 63)) <= (v2 / (((v2 - v2)) | 1))) {
                        v62 := GenS7{F0: "\n\t"}
                        var v63 bool = (-527 < 1000)
                        _ = v62
                        _ = v63
                    }
                }
            case -436:
                v64 := GenS6{F0: 357}
                _ = v64
            default:
                var v65 float32 = v1
                _ = v65
            }
        }
        _ = v27
    }
    for fi66 := 0; fi66 < 7; fi66++ {
        v15 = v16
        fmt.Println(fi66, v2)
        fmt.Println(fi66, v15, v2)
        _ = fi66
    }
    v1 = float32((1.5 + (0.1 - ((3.14159 - ((1e10 / 1.5) * 0.1)) + -2.25))))
    var v67 uint8 = 255
    if (int16(v2) >= 42) {
        if v15 {
            v16 = ((true || v15) && (v15 && v16))
        }
        v4 = append(v4, "\n\t")
        var v68 int32 = (v5 / (((v5 << ((602) & 63))) | 1))
        _ = v68
    } else {
        var v69 float32 = (v1 / ((v1 - v1) - v1))
        _ = v69
    }
    for fi70 := 0; fi70 < 3; fi70++ {
        fmt.Println(v5)
        v1 = v1
        _ = fi70
    }
    var v71 int32 = (v5 << ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v71
}
