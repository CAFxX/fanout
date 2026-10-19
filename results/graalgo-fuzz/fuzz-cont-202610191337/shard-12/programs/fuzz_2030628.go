package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []uint32{384, 4294967295}
    v2 := make(map[string]int)
    var v3 int8 = 0
    var v4 int16 = -32768
    var v5 int = 1
    v5 = (v5 & (-(v5)))
    var v6 uint16 = 956
    if (4294967295 == 206) {
        v2[""] = (v5 + 2147483647)
        switch "hello" {
        case "a":
            switch v5 {
            case 1:
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 float64 = (((float64((1.5 + (0.0 - 1e10))) - -2.25) - -2.25) * (0.0 * 0.0))
                    _ = fi7
                    _ = v8
                }
            case 138:
                switch "hello" {
                case "":
                    v9 := GenS1{F0: 2, F1: (-327 < -1), F2: int(v6)}
                    _ = v9
                case "hello":
                    v3 = 86
                default:
                    v6 = 1
                }
            case 2147483647:
                var v10 uint8 = 46
                _ = v10
            }
        case "hello":
            v5 = v5
        default:
            switch "" {
            case "hello":
                v3 += (v3 | v3)
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v5 = 127
            }
        }
    } else {
        fmt.Println(v6)
        fmt.Println(v4)
    }
    v11 := v2["a"]
    switch v11 {
    case 128:
        v5 -= 226
    }
    fmt.Println(v4, v5)
    v11 += (v5 - v5)
    _ = len(v2)
    var v12 uint16 = v6
    v13 := v2[(("a" + "\n\t") + ("\n\t" + "a"))]
    v6 = (v6 * (v12 % (((v12 ^ v12)) | 1)))
    v5 = (v5 | (^(v13)))
    v14 := GenS2{F0: uint32(v12)}
    v3 -= 1
    if false {
        switch v13 {
        case -1:
            for fi15 := 0; fi15 < 18; fi15++ {
                fmt.Println(v12, v13, v6)
                _ = fi15
            }
        case 2147483647:
            v6 = v6
        case 9223372036854775807:
            switch (v5 ^ (v11 ^ v5)) {
            case -889:
                for fi16 := 0; fi16 < 2; fi16++ {
                    v3 += 42
                    _ = fi16
                }
            }
        default:
            fmt.Println(v11, v12)
        }
        fmt.Println(v13)
    }
    v6 = 660
    var v17 float32 = 0.0
    v4 -= (-(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
