package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 127
    var v3 int8 = int8(v2)
    var v4 float32 = float32(0.1)
    v3 += v3
    v5 := []int{32767, -9223372036854775808, 32768, 2147483647}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int16 = (v1 | -32768)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        switch 127 {
        case 681:
            v9 := GenS1{F0: v2}
            _ = v9
        case 1:
            v10 := GenS1{F0: v2}
            _ = v10
        }
        _ = fi8
    }
    var v11 uint8 = v2
    v12 := GenS2{F0: 9223372036854775807, F1: 929, F2: -395}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(fi14, fi14)
            v4 = v4
            _ = fi14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint = 1
        _ = fi15
        _ = v16
    }
    v1 += (v1 ^ int16(v3))
    for fi17 := 0; fi17 < 15; fi17++ {
        switch "" {
        case "\n\t":
            v18 := GenS1{F0: 238}
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 4; fi19++ {
                switch (fi19 >> ((uint(v2)) & 63)) {
                case -387:
                    switch "a" {
                    case "hello":
                        var v20 uint = 1
                        _ = v20
                    case "foo bar":
                        v21 := GenS3{F0: 110, F1: fi19, F2: -2147483648}
                        _ = v21
                    default:
                        v11 += (v2 ^ 0)
                    }
                }
                var v22 int64 = 42
                switch ("a" + "") {
                case "":
                    var v23 int32 = 1
                    _ = v23
                case "\n\t":
                    v22 = (v22 % (((v22 / (((v22 | (v22 & v22))) | 1))) | 1))
                case "hello":
                    v24 := GenS1{F0: 64}
                    _ = v24
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v3 = (v3 >> ((uint(v2)) & 63))
                _ = fi19
                _ = v22
            }
        case "hello":
            var v25 int = (-(fi17))
            _ = v25
        default:
            v26 := GenS3{F0: 65535, F1: (fi17 * 4294967295), F2: 606}
            _ = v26
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
