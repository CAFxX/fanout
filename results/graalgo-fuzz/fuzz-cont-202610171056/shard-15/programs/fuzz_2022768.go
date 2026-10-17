package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint8 = 9
    v2 := []uint32{41, 2147483647, 42, 2}
    var v3 uint64 = 18446744073709551615
    var v4 bool = ((v3 + uint64(v1)) < v3)
    v5 := []int{}
    v6 := GenS1{F0: 42, F1: "foo bar", F2: -418}
    v7 := GenS2{F0: (^(v1)), F1: 765, F2: "\n\t"}
    var v8 int = -464
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint64 = (v3 / ((4294967295) | 1))
        v5 = append(v5, (fi9 << ((0) & 63)))
        _ = fi9
        _ = v10
    }
    v8 = (v8 * (v8 ^ v8))
    if v4 {
        if ((!(!v4)) || (255 == 42)) {
            v11 := GenS3{F0: (v3 ^ v3)}
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS4{F0: (42 > 42)}
            _ = fi12
            _ = v13
        }
        fmt.Println(v1, v3)
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v1 = (v1 * v1)
        v15 := GenS2{F0: v1, F1: 756, F2: ("" + "\n\t")}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        if (-128 == 17) {
            var v17 uint32 = 4294967295
            _ = v17
        }
        _ = fi16
    }
    if (-551 != -9223372036854775808) {
        var v18 int64 = 291
        var v19 uint32 = 1
        var v20 string = ("foo bar" + "a")
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int16 = -32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
