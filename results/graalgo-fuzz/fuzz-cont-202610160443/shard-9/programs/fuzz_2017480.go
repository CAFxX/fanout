package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = -345
    var v2 int8 = 1
    var v3 int8 = v2
    v4 := make(map[string]int)
    var v5 uint16 = 707
    v6 := []int{749, 9223372036854775807}
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            _ = len(v6)
            var v8 uint16 = 162
            _ = len(v6)
            _ = v8
        }
        _ = fi7
    }
    var v9 uint32 = 671
    v10 := GenS1{F0: 0, F1: 65535}
    delete(v4, "a")
    v11 := GenS1{F0: 23, F1: 2}
    var v12 uint8 = 80
    v13 := GenS2{F0: 32767}
    var v14 uint16 = (v5 << ((42) & 63))
    var v15 int16 = 1
    switch -9223372036854775808 {
    case -80:
        var v16 uint32 = (v9 % ((v9) | 1))
        _ = v16
    default:
        v2 -= v3
    }
    fmt.Println(v9)
    fmt.Println(v1)
    v17 := GenS2{F0: (v14 / (((v14 & uint16(v9))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
