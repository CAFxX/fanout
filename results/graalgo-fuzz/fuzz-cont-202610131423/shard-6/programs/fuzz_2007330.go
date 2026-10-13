package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 4294967295
    v2 := []string{}
    var v3 int8 = 127
    v4 := make(map[string]uint32)
    var v5 uint8 = 255
    fmt.Println(v1)
    var v6 float32 = (0.1 - (0.0 + 1e10))
    if true {
        v6 -= 1e10
        var v7 int8 = v3
        _ = v7
    } else {
        v6 = (0.0 / ((v6 - 1e10) * v6))
        for fi8 := 0; fi8 < 6; fi8++ {
            v2 = append(v2, "foo bar")
            var v9 float64 = 0.0
            var v10 int16 = -256
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(v5, v6, v6)
    }
    v1 = 42
    fmt.Println(v1, v1)
    v3 += (v3 + v3)
    var v11 string = "\n\t"
    var v12 uint16 = 42
    if (v12 > v12) {
        fmt.Println(v1)
    }
    var v13 bool = (false && false)
    var v14 float64 = 0.0
    var v15 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
