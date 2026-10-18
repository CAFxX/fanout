package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 1000
    var v2 uint8 = 42
    var v3 string = "a"
    var v4 int16 = v1
    var v5 uint16 = 65535
    var v6 uint = 423
    fmt.Println(v5)
    v1 = (v1 / ((v4) | 1))
    var v7 int16 = v1
    switch 128 {
    case 127:
        var v8 int = 766
        _ = v8
    case 2:
        v4 = v1
    case 9223372036854775807:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 float64 = float64(0.1)
            var v11 int8 = 127
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 bool = (!(!true))
    fmt.Println(v7)
    v13 := []int{0, 32767, -9223372036854775808}
    switch "" {
    case "\n\t":
        var v14 uint32 = 1
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
