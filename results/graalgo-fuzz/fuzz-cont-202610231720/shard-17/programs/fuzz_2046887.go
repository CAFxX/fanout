package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int8; F2 int }
func genFn0(p0 int8, p1 int32) int64 {
    var v7 uint64 = 4294967295
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 = (v7 ^ 2)
        v9 := GenS1{F0: 1, F1: 529, F2: ((p0 + (p0 & (p0 >> ((233) & 63)))) == p0)}
        _ = fi8
        _ = v9
    }
    v7 = 911
    if (0 >= -1) {
        var v10 string = "\n\t"
        var v11 float64 = 0.1
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 int64) uint8 {
    v12 := []int{}
    if (941 == 255) {
        var v13 uint16 = (p0 << ((65535) & 63))
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 1
    v3 := []uint32{65535, 65535, 2147483647, 1}
    var v4 uint16 = 0
    v5 := []string{}
    var v6 float32 = float32(3.14159)
    var v14 int = -643
    if false {
        var v15 float32 = -2.25
        _ = v15
    } else {
        v4 -= v4
        switch (^(v14)) {
        case 42:
            v4 -= 32767
        case 127:
            fmt.Println(v14)
        case 9223372036854775807:
            var v16 uint64 = 173
            _ = v16
        }
    }
    var v17 int64 = 42
    if (255 > 732) {
        v14 += 903
    } else {
        var v18 int8 = 1
        if (!(!((true && ((v18 + (v18 & 0)) <= (v18 + v18))) || (!(2147483647 > 4294967295))))) {
            v2 += 277
        } else {
            var v19 uint32 = 2147483648
            if true {
                v17 -= v17
                var v20 uint = 42
                var v21 uint32 = v19
                _ = v20
                _ = v21
            }
            var v22 int32 = (v2 & v2)
            _ = v19
            _ = v22
        }
        if true {
            var v23 uint = 4294967295
            fmt.Println(v2, v23)
            _ = v23
        }
        _ = v18
    }
    var v24 int64 = genFn0(42, 743)
    for fi25 := 0; fi25 < 2; fi25++ {
        delete(v1, (v14 & fi25))
        _ = fi25
    }
    if ((int64(v14) > (v17 - genFn0(int8(v17), (v2 * v2)))) || ((0 <= 560) || false)) {
        fmt.Println(v17, v17)
    } else {
        var v26 string = "a"
        fmt.Println(v17, v4)
        v2 -= v2
        _ = v26
    }
    v27 := GenS1{F0: v4, F1: -32768, F2: false}
    switch 2147483647 {
    case -1:
        v6 += float32(1e10)
    case 815:
        switch (v14 + 4294967295) {
        case 1:
            if true {
                delete(v1, 32767)
            } else {
                v14 += v14
                v6 -= (v6 - 3.14159)
                var v28 uint32 = 2
                _ = v28
            }
        case 997:
            var v29 int = v14
            _ = v29
        default:
            var v30 uint64 = 1
            _ = v30
        }
    case 0:
        v31 := v1[32768]
        _ = v31
    }
    if true {
        var v32 uint = 18446744073709551615
        if true {
            fmt.Println(v6, v24)
            v33 := GenS2{F0: -1}
            _ = v33
        } else {
            var v34 int16 = 1000
            _ = v34
        }
        switch "foo bar" {
        case "a":
            var v35 uint64 = 4294967295
            _ = v35
        case "\n\t":
            if ((255 < 0) || ((!true) && false)) {
                if false {
                    var v36 int32 = 1
                    _ = v36
                } else {
                    fmt.Println(v14)
                    v17 = (v17 ^ v17)
                    v24 += (^(v24))
                }
                v1[(v14 & (v14 << ((v32) & 63)))] = (!(v4 == v4))
                if false {
                    for fi37 := 0; fi37 < 2; fi37++ {
                        if true {
                            fmt.Println(v6, v24, fi37)
                            var v38 uint32 = 23
                            var v39 uint64 = 42
                            _ = v38
                            _ = v39
                        }
                        switch int(v4) {
                        case 4294967295:
                            switch ("foo bar" + "") {
                            case "hello":
                                var v40 int64 = v17
                                _ = v40
                            }
                        default:
                            var v41 float32 = ((((v6 / ((v6 / (v6 * 0.0)) * v6)) / v6) * (v6 - 0.1)) + 3.14159)
                            _ = v41
                        }
                        _ = fi37
                    }
                    v42 := GenS3{F0: (v24 / (((v17 / ((-9223372036854775808) | 1))) | 1))}
                    v4 = 901
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 16; fi43++ {
                        for fi44 := 0; fi44 < 6; fi44++ {
                            v45 := GenS4{F0: -594, F1: 42, F2: fi44}
                            _ = len(v5)
                            _ = fi44
                            _ = v45
                        }
                        for fi46 := 0; fi46 < 12; fi46++ {
                            v4 = v4
                            _ = fi46
                        }
                        if true {
                            v2 = 42
                        } else {
                            for fi47 := 0; fi47 < 15; fi47++ {
                                v17 -= v24
                                _ = fi47
                            }
                            v1[fi43] = true
                        }
                        _ = fi43
                    }
                }
            }
        case "foo bar":
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 float64 = 1.5
                var v50 uint64 = 869
                for fi51 := 0; fi51 < 19; fi51++ {
                    v1[(fi48 % (((^(fi48))) | 1))] = (!false)
                    switch ("" + "") {
                    case "\n\t":
                        v3 = append(v3, 315)
                    case "":
                        v1[-260] = false
                        _ = len(v1)
                        if false {
                            v14 -= v14
                        } else {
                            fmt.Println(v2)
                            if ((2147483647 != 2147483647) && ((v6 - (-2.25 - ((v6 - -2.25) - v6))) > v6)) {
                                v52 := v1[(v14 - (v14 & v14))]
                                _ = v52
                            }
                            v32 += 18446744073709551615
                        }
                    }
                    _ = fi51
                }
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        _ = v32
    }
    if false {
        v4 = (v4 * (v4 | 1))
        if len(v3) > 0 {
            v3[0] = 42
        }
    } else {
        if (false && true) {
            v1[(v14 / (((v14 << ((1) & 63))) | 1))] = false
        }
    }
    switch (^(v14)) {
    case 32767:
        v53 := GenS3{F0: v17}
        _ = v53
    }
    v6 += float32(((3.14159 * 0.0) / 1.0))
    if ((false || false) && true) {
        var v54 int32 = (v2 * -1)
        _ = v54
    }
    switch -336 {
    case 128:
        v55 := GenS2{F0: int8(v2)}
        _ = v55
    case -207:
        var v56 int16 = -32768
        _ = v56
    }
    var v57 int = (^(v14))
    if (-32768 < -32768) {
        v2 -= (v2 << ((449) & 63))
        v6 += 1e10
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        var v59 int = 9223372036854775807
        if (255 > 18446744073709551615) {
            for fi60 := 0; fi60 < 13; fi60++ {
                v6 -= 1.5
                fmt.Println(v6)
                v61 := GenS2{F0: -128}
                _ = fi60
                _ = v61
            }
            if false {
                var v62 int = v59
                for fi63 := 0; fi63 < 18; fi63++ {
                    v64 := GenS4{F0: -32768, F1: 0, F2: -883}
                    _ = fi63
                    _ = v64
                }
                _ = v62
            } else {
                var v65 int8 = -82
                v57 = 542
                if (!(-72 == (v65 - (v65 / ((-128) | 1))))) {
                    v17 = (^(v24))
                    v24 += v17
                }
                _ = v65
            }
        } else {
            switch "" {
            case "a":
                for fi66 := 0; fi66 < 2; fi66++ {
                    v6 += (((v6 * float32(1.5)) + float32(1e10)) - v6)
                    _ = fi66
                }
            default:
                for fi67 := 0; fi67 < 19; fi67++ {
                    var v68 float32 = (v6 - v6)
                    v68 += v6
                    _ = fi67
                    _ = v68
                }
            }
            v2 = v2
            var v69 string = ("" + "a")
            _ = v69
        }
        v24 = 0
        _ = fi58
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
