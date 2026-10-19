package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int }
type GenS6 struct { F0 int; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0() int {
    var v6 int16 = 42
    v7 := []int{2147483647, 735, 127}
    v8 := GenS1{F0: 348, F1: 1, F2: 18446744073709551615}
    var v9 uint32 = 578
    var v10 int64 = -1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -834
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 57
    var v3 float32 = -2.25
    var v4 uint8 = v2
    var v5 uint64 = 764
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(fi11)
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := []int{fi12}
            v14 := GenS1{F0: 1, F1: v5, F2: uint(v5)}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := []int{(fi11 % (((fi11 & (fi11 | -1))) | 1)), 127, fi11, (fi11 % ((32768) | 1))}
        _ = fi11
        _ = v15
    }
    if true {
        fmt.Println(v1, v3)
    }
    var v16 float32 = float32(0.1)
    if (65535 != 997) {
        if true {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int64 = (v1 * v1)
                if (!(!false)) {
                    v19 := GenS2{F0: 973, F1: -79, F2: v2}
                    v20 := []int{}
                    v2 -= v2
                    _ = v19
                    _ = v20
                }
                v1 += (v1 << ((255) & 63))
                _ = fi17
                _ = v18
            }
            if ((true || (880 == 32767)) || (!true)) {
                for fi21 := 0; fi21 < 12; fi21++ {
                    if false {
                        v2 += (v2 * (v2 + (^(v2))))
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                v22 := GenS2{F0: -1, F1: 42, F2: v4}
                                _ = v22
                            case "":
                                fmt.Println(v4)
                            case "foo bar":
                                if false {
                                    v5 += (v5 % (((v5 | (v5 / (((v5 | 1)) | 1)))) | 1))
                                    for fi23 := 0; fi23 < 1; fi23++ {
                                        v2 -= (v4 + (v4 * (v2 - 95)))
                                        var v24 int = 42
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                            }
                            fmt.Println(v4, v1)
                        } else {
                            fmt.Println(v2)
                            v25 := GenS1{F0: -1, F1: (v5 * 326), F2: 4294967295}
                            _ = v25
                        }
                    } else {
                        if (false || false) {
                            var v26 float32 = (v3 - v3)
                            v26 = (1.5 + v16)
                            v5 -= (v5 ^ v5)
                            _ = v26
                        } else {
                            switch int(v5) {
                            case -9223372036854775808:
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    var v28 uint32 = 0
                                    var v29 uint32 = v28
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                            case -325:
                                var v30 string = "a"
                                _ = v30
                            default:
                                if true {
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        if false {
                                            v32 := []int{}
                                            switch "" {
                                            case "":
                                                if len(v32) > 0 {
                                                    v32[0] = 1
                                                }
                                            case "a":
                                                if len(v32) > 0 {
                                                    v32[0] = fi31
                                                }
                                            case "hello":
                                                var v33 uint32 = 2147483647
                                                _ = v33
                                            }
                                            switch "" {
                                            case "hello":
                                                var v34 uint32 = 153
                                                _ = v34
                                            case "\n\t":
                                                var v35 int16 = 39
                                                _ = v35
                                            }
                                            _ = v32
                                        }
                                        _ = fi31
                                    }
                                }
                            }
                            var v36 uint = 2
                            _ = v36
                        }
                        fmt.Println(v1)
                    }
                    switch fi21 {
                    case 326:
                        fmt.Println(v1, v4, v5)
                    case -383:
                        v37 := GenS2{F0: -563, F1: 100, F2: (v2 + (v4 | (v4 << ((2) & 63))))}
                        _ = v37
                    case 4294967295:
                        v38 := GenS3{F0: v5, F1: (^(v4))}
                        _ = v38
                    default:
                        fmt.Println(fi21, v1, v2)
                    }
                    switch "\n\t" {
                    case "":
                        var v39 uint8 = (v2 >> ((1) & 63))
                        _ = v39
                    case "hello":
                        if false {
                            switch "a" {
                            case "hello":
                                var v40 uint64 = 925
                                _ = v40
                            }
                            fmt.Println(v2, v2, fi21)
                        } else {
                            v4 = (v4 / ((v2) | 1))
                        }
                    default:
                        v41 := []int{0, (fi21 | (fi21 * (fi21 + fi21))), fi21, (fi21 << ((65535) & 63))}
                        _ = v41
                    }
                    _ = fi21
                }
                if (!(!true)) {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        var v43 float64 = 0.0
                        var v44 bool = (!((false && false) && (false || ((("\n\t" + "foo bar") + "") > "hello"))))
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    v45 := GenS1{F0: 42, F1: v5, F2: 918}
                    v46 := []int{genFn0(), 0}
                    _ = v45
                    _ = v46
                }
            } else {
                for fi47 := 0; fi47 < 8; fi47++ {
                    v3 = v3
                    _ = fi47
                }
            }
            v1 -= v1
        } else {
            switch -1 {
            case 2147483647:
                var v48 float64 = ((0.0 + (0.1 / 1e10)) + -2.25)
                _ = v48
            }
            v49 := []int{675}
            if ((-1 > -984) && false) {
                if true {
                    v50 := GenS4{F0: uint16(v5), F1: -128, F2: -128}
                    fmt.Println(v4)
                    v49 = append(v49, 32767)
                    _ = v50
                } else {
                    var v51 int64 = v1
                    if len(v49) > 0 {
                        v49[0] = 32768
                    }
                    _ = v51
                }
            }
            _ = v49
        }
        if true {
            v2 = v2
            for fi52 := 0; fi52 < 9; fi52++ {
                v2 += v2
                if true {
                    var v53 bool = true
                    _ = v53
                }
                _ = fi52
            }
        }
        v54 := []int{1, 32768, 9223372036854775807}
        _ = v54
    } else {
        v55 := []int{830, 42, genFn0(), 32767}
        for fi56 := 0; fi56 < 20; fi56++ {
            var v57 int16 = 0
            _ = len(v55)
            _ = fi56
            _ = v57
        }
        var v58 uint16 = 0
        _ = v55
        _ = v58
    }
    if ("hello" != "") {
        var v59 uint = 2
        _ = v59
    }
    v60 := GenS1{F0: 0, F1: (^(v5)), F2: 4294967295}
    switch -543 {
    case 480:
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v3, v1, v3)
        case "foo bar":
            v61 := GenS5{F0: v5, F1: -32768, F2: 127}
            _ = v61
        case "a":
            fmt.Println(v2)
        default:
            if ((v1 << ((325) & 63)) <= v1) {
                v62 := []int{581, 2147483647, 42}
                if (!false) {
                    switch -281 {
                    case 1:
                        v3 += (v16 - (3.14159 + v16))
                    default:
                        for fi63 := 0; fi63 < 16; fi63++ {
                            if len(v62) > 0 {
                                v62[0] = fi63
                            }
                            if len(v62) > 0 {
                                v62[0] = 4294967295
                            }
                            v64 := GenS6{F0: fi63, F1: 65535, F2: 2}
                            _ = fi63
                            _ = v64
                        }
                    }
                } else {
                    v65 := GenS7{F0: 1, F1: (v2 - (v4 & (v4 | v2))), F2: (^(v4))}
                    _ = v65
                }
                var v66 int64 = (^(v1))
                _ = v62
                _ = v66
            }
        }
    }
    v67 := []int{-340, 32768, 42, -9223372036854775808}
    fmt.Println(v2, v5)
    var v68 int64 = (-(v1))
    for fi69 := 0; fi69 < 18; fi69++ {
        for fi70 := 0; fi70 < 6; fi70++ {
            var v71 int32 = 0
            _ = fi70
            _ = v71
        }
        var v72 uint = 531
        var v73 uint16 = 91
        _ = fi69
        _ = v72
        _ = v73
    }
    fmt.Println(v16, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v67
    _ = v68
}
