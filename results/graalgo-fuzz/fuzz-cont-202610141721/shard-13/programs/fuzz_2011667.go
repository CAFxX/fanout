package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{}
    var v2 int8 = 42
    var v3 uint8 = 141
    var v4 uint16 = 65535
    v5 := make(map[uint32]bool)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: v4}
    switch 42 {
    case 2147483647:
        v8 := GenS1{F0: (v4 | v4)}
        _ = v8
    }
    v9 := v6[9223372036854775807]
    v10 := GenS2{F0: 9223372036854775807, F1: v9, F2: 2}
    var v11 uint8 = 42
    fmt.Println(v4, v3)
    switch 1 {
    case -1:
        delete(v6, 42)
    }
    if false {
        var v12 int64 = 0
        fmt.Println(v9, v3, v4)
        _ = v12
    } else {
        if len(v1) > 0 {
            v1[0] = 3
        }
    }
    var v13 int32 = 42
    var v14 float64 = (1.5 / ((3.14159 / 0.1) - float64(1.5)))
    v15 := GenS2{F0: -1, F1: 2147483647, F2: (v11 + 0)}
    v11 += 32
    switch ("hello" + ("\n\t" + "a")) {
    case "foo bar":
        v14 = (v14 / (1.5 - ((v14 * 1.5) - ((v14 * ((((v14 / 0.1) / 1e10) / v14) * v14)) + (v14 * 0.1)))))
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 655
        }
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        fmt.Println(v4, v9)
    }
    _ = len(v1)
    v16 := GenS3{F0: 431, F1: 1000}
    if false {
        v11 += (^(v3))
        for fi17 := 0; fi17 < 1; fi17++ {
            switch fi17 {
            case 0:
                if ((v13 - (v13 - v13)) == 1) {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        if false {
                            v9 -= v9
                            var v19 uint64 = 157
                            v19 += (^(v19))
                            _ = v19
                        }
                        v20 := GenS4{F0: v4}
                        _ = fi18
                        _ = v20
                    }
                }
            case 128:
                fmt.Println(v9, v2, v13)
            }
            v21 := GenS2{F0: -9223372036854775808, F1: (v9 / ((269) | 1)), F2: (v11 / (((v3 >> ((255) & 63))) | 1))}
            _ = fi17
            _ = v21
        }
    } else {
        v4 = 828
        switch "foo bar" {
        case "hello":
            fmt.Println(v11, v9, v4)
        case "foo bar":
            var v22 int8 = -1
            _ = v22
        }
    }
    v4 -= (^(v4))
    fmt.Println(v2, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
