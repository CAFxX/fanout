package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 int16 = 1
    var v5 uint32 = 694
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v4, v5)
        fmt.Println(v5)
        var v7 float64 = 0.1
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 -= (v5 | 2)
        _ = fi8
    }
    var v9 float32 = 0.0
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10, fi10, fi10)
        _ = fi10
    }
    var v11 uint64 = 18446744073709551615
    for fi12 := 0; fi12 < 2; fi12++ {
        switch ("a" + "foo bar") {
        case "":
            if (!(!(fi12 != 0))) {
                v13 := GenS1{F0: (v4 << ((18446744073709551615) & 63)), F1: 9223372036854775807}
                _ = v13
            }
        case "a":
            v14 := v2[(v5 ^ v5)]
            _ = v14
        }
        var v15 int = -9223372036854775808
        _ = fi12
        _ = v15
    }
    if (1 < 30) {
        v16 := GenS1{F0: (v4 | v4), F1: -1}
        _ = v16
    } else {
        if true {
            fmt.Println(v9)
            fmt.Println(v9)
        } else {
            var v17 int64 = -9223372036854775808
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
