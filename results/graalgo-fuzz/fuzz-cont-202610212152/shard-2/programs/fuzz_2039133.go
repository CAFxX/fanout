package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 string; F1 int8; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 bool; F2 int64 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 uint8; F1 int64; F2 int32 }
func genFn0(p0 int32, p1 int32, p2 uint8) uint16 {
    var v6 bool = false
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        if (42 <= -1) {
            for fi9 := 0; fi9 < 10; fi9++ {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    if v6 {
                        fmt.Println(fi9)
                        p2 = 17
                        if (0 == 458) {
                            v10 := GenS1{F0: 331, F1: 497}
                            var v11 int16 = -32768
                            var v12 int64 = -1
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int8 = -1
                            p2 = 42
                            fmt.Println(v13, fi8, p0)
                            _ = v13
                        }
                        p1 = 42
                    }
                case "a":
                    v14 := GenS2{F0: 0, F1: "hello"}
                    _ = v14
                case "\n\t":
                    v15 := GenS1{F0: 2, F1: 42}
                    _ = v15
                default:
                    var v16 int64 = 1
                    _ = v16
                }
                p0 = (p0 + p0)
                _ = fi9
            }
            if len(v7) > 0 {
                v7[0] = fi8
            }
        } else {
            if len(v7) > 0 {
                v7[0] = fi8
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint = 699
                v18 = (v18 & (v18 << ((255) & 63)))
                _ = fi17
                _ = v18
            }
            v19 := GenS2{F0: 42, F1: (("foo bar" + "a") + ("\n\t" + ("hello" + "\n\t")))}
            _ = v19
        }
        switch fi8 {
        case 128:
            var v20 uint32 = 2147483648
            _ = v20
        case 195:
            if ((3.14159 * 1.5) != ((-2.25 * 3.14159) + 1e10)) {
                p2 = (p2 + p2)
            } else {
                var v21 uint32 = 42
                _ = v21
            }
        default:
            v22 := GenS1{F0: 255, F1: 0}
            _ = v22
        }
        fmt.Println(p1, p2)
        _ = fi8
    }
    if (0.0 == (-2.25 - (float64(float32((0.0 - 0.0))) + 0.0))) {
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                if len(v7) > 0 {
                    v7[0] = fi24
                }
                var v25 uint64 = 2
                v26 := GenS3{F0: "hello", F1: 1, F2: (-(fi24))}
                p0 = p0
                _ = fi24
                _ = v25
                _ = v26
            }
            p0 -= (p0 >> ((0) & 63))
            _ = fi23
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{42, 4294967295, 0}
    v2 := []uint32{}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 uint8 = 16
    v5 -= (v5 << ((1) & 63))
    v5 -= (v5 | (v5 + (v5 | 224)))
    v1 = append(v1, 2147483647)
    v27 := GenS1{F0: 0, F1: 0}
    v28 := GenS3{F0: ("foo bar" + "foo bar"), F1: 100, F2: 32768}
    delete(v3, "foo bar")
    _ = len(v4)
    if true {
        v5 -= 128
        var v29 uint16 = 2
        _ = v29
    } else {
        v30 := GenS4{F0: 42}
        for fi31 := 0; fi31 < 12; fi31++ {
            if true {
                var v32 uint8 = (v5 >> ((255) & 63))
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 float64 = 1e10
                    _ = fi33
                    _ = v34
                }
                v32 -= v5
                _ = v32
            }
            var v35 int32 = 42
            _ = fi31
            _ = v35
        }
        switch "foo bar" {
        case "a":
            switch "foo bar" {
            case "hello":
                v36 := GenS3{F0: "foo bar", F1: -1, F2: -58}
                _ = v36
            case "a":
                v37 := GenS4{F0: 4294967295}
                _ = v37
            case "":
                if true {
                    v5 = (v5 & 1)
                }
            }
        }
        _ = v30
    }
    v38 := GenS1{F0: 572, F1: 0}
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 7; fi40++ {
            v3[("hello" + "a")] = (v5 > (v5 * v5))
            _ = fi40
        }
        if false {
            var v41 float64 = float64(((0.0 * (0.1 * 1e10)) * 1e10))
            _ = v41
        } else {
            v42 := GenS4{F0: 2147483647}
            v5 = 193
            v43 := GenS5{F0: 65535}
            _ = v42
            _ = v43
        }
        v5 -= (^(v5))
        _ = fi39
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 uint = 338
        _ = fi44
        _ = v45
    }
    v5 = v5
    switch -9223372036854775808 {
    case 32767:
        v46 := v4[0]
        _ = v46
    default:
        if (!false) {
            switch -401 {
            case 9223372036854775807:
                var v47 float32 = float32((3.14159 - -2.25))
                _ = v47
            case -2147483648:
                var v48 float64 = (3.14159 / 3.14159)
                _ = v48
            }
            var v49 float64 = 3.14159
            if (true && true) {
                v49 += -2.25
            } else {
                v50 := v4[2]
                for fi51 := 0; fi51 < 11; fi51++ {
                    switch 127 {
                    case 32767:
                        delete(v4, 2147483648)
                    }
                    v50 -= (fi51 - fi51)
                    if (!(true && true)) {
                        var v52 float32 = 1.5
                        var v53 int = (fi51 / ((128) | 1))
                        if false {
                            v49 = 1e10
                        } else {
                            var v54 float32 = v52
                            _ = v54
                        }
                        _ = v52
                        _ = v53
                    }
                    _ = fi51
                }
                v50 -= 159
                _ = v50
            }
            _ = v49
        }
    }
    fmt.Println(v5)
    if false {
        if ((v5 % ((v5) | 1)) <= v5) {
            if (-2147483648 >= 2147483647) {
                fmt.Println(v5)
                if ((100 < 1) || false) {
                    fmt.Println(v5)
                    v55 := GenS6{F0: 2147483648, F1: true, F2: -717}
                    _ = v55
                } else {
                    v5 = (^(v5))
                    for fi56 := 0; fi56 < 4; fi56++ {
                        v5 = 183
                        var v57 uint16 = 1
                        _ = fi56
                        _ = v57
                    }
                }
            } else {
                for fi58 := 0; fi58 < 8; fi58++ {
                    v5 -= 42
                    switch "foo bar" {
                    case "a":
                        var v59 bool = ((0.1 + -2.25) > 0.1)
                        _ = v59
                    case "\n\t":
                        var v60 int16 = 962
                        _ = v60
                    case "hello":
                        if (true && true) {
                            if len(v2) > 0 {
                                v2[0] = uint32(v5)
                            }
                        } else {
                            var v61 string = ("hello" + "")
                            _ = v61
                        }
                    }
                    for fi62 := 0; fi62 < 6; fi62++ {
                        switch (fi62 & (fi62 & fi58)) {
                        case 0:
                            if false {
                                v63 := GenS6{F0: 42, F1: true, F2: -157}
                                _ = v63
                            }
                        case 128:
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                        }
                        var v64 int32 = -97
                        v3[("a" + ("\n\t" + "hello"))] = false
                        _ = fi62
                        _ = v64
                    }
                    _ = fi58
                }
                for fi65 := 0; fi65 < 18; fi65++ {
                    v66 := v3[""]
                    if (18446744073709551615 != 1) {
                        if v66 {
                            switch "" {
                            case "a":
                                for fi67 := 0; fi67 < 5; fi67++ {
                                    fmt.Println(v5)
                                    if true {
                                        v68 := GenS2{F0: -1, F1: ""}
                                        _ = len(v3)
                                        var v69 uint = 4294967295
                                        _ = v68
                                        _ = v69
                                    }
                                    fmt.Println(fi65, fi65)
                                    _ = fi67
                                }
                            }
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = 429
                    }
                    _ = fi65
                    _ = v66
                }
                fmt.Println(v5)
            }
        } else {
            v70 := GenS7{F0: 73, F1: ""}
            v5 -= 0
            _ = v70
        }
        v71 := GenS8{F0: v5, F1: -1, F2: 2147483647}
        _ = v71
    }
    delete(v4, 2)
    v72 := GenS7{F0: -5, F1: ("" + "\n\t")}
    v73 := v3["a"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v72
    _ = v73
}
