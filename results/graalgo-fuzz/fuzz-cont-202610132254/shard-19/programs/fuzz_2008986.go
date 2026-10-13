package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 int32; F2 string }
type GenS10 struct { F0 uint }

func main() {
    v1 := []uint32{65535, 2147483648, 321}
    v2 := []int{}
    var v3 int16 = -32768
    var v4 int16 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v5 uint16 = 0
        _ = v5
    case -987:
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 float32 = 1.5
            v4 = v3
            _ = fi6
            _ = v7
        }
    case 32768:
        if (false && false) {
            var v8 int16 = (v3 % ((345) | 1))
            switch "hello" {
            case "hello":
                var v9 int8 = -128
                _ = v9
            case "a":
                v10 := GenS1{F0: 972, F1: 439}
                _ = v10
            default:
                v3 = v4
            }
            _ = v8
        }
    }
    v4 = (v4 - v4)
    if len(v1) > 0 {
        v1[0] = 1
    }
    if false {
        if (1 < 9223372036854775807) {
            for fi11 := 0; fi11 < 2; fi11++ {
                v4 -= (v4 * v4)
                if (!false) {
                    _ = len(v2)
                    if true {
                        v3 += v4
                        for fi12 := 0; fi12 < 12; fi12++ {
                            if (-2.25 == 0.0) {
                                var v13 int32 = 42
                                _ = v13
                            } else {
                                fmt.Println(fi11, v4, fi12)
                                v4 = (v3 | (v3 << ((0) & 63)))
                                switch -131 {
                                case 42:
                                    var v14 bool = false
                                    _ = v14
                                case 127:
                                    v15 := GenS1{F0: 2, F1: 935}
                                    _ = v15
                                default:
                                    fmt.Println(fi12)
                                }
                            }
                            var v16 float32 = float32(0.1)
                            _ = fi12
                            _ = v16
                        }
                    }
                    var v17 int16 = v4
                    _ = v17
                }
                _ = fi11
            }
            if false {
                var v18 float32 = 0.0
                if (18446744073709551615 != 0) {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if len(v1) > 0 {
                            v1[0] = 39
                        }
                        fmt.Println(v4)
                        _ = fi19
                    }
                }
                _ = v18
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                if false {
                    var v21 float32 = float32((3.14159 + ((0.0 + float64((((0.1 * ((1.5 / float32(1.5)) * 3.14159)) + -2.25) / 1.0))) - 1.5)))
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v3 += v3
                        v4 += 0
                        var v23 string = ("hello" + "")
                        _ = fi22
                        _ = v23
                    }
                    var v24 uint16 = 42
                    _ = v21
                    _ = v24
                } else {
                    v25 := GenS2{F0: -9223372036854775808}
                    _ = v25
                }
                _ = fi20
            }
        } else {
            var v26 float32 = 1e10
            switch "a" {
            case "foo bar":
                _ = len(v1)
            case "\n\t":
                var v27 int8 = 96
                _ = v27
            case "":
                var v28 int8 = -4
                _ = v28
            default:
                v2 = append(v2, -2147483648)
            }
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = v26
        }
        v29 := GenS1{F0: 2, F1: 271}
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v29
    } else {
        for fi30 := 0; fi30 < 13; fi30++ {
            v3 -= 0
            if false {
                var v31 string = "a"
                _ = v31
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(fi30, v3, v4)
                case "foo bar":
                    v32 := GenS2{F0: -9223372036854775808}
                    _ = v32
                default:
                    v3 -= (v3 + 1)
                }
            }
            _ = fi30
        }
        if true {
            switch 9223372036854775807 {
            case -103:
                var v33 int16 = (-(v4))
                _ = v33
            case 32768:
                v34 := GenS1{F0: 0, F1: 503}
                _ = v34
            }
            fmt.Println(v3, v4)
        } else {
            switch "hello" {
            case "a":
                for fi35 := 0; fi35 < 8; fi35++ {
                    fmt.Println(fi35, v4, fi35)
                    var v36 int = fi35
                    fmt.Println(fi35)
                    _ = fi35
                    _ = v36
                }
            }
        }
    }
    var v37 int64 = 9223372036854775807
    switch 0 {
    case -295:
        var v38 int16 = v3
        _ = v38
    }
    var v39 uint32 = 42
    var v40 uint = 4294967295
    v41 := GenS2{F0: (^(v37))}
    switch "foo bar" {
    case "hello":
        var v42 bool = true
        _ = v42
    case "a":
        v40 -= v40
    }
    if true {
        var v43 uint = v40
        v43 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v43
    }
    v44 := GenS3{F0: "", F1: "a"}
    v1 = append(v1, 450)
    var v45 int16 = (v4 / (((v4 ^ v3)) | 1))
    if (63 != 2147483647) {
        v40 = 708
        if ((!((("hello" + ("foo bar" + "hello")) == "") && true)) || (42 != 1)) {
            if false {
                var v46 int16 = 0
                if true {
                    if (-128 == -128) {
                        v3 = 1000
                        v47 := GenS4{F0: 42, F1: ((1 <= 42) && (int8(v39) < 0))}
                        for fi48 := 0; fi48 < 8; fi48++ {
                            switch "foo bar" {
                            case "a":
                                var v49 uint64 = 18446744073709551615
                                _ = v49
                            default:
                                var v50 uint16 = 630
                                _ = v50
                            }
                            if true {
                                for fi51 := 0; fi51 < 7; fi51++ {
                                    switch (-(fi48)) {
                                    case -9223372036854775808:
                                        var v52 uint8 = 128
                                        _ = v52
                                    }
                                    v53 := GenS5{F0: fi48}
                                    _ = fi51
                                    _ = v53
                                }
                                v39 = v39
                                for fi54 := 0; fi54 < 20; fi54++ {
                                    switch -845 {
                                    case 32768:
                                        switch 32767 {
                                        case 9223372036854775807:
                                            _ = len(v2)
                                        case 692:
                                            v55 := GenS1{F0: 316, F1: 2}
                                            _ = v55
                                        }
                                    case 0:
                                        v1 = append(v1, v39)
                                    }
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v40, fi54, fi54)
                                    case "\n\t":
                                        v37 = -9223372036854775808
                                    case "hello":
                                        _ = len(v2)
                                    default:
                                        v56 := GenS1{F0: uint64(v39), F1: (^(v39))}
                                        _ = v56
                                    }
                                    if (!false) {
                                        v57 := GenS6{F0: 2, F1: -106}
                                        fmt.Println(v4, v40, v45)
                                        v39 -= (^(v39))
                                        _ = v57
                                    } else {
                                        switch 32768 {
                                        case -613:
                                            v3 = v3
                                        case 32767:
                                            v58 := GenS7{F0: 2}
                                            _ = v58
                                        default:
                                            var v59 int16 = v46
                                            _ = v59
                                        }
                                    }
                                    _ = fi54
                                }
                            } else {
                                switch "hello" {
                                case "":
                                    for fi60 := 0; fi60 < 10; fi60++ {
                                        v61 := GenS4{F0: 1, F1: (-2.25 != -2.25)}
                                        if false {
                                            v62 := GenS6{F0: 255, F1: (^(v37))}
                                            _ = v62
                                        }
                                        switch "" {
                                        case "":
                                            var v63 uint8 = 0
                                            _ = len(v2)
                                            var v64 bool = false
                                            _ = v63
                                            _ = v64
                                        case "foo bar":
                                            _ = len(v2)
                                        case "a":
                                            v2 = append(v2, fi48)
                                        }
                                        if ((0.1 / (1.5 / (1e10 / (0.1 - (1e10 / 1e10))))) < 1e10) {
                                            if ((!(!true)) || (-588 > 1)) {
                                                switch "" {
                                                case "a":
                                                    v65 := GenS8{F0: (v40 | (v40 + 0))}
                                                    _ = v65
                                                case "hello":
                                                    fmt.Println(v45, v4)
                                                }
                                                fmt.Println(v3)
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v66 uint32 = (v39 ^ (v39 ^ 42))
                                            _ = v66
                                        }
                                        _ = fi60
                                        _ = v61
                                    }
                                case "hello":
                                    fmt.Println(v3, v39)
                                default:
                                    v3 -= v4
                                }
                                switch "\n\t" {
                                case "a":
                                    var v67 int = fi48
                                    _ = v67
                                case "hello":
                                    v68 := GenS9{F0: 1, F1: 0, F2: (("\n\t" + ("hello" + (("hello" + ("" + "foo bar")) + ("" + "hello")))) + "")}
                                    _ = v68
                                default:
                                    _ = len(v1)
                                }
                                v69 := GenS4{F0: 719, F1: (!false)}
                                _ = v69
                            }
                            v70 := GenS4{F0: 1, F1: true}
                            _ = fi48
                            _ = v70
                        }
                        _ = v47
                    } else {
                        for fi71 := 0; fi71 < 16; fi71++ {
                            fmt.Println(v3)
                            v45 = v3
                            _ = fi71
                        }
                        v2 = append(v2, -1)
                    }
                    if ((((float64((0.1 + 0.0)) * 1.5) + 0.1) + (0.0 * -2.25)) < 1e10) {
                        var v72 uint16 = 32767
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v73 int64 = v37
                            fmt.Println(v4)
                            _ = v73
                        }
                        _ = v72
                    }
                    v40 -= v40
                }
                v37 = (-(v37))
                _ = v46
            }
            _ = len(v1)
            if false {
                if (!(false && false)) {
                    switch (("foo bar" + ("a" + ("a" + ""))) + "hello") {
                    case "":
                        var v74 uint64 = 42
                        _ = v74
                    case "hello":
                        v75 := GenS4{F0: 32767, F1: (!false)}
                        _ = v75
                    }
                    v40 -= (v40 << (((v40 ^ 255)) & 63))
                } else {
                    for fi76 := 0; fi76 < 20; fi76++ {
                        v39 = (v39 * (v39 / ((42) | 1)))
                        _ = fi76
                    }
                }
                fmt.Println(v4, v45)
            }
        } else {
            var v77 uint8 = 77
            _ = v77
        }
    } else {
        v78 := GenS10{F0: (v40 >> ((v40) & 63))}
        var v79 uint64 = 1
        _ = v78
        _ = v79
    }
    v3 = int16(v39)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
}
