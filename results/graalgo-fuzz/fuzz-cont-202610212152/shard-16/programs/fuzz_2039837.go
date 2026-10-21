package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint = 255
    v3 := make(map[uint32]string)
    var v4 uint32 = 0
    v5 := []int{128, 32768, -1}
    var v6 int16 = -455
    v1 = "foo bar"
    v7 := GenS1{F0: 1, F1: 9223372036854775807, F2: (v6 + (v6 << ((v2) & 63)))}
    var v8 string = (v1 + ("a" + v1))
    for fi9 := 0; fi9 < 19; fi9++ {
        if (int64(fi9) >= -840) {
            v6 = v6
            v10 := v3[2147483647]
            _ = v10
        }
        _ = fi9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (0.1 != ((0.1 * -2.25) - (((3.14159 - (0.0 / 0.1)) - -2.25) / 3.14159))) {
        v2 -= v2
    } else {
        fmt.Println(v4, v4, v8)
    }
    fmt.Println(v1, v6)
    var v11 uint16 = 65535
    var v12 string = (v1 + "a")
    _ = len(v5)
    v4 -= v4
    var v13 uint16 = v11
    fmt.Println(v4, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
