package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 string; F1 int64; F2 bool }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int; F1 uint8; F2 string }
type GenS8 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int = 32767
    var v2 int = (v1 ^ v1)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    v1 += (v2 % ((v1) | 1))
    if ((true || (904 <= 42)) && (true || false)) {
        v5 := GenS1{F0: 971}
        v1 += v1
        for fi6 := 0; fi6 < 8; fi6++ {
            if true {
                var v7 int16 = 42
                if true {
                    v8 := GenS1{F0: 2}
                    v3["a"] = 65535
                    _ = len(v4)
                    _ = v8
                }
                for fi9 := 0; fi9 < 5; fi9++ {
                    if true {
                        var v10 float64 = ((0.1 + 0.1) - (1.5 - 3.14159))
                        v10 -= (v10 + v10)
                        _ = v10
                    } else {
                        fmt.Println(v1)
                        fmt.Println(v7, fi9)
                        if (0.1 > -2.25) {
                            fmt.Println(v7)
                        } else {
                            if (0 > -4) {
                                var v11 string = ("hello" + (("\n\t" + "") + ("hello" + ("foo bar" + "foo bar"))))
                                _ = v11
                            }
                        }
                    }
                    _ = fi9
                }
                _ = v7
            } else {
                switch "a" {
                case "a":
                    var v12 uint = 42
                    _ = v12
                case "":
                    fmt.Println(fi6, v1)
                case "\n\t":
                    v13 := GenS2{F0: -115, F1: 65535, F2: 1}
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        if (4294967295 != 616) {
                            if false {
                                v15 := GenS3{F0: -486, F1: 127, F2: 65535}
                                _ = v15
                            }
                            if true {
                                v16 := GenS4{F0: 1, F1: -32768}
                                if ((100 != 100) && true) {
                                    v17 := GenS3{F0: -430, F1: 47, F2: 1}
                                    _ = v17
                                }
                                v18 := GenS4{F0: 2, F1: -190}
                                _ = v16
                                _ = v18
                            }
                        }
                        delete(v3, "\n\t")
                        switch v1 {
                        case 181:
                            v2 -= 4294967295
                        case -918:
                            var v19 bool = (-2147483648 == 32767)
                            _ = v19
                        case 128:
                            if (1 != 9223372036854775807) {
                                v20 := GenS1{F0: 4294967295}
                                if true {
                                    v21 := GenS3{F0: 42, F1: 0, F2: 64}
                                    switch "foo bar" {
                                    case "hello":
                                        var v22 int16 = -32768
                                        _ = v22
                                    }
                                    _ = v21
                                }
                                _ = v20
                            } else {
                                if false {
                                    var v23 int = v1
                                    _ = v23
                                }
                                var v24 bool = true
                                _ = v24
                            }
                        }
                        _ = fi14
                    }
                }
            }
            delete(v4, 0)
            var v25 string = "foo bar"
            _ = fi6
            _ = v25
        }
        _ = v5
    }
    switch "\n\t" {
    case "hello":
        v26 := v4[419]
        _ = v26
    case "a":
        v1 = v2
    default:
        v27 := GenS5{F0: ("" + ("hello" + "a")), F1: 130, F2: (1 >= -1)}
        _ = v27
    }
    var v28 uint16 = 2
    v29 := GenS4{F0: (v28 ^ (v28 + 255)), F1: -418}
    switch (-(v2)) {
    case -1:
        for fi30 := 0; fi30 < 13; fi30++ {
            v1 -= v1
            v31 := []int{0, -890, -2147483648, -278}
            v3["hello"] = 0
            _ = fi30
            _ = v31
        }
    case -136:
        if ((-2.25 - float32((float64(float32(1e10)) * 0.1))) > 0.1) {
            var v32 int32 = 0
            if true {
                fmt.Println(v32, v2, v1)
                v33 := v4[884]
                _ = v33
            }
            v32 = 1
            _ = v32
        } else {
            var v34 float64 = 1.5
            _ = v34
        }
    case 32767:
        var v35 int = (v2 - v1)
        _ = v35
    default:
        switch (^(v2)) {
        case 32767:
            fmt.Println(v1, v2)
        case 4294967295:
            switch (v2 << ((18446744073709551615) & 63)) {
            case 128:
                var v36 string = "hello"
                _ = v36
            default:
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch ("\n\t" + "") {
                    case "":
                        switch (^(v1)) {
                        case 0:
                            var v38 int = (fi37 % ((v2) | 1))
                            _ = v38
                        case 1:
                            for fi39 := 0; fi39 < 4; fi39++ {
                                _ = len(v4)
                                var v40 int8 = 35
                                var v41 float32 = 1e10
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        default:
                            for fi42 := 0; fi42 < 19; fi42++ {
                                if ("hello" > "\n\t") {
                                    var v43 bool = (279 != 42)
                                    if v43 {
                                        var v44 int8 = 91
                                        var v45 float32 = 0.1
                                        v46 := []int{(v2 << ((292) & 63)), fi42, (fi42 << ((2) & 63))}
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    switch fi42 {
                                    case -835:
                                        fmt.Println(v43, v28, v2)
                                    case -9223372036854775808:
                                        v47 := GenS5{F0: "", F1: 749, F2: (v43 && (2 != 255))}
                                        _ = v47
                                    default:
                                        if (!false) {
                                            v28 -= (^(v28))
                                            if (2 < 42) {
                                                v43 = (1 > 1)
                                                v48 := GenS4{F0: 786, F1: 42}
                                                _ = v48
                                            } else {
                                                delete(v4, 855)
                                                switch "a" {
                                                case "a":
                                                    if false {
                                                        v49 := GenS6{F0: 2, F1: 18446744073709551615}
                                                        _ = v49
                                                    }
                                                case "":
                                                    var v50 uint32 = 1
                                                    _ = v50
                                                case "foo bar":
                                                    v3[("\n\t" + "hello")] = 2
                                                default:
                                                    var v51 uint8 = 127
                                                    _ = v51
                                                }
                                            }
                                        } else {
                                            v52 := []int{v2, -1, fi37, (fi37 & (fi42 * fi37))}
                                            v1 = (^(fi37))
                                            _ = v52
                                        }
                                    }
                                    _ = v43
                                }
                                var v53 bool = false
                                switch v2 {
                                case 9223372036854775807:
                                    var v54 uint32 = 65535
                                    _ = v54
                                }
                                _ = fi42
                                _ = v53
                            }
                        }
                    case "hello":
                        v2 += fi37
                    case "foo bar":
                        fmt.Println(v1, v28)
                    default:
                        v55 := GenS6{F0: 2, F1: 42}
                        _ = v55
                    }
                    var v56 int32 = -1
                    v1 += fi37
                    _ = fi37
                    _ = v56
                }
            }
        case 505:
            if (!(!(1 == v2))) {
                for fi57 := 0; fi57 < 2; fi57++ {
                    v2 -= v2
                    var v58 int16 = -656
                    _ = fi57
                    _ = v58
                }
                switch (v1 >> ((65535) & 63)) {
                case 4294967295:
                    v59 := GenS5{F0: (("foo bar" + "hello") + "a"), F1: 781, F2: false}
                    _ = v59
                case 32768:
                    var v60 uint = 214
                    _ = v60
                case 1:
                    for fi61 := 0; fi61 < 14; fi61++ {
                        v62 := GenS5{F0: "\n\t", F1: 42, F2: true}
                        _ = fi61
                        _ = v62
                    }
                default:
                    for fi63 := 0; fi63 < 4; fi63++ {
                        v64 := GenS7{F0: v2, F1: 100, F2: ("hello" + "foo bar")}
                        v65 := []int{-1, v1, -9223372036854775808}
                        v28 += (v28 - (v28 | v28))
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                }
                var v66 float32 = (((0.0 - 0.1) + 0.1) / 1e10)
                _ = v66
            }
        }
    }
    v1 = (v1 & (v1 >> ((1) & 63)))
    fmt.Println(v2)
    var v67 int8 = 100
    v68 := []int{}
    var v69 uint16 = (v28 % (((v28 & 1)) | 1))
    v3["\n\t"] = 42
    fmt.Println(v69, v1)
    v70 := GenS8{F0: (v28 << ((18446744073709551615) & 63)), F1: 42}
    v2 = 128
    switch "a" {
    case "foo bar":
        fmt.Println(v28, v67)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
