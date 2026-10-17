package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    var v3 uint = 90
    v4 := make(map[uint32]int)
    var v5 bool = false
    var v6 int16 = 32767
    v7 := GenS1{F0: 753, F1: 2147483647, F2: "a"}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch "foo bar" {
        case "\n\t":
            v3 = 65535
        case "hello":
            for fi9 := 0; fi9 < 5; fi9++ {
                v5 = ((((3.14159 * 0.0) >= (-2.25 / 3.14159)) && (v5 || (2147483647 > 42))) && (570 <= 991))
                v10 := GenS1{F0: 784, F1: 42, F2: "foo bar"}
                _ = fi9
                _ = v10
            }
        default:
            fmt.Println(v6, v1)
        }
        v5 = true
        var v11 uint32 = 342
        _ = fi8
        _ = v11
    }
    v12 := GenS2{F0: uint16(v3), F1: -9223372036854775808, F2: 65535}
    var v13 string = "a"
    for fi14 := 0; fi14 < 15; fi14++ {
        v5 = v5
        v5 = true
        _ = fi14
    }
    v1 = v6
    v15 := []int{127, -2147483648, -427}
    switch v13 {
    case "hello":
        v16 := GenS1{F0: 1, F1: 1, F2: v13}
        _ = v16
    default:
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int32 = 0
            v5 = false
            _ = fi17
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
