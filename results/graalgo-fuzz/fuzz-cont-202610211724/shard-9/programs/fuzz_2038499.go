package main

import (
    "fmt"
)


func main() {
    var v1 uint = 255
    v2 := make(map[uint32]string)
    var v3 uint64 = 468
    v3 += 18446744073709551615
    v1 += (v1 + (v1 >> (((v1 % ((0) | 1))) & 63)))
    v1 -= (v1 << ((v1) & 63))
    fmt.Println(v3, v3)
    var v4 uint8 = 127
    v1 = 4294967295
    if (false && (1e10 <= -2.25)) {
        v3 = 4294967295
        var v5 bool = (891 < -32768)
        _ = v5
    }
    var v6 int64 = -665
    v4 -= v4
    var v7 int32 = 135
    v1 += v1
    var v8 uint = 930
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
