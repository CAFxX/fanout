package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int64 = 558
    v2 := make(map[uint32]bool)
    var v3 uint = 18446744073709551615
    v4 := []string{}
    var v5 float32 = 0.1
    var v6 uint = v3
    v5 = ((float32(1.5) / 1.0) * (1e10 * v5))
    v7 := GenS1{F0: 4294967295}
    v4 = append(v4, ("\n\t" + "\n\t"))
    var v8 int64 = (-(v1))
    if (127 < -66) {
        var v9 uint64 = 0
        var v10 uint16 = 0
        switch 0 {
        case -218:
            var v11 float64 = 0.0
            _ = v11
        case -2147483648:
            var v12 uint32 = 157
            _ = v12
        case 973:
            fmt.Println(v3)
        }
        _ = v10
        _ = v9
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint8 = 255
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            v6 += 816
            _ = fi15
        }
        var v16 int32 = -2147483648
        _ = v16
    }
    _ = len(v4)
    v6 += v3
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint32 = 65535
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 255, F1: 763}
    var v20 uint = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
