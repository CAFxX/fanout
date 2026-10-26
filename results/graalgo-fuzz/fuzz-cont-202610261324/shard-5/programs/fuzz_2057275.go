package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int16 = 32767
    var v3 int32 = 1
    var v4 float32 = v1
    v5 := make(map[string]string)
    v6 := []int{}
    if false {
        v2 = (v2 << ((792) & 63))
        var v7 int64 = 42
        var v8 float64 = -2.25
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int32 = (v3 - (v3 * (v3 % ((v3) | 1))))
            var v12 string = "foo bar"
            switch v12 {
            case "foo bar":
                var v13 int = (fi10 % (((fi9 / ((2147483647) | 1))) | 1))
                _ = v13
            }
            v14 := GenS1{F0: 128, F1: 0, F2: -1}
            v3 += -2147483648
            _ = fi10
            _ = v11
            _ = v12
            _ = v14
        }
        var v15 int8 = 127
        var v16 string = "\n\t"
        _ = fi9
        _ = v15
        _ = v16
    }
    switch -412 {
    case 9223372036854775807:
        fmt.Println(v4, v4)
    case -1:
        var v17 int64 = 42
        _ = v17
    }
    var v18 uint32 = 42
    var v19 float64 = 0.1
    var v20 bool = true
    switch "a" {
    case "a":
        var v21 int16 = (-(v2))
        _ = v21
    case "\n\t":
        var v22 int32 = v3
        _ = v22
    case "foo bar":
        fmt.Println(v3, v19)
    default:
        switch -2147483648 {
        case 2147483647:
            var v23 int = 0
            _ = v23
        }
    }
    v24 := GenS2{F0: (v2 & v2), F1: (v3 ^ v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
