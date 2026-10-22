package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 0
    var v3 int64 = 1
    var v4 int16 = v2
    var v5 uint = 1
    v5 = (^(v5))
    var v6 int16 = 42
    v3 = -295
    v7 := []int{-495, 128, -322, 9223372036854775807}
    for fi8 := 0; fi8 < 6; fi8++ {
        v2 = v4
        _ = fi8
    }
    switch "\n\t" {
    case "foo bar":
        v9 := v1[42]
        _ = v9
    case "\n\t":
        var v10 float64 = 3.14159
        _ = v10
    default:
        v4 = (v6 >> ((128) & 63))
    }
    var v11 uint64 = 4294967295
    var v12 uint32 = 2147483648
    var v13 int64 = v3
    v14 := GenS1{F0: uint16(v11), F1: ((896 < -985) || false)}
    var v15 int64 = (v3 << (((^(v5))) & 63))
    var v16 string = ("a" + "foo bar")
    v13 -= (v3 + (v3 / ((-439) | 1)))
    for fi17 := 0; fi17 < 13; fi17++ {
        if len(v7) > 0 {
            v7[0] = (^(fi17))
        }
        v11 = 18446744073709551615
        if false {
            for fi18 := 0; fi18 < 12; fi18++ {
                if false {
                    if true {
                        v11 -= v11
                        v3 += (v13 % ((v3) | 1))
                        var v19 float32 = (0.0 - (0.0 + 0.1))
                        if true {
                            fmt.Println(v3, v12, v5)
                            switch fi18 {
                            case 0:
                                fmt.Println(v5, fi17)
                            case 128:
                                switch v16 {
                                case "a":
                                    v11 += (^(v11))
                                    v13 = (v13 * 0)
                                case "foo bar":
                                    var v20 int = (fi17 - 97)
                                    _ = v20
                                default:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                            case 9223372036854775807:
                                var v21 int8 = 99
                                _ = v21
                            default:
                                var v22 int8 = 0
                                _ = v22
                            }
                            if ((-2.25 + v19) == -2.25) {
                                v1[uint32(v5)] = false
                            }
                        } else {
                            v2 = -359
                        }
                        v13 -= v13
                        v19 += ((v19 * 1.5) - 1e10)
                        _ = v19
                    }
                    v23 := GenS1{F0: 0, F1: false}
                    _ = v23
                }
                _ = fi18
            }
            var v24 uint = (v5 + 2)
            v25 := GenS1{F0: 997, F1: false}
            _ = v24
            _ = v25
        }
        _ = fi17
    }
    var v26 float32 = (1e10 - 0.1)
    switch -68 {
    case 42:
        fmt.Println(v6)
    case -9223372036854775808:
        v11 += uint64(v12)
    case -167:
        v4 = (v6 | v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
