package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 int16, p1 string, p2 string) uint64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int64 = -9223372036854775808
        var v6 uint = 65535
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 969
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := []uint32{532}
    var v9 uint8 = 128
    v10 := GenS1{F0: (true && true)}
    switch 42 {
    case 0:
        v11 := GenS2{F0: -32768, F1: 1000}
        _ = v11
    default:
        v9 = v9
    }
    v12 := v1[-260]
    var v13 uint64 = 1
    v12 = (3.14159 <= 1.5)
    var v14 string = ""
    if ((genFn0(-118, v14, "\n\t") != (v13 ^ 4294967295)) && (v12 && (3.14159 == (1e10 * 1e10)))) {
        switch -2147483648 {
        case 127:
            switch v14 {
            case "hello":
                v9 -= (v9 % ((v9) | 1))
            default:
                var v15 int16 = -32768
                _ = v15
            }
        default:
            if (!(0.1 != 1e10)) {
                fmt.Println(v12, v13, v13)
            }
        }
    }
    fmt.Println(v13, v12)
    var v16 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
