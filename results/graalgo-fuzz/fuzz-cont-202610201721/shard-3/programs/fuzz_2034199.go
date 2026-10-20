package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint64; F2 string }
func genFn0(p0 int32, p1 uint) uint8 {
    v5 := GenS1{F0: (p1 / (((p1 & 1)) | 1)), F1: 2, F2: (p0 - 394)}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: p1, F1: 4294967295, F2: 1}
        _ = fi6
        _ = v7
    }
    if (42 >= 175) {
        switch "hello" {
        case "a":
            switch "foo bar" {
            case "\n\t":
                p0 = p0
            }
        case "\n\t":
            p1 -= (p1 & (p1 & (p1 & p1)))
        case "foo bar":
            v8 := []int{-536, -1, 128}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                p1 = (p1 >> (((p1 & 255)) & 63))
                p0 = p0
                _ = fi10
            }
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := []int{}
                var v13 float32 = ((0.0 * 0.1) * -2.25)
                v14 := GenS1{F0: (p1 | p1), F1: 2, F2: p0}
                if (true && (((p1 >> (((p1 >> ((p1) & 63))) & 63)) > 1) && false)) {
                    var v15 uint32 = 972
                    switch -350 {
                    case 128:
                        fmt.Println(p0, p1)
                    case 9223372036854775807:
                        switch "foo bar" {
                        case "\n\t":
                            if (458 != uint(v15)) {
                                _ = len(v12)
                                var v16 uint8 = 2
                                v15 += v15
                                _ = v16
                            } else {
                                var v17 string = "foo bar"
                                var v18 uint32 = v15
                                p0 = 0
                                _ = v17
                                _ = v18
                            }
                            switch 9223372036854775807 {
                            case 2147483647:
                                if true {
                                    fmt.Println(p0)
                                    p0 = p0
                                }
                            default:
                                var v19 int8 = -128
                                _ = v19
                            }
                            switch "" {
                            case "":
                                var v20 float64 = -2.25
                                _ = v20
                            case "hello":
                                _ = len(v12)
                            case "foo bar":
                                v21 := GenS1{F0: (p1 - p1), F1: 0, F2: p0}
                                _ = v21
                            }
                            v13 = ((float32(float64(1e10)) / v13) - v13)
                        case "":
                            if (true || (false && false)) {
                                var v22 uint = p1
                                var v23 int = (^(fi9))
                                _ = len(v12)
                                _ = v22
                                _ = v23
                            } else {
                                var v24 bool = false
                                v25 := GenS2{F0: uint8(v15)}
                                _ = v24
                                _ = v25
                            }
                        }
                    }
                    _ = v15
                } else {
                    switch "\n\t" {
                    case "a":
                        var v26 uint8 = 114
                        _ = v26
                    }
                    switch (^(fi9)) {
                    case 42:
                        fmt.Println(v13)
                    case 0:
                        v13 = (v13 * v13)
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            if ("foo bar" > ("\n\t" + "foo bar")) {
                v27 := []int{32767, (fi9 - 2147483647), 4294967295}
                for fi28 := 0; fi28 < 7; fi28++ {
                    if true {
                        v27 = append(v27, -2147483648)
                        var v29 uint16 = 32767
                        p0 += (-(p0))
                        _ = v29
                    }
                    _ = fi28
                }
                var v30 int32 = (^(p0))
                _ = v27
                _ = v30
            } else {
                p0 = p0
                p1 += 296
                v31 := []int{(^(fi9))}
                _ = v31
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 16
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 uint32 = 65535
    v4 := make(map[uint32]bool)
    switch 154 {
    case 42:
        fmt.Println(v3)
    }
    v32 := GenS3{F0: 255, F1: 18446744073709551615}
    v33 := GenS3{F0: uint16(v3), F1: 24}
    v34 := GenS4{F0: 908, F1: -2147483648}
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            if false {
                v3 -= 0
                v3 = (v3 - 347)
                var v37 uint = 18446744073709551615
                _ = v37
            }
            _ = fi36
        }
        var v38 int = fi35
        delete(v4, v3)
        _ = fi35
        _ = v38
    }
    v3 -= v3
    v39 := GenS4{F0: 42, F1: -1}
    v3 = (v3 << ((4294967295) & 63))
    v3 = (v3 + 2)
    switch "a" {
    case "hello":
        var v40 string = "foo bar"
        _ = v40
    default:
        v41 := GenS5{F0: 1, F1: -764}
        _ = v41
    }
    var v42 string = (((((("hello" + "hello") + "a") + "\n\t") + ("" + "")) + ("\n\t" + "\n\t")) + ("\n\t" + ("hello" + ("a" + "hello"))))
    v42 = v42
    for fi43 := 0; fi43 < 14; fi43++ {
        for fi44 := 0; fi44 < 10; fi44++ {
            v3 = (^(v3))
            switch v42 {
            case "a":
                delete(v4, (v3 ^ v3))
            case "":
                if false {
                    fmt.Println(fi43)
                    v45 := v4[(^(v3))]
                    if true {
                        var v46 float64 = 0.0
                        if v45 {
                            var v47 int = fi44
                            if (42 == int32(v3)) {
                                var v48 int8 = 5
                                _ = v48
                            } else {
                                fmt.Println(v42)
                            }
                            v3 = 4294967295
                            var v49 uint32 = (v3 + 291)
                            var v50 int16 = 222
                            _ = v47
                            _ = v49
                            _ = v50
                        } else {
                            switch fi43 {
                            case 373:
                                if ((true || v45) && v45) {
                                    var v51 uint64 = 570
                                    _ = v51
                                } else {
                                    var v52 uint64 = 516
                                    _ = v52
                                }
                            case 189:
                                var v53 int64 = 1
                                _ = v53
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (-(fi44))
                                }
                            }
                            var v54 int8 = 127
                            v55 := GenS6{F0: (v42 + v42), F1: 42}
                            _ = len(v1)
                            _ = v54
                            _ = v55
                        }
                        _ = v46
                    }
                    _ = v45
                }
            case "foo bar":
                var v56 uint8 = 90
                _ = v56
            default:
                fmt.Println(v42)
            }
            _ = fi44
        }
        if (1 <= 2) {
            v42 = v42
            if ((927 < 42) && (((-2.25 + 1.5) / 1.5) > -2.25)) {
                v3 += v3
            } else {
                var v57 string = ((("" + ((v42 + (v42 + ((v42 + (v42 + ((("hello" + (((v42 + "a") + ("foo bar" + v42)) + v42)) + v42) + v42))) + "a"))) + v42)) + v42) + "\n\t")
                fmt.Println(v42, v3, fi43)
                v58 := GenS1{F0: 42, F1: 42, F2: -810}
                _ = v57
                _ = v58
            }
        }
        fmt.Println(v3, v3)
        _ = fi43
    }
    for fi59 := 0; fi59 < 15; fi59++ {
        var v60 uint = 4294967295
        if true {
            for fi61 := 0; fi61 < 13; fi61++ {
                var v62 uint8 = 0
                if false {
                    fmt.Println(v62, fi59, v60)
                    if true {
                        v62 -= (v62 << ((0) & 63))
                    }
                }
                v63 := GenS7{F0: (v3 << (((v60 ^ v60)) & 63))}
                switch (fi59 ^ (-(fi61))) {
                case 42:
                    v64 := GenS8{F0: v42, F1: 0, F2: v42}
                    _ = v64
                case 2147483647:
                    if ((true || (!(0 != int32(fi61)))) || (0 >= 42)) {
                        v62 = (v62 & (v62 / ((v62) | 1)))
                    }
                    var v65 int = fi59
                    _ = v65
                case 914:
                    var v66 int16 = -537
                    _ = v66
                }
                _ = fi61
                _ = v62
                _ = v63
            }
            var v67 uint = (v60 % ((v60) | 1))
            _ = v67
        }
        _ = fi59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v42
}
