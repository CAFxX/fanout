package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int16 = 0
    var v2 bool = (!true)
    var v3 uint64 = 4294967295
    var v4 uint = 461
    var v5 uint = (v4 >> ((v4) & 63))
    var v6 uint32 = 1
    switch 4294967295 {
    case -2147483648:
        v7 := []int{647, 309, int(v3), 127}
        _ = v7
    default:
        fmt.Println(v2, v3, v6)
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (fi9 % ((fi8) | 1)) {
            case 32767:
                v10 := GenS1{F0: -616}
                _ = v10
            }
            v6 += v6
            _ = fi9
        }
        v4 = v5
        v11 := []int{(fi8 + fi8), 128, (-(fi8)), 0}
        _ = fi8
        _ = v11
    }
    v12 := []int{}
    var v13 bool = true
    v6 += 65535
    var v14 bool = (uint8(v6) == 1)
    for fi15 := 0; fi15 < 5; fi15++ {
        _ = len(v12)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi15
    }
    fmt.Println(v13, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
