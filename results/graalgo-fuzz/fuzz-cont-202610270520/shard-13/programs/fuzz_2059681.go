package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[int]uint32)
    var v5 int8 = 1
    var v6 int = -9223372036854775808
    var v7 uint32 = 2147483648
    switch (v6 / ((v6) | 1)) {
    case 32767:
        v7 = (v7 << ((847) & 63))
    case 723:
        v8 := GenS1{F0: 42, F1: -90, F2: 1}
        _ = v8
    case 128:
        var v9 uint = 438
        _ = v9
    }
    v2[2] = true
    v10 := []int{}
    if false {
        v6 += v6
    }
    var v11 int8 = 61
    switch (v6 / (((v6 - (-(v6)))) | 1)) {
    case -266:
        for fi12 := 0; fi12 < 4; fi12++ {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v5)
            case "hello":
                v1[v7] = v3
            }
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v6 = 93
        _ = fi13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            v11 += v5
            _ = fi15
        }
        v16 := GenS2{F0: (v6 >> ((255) & 63))}
        if len(v10) > 0 {
            v10[0] = (^(fi14))
        }
        _ = fi14
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
