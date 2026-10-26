package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int8; F2 uint }

func main() {
    v1 := []uint32{794, 282}
    var v2 uint32 = 1
    var v3 uint16 = 366
    v4 := make(map[uint32]bool)
    var v5 uint32 = (v2 ^ (v2 % (((v2 >> ((42) & 63))) | 1)))
    v3 -= 0
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = 2147483647
        switch "a" {
        case "foo bar":
            for fi7 := 0; fi7 < 10; fi7++ {
                if true {
                    if false {
                        fmt.Println(fi7, v5, v5)
                    }
                    var v8 uint32 = (^(v2))
                    _ = v8
                } else {
                    if true {
                        var v9 uint16 = 428
                        _ = v9
                    } else {
                        v3 = (v3 >> ((uint(v5)) & 63))
                        var v10 float64 = 1e10
                        _ = v10
                    }
                }
                _ = fi7
            }
        case "hello":
            v11 := GenS1{F0: 1000, F1: 42, F2: false}
            _ = v11
        }
        _ = fi6
    }
    switch 299 {
    case 9223372036854775807:
        var v12 float32 = (1.5 - ((0.0 + (0.0 + 1.5)) + 0.1))
        _ = v12
    }
    if ((1 > 33) && true) {
        v4[(^(v5))] = true
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS2{F0: (v5 + v2)}
            _ = fi14
            _ = v15
        }
        fmt.Println(fi13)
        switch (fi13 % (((fi13 - fi13)) | 1)) {
        case 32767:
            fmt.Println(v5, v3, v5)
        case -713:
            switch 127 {
            case 9223372036854775807:
                if false {
                    switch (fi13 / ((fi13) | 1)) {
                    case 0:
                        fmt.Println(v3, v3, v3)
                    case 997:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        fmt.Println(fi13)
                    }
                    v16 := GenS3{F0: 0}
                    _ = v16
                } else {
                    var v17 int8 = 0
                    var v18 int = fi13
                    _ = v17
                    _ = v18
                }
            case 127:
                v3 = (v3 / ((v3) | 1))
            case 42:
                for fi19 := 0; fi19 < 9; fi19++ {
                    v2 = v5
                    v5 -= (v2 + v5)
                    _ = fi19
                }
            default:
                v20 := GenS2{F0: 2147483648}
                _ = v20
            }
        }
        _ = fi13
    }
    var v21 int64 = -129
    fmt.Println(v2, v3, v3)
    var v22 uint16 = 867
    fmt.Println(v22)
    var v23 bool = ((177 == 2) || false)
    v24 := GenS4{F0: 127, F1: 34, F2: 505}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v25 uint8 = uint8(v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
