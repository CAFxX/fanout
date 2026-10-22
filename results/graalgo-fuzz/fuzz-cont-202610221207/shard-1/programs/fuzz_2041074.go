package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    var v1 int = 2147483647
    var v2 uint64 = 0
    v3 := make(map[string]string)
    v4 := GenS1{F0: (0.1 > (1e10 / 0.1))}
    fmt.Println(v2)
    fmt.Println(v1)
    v2 += v2
    v2 = v2
    if true {
        v1 = (v1 - v1)
        v2 = 4294967295
    }
    if (!(!true)) {
        v5 := GenS1{F0: (2 != 901)}
        var v6 int64 = 9223372036854775807
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            v1 += (v1 | 750)
            _ = fi7
        }
    }
    v8 := GenS2{F0: 42, F1: ((v2 + (v2 ^ (v2 - 1))) != v2), F2: 255}
    var v9 uint32 = 65535
    delete(v3, "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
