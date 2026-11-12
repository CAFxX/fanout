package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 uint16 }
func genFn0(p0 string, p1 uint64) int {
    fmt.Println(p0, p1)
    var v7 int64 = -416
    _ = p0
    _ = p1
    _ = v7
    return 303
}
func genFn1(p0 uint) uint {
    v8 := []int{-931, 1}
    var v9 uint64 = 18446744073709551615
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int16 = 0
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v9 = v9
            _ = fi10
            _ = v11
        }
        var v12 float64 = 1.5
        _ = v12
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println(p0)
    _ = p0
    _ = v8
    _ = v9
    return (p0 >> ((p0) & 63))
}

func main() {
    v1 := []uint32{1, 42}
    var v2 int = -82
    var v3 int16 = 1
    var v4 uint64 = 0
    var v5 uint64 = v4
    v6 := make(map[string]int)
    var v13 uint64 = v4
    var v14 uint64 = (^(v13))
    if false {
        var v15 int8 = -1
        _ = v15
    }
    if ("foo bar" != "\n\t") {
        switch -1 {
        case 42:
            if (127 <= 127) {
                fmt.Println(v3, v2)
                _ = len(v6)
            }
        case 32768:
            var v16 uint8 = 115
            _ = v16
        case 700:
            var v17 int8 = -101
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 164, F1: 1}
        if (!(-564 < 570)) {
            v19 := v6["\n\t"]
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                if (0.1 >= 1.5) {
                    switch (^(fi20)) {
                    case 0:
                        var v21 float64 = ((1e10 * (0.1 * float64(float32(float64(-2.25))))) * float64(((1.5 - (((3.14159 / ((float32(float64(float32(-2.25))) / 3.14159) * 1e10)) + ((((float32(0.0) + 1e10) - -2.25) + (((-2.25 - (1e10 - 0.0)) / 1.0) - 0.0)) - 0.1)) + float32(float64(float32((float64((3.14159 * (-2.25 - 0.0))) + (float64(0.0) - 3.14159))))))) * 1e10)))
                        _ = v21
                    case -9223372036854775808:
                        switch "" {
                        case "hello":
                            for fi22 := 0; fi22 < 16; fi22++ {
                                var v23 int16 = v3
                                _ = fi22
                                _ = v23
                            }
                        }
                    case 32767:
                        v3 = (^(v3))
                    }
                    if (2 != 65535) {
                        v24 := GenS2{F0: uint32(v13), F1: 794, F2: 65535}
                        v25 := GenS3{F0: 1, F1: (v5 * (v13 >> ((20) & 63)))}
                        if ((!(("hello" + "foo bar") != "")) || (53 <= 1)) {
                            v26 := GenS4{F0: ("a" + ("a" + (((("hello" + (("" + ("a" + "")) + "foo bar")) + "a") + "\n\t") + ("foo bar" + "a")))), F1: 42}
                            v27 := GenS2{F0: 2, F1: 1, F2: 0}
                            var v28 int32 = 0
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v24
                        _ = v25
                    } else {
                        _ = len(v1)
                    }
                } else {
                    v13 = (v13 | (v13 ^ (v14 * (^(v13)))))
                }
                _ = fi20
            }
        }
        _ = v18
    }
    v29 := GenS5{F0: 1}
    v3 = (v3 >> ((genFn1(458)) & 63))
    if ((-(v2)) == (v2 >> ((genFn1(1)) & 63))) {
        switch ("\n\t" + "\n\t") {
        case "":
            v30 := v6[("a" + ("foo bar" + ("hello" + ("\n\t" + ("a" + "\n\t")))))]
            _ = v30
        case "a":
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 uint64 = v13
                var v33 float64 = -2.25
                _ = fi31
                _ = v32
                _ = v33
            }
        default:
            if false {
                fmt.Println(v2)
                v14 -= 479
            } else {
                v34 := GenS6{F0: 639}
                _ = v34
            }
        }
        fmt.Println(v13, v13)
    } else {
        var v35 int16 = 576
        var v36 float64 = 3.14159
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = fi37
    }
    v2 += (v2 - genFn0(("hello" + ""), (^(v4))))
    switch -739 {
    case -695:
        for fi38 := 0; fi38 < 1; fi38++ {
            if false {
                var v39 bool = (v3 < v3)
                if v39 {
                    var v40 uint16 = uint16(v4)
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 8; fi41++ {
                        if v39 {
                            if true {
                                fmt.Println(fi38, fi41)
                                var v42 int8 = 42
                                _ = v42
                            } else {
                                if ((v14 % (((v13 >> ((genFn1(736)) & 63))) | 1)) != (v13 & v4)) {
                                    var v43 int64 = 529
                                    _ = v43
                                } else {
                                    if v39 {
                                        v39 = false
                                        switch "\n\t" {
                                        case "a":
                                            if true {
                                                var v44 int64 = 0
                                                fmt.Println(v14, fi38, v3)
                                                v3 -= (v3 - v3)
                                                _ = v44
                                            } else {
                                                var v45 uint16 = 42
                                                _ = v45
                                            }
                                            if true {
                                                var v46 bool = (65535 >= 32767)
                                                if true {
                                                    if v46 {
                                                        v2 += v2
                                                    } else {
                                                        switch fi38 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v3, v5, fi41)
                                                        case -897:
                                                            switch -1 {
                                                            case 32767:
                                                                v47 := GenS1{F0: -1, F1: 1}
                                                                _ = v47
                                                            case 42:
                                                                var v48 uint8 = 0
                                                                v3 += 0
                                                                _ = v48
                                                            }
                                                        }
                                                        switch "hello" {
                                                        case "a":
                                                            if (!false) {
                                                                v3 -= v3
                                                                v49 := GenS1{F0: -1, F1: 255}
                                                                _ = v49
                                                            } else {
                                                                v13 = (v14 + v4)
                                                            }
                                                        case "hello":
                                                            if ((fi41 % ((fi41) | 1)) <= (fi38 % ((fi41) | 1))) {
                                                                v13 = (v4 | (v13 | v5))
                                                                _ = len(v6)
                                                            }
                                                        default:
                                                            var v50 int32 = -480
                                                            _ = v50
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        v39 = (!false)
                                                        v2 += (-(v2))
                                                        v6[((("" + "\n\t") + "\n\t") + "")] = 127
                                                    case "\n\t":
                                                        var v51 uint64 = 42
                                                        _ = v51
                                                    case "":
                                                        v52 := GenS2{F0: 2147483648, F1: 65535, F2: 32767}
                                                        _ = v52
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                }
                                                var v53 uint8 = 190
                                                _ = v46
                                                _ = v53
                                            } else {
                                                v4 += 1
                                                v39 = v39
                                            }
                                            v54 := GenS5{F0: 2}
                                            fmt.Println(v3)
                                            _ = v54
                                        case "foo bar":
                                            fmt.Println(v3)
                                        case "":
                                            if ((!v39) && (!(!(1 == 100)))) {
                                                var v55 uint = 196
                                                v4 += v13
                                                v39 = (uint(v13) > (v55 / ((0) | 1)))
                                                _ = v55
                                            } else {
                                                var v56 int64 = -65
                                                v57 := GenS1{F0: -579, F1: 140}
                                                var v58 uint8 = 127
                                                if v39 {
                                                    v39 = ((true && v39) || v39)
                                                }
                                                if ((fi41 % (((fi41 & (fi41 + 327))) | 1)) < 922) {
                                                    fmt.Println(fi38)
                                                    v56 = v56
                                                }
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        default:
                                            if v39 {
                                                v59 := GenS7{F0: 4294967295}
                                                _ = v59
                                            }
                                            var v60 int8 = 52
                                            var v61 bool = (-2147483648 != int32(v60))
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    v39 = v39
                                    fmt.Println(v5)
                                }
                                _ = len(v1)
                            }
                        }
                        _ = fi41
                    }
                    for fi62 := 0; fi62 < 3; fi62++ {
                        v13 = (v5 >> ((42) & 63))
                        var v63 int64 = 71
                        _ = fi62
                        _ = v63
                    }
                }
                v64 := GenS8{F0: fi38, F1: 2}
                _ = v39
                _ = v64
            }
            _ = fi38
        }
    case 180:
        _ = len(v1)
    case 4294967295:
        switch 4294967295 {
        case 9223372036854775807:
            v5 = v13
        case 127:
            var v65 int16 = v3
            _ = v65
        }
    }
    v66 := GenS1{F0: 392, F1: uint16(v14)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
