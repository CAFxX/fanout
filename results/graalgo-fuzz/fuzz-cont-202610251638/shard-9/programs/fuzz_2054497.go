package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = -8
    v2 := make(map[uint32]string)
    v3 := []uint32{11, 42, 591}
    var v4 int8 = v1
    var v5 float32 = 0.0
    var v6 uint64 = 4294967295
    v6 = v6
    _ = len(v3)
    v1 -= v4
    v7 := GenS1{F0: v6}
    v8 := GenS1{F0: (v6 % (((v6 >> ((425) & 63))) | 1))}
    var v9 float32 = 3.14159
    if false {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int64 = 549
            _ = fi10
            _ = v11
        }
        switch ("a" + "") {
        case "a":
            if (221 != 18446744073709551615) {
                switch "hello" {
                case "\n\t":
                    if (18446744073709551615 != 336) {
                        switch 32767 {
                        case 82:
                            v9 = (-2.25 + -2.25)
                        case 2147483647:
                            switch -2147483648 {
                            case 4294967295:
                                switch 128 {
                                case -161:
                                    var v12 uint16 = 32767
                                    _ = v12
                                case 798:
                                    _ = len(v3)
                                default:
                                    _ = len(v3)
                                }
                            default:
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    v14 := GenS1{F0: (v6 | (v6 | v6))}
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        }
                    }
                case "hello":
                    v15 := GenS1{F0: v6}
                    _ = v15
                default:
                    v5 += float32(-2.25)
                }
                var v16 string = ""
                _ = v16
            } else {
                fmt.Println(v5)
            }
        case "":
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    v9 = 1.5
                    var v19 uint16 = 42
                    if (((float32(3.14159) * v9) - float32(3.14159)) <= (float32(float64(v9)) - v9)) {
                        fmt.Println(v6, v9)
                    } else {
                        v1 = (v1 * v4)
                        _ = len(v3)
                        v20 := GenS1{F0: v6}
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        default:
            var v21 int8 = (v1 << ((4294967295) & 63))
            _ = v21
        }
    } else {
        var v22 int32 = 2147483647
        if (2147483648 == 422) {
            v22 = v22
            v22 = (v22 & (v22 + 2147483647))
        }
        switch -2147483648 {
        case 51:
            var v23 float64 = 0.0
            _ = v23
        default:
            var v24 uint32 = 2
            _ = v24
        }
        _ = v22
    }
    var v25 int64 = 9223372036854775807
    var v26 float64 = 0.1
    v26 = ((float64(v9) - 1.5) - v26)
    for fi27 := 0; fi27 < 2; fi27++ {
        v25 = (v25 - v25)
        _ = len(v2)
        v25 = v25
        _ = fi27
    }
    var v28 uint16 = 255
    v2[1] = "hello"
    var v29 int8 = (v1 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
