package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint64; F2 int16 }
type GenS7 struct { F0 int32; F1 uint16; F2 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int16; F2 string }
type GenS10 struct { F0 uint8; F1 uint; F2 bool }

func main() {
    var v1 bool = (!true)
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 int16 = 32767
    if false {
        _ = len(v3)
    } else {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 bool = v1
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: 100}
    v1 = ((false || (-128 != 48)) || v1)
    v2 = 127
    if false {
        delete(v3, 239)
        if (!((0.1 / 1e10) != ((((3.14159 * (1.5 + 1.5)) + 1e10) * (((-2.25 - 1e10) - 0.0) - (1e10 - (0.1 / 1.0)))) * 1.5))) {
            v1 = (((-169 < -1) || v1) || ((42 >= 0) || (-1 < 27)))
        } else {
            v9 := GenS2{F0: -388, F1: (v5 + -1)}
            var v10 int = -1
            v11 := GenS3{F0: 30, F1: 128}
            _ = v10
            _ = v11
            _ = v9
        }
        v2 = 128
    } else {
        if v1 {
            v12 := GenS3{F0: 1, F1: 194}
            _ = v12
        }
    }
    v13 := GenS4{F0: v5, F1: 249}
    if ((!true) && (1 >= 1)) {
        if v1 {
            var v14 int16 = v5
            var v15 uint = 4294967295
            if (!(!v1)) {
                var v16 uint8 = 195
                var v17 int64 = 0
                var v18 float32 = ((-2.25 + float32((0.0 / float64(0.1)))) * (0.1 * 1e10))
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v14
            _ = v15
        } else {
            var v19 uint16 = 864
            fmt.Println(v2)
            v20 := []int{(v2 + v2), v2, v2, (v2 << ((569) & 63))}
            _ = v19
            _ = v20
        }
        if v1 {
            var v21 int32 = -782
            if ("a" == (("\n\t" + ("hello" + "\n\t")) + ("" + "\n\t"))) {
                if false {
                    v1 = (76 <= -128)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(v1, fi22)
                        _ = fi22
                    }
                } else {
                    v23 := GenS5{F0: 1}
                    _ = v23
                }
            } else {
                if (0 == 0) {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        if ("\n\t" != "") {
                            for fi25 := 0; fi25 < 8; fi25++ {
                                v26 := GenS6{F0: 111, F1: 103, F2: (-(v5))}
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            v27 := GenS7{F0: -1, F1: 42, F2: ((!(v2 < (fi24 >> ((4294967295) & 63)))) || v1)}
                            _ = v27
                        }
                        _ = len(v4)
                        if false {
                            switch "" {
                            case "\n\t":
                                switch ((((("foo bar" + "foo bar") + "hello") + ("hello" + (("foo bar" + ("foo bar" + ("" + ("hello" + "hello")))) + ("foo bar" + "a")))) + "foo bar") + "") {
                                case "":
                                    v5 -= v5
                                case "hello":
                                    if v1 {
                                        var v28 float32 = (1.5 - (float32(3.14159) + (0.1 * 1.5)))
                                        _ = v28
                                    }
                                case "a":
                                    v5 += (-(v5))
                                default:
                                    for fi29 := 0; fi29 < 16; fi29++ {
                                        v30 := GenS8{F0: 4294967295}
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v5, fi24, fi24)
                                        default:
                                            fmt.Println(v1, v2)
                                        }
                                        v31 := v4[(fi29 & 863)]
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            }
                        } else {
                            if false {
                                var v32 float32 = 1e10
                                _ = v32
                            } else {
                                v1 = true
                                v21 = (v21 ^ (^(v21)))
                                v21 = (v21 + -437)
                            }
                            if (v1 || ("\n\t" <= "\n\t")) {
                                if (494 != 2) {
                                    for fi33 := 0; fi33 < 11; fi33++ {
                                        v1 = v1
                                        _ = fi33
                                    }
                                } else {
                                    v21 += (v21 & (v21 & (v21 - (-(v21)))))
                                    for fi34 := 0; fi34 < 1; fi34++ {
                                        var v35 int64 = 728
                                        _ = fi34
                                        _ = v35
                                    }
                                    delete(v3, 42)
                                }
                                if v1 {
                                    for fi36 := 0; fi36 < 5; fi36++ {
                                        var v37 int = 1
                                        var v38 int8 = -1
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                    v39 := GenS9{F0: v5, F1: v5, F2: "\n\t"}
                                    _ = v39
                                } else {
                                    v2 = (v2 / ((732) | 1))
                                    for fi40 := 0; fi40 < 4; fi40++ {
                                        if v1 {
                                            switch fi24 {
                                            case 127:
                                                v21 += (v21 + v21)
                                            default:
                                                v5 -= v5
                                            }
                                        } else {
                                            v41 := GenS10{F0: 1, F1: 4294967295, F2: v1}
                                            fmt.Println(fi24, v2, fi24)
                                            var v42 string = "\n\t"
                                            _ = v41
                                            _ = v42
                                        }
                                        if (532 == 337) {
                                            var v43 int32 = v21
                                            fmt.Println(v1)
                                            v44 := []int{32767, v2, 42}
                                            _ = v43
                                            _ = v44
                                        }
                                        v45 := GenS4{F0: 0, F1: 255}
                                        switch "\n\t" {
                                        case "a":
                                            v46 := GenS7{F0: v21, F1: 2, F2: (!(18446744073709551615 == 4294967295))}
                                            _ = v46
                                        case "":
                                            var v47 int64 = -347
                                            _ = v47
                                        }
                                        if (18446744073709551615 >= 229) {
                                            v2 = fi40
                                            v48 := GenS8{F0: 255}
                                            _ = v48
                                        }
                                        if false {
                                            v49 := []int{}
                                            v1 = (v1 || v1)
                                            _ = v49
                                        } else {
                                            fmt.Println(fi24, v1)
                                            v21 -= -1
                                        }
                                        var v50 float64 = 0.1
                                        v21 = v21
                                        _ = fi40
                                        _ = v45
                                        _ = v50
                                    }
                                }
                                _ = len(v4)
                            }
                        }
                        _ = fi24
                    }
                    var v51 uint16 = 255
                    _ = v51
                } else {
                    v52 := []int{v2, (-(v2))}
                    _ = v52
                }
                var v53 int8 = 29
                switch v2 {
                case 1:
                    delete(v3, 449)
                default:
                    v54 := GenS8{F0: 65535}
                    _ = v54
                }
                _ = v53
            }
            _ = v21
        }
    }
    v1 = v1
    var v55 int16 = (v5 + 945)
    v56 := []int{}
    var v57 float32 = 1.5
    for fi58 := 0; fi58 < 16; fi58++ {
        v59 := GenS1{F0: 5}
        switch (-(fi58)) {
        case -528:
            v55 = (v55 & (v55 ^ (v5 / ((v5) | 1))))
        case -9223372036854775808:
            v2 += (v2 * fi58)
        case 32768:
            if (false && v1) {
                v57 = v57
                v55 += (v55 - (v5 % ((v55) | 1)))
                var v60 int32 = -893
                _ = v60
            }
        default:
            v61 := v4[(fi58 >> ((467) & 63))]
            _ = v61
        }
        _ = fi58
        _ = v59
    }
    for fi62 := 0; fi62 < 14; fi62++ {
        v63 := GenS4{F0: v5, F1: 4294967295}
        v57 += v57
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 1; fi64++ {
        v1 = true
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v8
}
