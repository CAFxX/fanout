package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := []uint32{2, 123, 1}
    v3 := []string{}
    var v4 uint64 = 855
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "hello"
        var v7 int16 = 32767
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v1, v4)
    v4 += v4
    fmt.Println(v4)
    var v8 int8 = -1
    var v9 uint64 = (v4 * (^(v4)))
    fmt.Println(v8)
    var v10 int16 = 1000
    fmt.Println(v8, v10, v9)
    v11 := GenS1{F0: v1, F1: 42, F2: 117}
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            _ = len(v2)
            v10 = (^(v10))
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
    case "":
        var v14 int8 = v8
        _ = v14
    case "foo bar":
        var v15 bool = (!true)
        _ = v15
    }
    if ((!false) || true) {
        var v16 int = 0
        v9 = 0
        _ = v16
    }
    var v17 int32 = 42
    v18 := GenS1{F0: (v1 ^ (v1 >> ((590) & 63))), F1: v1, F2: (^(v1))}
    var v19 float32 = 1.5
    v9 = v4
    fmt.Println(v4, v9, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
