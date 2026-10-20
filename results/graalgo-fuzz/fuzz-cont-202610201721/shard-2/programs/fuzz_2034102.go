package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }

func main() {
    var v1 string = "\n\t"
    var v2 string = ""
    v3 := make(map[int]string)
    var v4 string = ""
    v5 := []int{1, 9223372036854775807}
    v2 = v1
    v4 = (((v1 + v2) + v4) + v4)
    v1 = v2
    fmt.Println(v4, v2, v2)
    for fi6 := 0; fi6 < 4; fi6++ {
        v4 = "\n\t"
        v1 = (v4 + (v2 + v4))
        _ = fi6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v4 = v1
            var v9 uint = 871
            _ = fi8
            _ = v9
        }
        var v10 int64 = 9223372036854775807
        _ = fi7
        _ = v10
    }
    _ = len(v5)
    v11 := GenS1{F0: 0, F1: 127, F2: 474}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
