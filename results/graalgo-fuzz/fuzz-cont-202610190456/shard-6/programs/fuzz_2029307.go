package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (-9223372036854775808 == -9223372036854775808)
    var v3 uint16 = 2
    for fi4 := 0; fi4 < 19; fi4++ {
        if v2 {
            for fi5 := 0; fi5 < 11; fi5++ {
                var v6 int64 = 0
                _ = fi5
                _ = v6
            }
            v2 = v2
        }
        v7 := []int{32767, (fi4 * (fi4 % (((fi4 >> ((2) & 63))) | 1)))}
        _ = fi4
        _ = v7
    }
    v8 := []int{127, 32767}
    _ = len(v1)
    if true {
        var v9 bool = v2
        if v9 {
            fmt.Println(v3)
            var v10 bool = v2
            var v11 float64 = -2.25
            _ = v10
            _ = v11
        } else {
            var v12 uint = 2
            _ = v12
        }
        _ = v9
    }
    var v13 uint32 = 321
    v14 := GenS1{F0: 542}
    v15 := GenS1{F0: 2147483647}
    var v16 float32 = (0.1 - 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
