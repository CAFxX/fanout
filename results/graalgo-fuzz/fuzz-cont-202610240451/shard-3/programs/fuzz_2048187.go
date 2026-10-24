package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }

func main() {
    v1 := []int{2147483647, 1}
    v2 := make(map[string]string)
    v3 := []uint32{2, 907}
    var v4 int64 = 45
    var v5 uint16 = 727
    var v6 int = 32768
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = v5
        _ = fi7
        _ = v8
    }
    var v9 int16 = int16(v4)
    if (!(true && (v9 != (v9 * (v9 | int16(v4)))))) {
        v5 = (v5 / ((405) | 1))
    }
    v4 += v4
    var v10 uint64 = 428
    var v11 uint16 = 0
    var v12 uint = 65535
    for fi13 := 0; fi13 < 8; fi13++ {
        switch (fi13 | (v6 / (((fi13 / ((0) | 1))) | 1))) {
        case -875:
            v9 += -1
        case 640:
            if (58 == 128) {
                v4 = 42
                var v14 uint = (^(v12))
                _ = v14
            }
        default:
            var v15 int32 = -993
            _ = v15
        }
        v2[("foo bar" + ("" + "hello"))] = "\n\t"
        fmt.Println(fi13, v9)
        _ = fi13
    }
    v5 += (v5 & v5)
    v11 = v11
    v10 -= (v10 << ((v12) & 63))
    v3 = append(v3, uint32(v11))
    switch v6 {
    case 42:
        switch "hello" {
        case "a":
            v5 = (v5 << ((v12) & 63))
        case "hello":
            var v16 int = (v6 | 563)
            _ = v16
        default:
            fmt.Println(v4, v11, v9)
        }
    case 0:
        v9 = v9
    case 128:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v4)
            _ = len(v2)
            _ = fi17
        }
    default:
        switch (v6 - v6) {
        case 128:
            switch v6 {
            case 2147483647:
                v10 = uint64(v12)
            default:
                v5 += 255
            }
        }
    }
    v18 := GenS1{F0: -1, F1: v12, F2: (!true)}
    fmt.Println(v11, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
