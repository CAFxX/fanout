package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 string; F2 int16 }

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 uint = 16
    var v3 uint8 = 42
    v4 := make(map[int]uint32)
    var v5 float64 = (((1e10 - -2.25) * 1e10) - 1.5)
    if false {
        v4[4294967295] = 42
        for fi6 := 0; fi6 < 5; fi6++ {
            if false {
                var v7 float32 = (0.0 - (3.14159 * 0.0))
                _ = v7
            } else {
                v1 = (("foo bar" + v1) + v1)
                v8 := GenS1{F0: 885}
                _ = v8
            }
            _ = fi6
        }
    } else {
        var v9 int16 = -1
        fmt.Println(v2, v2)
        v10 := GenS2{F0: 42, F1: ("hello" + ""), F2: (v9 * (v9 & (v9 << ((v2) & 63))))}
        _ = v10
        _ = v9
    }
    var v11 int16 = -482
    v12 := GenS2{F0: 9223372036854775807, F1: v1, F2: (v11 << ((v2) & 63))}
    switch (v1 + (v1 + ("\n\t" + ("" + v1)))) {
    case "a":
        var v13 uint = 1
        _ = v13
    case "\n\t":
        fmt.Println(v5, v2, v3)
    case "hello":
        v14 := GenS1{F0: 9223372036854775807}
        _ = v14
    }
    v5 = (1.5 / v5)
    var v15 uint16 = 2
    _ = len(v4)
    v15 = v15
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
