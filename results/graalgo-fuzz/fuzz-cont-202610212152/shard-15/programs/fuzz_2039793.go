package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 54
    v3 := []string{}
    v4 := []int{4294967295, -1, -9223372036854775808, -9223372036854775808}
    v2 = 255
    if len(v4) > 0 {
        v4[0] = 210
    }
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: int32(v2)}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: -1}
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(v1, fi7)
            var v10 int32 = 0
            fmt.Println(fi9)
            v2 = (v2 | (v2 | v2))
            v3 = append(v3, ("a" + "a"))
            v11 := GenS1{F0: v10}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS2{F0: (v2 + 77), F1: 255}
    var v13 int8 = 72
    for fi14 := 0; fi14 < 7; fi14++ {
        v13 -= (v13 << ((362) & 63))
        _ = fi14
    }
    var v15 int64 = 9223372036854775807
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = v13
        _ = fi16
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 128)
    v4 = append(v4, -227)
    v2 = (v2 << ((122) & 63))
    var v18 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
