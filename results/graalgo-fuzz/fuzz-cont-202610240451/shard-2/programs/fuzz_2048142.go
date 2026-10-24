package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint; F2 int }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 string }
func genFn0() int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        if (true || false) {
            var v6 int = fi5
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 138, F1: 772, F2: fi7}
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint32 = 2
        switch (("" + ("a" + ("" + (("hello" + "a") + "hello")))) + "foo bar") {
        case "hello":
            var v12 int64 = -1
            _ = v12
        case "\n\t":
            switch fi10 {
            case 1:
                var v13 uint64 = 4294967295
                _ = v13
            case 42:
                v14 := GenS1{F0: 1, F1: 2, F2: (fi10 & 4294967295)}
                _ = v14
            }
        }
        _ = fi10
        _ = v11
    }
    var v15 uint32 = 2
    _ = v15
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 uint = 65535
    v3 := make(map[uint32]int)
    var v4 int32 = 1
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := []int{(fi16 | (^(fi16))), (fi16 % (((fi16 << (((v2 / (((v2 | 65535)) | 1))) & 63))) | 1)), (fi16 << ((v2) & 63))}
        v18 := GenS2{F0: 303, F1: 485, F2: 487}
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        var v19 uint = (v2 | v2)
        for fi20 := 0; fi20 < 8; fi20++ {
            if false {
                switch (fi20 | (fi20 << ((v19) & 63))) {
                case 623:
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 bool = (!(128 <= uint8(v2)))
                        fmt.Println(v22, v2)
                        if v22 {
                            if true {
                                if true {
                                    fmt.Println(fi21, fi21)
                                }
                                v23 := []int{}
                                switch fi20 {
                                case 32767:
                                    _ = len(v3)
                                case 2147483647:
                                    v2 = v2
                                    v3[726] = fi21
                                    var v24 string = ("\n\t" + ("hello" + "hello"))
                                    _ = v24
                                }
                                _ = v23
                            }
                            fmt.Println(v4)
                            v25 := GenS1{F0: 1, F1: (v19 + (v19 / ((v2) | 1))), F2: (fi20 >> ((v19) & 63))}
                            _ = v25
                        } else {
                            if (!v22) {
                                v26 := GenS3{F0: -9223372036854775808}
                                v22 = ((("\n\t" + "a") + "a") > "foo bar")
                                switch "a" {
                                case "hello":
                                    fmt.Println(v19, v1)
                                case "foo bar":
                                    if (true || v22) {
                                        v4 = v4
                                    }
                                case "a":
                                    if ((-721 != genFn0()) && false) {
                                        delete(v3, 0)
                                    }
                                default:
                                    _ = len(v3)
                                }
                                _ = v26
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                case -2147483648:
                    switch "" {
                    case "a":
                        fmt.Println(v19, v4)
                    case "":
                        v1 = (^(v1))
                    case "foo bar":
                        switch fi20 {
                        case -334:
                            v1 = v1
                        case -680:
                            v1 += 42
                        case 32768:
                            var v27 string = ""
                            _ = v27
                        default:
                            switch (fi20 << (((v2 - (v19 | v19))) & 63)) {
                            case 2147483647:
                                fmt.Println(fi20, v19)
                            case 1:
                                for fi28 := 0; fi28 < 18; fi28++ {
                                    v29 := v3[42]
                                    v1 -= (v1 * v1)
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        }
                    default:
                        v4 -= v4
                    }
                default:
                    fmt.Println(v1, v1)
                }
                switch 0 {
                case -9223372036854775808:
                    if (18446744073709551615 <= 446) {
                        for fi30 := 0; fi30 < 2; fi30++ {
                            var v31 uint = (v2 * v2)
                            if false {
                                var v32 float32 = (((((((0.0 + (float32(1e10) - 1.5)) - float32(((3.14159 * 1.5) * (float64(3.14159) - 0.1)))) / 1.5) + ((1.5 + -2.25) * -2.25)) * -2.25) / 1.0) + 3.14159)
                                if (0 > 1000) {
                                    if false {
                                        fmt.Println(fi20, v4)
                                    }
                                } else {
                                    if ((!false) || true) {
                                        delete(v3, 366)
                                        v19 += v31
                                    }
                                    v33 := GenS4{F0: "\n\t", F1: v19, F2: (fi20 % (((fi30 / ((fi20) | 1))) | 1))}
                                    _ = v33
                                }
                                _ = v32
                            }
                            var v34 int64 = -672
                            _ = fi30
                            _ = v31
                            _ = v34
                        }
                        v1 += (-(v1))
                        var v35 float64 = 1e10
                        _ = v35
                    } else {
                        var v36 uint32 = 689
                        v37 := GenS5{F0: (v1 & -128), F1: -1}
                        _ = v36
                        _ = v37
                    }
                case 42:
                    v38 := []int{fi20}
                    _ = v38
                }
            }
            var v39 int16 = 32767
            var v40 int8 = v1
            _ = fi20
            _ = v39
            _ = v40
        }
        _ = v19
    }
    v41 := []int{-2147483648}
    fmt.Println(v2, v1)
    v4 = v4
    for fi42 := 0; fi42 < 18; fi42++ {
        if true {
            v43 := GenS5{F0: (v1 | v1), F1: 0}
            _ = v43
        }
        _ = len(v41)
        _ = fi42
    }
    switch "" {
    case "hello":
        switch 409 {
        case 127:
            var v44 int64 = 943
            _ = v44
        case 128:
            delete(v3, 795)
        case 9223372036854775807:
            var v45 uint8 = 2
            _ = v45
        }
    default:
        var v46 int16 = -448
        _ = v46
    }
    switch "\n\t" {
    case "":
        _ = len(v41)
    case "a":
        v4 = (v4 * 2147483647)
    }
    fmt.Println(v4, v1, v2)
    fmt.Println(v4, v2, v1)
    for fi47 := 0; fi47 < 4; fi47++ {
        switch -1 {
        case 42:
            var v48 int32 = v4
            _ = v48
        default:
            var v49 uint8 = 128
            _ = v49
        }
        v2 = 1
        v50 := GenS6{F0: 128}
        _ = fi47
        _ = v50
    }
    var v51 int = -227
    if true {
        for fi52 := 0; fi52 < 16; fi52++ {
            v53 := GenS7{F0: "foo bar", F1: 4294967295}
            _ = fi52
            _ = v53
        }
        v1 -= 127
        if true {
            delete(v3, 2)
            for fi54 := 0; fi54 < 6; fi54++ {
                if false {
                    for fi55 := 0; fi55 < 16; fi55++ {
                        if true {
                            fmt.Println(v51, v4, v4)
                            var v56 uint16 = 42
                            _ = v56
                        }
                        var v57 int16 = 0
                        v58 := GenS3{F0: 0}
                        _ = fi55
                        _ = v57
                        _ = v58
                    }
                    v59 := GenS1{F0: 9223372036854775807, F1: (v2 << (((v2 << ((v2) & 63))) & 63)), F2: (fi54 << ((v2) & 63))}
                    var v60 bool = false
                    _ = v59
                    _ = v60
                }
                switch fi54 {
                case 42:
                    v61 := GenS6{F0: 1}
                    _ = v61
                default:
                    fmt.Println(v4)
                }
                var v62 int16 = -32768
                _ = fi54
                _ = v62
            }
        } else {
            v63 := GenS6{F0: 62}
            if (9223372036854775807 > int64(v4)) {
                v64 := GenS7{F0: ((("foo bar" + "") + ("hello" + (("a" + "hello") + "foo bar"))) + "foo bar"), F1: 1}
                fmt.Println(v51, v51, v1)
                v4 -= (v4 - v4)
                _ = v64
            }
            _ = v63
        }
    } else {
        if true {
            fmt.Println(v1)
            if false {
                for fi65 := 0; fi65 < 8; fi65++ {
                    switch "a" {
                    case "foo bar":
                        v66 := GenS8{F0: 43}
                        _ = v66
                    }
                    _ = fi65
                }
            } else {
                for fi67 := 0; fi67 < 8; fi67++ {
                    fmt.Println(fi67, v2)
                    if false {
                        v68 := GenS9{F0: 0, F1: "\n\t"}
                        _ = v68
                    } else {
                        v1 = 100
                    }
                    if len(v41) > 0 {
                        v41[0] = fi67
                    }
                    _ = fi67
                }
                var v69 float32 = float32(-2.25)
                var v70 string = (("foo bar" + ("\n\t" + ("foo bar" + "a"))) + "")
                _ = v69
                _ = v70
            }
            for fi71 := 0; fi71 < 13; fi71++ {
                v41 = append(v41, -1)
                fmt.Println(v1)
                if true {
                    switch (fi71 >> (((v2 | 18446744073709551615)) & 63)) {
                    case 4294967295:
                        v72 := GenS1{F0: 876, F1: (v2 * v2), F2: v51}
                        _ = v72
                    }
                    _ = len(v3)
                    v73 := GenS4{F0: "a", F1: v2, F2: (v51 << (((v2 - v2)) & 63))}
                    _ = v73
                } else {
                    fmt.Println(fi71, v51)
                    if ("" < "foo bar") {
                        v74 := GenS3{F0: -646}
                        _ = len(v3)
                        v4 -= (v4 | (v4 ^ 1))
                        _ = v74
                    }
                }
                _ = fi71
            }
        }
        v51 += v51
    }
    var v75 float32 = (float32(3.14159) + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v51:", v51)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v51
    _ = v75
}
