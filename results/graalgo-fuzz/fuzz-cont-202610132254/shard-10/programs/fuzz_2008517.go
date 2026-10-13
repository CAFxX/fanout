package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 string, p2 int) uint16 {
    var v5 uint32 = 669
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6, p1, p1)
        if (42 >= 4294967295) {
            v7 := []int{(p2 + (fi6 & (-(fi6))))}
            fmt.Println(p0, v5)
            _ = v7
        } else {
            v5 = 65535
        }
        _ = fi6
    }
    p2 += (p2 % ((p2) | 1))
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return p0
}

func main() {
    v1 := []int{42, 1, 42}
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{65535, 4294967295}
    var v9 int32 = -777
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(fi10, v9)
        _ = fi10
    }
    v9 -= (^(v9))
    var v11 float32 = 0.1
    switch "a" {
    case "a":
        var v12 bool = (v11 == (1e10 / v11))
        _ = v12
    case "":
        v9 += -2147483648
    case "\n\t":
        v9 -= v9
    default:
        v11 -= v11
    }
    v11 = v11
    if false {
        v11 += v11
        v11 += v11
    }
    var v13 uint8 = 222
    var v14 string = "foo bar"
    var v15 int16 = -1
    fmt.Println(v13, v15)
    if len(v4) > 0 {
        v4[0] = 65535
    }
    v14 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
