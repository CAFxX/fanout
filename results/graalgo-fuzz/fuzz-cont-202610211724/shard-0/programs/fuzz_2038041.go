package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 uint16 = 0
    var v4 uint16 = 65535
    var v5 float32 = ((1e10 - float32((((-2.25 + 0.0) / 1.0) * float64(1e10)))) / 1.0)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 42, F1: (fi6 | (fi6 * -2147483648)), F2: 904}
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 += (v1 % ((4294967295) | 1))
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    var v9 string = "a"
    if (534 != -85) {
        v1 -= (v1 % ((v1) | 1))
    }
    var v10 int = 2147483647
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(fi11, v9, v3)
            v13 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: 42}
            _ = fi12
            _ = v13
        }
        v2[v9] = fi11
        v14 := GenS1{F0: 1, F1: (v10 & fi11), F2: -9223372036854775808}
        _ = fi11
        _ = v14
    }
    var v15 int = (-(v10))
    v16 := GenS1{F0: 0, F1: 636, F2: int64(v1)}
    v17 := GenS1{F0: -330, F1: (v15 ^ -889), F2: -545}
    fmt.Println(v5, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
