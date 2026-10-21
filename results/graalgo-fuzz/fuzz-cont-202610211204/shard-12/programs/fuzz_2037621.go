package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint8, p2 int) uint32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5, fi4, p2)
            var v6 int64 = -1
            var v7 uint8 = (p1 | 186)
            fmt.Println(p0)
            v7 += p1
            v8 := GenS1{F0: 173}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = (255 <= 53)
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    v3 = v3
    var v9 int = 485
    var v10 uint = 886
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 13
        _ = len(v2)
        var v13 string = ""
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = v1
    if (!(871 == 2)) {
        for fi15 := 0; fi15 < 11; fi15++ {
            v10 -= v10
            _ = fi15
        }
    }
    var v16 int = v9
    var v17 uint64 = 1
    v1 = true
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
