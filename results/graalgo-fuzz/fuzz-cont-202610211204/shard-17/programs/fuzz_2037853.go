package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 float64 = -2.25
    v2 := []string{"a", "\n\t", "\n\t"}
    var v3 float32 = 1e10
    var v4 uint64 = 42
    var v5 int64 = 1
    var v6 int16 = 0
    switch "\n\t" {
    case "\n\t":
        var v7 uint32 = 4294967295
        _ = v7
    case "hello":
        v8 := GenS1{F0: 128, F1: (!false)}
        _ = v8
    case "a":
        fmt.Println(v5, v6, v4)
    default:
        fmt.Println(v6, v1, v4)
    }
    if false {
        fmt.Println(v3)
        switch 0 {
        case 13:
            if false {
                v5 = (v5 + v5)
            } else {
                fmt.Println(v4, v4, v3)
            }
        case 783:
            fmt.Println(v3, v1, v4)
        case 2147483647:
            switch 1 {
            case 128:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                var v9 uint32 = 0
                _ = v9
            case 127:
                switch 32768 {
                case 982:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -2147483648:
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if true {
                            v5 = 0
                            fmt.Println(v4, v1, v6)
                            for fi11 := 0; fi11 < 7; fi11++ {
                                var v12 float64 = (((0.0 + 0.1) - 1e10) / (float64(v3) + 0.1))
                                var v13 uint16 = 255
                                v1 = float64(3.14159)
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        }
                        v14 := GenS2{F0: "\n\t", F1: (-106 != 31), F2: 0}
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v4 = 42
                            fmt.Println(fi15)
                            _ = fi15
                        }
                        _ = fi10
                        _ = v14
                    }
                default:
                    var v16 int16 = 787
                    _ = v16
                }
            default:
                var v17 uint8 = 2
                _ = v17
            }
        default:
            var v18 float32 = 3.14159
            _ = v18
        }
    }
    v1 = (v1 - 1.5)
    if (-101 <= 127) {
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 uint = 0
                v5 = v5
                _ = fi19
                _ = v20
            }
            v5 = 1
            _ = len(v2)
        } else {
            var v21 uint32 = 2
            _ = len(v2)
            v22 := GenS3{F0: 18446744073709551615}
            _ = v21
            _ = v22
        }
        v5 = (v5 / (((^(v5))) | 1))
    }
    var v23 uint32 = uint32(v4)
    var v24 bool = false
    var v25 uint16 = 719
    var v26 int64 = v5
    var v27 int32 = -146
    v23 -= (^(v23))
    var v28 bool = (false && v24)
    fmt.Println(v24, v6, v3)
    if v28 {
        for fi29 := 0; fi29 < 14; fi29++ {
            v24 = (v28 && ((false && true) && ((((v23 | (v23 >> ((1) & 63))) != v23) && v24) && (v25 == (v25 << ((2) & 63))))))
            v24 = (100 <= -90)
            _ = fi29
        }
        var v30 uint8 = 216
        _ = v30
    } else {
        v31 := GenS3{F0: 940}
        _ = v31
    }
    v32 := GenS1{F0: 191, F1: v24}
    var v33 int32 = v27
    v34 := GenS2{F0: (("a" + "foo bar") + ""), F1: v24, F2: (^(v25))}
    if v24 {
        var v35 int32 = v27
        switch "foo bar" {
        case "a":
            fmt.Println(v5, v6)
        case "\n\t":
            if v24 {
                switch "\n\t" {
                case "\n\t":
                    for fi36 := 0; fi36 < 15; fi36++ {
                        for fi37 := 0; fi37 < 20; fi37++ {
                            _ = len(v2)
                            _ = fi37
                        }
                        _ = fi36
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                default:
                    switch 42 {
                    case -455:
                        v6 = v6
                    default:
                        v2 = append(v2, "hello")
                    }
                }
                var v38 float32 = float32(v1)
                _ = v38
            } else {
                for fi39 := 0; fi39 < 11; fi39++ {
                    v28 = v24
                    v40 := GenS2{F0: "foo bar", F1: (v25 >= 2), F2: 255}
                    if v24 {
                        for fi41 := 0; fi41 < 14; fi41++ {
                            v42 := GenS2{F0: "", F1: (v24 && false), F2: uint16(v23)}
                            _ = fi41
                            _ = v42
                        }
                        v2 = append(v2, "\n\t")
                    }
                    _ = fi39
                    _ = v40
                }
                var v43 string = "\n\t"
                _ = v43
            }
        case "foo bar":
            var v44 uint16 = 1
            _ = v44
        }
        v45 := GenS4{F0: (^(v25)), F1: v4}
        _ = v35
        _ = v45
    }
    if v24 {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v46 := GenS2{F0: ("a" + ""), F1: v24, F2: (v25 & 669)}
            _ = v46
        case "a":
            v26 -= (v5 % ((9223372036854775807) | 1))
        }
        v5 = (v5 ^ -1)
        if (1 > 24) {
            v23 = (v23 - v23)
        }
    } else {
        fmt.Println(v23, v23, v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
