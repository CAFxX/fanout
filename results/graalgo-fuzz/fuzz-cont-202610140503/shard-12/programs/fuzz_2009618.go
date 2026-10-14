package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64, p1 bool) int8 {
    var v4 uint = 946
    var v5 uint8 = 0
    v4 += (v4 % ((42) | 1))
    v6 := []int{267, 127}
    v4 = (v4 & (v4 % ((255) | 1)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint16) string {
    v7 := []int{128}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    var v3 uint32 = 4294967295
    var v8 int16 = int16(v3)
    var v9 uint8 = 42
    switch 616 {
    case -448:
        v10 := GenS1{F0: (v3 | 2147483647), F1: -1}
        _ = v10
    default:
        if (-1 != genFn0(42, (!false))) {
            var v11 uint64 = 2
            var v12 int32 = 2147483647
            v3 -= 418
            _ = v11
            _ = v12
        } else {
            v13 := v2[108]
            _ = v13
        }
    }
    v1 = v1
    fmt.Println(v9, v8, v9)
    switch ("foo bar" + "foo bar") {
    case "a":
        fmt.Println(v1, v1, v1)
    case "hello":
        for fi14 := 0; fi14 < 16; fi14++ {
            switch (^(fi14)) {
            case 42:
                if true {
                    var v15 int64 = -295
                    _ = v15
                } else {
                    v3 = v3
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v8 = (v8 & v8)
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS1{F0: v3, F1: 0}
                        if true {
                            var v19 int64 = 9223372036854775807
                            _ = v19
                        } else {
                            if false {
                                var v20 int = 4294967295
                                _ = v20
                            }
                            var v21 int8 = (v1 + genFn0(862, false))
                            if (("a" > "a") && false) {
                                v1 -= 0
                                fmt.Println(fi17)
                                v9 = uint8(v3)
                            } else {
                                var v22 float64 = float64(0.0)
                                v23 := GenS1{F0: (v3 | (v3 * (v3 << ((0) & 63)))), F1: 42}
                                v24 := []int{105, (^(fi17))}
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            }
            _ = fi14
        }
    case "":
        v2[634] = (v3 + (^(v3)))
    default:
        if false {
            var v25 float64 = -2.25
            v26 := []int{-1, 32767, 97, -9223372036854775808}
            v27 := v2[793]
            _ = v25
            _ = v26
            _ = v27
        } else {
            v28 := GenS2{F0: (65535 <= 2)}
            if (false && (int64(v8) != 1)) {
                var v29 uint64 = 4294967295
                if false {
                    fmt.Println(v3, v8)
                    for fi30 := 0; fi30 < 20; fi30++ {
                        if ((271 >= 32767) && (!true)) {
                            if false {
                                if ((306 > -1) || true) {
                                    var v31 int32 = 42
                                    var v32 int = (fi30 | 2147483647)
                                    v33 := []int{}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    switch fi30 {
                                    case 4294967295:
                                        v9 -= v9
                                    case 42:
                                        v34 := GenS3{F0: 2147483647, F1: 18446744073709551615}
                                        _ = v34
                                    case 13:
                                        v3 += 4294967295
                                    }
                                    delete(v2, (-(fi30)))
                                }
                                for fi35 := 0; fi35 < 6; fi35++ {
                                    if (42 <= -32) {
                                        if true {
                                            v36 := GenS1{F0: (v3 ^ v3), F1: -2147483648}
                                            switch "" {
                                            case "a":
                                                v29 += (v29 / (((v29 >> ((65535) & 63))) | 1))
                                            case "foo bar":
                                                var v37 uint = 2
                                                _ = v37
                                            case "":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "":
                                                        var v38 float32 = -2.25
                                                        _ = v38
                                                    case "foo bar":
                                                        v39 := GenS3{F0: -450, F1: 233}
                                                        _ = v39
                                                    case "hello":
                                                        v40 := []int{(fi35 & (fi30 | fi35))}
                                                        _ = v40
                                                    }
                                                case "\n\t":
                                                    var v41 uint32 = (v3 >> ((18) & 63))
                                                    var v42 bool = (true || false)
                                                    _ = v41
                                                    _ = v42
                                                case "":
                                                    v43 := GenS3{F0: 390, F1: 535}
                                                    _ = v43
                                                }
                                            default:
                                                fmt.Println(v29)
                                            }
                                            v8 -= v8
                                            _ = v36
                                        } else {
                                            v44 := v2[(fi35 / ((int(v9)) | 1))]
                                            _ = v44
                                        }
                                        v45 := GenS4{F0: (v9 / ((v9) | 1))}
                                        if (!(((0.0 - 1.5) * 0.0) >= 3.14159)) {
                                            fmt.Println(v3)
                                            var v46 string = "foo bar"
                                            var v47 int64 = 9223372036854775807
                                            _ = v46
                                            _ = v47
                                        } else {
                                            fmt.Println(v1, v9, v3)
                                            v48 := v2[-2147483648]
                                            _ = v48
                                        }
                                        _ = v45
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            v49 := GenS5{F0: v29, F1: 1, F2: (v8 & 903)}
                                            _ = v49
                                        case "a":
                                            v3 = 480
                                        }
                                        var v50 float32 = (1e10 * 1e10)
                                        _ = v50
                                    }
                                    _ = fi35
                                }
                            } else {
                                for fi51 := 0; fi51 < 5; fi51++ {
                                    var v52 bool = true
                                    if (100 <= (v1 << ((132) & 63))) {
                                        var v53 int = 32768
                                        _ = v53
                                    }
                                    _ = fi51
                                    _ = v52
                                }
                                _ = len(v2)
                            }
                        }
                        if (!false) {
                            v3 += (v3 / ((2147483647) | 1))
                            v9 = (v9 % ((v9) | 1))
                        }
                        _ = fi30
                    }
                }
                _ = v29
            }
            _ = v28
        }
    }
    for fi54 := 0; fi54 < 15; fi54++ {
        v1 -= v1
        for fi55 := 0; fi55 < 3; fi55++ {
            var v56 string = "\n\t"
            _ = fi55
            _ = v56
        }
        if (2 <= 209) {
            v9 = (v9 >> ((uint(v9)) & 63))
            v1 = int8(fi54)
        }
        _ = fi54
    }
    var v57 uint16 = 205
    v3 -= uint32(v57)
    for fi58 := 0; fi58 < 19; fi58++ {
        if (true || (true || true)) {
            v8 = (-(v8))
            if true {
                v8 = (-(v8))
            }
        } else {
            v59 := GenS2{F0: (false && ((false && false) && false))}
            v57 += v57
            var v60 uint64 = 0
            _ = v59
            _ = v60
        }
        _ = fi58
    }
    var v61 uint8 = (^(v9))
    for fi62 := 0; fi62 < 10; fi62++ {
        var v63 int16 = (^(v8))
        _ = fi62
        _ = v63
    }
    v64 := GenS6{F0: v3, F1: v3}
    var v65 uint32 = (v3 % ((42) | 1))
    v66 := GenS6{F0: (v3 | (v3 & (^(v65)))), F1: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v61
    _ = v64
    _ = v65
    _ = v66
    _ = v8
    _ = v9
}
