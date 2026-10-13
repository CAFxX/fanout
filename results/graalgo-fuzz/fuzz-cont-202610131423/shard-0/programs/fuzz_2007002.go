package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = -1
    var v3 uint = 42
    if (!(42 != 797)) {
        fmt.Println(v2)
        for fi4 := 0; fi4 < 13; fi4++ {
            fmt.Println(fi4, v3, v2)
            var v5 int16 = v2
            _ = fi4
            _ = v5
        }
    }
    v1 += v1
    var v6 int32 = 2147483647
    v7 := []int{0, -684}
    v8 := GenS1{F0: 2147483647, F1: v3, F2: 0}
    var v9 int8 = -9
    switch 32768 {
    case -315:
        switch "" {
        case "hello":
            var v10 uint16 = 2
            _ = v10
        case "":
            v2 = (-(v2))
        case "\n\t":
            var v11 int16 = (^(v2))
            _ = v11
        }
    }
    v12 := GenS1{F0: 853, F1: (v3 << ((v3) & 63)), F2: 228}
    v1 += v1
    fmt.Println(v2, v2, v3)
    v1 = 0.0
    _ = len(v7)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint32 = 65535
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v1 += v1
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
