package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 2
    v3 := []string{"a", "a"}
    var v4 uint16 = 230
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    var v7 uint16 = 2
    v1 -= (v1 >> ((65535) & 63))
    if ((^(v2)) <= 65535) {
        v8 := GenS1{F0: v1}
        fmt.Println(v4)
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := v6[v2]
            if true {
                v11 := GenS1{F0: v1}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v2, fi12)
        _ = fi12
    }
    var v13 int64 = 0
    var v14 int64 = (-(v13))
    v4 = 65535
    var v15 uint64 = 627
    switch "\n\t" {
    case "":
        v14 += (^(v13))
    case "a":
        var v16 uint = uint(v7)
        _ = v16
    }
    var v17 bool = ((65535 <= uint32(v15)) && false)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
