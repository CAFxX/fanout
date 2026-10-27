package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 int32 = 672
    var v3 uint32 = 2147483647
    v4 := []uint32{}
    var v5 int32 = -1
    v2 = v5
    var v6 string = ""
    v7 := GenS1{F0: 202}
    v3 = v3
    if false {
        v8 := GenS2{F0: 2147483647, F1: -9223372036854775808, F2: 1}
        if len(v4) > 0 {
            v4[0] = v3
        }
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int = 4294967295
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v11 int = (v10 ^ v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    v3 += (v3 + v3)
    v12 := GenS3{F0: -1, F1: 958}
    if (v1 > (v1 / (((v1 % (((v1 >> ((4294967295) & 63))) | 1))) | 1))) {
        switch "hello" {
        case "foo bar":
            if true {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS4{F0: (v1 & v1), F1: (((0 != (v5 << ((378) & 63))) || false) || true)}
                    switch v6 {
                    case "foo bar":
                        fmt.Println(v5, v6, fi13)
                    case "\n\t":
                        if (128 > 2) {
                            for fi15 := 0; fi15 < 13; fi15++ {
                                var v16 string = "foo bar"
                                _ = fi15
                                _ = v16
                            }
                        }
                    case "hello":
                        v1 = 38
                    default:
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v4 = append(v4, 0)
                            if ((v6 < ("hello" + (v6 + v6))) && (!true)) {
                                var v18 bool = (!true)
                                v1 -= (v1 >> ((255) & 63))
                                v19 := GenS3{F0: -630, F1: 255}
                                _ = v18
                                _ = v19
                            } else {
                                fmt.Println(v6, fi17)
                            }
                            if true {
                                fmt.Println(v5, v1)
                            }
                            _ = fi17
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                v2 -= 42
                v2 = (v5 * -2147483648)
            } else {
                switch "hello" {
                case "hello":
                    var v20 uint32 = (v3 % ((v3) | 1))
                    _ = v20
                case "a":
                    var v21 float32 = 1.5
                    _ = v21
                }
                v22 := GenS3{F0: -159, F1: 32767}
                v1 = v1
                _ = v22
            }
        case "a":
            if (!(100 < 127)) {
                if true {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if (v1 > v1) {
                            v2 = (v5 + (v5 - -2147483648))
                            v3 = (v3 >> ((uint(v3)) & 63))
                        } else {
                            v3 = v3
                            fmt.Println(v6)
                        }
                        fmt.Println(v5, fi23)
                        if (1 <= 65535) {
                            v24 := GenS2{F0: (v2 * (-(v2))), F1: 127, F2: 1}
                            _ = len(v4)
                            var v25 int64 = -9223372036854775808
                            _ = v24
                            _ = v25
                        } else {
                            v3 = (^(v3))
                            v3 = 301
                        }
                        _ = fi23
                    }
                }
            }
        }
        v1 -= v1
    }
    _ = len(v4)
    _ = len(v4)
    var v26 uint64 = 864
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
