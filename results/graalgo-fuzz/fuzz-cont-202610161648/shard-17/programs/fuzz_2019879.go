package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 bool }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 uint }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{633, 2, 2147483647}
    v3 := make(map[uint32]bool)
    var v4 float64 = v1
    v5 := make(map[int]int)
    if (!true) {
        var v6 uint32 = 497
        _ = v6
    } else {
        var v7 bool = (true && true)
        v8 := GenS1{F0: 0, F1: -2147483648, F2: 4294967295}
        _ = v7
        _ = v8
    }
    switch ("a" + ("foo bar" + ("hello" + ("" + "foo bar")))) {
    case "foo bar":
        v9 := GenS1{F0: 390, F1: 248, F2: 0}
        _ = v9
    case "a":
        v10 := GenS2{F0: 819, F1: 18446744073709551615}
        _ = v10
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int8 = 0
            switch "foo bar" {
            case "":
                var v13 float64 = v1
                _ = v13
            case "a":
                v12 += v12
            default:
                var v14 uint32 = 2147483647
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    if false {
        _ = len(v5)
        switch ("foo bar" + "a") {
        case "foo bar":
            for fi15 := 0; fi15 < 12; fi15++ {
                switch ("hello" + ((((("foo bar" + "a") + "\n\t") + "foo bar") + "\n\t") + "foo bar")) {
                case "\n\t":
                    switch (fi15 ^ (-(fi15))) {
                    case 4294967295:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            var v17 float32 = (0.0 - 0.0)
                            switch (fi16 / (((fi16 >> ((1) & 63))) | 1)) {
                            case 1:
                                var v18 uint32 = 2
                                _ = v18
                            case 128:
                                fmt.Println(fi15)
                            case -2147483648:
                                switch "hello" {
                                case "":
                                    var v19 float64 = (float64(((float32(0.1) / float32(v4)) - 1.5)) + 3.14159)
                                    _ = v19
                                case "foo bar":
                                    v20 := v5[(-(fi16))]
                                    _ = v20
                                }
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                    case 88:
                        if (("foo bar" == ((("hello" + "") + (("foo bar" + "foo bar") + ("" + ("\n\t" + "foo bar")))) + (("foo bar" + "\n\t") + "\n\t"))) && true) {
                            var v21 uint32 = 65535
                            _ = v21
                        }
                    case 83:
                        v22 := GenS3{F0: -513, F1: 746}
                        _ = v22
                    }
                default:
                    v23 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                    _ = v23
                }
                if false {
                    delete(v5, 1)
                    v24 := GenS4{F0: 240, F1: "", F2: (2 <= 2)}
                    var v25 uint16 = 65535
                    _ = v24
                    _ = v25
                } else {
                    if true {
                        v2 = append(v2, 0)
                        switch "hello" {
                        case "":
                            v4 -= v4
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                        case "a":
                            fmt.Println(v4)
                        }
                        v1 -= v4
                    }
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 uint64 = 2
                        v4 += (((v1 + ((v4 - v1) * v1)) / v1) - ((1.5 + (float64(((((((1.5 * (3.14159 - 1e10)) * (0.0 * 1.5)) + 0.1) / 0.1) * -2.25) * float32(3.14159))) - float64(0.0))) - (v1 / v1)))
                        _ = fi26
                        _ = v27
                    }
                }
                var v28 uint32 = 2147483647
                _ = fi15
                _ = v28
            }
        }
    } else {
        v4 = 0.1
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 uint8 = 127
            if (993 == 42) {
                if len(v2) > 0 {
                    v2[0] = 0
                }
            }
            _ = fi30
            _ = v31
        }
        if ("hello" >= "hello") {
            if false {
                switch "hello" {
                case "hello":
                    if (737 >= 4294967295) {
                        var v32 uint64 = 2
                        _ = v32
                    }
                case "\n\t":
                    var v33 bool = (1e10 < -2.25)
                    _ = v33
                }
                if ((!false) && (float32((1e10 - (float64(0.0) * (0.0 * v1)))) < 0.1)) {
                    fmt.Println(v1)
                    var v34 uint32 = 814
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v36 := GenS5{F0: 255, F1: v34, F2: 1}
                        v37 := GenS6{F0: 0}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                } else {
                    var v38 int16 = -32768
                    v39 := GenS5{F0: 552, F1: 974, F2: 73}
                    var v40 int = fi29
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
            var v41 int = (fi29 / ((fi29) | 1))
            _ = v41
        } else {
            switch (fi29 % ((fi29) | 1)) {
            case 4294967295:
                fmt.Println(v4, v4)
            case 9223372036854775807:
                var v42 uint = 788
                _ = v42
            case 32768:
                for fi43 := 0; fi43 < 19; fi43++ {
                    if true {
                        var v44 uint16 = 2
                        if false {
                            _ = len(v5)
                            switch ("a" + "a") {
                            case "":
                                var v45 uint32 = 4294967295
                                _ = v45
                            case "hello":
                                var v46 bool = (!false)
                                _ = v46
                            default:
                                v4 = float64(0.1)
                            }
                            v4 += v1
                            v47 := GenS7{F0: 65535, F1: 1}
                            _ = v47
                        } else {
                            if (!(false && true)) {
                                fmt.Println(fi29, v44, v44)
                                v48 := GenS6{F0: -1}
                                _ = v48
                            } else {
                                v4 += v1
                                switch "hello" {
                                case "a":
                                    var v49 int32 = 0
                                    v2 = append(v2, 2147483647)
                                    var v50 int32 = 42
                                    if (255 < 128) {
                                        _ = len(v3)
                                        var v51 uint = 4294967295
                                        v44 = 65535
                                        _ = v51
                                    }
                                    _ = v49
                                    _ = v50
                                case "foo bar":
                                    var v52 uint64 = 1
                                    _ = v52
                                case "\n\t":
                                    v4 = 0.0
                                }
                            }
                            var v53 uint64 = 42
                            var v54 string = "a"
                            _ = v53
                            _ = v54
                        }
                        var v55 int64 = -1
                        _ = v44
                        _ = v55
                    } else {
                        var v56 bool = true
                        fmt.Println(fi43, fi29)
                        _ = v56
                    }
                    fmt.Println(fi43, v4)
                    v57 := GenS8{F0: 42}
                    v5[(fi29 ^ 2147483647)] = fi43
                    var v58 uint16 = 0
                    v4 += ((v4 + 1.5) / v4)
                    _ = fi43
                    _ = v57
                    _ = v58
                }
            }
            switch "" {
            case "a":
                var v59 uint32 = 2147483647
                _ = v59
            default:
                var v60 uint = 255
                _ = v60
            }
            v1 = v1
        }
        _ = fi29
    }
    if false {
        if (4294967295 < 18446744073709551615) {
            fmt.Println(v4, v1)
            var v61 uint8 = 0
            _ = v61
        }
        v62 := GenS5{F0: 42, F1: 42, F2: 32767}
        _ = v62
    }
    if true {
        if ("\n\t" <= ("" + "\n\t")) {
            switch "\n\t" {
            case "a":
                v4 += (1.5 * v1)
            case "":
                var v63 uint64 = 868
                _ = v63
            default:
                v64 := GenS4{F0: 2147483647, F1: "\n\t", F2: false}
                _ = v64
            }
            fmt.Println(v1, v1)
            var v65 float64 = 0.0
            _ = v65
        }
    }
    v66 := GenS2{F0: 0, F1: 4294967295}
    var v67 uint16 = 65535
    if true {
        var v68 bool = true
        _ = v68
    } else {
        switch -2147483648 {
        case 9223372036854775807:
            var v69 float32 = (((float32((0.1 / v1)) - -2.25) / (-2.25 * float32(v4))) * (1.5 + float32(float64(-2.25))))
            _ = v69
        case 4294967295:
            var v70 int32 = -1
            _ = v70
        }
        if (1 >= 0) {
            for fi71 := 0; fi71 < 2; fi71++ {
                fmt.Println(fi71, v67, fi71)
                var v72 int16 = -1
                v1 += float64(float32(v4))
                _ = fi71
                _ = v72
            }
        }
    }
    if false {
        v73 := GenS3{F0: 484, F1: 2}
        var v74 uint32 = 1
        v74 = v74
        _ = v73
        _ = v74
    }
    var v75 int16 = -202
    var v76 uint32 = 2147483647
    if false {
        var v77 string = ("a" + "hello")
        _ = v77
    }
    if ((v76 ^ v76) >= 1) {
        var v78 int = 9223372036854775807
        v76 = (v76 - v76)
        _ = v78
    }
    v75 += (v75 / ((-955) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v75
    _ = v76
}
