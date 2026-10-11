package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 uint16) int64 {
    v6 := []int{4294967295, 32767}
    if len(v6) > 0 {
        v6[0] = 32768
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]uint32)
    var v3 uint32 = 4294967295
    var v4 int32 = (-(v1))
    var v5 bool = false
    v3 = (v3 << ((18446744073709551615) & 63))
    v1 = (v1 - (^(v4)))
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: 42}
        v9 := []int{}
        if false {
            _ = len(v2)
            v1 -= v1
            if (42 <= 694) {
                var v10 float64 = 1e10
                _ = v10
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = (-(v4))
    v1 = 1
    fmt.Println(v5, v1)
    v3 -= 65535
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v4, v4, v3)
        _ = fi11
    }
    if true {
        var v12 int8 = -106
        v3 = (v3 % (((v3 | v3)) | 1))
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{fi13, fi13}
            v4 = 2147483647
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
