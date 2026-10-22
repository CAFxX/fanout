package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []uint32{1, 65535}
    v2 := make(map[int]int)
    var v3 int32 = 2147483647
    var v4 int64 = 0
    var v5 int8 = 100
    v6 := make(map[int]string)
    if (-757 < 0) {
        var v7 int32 = v3
        for fi8 := 0; fi8 < 1; fi8++ {
            if false {
                var v9 int64 = -9223372036854775808
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v2, (fi8 ^ (fi8 + (^(fi8)))))
                _ = v9
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v10 float64 = -2.25
                v5 += (^(v5))
                _ = v10
            }
            _ = fi8
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 = append(v1, 42)
        _ = fi11
    }
    if ("foo bar" <= (("hello" + "") + "a")) {
        var v12 uint = 42
        _ = v12
    } else {
        if false {
            v13 := v6[128]
            switch v13 {
            case "a":
                if (false || false) {
                    fmt.Println(v3, v5)
                } else {
                    v14 := GenS1{F0: 4294967295}
                    v3 += int32(v4)
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v6[(fi15 & 9223372036854775807)] = ""
                        v5 = 127
                        _ = fi15
                    }
                    _ = v14
                }
            case "hello":
                v5 = v5
            default:
                if (v13 <= v13) {
                    if false {
                        if (!(556 >= -1)) {
                            _ = len(v2)
                            for fi16 := 0; fi16 < 4; fi16++ {
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    v18 := GenS2{F0: 450, F1: 1}
                                    if (!(false && (407 != 4294967295))) {
                                        var v19 int32 = (v3 % ((v3) | 1))
                                        v20 := v6[1]
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                                if (v13 < v13) {
                                    var v21 int = (^(fi16))
                                    switch fi16 {
                                    case 2147483647:
                                        var v22 int64 = v4
                                        _ = v22
                                    case 9223372036854775807:
                                        for fi23 := 0; fi23 < 1; fi23++ {
                                            var v24 uint = 420
                                            v2[fi16] = fi16
                                            _ = fi23
                                            _ = v24
                                        }
                                    default:
                                        v25 := v6[(v21 | 0)]
                                        _ = v25
                                    }
                                    var v26 float64 = -2.25
                                    _ = v21
                                    _ = v26
                                }
                                _ = fi16
                            }
                        } else {
                            v4 += (v4 & v4)
                            v3 += (v3 - (v3 + int32(v4)))
                            var v27 float32 = 0.0
                            _ = v27
                        }
                    }
                } else {
                    v28 := GenS1{F0: 2}
                    v29 := GenS3{F0: 1, F1: (v3 - v3), F2: (^(v5))}
                    fmt.Println(v13, v13, v5)
                    _ = v28
                    _ = v29
                }
            }
            _ = v13
        }
        v30 := GenS1{F0: 1}
        v3 = v3
        _ = v30
    }
    var v31 float32 = 1e10
    v5 = (^(v5))
    var v32 float32 = v31
    if ((true || true) || ((true && true) && (462 < 255))) {
        v33 := GenS4{F0: 42}
        _ = v33
    } else {
        v5 = (v5 + (v5 + 88))
        v32 = (v31 / v32)
    }
    switch "a" {
    case "a":
        if true {
            for fi34 := 0; fi34 < 17; fi34++ {
                fmt.Println(v31, fi34, v5)
                for fi35 := 0; fi35 < 4; fi35++ {
                    switch 984 {
                    case 32767:
                        v31 += (0.1 - v31)
                    case 897:
                        v4 -= 1
                    case 128:
                        var v36 int16 = -859
                        _ = v36
                    }
                    if true {
                        v31 -= v31
                    }
                    _ = fi35
                }
                _ = fi34
            }
            fmt.Println(v32, v32, v31)
            v32 = v32
        }
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        v38 := GenS2{F0: 990, F1: (fi37 + 170)}
        switch fi37 {
        case 42:
            var v39 uint8 = 128
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    v3 = v3
    v2[890] = -9223372036854775808
    v6[1] = "foo bar"
    if (!true) {
        if false {
            for fi40 := 0; fi40 < 1; fi40++ {
                var v41 float32 = v31
                switch fi40 {
                case 32768:
                    var v42 bool = (2 == 4294967295)
                    _ = v42
                case 32767:
                    var v43 string = ""
                    _ = v43
                case -9223372036854775808:
                    switch 42 {
                    case 128:
                        v4 += v4
                    }
                default:
                    for fi44 := 0; fi44 < 8; fi44++ {
                        v2[fi40] = fi44
                        if false {
                            v45 := GenS2{F0: 32, F1: 32768}
                            var v46 int = fi44
                            v4 -= v4
                            _ = v45
                            _ = v46
                        } else {
                            var v47 float32 = 0.1
                            v48 := GenS2{F0: 116, F1: fi44}
                            _ = v47
                            _ = v48
                        }
                        _ = fi44
                    }
                }
                if true {
                    v41 = float32(0.0)
                    var v49 int32 = -1
                    _ = len(v1)
                    _ = v49
                }
                _ = fi40
                _ = v41
            }
            fmt.Println(v4, v31)
        }
        switch 803 {
        case 2147483647:
            v50 := GenS1{F0: 726}
            _ = v50
        case 38:
            v1 = append(v1, 2147483648)
        }
    } else {
        v1 = append(v1, 2147483648)
        if (v5 > (v5 * (v5 % (((v5 / ((v5) | 1))) | 1)))) {
            if (255 < 399) {
                v31 -= v32
                v4 = 650
            }
        }
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 float32 = (3.14159 / float32(1.5))
            var v53 uint = 1
            fmt.Println(v52, v32, v31)
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    switch "hello" {
    case "a":
        for fi54 := 0; fi54 < 7; fi54++ {
            v5 += (v5 >> ((42) & 63))
            _ = fi54
        }
    default:
        switch int(v4) {
        case -9223372036854775808:
            var v55 uint8 = 1
            _ = v55
        case 128:
            if (369 != 899) {
                v5 = v5
            }
        default:
            v4 = 512
        }
    }
    switch -215 {
    case 32767:
        var v56 int64 = (v4 / ((-9223372036854775808) | 1))
        _ = v56
    }
    v32 -= v32
    fmt.Println(v32, v31, v32)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
