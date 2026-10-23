package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64) uint64 {
    fmt.Println(p0)
    p0 = p0
    p0 += p0
    v5 := []int{}
    p0 -= (p0 + p0)
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 int16, p1 bool, p2 uint16) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 = (p0 * (p0 ^ (p0 & int16(p2))))
        _ = fi6
    }
    v7 := []int{9223372036854775807, 32767, 0}
    var v8 bool = (0 != 0)
    var v9 uint8 = 127
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: uint(v9)}
            v8 = (p2 >= p2)
            _ = fi11
            _ = v12
        }
        if (181 <= (p2 % ((p2) | 1))) {
            v13 := GenS1{F0: 4294967295}
            fmt.Println(p1)
            _ = v13
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 506
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 1
    var v3 int32 = -2147483648
    var v4 uint = 18446744073709551615
    fmt.Println(v2)
    fmt.Println(v3, v2)
    v14 := GenS1{F0: v4}
    fmt.Println(v2, v4, v3)
    fmt.Println(v2, v2, v4)
    fmt.Println(v4, v2, v3)
    var v15 uint = v4
    if (true || true) {
        v15 = v15
        v15 = v15
    } else {
        v16 := []int{128}
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 bool = ((fi17 ^ (-(fi18))) == (fi18 / ((int(v3)) | 1)))
                _ = fi18
                _ = v19
            }
            var v20 uint = v15
            _ = fi17
            _ = v20
        }
        _ = v16
    }
    v21 := []int{104, 128}
    if true {
        if (!(18446744073709551615 == uint64(v15))) {
            v15 = 270
        } else {
            var v22 int32 = v3
            if false {
                v22 += v3
            }
            _ = v22
        }
        fmt.Println(v2, v3, v3)
        fmt.Println(v3, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
