package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint; F1 uint8; F2 int }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 string; F1 uint; F2 bool }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int; F1 int64 }
func genFn0(p0 int) uint8 {
    v5 := []int{p0, (p0 - (p0 ^ -1)), (p0 | p0), p0}
    var v6 bool = false
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16) bool {
    v7 := GenS1{F0: 2147483647}
    v8 := []int{42, 1}
    if true {
        if (false || ("\n\t" <= "")) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            switch 4294967295 {
            case 32767:
                var v9 int16 = (p0 << ((0) & 63))
                _ = v9
            case 2147483647:
                var v10 int64 = -1
                _ = v10
            case -1:
                var v11 int64 = 42
                _ = v11
            default:
                var v12 uint32 = 42
                _ = v12
            }
            p0 -= p0
        }
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            p0 += p0
            for fi14 := 0; fi14 < 10; fi14++ {
                p0 = p0
                var v15 int32 = -1
                v16 := GenS2{F0: (((v15 | (v15 | -1)) <= (v15 >> ((894) & 63))) && false), F1: (p0 | p0), F2: 387}
                _ = fi14
                _ = v15
                _ = v16
            }
            switch ((("\n\t" + "a") + "\n\t") + "hello") {
            case "":
                switch "hello" {
                case "\n\t":
                    v17 := GenS2{F0: (!(-2.25 < 0.1)), F1: p0, F2: 1}
                    _ = v17
                }
            case "a":
                if true {
                    v18 := GenS1{F0: fi13}
                    _ = v18
                } else {
                    if (2147483648 == 287) {
                        var v19 string = "foo bar"
                        switch (fi13 - 127) {
                        case 4294967295:
                            switch fi13 {
                            case -9223372036854775808:
                                v20 := GenS1{F0: (fi13 / (((fi13 + (fi13 | (fi13 / (((-(fi13))) | 1))))) | 1))}
                                _ = v20
                            case -277:
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    if false {
                                        var v22 int32 = 42
                                        var v23 int = 9223372036854775807
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = len(v8)
                                    p0 += (p0 * p0)
                                    switch v19 {
                                    case "hello":
                                        if false {
                                            var v24 int32 = -554
                                            var v25 uint32 = 4294967295
                                            switch fi21 {
                                            case -53:
                                                v26 := GenS2{F0: (true && ((v24 == v24) && true)), F1: p0, F2: 0}
                                                _ = v26
                                            default:
                                                var v27 int8 = -21
                                                _ = v27
                                            }
                                            if len(v8) > 0 {
                                                v8[0] = (fi21 ^ 127)
                                            }
                                            if (int8(fi21) < 42) {
                                                v25 -= v25
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                                v28 := GenS2{F0: (true && (42 != -2147483648)), F1: (p0 & (p0 & p0)), F2: 129}
                                                v29 := GenS3{F0: 112, F1: 720}
                                                _ = v28
                                                _ = v29
                                            } else {
                                                p0 -= 257
                                                if (!true) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v25 -= (v25 >> ((494) & 63))
                                                    default:
                                                        var v30 uint16 = 255
                                                        _ = v30
                                                    }
                                                    v25 = 0
                                                } else {
                                                    v25 -= (v25 << ((4294967295) & 63))
                                                    var v31 uint8 = 42
                                                    v32 := GenS4{F0: v19}
                                                    _ = v31
                                                    _ = v32
                                                }
                                            }
                                            v24 = (^(v24))
                                            v33 := GenS5{F0: 2}
                                            v8 = append(v8, fi13)
                                            var v34 int = (fi21 * 492)
                                            var v35 bool = ((((true && (true && (true || (false && true)))) && (false && (!false))) && (42 == 9223372036854775807)) || false)
                                            _ = v24
                                            _ = v25
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                    case "":
                                        v19 = "hello"
                                    case "\n\t":
                                        p0 -= (p0 * (p0 / (((p0 & (p0 << ((18446744073709551615) & 63)))) | 1)))
                                    }
                                    p0 = -774
                                    _ = fi21
                                }
                            case 32768:
                                v36 := GenS5{F0: 296}
                                _ = v36
                            default:
                                var v37 float64 = 0.0
                                _ = v37
                            }
                        case 32767:
                            fmt.Println(p0, v19)
                        case 2147483647:
                            v38 := GenS6{F0: -1, F1: 0}
                            _ = v38
                        }
                        _ = v19
                    } else {
                        if (false && ("foo bar" != "\n\t")) {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                v8 = append(v8, -9223372036854775808)
                                _ = fi39
                            }
                            _ = len(v8)
                            var v40 string = ""
                            _ = v40
                        }
                        var v41 float64 = 3.14159
                        _ = v41
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
            case "foo bar":
                var v42 int32 = -203
                _ = v42
            default:
                p0 += (p0 * p0)
            }
            _ = fi13
        }
    }
    v43 := GenS5{F0: 256}
    _ = p0
    _ = v43
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 string = "hello"
    var v2 float32 = 0.1
    v3 := []uint32{344, 4294967295, 311}
    var v4 uint = 774
    v44 := GenS2{F0: (42 < 65535), F1: -548, F2: 874}
    fmt.Println(v1, v2, v1)
    if false {
        if (!true) {
            var v45 uint64 = 362
            _ = v45
        }
    }
    var v46 float64 = 0.0
    for fi47 := 0; fi47 < 13; fi47++ {
        v46 = float64(v2)
        if false {
            for fi48 := 0; fi48 < 20; fi48++ {
                v1 = "\n\t"
                v2 += 1.5
                var v49 string = ""
                _ = fi48
                _ = v49
            }
            fmt.Println(v1, v2, v4)
            v50 := GenS7{F0: v4, F1: 42, F2: 138}
            _ = v50
        }
        _ = fi47
    }
    var v51 int8 = 42
    switch v1 {
    case "hello":
        switch ("" + "a") {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                var v52 uint8 = genFn0(4294967295)
                _ = v52
            case "foo bar":
                v53 := GenS5{F0: 1}
                _ = v53
            }
        case "":
            switch 42 {
            case 42:
                v46 = float64((((v2 + v2) * v2) + float32(v46)))
            case 128:
                for fi54 := 0; fi54 < 11; fi54++ {
                    v4 -= 18446744073709551615
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    switch v1 {
                    case "":
                        for fi55 := 0; fi55 < 10; fi55++ {
                            var v56 uint64 = 967
                            var v57 uint64 = v56
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    case "\n\t":
                        for fi58 := 0; fi58 < 18; fi58++ {
                            if false {
                                v2 -= (v2 - 0.1)
                            } else {
                                v3 = append(v3, 4294967295)
                                fmt.Println(fi54, v51)
                                if (v1 >= v1) {
                                    fmt.Println(v4, v4, v4)
                                } else {
                                    v59 := GenS8{F0: 1, F1: 9223372036854775807}
                                    _ = v59
                                }
                            }
                            var v60 int32 = 0
                            _ = len(v3)
                            _ = fi58
                            _ = v60
                        }
                    default:
                        v61 := GenS9{F0: v1, F1: (v4 & 767), F2: (int64(fi54) > -1)}
                        _ = v61
                    }
                    _ = fi54
                }
            default:
                for fi62 := 0; fi62 < 19; fi62++ {
                    for fi63 := 0; fi63 < 15; fi63++ {
                        fmt.Println(v2, v51)
                        _ = fi63
                    }
                    _ = fi62
                }
            }
        case "a":
            v3 = append(v3, 42)
        }
    default:
        if (v4 == v4) {
            var v64 string = ""
            _ = v64
        } else {
            v3 = append(v3, 0)
            _ = len(v3)
        }
    }
    v65 := GenS10{F0: 155, F1: 32767}
    v66 := GenS11{F0: 9223372036854775807, F1: 428}
    for fi67 := 0; fi67 < 8; fi67++ {
        if true {
            if true {
                var v68 int = fi67
                _ = v68
            } else {
                var v69 uint32 = 375
                _ = v69
            }
        } else {
            v70 := GenS7{F0: (v4 + v4), F1: 127, F2: (fi67 - fi67)}
            _ = v70
        }
        fmt.Println(v1, v46)
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v51
    _ = v65
    _ = v66
}
