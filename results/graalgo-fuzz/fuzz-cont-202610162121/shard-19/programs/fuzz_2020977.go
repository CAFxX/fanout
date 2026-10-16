package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 uint16 = 2
    v5 := GenS1{F0: v1}
    var v6 uint = 609
    var v7 int = 2147483647
    if (-850 != 688) {
        v1 = "foo bar"
    } else {
        v4 = v4
    }
    v7 = v7
    v8 := GenS2{F0: v4, F1: -9223372036854775808}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: ""}
        fmt.Println(v7, v6)
        for fi11 := 0; fi11 < 12; fi11++ {
            v7 -= fi9
            fmt.Println(v4, v1)
            v4 += (v4 + (^(v4)))
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
