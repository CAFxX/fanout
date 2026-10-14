package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = (("" + ("hello" + ("\n\t" + ("hello" + "")))) + "a")
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    var v5 int64 = 1
    delete(v3, 0)
    if ((!false) || false) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v5 = 42
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(fi7)
                fmt.Println(v5)
                v5 += v5
                _ = fi7
            }
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int8 = 42
                _ = fi8
                _ = v9
            }
            _ = fi6
        }
        var v10 float32 = -2.25
        if (!(1.5 >= float64(v10))) {
            switch -2147483648 {
            case 32767:
                if true {
                    switch "\n\t" {
                    case "a":
                        v5 += (v5 + (v5 - v5))
                    }
                    v10 -= (-2.25 + v10)
                }
            case 42:
                for fi11 := 0; fi11 < 2; fi11++ {
                    fmt.Println(fi11, v5)
                    _ = fi11
                }
            case 99:
                var v12 int = 32767
                _ = v12
            default:
                switch 42 {
                case -666:
                    var v13 int8 = 127
                    _ = v13
                case 4294967295:
                    switch (v2 + "hello") {
                    case "":
                        if (0.1 == ((v10 * v10) / (1.5 - v10))) {
                            v5 = (v5 | (^(v5)))
                            v5 = 0
                        } else {
                            var v14 float32 = 0.1
                            _ = v14
                        }
                    }
                case -1:
                    var v15 int16 = 1000
                    _ = v15
                }
            }
        } else {
            if true {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := []int{32767, fi16, (fi16 - (fi16 % ((fi16) | 1)))}
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 uint32 = 2147483647
                        if true {
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                v2 = ("a" + (v2 + (v2 + "a")))
                            }
                            if (v19 > v19) {
                                var v20 float32 = (v10 / v10)
                                v21 := GenS1{F0: 18446744073709551615, F1: -97, F2: 0}
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(v5, v10, fi16)
                        } else {
                            v10 = (float32(-2.25) - 0.1)
                            if true {
                                var v22 int16 = 32767
                                fmt.Println(fi18, v2, v10)
                                _ = v22
                            } else {
                                var v23 uint = 255
                                if true {
                                    var v24 bool = (1 == 1)
                                    if true {
                                        v23 = 228
                                    }
                                    _ = v24
                                } else {
                                    v3[(^(v19))] = (false && false)
                                    var v25 float32 = v10
                                    var v26 uint = (v23 + (v23 % ((492) | 1)))
                                    fmt.Println(v19, fi18)
                                    _ = v25
                                    _ = v26
                                }
                                if false {
                                    switch "" {
                                    case "hello":
                                        v23 = (^(v23))
                                        fmt.Println(v10, v19)
                                    case "\n\t":
                                        v27 := GenS1{F0: 2, F1: 127, F2: 255}
                                        _ = v27
                                    default:
                                        var v28 int = 205
                                        _ = v28
                                    }
                                    if true {
                                        if (!(v10 > v10)) {
                                            v29 := GenS2{F0: v2, F1: 1}
                                            var v30 float32 = (v10 - v10)
                                            var v31 string = (((v2 + (v2 + (("hello" + ((("a" + v2) + ("\n\t" + v2)) + "a")) + v2))) + "a") + "a")
                                            fmt.Println(fi18, v2)
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        } else {
                                            fmt.Println(fi18, v5)
                                            var v32 uint = 4294967295
                                            _ = v32
                                        }
                                    } else {
                                        _ = len(v3)
                                        var v33 int32 = 350
                                        v33 = (v33 | v33)
                                        _ = v33
                                    }
                                }
                                _ = v23
                            }
                            if false {
                                v19 -= 83
                            } else {
                                var v34 uint16 = 42
                                _ = v34
                            }
                            var v35 uint8 = 205
                            _ = v35
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            } else {
                var v36 float32 = ((((((float32(0.0) + v10) * v10) + v10) - (0.0 - float32(0.0))) - v10) - 1.5)
                v2 = (v2 + "a")
                _ = v36
            }
        }
        _ = v10
    } else {
        delete(v3, 2)
        v37 := GenS2{F0: ("hello" + v2), F1: 7}
        _ = v37
    }
    v2 = ("" + v2)
    v38 := GenS3{F0: 127, F1: 42, F2: 1}
    switch v2 {
    case "\n\t":
        for fi39 := 0; fi39 < 6; fi39++ {
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 int = (fi40 << ((867) & 63))
                _ = fi40
                _ = v41
            }
            if (-2147483648 > 42) {
                var v42 uint32 = 474
                v2 = "foo bar"
                if false {
                    v1[v42] = 42
                    switch (^(fi39)) {
                    case 4294967295:
                        v43 := []int{}
                        _ = v43
                    default:
                        v5 -= v5
                    }
                }
                _ = v42
            }
            if (-443 > 753) {
                for fi44 := 0; fi44 < 10; fi44++ {
                    var v45 float64 = (((3.14159 / (float64(float32((((((-2.25 + (3.14159 - 0.1)) + 1e10) + 1e10) + ((float64(((0.0 * (1.5 + -2.25)) / 1.0)) - float64(1.5)) + 1.5)) + (1.5 * 3.14159)))) / ((0.0 - 0.0) + float64((1e10 / 1.0))))) * 0.0) / float64(1.5))
                    var v46 int8 = 123
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                v2 = v2
            } else {
                switch fi39 {
                case 32767:
                    var v47 int64 = 42
                    _ = v47
                default:
                    for fi48 := 0; fi48 < 2; fi48++ {
                        var v49 float64 = (-2.25 + 1e10)
                        var v50 float32 = -2.25
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
                v5 = v5
                switch v2 {
                case "hello":
                    if false {
                        if true {
                            v2 = ("" + "a")
                        }
                    }
                case "\n\t":
                    for fi51 := 0; fi51 < 9; fi51++ {
                        fmt.Println(v5, fi51, fi51)
                        switch (v2 + v2) {
                        case "foo bar":
                            if (0 <= 0) {
                                switch (((("" + v2) + (v2 + v2)) + v2) + v2) {
                                case "\n\t":
                                    v52 := GenS4{F0: fi51, F1: -32768, F2: 18446744073709551615}
                                    _ = v52
                                case "foo bar":
                                    var v53 int32 = -503
                                    _ = v53
                                }
                                if true {
                                    var v54 float64 = 1.5
                                    _ = v54
                                }
                                _ = len(v4)
                            }
                        }
                        v5 = (v5 - v5)
                        _ = fi51
                    }
                default:
                    if false {
                        fmt.Println(fi39)
                        v55 := []int{fi39, 127, (-(fi39))}
                        _ = v55
                    } else {
                        v5 -= (v5 % ((42) | 1))
                        v2 = (v2 + v2)
                        v5 += v5
                    }
                }
            }
            _ = fi39
        }
    case "":
        for fi56 := 0; fi56 < 20; fi56++ {
            switch 127 {
            case 32767:
                switch -9223372036854775808 {
                case 32768:
                    fmt.Println(v2)
                case 1:
                    var v57 int = 32768
                    _ = v57
                case 32767:
                    for fi58 := 0; fi58 < 5; fi58++ {
                        if (1 >= 157) {
                            var v59 int32 = -1
                            if (0 <= 31) {
                                if false {
                                    v5 -= -9223372036854775808
                                    fmt.Println(v5, v5)
                                    if true {
                                        v59 = 0
                                        v5 = v5
                                        var v60 float64 = float64(3.14159)
                                        _ = v60
                                    }
                                }
                                v61 := GenS5{F0: 0}
                                _ = v61
                            }
                            var v62 bool = true
                            _ = v59
                            _ = v62
                        }
                        v2 = v2
                        _ = fi58
                    }
                }
            }
            var v63 uint32 = 42
            switch fi56 {
            case 2147483647:
                for fi64 := 0; fi64 < 12; fi64++ {
                    if (4294967295 >= 312) {
                        var v65 string = v2
                        _ = v65
                    }
                    _ = fi64
                }
            default:
                v66 := []int{127, -414}
                _ = v66
            }
            _ = fi56
            _ = v63
        }
    case "foo bar":
        fmt.Println(v5, v2)
    }
    v67 := []int{127}
    if (((1.5 * 1e10) - 0.1) <= float64(0.1)) {
        if false {
            if len(v67) > 0 {
                v67[0] = -128
            }
        } else {
            var v68 int64 = v5
            _ = v68
        }
        fmt.Println(v5, v5)
        v2 = (v2 + v2)
    } else {
        var v69 uint16 = 0
        for fi70 := 0; fi70 < 5; fi70++ {
            switch v2 {
            case "hello":
                v71 := v3[uint32(v69)]
                _ = v71
            case "\n\t":
                v72 := v3[37]
                _ = v72
            case "foo bar":
                fmt.Println(fi70, v5, v5)
            default:
                v67 = append(v67, (fi70 >> ((611) & 63)))
            }
            if len(v67) >= 2 {
                v67 = v67[1:len(v67)]
            }
            _ = fi70
        }
        for fi73 := 0; fi73 < 15; fi73++ {
            v69 -= v69
            if true {
                var v74 string = (((("foo bar" + ("\n\t" + "foo bar")) + "hello") + ("hello" + v2)) + v2)
                _ = v74
            }
            _ = fi73
        }
        _ = v69
    }
    switch "a" {
    case "\n\t":
        v75 := GenS3{F0: 255, F1: -1, F2: 65535}
        _ = v75
    case "a":
        v76 := GenS3{F0: 217, F1: -62, F2: 846}
        _ = v76
    case "hello":
        if (128 <= 32768) {
            var v77 uint8 = 67
            switch -2147483648 {
            case 1:
                for fi78 := 0; fi78 < 2; fi78++ {
                    if false {
                        v67 = append(v67, fi78)
                        var v79 uint8 = (v77 + (v77 | 127))
                        _ = v79
                    }
                    if true {
                        if len(v67) > 0 {
                            v67[0] = (fi78 | fi78)
                        }
                        var v80 int64 = (^(v5))
                        _ = v80
                    }
                    v5 = 42
                    _ = fi78
                }
            case 602:
                switch "a" {
                case "a":
                    if (42 > 894) {
                        var v81 uint = 18446744073709551615
                        _ = v81
                    }
                default:
                    v77 += (v77 / (((v77 & 1)) | 1))
                }
            case 423:
                v2 = "\n\t"
            }
            fmt.Println(v77, v5)
            _ = v77
        }
    }
    if len(v67) > 0 {
        v67[0] = 4294967295
    }
    if (true && ((2 >= 2147483647) && (!(true && (857 == 42))))) {
        v82 := GenS6{F0: -525, F1: -128}
        if true {
            fmt.Println(v2, v5)
            var v83 uint64 = 4294967295
            _ = v83
        } else {
            for fi84 := 0; fi84 < 11; fi84++ {
                v85 := GenS7{F0: 1}
                v4[((v2 + ((((v2 + (v2 + "hello")) + (v2 + "\n\t")) + v2) + v2)) + v2)] = "foo bar"
                _ = fi84
                _ = v85
            }
            var v86 int = 128
            _ = v86
        }
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v67
}
