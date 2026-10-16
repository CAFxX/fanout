package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    v4 := GenS1{F0: 1, F1: 2147483647}
    v5 := []int{128, 42}
    var v6 string = ""
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            var v8 float32 = -2.25
            _ = v8
        } else {
            v9 := GenS1{F0: 255, F1: (^(fi7))}
            _ = v9
        }
        v10 := GenS2{F0: 1, F1: 72, F2: (fi7 >> ((18446744073709551615) & 63))}
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "" {
        case "\n\t":
            var v12 uint32 = 2147483647
            _ = v12
        default:
            fmt.Println(v3)
        }
        _ = fi11
    }
    v1 += v1
    v6 = "a"
    var v13 int64 = -1
    fmt.Println(v1, v1, v1)
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS3{F0: v1, F1: (1000 == -32768)}
        v16 := GenS3{F0: (v1 - 92), F1: false}
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
