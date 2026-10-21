package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int16 = 1000
    var v2 int64 = 0
    v3 := make(map[int]string)
    var v4 uint16 = 255
    fmt.Println(v2, v2, v4)
    v2 = 1
    v5 := []int{int(v2), 81, 127}
    if true {
        v6 := v3[int(v4)]
        fmt.Println(v4, v6)
        _ = v6
    }
    if true {
        v1 = v1
    } else {
        v2 = v2
    }
    v1 -= v1
    switch 42 {
    case 9223372036854775807:
        var v7 int16 = (v1 & 712)
        _ = v7
    case -597:
        switch 0 {
        case -747:
            v3[4294967295] = ("foo bar" + "hello")
        default:
            if (!false) {
                var v8 uint16 = v4
                _ = v8
            } else {
                v9 := GenS1{F0: 90}
                _ = v9
            }
        }
    case 194:
        var v10 uint8 = 127
        _ = v10
    }
    _ = len(v3)
    var v11 uint32 = 4294967295
    var v12 int8 = 81
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
