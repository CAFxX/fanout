package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 string; F2 int }
type GenS5 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS6 struct { F0 int; F1 int16; F2 uint64 }
func genFn0(p0 string, p1 int64, p2 uint32) uint64 {
    if ("a" < p0) {
        v6 := GenS1{F0: (p2 % (((^(p2))) | 1)), F1: false, F2: 180}
        p0 = (p0 + "foo bar")
        var v7 uint32 = (p2 % (((p2 & (p2 ^ p2))) | 1))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (^(p2)), F1: true, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return uint64(p2)
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{-1, 127}
    v3 := make(map[string]string)
    var v4 int32 = 1
    v5 := make(map[int]string)
    if false {
        var v9 float64 = (1e10 * -2.25)
        _ = v9
    }
    switch "foo bar" {
    case "":
        fmt.Println(v4)
    case "foo bar":
        v4 = -42
    case "\n\t":
        var v10 uint16 = 2
        _ = v10
    }
    fmt.Println(v4)
    var v11 int16 = 586
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: 0, F1: (fi12 >= (fi12 / ((fi12) | 1))), F2: v11}
        _ = fi12
        _ = v13
    }
    var v14 uint = 4294967295
    if false {
        var v15 int8 = -97
        _ = v15
    } else {
        v11 = v11
        v4 = (v4 >> ((2) & 63))
        var v16 int8 = -1
        _ = v16
    }
    if ((((-2.25 * 0.1) * 0.1) - 3.14159) > 1e10) {
        if true {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "" {
                case "\n\t":
                    var v18 string = "a"
                    _ = v18
                }
                if false {
                    fmt.Println(v14, v11, v4)
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if true {
                            v20 := GenS1{F0: 387, F1: (!true), F2: v11}
                            v21 := GenS1{F0: 2147483648, F1: (true || (1 < 1)), F2: (-(v11))}
                            _ = v20
                            _ = v21
                        }
                        if true {
                            v14 += (v14 / ((42) | 1))
                            v14 = (v14 | 1)
                        } else {
                            switch ("foo bar" + ("\n\t" + "\n\t")) {
                            case "\n\t":
                                switch "\n\t" {
                                case "\n\t":
                                    var v22 uint = v14
                                    _ = v22
                                case "hello":
                                    if (0.0 > -2.25) {
                                        if (42 >= 149) {
                                            fmt.Println(fi17, fi17, fi17)
                                            var v23 uint = v14
                                            _ = v23
                                        } else {
                                            v14 = (v14 / ((v14) | 1))
                                            v24 := GenS1{F0: 2147483648, F1: false, F2: -426}
                                            var v25 int = fi17
                                            _ = v24
                                            _ = v25
                                        }
                                        fmt.Println(fi19, v4)
                                        _ = len(v1)
                                    } else {
                                        var v26 bool = false
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if false {
                                            var v27 string = "\n\t"
                                            var v28 float32 = float32(0.1)
                                            var v29 uint32 = 532
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int32 = -937
                                            var v31 int = int(v11)
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                if (v14 > v14) {
                                                    var v32 int16 = (v11 + v11)
                                                    switch v31 {
                                                    case 9223372036854775807:
                                                        v4 = v30
                                                    case 127:
                                                        if (v26 && v26) {
                                                            var v33 int8 = 0
                                                            v30 += -607
                                                            fmt.Println(v4, v32)
                                                            _ = v33
                                                        }
                                                    default:
                                                        switch "hello" {
                                                        case "":
                                                            var v34 uint64 = 623
                                                            _ = v34
                                                        default:
                                                            v4 += (v30 ^ (v30 / ((v30) | 1)))
                                                        }
                                                    }
                                                    var v35 bool = (3.14159 >= float32(1e10))
                                                    _ = v32
                                                    _ = v35
                                                } else {
                                                    fmt.Println(fi19, v11, v30)
                                                }
                                            case "hello":
                                                v3["hello"] = "a"
                                            }
                                            if v26 {
                                                var v36 bool = true
                                                if (234 <= 128) {
                                                    v36 = ((-102 != 100) && (346 != genFn0(("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))), int64(v14), 0)))
                                                    v37 := GenS2{F0: (fi19 ^ fi19), F1: 2, F2: 2}
                                                    var v38 int64 = 848
                                                    _ = v37
                                                    _ = v38
                                                }
                                                _ = v36
                                            } else {
                                                v39 := GenS1{F0: 655, F1: false, F2: 1}
                                                _ = v39
                                            }
                                            v40 := GenS1{F0: 2147483648, F1: v26, F2: (-(v11))}
                                            _ = v30
                                            _ = v31
                                            _ = v40
                                        }
                                        fmt.Println(v11)
                                        _ = v26
                                    }
                                    switch ("\n\t" + ((("" + "foo bar") + "\n\t") + "hello")) {
                                    case "\n\t":
                                        v4 = -2147483648
                                    case "a":
                                        v41 := GenS2{F0: (^(fi19)), F1: 42, F2: 42}
                                        _ = v41
                                    case "foo bar":
                                        var v42 int16 = -932
                                        v42 = (v11 + (v42 % ((0) | 1)))
                                        _ = v42
                                    }
                                    v43 := GenS1{F0: 1, F1: (4294967295 <= 2147483648), F2: v11}
                                    _ = v43
                                case "a":
                                    v44 := GenS2{F0: int(v14), F1: 18446744073709551615, F2: 2}
                                    _ = v44
                                default:
                                    v45 := GenS3{F0: -1}
                                    _ = v45
                                }
                            case "a":
                                switch 0 {
                                case -343:
                                    v11 -= v11
                                case -721:
                                    var v46 int = (fi19 % (((fi17 & fi17)) | 1))
                                    _ = v46
                                case 962:
                                    if false {
                                        v11 = (v11 % ((v11) | 1))
                                        v11 = 42
                                    } else {
                                        v11 = v11
                                        v2 = append(v2, fi17)
                                        v14 -= 42
                                    }
                                }
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    var v47 float32 = -2.25
                                    _ = v47
                                }
                            }
                            v11 -= (v11 / ((v11) | 1))
                        }
                        _ = fi19
                    }
                } else {
                    fmt.Println(v14)
                }
                _ = fi17
            }
        } else {
            switch 1 {
            case 511:
                fmt.Println(v14, v14, v4)
            default:
                for fi48 := 0; fi48 < 18; fi48++ {
                    v4 += (v4 - v4)
                    var v49 uint8 = 1
                    var v50 string = "a"
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
            var v51 string = "\n\t"
            _ = v51
        }
        switch -851 {
        case 969:
            v52 := v5[-1]
            _ = v52
        }
    } else {
        fmt.Println(v4, v4, v11)
    }
    switch 128 {
    case -2147483648:
        var v53 bool = (((3.14159 / 1.5) * (-2.25 - (float32(3.14159) + 0.1))) != 3.14159)
        _ = v53
    }
    var v54 bool = (!(!false))
    for fi55 := 0; fi55 < 7; fi55++ {
        v56 := GenS4{F0: v14, F1: ("foo bar" + (("a" + ("hello" + "")) + (("" + "") + ((((("a" + ("\n\t" + ("hello" + ""))) + "\n\t") + ("" + "\n\t")) + "a") + "")))), F2: 127}
        var v57 int = (fi55 % ((fi55) | 1))
        v54 = false
        _ = fi55
        _ = v56
        _ = v57
    }
    v58 := GenS5{F0: 1, F1: 113, F2: v11}
    if (!(!true)) {
        v14 -= v14
        for fi59 := 0; fi59 < 1; fi59++ {
            v4 -= 42
            var v60 int32 = (v4 ^ (v4 + (v4 + 0)))
            var v61 float32 = (0.1 + 0.1)
            _ = fi59
            _ = v60
            _ = v61
        }
        if v54 {
            for fi62 := 0; fi62 < 10; fi62++ {
                if (-128 <= 1) {
                    v11 = (v11 << (((v14 & v14)) & 63))
                    for fi63 := 0; fi63 < 12; fi63++ {
                        var v64 string = ("a" + "hello")
                        _ = len(v1)
                        _ = fi63
                        _ = v64
                    }
                } else {
                    v11 = (v11 ^ 0)
                }
                _ = fi62
            }
        } else {
            switch -753 {
            case -940:
                var v65 float64 = (((1e10 * (-2.25 + 3.14159)) * (0.1 / ((0.1 - -2.25) + (((-2.25 - 0.0) + ((1.5 - 0.0) * ((0.1 - 0.1) - float64((float32(((((float64(0.0) - -2.25) / -2.25) + 1e10) * (0.1 + (1e10 + (1e10 * -2.25))))) + 1e10))))) / 1.0)))) / 1e10)
                _ = v65
            case 127:
                var v66 int = 42
                _ = v66
            case 2147483647:
                switch 0 {
                case 128:
                    v11 -= 42
                default:
                    fmt.Println(v14, v11)
                }
            }
            switch "" {
            case "":
                fmt.Println(v14)
            case "\n\t":
                var v67 int32 = v4
                _ = v67
            case "foo bar":
                v54 = v54
            }
        }
    } else {
        for fi68 := 0; fi68 < 8; fi68++ {
            v54 = (("a" + "a") > "")
            var v69 int = 42
            _ = fi68
            _ = v69
        }
        for fi70 := 0; fi70 < 10; fi70++ {
            var v71 string = ("foo bar" + (("" + "a") + "\n\t"))
            for fi72 := 0; fi72 < 2; fi72++ {
                v73 := GenS2{F0: (fi72 ^ (^(fi70))), F1: 843, F2: 255}
                _ = fi72
                _ = v73
            }
            _ = fi70
            _ = v71
        }
    }
    switch ("" + "\n\t") {
    case "hello":
        for fi74 := 0; fi74 < 14; fi74++ {
            for fi75 := 0; fi75 < 10; fi75++ {
                if v54 {
                    switch "" {
                    case "\n\t":
                        var v76 uint = 42
                        _ = v76
                    default:
                        v54 = false
                        v77 := GenS6{F0: fi74, F1: (v11 + v11), F2: 0}
                        _ = v77
                    }
                }
                _ = fi75
            }
            v11 += v11
            v11 -= (v11 - 1)
            _ = fi74
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v58
}
