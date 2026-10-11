package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }

func main() {
    var v1 string = ("a" + "a")
    var v2 int = 1
    var v3 float64 = float64(0.1)
    var v4 uint16 = 1
    v5 := GenS1{F0: v1, F1: 42, F2: 93}
    v6 := []int{(v2 ^ v2), v2, (v2 << ((18446744073709551615) & 63))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 8; fi8++ {
        v3 -= v3
        switch "a" {
        case "":
            var v9 int = fi8
            _ = v9
        case "a":
            switch fi8 {
            case -761:
                fmt.Println(v7)
            case 461:
                v1 = v1
            case 171:
                switch fi8 {
                case -609:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        switch v1 {
                        case "foo bar":
                            var v11 int16 = 0
                            _ = v11
                        case "a":
                            var v12 uint32 = 189
                            _ = v12
                        default:
                            if ((((0.1 + v3) + (-2.25 - v3)) - v3) != ((0.0 * float64(1.5)) - v3)) {
                                fmt.Println(v2, v1)
                            }
                        }
                        _ = fi10
                    }
                }
            }
        case "hello":
            var v13 uint64 = v7
            _ = v13
        default:
            var v14 float32 = (1e10 / 1.5)
            _ = v14
        }
        var v15 float32 = (((float32(v3) * 0.1) - 0.1) * float32(v3))
        _ = fi8
        _ = v15
    }
    if true {
        var v16 int8 = -128
        _ = v16
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            if len(v6) > 0 {
                v6[0] = (fi17 * fi17)
            }
            v2 = v2
            if true {
                fmt.Println(fi17, v7)
                var v18 uint16 = v4
                if (-128 <= -1) {
                    var v19 uint8 = 0
                    _ = v19
                }
                _ = v18
            } else {
                v7 += v7
                switch (fi17 * fi17) {
                case -2147483648:
                    var v20 int16 = 1
                    _ = v20
                case -9223372036854775808:
                    var v21 int = (v2 - (-(fi17)))
                    _ = v21
                case -425:
                    var v22 uint32 = 0
                    _ = v22
                }
            }
            _ = fi17
        }
        var v23 int = 0
        _ = v23
    }
    var v24 int64 = 206
    v24 = (v24 / ((-1) | 1))
    v7 = v7
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
