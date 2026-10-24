package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("hello" + ("foo bar" + "foo bar"))
    v3 := []uint32{447}
    var v4 int16 = 42
    v5 := []int{0}
    v6 := []int{int(v4)}
    v4 -= v4
    switch v2 {
    case "hello":
        var v7 uint16 = 0
        _ = v7
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            v1[779] = 0
            _ = fi8
        }
    }
    var v9 int32 = -110
    var v10 int16 = v4
    _ = len(v6)
    v4 += 0
    v10 += v4
    v9 += 42
    var v11 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
