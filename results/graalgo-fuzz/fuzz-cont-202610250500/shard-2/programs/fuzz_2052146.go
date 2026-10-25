package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 uint64) uint {
    var v4 uint16 = 2
    p2 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    var v5 int64 = -1
    v5 += v5
    var v6 int = 128
    fmt.Println(v6)
    v5 = v5
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{fi7, fi8, 4294967295, fi7}
            v10 := GenS1{F0: 255}
            fmt.Println(fi8)
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v6, fi11, v5)
            v6 = (fi11 << ((4294967295) & 63))
            _ = fi11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 string = ("a" + "\n\t")
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := []int{v6, 2147483647, (v6 + (-(fi15)))}
            var v17 uint64 = 800
            var v18 int64 = 42
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v19 := v2[42]
        _ = fi14
        _ = v19
    }
    if (1 != -642) {
        var v20 bool = (!true)
        _ = v20
    }
    var v21 int32 = 42
    var v22 uint64 = 1
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 float64 = -2.25
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v6
}
