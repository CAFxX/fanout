package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{}
    var v3 int16 = 32767
    var v4 int16 = v3
    if true {
        fmt.Println(v1, v4)
        var v5 int32 = 1
        _ = v5
    } else {
        v4 = v3
    }
    var v6 string = "hello"
    if (386 != 0) {
        v6 = "foo bar"
        var v7 uint = 1
        _ = v7
    } else {
        v8 := GenS1{F0: 2147483647, F1: 0}
        _ = v8
    }
    var v9 float32 = ((((v1 - v1) / v1) + v1) + ((0.0 - 3.14159) / 1.5))
    v10 := GenS1{F0: 65535, F1: 442}
    _ = len(v2)
    var v11 uint64 = 365
    v12 := GenS2{F0: 2, F1: -109}
    v13 := GenS2{F0: 127, F1: -2}
    switch ("\n\t" + "") {
    case "hello":
        var v14 float32 = 1e10
        _ = v14
    default:
        v6 = v6
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 591, F1: 602}
        var v17 float64 = float64((-2.25 / float32(-2.25)))
        switch v6 {
        case "a":
            v9 -= v9
        case "":
            var v18 uint64 = v11
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
