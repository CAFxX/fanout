package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{65535}
    v2 := make(map[int]int)
    var v3 uint = 843
    var v4 int16 = -79
    var v5 int16 = (v4 - (v4 << ((42) & 63)))
    var v6 uint8 = 255
    v3 += (v3 * (v3 ^ v3))
    if false {
        var v7 int64 = 1
        v8 := GenS1{F0: (-(v4)), F1: 1, F2: 230}
        _ = v7
        _ = v8
    }
    v5 -= (v5 % ((v4) | 1))
    _ = len(v1)
    for fi9 := 0; fi9 < 7; fi9++ {
        v6 += v6
        _ = len(v1)
        _ = fi9
    }
    _ = len(v1)
    var v10 int8 = -79
    if true {
        v11 := GenS1{F0: v5, F1: 1, F2: int32(v6)}
        var v12 uint64 = 741
        _ = v11
        _ = v12
    } else {
        v3 += (v3 - (v3 * v3))
        v13 := GenS2{F0: v3}
        v5 = 1
        _ = v13
    }
    fmt.Println(v6)
    for fi14 := 0; fi14 < 5; fi14++ {
        v10 += (v10 - -1)
        switch "foo bar" {
        case "a":
            v1 = append(v1, 735)
        }
        _ = fi14
    }
    v10 = v10
    fmt.Println(v6, v6)
    if (119 == 32767) {
        v6 += (v6 % ((v6) | 1))
        v2[32768] = 529
        v3 -= (v3 % ((42) | 1))
    }
    delete(v2, 128)
    switch -1 {
    case 4294967295:
        var v15 int32 = -1
        _ = v15
    default:
        var v16 uint = v3
        _ = v16
    }
    var v17 int8 = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
