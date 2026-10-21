package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 uint8 = 1
    v5 := []int{2147483647, 85}
    var v6 float32 = 0.1
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (fi7 ^ fi7) {
        case 1:
            var v8 int8 = 127
            _ = v8
        case -919:
            var v9 int32 = 1
            _ = v9
        case 2147483647:
            var v10 uint = 510
            _ = v10
        }
        _ = fi7
    }
    v6 -= v6
    var v11 uint = 818
    var v12 uint64 = 2
    v2[-142] = 127
    fmt.Println(v11, v11, v6)
    delete(v2, 42)
    v12 -= v12
    fmt.Println(v4)
    delete(v2, 227)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
