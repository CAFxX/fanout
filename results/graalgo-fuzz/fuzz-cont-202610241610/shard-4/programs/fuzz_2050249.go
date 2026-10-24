package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 uint16; F1 int32; F2 uint }
func genFn0() uint64 {
    v4 := GenS1{F0: 65535, F1: 1}
    if (!false) {
        v5 := []int{32768, 2147483647, 0}
        if false {
            var v6 string = ((("foo bar" + "a") + "a") + "hello")
            switch -2147483648 {
            case 1:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                if true {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        v8 := GenS2{F0: 226, F1: 0}
                        v9 := GenS1{F0: 42, F1: 2}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 string = "hello"
                        if true {
                            for fi12 := 0; fi12 < 18; fi12++ {
                                var v13 uint16 = 0
                                v5 = append(v5, (fi10 >> ((36) & 63)))
                                _ = fi12
                                _ = v13
                            }
                        } else {
                            v11 = (v11 + v6)
                            v6 = v11
                            v14 := GenS3{F0: 42, F1: -128, F2: 0}
                            _ = v14
                        }
                        _ = fi10
                        _ = v11
                    }
                }
            }
            v15 := GenS2{F0: 0, F1: 665}
            _ = v15
            _ = v6
        }
        _ = v5
    } else {
        var v16 int = 798
        fmt.Println(v16)
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := []int{(fi17 - fi17), 128, fi17, -2147483648}
        var v19 uint64 = 42
        if true {
            fmt.Println(fi17)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint8 = 42
    for fi21 := 0; fi21 < 7; fi21++ {
        switch fi21 {
        case 437:
            var v22 bool = (3.14159 < float32((0.1 - 3.14159)))
            _ = v22
        }
        _ = fi21
    }
    _ = v20
    _ = v4
    return 2
}
func genFn1(p0 uint32, p1 int32) int16 {
    var v23 int16 = -800
    for fi24 := 0; fi24 < 2; fi24++ {
        if false {
            var v25 float32 = 0.1
            if false {
                fmt.Println(fi24)
                v23 += (-(v23))
            }
            _ = v25
        }
        _ = fi24
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int32 = (-(p1))
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v23
    return (v23 + (^(v23)))
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{0, 0, -2147483648, 914}
    var v3 int16 = 164
    v3 = (-(v3))
    if true {
        for fi28 := 0; fi28 < 5; fi28++ {
            v1[fi28] = (-(fi28))
            _ = fi28
        }
        fmt.Println(v3)
    }
    v29 := GenS4{F0: -9223372036854775808}
    v3 = v3
    if (false && (4294967295 <= 899)) {
        switch -1 {
        case -585:
            v30 := GenS2{F0: 213, F1: 18446744073709551615}
            _ = v30
        case -1:
            switch "" {
            case "hello":
                v31 := GenS5{F0: 1}
                _ = v31
            case "a":
                if true {
                    fmt.Println(v3)
                    fmt.Println(v3)
                    if true {
                        fmt.Println(v3)
                        _ = len(v2)
                        v3 -= (-(v3))
                    }
                } else {
                    v3 -= v3
                }
            default:
                var v32 int = 32768
                _ = v32
            }
        case 127:
            v33 := v1[127]
            _ = v33
        default:
            if true {
                var v34 uint = 2
                fmt.Println(v34)
                _ = v34
            }
        }
        v3 = v3
        var v35 int8 = 53
        _ = v35
    } else {
        var v36 int32 = int32(v3)
        for fi37 := 0; fi37 < 20; fi37++ {
            _ = len(v1)
            if ((v36 == 42) || false) {
                if false {
                    if (1 == fi37) {
                        v1[(-(fi37))] = (-(fi37))
                    } else {
                        switch ("foo bar" + (("" + "foo bar") + "\n\t")) {
                        case "\n\t":
                            var v38 int16 = (v3 / ((genFn1(1, (v36 / (((v36 / (((v36 & v36)) | 1))) | 1)))) | 1))
                            _ = v38
                        }
                        _ = len(v2)
                        if ((!false) && false) {
                            var v39 uint = 2
                            if len(v2) > 0 {
                                v2[0] = 32767
                            }
                            _ = v39
                        }
                    }
                    if (1 < -128) {
                        v36 += v36
                    }
                }
                if false {
                    v40 := GenS6{F0: 793, F1: 4294967295, F2: 2}
                    v41 := GenS7{F0: 42}
                    _ = v40
                    _ = v41
                }
                v36 = v36
            }
            if true {
                var v42 uint = 0
                _ = v42
            } else {
                for fi43 := 0; fi43 < 8; fi43++ {
                    v3 = (v3 | (v3 | v3))
                    _ = fi43
                }
            }
            _ = fi37
        }
        _ = v36
    }
    var v44 int8 = 100
    if len(v2) > 0 {
        v2[0] = 882
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 int = 42
        _ = fi45
        _ = v46
    }
    v47 := GenS4{F0: 487}
    switch "\n\t" {
    case "a":
        for fi48 := 0; fi48 < 3; fi48++ {
            var v49 bool = false
            _ = fi48
            _ = v49
        }
    case "foo bar":
        v50 := GenS2{F0: 128, F1: 18446744073709551615}
        _ = v50
    default:
        if false {
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 uint64 = genFn0()
                for fi53 := 0; fi53 < 12; fi53++ {
                    var v54 float64 = 0.1
                    _ = fi53
                    _ = v54
                }
                switch fi51 {
                case 1:
                    _ = len(v2)
                case 204:
                    if false {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v52, v52, v52)
                        case "\n\t":
                            var v55 int64 = 915
                            _ = v55
                        }
                    } else {
                        var v56 string = "a"
                        switch v56 {
                        case "a":
                            _ = len(v1)
                        case "\n\t":
                            var v57 int64 = -1
                            _ = v57
                        default:
                            var v58 float32 = (((1e10 - 1.5) + (float32(-2.25) + (1.5 / -2.25))) - 1.5)
                            _ = v58
                        }
                        _ = v56
                    }
                case 42:
                    switch fi51 {
                    case -612:
                        var v59 int64 = 0
                        _ = v59
                    case 128:
                        switch fi51 {
                        case 32768:
                            var v60 uint32 = 712
                            _ = v60
                        case -2147483648:
                            var v61 int16 = v3
                            _ = v61
                        case 434:
                            var v62 int8 = v44
                            _ = v62
                        default:
                            switch -224 {
                            case 34:
                                var v63 int8 = -128
                                _ = v63
                            }
                        }
                    case -1:
                        switch "" {
                        case "a":
                            if false {
                                if (-1 == 2147483647) {
                                    v52 += v52
                                    v3 = v3
                                    for fi64 := 0; fi64 < 17; fi64++ {
                                        var v65 int64 = 9223372036854775807
                                        v44 -= (v44 << ((65535) & 63))
                                        var v66 int16 = (-(v3))
                                        if (0 >= 1) {
                                            fmt.Println(v65, fi64)
                                        }
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    v67 := GenS8{F0: true, F1: ("" + "foo bar")}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v3 -= genFn1(1, 84)
                                    _ = v67
                                }
                                fmt.Println(fi51, v52)
                                v3 = genFn1(65535, 788)
                            } else {
                                if true {
                                    switch fi51 {
                                    case -2147483648:
                                        var v68 int32 = 2147483647
                                        _ = v68
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                } else {
                                    var v69 uint16 = 1
                                    _ = v69
                                }
                                fmt.Println(v3)
                            }
                        case "":
                            v2 = append(v2, (fi51 | fi51))
                        }
                    }
                default:
                    if (!true) {
                        var v70 int16 = int16(v52)
                        _ = v70
                    } else {
                        v52 = genFn0()
                    }
                }
                _ = fi51
                _ = v52
            }
        } else {
            if len(v2) > 0 {
                v2[0] = -586
            }
        }
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        if ((true && false) || (1.5 == (float64((0.1 + (float32(0.1) + -2.25))) - 0.1))) {
            v72 := GenS9{F0: 0, F1: 2147483647, F2: 1}
            v2 = append(v2, int(v44))
            _ = v72
        } else {
            for fi73 := 0; fi73 < 5; fi73++ {
                v44 = (v44 >> ((345) & 63))
                _ = fi73
            }
            v3 += (-(v3))
            for fi74 := 0; fi74 < 3; fi74++ {
                var v75 int8 = int8(v3)
                var v76 int32 = -302
                v77 := GenS7{F0: 4294967295}
                var v78 float32 = (1.5 * (0.0 / 0.1))
                _ = fi74
                _ = v75
                _ = v76
                _ = v77
                _ = v78
            }
        }
        if ("hello" < "a") {
            var v79 float32 = -2.25
            _ = v79
        } else {
            v80 := v1[(fi71 ^ 32768)]
            if (-536 >= 42) {
                v2 = append(v2, 2147483647)
            }
            switch "\n\t" {
            case "hello":
                switch 127 {
                case 478:
                    if (!(2 <= 1)) {
                        var v81 int8 = (v44 - (v44 >> ((1) & 63)))
                        _ = v81
                    }
                case 220:
                    v3 = (v3 / (((^(v3))) | 1))
                case 0:
                    if true {
                        var v82 uint = 798
                        _ = v82
                    }
                }
            }
            _ = v80
        }
        if true {
            if false {
                v44 += -1
                fmt.Println(v44, fi71, fi71)
                if false {
                    var v83 uint16 = 32767
                    var v84 uint32 = 42
                    _ = v83
                    _ = v84
                } else {
                    var v85 float64 = 3.14159
                    _ = v85
                }
            }
        } else {
            if true {
                var v86 uint8 = 1
                v44 -= -89
                _ = v86
            }
        }
        _ = fi71
    }
    var v87 int32 = -2147483648
    fmt.Println(v3)
    v88 := GenS3{F0: v3, F1: (v44 / ((v44) | 1)), F2: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v44
    _ = v47
    _ = v87
    _ = v88
}
