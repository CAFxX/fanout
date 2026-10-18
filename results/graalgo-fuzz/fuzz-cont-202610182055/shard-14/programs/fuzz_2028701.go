package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    var v5 uint8 = 2
    var v6 uint32 = 65535
    var v7 int = 9223372036854775807
    v3[42] = -9223372036854775808
    fmt.Println(v6, v5, v7)
    var v8 float32 = v2
    fmt.Println(v6, v4)
    for fi9 := 0; fi9 < 14; fi9++ {
        if true {
            v10 := GenS1{F0: 127, F1: v6, F2: 100}
            if false {
                v11 := GenS2{F0: true, F1: -1, F2: 4294967295}
                _ = v11
            } else {
                var v12 int64 = -1
                _ = v12
            }
            var v13 uint64 = 244
            _ = v10
            _ = v13
        } else {
            if ("a" == v4) {
                fmt.Println(fi9)
            }
        }
        v6 -= (v6 << ((0) & 63))
        var v14 int16 = 1000
        _ = fi9
        _ = v14
    }
    v15 := GenS2{F0: true, F1: -9223372036854775808, F2: 65535}
    if (true || false) {
        delete(v1, v4)
        var v16 float64 = -2.25
        var v17 float32 = float32(float64(v8))
        _ = v16
        _ = v17
    } else {
        var v18 uint = 617
        _ = v18
    }
    var v19 float64 = 0.1
    switch "" {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        for fi20 := 0; fi20 < 6; fi20++ {
            v5 -= v5
            v19 = v19
            _ = fi20
        }
    }
    v21 := GenS2{F0: false, F1: 0, F2: 758}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
