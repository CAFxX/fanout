package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 int64; F1 int8; F2 int8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int16; F1 int8 }
func genFn0() uint32 {
    switch -1 {
    case 416:
        var v6 uint64 = 4294967295
        _ = v6
    case 559:
        if (((959 >= 18446744073709551615) && ((65535 == 382) && true)) || (!(127 >= 1))) {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 string = ("foo bar" + ("foo bar" + (((("foo bar" + "\n\t") + "") + "") + "foo bar")))
                _ = fi7
                _ = v8
            }
            var v9 uint = 4294967295
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 string = (("a" + "a") + ("" + ""))
            _ = fi10
            _ = v11
        }
    }
    var v12 int64 = 42
    v12 -= (^(v12))
    v13 := []int{615, 848, -50}
    for fi14 := 0; fi14 < 6; fi14++ {
        _ = len(v13)
        fmt.Println(v12, v12)
        _ = fi14
    }
    _ = v12
    _ = v13
    return 126
}

func main() {
    var v1 int8 = 93
    v2 := []string{""}
    v3 := []int{}
    v4 := []string{(("" + (("a" + "\n\t") + "hello")) + ("a" + "")), "\n\t"}
    v5 := make(map[string]string)
    var v15 int16 = 32767
    if false {
        var v16 uint = 0
        var v17 uint64 = 948
        _ = v16
        _ = v17
    } else {
        v1 = 100
        if (632 < 2147483647) {
            v18 := GenS1{F0: 883}
            _ = v18
        }
        fmt.Println(v1)
    }
    var v19 bool = (1.5 < 0.1)
    v20 := GenS1{F0: -360}
    v21 := GenS1{F0: 2147483647}
    if v19 {
        var v22 uint8 = 98
        for fi23 := 0; fi23 < 5; fi23++ {
            switch ("foo bar" + "\n\t") {
            case "a":
                var v24 string = ("hello" + "foo bar")
                _ = v24
            case "hello":
                if (true || (!true)) {
                    fmt.Println(v22, fi23, v19)
                    if (!((fi23 != -117) || false)) {
                        v25 := GenS1{F0: 1}
                        _ = v25
                    } else {
                        v22 += (v22 % (((^(v22))) | 1))
                        var v26 uint = 2
                        _ = v26
                    }
                    v27 := GenS1{F0: 2147483647}
                    _ = v27
                } else {
                    var v28 uint32 = 2147483648
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 int = (fi29 << ((838) & 63))
                        v22 += v22
                        if (0 != 42) {
                            var v31 uint = 0
                            switch "a" {
                            case "a":
                                if true {
                                    delete(v5, "a")
                                }
                            default:
                                v19 = true
                            }
                            if true {
                                v31 += 42
                                if v19 {
                                    fmt.Println(v19, v19, v1)
                                } else {
                                    if (uint64(v31) != 626) {
                                        v32 := GenS1{F0: 42}
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: 2147483647}
                                        var v34 float32 = ((0.0 * (float32(float64(float32(((1e10 + (1e10 + 1.5)) * 0.1)))) + ((((-2.25 + -2.25) + 0.0) * 1.5) / (float32(1.5) / (1e10 + 1e10))))) + (float32(0.1) * 1e10))
                                        v5[("foo bar" + (("foo bar" + "\n\t") + ""))] = ("a" + "\n\t")
                                        v19 = (65535 >= v28)
                                        _ = v33
                                        _ = v34
                                    }
                                    if (-2.25 >= 0.1) {
                                        _ = len(v4)
                                        if ((v22 / ((v22) | 1)) > (v22 << ((593) & 63))) {
                                            if true {
                                                _ = len(v5)
                                                var v35 int8 = v1
                                                v28 += (v28 - (v28 | v28))
                                                _ = v35
                                            } else {
                                                v36 := GenS2{F0: (v1 % ((v1) | 1)), F1: v22}
                                                _ = len(v4)
                                                _ = v36
                                            }
                                            v28 = genFn0()
                                        }
                                        v22 = 0
                                    }
                                }
                                v37 := GenS3{F0: (v30 << ((v31) & 63)), F1: (v28 * (v28 * v28)), F2: 154}
                                var v38 float32 = 0.1
                                _ = v37
                                _ = v38
                            }
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            case "":
                v15 = (v15 & v15)
            default:
                var v39 uint = 255
                _ = v39
            }
            for fi40 := 0; fi40 < 6; fi40++ {
                delete(v5, ("foo bar" + "\n\t"))
                v41 := GenS4{F0: -1, F1: int8(v15), F2: v1}
                _ = fi40
                _ = v41
            }
            v42 := GenS2{F0: -128, F1: (v22 & v22)}
            _ = fi23
            _ = v42
        }
        if (2 != 4294967295) {
            for fi43 := 0; fi43 < 17; fi43++ {
                if (318 != 346) {
                    v3 = append(v3, fi43)
                    for fi44 := 0; fi44 < 6; fi44++ {
                        var v45 int16 = v15
                        var v46 bool = v19
                        v47 := GenS5{F0: "a", F1: 0}
                        if ("" > ("" + "a")) {
                            v1 -= 1
                            _ = len(v5)
                            v48 := GenS2{F0: (-(v1)), F1: v22}
                            _ = v48
                        }
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                }
                for fi49 := 0; fi49 < 8; fi49++ {
                    var v50 uint32 = 580
                    v22 -= v22
                    _ = fi49
                    _ = v50
                }
                fmt.Println(fi43, fi43)
                _ = fi43
            }
        } else {
            switch -1 {
            case -1:
                var v51 int32 = 0
                _ = v51
            case 373:
                if len(v3) > 0 {
                    v3[0] = -9223372036854775808
                }
            case 1:
                v52 := GenS6{F0: 150, F1: 337}
                _ = v52
            default:
                fmt.Println(v19, v19, v22)
            }
            for fi53 := 0; fi53 < 19; fi53++ {
                fmt.Println(v22, v19)
                v54 := GenS7{F0: 0, F1: 4294967295}
                v22 += (v22 | 42)
                _ = fi53
                _ = v54
            }
            v2 = append(v2, "\n\t")
        }
        _ = v22
    }
    v3 = append(v3, 127)
    for fi55 := 0; fi55 < 14; fi55++ {
        v15 -= (v15 + 0)
        _ = fi55
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        v57 := GenS8{F0: 42, F1: 0}
        v19 = (672 != -342)
        _ = fi56
        _ = v57
    }
    v1 += v1
    v58 := GenS3{F0: 4294967295, F1: 2, F2: -1}
    _ = len(v3)
    switch -675 {
    case -2147483648:
        fmt.Println(v15, v15, v19)
    case -425:
        for fi59 := 0; fi59 < 13; fi59++ {
            v1 += v1
            _ = fi59
        }
    default:
        v60 := GenS9{F0: 1, F1: (v1 ^ 98)}
        _ = v60
    }
    v5[""] = (("a" + "foo bar") + "")
    fmt.Println(v19, v15, v19)
    v19 = ((-2147483648 != 0) && (false || v19))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v58
}
