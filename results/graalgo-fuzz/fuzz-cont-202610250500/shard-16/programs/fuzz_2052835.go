package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint16 = 552
    var v2 uint8 = 127
    var v3 int8 = -1
    var v4 int = 32768
    var v5 string = ("hello" + "\n\t")
    v6 := []int{}
    _ = len(v6)
    v2 = v2
    v7 := GenS1{F0: -1, F1: (^(v3))}
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v8 uint = 2
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 bool = ((v5 <= v5) || false)
        switch v4 {
        case 4294967295:
            var v11 int = -9223372036854775808
            _ = v11
        case -334:
            fmt.Println(v1, v5)
        case 62:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            for fi12 := 0; fi12 < 7; fi12++ {
                v10 = v10
                v13 := GenS2{F0: (v2 / ((v2) | 1)), F1: (-1 < -975)}
                v14 := GenS1{F0: -1, F1: (^(v3))}
                v6 = append(v6, (fi9 ^ v4))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    var v15 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
