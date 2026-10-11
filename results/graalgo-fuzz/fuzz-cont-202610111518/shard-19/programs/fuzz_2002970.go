package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 704
    var v3 int = 4294967295
    v4 := []int{(^(v3)), v3, 127, (v3 % (((v3 >> ((1) & 63))) | 1))}
    if ((true && false) || true) {
        var v5 int8 = 38
        v6 := GenS1{F0: 42}
        _ = v5
        _ = v6
    }
    var v7 float64 = 0.0
    switch (v3 & -9223372036854775808) {
    case -2147483648:
        v8 := GenS2{F0: 2, F1: 2}
        _ = v8
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint64 = 156
            var v11 uint16 = 828
            var v12 float32 = ((-2.25 - (3.14159 - 1.5)) + float32(v7))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13)
        fmt.Println(v3, v3, v2)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi13
    }
    v14 := GenS2{F0: 786, F1: 0}
    var v15 uint = 724
    v16 := GenS1{F0: 2}
    for fi17 := 0; fi17 < 20; fi17++ {
        v15 = 0
        v18 := GenS1{F0: 32767}
        _ = fi17
        _ = v18
    }
    var v19 int = (v3 * v3)
    if ((!true) && (42 != 42)) {
        v1 += (v1 + v1)
        v2 = v2
    } else {
        var v20 float32 = (3.14159 * (((1e10 - -2.25) * ((((((0.0 - (0.0 / 3.14159)) + 1e10) + 1e10) * 0.0) - ((1e10 * (1e10 / 1.0)) + 3.14159)) * (0.0 * float32(1e10)))) + (float32(v7) * 0.1)))
        var v21 string = "hello"
        if true {
            for fi22 := 0; fi22 < 7; fi22++ {
                switch 9223372036854775807 {
                case 32768:
                    for fi23 := 0; fi23 < 13; fi23++ {
                        v3 -= fi23
                        if true {
                            var v24 uint32 = v2
                            var v25 uint32 = v24
                            v1 = (v1 ^ (v1 % ((42) | 1)))
                            _ = v24
                            _ = v25
                        } else {
                            _ = len(v4)
                            v20 = 3.14159
                        }
                        _ = fi23
                    }
                default:
                    v20 = ((v20 * v20) / v20)
                }
                _ = fi22
            }
            switch (v19 << ((503) & 63)) {
            case 1:
                if (true && true) {
                    v4 = append(v4, (^(v19)))
                }
            case 128:
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := GenS1{F0: 266}
                    _ = fi26
                    _ = v27
                }
            default:
                if false {
                    v28 := GenS1{F0: 32767}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 bool = (false || (((1.5 - 1.5) / float64(float32(float64((v20 + float32(v7)))))) <= 0.0))
                        var v31 uint16 = 758
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if true {
                            if false {
                                fmt.Println(v7)
                                _ = len(v4)
                                v1 += v1
                            }
                            var v33 int8 = 42
                            switch (-(v3)) {
                            case 32768:
                                var v34 int16 = -789
                                _ = v34
                            }
                            _ = v33
                        }
                        var v35 uint32 = v2
                        _ = fi32
                        _ = v35
                    }
                    switch (v19 + (v19 ^ (v19 & (-(v3))))) {
                    case 42:
                        var v36 int32 = -947
                        _ = v36
                    case -2147483648:
                        switch "hello" {
                        case "foo bar":
                            for fi37 := 0; fi37 < 8; fi37++ {
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                if ((^(v3)) >= (v3 << (((v15 | v15)) & 63))) {
                                    var v38 uint32 = 3
                                    for fi39 := 0; fi39 < 12; fi39++ {
                                        if (v20 != (-2.25 + (1e10 * v20))) {
                                            if (-128 > 0) {
                                                var v40 int16 = 1
                                                _ = v40
                                            }
                                            var v41 string = ""
                                            fmt.Println(v2, v41, fi39)
                                            v42 := GenS1{F0: 2}
                                            _ = v41
                                            _ = v42
                                        }
                                        v38 = (v2 / ((v38) | 1))
                                        fmt.Println(fi39, v2)
                                        _ = fi39
                                    }
                                    fmt.Println(v19)
                                    _ = v38
                                } else {
                                    for fi43 := 0; fi43 < 4; fi43++ {
                                        v7 -= (v7 + v7)
                                        v44 := GenS3{F0: -2147483648, F1: v21, F2: (v2 - v2)}
                                        _ = fi43
                                        _ = v44
                                    }
                                    switch v19 {
                                    case 772:
                                        v21 = ""
                                    case 32767:
                                        var v45 string = ""
                                        _ = v45
                                    }
                                }
                                v7 = (v7 / 0.1)
                                _ = fi37
                            }
                        case "":
                            v21 = "a"
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if (-2147483648 == 2147483647) {
                                v46 := GenS2{F0: 705, F1: (v15 ^ v15)}
                                v20 -= 1e10
                                _ = v46
                            } else {
                                var v47 int64 = -9223372036854775808
                                for fi48 := 0; fi48 < 19; fi48++ {
                                    for fi49 := 0; fi49 < 8; fi49++ {
                                        _ = len(v4)
                                        var v50 int8 = -1
                                        v19 = (fi48 * 127)
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = fi48
                                }
                                for fi51 := 0; fi51 < 14; fi51++ {
                                    var v52 bool = false
                                    _ = fi51
                                    _ = v52
                                }
                                _ = v47
                            }
                        }
                    case 4294967295:
                        v2 = 1
                    default:
                        if (!false) {
                            for fi53 := 0; fi53 < 1; fi53++ {
                                for fi54 := 0; fi54 < 4; fi54++ {
                                    v20 += (v20 - (v20 - 0.1))
                                    v19 = fi53
                                    _ = len(v4)
                                    _ = fi54
                                }
                                _ = fi53
                            }
                        } else {
                            if true {
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    switch (v21 + ((("" + v21) + v21) + "")) {
                                    case "\n\t":
                                        v56 := GenS1{F0: 42}
                                        _ = v56
                                    }
                                    v2 = v2
                                    _ = fi55
                                }
                                var v57 uint64 = 0
                                fmt.Println(v2)
                                _ = v57
                            } else {
                                var v58 uint = v15
                                for fi59 := 0; fi59 < 2; fi59++ {
                                    if ((((v21 + v21) + "") + "foo bar") <= v21) {
                                        switch fi59 {
                                        case -519:
                                            switch v21 {
                                            case "hello":
                                                v60 := GenS2{F0: 18446744073709551615, F1: (^(v15))}
                                                _ = v60
                                            }
                                        case 69:
                                            v4 = append(v4, fi59)
                                        default:
                                            switch v21 {
                                            case "foo bar":
                                                v61 := GenS1{F0: 255}
                                                _ = v61
                                            }
                                        }
                                        fmt.Println(v21, v3, fi59)
                                        fmt.Println(v2)
                                    } else {
                                        switch 360 {
                                        case 1:
                                            for fi62 := 0; fi62 < 11; fi62++ {
                                                if (false || false) {
                                                    v20 -= v20
                                                }
                                                v3 += fi59
                                                _ = fi62
                                            }
                                        }
                                        v63 := GenS1{F0: 2}
                                        _ = v63
                                    }
                                    v2 -= (v2 * (v2 % ((2147483647) | 1)))
                                    _ = fi59
                                }
                                _ = v58
                            }
                            var v64 uint32 = (v2 / (((v2 % (((v2 * v2)) | 1))) | 1))
                            var v65 string = "hello"
                            _ = v64
                            _ = v65
                        }
                    }
                }
            }
            v15 -= v15
        }
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
