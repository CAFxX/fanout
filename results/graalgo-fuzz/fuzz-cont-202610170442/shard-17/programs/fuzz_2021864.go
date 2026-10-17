package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }

func main() {
    var v1 bool = true
    var v2 int64 = 42
    var v3 uint8 = 255
    var v4 int = -416
    v5 := []int{-199, v4, (-(v4)), v4}
    var v6 int32 = 1
    v1 = (!false)
    v6 += (v6 * (-(v6)))
    v7 := GenS1{F0: 42, F1: 1, F2: 0}
    for fi8 := 0; fi8 < 1; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v9 := GenS1{F0: 2, F1: 0, F2: 0}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi8
        _ = v9
    }
    v3 += (v3 ^ 42)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 uint32 = 268
    switch "foo bar" {
    case "hello":
        v4 = -887
    }
    v10 = 0
    switch v4 {
    case -152:
        _ = len(v5)
    }
    var v11 int16 = 32767
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = v2
        for fi14 := 0; fi14 < 10; fi14++ {
            v11 = 0
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v15 int = v4
    fmt.Println(v10)
    var v16 int = (v4 & v4)
    v17 := GenS2{F0: false, F1: 240, F2: -384}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
