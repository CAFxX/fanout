package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 1
    var v3 uint32 = 223
    delete(v1, v3)
    v2 -= 65535
    fmt.Println(v2, v3)
    v4 := GenS1{F0: 293, F1: (v3 << ((0) & 63))}
    if true {
        v5 := []int{108, 0}
        v6 := GenS1{F0: 341, F1: 467}
        _ = v5
        _ = v6
    }
    switch 623 {
    case 32767:
        var v7 int8 = -1
        _ = v7
    default:
        v8 := GenS1{F0: 0, F1: v3}
        _ = v8
    }
    v9 := GenS1{F0: 1000, F1: 65535}
    for fi10 := 0; fi10 < 12; fi10++ {
        if (true || (!true)) {
            v3 += v3
        } else {
            var v11 string = ("" + "hello")
            var v12 int64 = -9223372036854775808
            var v13 bool = true
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 uint16 = 729
    v3 = 2147483647
    v14 += v2
    for fi15 := 0; fi15 < 16; fi15++ {
        switch (-(fi15)) {
        case 1:
            v16 := []int{42}
            _ = v16
        case -9223372036854775808:
            v3 -= 18
        default:
            v14 = (v2 & (v2 / ((v2) | 1)))
        }
        _ = fi15
    }
    v17 := GenS1{F0: 0, F1: (v3 ^ (v3 << ((uint(v3)) & 63)))}
    switch "a" {
    case "hello":
        v18 := GenS1{F0: 1, F1: v3}
        _ = v18
    }
    fmt.Println(v3, v14)
    v19 := v1[(v3 - v3)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
