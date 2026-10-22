package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []uint32{0, 946, 65535, 42}
    var v2 int16 = 1000
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 int64 = 1
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 13; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v8 int = -2147483648
        v9 := GenS1{F0: 1, F1: 111, F2: v5}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = (-(v5))
    var v11 int64 = (v10 + int64(v6))
    fmt.Println(v6)
    var v12 int8 = 0
    v6 += v6
    _ = len(v4)
    if ((v10 / (((^(v5))) | 1)) >= (v11 % ((0) | 1))) {
        v5 -= (^(v11))
        switch -9223372036854775808 {
        case 127:
            v13 := v3[("\n\t" + "\n\t")]
            _ = v13
        case -9223372036854775808:
            _ = len(v4)
        case -640:
            var v14 uint32 = 413
            _ = v14
        }
    } else {
        if (v6 > v6) {
            fmt.Println(v11, v2)
            fmt.Println(v11)
            if (true || ((v12 < 100) && true)) {
                if (!false) {
                    v6 = (v6 + (v6 | v6))
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v15 float32 = (0.0 * 0.0)
                _ = v15
            }
        } else {
            var v16 uint8 = 2
            fmt.Println(v16, v6, v2)
            fmt.Println(v5)
            _ = v16
        }
        v17 := GenS1{F0: (v12 << (((v6 >> ((v6) & 63))) & 63)), F1: 1, F2: v10}
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        var v18 uint32 = 1
        _ = v18
    default:
        v19 := GenS2{F0: 344, F1: (v10 * -1)}
        _ = v19
    }
    var v20 uint8 = 79
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
