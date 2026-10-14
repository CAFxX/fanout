package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -682
    var v3 uint16 = 1
    var v4 uint16 = (v3 / ((775) | 1))
    v5 := GenS1{F0: -295}
    var v6 int8 = 42
    v6 = v6
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := v1[649]
        fmt.Println(v2, v6, v8)
        v9 := GenS1{F0: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v6)
    v10 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: 0, F2: 0}
    switch (v2 % (((v2 ^ 32768)) | 1)) {
    case 127:
        v4 = (v3 % ((834) | 1))
    default:
        var v11 int64 = -166
        _ = v11
    }
    v2 = v2
    switch "" {
    case "hello":
        fmt.Println(v6)
    }
    var v12 bool = (!true)
    v3 = v4
    var v13 bool = v12
    v12 = (((v4 * 423) != 32767) || (!v13))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
