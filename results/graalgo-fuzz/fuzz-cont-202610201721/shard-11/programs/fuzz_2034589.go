package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 int16 = 1
    var v2 int = -603
    var v3 string = "hello"
    var v4 int8 = 127
    v1 -= 177
    var v5 uint64 = 287
    v6 := []int{73, (-(v2))}
    v7 := GenS1{F0: -2147483648, F1: true, F2: 42}
    v8 := GenS2{F0: 1}
    v5 -= (v5 << ((621) & 63))
    var v9 int32 = 979
    for fi10 := 0; fi10 < 13; fi10++ {
        v6 = append(v6, (fi10 % (((v2 >> ((18446744073709551615) & 63))) | 1)))
        if (true && (false && false)) {
            v11 := GenS2{F0: (v9 + v9)}
            fmt.Println(v5, v2, v5)
            v1 = (v1 / (((v1 * v1)) | 1))
            _ = v11
        }
        _ = fi10
    }
    var v12 string = "a"
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint = 839
        if (!false) {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint64 = v5
                _ = fi15
                _ = v16
            }
            if (65535 <= 760) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 string = v3
                    var v19 bool = true
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int16 = (v1 * (v1 - (v1 | v1)))
                _ = v20
            } else {
                v21 := GenS3{F0: 9223372036854775807}
                var v22 bool = (0 == 1)
                fmt.Println(v3, v4)
                _ = v21
                _ = v22
            }
        }
        var v23 int8 = v4
        _ = fi13
        _ = v14
        _ = v23
    }
    var v24 int16 = (v1 % ((v1) | 1))
    v4 = (v4 % (((v4 | 14)) | 1))
    fmt.Println(v5, v24, v5)
    switch ((v3 + v12) + "a") {
    case "hello":
        v25 := GenS1{F0: 0, F1: true, F2: 357}
        _ = v25
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS3{F0: 42}
            if false {
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 int32 = int32(v4)
                    if ((true && (!false)) || (1 > 559)) {
                        var v30 int = (fi26 - int(v4))
                        v31 := GenS4{F0: v12}
                        switch -9223372036854775808 {
                        case -668:
                            v24 -= v24
                        default:
                            _ = len(v6)
                        }
                        var v32 int64 = -39
                        _ = v30
                        _ = v31
                        _ = v32
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v9 -= (v29 & v9)
                    _ = fi28
                    _ = v29
                }
            } else {
                v24 = int16(v5)
            }
            _ = fi26
            _ = v27
        }
    }
    var v33 int32 = (v9 ^ v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
