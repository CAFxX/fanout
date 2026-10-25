package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }
func genFn0() string {
    var v7 bool = true
    var v8 int = 0
    v8 += (v8 ^ v8)
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10)
            var v11 int16 = 0
            fmt.Println(fi10)
            v12 := []int{fi9, (-(v8)), v8, (-(fi9))}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int8 = -1
        v8 = (fi9 >> ((4294967295) & 63))
        _ = fi9
        _ = v13
    }
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 float64 = (((3.14159 / 1.5) - (0.0 - (3.14159 - 0.1))) * 3.14159)
    v4 := make(map[int]uint32)
    var v5 int16 = 608
    var v6 bool = true
    if v6 {
        switch "" {
        case "":
            var v14 uint64 = 42
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 13; fi15++ {
                if (18446744073709551615 >= 1) {
                    v16 := GenS1{F0: 1, F1: (v5 - 665), F2: 180}
                    var v17 float64 = ((float64((-2.25 * 3.14159)) - 0.0) - v3)
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
        default:
            if (!v6) {
                v18 := GenS1{F0: 1, F1: (v5 >> ((2) & 63)), F2: 1}
                if (((("foo bar" + "\n\t") + ("" + "\n\t")) + "") != ((("hello" + ("foo bar" + ("foo bar" + "hello"))) + ("hello" + ("\n\t" + "hello"))) + "hello")) {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if (176 <= 239) {
                            v3 = float64(-2.25)
                            var v20 string = "\n\t"
                            v21 := GenS1{F0: 18446744073709551615, F1: (v5 & v5), F2: 127}
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                    if v6 {
                        if v6 {
                            var v22 uint = 161
                            switch ("hello" + "foo bar") {
                            case "hello":
                                fmt.Println(v5, v22)
                            }
                            _ = v22
                        } else {
                            switch 486 {
                            case 2147483647:
                                v3 += ((0.0 + v3) - v3)
                            }
                            for fi23 := 0; fi23 < 14; fi23++ {
                                fmt.Println(v3, v5)
                                fmt.Println(fi23, v6)
                                _ = fi23
                            }
                            for fi24 := 0; fi24 < 12; fi24++ {
                                v6 = v6
                                if v6 {
                                    for fi25 := 0; fi25 < 10; fi25++ {
                                        var v26 int32 = -1
                                        if (!(!v6)) {
                                            if v6 {
                                                v3 += v3
                                                if v6 {
                                                    _ = len(v1)
                                                    switch "a" {
                                                    case "":
                                                        var v27 float64 = v3
                                                        _ = v27
                                                    case "a":
                                                        v5 += v5
                                                    default:
                                                        var v28 bool = false
                                                        _ = v28
                                                    }
                                                }
                                                v29 := GenS2{F0: 985}
                                                _ = v29
                                            }
                                            if (!((true || (-48 == -1)) || (!false))) {
                                                var v30 float32 = (0.1 + 0.1)
                                                v31 := v4[(-(fi24))]
                                                var v32 string = "foo bar"
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                        if (-9223372036854775808 <= 42) {
                                            v33 := GenS3{F0: (v26 / ((42) | 1)), F1: (v26 * v26)}
                                            fmt.Println(v5, v5)
                                            if (!((375 != 42) || v6)) {
                                                v5 = (^(v5))
                                            }
                                            _ = v33
                                        }
                                        var v34 bool = ("\n\t" < "a")
                                        _ = fi25
                                        _ = v26
                                        _ = v34
                                    }
                                }
                                var v35 uint16 = 32767
                                _ = fi24
                                _ = v35
                            }
                        }
                    }
                }
                var v36 bool = true
                _ = v18
                _ = v36
            } else {
                switch -1 {
                case 989:
                    var v37 uint64 = 1
                    _ = v37
                case -804:
                    v5 = -151
                case 1:
                    for fi38 := 0; fi38 < 15; fi38++ {
                        fmt.Println(v6, v6)
                        _ = len(v1)
                        _ = fi38
                    }
                }
                v6 = false
            }
        }
        v39 := GenS2{F0: 2}
        if v6 {
            switch 1 {
            case -1:
                if (v6 || v6) {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v2[fi40] = 424
                        _ = fi40
                    }
                }
            case 9223372036854775807:
                fmt.Println(v5)
            default:
                switch -85 {
                case 127:
                    for fi41 := 0; fi41 < 20; fi41++ {
                        v42 := GenS1{F0: 2, F1: 42, F2: 141}
                        _ = fi41
                        _ = v42
                    }
                case 32768:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v44 := GenS3{F0: -1, F1: -661}
                        v6 = (v6 && false)
                        _ = fi43
                        _ = v44
                    }
                case -42:
                    v3 -= float64(1e10)
                }
            }
            var v45 uint32 = 10
            var v46 float64 = (((float64(3.14159) - -2.25) / v3) * (-2.25 / v3))
            _ = v45
            _ = v46
        } else {
            var v47 uint16 = 642
            if (34 <= 2) {
                v48 := GenS1{F0: 1, F1: (v5 >> ((255) & 63)), F2: 227}
                for fi49 := 0; fi49 < 7; fi49++ {
                    v50 := v4[1]
                    _ = fi49
                    _ = v50
                }
                v51 := []int{-856}
                _ = v48
                _ = v51
            } else {
                v52 := []int{128, 9223372036854775807, 965}
                switch 0 {
                case 248:
                    for fi53 := 0; fi53 < 3; fi53++ {
                        v54 := GenS4{F0: 1, F1: 127}
                        v3 = 3.14159
                        var v55 float32 = ((((float32(((1e10 / v3) / 1.5)) + ((-2.25 / 1.0) * (1e10 + (1e10 * (0.1 * (((float32(v3) - 0.1) + 3.14159) - 1e10)))))) * 0.0) * (0.1 / 1.0)) - 3.14159)
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                case 127:
                    var v56 uint = 65535
                    _ = v56
                case 32768:
                    v57 := GenS1{F0: 18446744073709551615, F1: (v5 % (((v5 ^ 32767)) | 1)), F2: 128}
                    _ = v57
                default:
                    v6 = false
                }
                var v58 int = 9223372036854775807
                _ = v52
                _ = v58
            }
            _ = v47
        }
        _ = v39
    }
    v59 := GenS4{F0: -1, F1: -960}
    switch ("foo bar" + "foo bar") {
    case "hello":
        v60 := []int{}
        _ = v60
    }
    v61 := GenS4{F0: 1, F1: 267}
    var v62 string = ""
    fmt.Println(v5, v3, v62)
    switch "foo bar" {
    case "hello":
        v63 := GenS5{F0: 1}
        _ = v63
    case "":
        var v64 uint64 = 0
        _ = v64
    case "\n\t":
        v6 = (v62 < "foo bar")
    default:
        var v65 int = 298
        _ = v65
    }
    var v66 uint32 = 65535
    v67 := []int{-2147483648, -9223372036854775808, 127}
    v68 := GenS3{F0: 168, F1: 42}
    var v69 string = v62
    v62 = v69
    for fi70 := 0; fi70 < 18; fi70++ {
        var v71 int16 = 448
        if len(v67) >= 2 {
            v67 = v67[1:len(v67)]
        }
        switch (fi70 % (((fi70 >> ((2) & 63))) | 1)) {
        case -360:
            v72 := GenS2{F0: v66}
            _ = v72
        case -80:
            switch (fi70 - (fi70 & 2147483647)) {
            case 128:
                v69 = ("\n\t" + v69)
            default:
                v4[fi70] = v66
            }
        case -431:
            var v73 int64 = -940
            _ = v73
        default:
            for fi74 := 0; fi74 < 4; fi74++ {
                var v75 float32 = 1.5
                v5 = v5
                var v76 int16 = (v5 ^ -1)
                _ = fi74
                _ = v75
                _ = v76
            }
        }
        _ = fi70
        _ = v71
    }
    v66 += v66
    v77 := GenS3{F0: 161, F1: -350}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v61
    _ = v62
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v77
}
