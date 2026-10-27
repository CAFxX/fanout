package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0() uint32 {
    var v7 uint64 = 2
    if false {
        v8 := GenS1{F0: 1, F1: 9223372036854775807}
        if ("\n\t" > "") {
            v9 := GenS1{F0: 42, F1: 42}
            v10 := GenS1{F0: -1, F1: 9223372036854775807}
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: 42, F1: 1}
        _ = v11
        _ = v8
    } else {
        switch 800 {
        case 42:
            for fi12 := 0; fi12 < 17; fi12++ {
                v7 = v7
                v7 = (^(v7))
                if (4294967295 < 24) {
                    fmt.Println(fi12)
                } else {
                    var v13 bool = (float32(0.1) == float32(((1.5 / 0.1) + -2.25)))
                    _ = v13
                }
                _ = fi12
            }
        case 2147483647:
            if false {
                var v14 string = ""
                _ = v14
            }
        case 9223372036854775807:
            v15 := []int{1, 55, 42, -59}
            _ = v15
        }
        v16 := []int{448, 127, -1, 2147483647}
        _ = v16
    }
    v7 = (^(v7))
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 string = "a"
    var v4 uint16 = 2
    var v5 uint16 = (v4 + (v4 + (v4 + (v4 >> ((255) & 63)))))
    v6 := make(map[uint32]string)
    var v17 uint32 = 2147483648
    v3 = v3
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint16 = (v5 ^ (v4 | 943))
        v2 = append(v2, (v17 & genFn0()))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(fi20, v5, v5)
        fmt.Println(v3, v3)
        _ = fi20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int64 = 0
        v5 = v5
        var v23 int64 = (v22 ^ v22)
        _ = fi21
        _ = v22
        _ = v23
    }
    if (!false) {
        v1[(^(v17))] = v17
    }
    v5 += 0
    v5 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
