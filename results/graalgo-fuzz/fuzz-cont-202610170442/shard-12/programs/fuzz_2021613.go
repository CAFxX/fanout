package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 int64 }

func main() {
    var v1 uint = 255
    v2 := []string{"", "a"}
    var v3 uint = (v1 + 4294967295)
    v4 := make(map[uint32]bool)
    var v5 int = 2147483647
    var v6 uint64 = 4294967295
    v3 -= v3
    v7 := GenS1{F0: v1}
    switch "hello" {
    case "\n\t":
        fmt.Println(v6)
    }
    v1 -= v3
    fmt.Println(v6, v6)
    if (false || ((false && false) && true)) {
        var v8 int = v5
        v6 = (v6 << (((^(v1))) & 63))
        var v9 float64 = float64(-2.25)
        _ = v8
        _ = v9
    } else {
        switch "" {
        case "\n\t":
            v10 := GenS2{F0: 255, F1: (v1 * (v3 << ((18446744073709551615) & 63))), F2: -526}
            _ = v10
        case "a":
            v1 = uint(v6)
        default:
            switch "a" {
            case "hello":
                v6 = v6
            case "\n\t":
                fmt.Println(v5, v5)
            default:
                if (v1 >= 1) {
                    if false {
                        var v11 int8 = -33
                        v12 := v4[2]
                        _ = v11
                        _ = v12
                    }
                    if (483 >= 65535) {
                        if false {
                            fmt.Println(v3)
                            v1 = v3
                        } else {
                            for fi13 := 0; fi13 < 3; fi13++ {
                                switch (("" + ("\n\t" + "\n\t")) + "\n\t") {
                                case "a":
                                    _ = len(v4)
                                }
                                for fi14 := 0; fi14 < 9; fi14++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi14
                                }
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    v6 = 4294967295
                                    v16 := GenS2{F0: 2, F1: 743, F2: -1}
                                    _ = fi15
                                    _ = v16
                                }
                                _ = fi13
                            }
                            v1 = uint(v6)
                            var v17 float64 = float64(3.14159)
                            _ = v17
                        }
                    } else {
                        fmt.Println(v6, v1, v3)
                        v6 = 123
                        v18 := GenS2{F0: 128, F1: v3, F2: 707}
                        _ = v18
                    }
                } else {
                    if (!true) {
                        switch (-(v5)) {
                        case -9223372036854775808:
                            v19 := GenS3{F0: -832, F1: (81 <= 100), F2: 1}
                            _ = v19
                        default:
                            fmt.Println(v1, v1)
                        }
                        var v20 uint32 = 42
                        _ = v20
                    }
                    v6 += v6
                }
            }
        }
        if true {
            v1 = v1
            var v21 uint32 = 42
            v2 = append(v2, "")
            _ = v21
        }
    }
    var v22 int64 = -902
    if (!true) {
        var v23 uint = (v1 / ((v3) | 1))
        delete(v4, 205)
        v5 += int(v1)
        _ = v23
    } else {
        var v24 uint16 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v6 -= (v6 & 0)
        v1 += (^(v1))
        _ = fi25
    }
    v4[4294967295] = (((false || true) && (344 != 2)) && true)
    v22 -= v22
    var v26 uint32 = uint32(v6)
    v1 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
