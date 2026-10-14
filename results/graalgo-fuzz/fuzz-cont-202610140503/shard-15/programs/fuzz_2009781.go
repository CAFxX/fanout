package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 0
    var v2 uint = 65535
    var v3 float32 = (((0.0 / 1.5) + 0.1) / -2.25)
    var v4 uint16 = v1
    var v5 string = "\n\t"
    v6 := []string{v5, "\n\t"}
    var v7 int8 = 1
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int64 = -933
        _ = fi8
        _ = v9
    }
    v2 += (v2 - 255)
    var v10 string = ((v5 + "hello") + ("a" + "hello"))
    v10 = v10
    fmt.Println(v1, v1, v7)
    if len(v6) > 0 {
        v6[0] = (v5 + v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
