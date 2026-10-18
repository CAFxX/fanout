package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint; F2 uint }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint32) bool {
    v5 := []int{-9223372036854775808, 1, 2147483647}
    switch 32768 {
    case 1:
        v6 := GenS1{F0: ((42 == 1) || (!(!true))), F1: "foo bar"}
        _ = v6
    case -1:
        var v7 int = -9223372036854775808
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (733 < 2147483647)
}
func genFn1(p0 uint8) uint8 {
    var v8 uint8 = p0
    if (((0.1 >= 3.14159) && false) || (true || ((1e10 - (3.14159 / 1.0)) >= (float32(0.0) - 0.1)))) {
        var v9 int = 42
        _ = v9
    }
    _ = p0
    _ = v8
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = 1
    var v3 string = v1
    var v4 int16 = -32768
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: (42 == 0), F1: v3}
        v12 := []int{}
        if (false && (3.14159 == (3.14159 - 1.5))) {
            if ((88 <= 949) || true) {
                var v13 uint16 = 909
                fmt.Println(v13, v2)
                _ = v13
            }
            v14 := GenS1{F0: false, F1: "hello"}
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    fmt.Println(v1, v1, v3)
    if (3.14159 == 3.14159) {
        if false {
            var v16 int64 = (v2 ^ -452)
            var v17 int = -9223372036854775808
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 uint16 = 0
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            var v20 int16 = v4
            if false {
                v21 := []int{652, 32768, 995}
                _ = v21
            }
            v4 = (v4 / ((v4) | 1))
            _ = v20
        }
        fmt.Println(v1, v3)
        var v22 uint32 = 2
        _ = v22
    } else {
        var v23 int16 = v4
        switch 127 {
        case -1:
            var v24 int16 = v23
            _ = v24
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 bool = (223 >= 2)
        _ = fi25
        _ = v26
    }
    if (!false) {
        fmt.Println(v4, v2, v1)
        switch "\n\t" {
        case "":
            fmt.Println(v3, v3)
        case "\n\t":
            v27 := GenS1{F0: (0.1 > ((1e10 * 0.0) / (0.1 / 1.5))), F1: (v3 + "\n\t")}
            _ = v27
        }
        var v28 uint64 = 4294967295
        _ = v28
    }
    v3 = "hello"
    v29 := []int{127}
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            _ = len(v29)
            v2 = v2
            var v32 int8 = -128
            _ = fi31
            _ = v32
        }
        if false {
            if len(v29) > 0 {
                v29[0] = 1
            }
            v33 := GenS2{F0: 822, F1: 887, F2: (v2 % (((v2 / ((v2) | 1))) | 1))}
            v34 := GenS3{F0: false}
            _ = v33
            _ = v34
        } else {
            var v35 uint8 = 2
            _ = v35
        }
        if (!(!((fi30 & (-(fi30))) == (fi30 ^ (-(fi30)))))) {
            for fi36 := 0; fi36 < 3; fi36++ {
                switch fi30 {
                case 32767:
                    v1 = v3
                case 391:
                    switch (fi36 ^ 4294967295) {
                    case 128:
                        v37 := GenS3{F0: (0 >= 0)}
                        _ = v37
                    case 32767:
                        if (-2147483648 != 1) {
                            v4 = v4
                            fmt.Println(fi36, v1)
                        }
                    default:
                        v29 = append(v29, (fi36 + fi36))
                    }
                case 128:
                    if true {
                        _ = len(v29)
                    }
                }
                fmt.Println(fi36, v3, v3)
                if false {
                    switch "hello" {
                    case "":
                        var v38 float64 = (((-2.25 / 1.5) + float64(1e10)) / 1.0)
                        _ = v38
                    case "foo bar":
                        var v39 uint8 = 1
                        _ = v39
                    case "\n\t":
                        if (159 >= 127) {
                            fmt.Println(fi30, v3, v4)
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                        } else {
                            switch (fi36 ^ fi30) {
                            case -1:
                                v3 = ((("\n\t" + (v1 + "")) + "a") + v1)
                            case -9223372036854775808:
                                v40 := GenS4{F0: (v4 << ((2) & 63)), F1: v3}
                                _ = v40
                            }
                            v2 = v2
                            v41 := GenS5{F0: v2}
                            _ = v41
                        }
                    default:
                        if len(v29) > 0 {
                            v29[0] = (-(fi36))
                        }
                    }
                    if false {
                        switch fi30 {
                        case -250:
                            var v42 uint = 1
                            _ = v42
                        }
                    }
                }
                if (-9223372036854775808 >= 1) {
                    v1 = ((v3 + v1) + v3)
                    v43 := GenS4{F0: (v4 & v4), F1: v1}
                    v44 := GenS6{F0: 4294967295, F1: 42, F2: 567}
                    _ = v43
                    _ = v44
                }
                _ = fi36
            }
            for fi45 := 0; fi45 < 3; fi45++ {
                v1 = v1
                _ = fi45
            }
            var v46 int8 = -65
            _ = v46
        } else {
            for fi47 := 0; fi47 < 8; fi47++ {
                var v48 int8 = -77
                _ = fi47
                _ = v48
            }
            fmt.Println(v4, v3)
        }
        _ = fi30
    }
    var v49 uint = 42
    _ = len(v29)
    v4 += 42
    v50 := GenS2{F0: 86, F1: 1, F2: (v2 & 9223372036854775807)}
    v4 += 1
    v2 += (v2 & v2)
    v2 += v2
    if true {
        v4 += 1000
        v1 = ((v3 + ((v3 + v3) + (v1 + v1))) + "\n\t")
        v51 := GenS4{F0: (v4 + (v4 >> ((v49) & 63))), F1: (v1 + v1)}
        _ = v51
    }
    switch (v3 + (v1 + v1)) {
    case "\n\t":
        if (0.1 >= 0.0) {
            if len(v29) > 0 {
                v29[0] = 32767
            }
            for fi52 := 0; fi52 < 9; fi52++ {
                if true {
                    var v53 int8 = -1
                    if (((v4 >> ((v49) & 63)) >= (v4 - v4)) || (false || false)) {
                        var v54 int16 = v4
                        v55 := GenS7{F0: (v53 << ((v49) & 63))}
                        _ = v54
                        _ = v55
                    } else {
                        v56 := GenS3{F0: (v4 == v4)}
                        switch v3 {
                        case "a":
                            for fi57 := 0; fi57 < 8; fi57++ {
                                if (((v2 % ((v2) | 1)) >= v2) || false) {
                                    v29 = append(v29, (fi52 * fi57))
                                    fmt.Println(fi52)
                                } else {
                                    fmt.Println(v1)
                                }
                                _ = fi57
                            }
                        default:
                            if false {
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    var v59 uint16 = 2
                                    switch (v1 + v3) {
                                    case "\n\t":
                                        v2 += (v2 << (((v49 >> (((v49 << ((v49) & 63))) & 63))) & 63))
                                    case "hello":
                                        if (((((((1e10 + 0.0) / float32(0.1)) + 1.5) / 1.0) - 0.0) * 1.5) != -2.25) {
                                            if ((!true) || true) {
                                                v29 = append(v29, (fi58 << (((v49 * v49)) & 63)))
                                            }
                                            v60 := GenS4{F0: 0, F1: "a"}
                                            var v61 bool = (false && true)
                                            fmt.Println(v2)
                                            v29 = append(v29, 781)
                                            _ = v60
                                            _ = v61
                                        } else {
                                            v62 := GenS2{F0: (v59 % ((uint16(v49)) | 1)), F1: 65535, F2: v2}
                                            _ = v62
                                        }
                                    case "foo bar":
                                        if (!((v4 + (v4 / (((v4 | 910)) | 1))) <= v4)) {
                                            switch fi58 {
                                            case 0:
                                                fmt.Println(v53, v3)
                                            case -9223372036854775808:
                                                var v63 uint8 = 128
                                                _ = v63
                                            default:
                                                if (uint(v59) != (^(v49))) {
                                                    v64 := GenS1{F0: ((1 == -2147483648) || true), F1: v1}
                                                    v53 += int8(v2)
                                                    var v65 string = v1
                                                    _ = v64
                                                    _ = v65
                                                }
                                            }
                                            v66 := GenS6{F0: uint64(v49), F1: (v49 >> ((70) & 63)), F2: v49}
                                            _ = v66
                                        } else {
                                            var v67 uint = (v49 + 810)
                                            _ = v67
                                        }
                                    }
                                    var v68 uint32 = 2147483648
                                    _ = fi58
                                    _ = v59
                                    _ = v68
                                }
                                v29 = append(v29, fi52)
                            }
                        }
                        v69 := GenS5{F0: v2}
                        _ = v56
                        _ = v69
                    }
                    v1 = "\n\t"
                    _ = v53
                }
                _ = fi52
            }
            var v70 string = "\n\t"
            _ = v70
        }
    case "hello":
        if false {
            v71 := GenS5{F0: -329}
            switch 32767 {
            case 0:
                v72 := GenS4{F0: (v4 % ((-1) | 1)), F1: v3}
                _ = v72
            case 32767:
                v3 = v1
            }
            _ = v71
        } else {
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
            v4 = (v4 << (((v49 << (((v49 / ((v49) | 1))) & 63))) & 63))
            var v73 uint32 = 316
            _ = v73
        }
    case "a":
        switch ((v3 + v1) + (v3 + v1)) {
        case "a":
            v74 := GenS5{F0: (v2 % ((v2) | 1))}
            _ = v74
        case "hello":
            for fi75 := 0; fi75 < 12; fi75++ {
                var v76 int16 = v4
                _ = fi75
                _ = v76
            }
        case "foo bar":
            var v77 uint = v49
            _ = v77
        default:
            v3 = v3
        }
    default:
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v49
    _ = v50
}
