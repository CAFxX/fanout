package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 int32 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    v4 := make(map[string]string)
    var v5 bool = false
    switch "a" {
    case "\n\t":
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: (v1 + ((v1 + v1) + v1))}
            var v8 int32 = 1
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v1 = "foo bar"
    v9 := GenS1{F0: -1, F1: 2147483647, F2: ("foo bar" + (((("a" + "foo bar") + ((v1 + "a") + v1)) + "foo bar") + v1))}
    _ = len(v2)
    v10 := GenS1{F0: 0, F1: 2, F2: "foo bar"}
    v5 = v5
    v11 := GenS2{F0: -1}
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint = (v3 >> ((v3) & 63))
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
