package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 int32, p1 int8, p2 int32) int64 {
    p0 += -600
    if false {
        var v5 uint = 18446744073709551615
        v6 := GenS1{F0: 2}
        v7 := GenS1{F0: 424}
        _ = v5
        _ = v6
        _ = v7
    } else {
        switch ("a" + "a") {
        case "foo bar":
            var v8 bool = true
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 19; fi9++ {
                p1 = (-(p1))
                fmt.Println(p2, fi9)
                var v10 uint32 = 2147483647
                _ = fi9
                _ = v10
            }
        case "hello":
            var v11 uint8 = 2
            _ = v11
        }
        v12 := GenS1{F0: 0}
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int16 = 32767
            v15 := GenS1{F0: 0}
            var v16 int8 = (p1 % ((int8(v14)) | 1))
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v12
    }
    var v17 int8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -4
}
func genFn1(p0 int64, p1 int32) int32 {
    p1 = -1
    p0 = 9223372036854775807
    switch ("hello" + (("hello" + "") + "a")) {
    case "":
        var v18 float32 = 3.14159
        _ = v18
    case "\n\t":
        if false {
            v19 := []int{1}
            if false {
                if true {
                    if false {
                        var v20 int8 = -1
                        p0 = (p0 / (((p0 ^ (p0 << ((944) & 63)))) | 1))
                        p1 = p1
                        _ = v20
                    }
                    p0 += p0
                }
                v21 := GenS2{F0: 405, F1: 255, F2: -1}
                var v22 uint16 = 621
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 float32 = 0.1
                    _ = fi23
                    _ = v24
                }
                switch "" {
                case "foo bar":
                    p0 += (p0 / ((genFn0((p1 * p1), -128, p1)) | 1))
                case "hello":
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                }
            }
            _ = v19
        } else {
            if (!false) {
                p0 += (p0 / (((p0 / (((p0 - genFn0(0, 42, 1))) | 1))) | 1))
            } else {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v26 := GenS2{F0: 65535, F1: 2, F2: (p0 + (^(p0)))}
                    var v27 float32 = 1e10
                    v28 := GenS1{F0: 1}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
        }
    case "hello":
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(p0, fi29)
            v30 := []int{fi29, fi29, fi29, int(p0)}
            _ = fi29
            _ = v30
        }
    }
    if true {
        fmt.Println(p0, p0)
    } else {
        v31 := []int{127}
        for fi32 := 0; fi32 < 1; fi32++ {
            switch (fi32 % (((fi32 | fi32)) | 1)) {
            case 329:
                for fi33 := 0; fi33 < 6; fi33++ {
                    if false {
                        var v34 uint = 65535
                        _ = v34
                    } else {
                        if (0 != 71) {
                            fmt.Println(fi33, p1, fi33)
                            var v35 float32 = ((1e10 - 3.14159) / ((1e10 - 3.14159) + 1.5))
                            v36 := GenS3{F0: p1}
                            v37 := GenS2{F0: 42, F1: 506, F2: 42}
                            if len(v31) > 0 {
                                v31[0] = (fi33 + 416)
                            }
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        p1 -= p1
                        p0 -= genFn0(42, int8(p1), 877)
                    }
                    if (fi33 == (-(fi32))) {
                        if true {
                            if (65535 > 345) {
                                fmt.Println(p0, p1)
                            }
                            p0 += p0
                            var v38 int16 = 42
                            fmt.Println(p0, fi32, fi32)
                            v39 := GenS1{F0: 843}
                            v40 := GenS1{F0: 604}
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        var v41 int32 = 930
                        _ = v41
                    } else {
                        v42 := GenS1{F0: 18446744073709551615}
                        v43 := GenS1{F0: 4294967295}
                        _ = len(v31)
                        _ = v42
                        _ = v43
                    }
                    _ = fi33
                }
            case -2147483648:
                p0 -= p0
            default:
                v44 := GenS1{F0: 18446744073709551615}
                _ = v44
            }
            v45 := GenS1{F0: 321}
            _ = fi32
            _ = v45
        }
        _ = v31
    }
    var v46 float32 = (-2.25 * -2.25)
    _ = p0
    _ = p1
    _ = v46
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 18446744073709551615
    var v3 int32 = -1
    v4 := make(map[int]bool)
    v3 = (v3 >> ((v2) & 63))
    switch -1 {
    case 4294967295:
        for fi47 := 0; fi47 < 6; fi47++ {
            fmt.Println(v3, v2)
            for fi48 := 0; fi48 < 12; fi48++ {
                if true {
                    v3 = (v3 | v3)
                    switch (fi47 >> ((v2) & 63)) {
                    case -158:
                        v49 := GenS3{F0: v3}
                        _ = v49
                    case 1:
                        if true {
                            v50 := GenS3{F0: genFn1(42, (v3 | (v3 & v3)))}
                            var v51 uint = 4294967295
                            _ = v50
                            _ = v51
                        } else {
                            v2 += (v2 % (((v2 | (v2 ^ 2))) | 1))
                            switch ("" + "a") {
                            case "":
                                v52 := v4[fi47]
                                _ = v52
                            case "foo bar":
                                var v53 int64 = 1
                                _ = v53
                            case "hello":
                                var v54 int8 = 42
                                _ = v54
                            default:
                                v55 := GenS4{F0: 626, F1: -796, F2: v1}
                                _ = v55
                            }
                            switch "\n\t" {
                            case "hello":
                                v56 := GenS1{F0: 2}
                                _ = v56
                            }
                        }
                    case 0:
                        switch (fi47 & (fi48 >> (((v2 * (v2 | v2))) & 63))) {
                        case 118:
                            var v57 float32 = 3.14159
                            _ = v57
                        case 780:
                            v1 += (v1 | v1)
                        case 103:
                            v2 += v2
                        }
                    }
                    if (!(true && false)) {
                        var v58 uint32 = 2
                        _ = v58
                    }
                }
                v59 := GenS1{F0: 646}
                _ = fi48
                _ = v59
            }
            _ = fi47
        }
    case -892:
        for fi60 := 0; fi60 < 15; fi60++ {
            v3 = (v3 & (v3 << ((v2) & 63)))
            v3 += int32(v1)
            _ = fi60
        }
    default:
        for fi61 := 0; fi61 < 4; fi61++ {
            v2 += (v2 << ((v2) & 63))
            _ = fi61
        }
    }
    v62 := v4[42]
    for fi63 := 0; fi63 < 7; fi63++ {
        v64 := GenS2{F0: 42, F1: v1, F2: 0}
        v65 := GenS3{F0: -2147483648}
        if (v62 && (2147483647 == int(v1))) {
            switch ("foo bar" + (((("\n\t" + ("a" + "foo bar")) + "a") + ("\n\t" + "foo bar")) + "hello")) {
            case "":
                v4[(fi63 % ((fi63) | 1))] = false
            }
        } else {
            switch "" {
            case "\n\t":
                for fi66 := 0; fi66 < 20; fi66++ {
                    if (v2 != 179) {
                        var v67 uint8 = 128
                        fmt.Println(fi63, v3, v62)
                        _ = v67
                    }
                    _ = fi66
                }
            case "foo bar":
                if (false && false) {
                    v68 := []int{}
                    v2 -= (v2 << ((255) & 63))
                    fmt.Println(v3, v3, v62)
                    _ = v68
                }
            case "":
                if v62 {
                    v69 := []int{383, int(v2)}
                    _ = v69
                } else {
                    v1 -= v1
                    fmt.Println(v2, v62)
                }
            }
        }
        _ = fi63
        _ = v64
        _ = v65
    }
    v70 := GenS2{F0: 255, F1: (v1 >> (((v2 % (((v2 * (v2 ^ (v2 & v2)))) | 1))) & 63)), F2: -9223372036854775808}
    fmt.Println(v62, v2, v3)
    v71 := GenS5{F0: -9223372036854775808, F1: 255, F2: -2147483648}
    fmt.Println(v3)
    v62 = v62
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v62
    _ = v70
    _ = v71
}
