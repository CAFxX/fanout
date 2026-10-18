package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 uint32) int64 {
    var v6 float64 = ((0.1 * 1e10) - 1e10)
    var v7 float64 = v6
    v8 := []int{32767, 989, -949, -2147483648}
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 82
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 127
    v3 := make(map[int]bool)
    var v4 int64 = -1
    var v5 int64 = (v4 ^ v4)
    for fi9 := 0; fi9 < 5; fi9++ {
        delete(v3, fi9)
        _ = fi9
    }
    switch 42 {
    case 9223372036854775807:
        var v10 uint32 = 65535
        _ = v10
    case -1:
        v2 += v2
    case 253:
        v5 -= (v4 / (((v5 >> ((18446744073709551615) & 63))) | 1))
    }
    fmt.Println(v4, v4)
    var v11 int64 = -466
    var v12 bool = ((255 > 906) || false)
    var v13 uint16 = 42
    fmt.Println(v4, v2)
    v14 := []int{2147483647}
    delete(v3, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
