package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := []string{"a", "\n\t", "", "a"}
    var v2 uint32 = 4294967295
    v3 := []uint32{2147483648}
    v4 := make(map[int]bool)
    var v5 float64 = float64(1e10)
    var v6 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS1{F0: 0, F1: 1}
    for fi8 := 0; fi8 < 16; fi8++ {
        v6 = 1
        _ = fi8
    }
    var v9 uint64 = 1
    v10 := GenS1{F0: -9223372036854775808, F1: 100}
    var v11 int8 = -128
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v5 += float64(-2.25)
            _ = fi13
        }
        var v14 uint32 = (^(v2))
        _ = fi12
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: 0}
    _ = len(v3)
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS1{F0: 91, F1: v11}
            _ = fi17
            _ = v18
        }
        if (42 <= 796) {
            v3 = append(v3, (^(v2)))
            var v19 uint8 = (v6 >> ((718) & 63))
            v1 = append(v1, ("a" + "a"))
            _ = v19
        }
        v1 = append(v1, (("" + ("a" + "hello")) + "\n\t"))
        _ = fi16
    }
    _ = len(v4)
    switch "\n\t" {
    case "\n\t":
        switch 32767 {
        case 1:
            fmt.Println(v2)
        case 2:
            fmt.Println(v5, v11)
        case 4294967295:
            var v20 int8 = v11
            _ = v20
        }
    case "a":
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 float32 = 0.1
            switch ("a" + "foo bar") {
            case "\n\t":
                v6 = (v6 * v6)
            case "hello":
                v9 -= v9
            case "a":
                v22 = (v22 - -2.25)
            default:
                var v23 int = (^(fi21))
                _ = v23
            }
            v24 := GenS1{F0: -9223372036854775808, F1: 120}
            _ = fi21
            _ = v22
            _ = v24
        }
    }
    fmt.Println(v5, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
