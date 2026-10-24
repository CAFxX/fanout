package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int32; F1 string }

func main() {
    v1 := []uint32{4294967295, 65535, 611}
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    v5 := GenS1{F0: (v4 % ((v4) | 1))}
    v6 := GenS1{F0: 42}
    var v7 uint = 4294967295
    var v8 float32 = (((1e10 / -2.25) / float32(-2.25)) + float32(((1.5 * 0.1) + (0.0 / 1.0))))
    var v9 int8 = 0
    var v10 uint64 = 2
    v9 = v9
    fmt.Println(v10, v8, v9)
    switch -9223372036854775808 {
    case -2147483648:
        v4 += (^(v4))
    case 4294967295:
        var v11 uint = (v7 & v7)
        _ = v11
    }
    delete(v3, ("\n\t" + ("\n\t" + (("\n\t" + ((("\n\t" + ("" + "a")) + "foo bar") + "a")) + "\n\t"))))
    v10 = (^(v10))
    for fi12 := 0; fi12 < 14; fi12++ {
        switch "foo bar" {
        case "a":
            if ((v4 / ((127) | 1)) == (v4 << ((956) & 63))) {
                if false {
                    v13 := GenS2{F0: -32768, F1: 442}
                    if false {
                        v7 = 1
                        fmt.Println(fi12, v8, v7)
                        for fi14 := 0; fi14 < 9; fi14++ {
                            var v15 int = fi12
                            var v16 int32 = -2147483648
                            switch "foo bar" {
                            case "a":
                                if false {
                                    var v17 int32 = v16
                                    fmt.Println(v16, fi14, v7)
                                    _ = v17
                                } else {
                                    switch ((("" + "\n\t") + "hello") + ("\n\t" + "foo bar")) {
                                    case "hello":
                                        if (v15 <= (fi12 & -299)) {
                                            var v18 int64 = 9223372036854775807
                                            v7 += v7
                                            v19 := v3["\n\t"]
                                            _ = v18
                                            _ = v19
                                        } else {
                                            var v20 uint8 = v4
                                            _ = v20
                                        }
                                    default:
                                        v4 = v4
                                    }
                                }
                            case "\n\t":
                                v21 := GenS1{F0: v4}
                                _ = v21
                            case "":
                                v16 -= (^(v16))
                            default:
                                if (-1 <= 9223372036854775807) {
                                    var v22 string = "hello"
                                    _ = v22
                                }
                            }
                            _ = len(v1)
                            fmt.Println(v4)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    v23 := GenS2{F0: 42, F1: 202}
                    _ = v13
                    _ = v23
                } else {
                    if (1 < 0) {
                        if (!true) {
                            v3[("\n\t" + "foo bar")] = 4294967295
                            v4 = 0
                            v24 := GenS3{F0: "a"}
                            _ = v24
                        }
                        var v25 int64 = int64(fi12)
                        var v26 float64 = (3.14159 / (1e10 * 1.5))
                        _ = v25
                        _ = v26
                    } else {
                        var v27 float32 = (v8 / ((3.14159 * v8) + (float32(0.0) - (v8 / v8))))
                        var v28 int16 = 42
                        for fi29 := 0; fi29 < 17; fi29++ {
                            fmt.Println(v9, fi29, v27)
                            var v30 float32 = (1e10 / v27)
                            _ = fi29
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    }
                    v2[(fi12 % ((-1) | 1))] = (230 != uint32(v10))
                    var v31 int64 = -9223372036854775808
                    _ = v31
                }
            }
        case "":
            if (float32(0.0) <= (1.5 - v8)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v32 int64 = -905
                v10 -= v10
                _ = v32
            }
        case "hello":
            if false {
                var v33 uint = (v7 % ((0) | 1))
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 uint64 = 18446744073709551615
                    v4 += 145
                    _ = fi34
                    _ = v35
                }
                fmt.Println(v33, v33)
                _ = v33
            }
        }
        var v36 uint16 = 42
        var v37 uint = 18446744073709551615
        _ = fi12
        _ = v36
        _ = v37
    }
    var v38 float64 = 1e10
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 uint = (^(v7))
        _ = fi39
        _ = v40
    }
    if true {
        if (!(0.1 <= 1e10)) {
            v7 += (v7 - (v7 >> (((v7 * v7)) & 63)))
        } else {
            var v41 uint64 = 2
            delete(v3, (("a" + "") + ""))
            _ = v41
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            _ = len(v2)
            var v43 float64 = v38
            if (42 < 591) {
                if true {
                    switch fi42 {
                    case 543:
                        v8 = 3.14159
                    case 4294967295:
                        v44 := GenS4{F0: 70}
                        _ = v44
                    default:
                        v7 += 1
                    }
                } else {
                    v7 = 810
                    switch (fi42 / (((-(fi42))) | 1)) {
                    case 622:
                        for fi45 := 0; fi45 < 7; fi45++ {
                            var v46 uint64 = uint64(v7)
                            var v47 int64 = -231
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    case 1:
                        fmt.Println(v43, v8)
                    case 128:
                        var v48 string = "\n\t"
                        _ = v48
                    default:
                        var v49 float64 = v38
                        _ = v49
                    }
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        for fi50 := 0; fi50 < 11; fi50++ {
                            switch "hello" {
                            case "\n\t":
                                v43 += float64((v8 * v8))
                            default:
                                if false {
                                    v51 := GenS5{F0: (v10 / (((v10 * 4294967295)) | 1))}
                                    var v52 uint64 = 1
                                    _ = v51
                                    _ = v52
                                }
                            }
                            if (65535 == 255) {
                                var v53 int = fi42
                                _ = v53
                            }
                            v1 = append(v1, 2)
                            _ = fi50
                        }
                    case "":
                        v43 += v43
                    default:
                        var v54 uint32 = 0
                        _ = v54
                    }
                }
                if (752 <= 32767) {
                    var v55 int8 = 127
                    for fi56 := 0; fi56 < 18; fi56++ {
                        fmt.Println(v9)
                        _ = fi56
                    }
                    switch fi42 {
                    case 959:
                        v55 -= 127
                    }
                    _ = v55
                }
            }
            _ = fi42
            _ = v43
        }
        if false {
            for fi57 := 0; fi57 < 15; fi57++ {
                _ = len(v1)
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v58 := GenS2{F0: -32768, F1: v7}
                    _ = v58
                case -1:
                    for fi59 := 0; fi59 < 11; fi59++ {
                        v60 := v2[fi57]
                        v38 += ((v38 + v38) * (3.14159 / 1.5))
                        if false {
                            fmt.Println(v38)
                            if (!(42 >= 2)) {
                                var v61 int8 = v9
                                _ = v61
                            }
                            var v62 string = "a"
                            _ = v62
                        }
                        _ = fi59
                        _ = v60
                    }
                case -106:
                    v63 := v2[fi57]
                    _ = v63
                }
                _ = fi57
            }
            v64 := GenS3{F0: ("a" + "a")}
            fmt.Println(v7, v9)
            _ = v64
        } else {
            switch int(v9) {
            case -449:
                v9 = -117
            }
        }
    } else {
        if (((v10 + (v10 + uint64(v4))) != v10) || true) {
            var v65 float32 = ((float32((v38 + (v38 / v38))) * v8) / v8)
            if ((1 < 2147483647) || false) {
                v4 = v4
                v66 := v3[(("foo bar" + ("\n\t" + "\n\t")) + "\n\t")]
                v67 := GenS2{F0: 334, F1: 503}
                _ = v66
                _ = v67
            } else {
                var v68 int64 = -883
                v69 := GenS2{F0: 1, F1: 959}
                if (42 < v10) {
                    v7 -= (v7 - (^(v7)))
                    v9 -= v9
                } else {
                    if false {
                        var v70 int64 = (v68 << ((v7) & 63))
                        for fi71 := 0; fi71 < 16; fi71++ {
                            fmt.Println(v9, v9)
                            _ = fi71
                        }
                        v72 := v3[("hello" + "a")]
                        _ = v70
                        _ = v72
                    }
                    v73 := GenS1{F0: 102}
                    for fi74 := 0; fi74 < 17; fi74++ {
                        for fi75 := 0; fi75 < 6; fi75++ {
                            v7 += (v7 | (v7 & (v7 + 2)))
                            v9 += v9
                            var v76 int16 = 1
                            _ = fi75
                            _ = v76
                        }
                        _ = fi74
                    }
                    _ = v73
                }
                _ = v68
                _ = v69
            }
            if len(v1) > 0 {
                v1[0] = 296
            }
            _ = v65
        } else {
            if len(v1) > 0 {
                v1[0] = 848
            }
            var v77 uint64 = (v10 - 0)
            _ = v77
        }
    }
    switch "\n\t" {
    case "hello":
        v78 := GenS6{F0: 32767, F1: 18446744073709551615}
        _ = v78
    case "foo bar":
        v8 += (v8 + v8)
    }
    if (-740 == 1000) {
        v79 := GenS7{F0: 42, F1: "hello"}
        _ = v79
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
