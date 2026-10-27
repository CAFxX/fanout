package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint64 }

func main() {
    var v1 int = 941
    var v2 uint16 = 42
    v3 := make(map[int]int)
    var v4 uint8 = 0
    v5 := GenS1{F0: 165, F1: uint32(v2)}
    v1 = (v1 << ((4294967295) & 63))
    var v6 uint = 18446744073709551615
    v7 := []int{9223372036854775807, (v1 ^ -946), (v1 + -735), -1}
    v1 = (v1 << ((4294967295) & 63))
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS2{F0: (true || true)}
        _ = fi8
        _ = v9
    }
    v2 += (v2 >> ((0) & 63))
    for fi10 := 0; fi10 < 15; fi10++ {
        v2 -= v2
        var v11 uint32 = 404
        v1 += 2147483647
        _ = fi10
        _ = v11
    }
    if (!false) {
        var v12 bool = ((("foo bar" + ("" + ("\n\t" + "hello"))) + "a") > "foo bar")
        _ = v12
    }
    v3[(-(v1))] = (v1 * (-(v1)))
    v13 := GenS3{F0: 1, F1: 4294967295, F2: 18446744073709551615}
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float32 = 1e10
        var v16 uint16 = (v2 << ((v6) & 63))
        _ = fi14
        _ = v15
        _ = v16
    }
    v1 -= (v1 - (v1 + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
