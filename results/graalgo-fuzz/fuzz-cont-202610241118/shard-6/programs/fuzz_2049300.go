package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 int16; F1 uint; F2 int8 }

func main() {
    var v1 float32 = float32(0.0)
    var v2 int = 0
    var v3 int8 = 42
    v1 += (v1 - -2.25)
    var v4 int64 = 42
    var v5 int = (v2 << ((970) & 63))
    if true {
        var v6 uint16 = 0
        var v7 uint64 = 42
        _ = v6
        _ = v7
    } else {
        if ((v3 << ((65535) & 63)) <= (v3 * v3)) {
            switch (("a" + "\n\t") + "a") {
            case "foo bar":
                for fi8 := 0; fi8 < 9; fi8++ {
                    fmt.Println(v5)
                    var v9 uint8 = 0
                    v10 := GenS1{F0: (v3 + 127), F1: "hello", F2: 32767}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                var v11 bool = true
                _ = v11
            }
            var v12 float64 = 0.0
            _ = v12
        }
    }
    switch "hello" {
    case "hello":
        v13 := GenS2{F0: -322, F1: (v4 + (v4 % ((v4) | 1)))}
        _ = v13
    }
    var v14 int16 = 292
    switch "a" {
    case "\n\t":
        v15 := GenS3{F0: 190, F1: 9223372036854775807, F2: ((0.1 / (3.14159 * ((v1 * 3.14159) + 1.5))) > v1)}
        _ = v15
    }
    var v16 int16 = v14
    v17 := []int{32768, (v5 ^ v5), (-(v5)), (v2 / ((196) | 1))}
    v1 += 0.1
    var v18 uint32 = 1
    switch (v5 * v2) {
    case 127:
        if (!(("hello" != (("" + (("a" + "foo bar") + "hello")) + ("foo bar" + ("a" + "\n\t")))) && true)) {
            if false {
                if ((-(v16)) == (v14 + (v16 & (v16 & (v14 >> ((1) & 63)))))) {
                    v5 -= (v2 | v5)
                } else {
                    switch "hello" {
                    case "hello":
                        var v19 int32 = 42
                        _ = v19
                    case "\n\t":
                        var v20 float32 = (v1 * 1.5)
                        _ = v20
                    case "a":
                        fmt.Println(v3, v18)
                    }
                    var v21 float32 = float32(3.14159)
                    _ = v21
                }
                var v22 int = 32768
                _ = v22
            } else {
                for fi23 := 0; fi23 < 1; fi23++ {
                    v4 = (^(v4))
                    for fi24 := 0; fi24 < 15; fi24++ {
                        fmt.Println(v2, v4, fi23)
                        v25 := GenS4{F0: (^(v14)), F1: 1, F2: (v3 % ((v3) | 1))}
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
                switch v5 {
                case 128:
                    v3 -= (v3 << ((18446744073709551615) & 63))
                case -237:
                    var v26 uint32 = 65535
                    _ = v26
                default:
                    v4 = (v4 ^ (v4 % ((v4) | 1)))
                }
                var v27 int32 = 2147483647
                _ = v27
            }
            var v28 int32 = 1
            _ = v28
        }
    case 1:
        _ = len(v17)
    default:
        v14 = 188
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
