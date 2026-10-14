package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint64; F1 int32; F2 int8 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int; F1 int }
func genFn0(p0 int) string {
    p0 = p0
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = p0
            v8 := GenS1{F0: -571}
            p0 = fi6
            _ = fi7
            _ = v8
        }
        var v9 int32 = -1
        var v10 int32 = v9
        _ = fi6
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    p0 += p0
    _ = p0
    return "\n\t"
}
func genFn1(p0 int32, p1 int8) int16 {
    switch "foo bar" {
    case "foo bar":
        v11 := []int{-9223372036854775808}
        _ = v11
    case "\n\t":
        v12 := []int{9223372036854775807, -135, -1}
        _ = v12
    default:
        var v13 uint8 = 42
        _ = v13
    }
    switch -495 {
    case -463:
        var v14 int = 521
        _ = v14
    case 2147483647:
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, p1, fi15)
            if ((false || true) && true) {
                p1 = (p1 * (p1 | (p1 * (p1 << ((866) & 63)))))
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS1{F0: 0}
                    var v18 uint64 = 18446744073709551615
                    fmt.Println(fi15)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            v19 := []int{fi15}
            _ = fi15
            _ = v19
        }
    default:
        v20 := []int{0, 32767}
        _ = v20
    }
    _ = p0
    _ = p1
    return 1000
}

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 14
    v3 := make(map[uint32]string)
    var v4 int32 = 1
    v5 := make(map[string]bool)
    switch 32768 {
    case -2147483648:
        var v21 int8 = 79
        _ = v21
    case -9223372036854775808:
        v1 -= v1
    default:
        fmt.Println(v4)
    }
    var v22 uint8 = (v2 - (v2 % (((^(v2))) | 1)))
    v23 := GenS1{F0: 0}
    if true {
        v24 := v5["foo bar"]
        delete(v5, ((("" + "a") + "hello") + ("" + (("hello" + "\n\t") + "hello"))))
        _ = v24
    } else {
        v22 = (^(v22))
        for fi25 := 0; fi25 < 5; fi25++ {
            v4 = (v4 & (v4 / (((v4 & (v4 + v4))) | 1)))
            v1 += (v1 ^ (v1 / ((v1) | 1)))
            v4 -= 0
            _ = fi25
        }
    }
    var v26 int8 = 1
    switch 1 {
    case 32767:
        v3[671] = ("a" + ("\n\t" + "foo bar"))
    case 267:
        v27 := GenS2{F0: v4, F1: 32767, F2: v4}
        _ = v27
    case 42:
        var v28 uint8 = 0
        _ = v28
    default:
        if (v1 < v1) {
            switch (("a" + "\n\t") + "hello") {
            case "foo bar":
                var v29 int16 = -214
                _ = v29
            case "":
                for fi30 := 0; fi30 < 7; fi30++ {
                    v31 := GenS3{F0: 4294967295}
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v33 := v5["a"]
                        _ = fi32
                        _ = v33
                    }
                    v34 := GenS1{F0: 0}
                    _ = fi30
                    _ = v31
                    _ = v34
                }
            default:
                var v35 string = "hello"
                _ = v35
            }
        } else {
            switch "a" {
            case "hello":
                var v36 float32 = 3.14159
                _ = v36
            case "a":
                v37 := GenS4{F0: 2147483647, F1: uint32(v1), F2: -1}
                _ = v37
            }
            v38 := GenS5{F0: 2, F1: 541}
            _ = v38
        }
    }
    fmt.Println(v22, v26, v1)
    v39 := []int{1, 932}
    v1 -= (v1 << ((4294967295) & 63))
    v22 += v2
    if true {
        switch (("a" + "foo bar") + "foo bar") {
        case "a":
            for fi40 := 0; fi40 < 15; fi40++ {
                for fi41 := 0; fi41 < 5; fi41++ {
                    fmt.Println(v26)
                    if (0 != 0) {
                        var v42 int16 = 1000
                        var v43 uint64 = 18446744073709551615
                        switch fi40 {
                        case 128:
                            var v44 uint32 = 491
                            _ = v44
                        case 2147483647:
                            v45 := GenS6{F0: uint64(v1), F1: 0, F2: 100}
                            _ = v45
                        default:
                            v5["\n\t"] = true
                        }
                        _ = v42
                        _ = v43
                    }
                    var v46 uint8 = (v2 ^ v22)
                    var v47 uint = 255
                    _ = fi41
                    _ = v46
                    _ = v47
                }
                _ = fi40
            }
        case "":
            if false {
                v4 -= (v4 >> ((18446744073709551615) & 63))
            } else {
                for fi48 := 0; fi48 < 2; fi48++ {
                    v1 = 2
                    for fi49 := 0; fi49 < 7; fi49++ {
                        v50 := GenS7{F0: "a", F1: 255}
                        var v51 string = "a"
                        v52 := GenS3{F0: 760}
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    _ = fi48
                }
            }
        case "hello":
            if ((v22 % ((0) | 1)) < (v2 % ((v2) | 1))) {
                var v53 uint16 = (v1 * v1)
                var v54 uint16 = v1
                _ = v53
                _ = v54
            }
        default:
            if true {
                for fi55 := 0; fi55 < 6; fi55++ {
                    _ = len(v39)
                    for fi56 := 0; fi56 < 11; fi56++ {
                        var v57 float32 = float32(float64(1e10))
                        v22 = (v22 >> ((uint(v22)) & 63))
                        if (0 > 4294967295) {
                            var v58 float32 = float32(0.0)
                            _ = v58
                        }
                        _ = fi56
                        _ = v57
                    }
                    for fi59 := 0; fi59 < 17; fi59++ {
                        if (109 > 42) {
                            var v60 int8 = (v26 - (v26 ^ v26))
                            switch ("a" + "a") {
                            case "\n\t":
                                _ = len(v39)
                            default:
                                var v61 int64 = -583
                                _ = v61
                            }
                            fmt.Println(fi55, v1)
                            v4 = (-(v4))
                            if (3.14159 != (float64((1e10 / (((float32(float64(0.1)) - 3.14159) * 3.14159) - 0.1))) / 1e10)) {
                                fmt.Println(v4, v22)
                            } else {
                                v62 := GenS6{F0: 0, F1: v4, F2: (v26 ^ -128)}
                                fmt.Println(v2, fi55)
                                v63 := GenS3{F0: uint64(v1)}
                                _ = v62
                                _ = v63
                            }
                            _ = v60
                        } else {
                            v26 = v26
                        }
                        _ = len(v3)
                        if (-474 != -1) {
                            switch "foo bar" {
                            case "\n\t":
                                switch fi59 {
                                case -9223372036854775808:
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        v3[0] = "hello"
                                    default:
                                        delete(v3, 2147483647)
                                    }
                                case 1:
                                    fmt.Println(fi55)
                                default:
                                    v26 -= (v26 >> ((932) & 63))
                                    delete(v3, 65535)
                                    switch fi55 {
                                    case 127:
                                        v22 -= 255
                                    }
                                }
                            case "":
                                v22 = (v2 - v2)
                            default:
                                if (0 <= 1) {
                                    var v64 int = 1
                                    var v65 uint32 = 129
                                    v26 = v26
                                    v64 -= 4294967295
                                    if false {
                                        v66 := GenS8{F0: int(v4), F1: 128}
                                        _ = v66
                                    } else {
                                        v1 = v1
                                    }
                                    _ = v64
                                    _ = v65
                                }
                            }
                            switch (fi59 % ((fi59) | 1)) {
                            case 128:
                                var v67 uint32 = 0
                                _ = v67
                            default:
                                fmt.Println(fi55, v22)
                            }
                            var v68 int64 = -875
                            _ = v68
                        }
                        var v69 int16 = genFn1(42, (v26 / ((-1) | 1)))
                        _ = fi59
                        _ = v69
                    }
                    _ = fi55
                }
            }
        }
        switch 42 {
        case 2147483647:
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
        case 4294967295:
            for fi70 := 0; fi70 < 5; fi70++ {
                var v71 bool = false
                v1 = v1
                var v72 float32 = float32(1e10)
                _ = fi70
                _ = v71
                _ = v72
            }
        default:
            v73 := GenS4{F0: 737, F1: 371, F2: -322}
            _ = v73
        }
        switch 38 {
        case 9223372036854775807:
            for fi74 := 0; fi74 < 3; fi74++ {
                v75 := GenS7{F0: (("hello" + ("" + "hello")) + "foo bar"), F1: 0}
                delete(v3, 2147483647)
                _ = fi74
                _ = v75
            }
        default:
            _ = len(v39)
        }
    }
    v1 = 904
    fmt.Println(v1, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
