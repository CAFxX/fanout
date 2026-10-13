package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 string; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 string }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint; F1 int64; F2 int32 }
type GenS9 struct { F0 string; F1 string; F2 int8 }
type GenS10 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -23
    v3 := make(map[uint32]bool)
    var v4 uint64 = 0
    var v5 uint32 = 1
    v6 := make(map[uint32]uint32)
    if false {
        v7 := GenS1{F0: 1, F1: 127, F2: 42}
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi8, v2)
            for fi9 := 0; fi9 < 15; fi9++ {
                if ((^(v5)) < (v5 << ((0) & 63))) {
                    v10 := GenS2{F0: "\n\t", F1: -2147483648, F2: ""}
                    fmt.Println(v4)
                    _ = v10
                }
                if true {
                    var v11 int8 = -62
                    _ = v11
                } else {
                    v12 := GenS2{F0: ("foo bar" + "\n\t"), F1: -2147483648, F2: "a"}
                    var v13 string = "foo bar"
                    fmt.Println(fi9)
                    _ = v12
                    _ = v13
                }
                v14 := GenS3{F0: (v4 ^ v4), F1: 2, F2: ((("hello" + ("\n\t" + "a")) + (("hello" + "hello") + "foo bar")) + ("a" + ("a" + ("" + ("" + "a")))))}
                switch (fi9 << ((389) & 63)) {
                case -9223372036854775808:
                    var v15 int16 = (-(v2))
                    _ = v15
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        if ((("\n\t" + "") + "foo bar") != "") {
                            v6[(v5 * v5)] = 533
                        } else {
                            v16 := GenS3{F0: 111, F1: 42, F2: "\n\t"}
                            var v17 int16 = v2
                            if false {
                                if false {
                                    if (!false) {
                                        v18 := []int{fi8, 9223372036854775807, 42, -2147483648}
                                        _ = v18
                                    } else {
                                        switch fi9 {
                                        case -2147483648:
                                            switch (("\n\t" + ("hello" + ("a" + "hello"))) + "a") {
                                            case "":
                                                switch fi9 {
                                                case 9223372036854775807:
                                                    fmt.Println(v2)
                                                case -482:
                                                    if true {
                                                        v19 := []int{4294967295}
                                                        var v20 float32 = -2.25
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        v5 += v5
                                                        var v21 int16 = (v2 ^ v2)
                                                        if (!false) {
                                                            v6[v5] = (^(v5))
                                                            v22 := GenS1{F0: 1, F1: 127, F2: 249}
                                                            v2 += 42
                                                            _ = v22
                                                        }
                                                        fmt.Println(v5, v17)
                                                        _ = v21
                                                    }
                                                case 127:
                                                    fmt.Println(fi8)
                                                default:
                                                    if (128 < 2) {
                                                        var v23 uint8 = 2
                                                        _ = v23
                                                    }
                                                }
                                                v24 := GenS1{F0: 1, F1: -1, F2: 255}
                                                fmt.Println(v2, v2, fi9)
                                                v2 += v2
                                                _ = v24
                                            case "foo bar":
                                                v2 = -338
                                            case "a":
                                                if false {
                                                    _ = len(v3)
                                                    var v25 uint = 2
                                                    _ = v25
                                                } else {
                                                    v4 = v4
                                                }
                                            }
                                        case 0:
                                            v26 := GenS3{F0: 4294967295, F1: 255, F2: ""}
                                            _ = v26
                                        }
                                        v27 := []int{(fi9 & fi9), (fi9 & fi8), (fi8 << ((65535) & 63))}
                                        var v28 float64 = (0.0 - (1.5 - float64(1.5)))
                                        _ = v27
                                        _ = v28
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if (65535 == 32767) {
                                            var v29 bool = (32767 != 65535)
                                            v2 = (-(v2))
                                            _ = v29
                                        } else {
                                            v4 = v4
                                            if false {
                                                switch (fi8 / (((fi9 * fi9)) | 1)) {
                                                case -454:
                                                    var v30 int8 = 1
                                                    _ = v30
                                                case 9223372036854775807:
                                                    if (v17 > v2) {
                                                        v31 := GenS1{F0: -26, F1: 42, F2: 659}
                                                        v2 = -984
                                                        _ = v31
                                                    }
                                                default:
                                                    if (v4 != (v4 << ((0) & 63))) {
                                                        v32 := GenS3{F0: (v4 / (((v4 - 0)) | 1)), F1: uint8(v5), F2: "foo bar"}
                                                        var v33 float32 = 1.5
                                                        if ((true || ((^(v4)) <= (v4 % ((v4) | 1)))) && (!false)) {
                                                            switch 0 {
                                                            case 32767:
                                                                v2 += v17
                                                            default:
                                                                if (0 == 255) {
                                                                    v34 := v6[2]
                                                                    _ = v34
                                                                } else {
                                                                    var v35 uint32 = (v5 % (((v5 * v5)) | 1))
                                                                    if (0.1 != (float64(v33) / 1.5)) {
                                                                        if false {
                                                                            v5 += (v35 << ((18446744073709551615) & 63))
                                                                            v17 = (v17 - 677)
                                                                        } else {
                                                                            v33 += (v33 + (float32(3.14159) / v33))
                                                                        }
                                                                        switch 567 {
                                                                        case 66:
                                                                            delete(v6, (v35 + (v35 % ((v35) | 1))))
                                                                            fmt.Println(v5, v17, v33)
                                                                            v36 := GenS4{F0: (fi9 - fi8), F1: uint8(v4), F2: (v4 - v4)}
                                                                            _ = v36
                                                                        default:
                                                                            v37 := GenS5{F0: 2}
                                                                            _ = v37
                                                                        }
                                                                        v5 = v35
                                                                    }
                                                                    _ = v35
                                                                }
                                                            }
                                                            if (!("" > "")) {
                                                                _ = len(v3)
                                                                if false {
                                                                    var v38 uint16 = 65535
                                                                    _ = v38
                                                                }
                                                                var v39 uint = 0
                                                                fmt.Println(v33, v5)
                                                                _ = v39
                                                            } else {
                                                                if true {
                                                                    switch -953 {
                                                                    case 128:
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            var v40 string = "\n\t"
                                                                            _ = v40
                                                                        case "\n\t":
                                                                            v4 = (^(v4))
                                                                        default:
                                                                            var v41 int32 = 0
                                                                            _ = v41
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            var v42 uint32 = (v5 >> ((18446744073709551615) & 63))
                                                                            _ = v42
                                                                        }
                                                                    }
                                                                    v5 = (v5 | v5)
                                                                    if false {
                                                                        fmt.Println(v33, v2, v17)
                                                                        v43 := GenS6{F0: 0, F1: v4}
                                                                        var v44 int16 = (v17 + (-(v2)))
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    var v45 int8 = 1
                                                                    v46 := GenS1{F0: 42, F1: (-(v45)), F2: 829}
                                                                    var v47 float64 = (0.1 - 1e10)
                                                                    v17 = 42
                                                                    _ = v45
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                var v48 int16 = (v17 ^ v2)
                                                                _ = v48
                                                            }
                                                            v17 = (v17 | (v17 + (v17 / ((v2) | 1))))
                                                        } else {
                                                            v49 := []int{fi9}
                                                            _ = v49
                                                        }
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    if false {
                                                        if true {
                                                            v50 := GenS3{F0: v4, F1: 1, F2: "hello"}
                                                            var v51 bool = (true && true)
                                                            var v52 float32 = 0.0
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        v4 = (^(v4))
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v17, v4, fi9)
                                            }
                                            v6[v5] = v5
                                        }
                                    default:
                                        v4 = 4294967295
                                    }
                                }
                            }
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        v53 := []int{(fi9 + fi8)}
                        _ = v53
                    }
                case 128:
                    v54 := v1[-1]
                    _ = v54
                default:
                    var v55 float64 = (((-2.25 / 1e10) / (0.1 / 1e10)) - (3.14159 - float64(3.14159)))
                    _ = v55
                }
                v56 := GenS7{F0: -346, F1: 65535, F2: uint32(v4)}
                switch fi8 {
                case -500:
                    fmt.Println(v5)
                }
                _ = fi9
                _ = v14
                _ = v56
            }
            _ = fi8
        }
        v2 += -1
        _ = v7
    }
    var v57 bool = (!false)
    v58 := GenS8{F0: 18446744073709551615, F1: 9223372036854775807, F2: -2147483648}
    for fi59 := 0; fi59 < 19; fi59++ {
        v5 += 4294967295
        if v57 {
            var v60 uint32 = (v5 ^ (v5 + uint32(v4)))
            fmt.Println(v5)
            _ = v60
        } else {
            fmt.Println(v57)
            var v61 int64 = -9223372036854775808
            _ = v61
        }
        _ = fi59
    }
    v57 = (!v57)
    var v62 bool = false
    switch -9223372036854775808 {
    case 296:
        v63 := []int{127, 32768}
        _ = v63
    case 32767:
        fmt.Println(v57)
    default:
        v62 = v57
    }
    v4 += 439
    delete(v6, 1)
    _ = len(v6)
    var v64 uint8 = 128
    v65 := GenS8{F0: 18446744073709551615, F1: 83, F2: -1}
    var v66 int16 = v2
    if (!true) {
        for fi67 := 0; fi67 < 3; fi67++ {
            v68 := GenS9{F0: "hello", F1: (("" + "foo bar") + "a"), F2: 42}
            _ = fi67
            _ = v68
        }
    }
    if v57 {
        v69 := GenS9{F0: ("hello" + "\n\t"), F1: "\n\t", F2: -1}
        _ = v69
    }
    v70 := GenS1{F0: 100, F1: 77, F2: 861}
    v71 := GenS10{F0: (v5 % ((v5) | 1)), F1: (2 != 127)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v62
    _ = v64
    _ = v65
    _ = v66
    _ = v70
    _ = v71
}
