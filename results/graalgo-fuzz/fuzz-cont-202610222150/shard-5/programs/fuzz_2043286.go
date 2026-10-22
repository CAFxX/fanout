package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
type GenS5 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 float32 = 1.5
    var v2 int64 = -1
    v3 := make(map[uint32]string)
    v4 := []string{"", ("foo bar" + "foo bar"), "", "a"}
    fmt.Println(v1, v1)
    if (v1 != 1.5) {
        if (0 != 42) {
            var v5 int64 = 424
            v6 := GenS1{F0: 1, F1: 196, F2: 0}
            v7 := GenS2{F0: 1, F1: 1, F2: 42}
            _ = v5
            _ = v6
            _ = v7
        } else {
            v1 = 0.1
            v2 = (v2 >> ((65535) & 63))
            v8 := GenS3{F0: 65535, F1: 100}
            _ = v8
        }
    }
    var v9 uint16 = 32767
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := v3[1]
        var v12 uint = 3
        fmt.Println(fi10, v2, v1)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int32 = 42
    switch 128 {
    case 4294967295:
        v14 := GenS1{F0: uint8(v9), F1: 1, F2: 892}
        _ = v14
    case 128:
        switch 32767 {
        case 829:
            fmt.Println(v13)
        case 42:
            v9 = v9
        case -876:
            var v15 uint16 = v9
            _ = v15
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                if ((fi16 % ((1) | 1)) <= 9223372036854775807) {
                    v17 := GenS2{F0: 32767, F1: -160, F2: 32767}
                    _ = v17
                }
                var v18 int64 = (v2 + 1)
                fmt.Println(fi16, v9, v2)
                _ = fi16
                _ = v18
            }
        }
    case 42:
        var v19 float32 = v1
        _ = v19
    default:
        fmt.Println(v1)
    }
    var v20 uint = 4294967295
    switch 1 {
    case -2147483648:
        var v21 uint16 = (v9 ^ 713)
        _ = v21
    }
    fmt.Println(v9, v1, v20)
    v22 := GenS4{F0: (v20 / (((v20 / (((v20 ^ v20)) | 1))) | 1)), F1: uint64(v9), F2: false}
    delete(v3, 2147483647)
    v23 := GenS4{F0: (v20 & 18446744073709551615), F1: 769, F2: false}
    v2 -= -272
    if true {
        for fi24 := 0; fi24 < 3; fi24++ {
            _ = len(v3)
            if (42 == 1) {
                for fi25 := 0; fi25 < 6; fi25++ {
                    if true {
                        _ = len(v3)
                        var v26 bool = true
                        _ = v26
                    }
                    if (((1.5 * (1.5 / v1)) / v1) > float32(1e10)) {
                        var v27 float32 = v1
                        v27 = (v27 - ((float32((((0.1 * (0.1 * (float64(v1) / 1.0))) + float64(v27)) * 0.0)) * float32(0.0)) / (0.1 * -2.25)))
                        v9 = (^(v9))
                        _ = v27
                    } else {
                        var v28 string = (("hello" + ("foo bar" + "foo bar")) + "\n\t")
                        _ = v28
                    }
                    v29 := GenS1{F0: 1, F1: 0, F2: v20}
                    if false {
                        _ = len(v4)
                        if false {
                            v3[uint32(v20)] = "\n\t"
                        }
                        var v30 uint = v20
                        _ = v30
                    }
                    var v31 uint64 = 510
                    v32 := GenS5{F0: 18446744073709551615, F1: 255}
                    delete(v3, 917)
                    _ = fi25
                    _ = v29
                    _ = v31
                    _ = v32
                }
            } else {
                v13 -= (v13 ^ (v13 / ((-363) | 1)))
            }
            v33 := v3[1]
            _ = fi24
            _ = v33
        }
        var v34 int32 = (v13 - v13)
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
