package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 15; fi4++ {
        fmt.Println(fi4)
        v5 := GenS1{F0: 2, F1: 4294967295}
        v6 := GenS1{F0: 1, F1: (fi4 ^ (fi4 - fi4))}
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        var v7 uint64 = 2
        _ = v7
    }
    if (42 <= 852) {
        v8 := []int{32768, 4294967295, 9223372036854775807}
        var v9 float64 = ((1e10 / 1e10) - -2.25)
        var v10 int = 2147483647
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 bool = true
    v11 = v11
    v11 = (982 >= 0)
    var v12 int64 = 861
    v11 = v11
    var v13 uint = 567
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int64 = v12
        v12 = v12
        if false {
            _ = len(v1)
            var v16 uint64 = 1
            _ = v16
        } else {
            if v11 {
                var v17 uint64 = 4294967295
                v18 := GenS1{F0: 33, F1: fi14}
                _ = v17
                _ = v18
            }
        }
        _ = fi14
        _ = v15
    }
    var v19 int16 = 0
    v12 += (v12 / ((v12) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
}
