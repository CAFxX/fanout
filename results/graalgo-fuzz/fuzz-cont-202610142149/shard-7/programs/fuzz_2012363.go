package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint64 = 0
    var v2 uint = 785
    v3 := []uint32{0, 65535}
    v4 := []int{-2147483648, 128, 863}
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint32 = 65535
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 float32 = 1e10
            if false {
                var v9 uint = 779
                if (!((0 < 2) && true)) {
                    v6 = v6
                }
                var v10 float64 = (3.14159 - 1e10)
                _ = v10
                _ = v9
            } else {
                v6 -= (v6 << ((v2) & 63))
                v11 := GenS1{F0: (v1 + v1)}
                v8 -= v8
                _ = v11
            }
            var v12 uint = 65535
            _ = fi7
            _ = v12
            _ = v8
        }
        v13 := GenS1{F0: 430}
        _ = fi5
        _ = v13
        _ = v6
    }
    var v14 uint16 = 32767
    var v15 uint = 65535
    for fi16 := 0; fi16 < 2; fi16++ {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        if true {
            fmt.Println(v1)
        }
        var v17 int16 = 32767
        _ = fi16
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int = (^(fi18))
            switch "hello" {
            case "\n\t":
                switch "\n\t" {
                case "":
                    v1 -= 190
                case "foo bar":
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v21 := GenS1{F0: v1}
                        _ = fi20
                        _ = v21
                    }
                case "hello":
                    fmt.Println(v1, v15)
                default:
                    _ = len(v4)
                }
            }
            var v22 uint64 = (v1 - (v1 % ((v1) | 1)))
            _ = fi18
            _ = v19
            _ = v22
        }
        var v23 int = 42
        _ = v23
    } else {
        if (true || ((("\n\t" + ("hello" + "")) + "foo bar") >= ("hello" + ""))) {
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(fi24)
                for fi25 := 0; fi25 < 15; fi25++ {
                    fmt.Println(fi25, v2, v2)
                    switch 42 {
                    case 0:
                        fmt.Println(fi25, v15)
                        fmt.Println(v2, fi24, fi24)
                        if (1e10 <= ((((3.14159 * (float32(((1e10 - (float64(float32(1e10)) / float64(1e10))) + 1.5)) * 1e10)) + float32((1.5 * 3.14159))) / (-2.25 + 1e10)) + 0.0)) {
                            v15 -= v15
                        }
                        v2 = v2
                        var v26 int32 = 42
                        var v27 uint8 = 255
                        if (false && (42 != 0)) {
                            var v28 int64 = 988
                            _ = v28
                        } else {
                            var v29 int64 = 9223372036854775807
                            v30 := GenS2{F0: v14, F1: v26}
                            v31 := GenS2{F0: (v14 + (^(v14))), F1: v26}
                            var v32 float32 = 0.1
                            var v33 int8 = 42
                            if len(v3) > 0 {
                                v3[0] = 129
                            }
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = v26
                        _ = v27
                    case 2147483647:
                        v34 := GenS2{F0: 1, F1: 1}
                        _ = v34
                    }
                    _ = fi25
                }
                _ = fi24
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                if true {
                    v36 := GenS3{F0: 0}
                    v14 += (v14 | v14)
                    _ = v36
                } else {
                    if (!(v1 == v1)) {
                        if true {
                            for fi37 := 0; fi37 < 15; fi37++ {
                                _ = len(v4)
                                switch "" {
                                case "a":
                                    fmt.Println(v2)
                                    if true {
                                        var v38 int16 = -1
                                        v1 += uint64(v14)
                                        var v39 float64 = (1e10 * float64((float32(3.14159) + (1e10 + 3.14159))))
                                        _ = v38
                                        _ = v39
                                    }
                                case "":
                                    if (fi37 == 32768) {
                                        v2 += (v2 / (((v2 ^ (v2 - v15))) | 1))
                                    } else {
                                        v40 := GenS3{F0: 2147483647}
                                        _ = v40
                                    }
                                }
                                _ = fi37
                            }
                            for fi41 := 0; fi41 < 8; fi41++ {
                                if (42 < -32768) {
                                    if true {
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = (^(fi41))
                                            }
                                            v2 += 4294967295
                                            v2 += (v2 - 1)
                                        } else {
                                            var v42 float32 = (1.5 - 1.5)
                                            v14 = 65535
                                            var v43 int8 = -128
                                            v2 -= v15
                                            if (!true) {
                                                fmt.Println(v43, v43, fi41)
                                                switch "" {
                                                case "foo bar":
                                                    if len(v3) > 0 {
                                                        v3[0] = 2147483648
                                                    }
                                                    fmt.Println(fi41, v43)
                                                }
                                                v42 += (1e10 / (float32(-2.25) * ((float32((0.1 + 3.14159)) + v42) * (v42 + v42))))
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if ((v2 | v2) >= (^(v15))) {
                                                    v14 += (v14 | (v14 << (((v15 >> ((uint(v14)) & 63))) & 63)))
                                                    var v44 uint = v2
                                                    _ = v44
                                                }
                                                var v45 uint16 = (v14 % (((v14 >> ((v2) & 63))) | 1))
                                                var v46 uint64 = v1
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v42
                                            _ = v43
                                        }
                                        var v47 float64 = 1.5
                                        _ = v47
                                    }
                                    v3 = append(v3, 579)
                                }
                                switch "a" {
                                case "a":
                                    v14 = 863
                                default:
                                    if false {
                                        var v48 uint32 = 42
                                        fmt.Println(v48)
                                        var v49 float64 = float64(1.5)
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                var v50 float32 = (0.0 * (1e10 + (1.5 + 0.0)))
                                _ = fi41
                                _ = v50
                            }
                        } else {
                            if false {
                                switch "hello" {
                                case "hello":
                                    v2 = 18446744073709551615
                                case "a":
                                    var v51 float64 = 3.14159
                                    _ = v51
                                }
                                if true {
                                    v52 := GenS3{F0: 2147483648}
                                    _ = v52
                                }
                                switch fi35 {
                                case 32768:
                                    v14 += 0
                                case 4294967295:
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                }
                            } else {
                                v14 = (v14 | v14)
                            }
                            var v53 float32 = 1.5
                            _ = v53
                        }
                        var v54 int = fi35
                        v3 = append(v3, 2147483647)
                        _ = v54
                    }
                    v2 = (v2 + uint(v1))
                    switch (^(fi35)) {
                    case 42:
                        for fi55 := 0; fi55 < 19; fi55++ {
                            if (false || false) {
                                v4 = append(v4, fi55)
                                var v56 uint8 = 198
                                fmt.Println(fi55, v2, fi35)
                                _ = v56
                            } else {
                                v2 = (v15 | (v2 + (v15 / ((524) | 1))))
                                v14 = (v14 * v14)
                                v1 -= v1
                            }
                            fmt.Println(v14)
                            v15 = v2
                            v57 := GenS1{F0: v1}
                            _ = fi55
                            _ = v57
                        }
                    case -545:
                        var v58 int64 = -9223372036854775808
                        _ = v58
                    case -1:
                        v59 := GenS1{F0: v1}
                        _ = v59
                    }
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    v3 = append(v3, 2)
                    _ = fi60
                }
                for fi61 := 0; fi61 < 10; fi61++ {
                    v62 := GenS3{F0: 281}
                    switch "foo bar" {
                    case "hello":
                        var v63 int = (-(fi35))
                        _ = v63
                    case "a":
                        if (168 <= 2) {
                            v1 -= (v1 & 832)
                            v2 += (v15 << ((65535) & 63))
                            v64 := GenS4{F0: v1, F1: 100}
                            _ = v64
                        }
                    default:
                        var v65 float64 = (float64(((3.14159 + (((((-2.25 * 1e10) + 1.5) + (0.0 + (0.0 / (float32(-2.25) - (-2.25 / (1.5 / -2.25)))))) * float32(-2.25)) / float32(1e10))) / 1e10)) + 0.1)
                        _ = v65
                    }
                    _ = fi61
                    _ = v62
                }
                _ = fi35
            }
            var v66 int = -2147483648
            _ = v66
        } else {
            var v67 uint = uint(v1)
            _ = v67
        }
    }
    v1 = (v1 & (v1 / (((v1 + (v1 / ((v1) | 1)))) | 1)))
    v14 -= v14
    var v68 uint16 = (v14 / (((v14 & v14)) | 1))
    if false {
        if true {
            if true {
                _ = len(v4)
            }
            if false {
                fmt.Println(v1, v15, v15)
                fmt.Println(v2, v2, v1)
                v68 = v68
            }
            if true {
                v69 := GenS5{F0: (v15 > (v2 * (v2 ^ v2)))}
                v14 = 645
                _ = v69
            }
        } else {
            var v70 int = 2147483647
            v15 += (v15 >> (((v2 + (^(v2)))) & 63))
            _ = v70
        }
        if false {
            for fi71 := 0; fi71 < 15; fi71++ {
                v1 -= v1
                switch fi71 {
                case 128:
                    v68 = v68
                case 0:
                    var v72 bool = (0 == 1000)
                    _ = v72
                }
                _ = fi71
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v68
}
