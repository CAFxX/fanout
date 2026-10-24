package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 int = 32767
    var v3 uint32 = 450
    var v4 int64 = 42
    var v5 string = "\n\t"
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: (^(v1)), F1: 0, F2: 116}
        _ = fi7
        _ = v8
    }
    var v9 bool = (-10 <= -1)
    v1 += v1
    v9 = (17 != 2147483647)
    v10 := GenS1{F0: (^(v1)), F1: int16(v3), F2: 127}
    var v11 uint32 = (v3 ^ (v3 / ((v3) | 1)))
    v12 := GenS1{F0: v1, F1: -1, F2: 114}
    switch v5 {
    case "hello":
        v11 = v3
    default:
        v3 += (v11 - 2)
    }
    fmt.Println(v1, v1, v11)
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{(v2 % ((0) | 1)), (-(fi14)), fi14}
        switch 2147483647 {
        case -2147483648:
            v4 = (v4 - 42)
        case 2147483647:
            fmt.Println(v3, v4)
        }
        v3 = v3
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 65535, F1: -1, F2: 42}
    fmt.Println(v5, v3)
    fmt.Println(v11, v5)
    var v17 float64 = 1.5
    var v18 bool = (42 == 42)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
