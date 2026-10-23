package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]bool)
    var v3 int = 708
    var v4 string = (("" + "foo bar") + "foo bar")
    switch ("a" + v4) {
    case "a":
        var v5 string = "a"
        _ = v5
    default:
        v4 = v4
    }
    var v6 bool = (1 >= 0)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: -1, F1: 65535, F2: -9223372036854775808}
        _ = fi7
        _ = v8
    }
    v9 := v2[4294967295]
    var v10 int8 = 0
    switch v4 {
    case "":
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := v2[142]
            _ = fi11
            _ = v12
        }
    default:
        v1 = v1
    }
    v4 = v4
    var v13 int32 = v1
    v13 -= (^(v13))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
