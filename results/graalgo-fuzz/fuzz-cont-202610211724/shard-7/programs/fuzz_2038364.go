package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 1
    var v2 int16 = 0
    var v3 uint = 0
    v3 += uint(v1)
    v4 := []int{2147483647, -2147483648, 1, 879}
    switch 0 {
    case 9223372036854775807:
        v2 = v2
    }
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    if true {
        v3 = 0
        fmt.Println(v1)
        var v5 int = -165
        _ = v5
    }
    v2 += (^(v2))
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
