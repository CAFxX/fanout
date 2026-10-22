package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }

func main() {
    v1 := []int{}
    var v2 uint64 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    var v5 int = 2147483647
    for fi6 := 0; fi6 < 16; fi6++ {
        switch (-(fi6)) {
        case 4294967295:
            switch (fi6 & (v5 + v5)) {
            case 21:
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 bool = ("a" <= (("" + "foo bar") + "hello"))
                    v9 := GenS1{F0: 42, F1: 755}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        case 128:
            v10 := v4[("" + "a")]
            _ = v10
        case -2147483648:
            v11 := GenS2{F0: 1, F1: 669, F2: -70}
            _ = v11
        }
        if len(v1) > 0 {
            v1[0] = (^(v5))
        }
        fmt.Println(v2, v2, fi6)
        _ = fi6
    }
    fmt.Println(v2, v2)
    var v12 bool = ((0.1 * 3.14159) != 0.0)
    v13 := GenS2{F0: 1, F1: 0, F2: 127}
    v1 = append(v1, 127)
    v2 = (v2 ^ (v2 | v2))
    fmt.Println(v5, v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
