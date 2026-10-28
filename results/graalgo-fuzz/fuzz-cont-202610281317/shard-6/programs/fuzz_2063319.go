package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 int32; F2 int32 }

func main() {
    var v1 float64 = (1e10 - float64(float32((1.5 * 0.1))))
    v2 := []uint32{}
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 924, F1: (!true)}
        var v7 float32 = 0.1
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 += v1
    var v8 float64 = v1
    switch "\n\t" {
    case "foo bar":
        v9 := GenS2{F0: 4294967295, F1: 0, F2: 2}
        _ = v9
    case "a":
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                v2 = append(v2, 2147483648)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v11 int16 = 505
                _ = fi10
                _ = v11
            }
        } else {
            v1 += float64(1e10)
            if (1.5 < ((float32(((v1 + v8) + v8)) - 1e10) + float32(v1))) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v12 := GenS2{F0: 75, F1: 1, F2: 42}
                    _ = v12
                case "\n\t":
                    v8 = v8
                }
            } else {
                if (!(0 < 65535)) {
                    v13 := GenS1{F0: 65535, F1: false}
                    v2 = append(v2, 0)
                    _ = v13
                } else {
                    switch "foo bar" {
                    case "":
                        if false {
                            var v14 uint64 = 4294967295
                            _ = v14
                        } else {
                            var v15 uint8 = 255
                            var v16 int8 = 42
                            v17 := GenS3{F0: (^(v16)), F1: 42}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        fmt.Println(v1)
                    case "a":
                        if false {
                            var v18 int16 = -32768
                            if (42 == 262) {
                                v19 := GenS2{F0: 1, F1: 42, F2: 76}
                                var v20 float64 = (v1 / (((v8 / v1) * v8) / 0.1))
                                _ = v19
                                _ = v20
                            }
                            _ = v18
                        } else {
                            v1 = (0.0 + -2.25)
                        }
                    default:
                        switch 42 {
                        case 1:
                            for fi21 := 0; fi21 < 8; fi21++ {
                                switch "hello" {
                                case "hello":
                                    v8 += 0.0
                                case "\n\t":
                                    var v22 uint64 = 2
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    switch "a" {
                                    case "a":
                                        switch fi21 {
                                        case 602:
                                            v24 := GenS4{F0: 65535, F1: 1, F2: 224}
                                            _ = v24
                                        case 32768:
                                            v1 += v1
                                        case 128:
                                            v25 := GenS5{F0: -19}
                                            _ = v25
                                        }
                                    }
                                    _ = fi23
                                }
                                _ = fi21
                            }
                        case 0:
                            v1 = v8
                        default:
                            var v26 int32 = -2147483648
                            _ = v26
                        }
                    }
                    v8 -= ((v8 * (float64((1.5 - (-2.25 + 3.14159))) / v1)) * float64(0.0))
                    if false {
                        v27 := GenS4{F0: 32767, F1: 255, F2: 1}
                        var v28 float64 = v8
                        var v29 int64 = 0
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                if false {
                    var v30 uint8 = 2
                    switch 42 {
                    case -1:
                        v31 := GenS4{F0: 0, F1: 578, F2: v30}
                        _ = v31
                    case -213:
                        for fi32 := 0; fi32 < 12; fi32++ {
                            if true {
                                var v33 float32 = 0.0
                                if (((-2.25 / ((0.0 * ((float32(((v1 * float64(((1e10 * v33) + v33))) - 0.0)) - (1e10 - 1.5)) * (v33 * v33))) + 1.5)) + v33) < (v33 / (0.0 / (v33 / 1.5)))) {
                                    var v34 int = (^(fi32))
                                    var v35 int64 = -9223372036854775808
                                    _ = v34
                                    _ = v35
                                }
                                _ = v33
                            }
                            v36 := GenS6{F0: (v30 & (v30 ^ (v30 - v30))), F1: 547}
                            _ = fi32
                            _ = v36
                        }
                    case 313:
                        v8 += v1
                    default:
                        for fi37 := 0; fi37 < 19; fi37++ {
                            var v38 uint32 = 1
                            _ = fi37
                            _ = v38
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v30
                }
            }
            for fi39 := 0; fi39 < 16; fi39++ {
                fmt.Println(fi39)
                _ = fi39
            }
        }
    case "hello":
        var v40 int32 = -2147483648
        _ = v40
    default:
        var v41 bool = false
        _ = v41
    }
    _ = len(v3)
    if true {
        v42 := GenS4{F0: 255, F1: 255, F2: 127}
        switch 1 {
        case 1:
            v43 := GenS5{F0: 42}
            _ = v43
        }
        _ = v42
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if true {
            v2 = append(v2, 187)
        }
        v1 = 1e10
        _ = fi44
    }
    switch "" {
    case "hello":
        if (4294967295 != 18446744073709551615) {
            for fi45 := 0; fi45 < 6; fi45++ {
                v1 = v1
                fmt.Println(v8, v1)
                _ = fi45
            }
            var v46 uint8 = 251
            for fi47 := 0; fi47 < 17; fi47++ {
                v48 := GenS7{F0: -1}
                switch "a" {
                case "\n\t":
                    fmt.Println(v1, fi47)
                default:
                    var v49 int16 = 0
                    _ = v49
                }
                v2 = append(v2, 1)
                _ = fi47
                _ = v48
            }
            _ = v46
        } else {
            var v50 uint16 = 501
            for fi51 := 0; fi51 < 16; fi51++ {
                var v52 int32 = 42
                _ = fi51
                _ = v52
            }
            switch "a" {
            case "\n\t":
                fmt.Println(v50, v50, v50)
            default:
                var v53 float32 = 1e10
                _ = v53
            }
            _ = v50
        }
    }
    var v54 int16 = -477
    if true {
        v55 := GenS5{F0: v54}
        _ = len(v2)
        if (505 < 65535) {
            _ = len(v2)
        }
        _ = v55
    }
    v54 = (v54 + (v54 % (((-(v54))) | 1)))
    for fi56 := 0; fi56 < 15; fi56++ {
        v57 := GenS8{F0: 4294967295}
        fmt.Println(v1)
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 16; fi58++ {
        for fi59 := 0; fi59 < 13; fi59++ {
            if (4294967295 >= 0) {
                var v60 float64 = 0.0
                _ = v60
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi59
        }
        _ = fi58
    }
    var v61 uint16 = 2
    var v62 string = (("\n\t" + "") + ("a" + "foo bar"))
    if false {
        if (!(1 > 1)) {
            var v63 string = v62
            _ = len(v2)
            _ = v63
        } else {
            v64 := GenS1{F0: 966, F1: (877 != 1)}
            v1 -= 1e10
            _ = v64
        }
        var v65 int64 = 0
        _ = v65
    }
    if true {
        v66 := GenS6{F0: 0, F1: 42}
        v67 := GenS9{F0: 42, F1: 42, F2: -1}
        var v68 int16 = v54
        _ = v66
        _ = v67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v61
    _ = v62
    _ = v8
}
