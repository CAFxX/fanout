package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int; F1 int; F2 uint16 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int64; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint8 }
func genFn0(p0 bool, p1 uint16) uint64 {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: -8, F1: 0, F2: 42}
        if p0 {
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{(fi7 - fi5), fi5}
                switch (fi7 / ((-786) | 1)) {
                case 9223372036854775807:
                    var v9 float32 = float32((3.14159 + 0.0))
                    _ = v9
                case -9223372036854775808:
                    p1 = p1
                case -1:
                    v10 := GenS2{F0: (fi7 & fi7)}
                    p0 = p0
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v11 uint32 = 4294967295
                    _ = len(v8)
                    v12 := GenS1{F0: 1, F1: 255, F2: 4294967295}
                    _ = v10
                    _ = v11
                    _ = v12
                default:
                    if len(v8) > 0 {
                        v8[0] = (fi7 - (fi5 + 42))
                    }
                }
                v13 := GenS3{F0: 32767}
                _ = fi7
                _ = v13
                _ = v8
            }
            v14 := []int{}
            v15 := GenS2{F0: fi5}
            _ = v14
            _ = v15
        }
        _ = fi5
        _ = v6
    }
    var v16 bool = (!(1 >= 894))
    _ = p0
    _ = p1
    _ = v16
    return 0
}
func genFn1() int64 {
    var v17 uint32 = 42
    if true {
        var v18 int32 = -2147483648
        _ = v18
    }
    var v19 uint16 = 898
    _ = v17
    _ = v19
    return -9223372036854775808
}

func main() {
    v1 := []string{"", "", "foo bar", ""}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 int = 9223372036854775807
    fmt.Println(v4)
    var v20 uint16 = 956
    v4 += (v4 & v4)
    for fi21 := 0; fi21 < 2; fi21++ {
        v4 -= 9223372036854775807
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int32 = -1
            v23 -= (v23 << ((18446744073709551615) & 63))
            _ = fi22
            _ = v23
        }
        var v24 string = ("hello" + (("hello" + "a") + ""))
        _ = fi21
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v20 = (v20 % ((v20) | 1))
        v4 -= 127
        for fi26 := 0; fi26 < 4; fi26++ {
            switch (fi26 << ((447) & 63)) {
            case -909:
                if (!(477 > fi26)) {
                    v20 += (v20 * 2)
                }
            default:
                var v27 bool = ((815 < 65535) && false)
                _ = v27
            }
            fmt.Println(v20, fi25)
            v4 = 800
            if false {
                var v28 float64 = ((-2.25 * float64((3.14159 * 3.14159))) * 3.14159)
                _ = v28
            }
            _ = fi26
        }
        _ = fi25
    }
    v29 := GenS4{F0: v4, F1: 1, F2: 226}
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS5{F0: -2147483648, F1: v4, F2: v20}
        switch -138 {
        case -761:
            var v32 int = fi30
            _ = v32
        case -903:
            switch "\n\t" {
            case "a":
                fmt.Println(fi30, fi30, fi30)
            case "foo bar":
                var v33 uint8 = 42
                _ = v33
            default:
                for fi34 := 0; fi34 < 19; fi34++ {
                    delete(v2, ("hello" + "a"))
                    if (4294967295 == 2) {
                        if true {
                            v1 = append(v1, ("foo bar" + "\n\t"))
                        } else {
                            v4 = (v4 / (((fi34 ^ 4294967295)) | 1))
                            v20 -= (v20 & (v20 | 2))
                        }
                    }
                    if (1 > 1) {
                        v35 := GenS4{F0: 127, F1: 0, F2: 2}
                        _ = v35
                    } else {
                        var v36 bool = true
                        _ = v36
                    }
                    if ((1.5 * 0.1) < (((3.14159 / float64((1e10 / 0.1))) / 1.0) - (float64((0.1 - -2.25)) + (3.14159 * float64(-2.25))))) {
                        v4 -= 595
                        v4 = v4
                        if true {
                            v37 := v2[("a" + "")]
                            _ = v37
                        }
                        v4 += v4
                        if true {
                            switch (v4 * fi34) {
                            case 983:
                                var v38 float64 = 1.5
                                _ = v38
                            default:
                                switch v4 {
                                case -9223372036854775808:
                                    var v39 float64 = (((1e10 / 1.0) * (float64(1e10) * float64(0.1))) + float64(1.5))
                                    _ = v39
                                case 127:
                                    v40 := GenS5{F0: v4, F1: fi30, F2: 903}
                                    _ = v40
                                case 32767:
                                    v20 -= (^(v20))
                                default:
                                    v20 = v20
                                }
                            }
                            switch ("\n\t" + "\n\t") {
                            case "foo bar":
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                    if (true || (true || (false || (!((^(fi30)) != (fi34 | (v4 & (fi34 / (((fi30 << ((18446744073709551615) & 63))) | 1))))))))) {
                                        fmt.Println(fi34)
                                        v20 += 0
                                        v41 := GenS6{F0: fi30, F1: v20}
                                        _ = len(v1)
                                        _ = v41
                                    } else {
                                        switch "" {
                                        case "a":
                                            v42 := GenS3{F0: 0}
                                            _ = v42
                                        case "foo bar":
                                            if (2 == 904) {
                                                fmt.Println(fi34, fi34, fi34)
                                                v43 := GenS4{F0: 32767, F1: 0, F2: 1}
                                                _ = v43
                                            } else {
                                                v4 += fi30
                                            }
                                        }
                                        v4 = 4294967295
                                    }
                                    var v44 int = fi30
                                    switch (fi30 >> ((18446744073709551615) & 63)) {
                                    case -9223372036854775808:
                                        v4 = 42
                                    case 349:
                                        var v45 uint32 = 1
                                        _ = v45
                                    }
                                    if true {
                                        if true {
                                            var v46 int64 = 0
                                            var v47 float32 = 0.1
                                            _ = v46
                                            _ = v47
                                        }
                                        fmt.Println(v20)
                                        var v48 int = 379
                                        v20 -= v20
                                        _ = v48
                                    } else {
                                        if true {
                                            v49 := GenS7{F0: 1, F1: -1, F2: -32768}
                                            var v50 int64 = 42
                                            _ = v49
                                            _ = v50
                                        } else {
                                            v51 := GenS1{F0: 127, F1: 97, F2: 4294967295}
                                            fmt.Println(fi34, v4, v20)
                                            var v52 uint16 = v20
                                            _ = v51
                                            _ = v52
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v53 := GenS5{F0: v44, F1: fi34, F2: (v20 - (v20 * 255))}
                                    _ = v44
                                    _ = v53
                                }
                            case "\n\t":
                                v54 := GenS8{F0: true}
                                _ = v54
                            case "hello":
                                var v55 uint16 = v20
                                _ = v55
                            }
                        } else {
                            v20 = (v20 / (((v20 ^ v20)) | 1))
                            if true {
                                switch (fi30 >> ((255) & 63)) {
                                case 1:
                                    v20 += (^(v20))
                                case 0:
                                    v20 -= (v20 * v20)
                                case 996:
                                    v56 := v3[(v4 % ((801) | 1))]
                                    _ = v56
                                default:
                                    switch "hello" {
                                    case "hello":
                                        var v57 int32 = -1
                                        _ = v57
                                    case "a":
                                        v58 := GenS7{F0: 65535, F1: -1, F2: 0}
                                        _ = v58
                                    default:
                                        fmt.Println(v4, v20)
                                    }
                                }
                                if (1 >= 42) {
                                    var v59 int64 = -1
                                    _ = v59
                                }
                                var v60 int8 = -128
                                _ = v60
                            } else {
                                v4 = 241
                                var v61 int64 = 1
                                _ = v61
                            }
                        }
                    } else {
                        v20 = v20
                        var v62 uint16 = v20
                        _ = v62
                    }
                    _ = fi34
                }
            }
        default:
            var v63 uint16 = v20
            _ = v63
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v4, v4)
    if false {
        switch 9223372036854775807 {
        case -1:
            v1 = append(v1, ("" + ""))
        default:
            var v64 int16 = -32768
            _ = v64
        }
        if len(v1) > 0 {
            v1[0] = "a"
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v20 = (^(v20))
        switch "\n\t" {
        case "foo bar":
            if (false && true) {
                for fi65 := 0; fi65 < 1; fi65++ {
                    fmt.Println(v20, v20, v4)
                    fmt.Println(v20, v20)
                    v66 := GenS3{F0: 0}
                    _ = fi65
                    _ = v66
                }
            }
        case "":
            v4 += (v4 & (v4 - -596))
        case "hello":
            v67 := GenS9{F0: -783, F1: 1}
            _ = v67
        default:
            if true {
                var v68 int64 = 53
                _ = v68
            } else {
                for fi69 := 0; fi69 < 15; fi69++ {
                    v20 += 355
                    delete(v2, ("\n\t" + ("" + "")))
                    _ = fi69
                }
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                var v70 int8 = -22
                _ = v70
            }
        }
        for fi71 := 0; fi71 < 17; fi71++ {
            for fi72 := 0; fi72 < 3; fi72++ {
                switch v4 {
                case 42:
                    v20 = v20
                case -9223372036854775808:
                    fmt.Println(fi71, v20)
                case 128:
                    v4 = (v4 % (((fi71 ^ (fi71 / ((127) | 1)))) | 1))
                }
                v20 -= (v20 * v20)
                _ = fi72
            }
            _ = fi71
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
}
