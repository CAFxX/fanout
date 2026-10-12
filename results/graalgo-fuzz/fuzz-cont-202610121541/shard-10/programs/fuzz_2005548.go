package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS9 struct { F0 int; F1 bool; F2 uint64 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint }
type GenS11 struct { F0 int; F1 int32 }
func genFn0(p0 int16, p1 bool, p2 int16) bool {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: "\n\t", F1: 0}
        for fi6 := 0; fi6 < 2; fi6++ {
            p0 = p0
            switch (fi6 / ((fi6) | 1)) {
            case -2147483648:
                var v7 int64 = 0
                _ = v7
            case 4294967295:
                if (1 < 65535) {
                    switch (fi4 & fi4) {
                    case 806:
                        v8 := []int{fi6, fi4, (fi6 + (fi4 - 32768))}
                        _ = v8
                    }
                    v9 := GenS1{F0: "", F1: -1}
                    _ = v9
                } else {
                    p0 -= 1000
                }
            case 897:
                var v10 int16 = 820
                _ = v10
            }
            v11 := GenS2{F0: (p0 + (p0 | p0)), F1: 100}
            _ = fi6
            _ = v11
        }
        p1 = (true || ((0 == 1) && false))
        _ = fi4
        _ = v5
    }
    var v12 float64 = -2.25
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint32 = 891
        var v15 float32 = float32((v12 * float64(-2.25)))
        var v16 float32 = float32(v12)
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (true || p1)
}
func genFn1(p0 uint32, p1 int64, p2 uint64) uint8 {
    p1 = p1
    if true {
        p1 = p1
        var v17 uint8 = 42
        p0 = (p0 >> ((42) & 63))
        _ = v17
    } else {
        var v18 string = ("hello" + ("hello" + ("foo bar" + "a")))
        if true {
            if (true && (0 >= 417)) {
                if true {
                    fmt.Println(p2, p1, p1)
                    p2 = (p2 - 4294967295)
                } else {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v20 := []int{127, 42, 128}
                        _ = fi19
                        _ = v20
                    }
                }
                var v21 uint16 = 594
                var v22 uint8 = uint8(p2)
                _ = v21
                _ = v22
            }
        }
        switch (v18 + v18) {
        case "foo bar":
            v23 := []int{-886, 4294967295, 127}
            _ = v23
        case "hello":
            switch 492 {
            case 733:
                v24 := []int{-485, 32767, 32768, -540}
                _ = v24
            case 411:
                p2 += (p2 - 4294967295)
            case 128:
                var v25 int8 = -128
                _ = v25
            default:
                for fi26 := 0; fi26 < 5; fi26++ {
                    p0 = 194
                    _ = fi26
                }
            }
        case "\n\t":
            v27 := GenS3{F0: (-(p1))}
            _ = v27
        }
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 622
    var v3 float64 = (float64((-2.25 + (1.5 + ((float32(float64(float32((1.5 - float64((0.0 * -2.25)))))) / 0.1) - ((0.0 * -2.25) + 3.14159))))) * 0.0)
    var v28 int8 = 0
    v3 = (v3 / float64((float32(0.1) + (0.1 * 1e10))))
    for fi29 := 0; fi29 < 4; fi29++ {
        v28 = 42
        for fi30 := 0; fi30 < 8; fi30++ {
            v3 += 0.0
            _ = fi30
        }
        fmt.Println(fi29, v3)
        _ = fi29
    }
    v31 := GenS2{F0: v2, F1: (v28 % ((v28) | 1))}
    fmt.Println(v28, v2, v3)
    v32 := []int{42}
    v28 = (-(v28))
    switch 32768 {
    case -2147483648:
        v3 = (float64((((0.1 / 1.0) + -2.25) * (((3.14159 + 1e10) - 0.0) * 1.5))) * ((0.1 / v3) / v3))
    }
    var v33 string = "a"
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS1{F0: ((v33 + v33) + v33), F1: v28}
        v36 := GenS1{F0: v33, F1: v28}
        for fi37 := 0; fi37 < 15; fi37++ {
            if (!((!true) || true)) {
                fmt.Println(v2)
            }
            switch (fi34 - fi34) {
            case 32767:
                var v38 int64 = -9223372036854775808
                _ = v38
            default:
                switch v33 {
                case "a":
                    v39 := GenS2{F0: (v2 ^ v2), F1: v28}
                    _ = v39
                case "hello":
                    switch 127 {
                    case 32768:
                        v40 := GenS4{F0: fi34}
                        _ = v40
                    case 9223372036854775807:
                        switch v33 {
                        case "a":
                            var v41 int = (fi34 << ((4294967295) & 63))
                            _ = v41
                        case "":
                            v2 += (^(v2))
                        default:
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        }
                    case -1:
                        switch 0 {
                        case -708:
                            fmt.Println(v28, fi37, fi37)
                        case -399:
                            fmt.Println(v2)
                        default:
                            fmt.Println(fi34, fi34)
                        }
                    default:
                        v2 = v2
                    }
                default:
                    if ((2 <= 42) || true) {
                        if (!true) {
                            switch v33 {
                            case "\n\t":
                                v32 = append(v32, 32768)
                            case "foo bar":
                                v28 = (v28 << ((32) & 63))
                            case "a":
                                v42 := GenS2{F0: v2, F1: (v28 << ((18446744073709551615) & 63))}
                                _ = v42
                            }
                            v2 = (v2 - (-(v2)))
                            switch fi37 {
                            case 693:
                                _ = len(v1)
                            case 9223372036854775807:
                                var v43 uint16 = 42
                                _ = v43
                            case 345:
                                delete(v1, v33)
                                var v44 uint8 = 0
                                _ = v44
                            default:
                                switch fi34 {
                                case 127:
                                    var v45 int64 = 95
                                    _ = v45
                                case 42:
                                    v2 += -563
                                }
                            }
                        } else {
                            delete(v1, v33)
                            var v46 int = (fi37 | fi37)
                            switch (fi37 << ((2) & 63)) {
                            case 128:
                                switch 128 {
                                case 0:
                                    switch v33 {
                                    case "hello":
                                        v47 := v1[("" + "foo bar")]
                                        _ = v47
                                    case "a":
                                        v48 := GenS5{F0: 999}
                                        _ = v48
                                    case "":
                                        _ = len(v32)
                                    default:
                                        switch v46 {
                                        case 32767:
                                            v3 = float64(-2.25)
                                        default:
                                            v49 := GenS6{F0: -926, F1: v2}
                                            _ = v49
                                        }
                                    }
                                    switch v33 {
                                    case "\n\t":
                                        v33 = v33
                                    case "foo bar":
                                        switch 425 {
                                        case -2147483648:
                                            if false {
                                                switch "hello" {
                                                case "a":
                                                    v33 = "\n\t"
                                                case "foo bar":
                                                    _ = len(v32)
                                                case "":
                                                    if (78 != 1) {
                                                        if ((^(v28)) != v28) {
                                                            if false {
                                                                var v50 uint = 535
                                                                var v51 uint16 = 65535
                                                                var v52 uint = 222
                                                                _ = v50
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                switch v33 {
                                                                case "\n\t":
                                                                    v53 := GenS7{F0: (v2 / (((^(v2))) | 1))}
                                                                    _ = v53
                                                                }
                                                            }
                                                            v2 = v2
                                                            v1[("" + (v33 + ("" + v33)))] = false
                                                        } else {
                                                            v46 = 2147483647
                                                            var v54 float32 = float32((v3 * v3))
                                                            _ = v54
                                                        }
                                                    } else {
                                                        v46 += fi34
                                                    }
                                                default:
                                                    v3 = (v3 - (v3 / v3))
                                                    if (0 > -2147483648) {
                                                        var v55 int32 = 1
                                                        v2 = 42
                                                        v56 := GenS3{F0: 495}
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        var v57 int = -9223372036854775808
                                                        delete(v1, v33)
                                                        _ = v57
                                                    }
                                                }
                                                var v58 int8 = 1
                                                _ = v58
                                            }
                                        }
                                    default:
                                        var v59 float64 = (v3 * 1.5)
                                        if (65535 < 2) {
                                            v60 := GenS3{F0: -132}
                                            _ = len(v1)
                                            if (871 != 2) {
                                                switch v33 {
                                                case "\n\t":
                                                    delete(v1, v33)
                                                case "a":
                                                    v59 = v3
                                                }
                                                if (!(42 >= 2147483647)) {
                                                    if len(v32) > 0 {
                                                        v32[0] = 42
                                                    }
                                                    v3 += 0.1
                                                } else {
                                                    v59 -= ((3.14159 / (-2.25 / v3)) * 1e10)
                                                    var v61 float32 = -2.25
                                                    fmt.Println(v2, fi34)
                                                    var v62 string = ((v33 + (v33 + v33)) + v33)
                                                    _ = v61
                                                    _ = v62
                                                }
                                            } else {
                                                v33 = v33
                                                fmt.Println(v46)
                                                v63 := GenS7{F0: 1000}
                                                _ = v63
                                            }
                                            _ = v60
                                        } else {
                                            v64 := GenS7{F0: v2}
                                            v33 = "a"
                                            v65 := GenS8{F0: 4294967295, F1: 255, F2: 243}
                                            _ = v64
                                            _ = v65
                                        }
                                        switch v33 {
                                        case "a":
                                            v66 := GenS2{F0: v2, F1: v28}
                                            _ = v66
                                        }
                                        _ = v59
                                    }
                                case 9223372036854775807:
                                    fmt.Println(v28)
                                case -549:
                                    v67 := GenS9{F0: (fi37 % ((-9223372036854775808) | 1)), F1: (!(42 <= 42)), F2: 18446744073709551615}
                                    _ = v67
                                default:
                                    switch v33 {
                                    case "\n\t":
                                        v68 := GenS10{F0: 42, F1: 127, F2: 2}
                                        _ = v68
                                    case "foo bar":
                                        var v69 int = fi34
                                        _ = v69
                                    case "hello":
                                        v70 := GenS7{F0: (v2 % (((^(v2))) | 1))}
                                        _ = v70
                                    }
                                }
                            }
                            _ = v46
                        }
                        if false {
                            var v71 int64 = -9223372036854775808
                            v72 := GenS11{F0: (fi34 >> ((255) & 63)), F1: -649}
                            v3 += float64(1.5)
                            _ = v71
                            _ = v72
                        }
                        var v73 uint64 = 42
                        var v74 float64 = v3
                        v75 := GenS9{F0: fi34, F1: true, F2: (v73 - (v73 - v73))}
                        _ = v73
                        _ = v74
                        _ = v75
                    } else {
                        v76 := GenS5{F0: 770}
                        _ = v76
                    }
                }
            }
            v33 = "foo bar"
            _ = fi37
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v33 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
