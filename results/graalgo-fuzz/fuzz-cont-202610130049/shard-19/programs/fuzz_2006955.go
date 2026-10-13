package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }

func main() {
    var v1 int = 1
    v2 := []int{-2147483648}
    var v3 uint16 = 174
    var v4 int64 = 1
    v4 = v4
    var v5 int = v1
    v5 = 0
    v1 -= 32767
    v1 += (v5 >> ((4294967295) & 63))
    v6 := GenS1{F0: -788, F1: 255, F2: -1}
    var v7 uint32 = 65535
    v5 = v5
    var v8 string = ""
    fmt.Println(v5, v7)
    v7 = (v7 >> ((65535) & 63))
    for fi9 := 0; fi9 < 2; fi9++ {
        if (false && (3.14159 > (1.5 * 0.1))) {
            var v10 float32 = ((-2.25 * 1.5) / 3.14159)
            _ = v10
        }
        v8 = "a"
        _ = fi9
    }
    v4 -= 1
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
