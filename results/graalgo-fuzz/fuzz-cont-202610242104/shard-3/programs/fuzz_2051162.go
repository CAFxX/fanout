package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    v4 := []uint32{2}
    v3[int(v1)] = ("\n\t" + ("\n\t" + "\n\t"))
    switch "\n\t" {
    case "a":
        v1 = (v1 | (v1 % ((v1) | 1)))
    }
    delete(v3, 127)
    v1 = 42
    v5 := GenS1{F0: v1, F1: 42, F2: 788}
    var v6 float64 = 0.1
    v1 = (^(v1))
    var v7 float64 = 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6)
    var v8 int16 = -1
    for fi9 := 0; fi9 < 12; fi9++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v8 = (v8 % (((-(v8))) | 1))
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := v3[(fi9 & (fi10 ^ (^(fi10))))]
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    delete(v3, 9223372036854775807)
    var v12 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
