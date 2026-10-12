package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 15; fi6++ {
        v5 = (v5 + (v5 * (^(v5))))
        var v7 int16 = 32767
        _ = fi6
        _ = v7
    }
    if false {
        v5 = (^(v5))
    } else {
        v5 = v5
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v5 -= (v5 & (v5 << ((2) & 63)))
        v9 := v3[1]
        v4 = append(v4, 49)
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 1
    v5 -= (v5 / (((v5 | 42)) | 1))
    _ = len(v4)
    var v11 int64 = 0
    fmt.Println(v5)
    switch 42 {
    case 127:
        fmt.Println(v11, v5)
    default:
        if false {
            for fi12 := 0; fi12 < 5; fi12++ {
                v2[2147483647] = 2
                var v13 string = "foo bar"
                _ = fi12
                _ = v13
            }
            v11 = (v11 & (v11 << ((42) & 63)))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
