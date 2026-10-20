package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    v5 := []int{}
    var v6 uint64 = 1
    var v7 float32 = (-2.25 + (3.14159 / 1e10))
    switch 0 {
    case 32768:
        v8 := v3[-9223372036854775808]
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint8 = 0
        _ = fi9
        _ = v10
    }
    v6 = v6
    for fi11 := 0; fi11 < 6; fi11++ {
        if len(v5) > 0 {
            v5[0] = 1
        }
        _ = len(v4)
        var v12 int8 = 1
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v7 = v7
        _ = fi13
    }
    switch 550 {
    case 32768:
        v14 := v1[2147483648]
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 4; fi15++ {
            v7 = v7
            fmt.Println(v6)
            v16 := GenS1{F0: 1, F1: "foo bar", F2: 1}
            _ = fi15
            _ = v16
        }
    case -1:
        var v17 uint = 65535
        _ = v17
    default:
        if (-369 >= 42) {
            if (!false) {
                v18 := GenS1{F0: 1, F1: "hello", F2: 42}
                _ = v18
            } else {
                var v19 uint8 = 255
                var v20 float64 = -2.25
                _ = v19
                _ = v20
            }
            v6 = v6
            var v21 string = "a"
            _ = v21
        }
    }
    var v22 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
