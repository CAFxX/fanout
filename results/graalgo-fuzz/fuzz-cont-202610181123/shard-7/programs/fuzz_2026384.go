package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 uint8 = 173
    v4 := make(map[int]bool)
    v5 := []int{}
    var v6 uint32 = 779
    _ = len(v5)
    var v7 int32 = 28
    fmt.Println(v3)
    if (1e10 == 0.1) {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: v7}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                v5 = append(v5, -669)
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    } else {
        _ = len(v5)
    }
    fmt.Println(v3, v6)
    v3 += 127
    v3 = (v3 & v3)
    for fi11 := 0; fi11 < 1; fi11++ {
        _ = len(v5)
        var v12 uint8 = (v3 ^ (v3 << ((65535) & 63)))
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 1}
    v3 = v3
    v7 += v7
    var v14 uint32 = (v6 * 1)
    var v15 uint = 18446744073709551615
    fmt.Println(v15, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
