package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    v1 := []uint32{65535, 42, 2}
    v2 := []uint32{4294967295, 742, 2}
    v3 := make(map[string]uint32)
    var v4 uint8 = 155
    var v5 int8 = -48
    var v6 float32 = 1.5
    var v7 uint64 = 2
    v8 := GenS1{F0: 2147483647, F1: v5}
    var v9 uint16 = 85
    if len(v1) > 0 {
        v1[0] = 2
    }
    if ((v4 == (v4 ^ 0)) || (false || false)) {
        fmt.Println(v4, v6, v9)
        v5 = (v5 | (-(v5)))
        v5 = -1
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            if ((v5 - v5) < v5) {
                switch fi11 {
                case 0:
                    var v12 uint8 = v4
                    _ = v12
                case -2147483648:
                    var v13 float64 = (float64(v6) * 1e10)
                    _ = v13
                case -727:
                    fmt.Println(v7, fi10, fi11)
                    v14 := GenS1{F0: 65535, F1: v5}
                    v3["a"] = 2147483648
                    _ = v14
                }
                var v15 int8 = (v5 - v5)
                _ = v15
            }
            v16 := GenS2{F0: (v4 % (((v4 & (v4 % ((v4) | 1)))) | 1)), F1: (false || (v6 < v6))}
            _ = fi11
            _ = v16
        }
        var v17 uint16 = v9
        _ = fi10
        _ = v17
    }
    if (((v9 << ((2) & 63)) >= (v9 << ((505) & 63))) || true) {
        var v18 uint = 42
        var v19 int64 = 9223372036854775807
        _ = v18
        _ = v19
    } else {
        v3[((("\n\t" + "foo bar") + "\n\t") + ("hello" + ("hello" + "a")))] = 2147483647
        if false {
            v4 += v4
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
