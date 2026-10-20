package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 6
    v3 := make(map[uint32]bool)
    var v4 uint8 = 5
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: (1 >= 9223372036854775807), F1: "a"}
        var v8 int = (^(fi6))
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v4)
    fmt.Println(v2, v4)
    var v9 bool = true
    fmt.Println(v2, v9)
    v9 = v9
    if (0 >= 1) {
        if (-9223372036854775808 < 42) {
            fmt.Println(v4, v9, v2)
            if false {
                v9 = v9
            }
            var v10 uint8 = (v4 >> ((201) & 63))
            _ = v10
        }
        v4 += (^(v2))
    } else {
        v4 = v4
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint16 = 0
            _ = fi11
            _ = v12
        }
        v13 := []int{-1}
        _ = v13
    }
    var v14 uint8 = v2
    if false {
        v2 = v4
    }
    switch 4294967295 {
    case 32767:
        v2 = v14
    case -904:
        var v15 float32 = ((1.5 / 3.14159) * float32(1.5))
        _ = v15
    case 128:
        v9 = (-1 <= 662)
    default:
        v2 += 1
    }
    v14 -= v4
    v5[42] = 2
    v2 += 128
    delete(v5, 128)
    v9 = (42 >= 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
