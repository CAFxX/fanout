package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32) int8 {
    v7 := []int{-2147483648}
    p0 -= (p0 & (p0 * p0))
    var v8 uint16 = 65535
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 += v8
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 2147483647
    v2 := []uint32{671, 112, 2, 2147483648}
    var v3 uint64 = 421
    var v4 uint16 = 1
    v5 := make(map[uint32]int)
    v6 := make(map[string]string)
    v9 := GenS1{F0: -979}
    v4 = v4
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := v6[("foo bar" + "\n\t")]
        var v13 int16 = 0
        v12 = v12
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 299}
    v5[244] = (v1 + (v1 + -1))
    var v15 int32 = (v10 >> ((759) & 63))
    fmt.Println(v15, v4, v4)
    fmt.Println(v10, v10)
    var v16 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
