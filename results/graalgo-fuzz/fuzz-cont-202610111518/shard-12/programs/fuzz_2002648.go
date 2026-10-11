package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 string) bool {
    var v4 int64 = 42
    fmt.Println(v4, v4, v4)
    v5 := []int{32767, 1, 2147483647, -401}
    if len(v5) > 0 {
        v5[0] = -1
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 || true)
}
func genFn1() uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = 1
        v7 = -1
        var v8 uint = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = ""
    var v10 int64 = 42
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    if false {
        v3 += (v3 << ((178) & 63))
        var v11 int = 9223372036854775807
        _ = v11
    }
    v3 += 1
    var v12 float64 = float64(1e10)
    v12 -= 0.1
    var v13 float32 = 1.5
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int32 = v3
        v3 = v3
        delete(v2, 2)
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    fmt.Println(v12, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
