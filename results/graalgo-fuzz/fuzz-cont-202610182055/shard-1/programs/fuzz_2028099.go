package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint64 = 2
    var v2 string = ""
    v3 := make(map[int]uint32)
    if (255 >= 996) {
        v1 -= v1
        v4 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
        if false {
            v1 += 0
            var v5 int = 127
            fmt.Println(v5)
            _ = v5
        }
        _ = v4
    } else {
        v3[42] = 0
        fmt.Println(v2)
        v3[0] = 181
    }
    v6 := []int{4294967295, 32768, 128, 128}
    var v7 bool = (float64(0.1) >= float64(1.5))
    v1 = (v1 - v1)
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = v2
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 float64 = 0.1
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 int8 = 9
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
