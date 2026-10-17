package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }

func main() {
    var v1 uint8 = 128
    v2 := []uint32{0}
    var v3 int = -9223372036854775808
    var v4 int64 = 0
    var v5 float64 = ((1e10 - -2.25) * 0.1)
    v4 += v4
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 += (v6 | v6)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = 493
        }
        _ = fi7
    }
    var v8 int = v3
    v3 = v8
    for fi9 := 0; fi9 < 14; fi9++ {
        switch ("\n\t" + "\n\t") {
        case "":
            v3 -= fi9
        case "\n\t":
            v10 := GenS1{F0: 608, F1: 2147483647, F2: v6}
            _ = v10
        default:
            var v11 uint8 = (v1 - (v1 >> ((65535) & 63)))
            _ = v11
        }
        v2 = append(v2, 736)
        v12 := GenS1{F0: 42, F1: 1, F2: -1}
        _ = fi9
        _ = v12
    }
    v1 += (v1 % ((v1) | 1))
    v5 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
