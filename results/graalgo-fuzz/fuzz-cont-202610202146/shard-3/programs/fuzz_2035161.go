package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int8 = 42
    var v2 string = "hello"
    var v3 int32 = -1
    var v4 int8 = (-(v1))
    v5 := make(map[int]uint32)
    v1 = v1
    var v6 bool = ((1.5 * float64(-2.25)) != float64((0.0 - ((-2.25 * ((float32(1.5) / 1.5) / 0.1)) - (float32((1e10 * 1.5)) / 0.1)))))
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = 844
        if v6 {
            v2 = ""
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v1)
    if (v6 && ((0.0 * (3.14159 / ((1.5 / (1.5 * 3.14159)) * (((-2.25 - 0.0) * -2.25) + 1e10)))) >= 0.1)) {
        var v9 uint64 = 2
        var v10 string = v2
        v1 = 0
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v1, v3)
            if v6 {
                v2 = v2
                v12 := []int{fi11, (-(fi11)), -701}
                _ = v12
            } else {
                v13 := GenS1{F0: 1, F1: v3}
                _ = v13
            }
            _ = fi11
        }
    }
    delete(v5, 127)
    if true {
        v3 -= (v3 % ((42) | 1))
        v14 := GenS1{F0: 2, F1: v3}
        v5[9223372036854775807] = 65535
        _ = v14
    }
    v6 = (v6 && (v6 || (!(v3 >= (-(v3))))))
    switch v2 {
    case "a":
        var v15 float64 = (1.5 + -2.25)
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 20; fi16++ {
            v3 -= v3
            if v6 {
                var v17 float32 = -2.25
                v5[(fi16 | fi16)] = 65535
                v18 := GenS2{F0: 4294967295}
                _ = v17
                _ = v18
            }
            v2 = v2
            _ = fi16
        }
    default:
        v19 := []int{-1, 32767, -9223372036854775808, -2147483648}
        _ = v19
    }
    v4 = 0
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
