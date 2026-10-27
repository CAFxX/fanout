package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 0
    var v3 float64 = 0.1
    var v4 uint = 42
    var v5 uint32 = (v2 << ((42) & 63))
    var v6 int64 = 0
    v2 += (v2 >> ((v4) & 63))
    var v7 int = -9223372036854775808
    if (!false) {
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v6, v6, v4)
                _ = fi9
            }
            var v10 uint = (^(v4))
            _ = fi8
            _ = v10
        }
        var v11 uint8 = 41
        var v12 uint64 = 553
        _ = v11
        _ = v12
    }
    v6 = (^(v6))
    v13 := GenS1{F0: v7}
    v14 := []int{v7, (v7 % (((^(v7))) | 1))}
    v14 = append(v14, int(v4))
    v4 = v4
    fmt.Println(v5)
    v15 := GenS2{F0: 339, F1: ("foo bar" + "foo bar"), F2: 4294967295}
    var v16 uint = v4
    v4 += (v16 << ((v4) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
