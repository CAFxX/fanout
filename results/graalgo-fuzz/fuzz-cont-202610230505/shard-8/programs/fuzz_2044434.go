package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2
    for fi4 := 0; fi4 < 17; fi4++ {
        v1 -= v1
        _ = fi4
    }
    var v5 uint32 = v3
    v6 := []int{-9223372036854775808, 42, 2147483647}
    var v7 float32 = (1e10 - 3.14159)
    v1 -= (v1 - 18446744073709551615)
    v7 = 0.0
    if false {
        v3 -= v3
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint = 1
                _ = fi9
                _ = v10
            }
            delete(v2, v3)
            var v11 float32 = v7
            _ = fi8
            _ = v11
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
