package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -2147483648
    var v3 uint64 = 2
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    var v6 string = "\n\t"
    v7 := []int{}
    v8 := GenS1{F0: 42, F1: (v6 + v6)}
    var v9 bool = (((2147483648 < 2147483647) && (-2.25 <= 1e10)) || (2147483647 != 65535))
    var v10 uint64 = v3
    switch 566 {
    case 127:
        fmt.Println(v3, v6, v4)
    }
    v11 := GenS1{F0: 127, F1: "foo bar"}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v1[int(v3)] = v6
            _ = fi13
        }
        v3 = (v10 | v3)
        _ = fi12
    }
    fmt.Println(v3)
    var v14 int = 760
    if v9 {
        var v15 uint64 = (v3 ^ 4294967295)
        fmt.Println(v3)
        v9 = ((42 == 858) && true)
        _ = v15
    } else {
        var v16 int32 = -612
        _ = v16
    }
    v17 := GenS2{F0: -54, F1: -1}
    v7 = append(v7, (v14 / ((int(v2)) | 1)))
    fmt.Println(v4)
    v5[2] = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
