package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    if (-9223372036854775808 != -9223372036854775808) {
        var v4 bool = false
        _ = v4
    }
    v5 := GenS1{F0: 32767, F1: true, F2: 952}
    switch v1 {
    case "\n\t":
        if false {
            var v6 int16 = 0
            v1 = v1
            _ = v6
        }
    default:
        v1 = (v1 + v1)
    }
    var v7 int16 = 183
    var v8 uint64 = 753
    if false {
        var v9 uint = uint(v8)
        _ = v9
    } else {
        fmt.Println(v7)
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        _ = len(v2)
        _ = fi10
    }
    v11 := []int{-1, -9223372036854775808}
    _ = len(v11)
    if ((((("" + v1) + v1) + "foo bar") + v1) == v1) {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 bool = (!(1 > 175))
            v7 -= v7
            _ = fi12
            _ = v13
        }
        var v14 uint64 = (v8 << ((18446744073709551615) & 63))
        _ = v14
    } else {
        var v15 int32 = 42
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
}
