package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 2
    var v2 uint = 456
    v3 := []int{1, 20, 42, 1}
    v4 := make(map[int]string)
    var v5 int = -9223372036854775808
    var v6 float64 = (0.0 * 1.5)
    var v7 string = "foo bar"
    if ((0 != 1) && (v2 > (v2 - (v2 * v2)))) {
        var v8 uint32 = 2147483648
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint = 18446744073709551615
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v7 = v7
    v1 -= v1
    for fi11 := 0; fi11 < 17; fi11++ {
        v2 += v2
        for fi12 := 0; fi12 < 18; fi12++ {
            v1 -= (v1 / ((1) | 1))
            fmt.Println(fi12)
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v1 = v1
    var v14 float32 = (float32(v6) + (-2.25 - 0.1))
    delete(v4, (v5 & 572))
    v4[v5] = v7
    var v15 float64 = (0.0 + 1.5)
    fmt.Println(v1, v2, v6)
    for fi16 := 0; fi16 < 20; fi16++ {
        v5 += v5
        _ = fi16
    }
    if true {
        var v17 uint = (^(v2))
        v5 += (-(v5))
        _ = v17
    }
    v18 := v4[v5]
    var v19 int8 = 42
    fmt.Println(v1)
    delete(v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
