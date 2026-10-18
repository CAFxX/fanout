package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16; F1 int8; F2 int64 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 int; F1 uint32; F2 int32 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := []string{("foo bar" + (("\n\t" + "") + ("\n\t" + "\n\t"))), "\n\t"}
    v2 := []uint32{4294967295, 980, 4294967295, 1}
    v3 := make(map[uint32]string)
    var v4 uint32 = 65535
    var v5 uint = 4294967295
    v6 := make(map[int]string)
    var v7 int16 = 1000
    if true {
        v8 := GenS1{F0: -47, F1: (^(v4)), F2: -427}
        if ((!true) && true) {
            v7 = -1
        } else {
            _ = len(v1)
            var v9 float64 = 0.0
            _ = v9
        }
        if (false && ((false || (1 >= 127)) || (!(-523 == 42)))) {
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 int64 = 0
                v12 := GenS1{F0: 14, F1: (v4 / (((v4 << (((v5 - v5)) & 63))) | 1)), F2: v7}
                var v13 uint = 0
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = v8
    } else {
        switch "foo bar" {
        case "":
            v7 = (v7 - (^(v7)))
        case "\n\t":
            v5 += (v5 | v5)
        case "foo bar":
            var v14 int16 = v7
            _ = v14
        }
        var v15 float64 = ((0.0 * (((float64((0.0 + -2.25)) / (float64((1.5 + 1.5)) - 3.14159)) - ((-2.25 * 1e10) + 0.1)) - 1e10)) + 3.14159)
        _ = v15
    }
    v16 := GenS2{F0: false, F1: 0, F2: 42}
    if (true && true) {
        for fi17 := 0; fi17 < 3; fi17++ {
            if true {
                v5 -= v5
                if (255 < 255) {
                    if (!(!(3.14159 >= (-2.25 - 1e10)))) {
                        v5 += v5
                        var v18 float64 = 3.14159
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 int8 = 1
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        v2 = append(v2, v4)
                        for fi21 := 0; fi21 < 4; fi21++ {
                            var v22 int64 = int64(v4)
                            _ = fi21
                            _ = v22
                        }
                        v4 -= v4
                    }
                } else {
                    switch -1 {
                    case 42:
                        fmt.Println(v7, fi17, fi17)
                    case 32767:
                        if true {
                            fmt.Println(v5, v7, fi17)
                            fmt.Println(v7, v4, v5)
                            for fi23 := 0; fi23 < 3; fi23++ {
                                fmt.Println(fi17, fi17, v5)
                                var v24 uint = (^(v5))
                                switch "foo bar" {
                                case "foo bar":
                                    var v25 string = "hello"
                                    _ = v25
                                case "hello":
                                    v24 -= (v5 - 65535)
                                default:
                                    fmt.Println(fi23, fi23, v7)
                                }
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            var v26 float64 = 1e10
                            _ = len(v2)
                            _ = v26
                        }
                    }
                    if (!(963 > 557)) {
                        var v27 float64 = float64(1e10)
                        _ = v27
                    }
                }
                v5 = v5
            } else {
                switch 32767 {
                case 128:
                    var v28 uint = (v5 % ((v5) | 1))
                    _ = v28
                case 0:
                    v29 := GenS3{F0: 1, F1: ("" + "a"), F2: "foo bar"}
                    _ = v29
                }
                for fi30 := 0; fi30 < 13; fi30++ {
                    v7 = (v7 + 32767)
                    _ = fi30
                }
            }
            _ = len(v1)
            var v31 int16 = 32767
            _ = fi17
            _ = v31
        }
        var v32 int32 = 787
        v33 := GenS4{F0: (v4 % (((v4 * 2147483647)) | 1)), F1: (-(v7))}
        _ = v32
        _ = v33
    } else {
        v5 -= (v5 << ((255) & 63))
        if true {
            if false {
                if (true || true) {
                    switch (("a" + ("\n\t" + "hello")) + "") {
                    case "hello":
                        switch -9223372036854775808 {
                        case 42:
                            for fi34 := 0; fi34 < 2; fi34++ {
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    switch fi34 {
                                    case 2147483647:
                                        v36 := GenS5{F0: 42, F1: 42, F2: 42}
                                        _ = v36
                                    case 705:
                                        fmt.Println(fi35, fi34)
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (32768 <= (fi35 | fi34)) {
                                        if true {
                                            var v37 uint16 = 1
                                            _ = v37
                                        }
                                        fmt.Println(v7, v4)
                                    }
                                    _ = fi35
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi34
                            }
                        case 127:
                            fmt.Println(v5, v4, v7)
                        }
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 588
                        }
                    case "\n\t":
                        for fi38 := 0; fi38 < 6; fi38++ {
                            for fi39 := 0; fi39 < 1; fi39++ {
                                switch -9223372036854775808 {
                                case -2147483648:
                                    var v40 int8 = -128
                                    _ = v40
                                case 42:
                                    var v41 uint = 255
                                    _ = v41
                                }
                                v5 = 0
                                var v42 float64 = -2.25
                                _ = fi39
                                _ = v42
                            }
                            v43 := GenS6{F0: 189, F1: 605}
                            v44 := v3[v4]
                            _ = fi38
                            _ = v43
                            _ = v44
                        }
                    }
                }
                v45 := GenS1{F0: -31, F1: v4, F2: (v7 | v7)}
                _ = v45
            }
        } else {
            var v46 uint16 = 357
            _ = v46
        }
    }
    v47 := GenS6{F0: 735, F1: 0}
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 uint16 = 0
        _ = fi48
        _ = v49
    }
    if (65535 == 42) {
        v5 -= (^(v5))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (222 <= -1) {
        switch 216 {
        case 32768:
            for fi50 := 0; fi50 < 19; fi50++ {
                v51 := GenS7{F0: -9223372036854775808, F1: v4, F2: -184}
                if len(v2) > 0 {
                    v2[0] = 2
                }
                fmt.Println(fi50)
                _ = fi50
                _ = v51
            }
        }
    } else {
        if (true || ((v5 >> ((v5) & 63)) > v5)) {
            if (-714 >= -2147483648) {
                v5 = (v5 >> (((v5 * (v5 ^ (^(v5))))) & 63))
                var v52 float64 = 0.0
                for fi53 := 0; fi53 < 16; fi53++ {
                    for fi54 := 0; fi54 < 4; fi54++ {
                        v55 := GenS4{F0: v4, F1: (^(v7))}
                        v56 := GenS8{F0: 307}
                        var v57 int16 = v7
                        _ = fi54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 10; fi58++ {
                        if ("a" > "hello") {
                            v4 = v4
                        } else {
                            var v59 int8 = -23
                            var v60 int64 = 0
                            v52 = 0.1
                            _ = v59
                            _ = v60
                        }
                        var v61 int8 = 71
                        var v62 uint64 = uint64(v4)
                        _ = fi58
                        _ = v61
                        _ = v62
                    }
                    _ = fi53
                }
                _ = v52
            }
        }
        fmt.Println(v4)
    }
    fmt.Println(v5, v5, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
