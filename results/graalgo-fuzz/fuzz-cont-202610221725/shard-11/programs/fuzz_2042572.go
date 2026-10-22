package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 925
    v3 := make(map[string]bool)
    var v4 int16 = (v2 + v2)
    var v5 string = "a"
    var v6 float32 = 1e10
    v7 := v3[(("\n\t" + "foo bar") + v5)]
    fmt.Println(v4, v2)
    var v8 int = 127
    v7 = v7
    v9 := GenS1{F0: 4294967295, F1: v4, F2: -237}
    fmt.Println(v4)
    if ((v4 > 518) && (!false)) {
        switch v5 {
        case "foo bar":
            var v10 int64 = -1
            _ = v10
        case "":
            v8 -= v8
        default:
            switch (v8 ^ 0) {
            case -966:
                var v11 uint64 = 1
                _ = v11
            case -94:
                v12 := GenS1{F0: 0, F1: (-(v4)), F2: v8}
                _ = v12
            }
        }
    }
    var v13 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
