package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 int8, p1 string) bool {
    var v6 int64 = 9223372036854775807
    fmt.Println(v6, p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1() uint8 {
    var v7 float64 = 1.5
    var v8 uint64 = 4294967295
    switch ("a" + "foo bar") {
    case "hello":
        v9 := GenS1{F0: 285, F1: 65535}
        _ = v9
    }
    var v10 bool = (1e10 > -2.25)
    _ = v10
    _ = v7
    _ = v8
    return 238
}

func main() {
    v1 := []int{-468, 4294967295, 127}
    var v2 uint16 = 32767
    var v3 uint16 = v2
    var v4 float32 = 0.1
    var v5 uint32 = 2147483648
    v11 := GenS1{F0: 0, F1: 774}
    v1 = append(v1, 9223372036854775807)
    v3 = 2
    v4 = v4
    if (316 > 0) {
        var v12 int = 32767
        if ((true && ((3.14159 < 1.5) && (4294967295 < 18446744073709551615))) && (true || true)) {
            var v13 uint8 = 0
            v12 -= (v12 / (((-(v12))) | 1))
            _ = v13
        } else {
            v12 = v12
            v12 -= v12
            v14 := GenS1{F0: -32768, F1: 4294967295}
            _ = v14
        }
        var v15 uint64 = 93
        _ = v12
        _ = v15
    } else {
        var v16 int32 = int32(v3)
        if len(v1) > 0 {
            v1[0] = 1
        }
        fmt.Println(v4, v2, v3)
        _ = v16
    }
    if (2 == 18446744073709551615) {
        fmt.Println(v2)
        switch "hello" {
        case "foo bar":
            var v17 int32 = -2147483648
            _ = v17
        case "\n\t":
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 float32 = v4
                _ = fi18
                _ = v19
            }
        case "hello":
            switch -895 {
            case -9223372036854775808:
                switch (("hello" + "a") + ("" + (((("a" + ("hello" + "")) + (("foo bar" + ("\n\t" + (("foo bar" + ((((("foo bar" + "\n\t") + ("a" + "\n\t")) + "") + ("hello" + "foo bar")) + "")) + ((("\n\t" + "foo bar") + "hello") + ("hello" + "hello"))))) + "")) + "\n\t") + ("foo bar" + "hello")))) {
                case "a":
                    switch -9223372036854775808 {
                    case 128:
                        if false {
                            v4 -= 3.14159
                            var v20 int32 = int32(v2)
                            v21 := GenS2{F0: v20}
                            _ = v20
                            _ = v21
                        } else {
                            var v22 uint = 453
                            _ = v22
                        }
                    default:
                        v23 := GenS3{F0: 2147483647, F1: 32767}
                        _ = v23
                    }
                default:
                    if false {
                        v1 = append(v1, 4294967295)
                        if true {
                            var v24 int64 = 9223372036854775807
                            v2 = v3
                            var v25 float32 = float32(3.14159)
                            _ = v24
                            _ = v25
                        }
                    } else {
                        if true {
                            var v26 uint = 735
                            v27 := GenS3{F0: v5, F1: 32767}
                            _ = v26
                            _ = v27
                        } else {
                            v5 += (v5 * (v5 * (v5 - (v5 * (v5 * v5)))))
                            for fi28 := 0; fi28 < 11; fi28++ {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    var v30 bool = ((670 <= -2147483648) || false)
                                    v30 = false
                                    var v31 uint16 = (v3 / (((v3 >> ((2) & 63))) | 1))
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi28
                            }
                            switch 128 {
                            case 2147483647:
                                if false {
                                    v1 = append(v1, 32768)
                                    switch "hello" {
                                    case "\n\t":
                                        switch int(v3) {
                                        case 127:
                                            if true {
                                                for fi32 := 0; fi32 < 2; fi32++ {
                                                    v4 -= (3.14159 * float32((0.1 - 1e10)))
                                                    if true {
                                                        if false {
                                                            var v33 float32 = (float32((1.5 + 0.1)) - 1.5)
                                                            v4 = float32(0.1)
                                                            if (0.0 < 3.14159) {
                                                                v5 = (v5 ^ (v5 & v5))
                                                            } else {
                                                                v34 := GenS2{F0: -62}
                                                                _ = v34
                                                            }
                                                            _ = v33
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = (fi32 - (-(fi32)))
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            switch (fi32 % ((-9223372036854775808) | 1)) {
                                                            case -1:
                                                                v1 = append(v1, (fi32 - fi32))
                                                            case 2147483647:
                                                                var v35 float32 = v4
                                                                _ = v35
                                                            case 32767:
                                                                if false {
                                                                    v2 += (v2 & 444)
                                                                    fmt.Println(v4, v3, fi32)
                                                                }
                                                            }
                                                        }
                                                        if (0 > 4294967295) {
                                                            fmt.Println(v2, fi32)
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    _ = fi32
                                                }
                                                var v36 int32 = -1
                                                _ = v36
                                            } else {
                                                for fi37 := 0; fi37 < 6; fi37++ {
                                                    v3 = 42
                                                    for fi38 := 0; fi38 < 17; fi38++ {
                                                        fmt.Println(v3, v4, v4)
                                                        if true {
                                                            var v39 uint = 255
                                                            _ = v39
                                                        }
                                                        if (v5 == (v5 >> ((1) & 63))) {
                                                            switch "hello" {
                                                            case "":
                                                                if false {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                var v40 bool = ((genFn1() <= 128) && false)
                                                                v1 = append(v1, (fi38 / (((fi37 | fi38)) | 1)))
                                                                var v41 uint32 = v5
                                                                _ = v40
                                                                _ = v41
                                                            } else {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v42 := GenS3{F0: (v5 & (v5 - (v5 - v5))), F1: 42}
                                                                    _ = v42
                                                                case "hello":
                                                                    if true {
                                                                        var v43 string = ""
                                                                        var v44 float64 = ((1.5 / (float64(v4) / 0.1)) * float64(((v4 - -2.25) + v4)))
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    switch (fi38 / (((fi37 % (((fi38 + fi38)) | 1))) | 1)) {
                                                                    case -168:
                                                                        if (false && (3.14159 > (float32(float64(-2.25)) + ((((0.1 / v4) / float32(3.14159)) * v4) / 1.0)))) {
                                                                            var v45 int8 = 100
                                                                            var v46 int64 = -9223372036854775808
                                                                            v46 = v46
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            fmt.Println(v4, v4, v5)
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi37, v3, v3)
                                                                    }
                                                                case "foo bar":
                                                                    v4 = ((1e10 + (v4 * (0.1 - (-2.25 * v4)))) / v4)
                                                                default:
                                                                    var v47 uint32 = v5
                                                                    _ = v47
                                                                }
                                                                var v48 string = (((("foo bar" + "foo bar") + ("" + "hello")) + "a") + (("foo bar" + "foo bar") + "a"))
                                                                if (1 >= 123) {
                                                                    var v49 string = v48
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            }
                                                            var v50 float32 = 1.5
                                                            _ = v50
                                                        }
                                                        _ = fi38
                                                    }
                                                    _ = fi37
                                                }
                                            }
                                        default:
                                            for fi51 := 0; fi51 < 12; fi51++ {
                                                var v52 int64 = -407
                                                _ = fi51
                                                _ = v52
                                            }
                                        }
                                    case "":
                                        for fi53 := 0; fi53 < 13; fi53++ {
                                            for fi54 := 0; fi54 < 16; fi54++ {
                                                v3 += v2
                                                _ = fi54
                                            }
                                            _ = fi53
                                        }
                                    }
                                }
                            case -2147483648:
                                var v55 int8 = 1
                                _ = v55
                            case 128:
                                v3 = 2
                            default:
                                fmt.Println(v3)
                            }
                        }
                        if false {
                            var v56 int16 = -288
                            var v57 int = -2147483648
                            _ = v56
                            _ = v57
                        } else {
                            var v58 float64 = 0.0
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch 127 {
                            case 32768:
                                _ = len(v1)
                            default:
                                for fi59 := 0; fi59 < 12; fi59++ {
                                    for fi60 := 0; fi60 < 19; fi60++ {
                                        v3 = 1
                                        var v61 uint16 = v2
                                        _ = fi60
                                        _ = v61
                                    }
                                    v5 = v5
                                    v4 = v4
                                    _ = fi59
                                }
                            }
                            _ = v58
                        }
                    }
                }
            case 851:
                switch 363 {
                case -9223372036854775808:
                    v3 = v2
                case 9223372036854775807:
                    if false {
                        var v62 string = ("hello" + "foo bar")
                        _ = v62
                    } else {
                        v5 += (v5 << ((1) & 63))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v3)
                    }
                case -767:
                    if len(v1) > 0 {
                        v1[0] = 32767
                    }
                default:
                    for fi63 := 0; fi63 < 10; fi63++ {
                        var v64 int = fi63
                        _ = fi63
                        _ = v64
                    }
                }
            }
        }
    }
    fmt.Println(v4, v4)
    fmt.Println(v5)
    switch 128 {
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -9223372036854775808:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -265:
        var v65 bool = (true && true)
        _ = v65
    }
    switch "\n\t" {
    case "":
        var v66 float32 = 1.5
        _ = v66
    case "hello":
        if (0.0 == ((1e10 - 0.0) * 0.1)) {
            fmt.Println(v2)
        } else {
            v1 = append(v1, 127)
            var v67 string = ("\n\t" + "a")
            if false {
                switch (v67 + "foo bar") {
                case "":
                    switch v67 {
                    case "":
                        v67 = v67
                    case "hello":
                        if (((v67 + v67) > v67) && ((-874 > -1) || false)) {
                            if (1 > -1) {
                                var v68 string = v67
                                v69 := GenS1{F0: 42, F1: 1}
                                _ = v68
                                _ = v69
                            } else {
                                switch "" {
                                case "hello":
                                    v2 = 255
                                case "a":
                                    v4 = (v4 - 1.5)
                                default:
                                    var v70 float32 = v4
                                    _ = v70
                                }
                            }
                            var v71 uint64 = 448
                            _ = v71
                        }
                    case "foo bar":
                        v2 += uint16(v5)
                    default:
                        v5 = (v5 % ((v5) | 1))
                    }
                case "foo bar":
                    if true {
                        var v72 uint64 = 18446744073709551615
                        v4 -= float32(3.14159)
                        _ = v72
                    } else {
                        for fi73 := 0; fi73 < 10; fi73++ {
                            var v74 uint64 = 1
                            v1 = append(v1, 2147483647)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi73
                            _ = v74
                        }
                        if (true || false) {
                            v67 = "foo bar"
                        } else {
                            if false {
                                for fi75 := 0; fi75 < 12; fi75++ {
                                    for fi76 := 0; fi76 < 8; fi76++ {
                                        if (12 == 156) {
                                            fmt.Println(fi75, v5)
                                            var v77 int16 = -879
                                            _ = v77
                                        }
                                        _ = len(v1)
                                        switch fi76 {
                                        case -2147483648:
                                            switch 0 {
                                            case 0:
                                                switch fi76 {
                                                case -1:
                                                    fmt.Println(v2, v67)
                                                case 347:
                                                    v78 := GenS1{F0: 0, F1: 2}
                                                    _ = v78
                                                case 902:
                                                    fmt.Println(v4, v2, fi75)
                                                default:
                                                    v3 = 13
                                                }
                                                v79 := GenS1{F0: 1000, F1: 2}
                                                _ = v79
                                            case 32768:
                                                v80 := GenS1{F0: -379, F1: 42}
                                                _ = v80
                                            default:
                                                var v81 int64 = 1
                                                _ = v81
                                            }
                                        case 32767:
                                            v5 += v5
                                        case 0:
                                            switch v67 {
                                            case "hello":
                                                v3 = 2
                                            default:
                                                v82 := GenS2{F0: 2147483647}
                                                _ = v82
                                            }
                                        default:
                                            v83 := GenS4{F0: 4294967295, F1: 0, F2: 0}
                                            _ = v83
                                        }
                                        _ = fi76
                                    }
                                    v2 += v3
                                    v1 = append(v1, fi75)
                                    _ = fi75
                                }
                                if len(v1) > 0 {
                                    v1[0] = 761
                                }
                                var v84 uint32 = 0
                                _ = v84
                            }
                            if len(v1) > 0 {
                                v1[0] = 9223372036854775807
                            }
                        }
                        v85 := GenS5{F0: genFn1(), F1: -1, F2: 0}
                        _ = v85
                    }
                }
                var v86 uint32 = 4294967295
                _ = v86
            }
            _ = v67
        }
    default:
        v87 := GenS3{F0: v5, F1: 0}
        _ = v87
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
