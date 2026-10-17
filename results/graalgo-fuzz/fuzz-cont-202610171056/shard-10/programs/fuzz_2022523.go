package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
func genFn0(p0 uint) int8 {
    var v5 uint64 = 505
    v6 := []int{331, -2147483648, -205}
    v7 := GenS1{F0: 134, F1: (p0 - p0), F2: 1}
    switch -402 {
    case 511:
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 72
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 242
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    delete(v3, 128)
    v2 = v2
    delete(v4, "hello")
    var v9 string = "a"
    var v10 string = v9
    fmt.Println(v10, v2)
    v11 := []int{-647}
    v12 := GenS1{F0: (v2 >> ((425) & 63)), F1: 18446744073709551615, F2: 1}
    if ((3.14159 / (1.5 / 1.0)) == 1.5) {
        v13 := GenS2{F0: 42, F1: 0, F2: 42}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 float32 = (0.0 - 0.0)
                _ = fi15
                _ = v16
            }
            v10 = "a"
            _ = fi14
        }
        fmt.Println(v9, v10, v10)
        fmt.Println(v10)
    }
    var v17 uint16 = 553
    v18 := GenS1{F0: v2, F1: 4294967295, F2: 127}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
