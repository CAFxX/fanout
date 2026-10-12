package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    var v1 int8 = 100
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(v1, v1)
        _ = fi5
    }
    if (1 >= 179) {
        v6 := GenS1{F0: 18446744073709551615, F1: -854}
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{}
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(fi9)
                v1 = v1
                fmt.Println(fi7, fi9)
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(v1)
        _ = v6
    } else {
        fmt.Println(v1)
        switch "a" {
        case "foo bar":
            v10 := GenS1{F0: 42, F1: 1}
            _ = v10
        case "":
            if false {
                v11 := GenS1{F0: 895, F1: -582}
                fmt.Println(v1)
                v1 += 127
                _ = v11
            } else {
                if true {
                    if false {
                        v12 := GenS1{F0: 214, F1: 439}
                        _ = v12
                    } else {
                        var v13 int64 = -9223372036854775808
                        _ = v13
                    }
                } else {
                    fmt.Println(v1)
                }
            }
        case "\n\t":
            var v14 float64 = (3.14159 - (((1e10 * (1.5 * (float64(1.5) + 0.1))) / 1e10) - (float64((0.0 * (((((float32(1e10) - ((0.0 / 1.0) * 1e10)) / 3.14159) * 0.1) * 3.14159) * 1e10))) / float64(1.5))))
            _ = v14
        }
        v15 := []int{32767, 1}
        _ = v15
    }
    v1 -= (v1 - 42)
    v1 = v1
    fmt.Println(v1)
    switch ("hello" + "hello") {
    case "hello":
        var v16 float32 = 0.1
        _ = v16
    default:
        var v17 uint16 = 1
        _ = v17
    }
    if (!((v1 >> ((496) & 63)) <= v1)) {
        var v18 float64 = float64(-2.25)
        if true {
            if false {
                v19 := []int{2147483647, 1, 311, 1}
                v1 += 127
                if (!(!true)) {
                    var v20 int32 = 42
                    _ = v20
                } else {
                    v21 := GenS2{F0: 42, F1: 2}
                    var v22 uint32 = 2147483647
                    _ = v21
                    _ = v22
                }
                _ = v19
            } else {
                var v23 uint16 = 32767
                for fi24 := 0; fi24 < 1; fi24++ {
                    if true {
                        var v25 int = fi24
                        fmt.Println(v25, v18, v1)
                        var v26 uint = 549
                        _ = v25
                        _ = v26
                    } else {
                        if (!(-32768 != -1)) {
                            switch fi24 {
                            case 4294967295:
                                fmt.Println(v1, v1)
                            case 999:
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    if ((!true) && ((0.0 * (v18 * float64((1e10 - -2.25)))) != 1.5)) {
                                        v28 := []int{(fi24 + (fi24 ^ (fi27 & 42)))}
                                        _ = v28
                                    } else {
                                        v29 := GenS2{F0: 2, F1: v23}
                                        if (1 <= 42) {
                                            v30 := v4["hello"]
                                            _ = v30
                                        } else {
                                            v18 += (((1e10 + v18) + v18) * float64(-2.25))
                                            var v31 int16 = 42
                                            _ = v31
                                        }
                                        _ = v29
                                    }
                                    _ = fi27
                                }
                            default:
                                if (0 <= 1) {
                                    for fi32 := 0; fi32 < 16; fi32++ {
                                        v1 -= (v1 << ((0) & 63))
                                        v33 := v4[("hello" + "foo bar")]
                                        fmt.Println(v33, v23)
                                        _ = fi32
                                        _ = v33
                                    }
                                    var v34 float32 = 1.5
                                    for fi35 := 0; fi35 < 3; fi35++ {
                                        var v36 float64 = (((float64(1.5) / float64(v34)) * v18) * (((0.0 / -2.25) + v18) + v18))
                                        var v37 uint8 = 255
                                        v18 = (v36 / (0.1 + 0.1))
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v34
                                }
                            }
                            for fi38 := 0; fi38 < 4; fi38++ {
                                v39 := []int{(fi38 + fi24), 42, fi24}
                                v40 := GenS2{F0: 128, F1: v23}
                                var v41 uint8 = 54
                                if len(v39) >= 2 {
                                    v39 = v39[1:len(v39)]
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        } else {
                            for fi42 := 0; fi42 < 6; fi42++ {
                                fmt.Println(v1, fi42)
                                v43 := GenS1{F0: 42, F1: 0}
                                _ = fi42
                                _ = v43
                            }
                            v23 += (v23 / ((v23) | 1))
                        }
                    }
                    if ((0.0 + (-2.25 - (-2.25 - -2.25))) > (3.14159 * 0.1)) {
                        var v44 bool = false
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(fi24, v1, v18)
                        default:
                            v44 = (!(1 != 776))
                        }
                        _ = v44
                    }
                    for fi45 := 0; fi45 < 6; fi45++ {
                        var v46 string = "foo bar"
                        switch fi45 {
                        case -445:
                            var v47 float32 = 1.5
                            _ = v47
                        case 418:
                            if true {
                                var v48 float64 = float64(((3.14159 * 1e10) + float32(3.14159)))
                                switch ((((v46 + ((v46 + "a") + v46)) + v46) + "foo bar") + "\n\t") {
                                case "foo bar":
                                    var v49 int8 = v1
                                    _ = v49
                                default:
                                    v50 := []int{(fi24 | (fi45 & (fi45 & fi45))), (fi45 ^ (fi45 * (fi45 % ((fi24) | 1))))}
                                    _ = v50
                                }
                                v51 := []int{fi24, fi45, fi45, (fi24 * 4294967295)}
                                _ = v48
                                _ = v51
                            } else {
                                var v52 int = fi45
                                switch (fi24 + (-(fi45))) {
                                case 2147483647:
                                    fmt.Println(v46, v46, fi45)
                                case -9223372036854775808:
                                    var v53 float32 = 0.1
                                    switch "hello" {
                                    case "\n\t":
                                        v1 += v1
                                    case "hello":
                                        v54 := GenS1{F0: 18446744073709551615, F1: (v52 * 32768)}
                                        _ = v54
                                    case "foo bar":
                                        var v55 uint64 = 585
                                        _ = v55
                                    default:
                                        v1 = v1
                                    }
                                    var v56 uint16 = (^(v23))
                                    switch (fi45 % ((fi45) | 1)) {
                                    case -566:
                                        var v57 uint64 = 4294967295
                                        _ = v57
                                    }
                                    v58 := GenS3{F0: 0}
                                    _ = v53
                                    _ = v56
                                    _ = v58
                                }
                                _ = v52
                            }
                        default:
                            v23 = (v23 & (v23 % (((^(v23))) | 1)))
                        }
                        _ = fi45
                        _ = v46
                    }
                    _ = fi24
                }
                _ = v23
            }
            var v59 uint32 = 356
            _ = v59
        }
        switch -9223372036854775808 {
        case 950:
            var v60 int64 = 1
            _ = v60
        case -9223372036854775808:
            v61 := GenS4{F0: 42, F1: 1, F2: 18446744073709551615}
            _ = v61
        }
        _ = v18
    }
    if ((true && true) || true) {
        v1 = (v1 >> ((65535) & 63))
        v62 := []int{128}
        _ = len(v3)
        _ = v62
    } else {
        var v63 int = 527
        _ = v63
    }
    var v64 float32 = 1e10
    var v65 float64 = 0.0
    for fi66 := 0; fi66 < 7; fi66++ {
        for fi67 := 0; fi67 < 15; fi67++ {
            v64 += 0.1
            _ = len(v3)
            _ = fi67
        }
        _ = len(v3)
        _ = fi66
    }
    var v68 int64 = 238
    v69 := GenS1{F0: 65535, F1: 42}
    var v70 uint = 255
    if true {
        var v71 uint64 = 0
        var v72 int8 = v1
        _ = v71
        _ = v72
    } else {
        delete(v3, 37)
        v65 = v65
        var v73 uint = v70
        _ = v73
    }
    for fi74 := 0; fi74 < 19; fi74++ {
        if (4294967295 >= 18446744073709551615) {
            v3[fi74] = (-(fi74))
        } else {
            for fi75 := 0; fi75 < 16; fi75++ {
                v76 := GenS1{F0: (v70 + v70), F1: fi75}
                var v77 uint16 = 42
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        if true {
            v64 = v64
            if (2 <= (v70 >> (((v70 % (((v70 + (v70 + v70))) | 1))) & 63))) {
                var v78 uint = (v70 % ((v70) | 1))
                var v79 int8 = v1
                _ = v78
                _ = v79
            } else {
                v68 += 0
                var v80 float32 = (((float32(v65) - ((((3.14159 * v64) / 1.5) / ((1.5 / v64) / v64)) * -2.25)) + (float32(v65) - v64)) * v64)
                _ = v80
            }
            v1 = -128
        } else {
            fmt.Println(v65, v68)
            v68 = v68
        }
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v68
    _ = v69
    _ = v70
}
