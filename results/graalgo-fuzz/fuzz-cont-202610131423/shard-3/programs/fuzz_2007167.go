package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    v1 := []uint32{388}
    var v2 uint = 42
    var v3 uint = 42
    var v4 bool = (!(-32768 != -1))
    var v5 int = 121
    v3 = v3
    for fi6 := 0; fi6 < 8; fi6++ {
        switch fi6 {
        case -1:
            v7 := GenS1{F0: -128}
            _ = v7
        case 127:
            var v8 int64 = 997
            _ = v8
        case 1:
            if v4 {
                fmt.Println(v2, fi6, v3)
                var v9 uint = (v3 % ((v2) | 1))
                v5 = (fi6 - 127)
                _ = v9
            }
        default:
            var v10 int = 0
            _ = v10
        }
        v3 = v3
        _ = fi6
    }
    switch (("hello" + ("" + "")) + "a") {
    case "a":
        fmt.Println(v5, v5)
    default:
        var v11 int32 = 0
        _ = v11
    }
    var v12 uint32 = 681
    v2 = 4294967295
    if false {
        var v13 uint8 = 37
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint64 = 1
            var v16 uint8 = v13
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v4 = (2147483647 > -1)
        _ = fi17
    }
    v4 = ((v3 / ((767) | 1)) != v3)
    v1 = append(v1, 957)
    v12 = (^(v12))
    v12 = 2147483647
    _ = len(v1)
    v18 := GenS2{F0: (^(v12)), F1: v5}
    fmt.Println(v3)
    v3 = (v3 & 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
