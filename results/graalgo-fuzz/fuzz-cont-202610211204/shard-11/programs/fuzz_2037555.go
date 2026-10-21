package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() int16 {
    v7 := GenS1{F0: -830}
    var v8 uint16 = 348
    v8 += v8
    var v9 uint64 = 1
    _ = v7
    _ = v8
    _ = v9
    return 212
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -346
    v3 := []uint32{}
    var v4 uint32 = 65535
    v5 := make(map[int]uint32)
    v6 := []int{}
    for fi10 := 0; fi10 < 9; fi10++ {
        switch ("foo bar" + "foo bar") {
        case "hello":
            var v11 int64 = 1
            _ = v11
        case "a":
            fmt.Println(fi10, fi10, v2)
        case "":
            v12 := GenS1{F0: -9223372036854775808}
            _ = v12
        default:
            v4 -= v4
        }
        _ = fi10
    }
    var v13 uint = 4294967295
    v13 = 2
    var v14 int16 = v2
    var v15 uint8 = 69
    v16 := GenS1{F0: 42}
    fmt.Println(v15, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = 310
    }
    var v17 string = "\n\t"
    fmt.Println(v17, v2, v2)
    v18 := GenS2{F0: 42, F1: (v4 & uint32(v15))}
    v19 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
