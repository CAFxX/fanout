package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 int16 = -1
    v3 := make(map[int]int)
    var v4 int16 = 950
    v1 += v1
    v5 := []int{2147483647, -2147483648, 127}
    v6 := GenS1{F0: -1}
    v7 := GenS2{F0: (^(v1))}
    if (140 > 2147483648) {
        var v8 int = -174
        _ = v8
    } else {
        var v9 int32 = -2147483648
        fmt.Println(v4, v9)
        _ = v9
    }
    if true {
        var v10 int = 2147483647
        fmt.Println(v2)
        var v11 uint = 18446744073709551615
        _ = v10
        _ = v11
    } else {
        v2 -= (-(v4))
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v12 := GenS2{F0: v1}
    switch (("hello" + "a") + "") {
    case "":
        v13 := GenS1{F0: -128}
        _ = v13
    case "hello":
        var v14 uint16 = 65535
        _ = v14
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        v2 = (v4 - 0)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        switch fi15 {
        case 0:
            var v16 uint64 = 42
            _ = v16
        }
        var v17 bool = (0.0 <= float64(1e10))
        var v18 int32 = 1
        _ = fi15
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
