package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []int{42}
    var v2 int8 = -89
    v3 := []int{}
    var v4 uint16 = 266
    v5 := make(map[uint32]bool)
    v6 := []string{"", (("" + (("a" + "hello") + "\n\t")) + "foo bar"), ("foo bar" + "hello")}
    switch -1 {
    case 1:
        v7 := GenS1{F0: 889}
        _ = v7
    case 42:
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS2{F0: int32(v2), F1: -374, F2: v2}
            var v10 bool = ((fi8 / (((fi8 & (-(fi8)))) | 1)) < fi8)
            v11 := GenS2{F0: 42, F1: 0, F2: v2}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    v2 = -128
    var v12 int16 = -714
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v4)
        v12 -= (v12 << ((4294967295) & 63))
        _ = fi13
    }
    var v14 int = 32768
    v2 -= 81
    v15 := GenS2{F0: 2147483647, F1: 39, F2: (v2 + v2)}
    fmt.Println(v14, v14)
    switch (^(v14)) {
    case -2147483648:
        var v16 float32 = 0.1
        _ = v16
    default:
        if (0 < 42) {
            switch (("\n\t" + "") + "hello") {
            case "foo bar":
                v12 += int16(v14)
            default:
                v17 := GenS3{F0: 2}
                _ = v17
            }
        }
    }
    var v18 float32 = 0.0
    var v19 uint64 = 722
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
