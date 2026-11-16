package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 0
    var v3 uint = 4294967295
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 13; fi6++ {
        v2 = (v2 ^ v2)
        var v7 int16 = 109
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: -904, F2: 2147483648}
    var v9 int16 = -582
    v9 -= (v9 / ((-32768) | 1))
    v10 := []int{32767}
    v11 := GenS1{F0: -253, F1: v9, F2: 1}
    if false {
        v2 = v2
        var v12 int8 = 42
        _ = v12
    }
    v13 := GenS1{F0: 2147483647, F1: (v9 / ((1) | 1)), F2: 983}
    if (!false) {
        if ((false || true) && true) {
            var v14 uint = (v3 / ((2) | 1))
            var v15 uint16 = 0
            var v16 string = "\n\t"
            _ = v14
            _ = v15
            _ = v16
        }
    } else {
        var v17 uint32 = 536
        v3 += 65535
        _ = v17
    }
    v10 = append(v10, -146)
    if false {
        var v18 int16 = (v9 - v9)
        fmt.Println(v3, v4)
        _ = v18
    }
    switch ("hello" + ("hello" + "a")) {
    case "hello":
        v3 = (v3 + 0)
    case "\n\t":
        v19 := GenS1{F0: 501, F1: v9, F2: 2147483648}
        _ = v19
    }
    var v20 float64 = ((-2.25 + 1.5) * float64((1e10 * (0.1 + (0.1 + 0.1)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
