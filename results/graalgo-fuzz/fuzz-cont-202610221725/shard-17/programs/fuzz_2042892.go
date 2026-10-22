package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 string; F1 int8; F2 int64 }
type GenS5 struct { F0 bool; F1 uint16; F2 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int32; F2 int64 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 bool, p1 int32) int8 {
    var v7 int = 1
    v8 := GenS1{F0: ("" + ("foo bar" + "\n\t")), F1: 42, F2: p0}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}
func genFn1() uint {
    v9 := []int{1, 783}
    if false {
        var v10 int32 = -1
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v10
    }
    _ = v9
    return 544
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 574
    v3 := make(map[string]bool)
    var v4 string = "a"
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    v11 := []int{}
    var v12 int32 = 42
    switch ("foo bar" + v4) {
    case "":
        fmt.Println(v1, v12, v2)
    case "a":
        v13 := GenS1{F0: "hello", F1: 0, F2: false}
        _ = v13
    case "hello":
        var v14 int64 = 1
        _ = v14
    }
    v15 := v5[938]
    v12 += v12
    for fi16 := 0; fi16 < 15; fi16++ {
        if ((false && true) || (0.0 == 1e10)) {
            _ = len(v6)
            for fi17 := 0; fi17 < 20; fi17++ {
                v2 -= v2
                var v18 int32 = -1
                v4 = ("hello" + v4)
                switch v4 {
                case "a":
                    if (false || true) {
                        var v19 float64 = 0.1
                        v4 = (("\n\t" + (v4 + "foo bar")) + v4)
                        var v20 bool = true
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi17
                _ = v18
            }
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int32 = v12
            v22 = v12
            if (!true) {
                fmt.Println(v4)
                v2 += v2
            } else {
                switch v4 {
                case "a":
                    switch ("a" + ("foo bar" + v4)) {
                    case "foo bar":
                        var v23 int = fi21
                        _ = v23
                    case "hello":
                        v1 += v1
                    case "a":
                        var v24 string = "\n\t"
                        _ = v24
                    }
                }
            }
            if (1.5 < ((1.5 / float64(3.14159)) * 1e10)) {
                switch (^(v15)) {
                case 128:
                    if (false && false) {
                        v25 := GenS1{F0: v4, F1: -962, F2: (true || (!(!false)))}
                        v26 := GenS2{F0: false, F1: false}
                        _ = v25
                        _ = v26
                    }
                case -226:
                    fmt.Println(v4, v22)
                case 127:
                    if false {
                        switch "a" {
                        case "foo bar":
                            if (((float64((-2.25 - 3.14159)) - (0.0 * 1e10)) == 0.0) || (-2.25 > (float64(1e10) / ((((1.5 / 1.5) / float64(3.14159)) / 0.1) - (((3.14159 - 1.5) * 0.1) - 0.1))))) {
                                v15 -= -9223372036854775808
                                var v27 bool = (3.14159 <= float32(0.0))
                                _ = v27
                            }
                            switch "hello" {
                            case "\n\t":
                                if ((!true) || (32767 > 42)) {
                                    v28 := GenS2{F0: false, F1: (!true)}
                                    switch v4 {
                                    case "\n\t":
                                        if len(v11) > 0 {
                                            v11[0] = (fi21 * -2147483648)
                                        }
                                    case "":
                                        v2 = v2
                                        var v29 int64 = -407
                                        _ = v29
                                    default:
                                        v1 = 1
                                    }
                                    _ = v28
                                }
                            case "foo bar":
                                v30 := GenS3{F0: v15, F1: 2}
                                _ = v30
                            default:
                                switch v15 {
                                case 32767:
                                    fmt.Println(v4, fi21)
                                }
                            }
                            switch int(v22) {
                            case 439:
                                switch (((v4 + (v4 + ("" + v4))) + v4) + (v4 + (v4 + v4))) {
                                case "foo bar":
                                    var v31 float32 = ((1.5 - 1e10) * 0.1)
                                    _ = v31
                                case "a":
                                    _ = len(v11)
                                case "hello":
                                    v12 = v22
                                }
                            case 32767:
                                var v32 string = ""
                                _ = v32
                            case -173:
                                v22 = -2147483648
                            }
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            var v33 float32 = 0.1
                            switch ((v4 + v4) + (("foo bar" + v4) + v4)) {
                            case "":
                                fmt.Println(v4, fi16, v2)
                            }
                            v34 := v5[-1]
                            _ = v33
                            _ = v34
                        default:
                            var v35 bool = ((fi21 % ((fi16) | 1)) == v15)
                            _ = v35
                        }
                        v1 = v1
                    }
                default:
                    v5[(fi21 | fi16)] = (fi16 % ((-998) | 1))
                }
                var v36 int = (fi16 | fi16)
                _ = v36
            } else {
                v22 -= v12
            }
            _ = fi21
            _ = v22
        }
        if true {
            v37 := GenS4{F0: v4, F1: 49, F2: -9223372036854775808}
            v38 := v5[(v15 / ((-788) | 1))]
            _ = v37
            _ = v38
        }
        _ = fi16
    }
    if true {
        var v39 int8 = genFn0(true, v12)
        var v40 string = "a"
        _ = v39
        _ = v40
    } else {
        v15 -= v15
        for fi41 := 0; fi41 < 11; fi41++ {
            v42 := GenS5{F0: true, F1: (v2 << ((genFn1()) & 63)), F2: 1}
            v6[(v15 - 4294967295)] = fi41
            _ = fi41
            _ = v42
        }
        if true {
            if true {
                var v43 int64 = -1
                _ = v43
            }
        } else {
            fmt.Println(v2)
        }
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        switch (v4 + v4) {
        case "":
            for fi45 := 0; fi45 < 17; fi45++ {
                v46 := GenS1{F0: v4, F1: 42, F2: false}
                fmt.Println(v1, v12, v1)
                var v47 uint32 = (^(v1))
                _ = fi45
                _ = v46
                _ = v47
            }
        default:
            for fi48 := 0; fi48 < 4; fi48++ {
                var v49 float32 = -2.25
                if (true || (fi48 > fi48)) {
                    if false {
                        v1 += (v1 * (v1 + v1))
                        switch "\n\t" {
                        case "a":
                            switch v4 {
                            case "hello":
                                var v50 int64 = -779
                                _ = v50
                            case "a":
                                switch (fi48 >> ((65535) & 63)) {
                                case 127:
                                    v51 := GenS5{F0: (!false), F1: 65535, F2: 0}
                                    _ = v51
                                case 32767:
                                    var v52 uint8 = 42
                                    _ = v52
                                case -9223372036854775808:
                                    v53 := GenS6{F0: "\n\t"}
                                    _ = v53
                                default:
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(fi48, v2, v1)
                                    case "":
                                        var v54 uint8 = 255
                                        _ = v54
                                    case "foo bar":
                                        var v55 bool = (!true)
                                        _ = v55
                                    default:
                                        v56 := GenS2{F0: (-1 == 0), F1: true}
                                        _ = v56
                                    }
                                }
                            case "":
                                if (737 == int64(v1)) {
                                    if true {
                                        v15 += (v15 * fi44)
                                        switch "" {
                                        case "foo bar":
                                            v12 += v12
                                        case "a":
                                            v15 += -832
                                        case "":
                                            fmt.Println(fi44)
                                        }
                                    } else {
                                        v11 = append(v11, fi48)
                                        switch "hello" {
                                        case "hello":
                                            v57 := v6[fi44]
                                            _ = v57
                                        }
                                        v58 := GenS7{F0: v4, F1: (v12 | v12), F2: -9223372036854775808}
                                        _ = v58
                                    }
                                }
                            }
                        case "foo bar":
                            if false {
                                var v59 int16 = int16(fi48)
                                v60 := GenS8{F0: 255, F1: 42}
                                v2 = v2
                                _ = v59
                                _ = v60
                            } else {
                                v15 = (v15 << ((1) & 63))
                            }
                        }
                        var v61 uint = 637
                        _ = v61
                    } else {
                        var v62 int8 = int8(v15)
                        _ = v62
                    }
                    fmt.Println(fi44)
                    delete(v5, 2147483647)
                    _ = len(v11)
                }
                var v63 uint16 = v2
                fmt.Println(v12)
                _ = fi48
                _ = v49
                _ = v63
            }
        }
        fmt.Println(fi44, v12, v4)
        _ = fi44
    }
    v1 = v1
    if false {
        for fi64 := 0; fi64 < 13; fi64++ {
            v4 = v4
            _ = fi64
        }
    }
    fmt.Println(v4)
    if true {
        v2 = (v2 * v2)
    }
    var v65 int64 = -49
    v66 := GenS9{F0: 628}
    fmt.Println(v1)
    for fi67 := 0; fi67 < 5; fi67++ {
        v15 += v15
        v2 = (v2 << ((uint(v2)) & 63))
        var v68 uint64 = 0
        _ = fi67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
