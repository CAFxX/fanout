package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 467
    var v3 uint32 = uint32(v2)
    v2 -= (^(v2))
    switch 127 {
    case 379:
        var v4 uint64 = 4294967295
        _ = v4
    case -2147483648:
        var v5 uint16 = 42
        _ = v5
    case 587:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6, fi6)
            _ = fi6
        }
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = (v2 | v2)
            v9 := v1[v3]
            var v10 uint8 = 42
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    }
    if (true && true) {
        fmt.Println(v2)
    } else {
        v3 -= v3
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = -109
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{128, (fi13 | 799), (fi11 * fi13), (fi13 ^ fi11)}
                var v15 uint16 = 0
                v15 -= (v15 << ((1) & 63))
                var v16 uint64 = 951
                switch (fi11 / ((fi11) | 1)) {
                case 592:
                    var v17 uint8 = 0
                    _ = v17
                case 0:
                    v2 = (v2 + 0)
                case -806:
                    fmt.Println(v16, v3, v3)
                default:
                    _ = len(v14)
                }
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        fmt.Println(v3, v3)
    }
    v3 += (v3 & (^(v3)))
    v3 = (v3 << ((0) & 63))
    var v18 int64 = 0
    v2 = (v2 % (((v2 % (((v2 - (^(v2)))) | 1))) | 1))
    switch ("\n\t" + "foo bar") {
    case "":
        v19 := []int{128, -2147483648, -882}
        _ = v19
    case "hello":
        var v20 uint32 = (v3 % (((^(v3))) | 1))
        _ = v20
    default:
        var v21 int32 = -538
        _ = v21
    }
    v1[2147483648] = "a"
    _ = len(v1)
    v22 := []int{128, -375, 354}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
