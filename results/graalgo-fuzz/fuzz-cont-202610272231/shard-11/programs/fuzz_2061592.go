package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint8; F1 string; F2 int32 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 37
    var v3 uint32 = 4294967295
    if (1000 >= 42) {
        if false {
            for fi4 := 0; fi4 < 17; fi4++ {
                for fi5 := 0; fi5 < 6; fi5++ {
                    v2 = (v2 % ((v2) | 1))
                    _ = fi5
                }
                v2 = v2
                _ = fi4
            }
        }
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -410
        v8 := GenS1{F0: v7, F1: (v3 & (v3 - 2)), F2: 9223372036854775807}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 930, F1: 4294967295, F2: -9223372036854775808}
    var v10 int8 = 42
    fmt.Println(v10)
    var v11 float64 = (1e10 * -2.25)
    fmt.Println(v2, v11)
    v11 -= (v11 * 3.14159)
    v12 := []int{128}
    var v13 int = -2147483648
    var v14 string = ("" + "\n\t")
    _ = len(v12)
    switch v14 {
    case "hello":
        fmt.Println(v13, v11, v13)
    case "a":
        if true {
            fmt.Println(v3)
            if true {
                var v15 int64 = -1
                if (!((v15 > (^(v15))) || (761 >= -74))) {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 int = -822
                        if len(v12) > 0 {
                            v12[0] = 128
                        }
                        v11 = 1.5
                        _ = fi16
                        _ = v17
                    }
                    if false {
                        var v18 uint = 65535
                        _ = v18
                    } else {
                        if false {
                            if (853 == -32768) {
                                _ = len(v12)
                            }
                        } else {
                            fmt.Println(v13, v10, v3)
                            var v19 int64 = (v15 % ((v15) | 1))
                            for fi20 := 0; fi20 < 2; fi20++ {
                                fmt.Println(v10)
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 float32 = 3.14159
                                    _ = fi21
                                    _ = v22
                                }
                                _ = fi20
                            }
                            _ = v19
                        }
                    }
                    v13 = v13
                }
                v23 := GenS2{F0: (("hello" + v14) + ("a" + v14)), F1: (v10 | v10)}
                _ = v15
                _ = v23
            }
        }
    case "\n\t":
        var v24 float32 = 0.1
        _ = v24
    }
    v10 += v10
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS1{F0: -940, F1: v3, F2: 735}
        if false {
            if true {
                var v27 uint64 = 656
                _ = v27
            }
            if (!true) {
                var v28 uint32 = (v3 - 2147483648)
                switch v14 {
                case "foo bar":
                    if (!true) {
                        if (-2.25 >= v11) {
                            v3 = 4294967295
                        }
                        var v29 uint16 = 577
                        v30 := GenS3{F0: (v3 | 2147483647), F1: 0, F2: (v14 + "\n\t")}
                        _ = v29
                        _ = v30
                    }
                case "":
                    fmt.Println(v13, v10, v10)
                case "hello":
                    var v31 float32 = (0.1 * -2.25)
                    _ = v31
                default:
                    fmt.Println(fi25, v1, fi25)
                }
                _ = len(v12)
                _ = v28
            } else {
                if (uint16(v2) >= v1) {
                    if len(v12) > 0 {
                        v12[0] = 66
                    }
                    if false {
                        var v32 int16 = -586
                        switch -9223372036854775808 {
                        case -507:
                            for fi33 := 0; fi33 < 9; fi33++ {
                                var v34 int16 = (v32 * v32)
                                _ = fi33
                                _ = v34
                            }
                        case 711:
                            if (v14 <= v14) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v12 = append(v12, fi25)
                                if (true || false) {
                                    v10 = (v10 & (v10 * v10))
                                    if len(v12) > 0 {
                                        v12[0] = (-(v13))
                                    }
                                    var v35 string = (v14 + v14)
                                    _ = v35
                                } else {
                                    if (!false) {
                                        fmt.Println(v14, v32)
                                    }
                                    v36 := GenS4{F0: 541}
                                    _ = v36
                                }
                            } else {
                                var v37 int64 = 42
                                _ = v37
                            }
                        case -2147483648:
                            for fi38 := 0; fi38 < 6; fi38++ {
                                v12 = append(v12, (-(v13)))
                                fmt.Println(v2, v10)
                                v3 -= (^(v3))
                                _ = fi38
                            }
                        }
                        _ = v32
                    }
                }
                switch (v14 + "a") {
                case "hello":
                    v39 := GenS4{F0: (v1 & (^(v1)))}
                    _ = v39
                case "":
                    var v40 int8 = 127
                    _ = v40
                case "foo bar":
                    if false {
                        v41 := GenS5{F0: 127, F1: 1}
                        if false {
                            v42 := GenS1{F0: 32767, F1: 42, F2: int64(v10)}
                            _ = v42
                        }
                        var v43 uint64 = 42
                        _ = v41
                        _ = v43
                    } else {
                        v12 = append(v12, -9223372036854775808)
                        var v44 uint = 65535
                        for fi45 := 0; fi45 < 3; fi45++ {
                            fmt.Println(fi45)
                            _ = fi45
                        }
                        _ = v44
                    }
                default:
                    for fi46 := 0; fi46 < 3; fi46++ {
                        fmt.Println(v10, v10, v14)
                        _ = fi46
                    }
                }
                fmt.Println(v1)
            }
        } else {
            var v47 float64 = (v11 * v11)
            v48 := GenS6{F0: v2, F1: v14, F2: -1}
            v1 = uint16(v2)
            _ = v47
            _ = v48
        }
        _ = fi25
        _ = v26
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        if false {
            if false {
                var v50 bool = (-2147483648 > 42)
                if (2 < 998) {
                    if false {
                        for fi51 := 0; fi51 < 14; fi51++ {
                            if ((v1 != (v1 >> ((287) & 63))) && (v50 || v50)) {
                                fmt.Println(v10, v14, v2)
                                var v52 int32 = int32(v10)
                                _ = v52
                            }
                            v12 = append(v12, v13)
                            v13 = (v13 % (((fi51 & (v13 - (v13 << ((0) & 63))))) | 1))
                            _ = fi51
                        }
                    }
                } else {
                    var v53 float64 = (v11 * (v11 - 3.14159))
                    var v54 uint32 = (v3 << ((255) & 63))
                    var v55 uint = 42
                    _ = v53
                    _ = v54
                    _ = v55
                }
                _ = v50
            } else {
                if true {
                    v56 := GenS1{F0: 0, F1: v3, F2: int64(v1)}
                    _ = v56
                }
            }
            switch v13 {
            case 9223372036854775807:
                if true {
                    v10 -= v10
                    v10 = (^(v10))
                    v12 = append(v12, v13)
                } else {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    for fi57 := 0; fi57 < 1; fi57++ {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        var v58 int = fi49
                        _ = fi57
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 17; fi59++ {
                        fmt.Println(v14)
                        v60 := GenS7{F0: 42}
                        fmt.Println(fi49, v11)
                        _ = fi59
                        _ = v60
                    }
                }
            default:
                v61 := GenS1{F0: 1, F1: v3, F2: 9223372036854775807}
                _ = v61
            }
            if (v2 <= v2) {
                var v62 int16 = 42
                _ = v62
            } else {
                fmt.Println(fi49)
                fmt.Println(v10)
                if false {
                    v63 := GenS4{F0: (^(v1))}
                    switch (v14 + "foo bar") {
                    case "foo bar":
                        var v64 bool = true
                        _ = v64
                    case "\n\t":
                        v1 = v1
                    case "":
                        var v65 uint32 = (v3 + (v3 << ((uint(v2)) & 63)))
                        _ = v65
                    default:
                        if len(v12) > 0 {
                            v12[0] = fi49
                        }
                    }
                    if false {
                        v14 = "\n\t"
                        var v66 float64 = (0.1 / v11)
                        _ = v66
                    }
                    _ = v63
                } else {
                    for fi67 := 0; fi67 < 13; fi67++ {
                        v11 += (float64(0.1) * (float64((float32(v11) - 1e10)) - float64((((1e10 - 0.1) / float32(float64(0.1))) * (3.14159 - -2.25)))))
                        _ = len(v12)
                        _ = fi67
                    }
                }
            }
        }
        v2 = 127
        var v68 int32 = 747
        _ = fi49
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
