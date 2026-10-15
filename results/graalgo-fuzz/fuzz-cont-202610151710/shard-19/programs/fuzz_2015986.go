package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 string }
func genFn0(p0 int64, p1 int64, p2 uint8) uint8 {
    switch 279 {
    case 0:
        var v6 uint64 = 2
        _ = v6
    default:
        v7 := []int{-295}
        _ = v7
    }
    p2 += (p2 | (p2 - (^(p2))))
    v8 := GenS1{F0: 42, F1: 733, F2: 107}
    switch -946 {
    case 363:
        var v9 int32 = 42
        _ = v9
    case -2147483648:
        p1 -= p0
    case -1:
        if (128 >= -285) {
            v10 := GenS2{F0: (true || true), F1: 339, F2: 4294967295}
            _ = v10
        } else {
            var v11 int64 = p1
            v12 := GenS3{F0: (3.14159 < 1e10)}
            if (!true) {
                v13 := GenS4{F0: 917, F1: 42}
                v14 := []int{}
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 string = "foo bar"
                    for fi17 := 0; fi17 < 4; fi17++ {
                        p0 += (p0 - 42)
                        v18 := GenS3{F0: (65535 != 448)}
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(p1, v16, fi15)
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
            _ = v11
            _ = v12
        }
    default:
        v19 := GenS5{F0: 14}
        _ = v19
    }
    v20 := []int{-75, 32768, 233, -481}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v8
    return 127
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]string)
    var v3 bool = (-159 >= -888)
    var v4 int16 = 1000
    v5 := []int{}
    var v21 int16 = v4
    v1 = 264
    if v3 {
        fmt.Println(v1, v21)
        v2[""] = "\n\t"
    }
    if (1 > 42) {
        var v22 int32 = -97
        _ = v22
    } else {
        fmt.Println(v3)
        var v23 bool = v3
        v21 = 32767
        _ = v23
    }
    fmt.Println(v1, v4, v21)
    v21 -= v21
    if (0.0 >= ((3.14159 - (((3.14159 * 1.5) - 3.14159) * 1e10)) - 3.14159)) {
        v5 = append(v5, 42)
    } else {
        var v24 float64 = 1e10
        fmt.Println(v21)
        _ = v24
    }
    if (!(!true)) {
        var v25 uint8 = 1
        switch "a" {
        case "\n\t":
            switch -1 {
            case 32768:
                var v26 int16 = (v21 ^ 0)
                _ = v26
            }
        case "":
            if ((^(v25)) != 4) {
                if v3 {
                    v1 -= v1
                }
            } else {
                if true {
                    v1 += 42
                    var v27 int32 = -897
                    var v28 uint8 = 128
                    _ = v27
                    _ = v28
                }
                switch "foo bar" {
                case "foo bar":
                    switch ((("hello" + "foo bar") + "\n\t") + "a") {
                    case "\n\t":
                        var v29 int = 32768
                        _ = v29
                    default:
                        v30 := GenS4{F0: -2147483648, F1: (v25 + 128)}
                        _ = v30
                    }
                case "a":
                    if ((v1 % (((v1 - v1)) | 1)) == 2) {
                        for fi31 := 0; fi31 < 5; fi31++ {
                            var v32 int = (fi31 ^ 2147483647)
                            if false {
                                v33 := GenS4{F0: 2147483647, F1: v25}
                                _ = v33
                            }
                            _ = fi31
                            _ = v32
                        }
                        v34 := GenS2{F0: (1 < 99), F1: 65535, F2: 2147483647}
                        _ = v34
                    }
                case "":
                    v3 = (869 <= 386)
                default:
                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                    case "a":
                        v35 := GenS6{F0: 100, F1: -396}
                        _ = v35
                    case "":
                        v36 := GenS7{F0: (-79 < 42), F1: 100}
                        _ = v36
                    case "hello":
                        v37 := GenS2{F0: (!(0 == 0)), F1: 32767, F2: 508}
                        _ = v37
                    default:
                        switch 1 {
                        case -9223372036854775808:
                            fmt.Println(v3, v21, v1)
                        case -115:
                            v38 := GenS8{F0: 741}
                            _ = v38
                        case 9223372036854775807:
                            v39 := v2[("hello" + "hello")]
                            _ = v39
                        default:
                            var v40 int16 = (v4 / ((619) | 1))
                            _ = v40
                        }
                    }
                }
            }
        case "foo bar":
            for fi41 := 0; fi41 < 12; fi41++ {
                fmt.Println(v1)
                _ = fi41
            }
        default:
            v3 = v3
        }
        _ = v25
    } else {
        v42 := GenS8{F0: 65535}
        v43 := GenS9{F0: v1, F1: 9223372036854775807}
        _ = v42
        _ = v43
    }
    if len(v5) > 0 {
        v5[0] = 42
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        if (0.0 < (1.5 + (3.14159 / 1e10))) {
            if len(v5) > 0 {
                v5[0] = 1
            }
            for fi45 := 0; fi45 < 4; fi45++ {
                if v3 {
                    var v46 uint8 = 127
                    _ = v46
                }
                var v47 int8 = -1
                var v48 uint64 = v1
                _ = fi45
                _ = v47
                _ = v48
            }
            v5 = append(v5, (^(fi44)))
        } else {
            for fi49 := 0; fi49 < 2; fi49++ {
                if v3 {
                    if len(v5) > 0 {
                        v5[0] = (fi44 & (fi44 % (((fi49 << ((4294967295) & 63))) | 1)))
                    }
                    v50 := GenS10{F0: 4294967295, F1: 127, F2: (v1 >> ((676) & 63))}
                    if v3 {
                        v1 = (^(v1))
                        var v51 int = fi44
                        _ = v51
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        delete(v2, ("\n\t" + ((("foo bar" + "foo bar") + "hello") + "hello")))
                    }
                    var v52 uint32 = 1
                    var v53 int32 = 0
                    if true {
                        if false {
                            var v54 string = ("a" + "\n\t")
                            switch (-(fi49)) {
                            case -694:
                                v52 -= v52
                            default:
                                v21 = v21
                            }
                            _ = v54
                        } else {
                            var v55 int64 = 1
                            v4 -= -744
                            _ = v55
                        }
                        switch "\n\t" {
                        case "":
                            var v56 bool = true
                            _ = v56
                        case "\n\t":
                            v4 -= 42
                        }
                        v57 := GenS10{F0: v52, F1: -1, F2: (v1 - (v1 - v1))}
                        _ = v57
                    } else {
                        v58 := GenS11{F0: genFn0(42, -213, genFn0(-810, 394, 233))}
                        v4 -= (v4 | v21)
                        v1 -= v1
                        _ = v58
                    }
                    _ = len(v2)
                    _ = v50
                    _ = v52
                    _ = v53
                } else {
                    delete(v2, "foo bar")
                    var v59 int64 = 42
                    _ = v59
                }
                _ = fi49
            }
        }
        v60 := GenS12{F0: 937, F1: 1}
        _ = fi44
        _ = v60
    }
    var v61 float32 = 1e10
    for fi62 := 0; fi62 < 16; fi62++ {
        if (!true) {
            v3 = (v3 || v3)
            v61 = v61
            v1 -= v1
        } else {
            v61 += -2.25
            var v63 uint16 = 327
            v64 := GenS13{F0: "\n\t"}
            _ = v63
            _ = v64
        }
        _ = fi62
    }
    var v65 int8 = 42
    var v66 float64 = (1.5 + 0.1)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v65
    _ = v66
}
