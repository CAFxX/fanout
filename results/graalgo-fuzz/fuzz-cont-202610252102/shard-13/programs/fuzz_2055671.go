package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{""}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 int = 32768
    var v6 uint8 = 186
    if (!true) {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := v3["\n\t"]
            if true {
                v9 := GenS1{F0: -1, F1: -19, F2: 828}
                _ = v9
            }
            v1 += v1
            _ = fi7
            _ = v8
        }
        _ = len(v2)
    }
    switch -53 {
    case -225:
        if len(v2) > 0 {
            v2[0] = ""
        }
    case 13:
        v10 := GenS1{F0: 42, F1: -496, F2: int16(v6)}
        _ = v10
    default:
        v1 -= 1
    }
    v4[2147483648] = 2147483647
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = -9223372036854775808
        switch fi12 {
        case 660:
            if (!true) {
                switch "" {
                case "hello":
                    var v14 uint8 = v6
                    _ = v14
                }
            }
        case 968:
            v5 = (fi12 << (((v1 >> ((v1) & 63))) & 63))
        default:
            var v15 string = "foo bar"
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := v4[(v11 + 766)]
    fmt.Println(v6)
    v6 = v6
    if (1000 > -517) {
        var v17 int8 = -126
        var v18 bool = true
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
