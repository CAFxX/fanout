package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 int64) int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = p0
        _ = fi7
        _ = v8
    }
    p0 = (-(p0))
    _ = p0
    _ = p1
    return (p0 + 127)
}

func main() {
    var v1 uint = 65535
    var v2 bool = (100 >= -125)
    var v3 int = 26
    var v4 uint = 4294967295
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    var v9 uint16 = 32767
    v4 = (v4 + (v4 >> (((v1 - 333)) & 63)))
    var v10 float32 = float32(0.0)
    for fi11 := 0; fi11 < 10; fi11++ {
        v3 += (fi11 * 4294967295)
        var v12 uint = v4
        _ = fi11
        _ = v12
    }
    v13 := []int{v3}
    v14 := v6["a"]
    fmt.Println(v3)
    var v15 float64 = 3.14159
    if v2 {
        switch v3 {
        case -1:
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 int8 = 42
                _ = fi16
                _ = v17
            }
        case 292:
            var v18 bool = v2
            _ = v18
        default:
            var v19 uint16 = (v9 & (v9 % ((v9) | 1)))
            _ = v19
        }
        v10 = v10
        v9 -= v9
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
