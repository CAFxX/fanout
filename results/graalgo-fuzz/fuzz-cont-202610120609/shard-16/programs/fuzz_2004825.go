package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    fmt.Println(v2)
    var v5 int8 = -128
    v4 = (v4 / ((v4) | 1))
    switch 892 {
    case 1:
        v6 := []int{}
        _ = v6
    case 9223372036854775807:
        var v7 string = ""
        _ = v7
    case -328:
        var v8 uint16 = 0
        _ = v8
    }
    switch (((v2 + v2) + "foo bar") + v2) {
    case "hello":
        v4 = v4
    case "foo bar":
        fmt.Println(v4, v4, v2)
    case "":
        v4 -= 842
    default:
        var v9 int16 = 42
        _ = v9
    }
    v4 += v4
    switch 127 {
    case 2147483647:
        var v10 float64 = (1e10 - 3.14159)
        _ = v10
    case 127:
        if (!((1.5 == (((float64(1e10) - 1.5) * (0.1 * (3.14159 - 0.0))) + (3.14159 / 3.14159))) || (41 == 127))) {
            switch -606 {
            case -2147483648:
                fmt.Println(v2, v5)
            case 4294967295:
                var v11 uint8 = 42
                _ = v11
            case -669:
                if false {
                    v12 := GenS1{F0: 1, F1: 42}
                    _ = v12
                } else {
                    var v13 uint32 = 1
                    v14 := GenS1{F0: 42, F1: (v13 + 2147483647)}
                    fmt.Println(v2)
                    _ = v13
                    _ = v14
                }
            default:
                var v15 int32 = -302
                _ = v15
            }
        }
    default:
        for fi16 := 0; fi16 < 18; fi16++ {
            switch v2 {
            case "hello":
                var v17 int32 = 42
                _ = v17
            case "":
                fmt.Println(v4)
            case "a":
                v2 = ("foo bar" + ((v2 + v2) + "\n\t"))
            }
            _ = fi16
        }
    }
    v5 = v5
    v4 -= (v4 >> ((42) & 63))
    v18 := GenS2{F0: 127, F1: 2147483647, F2: 1}
    var v19 uint32 = 667
    var v20 int = 2147483647
    v21 := []int{v20}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
