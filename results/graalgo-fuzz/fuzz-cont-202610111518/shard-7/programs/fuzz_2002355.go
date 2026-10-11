package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint64 = 886
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{-989, 9223372036854775807}
    var v5 int64 = 1
    v6 := make(map[uint32]uint32)
    var v7 string = ((("hello" + "") + ((("\n\t" + "") + "\n\t") + "hello")) + "a")
    var v8 int = -1
    fmt.Println(v8)
    v1 -= (v1 & 42)
    if len(v4) > 0 {
        v4[0] = (^(v8))
    }
    _ = len(v4)
    if true {
        var v9 uint64 = v1
        switch "a" {
        case "":
            for fi10 := 0; fi10 < 6; fi10++ {
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint64 = (v1 << ((1) & 63))
                    _ = fi11
                    _ = v12
                }
                v13 := v6[262]
                _ = fi10
                _ = v13
            }
        case "\n\t":
            var v14 uint32 = 85
            _ = v14
        }
        var v15 int16 = 1000
        _ = v15
        _ = v9
    } else {
        v16 := GenS1{F0: 1, F1: 737}
        _ = v16
    }
    var v17 int16 = -647
    switch v8 {
    case 32768:
        fmt.Println(v17, v5)
    case 128:
        v7 = (((v7 + (("foo bar" + v7) + v7)) + "a") + "")
    case 105:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
