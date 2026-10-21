package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint32 = 1
    var v2 bool = false
    v3 := []uint32{v1}
    v4 := []uint32{(v1 - (v1 << ((598) & 63)))}
    v2 = (2 <= 2)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: (^(v1)), F1: -1}
        if true {
            v2 = (998 != 1)
            v1 = (v1 & (^(v1)))
        }
        _ = fi5
        _ = v6
    }
    var v7 int64 = 9223372036854775807
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v4)
    v8 := GenS2{F0: 100, F1: 103, F2: 2}
    v2 = v2
    v7 -= 42
    fmt.Println(v1)
    switch 128 {
    case 862:
        var v9 int16 = -82
        _ = v9
    case 0:
        if (1e10 >= (float32(3.14159) / float32(float64(3.14159)))) {
            for fi10 := 0; fi10 < 15; fi10++ {
                if len(v4) > 0 {
                    v4[0] = 206
                }
                fmt.Println(v1, v7, v1)
                fmt.Println(v7, v1, v7)
                _ = fi10
            }
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2 = true
    switch ("hello" + "") {
    case "hello":
        v3 = append(v3, 2147483647)
    default:
        v2 = false
    }
    if (!false) {
        var v11 int = 140
        if (((1 >= 1) && (255 != 65535)) && (v1 >= (v1 - (v1 * (v1 >> ((1) & 63)))))) {
            var v12 uint16 = 663
            _ = v12
        } else {
            v13 := GenS1{F0: v1, F1: v7}
            fmt.Println(v11, v7, v2)
            fmt.Println(v1, v2)
            _ = v13
        }
        v7 += -1
        _ = v11
    } else {
        var v14 int = 4294967295
        v3 = append(v3, (v1 ^ (v1 - 65535)))
        v14 = -661
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS3{F0: (1 <= -1)}
            if true {
                v7 -= (v7 >> ((4294967295) & 63))
                var v18 uint64 = 0
                if (false && v2) {
                    var v19 int32 = 42
                    _ = v19
                } else {
                    v20 := GenS2{F0: 127, F1: 126, F2: 128}
                    v7 += (v7 / (((v7 ^ -234)) | 1))
                    _ = v20
                }
                _ = v18
            } else {
                v21 := GenS3{F0: (false || v2)}
                var v22 int16 = -32768
                v23 := GenS3{F0: (v2 && v2)}
                _ = v21
                _ = v22
                _ = v23
            }
            switch "a" {
            case "a":
                switch "\n\t" {
                case "hello":
                    switch (fi16 + 42) {
                    case -1:
                        v2 = (!(!(431 > 714)))
                    default:
                        v24 := GenS4{F0: 393, F1: 1}
                        _ = v24
                    }
                case "":
                    v25 := GenS1{F0: (v1 / ((v1) | 1)), F1: v7}
                    _ = v25
                case "foo bar":
                    v26 := GenS5{F0: -1}
                    _ = v26
                }
            case "foo bar":
                var v27 int32 = 2147483647
                _ = v27
            case "hello":
                switch "\n\t" {
                case "hello":
                    var v28 int = fi15
                    _ = v28
                case "\n\t":
                    var v29 uint16 = 65535
                    _ = v29
                default:
                    if v2 {
                        var v30 uint64 = 42
                        v31 := GenS6{F0: 42}
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int16 = 42
                        _ = v32
                    }
                }
            default:
                v1 = (v1 << ((42) & 63))
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    switch "a" {
    case "\n\t":
        if true {
            if len(v4) > 0 {
                v4[0] = 355
            }
            v1 -= v1
        } else {
            v7 = (v7 ^ 42)
            v4 = append(v4, (v1 ^ (v1 << ((65535) & 63))))
            if ((18446744073709551615 != 42) || (v2 || (2 != 90))) {
                v2 = v2
            }
        }
    default:
        if v2 {
            v1 += (v1 & (^(v1)))
        }
    }
    if (v2 || v2) {
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS7{F0: (v1 & (v1 / (((v1 | (v1 & 957))) | 1))), F1: -1}
            v7 = (v7 >> ((255) & 63))
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = fi33
            _ = v34
        }
        v2 = (v1 >= (v1 >> ((73) & 63)))
        v35 := GenS3{F0: v2}
        _ = v35
    } else {
        var v36 int32 = -1
        v2 = false
        v36 -= v36
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
