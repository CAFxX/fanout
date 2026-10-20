package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v5 uint16 = 32767
    var v6 string = ("\n\t" + "")
    _ = v5
    _ = v6
    return (v5 << ((0) & 63))
}
func genFn1(p0 string, p1 uint32, p2 int8) int {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int8 = p2
        _ = fi7
        _ = v8
    }
    p1 = (p1 * 4294967295)
    v9 := []int{-9223372036854775808, -212, 32767, 128}
    var v10 int = 4294967295
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (v10 | -926)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (545 != 582)
    v3 := []int{}
    v4 := make(map[string]uint32)
    for fi11 := 0; fi11 < 5; fi11++ {
        v2 = ((2 <= 653) || (-1 <= 1000))
        _ = fi11
    }
    var v12 uint8 = 0
    var v13 int64 = -1
    var v14 uint64 = 617
    v14 -= 2
    var v15 uint8 = v12
    var v16 uint64 = (v14 + 0)
    v16 = (^(v16))
    var v17 float32 = 1e10
    _ = len(v3)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
