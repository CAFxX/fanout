package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{"", ("\n\t" + ("a" + "a"))}
    var v2 int32 = 268
    var v3 float64 = 0.1
    var v4 uint8 = 203
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 2, F1: 42, F2: 0}
    v2 += (v2 << ((42) & 63))
    v7 := GenS2{F0: 32767}
    fmt.Println(v3)
    fmt.Println(v3, v2)
    var v8 uint8 = (v4 + (v4 * (v4 ^ (v4 * (^(v4))))))
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint32 = 0
        _ = fi9
        _ = v10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (18446744073709551615 != 18446744073709551615) {
        var v11 int16 = -32768
        _ = v11
    }
    v8 = 140
    fmt.Println(v3)
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 -= (v4 + (v8 % ((v8) | 1)))
        _ = fi12
    }
    v2 -= (v2 ^ v2)
    v2 = (v2 & 2147483647)
    var v13 int64 = 0
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        v14 := GenS2{F0: -32768}
        _ = v14
    }
    v8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
