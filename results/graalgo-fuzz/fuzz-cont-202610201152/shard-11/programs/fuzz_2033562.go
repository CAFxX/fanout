package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS5 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int }
type GenS9 struct { F0 string; F1 uint32; F2 bool }
func genFn0() string {
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = 0.1
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 float32 = -2.25
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v5 = v5
            _ = fi9
        }
        _ = fi6
    }
    if (1000 < 1) {
        v10 := GenS1{F0: 1}
        if false {
            v11 := GenS1{F0: 131}
            switch "" {
            case "\n\t":
                if (!(true || false)) {
                    var v12 bool = (1000 == 42)
                    var v13 float64 = (1e10 / -2.25)
                    _ = v12
                    _ = v13
                }
            }
            _ = v11
        } else {
            v14 := GenS2{F0: 0}
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    v5 = 0.1
                    switch "\n\t" {
                    case "foo bar":
                        v17 := []int{(^(fi15)), 709, (fi16 / ((1) | 1)), (^(fi15))}
                        _ = v17
                    case "\n\t":
                        v5 -= v5
                    }
                    if false {
                        v5 = (float32(float64(-2.25)) * (v5 * v5))
                    }
                    _ = fi16
                }
                var v18 int8 = 127
                _ = fi15
                _ = v18
            }
            _ = v14
        }
        _ = v10
    } else {
        v19 := []int{1, -1}
        var v20 int64 = 0
        if false {
            var v21 float64 = 1e10
            switch "hello" {
            case "\n\t":
                v20 = v20
            case "":
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS2{F0: 255}
                    _ = fi22
                    _ = v23
                }
            default:
                var v24 int16 = -1
                _ = v24
            }
            _ = len(v19)
            _ = v21
        } else {
            if true {
                v5 = ((float32((((0.0 - float64(v5)) - 0.0) * float64((v5 + v5)))) / float32(0.1)) - (((v5 - v5) / v5) * (v5 * v5)))
                for fi25 := 0; fi25 < 2; fi25++ {
                    fmt.Println(fi25)
                    if (18446744073709551615 >= 0) {
                        var v26 int64 = v20
                        v26 = v20
                        _ = v26
                    }
                    _ = fi25
                }
                if ((!(32767 == 964)) && false) {
                    for fi27 := 0; fi27 < 3; fi27++ {
                        if (false && (4294967295 < 153)) {
                            if true {
                                var v28 uint32 = 2147483647
                                _ = v28
                            } else {
                                fmt.Println(v5)
                                if len(v19) > 0 {
                                    v19[0] = (fi27 * -9223372036854775808)
                                }
                            }
                            switch -1 {
                            case -1:
                                v20 = (v20 - 42)
                            case -556:
                                fmt.Println(fi27, v20)
                            }
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v20 -= v20
                            case "hello":
                                v29 := GenS3{F0: (241 <= 127), F1: 532, F2: 107}
                                _ = v29
                            default:
                                v5 += v5
                            }
                            switch "a" {
                            case "":
                                var v30 float64 = 1e10
                                _ = v30
                            }
                        }
                        v31 := GenS4{F0: 32767, F1: 32767, F2: 1}
                        _ = fi27
                        _ = v31
                    }
                    v20 -= (v20 << ((77) & 63))
                }
            }
            if false {
                if true {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        switch (fi32 & -539) {
                        case -2147483648:
                            var v33 int32 = -2147483648
                            _ = v33
                        case 0:
                            v34 := GenS5{F0: int32(fi32), F1: 65535, F2: 0}
                            _ = v34
                        case 4294967295:
                            fmt.Println(v20, v5)
                        }
                        if (65535 == 2) {
                            if (true || ((294 <= 317) && false)) {
                                fmt.Println(v5, fi32, v5)
                            }
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    var v35 uint = 65535
                                    v36 := GenS6{F0: 2147483647}
                                    _ = v35
                                    _ = v36
                                } else {
                                    if (true && false) {
                                        fmt.Println(v5, fi32)
                                        v5 -= ((v5 + 1e10) - v5)
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v5 += 1e10
                                    }
                                    v20 = -9223372036854775808
                                }
                            case "\n\t":
                                var v37 int16 = 32767
                                _ = v37
                            case "":
                                v38 := GenS7{F0: -248, F1: 486}
                                _ = v38
                            default:
                                _ = len(v19)
                            }
                        } else {
                            v20 += v20
                            if false {
                                if (!true) {
                                    v39 := GenS8{F0: 1, F1: 8, F2: (fi32 + (fi32 % (((fi32 * 1)) | 1)))}
                                    if false {
                                        v20 -= (v20 | v20)
                                        if len(v19) > 0 {
                                            v19[0] = (fi32 - fi32)
                                        }
                                    } else {
                                        v5 -= 0.1
                                        for fi40 := 0; fi40 < 11; fi40++ {
                                            v41 := GenS1{F0: 127}
                                            _ = fi40
                                            _ = v41
                                        }
                                    }
                                    var v42 float32 = v5
                                    _ = v39
                                    _ = v42
                                } else {
                                    v20 = (v20 - (^(v20)))
                                    if true {
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            v43 := GenS9{F0: "a", F1: 0, F2: false}
                                            _ = v43
                                        case "\n\t":
                                            for fi44 := 0; fi44 < 10; fi44++ {
                                                if false {
                                                    v20 -= v20
                                                    var v45 int8 = 127
                                                    v45 -= 3
                                                    _ = v45
                                                }
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                var v46 uint8 = 20
                                                _ = fi44
                                                _ = v46
                                            }
                                        default:
                                            _ = len(v19)
                                        }
                                    }
                                    if false {
                                        v20 = -1
                                        v47 := GenS1{F0: 35}
                                        _ = v47
                                    } else {
                                        fmt.Println(v20, v20, v20)
                                    }
                                }
                            } else {
                                for fi48 := 0; fi48 < 20; fi48++ {
                                    fmt.Println(v5, fi32, fi48)
                                    v20 -= v20
                                    var v49 uint8 = 42
                                    _ = fi48
                                    _ = v49
                                }
                                v50 := GenS8{F0: 0, F1: 0, F2: 1}
                                switch ("a" + "foo bar") {
                                case "foo bar":
                                    var v51 uint32 = 65535
                                    _ = v51
                                }
                                _ = v50
                            }
                            fmt.Println(v5, v5, v20)
                        }
                        _ = fi32
                    }
                    v5 -= v5
                    v5 = float32(0.1)
                } else {
                    v20 = (v20 + (v20 * v20))
                    v52 := GenS1{F0: 207}
                    _ = v52
                }
            }
            if true {
                switch "hello" {
                case "":
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        var v53 uint16 = 1
                        _ = v53
                    case "hello":
                        var v54 float32 = -2.25
                        _ = v54
                    case "":
                        v20 = 0
                    default:
                        _ = len(v19)
                    }
                }
                fmt.Println(v20)
            }
        }
        _ = v19
        _ = v20
    }
    _ = v5
    return ("foo bar" + (("" + ("\n\t" + "")) + ("a" + "")))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 uint16 = 201
    v4 := make(map[int]uint32)
    v2 += v2
    var v55 string = "foo bar"
    var v56 uint = 657
    fmt.Println(v55, v2, v3)
    var v57 int8 = -128
    v55 = v55
    v57 = (-(v57))
    var v58 float32 = ((1e10 - 0.1) - float32(0.0))
    v59 := []int{-2147483648, 4294967295}
    if false {
        var v60 bool = (!false)
        fmt.Println(v58, v55, v56)
        fmt.Println(v56, v58)
        _ = v60
    } else {
        if len(v59) > 0 {
            v59[0] = 818
        }
    }
    switch -2147483648 {
    case 32768:
        switch int(v57) {
        case 128:
            var v61 string = ((((("hello" + v55) + ((v55 + v55) + "hello")) + v55) + (((v55 + (v55 + "foo bar")) + (v55 + ("\n\t" + v55))) + v55)) + (v55 + ("" + "foo bar")))
            _ = v61
        case -2147483648:
            fmt.Println(v55, v58, v3)
        }
    case 475:
        var v62 uint8 = 75
        _ = v62
    }
    v4[1] = 2
    if true {
        var v63 int16 = 567
        _ = v63
    }
    if len(v59) >= 2 {
        v59 = v59[1:len(v59)]
    }
    switch v55 {
    case "\n\t":
        v57 -= v57
    case "foo bar":
        if ((v2 - v2) == -364) {
            var v64 int16 = 32767
            _ = v64
        } else {
            v58 += v58
            v65 := GenS3{F0: false, F1: 0, F2: 1000}
            _ = v65
        }
    }
    v3 -= (v3 / (((v3 / ((65535) | 1))) | 1))
    fmt.Println(v3, v2, v56)
    if ((!(-1 == 665)) && false) {
        var v66 uint64 = 2
        if len(v59) > 0 {
            v59[0] = int(v56)
        }
        _ = v66
    } else {
        fmt.Println(v3, v56)
        _ = len(v59)
        for fi67 := 0; fi67 < 8; fi67++ {
            switch int(v2) {
            case -823:
                switch "foo bar" {
                case "\n\t":
                    switch 195 {
                    case 2147483647:
                        var v68 int16 = 1000
                        _ = v68
                    }
                default:
                    v56 = v56
                }
            case -302:
                fmt.Println(v58)
            }
            var v69 uint16 = 65535
            for fi70 := 0; fi70 < 15; fi70++ {
                v56 = (v56 & v56)
                _ = fi70
            }
            _ = fi67
            _ = v69
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
