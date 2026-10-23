package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[string]string)
    var v2 int = -2147483648
    var v3 int8 = 100
    var v4 bool = (!(((105 <= 55) && (255 >= 256)) && false))
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int32 = -2147483648
        fmt.Println(v3, v2)
        fmt.Println(v6)
        _ = fi5
        _ = v6
    }
    v7 := []int{2147483647, 42}
    var v8 int8 = (v3 ^ (^(v3)))
    fmt.Println(v4, v3)
    if (127 < (v8 << ((65535) & 63))) {
        fmt.Println(v2, v4, v2)
    }
    var v9 int32 = 2147483647
    v10 := GenS1{F0: (0 <= 255), F1: "hello"}
    v11 := GenS1{F0: (((!(-1 < (v2 * 358))) && (0.1 <= ((0.1 - 1.5) + (3.14159 * 0.0)))) && (!v4)), F1: ""}
    switch "hello" {
    case "hello":
        fmt.Println(v3, v4)
    case "a":
        var v12 int8 = (v8 >> ((0) & 63))
        _ = v12
    case "foo bar":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
