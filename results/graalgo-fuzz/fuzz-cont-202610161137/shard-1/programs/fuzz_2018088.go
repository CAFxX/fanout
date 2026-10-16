package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 int32) uint {
    var v7 uint32 = 2147483647
    if (-107 >= 1) {
        fmt.Println(v7, v7, p0)
        fmt.Println(p1, v7, p1)
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                v9 := []int{128}
                _ = v9
            }
            var v10 int8 = 38
            _ = fi8
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 bool) uint8 {
    p1 = ((0.1 + 0.0) <= 1.5)
    if p1 {
        p0 += p0
        fmt.Println(p1)
    }
    p1 = true
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []uint32{2147483648}
    var v2 int32 = -2147483648
    var v3 int16 = 478
    var v4 float32 = 0.1
    v5 := []string{}
    var v6 int64 = -9223372036854775808
    v5 = append(v5, ((("foo bar" + "") + "") + "a"))
    var v11 bool = true
    if (float64(v4) >= 0.1) {
        var v12 bool = (((false || v11) && (!v11)) && v11)
        for fi13 := 0; fi13 < 1; fi13++ {
            v4 = (3.14159 * v4)
            _ = fi13
        }
        _ = v12
    }
    _ = len(v1)
    v2 += (v2 % ((v2) | 1))
    var v14 uint64 = 13
    v6 = v6
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
