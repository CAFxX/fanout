package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 float64 = (0.0 - -2.25)
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    if (v3 < 1e10) {
        var v6 uint64 = 394
        var v7 float64 = 0.1
        _ = v6
        _ = v7
    } else {
        v1 = v1
    }
    v1 = v1
    fmt.Println(v1)
    v8 := GenS1{F0: -600, F1: 2147483647}
    v1 -= (v1 << ((1) & 63))
    v9 := []int{32767, -368, 32768}
    var v10 uint8 = 127
    _ = len(v2)
    if true {
        var v11 uint = 167
        for fi12 := 0; fi12 < 7; fi12++ {
            v2[("foo bar" + "hello")] = "foo bar"
            v9 = append(v9, (fi12 * (^(fi12))))
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint = (v11 & (v11 & 556))
                var v15 uint16 = 2
                v10 += v10
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        fmt.Println(v10, v11)
        _ = v11
    }
    v16 := v4[2147483647]
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v16, v3)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
