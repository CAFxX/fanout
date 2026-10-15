package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32) int64 {
    p0 += p0
    for fi4 := 0; fi4 < 20; fi4++ {
        p0 = 2
        var v5 float32 = 3.14159
        _ = fi4
        _ = v5
    }
    var v6 int64 = 1
    var v7 uint = 42
    _ = p0
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := []int{-2147483648, 499}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 0
    switch "\n\t" {
    case "\n\t":
        if (!(!false)) {
            v1 = append(v1, 1)
            _ = len(v1)
        }
    default:
        v3 += (v3 ^ 65535)
    }
    v1 = append(v1, 32768)
    v3 = (v3 - (^(v3)))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint = 660
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint64 = 2
            var v12 int64 = 586
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int64 = int64(v9)
        _ = fi8
        _ = v13
        _ = v9
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        fmt.Println(v3)
    }
    var v14 float64 = 3.14159
    if (1 >= 2) {
        v15 := GenS1{F0: -1, F1: (((2 >= 128) || (838 < (v3 ^ (^(v3))))) || ((v3 % (((v3 << ((0) & 63))) | 1)) == v3))}
        v16 := GenS2{F0: 551}
        _ = v15
        _ = v16
    } else {
        switch 0 {
        case -2147483648:
            v17 := GenS1{F0: 0, F1: ((!(false || true)) || true)}
            _ = v17
        }
    }
    var v18 int = -328
    switch "" {
    case "hello":
        v18 -= v18
    }
    var v19 int8 = 0
    var v20 uint32 = 2147483648
    var v21 string = "\n\t"
    v3 = (v20 >> ((772) & 63))
    fmt.Println(v18)
    var v22 uint = 65535
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
