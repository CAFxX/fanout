package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }

func main() {
    var v1 uint8 = 255
    var v2 float64 = 1.5
    var v3 float32 = (0.1 + (float32((0.1 * 1e10)) + 1e10))
    var v4 int8 = 82
    v3 = v3
    v5 := []int{2147483647}
    fmt.Println(v1)
    fmt.Println(v4)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: (-(fi6)), F1: -1, F2: 888}
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "hello":
        var v8 int = 32767
        _ = v8
    case "foo bar":
        switch 32767 {
        case 128:
            v3 = v3
        case -1:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 0:
            var v9 int8 = (^(v4))
            _ = v9
        }
    }
    var v10 string = ""
    fmt.Println(v2, v3, v3)
    fmt.Println(v2, v1, v1)
    v4 = (v4 << ((105) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
