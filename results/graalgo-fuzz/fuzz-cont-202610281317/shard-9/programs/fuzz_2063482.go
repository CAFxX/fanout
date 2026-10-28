package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 59
    var v2 uint32 = 556
    v3 := make(map[int]string)
    var v4 uint32 = v2
    var v5 int32 = -730
    var v6 int64 = -9223372036854775808
    v2 = v2
    v6 -= (v6 * v6)
    var v7 int8 = 100
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            v6 -= (v6 & (v6 >> ((18446744073709551615) & 63)))
            if (-32768 <= 32767) {
                var v10 bool = false
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    v5 = v5
    var v11 bool = false
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{fi12, fi12, (fi12 << ((42) & 63))}
        var v14 uint64 = 42
        var v15 bool = v11
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v6 += (v6 << ((18446744073709551615) & 63))
    var v16 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
