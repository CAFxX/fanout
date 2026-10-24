package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 int; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 string; F1 uint64; F2 uint }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0() uint32 {
    var v5 uint8 = 128
    v6 := []int{}
    var v7 uint32 = 2147483647
    if true {
        if ((uint64(v5) < 42) && false) {
            var v8 int64 = 0
            _ = v8
        } else {
            v7 = v7
            var v9 uint32 = (^(v7))
            var v10 int32 = 42
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                v7 = 630
                _ = fi12
            }
            _ = fi11
        }
        v13 := GenS1{F0: 673}
        _ = v13
    } else {
        switch -888 {
        case 377:
            v5 -= (^(v5))
        }
        switch "foo bar" {
        case "\n\t":
            for fi14 := 0; fi14 < 16; fi14++ {
                if len(v6) > 0 {
                    v6[0] = (fi14 - 32767)
                }
                v15 := GenS1{F0: 42}
                _ = fi14
                _ = v15
            }
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return (v7 - v7)
}

func main() {
    v1 := []uint32{2147483647}
    var v2 int64 = 9223372036854775807
    var v3 float64 = -2.25
    v4 := make(map[uint32]string)
    var v16 bool = false
    switch "a" {
    case "hello":
        var v17 int32 = -998
        _ = v17
    }
    var v18 int8 = -128
    if true {
        _ = len(v1)
        var v19 string = "foo bar"
        v1 = append(v1, genFn0())
        _ = v19
    } else {
        v20 := GenS2{F0: -2147483648, F1: 128, F2: -564}
        for fi21 := 0; fi21 < 2; fi21++ {
            v18 = v18
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 float32 = (0.0 * (1e10 - (1.5 + 1.5)))
                if (65535 != 32767) {
                    switch (fi21 % (((fi22 + (fi22 + fi22))) | 1)) {
                    case -9223372036854775808:
                        var v24 int = (fi21 + 1)
                        _ = v24
                    case -1:
                        v25 := GenS2{F0: fi22, F1: 255, F2: 4294967295}
                        _ = v25
                    case 32767:
                        if (v16 && true) {
                            v16 = (42 != 29)
                            var v26 string = "foo bar"
                            v27 := v4[2]
                            _ = v26
                            _ = v27
                        }
                    default:
                        v16 = (255 >= 112)
                    }
                    v3 = v3
                    var v28 int16 = -453
                    _ = v28
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v4[2] = "hello"
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 float32 = -2.25
                _ = fi30
                _ = v31
            }
            switch "" {
            case "a":
                for fi32 := 0; fi32 < 20; fi32++ {
                    switch -297 {
                    case -1:
                        var v33 float32 = ((0.1 * float32((0.1 - v3))) / 1.5)
                        _ = v33
                    case 4294967295:
                        fmt.Println(v16)
                    default:
                        var v34 int16 = 235
                        _ = v34
                    }
                    switch ((("hello" + "foo bar") + (("foo bar" + "a") + "\n\t")) + "a") {
                    case "a":
                        var v35 uint8 = 0
                        _ = v35
                    case "\n\t":
                        v36 := GenS3{F0: (v2 / (((v2 ^ -888)) | 1)), F1: (fi29 % ((4294967295) | 1)), F2: 2}
                        _ = v36
                    }
                    _ = fi32
                }
            case "":
                var v37 int8 = v18
                _ = v37
            default:
                var v38 float64 = 1e10
                _ = v38
            }
            _ = fi29
        }
        _ = v20
    }
    var v39 int32 = 5
    for fi40 := 0; fi40 < 7; fi40++ {
        v41 := GenS4{F0: (-(v18)), F1: (true || (v16 && false))}
        for fi42 := 0; fi42 < 8; fi42++ {
            v3 += (float64(1e10) - v3)
            var v43 uint16 = 255
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    if (2 == 255) {
        v44 := v4[2]
        _ = v44
    } else {
        v45 := GenS5{F0: "a", F1: 42, F2: 2}
        v18 -= 1
        _ = v45
    }
    if v16 {
        if ("hello" != ((((("hello" + ((((("" + "a") + "foo bar") + "a") + "\n\t") + "\n\t")) + ("a" + "")) + ("\n\t" + "hello")) + "") + "")) {
            if len(v1) > 0 {
                v1[0] = genFn0()
            }
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 int32 = (^(v39))
                switch 590 {
                case -1:
                    if (2 >= 2) {
                        v1 = append(v1, 919)
                        if (v18 > (^(v18))) {
                            if v16 {
                                for fi48 := 0; fi48 < 8; fi48++ {
                                    switch fi48 {
                                    case -518:
                                        fmt.Println(v3, fi48)
                                    case 808:
                                        fmt.Println(v2, v16, v16)
                                    }
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        v39 = (v39 >> ((2) & 63))
                                    case "\n\t":
                                        v2 += v2
                                    case "a":
                                        v16 = v16
                                    default:
                                        var v49 bool = v16
                                        _ = v49
                                    }
                                    delete(v4, genFn0())
                                    _ = fi48
                                }
                                v16 = (!false)
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 2147483648
                                }
                                v50 := GenS6{F0: (^(v18)), F1: 18446744073709551615}
                                _ = v50
                            }
                            v47 = 2147483647
                            switch ("" + "") {
                            case "\n\t":
                                if v16 {
                                    if (42 == 0) {
                                        var v51 uint16 = 65535
                                        _ = v51
                                    }
                                    if (v16 && v16) {
                                        var v52 int32 = 0
                                        v53 := GenS3{F0: (^(v2)), F1: fi46, F2: 255}
                                        var v54 uint16 = 448
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                } else {
                                    v1 = append(v1, 42)
                                }
                            case "":
                                var v55 bool = v16
                                _ = v55
                            case "a":
                                v16 = (((true || true) || (!((float64(1e10) >= v3) || true))) && v16)
                            }
                        }
                    }
                }
                _ = fi46
                _ = v47
            }
        }
        var v56 int64 = v2
        switch 9223372036854775807 {
        case 515:
            v3 += ((1.5 / 1.0) * (float64(0.0) - 3.14159))
        default:
            _ = len(v4)
        }
        _ = v56
    } else {
        if true {
            if (!v16) {
                if (32767 == 42) {
                    if true {
                        v16 = (65535 > 65535)
                        v39 += 42
                        fmt.Println(v3, v3, v18)
                    }
                    var v57 uint32 = 732
                    _ = v57
                }
                var v58 int = -2147483648
                if (v2 <= (v2 & (v2 - (^(v2))))) {
                    switch v58 {
                    case 2147483647:
                        v18 += v18
                    case 127:
                        fmt.Println(v16, v3)
                    case -653:
                        fmt.Println(v2)
                    default:
                        if v16 {
                            v16 = ((true || true) && v16)
                        }
                    }
                } else {
                    fmt.Println(v58)
                    v16 = (v39 == (-(v39)))
                    var v59 float64 = float64(0.0)
                    _ = v59
                }
                _ = v58
            } else {
                v3 -= float64((0.0 + (((float32(v3) * (0.1 / -2.25)) * 0.0) - -2.25)))
            }
            for fi60 := 0; fi60 < 9; fi60++ {
                if (((v3 * v3) / (v3 - (v3 * (float64((1e10 + (-2.25 - 1e10))) * (-2.25 + 0.0))))) != 1e10) {
                    var v61 int64 = v2
                    v62 := GenS7{F0: true, F1: 0, F2: 255}
                    v39 = (v39 % ((v39) | 1))
                    _ = v61
                    _ = v62
                }
                _ = fi60
            }
            fmt.Println(v39)
        } else {
            var v63 uint64 = 844
            var v64 uint = 65535
            v65 := GenS8{F0: 2, F1: (^(v64)), F2: 876}
            _ = v63
            _ = v64
            _ = v65
        }
    }
    var v66 int16 = -510
    v67 := v4[446]
    var v68 int = 32767
    fmt.Println(v66, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v66
    _ = v67
    _ = v68
}
