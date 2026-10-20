package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int16; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint; F2 bool }
type GenS8 struct { F0 int8 }
func genFn0(p0 int16) int32 {
    var v6 int = 172
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS1{F0: "\n\t"}
            _ = fi9
            _ = v10
        }
        var v11 float32 = 0.0
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 3.14159
    var v3 uint64 = 884
    var v4 int16 = -973
    var v5 uint16 = (v1 & v1)
    if false {
        v12 := GenS1{F0: "hello"}
        fmt.Println(v1, v3)
        _ = v12
    }
    var v13 string = ((("a" + ("\n\t" + "a")) + ("foo bar" + "")) + ("\n\t" + ("\n\t" + "\n\t")))
    if (false || (!(0 > (v4 % (((v4 - (v4 / ((v4) | 1)))) | 1))))) {
        var v14 uint64 = v3
        if (true && false) {
            v15 := []int{-957, 32767}
            switch "a" {
            case "a":
                if false {
                    if false {
                        if len(v15) > 0 {
                            v15[0] = 0
                        }
                    } else {
                        var v16 int64 = -844
                        var v17 uint8 = 31
                        v18 := GenS1{F0: (v13 + v13)}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 string = "a"
                    v20 := GenS1{F0: "hello"}
                    _ = v19
                    _ = v20
                }
            default:
                var v21 int32 = genFn0((-(v4)))
                _ = v21
            }
            v22 := GenS2{F0: v14, F1: -963, F2: 2147483647}
            _ = v15
            _ = v22
        } else {
            v23 := GenS1{F0: v13}
            var v24 int32 = -1
            if true {
                var v25 int = 42
                _ = v25
            } else {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := []int{}
                    _ = fi26
                    _ = v27
                }
            }
            _ = v23
            _ = v24
        }
        if (!(true || true)) {
            fmt.Println(v13)
            v5 += v1
            for fi28 := 0; fi28 < 12; fi28++ {
                if (v4 >= v4) {
                    if (false || (1 <= 0)) {
                        v29 := []int{}
                        switch "" {
                        case "foo bar":
                            if (888 > 776) {
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    var v31 int64 = 1
                                    v32 := GenS2{F0: (v14 & 42), F1: genFn0(v4), F2: fi28}
                                    v4 = v4
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                if true {
                                    switch (fi28 * (fi28 | (fi28 / (((fi28 - (fi28 ^ (fi28 % (((fi28 - (fi28 / ((-2147483648) | 1)))) | 1))))) | 1)))) {
                                    case -849:
                                        v33 := GenS2{F0: v14, F1: -2147483648, F2: (fi28 * fi28)}
                                        _ = v33
                                    case 1:
                                        v34 := GenS3{F0: 608, F1: 18446744073709551615}
                                        _ = v34
                                    case 0:
                                        var v35 uint = 266
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 12; fi36++ {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            v4 += v4
                                            _ = fi36
                                        }
                                    }
                                }
                            }
                        case "\n\t":
                            for fi37 := 0; fi37 < 5; fi37++ {
                                switch (fi28 * fi37) {
                                case 42:
                                    var v38 int64 = -78
                                    _ = v38
                                case 32768:
                                    if (uint8(v14) >= 9) {
                                        if false {
                                            var v39 int8 = 100
                                            switch -9223372036854775808 {
                                            case 4294967295:
                                                v40 := GenS1{F0: ((("\n\t" + "") + "\n\t") + v13)}
                                                _ = v40
                                            case -1:
                                                v29 = append(v29, 1)
                                            }
                                            _ = v39
                                        }
                                    } else {
                                        fmt.Println(v13)
                                        var v41 string = v13
                                        _ = v41
                                    }
                                case -2147483648:
                                    switch (^(fi28)) {
                                    case 0:
                                        var v42 int8 = -26
                                        _ = v42
                                    case -9223372036854775808:
                                        if (v2 <= float64((0.0 - 1.5))) {
                                            v43 := GenS2{F0: (v14 ^ (v14 / ((v3) | 1))), F1: genFn0(42), F2: -303}
                                            _ = v43
                                        } else {
                                            var v44 bool = true
                                            if len(v29) > 0 {
                                                v29[0] = (fi28 - fi28)
                                            }
                                            _ = v44
                                        }
                                    }
                                default:
                                    v45 := GenS2{F0: v14, F1: genFn0(1000), F2: (fi28 >> ((857) & 63))}
                                    _ = v45
                                }
                                _ = fi37
                            }
                        default:
                            v1 += (v5 * v1)
                        }
                        var v46 float64 = v2
                        _ = v29
                        _ = v46
                    } else {
                        v3 += 555
                    }
                    var v47 bool = (128 != 127)
                    v48 := GenS1{F0: ("hello" + (v13 + v13))}
                    _ = v47
                    _ = v48
                }
                var v49 uint32 = 42
                _ = fi28
                _ = v49
            }
        } else {
            if (true || false) {
                v50 := []int{-9223372036854775808, 9223372036854775807}
                _ = v50
            }
            if ((v5 % ((v5) | 1)) < 53) {
                var v51 uint64 = v14
                v51 -= v3
                _ = v51
            } else {
                fmt.Println(v2)
                v52 := GenS2{F0: v14, F1: -1, F2: -219}
                _ = v52
            }
            if true {
                v1 += (v5 + (v5 & 956))
                for fi53 := 0; fi53 < 8; fi53++ {
                    v54 := GenS1{F0: (v13 + "foo bar")}
                    if false {
                        var v55 bool = (true || false)
                        var v56 int16 = 42
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS4{F0: (^(v14))}
                    _ = fi53
                    _ = v54
                    _ = v57
                }
                switch 0 {
                case 1:
                    var v58 int = 1
                    _ = v58
                case 42:
                    if ("a" <= ((v13 + (((((("hello" + (v13 + v13)) + "a") + v13) + v13) + (v13 + v13)) + v13)) + (v13 + v13))) {
                        for fi59 := 0; fi59 < 1; fi59++ {
                            switch fi59 {
                            case 128:
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    v61 := GenS5{F0: (^(v5)), F1: v4, F2: v13}
                                    _ = fi60
                                    _ = v61
                                }
                            case 2147483647:
                                fmt.Println(v14)
                            case 1:
                                v62 := []int{fi59, fi59, 4294967295}
                                _ = v62
                            }
                            v63 := GenS2{F0: 42, F1: -2147483648, F2: fi59}
                            v13 = v13
                            _ = fi59
                            _ = v63
                        }
                    } else {
                        fmt.Println(v2, v3)
                        for fi64 := 0; fi64 < 18; fi64++ {
                            var v65 uint8 = 88
                            _ = fi64
                            _ = v65
                        }
                    }
                case 2147483647:
                    v66 := GenS6{F0: v3}
                    _ = v66
                default:
                    var v67 uint64 = 2
                    _ = v67
                }
            } else {
                if true {
                    var v68 uint16 = v1
                    var v69 uint8 = 1
                    for fi70 := 0; fi70 < 8; fi70++ {
                        v71 := GenS4{F0: (v3 * (^(v3)))}
                        if false {
                            switch "a" {
                            case "foo bar":
                                var v72 float64 = 0.0
                                _ = v72
                            case "hello":
                                switch fi70 {
                                case -1:
                                    for fi73 := 0; fi73 < 19; fi73++ {
                                        v74 := GenS7{F0: 623, F1: 255, F2: (true || false)}
                                        _ = fi73
                                        _ = v74
                                    }
                                case -406:
                                    v75 := []int{(fi70 & fi70)}
                                    _ = v75
                                case 9223372036854775807:
                                    if (true || false) {
                                        switch (v13 + v13) {
                                        case "hello":
                                            v76 := []int{int(v69), fi70, (fi70 ^ (fi70 | 842))}
                                            _ = v76
                                        default:
                                            v2 = ((float64(1e10) - (v2 * v2)) - (v2 / 1.5))
                                        }
                                    }
                                }
                            }
                            for fi77 := 0; fi77 < 16; fi77++ {
                                fmt.Println(fi77, v69, fi77)
                                _ = fi77
                            }
                            v1 = v1
                        }
                        _ = fi70
                        _ = v71
                    }
                    _ = v68
                    _ = v69
                } else {
                    switch (v13 + (v13 + v13)) {
                    case "a":
                        v78 := GenS2{F0: v14, F1: 1, F2: 1}
                        _ = v78
                    default:
                        v79 := GenS8{F0: -81}
                        _ = v79
                    }
                    var v80 float64 = float64(((0.1 - 0.0) - ((0.1 - 0.1) * 0.0)))
                    v81 := []int{0, -9223372036854775808}
                    _ = v80
                    _ = v81
                }
                v82 := GenS5{F0: (v1 & 1), F1: v4, F2: v13}
                _ = v82
            }
        }
        _ = v14
    }
    v83 := []int{4294967295, -901, 32767, 0}
    if len(v83) >= 2 {
        v83 = v83[1:len(v83)]
    }
    var v84 string = v13
    v3 = (v3 * v3)
    if true {
        v83 = append(v83, 69)
        switch "" {
        case "foo bar":
            v85 := GenS7{F0: v4, F1: 617, F2: false}
            _ = v85
        case "hello":
            var v86 int64 = -1
            _ = v86
        case "a":
            if true {
                if len(v83) >= 2 {
                    v83 = v83[1:len(v83)]
                }
            }
        }
        fmt.Println(v13, v84, v2)
    } else {
        var v87 int32 = 380
        _ = v87
    }
    for fi88 := 0; fi88 < 11; fi88++ {
        v13 = "a"
        _ = fi88
    }
    fmt.Println(v3)
    v2 += float64(-2.25)
    var v89 bool = ((2 != 800) && (v3 < v3))
    var v90 int = 32767
    if v89 {
        for fi91 := 0; fi91 < 12; fi91++ {
            var v92 uint = 2
            var v93 bool = true
            _ = fi91
            _ = v92
            _ = v93
        }
        if len(v83) >= 2 {
            v83 = v83[1:len(v83)]
        }
    } else {
        fmt.Println(v5, v3, v13)
        var v94 uint8 = 40
        fmt.Println(v13)
        _ = v94
    }
    if true {
        v90 -= 317
    } else {
        if len(v83) > 0 {
            v83[0] = 192
        }
        v4 += (v4 ^ (v4 + v4))
    }
    if (("foo bar" >= "") || true) {
        v90 = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v84:", v84)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v89
    _ = v90
}
