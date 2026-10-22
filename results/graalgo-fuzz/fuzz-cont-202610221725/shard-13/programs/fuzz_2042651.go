package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int16; F1 int16; F2 int32 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint, p1 uint16) uint32 {
    if (p1 != (^(p1))) {
        switch 9223372036854775807 {
        case 4294967295:
            var v6 int16 = 1
            _ = v6
        case 2147483647:
            v7 := []int{-9223372036854775808, 709, 32768}
            _ = v7
        }
        var v8 bool = (!(true || (!(uint64(p0) >= 0))))
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float32 = 0.0
        _ = fi9
        _ = v10
    }
    v11 := []int{-2147483648, 0, 797}
    _ = p0
    _ = p1
    _ = v11
    return 0
}
func genFn1() uint16 {
    v12 := []int{0, -635}
    if len(v12) > 0 {
        v12[0] = -2147483648
    }
    _ = v12
    return 543
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 127
    var v3 uint8 = v1
    v4 := []int{-1}
    var v5 string = "hello"
    if false {
        v3 = v1
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint32 = 2147483648
            if true {
                switch "\n\t" {
                case "":
                    switch (^(fi13)) {
                    case -375:
                        v3 += v3
                    case 159:
                        var v15 float64 = ((3.14159 - 1e10) - 1e10)
                        _ = v15
                    case 1:
                        if true {
                            var v16 string = "foo bar"
                            v17 := GenS1{F0: 1, F1: (^(v14))}
                            var v18 uint = 0
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                case "foo bar":
                    v1 = 2
                }
                var v19 bool = true
                var v20 uint = 510
                _ = v19
                _ = v20
            } else {
                fmt.Println(fi13, v1)
                v1 = (v3 - v1)
                _ = len(v4)
            }
            var v21 uint64 = 224
            _ = fi13
            _ = v14
            _ = v21
        }
        var v22 uint = 18446744073709551615
        _ = v22
    }
    v4 = append(v4, 2147483647)
    fmt.Println(v3, v1, v5)
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 string = "\n\t"
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 65535
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            if true {
                if true {
                    var v28 int = fi26
                    v3 = (^(v3))
                    _ = v28
                } else {
                    var v29 int32 = 2147483647
                    v30 := GenS2{F0: 0, F1: -1, F2: true}
                    v5 = ((v5 + "\n\t") + v5)
                    _ = v29
                    _ = v30
                }
                var v31 uint = 4294967295
                v1 -= (v3 << (((v31 ^ 656)) & 63))
                _ = v31
            }
            switch v5 {
            case "a":
                switch fi26 {
                case 127:
                    var v32 uint16 = genFn1()
                    _ = v32
                case -251:
                    var v33 uint = 42
                    _ = v33
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case "hello":
                v25 = v25
            case "\n\t":
                v34 := GenS3{F0: false}
                var v35 int32 = 736
                v36 := GenS4{F0: 2, F1: v5, F2: 42}
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi27
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            v3 = v3
            _ = fi37
        }
        var v38 uint64 = 31
        _ = fi26
        _ = v38
    }
    v25 += 0
    if true {
        var v39 int64 = (v2 | v2)
        _ = v39
    }
    switch 127 {
    case 32768:
        var v40 string = "\n\t"
        _ = v40
    case 4294967295:
        for fi41 := 0; fi41 < 12; fi41++ {
            v2 += v2
            _ = fi41
        }
    case 9223372036854775807:
        v42 := GenS2{F0: 1, F1: 2147483647, F2: ((((1.5 + (0.1 * -2.25)) + float32(((float64(((0.0 / (float32(float64((3.14159 - (float32(3.14159) + (1.5 + 3.14159))))) * -2.25)) * (0.0 * float32((1e10 * 1.5))))) - 1e10) + (3.14159 + 0.0)))) < 0.0) || false)}
        _ = v42
    }
    v1 = (v1 >> ((466) & 63))
    if (((!true) && (4294967295 >= 65535)) && (!(460 == 25))) {
        if true {
            var v43 int32 = 953
            _ = v43
        } else {
            if false {
                for fi44 := 0; fi44 < 19; fi44++ {
                    if (true && (!((!((false || (18446744073709551615 == 987)) && true)) && false))) {
                        v45 := GenS2{F0: 0, F1: 2147483647, F2: ("" < ("\n\t" + v5))}
                        _ = v45
                    } else {
                        switch (fi44 % (((fi44 / (((-(fi44))) | 1))) | 1)) {
                        case 127:
                            for fi46 := 0; fi46 < 11; fi46++ {
                                var v47 uint16 = 255
                                v5 = v5
                                switch fi46 {
                                case 918:
                                    if len(v4) > 0 {
                                        v4[0] = fi46
                                    }
                                case 743:
                                    fmt.Println(v5, v25)
                                    if true {
                                        var v48 float64 = 3.14159
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v49 bool = (!(v48 >= v48))
                                        _ = v48
                                        _ = v49
                                    }
                                case 32768:
                                    if false {
                                        if (1.5 > 0.1) {
                                            switch fi44 {
                                            case -33:
                                                var v50 float64 = 1.5
                                                _ = v50
                                            case 1:
                                                if ((v5 + v5) <= ((v5 + "hello") + ((v5 + "a") + ("" + v5)))) {
                                                    _ = len(v4)
                                                    if (!(("" + v5) >= v5)) {
                                                        var v51 uint = 255
                                                        var v52 float32 = 0.0
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                            default:
                                                var v53 int64 = (^(v2))
                                                _ = v53
                                            }
                                            fmt.Println(v2, v25)
                                            fmt.Println(v25, v47)
                                        } else {
                                            var v54 float64 = 1.5
                                            v2 = 0
                                            _ = v54
                                        }
                                    } else {
                                        v3 -= v3
                                        switch (v5 + v5) {
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                default:
                                    fmt.Println(v47, fi46, v5)
                                }
                                v25 = (v25 >> ((uint(v25)) & 63))
                                v25 = genFn0(255, v47)
                                _ = fi46
                                _ = v47
                            }
                        default:
                            fmt.Println(v5, v3)
                        }
                    }
                    _ = fi44
                }
                switch "foo bar" {
                case "a":
                    var v55 bool = (752 != 1)
                    _ = v55
                case "\n\t":
                    if (("" + ((v5 + (((v5 + "\n\t") + v5) + ("hello" + v5))) + (("foo bar" + (v5 + ("hello" + v5))) + v5))) == v5) {
                        if (!false) {
                            if (278 < 18446744073709551615) {
                                for fi56 := 0; fi56 < 2; fi56++ {
                                    var v57 int32 = 42
                                    _ = fi56
                                    _ = v57
                                }
                            } else {
                                for fi58 := 0; fi58 < 14; fi58++ {
                                    v59 := GenS2{F0: 1, F1: 0, F2: false}
                                    for fi60 := 0; fi60 < 18; fi60++ {
                                        v61 := GenS3{F0: true}
                                        v3 -= (^(v3))
                                        _ = fi60
                                        _ = v61
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                            }
                        } else {
                            switch 2147483647 {
                            case 1:
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    v63 := GenS4{F0: 776, F1: v5, F2: v2}
                                    v1 = 1
                                    _ = fi62
                                    _ = v63
                                }
                            case 32768:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            switch 0 {
                            case -1:
                                switch -1 {
                                case -1:
                                    var v64 int = 517
                                    _ = v64
                                }
                            case 412:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 4294967295:
                                _ = len(v4)
                            }
                        }
                        fmt.Println(v1, v25, v1)
                    } else {
                        if len(v4) > 0 {
                            v4[0] = -9223372036854775808
                        }
                        var v65 int = -115
                        _ = v65
                    }
                }
                switch 640 {
                case 128:
                    v2 -= -9223372036854775808
                default:
                    v1 = uint8(v25)
                }
            }
            var v66 int64 = (v2 % (((v2 >> ((42) & 63))) | 1))
            for fi67 := 0; fi67 < 2; fi67++ {
                if (!(-499 >= 1)) {
                    if (true && true) {
                        v68 := GenS5{F0: 0, F1: 18446744073709551615}
                        var v69 uint64 = 1
                        _ = v68
                        _ = v69
                    } else {
                        _ = len(v4)
                    }
                    var v70 uint32 = v25
                    _ = v70
                } else {
                    if true {
                        v71 := GenS6{F0: 42, F1: -470, F2: -668}
                        _ = v71
                    } else {
                        v72 := GenS7{F0: (v25 ^ v25), F1: v1, F2: 635}
                        _ = v72
                    }
                }
                var v73 int8 = -128
                var v74 uint = 255
                _ = fi67
                _ = v73
                _ = v74
            }
            _ = v66
        }
        v4 = append(v4, -330)
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        v25 = (v25 * (v25 + (v25 >> ((2) & 63))))
        v3 -= 255
        if false {
            var v76 int = fi75
            _ = len(v4)
            v2 += 1
            _ = v76
        }
        _ = fi75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
