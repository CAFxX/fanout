package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }

func main() {
    var v1 int16 = -468
    var v2 uint16 = 2
    var v3 int8 = -128
    var v4 string = "foo bar"
    switch v4 {
    case "":
        var v5 int16 = (v1 - -684)
        _ = v5
    case "\n\t":
        fmt.Println(v4, v2)
    }
    v6 := GenS1{F0: (v3 ^ (v3 >> ((2) & 63))), F1: 42}
    v7 := GenS2{F0: (^(v2)), F1: 42, F2: 2147483647}
    var v8 int = 746
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{}
        switch fi9 {
        case 0:
            fmt.Println(v2, v2)
        case 911:
            v1 += v1
        case -808:
            var v11 uint = 255
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS2{F0: v2, F1: -9223372036854775808, F2: -550}
    v3 = v3
    switch v8 {
    case 128:
        v13 := GenS1{F0: 0, F1: -1}
        _ = v13
    case 1:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
