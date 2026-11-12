package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 int64; F2 int32 }

func main() {
    var v1 float32 = 0.1
    var v2 int = -834
    v3 := make(map[uint32]string)
    v2 -= v2
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 float64 = float64(1e10)
        var v6 uint32 = 1
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{v2, (v2 << ((255) & 63))}
        v9 := GenS1{F0: 2, F1: -128, F2: 100}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-568, v2, 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        v1 = ((float32(1e10) * v1) + (v1 - 0.1))
        _ = fi11
    }
    v12 := GenS1{F0: 847, F1: -1, F2: -66}
    v13 := GenS1{F0: 65535, F1: -73, F2: 65}
    _ = len(v3)
    v1 = v1
    v14 := GenS2{F0: -2147483648, F1: -528, F2: -198}
    v3[2147483647] = ("hello" + "\n\t")
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v2, fi15, v2)
            var v17 int16 = -32768
            switch "\n\t" {
            case "":
                v3[253] = "a"
            case "a":
                var v18 uint64 = 42
                _ = v18
            case "\n\t":
                var v19 string = "foo bar"
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        v20 := GenS1{F0: 2147483648, F1: -128, F2: -105}
        _ = fi15
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
