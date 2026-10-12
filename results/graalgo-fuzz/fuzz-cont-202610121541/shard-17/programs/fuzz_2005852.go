package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
type GenS4 struct { F0 bool; F1 int32; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int16; F2 bool }
func genFn0(p0 uint16, p1 int64) uint {
    var v7 int32 = 668
    switch 4294967295 {
    case 42:
        switch ("" + "a") {
        case "hello":
            p0 = (p0 + p0)
        default:
            var v8 uint32 = 1
            _ = v8
        }
    case 2147483647:
        v9 := []int{9223372036854775807, 32767}
        _ = v9
    }
    if (2 != 746) {
        var v10 int64 = p1
        _ = v10
    } else {
        var v11 int8 = 127
        var v12 int8 = (v11 + v11)
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        p1 = p1
        _ = fi13
    }
    var v14 int16 = 916
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 255
}

func main() {
    var v1 uint16 = 802
    var v2 string = "a"
    v3 := []uint32{1, 65535}
    v4 := make(map[int]uint32)
    var v5 uint64 = 42
    var v6 uint8 = 1
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 11; fi16++ {
        v6 -= (v6 % (((v6 * v6)) | 1))
        if false {
            var v17 float64 = 0.1
            fmt.Println(v6, v15, v2)
            for fi18 := 0; fi18 < 18; fi18++ {
                _ = len(v4)
                v5 -= uint64(v6)
                switch (v2 + "hello") {
                case "a":
                    switch "a" {
                    case "foo bar":
                        v19 := GenS1{F0: v15, F1: -403, F2: (fi16 % ((fi16) | 1))}
                        _ = v19
                    case "":
                        v17 = ((((0.0 + (v17 - (0.0 / v17))) / v17) - (v17 / v17)) + float64(((float32(v17) * 0.0) + 0.1)))
                    case "hello":
                        v5 = (v5 << ((uint(v5)) & 63))
                    }
                case "foo bar":
                    v15 -= 42
                    v20 := GenS1{F0: (v15 >> (((^(v15))) & 63)), F1: 933, F2: (fi18 + (fi16 + fi16))}
                    switch ("hello" + (v2 + (v2 + "foo bar"))) {
                    case "foo bar":
                        if ((true && ((v1 & v1) < v1)) && false) {
                            v21 := GenS2{F0: v6}
                            _ = v21
                        }
                    case "":
                        if false {
                            var v22 float64 = 0.1
                            switch v2 {
                            case "hello":
                                v23 := GenS1{F0: genFn0((v1 ^ v1), 1), F1: 1, F2: fi18}
                                _ = v23
                            case "a":
                                var v24 bool = ((v2 + v2) == v2)
                                _ = v24
                            default:
                                v17 += ((((0.1 / 0.1) + v22) + (0.0 - v17)) + (v17 * 0.1))
                            }
                            fmt.Println(v2)
                            _ = v22
                        } else {
                            var v25 int8 = 127
                            v26 := GenS2{F0: v6}
                            v15 += (v15 & v15)
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v20
                case "":
                    if ((v6 > 249) || false) {
                        var v27 string = v2
                        _ = v27
                    } else {
                        if false {
                            var v28 int8 = 20
                            _ = v28
                        } else {
                            v29 := GenS1{F0: (v15 / (((v15 + (v15 << ((uint(v6)) & 63)))) | 1)), F1: 42, F2: fi16}
                            _ = v29
                        }
                        switch v2 {
                        case "a":
                            var v30 bool = false
                            _ = v30
                        case "foo bar":
                            fmt.Println(v5, fi18, v6)
                        case "\n\t":
                            var v31 int8 = 105
                            v3 = append(v3, 0)
                            switch "foo bar" {
                            case "a":
                                switch -860 {
                                case 128:
                                    _ = len(v3)
                                case 9223372036854775807:
                                    delete(v4, fi16)
                                case -1:
                                    var v32 int64 = -9223372036854775808
                                    _ = v32
                                }
                            }
                            var v33 uint16 = v1
                            if ((((v2 + ((v2 + v2) + "\n\t")) + "hello") + v2) <= (("" + v2) + "hello")) {
                                var v34 uint16 = (v1 << (((v15 >> ((18446744073709551615) & 63))) & 63))
                                var v35 string = v2
                                _ = v34
                                _ = v35
                            }
                            _ = v31
                            _ = v33
                        default:
                            switch "" {
                            case "a":
                                v4[fi16] = uint32(v5)
                            case "":
                                var v36 int8 = 109
                                _ = v36
                            default:
                                if false {
                                    v37 := GenS3{F0: v1, F1: (false && (!false)), F2: v15}
                                    v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
                                    _ = v37
                                } else {
                                    var v38 float64 = 0.0
                                    v39 := GenS1{F0: (v15 << ((v15) & 63)), F1: 1, F2: (-(fi16))}
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                    }
                }
                _ = fi18
            }
            _ = v17
        }
        _ = fi16
    }
    var v40 int64 = -9223372036854775808
    var v41 uint8 = 19
    v42 := GenS1{F0: v15, F1: v40, F2: 127}
    fmt.Println(v40, v2, v41)
    var v43 float64 = (0.0 - 1e10)
    if false {
        v1 = v1
    } else {
        v44 := GenS4{F0: (true && (!true)), F1: -1, F2: (v40 / ((-1) | 1))}
        var v45 int = 32767
        var v46 float64 = v43
        _ = v44
        _ = v45
        _ = v46
    }
    var v47 float64 = 1.5
    for fi48 := 0; fi48 < 6; fi48++ {
        fmt.Println(v40, fi48, v40)
        v2 = v2
        _ = fi48
    }
    var v49 uint32 = 42
    var v50 uint16 = v1
    switch "\n\t" {
    case "":
        var v51 int32 = 448
        _ = v51
    }
    switch 127 {
    case -489:
        for fi52 := 0; fi52 < 15; fi52++ {
            for fi53 := 0; fi53 < 18; fi53++ {
                var v54 uint16 = (^(v1))
                switch v2 {
                case "hello":
                    v47 -= (((1.5 * 0.0) + ((v47 - v47) - (-2.25 * 3.14159))) + v43)
                case "a":
                    var v55 int32 = int32(v54)
                    _ = v55
                case "":
                    v3 = append(v3, (v49 | (v49 << (((v15 | (v15 & 42))) & 63))))
                default:
                    fmt.Println(v1)
                }
                v56 := v4[fi52]
                _ = fi53
                _ = v54
                _ = v56
            }
            if true {
                switch fi52 {
                case 496:
                    if ((v49 | v49) <= (v49 - (^(v49)))) {
                        var v57 uint16 = (v1 - 65535)
                        for fi58 := 0; fi58 < 18; fi58++ {
                            if ((!true) || false) {
                                v59 := GenS5{F0: (v57 - v57), F1: v15, F2: (v5 | v5)}
                                _ = len(v4)
                                _ = v59
                            } else {
                                _ = len(v3)
                            }
                            _ = fi58
                        }
                        _ = v57
                    }
                case 1:
                    var v60 int = (fi52 - fi52)
                    _ = v60
                case -2147483648:
                    v43 = v43
                }
                v61 := GenS6{F0: int8(v1)}
                _ = v61
            }
            _ = fi52
        }
    case 32767:
        switch -554 {
        case -307:
            var v62 bool = (42 == 42)
            _ = v62
        case -1:
            switch 9223372036854775807 {
            case 128:
                for fi63 := 0; fi63 < 19; fi63++ {
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v49 -= v49
                        var v65 string = "a"
                        v6 = (v6 * v6)
                        _ = fi64
                        _ = v65
                    }
                    var v66 int8 = 100
                    _ = fi63
                    _ = v66
                }
            case -1:
                _ = len(v4)
            }
        case 1:
            v67 := GenS7{F0: -1, F1: 0, F2: (uint16(v15) > v1)}
            _ = v67
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
