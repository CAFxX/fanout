package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    var v4 int64 = (v1 % (((^(v1))) | 1))
    var v5 uint64 = 42
    v3 = v3
    v6 := []int{int(v4), 32768, 4294967295}
    v1 = 9223372036854775807
    v4 -= (^(v1))
    v5 += (^(v5))
    v4 += v1
    var v7 uint8 = 126
    if (true && (!(2 == 153))) {
        var v8 int = -750
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int32 = 994
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 434
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
