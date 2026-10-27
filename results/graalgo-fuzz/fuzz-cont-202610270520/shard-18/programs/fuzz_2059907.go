package main

import (
    "fmt"
)

func genFn0(p0 int8) int {
    p0 = p0
    var v6 int8 = 1
    _ = p0
    _ = v6
    return -231
}
func genFn1() int64 {
    var v7 bool = true
    v8 := []int{-453, 32767, 42, -2147483648}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 uint = 18446744073709551615
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 989
    var v2 uint8 = 42
    v3 := []uint32{690, 4294967295}
    var v4 int32 = 279
    v5 := make(map[int]uint32)
    var v10 int64 = 999
    var v11 int8 = 17
    var v12 uint = (v1 % ((v1) | 1))
    v4 -= (-(v4))
    fmt.Println(v1)
    for fi13 := 0; fi13 < 11; fi13++ {
        v11 -= (v11 + int8(v10))
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint16 = 0
            _ = fi14
            _ = v15
        }
        fmt.Println(v12, v4)
        _ = fi13
    }
    fmt.Println(v11, v11)
    if false {
        var v16 uint32 = 869
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                fmt.Println(v10)
                var v19 uint = (^(v12))
                fmt.Println(v4, v19)
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
