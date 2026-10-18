package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int8 = -66
    var v2 int8 = -64
    v3 := make(map[int]string)
    var v4 int64 = 205
    var v5 bool = (float64(float32((0.1 - 0.1))) < 1.5)
    var v6 string = "a"
    switch "\n\t" {
    case "a":
        var v7 uint64 = 286
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            v1 += (v2 * (v2 & v1))
            _ = fi8
        }
    case "":
        fmt.Println(v1, v6, v2)
    default:
        var v9 uint32 = 42
        _ = v9
    }
    v4 = (^(v4))
    v4 += v4
    var v10 uint = 65535
    var v11 int16 = -32768
    if v5 {
        v6 = ("a" + ("foo bar" + "a"))
        var v12 string = (((v6 + ((("" + "\n\t") + "\n\t") + v6)) + v6) + (v6 + "hello"))
        var v13 uint8 = 255
        _ = v12
        _ = v13
    }
    if ((((v6 + v6) + "\n\t") + v6) >= v6) {
        for fi14 := 0; fi14 < 1; fi14++ {
            fmt.Println(fi14, v10, v6)
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint16 = 0
                switch v6 {
                case "foo bar":
                    v4 = -9223372036854775808
                default:
                    v16 = v16
                }
                var v17 uint32 = 4294967295
                v18 := GenS1{F0: (fi14 & (fi15 % ((4294967295) | 1))), F1: v16, F2: (v6 + ((("hello" + "foo bar") + v6) + ""))}
                if v5 {
                    v19 := GenS2{F0: (^(fi14))}
                    _ = v19
                }
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
        v11 = int16(v10)
        v5 = true
    } else {
        v20 := []int{267, 127, 0, -428}
        _ = v20
    }
    var v21 int64 = (v4 / ((9223372036854775807) | 1))
    v11 += (v11 & 0)
    v22 := GenS2{F0: 2147483647}
    v23 := GenS3{F0: (v6 + v6), F1: v11, F2: v1}
    switch 32768 {
    case -2147483648:
        for fi24 := 0; fi24 < 17; fi24++ {
            v10 = 603
            switch "\n\t" {
            case "a":
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 bool = (!v5)
                    v27 := GenS4{F0: 42}
                    fmt.Println(v11, fi24, v4)
                    switch "hello" {
                    case "hello":
                        v5 = v26
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            v28 := GenS1{F0: (fi24 >> ((18446744073709551615) & 63)), F1: 0, F2: "foo bar"}
            _ = fi24
            _ = v28
        }
    case 1:
        v29 := []int{-520, -9223372036854775808, 185}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
