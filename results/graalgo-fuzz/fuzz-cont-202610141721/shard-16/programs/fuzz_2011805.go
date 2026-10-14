package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int64, p1 int) int16 {
    var v4 int8 = 0
    for fi5 := 0; fi5 < 10; fi5++ {
        p1 = p1
        p0 += p0
        _ = fi5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 445
}

func main() {
    var v1 bool = true
    var v2 int32 = -10
    v3 := make(map[string]bool)
    var v6 uint8 = 2
    v6 += (v6 / ((v6) | 1))
    switch -2147483648 {
    case -625:
        v7 := GenS1{F0: -11, F1: v6}
        _ = v7
    default:
        v8 := GenS1{F0: -43, F1: v6}
        _ = v8
    }
    v6 = 42
    v9 := []int{128}
    switch -2147483648 {
    case 677:
        v2 = 102
    case 0:
        v3["\n\t"] = (0 >= -109)
    case 32768:
        if (65535 < 18446744073709551615) {
            for fi10 := 0; fi10 < 13; fi10++ {
                switch "a" {
                case "":
                    v2 = (v2 | (v2 | -1))
                case "\n\t":
                    if ((2 > 2147483648) || v1) {
                        v11 := v3[(("\n\t" + ("" + "foo bar")) + ("hello" + ""))]
                        _ = v11
                    }
                case "hello":
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 uint16 = 102
                        _ = fi12
                        _ = v13
                    }
                default:
                    v2 += (v2 / (((v2 & (-(v2)))) | 1))
                }
                _ = fi10
            }
            if (!v1) {
                var v14 float64 = -2.25
                _ = v14
            } else {
                v1 = ((1 <= 9223372036854775807) && (v2 >= -1))
            }
        } else {
            if v1 {
                v15 := GenS2{F0: (v6 % ((v6) | 1))}
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 float64 = 0.1
                    _ = fi16
                    _ = v17
                }
                _ = v15
            } else {
                switch "" {
                case "hello":
                    var v18 int8 = -1
                    _ = v18
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    if v1 {
                        v1 = (v6 > v6)
                        v1 = true
                        var v20 int = fi19
                        _ = v20
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v21 string = "foo bar"
                            _ = v21
                        case "hello":
                            fmt.Println(v6, v6)
                        default:
                            var v22 int32 = v2
                            _ = v22
                        }
                        if v1 {
                            switch ("\n\t" + "") {
                            case "hello":
                                var v23 uint64 = 4294967295
                                _ = v23
                            case "foo bar":
                                if v1 {
                                    switch "a" {
                                    case "a":
                                        if v1 {
                                            for fi24 := 0; fi24 < 2; fi24++ {
                                                switch fi19 {
                                                case 887:
                                                    var v25 float32 = (((1.5 - 0.0) / (-2.25 - ((-2.25 + (-2.25 * -2.25)) / (-2.25 - (0.0 - float32(0.1)))))) / (3.14159 - (1.5 * (1.5 + 1e10))))
                                                    _ = v25
                                                default:
                                                    v26 := GenS3{F0: fi24, F1: 42}
                                                    _ = v26
                                                }
                                                _ = fi24
                                            }
                                            v27 := GenS2{F0: 1}
                                            _ = v27
                                        } else {
                                            fmt.Println(v6, v1, v1)
                                            for fi28 := 0; fi28 < 14; fi28++ {
                                                fmt.Println(fi28, fi19)
                                                v29 := GenS4{F0: (-(fi19)), F1: 127}
                                                var v30 string = "\n\t"
                                                _ = fi28
                                                _ = v29
                                                _ = v30
                                            }
                                            v31 := GenS5{F0: (v2 << ((42) & 63)), F1: "\n\t", F2: -305}
                                            _ = v31
                                        }
                                    case "\n\t":
                                        if v1 {
                                            if (v6 != (v6 * v6)) {
                                                for fi32 := 0; fi32 < 16; fi32++ {
                                                    v1 = v1
                                                    _ = fi32
                                                }
                                                switch 1 {
                                                case 4294967295:
                                                    for fi33 := 0; fi33 < 12; fi33++ {
                                                        v6 = (^(v6))
                                                        var v34 float32 = 0.1
                                                        _ = fi33
                                                        _ = v34
                                                    }
                                                default:
                                                    for fi35 := 0; fi35 < 6; fi35++ {
                                                        v36 := GenS5{F0: v2, F1: ("" + ("hello" + (("" + "") + (("" + ("foo bar" + "\n\t")) + "hello")))), F2: -68}
                                                        var v37 int64 = 0
                                                        _ = fi35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                                if v1 {
                                                    if v1 {
                                                        v9 = append(v9, 0)
                                                    }
                                                    v38 := GenS6{F0: (42 != -9223372036854775808)}
                                                    _ = v38
                                                }
                                            } else {
                                                var v39 uint = 748
                                                if (!false) {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                }
                                                _ = v39
                                            }
                                            var v40 float32 = -2.25
                                            switch "" {
                                            case "\n\t":
                                                var v41 int32 = v2
                                                _ = v41
                                            case "hello":
                                                fmt.Println(v40, v1)
                                            case "a":
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                            default:
                                                v40 += (1e10 + -2.25)
                                            }
                                            _ = v40
                                        } else {
                                            var v42 int64 = 0
                                            var v43 uint = 825
                                            switch "\n\t" {
                                            case "a":
                                                switch (fi19 / (((^(fi19))) | 1)) {
                                                case 903:
                                                    var v44 uint64 = 1
                                                    _ = v44
                                                case 128:
                                                    for fi45 := 0; fi45 < 7; fi45++ {
                                                        fmt.Println(v1, v6)
                                                        var v46 int = fi19
                                                        v1 = false
                                                        _ = fi45
                                                        _ = v46
                                                    }
                                                case 2147483647:
                                                    v2 = -2147483648
                                                default:
                                                    var v47 string = ""
                                                    _ = v47
                                                }
                                            }
                                            _ = v42
                                            _ = v43
                                        }
                                    case "hello":
                                        v48 := GenS7{F0: (65535 <= 1), F1: 248}
                                        _ = v48
                                    }
                                }
                            case "a":
                                switch (^(fi19)) {
                                case 9223372036854775807:
                                    var v49 uint = 255
                                    _ = v49
                                }
                            }
                            var v50 int32 = (v2 * int32(fi19))
                            _ = v50
                        } else {
                            fmt.Println(v1)
                        }
                        v51 := GenS4{F0: (fi19 & fi19), F1: 13}
                        _ = v51
                    }
                    _ = fi19
                }
            }
            v52 := v3["a"]
            var v53 uint64 = 0
            _ = v52
            _ = v53
        }
    default:
        v6 += 127
    }
    var v54 float64 = -2.25
    if ("\n\t" > "hello") {
        v55 := GenS8{F0: 4294967295}
        var v56 uint64 = 18446744073709551615
        var v57 int32 = 1
        _ = v55
        _ = v56
        _ = v57
    }
    if (61 <= v6) {
        if ((-1 == -1) && ("" >= "")) {
            v3[("" + "")] = ((false && false) || true)
            for fi58 := 0; fi58 < 1; fi58++ {
                if v1 {
                    v6 -= 42
                    if v1 {
                        if v1 {
                            var v59 int16 = 42
                            v3[("a" + "hello")] = v1
                            if ((1.5 - v54) != v54) {
                                switch 1 {
                                case -1:
                                    fmt.Println(v1, v59, v59)
                                case 32767:
                                    var v60 int16 = v59
                                    _ = v60
                                case -57:
                                    v6 = v6
                                }
                            }
                            _ = v59
                        }
                    }
                }
                for fi61 := 0; fi61 < 3; fi61++ {
                    fmt.Println(v1, v2)
                    if v1 {
                        fmt.Println(v54, v6)
                    } else {
                        _ = len(v9)
                        var v62 int32 = 0
                        var v63 uint32 = 4294967295
                        _ = v62
                        _ = v63
                    }
                    v64 := v3[("a" + "\n\t")]
                    v54 = float64(3.14159)
                    _ = fi61
                    _ = v64
                }
                _ = fi58
            }
        }
    } else {
        for fi65 := 0; fi65 < 6; fi65++ {
            v54 -= float64((((0.0 / 1.5) - 0.0) / (-2.25 + float32(v54))))
            if ((v1 && v1) || (((-2.25 * ((0.0 + 1.5) + 1e10)) > -2.25) || (uint16(v6) != 32767))) {
                v6 = 128
            }
            v66 := GenS8{F0: 1}
            _ = fi65
            _ = v66
        }
        fmt.Println(v2, v2)
    }
    v67 := GenS1{F0: -128, F1: 2}
    if v1 {
        v6 = 127
        v68 := GenS7{F0: (-1 >= -271), F1: 452}
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v6
    _ = v67
    _ = v9
}
