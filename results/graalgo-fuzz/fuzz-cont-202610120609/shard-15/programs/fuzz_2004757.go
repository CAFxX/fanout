package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    v3 := make(map[string]bool)
    var v4 int = -9223372036854775808
    v5 := GenS1{F0: 42, F1: 42, F2: -7}
    var v6 uint64 = 42
    if false {
        fmt.Println(v6, v2, v6)
        var v7 int16 = -129
        var v8 uint8 = 0
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 42, F1: -146, F2: 42}
        v10 := GenS2{F0: (v6 % ((v6) | 1)), F1: 0}
        _ = v10
        _ = v9
    }
    var v11 float64 = 1e10
    v12 := GenS3{F0: 2147483647, F1: 2147483648, F2: 187}
    v13 := []int{}
    if false {
        var v14 int64 = 9223372036854775807
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float32 = 1.5
            _ = fi15
            _ = v16
        }
        switch (^(v4)) {
        case -1:
            if false {
                v17 := GenS1{F0: 2, F1: -204, F2: -4}
                _ = v17
            } else {
                v3["hello"] = false
                v4 += v4
                v11 = 1.5
            }
        }
        _ = v14
    } else {
        v18 := GenS4{F0: (v6 << ((97) & 63))}
        _ = v18
    }
    fmt.Println(v4)
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS5{F0: 2}
        _ = fi19
        _ = v20
    }
    _ = len(v1)
    var v21 float32 = 0.1
    v11 = (((v11 / 1.5) + v11) - v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
