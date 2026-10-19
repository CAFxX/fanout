package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 127
    var v3 uint16 = 176
    var v4 int = 251
    var v5 int64 = 216
    delete(v1, -1)
    switch v4 {
    case -9223372036854775808:
        v2 += (v2 / ((v2) | 1))
    case -616:
        var v6 uint64 = 42
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint32 = 2
        var v9 string = ""
        fmt.Println(v2)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := v1[v4]
    switch 994 {
    case 2147483647:
        var v11 int32 = 42
        _ = v11
    case 4294967295:
        v5 -= -1
    case -1:
        for fi12 := 0; fi12 < 14; fi12++ {
            if true {
                if (-2147483648 >= 0) {
                    var v13 uint8 = 127
                    _ = v13
                }
                v1[128] = v10
            } else {
                v10 = 675
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                v2 = 1
                _ = fi14
            }
            _ = fi12
        }
    }
    v5 -= v5
    v15 := GenS1{F0: 382, F1: 42}
    v16 := []int{(v4 | -410), 0, (v4 << ((388) & 63))}
    var v17 int32 = -799
    switch v4 {
    case 192:
        var v18 bool = true
        _ = v18
    }
    fmt.Println(v3)
    var v19 bool = ("a" < "hello")
    var v20 uint16 = uint16(v10)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
