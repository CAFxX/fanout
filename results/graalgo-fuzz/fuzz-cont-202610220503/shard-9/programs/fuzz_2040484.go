package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8 }
func genFn0() uint {
    switch "hello" {
    case "a":
        var v4 int16 = -32768
        _ = v4
    default:
        if ((355 < 1) && false) {
            var v5 string = ""
            _ = v5
        } else {
            if (false || false) {
                switch "foo bar" {
                case "":
                    for fi6 := 0; fi6 < 15; fi6++ {
                        if false {
                            var v7 uint8 = 128
                            v8 := GenS1{F0: -2147483648, F1: v7, F2: 9223372036854775807}
                            var v9 string = ("\n\t" + "\n\t")
                            _ = v7
                            _ = v8
                            _ = v9
                        } else {
                            switch "a" {
                            case "foo bar":
                                var v10 uint16 = 1
                                _ = v10
                            case "a":
                                fmt.Println(fi6)
                            case "":
                                var v11 uint64 = 632
                                _ = v11
                            default:
                                var v12 int32 = 42
                                _ = v12
                            }
                            if (228 >= 255) {
                                var v13 uint32 = 468
                                _ = v13
                            } else {
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    var v15 int64 = 9223372036854775807
                                    v16 := []int{-1, fi14, (fi6 / ((fi14) | 1)), fi6}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 uint16 = 32767
                            _ = v17
                        }
                        switch ("" + ("a" + "foo bar")) {
                        case "":
                            var v18 uint16 = 65535
                            _ = v18
                        default:
                            var v19 string = "\n\t"
                            _ = v19
                        }
                        var v20 uint8 = 2
                        _ = fi6
                        _ = v20
                    }
                }
            }
        }
    }
    v21 := []int{-878, 32768, 32767, 9223372036854775807}
    var v22 uint16 = 980
    for fi23 := 0; fi23 < 13; fi23++ {
        v22 = v22
        _ = fi23
    }
    _ = v21
    _ = v22
    return 4294967295
}
func genFn1(p0 int16, p1 int, p2 int64) int64 {
    v24 := GenS1{F0: -2147483648, F1: 2, F2: p2}
    if (false && false) {
        var v25 uint16 = 89
        _ = v25
    } else {
        p0 = p0
        fmt.Println(p0, p1, p2)
    }
    v26 := GenS1{F0: -1, F1: 128, F2: (p2 * p2)}
    v27 := GenS2{F0: p1, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v26
    _ = v27
    return p2
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 float32 = -2.25
    switch 128 {
    case 9223372036854775807:
        var v28 int16 = 32767
        _ = v28
    case 4294967295:
        var v29 int8 = 40
        _ = v29
    default:
        v3 -= v3
    }
    _ = len(v2)
    if true {
        switch v1 {
        case "hello":
            v3 += v3
        }
        if (2 == 2) {
            v3 -= ((v3 / (1.5 / 1.0)) / v3)
        } else {
            fmt.Println(v1, v1)
            v30 := []int{1, -319, -306}
            var v31 int8 = 80
            _ = v30
            _ = v31
        }
        v32 := []int{-422, 4294967295}
        _ = v32
    } else {
        v33 := GenS1{F0: 525, F1: 127, F2: -1}
        v34 := GenS3{F0: 87, F1: -1}
        _ = len(v2)
        _ = v33
        _ = v34
    }
    v1 = "a"
    v1 = ""
    v1 = v1
    var v35 int32 = 42
    fmt.Println(v35, v1, v3)
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 uint16 = 65535
            v39 := GenS4{F0: -9223372036854775808, F1: 65535}
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi36
    }
    if true {
        v35 = v35
        v40 := GenS4{F0: 667, F1: 42}
        var v41 uint8 = 127
        _ = v40
        _ = v41
    } else {
        var v42 float64 = 1.5
        var v43 int32 = 1
        for fi44 := 0; fi44 < 15; fi44++ {
            v45 := GenS3{F0: 42, F1: genFn1(843, (fi44 | (-(fi44))), 9223372036854775807)}
            v46 := v2[(v1 + (v1 + "a"))]
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = v42
        _ = v43
    }
    var v47 uint = 876
    fmt.Println(v35)
    v3 = -2.25
    switch -1 {
    case 853:
        if (!false) {
            var v48 uint8 = 42
            delete(v2, v1)
            for fi49 := 0; fi49 < 16; fi49++ {
                for fi50 := 0; fi50 < 7; fi50++ {
                    v1 = (((v1 + v1) + v1) + v1)
                    var v51 int32 = v35
                    _ = fi50
                    _ = v51
                }
                var v52 uint8 = (v48 >> ((genFn0()) & 63))
                v53 := []int{(-(fi49)), fi49, 211}
                _ = fi49
                _ = v52
                _ = v53
            }
            _ = v48
        } else {
            for fi54 := 0; fi54 < 16; fi54++ {
                var v55 bool = (fi54 > (fi54 - (fi54 / (((fi54 << (((v47 % ((255) | 1))) & 63))) | 1))))
                _ = fi54
                _ = v55
            }
        }
    case 0:
        if true {
            v56 := GenS5{F0: v1}
            v57 := []int{}
            for fi58 := 0; fi58 < 15; fi58++ {
                var v59 int32 = v35
                _ = fi58
                _ = v59
            }
            _ = v56
            _ = v57
        }
    default:
        for fi60 := 0; fi60 < 19; fi60++ {
            var v61 float64 = float64((0.1 - (v3 + 1.5)))
            if false {
                var v62 int8 = -128
                if (-2.25 < (((-2.25 / -2.25) + ((float32(0.0) * v3) * (1.5 / v3))) - float32(v61))) {
                    v63 := v2[v1]
                    for fi64 := 0; fi64 < 11; fi64++ {
                        var v65 int8 = (v62 << ((441) & 63))
                        if (float64(v3) < v61) {
                            switch 32768 {
                            case 373:
                                if true {
                                    fmt.Println(v35, v1)
                                    v66 := GenS6{F0: 42, F1: genFn1(1000, fi60, 328)}
                                    _ = v66
                                }
                            }
                        } else {
                            v3 = v3
                            var v67 uint16 = 255
                            _ = v67
                        }
                        _ = fi64
                        _ = v65
                    }
                    _ = v63
                }
                var v68 float64 = float64(float32(v61))
                _ = v62
                _ = v68
            } else {
                switch ((v1 + (("foo bar" + v1) + v1)) + "\n\t") {
                case "a":
                    v69 := []int{fi60}
                    _ = v69
                case "\n\t":
                    v70 := []int{}
                    _ = v70
                default:
                    v1 = v1
                }
                for fi71 := 0; fi71 < 10; fi71++ {
                    v72 := GenS7{F0: v35}
                    fmt.Println(v1, fi60)
                    var v73 int16 = 42
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            }
            v74 := GenS8{F0: -101}
            _ = fi60
            _ = v61
            _ = v74
        }
    }
    var v75 int = 1
    for fi76 := 0; fi76 < 14; fi76++ {
        var v77 uint = v47
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v47:", v47)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v47
    _ = v75
}
