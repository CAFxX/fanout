package main

import (
    "fmt"
)

func genFn0(p0 int, p1 string, p2 int16) bool {
    v7 := []int{}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ((true || ("hello" > ("hello" + p1))) && (p1 <= ""))
}

func main() {
    var v1 int64 = 1
    var v2 int32 = -1
    v3 := []string{"\n\t"}
    var v4 uint64 = 723
    v5 := make(map[uint32]uint32)
    var v6 int16 = -1
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            _ = len(v5)
            _ = fi9
        }
        _ = fi8
    }
    v4 = (v4 >> ((0) & 63))
    delete(v5, 657)
    var v10 int = 977
    fmt.Println(v6, v10)
    var v11 float32 = ((1.5 * float32(0.0)) / 1.0)
    v6 += (v6 & (v6 >> ((969) & 63)))
    fmt.Println(v11)
    var v12 int8 = -64
    var v13 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
