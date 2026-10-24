package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 int8 }

func main() {
    var v1 uint32 = 189
    var v2 uint64 = 42
    v3 := make(map[int]uint32)
    var v4 bool = true
    fmt.Println(v2)
    var v5 int8 = 42
    v6 := GenS1{F0: 0}
    _ = len(v3)
    var v7 int8 = (^(v5))
    v8 := []int{32768, 4294967295, -9223372036854775808}
    switch "" {
    case "":
        var v9 uint64 = (v2 << ((42) & 63))
        _ = v9
    default:
        v3[128] = v1
    }
    var v10 string = ("hello" + "a")
    fmt.Println(v2)
    var v11 uint8 = 156
    for fi12 := 0; fi12 < 18; fi12++ {
        if ("\n\t" != v10) {
            fmt.Println(v10, v7, v2)
            v4 = (1 < 4294967295)
            fmt.Println(v5, v2)
        }
        switch v10 {
        case "a":
            v13 := GenS1{F0: 0}
            _ = v13
        case "hello":
            var v14 uint64 = (v2 % ((v2) | 1))
            _ = v14
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                v10 = v10
                v16 := GenS2{F0: 54, F1: (v5 << ((986) & 63)), F2: v7}
                _ = fi15
                _ = v16
            }
        }
        if (v4 && v4) {
            switch v10 {
            case "foo bar":
                var v17 uint16 = 2
                _ = v17
            }
            switch ((v10 + v10) + "foo bar") {
            case "a":
                fmt.Println(v7, v7, v11)
            default:
                var v18 bool = ((v1 == v1) || (!(!v4)))
                _ = v18
            }
            v19 := GenS2{F0: 282, F1: v7, F2: v5}
            _ = v19
        }
        _ = fi12
    }
    fmt.Println(v1, v7)
    var v20 int8 = (v5 << ((42) & 63))
    v21 := GenS1{F0: -1}
    fmt.Println(v7, v2, v11)
    v4 = true
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v20, v11, v1)
    default:
        switch 4294967295 {
        case -9223372036854775808:
            var v22 uint64 = v2
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
