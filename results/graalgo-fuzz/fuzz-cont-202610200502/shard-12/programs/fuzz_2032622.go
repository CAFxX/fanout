package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 uint) uint {
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 = p1
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{fi9, (fi8 ^ -1), (-(fi8))}
            fmt.Println(v7, fi8, v7)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 uint = 553
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]bool)
    var v4 int = -137
    var v5 bool = (255 > 255)
    var v6 int = (v4 | v4)
    if (-1 > 0) {
        fmt.Println(v4, v2, v6)
    }
    if v5 {
        if ((!true) || (42 <= -80)) {
            delete(v3, (v2 % ((v2) | 1)))
            v5 = v5
            var v11 int32 = 42
            _ = v11
        }
        switch v4 {
        case -2147483648:
            if true {
                v12 := []int{}
                _ = v12
            }
        case 851:
            var v13 int32 = -342
            _ = v13
        case 42:
            fmt.Println(v1)
        default:
            v2 = v2
        }
    } else {
        var v14 float64 = float64((3.14159 / 1e10))
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int64 = 1
        _ = fi15
        _ = v16
    }
    v17 := []int{v6, (v4 << ((v1) & 63)), (^(v6))}
    v6 = v4
    v2 += (v2 - (v2 % ((v2) | 1)))
    if (!false) {
        v5 = (255 != 255)
        var v18 uint32 = v2
        v18 = 2
        _ = v18
    }
    fmt.Println(v1, v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
