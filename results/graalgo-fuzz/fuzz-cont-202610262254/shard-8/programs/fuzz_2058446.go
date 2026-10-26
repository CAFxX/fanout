package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 float64 = 0.0
    var v2 int = 2147483647
    var v3 string = "hello"
    v4 := []uint32{471, 119}
    var v5 uint32 = 353
    var v6 uint8 = 128
    var v7 uint8 = v6
    v3 = v3
    v8 := GenS1{F0: 32767}
    if (((v3 + "foo bar") + v3) != ((v3 + (v3 + v3)) + (v3 + v3))) {
        fmt.Println(v1, v7)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch v2 {
        case -1:
            v9 := GenS2{F0: 1000, F1: (^(v5)), F2: (v6 / ((v6) | 1))}
            _ = v9
        default:
            switch ("hello" + ((v3 + v3) + (("\n\t" + "hello") + v3))) {
            case "foo bar":
                if (127 < -124) {
                    v7 = v6
                }
            case "a":
                v10 := GenS2{F0: -1, F1: v5, F2: v6}
                _ = v10
            case "\n\t":
                if false {
                    _ = len(v4)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    if false {
                        var v11 int32 = 0
                        _ = v11
                    }
                    v4 = append(v4, v5)
                    v4 = append(v4, (v5 >> ((42) & 63)))
                }
            }
        }
        if true {
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    switch "hello" {
                    case "foo bar":
                        v5 = 2147483648
                    case "a":
                        if true {
                            if (973 > 42) {
                                v14 := GenS2{F0: -961, F1: v5, F2: v7}
                                _ = v14
                            }
                            var v15 float32 = 1e10
                            _ = v15
                        }
                    case "":
                        if true {
                            if (4294967295 == 0) {
                                v16 := GenS2{F0: -69, F1: (v5 >> ((4294967295) & 63)), F2: (v7 % (((^(v6))) | 1))}
                                if false {
                                    var v17 string = v3
                                    _ = v17
                                }
                                var v18 uint64 = 411
                                _ = v16
                                _ = v18
                            }
                            v19 := GenS3{F0: 42, F1: int64(v5), F2: 32767}
                            _ = v19
                        }
                    default:
                        fmt.Println(v7, fi13)
                    }
                    _ = fi13
                }
                switch v2 {
                case 2147483647:
                    var v20 int64 = 0
                    _ = v20
                case 629:
                    switch v3 {
                    case "":
                        for fi21 := 0; fi21 < 15; fi21++ {
                            var v22 int32 = 625
                            switch (fi12 / (((v2 | (fi12 / (((v2 / ((fi21) | 1))) | 1)))) | 1)) {
                            case 2147483647:
                                var v23 bool = ((790 >= 4294967295) || (true && (2 < 737)))
                                var v24 uint32 = 2147483647
                                var v25 float64 = ((v1 / (0.1 / v1)) - float64(float32(3.14159)))
                                _ = v23
                                _ = v24
                                _ = v25
                            case -220:
                                v7 -= (v7 & v7)
                                fmt.Println(v6, v3, v5)
                                switch (v2 | 32768) {
                                case -595:
                                    var v26 uint16 = 839
                                    _ = v26
                                case 727:
                                    var v27 bool = ((float32(v1) - 3.14159) != 1e10)
                                    _ = v27
                                default:
                                    v6 -= v6
                                }
                                if true {
                                    if ((true && (!true)) && true) {
                                        v22 -= (-(v22))
                                    } else {
                                        v28 := GenS3{F0: (v22 - (^(v22))), F1: 9223372036854775807, F2: 32767}
                                        var v29 float32 = 0.1
                                        _ = v28
                                        _ = v29
                                    }
                                } else {
                                    v4 = append(v4, 4294967295)
                                }
                                v22 = v22
                                v30 := GenS4{F0: (v6 * 0), F1: 594, F2: 0}
                                _ = v30
                            case 32768:
                                if (0.1 <= 3.14159) {
                                    v31 := GenS4{F0: v7, F1: 106, F2: 255}
                                    var v32 uint32 = (v5 | 196)
                                    if false {
                                        if false {
                                            _ = len(v4)
                                        } else {
                                            fmt.Println(fi21, v3)
                                            v1 = (v1 / v1)
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v33 float32 = 0.0
                                        if true {
                                            var v34 int32 = v22
                                            switch -2147483648 {
                                            case 110:
                                                v2 += (^(fi21))
                                            case 70:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            default:
                                                var v35 int64 = 0
                                                _ = v35
                                            }
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                    _ = v31
                                    _ = v32
                                }
                            }
                            if false {
                                v22 -= v22
                                v7 = 226
                                v36 := GenS3{F0: (v22 * -2147483648), F1: 255, F2: 442}
                                var v37 int = (v2 & fi12)
                                fmt.Println(v6)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v36
                                _ = v37
                            }
                            v1 += v1
                            var v38 uint64 = 736
                            v39 := GenS2{F0: -32768, F1: v5, F2: (v6 * v7)}
                            _ = fi21
                            _ = v22
                            _ = v38
                            _ = v39
                        }
                    case "foo bar":
                        v40 := GenS3{F0: -1, F1: 42, F2: 458}
                        _ = v40
                    default:
                        v7 -= 1
                    }
                default:
                    for fi41 := 0; fi41 < 19; fi41++ {
                        var v42 int8 = 42
                        _ = fi41
                        _ = v42
                    }
                }
                _ = fi12
            }
            v2 -= (^(v2))
        }
        if (521 != 0) {
            var v43 uint64 = 0
            _ = v43
        } else {
            switch -1 {
            case -9223372036854775808:
                v1 += v1
            case 42:
                var v44 bool = (191 >= -32768)
                _ = v44
            case 2147483647:
                fmt.Println(v6, v5, v3)
            }
        }
    }
    var v45 int8 = 57
    var v46 uint32 = v5
    if true {
        var v47 uint64 = 42
        for fi48 := 0; fi48 < 14; fi48++ {
            if ((((((1 <= -169) && ((-1 == 0) && false)) && true) && true) || (128 >= (^(v2)))) && false) {
                _ = len(v4)
            } else {
                for fi49 := 0; fi49 < 2; fi49++ {
                    v6 = (v7 + (v7 / ((v7) | 1)))
                    v50 := GenS3{F0: -1, F1: -9223372036854775808, F2: 1000}
                    v2 = 0
                    fmt.Println(v46)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi49
                    _ = v50
                }
                v47 += (v47 - v47)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            if true {
                v1 += 3.14159
                v46 = (v5 & (v5 << ((18446744073709551615) & 63)))
                v46 = 563
            }
            _ = fi48
        }
        _ = v47
    } else {
        v4 = append(v4, v5)
        var v51 bool = true
        _ = v51
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
