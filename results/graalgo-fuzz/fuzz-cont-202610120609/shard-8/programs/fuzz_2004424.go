package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int16 = -93
    var v2 int16 = v1
    var v3 string = ("" + "foo bar")
    var v4 uint32 = 365
    v5 := []int{128, 0}
    v3 = "a"
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 8, F1: -32768}
    v9 := GenS2{F0: 32768}
    var v10 uint16 = 42
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 9; fi11++ {
            v4 = v4
            v12 := GenS3{F0: 9223372036854775807}
            v13 := GenS1{F0: 42, F1: 42}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint16 = v10
        _ = fi14
        _ = v15
    }
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
    _ = v8
    _ = v9
}
