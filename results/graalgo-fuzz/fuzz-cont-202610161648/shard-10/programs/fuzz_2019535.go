package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    var v1 bool = true
    var v2 int16 = -165
    v3 := []uint32{}
    switch "foo bar" {
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v4 := GenS1{F0: 42}
        _ = v4
    }
    if v1 {
        var v5 uint16 = 554
        var v6 int16 = (v2 % ((v2) | 1))
        v7 := GenS2{F0: 4294967295, F1: 843}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS3{F0: 2147483647, F1: 428}
    var v9 string = ("a" + ("\n\t" + "\n\t"))
    var v10 int16 = (^(v2))
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v2, v1)
        _ = fi11
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    v3 = append(v3, 42)
    v2 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
