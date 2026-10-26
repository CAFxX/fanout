package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint8; F1 string; F2 bool }
type GenS7 struct { F0 int16; F1 uint16; F2 int32 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    var v1 int = 32767
    var v2 uint8 = 254
    v3 := []uint32{}
    var v4 uint32 = 715
    var v5 uint = 78
    switch (("hello" + "foo bar") + "\n\t") {
    case "\n\t":
        v4 -= 65535
    }
    var v6 int16 = int16(v4)
    var v7 uint8 = 0
    switch (("a" + "hello") + "\n\t") {
    case "foo bar":
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int16 = v6
            var v10 uint64 = uint64(v5)
            _ = fi8
            _ = v10
            _ = v9
        }
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            v4 -= (v4 % ((v4) | 1))
            fmt.Println(v4)
            for fi12 := 0; fi12 < 13; fi12++ {
                switch -9223372036854775808 {
                case 180:
                    v6 -= v6
                case -537:
                    v13 := GenS1{F0: (fi11 << (((v5 ^ 255)) & 63)), F1: 255, F2: 18446744073709551615}
                    _ = v13
                case 32768:
                    var v14 int64 = -260
                    _ = v14
                default:
                    _ = len(v3)
                }
                if len(v3) > 0 {
                    v3[0] = (v4 ^ v4)
                }
                if true {
                    v15 := GenS1{F0: fi11, F1: 2, F2: 0}
                    v16 := GenS2{F0: v4, F1: (v6 >> (((v5 / (((^(v5))) | 1))) & 63))}
                    v3 = append(v3, (^(v4)))
                    _ = v15
                    _ = v16
                }
                _ = fi12
            }
            _ = fi11
        }
    default:
        var v17 uint8 = v7
        _ = v17
    }
    v18 := GenS1{F0: v1, F1: 2, F2: 0}
    v19 := GenS3{F0: 225, F1: -1}
    fmt.Println(v7)
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float64 = 0.1
        if true {
            v3 = append(v3, uint32(v7))
        }
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 32767
    v23 := GenS4{F0: v4}
    var v24 uint32 = 710
    if false {
        var v25 int8 = -56
        v26 := GenS3{F0: 804, F1: 0}
        var v27 float64 = 0.0
        _ = v25
        _ = v26
        _ = v27
    } else {
        switch v1 {
        case 655:
            if (!(!(-1 != 1))) {
                switch v1 {
                case -1:
                    fmt.Println(v2, v5)
                }
                for fi28 := 0; fi28 < 13; fi28++ {
                    if true {
                        v29 := GenS1{F0: int(v6), F1: (v22 % ((v22) | 1)), F2: 923}
                        _ = v29
                    } else {
                        var v30 uint = v5
                        for fi31 := 0; fi31 < 12; fi31++ {
                            if (!((((("foo bar" + ("foo bar" + (("a" + "") + "hello"))) + "hello") + "\n\t") + (("a" + "a") + "hello")) == ("foo bar" + ("hello" + "")))) {
                                v1 = fi28
                                v2 += v7
                                if (((false || false) && false) || (!(2 >= 255))) {
                                    fmt.Println(fi28)
                                    v30 = (v30 ^ 592)
                                } else {
                                    var v32 float64 = 3.14159
                                    v32 = 3.14159
                                    switch fi28 {
                                    case 1:
                                        switch (fi28 >> (((v30 / ((uint(v7)) | 1))) & 63)) {
                                        case 127:
                                            if (("" == "hello") && false) {
                                                var v33 int8 = 42
                                                _ = v33
                                            } else {
                                                var v34 int8 = 42
                                                v32 = (v32 - v32)
                                                _ = v34
                                            }
                                        default:
                                            v22 = (v22 % (((^(v22))) | 1))
                                        }
                                    default:
                                        var v35 uint32 = v24
                                        _ = v35
                                    }
                                    _ = v32
                                }
                            } else {
                                _ = len(v3)
                                var v36 uint32 = (v4 % ((v4) | 1))
                                v6 += 32767
                                _ = v36
                            }
                            _ = fi31
                        }
                        _ = v30
                    }
                    for fi37 := 0; fi37 < 10; fi37++ {
                        v38 := GenS5{F0: (!(!false)), F1: 649, F2: 18446744073709551615}
                        switch 127 {
                        case 2147483647:
                            v39 := GenS6{F0: (v2 & v7), F1: "foo bar", F2: (true && false)}
                            _ = v39
                        case 1:
                            v40 := GenS7{F0: (^(v6)), F1: (v22 * v22), F2: -2147483648}
                            _ = v40
                        default:
                            switch "" {
                            case "hello":
                                if (1 >= 42) {
                                    switch (fi37 >> ((122) & 63)) {
                                    case 2147483647:
                                        v7 -= (v2 - (^(v7)))
                                    case 9223372036854775807:
                                        v41 := GenS8{F0: 1, F1: v2, F2: v7}
                                        _ = v41
                                    case 1:
                                        v3 = append(v3, 2)
                                    }
                                    if false {
                                        fmt.Println(v22, v2, v22)
                                    } else {
                                        var v42 uint = v5
                                        _ = v42
                                    }
                                }
                            default:
                                var v43 int16 = v6
                                _ = v43
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                    v1 = (v1 ^ v1)
                    _ = fi28
                }
                v44 := GenS4{F0: (v4 / (((v4 | 56)) | 1))}
                _ = v44
            }
        }
    }
    if false {
        v45 := GenS5{F0: true, F1: 495, F2: 18446744073709551615}
        v24 = uint32(v5)
        _ = len(v3)
        _ = v45
    } else {
        var v46 int = 32767
        switch v46 {
        case 42:
            v6 += v6
        case -9223372036854775808:
            fmt.Println(v46, v2, v6)
        }
        _ = v46
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        v3 = append(v3, v24)
        switch (fi47 >> ((255) & 63)) {
        case 986:
            var v48 int64 = 1
            _ = v48
        case 1:
            fmt.Println(v7, v2, v7)
        }
        _ = fi47
    }
    var v49 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
