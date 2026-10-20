package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0() uint16 {
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42, F1: 0, F2: 185}
            var v10 uint8 = 0
            v11 := GenS2{F0: 42, F1: 4294967295, F2: true}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    _ = v6
    return 32767
}
func genFn1() uint {
    v12 := []int{42, -539, 32768}
    for fi13 := 0; fi13 < 10; fi13++ {
        _ = len(v12)
        var v14 int32 = 2147483647
        var v15 uint = 42
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 int16 = -118
            switch "hello" {
            case "":
                var v19 int = fi17
                _ = v19
            }
            var v20 int = fi17
            _ = fi17
            _ = v18
            _ = v20
        }
        _ = fi16
    }
    _ = v12
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = 1
    v3 := []uint32{v1, v1}
    var v4 int = 821
    v5 := []uint32{v1, 4294967295, 2147483648}
    var v21 float64 = (3.14159 - (-2.25 / (float64(-2.25) + 0.0)))
    if len(v3) > 0 {
        v3[0] = (v1 << ((576) & 63))
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        _ = len(v5)
        _ = fi22
    }
    if ((1e10 >= 1e10) || (42 == -1)) {
        var v23 int32 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v23
    } else {
        if true {
            var v24 int32 = 0
            v21 -= v21
            v21 += v21
            _ = v24
        } else {
            if (false && true) {
                switch "hello" {
                case "a":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 int32 = 2147483647
                        fmt.Println(v26)
                        _ = fi25
                        _ = v26
                    }
                case "\n\t":
                    fmt.Println(v1)
                }
            }
        }
        switch "" {
        case "foo bar":
            v4 += (v4 - (v4 << ((4294967295) & 63)))
        case "a":
            v27 := GenS1{F0: uint64(v1), F1: 2147483647, F2: 255}
            _ = v27
        case "hello":
            fmt.Println(v2)
        default:
            switch v4 {
            case -1:
                fmt.Println(v21, v4, v2)
            default:
                var v28 uint32 = 42
                _ = v28
            }
        }
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        switch fi29 {
        case 2147483647:
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS1{F0: 42, F1: 0, F2: 128}
                _ = fi30
                _ = v31
            }
        case -918:
            if (630 > int16(v1)) {
                if false {
                    var v32 float64 = (((((0.1 + v21) / (3.14159 * v21)) - v21) / (((-2.25 * -2.25) + (((1e10 * (0.1 - float64(0.1))) - (v21 + (float64(-2.25) * v21))) + float64(0.1))) / 3.14159)) - (float64(0.0) + (3.14159 - v21)))
                    if false {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v32 -= (-2.25 / -2.25)
                            v34 := GenS2{F0: 42, F1: 755, F2: false}
                            _ = fi33
                            _ = v34
                        }
                        var v35 int = 961
                        fmt.Println(v1, v1, v4)
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 12; fi36++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v1 -= (v1 << ((genFn1()) & 63))
                        switch 9223372036854775807 {
                        case -243:
                            v32 = ((v32 - 0.0) * v32)
                            if (443 != 361) {
                                v32 = v21
                                switch "" {
                                case "hello":
                                    v1 -= 1
                                case "a":
                                    v2 += v2
                                case "foo bar":
                                    fmt.Println(v2, v1, v1)
                                }
                            }
                        case 9223372036854775807:
                            v37 := GenS1{F0: 42, F1: 1, F2: 42}
                            _ = v37
                        case 4294967295:
                            switch ("a" + "") {
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = 42
                                }
                                v38 := GenS2{F0: 2, F1: 1, F2: false}
                                _ = v38
                            case "\n\t":
                                v32 = ((((0.1 + v32) / v32) + (float64(1e10) * ((v21 + v32) + ((((float64((1e10 * float32(1.5))) - v21) * (((0.0 - 3.14159) + ((((float64(3.14159) - (v21 / (((v32 / v21) - (float64(float32((v21 / v21))) + v32)) + (0.0 * v21)))) / -2.25) + v21) - 3.14159)) + (v32 + v32))) * ((((float64(1.5) - ((1.5 * ((float64((1e10 * 3.14159)) + v21) * -2.25)) - v32)) * ((1.5 / -2.25) / v21)) / v32) + (v21 - (v21 / 1.5)))) - -2.25)))) + 1.5)
                            default:
                                v1 += (v1 + (v1 + v1))
                            }
                        default:
                            if true {
                                var v39 uint64 = 505
                                fmt.Println(v39, fi36)
                                var v40 uint = 873
                                v21 = float64((0.1 + (1e10 * (1.5 / float32(v32)))))
                                v41 := GenS2{F0: v39, F1: v39, F2: (255 <= (^(v40)))}
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = fi36
                    }
                    _ = v32
                }
                var v42 int = fi29
                v2 = (v2 & v2)
                _ = v42
            }
        }
        _ = fi29
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        var v44 int64 = (v2 & (v2 ^ v2))
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        var v46 int16 = 374
        var v47 int32 = -2147483648
        _ = fi45
        _ = v46
        _ = v47
    }
    v4 -= 0
    v4 = (v4 >> ((18446744073709551615) & 63))
    var v48 uint64 = 109
    for fi49 := 0; fi49 < 3; fi49++ {
        v50 := GenS1{F0: (v48 / ((uint64(v1)) | 1)), F1: -2147483648, F2: 255}
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        v52 := GenS1{F0: (v48 * (v48 * v48)), F1: -929, F2: 79}
        if (float64((0.0 / 1.0)) <= (v21 + 1.5)) {
            v48 = (v48 & (v48 * v48))
        }
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 7; fi53++ {
        v54 := GenS2{F0: v48, F1: v48, F2: true}
        _ = fi53
        _ = v54
    }
    if false {
        v1 += 2
        var v55 bool = (2147483648 != (^(v1)))
        var v56 uint64 = (^(v48))
        _ = v55
        _ = v56
    } else {
        var v57 bool = (false && false)
        _ = v57
    }
    for fi58 := 0; fi58 < 7; fi58++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v59 uint64 = 184
        v60 := GenS3{F0: (v1 >> ((255) & 63)), F1: v59}
        _ = fi58
        _ = v59
        _ = v60
    }
    for fi61 := 0; fi61 < 16; fi61++ {
        v62 := GenS4{F0: v1, F1: 0}
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 5; fi63++ {
        v64 := GenS2{F0: (v48 / ((v48) | 1)), F1: 2, F2: (0 > 1)}
        var v65 int32 = 1
        fmt.Println(fi63, v4)
        _ = fi63
        _ = v64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
