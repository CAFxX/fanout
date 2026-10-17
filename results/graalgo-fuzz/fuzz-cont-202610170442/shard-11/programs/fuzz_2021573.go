package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 string }
func genFn0() int16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := []int{(fi4 - fi4), 9223372036854775807, (fi4 >> ((593) & 63)), fi4}
        var v6 string = "foo bar"
        v5 = append(v5, fi4)
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -128, F1: 42}
    _ = v7
    return -1
}
func genFn1() uint {
    var v8 uint16 = 42
    v8 -= (v8 + (^(v8)))
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint8 = 128
        var v11 string = (("" + "a") + "a")
        if (false && (23 >= 42)) {
            v11 = v11
            var v12 int8 = -1
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -365
    v3 := make(map[uint32]string)
    v13 := []int{-950, -292, 0, 4294967295}
    v2 = -1
    for fi14 := 0; fi14 < 12; fi14++ {
        v2 += 792
        _ = len(v1)
        _ = fi14
    }
    if true {
        var v15 int8 = -1
        _ = v15
    }
    _ = len(v13)
    v16 := GenS2{F0: ("" + "a"), F1: 1}
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: 42, F1: 2}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!((float64(-2.25) * (1.5 + float64(1e10))) == 0.1)) {
            if (true && (!(((v2 & (v2 + 2147483647)) < (v2 << ((genFn1()) & 63))) && true))) {
                var v20 float64 = 0.1
                fmt.Println(v20)
                if len(v13) > 0 {
                    v13[0] = 2147483647
                }
                _ = v20
            }
        }
        v2 -= 42
        _ = fi19
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        switch fi21 {
        case -2147483648:
            v1[-541] = 2147483647
        case 32768:
            v22 := GenS2{F0: ((("" + "a") + "foo bar") + (("hello" + "\n\t") + ("hello" + "a"))), F1: 2147483647}
            _ = v22
        case 4294967295:
            switch "hello" {
            case "":
                if (1 > 18446744073709551615) {
                    fmt.Println(v2, fi21)
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 bool = (42 == 18446744073709551615)
                        if v24 {
                            var v25 float32 = 0.0
                            v26 := GenS2{F0: "a", F1: -478}
                            fmt.Println(fi21)
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS1{F0: 42, F1: 245}
                            _ = v27
                        }
                        var v28 int32 = v2
                        if ((1e10 * (3.14159 - 0.0)) == 1e10) {
                            v29 := GenS1{F0: -128, F1: 183}
                            var v30 float64 = float64(float32(0.0))
                            _ = v29
                            _ = v30
                        } else {
                            v31 := GenS2{F0: ("hello" + "hello"), F1: 1}
                            _ = v31
                        }
                        _ = fi23
                        _ = v24
                        _ = v28
                    }
                    var v32 int64 = 593
                    _ = v32
                }
            case "\n\t":
                var v33 float64 = ((0.0 * 1e10) + 1.5)
                _ = v33
            }
        default:
            if true {
                for fi34 := 0; fi34 < 17; fi34++ {
                    if true {
                        var v35 int32 = v2
                        v36 := GenS2{F0: ("hello" + "\n\t"), F1: v35}
                        if true {
                            v35 -= (v2 | v2)
                        }
                        _ = v35
                        _ = v36
                    }
                    _ = fi34
                }
                if (-32768 >= 32767) {
                    v2 += v2
                    for fi37 := 0; fi37 < 12; fi37++ {
                        switch fi37 {
                        case 2147483647:
                            v13 = append(v13, 42)
                        case -2147483648:
                            if (false && (!(0.1 == -2.25))) {
                                if (false && false) {
                                    var v38 float64 = 1.5
                                    var v39 string = "hello"
                                    _ = v38
                                    _ = v39
                                }
                            }
                        case -797:
                            if true {
                                if false {
                                    v1[(-(fi37))] = (fi21 << ((78) & 63))
                                    v2 = (v2 >> ((4294967295) & 63))
                                }
                                v2 = 2147483647
                                var v40 int64 = -1
                                _ = v40
                            }
                        default:
                            v2 -= v2
                        }
                        switch (fi37 << ((65535) & 63)) {
                        case 397:
                            v41 := GenS1{F0: 0, F1: 2}
                            _ = v41
                        case -379:
                            if (!(1 <= 1)) {
                                if (255 > 255) {
                                    if ("foo bar" == "\n\t") {
                                        v42 := v3[4294967295]
                                        var v43 float32 = ((0.1 * 1e10) / (-2.25 * 3.14159))
                                        v42 = v42
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if true {
                                            if (v2 == v2) {
                                                v2 -= (-(v2))
                                                v44 := GenS1{F0: -128, F1: 128}
                                                _ = v44
                                            }
                                            var v45 int16 = -32768
                                            var v46 int32 = v2
                                            _ = v45
                                            _ = v46
                                        }
                                        v47 := GenS1{F0: int8(v2), F1: 0}
                                        _ = v47
                                    }
                                    v2 = (v2 - -717)
                                }
                                if true {
                                    var v48 uint8 = 42
                                    if ((-269 != 42) || true) {
                                        if (65535 <= 1) {
                                            var v49 uint16 = 65535
                                            _ = v49
                                        } else {
                                            _ = len(v1)
                                            v48 = (v48 * 0)
                                            if len(v13) > 0 {
                                                v13[0] = fi21
                                            }
                                        }
                                        if false {
                                            switch fi21 {
                                            case 128:
                                                if true {
                                                    v50 := GenS3{F0: 4294967295}
                                                    switch (-(fi21)) {
                                                    case 9223372036854775807:
                                                        fmt.Println(v48, v48, fi21)
                                                    case 127:
                                                        v2 = (v2 << ((1) & 63))
                                                    case -9223372036854775808:
                                                        v2 = (v2 << ((1) & 63))
                                                    default:
                                                        var v51 int16 = 42
                                                        _ = v51
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v52 int64 = -1
                                                        _ = v52
                                                    default:
                                                        var v53 int8 = -21
                                                        _ = v53
                                                    }
                                                    _ = v50
                                                } else {
                                                    v48 = (^(v48))
                                                }
                                            case -9223372036854775808:
                                                if true {
                                                    var v54 int = fi37
                                                    _ = v54
                                                }
                                            }
                                        }
                                    }
                                    _ = v48
                                } else {
                                    v55 := GenS4{F0: genFn0()}
                                    v2 = -1
                                    v2 = v2
                                    _ = v55
                                }
                                if true {
                                    if (0.0 == 3.14159) {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v3[42] = ("" + ("hello" + ("" + "hello")))
                                        case "hello":
                                            v2 += (^(v2))
                                        default:
                                            if len(v13) > 0 {
                                                v13[0] = (fi21 | int(v2))
                                            }
                                        }
                                    }
                                    var v56 float32 = 0.1
                                    _ = v56
                                } else {
                                    if true {
                                        v57 := GenS4{F0: 42}
                                        _ = v57
                                    }
                                    v58 := GenS4{F0: 42}
                                    fmt.Println(fi37, fi21)
                                    v2 = (v2 & v2)
                                    _ = v58
                                }
                            }
                        }
                        _ = fi37
                    }
                } else {
                    switch "a" {
                    case "hello":
                        v59 := GenS3{F0: genFn1()}
                        _ = v59
                    case "\n\t":
                        for fi60 := 0; fi60 < 5; fi60++ {
                            v61 := GenS5{F0: 42}
                            _ = fi60
                            _ = v61
                        }
                    case "":
                        v2 += -1
                    default:
                        v62 := GenS1{F0: 127, F1: 2}
                        _ = v62
                    }
                    v2 += (v2 + (-(v2)))
                    fmt.Println(fi21, v2)
                }
            }
        }
        fmt.Println(fi21)
        _ = fi21
    }
    var v63 string = "foo bar"
    v2 = v2
    var v64 int = -1
    v65 := GenS6{F0: 853, F1: v63}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
}
