package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 1
    var v3 int8 = (v2 << ((476) & 63))
    v4 := make(map[uint32]int)
    v2 = (v3 / ((v2) | 1))
    switch "" {
    case "":
        v2 = (v3 | 103)
    }
    var v5 bool = true
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: 42, F1: 127, F2: 0}
    var v8 float64 = float64(1e10)
    var v9 int16 = -1
    fmt.Println(v3, v8, v2)
    switch -862 {
    case 0:
        if (!(0.1 <= 1e10)) {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v5)
                _ = fi10
            }
            fmt.Println(v1, v9, v2)
            v11 := v4[4294967295]
            _ = v11
        }
    case 128:
        switch 32768 {
        case -2147483648:
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS2{F0: 4294967295, F1: (fi12 % ((32768) | 1)), F2: -2147483648}
                v3 += -1
                _ = fi12
                _ = v13
            }
        case 128:
            v14 := []int{-9223372036854775808, 1, 32768, -2147483648}
            _ = v14
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
