package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 string = ((("hello" + "a") + "hello") + "foo bar")
    var v2 string = v1
    var v3 int = -9223372036854775808
    v4 := []string{"\n\t"}
    var v5 uint8 = 2
    var v6 int8 = -1
    var v7 uint32 = 379
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint16 = 0
        var v10 int64 = 625
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    v6 = -92
    v11 := GenS1{F0: 1}
    switch (v3 & 1) {
    case 32767:
        v4 = append(v4, v2)
    case -9223372036854775808:
        v3 += -1
    case -685:
        var v12 int64 = 42
        _ = v12
    }
    v7 += 2147483648
    v13 := GenS1{F0: (v5 << ((326) & 63))}
    v3 = (v3 + (v3 * v3))
    var v14 float64 = 1e10
    v7 = v7
    for fi15 := 0; fi15 < 12; fi15++ {
        if (true && true) {
            v16 := GenS2{F0: v1, F1: (v5 ^ (^(v5)))}
            var v17 bool = (true && true)
            _ = v16
            _ = v17
        }
        v6 = (v6 % ((v6) | 1))
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
