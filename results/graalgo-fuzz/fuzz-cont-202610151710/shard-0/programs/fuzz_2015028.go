package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 uint16 = 0
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        switch fi6 {
        case 127:
            var v7 int = (fi6 / ((fi6) | 1))
            _ = v7
        case 128:
            var v8 bool = false
            _ = v8
        default:
            v9 := GenS1{F0: -128}
            _ = v9
        }
        _ = fi6
    }
    var v10 bool = (!(false && true))
    v2[-205] = -1
    v1 += v1
    v1 = v1
    fmt.Println(v10, v1)
    fmt.Println(v1, v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int32 = -2147483648
        _ = fi11
        _ = v12
    }
    var v13 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
