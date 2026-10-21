package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    var v3 int = 9223372036854775807
    fmt.Println(v3, v1)
    var v4 int16 = v1
    var v5 uint = 4294967295
    if true {
        var v6 uint64 = 0
        v7 := v2[42]
        _ = v6
        _ = v7
    } else {
        switch ("a" + "hello") {
        case "":
            var v8 int8 = 100
            _ = v8
        case "a":
            if true {
                v9 := GenS1{F0: 0, F1: 127, F2: 9223372036854775807}
                _ = v9
            }
        case "hello":
            var v10 uint = v5
            _ = v10
        default:
            var v11 int32 = -182
            _ = v11
        }
    }
    v12 := GenS2{F0: (^(v5))}
    var v13 uint8 = 76
    v14 := []int{(v3 - (v3 & (v3 - (v3 >> ((v5) & 63))))), (v3 * v3), (v3 + (v3 / ((v3) | 1))), 939}
    if true {
        _ = len(v14)
    } else {
        switch v3 {
        case 468:
            var v15 int64 = -1
            _ = v15
        case 42:
            v1 += -507
        case 421:
            v16 := GenS1{F0: 397, F1: (v13 & v13), F2: -1}
            _ = v16
        }
    }
    var v17 uint = v5
    var v18 int8 = 120
    var v19 int = v3
    var v20 int8 = v18
    v13 = (v13 * v13)
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if false {
                fmt.Println(v3, v20)
                var v23 uint16 = 1
                _ = v23
            }
            v1 = 763
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println(v17)
    v18 = (v20 << (((^(v17))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
