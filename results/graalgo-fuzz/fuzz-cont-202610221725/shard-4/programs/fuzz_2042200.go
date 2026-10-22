package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int16 = -1
    var v2 float32 = ((float32((0.0 / (0.1 / 1e10))) + (3.14159 * (1e10 * 1e10))) * (0.1 * 1e10))
    v3 := []uint32{}
    var v4 int64 = 500
    var v5 int = 747
    v6 := []uint32{2, 42, 2147483648, 968}
    var v7 uint8 = 0
    fmt.Println(v1, v1)
    for fi8 := 0; fi8 < 10; fi8++ {
        _ = len(v3)
        v4 -= int64(v7)
        _ = fi8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            for fi10 := 0; fi10 < 15; fi10++ {
                v6 = append(v6, 1)
                _ = fi10
            }
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            if (0 > 42) {
                var v12 uint32 = 42
                _ = v12
            } else {
                if (true && false) {
                    switch v5 {
                    case 1:
                        var v13 bool = false
                        _ = v13
                    case 32768:
                        if false {
                            v5 += (fi11 / (((^(fi9))) | 1))
                            var v14 int16 = 204
                            v1 -= v1
                            if (((-2.25 > (float64((v2 * v2)) / 1.5)) && true) && true) {
                                fmt.Println(fi9, fi11, fi9)
                                var v15 float64 = -2.25
                                v1 = (v14 + v14)
                                _ = v15
                            } else {
                                var v16 bool = false
                                fmt.Println(v16, v4)
                                _ = v16
                            }
                            switch (v5 % ((282) | 1)) {
                            case 4294967295:
                                switch 147 {
                                case 9223372036854775807:
                                    v4 = 9223372036854775807
                                case -619:
                                    v14 = 81
                                default:
                                    if ((122 != 42) || ((v7 ^ v7) > (v7 >> ((uint(v7)) & 63)))) {
                                        var v17 string = ("" + "foo bar")
                                        fmt.Println(v7)
                                        var v18 float64 = float64(v2)
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            case 32768:
                                switch v5 {
                                case 738:
                                    if true {
                                        v2 -= -2.25
                                        switch "hello" {
                                        case "hello":
                                            var v19 uint64 = 905
                                            _ = v19
                                        case "\n\t":
                                            var v20 int8 = -1
                                            _ = v20
                                        case "":
                                            switch 4294967295 {
                                            case 127:
                                                v1 += v14
                                            default:
                                                fmt.Println(fi11, v2, v2)
                                            }
                                        default:
                                            var v21 uint64 = 4294967295
                                            _ = v21
                                        }
                                    }
                                case 0:
                                    v2 = -2.25
                                default:
                                    v22 := GenS1{F0: 2, F1: 989, F2: -2147483648}
                                    _ = v22
                                }
                                var v23 bool = true
                                _ = v23
                            default:
                                _ = len(v3)
                            }
                            _ = v14
                        } else {
                            var v24 float64 = 3.14159
                            _ = v24
                        }
                    }
                }
            }
            var v25 uint64 = 883
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v26 int64 = v4
            _ = fi11
            _ = v25
            _ = v26
        }
        v1 = -184
        _ = fi9
    }
    var v27 int32 = 0
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS1{F0: 0, F1: 4294967295, F2: v27}
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 126, F1: 2, F2: 2147483647}
    var v31 float64 = -2.25
    var v32 uint8 = (^(v7))
    fmt.Println(v4, v1)
    fmt.Println(v4)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 int64 = -516
        switch "\n\t" {
        case "":
            if true {
                v34 = (v4 % ((v4) | 1))
                var v35 int32 = 0
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v35
            }
        case "foo bar":
            if (-77 == -128) {
                if (true || (4294967295 > 930)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if (!(65535 <= uint32(v32))) {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            fmt.Println(v34, v31, v32)
                            _ = fi36
                        }
                        v3 = append(v3, 687)
                    } else {
                        v37 := GenS1{F0: 100, F1: 65535, F2: 76}
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 17; fi38++ {
                    v32 += 9
                    var v39 int32 = 42
                    _ = fi38
                    _ = v39
                }
            } else {
                v40 := GenS3{F0: 2, F1: (-(v4)), F2: (v1 - (v1 | v1))}
                var v41 uint16 = 255
                v6 = append(v6, 4294967295)
                _ = v40
                _ = v41
            }
        }
        var v42 uint8 = (v7 * (v32 >> ((18446744073709551615) & 63)))
        _ = fi33
        _ = v34
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        switch ("\n\t" + "hello") {
        case "":
            switch v5 {
            case 1:
                var v44 int8 = -1
                _ = v44
            case 2147483647:
                if (!true) {
                    if len(v6) > 0 {
                        v6[0] = 2147483648
                    }
                    v5 -= v5
                }
            case -9223372036854775808:
                v1 = v1
            }
        }
        v45 := GenS3{F0: 9, F1: -932, F2: v1}
        switch "a" {
        case "\n\t":
            v46 := GenS4{F0: (fi43 ^ v5), F1: true}
            _ = v46
        case "hello":
            if ("foo bar" == "a") {
                var v47 uint64 = 0
                var v48 int32 = v27
                v49 := GenS2{F0: -91, F1: 1, F2: 2}
                _ = v47
                _ = v48
                _ = v49
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v7 += v32
            }
        default:
            v1 = -87
        }
        _ = fi43
        _ = v45
    }
    var v50 uint16 = 32767
    v51 := GenS3{F0: 42, F1: 42, F2: 0}
    v31 = (v31 - (v31 / float64(v2)))
    switch "a" {
    case "foo bar":
        var v52 int64 = (v4 ^ int64(v50))
        _ = v52
    case "":
        if true {
            for fi53 := 0; fi53 < 4; fi53++ {
                if (0 >= 0) {
                    v1 -= (v1 / (((^(v1))) | 1))
                }
                if (true || (2147483648 <= 1)) {
                    v31 = 1e10
                    v54 := GenS3{F0: 1, F1: int64(v27), F2: v1}
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 13; fi55++ {
                        v56 := GenS5{F0: v4}
                        _ = fi55
                        _ = v56
                    }
                    v57 := GenS3{F0: 14, F1: v4, F2: 0}
                    _ = v57
                }
                for fi58 := 0; fi58 < 9; fi58++ {
                    var v59 bool = (!(1.5 > (v31 * v31)))
                    _ = fi58
                    _ = v59
                }
                _ = fi53
            }
            var v60 int = v5
            v61 := GenS2{F0: 8, F1: 65535, F2: 2147483647}
            _ = v60
            _ = v61
        } else {
            v7 -= v7
        }
    default:
        v4 = -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
}
