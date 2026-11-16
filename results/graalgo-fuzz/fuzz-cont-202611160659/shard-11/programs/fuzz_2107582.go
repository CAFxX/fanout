package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 string; F1 uint; F2 uint }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint }
func genFn0(p0 int32, p1 int16, p2 int) string {
    p2 = (p2 - p2)
    if (!(p2 < (-(p2)))) {
        p1 = p1
    } else {
        p0 += 2147483647
        p2 = p2
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(p0, p1, p2)
            for fi8 := 0; fi8 < 9; fi8++ {
                fmt.Println(fi7)
                fmt.Println(fi7, p1, fi8)
                _ = fi8
            }
            _ = fi7
        }
    }
    var v9 uint32 = 717
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "foo bar"
}
func genFn1() int8 {
    var v10 uint32 = 0
    var v11 float32 = 0.1
    var v12 int16 = -1
    var v13 uint16 = uint16(v10)
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    v3 := make(map[string]bool)
    var v4 uint = 4294967295
    var v5 uint64 = uint64(v4)
    var v6 uint64 = v5
    v14 := []int{}
    v4 = (v4 ^ (v4 >> ((v4) & 63)))
    var v15 uint16 = 0
    v16 := GenS1{F0: 0, F1: 32767}
    if false {
        var v17 float64 = (1e10 / 3.14159)
        _ = v17
    }
    switch "\n\t" {
    case "\n\t":
        if (214 == -1) {
            v4 += v4
        } else {
            v4 = (v4 | 42)
        }
    case "foo bar":
        var v18 bool = true
        _ = v18
    case "hello":
        v19 := GenS1{F0: (v15 - (v15 * (v15 % ((v15) | 1)))), F1: 85}
        _ = v19
    default:
        v15 += (v15 - v15)
    }
    switch ("foo bar" + "hello") {
    case "hello":
        var v20 uint8 = 1
        _ = v20
    case "\n\t":
        v15 += (v15 / ((uint16(v4)) | 1))
    }
    var v21 int16 = -31
    for fi22 := 0; fi22 < 18; fi22++ {
        v15 = (v15 << (((v4 * (v4 - (v4 & v4)))) & 63))
        _ = fi22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        if true {
            if true {
                fmt.Println(v21, v21)
                if (!false) {
                    switch (^(fi23)) {
                    case -63:
                        switch (fi23 >> (((v4 | 204)) & 63)) {
                        case -635:
                            for fi24 := 0; fi24 < 9; fi24++ {
                                v5 += 238
                                v15 = v15
                                _ = fi24
                            }
                        case 43:
                            v25 := GenS2{F0: "\n\t", F1: (fi23 >> ((1) & 63))}
                            _ = v25
                        }
                    case -9223372036854775808:
                        if len(v14) > 0 {
                            v14[0] = 931
                        }
                    case -962:
                        fmt.Println(v2)
                    default:
                        switch ("" + "") {
                        case "foo bar":
                            v6 = v5
                        case "":
                            delete(v1, 4294967295)
                        case "a":
                            var v26 float32 = 1.5
                            _ = v26
                        }
                    }
                }
                switch (fi23 - 9223372036854775807) {
                case 4294967295:
                    switch 42 {
                    case -1:
                        if ((fi23 & fi23) > fi23) {
                            fmt.Println(v6)
                        } else {
                            var v27 float32 = (float32(0.1) + 0.0)
                            _ = v27
                        }
                    case -720:
                        var v28 float64 = 1.5
                        _ = v28
                    default:
                        switch (("foo bar" + "\n\t") + (("\n\t" + "foo bar") + ("a" + "hello"))) {
                        case "hello":
                            for fi29 := 0; fi29 < 6; fi29++ {
                                v30 := GenS1{F0: 42, F1: (^(v15))}
                                if (false && false) {
                                    switch (fi23 | (^(fi29))) {
                                    case -2147483648:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    case 2147483647:
                                        if ((v6 << ((v4) & 63)) >= (v6 % (((^(v6))) | 1))) {
                                            switch (fi29 & fi29) {
                                            case -214:
                                                if (!(1 >= 1)) {
                                                    v21 += (v21 * (v21 & 1))
                                                    if true {
                                                        v31 := GenS3{F0: v2, F1: 128}
                                                        var v32 uint = v4
                                                        v33 := GenS1{F0: 65535, F1: 2}
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        var v34 int32 = v2
                                                        _ = v34
                                                    }
                                                }
                                            case 32767:
                                                var v35 int16 = (v21 | (v21 >> (((v4 / (((v4 - 0)) | 1))) & 63)))
                                                _ = v35
                                            }
                                        } else {
                                            var v36 int16 = v21
                                            _ = v36
                                        }
                                    case 0:
                                        v6 -= v5
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        if (false || (65535 >= 0)) {
                                            if (false || true) {
                                                v37 := GenS1{F0: uint16(v4), F1: v15}
                                                switch "" {
                                                case "\n\t":
                                                    v14 = append(v14, fi29)
                                                case "":
                                                    fmt.Println(v2)
                                                case "hello":
                                                    var v38 uint16 = (v15 / ((v15) | 1))
                                                    _ = v38
                                                default:
                                                    var v39 float32 = 1e10
                                                    _ = v39
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch fi29 {
                                                    case -2147483648:
                                                        v40 := v3["foo bar"]
                                                        _ = v40
                                                    case -540:
                                                        v41 := GenS4{F0: 42, F1: -1}
                                                        _ = v41
                                                    case 2147483647:
                                                        if (v2 < v2) {
                                                            if (!true) {
                                                                var v42 int16 = v21
                                                                _ = v42
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    }
                                                case "hello":
                                                    if (!true) {
                                                        if true {
                                                            fmt.Println(fi29)
                                                            fmt.Println(v5, v4)
                                                        } else {
                                                            v21 -= (v21 ^ 32767)
                                                            v43 := GenS1{F0: v15, F1: v15}
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                            _ = v43
                                                        }
                                                        if (0.1 <= ((float32(0.1) * ((1e10 * float32(float64((float32(3.14159) / 1e10)))) / -2.25)) * ((1.5 - 0.0) * 0.0))) {
                                                            fmt.Println(v2, v15)
                                                            v4 = v4
                                                            v44 := GenS1{F0: (^(v15)), F1: v15}
                                                            switch fi23 {
                                                            case 0:
                                                                v15 = v15
                                                                v21 -= -346
                                                                if true {
                                                                    var v45 float32 = 0.0
                                                                    var v46 string = ("hello" + "hello")
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            case -270:
                                                                var v47 uint32 = 0
                                                                _ = v47
                                                            case -825:
                                                                var v48 uint32 = 2147483648
                                                                _ = v48
                                                            }
                                                            _ = v44
                                                        }
                                                    } else {
                                                        v4 = (^(v4))
                                                    }
                                                default:
                                                    var v49 uint = v4
                                                    _ = v49
                                                }
                                                _ = v37
                                            } else {
                                                if true {
                                                    var v50 int32 = v2
                                                    _ = v50
                                                } else {
                                                    if ((^(v6)) >= (v6 >> ((1) & 63))) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v51 int64 = 0
                                                            _ = v51
                                                        case "foo bar":
                                                            var v52 bool = (!(!false))
                                                            _ = v52
                                                        default:
                                                            v53 := v3["hello"]
                                                            _ = v53
                                                        }
                                                        v4 -= 255
                                                        switch (-(fi29)) {
                                                        case 32767:
                                                            var v54 int16 = v21
                                                            _ = v54
                                                        case -9223372036854775808:
                                                            switch 42 {
                                                            case -2147483648:
                                                                v4 += v4
                                                            case 4294967295:
                                                                if false {
                                                                    v55 := GenS5{F0: (("\n\t" + "foo bar") + "hello"), F1: (v4 + (v4 + v4)), F2: v4}
                                                                    if len(v14) > 0 {
                                                                        v14[0] = fi29
                                                                    }
                                                                    v5 -= v5
                                                                    v4 -= (v4 ^ (v4 | v4))
                                                                    v5 += 18446744073709551615
                                                                    fmt.Println(v2, v15)
                                                                    _ = v55
                                                                }
                                                            case -1:
                                                                v21 += v21
                                                            }
                                                        case -206:
                                                            if (9223372036854775807 != 1) {
                                                                v4 = v4
                                                                delete(v3, ("a" + "\n\t"))
                                                            }
                                                        default:
                                                            v5 -= 131
                                                        }
                                                    } else {
                                                        var v56 uint16 = 2
                                                        _ = v56
                                                    }
                                                    v57 := GenS1{F0: 42, F1: 736}
                                                    _ = v57
                                                }
                                            }
                                            fmt.Println(fi29)
                                            var v58 float64 = 1.5
                                            if false {
                                                var v59 uint8 = 255
                                                _ = v59
                                            } else {
                                                if ((v21 < (v21 | v21)) && false) {
                                                    fmt.Println(v6, fi29, v58)
                                                    v2 += (^(v2))
                                                    v6 -= 1
                                                    v21 = (v21 ^ -1)
                                                } else {
                                                    var v60 int = -275
                                                    _ = v60
                                                }
                                            }
                                            _ = v58
                                        }
                                    }
                                    v15 -= (^(v15))
                                }
                                if (!false) {
                                    var v61 int64 = 609
                                    var v62 int32 = (^(v2))
                                    if true {
                                        v6 -= (v5 & (v6 >> (((v4 & (v4 / (((v4 / ((255) | 1))) | 1)))) & 63)))
                                        if (128 == 0) {
                                            delete(v3, "foo bar")
                                        }
                                        fmt.Println(v15, v2, v61)
                                    } else {
                                        v2 -= -830
                                    }
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS6{F0: (v4 - uint(v5)), F1: 0}
                                    v64 := v3["hello"]
                                    v65 := GenS7{F0: 0}
                                    fmt.Println(fi23, v6, v6)
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                var v66 float32 = 1e10
                                var v67 float32 = (0.0 / ((v66 - (((1.5 + 1e10) + (v66 - float32(((0.1 - 1e10) * 3.14159)))) / v66)) / (v66 / 3.14159)))
                                _ = fi29
                                _ = v30
                                _ = v66
                                _ = v67
                            }
                        case "foo bar":
                            var v68 int32 = (v2 % (((v2 + 42)) | 1))
                            _ = v68
                        case "a":
                            v4 = 747
                        }
                    }
                }
            } else {
                var v69 uint32 = 1
                _ = v69
            }
            v4 = (^(v4))
        }
        v5 = 18446744073709551615
        _ = fi23
    }
    v70 := GenS7{F0: (v4 << (((v4 << ((18446744073709551615) & 63))) & 63))}
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
