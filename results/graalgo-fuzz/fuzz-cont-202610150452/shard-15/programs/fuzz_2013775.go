package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 407
    var v2 string = "\n\t"
    var v3 int32 = -954
    var v4 uint = 821
    var v5 uint16 = 665
    v6 := make(map[uint32]bool)
    fmt.Println(v4, v2)
    _ = len(v6)
    switch v2 {
    case "foo bar":
        fmt.Println(v1, v4)
    default:
        if false {
            v1 = (v1 / ((32767) | 1))
        } else {
            if ((100 <= -1) || false) {
                var v7 int8 = -85
                _ = v7
            }
        }
    }
    var v8 float64 = float64(float32(3.14159))
    var v9 uint16 = v5
    var v10 int32 = v3
    var v11 uint32 = 1
    switch v2 {
    case "":
        v3 = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
