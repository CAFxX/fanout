package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int64) uint {
    for fi4 := 0; fi4 < 13; fi4++ {
        p0 -= p0
        var v5 uint16 = 42
        v6 := GenS1{F0: (v5 >= v5), F1: 2}
        _ = fi4
        _ = v5
        _ = v6
    }
    switch 32767 {
    case 2147483647:
        var v7 int16 = 396
        _ = v7
    case 32767:
        v8 := []int{4294967295, 128, -602}
        _ = v8
    default:
        v9 := GenS1{F0: (-1 <= 42), F1: 1}
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS2{F0: 127, F1: 1}
        var v12 int32 = -2147483648
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = (65535 < 955)
    _ = p0
    _ = v13
    return 255
}

func main() {
    var v1 int = -2147483648
    var v2 uint = 2
    v3 := make(map[int]uint32)
    v1 = v1
    v14 := []int{}
    switch "foo bar" {
    case "\n\t":
        for fi15 := 0; fi15 < 14; fi15++ {
            switch "" {
            case "":
                if true {
                    v16 := GenS3{F0: 650, F1: "a"}
                    _ = v16
                }
            case "\n\t":
                var v17 uint16 = 42
                _ = v17
            default:
                var v18 uint32 = 1
                _ = v18
            }
            switch "a" {
            case "foo bar":
                var v19 int16 = 1
                _ = v19
            case "\n\t":
                if (4294967295 <= 42) {
                    v20 := GenS2{F0: 1, F1: 255}
                    if len(v14) > 0 {
                        v14[0] = (v1 * fi15)
                    }
                    _ = v20
                }
            case "":
                if false {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v3[(fi15 + (fi21 + (v1 * (v1 << (((v2 & 255)) & 63)))))] = 1
                        _ = len(v14)
                        _ = fi21
                    }
                    _ = len(v14)
                }
            default:
                v2 -= 1
            }
            _ = fi15
        }
    }
    var v22 uint64 = 42
    _ = len(v3)
    v23 := GenS2{F0: 0, F1: 128}
    fmt.Println(v2, v1)
    if (false || (122 > 184)) {
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(v2, fi24, v2)
            v25 := GenS2{F0: 42, F1: 0}
            _ = fi24
            _ = v25
        }
        v26 := GenS4{F0: 42, F1: 0, F2: (v22 - v22)}
        _ = v26
    } else {
        var v27 float32 = 0.1
        var v28 float32 = 1e10
        if (104 == 179) {
            if (3.14159 > float64(v27)) {
                var v29 int64 = -1
                v1 += (v1 << ((v2) & 63))
                _ = v29
            }
        } else {
            var v30 uint32 = 4294967295
            var v31 string = "a"
            _ = v30
            _ = v31
        }
        _ = v27
        _ = v28
    }
    if (!false) {
        if false {
            if (!(281 > 196)) {
                v32 := GenS2{F0: 78, F1: 2}
                var v33 float64 = ((float64(1e10) - 3.14159) - (1.5 * (0.1 * 1.5)))
                _ = v32
                _ = v33
            } else {
                v14 = append(v14, (v1 << ((1) & 63)))
                v22 -= 18446744073709551615
            }
            v34 := GenS4{F0: 0, F1: 2147483648, F2: (^(v22))}
            for fi35 := 0; fi35 < 9; fi35++ {
                fmt.Println(v22)
                if false {
                    if true {
                        v1 = 42
                        fmt.Println(v1)
                        v3[(v1 >> ((genFn0(1)) & 63))] = 768
                    }
                    var v36 int16 = -280
                    _ = v36
                } else {
                    var v37 uint = v2
                    _ = v37
                }
                _ = fi35
            }
            _ = v34
        } else {
            _ = len(v14)
            switch (("\n\t" + ((("" + "foo bar") + "hello") + "a")) + "hello") {
            case "hello":
                v2 -= 94
            case "\n\t":
                v38 := GenS2{F0: 10, F1: 42}
                _ = v38
            default:
                var v39 uint16 = 2
                _ = v39
            }
            var v40 int32 = -504
            _ = v40
        }
    }
    switch "a" {
    case "":
        for fi41 := 0; fi41 < 13; fi41++ {
            v42 := GenS5{F0: 66, F1: (v2 & v2)}
            _ = fi41
            _ = v42
        }
    case "foo bar":
        if false {
            _ = len(v3)
            var v43 int16 = -1
            _ = v43
        }
    case "a":
        var v44 int = (v1 | (v1 % ((v1) | 1)))
        _ = v44
    }
    switch "" {
    case "a":
        if (!(true && false)) {
            var v45 int8 = 1
            delete(v3, v1)
            _ = v45
        }
    case "foo bar":
        fmt.Println(v22, v1)
    case "hello":
        for fi46 := 0; fi46 < 6; fi46++ {
            switch "\n\t" {
            case "":
                for fi47 := 0; fi47 < 7; fi47++ {
                    v48 := GenS6{F0: 42, F1: 1}
                    _ = fi47
                    _ = v48
                }
            case "\n\t":
                fmt.Println(fi46)
            default:
                if len(v14) > 0 {
                    v14[0] = 32768
                }
            }
            _ = fi46
        }
    default:
        var v49 int = (-(v1))
        _ = v49
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        fmt.Println(v22, v2, v22)
        v51 := GenS5{F0: 56, F1: (v2 | v2)}
        _ = fi50
        _ = v51
    }
    v1 = (v1 | v1)
    for fi52 := 0; fi52 < 20; fi52++ {
        v2 -= (v2 / (((^(v2))) | 1))
        fmt.Println(fi52, fi52, v1)
        _ = fi52
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        if false {
            if false {
                v54 := GenS4{F0: 42, F1: 0, F2: v22}
                _ = v54
            }
            switch "hello" {
            case "hello":
                v1 = fi53
            case "":
                v55 := GenS7{F0: v1, F1: 0}
                _ = v55
            default:
                if true {
                    v2 += 803
                    var v56 bool = (987 > 1)
                    delete(v3, (-(v1)))
                    _ = v56
                }
            }
        } else {
            v3[(fi53 >> (((v2 + uint(v22))) & 63))] = 299
        }
        switch ("" + ("\n\t" + ("\n\t" + "foo bar"))) {
        case "a":
            if (false && (32767 == 65535)) {
                v2 = (^(v2))
                for fi57 := 0; fi57 < 8; fi57++ {
                    v14 = append(v14, (fi53 >> ((212) & 63)))
                    _ = fi57
                }
                v1 -= (v1 >> ((v2) & 63))
            }
        default:
            if (true || true) {
                switch ("foo bar" + "\n\t") {
                case "a":
                    for fi58 := 0; fi58 < 15; fi58++ {
                        v2 = genFn0(9223372036854775807)
                        fmt.Println(v2, v1, fi58)
                        _ = fi58
                    }
                case "hello":
                    if true {
                        v22 = 42
                        var v59 uint32 = 2
                        _ = v59
                    } else {
                        for fi60 := 0; fi60 < 14; fi60++ {
                            _ = len(v3)
                            if false {
                                v61 := GenS2{F0: 255, F1: 127}
                                var v62 uint64 = v22
                                _ = v61
                                _ = v62
                            }
                            _ = fi60
                        }
                        var v63 float32 = 0.0
                        _ = v63
                    }
                case "\n\t":
                    v64 := GenS3{F0: 65535, F1: ("a" + "hello")}
                    _ = v64
                default:
                    if false {
                        fmt.Println(fi53, v2)
                    }
                }
                for fi65 := 0; fi65 < 1; fi65++ {
                    var v66 int8 = 0
                    if (!(!false)) {
                        v3[v1] = 289
                        var v67 int16 = 32767
                        var v68 uint64 = v22
                        _ = v67
                        _ = v68
                    }
                    v3[fi65] = 2
                    _ = fi65
                    _ = v66
                }
            }
        }
        var v69 uint32 = 2
        _ = fi53
        _ = v69
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        _ = len(v14)
    }
    v14 = append(v14, v1)
    if false {
        var v70 int16 = 32767
        if false {
            var v71 uint = v2
            _ = v71
        } else {
            switch "hello" {
            case "":
                var v72 int16 = v70
                _ = v72
            case "\n\t":
                v1 -= -25
            case "hello":
                v73 := GenS8{F0: 2}
                _ = v73
            }
            if (false && (((65535 != 677) || false) && false)) {
                if len(v14) > 0 {
                    v14[0] = (-(v1))
                }
                var v74 float64 = 0.0
                _ = len(v14)
                _ = v74
            } else {
                for fi75 := 0; fi75 < 17; fi75++ {
                    if true {
                        v14 = append(v14, fi75)
                        if (-86 <= 1) {
                            v76 := GenS3{F0: 2, F1: ""}
                            _ = v76
                        }
                    }
                    if (!((-(v1)) < v1)) {
                        v70 = (v70 / ((1000) | 1))
                        for fi77 := 0; fi77 < 20; fi77++ {
                            v22 = (v22 & v22)
                            switch 127 {
                            case -447:
                                v78 := GenS6{F0: 0, F1: (v70 ^ (^(v70)))}
                                _ = v78
                            }
                            _ = fi77
                        }
                    }
                    _ = fi75
                }
                fmt.Println(v22, v22, v70)
            }
            v14 = append(v14, v1)
        }
        _ = v70
    } else {
        switch 4294967295 {
        case -1:
            v3[v1] = uint32(v2)
        case 42:
            v1 -= -598
        case 502:
            var v79 int64 = 42
            _ = v79
        default:
            v2 -= genFn0(9223372036854775807)
        }
        var v80 float64 = (0.0 - 0.1)
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
