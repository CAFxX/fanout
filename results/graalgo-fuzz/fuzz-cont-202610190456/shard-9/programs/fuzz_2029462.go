package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint8, p1 bool, p2 bool) int8 {
    if (false && (0 != (p0 >> ((18446744073709551615) & 63)))) {
        p0 -= (p0 >> ((844) & 63))
    } else {
        p1 = ((("\n\t" == ("a" + "a")) || p1) || p2)
    }
    p2 = (true || p2)
    var v5 uint8 = p0
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            fmt.Println(p2, p1)
        }
        _ = fi6
    }
    v7 := GenS1{F0: 1, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 141
    var v8 int = v2
    v2 = v2
    v9 := []int{899}
    fmt.Println(v8, v2)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint = 953
        v8 -= fi10
        _ = fi10
        _ = v11
    }
    if (0 > 9223372036854775807) {
        v8 = -187
    }
    var v12 float32 = float32(((1.5 - 3.14159) - 0.0))
    fmt.Println(v4, v12, v12)
    var v13 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
