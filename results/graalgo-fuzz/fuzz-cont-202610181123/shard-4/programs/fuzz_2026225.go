package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{"a", (("" + "") + ("hello" + "")), "foo bar", ("" + "hello")}
    var v3 float64 = 1.5
    var v4 uint16 = 42
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    v6[0] = "a"
    v7 := GenS1{F0: 1}
    if true {
        v3 = v3
    } else {
        fmt.Println(v4)
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS2{F0: uint16(v1)}
        v4 = 331
        fmt.Println(v1, v4, v4)
        _ = fi8
        _ = v9
    }
    fmt.Println(v3, v4, v3)
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch -9223372036854775808 {
        case 9223372036854775807:
            v1 -= (v1 - v1)
        }
    }
    var v10 float64 = v3
    switch 0 {
    case -9223372036854775808:
        v11 := GenS3{F0: true, F1: int8(v1), F2: 18446744073709551615}
        _ = v11
    case 127:
        if false {
            var v12 int32 = 0
            _ = v12
        } else {
            switch 2147483647 {
            case 53:
                v13 := GenS4{F0: v1}
                _ = v13
            case -1:
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        if false {
                            var v16 int32 = 42
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + "\n\t")
                            }
                            _ = v16
                        }
                        _ = fi15
                    }
                    v3 += float64((float32((v3 / ((v10 * v3) * v10))) - (((1e10 * (0.0 + ((float32((v10 / v10)) / 1.0) * 0.0))) * -2.25) * 1.5)))
                    v17 := GenS4{F0: 0}
                    _ = fi14
                    _ = v17
                }
            default:
                v10 += v10
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                v10 -= v10
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS1{F0: 205}
                    _ = len(v6)
                    switch "a" {
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            var v21 float32 = 3.14159
                            _ = v21
                        default:
                            var v22 uint32 = (v1 - (v1 | v1))
                            _ = v22
                        }
                    default:
                        if (((2147483647 <= 1) && (false && false)) && true) {
                            v4 = (v4 % ((659) | 1))
                            var v23 float32 = (0.0 - 1e10)
                            switch -1 {
                            case 2147483647:
                                if true {
                                    var v24 uint8 = 1
                                    var v25 uint64 = 4294967295
                                    _ = v24
                                    _ = v25
                                }
                            }
                            _ = v23
                        } else {
                            fmt.Println(fi19)
                            var v26 string = ("a" + ("hello" + ("foo bar" + "foo bar")))
                            var v27 uint16 = (^(v4))
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
    case -1:
        for fi28 := 0; fi28 < 17; fi28++ {
            _ = len(v5)
            v29 := v5[("a" + ("hello" + ("hello" + "a")))]
            _ = fi28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v10, v1)
        var v31 uint32 = (v1 + (v1 - (^(v1))))
        _ = fi30
        _ = v31
    }
    v32 := GenS4{F0: 225}
    switch "a" {
    case "\n\t":
        v33 := GenS5{F0: 42, F1: 28}
        _ = v33
    case "foo bar":
        var v34 uint8 = 127
        _ = v34
    case "a":
        fmt.Println(v10, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
