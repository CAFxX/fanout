package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32, p1 int8, p2 bool) uint {
    p1 += (p1 % ((p1) | 1))
    p0 -= 2147483647
    v5 := GenS1{F0: 1}
    var v6 uint64 = 674
    v7 := []int{-2147483648, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 805
}
func genFn1(p0 string, p1 bool) uint32 {
    for fi8 := 0; fi8 < 16; fi8++ {
        p1 = (!((p1 && ((p1 && p1) && (1 < 18446744073709551615))) || true))
        v9 := GenS1{F0: 817}
        _ = fi8
        _ = v9
    }
    p1 = (1 == 741)
    _ = p0
    _ = p1
    return 818
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v10 := GenS2{F0: -150, F1: 408}
    var v11 uint = 4294967295
    v1[("\n\t" + (("a" + "hello") + (("" + ("a" + "foo bar")) + "foo bar")))] = 2147483647
    var v12 int32 = 42
    var v13 int64 = 1
    var v14 float32 = 3.14159
    switch "" {
    case "foo bar":
        v11 = (v11 >> (((^(v11))) & 63))
    }
    var v15 bool = (("foo bar" + "") != "hello")
    v16 := []int{9223372036854775807, 1}
    var v17 uint = 1
    if v15 {
        v18 := GenS2{F0: (v13 ^ (^(v13))), F1: 1}
        v14 -= ((v14 + v14) / v14)
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
