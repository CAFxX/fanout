package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 0
    var v2 uint16 = 2
    var v3 int8 = 1
    v4 := make(map[string]bool)
    var v5 uint = 255
    var v6 string = "\n\t"
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5, v6, v5)
    default:
        v7 := GenS1{F0: 1}
        _ = v7
    }
    var v8 float32 = (1e10 + 3.14159)
    v9 := GenS1{F0: 929}
    if true {
        fmt.Println(v6)
    } else {
        v6 = ""
        fmt.Println(v3, v2)
    }
    var v10 bool = ((-442 <= 1) && true)
    v5 = 255
    v11 := GenS1{F0: -2147483648}
    if (553 < -613) {
        v12 := GenS1{F0: 2147483647}
        switch 4294967295 {
        case 127:
            v5 = 65535
        case 42:
            v13 := []int{-2147483648}
            _ = v13
        case 9223372036854775807:
            v3 = (v3 + v3)
        }
        v5 += v5
        _ = v12
    } else {
        var v14 int = v1
        _ = v14
    }
    fmt.Println(v8, v1, v6)
    v2 += (v2 ^ v2)
    var v15 uint64 = 42
    var v16 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
