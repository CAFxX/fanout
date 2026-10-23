package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 int64) int {
    v6 := []int{-9223372036854775808}
    if (!true) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v7 := GenS1{F0: 2147483647, F1: -1}
        if (42 > -1) {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 string = ""
                _ = fi8
                _ = v9
            }
        }
        _ = v7
    } else {
        switch 128 {
        case 128:
            p0 = p0
        case 42:
            v10 := GenS1{F0: 1, F1: -1}
            _ = v10
        }
    }
    _ = p0
    _ = v6
    return int(p0)
}

func main() {
    var v1 int64 = -1
    var v2 uint = 2
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 int8 = 110
    fmt.Println(v2)
    var v11 int64 = -70
    v1 += 42
    fmt.Println(v2, v11, v5)
    v2 = v2
    v4[int(v1)] = 65535
    fmt.Println(v1)
    if true {
        fmt.Println(v2)
    } else {
        var v12 uint32 = 2
        _ = v12
    }
    v1 -= v11
    fmt.Println(v11, v11, v5)
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v2, v11)
        _ = fi13
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
