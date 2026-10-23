package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0(p0 bool, p1 int32, p2 bool) uint {
    var v6 uint64 = 0
    p2 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 868
}

func main() {
    var v1 uint16 = 42
    v2 := []int{9223372036854775807, 2147483647, 763}
    v3 := make(map[uint32]string)
    var v4 int = -51
    v5 := make(map[uint32]uint32)
    v7 := GenS1{F0: 282, F1: false, F2: 4294967295}
    if false {
        var v8 int64 = -1
        fmt.Println(v8)
        _ = v8
    } else {
        if (false || false) {
            v9 := GenS1{F0: 0, F1: (false || (695 != 720)), F2: 2147483648}
            var v10 uint64 = 4294967295
            _ = v10
            _ = v9
        }
        v4 += (v4 >> ((874) & 63))
    }
    _ = len(v2)
    v5[4294967295] = 4294967295
    fmt.Println(v4)
    v1 = v1
    var v11 uint32 = 1
    v12 := v5[v11]
    v11 = 42
    var v13 int16 = 1
    v4 += 2147483647
    var v14 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
