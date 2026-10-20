package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    var v2 float64 = float64(3.14159)
    var v3 bool = false
    v4 := make(map[int]string)
    var v5 float64 = (3.14159 + (v2 - v2))
    var v6 float32 = 1e10
    v3 = (v3 && v3)
    fmt.Println(v6, v2)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = -1
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v1)
    v6 += v6
    if v3 {
        var v9 int8 = -128
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int32 = 904
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 int16 = 348
    if v3 {
        var v13 int8 = 127
        v12 -= (v12 & v12)
        var v14 int64 = 42
        _ = v13
        _ = v14
    } else {
        var v15 int16 = (v12 >> ((18446744073709551615) & 63))
        v12 = 375
        var v16 string = ((v1 + "foo bar") + (("\n\t" + v1) + v1))
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
