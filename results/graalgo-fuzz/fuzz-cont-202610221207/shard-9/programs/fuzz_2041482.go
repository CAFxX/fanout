package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 int8 = 19
    var v2 int = -1
    var v3 int16 = 1000
    var v4 int32 = -339
    var v5 bool = (!false)
    var v6 bool = (!(!v5))
    v7 := GenS1{F0: 42, F1: 42}
    var v8 int32 = (v4 + 837)
    v9 := GenS1{F0: 4294967295, F1: 65535}
    v10 := []int{2147483647}
    var v11 int8 = (v1 & (v1 ^ (v1 & v1)))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v12 := GenS1{F0: 2, F1: 419}
    switch "hello" {
    case "hello":
        v13 := GenS1{F0: 4294967295, F1: 944}
        _ = v13
    default:
        switch "foo bar" {
        case "a":
            v1 = v11
        case "foo bar":
            for fi14 := 0; fi14 < 16; fi14++ {
                v4 -= (v8 & (v4 << ((42) & 63)))
                _ = fi14
            }
        }
    }
    v15 := GenS1{F0: 1, F1: 2}
    var v16 uint = 18446744073709551615
    v1 += -124
    switch "hello" {
    case "":
        var v17 bool = v5
        _ = v17
    case "\n\t":
        v18 := GenS2{F0: v1, F1: 636}
        _ = v18
    case "foo bar":
        for fi19 := 0; fi19 < 9; fi19++ {
            v5 = true
            if false {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v2 -= (v2 & (fi19 * (fi19 % ((fi20) | 1))))
                    if (!true) {
                        _ = len(v10)
                        if len(v10) > 0 {
                            v10[0] = -1
                        }
                        if ((v1 == v1) || ((-(v3)) >= (v3 + (^(v3))))) {
                            fmt.Println(v1, v3, v6)
                            _ = len(v10)
                        } else {
                            v21 := GenS2{F0: -128, F1: v3}
                            var v22 int64 = -702
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v23 := GenS2{F0: (v1 % ((100) | 1)), F1: -461}
                        _ = v23
                    }
                    _ = fi20
                }
                v24 := GenS2{F0: 100, F1: (v3 + (v3 ^ v3))}
                var v25 uint8 = 127
                _ = v24
                _ = v25
            }
            _ = fi19
        }
    }
    v26 := GenS3{F0: (v11 % ((v1) | 1)), F1: v2, F2: 476}
    for fi27 := 0; fi27 < 11; fi27++ {
        switch "foo bar" {
        case "hello":
            v28 := GenS1{F0: 4294967295, F1: v16}
            _ = v28
        }
        v29 := GenS1{F0: 539, F1: v16}
        _ = fi27
        _ = v29
    }
    if true {
        v4 += v4
        switch (^(v2)) {
        case 32767:
            if (589 != 42) {
                v1 -= v11
                if (!((!v6) || (1.5 != (0.0 - 0.1)))) {
                    v10 = append(v10, (v2 ^ v2))
                    fmt.Println(v3)
                } else {
                    var v30 float64 = float64((0.1 + 1.5))
                    for fi31 := 0; fi31 < 19; fi31++ {
                        for fi32 := 0; fi32 < 4; fi32++ {
                            if v5 {
                                var v33 uint = 255
                                if (!((^(v2)) >= (v2 * 2147483647))) {
                                    v34 := GenS3{F0: (v11 + 100), F1: 32767, F2: 32767}
                                    if v6 {
                                        var v35 int64 = -9223372036854775808
                                        if ((false && (!true)) && true) {
                                            v36 := GenS1{F0: 576, F1: v16}
                                            _ = v36
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                v30 = v30
                                            case "foo bar":
                                                var v37 int = fi31
                                                _ = v37
                                            default:
                                                v33 = (v33 | (v16 | v16))
                                            }
                                            v38 := GenS4{F0: 1, F1: uint16(v33), F2: (^(v8))}
                                            _ = v38
                                        }
                                        _ = v35
                                    }
                                    switch fi32 {
                                    case -2147483648:
                                        v10 = append(v10, v2)
                                    }
                                    _ = v34
                                } else {
                                    var v39 float32 = float32(((float64((0.0 * -2.25)) + ((v30 + v30) / v30)) / ((float64((3.14159 + 0.0)) * 0.1) - v30)))
                                    v2 = (fi31 | (fi32 & v2))
                                    _ = v39
                                }
                                fmt.Println(v1)
                                _ = v33
                            }
                            var v40 uint8 = 209
                            var v41 int16 = 894
                            _ = fi32
                            _ = v40
                            _ = v41
                        }
                        v10 = append(v10, (v2 << ((v16) & 63)))
                        _ = fi31
                    }
                    _ = v30
                }
                if v5 {
                    var v42 int64 = 0
                    _ = v42
                }
            } else {
                _ = len(v10)
            }
        case 2147483647:
            switch "foo bar" {
            case "\n\t":
                if (42 < 0) {
                    fmt.Println(v3)
                    v10 = append(v10, (-(v2)))
                } else {
                    for fi43 := 0; fi43 < 1; fi43++ {
                        if true {
                            for fi44 := 0; fi44 < 2; fi44++ {
                                fmt.Println(v5)
                                v16 += 4294967295
                                _ = fi44
                            }
                            if true {
                                var v45 uint32 = 1
                                switch "hello" {
                                case "a":
                                    v8 = v8
                                }
                                _ = v45
                            } else {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        }
                        v46 := GenS5{F0: 42}
                        for fi47 := 0; fi47 < 8; fi47++ {
                            var v48 bool = (0 >= 752)
                            if ((v4 ^ v4) < (v4 | v4)) {
                                if (-988 >= (v4 + (v4 - v8))) {
                                    switch fi43 {
                                    case -1:
                                        v11 += v1
                                    case -2147483648:
                                        var v49 uint32 = 1
                                        _ = v49
                                    case 127:
                                        v5 = (42 >= 380)
                                    }
                                    v1 = (-(v11))
                                }
                                var v50 int32 = (^(v8))
                                _ = v50
                            }
                            if len(v10) > 0 {
                                v10[0] = (fi43 >> (((v16 & v16)) & 63))
                            }
                            _ = fi47
                            _ = v48
                        }
                        _ = fi43
                        _ = v46
                    }
                    _ = len(v10)
                }
            case "a":
                var v51 int64 = 1
                _ = v51
            case "foo bar":
                switch ("\n\t" + ("" + "hello")) {
                case "hello":
                    v4 += (v8 % (((v4 % ((int32(v16)) | 1))) | 1))
                }
            }
        case -886:
            v8 = 293
        }
        var v52 int64 = 9223372036854775807
        _ = v52
    } else {
        v16 += v16
        var v53 float64 = ((-2.25 / 3.14159) + 3.14159)
        _ = v53
    }
    for fi54 := 0; fi54 < 17; fi54++ {
        var v55 int16 = (-(v3))
        v2 += (^(fi54))
        fmt.Println(fi54, v3, v1)
        _ = fi54
        _ = v55
    }
    if v6 {
        var v56 uint8 = 0
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
