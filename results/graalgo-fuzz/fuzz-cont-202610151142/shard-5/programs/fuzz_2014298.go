package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int64 = 1
    var v2 int = 867
    var v3 int32 = 1
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    switch v2 {
    case -315:
        v6 := GenS1{F0: -128, F1: 745, F2: 42}
        _ = v6
    default:
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 0, F1: 2, F2: 42}
            _ = fi7
            _ = v8
        }
    }
    switch (v2 & -1) {
    case 766:
        v2 = (v2 >> ((2) & 63))
    }
    delete(v4, "hello")
    v9 := v5[("hello" + "hello")]
    v10 := []int{v2, v2}
    var v11 int16 = 1
    v2 -= -2147483648
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: v11, F1: 177}
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int16 = v11
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if (v2 <= (v2 * v2)) {
        v11 = -32768
        v16 := GenS1{F0: -76, F1: 336, F2: 128}
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint32 = 4294967295
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    switch "\n\t" {
    case "":
        var v19 int64 = (v1 + 1)
        _ = v19
    case "\n\t":
        var v20 uint64 = 42
        _ = v20
    case "hello":
        var v21 uint = 0
        _ = v21
    default:
        var v22 string = "a"
        _ = v22
    }
    if false {
        v3 += v3
    } else {
        var v23 int16 = -563
        for fi24 := 0; fi24 < 13; fi24++ {
            fmt.Println(v23, v3)
            _ = fi24
        }
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
