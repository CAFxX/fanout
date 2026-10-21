package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 100
    var v2 bool = false
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 uint16 = 1
    v2 = false
    v2 = (1 == -9223372036854775808)
    _ = len(v4)
    switch -1 {
    case 127:
        var v7 uint8 = 255
        _ = v7
    case 32768:
        v6 += (^(v6))
    }
    var v8 uint = 537
    v1 -= (v1 & (-(v1)))
    if true {
        v9 := []int{2147483647, 9223372036854775807}
        fmt.Println(v8, v8, v6)
        fmt.Println(v8)
        _ = v9
    }
    switch 32767 {
    case 127:
        fmt.Println(v1)
    case 9223372036854775807:
        var v10 int8 = (v1 + v1)
        _ = v10
    }
    var v11 int32 = -2147483648
    fmt.Println(v11, v1)
    switch 128 {
    case 42:
        v2 = (!(0 > 114))
    }
    v1 -= (v1 ^ 79)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
