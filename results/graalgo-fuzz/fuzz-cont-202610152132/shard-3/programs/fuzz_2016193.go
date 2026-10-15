package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 int32; F2 int32 }
type GenS7 struct { F0 int; F1 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{4294967295, 1}
    var v3 int16 = 32767
    var v4 uint16 = 409
    v5 := make(map[uint32]bool)
    var v6 int8 = 100
    switch 32767 {
    case 42:
        if ((727 == (v4 >> ((0) & 63))) && true) {
            var v7 float32 = float32(3.14159)
            _ = len(v5)
            var v8 int16 = 32767
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: 18446744073709551615, F2: v1}
            var v10 int16 = (v3 % ((-964) | 1))
            _ = v10
            _ = v9
        }
    case 9223372036854775807:
        fmt.Println(v4)
    case 32768:
        var v11 bool = (-1 >= 42)
        _ = v11
    }
    switch ("hello" + ("\n\t" + "a")) {
    case "":
        fmt.Println(v3)
    }
    if true {
        v12 := GenS1{F0: 882, F1: v1, F2: (v1 / (((^(v1))) | 1))}
        v13 := GenS1{F0: 998, F1: 392, F2: 18446744073709551615}
        if false {
            v14 := GenS1{F0: 30, F1: v1, F2: (v1 / ((v1) | 1))}
            if ((("foo bar" + "") + ("hello" + "a")) < "foo bar") {
                var v15 uint32 = 42
                v16 := GenS2{F0: 241, F1: (v4 % ((uint16(v15)) | 1)), F2: (v6 - 100)}
                if (975 > 0) {
                    var v17 int32 = 42
                    _ = v17
                } else {
                    _ = len(v2)
                }
                _ = v15
                _ = v16
            } else {
                var v18 float32 = float32(1.5)
                switch 32768 {
                case 868:
                    if false {
                        delete(v5, 777)
                        var v19 uint32 = 2147483647
                        var v20 int64 = int64(v19)
                        _ = v19
                        _ = v20
                    }
                case 32767:
                    switch 127 {
                    case 467:
                        v3 -= int16(v6)
                    default:
                        fmt.Println(v6)
                    }
                }
                _ = v18
            }
            _ = v14
        }
        _ = v12
        _ = v13
    } else {
        if true {
            if true {
                fmt.Println(v3)
                if (v4 == v4) {
                    v3 += 42
                    if false {
                        var v21 string = (("hello" + "a") + ("hello" + "hello"))
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            fmt.Println(v6, v6)
                            _ = fi22
                        }
                    }
                } else {
                    v23 := GenS3{F0: (-(v6))}
                    v24 := GenS4{F0: 430, F1: -2147483648}
                    _ = v23
                    _ = v24
                }
                v6 = v6
            }
            v4 += (v4 | (v4 >> ((42) & 63)))
        }
    }
    if false {
        v25 := GenS5{F0: false}
        _ = v25
    }
    var v26 int = 32767
    delete(v5, 512)
    fmt.Println(v26, v1, v6)
    var v27 string = "foo bar"
    var v28 int32 = -1
    v5[42] = (65535 >= uint(v1))
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS6{F0: 369, F1: (^(v28)), F2: (v28 & v28)}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    var v32 int16 = v3
    v2 = append(v2, 746)
    if (int16(v28) < int16(v6)) {
        if true {
            if (-2147483648 >= (^(v28))) {
                if ((1e10 >= 1.5) || false) {
                    for fi33 := 0; fi33 < 17; fi33++ {
                        switch (v26 + (fi33 ^ v26)) {
                        case -788:
                            _ = len(v2)
                        case 32767:
                            if ((v6 - (v6 << ((uint(v4)) & 63))) >= 0) {
                                v4 = v4
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v34 uint16 = (v4 & v4)
                                v3 += (v32 | 32767)
                                fmt.Println(v28, fi33)
                                _ = v34
                            }
                        }
                        _ = fi33
                    }
                } else {
                    v4 = 1
                }
                v35 := GenS5{F0: (!true)}
                var v36 float32 = 0.1
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int32 = (v28 ^ 0)
                v1 -= (^(v1))
                var v39 float32 = (-2.25 - 1.5)
                _ = fi37
                _ = v38
                _ = v39
            }
        }
    } else {
        v27 = (("foo bar" + ((("a" + "foo bar") + (v27 + ("a" + v27))) + v27)) + v27)
    }
    var v40 uint8 = 127
    switch v26 {
    case 32768:
        v3 = v32
    default:
        var v41 string = "\n\t"
        _ = v41
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        v40 = 255
        v43 := GenS6{F0: 2147483648, F1: v28, F2: v28}
        if (v27 <= "foo bar") {
            fmt.Println(v3, v3, v32)
            for fi44 := 0; fi44 < 4; fi44++ {
                switch v27 {
                case "":
                    v40 += v40
                case "hello":
                    var v45 string = (v27 + v27)
                    _ = v45
                default:
                    v26 = fi42
                }
                v46 := GenS7{F0: 597, F1: -968}
                _ = fi44
                _ = v46
            }
            v6 -= (v6 + v6)
        } else {
            if false {
                if false {
                    switch fi42 {
                    case -29:
                        if true {
                            v1 = v1
                            fmt.Println(v40)
                            switch v26 {
                            case 1:
                                _ = len(v5)
                            case 127:
                                if true {
                                    v40 -= (^(v40))
                                }
                            case 566:
                                var v47 float32 = 0.0
                                _ = v47
                            }
                        }
                    default:
                        var v48 uint32 = uint32(v40)
                        _ = v48
                    }
                } else {
                    fmt.Println(v40, v28)
                }
                for fi49 := 0; fi49 < 5; fi49++ {
                    if false {
                        v26 = v26
                    }
                    var v50 uint64 = 1
                    _ = fi49
                    _ = v50
                }
            } else {
                v5[613] = true
            }
            switch v27 {
            case "a":
                if (!false) {
                    switch fi42 {
                    case 32768:
                        for fi51 := 0; fi51 < 12; fi51++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v52 uint32 = 65535
                            v32 -= (v32 / ((v32) | 1))
                            v2 = append(v2, v52)
                            fmt.Println(v52, fi42, fi42)
                            _ = fi51
                            _ = v52
                        }
                    }
                    v28 += v28
                } else {
                    v2 = append(v2, 2147483647)
                }
            }
        }
        _ = fi42
        _ = v43
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        switch v27 {
        case "hello":
            var v54 uint8 = (v40 / ((v40) | 1))
            _ = v54
        case "foo bar":
            v55 := GenS1{F0: 2, F1: 18446744073709551615, F2: (v1 >> ((0) & 63))}
            _ = v55
        default:
            for fi56 := 0; fi56 < 4; fi56++ {
                var v57 uint = 2
                v32 = v3
                _ = fi56
                _ = v57
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi58 := 0; fi58 < 18; fi58++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v59 bool = false
            _ = len(v2)
            _ = fi58
            _ = v59
        }
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
