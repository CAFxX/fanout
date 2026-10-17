package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }

func main() {
    var v1 uint = 42
    var v2 float32 = 3.14159
    var v3 int = 65
    v4 := make(map[int]string)
    var v5 string = "a"
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 0, F1: -1}
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{(v3 + (fi8 * fi8)), (v3 << ((v1) & 63)), (^(fi8)), fi8}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 0
    v11 := []int{(v3 & 9223372036854775807), (v3 * v3), v3, (^(v3))}
    switch (v3 - v3) {
    case 32767:
        var v12 int64 = 1
        _ = v12
    case 0:
        var v13 uint32 = 1
        _ = v13
    }
    v14 := GenS1{F0: v10, F1: v10}
    switch ("\n\t" + (v5 + v5)) {
    case "\n\t":
        switch (v3 | v3) {
        case 4294967295:
            switch -754 {
            case 32768:
                var v15 uint16 = 65535
                _ = v15
            case -1:
                fmt.Println(v10)
            }
        default:
            v16 := GenS2{F0: 0, F1: -1, F2: (v1 * v1)}
            _ = v16
        }
    case "hello":
        var v17 bool = false
        _ = v17
    case "":
        v18 := GenS2{F0: 127, F1: 42, F2: (v1 * v1)}
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v2 += (v2 + 0.1)
        var v20 float64 = -2.25
        var v21 string = v5
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
