package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 uint64) string {
    v6 := GenS1{F0: (p1 - p1), F1: 663}
    fmt.Println(p0, p1)
    v7 := GenS1{F0: (p1 ^ p1), F1: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 int64) int64 {
    v8 := GenS1{F0: 18446744073709551615, F1: p0}
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "hello" {
        case "\n\t":
            if false {
                var v10 uint8 = 140
                var v11 uint = 305
                _ = v10
                _ = v11
            }
        case "":
            p0 = (p0 >> ((787) & 63))
        }
        v12 := GenS1{F0: 42, F1: p0}
        if (0.1 > 1e10) {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 string = ("foo bar" + "\n\t")
                v15 := []int{(fi13 & -9223372036854775808), (fi9 * (fi9 >> ((42) & 63)))}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = v8
    return (p0 - (p0 | 720))
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{"", "hello", "hello"}
    var v3 float32 = float32(0.1)
    var v4 float64 = (v1 + v1)
    v5 := []string{("hello" + "\n\t"), ("hello" + ""), ("hello" + ("\n\t" + ("foo bar" + "a"))), ("foo bar" + "a")}
    v16 := GenS1{F0: 128, F1: -714}
    var v17 uint64 = 571
    switch "foo bar" {
    case "":
        if (706 != 65535) {
            switch "a" {
            case "hello":
                fmt.Println(v17, v17, v17)
            case "\n\t":
                fmt.Println(v3, v3)
            }
        } else {
            var v18 float64 = v4
            v1 += v18
            var v19 string = ((("hello" + "") + "") + "hello")
            _ = v18
            _ = v19
        }
    case "hello":
        v4 += v1
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(fi20, fi20)
            v5 = append(v5, ("" + ("" + "foo bar")))
            _ = fi20
        }
    }
    var v21 uint8 = 42
    var v22 int32 = -1
    switch 2147483647 {
    case -1:
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 19; fi24++ {
                v21 = (v21 * (v21 / (((v21 >> ((255) & 63))) | 1)))
                var v25 int = -9223372036854775808
                _ = fi24
                _ = v25
            }
            var v26 int16 = 42
            for fi27 := 0; fi27 < 8; fi27++ {
                if false {
                    v28 := GenS2{F0: v17, F1: 4294967295}
                    var v29 int64 = -612
                    if true {
                        v30 := GenS3{F0: 42, F1: (v26 & v26)}
                        _ = len(v5)
                        _ = v30
                    } else {
                        v17 += v17
                        fmt.Println(fi27)
                    }
                    v4 = 3.14159
                    _ = v28
                    _ = v29
                } else {
                    if (true && (617 > v17)) {
                        v22 += 2147483647
                    } else {
                        if (-9223372036854775808 != 166) {
                            v31 := GenS4{F0: (!(0.1 != 0.0)), F1: v21, F2: (v22 >> ((42) & 63))}
                            v1 -= ((1e10 + v1) * v4)
                            _ = v31
                        } else {
                            var v32 uint8 = v21
                            _ = v32
                        }
                        var v33 float64 = 3.14159
                        if (!(v17 <= 903)) {
                            var v34 float64 = (0.1 + (float64(v3) * 0.1))
                            _ = v34
                        } else {
                            var v35 int32 = v22
                            if false {
                                v4 -= (v1 + 3.14159)
                                v26 = (v26 % ((v26) | 1))
                                var v36 uint32 = 65535
                                var v37 int64 = 1
                                _ = len(v5)
                                _ = v36
                                _ = v37
                            } else {
                                if (!(true || (true && true))) {
                                    if ((v3 / 1e10) > v3) {
                                        var v38 int8 = -21
                                        _ = v38
                                    }
                                }
                                if (!(!(true || true))) {
                                    switch -510 {
                                    case 330:
                                        v1 = v33
                                    case -9223372036854775808:
                                        var v39 uint16 = 0
                                        _ = v39
                                    case 32767:
                                        v40 := GenS5{F0: (v26 + v26)}
                                        _ = v40
                                    }
                                } else {
                                    v5 = append(v5, "foo bar")
                                    v22 = int32(v26)
                                    _ = len(v5)
                                }
                            }
                            v3 -= ((float32(((v33 + float64(v3)) / float64(-2.25))) + 0.1) * float32(3.14159))
                            _ = v35
                        }
                        _ = v33
                    }
                }
                v1 -= 0.0
                _ = fi27
            }
            _ = fi23
            _ = v26
        }
    default:
        var v41 uint64 = (v17 - (v17 | v17))
        _ = v41
    }
    if false {
        var v42 uint32 = 42
        v2 = append(v2, ("foo bar" + "foo bar"))
        switch ("" + "a") {
        case "hello":
            v1 = v1
        case "\n\t":
            switch 9223372036854775807 {
            case 128:
                if true {
                    if (false || (false && false)) {
                        if true {
                            fmt.Println(v4, v21, v21)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch "" {
                            case "\n\t":
                                if true {
                                    var v43 uint32 = v42
                                    v44 := GenS5{F0: 955}
                                    _ = v43
                                    _ = v44
                                } else {
                                    switch "" {
                                    case "":
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            if true {
                                                fmt.Println(v3, v21)
                                            } else {
                                                v46 := GenS4{F0: ((v17 & (^(v17))) == v17), F1: v21, F2: (v22 >> ((uint(v42)) & 63))}
                                                v21 += (v21 % ((255) | 1))
                                                _ = v46
                                            }
                                            _ = fi45
                                        }
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "\n\t":
                                        var v47 uint32 = 1
                                        _ = v47
                                    }
                                    v42 += v42
                                    switch 1 {
                                    case 4294967295:
                                        if true {
                                            for fi48 := 0; fi48 < 17; fi48++ {
                                                switch (^(fi48)) {
                                                case 1:
                                                    if (v3 <= v3) {
                                                        for fi49 := 0; fi49 < 5; fi49++ {
                                                            v3 += -2.25
                                                            _ = fi49
                                                        }
                                                        if false {
                                                            for fi50 := 0; fi50 < 12; fi50++ {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "\n\t"
                                                                }
                                                                v1 = (0.0 - float64(0.1))
                                                                var v51 int8 = -128
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v52 uint16 = 1
                                                                    _ = v52
                                                                case "a":
                                                                    var v53 float32 = (v3 * v3)
                                                                    _ = v53
                                                                case "":
                                                                    var v54 float32 = 1.5
                                                                    _ = v54
                                                                default:
                                                                    fmt.Println(fi48)
                                                                }
                                                                fmt.Println(v4)
                                                                _ = fi50
                                                                _ = v51
                                                            }
                                                            var v55 int32 = -1
                                                            v17 -= (v17 << ((251) & 63))
                                                            _ = v55
                                                        }
                                                    }
                                                case -809:
                                                    fmt.Println(v4, v21)
                                                case 917:
                                                    for fi56 := 0; fi56 < 1; fi56++ {
                                                        var v57 float64 = (((0.1 + (v4 / v4)) + ((float64((v3 + float32(0.0))) / (v1 + (float64(((-2.25 + 3.14159) + ((v3 + v3) - 1e10))) - v1))) / v1)) - 1e10)
                                                        var v58 int32 = v22
                                                        v59 := GenS6{F0: 1, F1: 42}
                                                        _ = fi56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                }
                                                var v60 int8 = -1
                                                _ = fi48
                                                _ = v60
                                            }
                                            v61 := GenS7{F0: 299, F1: (v17 / ((v17) | 1)), F2: 224}
                                            var v62 string = "foo bar"
                                            _ = v61
                                            _ = v62
                                        }
                                    case 128:
                                        var v63 uint8 = 255
                                        _ = v63
                                    default:
                                        v64 := GenS4{F0: true, F1: v21, F2: 767}
                                        _ = v64
                                    }
                                }
                            default:
                                fmt.Println(v22, v42)
                            }
                        } else {
                            v65 := GenS4{F0: true, F1: (v21 / (((v21 * (v21 + v21))) | 1)), F2: (v22 << ((65535) & 63))}
                            fmt.Println(v17, v1, v4)
                            var v66 int64 = 218
                            _ = v65
                            _ = v66
                        }
                        switch ("" + "foo bar") {
                        case "":
                            var v67 uint = 255
                            _ = v67
                        case "a":
                            v68 := GenS8{F0: 42}
                            _ = v68
                        }
                        fmt.Println(v21, v22)
                    } else {
                        fmt.Println(v1, v22)
                    }
                    fmt.Println(v17)
                }
            case -1:
                v1 = v1
            case 4294967295:
                v22 = (v22 + (v22 / (((v22 - (^(v22)))) | 1)))
            }
        }
        _ = v42
    } else {
        var v69 uint64 = v17
        v3 -= float32(0.0)
        v21 = 80
        _ = v69
    }
    var v70 float64 = v1
    for fi71 := 0; fi71 < 15; fi71++ {
        if false {
            _ = len(v2)
            _ = len(v2)
        }
        var v72 int64 = genFn1(0)
        v22 += (v22 / (((v22 | v22)) | 1))
        _ = fi71
        _ = v72
    }
    v22 -= v22
    v73 := GenS9{F0: ("foo bar" + "foo bar"), F1: -466}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v73
}
