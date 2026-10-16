package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{0, 42, 2147483648}
    v3 := []uint32{}
    var v4 uint32 = 2147483648
    var v5 bool = false
    var v6 uint16 = 551
    v7 := GenS1{F0: (v4 + (^(v4)))}
    if true {
        fmt.Println(v4)
        fmt.Println(v6)
    } else {
        if v5 {
            fmt.Println(v5, v4)
        } else {
            fmt.Println(v6, v6)
        }
        if (!(-9223372036854775808 == -185)) {
            fmt.Println(v5, v6)
        } else {
            v8 := GenS1{F0: v4}
            fmt.Println(v6)
            _ = v8
        }
    }
    v9 := GenS1{F0: (v4 >> ((18446744073709551615) & 63))}
    v6 += v6
    var v10 string = ""
    v11 := GenS1{F0: 2147483647}
    fmt.Println(v10)
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(fi12, v10)
        v10 = "foo bar"
        _ = fi12
    }
    if len(v2) > 0 {
        v2[0] = 931
    }
    switch 128 {
    case 127:
        var v13 uint8 = 0
        _ = v13
    case 9223372036854775807:
        switch v10 {
        case "a":
            v14 := GenS1{F0: (v4 / (((v4 >> ((255) & 63))) | 1))}
            _ = v14
        case "hello":
            switch "" {
            case "hello":
                var v15 uint8 = 0
                _ = v15
            case "foo bar":
                var v16 int32 = 1
                _ = v16
            case "":
                _ = len(v2)
            default:
                v17 := GenS1{F0: (^(v4))}
                _ = v17
            }
        case "foo bar":
            if ((0.0 / ((1.5 * (((((((-2.25 + float32((3.14159 * ((1.5 + 3.14159) / 1e10)))) / (-2.25 - 3.14159)) * 1e10) - 0.0) - 0.0) * (3.14159 / (float32(1e10) * 1.5))) * 0.1)) + 0.0)) != -2.25) {
                fmt.Println(v4, v6)
            }
        }
    case -1:
        if v5 {
            var v18 uint64 = 2
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
