package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint32 = 65535
    var v2 uint32 = v1
    var v3 uint8 = 42
    var v4 int = 9223372036854775807
    v5 := []string{((((((("\n\t" + "foo bar") + "hello") + "hello") + "a") + ((((("hello" + "foo bar") + "a") + "foo bar") + "\n\t") + "\n\t")) + "hello") + "hello")}
    var v6 int16 = -996
    var v7 int8 = 0
    v7 -= 39
    switch 127 {
    case 492:
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 += 65535
            var v9 uint = 65535
            var v10 int16 = v6
            _ = fi8
            _ = v10
            _ = v9
        }
    case -1:
        var v11 uint32 = (v2 | v2)
        _ = v11
    }
    v12 := GenS1{F0: (v2 << ((2) & 63)), F1: (v3 / ((255) | 1)), F2: (v6 * v6)}
    var v13 bool = ((2147483647 == 519) || false)
    var v14 uint64 = uint64(v3)
    var v15 uint16 = 0
    if v13 {
        v13 = true
    }
    v16 := GenS2{F0: (v2 >> ((42) & 63)), F1: v3}
    fmt.Println(v15, v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
