package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    v3 := []int{4294967295, 233}
    var v4 uint8 = 12
    var v5 uint = 348
    if (1 == 951) {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 float64 = -2.25
                var v9 uint8 = (^(v4))
                fmt.Println(fi6, v9)
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        if true {
            fmt.Println(v2, v4)
        }
    }
    v2 = ""
    v10 := GenS1{F0: 42}
    v5 = (v5 * 65535)
    if true {
        var v11 float32 = 0.1
        _ = v11
    }
    if true {
        fmt.Println(v4, v4, v5)
    }
    fmt.Println(v5)
    v5 = 845
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
