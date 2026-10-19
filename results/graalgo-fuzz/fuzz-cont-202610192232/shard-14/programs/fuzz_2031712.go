package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", (("foo bar" + "foo bar") + ("\n\t" + ""))}
    var v3 uint64 = 101
    v4 := []uint32{47, 0, 0}
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    v3 += (v3 - (v3 % ((4294967295) | 1)))
    var v7 uint64 = v3
    v8 := GenS1{F0: 42, F1: 635, F2: (("\n\t" + "hello") + "a")}
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            if len(v4) > 0 {
                v4[0] = 806
            }
            var v11 int8 = 127
            delete(v5, fi10)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch "\n\t" {
    case "\n\t":
        var v12 uint64 = v7
        _ = v12
    case "hello":
        fmt.Println(v3, v7)
    }
    var v13 int16 = 1
    v7 -= (v7 >> ((42) & 63))
    switch -877 {
    case 42:
        v3 -= (v3 / ((1) | 1))
    case 2147483647:
        v7 += v7
    case -9223372036854775808:
        v14 := GenS2{F0: 432, F1: 0}
        _ = v14
    }
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 3; fi16++ {
        v3 += 385
        _ = fi16
    }
    v15 = 1
    fmt.Println(v15, v15, v13)
    v15 = (^(v13))
    v7 -= (v7 >> ((233) & 63))
    var v17 float64 = 1e10
    v6[678] = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
