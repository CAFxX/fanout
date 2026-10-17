package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 int = 32767
    var v3 uint64 = 347
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(v3, v2, v3)
        fmt.Println(v2)
        fmt.Println(v2, v2)
        _ = fi4
    }
    if true {
        v3 = 18446744073709551615
        v5 := []int{32768, (v2 >> ((18446744073709551615) & 63)), v2, v2}
        _ = v5
    }
    v6 := []int{(-(v2))}
    v3 -= 0
    switch v2 {
    case -934:
        v3 += (v3 + (v3 >> ((107) & 63)))
    case 4294967295:
        v3 = (v3 & (v3 << ((4294967295) & 63)))
    case 9223372036854775807:
        v3 += (v3 % ((2) | 1))
    default:
        if false {
            if (false && (!false)) {
                var v7 int32 = -2147483648
                fmt.Println(v3, v3)
                _ = v7
            } else {
                _ = len(v6)
            }
            fmt.Println(v2)
        }
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: -1, F1: "a"}
        _ = len(v1)
        _ = fi8
        _ = v9
    }
    _ = len(v6)
    var v10 bool = true
    v3 += (v3 % ((v3) | 1))
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint8 = 162
        v13 := GenS1{F0: -884, F1: ""}
        var v14 int64 = 0
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v6 = append(v6, (^(v2)))
    _ = len(v1)
    for fi15 := 0; fi15 < 6; fi15++ {
        v3 += v3
        v16 := GenS1{F0: -238, F1: ""}
        _ = fi15
        _ = v16
    }
    fmt.Println(v3, v2)
    var v17 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
