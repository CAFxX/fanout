package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8, p1 uint32, p2 uint) int64 {
    v7 := []int{251, -2147483648}
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 float32 = (1.5 * 1e10)
        v11 := GenS1{F0: p1}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 string) uint32 {
    fmt.Println(p0)
    if (0 == 1) {
        p0 = p0
        fmt.Println(p0)
    }
    _ = p0
    return 65535
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 32767
    var v3 uint = 18446744073709551615
    var v4 int16 = v2
    var v5 string = ""
    v6 := make(map[string]bool)
    v12 := GenS1{F0: 2147483648}
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v1 = v1
            v15 := GenS1{F0: 42}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            if false {
                switch (v5 + v5) {
                case "\n\t":
                    var v17 int32 = 42
                    var v18 int64 = 834
                    _ = v17
                    _ = v18
                case "a":
                    v5 = (v5 + "")
                case "foo bar":
                    v3 -= 812
                default:
                    if (0.0 >= 0.0) {
                        v19 := GenS1{F0: 705}
                        if true {
                            var v20 int = 904
                            _ = v20
                        } else {
                            fmt.Println(fi16, v3)
                            fmt.Println(v5, fi16, fi13)
                            var v21 uint = (^(v3))
                            if (false || false) {
                                v5 = v5
                            }
                            v22 := GenS2{F0: 0, F1: v1, F2: 9223372036854775807}
                            if true {
                                v2 += (v4 - v4)
                                if (-1 != 1) {
                                    if true {
                                        v3 += (^(v21))
                                        v23 := []int{(^(fi16)), (fi13 + fi16), -9223372036854775808, (fi16 % ((fi13) | 1))}
                                        _ = v23
                                    } else {
                                        var v24 float32 = 0.0
                                        _ = v24
                                    }
                                    fmt.Println(v1, fi13, v21)
                                    v25 := GenS2{F0: 42, F1: (v1 % ((int8(fi16)) | 1)), F2: 1}
                                    _ = v25
                                }
                                v26 := GenS1{F0: genFn1(v5)}
                                _ = v26
                            } else {
                                v4 = v4
                                v21 += (v3 % (((v3 << (((v3 | (^(v3)))) & 63))) | 1))
                                v5 = ("a" + (v5 + v5))
                                fmt.Println(v4, v3)
                            }
                            _ = v21
                            _ = v22
                        }
                        v27 := GenS1{F0: 1}
                        _ = v19
                        _ = v27
                    } else {
                        if false {
                            _ = len(v6)
                            v1 += v1
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    var v28 int32 = 123
                                    v29 := []int{fi13, (fi13 % ((int(v3)) | 1)), (fi13 + 128), (fi13 % ((fi16) | 1))}
                                    v29 = append(v29, (fi16 % (((fi16 + fi16)) | 1)))
                                    _ = v28
                                    _ = v29
                                case "":
                                    v30 := GenS2{F0: 42, F1: -1, F2: 9223372036854775807}
                                    _ = v30
                                }
                            }
                        }
                    }
                    v4 -= (v2 + v2)
                    v3 = v3
                }
                v31 := GenS3{F0: v3, F1: 128}
                _ = v31
            }
            switch "foo bar" {
            case "foo bar":
                if (true || ((v4 ^ v2) == 1000)) {
                    v6[v5] = false
                    v4 = (v2 & v2)
                    v32 := GenS4{F0: 255, F1: 128}
                    _ = v32
                }
            default:
                switch fi16 {
                case 29:
                    v33 := []int{(fi16 ^ 716), -534, (fi13 - (fi13 >> ((42) & 63)))}
                    _ = v33
                }
            }
            var v34 float64 = 0.0
            _ = fi16
            _ = v34
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            switch fi13 {
            case 127:
                v36 := GenS1{F0: 618}
                _ = v36
            }
            v4 -= v4
            _ = fi35
        }
        _ = fi13
    }
    if (0 != (v1 | 127)) {
        var v37 uint = (v3 * (^(v3)))
        v2 = (v2 + (^(v2)))
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 bool = true
        var v40 float64 = -2.25
        _ = fi38
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(v1, fi41)
        switch (fi41 | (fi41 - (fi41 ^ 79))) {
        case 147:
            var v42 int32 = 0
            _ = v42
        case 419:
            var v43 int64 = -452
            _ = v43
        case 9223372036854775807:
            var v44 float32 = 1.5
            _ = v44
        default:
            v45 := GenS5{F0: (^(v1))}
            _ = v45
        }
        _ = fi41
    }
    v3 -= (v3 / ((v3) | 1))
    if false {
        var v46 int8 = (v1 << (((v3 ^ (v3 >> (((v3 / ((v3) | 1))) & 63)))) & 63))
        if (v3 == (v3 | (v3 / ((255) | 1)))) {
            delete(v6, v5)
            if true {
                v47 := []int{42, 898}
                _ = v47
            } else {
                switch v5 {
                case "":
                    var v48 uint16 = 32767
                    _ = v48
                case "hello":
                    var v49 int32 = -880
                    _ = v49
                }
                if (!((-2147483648 != 0) || (-128 <= (v46 - v46)))) {
                    v50 := GenS4{F0: 0, F1: uint8(v3)}
                    var v51 uint32 = 651
                    v52 := []int{}
                    _ = v50
                    _ = v51
                    _ = v52
                } else {
                    fmt.Println(v4)
                }
                v4 = v2
            }
        }
        _ = v46
    }
    if true {
        fmt.Println(v3)
        v1 = v1
        var v53 bool = (uint16(v3) != 645)
        _ = v53
    } else {
        fmt.Println(v4, v3, v5)
        fmt.Println(v3, v4, v4)
    }
    var v54 uint32 = 65535
    switch v5 {
    case "a":
        if false {
            if (2147483647 < -740) {
                switch (v5 + ("a" + ("\n\t" + "a"))) {
                case "\n\t":
                    if ((!(42 == 174)) || true) {
                        for fi55 := 0; fi55 < 18; fi55++ {
                            v2 -= -32768
                            _ = fi55
                        }
                        v3 = 297
                        if false {
                            fmt.Println(v1, v3, v4)
                            if true {
                                if false {
                                    var v56 float64 = 0.0
                                    var v57 int32 = 1
                                    for fi58 := 0; fi58 < 18; fi58++ {
                                        var v59 int16 = int16(fi58)
                                        if (!false) {
                                            v5 = ""
                                            v4 -= -32768
                                            v60 := []int{(fi58 + fi58), (^(fi58)), fi58}
                                            _ = v60
                                        }
                                        switch 0 {
                                        case -686:
                                            v61 := GenS5{F0: v1}
                                            _ = v61
                                        case -2147483648:
                                            v62 := GenS4{F0: (^(v3)), F1: 127}
                                            _ = v62
                                        case 136:
                                            fmt.Println(v56)
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                    _ = v56
                                    _ = v57
                                }
                                var v63 uint8 = 127
                                _ = v63
                            }
                        } else {
                            var v64 uint64 = 2
                            _ = v64
                        }
                    }
                case "hello":
                    var v65 int64 = 9223372036854775807
                    _ = v65
                }
            } else {
                v3 = v3
                for fi66 := 0; fi66 < 6; fi66++ {
                    var v67 uint16 = 225
                    v3 = (v3 * v3)
                    _ = fi66
                    _ = v67
                }
            }
            for fi68 := 0; fi68 < 18; fi68++ {
                for fi69 := 0; fi69 < 11; fi69++ {
                    switch fi68 {
                    case -225:
                        v70 := GenS6{F0: 607, F1: v4}
                        _ = v70
                    case -915:
                        v71 := GenS6{F0: genFn1(("\n\t" + v5)), F1: (v4 << ((v3) & 63))}
                        _ = v71
                    }
                    v4 -= 1
                    _ = fi69
                }
                _ = fi68
            }
        } else {
            for fi72 := 0; fi72 < 10; fi72++ {
                if false {
                    v4 -= (v2 / ((v4) | 1))
                    v54 = v54
                    v1 += v1
                }
                _ = fi72
            }
            v1 -= -17
        }
    case "":
        var v73 float32 = (1e10 * 3.14159)
        _ = v73
    case "hello":
        var v74 uint64 = 0
        _ = v74
    default:
        v4 -= (v2 + (v2 & v2))
    }
    var v75 int16 = 1000
    for fi76 := 0; fi76 < 8; fi76++ {
        v3 = v3
        _ = fi76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v75
}
