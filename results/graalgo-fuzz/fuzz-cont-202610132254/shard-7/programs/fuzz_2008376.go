package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }

func main() {
    var v1 int64 = -1
    v2 := []string{"foo bar", ("foo bar" + "a"), ""}
    var v3 uint64 = 4294967295
    v4 := make(map[string]bool)
    var v5 uint32 = 2
    var v6 uint16 = 459
    if false {
        v5 -= v5
    } else {
        var v7 float32 = 3.14159
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 = 65535
        var v9 uint = 65535
        v1 += (v1 + v1)
        _ = fi8
        _ = v9
    }
    switch 2147483647 {
    case 0:
        v6 = 428
    case 127:
        fmt.Println(v3)
    }
    v6 -= (^(v6))
    v5 = v5
    v5 = v5
    v10 := GenS1{F0: (false || true), F1: (v3 + (v3 % ((v3) | 1))), F2: 927}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
