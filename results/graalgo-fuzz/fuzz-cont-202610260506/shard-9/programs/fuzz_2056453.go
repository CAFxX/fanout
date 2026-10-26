package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int, p2 uint16) int64 {
    p2 -= (p2 * 1)
    fmt.Println(p0, p0, p0)
    p1 -= (p1 | p1)
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 int = 4294967295
    var v2 int = -252
    var v3 uint8 = 127
    if (true || (false || (!(!((183 == 4294967295) && (true || true)))))) {
        var v4 uint8 = 0
        fmt.Println(v4)
        _ = v4
    }
    var v5 float32 = 0.1
    v6 := []int{}
    switch v2 {
    case 0:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v7 := GenS1{F0: 612}
        _ = v7
    }
    v3 -= v3
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v5, fi8, v2)
        _ = fi8
    }
    var v9 string = "foo bar"
    v1 = (v2 - -749)
    if false {
        for fi10 := 0; fi10 < 2; fi10++ {
            if true {
                var v11 float32 = v5
                _ = v11
            } else {
                v12 := GenS2{F0: 4294967295}
                _ = v12
            }
            v3 += (v3 & 127)
            _ = fi10
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = genFn0(2, v2, 455)
        _ = fi13
        _ = v14
    }
    var v15 uint32 = 42
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
