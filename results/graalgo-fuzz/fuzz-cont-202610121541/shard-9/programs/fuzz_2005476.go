package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }

func main() {
    v1 := []string{"", "hello"}
    var v2 int8 = 0
    v3 := []uint32{2, 820, 0, 42}
    if false {
        v4 := GenS1{F0: 4294967295, F1: true, F2: 0}
        _ = v4
    } else {
        var v5 uint32 = 806
        var v6 uint16 = 255
        _ = v5
        _ = v6
    }
    v2 -= (v2 << ((283) & 63))
    if false {
        v2 = v2
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int16 = -619
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 65535, F1: (2 != 828), F2: 42}
    v2 = v2
    v10 := GenS1{F0: 0, F1: false, F2: 32767}
    if (-834 >= 9223372036854775807) {
        fmt.Println(v2)
        v3 = append(v3, 4294967295)
        _ = len(v1)
    }
    _ = len(v1)
    fmt.Println(v2)
    fmt.Println(v2)
    var v11 uint = 4294967295
    fmt.Println(v11, v11)
    fmt.Println(v2, v11)
    var v12 uint = (^(v11))
    _ = len(v3)
    var v13 uint16 = 255
    switch 9223372036854775807 {
    case 1:
        var v14 uint32 = 791
        _ = v14
    case -9223372036854775808:
        fmt.Println(v2, v13, v11)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
